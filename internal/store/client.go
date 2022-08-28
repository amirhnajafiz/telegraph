package store

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var ClientCollection = "clients"

type Client struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Name  string             `bson:"name,omitempty"`
	Pass  string             `bson:"pass,omitempty"`
	Chats []string           `bson:"chats,omitempty"`
}

func (s *Store) NewClient(ctx context.Context, client *Client) error {
	col := s.Database.Collection(ClientCollection)

	_, err := col.InsertOne(ctx, client)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) GetClient(ctx context.Context, name string) (Client, error) {
	filter := bson.D{{"name", name}}
	col := s.Database.Collection(ClientCollection)

	cursor, _ := col.Find(ctx, filter)
	defer cursor.Close(ctx)

	var client Client

	if cursor.Next(ctx) {
		if err := cursor.Decode(&client); err != nil {
			return client, err
		}
	}

	return client, nil
}
