package main

import (
	"encoding/json"
	"net/http"
)

func (c *Controller) GetValue(UseDefault bool) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		value := "default"
		if !UseDefault {
			value = c.Storage.Get()
		}

		w.WriteHeader(http.StatusOK)
		p := Payload{Value: value}
		if payload, err := json.Marshal(p); err == nil {
			w.Write(payload)
		}

	}

}
