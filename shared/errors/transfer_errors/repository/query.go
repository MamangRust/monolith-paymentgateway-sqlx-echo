package transferrepositoryerrors

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

var (
	// ErrFindAllTransfersFailed indicates a failure when retrieving all transfer records.
	ErrFindAllTransfersFailed = errors.ErrInternal.WithMessage("Failed to find all transfers")

	// ErrFindActiveTransfersFailed indicates a failure when retrieving active (non-trashed) transfer records.
	ErrFindActiveTransfersFailed = errors.ErrInternal.WithMessage("Failed to find active transfers")

	// ErrFindTrashedTransfersFailed indicates a failure when retrieving trashed (soft-deleted) transfer records.
	ErrFindTrashedTransfersFailed = errors.ErrInternal.WithMessage("Failed to find trashed transfers")

	// ErrFindTransferByIdFailed indicates a failure when retrieving a transfer record by its ID.
	ErrFindTransferByIdFailed = errors.ErrInternal.WithMessage("Failed to find transfer by ID")

	// ErrFindTransferByTransferFromFailed indicates a failure when retrieving transfers by the sender (transfer from).
	ErrFindTransferByTransferFromFailed = errors.ErrInternal.WithMessage("Failed to find transfer by transfer from")

	// ErrFindTransferByTransferToFailed indicates a failure when retrieving transfers by the receiver (transfer to).
	ErrFindTransferByTransferToFailed = errors.ErrInternal.WithMessage("Failed to find transfer by transfer to")

	// ErrTransferNotFound is returned when a transfer is not found.
	ErrTransferNotFound = errors.ErrNotFound.WithMessage("Transfer not found")
)
