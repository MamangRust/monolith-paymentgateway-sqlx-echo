package cardrepositoryerrors

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

var (
	// ErrGetTotalBalancesFailed is returned when fetching total balances fails.
	ErrGetTotalBalancesFailed = errors.ErrInternal.WithMessage("Failed to get total balances")

	// ErrGetTotalTopAmountFailed is returned when fetching the total top-up amount fails.
	ErrGetTotalTopAmountFailed = errors.ErrInternal.WithMessage("Failed to get total topup amount")

	// ErrGetTotalWithdrawAmountFailed is returned when fetching the total withdrawal amount fails.
	ErrGetTotalWithdrawAmountFailed = errors.ErrInternal.WithMessage("Failed to get total withdraw amount")

	// ErrGetTotalTransactionAmountFailed is returned when fetching the total transaction amount fails.
	ErrGetTotalTransactionAmountFailed = errors.ErrInternal.WithMessage("Failed to get total transaction amount")

	// ErrGetTotalTransferAmountFailed is returned when fetching the total transfer amount fails.
	ErrGetTotalTransferAmountFailed = errors.ErrInternal.WithMessage("Failed to get total transfer amount")

	// ErrGetTotalBalanceByCardFailed is returned when fetching the total balance by card fails.
	ErrGetTotalBalanceByCardFailed = errors.ErrInternal.WithMessage("Failed to get total balance by card")

	// ErrGetTotalTopupAmountByCardFailed is returned when fetching the total top-up amount by card fails.
	ErrGetTotalTopupAmountByCardFailed = errors.ErrInternal.WithMessage("Failed to get total topup amount by card")

	// ErrGetTotalWithdrawAmountByCardFailed is returned when fetching the total withdrawal amount by card fails.
	ErrGetTotalWithdrawAmountByCardFailed = errors.ErrInternal.WithMessage("Failed to get total withdraw amount by card")

	// ErrGetTotalTransactionAmountByCardFailed is returned when fetching the total transaction amount by card fails.
	ErrGetTotalTransactionAmountByCardFailed = errors.ErrInternal.WithMessage("Failed to get total transaction amount by card")

	// ErrGetTotalTransferAmountBySenderFailed is returned when fetching the total transfer amount by sender fails.
	ErrGetTotalTransferAmountBySenderFailed = errors.ErrInternal.WithMessage("Failed to get total transfer amount by sender")

	// ErrGetTotalTransferAmountByReceiverFailed is returned when fetching the total transfer amount by receiver fails.
	ErrGetTotalTransferAmountByReceiverFailed = errors.ErrInternal.WithMessage("Failed to get total transfer amount by receiver")

	// ErrGetTotalWithdrawsFailed is returned when fetching total withdrawals fails.
	ErrGetTotalWithdrawsFailed = errors.ErrInternal.WithMessage("Failed to get total withdraws")

	// ErrGetTotalTopupsFailed is returned when fetching total topups fails.
	ErrGetTotalTopupsFailed = errors.ErrInternal.WithMessage("Failed to get total topups")

	// ErrGetTotalTransactionsFailed is returned when fetching total transactions fails.
	ErrGetTotalTransactionsFailed = errors.ErrInternal.WithMessage("Failed to get total transactions")

	// ErrGetTotalTransfersFailed is returned when fetching total transfers fails.
	ErrGetTotalTransfersFailed = errors.ErrInternal.WithMessage("Failed to get total transfers")

	// ErrGetTotalTransferSenderByCardFailed is returned when fetching total transfer sender by card number fails.
	ErrGetTotalTransferSenderByCardFailed = errors.ErrInternal.WithMessage("Failed to get total transfer sender by card")

	// ErrGetTotalTransferReceiverByCardFailed is returned when fetching total transfer receiver by card number fails.
	ErrGetTotalTransferReceiverByCardFailed = errors.ErrInternal.WithMessage("Failed to get total transfer receiver by card")
)
