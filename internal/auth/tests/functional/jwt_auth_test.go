package authFunctionalTest

import (
	"encoding/json"
	"fmt"
	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/suite"
	"log"
	"micro_auth/internal/auth/models"
	"micro_auth/internal/test"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type APITestSuite struct {
	test.ApiClientSuite
}

func TestAPITestSuite(t *testing.T) {
	suite.Run(t, new(APITestSuite))
}

func getUser() models.User {
	user := models.User{}
	err := faker.FakeData(&user)
	if err != nil {
		log.Fatalln(err)
	}
	err = user.ID.Set(faker.UUIDHyphenated())
	if err != nil {
		log.Fatalln(err)
	}
	return user
}

func decode(body *httptest.ResponseRecorder) map[string]string {
	response := make(map[string]string, 1)
	_ = json.Unmarshal([]byte(body.Body.String()), &response)
	return response
}

func (s *APITestSuite) getToken(email string) *httptest.ResponseRecorder {
	body := fmt.Sprintf(`'{"email": %s, "password": "password"}'`, email)
	reqBody, err := json.Marshal(body)
	if err != nil {
		log.Fatalln(err)
	}
	got := s.Client.Post("/api/v1/auth/token", reqBody)
	return got
}

//func randomHexString(length int) (string, error) {
//	_bytes := make([]byte, length)
//	if _, err := rand.Read(_bytes); err != nil {
//		return "", err
//	}
//	return hex.EncodeToString(_bytes), nil
//}

func (s *APITestSuite) TestGettingTokenOk() {
	got := s.getToken(getUser().Email)

	test.Equals(s.T(), got.Code, http.StatusCreated)
}

func (s *APITestSuite) TestGettingTokenIsToken() {
	got := decode(s.getToken(getUser().Email))

	test.Assert(s.T(), len(got["access"]) > 60, "Access token should be bigger than 60")
	test.Assert(s.T(), len(got["refresh"]) > 60, "Refresh token should be bigger than 60")
	test.Equals(s.T(), len(strings.Split(got["access"], ".")), 3)
	test.Equals(s.T(), len(strings.Split(got["refresh"], ".")), 3)
}

//
//func TestReceivedAccessTokenWorks(t *testing.T) {
//	tr := tokenResponse{}
//	got := json.Unmarshal(getToken(t, getUser().Email).Body.Bytes(), &tr)
//
//	token := got.access
//	c := test.API{Authorization: "Bearer" + token}
//
//	r := c.Get(t, app, "/v1/whoami")
//
//	test.Equals(t, r.Code, http.StatusOK)
//
//	decoded := decode(r)
//	test.Equals(t, decoded.id, getUser().id)
//}
//
//func TestInvalidAccessTokenDoesNotWork(t *testing.T) {
//	token, _ := randomHexString(32)
//	c := test.API{Authorization: "Bearer" + token}
//
//	got := c.Get(t, app, "/v1/whoami")
//	test.Equals(t, got.Code, http.StatusUnauthorized)
//}
//
//func TestRefreshAccessToken(t *testing.T)  {
//	got1 := decode(getToken(t, getUser().Email))
//
//	c := test.API{}
//
//	body := []byte(fmt.Sprintf(`'{"refresh": %s}'`, got1.refresh))
//	got2 := c.Post(t, app, "/v1/auth/token/refresh", bytes.NewBuffer(body))
//	test.Equals(t, got2.Code, http.StatusOK)
//
//	decodedResponse := decode(got2)
//
//	test.Assert(t, len(decodedResponse.access) > 32, "")
//	test.Assert(t, len(decodedResponse.refresh) > 32, "")
//	test.Assert(t, decodedResponse.access != got1.access, "")
//	test.Assert(t, decodedResponse.refresh != got1.refresh, "")
//}
