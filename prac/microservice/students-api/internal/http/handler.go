package handler

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	Router *mux.Router
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) InitRoutes() {
	h.Router = mux.NewRouter()
	h.Router.HandleFunc("/api/status", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Status Up!")
	})
}
