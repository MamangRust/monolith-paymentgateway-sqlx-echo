package middlewares

import (
	"strings"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

var whiteListPaths = []string{
	"/api/auth/login",
	"/api/auth/register", "/api/auth/hello",
	"/api/auth/verify-code",
	"/api/auth/refresh-token",
	"/api/auth/forgot-password",
	"/api/auth/reset-password",
	"/docs/",
	"/docs",
	"/swagger",
	"/metrics",
	"/health",
	"/live",
	"/ready",
}

func WebSecurityConfig(e *echo.Echo) {
	config := echojwt.Config{
		SigningKey: []byte(viper.GetString("SECRET_KEY")),
		Skipper:    skipAuth,
		SuccessHandler: func(c echo.Context) {
			user := c.Get("user").(*jwt.Token)

			if claims, ok := user.Claims.(jwt.MapClaims); ok {
				subject := claims["sub"]
				// Keep both keys during the transition: existing handlers use
				// userId while role-protected handlers use user_id.
				c.Set("userId", subject)
				c.Set("user_id", subject)
			}
		},
		ErrorHandler: func(c echo.Context, err error) error {
			return echo.ErrUnauthorized
		},
	}
	e.Use(echojwt.WithConfig(config))

}

func skipAuth(c echo.Context) bool {
	path := c.Request().URL.Path

	for _, p := range whiteListPaths {
		if path == p ||
			strings.HasPrefix(path, "/swagger") ||
			strings.HasPrefix(path, "/metrics") {
			return true
		}
	}

	return false
}
