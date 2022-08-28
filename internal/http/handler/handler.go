package handler

import (
	"github.com/amirhnajafiz/Telegraph/internal/nats"
	"github.com/amirhnajafiz/Telegraph/internal/validate"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type Handler struct {
	Database *mongo.Database
	Logger   *zap.Logger
	Nats     nats.Nats
	Validate validate.Validate
}

func (h Handler) Set(app *echo.Echo) {
	Join{
		Database: h.Database,
		Logger:   h.Logger.Named("join"),
		Validate: h.Validate,
	}.Register(app.Group("/api"))

	Publish{
		Database:       h.Database,
		Logger:         h.Logger.Named("publish"),
		NatsConnection: h.Nats.New(),
		Validate:       h.Validate,
	}.Register(app.Group("/api"))

	Suppress{
		Database: h.Database,
		Logger:   h.Logger.Named("suppress"),
	}.Register(app.Group("/api"))
}
