package topupserviceerrors

import (
	"net/http"

	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

// ErrFailedFindMonthTopupStatusSuccess indicates failure in retrieving monthly successful top-up status.
var ErrFailedFindMonthTopupStatusSuccess = errors.NewErrorResponse("Failed to get monthly topup success status", http.StatusInternalServerError)

// ErrFailedFindYearlyTopupStatusSuccess indicates failure in retrieving yearly successful top-up status.
var ErrFailedFindYearlyTopupStatusSuccess = errors.NewErrorResponse("Failed to get yearly topup success status", http.StatusInternalServerError)

// ErrFailedFindMonthTopupStatusFailed indicates failure in retrieving monthly failed top-up status.
var ErrFailedFindMonthTopupStatusFailed = errors.NewErrorResponse("Failed to get monthly topup failed status", http.StatusInternalServerError)

// ErrFailedFindYearlyTopupStatusFailed indicates failure in retrieving yearly failed top-up status.
var ErrFailedFindYearlyTopupStatusFailed = errors.NewErrorResponse("Failed to get yearly topup failed status", http.StatusInternalServerError)

// ErrFailedFindMonthTopupStatusSuccessByCard indicates failure in retrieving monthly successful top-up status by card.
var ErrFailedFindMonthTopupStatusSuccessByCard = errors.NewErrorResponse("Failed to get monthly topup success status by card", http.StatusInternalServerError)

// ErrFailedFindYearlyTopupStatusSuccessByCard indicates failure in retrieving yearly successful top-up status by card.
var ErrFailedFindYearlyTopupStatusSuccessByCard = errors.NewErrorResponse("Failed to get yearly topup success status by card", http.StatusInternalServerError)

// ErrFailedFindMonthTopupStatusFailedByCard indicates failure in retrieving monthly failed top-up status by card.
var ErrFailedFindMonthTopupStatusFailedByCard = errors.NewErrorResponse("Failed to get monthly topup failed status by card", http.StatusInternalServerError)

// ErrFailedFindYearlyTopupStatusFailedByCard indicates failure in retrieving yearly failed top-up status by card.
var ErrFailedFindYearlyTopupStatusFailedByCard = errors.NewErrorResponse("Failed to get yearly topup failed status by card", http.StatusInternalServerError)
