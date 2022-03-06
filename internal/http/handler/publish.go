package handler

import (
	"context"
	"fmt"
	"github.com/amirhnajafiz/Telegraph/pkg/jwt"
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

	if auth, err := jwt.ParseToken(c.Request().Header.Get("jwt-token")); err != nil || !auth {
		return fmt.Errorf("unauthorized user")
	}

	item := &store.Message{
		Sender: data["sender"].(string),
		Msg:    data["message"].(string),
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err := store.Message{}.Store(publish.Database, ctx, item)
	if err != nil {
		publish.Logger.Error("insert into database failed", zap.Error(err))
	}

	return c.JSON(http.StatusOK, data)
}

func (publish Publish) Register(g *echo.Group) {
	g.POST("/publish", publish.Handle)
}
