package topupgrpcerrors

import (
	"net/http"

	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

var (
	ErrGrpcTopupInvalidID    = errors.NewGrpcError("Invalid Topup ID", http.StatusBadRequest)
	ErrGrpcTopupInvalidMonth = errors.NewGrpcError("Invalid Topup Month", http.StatusBadRequest)
	ErrGrpcInvalidCardNumber = errors.NewGrpcError("Invalid card number", http.StatusBadRequest)
	ErrGrpcTopupInvalidYear  = errors.NewGrpcError("Invalid Topup Year", http.StatusBadRequest)
)
