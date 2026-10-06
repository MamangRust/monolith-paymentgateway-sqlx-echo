package transactonserviceerrors

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

var (
	// ErrFailedFindAllTransactions indicates a failure when retrieving all transaction records.
	ErrFailedFindAllTransactions = errors.ErrInternal.WithMessage("Failed to fetch all transactions")

	// ErrFailedFindAllByCardNumber indicates a failure when retrieving transactions filtered by card number.
	ErrFailedFindAllByCardNumber = errors.ErrInternal.WithMessage("Failed to fetch transactions by card number")

	// ErrTransactionNotFound indicates that the requested transaction could not be found.
	ErrTransactionNotFound = errors.ErrNotFound.WithMessage("Transaction not found")

	// ErrFailedFindByActiveTransactions indicates a failure when retrieving active (non-deleted) transactions.
	ErrFailedFindByActiveTransactions = errors.ErrInternal.WithMessage("Failed to fetch active transactions")

	// ErrFailedFindByTrashedTransactions indicates a failure when retrieving trashed (soft-deleted) transactions.
	ErrFailedFindByTrashedTransactions = errors.ErrInternal.WithMessage("Failed to fetch trashed transactions")

	// ErrFailedFindByMerchantID indicates a failure when retrieving transactions filtered by merchant ID.
	ErrFailedFindByMerchantID = errors.ErrInternal.WithMessage("Failed to fetch transactions by merchant ID")
)
