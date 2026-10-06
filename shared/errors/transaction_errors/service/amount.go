package transactonserviceerrors

import (
	"net/http"

	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

// ErrFailedFindMonthlyAmounts indicates a failure when retrieving the total monthly transaction amounts.
var ErrFailedFindMonthlyAmounts = errors.NewErrorResponse("Failed to fetch monthly amounts", http.StatusInternalServerError)

// ErrFailedFindYearlyAmounts indicates a failure when retrieving the total yearly transaction amounts.
var ErrFailedFindYearlyAmounts = errors.NewErrorResponse("Failed to fetch yearly amounts", http.StatusInternalServerError)

// ErrFailedFindMonthlyAmountsByCard indicates a failure when retrieving monthly transaction amounts filtered by card.
var ErrFailedFindMonthlyAmountsByCard = errors.NewErrorResponse("Failed to fetch monthly amounts by card", http.StatusInternalServerError)

// ErrFailedFindYearlyAmountsByCard indicates a failure when retrieving yearly transaction amounts filtered by card.
var ErrFailedFindYearlyAmountsByCard = errors.NewErrorResponse("Failed to fetch yearly amounts by card", http.StatusInternalServerError)
