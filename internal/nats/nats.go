package nats

import (
	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
)

type Config struct {
	Port string `koanf:"port"`
}

type Nats struct {
	Logger zap.Logger
	Conf   Config
}

func (n Nats) Setup() *nats.Conn {
	nc, err := nats.Connect(n.Conf.Port)
	if err != nil {
		n.Logger.Error("nats connection failed", zap.Error(err))
	}

	return nc
}
