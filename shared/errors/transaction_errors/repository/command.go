package transactionrepositoryerrors

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

var (
	// ErrCreateTransactionFailed indicates a failure when creating a new transaction.
	ErrCreateTransactionFailed = errors.ErrInternal.WithMessage("Failed to create transaction")

	// ErrUpdateTransactionFailed indicates a failure when updating a transaction.
	ErrUpdateTransactionFailed = errors.ErrInternal.WithMessage("Failed to update transaction")

	// ErrUpdateTransactionStatusFailed indicates a failure when updating the status of a transaction.
	ErrUpdateTransactionStatusFailed = errors.ErrInternal.WithMessage("Failed to update transaction status")

	// ErrTrashedTransactionFailed indicates a failure when soft-deleting (trashing) a transaction.
	ErrTrashedTransactionFailed = errors.ErrInternal.WithMessage("Failed to move transaction to trash")

	// ErrRestoreTransactionFailed indicates a failure when restoring a trashed transaction.
	ErrRestoreTransactionFailed = errors.ErrInternal.WithMessage("Failed to restore transaction from trash")

	// ErrDeleteTransactionPermanentFailed indicates a failure when permanently deleting a transaction.
	ErrDeleteTransactionPermanentFailed = errors.ErrInternal.WithMessage("Failed to permanently delete transaction")

	// ErrRestoreAllTransactionsFailed indicates a failure when restoring all trashed transactions.
	ErrRestoreAllTransactionsFailed = errors.ErrInternal.WithMessage("Failed to restore all transactions")

	// ErrDeleteAllTransactionsPermanentFailed indicates a failure when permanently deleting all transactions.
	ErrDeleteAllTransactionsPermanentFailed = errors.ErrInternal.WithMessage("Failed to permanently delete all transactions")
)
