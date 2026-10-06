package merchantdocumentgrpcerrors

import (
	"net/http"

	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

var (
	ErrGrpcValidateCreateMerchantDocument = errors.NewGrpcError("Invalid input for create merchant document", http.StatusBadRequest)
	ErrGrpcValidateUpdateMerchantDocument = errors.NewGrpcError("Invalid input for update merchant document", http.StatusBadRequest)
)
