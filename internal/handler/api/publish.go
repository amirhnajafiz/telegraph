package api

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"net/http"
)

type Publish struct {
	Logger *zap.Logger
}

type Request struct {
	Source string `json:"from"`
	Des    string `json:"to"`
	Msg    string `json:"message"`
}

func (publish Publish) Handle(c echo.Context) (err error) {
	req := new(Request)

	if err = c.Bind(req); err != nil {
		return err
	}

	// TODO 0: Data validation
	// TODO 1: Save the message into database
	// TODO 2: Send the message to the destination
	// TODO 3: Notify the destination

	return c.JSON(http.StatusOK, req)
}

func (publish Publish) Register(g *echo.Group) {
	g.POST("/publish", publish.Handle)
}
