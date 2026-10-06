package cardserviceerrors

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

var (
	// ErrCardNotFoundRes is an error response when a requested card was not found.
	ErrCardNotFoundRes = errors.ErrNotFound.WithMessage("Card not found")

	// ErrFailedFindAllCards is an error response when retrieving all card records fails.
	ErrFailedFindAllCards = errors.ErrInternal.WithMessage("Failed to fetch cards")

	// ErrFailedFindActiveCards is an error response when retrieving active card records fails.
	ErrFailedFindActiveCards = errors.ErrInternal.WithMessage("Failed to fetch active cards")

	// ErrFailedFindTrashedCards is an error response when retrieving trashed card records fails.
	ErrFailedFindTrashedCards = errors.ErrInternal.WithMessage("Failed to fetch trashed cards")

	// ErrFailedFindById is an error response when finding a card by its ID fails.
	ErrFailedFindById = errors.ErrInternal.WithMessage("Failed to find card by ID")

	// ErrFailedFindByUserID is an error response when finding a card by its user ID fails.
	ErrFailedFindByUserID = errors.ErrInternal.WithMessage("Failed to find card by user ID")

	// ErrFailedFindByCardNumber is an error response when finding a card by its card number fails.
	ErrFailedFindByCardNumber = errors.ErrInternal.WithMessage("Failed to find card by card number")
)
