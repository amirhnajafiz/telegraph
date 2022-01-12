package serve

import (
	"Telegraph/internal/http/handler"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type Tools struct {
	Database *mongo.Database
	Logger   *zap.Logger
}

func GetServer(t Tools) *echo.Echo {
	t.Logger.Info("Initialized server")

	e := echo.New()
	handler.Set(e, t.Logger.Named("handler"), t.Database)

	return e
}
