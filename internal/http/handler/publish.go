package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/amirhnajafiz/Telegraph/internal/db/store"
	"github.com/amirhnajafiz/Telegraph/pkg/validate"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type Publish struct {
	Database *mongo.Database
	Logger   *zap.Logger
	Validate validate.Validate
}

func (publish Publish) Handle(c echo.Context) error {
	valid, data := publish.Validate.PublishValidate(c)

	if valid.Encode() != "" {
		return c.JSON(http.StatusBadRequest, valid)
	}

	item := &store.Message{
		From: data["from"].(string),
		To:   data["to"].(string),
		Msg:  data["message"].(string),
	}

	ctx, endCtx := context.WithTimeout(context.Background(), 10*time.Second)
	defer endCtx()

	err := store.Message{}.Store(publish.Database, ctx, item)
	if err != nil {
		publish.Logger.Error("insert into database failed", zap.Error(err))
	}

	return c.JSON(http.StatusOK, data)
}

func (publish Publish) Register(g *echo.Group) {
	g.POST("/publish", publish.Handle)
}
