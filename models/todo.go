package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Todo struct {
	// When we call Title, it will come as title
	// omitempty: Does not allow empty data to be sent
	Id      primitive.ObjectID `json:"id,omitempty"`
	Title   string             `json:"title,omitempty"`
	Content string             `json:"content,omitempty"`
}
