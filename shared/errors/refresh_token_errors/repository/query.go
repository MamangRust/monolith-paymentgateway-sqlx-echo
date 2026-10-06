package refreshtokenrepositoryerrors

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

var (
	// ErrTokenNotFound indicates that the refresh token could not be found.
	ErrTokenNotFound = errors.ErrNotFound.WithMessage("Refresh token not found")

	// ErrFindByToken is returned when a lookup for the refresh token by token value fails.
	ErrFindByToken = errors.ErrInternal.WithMessage("Failed to find refresh token by token")

	// ErrFindByUserID is returned when a lookup for the refresh token by user ID fails.
	ErrFindByUserID = errors.ErrInternal.WithMessage("Failed to find refresh token by user ID")
)
