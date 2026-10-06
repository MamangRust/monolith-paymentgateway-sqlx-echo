package transactonserviceerrors

import (
	"net/http"

	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

// ErrFailedFindMonthTransactionSuccess indicates a failure when retrieving monthly successful transactions.
var ErrFailedFindMonthTransactionSuccess = errors.NewErrorResponse("Failed to fetch monthly successful transactions", http.StatusInternalServerError)

// ErrFailedFindYearTransactionSuccess indicates a failure when retrieving yearly successful transactions.
var ErrFailedFindYearTransactionSuccess = errors.NewErrorResponse("Failed to fetch yearly successful transactions", http.StatusInternalServerError)

// ErrFailedFindMonthTransactionFailed indicates a failure when retrieving monthly failed transactions.
var ErrFailedFindMonthTransactionFailed = errors.NewErrorResponse("Failed to fetch monthly failed transactions", http.StatusInternalServerError)

// ErrFailedFindYearTransactionFailed indicates a failure when retrieving yearly failed transactions.
var ErrFailedFindYearTransactionFailed = errors.NewErrorResponse("Failed to fetch yearly failed transactions", http.StatusInternalServerError)

// ErrFailedFindMonthTransactionSuccessByCard indicates a failure when retrieving monthly successful transactions filtered by card number.
var ErrFailedFindMonthTransactionSuccessByCard = errors.NewErrorResponse("Failed to fetch monthly successful transactions by card", http.StatusInternalServerError)

// ErrFailedFindYearTransactionSuccessByCard indicates a failure when retrieving yearly successful transactions filtered by card number.
var ErrFailedFindYearTransactionSuccessByCard = errors.NewErrorResponse("Failed to fetch yearly successful transactions by card", http.StatusInternalServerError)

// ErrFailedFindMonthTransactionFailedByCard indicates a failure when retrieving monthly failed transactions filtered by card number.
var ErrFailedFindMonthTransactionFailedByCard = errors.NewErrorResponse("Failed to fetch monthly failed transactions by card", http.StatusInternalServerError)

// ErrFailedFindYearTransactionFailedByCard indicates a failure when retrieving yearly failed transactions filtered by card number.
var ErrFailedFindYearTransactionFailedByCard = errors.NewErrorResponse("Failed to fetch yearly failed transactions by card", http.StatusInternalServerError)
