package saldorepositoryerrors

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

var (
	// ErrSaldoNotFound is returned when the requested saldo record is not found in the database.
	ErrSaldoNotFound = errors.ErrNotFound.WithMessage("Saldo not found")

	// ErrFindAllSaldosFailed is returned when fetching all saldo records fails.
	ErrFindAllSaldosFailed = errors.ErrInternal.WithMessage("Failed to find all saldo records")

	// ErrFindActiveSaldosFailed is returned when fetching active saldo records fails.
	ErrFindActiveSaldosFailed = errors.ErrInternal.WithMessage("Failed to find active saldo records")

	// ErrFindTrashedSaldosFailed is returned when fetching trashed saldo records fails.
	ErrFindTrashedSaldosFailed = errors.ErrInternal.WithMessage("Failed to find trashed saldo records")

	// ErrFindSaldoByIdFailed is returned when fetching a saldo by its ID fails.
	ErrFindSaldoByIdFailed = errors.ErrInternal.WithMessage("Failed to find saldo by ID")

	// ErrFindSaldoByCardNumberFailed is returned when fetching saldo by card number fails.
	ErrFindSaldoByCardNumberFailed = errors.ErrInternal.WithMessage("Failed to find saldo by card number")
)
