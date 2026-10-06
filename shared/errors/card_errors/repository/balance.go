package cardrepositoryerrors

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

var (
	// ErrGetMonthlyBalanceFailed is returned when fetching monthly balances fails.
	ErrGetMonthlyBalanceFailed = errors.ErrInternal.WithMessage("Failed to get monthly balance")

	// ErrGetYearlyBalanceFailed is returned when fetching yearly balances fails.
	ErrGetYearlyBalanceFailed = errors.ErrInternal.WithMessage("Failed to get yearly balance")

	// ErrGetMonthlyBalanceByCardFailed is returned when fetching monthly balances by card fails.
	ErrGetMonthlyBalanceByCardFailed = errors.ErrInternal.WithMessage("Failed to get monthly balance by card")

	// ErrGetYearlyBalanceByCardFailed is returned when fetching yearly balances by card fails.
	ErrGetYearlyBalanceByCardFailed = errors.ErrInternal.WithMessage("Failed to get yearly balance by card")
)
