package store

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var UserCollection = "clients"

type Client struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	Name string             `bson:"from,omitempty"`
}

func (c Client) Store(database *mongo.Database, ctx context.Context, client *Client) error {
	col := database.Collection(UserCollection)

	_, err := col.InsertOne(ctx, client)
	if err != nil {
		return err
	}

	return nil
}
