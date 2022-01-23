package subscribe

import (
	"Telegraph/internal/http/request"
	"github.com/labstack/echo/v4"
	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type Subscribe struct {
	Logger *zap.Logger
	Nats   *nats.Conn
}

var timeout = time.Second * 5

func (s Subscribe) Handle(c echo.Context) error {
	valid, _ := request.SuppressValidate(c)

	if valid.Encode() != "" {
		return c.JSON(http.StatusBadRequest, valid)
	}

	id := c.FormValue("receiver")

	sub, err := s.Nats.SubscribeSync(id)
	if err != nil {
		s.Logger.Error("nats subscribe sync failed", zap.Error(err))
		return c.String(http.StatusInternalServerError, "nats failed")
	}

	msg, err := sub.NextMsg(timeout)
	if err != nil {
		s.Logger.Error("nats subscribe message getting failed", zap.Error(err))
		return c.String(http.StatusNoContent, "no message")
	}

	return c.JSON(http.StatusOK, msg)
}

func (s Subscribe) Register(g *echo.Group) {
	g.GET("/", s.Handle)
}
