package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	log "github.com/sirupsen/logrus"
)

func main() {

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, AllowCredentials: true, AllowedMethods: []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"},
	})

	router := mux.NewRouter()

	handler := c.Handler(router)
	router.HandleFunc("/login", LoginHandler).Methods("POST")
	router.HandleFunc("/health", HealthCheckHandler).Methods("GET", "OPTIONS")

	// Jwks
	f := http.FileServer(http.Dir(".well-known"))
	router.PathPrefix("/.well-known/").Handler(http.StripPrefix("/.well-known/", f))

	log.Info(http.ListenAndServe(":8081", handler))

}
