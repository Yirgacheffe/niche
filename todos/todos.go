package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/gorilla/mux"
)

type Note struct {
	ID          string    `json:"id,omitempty"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedOn   time.Time `json:"created_on,omitempty"`
}

var noteStore = make(map[string]Note)
var id int = 0

// HTTP Post - /api/notes
func PostNoteHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	var note Note

	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		panic(err)
	}

	note.CreatedOn = time.Now()
	id++

	k := strconv.Itoa(id)
	noteStore[k] = note

	j, err := json.Marshal(note)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(j)

}

// HTTP Get - /api/notes
func GetNoteHandler(w http.ResponseWriter, r *http.Request) {

	var notes []Note

	for _, v := range noteStore {
		notes = append(notes, v)
	}

	w.Header().Set("Content-Type", "application/json")

	j, err := json.Marshal(notes)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(j)

}

// HTTP Put - /api/notes/{id}
func PutNoteHandler(w http.ResponseWriter, r *http.Request) {

	var err error
	vars := mux.Vars(r)
	k := vars["id"]

	var noteToUpd Note

	err = json.NewDecoder(r.Body).Decode(&noteToUpd)
	if err != nil {
		panic(err)
	}

	if note, ok := noteStore[k]; ok {
		noteToUpd.CreatedOn = note.CreatedOn
		delete(noteStore, k)
		noteStore[k] = noteToUpd
	} else {
		log.Printf("Cound not find key of Note %s to delete", k)
	}

	w.WriteHeader(http.StatusNoContent)

}

// HTTP Delete - /api/notes/{id}
func DeleteNoteHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	k := vars["id"]

	if _, ok := noteStore[k]; ok {
		delete(noteStore, k)
	} else {
		log.Printf("Could not find key of Note %s to delete", k)
	}

	w.WriteHeader(http.StatusNoContent)

}

func CallMongoDB(note Note) {

	log.Println(note)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_CONN")))
	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(nil)

	collection := client.Database("todo-app").Collection("notes")
	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)

	_, err = collection.InsertOne(ctx, note)
	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(ctx)

}

func InsertMongoDB(note Note) {

	log.Println(note)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_CONN"))

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(nil)

	collection := client.Database("todo-app").Collection("notes")
	_, err = collection.InsertOne(ctx, note)
	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(ctx)

}

func main() {

	r := mux.NewRouter().StrictSlash(false)

	r.HandleFunc("/api/notes", PostNoteHandler).Methods("POST")
	r.HandleFunc("/api/notes", GetNoteHandler).Methods("GET")
	r.HandleFunc("/api/notes/{id}", PutNoteHandler).Methods("PUT")
	r.HandleFunc("/api/notes/{id}", DeleteNoteHandler).Methods("DELETE")

	server := &http.Server{
		Addr:    ":8081",
		Handler: r,
	}

	log.Println("Listening...")
	server.ListenAndServe()

}
