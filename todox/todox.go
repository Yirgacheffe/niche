package main

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Notex struct {
	ID        primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Title     string             `json:"title"`
	Body      string             `json:"body"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at,omitempty"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at,omitempty"`
}
