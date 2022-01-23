package serve

import (
	"Telegraph/internal/http/handler"
	"github.com/labstack/echo/v4"
)

func GetServer(t Tools) *echo.Echo {
	t.Logger.Info("Initialized server")

	e := echo.New()
	handler.Handler{
		Database: t.Database,
		Logger:   t.Logger.Named("handler"),
		Nats:     t.Nats,
	}.Set(e)

	return e
}
