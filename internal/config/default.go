package config

import (
	"github.com/amirhnajafiz/telegraph/internal/db"
	"github.com/amirhnajafiz/telegraph/internal/logger"
	"github.com/amirhnajafiz/telegraph/internal/nats"
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
			Host: "localhost:4222",
		},
	}
}
