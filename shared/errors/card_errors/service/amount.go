package cardserviceerrors

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

var (
	// ErrFailedFindMonthlyTopupAmount indicates a failure in retrieving the monthly top-up amount.
	ErrFailedFindMonthlyTopupAmount = errors.ErrInternal.WithMessage("Failed to get monthly topup amount")

	// ErrFailedFindYearlyTopupAmount indicates a failure in retrieving the yearly top-up amount.
	ErrFailedFindYearlyTopupAmount = errors.ErrInternal.WithMessage("Failed to get yearly topup amount")

	// ErrFailedFindMonthlyWithdrawAmount indicates a failure in retrieving the monthly withdraw amount.
	ErrFailedFindMonthlyWithdrawAmount = errors.ErrInternal.WithMessage("Failed to get monthly withdraw amount")

	// ErrFailedFindYearlyWithdrawAmount indicates a failure in retrieving the yearly withdraw amount.
	ErrFailedFindYearlyWithdrawAmount = errors.ErrInternal.WithMessage("Failed to get yearly withdraw amount")

	// ErrFailedFindMonthlyTransactionAmount indicates a failure in retrieving the monthly transaction amount.
	ErrFailedFindMonthlyTransactionAmount = errors.ErrInternal.WithMessage("Failed to get monthly transaction amount")

	// ErrFailedFindYearlyTransactionAmount indicates a failure in retrieving the yearly transaction amount.
	ErrFailedFindYearlyTransactionAmount = errors.ErrInternal.WithMessage("Failed to get yearly transaction amount")

	// ErrFailedFindMonthlyTransferAmountSender indicates a failure in retrieving the monthly transfer amount by sender.
	ErrFailedFindMonthlyTransferAmountSender = errors.ErrInternal.WithMessage("Failed to get monthly transfer amount by sender")

	// ErrFailedFindYearlyTransferAmountSender indicates a failure in retrieving the yearly transfer amount by sender.
	ErrFailedFindYearlyTransferAmountSender = errors.ErrInternal.WithMessage("Failed to get yearly transfer amount by sender")

	// ErrFailedFindMonthlyTransferAmountReceiver indicates a failure in retrieving the monthly transfer amount by receiver.
	ErrFailedFindMonthlyTransferAmountReceiver = errors.ErrInternal.WithMessage("Failed to get monthly transfer amount by receiver")

	// ErrFailedFindYearlyTransferAmountReceiver indicates a failure in retrieving the yearly transfer amount by receiver.
	ErrFailedFindYearlyTransferAmountReceiver = errors.ErrInternal.WithMessage("Failed to get yearly transfer amount by receiver")

	// ErrFailedFindDailyTopupAmount indicates a failure in retrieving the daily top-up amount.
	ErrFailedFindDailyTopupAmount = errors.ErrInternal.WithMessage("Failed to get daily topup amount")

	// ErrFailedFindDailyWithdrawAmount indicates a failure in retrieving the daily withdraw amount.
	ErrFailedFindDailyWithdrawAmount = errors.ErrInternal.WithMessage("Failed to get daily withdraw amount")

	// ErrFailedFindDailyTransactionAmount indicates a failure in retrieving the daily transaction amount.
	ErrFailedFindDailyTransactionAmount = errors.ErrInternal.WithMessage("Failed to get daily transaction amount")

	// ErrFailedFindDailyTransferAmountSender indicates a failure in retrieving the daily transfer amount by sender.
	ErrFailedFindDailyTransferAmountSender = errors.ErrInternal.WithMessage("Failed to get daily transfer amount by sender")

	// ErrFailedFindDailyTransferAmountReceiver indicates a failure in retrieving the daily transfer amount by receiver.
	ErrFailedFindDailyTransferAmountReceiver = errors.ErrInternal.WithMessage("Failed to get daily transfer amount by receiver")
)
