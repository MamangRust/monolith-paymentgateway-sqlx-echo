package autherrors

import (
	"net/http"

	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

var (
	// ErrGrpcLogin is returned when login fails due to invalid credentials or arguments.
	ErrGrpcLogin = errors.NewGrpcError("Login failed: invalid credentials or arguments provided", http.StatusBadRequest)

	// ErrGrpcGetMe is returned when fetching the current user info fails due to lack of authentication.
	ErrGrpcGetMe = errors.NewGrpcError("Failed to fetch user info: unauthenticated", http.StatusUnauthorized)

	// ErrGrpcRegisterToken is returned when registration fails due to invalid arguments.
	ErrGrpcRegisterToken = errors.NewGrpcError("Registration failed: invalid arguments provided", http.StatusBadRequest)
)
