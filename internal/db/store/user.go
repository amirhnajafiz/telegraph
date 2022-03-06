package store

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var UserCollection = "clients"

type Client struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	Name string             `bson:"name,omitempty"`
	Pass string             `bson:"pass,omitempty"`
}

func (Client) Store(database *mongo.Database, ctx context.Context, client *Client) error {
	col := database.Collection(UserCollection)

	_, err := col.InsertOne(ctx, client)
	if err != nil {
		return err
	}

	return nil
}

func (Client) Find(database *mongo.Database, ctx context.Context, name string) (Client, error) {
	col := database.Collection(UserCollection)
	cursor, _ := col.Find(ctx, bson.M{"name": name})

	defer cursor.Close(ctx)

	var data bson.M
	var client Client
	if err := cursor.Decode(&data); err != nil {
		return client, err
	}

	bsonBytes, _ := bson.Marshal(data)
	_ = bson.Unmarshal(bsonBytes, client)

	return client, nil
}
