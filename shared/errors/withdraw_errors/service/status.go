package withdrawserviceerrors

import (
	"net/http"

	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

// ErrFailedFindMonthWithdrawStatusSuccess is used when failed to fetch monthly successful withdraws
var ErrFailedFindMonthWithdrawStatusSuccess = errors.NewErrorResponse("Failed to fetch monthly successful withdraws", http.StatusInternalServerError)

// ErrFailedFindYearWithdrawStatusSuccess is used when failed to fetch yearly successful withdraws
var ErrFailedFindYearWithdrawStatusSuccess = errors.NewErrorResponse("Failed to fetch yearly successful withdraws", http.StatusInternalServerError)

// ErrFailedFindMonthWithdrawStatusFailed is used when failed to fetch monthly failed withdraws
var ErrFailedFindMonthWithdrawStatusFailed = errors.NewErrorResponse("Failed to fetch monthly failed withdraws", http.StatusInternalServerError)

// ErrFailedFindYearWithdrawStatusFailed is used when failed to fetch yearly failed withdraws
var ErrFailedFindYearWithdrawStatusFailed = errors.NewErrorResponse("Failed to fetch yearly failed withdraws", http.StatusInternalServerError)

// ErrFailedFindMonthWithdrawStatusSuccessByCard is used when failed to fetch monthly successful withdraws by card
var ErrFailedFindMonthWithdrawStatusSuccessByCard = errors.NewErrorResponse("Failed to fetch monthly successful withdraws by card", http.StatusInternalServerError)

// ErrFailedFindYearWithdrawStatusSuccessByCard is used when failed to fetch yearly successful withdraws by card
var ErrFailedFindYearWithdrawStatusSuccessByCard = errors.NewErrorResponse("Failed to fetch yearly successful withdraws by card", http.StatusInternalServerError)

// ErrFailedFindMonthWithdrawStatusFailedByCard is used when failed to fetch monthly failed withdraws by card
var ErrFailedFindMonthWithdrawStatusFailedByCard = errors.NewErrorResponse("Failed to fetch monthly failed withdraws by card", http.StatusInternalServerError)

// ErrFailedFindYearWithdrawStatusFailedByCard is used when failed to fetch yearly failed withdraws by card
var ErrFailedFindYearWithdrawStatusFailedByCard = errors.NewErrorResponse("Failed to fetch yearly failed withdraws by card", http.StatusInternalServerError)
