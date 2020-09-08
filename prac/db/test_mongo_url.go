package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://{user}:{password}@10.110.73.215:27017"))

	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(nil)

	collection := client.Database("service-X").Collection("greetings")
	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)

	_, err = collection.InsertOne(ctx, bson.M{"hello": "world"})
	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(ctx)

}
