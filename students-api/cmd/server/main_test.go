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

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
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

func Test_GetStudentByIDAnotherWay(t *testing.T) {
	router.HandleFunc("/api/students/{id}", h.GetStudentByID)
	ts := httptest.NewServer(router)
	defer ts.Close()

	t.Run("not found", func(t *testing.T) {
		res, err := http.Get(ts.URL + "/api/students/100")
		if err != nil {
			t.Errorf("expected nil, received %s", err.Error())
		}

		if res.StatusCode != http.StatusNotFound {
			t.Errorf("expected %d, received %d", http.StatusNotFound, res.StatusCode)
		}
	})
}
