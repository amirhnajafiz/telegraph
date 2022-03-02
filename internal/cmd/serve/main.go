package serve

import (
	"github.com/amirhnajafiz/Telegraph/internal/http/handler"
	"github.com/amirhnajafiz/Telegraph/pkg/validate"
	"github.com/labstack/echo/v4"
)

func (s Serve) GetServer() *echo.Echo {
	s.Logger.Info("Initialized server")

	e := echo.New()
	handler.Handler{
		Database: s.Database,
		Logger:   s.Logger.Named("handler"),
		Validate: validate.Validate{},
	}.Set(e)

	return e
}
