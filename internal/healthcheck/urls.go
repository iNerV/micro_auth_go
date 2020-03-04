package healthcheck

import (
	healthcheckApi "micro_auth/internal/healthcheck/api"
	"micro_auth/internal/router/typings"
	"net/http"
)

var V1Routes = typings.Routes{
	typings.Route{
		Name:     "HealthCheckHandler",
		Methods:  []string{http.MethodGet, http.MethodOptions},
		Pattern:  "/v1/healthcheck",
		Handler:  healthcheckApi.HealthCheckHandler,
		AuthOnly: false,
	},
}
