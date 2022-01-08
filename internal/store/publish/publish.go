package publish

import (
	"Telegraph/internal/docs"
	"Telegraph/internal/handler/api"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

var collection = "publications"

func Store(database *mongo.Database, ctx context.Context, r api.Request) error {
	col := database.Collection(collection)

	item := &docs.Publish{
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
