package nats

import (
	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
)

type Nats struct {
	Logger *zap.Logger
	Cfg    Config
}

func (n Nats) New() *nats.Conn {
	nc, err := nats.Connect(n.Cfg.Host)
	if err != nil {
		n.Logger.Fatal("nats connection failed", zap.Error(err))
	}

	return nc
}
