# 📦 Package `topup_errors`

**Source Path:** `./shared/errors/topup_errors`

## 🏷️ Variables

**Var:**

ErrApiBindCreateTopup returns a 400 error when binding the create topup request fails.

```go
var ErrApiBindCreateTopup = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "bind failed: invalid create topup request", http.StatusBadRequest)
}
```

**Var:**

ErrApiBindUpdateTopup returns a 400 error when binding the update topup request fails.

```go
var ErrApiBindUpdateTopup = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "bind failed: invalid update topup request", http.StatusBadRequest)
}
```

**Var:**

ErrApiFailedCreateTopup returns a 500 error when creating a new topup fails.

```go
var ErrApiFailedCreateTopup = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to create topup", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedDeleteAllTopupPermanent returns a 500 error when permanently deleting all topups fails.

```go
var ErrApiFailedDeleteAllTopupPermanent = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to permanently delete all topups", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedDeletePermanentTopup returns a 500 error when permanently deleting a topup fails.

```go
var ErrApiFailedDeletePermanentTopup = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to permanently delete topup", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindAllByCardNumberTopup returns a 500 error when retrieving top-ups by card number fails.

```go
var ErrApiFailedFindAllByCardNumberTopup = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch topups by card number", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindAllTopups returns a 500 Internal Server Error when fetching all top-up records fails.

```go
var ErrApiFailedFindAllTopups = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch all topups", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindAllTopupsActive returns a 500 error when retrieving all active (non-deleted) top-ups fails.

```go
var ErrApiFailedFindAllTopupsActive = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch all topups active", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindAllTopupsTrashed returns a 500 error when retrieving all trashed (soft-deleted) top-ups fails.

```go
var ErrApiFailedFindAllTopupsTrashed = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch all topups trashed", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindByIdTopup returns a 500 error when retrieving a specific top-up by ID fails.

```go
var ErrApiFailedFindByIdTopup = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch topup by ID", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindMonthlyTopupAmounts returns a 500 error when fetching monthly topup amounts fails.

```go
var ErrApiFailedFindMonthlyTopupAmounts = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly topup amounts", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindMonthlyTopupAmountsByCardNumber returns a 500 error when fetching monthly topup amounts by card number fails.

```go
var ErrApiFailedFindMonthlyTopupAmountsByCardNumber = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly topup amounts by card number", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindMonthlyTopupMethods returns a 500 error when fetching monthly topup methods fails.

```go
var ErrApiFailedFindMonthlyTopupMethods = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly topup methods", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindMonthlyTopupMethodsByCardNumber returns a 500 error when fetching monthly topup methods by card number fails.

```go
var ErrApiFailedFindMonthlyTopupMethodsByCardNumber = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly topup methods by card number", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindMonthlyTopupStatusFailed returns a 500 error when fetching monthly failed top-up stats fails.

```go
var ErrApiFailedFindMonthlyTopupStatusFailed = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly failed topups", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindMonthlyTopupStatusFailedByCardNumber returns a 500 error when fetching monthly failed topups by card number fails.

```go
var ErrApiFailedFindMonthlyTopupStatusFailedByCardNumber = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly failed topups by card number", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindMonthlyTopupStatusSuccess returns a 500 error when fetching monthly successful top-up stats fails.

```go
var ErrApiFailedFindMonthlyTopupStatusSuccess = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly successful topups", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindMonthlyTopupStatusSuccessByCardNumber returns a 500 error when fetching monthly successful top-ups by card number fails.

```go
var ErrApiFailedFindMonthlyTopupStatusSuccessByCardNumber = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly successful topups by card number", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindYearlyTopupAmounts returns a 500 error when fetching yearly topup amounts fails.

```go
var ErrApiFailedFindYearlyTopupAmounts = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly topup amounts", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindYearlyTopupAmountsByCardNumber returns a 500 error when fetching yearly topup amounts by card number fails.

```go
var ErrApiFailedFindYearlyTopupAmountsByCardNumber = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly topup amounts by card number", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindYearlyTopupMethods returns a 500 error when fetching yearly topup methods fails.

```go
var ErrApiFailedFindYearlyTopupMethods = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly topup methods", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindYearlyTopupMethodsByCardNumber returns a 500 error when fetching yearly topup methods by card number fails.

```go
var ErrApiFailedFindYearlyTopupMethodsByCardNumber = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly topup methods by card number", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindYearlyTopupStatusFailed returns a 500 error when fetching yearly failed top-up stats fails.

```go
var ErrApiFailedFindYearlyTopupStatusFailed = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly failed topups", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindYearlyTopupStatusFailedByCardNumber returns a 500 error when fetching yearly failed topups by card number fails.

```go
var ErrApiFailedFindYearlyTopupStatusFailedByCardNumber = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly failed topups by card number", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindYearlyTopupStatusSuccess returns a 500 error when fetching yearly successful top-up stats fails.

```go
var ErrApiFailedFindYearlyTopupStatusSuccess = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly successful topups", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindYearlyTopupStatusSuccessByCardNumber returns a 500 error when fetching yearly successful top-ups by card number fails.

```go
var ErrApiFailedFindYearlyTopupStatusSuccessByCardNumber = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly successful topups by card number", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedRestoreAllTopup returns a 500 error when restoring all trashed topups fails.

```go
var ErrApiFailedRestoreAllTopup = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to restore all topups", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedRestoreTopup returns a 500 error when restoring a trashed topup fails.

```go
var ErrApiFailedRestoreTopup = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to restore topup", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedTrashTopup returns a 500 error when moving a topup to trash fails.

```go
var ErrApiFailedTrashTopup = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to trash topup", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedUpdateTopup returns a 500 error when updating a topup fails.

```go
var ErrApiFailedUpdateTopup = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to update topup", http.StatusInternalServerError)
}
```

**Var:**

ErrApiInvalidCardNumber returns a 400 Bad Request error when the provided card number is invalid.

```go
var ErrApiInvalidCardNumber = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Invalid card number", http.StatusBadRequest)
}
```

**Var:**

ErrApiInvalidMonth returns a 400 Bad Request error when the provided month value is invalid or missing.

```go
var ErrApiInvalidMonth = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Invalid month value", http.StatusBadRequest)
}
```

**Var:**

ErrApiInvalidTopupID returns a 400 Bad Request error when the provided top-up ID is invalid.

```go
var ErrApiInvalidTopupID = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Invalid topup ID", http.StatusBadRequest)
}
```

**Var:**

ErrApiInvalidYear returns a 400 Bad Request error when the provided year value is invalid or missing.

```go
var ErrApiInvalidYear = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Invalid year value", http.StatusBadRequest)
}
```

**Var:**

ErrApiValidateCreateTopup returns a 400 error when the create topup request fails validation.

```go
var ErrApiValidateCreateTopup = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "validation failed: invalid create topup request", http.StatusBadRequest)
}
```

**Var:**

ErrApiValidateUpdateTopup returns a 400 error when the update topup request fails validation.

```go
var ErrApiValidateUpdateTopup = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "validation failed: invalid update topup request", http.StatusBadRequest)
}
```

**Var:**

ErrCreateTopupFailed indicates failure in creating a new top-up record.

```go
var ErrCreateTopupFailed = errors.New("failed to create topup")
```

**Var:**

ErrDeleteAllTopupPermanentFailed indicates failure in permanently deleting all trashed top-ups.

```go
var ErrDeleteAllTopupPermanentFailed = errors.New("failed to permanently delete all topups")
```

**Var:**

ErrDeleteTopupPermanentFailed indicates failure in permanently deleting a top-up.

```go
var ErrDeleteTopupPermanentFailed = errors.New("failed to permanently delete topup")
```

**Var:**

ErrFailedCreateTopup indicates failure in creating a new top-up record.

```go
var ErrFailedCreateTopup = response.NewErrorResponse("Failed to create Topup", http.StatusInternalServerError)
```

**Var:**

ErrFailedDeleteAllTopups indicates failure in permanently deleting all trashed top-up records.

```go
var ErrFailedDeleteAllTopups = response.NewErrorResponse("Failed to delete all Topups permanently", http.StatusInternalServerError)
```

**Var:**

ErrFailedDeleteTopup indicates failure in permanently deleting a top-up.

```go
var ErrFailedDeleteTopup = response.NewErrorResponse("Failed to delete Topup permanently", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindActiveTopups indicates failure in retrieving active (non-trashed) top-up records.

```go
var ErrFailedFindActiveTopups = response.NewErrorResponse("Failed to fetch active Topups", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindAllTopups indicates failure in retrieving all top-up records.

```go
var ErrFailedFindAllTopups = response.NewErrorResponse("Failed to fetch Topups", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindAllTopupsByCardNumber indicates failure in retrieving top-ups by card number.

```go
var ErrFailedFindAllTopupsByCardNumber = response.NewErrorResponse("Failed to fetch Topups by card number", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindMonthTopupStatusFailed indicates failure in retrieving monthly failed top-up status.

```go
var ErrFailedFindMonthTopupStatusFailed = response.NewErrorResponse("Failed to get monthly topup failed status", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindMonthTopupStatusFailedByCard indicates failure in retrieving monthly failed top-up status by card.

```go
var ErrFailedFindMonthTopupStatusFailedByCard = response.NewErrorResponse("Failed to get monthly topup failed status by card", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindMonthTopupStatusSuccess indicates failure in retrieving monthly successful top-up status.

```go
var ErrFailedFindMonthTopupStatusSuccess = response.NewErrorResponse("Failed to get monthly topup success status", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindMonthTopupStatusSuccessByCard indicates failure in retrieving monthly successful top-up status by card.

```go
var ErrFailedFindMonthTopupStatusSuccessByCard = response.NewErrorResponse("Failed to get monthly topup success status by card", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindMonthlyTopupAmounts indicates failure in retrieving monthly top-up amounts.

```go
var ErrFailedFindMonthlyTopupAmounts = response.NewErrorResponse("Failed to get monthly topup amounts", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindMonthlyTopupAmountsByCard indicates failure in retrieving monthly top-up amounts by card.

```go
var ErrFailedFindMonthlyTopupAmountsByCard = response.NewErrorResponse("Failed to get monthly topup amounts by card", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindMonthlyTopupMethods indicates failure in retrieving monthly top-up methods.

```go
var ErrFailedFindMonthlyTopupMethods = response.NewErrorResponse("Failed to get monthly topup methods", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindMonthlyTopupMethodsByCard indicates failure in retrieving monthly top-up methods by card.

```go
var ErrFailedFindMonthlyTopupMethodsByCard = response.NewErrorResponse("Failed to get monthly topup methods by card", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindTopupById indicates failure in finding a top-up by its ID.

```go
var ErrFailedFindTopupById = response.NewErrorResponse("Failed to find Topup by ID", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindTrashedTopups indicates failure in retrieving trashed (soft-deleted) top-up records.

```go
var ErrFailedFindTrashedTopups = response.NewErrorResponse("Failed to fetch trashed Topups", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindYearlyTopupAmounts indicates failure in retrieving yearly top-up amounts.

```go
var ErrFailedFindYearlyTopupAmounts = response.NewErrorResponse("Failed to get yearly topup amounts", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindYearlyTopupAmountsByCard indicates failure in retrieving yearly top-up amounts by card.

```go
var ErrFailedFindYearlyTopupAmountsByCard = response.NewErrorResponse("Failed to get yearly topup amounts by card", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindYearlyTopupMethods indicates failure in retrieving yearly top-up methods.

```go
var ErrFailedFindYearlyTopupMethods = response.NewErrorResponse("Failed to get yearly topup methods", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindYearlyTopupMethodsByCard indicates failure in retrieving yearly top-up methods by card.

```go
var ErrFailedFindYearlyTopupMethodsByCard = response.NewErrorResponse("Failed to get yearly topup methods by card", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindYearlyTopupStatusFailed indicates failure in retrieving yearly failed top-up status.

```go
var ErrFailedFindYearlyTopupStatusFailed = response.NewErrorResponse("Failed to get yearly topup failed status", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindYearlyTopupStatusFailedByCard indicates failure in retrieving yearly failed top-up status by card.

```go
var ErrFailedFindYearlyTopupStatusFailedByCard = response.NewErrorResponse("Failed to get yearly topup failed status by card", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindYearlyTopupStatusSuccess indicates failure in retrieving yearly successful top-up status.

```go
var ErrFailedFindYearlyTopupStatusSuccess = response.NewErrorResponse("Failed to get yearly topup success status", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindYearlyTopupStatusSuccessByCard indicates failure in retrieving yearly successful top-up status by card.

```go
var ErrFailedFindYearlyTopupStatusSuccessByCard = response.NewErrorResponse("Failed to get yearly topup success status by card", http.StatusInternalServerError)
```

**Var:**

ErrFailedRestoreAllTopups indicates failure in restoring all trashed top-up records.

```go
var ErrFailedRestoreAllTopups = response.NewErrorResponse("Failed to restore all Topups", http.StatusInternalServerError)
```

**Var:**

ErrFailedRestoreTopup indicates failure in restoring a previously trashed top-up.

```go
var ErrFailedRestoreTopup = response.NewErrorResponse("Failed to restore Topup", http.StatusInternalServerError)
```

**Var:**

ErrFailedTrashTopup indicates failure in soft-deleting (trashing) a top-up.

```go
var ErrFailedTrashTopup = response.NewErrorResponse("Failed to trash Topup", http.StatusInternalServerError)
```

**Var:**

ErrFailedUpdateTopup indicates failure in updating an existing top-up record.

```go
var ErrFailedUpdateTopup = response.NewErrorResponse("Failed to update Topup", http.StatusInternalServerError)
```

**Var:**

ErrFindAllTopupsFailed indicates failure in retrieving all top-up records from the database.

```go
var ErrFindAllTopupsFailed = errors.New("failed to find all topups")
```

**Var:**

ErrFindTopupByIdFailed indicates failure in finding a top-up using its unique ID.

```go
var ErrFindTopupByIdFailed = errors.New("failed to find topup by ID")
```

**Var:**

ErrFindTopupsByActiveFailed indicates failure in retrieving only the active (non-deleted) top-ups.

```go
var ErrFindTopupsByActiveFailed = errors.New("failed to find active topups")
```

**Var:**

ErrFindTopupsByCardNumberFailed indicates failure in finding top-ups by a specific card number.

```go
var ErrFindTopupsByCardNumberFailed = errors.New("failed to find topups by card number")
```

**Var:**

ErrFindTopupsByTrashedFailed indicates failure in retrieving trashed (soft-deleted) top-ups.

```go
var ErrFindTopupsByTrashedFailed = errors.New("failed to find trashed topups")
```

**Var:**

ErrGetMonthTopupStatusFailedByCardFailed indicates failure in getting monthly failed top-up status by card number.

```go
var ErrGetMonthTopupStatusFailedByCardFailed = errors.New("failed to get monthly topup status failed by card number")
```

**Var:**

ErrGetMonthTopupStatusFailedFailed indicates failure in getting the monthly count of failed top-ups.

```go
var ErrGetMonthTopupStatusFailedFailed = errors.New("failed to get monthly topup status failed")
```

**Var:**

ErrGetMonthTopupStatusSuccessByCardFailed indicates failure in getting monthly successful top-up status by card number.

```go
var ErrGetMonthTopupStatusSuccessByCardFailed = errors.New("failed to get monthly topup status success by card number")
```

**Var:**

ErrGetMonthTopupStatusSuccessFailed indicates failure in getting the monthly count of successful top-ups.

```go
var ErrGetMonthTopupStatusSuccessFailed = errors.New("failed to get monthly topup status success")
```

**Var:**

ErrGetMonthlyTopupAmountsByCardFailed indicates failure in retrieving monthly top-up amount stats by card number.

```go
var ErrGetMonthlyTopupAmountsByCardFailed = errors.New("failed to get monthly topup amounts by card number")
```

**Var:**

ErrGetMonthlyTopupAmountsFailed indicates failure in retrieving the monthly top-up amounts.

```go
var ErrGetMonthlyTopupAmountsFailed = errors.New("failed to get monthly topup amounts")
```

**Var:**

ErrGetMonthlyTopupMethodsByCardFailed indicates failure in retrieving monthly payment method stats by card number.

```go
var ErrGetMonthlyTopupMethodsByCardFailed = errors.New("failed to get monthly topup methods by card number")
```

**Var:**

ErrGetMonthlyTopupMethodsFailed indicates failure in retrieving monthly top-up payment methods statistics.

```go
var ErrGetMonthlyTopupMethodsFailed = errors.New("failed to get monthly topup methods")
```

**Var:**

ErrGetYearlyTopupAmountsByCardFailed indicates failure in retrieving yearly top-up amount stats by card number.

```go
var ErrGetYearlyTopupAmountsByCardFailed = errors.New("failed to get yearly topup amounts by card number")
```

**Var:**

ErrGetYearlyTopupAmountsFailed indicates failure in retrieving the yearly top-up amounts.

```go
var ErrGetYearlyTopupAmountsFailed = errors.New("failed to get yearly topup amounts")
```

**Var:**

ErrGetYearlyTopupMethodsByCardFailed indicates failure in retrieving yearly payment method stats by card number.

```go
var ErrGetYearlyTopupMethodsByCardFailed = errors.New("failed to get yearly topup methods by card number")
```

**Var:**

ErrGetYearlyTopupMethodsFailed indicates failure in retrieving yearly top-up payment methods statistics.

```go
var ErrGetYearlyTopupMethodsFailed = errors.New("failed to get yearly topup methods")
```

**Var:**

ErrGetYearlyTopupStatusFailedByCardFailed indicates failure in getting yearly failed top-up status by card number.

```go
var ErrGetYearlyTopupStatusFailedByCardFailed = errors.New("failed to get yearly topup status failed by card number")
```

**Var:**

ErrGetYearlyTopupStatusFailedFailed indicates failure in getting the yearly count of failed top-ups.

```go
var ErrGetYearlyTopupStatusFailedFailed = errors.New("failed to get yearly topup status failed")
```

**Var:**

ErrGetYearlyTopupStatusSuccessByCardFailed indicates failure in getting yearly successful top-up status by card number.

```go
var ErrGetYearlyTopupStatusSuccessByCardFailed = errors.New("failed to get yearly topup status success by card number")
```

**Var:**

ErrGetYearlyTopupStatusSuccessFailed indicates failure in getting the yearly count of successful top-ups.

```go
var ErrGetYearlyTopupStatusSuccessFailed = errors.New("failed to get yearly topup status success")
```

**Var:**

ErrRestoreAllTopupFailed indicates failure in restoring all trashed top-ups.

```go
var ErrRestoreAllTopupFailed = errors.New("failed to restore all topups")
```

**Var:**

ErrRestoreTopupFailed indicates failure in restoring a previously trashed top-up.

```go
var ErrRestoreTopupFailed = errors.New("failed to restore topup")
```

**Var:**

ErrTopupNotFoundRes indicates that a requested top-up was not found.

```go
var ErrTopupNotFoundRes = response.NewErrorResponse("Topup not found", http.StatusNotFound)
```

**Var:**

ErrTrashedTopupFailed indicates failure in soft-deleting (trashing) a top-up.

```go
var ErrTrashedTopupFailed = errors.New("failed to soft-delete (trash) topup")
```

**Var:**

ErrUpdateTopupAmountFailed indicates failure in updating only the top-up amount.

```go
var ErrUpdateTopupAmountFailed = errors.New("failed to update topup amount")
```

**Var:**

ErrUpdateTopupFailed indicates failure in updating an existing top-up record.

```go
var ErrUpdateTopupFailed = errors.New("failed to update topup")
```

**Var:**

ErrUpdateTopupStatusFailed indicates failure in updating the top-up status (e.g., success/failed).

```go
var ErrUpdateTopupStatusFailed = errors.New("failed to update topup status")
```