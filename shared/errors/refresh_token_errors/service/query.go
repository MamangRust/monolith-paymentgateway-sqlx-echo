package refreshtokenserviceerrors

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

var (
	// ErrRefreshTokenNotFound indicates that the refresh token was not found.
	ErrRefreshTokenNotFound = errors.ErrNotFound.WithMessage("Refresh token not found")

	// ErrFailedFindByToken indicates failure when searching for a refresh token by its token value.
	ErrFailedFindByToken = errors.ErrInternal.WithMessage("Failed to find refresh token by token")

	// ErrFailedFindByUserID indicates failure when searching for a refresh token by user ID.
	ErrFailedFindByUserID = errors.ErrInternal.WithMessage("Failed to find refresh token by user ID")
)
