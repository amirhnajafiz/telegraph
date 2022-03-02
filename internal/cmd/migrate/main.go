package migrate

import (
	"github.com/amirhnajafiz/Telegraph/internal/store"
)

func (m Migrate) Do() {
	_ = m.Database.Collection(store.MessageCollection)
	_ = m.Database.Collection(store.UserCollection)

	m.Logger.Info("collections created")
}
