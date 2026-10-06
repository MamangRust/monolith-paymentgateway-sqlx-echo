package merchantserviceerrors

import (
	"net/http"

	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

// ErrFailedFindMonthlyAmountMerchant indicates failure fetching monthly amounts.
var ErrFailedFindMonthlyAmountMerchant = errors.NewErrorResponse("Failed to get monthly amount", http.StatusInternalServerError)

// ErrFailedFindYearlyAmountMerchant indicates failure fetching yearly amounts.
var ErrFailedFindYearlyAmountMerchant = errors.NewErrorResponse("Failed to get yearly amount", http.StatusInternalServerError)

// ErrFailedFindMonthlyAmountByMerchants indicates failure fetching monthly amounts by merchant.
var ErrFailedFindMonthlyAmountByMerchants = errors.NewErrorResponse("Failed to get monthly amount by Merchant", http.StatusInternalServerError)

// ErrFailedFindYearlyAmountByMerchants indicates failure fetching yearly amounts by merchant.
var ErrFailedFindYearlyAmountByMerchants = errors.NewErrorResponse("Failed to get yearly amount by Merchant", http.StatusInternalServerError)

// ErrFailedFindMonthlyAmountByApikeys indicates failure fetching monthly amounts by API key.
var ErrFailedFindMonthlyAmountByApikeys = errors.NewErrorResponse("Failed to get monthly amount by API key", http.StatusInternalServerError)

// ErrFailedFindYearlyAmountByApikeys indicates failure fetching yearly amounts by API key.
var ErrFailedFindYearlyAmountByApikeys = errors.NewErrorResponse("Failed to get yearly amount by API key", http.StatusInternalServerError)
