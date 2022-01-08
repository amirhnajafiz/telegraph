package suppress

import (
	"context"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"time"
)

type Suppress struct {
	Database *mongo.Database
	Logger   *zap.Logger
}

func (s Suppress) Handle(c echo.Context) error {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	return nil
}
