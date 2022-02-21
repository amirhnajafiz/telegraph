package migrate

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type Migrate struct {
	Database *mongo.Database
	Logger   *zap.Logger
}
