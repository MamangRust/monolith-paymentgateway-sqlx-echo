package withdrawrepositoryerrors

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

var (
	// ErrGetMonthlyWithdrawsFailed is used when the system fails to get monthly withdraw amounts
	ErrGetMonthlyWithdrawsFailed = errors.ErrInternal.WithMessage("Failed to get monthly withdraw amounts")

	// ErrGetYearlyWithdrawsFailed is used when the system fails to get yearly withdraw amounts
	ErrGetYearlyWithdrawsFailed = errors.ErrInternal.WithMessage("Failed to get yearly withdraw amounts")

	// ErrGetMonthlyWithdrawsByCardFailed indicates a failure when retrieving monthly withdraw amounts by card number.
	ErrGetMonthlyWithdrawsByCardFailed = errors.ErrInternal.WithMessage("Failed to get monthly withdraw amounts by card number")

	// ErrGetYearlyWithdrawsByCardFailed indicates a failure when retrieving yearly withdraw amounts by card number.
	ErrGetYearlyWithdrawsByCardFailed = errors.ErrInternal.WithMessage("Failed to get yearly withdraw amounts by card number")
)
