package serve

import (
	nats2 "Telegraph/internal/nats"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type Tools struct {
	Database *mongo.Database
	Logger   *zap.Logger
	Nats     nats2.Nats
}
