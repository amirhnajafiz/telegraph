package migrate

import (
	"Telegraph/internal/store/message"
	"Telegraph/internal/store/user"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type Requirements struct {
	Database *mongo.Database
	Logger   *zap.Logger
}

func Do(r Requirements) {
	_ = r.Database.Collection(message.Collection)
	_ = r.Database.Collection(user.Collection)

	r.Logger.Info("collections created")
}
