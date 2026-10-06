package saldogrpcerrors

import (
	"net/http"

	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

var (
	ErrGrpcValidateCreateSaldo         = errors.NewGrpcError("Invalid input for create saldo", http.StatusBadRequest)
	ErrGrpcValidateUpdateSaldo         = errors.NewGrpcError("Invalid input for update saldo", http.StatusBadRequest)
	ErrGrpcValidateUpdateSaldoWithdraw = errors.NewGrpcError("Invalid input for update saldo withdraw", http.StatusBadRequest)
)
