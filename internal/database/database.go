package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Connect
// make connection to mongodb.
func Connect(cfg Config) (*mongo.Database, error) {
	opts := options.Client()
	opts.ApplyURI(cfg.MongoURL)

	client, err := mongo.NewClient(opts)
	if err != nil {
		return nil, fmt.Errorf("create a new db client faild: %v", err)
	}

	ctx, _ := context.WithTimeout(context.Background(), time.Duration(cfg.ConnectionTimeout)*time.Second)
	if er := client.Connect(ctx); er != nil {
		return nil, fmt.Errorf("mongo connection failed: %v", er)
	}

	if er := client.Ping(context.TODO(), readpref.Primary()); er != nil {
		return nil, fmt.Errorf("monog ping failed: %v", er)
	}

	return client.Database(cfg.Database), nil
}
