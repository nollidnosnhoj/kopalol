package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func ApiKeyAuthMiddleware() echo.MiddlewareFunc {
	return middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
		KeyLookup: "header:X-API-KEY",
		Validator: func(key string, c echo.Context) (bool, error) {
			// TODO: implement a real API key validation
			if key == "" {
				return false, nil
			}
			return true, nil
		},
	})
}
