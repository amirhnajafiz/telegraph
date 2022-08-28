package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/amirhnajafiz/telegraph/internal/http/middleware"
	"github.com/amirhnajafiz/telegraph/internal/jwt"
	"github.com/amirhnajafiz/telegraph/internal/store"
	"github.com/amirhnajafiz/telegraph/internal/validate"
	"github.com/labstack/echo/v4"
	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
)

type Handler struct {
	Logger   *zap.Logger
	Nats     *nats.Conn
	Validate validate.Validate
	Store    store.Store
}

func (h *Handler) login(c echo.Context) error {
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

	client, err := h.Store.GetClient(ctx, user)
	if err != nil || client.Pass != pass {
		return c.String(http.StatusUnauthorized, err.Error())
	}

	token, err := jwt.GenerateToken(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "token generating failed")
	}

	return c.String(http.StatusOK, token)
}

func (h *Handler) join(c echo.Context) error {
	ctx, endCtx := context.WithTimeout(context.Background(), 10*time.Second)
	defer endCtx()

	valid, data := h.Validate.JoinValidate(c)
	if valid.Encode() != "" {
		return c.JSON(http.StatusBadRequest, valid)
	}

	user := data["username"].(string)
	chat := data["chat"].(string)

	if err := h.Store.AddChatToClient(ctx, user, chat); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusNoContent, "joined")
}

func (h *Handler) publish(c echo.Context) error {
	valid, data := h.Validate.PublishValidate(c)
	if valid.Encode() != "" {
		return c.JSON(http.StatusBadRequest, valid)
	}

	item := &store.Message{
		Client:  data["sender"].(string),
		Message: data["message"].(string),
	}
	ctx, endCtx := context.WithTimeout(context.Background(), 10*time.Second)
	defer endCtx()

	err := h.Store.InsertMessage(ctx, item)
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

	res, err := h.Store.GetChatMessages(ctx, user)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}

func (h Handler) Set(app *echo.Group) {
	app.POST("/login", h.login)

	app.POST("/join", middleware.Authenticate(h.join))
	app.POST("/publish", middleware.Authenticate(h.publish))
	app.GET("/suppress", middleware.Authenticate(h.suppress))
}
