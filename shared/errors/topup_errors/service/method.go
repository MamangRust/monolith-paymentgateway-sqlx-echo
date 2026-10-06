package topupserviceerrors

import (
	"net/http"

	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

// ErrFailedFindMonthlyTopupMethods indicates failure in retrieving monthly top-up methods.
var ErrFailedFindMonthlyTopupMethods = errors.NewErrorResponse("Failed to get monthly topup methods", http.StatusInternalServerError)

// ErrFailedFindYearlyTopupMethods indicates failure in retrieving yearly top-up methods.
var ErrFailedFindYearlyTopupMethods = errors.NewErrorResponse("Failed to get yearly topup methods", http.StatusInternalServerError)

// ErrFailedFindMonthlyTopupMethodsByCard indicates failure in retrieving monthly top-up methods by card.
var ErrFailedFindMonthlyTopupMethodsByCard = errors.NewErrorResponse("Failed to get monthly topup methods by card", http.StatusInternalServerError)

// ErrFailedFindYearlyTopupMethodsByCard indicates failure in retrieving yearly top-up methods by card.
var ErrFailedFindYearlyTopupMethodsByCard = errors.NewErrorResponse("Failed to get yearly topup methods by card", http.StatusInternalServerError)
