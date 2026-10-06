package merchantserviceerrors

import (
	"net/http"

	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

// ErrFailedFindAllTransactions is returned when fetching all transactions fails.
var ErrFailedFindAllTransactions = errors.NewErrorResponse("Failed to fetch Merchant transactions", http.StatusInternalServerError)

// ErrFailedFindAllTransactionsByMerchant is returned when fetching transactions by merchant fails.
var ErrFailedFindAllTransactionsByMerchant = errors.NewErrorResponse("Failed to fetch transactions by Merchant", http.StatusInternalServerError)

// ErrFailedFindAllTransactionsByApikey is returned when fetching transactions by API key fails.
var ErrFailedFindAllTransactionsByApikey = errors.NewErrorResponse("Failed to fetch transactions by API key", http.StatusInternalServerError)
