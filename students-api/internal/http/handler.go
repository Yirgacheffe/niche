package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strconv"

	log "github.com/sirupsen/logrus"

	"students-api/internal/service/student"

	"github.com/gorilla/mux"
)

type Info struct {
	HTTPStatus int    `json:"-"`
	MainCode   string `json:"code"`
	Message    string `json:"msg"`
}

type Response struct {
	Status      string      `json:"status,omitempty"`
	Code        string      `json:"code,omitempty"`
	Msg         string      `json:"msg,omitempty"`
	ErrorDetail string      `json:"error_detail,omitempty"`
	Data        interface{} `json:"data,omitempty"`
}

var (
	ErrDefault        = errors.New("API Error")
	ErrIdParamParsing = errors.New("error parsing parameter ID")
	ErrFetchStudent   = errors.New("error retrieve student")
)

var respErrFormatter = map[error]Info{
	ErrDefault:        {500, "STD000", "API Error"},
	ErrIdParamParsing: {400, "STD100", "Error parsing parameter ID."},
	ErrFetchStudent:   {404, "STD110", "Error retrieve student."},
}

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

	// h.Router.HandleFunc("/api/students/{school}", h.GetStudentBySchool).Methods("GET")
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

const (
	SUCCESS = "success"
	FAIL    = "fail"
)

func NewErrorResponse(code string, msg string, err error) Response {
	return Response{
		Status: FAIL,
		Code:   code,
		Msg:    msg, ErrorDetail: err.Error()}
}

func formatErrResp(err error) (response Response) {
	response = Response{Status: FAIL}
	defer func() {
		if err := recover(); err != nil {
			response.ErrorDetail = ""
		}
	}()

	if info, ok := respErrFormatter[err]; ok {
		response.Code = info.MainCode
		response.Msg = info.Message
		response.ErrorDetail = err.Error()
	}

	return response
}

func NewSuccessResponse(data interface{}) Response {
	return Response{Status: SUCCESS, Data: data}
}

func (h *Handler) GetStudentByID(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]
	studentID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		errResp := NewErrorResponse(
			"STD775",
			ErrIdParamParsing.Error(),
			err,
		)
		renderJSON(w, http.StatusBadRequest, errResp)
		return
	}

	s, err := h.Service.GetStudentByID(uint(studentID))
	if err != nil {
		errResp := NewErrorResponse(
			"STD790",
			ErrFetchStudent.Error(),
			err,
		)
		renderJSON(w, http.StatusNotFound, errResp)
		return
	}

	renderJSON(w, http.StatusOK, NewSuccessResponse(s))

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

func array(v interface{}) interface{} {
	if rv := reflect.ValueOf(v); rv.Kind() == reflect.Slice && rv.IsNil() {
		v = []struct{}{}
	}
	return v // render "[]" rather than "nil" if value is slice
}

func renderJSON(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	err := json.NewEncoder(w).Encode(array(v))
	if err != nil {
		log.WithFields(
			log.Fields{
				"module": "Student",
				"error":  err,
			}).Error("Error happened while writing Content-Type header using status")
	}
	// -------------------------------------------------------------------------------
}