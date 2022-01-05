package cmd

import (
	"Telegraph/internal/cmd/serve"
	"Telegraph/internal/config"
	"Telegraph/internal/db"
	"Telegraph/internal/logger"
	"go.uber.org/zap"
)

func Exec() {
	cfg := config.Load()

	log := logger.NewLogger(cfg.Logger)

	db, er := db.NewDB(cfg.Database)
	if er != nil {
		log.Fatal("database initiation failed", zap.Error(er))
	}

	e := serve.GetServer(log.Named("serve"))

	err := e.Start(":5000")
	log.Fatal("error starting server", zap.Error(err))
}
