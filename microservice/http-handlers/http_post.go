package main

import (
	"encoding/json"
	"net/http"
)

type Payload struct {
	Greeting string `json:"greeting,omitempty"`
	Name     string `json:"name,omitempty"`
}

type GreetingResponse struct {
	Successfully bool    `json:"successful"`
	Error        string  `json:"error,omitempty"`
	Payload      Payload `json:payload`
}

func GreetingHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var resp GreetingResponse

	if err := r.ParseForm(); err != nil {
		resp.Error = "Bad request!"
		if p, err := json.Marshal(resp); err == nil {
			w.Write(p)
		}
	}

	name := r.FormValue("name")
	greeting := r.FormValue("greeting")

	w.WriteHeader(http.StatusOK)

	resp.Successfully = true
	resp.Payload.Name = name
	resp.Payload.Greeting = greeting

	if p, err := json.Marshal(resp); err == nil {
		w.Write(p)
	}

}
