package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/amirhnajafiz/Telegraph/internal/db/store"
	"github.com/amirhnajafiz/Telegraph/pkg/jwt"
	"github.com/amirhnajafiz/Telegraph/pkg/validate"
	"github.com/labstack/echo/v4"
	"github.com/nats-io/nats.go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type Publish struct {
	Database       *mongo.Database
	Logger         *zap.Logger
	NatsConnection *nats.Conn
	Validate       validate.Validate
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
	ctx, endCtx := context.WithTimeout(context.Background(), 10*time.Second)
	defer endCtx()

	err := store.Message{}.Store(publish.Database, ctx, item)
	if err != nil {
		publish.Logger.Error("insert into database failed", zap.Error(err))
	}

	js, _ := json.Marshal(item)
	if err := publish.NatsConnection.Publish("telegraph/chat", js); err != nil {
		publish.Logger.Fatal("nats publish failed", zap.Error(err))
	}

	return c.JSON(http.StatusOK, item)
}

func (publish Publish) Register(g *echo.Group) {
	g.POST("/publish", publish.Handle)
}
