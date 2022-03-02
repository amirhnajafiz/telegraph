package migrate

import (
	store2 "github.com/amirhnajafiz/Telegraph/internal/db/store"
)

func (m Migrate) Do() {
	_ = m.Database.Collection(store2.MessageCollection)
	_ = m.Database.Collection(store2.UserCollection)

	m.Logger.Info("collections created")
}
