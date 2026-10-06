package merchantgrpcerrors

import (
	"net/http"

	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

var (
	ErrGrpcValidateCreateMerchant       = errors.NewGrpcError("Invalid input for create merchant", http.StatusBadRequest)
	ErrGrpcValidateUpdateMerchant       = errors.NewGrpcError("Invalid input for update merchant", http.StatusBadRequest)
	ErrGrpcValidateUpdateMerchantStatus = errors.NewGrpcError("Invalid input for update merchant status", http.StatusBadRequest)
)
