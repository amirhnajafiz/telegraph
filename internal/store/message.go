package store

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var MessageCollection = "messages"

type Message struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Sender string             `bson:"sender,omitempty"`
	Msg    string             `bson:"msg,omitempty"`
	Date   time.Time          `bson:"date,omitempty"`
}

func (s *Store) InsertMessage(ctx context.Context, item *Message) (interface{}, error) {
	col := s.Database.Collection(MessageCollection)
	item.Date = time.Now()

	res, err := col.InsertOne(ctx, item)
	if err != nil {
		return nil, err
	}

	return res.InsertedID, nil
}

func (s *Store) GetAllMessages(ctx context.Context, user string) []bson.M {
	col := s.Database.Collection(MessageCollection)
	cursor, _ := col.Find(ctx, bson.M{"sender": user})
	defer cursor.Close(ctx)

	var results []bson.M
	cursor.All(ctx, &results)

	return results
}
