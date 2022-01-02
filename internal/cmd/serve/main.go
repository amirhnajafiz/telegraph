package serve

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetServer() *echo.Echo {
	e := echo.New()
	e.GET("api/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	return e
}
