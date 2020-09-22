package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var (
	dbHost, dbPort, dbUser, dbPazz, dbName string
)

func main() {

	dbHost := os.Getenv("DB_HOST")
	if len(dbHost) == 0 {
		dbHost = "127.0.0.1"
	}

	dbPort := os.Getenv("DB_PORT")
	if len(dbPort) == 0 {
		dbPort = "3306"
	}

	dbUser := os.Getenv("DB_USER")
	if len(dbUser) == 0 {
		dbUser = "tollgate"
	}

	dbPazz := os.Getenv("DB_PASS")
	if len(dbPazz) == 0 {
		dbPazz = "zMUjpGW67o"
	}

	dbName := os.Getenv("DB_NAME")
	if len(dbName) == 0 {
		dbName = "TOLLGATE"
	}

	db, err := ConnectSQL(dbHost, dbPort, dbUser, dbPazz, dbName)
	if err != nil {
		log.Println(err)
		os.Exit(-1)
	}

	fh := NewFileHandler(db)

	// Start the server
	router := mux.NewRouter()

	api := router.PathPrefix("/api").Subrouter()
	api.HandleFunc("/files/{id:[0-9]+}", fh.DetailsHandler).Methods("HEAD")
	api.HandleFunc("/files/{id:[0-9]+}", fh.PatchFileHandler).Methods("PATCH")

	api.HandleFunc("/health", HealthHandler).Methods("GET")

	server := &http.Server{
		Addr:    ":8083",
		Handler: router,
	}

	log.Println("Tus application will start, listening on 8083 ...")
	server.ListenAndServe()

}
