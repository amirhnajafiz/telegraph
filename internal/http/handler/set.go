package handler

import (
	"Telegraph/internal/http/handler/publish"
	"Telegraph/internal/http/handler/root"
	"Telegraph/internal/http/handler/suppress"
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

	suppress.Suppress{
		Database: database,
		Logger:   logger.Named("suppress"),
	}.Register(app.Group("/api"))
}
