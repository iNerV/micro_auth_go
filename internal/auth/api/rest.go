package authRest

import (
	"encoding/json"
	"io"
	"micro_auth/internal/auth/services"
	"net/http"
)

type NewUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func UserRegistrationHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	var req NewUser
	_ = json.NewDecoder(r.Body).Decode(&req)
	services.CreateUser(req.Username, req.Email, req.Password)

	io.WriteString(w, `{"access": "dfdddddd", "refresh": "sdfsfsfd"}`)
}

func TokenObtainHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	io.WriteString(w, `{"access": "dfdddddd", "refresh": "sdfsfsfd"}`)
}
