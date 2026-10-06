package cardserviceerrors

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

var (
	// ErrFailedDashboardCard is an error response when retrieving card dashboard fails.
	ErrFailedDashboardCard = errors.ErrInternal.WithMessage("Failed to get card dashboard")

	// ErrFailedDashboardCardNumber is an error response when retrieving card dashboard by card number fails.
	ErrFailedDashboardCardNumber = errors.ErrInternal.WithMessage("Failed to get card dashboard by card number")
)
