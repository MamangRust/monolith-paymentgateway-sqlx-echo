package withdrawserviceerrors

import (
	"net/http"

	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

// ErrFailedFindMonthlyWithdraws is used when failed to fetch monthly withdraw amounts
var ErrFailedFindMonthlyWithdraws = errors.NewErrorResponse("Failed to fetch monthly withdraw amounts", http.StatusInternalServerError)

// ErrFailedFindYearlyWithdraws is used when failed to fetch yearly withdraw amounts
var ErrFailedFindYearlyWithdraws = errors.NewErrorResponse("Failed to fetch yearly withdraw amounts", http.StatusInternalServerError)

// ErrFailedFindMonthlyWithdrawsByCardNumber is used when failed to fetch monthly withdraw amounts by card
var ErrFailedFindMonthlyWithdrawsByCardNumber = errors.NewErrorResponse("Failed to fetch monthly withdraw amounts by card", http.StatusInternalServerError)

// ErrFailedFindYearlyWithdrawsByCardNumber is used when failed to fetch yearly withdraw amounts by card
var ErrFailedFindYearlyWithdrawsByCardNumber = errors.NewErrorResponse("Failed to fetch yearly withdraw amounts by card", http.StatusInternalServerError)
