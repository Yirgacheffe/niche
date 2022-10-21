package handler

import (
	"log"
	"net/http"
	"time"
)

// testing purpose, apply on single endpoint
func WithMetrics(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		began := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s took %s", r.Method, r.URL, time.Since(began))
	})
}
