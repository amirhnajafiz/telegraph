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

func (s *Store) InsertClient(ctx context.Context, client *Client) error {
	col := s.Database.Collection(UserCollection)

	_, err := col.InsertOne(ctx, client)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) FindClient(ctx context.Context, name string) (Client, error) {
	col := s.Database.Collection(UserCollection)
	cursor, _ := col.Find(ctx, bson.M{"name": name})
	defer cursor.Close(ctx)

	var results []bson.M
	cursor.All(ctx, &results)

	if len(results) == 0 {
		return Client{}, mongo.ErrEmptySlice
	}

	var client Client
	bytes, _ := bson.Marshal(results[0])
	bson.Unmarshal(bytes, &client)

	return client, nil
}
