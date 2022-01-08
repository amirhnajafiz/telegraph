package message

import (
	"Telegraph/internal/handler/api"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func Store(database *mongo.Database, ctx context.Context, r api.Request) error {
	col := database.Collection(Collection)

	item := &Message{
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
