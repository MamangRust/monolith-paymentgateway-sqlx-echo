package transferserviceerrors

import (
	"net/http"

	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

// ErrFailedFindMonthlyTransferAmounts indicates a failure when retrieving the total monthly transfer amounts.
var ErrFailedFindMonthlyTransferAmounts = errors.NewErrorResponse("Failed to fetch monthly transfer amounts", http.StatusInternalServerError)

// ErrFailedFindYearlyTransferAmounts indicates a failure when retrieving the total yearly transfer amounts.
var ErrFailedFindYearlyTransferAmounts = errors.NewErrorResponse("Failed to fetch yearly transfer amounts", http.StatusInternalServerError)

// ErrFailedFindMonthlyTransferAmountsBySenderCard indicates a failure when retrieving monthly transfer amounts by sender card number.
var ErrFailedFindMonthlyTransferAmountsBySenderCard = errors.NewErrorResponse("Failed to fetch monthly transfer amounts by sender card", http.StatusInternalServerError)

// ErrFailedFindMonthlyTransferAmountsByReceiverCard indicates a failure when retrieving monthly transfer amounts by receiver card number.
var ErrFailedFindMonthlyTransferAmountsByReceiverCard = errors.NewErrorResponse("Failed to fetch monthly transfer amounts by receiver card", http.StatusInternalServerError)

// ErrFailedFindYearlyTransferAmountsBySenderCard indicates a failure when retrieving yearly transfer amounts by sender card number.
var ErrFailedFindYearlyTransferAmountsBySenderCard = errors.NewErrorResponse("Failed to fetch yearly transfer amounts by sender card", http.StatusInternalServerError)

// ErrFailedFindYearlyTransferAmountsByReceiverCard indicates a failure when retrieving yearly transfer amounts by receiver card number.
var ErrFailedFindYearlyTransferAmountsByReceiverCard = errors.NewErrorResponse("Failed to fetch yearly transfer amounts by receiver card", http.StatusInternalServerError)
