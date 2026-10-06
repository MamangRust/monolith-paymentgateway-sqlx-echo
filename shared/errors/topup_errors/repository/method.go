package topuprepositoryerrors

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

var (
	// ErrGetMonthlyTopupMethodsFailed indicates failure in retrieving monthly top-up payment methods statistics.
	ErrGetMonthlyTopupMethodsFailed = errors.ErrInternal.WithMessage("Failed to get monthly topup methods")

	// ErrGetYearlyTopupMethodsFailed indicates failure in retrieving yearly top-up payment methods statistics.
	ErrGetYearlyTopupMethodsFailed = errors.ErrInternal.WithMessage("Failed to get yearly topup methods")

	// ErrGetMonthlyTopupMethodsByCardFailed indicates failure in retrieving monthly payment method stats by card number.
	ErrGetMonthlyTopupMethodsByCardFailed = errors.ErrInternal.WithMessage("Failed to get monthly topup methods by card number")

	// ErrGetYearlyTopupMethodsByCardFailed indicates failure in retrieving yearly payment method stats by card number.
	ErrGetYearlyTopupMethodsByCardFailed = errors.ErrInternal.WithMessage("Failed to get yearly topup methods by card number")
)
