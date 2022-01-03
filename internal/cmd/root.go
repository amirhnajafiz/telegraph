package cmd

import (
	"Telegraph/internal/cmd/serve"
	"Telegraph/internal/config"
	"Telegraph/internal/logger"
	"go.uber.org/zap"
)

func Exec() {
	cfg := config.Load()

	log := logger.NewLogger(cfg.Logger)

	e := serve.GetServer(log.Named("serve"))

	err := e.Start(":5000")
	log.Fatal("error starting server", zap.Error(err))
}
