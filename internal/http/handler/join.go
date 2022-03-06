package handler

import (
	"context"
	"github.com/amirhnajafiz/Telegraph/internal/db/store"
	"github.com/amirhnajafiz/Telegraph/pkg/jwt"
	"github.com/amirhnajafiz/Telegraph/pkg/validate"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"net/http"
)

type Join struct {
	Database *mongo.Database
	Logger   *zap.Logger
	Validate validate.Validate
}

func (j Join) Handle(c echo.Context) error {
	valid, data := j.Validate.PublishValidate(c)

	if valid.Encode() != "" {
		return c.JSON(http.StatusBadRequest, valid)
	}

	user := data["username"].(string)
	pass := data["pass"].(string)
	if user == "" {
		c.Response().Status = http.StatusBadRequest
		return c.String(http.StatusBadRequest, "no username found")
	}

	ctx := context.Background()
	client, err := store.Client{}.Find(j.Database, ctx, user)
	if err != nil {
		return err
	}

	if client.Pass != "" && client.Pass != pass {
		c.Response().Status = http.StatusUnauthorized
		return c.String(http.StatusUnauthorized, "user and pass don't match")
	}

	if client.Pass == "" {
		_ = store.Client{}.Store(j.Database, ctx, &store.Client{
			Name: user,
			Pass: pass,
		})
	}

	token, err := jwt.GenerateToken(user)
	if err != nil {
		c.Response().Status = http.StatusInternalServerError
		return c.JSON(http.StatusInternalServerError, "token not build")
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}

func (j Join) Register(g *echo.Group) {
	g.POST("/join", j.Handle)
}
