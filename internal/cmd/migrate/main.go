package migrate

import (
	"Telegraph/internal/store/message"
	"Telegraph/internal/store/user"
)

func Do(r Requirements) {
	_ = r.Database.Collection(message.Collection)
	_ = r.Database.Collection(user.Collection)

	r.Logger.Info("collections created")
}
