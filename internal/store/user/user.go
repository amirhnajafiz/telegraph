package user

import "go.mongodb.org/mongo-driver/bson/primitive"

var Collection = "clients"

type Client struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Name     string             `bson:"from,omitempty"`
	Password string             `bson:"from,omitempty"`
}
