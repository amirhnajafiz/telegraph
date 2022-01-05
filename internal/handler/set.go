package handler

import (
	"Telegraph/internal/handler/api"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func Set(app *echo.Echo, logger *zap.Logger) {
	api.Root{
		Logger: logger.Named("root"),
	}.Register(app.Group("/api"))

	api.Publish{
		Logger: logger.Named("publish"),
	}.Register(app.Group("/api"))
}
