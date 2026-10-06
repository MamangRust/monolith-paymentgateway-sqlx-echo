package transferrepositoryerrors

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

var (
	// ErrCreateTransferFailed indicates a failure when creating a new transfer.
	ErrCreateTransferFailed = errors.ErrInternal.WithMessage("Failed to create transfer")

	// ErrUpdateTransferFailed indicates a failure when updating an existing transfer.
	ErrUpdateTransferFailed = errors.ErrInternal.WithMessage("Failed to update transfer")

	// ErrUpdateTransferAmountFailed indicates a failure when updating the amount of a transfer.
	ErrUpdateTransferAmountFailed = errors.ErrInternal.WithMessage("Failed to update transfer amount")

	// ErrUpdateTransferStatusFailed indicates a failure when updating the status of a transfer.
	ErrUpdateTransferStatusFailed = errors.ErrInternal.WithMessage("Failed to update transfer status")

	// ErrTrashedTransferFailed indicates a failure when soft-deleting (trashing) a transfer.
	ErrTrashedTransferFailed = errors.ErrInternal.WithMessage("Failed to move transfer to trash")

	// ErrRestoreTransferFailed indicates a failure when restoring a previously trashed transfer.
	ErrRestoreTransferFailed = errors.ErrInternal.WithMessage("Failed to restore transfer from trash")

	// ErrDeleteTransferPermanentFailed indicates a failure when permanently deleting a transfer.
	ErrDeleteTransferPermanentFailed = errors.ErrInternal.WithMessage("Failed to permanently delete transfer")

	// ErrRestoreAllTransfersFailed indicates a failure when restoring all trashed transfers.
	ErrRestoreAllTransfersFailed = errors.ErrInternal.WithMessage("Failed to restore all transfers")

	// ErrDeleteAllTransfersPermanentFailed indicates a failure when permanently deleting all transfers.
	ErrDeleteAllTransfersPermanentFailed = errors.ErrInternal.WithMessage("Failed to permanently delete all transfers")
)
