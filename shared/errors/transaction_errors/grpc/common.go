package transactiongrpcerrors

import (
	"net/http"

	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

var (
	ErrGrpcTransactionInvalidID         = errors.NewGrpcError("Invalid Transaction ID", http.StatusBadRequest)
	ErrGrpcTransactionInvalidMerchantID = errors.NewGrpcError("Invalid Transaction Merchant ID", http.StatusBadRequest)
	ErrGrpcInvalidCardNumber            = errors.NewGrpcError("Invalid card number", http.StatusBadRequest)
	ErrGrpcInvalidMonth                 = errors.NewGrpcError("Invalid month", http.StatusBadRequest)
	ErrGrpcInvalidYear                  = errors.NewGrpcError("Invalid year", http.StatusBadRequest)
)
