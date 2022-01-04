package handler

import (
	"Telegraph/internal/handler/api"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func Set(app *echo.Echo, logger *zap.Logger) {
	api.Root{
		Logger: logger,
	}.Register(app.Group("/api"))
}
