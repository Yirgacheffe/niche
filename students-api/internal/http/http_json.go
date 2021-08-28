package handler

import (
	"encoding/json"
	"net/http"
	"reflect"

	log "github.com/sirupsen/logrus"
)

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
	// ------------------------------------------------------------
}
