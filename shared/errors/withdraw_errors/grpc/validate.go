package withdrawgrpcerrors

import (
	"net/http"

	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

var (
	ErrGrpcValidateCreateWithdrawRequest = errors.NewGrpcError("Invalid input for create withdraw", http.StatusBadRequest)
	ErrGrpcValidateUpdateWithdrawRequest = errors.NewGrpcError("Invalid input for update withdraw", http.StatusBadRequest)
)
