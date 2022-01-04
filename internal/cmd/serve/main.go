package serve

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func GetServer(logger *zap.Logger) *echo.Echo {
	logger.Info("Initialized server")

	e := echo.New()

	return e
}
