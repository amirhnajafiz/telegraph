package cmd

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Exec() {
	e := echo.New()
	e.GET("api/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":5000"))
}
