package saldorepositoryerrors

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

var (
	// ErrGetMonthlySaldoBalancesFailed is returned when fetching monthly saldo balances fails.
	ErrGetMonthlySaldoBalancesFailed = errors.ErrInternal.WithMessage("Failed to get monthly saldo balances")

	// ErrGetYearlySaldoBalancesFailed is returned when fetching yearly saldo balances fails.
	ErrGetYearlySaldoBalancesFailed = errors.ErrInternal.WithMessage("Failed to get yearly saldo balances")
)
