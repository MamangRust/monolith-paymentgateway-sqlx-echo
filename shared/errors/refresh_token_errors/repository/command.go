package refreshtokenrepositoryerrors

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

var (
	// ErrCreateRefreshToken indicates that an error occurred while creating a new refresh token.
	ErrCreateRefreshToken = errors.ErrInternal.WithMessage("Failed to create refresh token")

	// ErrUpdateRefreshToken is returned when the refresh token update process fails.
	ErrUpdateRefreshToken = errors.ErrInternal.WithMessage("Failed to update refresh token")

	// ErrDeleteRefreshToken is returned when deleting a refresh token fails.
	ErrDeleteRefreshToken = errors.ErrInternal.WithMessage("Failed to delete refresh token")

	// ErrDeleteByUserID indicates a failure when attempting to delete a refresh token using the user ID.
	ErrDeleteByUserID = errors.ErrInternal.WithMessage("Failed to delete refresh token by user ID")
)
