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
	_ = r.Database.Collection("podcasts")
	_ = r.Database.Collection("episodes")

	r.Logger.Info("collections created")
}
