package router

import (
	"encoding/json"
	"micro_auth/core"
	"micro_auth/internal/errors"
	"micro_auth/internal/router/typings"
	"net/http"
	"strings"
)

func authCheckMiddleware(next typings.HttpHandlerFunc, authOnly bool) typings.HttpHandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		if !authOnly {
			next(res, req)
			return
		}
		tokenHeader := req.Header.Get("Authorization")
		if tokenHeader == "" {
			res.WriteHeader(http.StatusUnauthorized)
			_ = json.NewEncoder(res).Encode(errors.Unauthorized("Missing auth token"))
			res.Header().Add("Content-Type", "application/json")
			return
		}
		next(res, req)
	}
}

func corsMiddleware(next typings.HttpHandlerFunc, allowedMethods []string) typings.HttpHandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		if req.Method == "OPTIONS" {
			res.Header().Set("Access-Control-Allow-Methods", strings.Join(allowedMethods[:], ","))
			res.Header().Set("Access-Control-Allow-Origin", strings.Join(core.AppConfig.AllowedHosts[:], ","))
			res.Header().Set("Content-Type", "application/json")
			return
		}
		next(res, req)
	}
}
