package serve

import (
	"github.com/amirhnajafiz/telegraph/internal/config"
	"github.com/amirhnajafiz/telegraph/internal/database"
	"github.com/amirhnajafiz/telegraph/internal/http/handler"
	"github.com/amirhnajafiz/telegraph/internal/logger"
	"github.com/amirhnajafiz/telegraph/internal/nats"
	"github.com/amirhnajafiz/telegraph/internal/validate"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func GetCommand() *cobra.Command {
	return &cobra.Command{
		Short: "serve",
		Run: func(_ *cobra.Command, _ []string) {
			main()
		},
	}
}

func main() {
	cfg := config.Load()
	log := logger.NewLogger(cfg.Logger)

	db, er := database.Connect(cfg.MongoDB)
	if er != nil {
		log.Fatal("database initiation failed", zap.Error(er))
	}

	nat, err := nats.New(cfg.Nats)
	if err != nil {
		log.Fatal("nats connection failed", zap.Error(err))
	}

	e := echo.New()

	handler.Handler{
		Database: db,
		Logger:   log.Named("handler"),
		Nats:     nat,
		Validate: validate.Validate{},
	}.Set(e.Group("/api"))

	if err := e.Start(":8080"); err != nil {
		log.Fatal("error starting server", zap.Error(err))
	}
}
