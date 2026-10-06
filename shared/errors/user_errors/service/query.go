package userserviceerrors

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

var (
	// ErrUserIDInValid is returned when an invalid user ID is provided.
	ErrUserIDInValid = errors.ErrBadRequest.WithMessage("Invalid user ID")

	// ErrUserNotFoundRes is returned when a user is not found.
	ErrUserNotFoundRes = errors.ErrNotFound.WithMessage("User not found")

	// ErrUserEmailAlready is returned when a user email already exists.
	// A duplicate email is a resource conflict (409), not a malformed request (400).
	ErrUserEmailAlready = errors.ErrConflict.WithMessage("User email already exists")

	// ErrFailedFindAll is returned when fetching users fails.
	ErrFailedFindAll = errors.ErrInternal.WithMessage("Failed to fetch users")

	// ErrFailedFindActive is returned when fetching active users fails.
	ErrFailedFindActive = errors.ErrInternal.WithMessage("Failed to fetch active users")

	// ErrFailedFindTrashed is returned when fetching trashed users fails.
	ErrFailedFindTrashed = errors.ErrInternal.WithMessage("Failed to fetch trashed users")
)
