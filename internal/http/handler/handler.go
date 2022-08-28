package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/amirhnajafiz/telegraph/internal/jwt"
	"github.com/amirhnajafiz/telegraph/internal/store"
	"github.com/amirhnajafiz/telegraph/internal/validate"
	"github.com/labstack/echo/v4"
	"github.com/nats-io/nats.go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type Handler struct {
	Database *mongo.Database
	Logger   *zap.Logger
	Nats     *nats.Conn
	Validate validate.Validate
}

func (h *Handler) join(c echo.Context) error {
	ctx, endCtx := context.WithTimeout(context.Background(), 10*time.Second)
	defer endCtx()

	valid, data := h.Validate.JoinValidate(c)
	if valid.Encode() != "" {
		return c.JSON(http.StatusBadRequest, valid)
	}

	user := data["username"].(string)
	pass := data["password"].(string)
	if user == "" {
		return c.String(http.StatusBadRequest, "must have a username")
	}

	client, err := store.Client{}.Find(h.Database, ctx, user)

	if err != mongo.ErrEmptySlice && client.Pass != pass {
		return c.String(http.StatusUnauthorized, "username and password mismatched")
	}

	if err == mongo.ErrEmptySlice {
		_ = store.Client{}.Store(h.Database, ctx, &store.Client{
			Name: user,
			Pass: pass,
		})
	}

	token, err := jwt.GenerateToken(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "token generating failed")
	}

	return c.String(http.StatusOK, token)
}

func (h *Handler) publish(c echo.Context) error {
	valid, data := h.Validate.PublishValidate(c)
	if valid.Encode() != "" {
		return c.JSON(http.StatusBadRequest, valid)
	}

	item := &store.Message{
		Sender: data["sender"].(string),
		Msg:    data["message"].(string),
	}
	ctx, endCtx := context.WithTimeout(context.Background(), 10*time.Second)
	defer endCtx()

	err := store.Message{}.Store(h.Database, ctx, item)
	if err != nil {
		h.Logger.Error("insert into database failed", zap.Error(err))
	}

	js, _ := json.Marshal(item)
	if err := h.Nats.Publish("telegraph/chat", js); err != nil {
		h.Logger.Fatal("nats publish failed", zap.Error(err))
	}

	return c.JSON(http.StatusOK, item)
}

func (h *Handler) suppress(c echo.Context) error {
	valid, _ := h.Validate.SuppressValidate(c)
	if valid.Encode() != "" {
		return c.JSON(http.StatusBadRequest, valid)
	}

	user := c.FormValue("sender")
	ctx, endCtx := context.WithTimeout(context.Background(), 10*time.Second)
	defer endCtx()

	res := store.Message{}.All(h.Database, ctx, user)

	return c.JSON(http.StatusOK, res)
}

func (h Handler) Set(app *echo.Group) {
	app.POST("/join", h.join)
	app.POST("/publish", h.publish)
	app.GET("/suppress", h.suppress)
}
