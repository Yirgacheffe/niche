package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

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

	h.Router.HandleFunc("/api/students/{school}", h.GetStudentBySchool).Methods("GET")
	h.Router.HandleFunc("/api/students", h.GetAllStudents).Methods("GET")
	h.Router.HandleFunc("/api/students", h.PostStudent).Methods("POST")
	h.Router.HandleFunc("/api/students/{id}", h.GetStudentByID).Methods("GET")
	h.Router.HandleFunc("/api/students/{id}", h.DeleteStudent).Methods("DELETE")
	h.Router.HandleFunc("/api/students/{id}", h.UpdateStudent).Methods("PUT")

	// h.Router.HandleFunc("/api/tag/{tag}/", h.GetStudentsByTag).Methods("GET")
	// h.Router.HandleFunc("/api/due/{year:[0-9]+}/{month:[0-9]+}/{day:[0-9]+}", h.GetStudentsByDate).Methods("GET")

	h.Router.HandleFunc("/api/status", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Status Up!")
	})
}

func (h *Handler) GetStudentByID(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	id := mux.Vars(r)["id"]
	studentID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		fmt.Fprintf(
			w,
			"Error Parsing ID to UINT.",
		)
	}

	student, err := h.Service.GetStudentByID(uint(studentID))
	if err != nil {
		fmt.Fprintf(
			w,
			"Error Retrieving Student by ID.",
		)
	}

	if err := json.NewEncoder(w).Encode(student); err != nil {
		panic(err)
	}

}

func (h *Handler) DeleteStudent(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	id := mux.Vars(r)["id"]
	studentID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		fmt.Fprintf(w, "Error Parsing ID to UINT.")
	}

	err = h.Service.DeleteStudent(uint(studentID))
	if err != nil {
		fmt.Fprintf(w, "Failed to delete student by ID.")
	}

	message := `{Message: "Student deleted succeed!"}`
	if err := json.NewEncoder(w).Encode(message); err != nil {
		panic(err)
	}

}

func (h *Handler) PostStudent(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	var s student.Student
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		fmt.Fprintf(w, "Json body decode failed.")
	}

	s, err := h.Service.PostStudent(s)
	if err != nil {
		fmt.Fprintf(w, "Failed to create new student.")
	}

	if err = json.NewEncoder(w).Encode(s); err != nil {
		panic(err)
	}

}

func (h *Handler) GetStudentBySchool(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	school := vars["school"]

	students, err := h.Service.GetStudentBySchool(school)
	if err != nil {
		fmt.Fprintf(
			w,
			"Error Retrieving Students by School.",
		)
	}

	if err := json.NewEncoder(w).Encode(students); err != nil {
		panic(err)
	}

}

func (h *Handler) UpdateStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	id := mux.Vars(r)["id"]
	studentID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		fmt.Fprint(w, "Error Parsing ID to UINT.")
	}

	var s student.Student
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		fmt.Fprintf(
			w,
			"Decode Json body failed.",
		)
	}

	s, err = h.Service.UpdateStudent(uint(studentID), s)
	if err != nil {
		fmt.Fprintf(w, "Failed to update student.")
	}

	if err = json.NewEncoder(w).Encode(s); err != nil {
		panic(err)
	}
}

func (h *Handler) GetAllStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	students, err := h.Service.GetAllStudents()
	if err != nil {
		fmt.Fprintf(w, "Failed to retrieve students.")
	}

	if err := json.NewEncoder(w).Encode(students); err != nil {
		panic(err)
	}
}

func renderJSON(w http.ResponseWriter, v interface{}) {
	js, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "appliction/json")
	w.Write(js)
	// -------------------------------------------------------------------------------
}
