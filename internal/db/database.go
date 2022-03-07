package db

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const connectionTimeout = 10 * time.Second

func NewDB(cfg Config) (*mongo.Database, error) {
	opts := options.Client()
	opts.ApplyURI(cfg.URL)

	client, err := mongo.NewClient(opts)
	if err != nil {
		return nil, fmt.Errorf("create a new db client faild: %w", err)
	}

	if err := client.Connect(context.Background()); err != nil {
		return nil, err
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		// Can't connect to Mongo server
		return nil, err
	}

	return client.Database(cfg.Name), nil
}
