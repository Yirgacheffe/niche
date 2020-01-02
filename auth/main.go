package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, AllowCredentials: true, AllowedMethods: []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"},
	})

	router := mux.NewRouter()
	router.HandleFunc("/login", LoginHandler).Methods("POST")
	handler := c.Handler(router)

	log.Fatal(http.ListenAndServe(":8081", handler))
}
