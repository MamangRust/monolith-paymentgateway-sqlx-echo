package saldoserviceerror

import (
	"net/http"

	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

// ErrFailedFindMonthlyTotalSaldoBalance returns a 500 error when fetching monthly total saldo balance fails.
var ErrFailedFindMonthlyTotalSaldoBalance = errors.NewErrorResponse("Failed to fetch monthly total saldo balance", http.StatusInternalServerError)

// ErrFailedFindYearTotalSaldoBalance returns a 500 error when fetching yearly total saldo balance fails.
var ErrFailedFindYearTotalSaldoBalance = errors.NewErrorResponse("Failed to fetch yearly total saldo balance", http.StatusInternalServerError)
