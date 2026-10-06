package saldoserviceerror

import (
	"net/http"

	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
)

// ErrFailedFindMonthlySaldoBalances returns a 500 error when fetching monthly saldo balances fails.
var ErrFailedFindMonthlySaldoBalances = errors.NewErrorResponse("Failed to fetch monthly saldo balances", http.StatusInternalServerError)

// ErrFailedFindYearlySaldoBalances returns a 500 error when fetching yearly saldo balances fails.
var ErrFailedFindYearlySaldoBalances = errors.NewErrorResponse("Failed to fetch yearly saldo balances", http.StatusInternalServerError)
