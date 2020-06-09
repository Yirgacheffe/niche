package main

import (
	"encoding/json"
	"io"
	"net/http"
)

// LoginResponse - Also the authentication response
type LoginResponse struct {
	Token string `json:"token"`
}

// LoginHandler - Do user login
func LoginHandler(w http.ResponseWriter, r *http.Request) {

	// Invoke parset form before get the form value
	r.ParseForm()

	userName := r.FormValue("username")
	password := r.FormValue("password")

	if userName != "xyz1234" || password != "1qaz2wsx" {
		DisplayAppError(w, 401, "Invalid login credentials.")
		return
	}

	jwt, err := GenerateJWT(userName, "member")
	if err != nil {
		DisplayAppError(w, 500, "Error while generating JWT.")
		return
	}

	j, err := json.Marshal(LoginResponse{Token: jwt})
	if err != nil {
		DisplayAppError(w, 500, "An unexpected error occured.")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)

}

// JwksHandler - Provide JWK url for token verification
func JwksHandler(w http.ResponseWriter, r *http.Request) {

}

// HealthCheckHandler - monitor purpose
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	io.WriteString(w, `{"alive": true}`)
}
