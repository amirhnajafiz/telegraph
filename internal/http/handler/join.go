package handler

import (
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
	user := c.QueryParam("username")
	if user == "" {
		c.Response().Status = http.StatusBadRequest
		return c.String(http.StatusBadRequest, "no username found")
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
	g.GET("/join", j.Handle)
}
