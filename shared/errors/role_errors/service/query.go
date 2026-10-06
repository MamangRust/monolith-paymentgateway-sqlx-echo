package roleserviceerrors

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

var (
	// ErrRoleNotFoundRes is returned when the requested role is not found.
	ErrRoleNotFoundRes = errors.ErrNotFound.WithMessage("Role not found")

	// ErrFailedFindAll is returned when there is a failure in fetching all roles.
	ErrFailedFindAll = errors.ErrInternal.WithMessage("Failed to fetch roles")

	// ErrFailedFindActive is returned when there is a failure in fetching active roles.
	ErrFailedFindActive = errors.ErrInternal.WithMessage("Failed to fetch active roles")

	// ErrFailedFindTrashed is returned when there is a failure in fetching trashed roles.
	ErrFailedFindTrashed = errors.ErrInternal.WithMessage("Failed to fetch trashed roles")
)
