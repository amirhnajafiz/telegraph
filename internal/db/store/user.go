package store

import (
	"context"
	"fmt"

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
	var (
		data   bson.M
		client Client
	)

	col := database.Collection(UserCollection)
	cursor, err := col.Find(ctx, bson.M{"name": name})

	if err != nil {
		fmt.Println(err)
		return client, err
	}

	defer cursor.Close(ctx)
	if err := cursor.Decode(&data); err != nil {
		return client, err
	}

	bsonBytes, _ := bson.Marshal(data)
	_ = bson.Unmarshal(bsonBytes, client)

	return client, nil
}
