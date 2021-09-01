package main

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	DBName = "todox-app"
	URI    = "mongodb://root:password@localhost:27017"

	notesColl = "notes"
)

var ctx = context.Background()
var db *mongo.Database

func init() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Get MongoDB Connection
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

	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	db = client.Database(DBName)
}

// FindAllNotes - Query Notex from database
func FindAllNotes() []Notex {

	cursor, err := db.Collection(notesColl).Find(ctx, bson.M{})
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

// InsertMongoDB - Create new Notex
func InsertMongoDB(note Notex) string {
	log.Println(note)

	result, err := db.Collection(notesColl).InsertOne(ctx, note)
	if err != nil {
		log.Fatal(err)
	}

	/*
		notes := []interface{}{ note1, note2}
		res, err := db.Collection(notesColl).InsertMany(ctx, notes)
	*/

	id := result.InsertedID.(primitive.ObjectID)

	log.Println("Insert() with ID: ", id)
	return id.Hex()
}

func UpateNote(note Notex) {

	filter := bson.M{"_id": note.ID}
	update := bson.M{
		"$set": bson.M{
			"title": note.Title, "body": note.Body,
		},
	}

	r, err := db.Collection(notesColl).UpdateOne(ctx, filter, update)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("%d documents updated", r.ModifiedCount)
}

// FindByID - Grab note by ID
func FindByID(id string) Notex {

	log.Println(id)

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal("Bad ID parameter: ", id)
	}

	coll := db.Collection(notesColl)

	result := coll.FindOne(ctx, bson.M{"_id": objID})
	if err = result.Err(); err != nil {
		log.Fatal(err)
	}

	note := Notex{}
	err = result.Decode(&note)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(note.Body)
	return note

}

// Don't know why add this bottom line here -----------------
