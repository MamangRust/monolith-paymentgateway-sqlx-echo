package withdrawrepositoryerrors

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

var (
	// ErrGetMonthWithdrawStatusSuccessFailed is used when the system fails to get the monthly withdraw status success
	ErrGetMonthWithdrawStatusSuccessFailed = errors.ErrInternal.WithMessage("Failed to get monthly withdraw status success")

	// ErrGetYearlyWithdrawStatusSuccessFailed is used when the system fails to get the yearly withdraw status success
	ErrGetYearlyWithdrawStatusSuccessFailed = errors.ErrInternal.WithMessage("Failed to get yearly withdraw status success")

	// ErrGetMonthWithdrawStatusSuccessByCardFailed is used when the system fails to get the monthly withdraw status success by card number
	ErrGetMonthWithdrawStatusSuccessByCardFailed = errors.ErrInternal.WithMessage("Failed to get monthly withdraw status success by card number")

	// ErrGetYearlyWithdrawStatusSuccessByCardFailed is used when the system fails to get the yearly withdraw status success by card number
	ErrGetYearlyWithdrawStatusSuccessByCardFailed = errors.ErrInternal.WithMessage("Failed to get yearly withdraw status success by card number")

	// ErrGetMonthWithdrawStatusFailedFailed is used when the system fails to get the monthly withdraw status failed
	ErrGetMonthWithdrawStatusFailedFailed = errors.ErrInternal.WithMessage("Failed to get monthly withdraw status failed")

	// ErrGetYearlyWithdrawStatusFailedFailed is used when the system fails to get the yearly withdraw status failed
	ErrGetYearlyWithdrawStatusFailedFailed = errors.ErrInternal.WithMessage("Failed to get yearly withdraw status failed")

	// ErrGetMonthWithdrawStatusFailedByCardFailed is used when the system fails to get the monthly withdraw status failed by card number
	ErrGetMonthWithdrawStatusFailedByCardFailed = errors.ErrInternal.WithMessage("Failed to get monthly withdraw status failed by card number")

	// ErrGetYearlyWithdrawStatusFailedByCardFailed is used when the system fails to get the yearly withdraw status failed by card number
	ErrGetYearlyWithdrawStatusFailedByCardFailed = errors.ErrInternal.WithMessage("Failed to get yearly withdraw status failed by card number")
)
