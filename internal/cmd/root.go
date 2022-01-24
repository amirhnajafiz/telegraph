package cmd

import (
	"Telegraph/internal/cmd/migrate"
	"Telegraph/internal/cmd/serve"
	"Telegraph/internal/config"
	"Telegraph/internal/db"
	"Telegraph/internal/logger"
	"Telegraph/internal/nats"
	"go.uber.org/zap"
)

func Exec() {
	cfg := config.Load()

	log := logger.NewLogger(cfg.Logger)

	database, er := db.NewDB(cfg.Database)
	if er != nil {
		log.Fatal("database initiation failed", zap.Error(er))
	}

	migrate.Do(migrate.Requirements{
		Database: database,
		Logger:   log,
	})

	n := nats.Nats{
		Logger: *log.Named("nats"),
		Conf:   cfg.Nats,
	}
	n.Connection = n.Setup()

	e := serve.GetServer(serve.Tools{
		Database: database,
		Logger:   log.Named("serve"),
		Nats:     n,
	})

	err := e.Start(":5000")
	log.Fatal("error starting server", zap.Error(err))
}
