package middlewares

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

// testSecretKey must match the value read by WebSecurityConfig via viper so a
// token signed with it is accepted.
const testSecretKey = "test-secret-key-for-401-suite"

func setupAuthEcho() *echo.Echo {
	viper.Set("SECRET_KEY", testSecretKey)
	e := echo.New()
	WebSecurityConfig(e)
	e.GET("/api/topup-query/example", func(c echo.Context) error {
		return c.String(http.StatusOK, "protected")
	})
	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "healthy")
	})
	return e
}

func signToken(t *testing.T, secret string) string {
	t.Helper()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"sub": "1"})
	signed, err := token.SignedString([]byte(secret))
	if err != nil {
		t.Fatalf("failed to sign token: %v", err)
	}
	return signed
}

func TestAuthMiddlewareRejectsMissingToken(t *testing.T) {
	e := setupAuthEcho()
	req := httptest.NewRequest(http.MethodGet, "/api/topup-query/example", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("expected 401 for missing token, got %d", rec.Code)
	}
}

func TestAuthMiddlewareRejectsMalformedToken(t *testing.T) {
	e := setupAuthEcho()
	req := httptest.NewRequest(http.MethodGet, "/api/topup-query/example", nil)
	req.Header.Set(echo.HeaderAuthorization, "Bearer not-a-jwt")
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("expected 401 for malformed token, got %d", rec.Code)
	}
}

func TestAuthMiddlewareRejectsExpiredToken(t *testing.T) {
	e := setupAuthEcho()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": "1",
		"exp": time.Now().Add(-time.Hour).Unix(),
	})
	signed, err := token.SignedString([]byte(testSecretKey))
	if err != nil {
		t.Fatalf("failed to sign token: %v", err)
	}
	req := httptest.NewRequest(http.MethodGet, "/api/topup-query/example", nil)
	req.Header.Set(echo.HeaderAuthorization, "Bearer "+signed)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("expected 401 for expired token, got %d", rec.Code)
	}
}

func TestAuthMiddlewareRejectsTokenSignedWithWrongKey(t *testing.T) {
	e := setupAuthEcho()
	req := httptest.NewRequest(http.MethodGet, "/api/topup-query/example", nil)
	req.Header.Set(echo.HeaderAuthorization, "Bearer "+signToken(t, "wrong-secret-key"))
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("expected 401 for invalid signature, got %d", rec.Code)
	}
}

func TestAuthMiddlewareAllowsValidToken(t *testing.T) {
	e := setupAuthEcho()
	req := httptest.NewRequest(http.MethodGet, "/api/topup-query/example", nil)
	req.Header.Set(echo.HeaderAuthorization, "Bearer "+signToken(t, testSecretKey))
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200 for valid token, got %d", rec.Code)
	}
}

func TestAuthMiddlewareSkipsWhitelistedPathWithoutToken(t *testing.T) {
	e := setupAuthEcho()
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200 for whitelisted path without token, got %d", rec.Code)
	}
}
