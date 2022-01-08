package store

import (
	"Telegraph/internal/docs"
	"Telegraph/internal/handler/api"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func Publish(database *mongo.Database, ctx context.Context, r api.Request) error {
	col := database.Collection(docs.Collection)

	item := &docs.Message{
		From: r.Source,
		To:   r.Des,
		Msg:  r.Msg,
		Date: time.Now(),
	}

	_, err := col.InsertOne(ctx, item)
	if err != nil {
		return err
	}

	return nil
}
