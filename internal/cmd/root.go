package cmd

import (
	"Telegraph/internal/cmd/migrate"
	"Telegraph/internal/cmd/serve"
	"Telegraph/internal/config"
	"Telegraph/internal/db"
	"Telegraph/internal/logger"
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
	}.GetServer()

	err := e.Start(":5000")
	log.Fatal("error starting server", zap.Error(err))
}
