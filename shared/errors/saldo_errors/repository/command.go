package saldorepositoryerrors

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

var (
	// ErrCreateSaldoFailed is returned when creating a saldo record fails.
	ErrCreateSaldoFailed = errors.ErrInternal.WithMessage("Failed to create saldo record")

	// ErrUpdateSaldoFailed is returned when updating a saldo record fails.
	ErrUpdateSaldoFailed = errors.ErrInternal.WithMessage("Failed to update saldo record")

	// ErrUpdateSaldoBalanceFailed is returned when updating saldo balance fails.
	ErrUpdateSaldoBalanceFailed = errors.ErrInternal.WithMessage("Failed to update saldo balance")

	// ErrUpdateSaldoWithdrawFailed is returned when updating saldo for a withdrawal fails.
	ErrUpdateSaldoWithdrawFailed = errors.ErrInternal.WithMessage("Failed to update saldo withdrawal")

	// ErrTrashSaldoFailed is returned when soft-deleting (trashing) a saldo record fails.
	ErrTrashSaldoFailed = errors.ErrInternal.WithMessage("Failed to move saldo record to trash")

	// ErrRestoreSaldoFailed is returned when restoring a trashed saldo record fails.
	ErrRestoreSaldoFailed = errors.ErrInternal.WithMessage("Failed to restore saldo record from trash")

	// ErrDeleteSaldoPermanentFailed is returned when permanently deleting a saldo record fails.
	ErrDeleteSaldoPermanentFailed = errors.ErrInternal.WithMessage("Failed to permanently delete saldo record")

	// ErrRestoreAllSaldosFailed is returned when restoring all trashed saldo records fails.
	ErrRestoreAllSaldosFailed = errors.ErrInternal.WithMessage("Failed to restore all saldo records")

	// ErrDeleteAllSaldosPermanentFailed is returned when permanently deleting all saldo records fails.
	ErrDeleteAllSaldosPermanentFailed = errors.ErrInternal.WithMessage("Failed to permanently delete all saldo records")
)
