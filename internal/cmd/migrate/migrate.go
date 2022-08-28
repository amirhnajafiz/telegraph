package migrate

import (
	"github.com/amirhnajafiz/telegraph/internal/config"
	"github.com/amirhnajafiz/telegraph/internal/database"
	"github.com/amirhnajafiz/telegraph/internal/logger"
	"github.com/amirhnajafiz/telegraph/internal/store"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func GetCommand() *cobra.Command {
	return &cobra.Command{
		Short: "migrate",
		Run: func(_ *cobra.Command, _ []string) {
			main()
		},
	}
}

func main() {
	cfg := config.Load()
	log := logger.NewLogger(cfg.Logger)

	db, err := database.Connect(cfg.MongoDB)
	if err != nil {
		log.Error("mongo connection failed", zap.Error(err))

		return
	}

	_ = db.Collection(store.MessageCollection)
	_ = db.Collection(store.ChatCollection)
	_ = db.Collection(store.ClientCollection)

	log.Info("collections created")
}
