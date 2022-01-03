package serve

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"net/http"
)

func GetServer(logger *zap.Logger) *echo.Echo {
	logger.Info("Initialized server")

	e := echo.New()
	e.GET("api/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	return e
}
