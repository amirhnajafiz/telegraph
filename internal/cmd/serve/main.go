package serve

import (
	"Telegraph/internal/handler"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func GetServer(logger *zap.Logger, database *mongo.Database) *echo.Echo {
	logger.Info("Initialized server")

	e := echo.New()
	handler.Set(e, logger.Named("handler"), database)

	return e
}
