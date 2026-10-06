package cardgrpcerrors

import (
	"net/http"

	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

var (
	ErrGrpcValidateCreateCardRequest = errors.NewGrpcError("Invalid input for create card", http.StatusBadRequest)
	ErrGrpcValidateUpdateCardRequest = errors.NewGrpcError("Invalid input for update card", http.StatusBadRequest)
	ErrGrpcValidateToggleCardStatus  = errors.NewGrpcError("Invalid input for toggle card status", http.StatusBadRequest)
	ErrGrpcValidateUpdateCreditLimit = errors.NewGrpcError("Invalid input for update credit limit", http.StatusBadRequest)
	ErrGrpcValidateRedeemPoints      = errors.NewGrpcError("Invalid input for redeem points", http.StatusBadRequest)
)
