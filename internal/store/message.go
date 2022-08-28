package store

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var MessageCollection = "messages"

type Message struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Client  string             `bson:"client,omitempty"`
	Chat    string             `bson:"chat,omitempty"`
	Message string             `bson:"message,omitempty"`
	Date    time.Time          `bson:"date,omitempty"`
}

func (s *Store) InsertMessage(ctx context.Context, item *Message) error {
	col := s.Database.Collection(MessageCollection)

	item.Date = time.Now()

	if _, err := col.InsertOne(ctx, item); err != nil {
		return err
	}

	return nil
}

func (s *Store) GetChatMessages(ctx context.Context, chat string) ([]Message, error) {
	filter := bson.D{{"chat", chat}}
	col := s.Database.Collection(MessageCollection)

	cursor, _ := col.Find(ctx, filter)
	defer cursor.Close(ctx)

	var (
		message  Message
		messages []Message
	)

	for cursor.Next(ctx) {
		if err := cursor.Decode(&message); err != nil {
			return nil, err
		}

		messages = append(messages, message)
	}

	return messages, nil
}

func (s *Store) DeleteChatMessages(ctx context.Context, chat string) error {
	filter := bson.D{{"chat", chat}}
	col := s.Database.Collection(MessageCollection)

	if _, err := col.DeleteMany(ctx, filter); err != nil {
		return err
	}

	return nil
}
