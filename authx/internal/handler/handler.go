package handler

import (
	"fmt"
	"net/http"

	db "niche-auth/internal/db"
	"niche-auth/internal/jwt"
	"niche-auth/internal/model"
)

func NewAuthHandler(db *db.DB) *AuthHandler {
	return &AuthHandler{
		repo: model.NewAccountRepo(db.DB),
	}
}

type AuthHandler struct {
	repo model.AccountRepo
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

	tokenString, err := jwt.GenerateJWT(user.ID, u, user.Email)
	if err != nil {
		renderJson(w, 500, ErrResponse{"AUT003", "Authentication failed.", err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/jwt")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, tokenString)
}
