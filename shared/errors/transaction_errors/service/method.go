package transactonserviceerrors

import (
	"net/http"

	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

// ErrFailedFindMonthlyPaymentMethods indicates a failure when retrieving monthly statistics of payment methods used.
var ErrFailedFindMonthlyPaymentMethods = errors.NewErrorResponse("Failed to fetch monthly payment methods", http.StatusInternalServerError)

// ErrFailedFindYearlyPaymentMethods indicates a failure when retrieving yearly statistics of payment methods used.
var ErrFailedFindYearlyPaymentMethods = errors.NewErrorResponse("Failed to fetch yearly payment methods", http.StatusInternalServerError)

// ErrFailedFindMonthlyPaymentMethodsByCard indicates a failure when retrieving monthly payment methods filtered by card.
var ErrFailedFindMonthlyPaymentMethodsByCard = errors.NewErrorResponse("Failed to fetch monthly payment methods by card", http.StatusInternalServerError)

// ErrFailedFindYearlyPaymentMethodsByCard indicates a failure when retrieving yearly payment methods filtered by card.
var ErrFailedFindYearlyPaymentMethodsByCard = errors.NewErrorResponse("Failed to fetch yearly payment methods by card", http.StatusInternalServerError)
