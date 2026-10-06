package userrolerepositoryerrors

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

var (
	// ErrAssignRoleToUser is an error that is returned from the repository layer
	// when an error occurs while trying to assign a role to the user.
	ErrAssignRoleToUser = errors.ErrInternal.WithMessage("Failed to assign role to user")

	// ErrRemoveRole is an error that is returned from the repository layer
	// when an error occurs while trying to remove a role from the user.
	ErrRemoveRole = errors.ErrInternal.WithMessage("Failed to remove role from user")
)
