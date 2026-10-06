package transferserviceerrors

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

var (
	// ErrFailedCreateTransfer indicates a failure when attempting to create a new transfer record.
	ErrFailedCreateTransfer = errors.ErrInternal.WithMessage("Failed to create transfer")

	// ErrFailedUpdateTransfer indicates a failure when attempting to update an existing transfer record.
	ErrFailedUpdateTransfer = errors.ErrInternal.WithMessage("Failed to update transfer")

	// ErrFailedTrashedTransfer indicates a failure when attempting to soft-delete (trash) a transfer.
	ErrFailedTrashedTransfer = errors.ErrInternal.WithMessage("Failed to move transfer to trash")

	// ErrFailedRestoreTransfer indicates a failure when attempting to restore a previously trashed transfer.
	ErrFailedRestoreTransfer = errors.ErrInternal.WithMessage("Failed to restore transfer")

	// ErrFailedDeleteTransferPermanent indicates a failure when attempting to permanently delete a transfer.
	ErrFailedDeleteTransferPermanent = errors.ErrInternal.WithMessage("Failed to delete transfer permanently")

	// ErrFailedRestoreAllTransfers indicates a failure when attempting to restore all trashed transfers.
	ErrFailedRestoreAllTransfers = errors.ErrInternal.WithMessage("Failed to restore all transfers")

	// ErrFailedDeleteAllTransfersPermanent indicates a failure when attempting to permanently delete all transfer records.
	ErrFailedDeleteAllTransfersPermanent = errors.ErrInternal.WithMessage("Failed to delete all transfers permanently")
)
