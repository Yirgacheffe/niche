package main

import (
	"fmt"
	"net/http"

	db "niche-auth/db"
)

func NewAuthHandler(db *db.DB) *AuthHandler {
	return &AuthHandler{
		repo: NewAccountRepo(db.DB),
	}
}

type AuthHandler struct {
	repo AccountRepo
}

func (a *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	// Invoke parset form before get the form value
	r.ParseForm()
	u := r.FormValue("username")
	p := r.FormValue("password")

	user, err := a.repo.GetAccount(u, p)
	if err != nil {
		renderJson(w, 401, ErrResponse{"AUT001", "Authentication failed.", err.Error()})
		return
	}

	tokenString, err := GenerateJWT(user.ID, u, user.Email)
	if err != nil {
		renderJson(w, 500, ErrResponse{"AUT003", "Authentication failed.", err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/jwt")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, tokenString)
}

/*
func (a *AuthHandler) Auth(w http.ResponseWriter, r *http.Request) {

	token, err := jwt.ParseFromRequest(r, func(token *jwt.Token)(interface{}, error)) {
		return []byte, nil
	})

	if err != nil {
		switch err.(type) {
		default:
			log.Printf("token parse error: %v\n", err)
			return
		}
	}
}

func JwksHandler(w http.ResponseWriter, r *http.Request) {

	jwk := JSONWebKey{
		Kty: "RSA",
		Kid: "5UoLr4eKYMA3LeFszcWarhNxz",
		Use: "nichesoft.io",
		E:   "AQAB",
		N:   "",
		Alg: "RS256",
	}

	jwks := []JSONWebKey{}
	jwks = append(jwks, jwk)

	if j, err := json.Marshal(jwks); err != nil {
		panic(err)
	} else {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	}

}
*/
