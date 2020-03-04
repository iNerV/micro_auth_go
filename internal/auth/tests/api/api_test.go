package authApiTest

import (
	"encoding/json"
	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/suite"
	"log"
	"micro_auth/internal/auth/models"
	"micro_auth/internal/test"
	"testing"
)

type APITestSuite struct {
	test.ApiClientSuite
	user    *user
	reqBody []byte
}

func TestAPITestSuite(t *testing.T) {
	suite.Run(t, new(APITestSuite))
}

func (s *APITestSuite) SetupTest() {
	var err error
	s.ApiClientSuite.SetupTest()
	s.user = createUser()
	s.reqBody, err = json.Marshal(s.user)
	test.Ok(s.T(), err)
}

type user struct {
	Email    string `json:"email" faker:"email"`
	Username string `json:"username" faker:"username"`
	Password string `json:"password" faker:"password"`
}

func createUser() *user {
	user := &user{}
	err := faker.FakeData(&user)
	if err != nil {
		log.Fatal(err)
	}
	return user
}

func (s *APITestSuite) Test_registration_new_user() {

	_ = s.Client.Post("/api/v1/auth/register", s.reqBody)

	var response models.User
	err := s.DBConn.Get(&response, "SELECT username, email from users where username=$1", s.user.Username)
	test.Ok(s.T(), err)
	test.Equals(s.T(), response.Username, s.user.Username)
	test.Equals(s.T(), response.Email, s.user.Email)
}

func (s *APITestSuite) Test_registration_new_user_with_invalid_email() {
	s.user.Email = "invalid"

	_ = s.Client.Post("/api/v1/auth/register", s.reqBody)
	test.Assert(s.T(), false, "Not Impl error")
}

func (s *APITestSuite) Test_after_registration_new_user_we_should_have_only_one_user_in_db() {
	_ = s.Client.Post("/api/v1/auth/register", s.reqBody)

	var count int
	_ = s.DBConn.QueryRowx("SELECT count(*) as count from users").Scan(&count)
	test.Equals(s.T(), count, 1)
}
