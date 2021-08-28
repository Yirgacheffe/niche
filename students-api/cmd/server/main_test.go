package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"students-api/internal/db"
	internalHttp "students-api/internal/http"
	"students-api/internal/service/student"

	"github.com/gorilla/mux"
)

const dbFile = "/Users/aaron/proj/golang-exec/niche/students-api/students.db"

var students = []student.Student{
	{FirstName: "K", LastName: "Dash", Age: 20, School: "No. 1 middle school"},
	{FirstName: "M", LastName: "Dash", Age: 22, School: "No. 2 middle school"},
}

var (
	router *mux.Router
	rr     *httptest.ResponseRecorder
	s      *student.Service
	h      *internalHttp.Handler
)

func setUpTableEnv() {
	if err := os.Setenv("DB_TABLE", dbFile); err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	log.Println("----------------------- Clear table -------------------------")
}

func setUp() {
	conn, err := db.InitDatabase()
	if err != nil {
		log.Fatal(err)
	}
	s = student.NewService(conn)
	h = internalHttp.NewHandler(s)

	router = mux.NewRouter()
	rr = httptest.NewRecorder()
}

func TestMain(m *testing.M) {
	setUpTableEnv()
	setUp()

	code := m.Run()
	clearTable()
	os.Exit(code)
}

func Test_GetStudentByID(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/students/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	router.HandleFunc("/api/students/{id}", h.GetStudentByID)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf(
			"handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	type rawData struct {
		Status string
		Data   student.Student
	}

	var s rawData
	json.Unmarshal(rr.Body.Bytes(), &s)

	actual := s.Data.ID
	expect := 1

	if actual != uint(expect) {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expect)
	}
}
