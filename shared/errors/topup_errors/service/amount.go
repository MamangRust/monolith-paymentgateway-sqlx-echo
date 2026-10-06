package topupserviceerrors

import (
	"net/http"

	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

// ErrFailedFindMonthlyTopupAmounts indicates failure in retrieving monthly top-up amounts.
var ErrFailedFindMonthlyTopupAmounts = errors.NewErrorResponse("Failed to get monthly topup amounts", http.StatusInternalServerError)

// ErrFailedFindYearlyTopupAmounts indicates failure in retrieving yearly top-up amounts.
var ErrFailedFindYearlyTopupAmounts = errors.NewErrorResponse("Failed to get yearly topup amounts", http.StatusInternalServerError)

// ErrFailedFindMonthlyTopupAmountsByCard indicates failure in retrieving monthly top-up amounts by card.
var ErrFailedFindMonthlyTopupAmountsByCard = errors.NewErrorResponse("Failed to get monthly topup amounts by card", http.StatusInternalServerError)

// ErrFailedFindYearlyTopupAmountsByCard indicates failure in retrieving yearly top-up amounts by card.
var ErrFailedFindYearlyTopupAmountsByCard = errors.NewErrorResponse("Failed to get yearly topup amounts by card", http.StatusInternalServerError)
