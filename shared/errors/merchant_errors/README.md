# 📦 Package `merchant_errors`

**Source Path:** `shared/errors/merchant_errors`

## 🏷️ Variables

```go
var ErrApiBindCreateMerchant = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "bind failed: invalid create merchant request", http.StatusBadRequest)
}
```

```go
var ErrApiBindUpdateMerchant = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "bind failed: invalid update merchant request", http.StatusBadRequest)
}
```

```go
var ErrApiBindUpdateMerchantStatus = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "bind failed: invalid update merchant status request", http.StatusBadRequest)
}
```

```go
var ErrApiFailedCreateMerchant = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to create merchant", http.StatusInternalServerError)
}
```

```go
var ErrApiFailedDeleteAllMerchantPermanent = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to permanently delete all merchants", http.StatusInternalServerError)
}
```

```go
var ErrApiFailedDeleteMerchantPermanent = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to permanently delete merchant", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindAllMerchants is returned when failing to retrieve all merchants.

```go
var ErrApiFailedFindAllMerchants = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch all merchants", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindAllMerchantsActive is returned when failing to retrieve all active merchants.

```go
var ErrApiFailedFindAllMerchantsActive = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch all merchants active", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindAllMerchantsTrashed is returned when failing to retrieve all trashed merchants.

```go
var ErrApiFailedFindAllMerchantsTrashed = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch all merchants trashed", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindAllTransactionByApikey is returned when failing to retrieve transactions by API key.

```go
var ErrApiFailedFindAllTransactionByApikey = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch transactions by API key", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindAllTransactionByMerchant is returned when failing to retrieve transactions by merchant.

```go
var ErrApiFailedFindAllTransactionByMerchant = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch transactions by merchant", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindAllTransactions is returned when failing to retrieve all merchant transactions.

```go
var ErrApiFailedFindAllTransactions = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch all transactions", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindByApiKeyMerchant is returned when failing to retrieve a merchant by API key.

```go
var ErrApiFailedFindByApiKeyMerchant = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch merchant by api key", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindByIdMerchant is returned when failing to retrieve a merchant by ID.

```go
var ErrApiFailedFindByIdMerchant = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch merchant by ID", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindByUserId is returned when failing to retrieve a merchant by user ID.

```go
var ErrApiFailedFindByUserId = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch merchant by user ID", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindMonthlyAmountByApikeys is returned when failing to fetch monthly amounts by API key.

```go
var ErrApiFailedFindMonthlyAmountByApikeys = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly amounts by API key", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindMonthlyAmountByMerchants is returned when failing to fetch monthly amounts for all merchants.

```go
var ErrApiFailedFindMonthlyAmountByMerchants = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly amounts by merchant", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindMonthlyAmountMerchant is returned when failing to retrieve monthly transaction amount for a merchant.

```go
var ErrApiFailedFindMonthlyAmountMerchant = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly amount by merchant", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindMonthlyPaymentMethodByApikeys is returned when failing to fetch monthly payment methods by API key.

```go
var ErrApiFailedFindMonthlyPaymentMethodByApikeys = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly payment methods by API key", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindMonthlyPaymentMethodByMerchants is returned when failing to fetch monthly payment methods for all merchants.

```go
var ErrApiFailedFindMonthlyPaymentMethodByMerchants = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly payment methods by merchant", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindMonthlyPaymentMethodsMerchant is returned when failing to retrieve monthly payment methods for a merchant.

```go
var ErrApiFailedFindMonthlyPaymentMethodsMerchant = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly payment methods by merchant", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindMonthlyTotalAmountMerchant is returned when failing to retrieve total monthly transaction amount for a merchant.

```go
var ErrApiFailedFindMonthlyTotalAmountMerchant = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly total amount by merchant", http.StatusInternalServerError)
}
```

```go
var ErrApiFailedFindYearlyAmountByApikeys = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly amounts by API key", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindYearlyAmountByMerchants is returned when failing to fetch yearly amounts for all merchants.

```go
var ErrApiFailedFindYearlyAmountByMerchants = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly amounts by merchant", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindYearlyAmountMerchant is returned when failing to retrieve yearly transaction amount for a merchant.

```go
var ErrApiFailedFindYearlyAmountMerchant = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly amount by merchant", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindYearlyPaymentMethodByApikeys is returned when failing to fetch yearly payment methods by API key.

```go
var ErrApiFailedFindYearlyPaymentMethodByApikeys = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly payment methods by API key", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindYearlyPaymentMethodByMerchants is returned when failing to fetch yearly payment methods for all merchants.

```go
var ErrApiFailedFindYearlyPaymentMethodByMerchants = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly payment methods by merchant", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindYearlyPaymentMethodMerchant is returned when failing to retrieve yearly payment methods for a merchant.

```go
var ErrApiFailedFindYearlyPaymentMethodMerchant = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly payment methods by merchant", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindYearlyTotalAmountMerchant is returned when failing to retrieve total yearly transaction amount for a merchant.

```go
var ErrApiFailedFindYearlyTotalAmountMerchant = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly total amount by merchant", http.StatusInternalServerError)
}
```

```go
var ErrApiFailedRestoreAllMerchant = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to restore all merchants", http.StatusInternalServerError)
}
```

```go
var ErrApiFailedRestoreMerchant = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to restore merchant", http.StatusInternalServerError)
}
```

```go
var ErrApiFailedTrashMerchant = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to trash merchant", http.StatusInternalServerError)
}
```

```go
var ErrApiFailedUpdateMerchant = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to update merchant", http.StatusInternalServerError)
}
```

```go
var ErrApiFailedUpdateMerchantStatus = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to update merchant status", http.StatusInternalServerError)
}
```

**Var:**

ErrApiInvalidApiKey is returned when the API key provided is invalid or unauthorized.

```go
var ErrApiInvalidApiKey = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Invalid API key", http.StatusUnauthorized)
}
```

**Var:**

ErrApiInvalidCardID is returned when the card ID provided is invalid.

```go
var ErrApiInvalidCardID = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Invalid card ID", http.StatusBadRequest)
}
```

**Var:**

ErrApiInvalidMerchantID is returned when the merchant ID provided is invalid.

```go
var ErrApiInvalidMerchantID = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Invalid merchant ID", http.StatusBadRequest)
}
```

**Var:**

ErrApiInvalidMonth is returned when the month parameter is invalid or malformed.

```go
var ErrApiInvalidMonth = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Invalid month value", http.StatusBadRequest)
}
```

**Var:**

ErrApiInvalidUserID is returned when the user ID provided is invalid.

```go
var ErrApiInvalidUserID = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Invalid user ID", http.StatusBadRequest)
}
```

**Var:**

ErrApiInvalidYear is returned when the year parameter is invalid or malformed.

```go
var ErrApiInvalidYear = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Invalid year value", http.StatusBadRequest)
}
```

```go
var ErrApiValidateCreateMerchant = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "validation failed: invalid create merchant request", http.StatusBadRequest)
}
```

```go
var ErrApiValidateUpdateMerchant = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "validation failed: invalid update merchant request", http.StatusBadRequest)
}
```

```go
var ErrApiValidateUpdateMerchantStatus = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "validation failed: invalid update merchant status request", http.StatusBadRequest)
}
```

**Var:**

ErrCreateMerchantFailed indicates failure when creating a merchant.

```go
var ErrCreateMerchantFailed = errors.New("failed to create merchant")
```

**Var:**

ErrDeleteAllMerchantPermanentFailed indicates failure when permanently deleting all merchants.

```go
var ErrDeleteAllMerchantPermanentFailed = errors.New("failed to permanently delete all merchants")
```

**Var:**

ErrDeleteMerchantPermanentFailed indicates failure when permanently deleting a merchant.

```go
var ErrDeleteMerchantPermanentFailed = errors.New("failed to permanently delete merchant")
```

**Var:**

ErrFailedCreateMerchant indicates failure when creating a new merchant.

```go
var ErrFailedCreateMerchant = response.NewErrorResponse("Failed to create Merchant", http.StatusInternalServerError)
```

**Var:**

ErrFailedDeleteAllMerchants indicates failure when permanently deleting all trashed merchants.

```go
var ErrFailedDeleteAllMerchants = response.NewErrorResponse("Failed to delete all Merchants permanently", http.StatusInternalServerError)
```

**Var:**

ErrFailedDeleteMerchant indicates failure when permanently deleting a merchant.

```go
var ErrFailedDeleteMerchant = response.NewErrorResponse("Failed to delete Merchant permanently", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindActiveMerchants indicates failure in fetching active merchants.

```go
var ErrFailedFindActiveMerchants = response.NewErrorResponse("Failed to fetch active Merchants", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindAllMerchants indicates failure in fetching all merchants.

```go
var ErrFailedFindAllMerchants = response.NewErrorResponse("Failed to fetch Merchants", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindAllTransactions is returned when fetching all transactions fails.

```go
var ErrFailedFindAllTransactions = response.NewErrorResponse("Failed to fetch Merchant transactions", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindAllTransactionsByApikey is returned when fetching transactions by API key fails.

```go
var ErrFailedFindAllTransactionsByApikey = response.NewErrorResponse("Failed to fetch transactions by API key", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindAllTransactionsByMerchant is returned when fetching transactions by merchant fails.

```go
var ErrFailedFindAllTransactionsByMerchant = response.NewErrorResponse("Failed to fetch transactions by Merchant", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindByApiKey is returned when a merchant cannot be found by API key.

```go
var ErrFailedFindByApiKey = response.NewErrorResponse("Failed to find Merchant by API key", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindByMerchantUserId is returned when a merchant cannot be found by user ID.

```go
var ErrFailedFindByMerchantUserId = response.NewErrorResponse("Failed to find Merchant by User ID", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindMerchantById is returned when a merchant cannot be found by ID.

```go
var ErrFailedFindMerchantById = response.NewErrorResponse("Failed to find Merchant by ID", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindMonthlyAmountByApikeys indicates failure fetching monthly amounts by API key.

```go
var ErrFailedFindMonthlyAmountByApikeys = response.NewErrorResponse("Failed to get monthly amount by API key", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindMonthlyAmountByMerchants indicates failure fetching monthly amounts by merchant.

```go
var ErrFailedFindMonthlyAmountByMerchants = response.NewErrorResponse("Failed to get monthly amount by Merchant", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindMonthlyAmountMerchant indicates failure fetching monthly amounts.

```go
var ErrFailedFindMonthlyAmountMerchant = response.NewErrorResponse("Failed to get monthly amount", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindMonthlyPaymentMethodByApikeys indicates failure fetching monthly payment methods by API key.

```go
var ErrFailedFindMonthlyPaymentMethodByApikeys = response.NewErrorResponse("Failed to get monthly payment methods by API key", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindMonthlyPaymentMethodByMerchants indicates failure fetching monthly payment methods by merchant.

```go
var ErrFailedFindMonthlyPaymentMethodByMerchants = response.NewErrorResponse("Failed to get monthly payment methods by Merchant", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindMonthlyPaymentMethodsMerchant indicates failure fetching monthly payment methods.

```go
var ErrFailedFindMonthlyPaymentMethodsMerchant = response.NewErrorResponse("Failed to get monthly payment methods", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindMonthlyTotalAmountByApikeys indicates failure fetching monthly total amounts by API key.

```go
var ErrFailedFindMonthlyTotalAmountByApikeys = response.NewErrorResponse("Failed to get monthly total amount by API key", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindMonthlyTotalAmountByMerchants indicates failure fetching monthly total amounts by merchant.

```go
var ErrFailedFindMonthlyTotalAmountByMerchants = response.NewErrorResponse("Failed to get monthly total amount by Merchant", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindMonthlyTotalAmountMerchant indicates failure fetching monthly total amounts.

```go
var ErrFailedFindMonthlyTotalAmountMerchant = response.NewErrorResponse("Failed to get monthly total amount", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindTrashedMerchants indicates failure in fetching trashed merchants.

```go
var ErrFailedFindTrashedMerchants = response.NewErrorResponse("Failed to fetch trashed Merchants", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindYearlyAmountByApikeys indicates failure fetching yearly amounts by API key.

```go
var ErrFailedFindYearlyAmountByApikeys = response.NewErrorResponse("Failed to get yearly amount by API key", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindYearlyAmountByMerchants indicates failure fetching yearly amounts by merchant.

```go
var ErrFailedFindYearlyAmountByMerchants = response.NewErrorResponse("Failed to get yearly amount by Merchant", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindYearlyAmountMerchant indicates failure fetching yearly amounts.

```go
var ErrFailedFindYearlyAmountMerchant = response.NewErrorResponse("Failed to get yearly amount", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindYearlyPaymentMethodByApikeys indicates failure fetching yearly payment methods by API key.

```go
var ErrFailedFindYearlyPaymentMethodByApikeys = response.NewErrorResponse("Failed to get yearly payment method by API key", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindYearlyPaymentMethodByMerchants indicates failure fetching yearly payment methods by merchant.

```go
var ErrFailedFindYearlyPaymentMethodByMerchants = response.NewErrorResponse("Failed to get yearly payment method by Merchant", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindYearlyPaymentMethodMerchant indicates failure fetching yearly payment methods.

```go
var ErrFailedFindYearlyPaymentMethodMerchant = response.NewErrorResponse("Failed to get yearly payment method", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindYearlyTotalAmountByApikeys indicates failure fetching yearly total amounts by API key.

```go
var ErrFailedFindYearlyTotalAmountByApikeys = response.NewErrorResponse("Failed to get yearly total amount by API key", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindYearlyTotalAmountByMerchants indicates failure fetching yearly total amounts by merchant.

```go
var ErrFailedFindYearlyTotalAmountByMerchants = response.NewErrorResponse("Failed to get yearly total amount by Merchant", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindYearlyTotalAmountMerchant indicates failure fetching yearly total amounts.

```go
var ErrFailedFindYearlyTotalAmountMerchant = response.NewErrorResponse("Failed to get yearly total amount", http.StatusInternalServerError)
```

**Var:**

ErrFailedRestoreAllMerchants indicates failure when restoring all trashed merchants.

```go
var ErrFailedRestoreAllMerchants = response.NewErrorResponse("Failed to restore all Merchants", http.StatusInternalServerError)
```

**Var:**

ErrFailedRestoreMerchant indicates failure when restoring a trashed merchant.

```go
var ErrFailedRestoreMerchant = response.NewErrorResponse("Failed to restore Merchant", http.StatusInternalServerError)
```

**Var:**

ErrFailedSendEmail indicates a failure when sending an email related to merchant operations.

```go
var ErrFailedSendEmail = response.NewErrorResponse("Failed to send email", http.StatusInternalServerError)
```

**Var:**

ErrFailedTrashMerchant indicates failure when soft-deleting (trashing) a merchant.

```go
var ErrFailedTrashMerchant = response.NewErrorResponse("Failed to trash Merchant", http.StatusInternalServerError)
```

**Var:**

ErrFailedUpdateMerchant indicates failure when updating a merchant.

```go
var ErrFailedUpdateMerchant = response.NewErrorResponse("Failed to update Merchant", http.StatusInternalServerError)
```

**Var:**

ErrFindActiveMerchantsFailed is returned when fetching active merchants fails.

```go
var ErrFindActiveMerchantsFailed = errors.New("failed to find active merchants")
```

**Var:**

ErrFindAllMerchantsFailed is returned when fetching all merchants fails.

```go
var ErrFindAllMerchantsFailed = errors.New("failed to find all merchants")
```

**Var:**

ErrFindAllTransactionsByApiKeyFailed indicates failure fetching transactions by API key.

```go
var ErrFindAllTransactionsByApiKeyFailed = errors.New("failed to find merchant transactions by API key")
```

**Var:**

ErrFindAllTransactionsByMerchantFailed indicates failure fetching transactions by merchant ID.

```go
var ErrFindAllTransactionsByMerchantFailed = errors.New("failed to find merchant transactions by merchant ID")
```

**Var:**

ErrFindAllTransactionsFailed indicates failure fetching all transactions.

```go
var ErrFindAllTransactionsFailed = errors.New("failed to find all merchant transactions")
```

**Var:**

ErrFindMerchantByApiKeyFailed is returned when a merchant cannot be found by API key.

```go
var ErrFindMerchantByApiKeyFailed = errors.New("failed to find merchant by API key")
```

**Var:**

ErrFindMerchantByIdFailed is returned when a merchant cannot be found by ID.

```go
var ErrFindMerchantByIdFailed = errors.New("failed to find merchant by ID")
```

**Var:**

ErrFindMerchantByNameFailed is returned when a merchant cannot be found by name.

```go
var ErrFindMerchantByNameFailed = errors.New("failed to find merchant by name")
```

**Var:**

ErrFindMerchantByUserIdFailed is returned when a merchant cannot be found by user ID.

```go
var ErrFindMerchantByUserIdFailed = errors.New("failed to find merchant by user ID")
```

**Var:**

ErrFindTrashedMerchantsFailed is returned when fetching trashed merchants fails.

```go
var ErrFindTrashedMerchantsFailed = errors.New("failed to find trashed merchants")
```

**Var:**

ErrGetMonthlyAmountByApikeyFailed indicates failure fetching monthly amount by API key.

```go
var ErrGetMonthlyAmountByApikeyFailed = errors.New("failed to get monthly amount by API key")
```

**Var:**

ErrGetMonthlyAmountByMerchantsFailed indicates failure fetching monthly amount for all merchants.

```go
var ErrGetMonthlyAmountByMerchantsFailed = errors.New("failed to get monthly amount by merchants")
```

**Var:**

ErrGetMonthlyAmountMerchantFailed indicates failure fetching monthly amount for a merchant.

```go
var ErrGetMonthlyAmountMerchantFailed = errors.New("failed to get monthly amount of merchant")
```

**Var:**

ErrGetMonthlyPaymentMethodByApikeyFailed indicates failure fetching monthly payment methods by API key.

```go
var ErrGetMonthlyPaymentMethodByApikeyFailed = errors.New("failed to get monthly payment method by API key")
```

**Var:**

ErrGetMonthlyPaymentMethodByMerchantsFailed indicates failure fetching monthly payment methods for all merchants.

```go
var ErrGetMonthlyPaymentMethodByMerchantsFailed = errors.New("failed to get monthly payment method by merchants")
```

**Var:**

ErrGetMonthlyPaymentMethodsMerchantFailed indicates failure fetching monthly payment methods for a merchant.

```go
var ErrGetMonthlyPaymentMethodsMerchantFailed = errors.New("failed to get monthly payment methods of merchant")
```

**Var:**

ErrGetMonthlyTotalAmountByApikeyFailed indicates failure fetching monthly total amount by API key.

```go
var ErrGetMonthlyTotalAmountByApikeyFailed = errors.New("failed to get monthly total amount by API key")
```

**Var:**

ErrGetMonthlyTotalAmountByMerchantsFailed indicates failure fetching monthly total amount for all merchants.

```go
var ErrGetMonthlyTotalAmountByMerchantsFailed = errors.New("failed to get monthly total amount by merchants")
```

**Var:**

ErrGetMonthlyTotalAmountMerchantFailed indicates failure fetching monthly total amount for a merchant.

```go
var ErrGetMonthlyTotalAmountMerchantFailed = errors.New("failed to get monthly total amount of merchant")
```

**Var:**

ErrGetYearlyAmountByApikeyFailed indicates failure fetching yearly amount by API key.

```go
var ErrGetYearlyAmountByApikeyFailed = errors.New("failed to get yearly amount by API key")
```

**Var:**

ErrGetYearlyAmountByMerchantsFailed indicates failure fetching yearly amount for all merchants.

```go
var ErrGetYearlyAmountByMerchantsFailed = errors.New("failed to get yearly amount by merchants")
```

**Var:**

ErrGetYearlyAmountMerchantFailed indicates failure fetching yearly amount for a merchant.

```go
var ErrGetYearlyAmountMerchantFailed = errors.New("failed to get yearly amount of merchant")
```

**Var:**

ErrGetYearlyPaymentMethodByApikeyFailed indicates failure fetching yearly payment methods by API key.

```go
var ErrGetYearlyPaymentMethodByApikeyFailed = errors.New("failed to get yearly payment method by API key")
```

**Var:**

ErrGetYearlyPaymentMethodByMerchantsFailed indicates failure fetching yearly payment methods for all merchants.

```go
var ErrGetYearlyPaymentMethodByMerchantsFailed = errors.New("failed to get yearly payment method by merchants")
```

**Var:**

ErrGetYearlyPaymentMethodMerchantFailed indicates failure fetching yearly payment methods for a merchant.

```go
var ErrGetYearlyPaymentMethodMerchantFailed = errors.New("failed to get yearly payment method of merchant")
```

**Var:**

ErrGetYearlyTotalAmountByApikeyFailed indicates failure fetching yearly total amount by API key.

```go
var ErrGetYearlyTotalAmountByApikeyFailed = errors.New("failed to get yearly total amount by API key")
```

**Var:**

ErrGetYearlyTotalAmountByMerchantsFailed indicates failure fetching yearly total amount for all merchants.

```go
var ErrGetYearlyTotalAmountByMerchantsFailed = errors.New("failed to get yearly total amount by merchants")
```

**Var:**

ErrGetYearlyTotalAmountMerchantFailed indicates failure fetching yearly total amount for a merchant.

```go
var ErrGetYearlyTotalAmountMerchantFailed = errors.New("failed to get yearly total amount of merchant")
```

**Var:**

ErrMerchantNotFoundRes is returned when the merchant is not found.

```go
var ErrMerchantNotFoundRes = response.NewErrorResponse("Merchant not found", http.StatusNotFound)
```

**Var:**

ErrRestoreAllMerchantFailed indicates failure when restoring all trashed merchants.

```go
var ErrRestoreAllMerchantFailed = errors.New("failed to restore all merchants")
```

**Var:**

ErrRestoreMerchantFailed indicates failure when restoring a trashed merchant.

```go
var ErrRestoreMerchantFailed = errors.New("failed to restore merchant")
```

**Var:**

ErrTrashedMerchantFailed indicates failure when soft-deleting a merchant.

```go
var ErrTrashedMerchantFailed = errors.New("failed to soft-delete (trash) merchant")
```

**Var:**

ErrUpdateMerchantFailed indicates failure when updating a merchant.

```go
var ErrUpdateMerchantFailed = errors.New("failed to update merchant")
```

**Var:**

ErrUpdateMerchantStatusFailed indicates failure when updating merchant status.

```go
var ErrUpdateMerchantStatusFailed = errors.New("failed to update merchant status")
```

