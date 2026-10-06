package transactionrepositoryerrors

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

var (
	// ErrFindAllTransactionsFailed indicates a failure when attempting to retrieve all transactions.
	ErrFindAllTransactionsFailed = errors.ErrInternal.WithMessage("Failed to find all transactions")

	// ErrFindActiveTransactionsFailed indicates a failure when retrieving active (non-deleted) transactions.
	ErrFindActiveTransactionsFailed = errors.ErrInternal.WithMessage("Failed to find active transactions")

	// ErrFindTrashedTransactionsFailed indicates a failure when retrieving soft-deleted (trashed) transactions.
	ErrFindTrashedTransactionsFailed = errors.ErrInternal.WithMessage("Failed to find trashed transactions")

	// ErrFindTransactionsByCardNumberFailed indicates a failure when retrieving transactions by card number.
	ErrFindTransactionsByCardNumberFailed = errors.ErrInternal.WithMessage("Failed to find transactions by card number")

	// ErrFindTransactionByIdFailed indicates a failure when retrieving a transaction by its ID.
	ErrFindTransactionByIdFailed = errors.ErrInternal.WithMessage("Failed to find transaction by ID")

	// ErrFindTransactionByMerchantIdFailed indicates a failure when retrieving transactions by merchant ID.
	ErrFindTransactionByMerchantIdFailed = errors.ErrInternal.WithMessage("Failed to find transaction by merchant ID")

	// ErrTransactionNotFound is returned when a transaction is not found.
	ErrTransactionNotFound = errors.ErrNotFound.WithMessage("Transaction not found")
)
