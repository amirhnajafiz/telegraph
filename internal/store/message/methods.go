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

func All(database *mongo.Database, ctx context.Context, user string) []bson.M {
	col := database.Collection(Collection)

	cursor, _ := col.Find(ctx, bson.D{{"$or", []interface{}{
		bson.M{"from": user},
		bson.M{"to": user},
	}}})

	defer cursor.Close(ctx)

	var results []bson.M
	cursor.All(ctx, &results)

	return results
}
