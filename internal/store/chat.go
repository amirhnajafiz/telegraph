package store

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var ChatCollection = "chats"

type Chat struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	Name string             `bson:"name,omitempty"`
}

func (s *Store) NewChat(ctx context.Context, ch *Chat) (interface{}, error) {
	collection := s.Database.Collection(ChatCollection)

	id, err := collection.InsertOne(ctx, ch)
	if err != nil {
		return nil, err
	}

	return id.InsertedID, nil
}

func (s *Store) DeleteChat(ctx context.Context, name string) error {
	filter := bson.D{{"name", name}}

	collection := s.Database.Collection(ChatCollection)

	_, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}
