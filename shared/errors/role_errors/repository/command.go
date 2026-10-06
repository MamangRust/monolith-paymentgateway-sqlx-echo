package rolerepositoryerrors

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

var (
	// ErrCreateRole is returned when creating a new Role fails.
	ErrCreateRole = errors.ErrInternal.WithMessage("Failed to create role")

	// ErrUpdateRole is returned when updating an existing Role fails.
	ErrUpdateRole = errors.ErrInternal.WithMessage("Failed to update role")

	// ErrTrashedRole is returned when moving a Role to trash (soft-delete) fails.
	ErrTrashedRole = errors.ErrInternal.WithMessage("Failed to move role to trash")

	// ErrRestoreRole is returned when restoring a trashed Role fails.
	ErrRestoreRole = errors.ErrInternal.WithMessage("Failed to restore role from trash")

	// ErrDeleteRolePermanent is returned when permanently deleting a Role fails.
	ErrDeleteRolePermanent = errors.ErrInternal.WithMessage("Failed to permanently delete role")

	// ErrRestoreAllRoles is returned when restoring all trashed Roles fails.
	ErrRestoreAllRoles = errors.ErrInternal.WithMessage("Failed to restore all roles")

	// ErrDeleteAllRoles is returned when permanently deleting all Roles fails.
	ErrDeleteAllRoles = errors.ErrInternal.WithMessage("Failed to permanently delete all roles")
)
