package middleware

import (
	"github.com/amirhnajafiz/telegraph/internal/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Authenticate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Request().Header.Get("jwt-token") == "" {
			return c.String(http.StatusBadRequest, "no jwt-token in request header")
		}

		if auth, err := jwt.ParseToken(c.Request().Header.Get("jwt-token")); err != nil || !auth {
			return c.String(http.StatusUnauthorized, "not an authenticate user")
		}

		return next(c)
	}
}
