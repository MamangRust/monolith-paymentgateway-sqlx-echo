package transferserviceerrors

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

var (
	// ErrFailedFindAllTransfers indicates a failure when retrieving all transfer records.
	ErrFailedFindAllTransfers = errors.ErrInternal.WithMessage("Failed to fetch all transfers")

	// ErrTransferNotFound indicates that a specific transfer record was not found.
	ErrTransferNotFound = errors.ErrNotFound.WithMessage("Transfer not found")

	// ErrFailedFindActiveTransfers indicates a failure when retrieving active transfer records.
	ErrFailedFindActiveTransfers = errors.ErrInternal.WithMessage("Failed to fetch active transfers")

	// ErrFailedFindTrashedTransfers indicates a failure when retrieving trashed (soft-deleted) transfer records.
	ErrFailedFindTrashedTransfers = errors.ErrInternal.WithMessage("Failed to fetch trashed transfers")

	// ErrFailedFindTransfersBySender indicates a failure when retrieving transfers filtered by sender card.
	ErrFailedFindTransfersBySender = errors.ErrInternal.WithMessage("Failed to fetch transfers by sender")

	// ErrFailedFindTransfersByReceiver indicates a failure when retrieving transfers filtered by receiver card.
	ErrFailedFindTransfersByReceiver = errors.ErrInternal.WithMessage("Failed to fetch transfers by receiver")
)
