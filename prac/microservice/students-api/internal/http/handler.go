package handler

import (
	"encoding/json"
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

	h.Router.HandleFunc("/api/students", h.GetAllStudents).Methods("GET")
	h.Router.HandleFunc("/api/students", h.PostStudent).Methods("POST")
	h.Router.HandleFunc("/api/students/{id:[0-9]+}", h.GetStudentByID).Methods("GET")
	h.Router.HandleFunc("/api/students/{id:[0-9]+}", h.DeleteStudent).Methods("DELETE")
	h.Router.HandleFunc("/api/tag/{tag}/", h.GetStudentsByTag).Methods("GET")
	h.Router.HandleFunc("/api/due/{year:[0-9]+}/{month:[0-9]+}/{day:[0-9]+}", h.GetStudentsByDate).Methods("GET")

	h.Router.HandleFunc("/api/status", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Status Up!")
	})

}

func (h *Handler) GetAllStudents(w http.ResponseWriter, r *http.Request)     {}
func (h *Handler) PostStudent(w http.ResponseWriter, r *http.Request)        {}
func (h *Handler) GetStudentBySchool(w http.ResponseWriter, r *http.Request) {}
func (h *Handler) GetStudentByID(w http.ResponseWriter, r *http.Request)     {}
func (h *Handler) UpdateStudent(w http.ResponseWriter, r *http.Request)      {}
func (h *Handler) DeleteStudent(w http.ResponseWriter, r *http.Request)      {}
func (h *Handler) GetStudentsByTag(w http.ResponseWriter, r *http.Request)   {}
func (h *Handler) GetStudentsByDate(w http.ResponseWriter, r *http.Request)  {}

func renderJSON(w http.ResponseWriter, v interface{}) {

	js, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "appliction/json")
	w.Write(js)

}
