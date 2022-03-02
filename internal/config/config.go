package config

import (
	"encoding/json"
	"log"

	"github.com/amirhnajafiz/Telegraph/internal/db"
	"github.com/amirhnajafiz/Telegraph/internal/logger"
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/structs"
)

type (
	Config struct {
		Database db.Config     `koanf:"database"`
		Logger   logger.Config `koanf:"logger"`
	}
)

func Load() Config {
	var instance Config

	k := koanf.New(".")

	// load default configuration from default function
	if err := k.Load(structs.Provider(Default(), "koanf"), nil); err != nil {
		log.Fatalf("error loading default: %s", err)
	}

	// load configuration from file
	if err := k.Load(file.Provider("config.yml"), yaml.Parser()); err != nil {
		log.Printf("error loading config.yml: %s", err)
	}

	if err := k.Unmarshal("", &instance); err != nil {
		log.Fatalf("error unmarshalling config: %s", err)
	}

	indent, _ := json.MarshalIndent(instance, "", "\t")
	tmpl := `
	================ Loaded Configuration ================
	%s
	======================================================
	`
	log.Printf(tmpl, string(indent))

	return instance
}
