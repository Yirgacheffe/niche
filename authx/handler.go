package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var (
	def_username = "xyz1234"
	def_password = "1234"
)

type JwtResponse struct {
	Token string `json:"token"`
}

type User struct {
	ID       int
	UserName string
	Email    string
}

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	// Invoke parset form before get the form value
	r.ParseForm()

	u := r.FormValue("username")
	p := r.FormValue("password")

	user, err := getUserFromDB(u, p)
	if err != nil {
		renderJson(w, 403, ErrResponse{"AUT001", "Authentication failed.", err.Error()})
		return
	}

	tokenString, err := GenerateJWT(user.ID, user.UserName, user.Email)
	if err != nil {
		renderJson(w, 500, ErrResponse{"AUT003", "Authentication failed.", err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/jwt")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, tokenString)
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}

func getUserFromDB(user, pass string) (*User, error) {
	if user != def_username || pass != def_password {
		return nil, fmt.Errorf("Incorrect parameter.")
	}
	return &User{ID: 1, UserName: user, Email: "test@abc.com"}, nil
}
