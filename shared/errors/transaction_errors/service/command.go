package transactonserviceerrors

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

var (
	// ErrFailedCreateTransaction indicates a failure when creating a new transaction record.
	ErrFailedCreateTransaction = errors.ErrInternal.WithMessage("Failed to create transaction")

	// ErrFailedUpdateTransaction indicates a failure when updating an existing transaction record.
	ErrFailedUpdateTransaction = errors.ErrInternal.WithMessage("Failed to update transaction")

	// ErrFailedTrashedTransaction indicates a failure when soft-deleting (trashing) a transaction.
	ErrFailedTrashedTransaction = errors.ErrInternal.WithMessage("Failed to move transaction to trash")

	// ErrFailedRestoreTransaction indicates a failure when restoring a previously trashed transaction.
	ErrFailedRestoreTransaction = errors.ErrInternal.WithMessage("Failed to restore transaction")

	// ErrFailedDeleteTransactionPermanent indicates a failure when permanently deleting a transaction.
	ErrFailedDeleteTransactionPermanent = errors.ErrInternal.WithMessage("Failed to delete transaction permanently")

	// ErrFailedRestoreAllTransactions indicates a failure when restoring all trashed transactions.
	ErrFailedRestoreAllTransactions = errors.ErrInternal.WithMessage("Failed to restore all transactions")

	// ErrFailedDeleteAllTransactionsPermanent indicates a failure when permanently deleting all transactions.
	ErrFailedDeleteAllTransactionsPermanent = errors.ErrInternal.WithMessage("Failed to delete all transactions permanently")

	// ErrFailedAuthorizeTransaction is returned when authorization fails.
	ErrFailedAuthorizeTransaction = errors.ErrInternal.WithMessage("Failed to authorize transaction")

	// ErrFailedCaptureTransaction is returned when capture fails.
	ErrFailedCaptureTransaction = errors.ErrInternal.WithMessage("Failed to capture transaction")

	// ErrFailedVoidTransaction is returned when void fails.
	ErrFailedVoidTransaction = errors.ErrInternal.WithMessage("Failed to void transaction")

	// ErrFailedRefundTransaction is returned when refund fails.
	ErrFailedRefundTransaction = errors.ErrInternal.WithMessage("Failed to refund transaction")
)
