package store

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var MessageCollection = "messages"

type Message struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	From string             `bson:"from,omitempty"`
	To   string             `bson:"to,omitempty"`
	Msg  string             `bson:"msg,omitempty"`
	Date time.Time          `bson:"date,omitempty"`
}

func (Message) Store(database *mongo.Database, ctx context.Context, item *Message) error {
	col := database.Collection(MessageCollection)

	item.Date = time.Now()

	_, err := col.InsertOne(ctx, item)
	if err != nil {
		return err
	}

	return nil
}

func (Message) All(database *mongo.Database, ctx context.Context, user string) []bson.M {
	col := database.Collection(MessageCollection)

	cursor, _ := col.Find(ctx, bson.D{{"$or", []interface{}{
		bson.M{"from": user},
		bson.M{"to": user},
	}}})

	defer cursor.Close(ctx)

	var results []bson.M
	cursor.All(ctx, &results)

	return results
}
