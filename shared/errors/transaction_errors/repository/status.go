package transactionrepositoryerrors

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

var (
	// ErrGetMonthTransactionStatusSuccessFailed indicates a failure when retrieving monthly successful transaction status.
	ErrGetMonthTransactionStatusSuccessFailed = errors.ErrInternal.WithMessage("Failed to get monthly transaction status success")

	// ErrGetYearlyTransactionStatusSuccessFailed indicates a failure when retrieving yearly successful transaction status.
	ErrGetYearlyTransactionStatusSuccessFailed = errors.ErrInternal.WithMessage("Failed to get yearly transaction status success")

	// ErrGetMonthTransactionStatusSuccessByCardFailed indicates a failure when retrieving monthly successful transactions by card number.
	ErrGetMonthTransactionStatusSuccessByCardFailed = errors.ErrInternal.WithMessage("Failed to get monthly transaction status success by card number")

	// ErrGetYearlyTransactionStatusSuccessByCardFailed indicates a failure when retrieving yearly successful transactions by card number.
	ErrGetYearlyTransactionStatusSuccessByCardFailed = errors.ErrInternal.WithMessage("Failed to get yearly transaction status success by card number")

	// ErrGetMonthTransactionStatusFailedFailed indicates a failure when retrieving monthly failed transaction status.
	ErrGetMonthTransactionStatusFailedFailed = errors.ErrInternal.WithMessage("Failed to get monthly transaction status failed")

	// ErrGetYearlyTransactionStatusFailedFailed indicates a failure when retrieving yearly failed transaction status.
	ErrGetYearlyTransactionStatusFailedFailed = errors.ErrInternal.WithMessage("Failed to get yearly transaction status failed")

	// ErrGetMonthTransactionStatusFailedByCardFailed indicates a failure when retrieving monthly failed transactions by card number.
	ErrGetMonthTransactionStatusFailedByCardFailed = errors.ErrInternal.WithMessage("Failed to get monthly transaction status failed by card number")

	// ErrGetYearlyTransactionStatusFailedByCardFailed indicates a failure when retrieving yearly failed transactions by card number.
	ErrGetYearlyTransactionStatusFailedByCardFailed = errors.ErrInternal.WithMessage("Failed to get yearly transaction status failed by card number")
)
