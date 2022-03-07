package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/amirhnajafiz/Telegraph/internal/db/store"
	"github.com/amirhnajafiz/Telegraph/pkg/jwt"
	"github.com/amirhnajafiz/Telegraph/pkg/validate"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type Join struct {
	Database *mongo.Database
	Logger   *zap.Logger
	Validate validate.Validate
}

func (j Join) Handle(c echo.Context) error {
	ctx, endCtx := context.WithTimeout(context.Background(), 10*time.Second)
	defer endCtx()

	valid, data := j.Validate.JoinValidate(c)
	if valid.Encode() != "" {
		return c.JSON(http.StatusBadRequest, valid)
	}

	user := data["username"].(string)
	pass := data["password"].(string)
	if user == "" {
		return c.String(http.StatusBadRequest, "must have a username")
	}

	client, err := store.Client{}.Find(j.Database, ctx, user)

	if err != mongo.ErrEmptySlice && client.Pass != pass {
		return c.String(http.StatusUnauthorized, "username and password mismatched")
	}

	if err == mongo.ErrEmptySlice {
		_ = store.Client{}.Store(j.Database, ctx, &store.Client{
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

func (j Join) Register(g *echo.Group) {
	g.POST("/join", j.Handle)
}
