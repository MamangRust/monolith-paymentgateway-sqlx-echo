package transferrepositoryerrors

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

var (
	// ErrGetMonthTransferStatusSuccessFailed indicates a failure when retrieving monthly successful transfer statistics.
	ErrGetMonthTransferStatusSuccessFailed = errors.ErrInternal.WithMessage("Failed to get monthly transfer status success")

	// ErrGetYearlyTransferStatusSuccessFailed indicates a failure when retrieving yearly successful transfer statistics.
	ErrGetYearlyTransferStatusSuccessFailed = errors.ErrInternal.WithMessage("Failed to get yearly transfer status success")

	// ErrGetMonthTransferStatusSuccessByCardFailed indicates a failure when retrieving monthly successful transfers filtered by card number.
	ErrGetMonthTransferStatusSuccessByCardFailed = errors.ErrInternal.WithMessage("Failed to get monthly transfer status success by card number")

	// ErrGetYearlyTransferStatusSuccessByCardFailed indicates a failure when retrieving yearly successful transfers filtered by card number.
	ErrGetYearlyTransferStatusSuccessByCardFailed = errors.ErrInternal.WithMessage("Failed to get yearly transfer status success by card number")

	// ErrGetMonthTransferStatusFailedFailed indicates a failure when retrieving monthly failed transfer statistics.
	ErrGetMonthTransferStatusFailedFailed = errors.ErrInternal.WithMessage("Failed to get monthly transfer status failed")

	// ErrGetYearlyTransferStatusFailedFailed indicates a failure when retrieving yearly failed transfer statistics.
	ErrGetYearlyTransferStatusFailedFailed = errors.ErrInternal.WithMessage("Failed to get yearly transfer status failed")

	// ErrGetMonthTransferStatusFailedByCardFailed indicates a failure when retrieving monthly failed transfers filtered by card number.
	ErrGetMonthTransferStatusFailedByCardFailed = errors.ErrInternal.WithMessage("Failed to get monthly transfer status failed by card number")

	// ErrGetYearlyTransferStatusFailedByCardFailed indicates a failure when retrieving yearly failed transfers filtered by card number.
	ErrGetYearlyTransferStatusFailedByCardFailed = errors.ErrInternal.WithMessage("Failed to get yearly transfer status failed by card number")
)
