package transferserviceerrors

import (
	"net/http"

	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

// ErrFailedFindMonthTransferStatusSuccess indicates a failure when retrieving monthly successful transfer statistics.
var ErrFailedFindMonthTransferStatusSuccess = errors.NewErrorResponse("Failed to fetch monthly successful transfers", http.StatusInternalServerError)

// ErrFailedFindYearTransferStatusSuccess indicates a failure when retrieving yearly successful transfer statistics.
var ErrFailedFindYearTransferStatusSuccess = errors.NewErrorResponse("Failed to fetch yearly successful transfers", http.StatusInternalServerError)

// ErrFailedFindMonthTransferStatusFailed indicates a failure when retrieving monthly failed transfer statistics.
var ErrFailedFindMonthTransferStatusFailed = errors.NewErrorResponse("Failed to fetch monthly failed transfers", http.StatusInternalServerError)

// ErrFailedFindYearTransferStatusFailed indicates a failure when retrieving yearly failed transfer statistics.
var ErrFailedFindYearTransferStatusFailed = errors.NewErrorResponse("Failed to fetch yearly failed transfers", http.StatusInternalServerError)

// ErrFailedFindMonthTransferStatusSuccessByCard indicates a failure when retrieving monthly successful transfers by card number.
var ErrFailedFindMonthTransferStatusSuccessByCard = errors.NewErrorResponse("Failed to fetch monthly successful transfers by card", http.StatusInternalServerError)

// ErrFailedFindYearTransferStatusSuccessByCard indicates a failure when retrieving yearly successful transfers by card number.
var ErrFailedFindYearTransferStatusSuccessByCard = errors.NewErrorResponse("Failed to fetch yearly successful transfers by card", http.StatusInternalServerError)

// ErrFailedFindMonthTransferStatusFailedByCard indicates a failure when retrieving monthly failed transfers by card number.
var ErrFailedFindMonthTransferStatusFailedByCard = errors.NewErrorResponse("Failed to fetch monthly failed transfers by card", http.StatusInternalServerError)

// ErrFailedFindYearTransferStatusFailedByCard indicates a failure when retrieving yearly failed transfers by card number.
var ErrFailedFindYearTransferStatusFailedByCard = errors.NewErrorResponse("Failed to fetch yearly failed transfers by card", http.StatusInternalServerError)
