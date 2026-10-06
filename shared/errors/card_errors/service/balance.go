package cardserviceerrors

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

var (
	// ErrFailedFindMonthlyBalance indicates a failure in retrieving the monthly balance.
	ErrFailedFindMonthlyBalance = errors.ErrInternal.WithMessage("Failed to get monthly balance")

	// ErrFailedFindYearlyBalance indicates a failure in retrieving the yearly balance.
	ErrFailedFindYearlyBalance = errors.ErrInternal.WithMessage("Failed to get yearly balance")

	// ErrFailedFindMonthlyBalanceByCard returns an error response when retrieving monthly balance by card number fails.
	ErrFailedFindMonthlyBalanceByCard = errors.ErrInternal.WithMessage("Failed to get monthly balance by card")

	// ErrFailedFindYearlyBalanceByCard returns an error response when retrieving yearly balance by card number fails.
	ErrFailedFindYearlyBalanceByCard = errors.ErrInternal.WithMessage("Failed to get yearly balance by card")
)
