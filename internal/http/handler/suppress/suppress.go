package suppress

import (
	"context"
	"net/http"
	"time"

	"github.com/amirhnajafiz/Telegraph/internal/store/message"
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

	user := c.FormValue("sender")

	ctx, endCtx := context.WithTimeout(context.Background(), 10*time.Second)
	defer endCtx()

	res := message.All(s.Database, ctx, user)

	return c.JSON(http.StatusOK, res)
}

func (s Suppress) Register(g *echo.Group) {
	g.GET("/suppress", s.Handle)
}
