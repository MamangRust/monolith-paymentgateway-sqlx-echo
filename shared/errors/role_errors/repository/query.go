package rolerepositoryerrors

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

var (
	// ErrRoleNotFound indicates that the requested Role was not found in the database.
	ErrRoleNotFound = errors.ErrNotFound.WithMessage("Role not found")

	// ErrFindAllRoles is returned when retrieving all Roles from the database fails.
	ErrFindAllRoles = errors.ErrInternal.WithMessage("Failed to find all roles")

	// ErrFindActiveRoles is returned when retrieving all active Roles fails.
	ErrFindActiveRoles = errors.ErrInternal.WithMessage("Failed to find active roles")

	// ErrFindTrashedRoles is returned when retrieving trashed (soft-deleted) Roles fails.
	ErrFindTrashedRoles = errors.ErrInternal.WithMessage("Failed to find trashed roles")

	// ErrRoleConflict indicates a conflict where a Role already exists.
	ErrRoleConflict = errors.ErrConflict.WithMessage("Role already exists")
)
