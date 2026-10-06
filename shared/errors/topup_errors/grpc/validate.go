package topupgrpcerrors

import (
	"net/http"

	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

var (
	ErrGrpcValidateCreateTopup = errors.NewGrpcError("Invalid input for create topup", http.StatusBadRequest)
	ErrGrpcValidateUpdateTopup = errors.NewGrpcError("Invalid input for update topup", http.StatusBadRequest)
)
