package main

import (
	"log"
	"net/http"
	"time"
)

// Middleware wrapped by http.HandlerFunc
type Middleware func(http.HandlerFunc) http.HandlerFunc

func ApplyMiddleware(handler http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	applied := handler
	for _, m := range middlewares {
		applied = m(applied)
	}

	return applied
}

func Logger(l *log.Logger) Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			l.Printf("started request to %s with id %s", r.URL, GetID(r.Context()))
			next(w, r)
			l.Printf("complete request to %s with id %s in %s", r.URL, GetID(r.Context()), time.Since(start))
		}
	}
}
