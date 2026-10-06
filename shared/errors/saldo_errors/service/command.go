package saldoserviceerror

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

var (
	// ErrFailedCreateSaldo returns a 500 error when creating a saldo record fails.
	ErrFailedCreateSaldo = errors.ErrInternal.WithMessage("Failed to create saldo")

	// ErrFailedUpdateSaldo returns a 500 error when updating a saldo record fails.
	ErrFailedUpdateSaldo = errors.ErrInternal.WithMessage("Failed to update saldo")

	// ErrFailedTrashSaldo returns a 500 error when moving a saldo record to trash fails.
	ErrFailedTrashSaldo = errors.ErrInternal.WithMessage("Failed to move saldo to trash")

	// ErrFailedRestoreSaldo returns a 500 error when restoring a trashed saldo record fails.
	ErrFailedRestoreSaldo = errors.ErrInternal.WithMessage("Failed to restore saldo")

	// ErrFailedDeleteSaldoPermanent returns a 500 error when permanently deleting a saldo record fails.
	ErrFailedDeleteSaldoPermanent = errors.ErrInternal.WithMessage("Failed to delete saldo permanently")

	// ErrFailedRestoreAllSaldo returns a 500 error when restoring all trashed saldo records fails.
	ErrFailedRestoreAllSaldo = errors.ErrInternal.WithMessage("Failed to restore all saldos")

	// ErrFailedDeleteAllSaldoPermanent returns a 500 error when permanently deleting all trashed saldo records fails.
	ErrFailedDeleteAllSaldoPermanent = errors.ErrInternal.WithMessage("Failed to delete all saldos permanently")
)
