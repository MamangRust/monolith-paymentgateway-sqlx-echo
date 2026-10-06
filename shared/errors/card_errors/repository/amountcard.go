package cardrepositoryerrors

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

var (
	// ErrGetMonthlyTopupAmountByCardFailed is returned when fetching monthly top-up amounts by card fails.
	ErrGetMonthlyTopupAmountByCardFailed = errors.ErrInternal.WithMessage("Failed to get monthly topup amount by card")

	// ErrGetYearlyTopupAmountByCardFailed is returned when fetching yearly top-up amounts by card fails.
	ErrGetYearlyTopupAmountByCardFailed = errors.ErrInternal.WithMessage("Failed to get yearly topup amount by card")

	// ErrGetMonthlyWithdrawAmountByCardFailed is returned when fetching monthly withdrawal amounts by card fails.
	ErrGetMonthlyWithdrawAmountByCardFailed = errors.ErrInternal.WithMessage("Failed to get monthly withdraw amount by card")

	// ErrGetYearlyWithdrawAmountByCardFailed is returned when fetching yearly withdrawal amounts by card fails.
	ErrGetYearlyWithdrawAmountByCardFailed = errors.ErrInternal.WithMessage("Failed to get yearly withdraw amount by card")

	// ErrGetMonthlyTransactionAmountByCardFailed is returned when fetching monthly transaction amounts by card fails.
	ErrGetMonthlyTransactionAmountByCardFailed = errors.ErrInternal.WithMessage("Failed to get monthly transaction amount by card")

	// ErrGetYearlyTransactionAmountByCardFailed is returned when fetching yearly transaction amounts by card fails.
	ErrGetYearlyTransactionAmountByCardFailed = errors.ErrInternal.WithMessage("Failed to get yearly transaction amount by card")

	// ErrGetMonthlyTransferAmountBySenderFailed is returned when fetching monthly transfer amount by sender fails.
	ErrGetMonthlyTransferAmountBySenderFailed = errors.ErrInternal.WithMessage("Failed to get monthly transfer amount by sender")

	// ErrGetYearlyTransferAmountBySenderFailed is returned when fetching yearly transfer amount by sender fails.
	ErrGetYearlyTransferAmountBySenderFailed = errors.ErrInternal.WithMessage("Failed to get yearly transfer amount by sender")

	// ErrGetMonthlyTransferAmountByReceiverFailed is returned when fetching monthly transfer amount by receiver fails.
	ErrGetMonthlyTransferAmountByReceiverFailed = errors.ErrInternal.WithMessage("Failed to get monthly transfer amount by receiver")

	// ErrGetYearlyTransferAmountByReceiverFailed is returned when fetching yearly transfer amount by receiver fails.
	ErrGetYearlyTransferAmountByReceiverFailed = errors.ErrInternal.WithMessage("Failed to get yearly transfer amount by receiver")

	// ErrGetDailyTopupAmountByCardFailed is returned when fetching daily top-up amounts by card fails.
	ErrGetDailyTopupAmountByCardFailed = errors.ErrInternal.WithMessage("Failed to get daily topup amount by card")

	// ErrGetDailyWithdrawAmountByCardFailed is returned when fetching daily withdrawal amounts by card fails.
	ErrGetDailyWithdrawAmountByCardFailed = errors.ErrInternal.WithMessage("Failed to get daily withdraw amount by card")

	// ErrGetDailyTransactionAmountByCardFailed is returned when fetching daily transaction amounts by card fails.
	ErrGetDailyTransactionAmountByCardFailed = errors.ErrInternal.WithMessage("Failed to get daily transaction amount by card")

	// ErrGetDailyTransferAmountBySenderFailed is returned when fetching daily transfer amount by sender fails.
	ErrGetDailyTransferAmountBySenderFailed = errors.ErrInternal.WithMessage("Failed to get daily transfer amount by sender")

	// ErrGetDailyTransferAmountByReceiverFailed is returned when fetching daily transfer amount by receiver fails.
	ErrGetDailyTransferAmountByReceiverFailed = errors.ErrInternal.WithMessage("Failed to get daily transfer amount by receiver")
)
