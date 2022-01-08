package handler

import (
	"Telegraph/internal/handler/publish"
	"Telegraph/internal/handler/root"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func Set(app *echo.Echo, logger *zap.Logger, database *mongo.Database) {
	root.Root{
		Logger: logger.Named("root"),
	}.Register(app.Group("/api"))

	publish.Publish{
		Database: database,
		Logger:   logger.Named("publish"),
	}.Register(app.Group("/api"))
}
