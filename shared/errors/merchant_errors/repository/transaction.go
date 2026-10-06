package merchantrepositoryerrors

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

var (
	// ErrFindAllTransactionsFailed indicates failure fetching all transactions.
	ErrFindAllTransactionsFailed = errors.ErrInternal.WithMessage("failed to find all merchant transactions")

	// ErrFindAllTransactionsByMerchantFailed indicates failure fetching transactions by merchant ID.
	ErrFindAllTransactionsByMerchantFailed = errors.ErrInternal.WithMessage("failed to find merchant transactions by merchant ID")

	// ErrFindAllTransactionsByApiKeyFailed indicates failure fetching transactions by API key.
	ErrFindAllTransactionsByApiKeyFailed = errors.ErrInternal.WithMessage("failed to find merchant transactions by API key")
)
