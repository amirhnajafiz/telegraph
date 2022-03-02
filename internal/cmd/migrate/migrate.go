package migrate

import (
	store2 "github.com/amirhnajafiz/Telegraph/internal/db/store"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type Migrate struct {
	Database *mongo.Database
	Logger   *zap.Logger
}

func (m Migrate) Do() {
	_ = m.Database.Collection(store2.MessageCollection)
	_ = m.Database.Collection(store2.UserCollection)

	m.Logger.Info("collections created")
}
