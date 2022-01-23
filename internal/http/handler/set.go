package handler

import (
	"Telegraph/internal/http/handler/publish"
	"Telegraph/internal/http/handler/root"
	"Telegraph/internal/http/handler/subscribe"
	"Telegraph/internal/http/handler/suppress"
	"github.com/labstack/echo/v4"
	"github.com/nats-io/nats.go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type Handler struct {
	Database *mongo.Database
	Logger   *zap.Logger
	Nats     *nats.Conn
}

func (h Handler) Set(app *echo.Echo) {
	root.Root{
		Logger: h.Logger.Named("root"),
	}.Register(app.Group("/api"))

	publish.Publish{
		Database: h.Database,
		Logger:   h.Logger.Named("publish"),
	}.Register(app.Group("/api"))

	subscribe.Subscribe{
		Logger: h.Logger.Named("subscribe"),
		Nats:   h.Nats,
	}.Register(app.Group("/api"))

	suppress.Suppress{
		Database: h.Database,
		Logger:   h.Logger.Named("suppress"),
	}.Register(app.Group("/api"))
}
