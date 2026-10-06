package userserviceerrors

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

var (
	// ErrFailedCreateUser is returned when creating a user fails.
	ErrFailedCreateUser = errors.ErrInternal.WithMessage("Failed to create user")

	// ErrFailedUpdateUser is returned when updating a user fails.
	ErrFailedUpdateUser = errors.ErrInternal.WithMessage("Failed to update user")

	// ErrFailedTrashedUser is returned when moving a user to trash fails.
	ErrFailedTrashedUser = errors.ErrInternal.WithMessage("Failed to move user to trash")

	// ErrFailedRestoreUser is returned when restoring a user fails.
	ErrFailedRestoreUser = errors.ErrInternal.WithMessage("Failed to restore user")

	// ErrFailedDeletePermanent is returned when permanently deleting a user fails.
	ErrFailedDeletePermanent = errors.ErrInternal.WithMessage("Failed to delete user permanently")

	// ErrFailedRestoreAll is returned when restoring all users fails.
	ErrFailedRestoreAll = errors.ErrInternal.WithMessage("Failed to restore all users")

	// ErrFailedDeleteAll is returned when permanently deleting all users fails.
	ErrFailedDeleteAll = errors.ErrInternal.WithMessage("Failed to delete all users permanently")
)
