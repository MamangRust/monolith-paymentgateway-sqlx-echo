package refreshtokenserviceerrors

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

var (
	// ErrFailedCreateAccess is returned when the creation of an access token fails.
	ErrFailedCreateAccess = errors.ErrInternal.WithMessage("Failed to create access token")

	// ErrFailedCreateRefresh is returned when the creation of a refresh token fails.
	ErrFailedCreateRefresh = errors.ErrInternal.WithMessage("Failed to create refresh token")

	// ErrFailedCreateRefreshToken is returned when refresh token creation fails.
	ErrFailedCreateRefreshToken = errors.ErrInternal.WithMessage("Failed to create refresh token")

	// ErrFailedUpdateRefreshToken is returned when refresh token update fails.
	ErrFailedUpdateRefreshToken = errors.ErrInternal.WithMessage("Failed to update refresh token")

	// ErrFailedDeleteRefreshToken is returned when refresh token deletion fails.
	ErrFailedDeleteRefreshToken = errors.ErrInternal.WithMessage("Failed to delete refresh token")

	// ErrFailedDeleteByUserID is returned when deletion of a refresh token by user ID fails.
	ErrFailedDeleteByUserID = errors.ErrInternal.WithMessage("Failed to delete refresh token by user ID")
)
