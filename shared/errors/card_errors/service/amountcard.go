package cardserviceerrors

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

var (
	// ErrFailedFindMonthlyTopupAmountByCard returns an error response when retrieving monthly top-up amount by card number fails.
	ErrFailedFindMonthlyTopupAmountByCard = errors.ErrInternal.WithMessage("Failed to get monthly topup amount by card")

	// ErrFailedFindYearlyTopupAmountByCard returns an error response when retrieving yearly top-up amount by card number fails.
	ErrFailedFindYearlyTopupAmountByCard = errors.ErrInternal.WithMessage("Failed to get yearly topup amount by card")

	// ErrFailedFindMonthlyWithdrawAmountByCard returns an error response when retrieving monthly withdraw amount by card number fails.
	ErrFailedFindMonthlyWithdrawAmountByCard = errors.ErrInternal.WithMessage("Failed to get monthly withdraw amount by card")

	// ErrFailedFindYearlyWithdrawAmountByCard returns an error response when retrieving yearly withdraw amount by card number fails.
	ErrFailedFindYearlyWithdrawAmountByCard = errors.ErrInternal.WithMessage("Failed to get yearly withdraw amount by card")

	// ErrFailedFindMonthlyTransactionAmountByCard returns an error response when retrieving monthly transaction amount by card number fails.
	ErrFailedFindMonthlyTransactionAmountByCard = errors.ErrInternal.WithMessage("Failed to get monthly transaction amount by card")

	// ErrFailedFindYearlyTransactionAmountByCard returns an error response when retrieving yearly transaction amount by card number fails.
	ErrFailedFindYearlyTransactionAmountByCard = errors.ErrInternal.WithMessage("Failed to get yearly transaction amount by card")

	// ErrFailedFindMonthlyTransferAmountBySender returns an error response when retrieving monthly transfer amount by sender card number fails.
	ErrFailedFindMonthlyTransferAmountBySender = errors.ErrInternal.WithMessage("Failed to get monthly transfer amount by sender")

	// ErrFailedFindYearlyTransferAmountBySender returns an error response when retrieving yearly transfer amount by sender card number fails.
	ErrFailedFindYearlyTransferAmountBySender = errors.ErrInternal.WithMessage("Failed to get yearly transfer amount by sender")

	// ErrFailedFindMonthlyTransferAmountByReceiver returns an error response when retrieving monthly transfer amount by receiver card number fails.
	ErrFailedFindMonthlyTransferAmountByReceiver = errors.ErrInternal.WithMessage("Failed to get monthly transfer amount by receiver")

	// ErrFailedFindYearlyTransferAmountByReceiver returns an error response when retrieving yearly transfer amount by receiver card number fails.
	ErrFailedFindYearlyTransferAmountByReceiver = errors.ErrInternal.WithMessage("Failed to get yearly transfer amount by receiver")

	// ErrFailedFindDailyTopupAmountByCard returns an error response when retrieving daily top-up amount by card number fails.
	ErrFailedFindDailyTopupAmountByCard = errors.ErrInternal.WithMessage("Failed to get daily topup amount by card")

	// ErrFailedFindDailyWithdrawAmountByCard returns an error response when retrieving daily withdraw amount by card number fails.
	ErrFailedFindDailyWithdrawAmountByCard = errors.ErrInternal.WithMessage("Failed to get daily withdraw amount by card")

	// ErrFailedFindDailyTransactionAmountByCard returns an error response when retrieving daily transaction amount by card number fails.
	ErrFailedFindDailyTransactionAmountByCard = errors.ErrInternal.WithMessage("Failed to get daily transaction amount by card")

	// ErrFailedFindDailyTransferAmountBySender returns an error response when retrieving daily transfer amount by sender card number fails.
	ErrFailedFindDailyTransferAmountBySender = errors.ErrInternal.WithMessage("Failed to get daily transfer amount by sender")

	// ErrFailedFindDailyTransferAmountByReceiver returns an error response when retrieving daily transfer amount by receiver card number fails.
	ErrFailedFindDailyTransferAmountByReceiver = errors.ErrInternal.WithMessage("Failed to get daily transfer amount by receiver")
)
