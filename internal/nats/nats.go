package nats

import (
	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
	"time"
)

type Nats struct {
	Logger     *zap.Logger
	Conf       Config
	Connection *nats.Conn
}

var timeout = time.Second * 5

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
		n.Logger.Info("success")
		n.Logger.Error("nats publishing failed", zap.Error(e))
	}
}

func (n Nats) Subscribe(subject string) (*nats.Msg, error) {
	sub, err := n.Connection.SubscribeSync(subject)
	if err != nil {
		n.Logger.Error("nats subscribe sync failed", zap.Error(err))
		return nil, err
	}

	msg, err := sub.NextMsg(timeout)
	if err != nil {
		n.Logger.Error("nats subscribe message getting failed", zap.Error(err))
		return nil, err
	}

	return msg, nil
}
