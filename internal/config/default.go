package config

import (
	"Telegraph/internal/db"
	"Telegraph/internal/logger"
	"Telegraph/internal/nats"
)

func Default() Config {
	return Config{
		Logger: logger.Config{
			Level: "debug",
		},
		Database: db.Config{
			Name: "telegraph",
			URL:  "mongodb://127.0.0.1:27017",
		},
		Nats: nats.Config{
			URL: "127.0.0.1:4222",
		},
	}
}
