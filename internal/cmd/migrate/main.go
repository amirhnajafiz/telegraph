package migrate

import (
	"Telegraph/internal/store/message"
	"Telegraph/internal/store/user"
)

func (m Migrate) Do() {
	_ = m.Database.Collection(message.Collection)
	_ = m.Database.Collection(user.Collection)

	m.Logger.Info("collections created")
}
