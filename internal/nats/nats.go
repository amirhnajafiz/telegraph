package nats

import (
	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
)

type Nats struct {
	Logger zap.Logger
	Conf   Config
}

func (n Nats) Setup() *nats.Conn {
	nc, err := nats.Connect(n.Conf.URL)
	if err != nil {
		n.Logger.Error("nats connection failed", zap.Error(err))
	}

	return nc
}
