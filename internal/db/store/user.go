package store

import "go.mongodb.org/mongo-driver/bson/primitive"

var UserCollection = "clients"

type Client struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	Name string             `bson:"from,omitempty"`
}
