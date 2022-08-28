package migrate

import (
	"github.com/amirhnajafiz/telegraph/internal/store"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type Migrate struct {
	Database *mongo.Database
	Logger   *zap.Logger
}

func (m Migrate) Do() {
	_ = m.Database.Collection(store.MessageCollection)
	_ = m.Database.Collection(store.UserCollection)

	m.Logger.Info("collections created")
}
