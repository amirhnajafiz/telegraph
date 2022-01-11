package message

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func Store(database *mongo.Database, ctx context.Context, item *Message) error {
	col := database.Collection(Collection)

	item.Date = time.Now()

	_, err := col.InsertOne(ctx, item)
	if err != nil {
		return err
	}

	return nil
}

func All(database *mongo.Database, ctx context.Context) []Message {
	col := database.Collection(Collection)

	cursor, _ := col.Find(ctx, bson.D{})

	var results []Message
	cursor.All(ctx, results)

	return results
}
