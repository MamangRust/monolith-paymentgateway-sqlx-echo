package cardgrpcerrors

import (
	"net/http"

	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

var (
	ErrGrpcInvalidCardID     = errors.NewGrpcError("Invalid card ID", http.StatusBadRequest)
	ErrGrpcInvalidUserID     = errors.NewGrpcError("Invalid user ID", http.StatusBadRequest)
	ErrGrpcInvalidCardNumber = errors.NewGrpcError("Invalid card number", http.StatusBadRequest)
	ErrGrpcInvalidMonth      = errors.NewGrpcError("Invalid month", http.StatusBadRequest)
	ErrGrpcInvalidYear       = errors.NewGrpcError("Invalid year", http.StatusBadRequest)
)
