package typings

import "net/http"

type HttpHandlerFunc func(http.ResponseWriter, *http.Request)

type Route struct {
	Name     string
	Methods  []string
	Pattern  string
	Handler  HttpHandlerFunc
	AuthOnly bool
}

type Routes []Route
