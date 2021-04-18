package handler

import "github.com/gorilla/mux"

type Handler struct {
	Router *mux.Router
}

func NewHandler() *Handler {
	return &Handler{}
}
