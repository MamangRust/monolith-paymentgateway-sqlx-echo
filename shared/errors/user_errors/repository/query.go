package userrepositoryerrors

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

var (
	// ErrUserNotFound is returned when a user is not found.
	ErrUserNotFound = errors.ErrNotFound.WithMessage("User not found")

	// ErrFindAllUsers is returned when fetching all users fails.
	ErrFindAllUsers = errors.ErrInternal.WithMessage("Failed to find all users")

	// ErrFindActiveUsers is returned when fetching active users fails.
	ErrFindActiveUsers = errors.ErrInternal.WithMessage("Failed to find active users")

	// ErrFindTrashedUsers is returned when fetching trashed users fails.
	ErrFindTrashedUsers = errors.ErrInternal.WithMessage("Failed to find trashed users")
)
