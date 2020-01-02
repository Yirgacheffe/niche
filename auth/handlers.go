package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func isUserMatch( username, password string) boolean {
	return username == "xyz1234" && password == "1qwe4edc"
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		DisplayAppError(
			w, 
			err, 
			"Invalid login data.", 
			500,
		)
		return
	}

	if !isUserMatch(user.username, user.password) {
		DisplayAppError(
			w, 
			err, 
			"Invalid login credentials.", 
			401,
		)
		return
	}
	
	token, err := GenerateJWT(user.UserName, "member")
	if err != nil {
		DisplayAppError(
			w,
			err,
			"Error while generating access token.",
			500,
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	j, err := json.Marshal(token)
	if err != nil {
		DisplayAppError(
			w,
			err,
			"An unexpected error has occured.",
			500,
		)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(j)

}
