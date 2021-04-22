package handler

import (
	"fmt"
	"net/http"
	"students-api/internal/service/student"

	"github.com/gorilla/mux"
)

// Handler - students api http handler
type Handler struct {
	Router  *mux.Router
	Service *student.Service
}

func NewHandler(service *student.Service) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) InitRoutes() {
	h.Router = mux.NewRouter()

	h.Router.HandleFunc("/api/students", h.Service.GetAllStudents).Methods("GET")
	h.Router.HandleFunc("/api/students", h.Service.PostStudent).Methods("POST")

	h.Router.HandleFunc("/api/status", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Status Up!")
	})
}
