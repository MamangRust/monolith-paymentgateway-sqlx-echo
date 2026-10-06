package cardrepositoryerrors

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

var (
	// ErrFindAllCardsFailed is returned when fetching all card records fails.
	ErrFindAllCardsFailed = errors.ErrInternal.WithMessage("Failed to find all cards")

	// ErrFindActiveCardsFailed is returned when fetching active card records fails.
	ErrFindActiveCardsFailed = errors.ErrInternal.WithMessage("Failed to find active cards")

	// ErrFindTrashedCardsFailed is returned when fetching trashed card records fails.
	ErrFindTrashedCardsFailed = errors.ErrInternal.WithMessage("Failed to find trashed cards")

	// ErrFindCardByIdFailed is returned when fetching a card by its ID fails.
	ErrFindCardByIdFailed = errors.ErrInternal.WithMessage("Failed to find card by ID")

	// ErrFindCardByUserIdFailed is returned when fetching a card by user ID fails.
	ErrFindCardByUserIdFailed = errors.ErrInternal.WithMessage("Failed to find card by user ID")

	// ErrFindCardByCardNumberFailed is returned when fetching a card by card number fails.
	ErrFindCardByCardNumberFailed = errors.ErrInternal.WithMessage("Failed to find card by card number")

	// ErrCardAlreadyExists is returned when a card already exists.
	ErrCardAlreadyExists = errors.ErrConflict.WithMessage("Card already exists")

	// ErrCardNotFound is returned when a card is not found.
	ErrCardNotFound = errors.ErrNotFound.WithMessage("Card not found")

	// ErrFindUserCardByCardNumberFailed is returned when fetching a card by user and card number fails.
	ErrFindUserCardByCardNumberFailed = errors.ErrInternal.WithMessage("Failed to find card by user ID")
)
