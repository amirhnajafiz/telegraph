package cmd

import (
	"github.com/amirhnajafiz/telegraph/internal/cmd/migrate"
	"github.com/amirhnajafiz/telegraph/internal/cmd/serve"
	"github.com/amirhnajafiz/telegraph/internal/config"
	"github.com/amirhnajafiz/telegraph/internal/db"
	"github.com/amirhnajafiz/telegraph/internal/logger"
	"github.com/amirhnajafiz/telegraph/internal/nats"
	"go.uber.org/zap"
)

func Exec() {
	cfg := config.Load()
	log := logger.NewLogger(cfg.Logger)
	database, er := db.NewDB(cfg.Database)
	if er != nil {
		log.Fatal("database initiation failed", zap.Error(er))
	}

	migrate.Migrate{
		Database: database,
		Logger:   log.Named("migrate"),
	}.Do()

	e := serve.Serve{
		Database: database,
		Logger:   log.Named("serve"),
		Nats: nats.Nats{
			Logger: log.Named("nats"),
			Cfg:    cfg.Nats,
		},
	}.GetServer()

	err := e.Start(":8080")
	log.Fatal("error starting server", zap.Error(err))
}
