package roleserviceerrors

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

var (
	// ErrFailedCreateRole is returned when there is a failure in creating a role.
	ErrFailedCreateRole = errors.ErrInternal.WithMessage("Failed to create role")

	// ErrFailedUpdateRole is returned when there is a failure in updating a role.
	ErrFailedUpdateRole = errors.ErrInternal.WithMessage("Failed to update role")

	// ErrFailedTrashedRole is returned when there is a failure in moving a role to trash.
	ErrFailedTrashedRole = errors.ErrInternal.WithMessage("Failed to move role to trash")

	// ErrFailedRestoreRole is returned when there is a failure in restoring a trashed role.
	ErrFailedRestoreRole = errors.ErrInternal.WithMessage("Failed to restore role")

	// ErrFailedDeletePermanent is returned when there is a failure in permanently deleting a role.
	ErrFailedDeletePermanent = errors.ErrInternal.WithMessage("Failed to delete role permanently")

	// ErrFailedRestoreAll is returned when there is a failure in restoring all trashed roles.
	ErrFailedRestoreAll = errors.ErrInternal.WithMessage("Failed to restore all roles")

	// ErrFailedDeleteAll is returned when there is a failure in permanently deleting all roles.
	ErrFailedDeleteAll = errors.ErrInternal.WithMessage("Failed to delete all roles permanently")
)
