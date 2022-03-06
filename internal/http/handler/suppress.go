package handler

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/amirhnajafiz/Telegraph/internal/db/store"
	"github.com/amirhnajafiz/Telegraph/pkg/jwt"
	"github.com/amirhnajafiz/Telegraph/pkg/validate"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type Suppress struct {
	Database *mongo.Database
	Logger   *zap.Logger
	Validate validate.Validate
}

func (s Suppress) Handle(c echo.Context) error {
	valid, _ := s.Validate.SuppressValidate(c)

	if valid.Encode() != "" {
		return c.JSON(http.StatusBadRequest, valid)
	}

	if auth, err := jwt.ParseToken(c.Request().Header.Get("jwt-token")); err != nil || !auth {
		return fmt.Errorf("unauthorized user")
	}

	user := c.FormValue("sender")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	res := store.Message{}.All(s.Database, ctx, user)

	return c.JSON(http.StatusOK, res)
}

func (s Suppress) Register(g *echo.Group) {
	g.GET("/suppress", s.Handle)
}
