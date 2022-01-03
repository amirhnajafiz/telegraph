package config

import "Telegraph/internal/logger"

type (
	Config struct {
		Logger logger.Config `koanf:"logger"`
	}
)
