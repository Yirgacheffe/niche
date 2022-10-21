package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

type ErrResponse struct {
	Code      string `json:"code,omitempty"`
	Msg       string `json:"msg,omitempty"`
	ErrDetail string `json:"err_detail,omitempty"`
}

func renderJson(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		log.Println("Error happened while writing Content-Type header using status")
	}
}
