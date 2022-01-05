package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

const connectionTimeout = 10 * time.Second

func NewDB(cfg Config) (*mongo.Database, error) {
	opts := options.Client()
	opts.ApplyURI(cfg.URL)

	client, err := mongo.NewClient(opts)
	if err != nil {
		return nil, fmt.Errorf("create a new db client faild: %w", err)
	}

	// connect to the mongodb
	{
		ctx, done := context.WithTimeout(context.Background(), connectionTimeout)
		defer done()

		if err := client.Connect(ctx); err != nil {
			return nil, fmt.Errorf("database connection error: %w", err)
		}
	}
	// ping the mongodb
	{
		ctx, done := context.WithTimeout(context.Background(), connectionTimeout)
		defer done()

		if err := client.Ping(ctx, readpref.Primary()); err != nil {
			return nil, fmt.Errorf("database ping error: %w", err)
		}
	}

	return client.Database(cfg.Name), nil
}
