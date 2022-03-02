package subscribe

import (
	"net/http"

	"github.com/amirhnajafiz/Telegraph/internal/http/request"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type Subscribe struct {
	Logger *zap.Logger
}

func (s Subscribe) Handle(c echo.Context) error {
	valid, _ := request.SuppressValidate(c)

	if valid.Encode() != "" {
		return c.JSON(http.StatusBadRequest, valid)
	}

	id := c.FormValue("receiver")

	return c.JSON(http.StatusOK, id)
}

func (s Subscribe) Register(g *echo.Group) {
	g.GET("/", s.Handle)
}