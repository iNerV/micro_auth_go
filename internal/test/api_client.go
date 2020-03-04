package test

import (
	"bytes"
	"encoding/json"
	"micro_auth/internal/application"
	"net/http"
	"net/http/httptest"
	"testing"
)

type API struct {
	t             *testing.T
	app           *application.App
	godMode       bool
	isAnon        bool
	Authorization string
}

// T retrieves the current *testing.T context.
func (api *API) T() *testing.T {
	return api.t
}

// App retrieves the current *application.App context.
func (api *API) App() *application.App {
	return api.app
}

func (api API) Get(path string) *httptest.ResponseRecorder {
	resp, err := api.call("GET", path, nil)
	Ok(api.T(), err)
	Equals(api.T(), resp.Code, http.StatusOK)
	return resp
}

func (api *API) Post(path string, body []byte) *httptest.ResponseRecorder {
	resp, err := api.call("POST", path, body)
	Ok(api.T(), err)
	Equals(api.T(), resp.Code, http.StatusCreated)
	return resp
}

func (api *API) call(method, path string, body []byte) (*httptest.ResponseRecorder, error) {
	req, err := http.NewRequest(method, path, api.encode(body))
	rr := httptest.NewRecorder()

	api.App().Router.ServeHTTP(rr, req)

	return rr, err
}

func (api *API) decode(body *httptest.ResponseRecorder) map[string]string {
	response := make(map[string]string, 1)
	_ = json.Unmarshal([]byte(body.Body.String()), &response)
	return response
}

func (api *API) encode(request []byte) *bytes.Reader {
	return bytes.NewReader(request)
}
