package cardrepositoryerrors

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

var (
	// ErrInvalidCardRequest is returned when the card request data is invalid.
	ErrInvalidCardRequest = errors.ErrBadRequest.WithMessage("Invalid card request data")

	// ErrInvalidCardId is returned when the card ID is invalid.
	ErrInvalidCardId = errors.ErrBadRequest.WithMessage("Invalid card ID")

	// ErrInvalidUserId is returned when the user ID is invalid.
	ErrInvalidUserId = errors.ErrBadRequest.WithMessage("Invalid user ID")

	// ErrInvalidCardNumber is returned when the card number is invalid.
	ErrInvalidCardNumber = errors.ErrBadRequest.WithMessage("Invalid card number")
)
