package merchantgrpcerrors

import (
	"net/http"

	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

var (
	ErrGrpcMerchantInvalidID     = errors.NewGrpcError("Invalid Merchant ID", http.StatusBadRequest)
	ErrGrpcMerchantInvalidUserID = errors.NewGrpcError("Invalid Merchant User ID", http.StatusBadRequest)
	ErrGrpcMerchantInvalidApiKey = errors.NewGrpcError("Invalid Merchant Api Key", http.StatusBadRequest)
	ErrGrpcMerchantInvalidMonth  = errors.NewGrpcError("Invalid Merchant Month", http.StatusBadRequest)
	ErrGrpcMerchantInvalidYear   = errors.NewGrpcError("Invalid Merchant Year", http.StatusBadRequest)
)
