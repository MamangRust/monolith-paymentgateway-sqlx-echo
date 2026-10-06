package cardserviceerrors

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

var (
	// ErrFailedFindTotalBalances is an error response when retrieving total balances fails.
	ErrFailedFindTotalBalances = errors.ErrInternal.WithMessage("Failed to find total balances")

	// ErrFailedFindTotalTopAmount is an error response when retrieving total topup amount fails.
	ErrFailedFindTotalTopAmount = errors.ErrInternal.WithMessage("Failed to find total topup amount")

	// ErrFailedFindTotalWithdrawAmount is an error response when retrieving total withdraw amount fails.
	ErrFailedFindTotalWithdrawAmount = errors.ErrInternal.WithMessage("Failed to find total withdraw amount")

	// ErrFailedFindTotalTransactionAmount is an error response when retrieving total transaction amount fails.
	ErrFailedFindTotalTransactionAmount = errors.ErrInternal.WithMessage("Failed to find total transaction amount")

	// ErrFailedFindTotalTransferAmount is an error response when retrieving total transfer amount fails.
	ErrFailedFindTotalTransferAmount = errors.ErrInternal.WithMessage("Failed to find total transfer amount")

	// ErrFailedFindTotalBalanceByCard is an error response when retrieving total balance by card fails.
	ErrFailedFindTotalBalanceByCard = errors.ErrInternal.WithMessage("Failed to find total balance by card")

	// ErrFailedFindTotalTopupAmountByCard is an error response when retrieving total topup amount by card fails.
	ErrFailedFindTotalTopupAmountByCard = errors.ErrInternal.WithMessage("Failed to find total topup amount by card")

	// ErrFailedFindTotalWithdrawAmountByCard is an error response when retrieving total withdraw amount by card fails.
	ErrFailedFindTotalWithdrawAmountByCard = errors.ErrInternal.WithMessage("Failed to find total withdraw amount by card")

	// ErrFailedFindTotalTransactionAmountByCard is an error response when retrieving total transaction amount by card fails.
	ErrFailedFindTotalTransactionAmountByCard = errors.ErrInternal.WithMessage("Failed to find total transaction amount by card")

	// ErrFailedFindTotalTransferAmountBySender is an error response when retrieving total transfer amount by sender fails.
	ErrFailedFindTotalTransferAmountBySender = errors.ErrInternal.WithMessage("Failed to find total transfer amount by sender")

	// ErrFailedFindTotalTransferAmountByReceiver is an error response when retrieving total transfer amount by receiver fails.
	ErrFailedFindTotalTransferAmountByReceiver = errors.ErrInternal.WithMessage("Failed to find total transfer amount by receiver")
)
