package usergrpcerrors

import (
	"net/http"

	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

// ErrGrpcUserInvalidId is returned when an invalid user ID is provided.
var ErrGrpcUserInvalidId = errors.NewGrpcError("Invalid user ID", http.StatusBadRequest)

// ErrGrpcUserInvalidEmail is returned when an empty or malformed email is provided.
var ErrGrpcUserInvalidEmail = errors.NewGrpcError("Invalid email", http.StatusBadRequest)

// ErrGrpcUserInvalidVerificationCode is returned when an empty verification code is provided.
var ErrGrpcUserInvalidVerificationCode = errors.NewGrpcError("Invalid verification code", http.StatusBadRequest)
