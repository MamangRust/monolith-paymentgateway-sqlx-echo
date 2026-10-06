package userrepositoryerrors

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

var (
	// ErrUserConflict is returned when a user already exists.
	ErrUserConflict = errors.ErrConflict.WithMessage("User already exists")

	// ErrCreateUser is returned when creating a user fails.
	ErrCreateUser = errors.ErrInternal.WithMessage("Failed to create user")

	// ErrUpdateUser is returned when updating a user fails.
	ErrUpdateUser = errors.ErrInternal.WithMessage("Failed to update user")

	// ErrUpdateUserVerificationCode is returned when updating a user verification code fails.
	ErrUpdateUserVerificationCode = errors.ErrInternal.WithMessage("Failed to update user verification code")

	// ErrUpdateUserPassword is returned when updating a user password fails.
	ErrUpdateUserPassword = errors.ErrInternal.WithMessage("Failed to update user password")

	// ErrTrashedUser is returned when moving a user to trash fails.
	ErrTrashedUser = errors.ErrInternal.WithMessage("Failed to move user to trash")

	// ErrRestoreUser is returned when restoring a user from trash fails.
	ErrRestoreUser = errors.ErrInternal.WithMessage("Failed to restore user from trash")

	// ErrDeleteUserPermanent is returned when permanently deleting a user fails.
	ErrDeleteUserPermanent = errors.ErrInternal.WithMessage("Failed to permanently delete user")

	// ErrRestoreAllUsers is returned when restoring all users fails.
	ErrRestoreAllUsers = errors.ErrInternal.WithMessage("Failed to restore all users")

	// ErrDeleteAllUsers is returned when permanently deleting all users fails.
	ErrDeleteAllUsers = errors.ErrInternal.WithMessage("Failed to permanently delete all users")

	// ErrUserInternal is returned when an internal user error occurs.
	ErrUserInternal = errors.ErrInternal.WithMessage("User internal error")
)
