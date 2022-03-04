package handler

import (
	"fmt"
	"github.com/amirhnajafiz/Telegraph/pkg/jwt"
	"net/http"

	"github.com/amirhnajafiz/Telegraph/pkg/validate"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type Subscribe struct {
	Logger   *zap.Logger
	Validate validate.Validate
}

func (s Subscribe) Handle(c echo.Context) error {
	valid, _ := s.Validate.SuppressValidate(c)

	if valid.Encode() != "" {
		return c.JSON(http.StatusBadRequest, valid)
	}

	if auth, err := jwt.ParseToken(c.Request().Header.Get("jwt-token")); err != nil || !auth {
		return fmt.Errorf("unauthorized user")
	}

	id := c.FormValue("receiver")

	return c.JSON(http.StatusOK, id)
}

func (s Subscribe) Register(g *echo.Group) {
	g.GET("/", s.Handle)
}
