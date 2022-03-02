package migrate

import (
	"github.com/amirhnajafiz/Telegraph/internal/store/message"
	"github.com/amirhnajafiz/Telegraph/internal/store/user"
)

func (m Migrate) Do() {
	_ = m.Database.Collection(message.Collection)
	_ = m.Database.Collection(user.Collection)

	m.Logger.Info("collections created")
}
