package saldogrpcerrors

import (
	"net/http"

	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

var (
	ErrGrpcSaldoInvalidID         = errors.NewGrpcError("Invalid Saldo ID", http.StatusBadRequest)
	ErrGrpcSaldoInvalidCardNumber = errors.NewGrpcError("Invalid Saldo Card Number", http.StatusBadRequest)
	ErrGrpcSaldoInvalidMonth      = errors.NewGrpcError("Invalid Saldo Month", http.StatusBadRequest)
	ErrGrpcSaldoInvalidYear       = errors.NewGrpcError("Invalid Saldo Year", http.StatusBadRequest)
)
