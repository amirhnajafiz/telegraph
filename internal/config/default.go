package config

import (
	"github.com/amirhnajafiz/Telegraph/internal/db"
	"github.com/amirhnajafiz/Telegraph/internal/logger"
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
	}
}
