package main

import (
	"encoding/json"
	"net/http"
)

var (
	def_username = "xyz1234"
	def_password = "1234"
)

type JwtResponse struct {
	Token string `json:"token"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	// Invoke parset form before get the form value
	r.ParseForm()

	u := r.FormValue("username")
	p := r.FormValue("password")

	if u != def_password || p != def_password {
		renderJson(
			w,
			http.StatusUnauthorized,
			ErrResponse{Code: "AUT001", Msg: "Authentication failed."},
		)
		return
	}

	/*
		//
		jwt, err := GenerateJWT("")
		if err != nil {
			renderJson(
				w,
				http.StatusUnauthorized,
				ErrResponse{Code: "AUT002", Msg: "Authentication failed.", ErrDetail: err.Error()},
			)
			return
		}

		renderJson(w, http.StatusOK, JwtResponse{Token: jwt})
	*/
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}
