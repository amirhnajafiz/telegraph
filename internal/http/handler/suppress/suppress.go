package suppress

import (
	"Telegraph/internal/http/request"
	"Telegraph/internal/store/message"
	"context"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type Suppress struct {
	Database *mongo.Database
	Logger   *zap.Logger
}

func (s Suppress) Handle(c echo.Context) error {
	valid, _ := request.SuppressValidate(c)

	if valid.Encode() != "" {
		return c.JSON(http.StatusBadRequest, valid)
	}

	user := c.FormValue("sender")

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	res := message.All(s.Database, ctx, user)
	return c.JSON(http.StatusOK, res)
}

func (s Suppress) Register(g *echo.Group) {
	g.GET("/suppress", s.Handle)
}
