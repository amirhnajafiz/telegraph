package nats

import (
	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
)

type Nats struct {
	Logger     zap.Logger
	Conf       Config
	Connection *nats.Conn
}

func (n Nats) Setup() *nats.Conn {
	nc, err := nats.Connect(n.Conf.URL)
	if err != nil {
		n.Logger.Error("nats connection failed", zap.Error(err))
	}

	return nc
}

func (n Nats) Publish(subject string, message []byte) {
	e := n.Connection.Publish(subject, message)
	if e != nil {
		n.Logger.Error("nats publishing failed", zap.Error(e))
	}
}
