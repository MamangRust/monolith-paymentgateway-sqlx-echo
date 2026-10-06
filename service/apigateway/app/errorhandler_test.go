package app

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/MamangRust/monolith-payment-gateway-apigateway/middlewares"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

func TestMiddlewareErrorsWriteJSON(t *testing.T) {
	viper.Set("SECRET_KEY", "test-secret")

	e := createEchoServer(&ClientConfig{}, nil, nil)
	e.GET("/api/test-protected", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"ok": "1"})
	})

	req := httptest.NewRequest(http.MethodGet, "/api/test-protected", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	t.Logf("status=%d body=%q", rec.Code, rec.Body.String())

	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("expected 401 for missing token, got %d with body %q", rec.Code, rec.Body.String())
	}
	if rec.Body.Len() == 0 {
		t.Fatal("expected a JSON error body, got empty response")
	}
}

// TestErrorIsolation builds the echo chain incrementally to find the middleware
// that swallows error responses. Every case registers a route handler that
// returns a 401 error, simulating any route-level middleware rejection.
func TestErrorIsolation(t *testing.T) {
	viper.Set("SECRET_KEY", "test-secret")

	errorHandler := func(c echo.Context) error {
		return echo.NewHTTPError(http.StatusUnauthorized, "no token")
	}

	newEcho := func() *echo.Echo {
		e := echo.New()
		e.HTTPErrorHandler = func(err error, c echo.Context) {
			if c.Response().Committed {
				t.Logf("HTTPErrorHandler skipped: response already committed (status=%d)", c.Response().Status)
				return
			}
			code := http.StatusInternalServerError
			message := "Internal server error"
			if httpErr, ok := err.(*echo.HTTPError); ok {
				code = httpErr.Code
				if text, ok := httpErr.Message.(string); ok && text != "" {
					message = text
				}
			}
			_ = c.JSON(code, map[string]interface{}{
				"status":  "error",
				"message": message,
				"code":    code,
			})
		}
		return e
	}

	cases := []struct {
		name  string
		setup func(e *echo.Echo)
	}{
		{"baseline", func(e *echo.Echo) {}},
		{"recover", func(e *echo.Echo) { e.Use(middleware.Recover()) }},
		{"requestid", func(e *echo.Echo) { e.Use(middleware.RequestID()) }},
		{"guard", func(e *echo.Echo) {
			e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
				return func(c echo.Context) error {
					err := next(c)
					if err == nil && c.Path() == "" && !c.Response().Committed {
						return echo.NewHTTPError(http.StatusNotFound)
					}
					return err
				}
			})
		}},
		{"logger", func(e *echo.Echo) { e.Use(createLoggerMiddleware()) }},
		{"cors", func(e *echo.Echo) { e.Use(middleware.CORS()) }},
		{"gzip", func(e *echo.Echo) { e.Use(middleware.Gzip()) }},
		{"secure", func(e *echo.Echo) { e.Use(middleware.Secure()) }},
		{"recover+gzip+cors+secure", func(e *echo.Echo) {
			e.Use(middleware.Recover())
			e.Use(middleware.RequestID())
			e.Use(middleware.Gzip())
			e.Use(middleware.CORS())
			e.Use(middleware.Secure())
		}},
		{"pyroscope", func(e *echo.Echo) { e.Use(middlewares.PyroscopeMiddleware()) }},
		{"echojwt", func(e *echo.Echo) { middlewares.WebSecurityConfig(e) }},
		{"all-with-echojwt", func(e *echo.Echo) {
			e.Use(middleware.Recover())
			e.Use(middleware.RequestID())
			e.Use(middleware.Gzip())
			e.Use(middleware.CORS())
			e.Use(middleware.Secure())
			e.Use(middlewares.PyroscopeMiddleware())
			middlewares.WebSecurityConfig(e)
		}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			e := newEcho()
			tc.setup(e)
			e.GET("/api/test-protected", errorHandler)

			req := httptest.NewRequest(http.MethodGet, "/api/test-protected", nil)
			rec := httptest.NewRecorder()
			e.ServeHTTP(rec, req)

			t.Logf("status=%d body=%q", rec.Code, rec.Body.String())
			if rec.Code != http.StatusUnauthorized || rec.Body.Len() == 0 {
				t.Errorf("expected 401 with body, got status=%d body=%q", rec.Code, rec.Body.String())
			}
		})
	}
}
