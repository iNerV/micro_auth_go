package router_test

import (
	"github.com/stretchr/testify/suite"
	"micro_auth/internal/router"
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

func (s *APITestSuite) TestCorsForAllEndpoints() {
	for _, tt := range router.V1Routes {
		s.Run("Test CORS allowed for "+tt.Name, func() {
			s.T().Parallel()

			req, _ := http.NewRequest("OPTIONS", "/api"+tt.Pattern, nil)
			req.Header.Set("Origin", "https://origin.test")
			req.Header.Set("Access-Control-Request-Methods", tt.Methods[0])

			rr := httptest.NewRecorder()
			s.App().Router.ServeHTTP(rr, req)

			headerMethods := strings.Join(tt.Methods[:], ",")
			test.Equals(s.T(), rr.Code, http.StatusOK)
			test.Equals(s.T(), rr.Header().Get("Access-Control-Allow-Methods"), headerMethods)
		})
	}
}
