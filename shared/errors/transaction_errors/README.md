# 📦 Package `transaction_errors`

**Source Path:** `shared/errors/transaction_errors`

## 🏷️ Variables

**Var:**

ErrApiBindCreateTransaction returns an API error response
when binding the create transaction request body to struct fails.

```go
var ErrApiBindCreateTransaction = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "bind failed: invalid create transaction request", http.StatusBadRequest)
}
```

**Var:**

ErrApiBindUpdateTransaction returns an API error response
when binding the update transaction request body to struct fails.

```go
var ErrApiBindUpdateTransaction = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "bind failed: invalid update transaction request", http.StatusBadRequest)
}
```

**Var:**

ErrApiFailedCreateTransaction returns an API error response
when creating a new transaction fails.

```go
var ErrApiFailedCreateTransaction = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to create transaction", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedDeleteAllPermanent returns an API error response
when permanently deleting all trashed transactions fails.

```go
var ErrApiFailedDeleteAllPermanent = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to permanently delete all trashed transactions", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedDeletePermanent returns an API error response
when permanently deleting a transaction fails.

```go
var ErrApiFailedDeletePermanent = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to permanently delete transaction", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindActive returns an API error response
when retrieving active (non-deleted) transactions fails.

```go
var ErrApiFailedFindActive = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch active transactions", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindAllTransactions returns an API error response
when fetching all transactions from the database fails.

```go
var ErrApiFailedFindAllTransactions = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch all transactions", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindAllTransactionsActive returns an API error response
when fetching active (non-deleted) transactions fails.

```go
var ErrApiFailedFindAllTransactionsActive = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch all transactions active", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindAllTransactionsTrashed returns an API error response
when fetching trashed (soft-deleted) transactions fails.

```go
var ErrApiFailedFindAllTransactionsTrashed = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch all transactions active", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindByCardNumber returns an API error response
when fetching transactions filtered by card number fails.

```go
var ErrApiFailedFindByCardNumber = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch transactions by card number", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindById returns an API error response
when fetching a transaction by its ID fails.

```go
var ErrApiFailedFindById = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch transaction by ID", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindByMerchantID returns an API error response
when retrieving transactions by merchant ID fails.

```go
var ErrApiFailedFindByMerchantID = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch transactions by merchant ID", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindTrashed returns an API error response
when retrieving trashed (soft-deleted) transactions fails.

```go
var ErrApiFailedFindTrashed = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch trashed transactions", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedMonthlyAmounts returns an API error response
when retrieving total monthly transaction amounts fails.

```go
var ErrApiFailedMonthlyAmounts = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly amounts", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedMonthlyAmountsByCard returns an API error response
when retrieving total monthly transaction amounts by card number fails.

```go
var ErrApiFailedMonthlyAmountsByCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly amounts by card number", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedMonthlyFailed returns an API error response
when retrieving monthly failed transactions fails.

```go
var ErrApiFailedMonthlyFailed = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly failed transactions", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedMonthlyFailedByCard returns an API error response
when retrieving monthly failed transactions by card number fails.

```go
var ErrApiFailedMonthlyFailedByCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly failed transactions by card number", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedMonthlyMethods returns an API error response
when retrieving monthly payment methods fails.

```go
var ErrApiFailedMonthlyMethods = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly payment methods", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedMonthlyMethodsByCard returns an API error response
when retrieving monthly payment methods by card number fails.

```go
var ErrApiFailedMonthlyMethodsByCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly payment methods by card number", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedMonthlySuccess returns an API error response
when retrieving monthly successful transactions fails.

```go
var ErrApiFailedMonthlySuccess = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly successful transactions", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedMonthlySuccessByCard returns an API error response
when retrieving monthly successful transactions by card number fails.

```go
var ErrApiFailedMonthlySuccessByCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly successful transactions by card number", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedRestoreAllTransactions returns an API error response
when restoring all trashed transactions fails.

```go
var ErrApiFailedRestoreAllTransactions = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to restore all transactions", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedRestoreTransaction returns an API error response
when restoring a soft-deleted transaction fails.

```go
var ErrApiFailedRestoreTransaction = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to restore transaction", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedTrashTransaction returns an API error response
when moving a transaction to trash (soft-delete) fails.

```go
var ErrApiFailedTrashTransaction = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to move transaction to trash", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedUpdateTransaction returns an API error response
when updating an existing transaction fails.

```go
var ErrApiFailedUpdateTransaction = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to update transaction", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedYearlyAmounts returns an API error response
when retrieving total yearly transaction amounts fails.

```go
var ErrApiFailedYearlyAmounts = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly amounts", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedYearlyAmountsByCard returns an API error response
when retrieving total yearly transaction amounts by card number fails.

```go
var ErrApiFailedYearlyAmountsByCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly amounts by card number", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedYearlyFailed returns an API error response
when retrieving yearly failed transactions fails.

```go
var ErrApiFailedYearlyFailed = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly failed transactions", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedYearlyFailedByCard returns an API error response
when retrieving yearly failed transactions by card number fails.

```go
var ErrApiFailedYearlyFailedByCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly failed transactions by card number", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedYearlyMethods returns an API error response
when retrieving yearly payment methods fails.

```go
var ErrApiFailedYearlyMethods = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly payment methods", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedYearlyMethodsByCard returns an API error response
when retrieving yearly payment methods by card number fails.

```go
var ErrApiFailedYearlyMethodsByCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly payment methods by card number", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedYearlySuccess returns an API error response
when retrieving yearly successful transactions fails.

```go
var ErrApiFailedYearlySuccess = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly successful transactions", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedYearlySuccessByCard returns an API error response
when retrieving yearly successful transactions by card number fails.

```go
var ErrApiFailedYearlySuccessByCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly successful transactions by card number", http.StatusInternalServerError)
}
```

**Var:**

ErrApiInvalidMonth returns an API error response
when the month value is missing, invalid, or out of range.

```go
var ErrApiInvalidMonth = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Invalid month value", http.StatusBadRequest)
}
```

**Var:**

ErrApiInvalidTransactionApiKey returns an API error response
when the transaction API key is missing or invalid.

```go
var ErrApiInvalidTransactionApiKey = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Invalid transaction api-key", http.StatusBadRequest)
}
```

**Var:**

ErrApiInvalidTransactionCardNumber returns an API error response
when the provided transaction card number is invalid.

```go
var ErrApiInvalidTransactionCardNumber = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Invalid transaction Card Number", http.StatusBadRequest)
}
```

**Var:**

ErrApiInvalidTransactionID returns an API error response
when the transaction ID is invalid or improperly formatted.

```go
var ErrApiInvalidTransactionID = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Invalid transaction ID", http.StatusBadRequest)
}
```

**Var:**

ErrApiInvalidTransactionMerchantID returns an API error response
when the provided merchant ID for the transaction is invalid.

```go
var ErrApiInvalidTransactionMerchantID = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Invalid transaction merchant ID", http.StatusBadRequest)
}
```

**Var:**

ErrApiInvalidYear returns an API error response
when the year value is missing, invalid, or out of range.

```go
var ErrApiInvalidYear = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Invalid year value", http.StatusBadRequest)
}
```

**Var:**

ErrApiValidateCreateTransaction returns an API error response
when validation of the create transaction request fails.

```go
var ErrApiValidateCreateTransaction = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "validation failed: invalid create transaction request", http.StatusBadRequest)
}
```

**Var:**

ErrApiValidateUpdateTransaction returns an API error response
when validation of the update transaction request fails.

```go
var ErrApiValidateUpdateTransaction = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "validation failed: invalid update transaction request", http.StatusBadRequest)
}
```

**Var:**

ErrCreateTransactionFailed indicates a failure when creating a new transaction.

```go
var ErrCreateTransactionFailed = errors.New("failed to create transaction")
```

**Var:**

ErrDeleteAllTransactionsPermanentFailed indicates a failure when permanently deleting all transactions.

```go
var ErrDeleteAllTransactionsPermanentFailed = errors.New("failed to permanently delete all transactions")
```

**Var:**

ErrDeleteTransactionPermanentFailed indicates a failure when permanently deleting a transaction.

```go
var ErrDeleteTransactionPermanentFailed = errors.New("failed to permanently delete transaction")
```

**Var:**

ErrFailedCreateTransaction indicates a failure when creating a new transaction record.

```go
var ErrFailedCreateTransaction = response.NewErrorResponse("Failed to create transaction", http.StatusInternalServerError)
```

**Var:**

ErrFailedDeleteAllTransactionsPermanent indicates a failure when permanently deleting all transactions.

```go
var ErrFailedDeleteAllTransactionsPermanent = response.NewErrorResponse("Failed to permanently delete all transactions", http.StatusInternalServerError)
```

**Var:**

ErrFailedDeleteTransactionPermanent indicates a failure when permanently deleting a transaction.

```go
var ErrFailedDeleteTransactionPermanent = response.NewErrorResponse("Failed to permanently delete transaction", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindAllByCardNumber indicates a failure when retrieving transactions filtered by card number.

```go
var ErrFailedFindAllByCardNumber = response.NewErrorResponse("Failed to fetch transactions by card number", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindAllTransactions indicates a failure when retrieving all transaction records.

```go
var ErrFailedFindAllTransactions = response.NewErrorResponse("Failed to fetch all transactions", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindByActiveTransactions indicates a failure when retrieving active (non-deleted) transactions.

```go
var ErrFailedFindByActiveTransactions = response.NewErrorResponse("Failed to fetch active transactions", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindByMerchantID indicates a failure when retrieving transactions filtered by merchant ID.

```go
var ErrFailedFindByMerchantID = response.NewErrorResponse("Failed to fetch transactions by merchant ID", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindByTrashedTransactions indicates a failure when retrieving trashed (soft-deleted) transactions.

```go
var ErrFailedFindByTrashedTransactions = response.NewErrorResponse("Failed to fetch trashed transactions", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindMonthTransactionFailed indicates a failure when retrieving monthly failed transactions.

```go
var ErrFailedFindMonthTransactionFailed = response.NewErrorResponse("Failed to fetch monthly failed transactions", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindMonthTransactionFailedByCard indicates a failure when retrieving monthly failed transactions filtered by card number.

```go
var ErrFailedFindMonthTransactionFailedByCard = response.NewErrorResponse("Failed to fetch monthly failed transactions by card", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindMonthTransactionSuccess indicates a failure when retrieving monthly successful transactions.

```go
var ErrFailedFindMonthTransactionSuccess = response.NewErrorResponse("Failed to fetch monthly successful transactions", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindMonthTransactionSuccessByCard indicates a failure when retrieving monthly successful transactions filtered by card number.

```go
var ErrFailedFindMonthTransactionSuccessByCard = response.NewErrorResponse("Failed to fetch monthly successful transactions by card", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindMonthlyAmounts indicates a failure when retrieving the total monthly transaction amounts.

```go
var ErrFailedFindMonthlyAmounts = response.NewErrorResponse("Failed to fetch monthly amounts", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindMonthlyAmountsByCard indicates a failure when retrieving monthly transaction amounts filtered by card.

```go
var ErrFailedFindMonthlyAmountsByCard = response.NewErrorResponse("Failed to fetch monthly amounts by card", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindMonthlyPaymentMethods indicates a failure when retrieving monthly statistics of payment methods used.

```go
var ErrFailedFindMonthlyPaymentMethods = response.NewErrorResponse("Failed to fetch monthly payment methods", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindMonthlyPaymentMethodsByCard indicates a failure when retrieving monthly payment methods filtered by card.

```go
var ErrFailedFindMonthlyPaymentMethodsByCard = response.NewErrorResponse("Failed to fetch monthly payment methods by card", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindYearTransactionFailed indicates a failure when retrieving yearly failed transactions.

```go
var ErrFailedFindYearTransactionFailed = response.NewErrorResponse("Failed to fetch yearly failed transactions", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindYearTransactionFailedByCard indicates a failure when retrieving yearly failed transactions filtered by card number.

```go
var ErrFailedFindYearTransactionFailedByCard = response.NewErrorResponse("Failed to fetch yearly failed transactions by card", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindYearTransactionSuccess indicates a failure when retrieving yearly successful transactions.

```go
var ErrFailedFindYearTransactionSuccess = response.NewErrorResponse("Failed to fetch yearly successful transactions", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindYearTransactionSuccessByCard indicates a failure when retrieving yearly successful transactions filtered by card number.

```go
var ErrFailedFindYearTransactionSuccessByCard = response.NewErrorResponse("Failed to fetch yearly successful transactions by card", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindYearlyAmounts indicates a failure when retrieving the total yearly transaction amounts.

```go
var ErrFailedFindYearlyAmounts = response.NewErrorResponse("Failed to fetch yearly amounts", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindYearlyAmountsByCard indicates a failure when retrieving yearly transaction amounts filtered by card.

```go
var ErrFailedFindYearlyAmountsByCard = response.NewErrorResponse("Failed to fetch yearly amounts by card", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindYearlyPaymentMethods indicates a failure when retrieving yearly statistics of payment methods used.

```go
var ErrFailedFindYearlyPaymentMethods = response.NewErrorResponse("Failed to fetch yearly payment methods", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindYearlyPaymentMethodsByCard indicates a failure when retrieving yearly payment methods filtered by card.

```go
var ErrFailedFindYearlyPaymentMethodsByCard = response.NewErrorResponse("Failed to fetch yearly payment methods by card", http.StatusInternalServerError)
```

**Var:**

ErrFailedRestoreAllTransactions indicates a failure when restoring all trashed transactions.

```go
var ErrFailedRestoreAllTransactions = response.NewErrorResponse("Failed to restore all transactions", http.StatusInternalServerError)
```

**Var:**

ErrFailedRestoreTransaction indicates a failure when restoring a previously trashed transaction.

```go
var ErrFailedRestoreTransaction = response.NewErrorResponse("Failed to restore transaction", http.StatusInternalServerError)
```

**Var:**

ErrFailedTrashedTransaction indicates a failure when soft-deleting (trashing) a transaction.

```go
var ErrFailedTrashedTransaction = response.NewErrorResponse("Failed to trash transaction", http.StatusInternalServerError)
```

**Var:**

ErrFailedUpdateTransaction indicates a failure when updating an existing transaction record.

```go
var ErrFailedUpdateTransaction = response.NewErrorResponse("Failed to update transaction", http.StatusInternalServerError)
```

**Var:**

ErrFindActiveTransactionsFailed indicates a failure when retrieving active (non-deleted) transactions.

```go
var ErrFindActiveTransactionsFailed = errors.New("failed to find active transactions")
```

**Var:**

ErrFindAllTransactionsFailed indicates a failure when attempting to retrieve all transactions.

```go
var ErrFindAllTransactionsFailed = errors.New("failed to find all transactions")
```

**Var:**

ErrFindTransactionByIdFailed indicates a failure when retrieving a transaction by its ID.

```go
var ErrFindTransactionByIdFailed = errors.New("failed to find transaction by ID")
```

**Var:**

ErrFindTransactionByMerchantIdFailed indicates a failure when retrieving transactions by merchant ID.

```go
var ErrFindTransactionByMerchantIdFailed = errors.New("failed to find transaction by merchant ID")
```

**Var:**

ErrFindTransactionsByCardNumberFailed indicates a failure when retrieving transactions by card number.

```go
var ErrFindTransactionsByCardNumberFailed = errors.New("failed to find transactions by card number")
```

**Var:**

ErrFindTrashedTransactionsFailed indicates a failure when retrieving soft-deleted (trashed) transactions.

```go
var ErrFindTrashedTransactionsFailed = errors.New("failed to find trashed transactions")
```

**Var:**

ErrGetMonthTransactionStatusFailedByCardFailed indicates a failure when retrieving monthly failed transactions by card number.

```go
var ErrGetMonthTransactionStatusFailedByCardFailed = errors.New("failed to get monthly transaction status failed by card number")
```

**Var:**

ErrGetMonthTransactionStatusFailedFailed indicates a failure when retrieving monthly failed transaction status.

```go
var ErrGetMonthTransactionStatusFailedFailed = errors.New("failed to get monthly transaction status failed")
```

**Var:**

ErrGetMonthTransactionStatusSuccessByCardFailed indicates a failure when retrieving monthly successful transactions by card number.

```go
var ErrGetMonthTransactionStatusSuccessByCardFailed = errors.New("failed to get monthly transaction status success by card number")
```

**Var:**

ErrGetMonthTransactionStatusSuccessFailed indicates a failure when retrieving monthly successful transaction status.

```go
var ErrGetMonthTransactionStatusSuccessFailed = errors.New("failed to get monthly transaction status success")
```

**Var:**

ErrGetMonthlyAmountsByCardFailed indicates a failure when retrieving monthly amounts by card number.

```go
var ErrGetMonthlyAmountsByCardFailed = errors.New("failed to get monthly amounts by card number")
```

**Var:**

ErrGetMonthlyAmountsFailed indicates a failure when retrieving total monthly transaction amounts.

```go
var ErrGetMonthlyAmountsFailed = errors.New("failed to get monthly amounts")
```

**Var:**

ErrGetMonthlyPaymentMethodsByCardFailed indicates a failure when retrieving monthly payment methods by card number.

```go
var ErrGetMonthlyPaymentMethodsByCardFailed = errors.New("failed to get monthly payment methods by card number")
```

**Var:**

ErrGetMonthlyPaymentMethodsFailed indicates a failure when retrieving monthly payment method statistics.

```go
var ErrGetMonthlyPaymentMethodsFailed = errors.New("failed to get monthly payment methods")
```

**Var:**

ErrGetYearlyAmountsByCardFailed indicates a failure when retrieving yearly amounts by card number.

```go
var ErrGetYearlyAmountsByCardFailed = errors.New("failed to get yearly amounts by card number")
```

**Var:**

ErrGetYearlyAmountsFailed indicates a failure when retrieving total yearly transaction amounts.

```go
var ErrGetYearlyAmountsFailed = errors.New("failed to get yearly amounts")
```

**Var:**

ErrGetYearlyPaymentMethodsByCardFailed indicates a failure when retrieving yearly payment methods by card number.

```go
var ErrGetYearlyPaymentMethodsByCardFailed = errors.New("failed to get yearly payment methods by card number")
```

**Var:**

ErrGetYearlyPaymentMethodsFailed indicates a failure when retrieving yearly payment method statistics.

```go
var ErrGetYearlyPaymentMethodsFailed = errors.New("failed to get yearly payment methods")
```

**Var:**

ErrGetYearlyTransactionStatusFailedByCardFailed indicates a failure when retrieving yearly failed transactions by card number.

```go
var ErrGetYearlyTransactionStatusFailedByCardFailed = errors.New("failed to get yearly transaction status failed by card number")
```

**Var:**

ErrGetYearlyTransactionStatusFailedFailed indicates a failure when retrieving yearly failed transaction status.

```go
var ErrGetYearlyTransactionStatusFailedFailed = errors.New("failed to get yearly transaction status failed")
```

**Var:**

ErrGetYearlyTransactionStatusSuccessByCardFailed indicates a failure when retrieving yearly successful transactions by card number.

```go
var ErrGetYearlyTransactionStatusSuccessByCardFailed = errors.New("failed to get yearly transaction status success by card number")
```

**Var:**

ErrGetYearlyTransactionStatusSuccessFailed indicates a failure when retrieving yearly successful transaction status.

```go
var ErrGetYearlyTransactionStatusSuccessFailed = errors.New("failed to get yearly transaction status success")
```

**Var:**

ErrRestoreAllTransactionsFailed indicates a failure when restoring all trashed transactions.

```go
var ErrRestoreAllTransactionsFailed = errors.New("failed to restore all transactions")
```

**Var:**

ErrRestoreTransactionFailed indicates a failure when restoring a trashed transaction.

```go
var ErrRestoreTransactionFailed = errors.New("failed to restore transaction")
```

**Var:**

ErrTransactionNotFound indicates that the requested transaction could not be found.

```go
var ErrTransactionNotFound = response.NewErrorResponse("Transaction not found", http.StatusNotFound)
```

**Var:**

ErrTrashedTransactionFailed indicates a failure when soft-deleting (trashing) a transaction.

```go
var ErrTrashedTransactionFailed = errors.New("failed to soft-delete (trash) transaction")
```

**Var:**

ErrUpdateTransactionFailed indicates a failure when updating a transaction.

```go
var ErrUpdateTransactionFailed = errors.New("failed to update transaction")
```

**Var:**

ErrUpdateTransactionStatusFailed indicates a failure when updating the status of a transaction.

```go
var ErrUpdateTransactionStatusFailed = errors.New("failed to update transaction status")
```

