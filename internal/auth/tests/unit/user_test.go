package authUnitTest

import (
	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/suite"
	"log"
	"micro_auth/internal/auth/models"
	"micro_auth/internal/auth/services"
	"micro_auth/internal/test"
	"strings"
	"testing"
	"time"
)

type UnitTestSuite struct {
	test.PostgresSuite
}

func TestUnitTestSuite(t *testing.T) {
	suite.Run(t, new(UnitTestSuite))
}

func getFakeUser() *models.User {
	user := &models.User{}
	err := faker.FakeData(&user)
	if err != nil {
		log.Fatal(err)
	}
	err = user.ID.Set(faker.UUIDHyphenated())
	if err != nil {
		log.Fatal(err)
	}
	err = user.CreatedAt.Set(time.Now().UTC())
	if err != nil {
		log.Fatal(err)
	}
	return user
}

func (s *UnitTestSuite) Test_create_user() {
	user := getFakeUser()

	got, err := services.CreateUser(user.Username, user.Email, user.Password)
	//passwordEqual := bcrypt.CompareHashAndPassword([]byte(got.Password), []byte(user.Password))

	test.Ok(s.T(), err)
	test.Equals(s.T(), got.Username, user.Username)
	test.Equals(s.T(), got.Email, user.Email)
	//test.Ok(s.T(), passwordEqual)
}

func (s *UnitTestSuite) Test_create_user_email_domain_normalize_rfc3696() {
	cases := []struct {
		given    string
		expected string
	}{
		{`Abc\@DEF@EXAMPLE.com`, `Abc\@DEF@example.com`},
		{`"Abc\@DEF"@EXAMPLE.com`, `"Abc\@DEF"@example.com`},
		{`"Fred Bloggs"@EXAMPLE.com`, `"Fred Bloggs"@example.com`},
		{`"Joe\\Blow"@EXAMPLE.com`, `"Joe\\Blow"@example.com`},
		{`"Abc@def"@EXAMPLE.com`, `"Abc@def"@example.com`},
		{`customer/department=shipping@EXAMPLE.com`, `customer/department=shipping@example.com`},
		{`\$A12345@EXAMPLE.com`, `\$A12345@example.com`},
		{`\!def!xyz%abc@EXAMPLE.com`, `\!def!xyz%abc@example.com`},
		{`_somename@EXAMPLE.com`, `_somename@example.com`},
		{`somename@example.com`, `somename@example.com`},
	}
	for _, c := range cases {
		c := c
		s.Run(c.given, func() {
			s.T().Parallel()
			returned := services.NormalizeEmail(c.given)

			test.Equals(s.T(), returned, c.expected)
		})
	}
}

func (s *UnitTestSuite) Test_make_random_password() {
	allowedCharacters := "abcdefg"
	password := services.MakeRandomPassword(5, allowedCharacters)

	test.Equals(s.T(), len(password), 5)
	for _, char := range password {
		test.Assert(s.T(), strings.Contains(password, string(char)), "Password contains not allowed characters")
	}
}

func (s *UnitTestSuite) Test_get_user_by_username() {
	user := getFakeUser()

	_ = s.DBConn.QueryRowx(
		"INSERT INTO users (id, username, email, password, created_at) VALUES($1, $2, $3, $4, $5)",
		user.ID, user.Username, user.Email, user.Password, user.CreatedAt,
	)

	got := services.GetByUsername(user.Username)
	test.Equals(s.T(), got.Username, user.Username)
	test.Equals(s.T(), got.Email, user.Email)
}

func (s *UnitTestSuite) Test_get_user_by_username_with_invalid_username() {
	user := getFakeUser()

	_ = s.DBConn.QueryRowx(
		"INSERT INTO users (id, username, email, password, created_at) VALUES($1, $2, $3, $4, $5)",
		user.ID, user.Username, user.Email, user.Password, user.CreatedAt,
	)

	got := services.GetByUsername("invalid_username")
	test.Equals(s.T(), got.Username, user.Username)
	test.Equals(s.T(), got.Email, user.Email)
}

func (s *UnitTestSuite) Test_get_user_by_email() {
	user := getFakeUser()

	_ = s.DBConn.QueryRowx(
		"INSERT INTO users (id, username, email, password, created_at) VALUES($1, $2, $3, $4, $5)",
		user.ID, user.Username, user.Email, user.Password, user.CreatedAt,
	)

	got := services.GetByEmail(user.Email)
	test.Equals(s.T(), got.Username, user.Username)
	test.Equals(s.T(), got.Email, user.Email)
}

func (s *UnitTestSuite) Test_get_user_by_email_with_invalid_email() {
	user := getFakeUser()

	_ = s.DBConn.QueryRowx(
		"INSERT INTO users (id, username, email, password, created_at) VALUES($1, $2, $3, $4, $5)",
		user.ID, user.Username, user.Email, user.Password, user.CreatedAt,
	)

	got := services.GetByEmail("invalid_email")
	test.Equals(s.T(), got.Username, user.Username)
	test.Equals(s.T(), got.Email, user.Email)
}
