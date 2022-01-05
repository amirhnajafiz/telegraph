package handler

import (
	"Telegraph/internal/handler/api"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func Set(app *echo.Echo, logger *zap.Logger, database *mongo.Database) {
	api.Root{
		Logger: logger.Named("root"),
	}.Register(app.Group("/api"))

	api.Publish{
		Database: database,
		Logger:   logger.Named("publish"),
	}.Register(app.Group("/api"))
}
