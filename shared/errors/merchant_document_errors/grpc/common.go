package merchantdocumentgrpcerrors

import (
	"net/http"

	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

var ErrGrpcMerchantInvalidID = errors.NewGrpcError("Invalid merchant id", http.StatusBadRequest)
