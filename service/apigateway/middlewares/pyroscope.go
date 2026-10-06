package middlewares

import (
	"context"

	"github.com/grafana/pyroscope-go"
	"github.com/labstack/echo/v4"
)

func PyroscopeMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			var nextErr error
			pyroscope.TagWrapper(c.Request().Context(), pyroscope.Labels(
				"endpoint", c.Path(),
				"method", c.Request().Method,
			), func(ctx context.Context) {
				c.SetRequest(c.Request().WithContext(ctx))
				// The error must be propagated: swallowing it turned every
				// middleware rejection (JWT 401, API key timeout, ...) into an
				// empty 200 response.
				nextErr = next(c)
			})
			return nextErr
		}
	}
}
