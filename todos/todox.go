package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

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

	w.Header().Set("Content-Type", "application/json")

	notes := FindAllNotesInMongoDB()

	notesInJson, err := json.Marshal(notes)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(notesInJson)

}

// GetNoteHandler - /api/notes/{id}
func GetNoteHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

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

func main() {

	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/api/notes", PostNoteHandler).Methods("POST")
	r.HandleFunc("/api/notes", ListNoteHandler).Methods("GET")
	r.HandleFunc("/api/notes/{id}", GetNoteHandler).Methods("GET")

	server := &http.Server{
		Addr:    ":8081",
		Handler: r,
	}

	log.Println("Listening on 8081 ...")
	server.ListenAndServe()

}
