package nats

import (
	"github.com/nats-io/nats.go"
)

func New(cfg Config) (*nats.Conn, error) {
	nc, err := nats.Connect(cfg.Host)
	if err != nil {
		return nil, err
	}

	return nc, nil
}
