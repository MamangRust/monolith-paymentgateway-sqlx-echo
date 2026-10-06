package auth

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// ErrTokenExpired is an error that is returned when a JWT token is expired.
var ErrTokenExpired = errors.New("token expired")

//go:generate mockgen -source=token.go -destination=mocks/token.go
type TokenManager interface {
	GenerateToken(userId int, audience string) (string, error)
	ValidateToken(tokenString string) (string, error)
}

type Manager struct {
	secretKey []byte
}

// NewManager creates a new Manager instance with the given secret key.
//
// The secret key is expected to be a non-empty string. If the secret key is
// empty, an error is returned.
func NewManager(secretKey string) (*Manager, error) {
	if secretKey == "" {
		return nil, errors.New("empty secret key")
	}
	return &Manager{secretKey: []byte(secretKey)}, nil
}

// GenerateToken generates a new JWT token for the given user ID and audience.
//
// The token is valid for 12 hours from the time it is generated.
// The subject claim is set to the given user ID.
// The audience claim is set to the given audience.
// The token is signed with the secret key set on the Manager during initialization.
//
// If the token cannot be generated, an error is returned.
func (m *Manager) GenerateToken(userId int, audience string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(12 * time.Hour)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(expireTime),
		Subject:   strconv.Itoa(userId),
		Audience:  []string{audience},
	})

	return token.SignedString([]byte(m.secretKey))
}

// ValidateToken validates a JWT token and returns the user ID string if the validation is successful.
// If the token is invalid or expired, an error is returned.
// The error is wrapped with jwt.ErrTokenExpired if the token is expired.
// The error is wrapped with jwt.ErrTokenExpired if the token is expired.
func (m *Manager) ValidateToken(accessToken string) (string, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (i interface{}, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(m.secretKey), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return "", ErrTokenExpired
		}
		return "", fmt.Errorf("failed to parse token: %w", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("error get user claims from token")
	}

	return claims["sub"].(string), nil
}
