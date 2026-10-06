package userroleserviceerrors

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

var (
	// ErrFailedAssignRoleToUser is an error that is returned from the service layer
	// when an error occurs while trying to assign a role to the user.
	ErrFailedAssignRoleToUser = errors.ErrInternal.WithMessage("Failed to assign role to user")

	// ErrFailedRemoveRole is an error that is returned from the service layer
	// when an error occurs while trying to remove a role from the user.
	ErrFailedRemoveRole = errors.ErrInternal.WithMessage("Failed to remove role from user")
)
