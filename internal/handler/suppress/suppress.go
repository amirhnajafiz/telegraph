package suppress

import (
	"Telegraph/internal/store/message"
	"context"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
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
	col := s.Database.Collection(message.Collection)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	cursor, _ := col.Find(ctx, bson.D{})
	var results []message.Message

	cursor.All(ctx, results)

	return c.JSON(http.StatusOK, results)
}
