package auth

import (
	authRest "micro_auth/internal/auth/api"
	"micro_auth/internal/router/typings"
	"net/http"
)

var V1Routes = typings.Routes{
	typings.Route{
		Name:     "TokenObtainHandler",
		Methods:  []string{http.MethodPost, http.MethodOptions},
		Pattern:  "/v1/auth/token",
		Handler:  authRest.TokenObtainHandler,
		AuthOnly: false,
	},
	typings.Route{
		Name:     "UserRegistrationHandler",
		Methods:  []string{http.MethodPost, http.MethodOptions},
		Pattern:  "/v1/auth/register",
		Handler:  authRest.UserRegistrationHandler,
		AuthOnly: false,
	},
}
