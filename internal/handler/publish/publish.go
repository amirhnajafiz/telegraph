package publish

import (
	"Telegraph/internal/store/message"
	"Telegraph/pkg/validate"
	"context"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type Publish struct {
	Database *mongo.Database
	Logger   *zap.Logger
}

type Request struct {
	Source string `json:"from"`
	Des    string `json:"to"`
	Msg    string `json:"message"`
}

func (publish Publish) Handle(c echo.Context) error {
	valid, data := validate.ValidatePublish(c)
	if valid.Encode() != "" {
		return c.JSON(http.StatusBadRequest, valid)
	}

	item := &message.Message{
		From: data["from"].(string),
		To:   data["to"].(string),
		Msg:  data["message"].(string),
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err := message.Store(publish.Database, ctx, item)
	if err != nil {
		publish.Logger.Error("insert into database failed", zap.Error(err))
		return err
	}

	// TODO 2: Send the message to the destination
	// TODO 3: Notify the destination

	return c.JSON(http.StatusOK, data)
}

func (publish Publish) Register(g *echo.Group) {
	g.POST("/publish", publish.Handle)
}
