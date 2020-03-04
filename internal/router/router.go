package router

import (
	"github.com/gorilla/mux"
	"micro_auth/internal/auth"
	"micro_auth/internal/healthcheck"
	"micro_auth/internal/router/typings"
)

var V1Routes typings.Routes

func configureRoutes() {
	V1Routes = append(V1Routes, healthcheck.V1Routes...)
	V1Routes = append(V1Routes, auth.V1Routes...)
}

func BuildRouter(router *mux.Router) *mux.Router {
	configureRoutes()
	router = router.StrictSlash(true)
	apiRouter := router.PathPrefix("/api").Subrouter()

	for _, route := range V1Routes {
		apiRouter.
			HandleFunc(
				route.Pattern,
				corsMiddleware(
					authCheckMiddleware(
						route.Handler,
						route.AuthOnly,
					),
					route.Methods,
				),
			).
			Methods(route.Methods...).
			Name(route.Name)
	}

	return router
}
