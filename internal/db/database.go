package db

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionTimeout = 10 * time.Second

func NewDB(cfg Config) (*mongo.Database, error) {
	opts := options.Client()
	opts.ApplyURI(cfg.URL)

	client, err := mongo.NewClient(opts)
	if err != nil {
		return nil, fmt.Errorf("create a new db client faild: %w", err)
	}

	return client.Database(cfg.Name), nil
}
