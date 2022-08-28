package serve

import (
	"github.com/amirhnajafiz/telegraph/internal/http/handler"
	"github.com/amirhnajafiz/telegraph/internal/nats"
	"github.com/amirhnajafiz/telegraph/internal/validate"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type Serve struct {
	Database *mongo.Database
	Logger   *zap.Logger
	Nats     nats.Nats
}

func (s Serve) GetServer() *echo.Echo {
	s.Logger.Info("Initialized server")

	e := echo.New()
	handler.Handler{
		Database: s.Database,
		Logger:   s.Logger.Named("handler"),
		Nats:     s.Nats,
		Validate: validate.Validate{},
	}.Set(e)

	return e
}
