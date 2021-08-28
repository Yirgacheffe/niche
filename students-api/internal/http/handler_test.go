package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func Test_GetStudentByID(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/students/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/api/students/{id}", GetStudentByID)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf(
			"handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{"alive": true}`
	if rr.Body.String() != expected {
		t.Errorf(
			"handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func Test_GetAllStudents(t *testing.T) {

	r := mux.NewRouter()
	r.Handle("/api/students/{id}", GetAllStudents)
	ts := httptest.NewServer(r)
	defer ts.Close()

	t.Run("not found", func(t *testing.T){
		res, err := http.Get(ts.URL + "/api/students/1")
		if  err != nil {
			t.Errorf("expected nil, received %s", err.Error())
		}
		if res.StatusCode != http.StatusNotFound {
			t.Errorf("expected %d, received %d", http.StatusNotFound, res.StatusCode)
		}
	})

	t.Run("found", func(t *testing.T){
		res, err := http.Get(ts.URL + "/api/students/2")
		if err != nil {
			t.Errorf("expected nil, received %s", err.Error())
		}
		if res.StatusCode != http.StatusOK {
			t.Errorf("expected %d, received %d", http.StatusOK, res.StatusCode)
		}
	})

}
