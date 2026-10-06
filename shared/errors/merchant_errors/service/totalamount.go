package merchantserviceerrors

import (
	"net/http"

	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

// ErrFailedFindMonthlyTotalAmountMerchant indicates failure fetching monthly total amounts.
var ErrFailedFindMonthlyTotalAmountMerchant = errors.NewErrorResponse("Failed to get monthly total amount", http.StatusInternalServerError)

// ErrFailedFindYearlyTotalAmountMerchant indicates failure fetching yearly total amounts.
var ErrFailedFindYearlyTotalAmountMerchant = errors.NewErrorResponse("Failed to get yearly total amount", http.StatusInternalServerError)

// ErrFailedFindMonthlyTotalAmountByMerchants indicates failure fetching monthly total amounts by merchant.
var ErrFailedFindMonthlyTotalAmountByMerchants = errors.NewErrorResponse("Failed to get monthly total amount by Merchant", http.StatusInternalServerError)

// ErrFailedFindYearlyTotalAmountByMerchants indicates failure fetching yearly total amounts by merchant.
var ErrFailedFindYearlyTotalAmountByMerchants = errors.NewErrorResponse("Failed to get yearly total amount by Merchant", http.StatusInternalServerError)

// ErrFailedFindMonthlyTotalAmountByApikeys indicates failure fetching monthly total amounts by API key.
var ErrFailedFindMonthlyTotalAmountByApikeys = errors.NewErrorResponse("Failed to get monthly total amount by API key", http.StatusInternalServerError)

// ErrFailedFindYearlyTotalAmountByApikeys indicates failure fetching yearly total amounts by API key.
var ErrFailedFindYearlyTotalAmountByApikeys = errors.NewErrorResponse("Failed to get yearly total amount by API key", http.StatusInternalServerError)
