package config

import "Telegraph/internal/logger"

func Default() Config {
	return Config{
		Logger: logger.Config{
			Level: "debug",
		},
	}
}
