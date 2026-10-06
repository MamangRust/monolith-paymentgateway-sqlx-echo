# 📦 Package `transfer_errors`

**Source Path:** `./shared/errors/transfer_errors`

## 🏷️ Variables

**Var:**

ErrApiBindCreateTransfer returns an API error response
when binding the request for creating a transfer fails due to invalid input.

```go
var ErrApiBindCreateTransfer = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "bind failed: invalid create transfer request", http.StatusBadRequest)
}
```

**Var:**

ErrApiBindUpdateTransfer returns an API error response
when binding the request for updating a transfer fails due to invalid input.

```go
var ErrApiBindUpdateTransfer = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "bind failed: invalid update transfer request", http.StatusBadRequest)
}
```

**Var:**

ErrApiFailedCreateTransfer returns an API error response
when creating a new transfer fails.

```go
var ErrApiFailedCreateTransfer = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to create transfer", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedDeleteAllTransferPermanent returns an API error response
when attempting to permanently delete all transfers fails.

```go
var ErrApiFailedDeleteAllTransferPermanent = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to permanently delete all transfers", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedDeleteTransferPermanent returns an API error response
when attempting to permanently delete a transfer fails.

```go
var ErrApiFailedDeleteTransferPermanent = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to permanently delete transfer", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindAllTransfers is an error response for failed to fetch all transfers.

```go
var ErrApiFailedFindAllTransfers = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch all transfers", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindByActiveTransfer is an error response for failed to fetch active transfers.

```go
var ErrApiFailedFindByActiveTransfer = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch active transfers", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindByIdTransfer is an error response for failed to fetch transfer by ID.

```go
var ErrApiFailedFindByIdTransfer = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch transfer by ID", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindByTransferFrom is an error response for failed to fetch transfers by transfer_from.

```go
var ErrApiFailedFindByTransferFrom = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch transfers by transfer_from", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindByTransferTo is an error response for failed to fetch transfers by transfer_to.

```go
var ErrApiFailedFindByTransferTo = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch transfers by transfer_to", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindByTrashedTransfer returns an API error response when fetching trashed (soft-deleted) transfers fails.

```go
var ErrApiFailedFindByTrashedTransfer = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch trashed transfers", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindMonthlyTransferAmounts returns an API error response
when fetching monthly total transfer amounts fails.

```go
var ErrApiFailedFindMonthlyTransferAmounts = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly transfer amounts", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindMonthlyTransferAmountsByReceiverCardNumber returns an API error response
when fetching monthly transfer amounts by receiver card number fails.

```go
var ErrApiFailedFindMonthlyTransferAmountsByReceiverCardNumber = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly transfer amounts by receiver card number", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindMonthlyTransferAmountsBySenderCardNumber returns an API error response
when fetching monthly transfer amounts by sender card number fails.

```go
var ErrApiFailedFindMonthlyTransferAmountsBySenderCardNumber = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly transfer amounts by sender card number", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindMonthlyTransferStatusFailed returns an API error response when fetching monthly failed transfer data fails.

```go
var ErrApiFailedFindMonthlyTransferStatusFailed = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly failed transfers", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindMonthlyTransferStatusFailedByCardNumber returns an API error response when fetching monthly failed transfers by card number fails.

```go
var ErrApiFailedFindMonthlyTransferStatusFailedByCardNumber = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly failed transfers by card number", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindMonthlyTransferStatusSuccess returns an API error response when fetching monthly successful transfer data fails.

```go
var ErrApiFailedFindMonthlyTransferStatusSuccess = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly successful transfers", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindMonthlyTransferStatusSuccessByCardNumber returns an API error response when fetching monthly successful transfers by card number fails.

```go
var ErrApiFailedFindMonthlyTransferStatusSuccessByCardNumber = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly successful transfers by card number", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindYearlyTransferAmounts returns an API error response
when fetching yearly total transfer amounts fails.

```go
var ErrApiFailedFindYearlyTransferAmounts = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly transfer amounts", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindYearlyTransferAmountsByReceiverCardNumber returns an API error response
when fetching yearly transfer amounts by receiver card number fails.

```go
var ErrApiFailedFindYearlyTransferAmountsByReceiverCardNumber = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly transfer amounts by receiver card number", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindYearlyTransferAmountsBySenderCardNumber returns an API error response
when fetching yearly transfer amounts by sender card number fails.

```go
var ErrApiFailedFindYearlyTransferAmountsBySenderCardNumber = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly transfer amounts by sender card number", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindYearlyTransferStatusFailed returns an API error response when fetching yearly failed transfer data fails.

```go
var ErrApiFailedFindYearlyTransferStatusFailed = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly failed transfers", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindYearlyTransferStatusFailedByCardNumber returns an API error response
when fetching yearly failed transfers by card number fails.

```go
var ErrApiFailedFindYearlyTransferStatusFailedByCardNumber = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly failed transfers by card number", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindYearlyTransferStatusSuccess returns an API error response when fetching yearly successful transfer data fails.

```go
var ErrApiFailedFindYearlyTransferStatusSuccess = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly successful transfers", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindYearlyTransferStatusSuccessByCardNumber returns an API error response when fetching yearly successful transfers by card number fails.

```go
var ErrApiFailedFindYearlyTransferStatusSuccessByCardNumber = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly successful transfers by card number", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedRestoreAllTransfer returns an API error response
when attempting to restore all trashed transfers fails.

```go
var ErrApiFailedRestoreAllTransfer = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to restore all transfers", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedRestoreTransfer returns an API error response
when attempting to restore a trashed transfer fails.

```go
var ErrApiFailedRestoreTransfer = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to restore transfer", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedTrashedTransfer returns an API error response
when attempting to trash (soft-delete) a transfer fails.

```go
var ErrApiFailedTrashedTransfer = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to trash transfer", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedUpdateTransfer returns an API error response
when updating a transfer fails.

```go
var ErrApiFailedUpdateTransfer = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to update transfer", http.StatusInternalServerError)
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
	return response.NewApiErrorResponse(c, "error", "Invalid year", http.StatusBadRequest)
}
```

**Var:**

ErrApiTransferInvalidID is an error response for invalid transfer ID.

```go
var ErrApiTransferInvalidID = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Invalid Transfer ID", http.StatusBadRequest)
}
```

**Var:**

ErrApiTransferInvalidMerchantID is an error response for invalid transfer merchant ID.

```go
var ErrApiTransferInvalidMerchantID = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Invalid Transfer Merchant ID", http.StatusBadRequest)
}
```

**Var:**

ErrApiValidateCreateTransfer returns an API error response
when validation fails on the create transfer request payload.

```go
var ErrApiValidateCreateTransfer = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "validation failed: invalid create transfer request", http.StatusBadRequest)
}
```

**Var:**

ErrApiValidateUpdateTransfer returns an API error response
when validation fails on the update transfer request payload.

```go
var ErrApiValidateUpdateTransfer = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "validation failed: invalid update transfer request", http.StatusBadRequest)
}
```

**Var:**

ErrCreateTransferFailed indicates a failure when creating a new transfer.

```go
var ErrCreateTransferFailed = errors.New("failed to create transfer")
```

**Var:**

ErrDeleteAllTransfersPermanentFailed indicates a failure when permanently deleting all transfers.

```go
var ErrDeleteAllTransfersPermanentFailed = errors.New("failed to permanently delete all transfers")
```

**Var:**

ErrDeleteTransferPermanentFailed indicates a failure when permanently deleting a transfer.

```go
var ErrDeleteTransferPermanentFailed = errors.New("failed to permanently delete transfer")
```

**Var:**

ErrFailedCreateTransfer indicates a failure when attempting to create a new transfer record.

```go
var ErrFailedCreateTransfer = response.NewErrorResponse("Failed to create transfer", http.StatusInternalServerError)
```

**Var:**

ErrFailedDeleteAllTransfersPermanent indicates a failure when attempting to permanently delete all transfer records.

```go
var ErrFailedDeleteAllTransfersPermanent = response.NewErrorResponse("Failed to permanently delete all transfers", http.StatusInternalServerError)
```

**Var:**

ErrFailedDeleteTransferPermanent indicates a failure when attempting to permanently delete a transfer.

```go
var ErrFailedDeleteTransferPermanent = response.NewErrorResponse("Failed to permanently delete transfer", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindActiveTransfers indicates a failure when retrieving active transfer records.

```go
var ErrFailedFindActiveTransfers = response.NewErrorResponse("Failed to fetch active transfers", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindAllTransfers indicates a failure when retrieving all transfer records.

```go
var ErrFailedFindAllTransfers = response.NewErrorResponse("Failed to fetch all transfers", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindMonthTransferStatusFailed indicates a failure when retrieving monthly failed transfer statistics.

```go
var ErrFailedFindMonthTransferStatusFailed = response.NewErrorResponse("Failed to fetch monthly failed transfers", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindMonthTransferStatusFailedByCard indicates a failure when retrieving monthly failed transfers by card number.

```go
var ErrFailedFindMonthTransferStatusFailedByCard = response.NewErrorResponse("Failed to fetch monthly failed transfers by card", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindMonthTransferStatusSuccess indicates a failure when retrieving monthly successful transfer statistics.

```go
var ErrFailedFindMonthTransferStatusSuccess = response.NewErrorResponse("Failed to fetch monthly successful transfers", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindMonthTransferStatusSuccessByCard indicates a failure when retrieving monthly successful transfers by card number.

```go
var ErrFailedFindMonthTransferStatusSuccessByCard = response.NewErrorResponse("Failed to fetch monthly successful transfers by card", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindMonthlyTransferAmounts indicates a failure when retrieving the total monthly transfer amounts.

```go
var ErrFailedFindMonthlyTransferAmounts = response.NewErrorResponse("Failed to fetch monthly transfer amounts", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindMonthlyTransferAmountsByReceiverCard indicates a failure when retrieving monthly transfer amounts by receiver card number.

```go
var ErrFailedFindMonthlyTransferAmountsByReceiverCard = response.NewErrorResponse("Failed to fetch monthly transfer amounts by receiver card", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindMonthlyTransferAmountsBySenderCard indicates a failure when retrieving monthly transfer amounts by sender card number.

```go
var ErrFailedFindMonthlyTransferAmountsBySenderCard = response.NewErrorResponse("Failed to fetch monthly transfer amounts by sender card", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindTransfersByReceiver indicates a failure when retrieving transfers filtered by receiver card.

```go
var ErrFailedFindTransfersByReceiver = response.NewErrorResponse("Failed to fetch transfers by receiver", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindTransfersBySender indicates a failure when retrieving transfers filtered by sender card.

```go
var ErrFailedFindTransfersBySender = response.NewErrorResponse("Failed to fetch transfers by sender", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindTrashedTransfers indicates a failure when retrieving trashed (soft-deleted) transfer records.

```go
var ErrFailedFindTrashedTransfers = response.NewErrorResponse("Failed to fetch trashed transfers", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindYearTransferStatusFailed indicates a failure when retrieving yearly failed transfer statistics.

```go
var ErrFailedFindYearTransferStatusFailed = response.NewErrorResponse("Failed to fetch yearly failed transfers", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindYearTransferStatusFailedByCard indicates a failure when retrieving yearly failed transfers by card number.

```go
var ErrFailedFindYearTransferStatusFailedByCard = response.NewErrorResponse("Failed to fetch yearly failed transfers by card", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindYearTransferStatusSuccess indicates a failure when retrieving yearly successful transfer statistics.

```go
var ErrFailedFindYearTransferStatusSuccess = response.NewErrorResponse("Failed to fetch yearly successful transfers", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindYearTransferStatusSuccessByCard indicates a failure when retrieving yearly successful transfers by card number.

```go
var ErrFailedFindYearTransferStatusSuccessByCard = response.NewErrorResponse("Failed to fetch yearly successful transfers by card", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindYearlyTransferAmounts indicates a failure when retrieving the total yearly transfer amounts.

```go
var ErrFailedFindYearlyTransferAmounts = response.NewErrorResponse("Failed to fetch yearly transfer amounts", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindYearlyTransferAmountsByReceiverCard indicates a failure when retrieving yearly transfer amounts by receiver card number.

```go
var ErrFailedFindYearlyTransferAmountsByReceiverCard = response.NewErrorResponse("Failed to fetch yearly transfer amounts by receiver card", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindYearlyTransferAmountsBySenderCard indicates a failure when retrieving yearly transfer amounts by sender card number.

```go
var ErrFailedFindYearlyTransferAmountsBySenderCard = response.NewErrorResponse("Failed to fetch yearly transfer amounts by sender card", http.StatusInternalServerError)
```

**Var:**

ErrFailedRestoreAllTransfers indicates a failure when attempting to restore all trashed transfers.

```go
var ErrFailedRestoreAllTransfers = response.NewErrorResponse("Failed to restore all transfers", http.StatusInternalServerError)
```

**Var:**

ErrFailedRestoreTransfer indicates a failure when attempting to restore a previously trashed transfer.

```go
var ErrFailedRestoreTransfer = response.NewErrorResponse("Failed to restore transfer", http.StatusInternalServerError)
```

**Var:**

ErrFailedTrashedTransfer indicates a failure when attempting to soft-delete (trash) a transfer.

```go
var ErrFailedTrashedTransfer = response.NewErrorResponse("Failed to trash transfer", http.StatusInternalServerError)
```

**Var:**

ErrFailedUpdateTransfer indicates a failure when attempting to update an existing transfer record.

```go
var ErrFailedUpdateTransfer = response.NewErrorResponse("Failed to update transfer", http.StatusInternalServerError)
```

**Var:**

ErrFindActiveTransfersFailed indicates a failure when retrieving active (non-trashed) transfer records.

```go
var ErrFindActiveTransfersFailed = errors.New("failed to find active transfers")
```

**Var:**

ErrFindAllTransfersFailed indicates a failure when retrieving all transfer records.

```go
var ErrFindAllTransfersFailed = errors.New("failed to find all transfers")
```

**Var:**

ErrFindTransferByIdFailed indicates a failure when retrieving a transfer record by its ID.

```go
var ErrFindTransferByIdFailed = errors.New("failed to find transfer by ID")
```

**Var:**

ErrFindTransferByTransferFromFailed indicates a failure when retrieving transfers by the sender (transfer from).

```go
var ErrFindTransferByTransferFromFailed = errors.New("failed to find transfer by transfer from")
```

**Var:**

ErrFindTransferByTransferToFailed indicates a failure when retrieving transfers by the receiver (transfer to).

```go
var ErrFindTransferByTransferToFailed = errors.New("failed to find transfer by transfer to")
```

**Var:**

ErrFindTrashedTransfersFailed indicates a failure when retrieving trashed (soft-deleted) transfer records.

```go
var ErrFindTrashedTransfersFailed = errors.New("failed to find trashed transfers")
```

**Var:**

ErrGetMonthTransferStatusFailedByCardFailed indicates a failure when retrieving monthly failed transfers filtered by card number.

```go
var ErrGetMonthTransferStatusFailedByCardFailed = errors.New("failed to get monthly transfer status failed by card number")
```

**Var:**

ErrGetMonthTransferStatusFailedFailed indicates a failure when retrieving monthly failed transfer statistics.

```go
var ErrGetMonthTransferStatusFailedFailed = errors.New("failed to get monthly transfer status failed")
```

**Var:**

ErrGetMonthTransferStatusSuccessByCardFailed indicates a failure when retrieving monthly successful transfers filtered by card number.

```go
var ErrGetMonthTransferStatusSuccessByCardFailed = errors.New("failed to get monthly transfer status success by card number")
```

**Var:**

ErrGetMonthTransferStatusSuccessFailed indicates a failure when retrieving monthly successful transfer statistics.

```go
var ErrGetMonthTransferStatusSuccessFailed = errors.New("failed to get monthly transfer status success")
```

**Var:**

ErrGetMonthlyTransferAmountsByReceiverCardFailed indicates a failure when retrieving monthly transfer amounts filtered by receiver card number.

```go
var ErrGetMonthlyTransferAmountsByReceiverCardFailed = errors.New("failed to get monthly transfer amounts by receiver card number")
```

**Var:**

ErrGetMonthlyTransferAmountsBySenderCardFailed indicates a failure when retrieving monthly transfer amounts filtered by sender card number.

```go
var ErrGetMonthlyTransferAmountsBySenderCardFailed = errors.New("failed to get monthly transfer amounts by sender card number")
```

**Var:**

ErrGetMonthlyTransferAmountsFailed indicates a failure when retrieving the total amount of monthly transfers.

```go
var ErrGetMonthlyTransferAmountsFailed = errors.New("failed to get monthly transfer amounts")
```

**Var:**

ErrGetYearlyTransferAmountsByReceiverCardFailed indicates a failure when retrieving yearly transfer amounts filtered by receiver card number.

```go
var ErrGetYearlyTransferAmountsByReceiverCardFailed = errors.New("failed to get yearly transfer amounts by receiver card number")
```

**Var:**

ErrGetYearlyTransferAmountsBySenderCardFailed indicates a failure when retrieving yearly transfer amounts filtered by sender card number.

```go
var ErrGetYearlyTransferAmountsBySenderCardFailed = errors.New("failed to get yearly transfer amounts by sender card number")
```

**Var:**

ErrGetYearlyTransferAmountsFailed indicates a failure when retrieving the total amount of yearly transfers.

```go
var ErrGetYearlyTransferAmountsFailed = errors.New("failed to get yearly transfer amounts")
```

**Var:**

ErrGetYearlyTransferStatusFailedByCardFailed indicates a failure when retrieving yearly failed transfers filtered by card number.

```go
var ErrGetYearlyTransferStatusFailedByCardFailed = errors.New("failed to get yearly transfer status failed by card number")
```

**Var:**

ErrGetYearlyTransferStatusFailedFailed indicates a failure when retrieving yearly failed transfer statistics.

```go
var ErrGetYearlyTransferStatusFailedFailed = errors.New("failed to get yearly transfer status failed")
```

**Var:**

ErrGetYearlyTransferStatusSuccessByCardFailed indicates a failure when retrieving yearly successful transfers filtered by card number.

```go
var ErrGetYearlyTransferStatusSuccessByCardFailed = errors.New("failed to get yearly transfer status success by card number")
```

**Var:**

ErrGetYearlyTransferStatusSuccessFailed indicates a failure when retrieving yearly successful transfer statistics.

```go
var ErrGetYearlyTransferStatusSuccessFailed = errors.New("failed to get yearly transfer status success")
```

**Var:**

ErrRestoreAllTransfersFailed indicates a failure when restoring all trashed transfers.

```go
var ErrRestoreAllTransfersFailed = errors.New("failed to restore all transfers")
```

**Var:**

ErrRestoreTransferFailed indicates a failure when restoring a previously trashed transfer.

```go
var ErrRestoreTransferFailed = errors.New("failed to restore transfer")
```

**Var:**

ErrTransferNotFound indicates that a specific transfer record was not found.

```go
var ErrTransferNotFound = response.NewErrorResponse("Transfer not found", http.StatusNotFound)
```

**Var:**

ErrTrashedTransferFailed indicates a failure when soft-deleting (trashing) a transfer.

```go
var ErrTrashedTransferFailed = errors.New("failed to soft-delete (trash) transfer")
```

**Var:**

ErrUpdateTransferAmountFailed indicates a failure when updating the amount of a transfer.

```go
var ErrUpdateTransferAmountFailed = errors.New("failed to update transfer amount")
```

**Var:**

ErrUpdateTransferFailed indicates a failure when updating an existing transfer.

```go
var ErrUpdateTransferFailed = errors.New("failed to update transfer")
```

**Var:**

ErrUpdateTransferStatusFailed indicates a failure when updating the status of a transfer.

```go
var ErrUpdateTransferStatusFailed = errors.New("failed to update transfer status")
```