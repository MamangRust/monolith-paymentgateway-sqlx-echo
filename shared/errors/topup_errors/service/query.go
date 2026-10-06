package topupserviceerrors

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

var (
	// ErrTopupNotFoundRes indicates that a requested top-up was not found.
	ErrTopupNotFoundRes = errors.ErrNotFound.WithMessage("Topup not found")

	// ErrFailedFindAllTopups indicates failure in retrieving all top-up records.
	ErrFailedFindAllTopups = errors.ErrInternal.WithMessage("Failed to fetch topups")

	// ErrFailedFindAllTopupsByCardNumber indicates failure in retrieving top-ups by card number.
	ErrFailedFindAllTopupsByCardNumber = errors.ErrInternal.WithMessage("Failed to fetch topups by card number")

	// ErrFailedFindTopupById indicates failure in finding a top-up by its ID.
	ErrFailedFindTopupById = errors.ErrInternal.WithMessage("Failed to find topup by ID")

	// ErrFailedFindActiveTopups indicates failure in retrieving active (non-trashed) top-up records.
	ErrFailedFindActiveTopups = errors.ErrInternal.WithMessage("Failed to fetch active topups")

	// ErrFailedFindTrashedTopups indicates failure in retrieving trashed (soft-deleted) top-up records.
	ErrFailedFindTrashedTopups = errors.ErrInternal.WithMessage("Failed to fetch trashed topups")
)
