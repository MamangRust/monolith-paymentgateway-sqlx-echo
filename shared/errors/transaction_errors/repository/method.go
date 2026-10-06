package transactionrepositoryerrors

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

var (
	// ErrGetMonthlyPaymentMethodsFailed indicates a failure when retrieving monthly payment method statistics.
	ErrGetMonthlyPaymentMethodsFailed = errors.ErrInternal.WithMessage("Failed to get monthly payment methods")

	// ErrGetYearlyPaymentMethodsFailed indicates a failure when retrieving yearly payment method statistics.
	ErrGetYearlyPaymentMethodsFailed = errors.ErrInternal.WithMessage("Failed to get yearly payment methods")

	// ErrGetMonthlyPaymentMethodsByCardFailed indicates a failure when retrieving monthly payment methods by card number.
	ErrGetMonthlyPaymentMethodsByCardFailed = errors.ErrInternal.WithMessage("Failed to get monthly payment methods by card number")

	// ErrGetYearlyPaymentMethodsByCardFailed indicates a failure when retrieving yearly payment methods by card number.
	ErrGetYearlyPaymentMethodsByCardFailed = errors.ErrInternal.WithMessage("Failed to get yearly payment methods by card number")
)
