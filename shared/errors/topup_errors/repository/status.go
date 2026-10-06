package topuprepositoryerrors

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

var (
	// ErrGetMonthTopupStatusSuccessFailed indicates failure in getting the monthly count of successful top-ups.
	ErrGetMonthTopupStatusSuccessFailed = errors.ErrInternal.WithMessage("Failed to get monthly topup status success")

	// ErrGetYearlyTopupStatusSuccessFailed indicates failure in getting the yearly count of successful top-ups.
	ErrGetYearlyTopupStatusSuccessFailed = errors.ErrInternal.WithMessage("Failed to get yearly topup status success")

	// ErrGetMonthTopupStatusSuccessByCardFailed indicates failure in getting monthly successful top-up status by card number.
	ErrGetMonthTopupStatusSuccessByCardFailed = errors.ErrInternal.WithMessage("Failed to get monthly topup status success by card number")

	// ErrGetYearlyTopupStatusSuccessByCardFailed indicates failure in getting yearly successful top-up status by card number.
	ErrGetYearlyTopupStatusSuccessByCardFailed = errors.ErrInternal.WithMessage("Failed to get yearly topup status success by card number")

	// ErrGetMonthTopupStatusFailedFailed indicates failure in getting the monthly count of failed top-ups.
	ErrGetMonthTopupStatusFailedFailed = errors.ErrInternal.WithMessage("Failed to get monthly topup status failed")

	// ErrGetYearlyTopupStatusFailedFailed indicates failure in getting the yearly count of failed top-ups.
	ErrGetYearlyTopupStatusFailedFailed = errors.ErrInternal.WithMessage("Failed to get yearly topup status failed")

	// ErrGetMonthTopupStatusFailedByCardFailed indicates failure in getting monthly failed top-up status by card number.
	ErrGetMonthTopupStatusFailedByCardFailed = errors.ErrInternal.WithMessage("Failed to get monthly topup status failed by card number")

	// ErrGetYearlyTopupStatusFailedByCardFailed indicates failure in getting yearly failed top-up status by card number.
	ErrGetYearlyTopupStatusFailedByCardFailed = errors.ErrInternal.WithMessage("Failed to get yearly topup status failed by card number")
)
