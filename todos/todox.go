package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Notex struct {
	ID        primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Title     string             `json:"title"`
	Body      string             `json:"body"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at,omitempty"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at,omitempty"`
}

const (
	DBName = "todos-app"
	URI    = "mongodb://root:DsoN4DVgY5@localhost:27017"

	notesCollection = "notes"
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

// GetNoteHandler - /api/notes
func GetNoteHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	notes := FindAllNotesInMongoDB()

	notesInJson, err := json.Marshal(notes)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(notesInJson)

}

func FindAllNotesInMongoDB() []Notex {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Get mongo connection string uri
	var connURI = os.Getenv("MONGO_CONN")

	if len(connURI) == 0 {
		connURI = URI
	}

	// Get connection
	clientOpts := options.Client().ApplyURI(connURI)
	client, err := mongo.Connect(context.Background(), clientOpts)
	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(ctx)

	// Get db and collection, create new note
	db := client.Database(DBName)
	collection := db.Collection(notesCollection)

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	// Query all notes collection & loop
	notesResult := []Notex{}
	note := Notex{}

	for cursor.Next(ctx) {
		cursor.Decode(&note)
		notesResult = append(notesResult, note)
	}

	cursor.Close(context.TODO())
	return notesResult
}

func InsertMongoDB(note Notex) {
	log.Println(note)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Get mongo connection string uri
	var connURI = os.Getenv("MONGO_CONN")

	if len(connURI) == 0 {
		connURI = URI
	}

	// Get connection
	clientOpts := options.Client().ApplyURI(connURI)
	client, err := mongo.Connect(context.Background(), clientOpts)
	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(ctx)

	// Get db and collection, create new note
	db := client.Database(DBName)
	collection := db.Collection(notesCollection)

	result, err := collection.InsertOne(ctx, note)
	if err != nil {
		log.Fatal(err)
	}

	objID := result.InsertedID.(primitive.ObjectID)
	log.Println(objID)
}

func main() {

	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/api/notes", PostNoteHandler).Methods("POST")
	r.HandleFunc("/api/notes", GetNoteHandler).Methods("GET")

	server := &http.Server{
		Addr:    ":8081",
		Handler: r,
	}

	log.Println("Listening on 8081 ...")
	server.ListenAndServe()

}
