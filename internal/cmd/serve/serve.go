package serve

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type Serve struct {
	Database *mongo.Database
	Logger   *zap.Logger
}
