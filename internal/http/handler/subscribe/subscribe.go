package subscribe

import (
	"Telegraph/internal/http/request"
	nats2 "Telegraph/internal/nats"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"net/http"
)

type Subscribe struct {
	Logger *zap.Logger
	Nats   nats2.Nats
}

func (s Subscribe) Handle(c echo.Context) error {
	valid, _ := request.SuppressValidate(c)

	if valid.Encode() != "" {
		return c.JSON(http.StatusBadRequest, valid)
	}

	id := c.FormValue("receiver")

	msg, err := s.Nats.Subscribe(id)
	if err != nil {
		return c.String(http.StatusNoContent, "")
	}

	return c.JSON(http.StatusOK, msg)
}

func (s Subscribe) Register(g *echo.Group) {
	g.GET("/", s.Handle)
}
