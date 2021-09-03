package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	log "github.com/sirupsen/logrus"
)

func main() {

	getEnv := func(key, fallback string) string {
		if v, ok := os.LookupEnv(key); ok {
			return v
		} else {
			return fallback
		}
	}

	dbHost := getEnv("DB_HOST", "127.0.0.1")
	dbPort := getEnv("DB_PORT", "5432")
	dbUser := getEnv("DB_USER", "postgres")
	dbPazz := getEnv("DB_PAZZ", "postgres")
	dbName := getEnv("DB_NAME", "niche_auth")

	db, err := ConnectSQL(dbHost, dbPort, dbUser, dbPazz, dbName)
	if err != nil {
		log.Println(err)
		os.Exit(-1)
	}

	ah := NewAuthHandler(db)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"},
	})

	router := mux.NewRouter()
	handler := c.Handler(router)

	router.HandleFunc("/oauth/auth", withMetrics(ah.Login)).Methods("POST")
	router.HandleFunc("/health",
		func(w http.ResponseWriter, r *http.Request) {
			json.NewEncoder(w).Encode(map[string]bool{"ok": true})
		},
	).Methods("GET", "OPTIONS")

	// Jwks
	f := http.FileServer(http.Dir(".well-known"))
	router.PathPrefix("/.well-known/").Handler(http.StripPrefix("/.well-known/", f))

	log.Info(http.ListenAndServe(":9010", handler))
}
