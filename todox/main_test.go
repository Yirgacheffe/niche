package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_ListNoteHandler(t *testing.T) {

	req, err := http.NewRequest("GET", "/api/notes", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ListNoteHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned a wrong status code: got %v expected %v",
			status,
			http.StatusOK)
	}

	expected := "{}"
	got := rr.Body.String()

	if got != expected {
		t.Errorf("Handler returned a unexpected result: got %v expected %v",
			got,
			expected)
	}

}
