package merchantserviceerrors

import (
	"net/http"

	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

// ErrFailedFindMonthlyPaymentMethodsMerchant indicates failure fetching monthly payment methods.
var ErrFailedFindMonthlyPaymentMethodsMerchant = errors.NewErrorResponse("Failed to get monthly payment methods", http.StatusInternalServerError)

// ErrFailedFindYearlyPaymentMethodMerchant indicates failure fetching yearly payment methods.
var ErrFailedFindYearlyPaymentMethodMerchant = errors.NewErrorResponse("Failed to get yearly payment method", http.StatusInternalServerError)

// ErrFailedFindMonthlyPaymentMethodByMerchants indicates failure fetching monthly payment methods by merchant.
var ErrFailedFindMonthlyPaymentMethodByMerchants = errors.NewErrorResponse("Failed to get monthly payment methods by Merchant", http.StatusInternalServerError)

// ErrFailedFindYearlyPaymentMethodByMerchants indicates failure fetching yearly payment methods by merchant.
var ErrFailedFindYearlyPaymentMethodByMerchants = errors.NewErrorResponse("Failed to get yearly payment method by Merchant", http.StatusInternalServerError)

// ErrFailedFindMonthlyPaymentMethodByApikeys indicates failure fetching monthly payment methods by API key.
var ErrFailedFindMonthlyPaymentMethodByApikeys = errors.NewErrorResponse("Failed to get monthly payment methods by API key", http.StatusInternalServerError)

// ErrFailedFindYearlyPaymentMethodByApikeys indicates failure fetching yearly payment methods by API key.
var ErrFailedFindYearlyPaymentMethodByApikeys = errors.NewErrorResponse("Failed to get yearly payment method by API key", http.StatusInternalServerError)
