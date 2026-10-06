package transactionrepositoryerrors

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

var (
	// ErrGetMonthlyAmountsFailed indicates a failure when retrieving total monthly transaction amounts.
	ErrGetMonthlyAmountsFailed = errors.ErrInternal.WithMessage("Failed to get monthly amounts")

	// ErrGetYearlyAmountsFailed indicates a failure when retrieving total yearly transaction amounts.
	ErrGetYearlyAmountsFailed = errors.ErrInternal.WithMessage("Failed to get yearly amounts")

	// ErrGetMonthlyAmountsByCardFailed indicates a failure when retrieving monthly amounts by card number.
	ErrGetMonthlyAmountsByCardFailed = errors.ErrInternal.WithMessage("Failed to get monthly amounts by card number")

	// ErrGetYearlyAmountsByCardFailed indicates a failure when retrieving yearly amounts by card number.
	ErrGetYearlyAmountsByCardFailed = errors.ErrInternal.WithMessage("Failed to get yearly amounts by card number")
)
