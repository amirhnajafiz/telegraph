package serve

import (
	"github.com/nats-io/nats.go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type Tools struct {
	Database *mongo.Database
	Logger   *zap.Logger
	Nats     *nats.Conn
}
