package userserviceerrors

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

var (
	// ErrInternalServerError is a generic internal server error.
	ErrInternalServerError = errors.ErrInternal.WithMessage("Internal server error")

	// ErrFailedSendEmail is returned when sending an email fails.
	ErrFailedSendEmail = errors.ErrInternal.WithMessage("Failed to send email")

	// ErrFailedPasswordNoMatch is returned when passwords do not match.
	ErrFailedPasswordNoMatch = errors.ErrUnauthorized.WithMessage("Password does not match")

	// ErrUserPassword is returned when there is an invalid password.
	ErrUserPassword = errors.ErrUnauthorized.WithMessage("Invalid password")

	// ErrAccountLocked is returned when the account is temporarily locked.
	ErrAccountLocked = errors.ErrForbidden.WithMessage("Account temporarily locked due to many failed attempts")
)
