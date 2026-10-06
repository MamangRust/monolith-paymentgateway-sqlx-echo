package saldoserviceerror

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

var (
	// ErrFailedFindAllSaldos returns a 500 error when fetching all saldo records fails.
	ErrFailedFindAllSaldos = errors.ErrInternal.WithMessage("Failed to fetch saldos")

	// ErrFailedSaldoNotFound returns a 404 error when a requested saldo record is not found.
	ErrFailedSaldoNotFound = errors.ErrNotFound.WithMessage("Saldo not found")

	// ErrFailedFindSaldoByCardNumber returns a 500 error when fetching saldo by card number fails.
	ErrFailedFindSaldoByCardNumber = errors.ErrInternal.WithMessage("Failed to find saldo by card number")

	// ErrFailedFindActiveSaldos returns a 500 error when fetching active saldo records fails.
	ErrFailedFindActiveSaldos = errors.ErrInternal.WithMessage("Failed to fetch active saldos")

	// ErrFailedFindTrashedSaldos returns a 500 error when fetching trashed saldo records fails.
	ErrFailedFindTrashedSaldos = errors.ErrInternal.WithMessage("Failed to fetch trashed saldos")
)
