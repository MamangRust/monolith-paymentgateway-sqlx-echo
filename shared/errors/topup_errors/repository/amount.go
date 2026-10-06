package topuprepositoryerrors

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

var (
	// ErrGetMonthlyTopupAmountsFailed indicates failure in retrieving the monthly top-up amounts.
	ErrGetMonthlyTopupAmountsFailed = errors.ErrInternal.WithMessage("Failed to get monthly topup amounts")

	// ErrGetYearlyTopupAmountsFailed indicates failure in retrieving the yearly top-up amounts.
	ErrGetYearlyTopupAmountsFailed = errors.ErrInternal.WithMessage("Failed to get yearly topup amounts")

	// ErrGetMonthlyTopupAmountsByCardFailed indicates failure in retrieving monthly top-up amount stats by card number.
	ErrGetMonthlyTopupAmountsByCardFailed = errors.ErrInternal.WithMessage("Failed to get monthly topup amounts by card number")

	// ErrGetYearlyTopupAmountsByCardFailed indicates failure in retrieving yearly top-up amount stats by card number.
	ErrGetYearlyTopupAmountsByCardFailed = errors.ErrInternal.WithMessage("Failed to get yearly topup amounts by card number")
)
