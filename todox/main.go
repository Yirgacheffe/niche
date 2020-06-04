package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// PostNoteHandler - /api/notes
func PostNoteHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	var noteReq, note Notex
	err := json.NewDecoder(r.Body).Decode(&noteReq)
	if err != nil {
		panic(err)
	}

	// Create new note instance
	note.ID = primitive.NewObjectID()
	note.Title = noteReq.Title
	note.Body = noteReq.Body
	note.CreatedAt = time.Now().Local()
	note.UpdatedAt = time.Now().Local()

	// Persistance
	InsertMongoDB(note)

	j, err := json.Marshal(note)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(j)

}

// ListNoteHandler - /api/notes
func ListNoteHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	notes := FindAllNotes()
	notesInJson, err := json.Marshal(notes)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(notesInJson)

}

// GetNoteHandler - /api/notes/{id}
func GetNoteHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	vars := mux.Vars(r)
	objID := vars["id"]

	if len(objID) == 0 {
		log.Fatal("Bad parameter!")
	}

	note := FindByID(objID)

	j, err := json.Marshal(note)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(j)

}

// HealthCheckHandler - Check if it is alive
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	io.WriteString(w, `{"alive": true}`)
}

func main() {

	router := mux.NewRouter()

	api := router.PathPrefix("/api").Subrouter()
	api.HandleFunc("/notes", PostNoteHandler).Methods("POST")
	api.HandleFunc("/notes", ListNoteHandler).Methods("GET")
	api.HandleFunc("/notes/{id}", GetNoteHandler).Methods("GET")
	api.HandleFunc("/health", HealthCheckHandler).Methods("GET")

	api.Handle("/metrics", promhttp.Handler())

	server := &http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	log.Println("Todox Application will start, listening on 8081 ...")
	server.ListenAndServe()

}
