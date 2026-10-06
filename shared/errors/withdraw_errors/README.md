# 📦 Package `withdraw_errors`

**Source Path:** `shared/errors/withdraw_errors`

## 🏷️ Variables

**Var:**

ErrApiBindCreateWithdraw is an error response for bind failed: invalid create withdraw request.

```go
var ErrApiBindCreateWithdraw = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "bind failed: invalid create withdraw request", http.StatusBadRequest)
}
```

**Var:**

ErrApiBindUpdateWithdraw is an error response for bind failed: invalid update withdraw request.

```go
var ErrApiBindUpdateWithdraw = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "bind failed: invalid update withdraw request", http.StatusBadRequest)
}
```

**Var:**

ErrApiFailedCreateWithdraw is an error response for failed to create withdraw.

```go
var ErrApiFailedCreateWithdraw = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to create withdraw", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedDeleteAllWithdrawPermanent is an error response for failed to permanently delete all withdraws.

```go
var ErrApiFailedDeleteAllWithdrawPermanent = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to permanently delete all withdraws", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedDeleteWithdrawPermanent is an error response for failed to permanently delete withdraw.

```go
var ErrApiFailedDeleteWithdrawPermanent = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to permanently delete withdraw", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindAllWithdraw is an error response for failed to fetch all withdraws.

```go
var ErrApiFailedFindAllWithdraw = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch all withdraws", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindAllWithdrawByCardNumber is an error response for failed to fetch withdraws by card number.

```go
var ErrApiFailedFindAllWithdrawByCardNumber = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch withdraws by card number", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindByActiveWithdraw is an error response for failed to fetch active withdraws.

```go
var ErrApiFailedFindByActiveWithdraw = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch active withdraws", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindByCardNumber is an error response for failed to fetch withdraws using card number.

```go
var ErrApiFailedFindByCardNumber = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch withdraws using card number", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindByIdWithdraw is an error response for failed to fetch withdraw by ID.

```go
var ErrApiFailedFindByIdWithdraw = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch withdraw by ID", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindByTrashedWithdraw is an error response for failed to fetch trashed withdraws.

```go
var ErrApiFailedFindByTrashedWithdraw = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch trashed withdraws", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindMonthlyWithdrawStatusFailed is an error response for failed to fetch monthly failed withdraws.

```go
var ErrApiFailedFindMonthlyWithdrawStatusFailed = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly failed withdraws", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindMonthlyWithdrawStatusFailedCardNumber is an error response for failed to fetch monthly failed withdraws by card number.

```go
var ErrApiFailedFindMonthlyWithdrawStatusFailedCardNumber = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly failed withdraws by card number", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindMonthlyWithdrawStatusSuccess is an error response for failed to fetch monthly successful withdraws.

```go
var ErrApiFailedFindMonthlyWithdrawStatusSuccess = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly successful withdraws", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindMonthlyWithdrawStatusSuccessCardNumber is an error response for failed to fetch monthly successful withdraws by card number.

```go
var ErrApiFailedFindMonthlyWithdrawStatusSuccessCardNumber = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly successful withdraws by card number", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindMonthlyWithdraws is an error response for failed to fetch monthly withdraw amounts.

```go
var ErrApiFailedFindMonthlyWithdraws = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly withdraw amounts", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindMonthlyWithdrawsByCardNumber is an error response for failed to fetch monthly withdraw amounts by card number.

```go
var ErrApiFailedFindMonthlyWithdrawsByCardNumber = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly withdraw amounts by card number", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindYearlyWithdrawStatusFailed is an error response for failed to fetch yearly failed withdraws.

```go
var ErrApiFailedFindYearlyWithdrawStatusFailed = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly failed withdraws", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindYearlyWithdrawStatusFailedCardNumber is an error response for failed to fetch yearly failed withdraws by card number.

```go
var ErrApiFailedFindYearlyWithdrawStatusFailedCardNumber = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly failed withdraws by card number", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindYearlyWithdrawStatusSuccess is an error response for failed to fetch yearly successful withdraws.

```go
var ErrApiFailedFindYearlyWithdrawStatusSuccess = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly successful withdraws", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindYearlyWithdrawStatusSuccessCardNumber is an error response for failed to fetch yearly successful withdraws by card number.

```go
var ErrApiFailedFindYearlyWithdrawStatusSuccessCardNumber = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly successful withdraws by card number", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindYearlyWithdraws is an error response for failed to fetch yearly withdraw amounts.

```go
var ErrApiFailedFindYearlyWithdraws = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly withdraw amounts", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindYearlyWithdrawsByCardNumber is an error response for failed to fetch yearly withdraw amounts by card number.

```go
var ErrApiFailedFindYearlyWithdrawsByCardNumber = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly withdraw amounts by card number", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedRestoreAllWithdraw is an error response for failed to restore all withdraws.

```go
var ErrApiFailedRestoreAllWithdraw = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to restore all withdraws", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedRestoreWithdraw is an error response for failed to restore withdraw.

```go
var ErrApiFailedRestoreWithdraw = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to restore withdraw", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedTrashedWithdraw is an error response for failed to move withdraw to trash.

```go
var ErrApiFailedTrashedWithdraw = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to move withdraw to trash", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedUpdateWithdraw is an error response for failed to update withdraw.

```go
var ErrApiFailedUpdateWithdraw = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to update withdraw", http.StatusInternalServerError)
}
```

**Var:**

ErrApiInvalidCardNumber is an error response for invalid card number.

```go
var ErrApiInvalidCardNumber = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Invalid card number", http.StatusBadRequest)
}
```

**Var:**

ErrApiInvalidMonth is an error response for invalid month.

```go
var ErrApiInvalidMonth = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Invalid month", http.StatusBadRequest)
}
```

**Var:**

ErrApiInvalidYear is an error response for invalid year.

```go
var ErrApiInvalidYear = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "year", "Invalid year", http.StatusBadRequest)
}
```

**Var:**

ErrApiValidateCreateWithdraw is an error response for validation failed: invalid create withdraw request.

```go
var ErrApiValidateCreateWithdraw = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "validation failed: invalid create withdraw request", http.StatusBadRequest)
}
```

**Var:**

ErrApiValidateUpdateWithdraw is an error response for validation failed: invalid update withdraw request.

```go
var ErrApiValidateUpdateWithdraw = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "validation failed: invalid update withdraw request", http.StatusBadRequest)
}
```

**Var:**

ErrApiWithdrawInvalidID is an error response for invalid withdraw ID.

```go
var ErrApiWithdrawInvalidID = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Invalid Withdraw ID", http.StatusBadRequest)
}
```

**Var:**

ErrApiWithdrawInvalidUserID is an error response for invalid withdraw merchant ID.

```go
var ErrApiWithdrawInvalidUserID = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Invalid Withdraw Merchant ID", http.StatusBadRequest)
}
```

**Var:**

ErrApiWithdrawNotFound is an error response for withdraw not found.

```go
var ErrApiWithdrawNotFound = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Withdraw not found", http.StatusNotFound)
}
```

**Var:**

ErrCreateWithdrawFailed indicates a failure when creating a new withdraw record.

```go
var ErrCreateWithdrawFailed = errors.New("failed to create withdraw")
```

**Var:**

ErrDeleteAllWithdrawsPermanentFailed indicates a failure when permanently deleting all withdraw records.

```go
var ErrDeleteAllWithdrawsPermanentFailed = errors.New("failed to permanently delete all withdraws")
```

**Var:**

ErrDeleteWithdrawPermanentFailed indicates a failure when permanently deleting a withdraw record.

```go
var ErrDeleteWithdrawPermanentFailed = errors.New("failed to permanently delete withdraw")
```

**Var:**

ErrFailedCreateWithdraw is used when failed to create withdraw

```go
var ErrFailedCreateWithdraw = response.NewErrorResponse("Failed to create withdraw", http.StatusInternalServerError)
```

**Var:**

ErrFailedDeleteAllWithdrawPermanent is used when failed to permanently delete all withdraws

```go
var ErrFailedDeleteAllWithdrawPermanent = response.NewErrorResponse("Failed to permanently delete all withdraws", http.StatusInternalServerError)
```

**Var:**

ErrFailedDeleteWithdrawPermanent is used when failed to permanently delete withdraw

```go
var ErrFailedDeleteWithdrawPermanent = response.NewErrorResponse("Failed to permanently delete withdraw", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindActiveWithdraws is used when failed to fetch active withdraws

```go
var ErrFailedFindActiveWithdraws = response.NewErrorResponse("Failed to fetch active withdraws", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindAllWithdraws is used when failed to fetch all withdraws

```go
var ErrFailedFindAllWithdraws = response.NewErrorResponse("Failed to fetch all withdraws", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindAllWithdrawsByCard is used when failed to fetch all withdraws by card number

```go
var ErrFailedFindAllWithdrawsByCard = response.NewErrorResponse("Failed to fetch all withdraws by card number", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindMonthWithdrawStatusFailed is used when failed to fetch monthly failed withdraws

```go
var ErrFailedFindMonthWithdrawStatusFailed = response.NewErrorResponse("Failed to fetch monthly failed withdraws", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindMonthWithdrawStatusFailedByCard is used when failed to fetch monthly failed withdraws by card

```go
var ErrFailedFindMonthWithdrawStatusFailedByCard = response.NewErrorResponse("Failed to fetch monthly failed withdraws by card", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindMonthWithdrawStatusSuccess is used when failed to fetch monthly successful withdraws

```go
var ErrFailedFindMonthWithdrawStatusSuccess = response.NewErrorResponse("Failed to fetch monthly successful withdraws", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindMonthWithdrawStatusSuccessByCard is used when failed to fetch monthly successful withdraws by card

```go
var ErrFailedFindMonthWithdrawStatusSuccessByCard = response.NewErrorResponse("Failed to fetch monthly successful withdraws by card", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindMonthlyWithdraws is used when failed to fetch monthly withdraw amounts

```go
var ErrFailedFindMonthlyWithdraws = response.NewErrorResponse("Failed to fetch monthly withdraw amounts", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindMonthlyWithdrawsByCardNumber is used when failed to fetch monthly withdraw amounts by card

```go
var ErrFailedFindMonthlyWithdrawsByCardNumber = response.NewErrorResponse("Failed to fetch monthly withdraw amounts by card", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindTrashedWithdraws is used when failed to fetch trashed withdraws

```go
var ErrFailedFindTrashedWithdraws = response.NewErrorResponse("Failed to fetch trashed withdraws", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindYearWithdrawStatusFailed is used when failed to fetch yearly failed withdraws

```go
var ErrFailedFindYearWithdrawStatusFailed = response.NewErrorResponse("Failed to fetch yearly failed withdraws", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindYearWithdrawStatusFailedByCard is used when failed to fetch yearly failed withdraws by card

```go
var ErrFailedFindYearWithdrawStatusFailedByCard = response.NewErrorResponse("Failed to fetch yearly failed withdraws by card", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindYearWithdrawStatusSuccess is used when failed to fetch yearly successful withdraws

```go
var ErrFailedFindYearWithdrawStatusSuccess = response.NewErrorResponse("Failed to fetch yearly successful withdraws", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindYearWithdrawStatusSuccessByCard is used when failed to fetch yearly successful withdraws by card

```go
var ErrFailedFindYearWithdrawStatusSuccessByCard = response.NewErrorResponse("Failed to fetch yearly successful withdraws by card", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindYearlyWithdraws is used when failed to fetch yearly withdraw amounts

```go
var ErrFailedFindYearlyWithdraws = response.NewErrorResponse("Failed to fetch yearly withdraw amounts", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindYearlyWithdrawsByCardNumber is used when failed to fetch yearly withdraw amounts by card

```go
var ErrFailedFindYearlyWithdrawsByCardNumber = response.NewErrorResponse("Failed to fetch yearly withdraw amounts by card", http.StatusInternalServerError)
```

**Var:**

ErrFailedRestoreAllWithdraw is used when failed to restore all withdraws

```go
var ErrFailedRestoreAllWithdraw = response.NewErrorResponse("Failed to restore all withdraws", http.StatusInternalServerError)
```

**Var:**

ErrFailedRestoreWithdraw is used when failed to restore withdraw

```go
var ErrFailedRestoreWithdraw = response.NewErrorResponse("Failed to restore withdraw", http.StatusInternalServerError)
```

**Var:**

ErrFailedSendEmail is used when failed to send email

```go
var ErrFailedSendEmail = response.NewErrorResponse("Failed to send email", http.StatusInternalServerError)
```

**Var:**

ErrFailedTrashedWithdraw is used when failed to trash withdraw

```go
var ErrFailedTrashedWithdraw = response.NewErrorResponse("Failed to trash withdraw", http.StatusInternalServerError)
```

**Var:**

ErrFailedUpdateWithdraw is used when failed to update withdraw

```go
var ErrFailedUpdateWithdraw = response.NewErrorResponse("Failed to update withdraw", http.StatusInternalServerError)
```

**Var:**

ErrFindActiveWithdrawsFailed is used when the system fails to find active withdraws

```go
var ErrFindActiveWithdrawsFailed = errors.New("failed to find active withdraws")
```

**Var:**

ErrFindAllWithdrawsFailed is used when the system fails to find all withdraws

```go
var ErrFindAllWithdrawsFailed = errors.New("failed to find all withdraws")
```

**Var:**

ErrFindTrashedWithdrawsFailed is used when the system fails to find trashed withdraws

```go
var ErrFindTrashedWithdrawsFailed = errors.New("failed to find trashed withdraws")
```

**Var:**

ErrFindWithdrawByIdFailed is used when the system fails to find a withdraw by ID

```go
var ErrFindWithdrawByIdFailed = errors.New("failed to find withdraw by ID")
```

**Var:**

ErrFindWithdrawsByCardNumberFailed is used when the system fails to find withdraws by card number

```go
var ErrFindWithdrawsByCardNumberFailed = errors.New("failed to find withdraws by card number")
```

**Var:**

ErrGetMonthWithdrawStatusFailedByCardFailed is used when the system fails to get the monthly withdraw status failed by card number

```go
var ErrGetMonthWithdrawStatusFailedByCardFailed = errors.New("failed to get monthly withdraw status failed by card number")
```

**Var:**

ErrGetMonthWithdrawStatusFailedFailed is used when the system fails to get the monthly withdraw status failed

```go
var ErrGetMonthWithdrawStatusFailedFailed = errors.New("failed to get monthly withdraw status failed")
```

**Var:**

ErrGetMonthWithdrawStatusSuccessByCardFailed is used when the system fails to get the monthly withdraw status success by card number

```go
var ErrGetMonthWithdrawStatusSuccessByCardFailed = errors.New("failed to get monthly withdraw status success by card number")
```

**Var:**

ErrGetMonthWithdrawStatusSuccessFailed is used when the system fails to get the monthly withdraw status success

```go
var ErrGetMonthWithdrawStatusSuccessFailed = errors.New("failed to get monthly withdraw status success")
```

**Var:**

ErrGetMonthlyWithdrawsByCardFailed indicates a failure when retrieving monthly withdraw amounts by card number.

```go
var ErrGetMonthlyWithdrawsByCardFailed = errors.New("failed to get monthly withdraw amounts by card number")
```

**Var:**

ErrGetMonthlyWithdrawsFailed is used when the system fails to get monthly withdraw amounts

```go
var ErrGetMonthlyWithdrawsFailed = errors.New("failed to get monthly withdraw amounts")
```

**Var:**

ErrGetYearlyWithdrawStatusFailedByCardFailed is used when the system fails to get the yearly withdraw status failed by card number

```go
var ErrGetYearlyWithdrawStatusFailedByCardFailed = errors.New("failed to get yearly withdraw status failed by card number")
```

**Var:**

ErrGetYearlyWithdrawStatusFailedFailed is used when the system fails to get the yearly withdraw status failed

```go
var ErrGetYearlyWithdrawStatusFailedFailed = errors.New("failed to get yearly withdraw status failed")
```

**Var:**

ErrGetYearlyWithdrawStatusSuccessByCardFailed is used when the system fails to get the yearly withdraw status success by card number

```go
var ErrGetYearlyWithdrawStatusSuccessByCardFailed = errors.New("failed to get yearly withdraw status success by card number")
```

**Var:**

ErrGetYearlyWithdrawStatusSuccessFailed is used when the system fails to get the yearly withdraw status success

```go
var ErrGetYearlyWithdrawStatusSuccessFailed = errors.New("failed to get yearly withdraw status success")
```

**Var:**

ErrGetYearlyWithdrawsByCardFailed indicates a failure when retrieving yearly withdraw amounts by card number.

```go
var ErrGetYearlyWithdrawsByCardFailed = errors.New("failed to get yearly withdraw amounts by card number")
```

**Var:**

ErrGetYearlyWithdrawsFailed is used when the system fails to get yearly withdraw amounts

```go
var ErrGetYearlyWithdrawsFailed = errors.New("failed to get yearly withdraw amounts")
```

**Var:**

ErrRestoreAllWithdrawsFailed indicates a failure when restoring all trashed withdraw records.

```go
var ErrRestoreAllWithdrawsFailed = errors.New("failed to restore all withdraws")
```

**Var:**

ErrRestoreWithdrawFailed indicates a failure when restoring a previously trashed withdraw record.

```go
var ErrRestoreWithdrawFailed = errors.New("failed to restore withdraw")
```

**Var:**

ErrTrashedWithdrawFailed indicates a failure when soft-deleting (trashing) a withdraw record.

```go
var ErrTrashedWithdrawFailed = errors.New("failed to soft-delete (trash) withdraw")
```

**Var:**

ErrUpdateWithdrawFailed indicates a failure when updating a withdraw record.

```go
var ErrUpdateWithdrawFailed = errors.New("failed to update withdraw")
```

**Var:**

ErrUpdateWithdrawStatusFailed indicates a failure when updating the status of a withdraw record.

```go
var ErrUpdateWithdrawStatusFailed = errors.New("failed to update withdraw status")
```

**Var:**

ErrWithdrawNotFound is used when withdraw is not found

```go
var ErrWithdrawNotFound = response.NewErrorResponse("Withdraw not found", http.StatusNotFound)
```