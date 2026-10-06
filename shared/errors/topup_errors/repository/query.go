package topuprepositoryerrors

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

var (
	// ErrFindAllTopupsFailed indicates failure in retrieving all top-up records from the database.
	ErrFindAllTopupsFailed = errors.ErrInternal.WithMessage("Failed to find all topups")

	// ErrFindTopupsByActiveFailed indicates failure in retrieving only the active (non-deleted) top-ups.
	ErrFindTopupsByActiveFailed = errors.ErrInternal.WithMessage("Failed to find active topups")

	// ErrFindTopupsByTrashedFailed indicates failure in retrieving trashed (soft-deleted) top-ups.
	ErrFindTopupsByTrashedFailed = errors.ErrInternal.WithMessage("Failed to find trashed topups")

	// ErrFindTopupsByCardNumberFailed indicates failure in finding top-ups by a specific card number.
	ErrFindTopupsByCardNumberFailed = errors.ErrInternal.WithMessage("Failed to find topups by card number")

	// ErrFindTopupByIdFailed indicates failure in finding a top-up using its unique ID.
	ErrFindTopupByIdFailed = errors.ErrInternal.WithMessage("Failed to find topup by ID")

	// ErrTopupNotFound is returned when a topup is not found.
	ErrTopupNotFound = errors.ErrNotFound.WithMessage("Topup not found")
)
