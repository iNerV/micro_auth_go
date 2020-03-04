package api_test

import (
	"github.com/stretchr/testify/suite"
	"micro_auth/internal/test"
	"testing"
)

type APITestSuite struct {
	test.ApiClientSuite
}

func TestAPITestSuite(t *testing.T) {
	suite.Run(t, new(APITestSuite))
}

func (s *APITestSuite) TestHealthCheckHandler() {
	got := s.Client.Get("/api/v1/healthcheck")

	expected := `{"alive": true}`

	test.Equals(s.T(), got.Body.String(), expected)
}
