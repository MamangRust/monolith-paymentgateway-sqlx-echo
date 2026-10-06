package transactiongrpcerrors

import (
	"net/http"

	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

var (
	ErrGrpcValidateCreateTransactionRequest = errors.NewGrpcError("Invalid input for create card", http.StatusBadRequest)
	ErrGrpcValidateUpdateTransactionRequest = errors.NewGrpcError("Invalid input for update card", http.StatusBadRequest)

	// ErrGrpcInvalidApiKey is returned when api key is invalid.
	ErrGrpcInvalidApiKey = errors.NewGrpcError("Invalid API key", http.StatusBadRequest)
)
