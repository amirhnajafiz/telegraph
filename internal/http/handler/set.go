package handler

import (
	"github.com/amirhnajafiz/Telegraph/internal/http/handler/publish"
	"github.com/amirhnajafiz/Telegraph/internal/http/handler/root"
	"github.com/amirhnajafiz/Telegraph/internal/http/handler/subscribe"
	"github.com/amirhnajafiz/Telegraph/internal/http/handler/suppress"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type Handler struct {
	Database *mongo.Database
	Logger   *zap.Logger
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
	}.Register(app.Group("/api"))

	suppress.Suppress{
		Database: h.Database,
		Logger:   h.Logger.Named("suppress"),
	}.Register(app.Group("/api"))
}
