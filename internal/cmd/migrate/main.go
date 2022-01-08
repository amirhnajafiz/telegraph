package migrate

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type Requirements struct {
	Database *mongo.Database
	Logger   *zap.Logger
}

func Do(r Requirements) {
	_ = r.Database.Collection("publications")
	_ = r.Database.Collection("users")

	r.Logger.Info("collections created")
}
