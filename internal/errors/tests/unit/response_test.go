package errors_test

import (
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v3"
	"micro_auth/internal/errors"
	"micro_auth/internal/test"
	"net/http"
	"testing"
)

func TestErrorResponse_Error(t *testing.T) {
	e := errors.ErrorResponse{
		Message: "abc",
	}
	test.Equals(t, "abc", e.Error())
}

func TestErrorResponse_StatusCode(t *testing.T) {
	e := errors.ErrorResponse{
		Status: 400,
	}
	test.Equals(t, 400, e.StatusCode())
}

func TestInternalServerError(t *testing.T) {
	res := errors.InternalServerError("test")
	test.Equals(t, http.StatusInternalServerError, res.StatusCode())
	test.Equals(t, "test", res.Error())
	res = errors.InternalServerError("")
	test.Assert(t, res.Error() != "", "Error is empty")
}

func TestNotFound(t *testing.T) {
	res := errors.NotFound("test")
	test.Equals(t, http.StatusNotFound, res.StatusCode())
	test.Equals(t, "test", res.Error())
	res = errors.NotFound("")
	test.Assert(t, res.Error() != "", "Error is empty")
}

func TestUnauthorized(t *testing.T) {
	res := errors.Unauthorized("test")
	test.Equals(t, http.StatusUnauthorized, res.StatusCode())
	test.Equals(t, "test", res.Error())
	res = errors.Unauthorized("")
	test.Assert(t, res.Error() != "", "Error is empty")
}

func TestForbidden(t *testing.T) {
	res := errors.Forbidden("test")
	test.Equals(t, http.StatusForbidden, res.StatusCode())
	test.Equals(t, "test", res.Error())
	res = errors.Forbidden("")
	test.Assert(t, res.Error() != "", "Error is empty")
}

func TestBadRequest(t *testing.T) {
	res := errors.BadRequest("test")
	test.Equals(t, http.StatusBadRequest, res.StatusCode())
	test.Equals(t, "test", res.Error())
	res = errors.BadRequest("")
	test.Assert(t, res.Error() != "", "Error is empty")
}

func TestInvalidInput(t *testing.T) {
	err := errors.InvalidInput(validation.Errors{
		"xyz": fmt.Errorf("2"),
		"abc": fmt.Errorf("1"),
	})
	test.Equals(t, http.StatusBadRequest, err.Status)
	test.Equals(t, []errors.InvalidField{{"abc", "1"}, {"xyz", "2"}}, err.Details)
}
