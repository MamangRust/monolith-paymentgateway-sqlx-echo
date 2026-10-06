# 📦 Package `saldo_errors`

**Source Path:** `shared/errors/saldo_errors`

## 🏷️ Variables

**Var:**

ErrApiBindCreateSaldo is returned when binding fails for saldo creation.

```go
var ErrApiBindCreateSaldo = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "bind failed: invalid create Saldo request", http.StatusBadRequest)
}
```

**Var:**

ErrApiBindUpdateSaldo is returned when binding fails for saldo update.

```go
var ErrApiBindUpdateSaldo = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "bind failed: invalid update Saldo request", http.StatusBadRequest)
}
```

```go
var ErrApiFailedCreateSaldo = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to create saldo", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedDeleteAllSaldoPermanent is returned when permanently deleting all saldos fails.

```go
var ErrApiFailedDeleteAllSaldoPermanent = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to permanently delete all saldos", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedDeleteSaldoPermanent is returned when permanently deleting saldo fails.

```go
var ErrApiFailedDeleteSaldoPermanent = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to permanently delete saldo", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindAllSaldo is returned when fetching all saldo records fails.

```go
var ErrApiFailedFindAllSaldo = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch all saldos", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindAllSaldoActive is returned when fetching active saldo records fails.

```go
var ErrApiFailedFindAllSaldoActive = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch all saldos active", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindAllSaldoTrashed is returned when fetching trashed saldo records fails.

```go
var ErrApiFailedFindAllSaldoTrashed = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch all saldos trashed", http.StatusInternalServerError)
}
```

```go
var ErrApiFailedFindByCardNumberSaldo = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch saldo by card number", http.StatusInternalServerError)
}
```

```go
var ErrApiFailedFindByIdSaldo = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch saldo by ID", http.StatusInternalServerError)
}
```

```go
var ErrApiFailedFindMonthlySaldoBalances = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly saldo balances", http.StatusInternalServerError)
}
```

```go
var ErrApiFailedFindMonthlyTotalSaldoBalance = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly total saldo balance", http.StatusInternalServerError)
}
```

```go
var ErrApiFailedFindYearTotalSaldoBalance = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly total saldo balance", http.StatusInternalServerError)
}
```

```go
var ErrApiFailedFindYearlySaldoBalances = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly saldo balances", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedRestoreAllSaldo is returned when restoring all saldos fails.

```go
var ErrApiFailedRestoreAllSaldo = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to restore all saldos", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedRestoreSaldo is returned when restoring saldo fails.

```go
var ErrApiFailedRestoreSaldo = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to restore saldo", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedTrashSaldo is returned when soft-deleting saldo fails.

```go
var ErrApiFailedTrashSaldo = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to trash saldo", http.StatusInternalServerError)
}
```

```go
var ErrApiFailedUpdateSaldo = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to update saldo", http.StatusInternalServerError)
}
```

**Var:**

ErrApiInvalidCardNumber indicates an invalid card number value in the request.

```go
var ErrApiInvalidCardNumber = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Invalid card-number value", http.StatusBadRequest)
}
```

**Var:**

ErrApiInvalidMonth indicates an invalid month value in the request.

```go
var ErrApiInvalidMonth = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Invalid month value", http.StatusBadRequest)
}
```

**Var:**

ErrApiInvalidSaldoID indicates an invalid saldo ID in the request.

```go
var ErrApiInvalidSaldoID = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Invalid saldo ID", http.StatusBadRequest)
}
```

**Var:**

ErrApiInvalidYear indicates an invalid year value in the request.

```go
var ErrApiInvalidYear = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Invalid year value", http.StatusBadRequest)
}
```

**Var:**

ErrApiValidateCreateSaldo is returned when validation fails for saldo creation.

```go
var ErrApiValidateCreateSaldo = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "validation failed: invalid create Saldo request", http.StatusBadRequest)
}
```

**Var:**

ErrApiValidateUpdateSaldo is returned when validation fails for saldo update.

```go
var ErrApiValidateUpdateSaldo = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "validation failed: invalid update Saldo request", http.StatusBadRequest)
}
```

**Var:**

ErrCreateSaldoFailed is returned when creating a saldo record fails.

```go
var ErrCreateSaldoFailed = errors.New("failed to create saldo record")
```

**Var:**

ErrDeleteAllSaldosPermanentFailed is returned when permanently deleting all saldo records fails.

```go
var ErrDeleteAllSaldosPermanentFailed = errors.New("failed to delete all saldo records permanently")
```

**Var:**

ErrDeleteSaldoPermanentFailed is returned when permanently deleting a saldo record fails.

```go
var ErrDeleteSaldoPermanentFailed = errors.New("failed to delete saldo permanently")
```

**Var:**

ErrFailedCreateSaldo returns a 500 error when creating a saldo record fails.

```go
var ErrFailedCreateSaldo = response.NewErrorResponse("Failed to create saldo", http.StatusInternalServerError)
```

**Var:**

ErrFailedDeleteAllSaldoPermanent returns a 500 error when permanently deleting all trashed saldo records fails.

```go
var ErrFailedDeleteAllSaldoPermanent = response.NewErrorResponse("Failed to permanently delete all saldos", http.StatusInternalServerError)
```

**Var:**

ErrFailedDeleteSaldoPermanent returns a 500 error when permanently deleting a saldo record fails.

```go
var ErrFailedDeleteSaldoPermanent = response.NewErrorResponse("Failed to permanently delete saldo", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindActiveSaldos returns a 500 error when fetching active saldo records fails.

```go
var ErrFailedFindActiveSaldos = response.NewErrorResponse("Failed to fetch active saldos", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindAllSaldos returns a 500 error when fetching all saldo records fails.

```go
var ErrFailedFindAllSaldos = response.NewErrorResponse("Failed to fetch saldos", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindMonthlySaldoBalances returns a 500 error when fetching monthly saldo balances fails.

```go
var ErrFailedFindMonthlySaldoBalances = response.NewErrorResponse("Failed to fetch monthly saldo balances", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindMonthlyTotalSaldoBalance returns a 500 error when fetching monthly total saldo balance fails.

```go
var ErrFailedFindMonthlyTotalSaldoBalance = response.NewErrorResponse("Failed to fetch monthly total saldo balance", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindSaldoByCardNumber returns a 500 error when fetching saldo by card number fails.

```go
var ErrFailedFindSaldoByCardNumber = response.NewErrorResponse("Failed to find saldo by card number", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindTrashedSaldos returns a 500 error when fetching trashed saldo records fails.

```go
var ErrFailedFindTrashedSaldos = response.NewErrorResponse("Failed to fetch trashed saldos", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindYearTotalSaldoBalance returns a 500 error when fetching yearly total saldo balance fails.

```go
var ErrFailedFindYearTotalSaldoBalance = response.NewErrorResponse("Failed to fetch yearly total saldo balance", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindYearlySaldoBalances returns a 500 error when fetching yearly saldo balances fails.

```go
var ErrFailedFindYearlySaldoBalances = response.NewErrorResponse("Failed to fetch yearly saldo balances", http.StatusInternalServerError)
```

**Var:**

ErrFailedInsuffientBalance returns a 400 error when a transaction is attempted with insufficient balance.

```go
var ErrFailedInsuffientBalance = response.NewErrorResponse("Insufficient balance", http.StatusBadRequest)
```

**Var:**

ErrFailedRestoreAllSaldo returns a 500 error when restoring all trashed saldo records fails.

```go
var ErrFailedRestoreAllSaldo = response.NewErrorResponse("Failed to restore all saldos", http.StatusInternalServerError)
```

**Var:**

ErrFailedRestoreSaldo returns a 500 error when restoring a trashed saldo record fails.

```go
var ErrFailedRestoreSaldo = response.NewErrorResponse("Failed to restore saldo", http.StatusInternalServerError)
```

**Var:**

ErrFailedSaldoNotFound returns a 404 error when a requested saldo record is not found.

```go
var ErrFailedSaldoNotFound = response.NewErrorResponse("Saldo not found", http.StatusNotFound)
```

**Var:**

ErrFailedTrashSaldo returns a 500 error when moving a saldo record to trash fails.

```go
var ErrFailedTrashSaldo = response.NewErrorResponse("Failed to trash saldo", http.StatusInternalServerError)
```

**Var:**

ErrFailedUpdateSaldo returns a 500 error when updating a saldo record fails.

```go
var ErrFailedUpdateSaldo = response.NewErrorResponse("Failed to update saldo", http.StatusInternalServerError)
```

**Var:**

ErrFindActiveSaldosFailed is returned when fetching active saldo records fails.

```go
var ErrFindActiveSaldosFailed = errors.New("failed to find active saldo records")
```

**Var:**

ErrFindAllSaldosFailed is returned when fetching all saldo records fails.

```go
var ErrFindAllSaldosFailed = errors.New("failed to find all saldo records")
```

**Var:**

ErrFindSaldoByCardNumberFailed is returned when fetching saldo by card number fails.

```go
var ErrFindSaldoByCardNumberFailed = errors.New("failed to find saldo by card number")
```

**Var:**

ErrFindSaldoByIdFailed is returned when fetching a saldo by its ID fails.

```go
var ErrFindSaldoByIdFailed = errors.New("failed to find saldo by ID")
```

**Var:**

ErrFindTrashedSaldosFailed is returned when fetching trashed saldo records fails.

```go
var ErrFindTrashedSaldosFailed = errors.New("failed to find trashed saldo records")
```

**Var:**

ErrGetMonthlySaldoBalancesFailed is returned when fetching monthly saldo balances fails.

```go
var ErrGetMonthlySaldoBalancesFailed = errors.New("failed to get monthly saldo balances")
```

**Var:**

ErrGetMonthlyTotalSaldoBalanceFailed is returned when fetching monthly total saldo balance fails.

```go
var ErrGetMonthlyTotalSaldoBalanceFailed = errors.New("failed to get monthly total saldo balance")
```

**Var:**

ErrGetYearTotalSaldoBalanceFailed is returned when fetching yearly total saldo balance fails.

```go
var ErrGetYearTotalSaldoBalanceFailed = errors.New("failed to get yearly total saldo balance")
```

**Var:**

ErrGetYearlySaldoBalancesFailed is returned when fetching yearly saldo balances fails.

```go
var ErrGetYearlySaldoBalancesFailed = errors.New("failed to get yearly saldo balances")
```

**Var:**

ErrRestoreAllSaldosFailed is returned when restoring all trashed saldo records fails.

```go
var ErrRestoreAllSaldosFailed = errors.New("failed to restore all saldo records")
```

**Var:**

ErrRestoreSaldoFailed is returned when restoring a trashed saldo record fails.

```go
var ErrRestoreSaldoFailed = errors.New("failed to restore saldo record")
```

**Var:**

ErrTrashSaldoFailed is returned when soft-deleting (trashing) a saldo record fails.

```go
var ErrTrashSaldoFailed = errors.New("failed to trash saldo record")
```

**Var:**

ErrUpdateSaldoBalanceFailed is returned when updating saldo balance fails.

```go
var ErrUpdateSaldoBalanceFailed = errors.New("failed to update saldo balance")
```

**Var:**

ErrUpdateSaldoFailed is returned when updating a saldo record fails.

```go
var ErrUpdateSaldoFailed = errors.New("failed to update saldo record")
```

**Var:**

ErrUpdateSaldoWithdrawFailed is returned when updating saldo for a withdrawal fails.

```go
var ErrUpdateSaldoWithdrawFailed = errors.New("failed to update saldo withdrawal")
```

