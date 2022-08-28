package config

import (
	"github.com/amirhnajafiz/telegraph/internal/database"
	"github.com/amirhnajafiz/telegraph/internal/logger"
	"github.com/amirhnajafiz/telegraph/internal/nats"
)

func Default() Config {
	return Config{
		Logger: logger.Config{
			Level: "debug",
		},
		MongoDB: database.Config{
			ConnectionTimeout: 10,
			Database:          "telegraph",
			MongoURL:          "mongodb://127.0.0.1:27017",
		},
		Nats: nats.Config{
			Host: "localhost:4222",
		},
	}
}
