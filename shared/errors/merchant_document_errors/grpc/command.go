package merchantdocumentgrpcerrors

import (
	"net/http"

	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

var (
	ErrGrpcFailedCreateMerchantDocument = errors.NewGrpcError("Failed to create merchant document", http.StatusInternalServerError)
	ErrGrpcFailedUpdateMerchantDocument = errors.NewGrpcError("Failed to update merchant document", http.StatusInternalServerError)
)
