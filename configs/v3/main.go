package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

const (
	healthJson = `{"alive": true}`
	bgColor    = "#483D8B"
)

type Config struct {
	BgColor string `json:"bg_color,omitempty"`
	Version string `json:"version,omitempty"`
	PodName string `json:"pod_name,omitempty"`
}

func newConfig() *Config {
	version := os.Getenv("IMAGE_TAG")
	if len(version) == 0 {
		version = "Not defined"
	}

	podName := os.Getenv("KUBE_POD_NAME")
	if len(podName) == 0 {
		podName = "Not defined"
	}

	return &Config{
		BgColor: bgColor, Version: version, PodName: podName,
	}
}

func configHandler(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())
	rdn := rand.Intn(30)

	if rdn <= 20 {
		time.Sleep(15 * time.Second)
	} else {
		w.Header().Set("Content-Type", "application/json")
		config := newConfig()

		if err := json.NewEncoder(w).Encode(config); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusOK)
	}
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write([]byte(healthJson))
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
	}
	w.WriteHeader(http.StatusOK)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %s", r.Method, r.RequestURI, time.Since(start))
	})
}

func main() {
	router := mux.NewRouter()

	api := router.PathPrefix("/api").Subrouter()
	api.HandleFunc("/configs", configHandler).Methods("GET")
	api.HandleFunc("/health", healthCheckHandler).Methods("GET")

	router.Use(loggingMiddleware)

	server := &http.Server{
		Addr:    ":8081",
		Handler: router,
	}
	log.Println("Config Service starting, listening on 8081 ... ...")
	server.ListenAndServe()
}
