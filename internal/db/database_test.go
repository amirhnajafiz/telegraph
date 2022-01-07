package db

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"testing"
)

func TestNewDB(t *testing.T) {
	opts := options.Client()
	opts.ApplyURI(cfg.URL)

	client, err := mongo.NewClient(opts)
	if err != nil {
		assert.Fail(t, "Creating db client failed", err)
	}

	// connect to the mongodb
	{
		ctx, done := context.WithTimeout(context.Background(), connectionTimeout)
		defer done()

		if err := client.Connect(ctx); err != nil {
			assert.Error(t, err, "database connection error")
		}
	}
	// ping the mongodb
	{
		ctx, done := context.WithTimeout(context.Background(), connectionTimeout)
		defer done()

		if err := client.Ping(ctx, readpref.Primary()); err != nil {
			assert.Error(t, err, "database ping error")
		}
	}
}
