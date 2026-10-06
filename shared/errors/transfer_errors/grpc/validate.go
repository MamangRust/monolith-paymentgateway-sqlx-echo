package transfergrpcerrors

import (
	"net/http"

	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

var (
	ErrGrpcValidateCreateTransferRequest = errors.NewGrpcError("Invalid input for create transfer", http.StatusBadRequest)
	ErrGrpcValidateUpdateTransferRequest = errors.NewGrpcError("Invalid input for update transfer", http.StatusBadRequest)
)
