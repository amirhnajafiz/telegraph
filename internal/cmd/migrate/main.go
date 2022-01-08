package migrate

import (
	"Telegraph/internal/docs"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type Requirements struct {
	Database *mongo.Database
	Logger   *zap.Logger
}

func Do(r Requirements) {
	_ = r.Database.Collection(docs.Collection)

	r.Logger.Info("collections created")
}
