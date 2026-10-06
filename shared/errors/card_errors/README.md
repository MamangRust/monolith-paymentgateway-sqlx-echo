# 📦 Package `card_errors`

**Source Path:** `shared/errors/card_errors`

## 🏷️ Variables

**Var:**

ErrApiBindCreateCard represents binding error for create card request.

```go
var ErrApiBindCreateCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "bind failed: invalid create Card request", http.StatusBadRequest)
}
```

**Var:**

ErrApiBindUpdateCard represents binding error for update card request.

```go
var ErrApiBindUpdateCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "bind failed: invalid update Card request", http.StatusBadRequest)
}
```

**Var:**

ErrApiFailedCreateCard represents error when failing to create a new card.

```go
var ErrApiFailedCreateCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to create card", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedDashboardCard returns a 500 Internal Server Error when fetching the card dashboard fails.

```go
var ErrApiFailedDashboardCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch card dashboard", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedDashboardCardByCardNumber returns a 500 Internal Server Error when fetching the dashboard by card number fails.

```go
var ErrApiFailedDashboardCardByCardNumber = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch dashboard by card number", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedDeleteAllCardPermanent represents error when failing to permanently delete all trashed cards.

```go
var ErrApiFailedDeleteAllCardPermanent = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to permanently delete all cards", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedDeleteCardPermanent represents error when failing to permanently delete a card.

```go
var ErrApiFailedDeleteCardPermanent = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to permanently delete card", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindAllCards returns a 500 Internal Server Error when fetching all cards fails.

```go
var ErrApiFailedFindAllCards = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch all cards", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindByActiveCard represents error when failing to fetch active cards.

```go
var ErrApiFailedFindByActiveCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch active cards", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindByCardNumber represents error when failing to fetch card by card number.

```go
var ErrApiFailedFindByCardNumber = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch card by card number", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindByIdCard returns a 500 Internal Server Error when fetching a card by ID fails.

```go
var ErrApiFailedFindByIdCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch card by ID", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindByTrashedCard represents error when failing to fetch trashed cards.

```go
var ErrApiFailedFindByTrashedCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch trashed cards", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindByUserIdCard represents error when failing to fetch cards by user ID.

```go
var ErrApiFailedFindByUserIdCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch card by user ID", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindMonthlyBalance returns an API error response when fetching the monthly balance fails.

```go
var ErrApiFailedFindMonthlyBalance = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly balance", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindMonthlyBalanceByCard returns an API error response when fetching the monthly balance by card fails.

```go
var ErrApiFailedFindMonthlyBalanceByCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly balance by card", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindMonthlyTopupAmount returns an API error response when fetching the monthly top-up amount fails.

```go
var ErrApiFailedFindMonthlyTopupAmount = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly topup amount", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindMonthlyTopupAmountByCard returns an API error response when fetching the monthly top-up amount by card fails.

```go
var ErrApiFailedFindMonthlyTopupAmountByCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly topup amount by card", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindMonthlyTransactionAmount returns an API error response when fetching the monthly transaction amount fails.

```go
var ErrApiFailedFindMonthlyTransactionAmount = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly transaction amount", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindMonthlyTransactionAmountByCard returns an API error response when fetching the monthly transaction amount by card fails.

```go
var ErrApiFailedFindMonthlyTransactionAmountByCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly transaction amount by card", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindMonthlyTransferReceiverAmount returns an API error response when fetching the monthly transfer receiver amount fails.

```go
var ErrApiFailedFindMonthlyTransferReceiverAmount = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly transfer receiver amount", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindMonthlyTransferReceiverAmountByCard represents error when failing to fetch monthly transfer amount by receiver card.

```go
var ErrApiFailedFindMonthlyTransferReceiverAmountByCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly transfer receiver amount by card", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindMonthlyTransferSenderAmount returns an API error response when fetching the monthly transfer sender amount fails.

```go
var ErrApiFailedFindMonthlyTransferSenderAmount = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly transfer sender amount", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindMonthlyTransferSenderAmountByCard represents error when failing to fetch monthly transfer amount by sender card.

```go
var ErrApiFailedFindMonthlyTransferSenderAmountByCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly transfer sender amount by card", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindMonthlyWithdrawAmount returns an API error response when fetching the monthly withdrawal amount fails.

```go
var ErrApiFailedFindMonthlyWithdrawAmount = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly withdraw amount", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindMonthlyWithdrawAmountByCard returns an API error response when fetching the monthly withdrawal amount by card fails.

```go
var ErrApiFailedFindMonthlyWithdrawAmountByCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly withdraw amount by card", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindYearlyBalance returns an API error response when fetching the yearly balance fails.

```go
var ErrApiFailedFindYearlyBalance = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly balance", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindYearlyBalanceByCard returns an API error response when fetching the yearly balance by card fails.

```go
var ErrApiFailedFindYearlyBalanceByCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly balance by card", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindYearlyTopupAmount returns an API error response when fetching the yearly top-up amount fails.

```go
var ErrApiFailedFindYearlyTopupAmount = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly topup amount", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindYearlyTopupAmountByCard returns an API error response when fetching the yearly top-up amount by card fails.

```go
var ErrApiFailedFindYearlyTopupAmountByCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly topup amount by card", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindYearlyTransactionAmount returns an API error response when fetching the yearly transaction amount fails.

```go
var ErrApiFailedFindYearlyTransactionAmount = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly transaction amount", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindYearlyTransactionAmountByCard returns an API error response when fetching the yearly transaction amount by card fails.

```go
var ErrApiFailedFindYearlyTransactionAmountByCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly transaction amount by card", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindYearlyTransferReceiverAmount returns an API error response when fetching the yearly transfer receiver amount fails.

```go
var ErrApiFailedFindYearlyTransferReceiverAmount = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly transfer receiver amount", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindYearlyTransferReceiverAmountByCard represents error when failing to fetch yearly transfer amount by receiver card.

```go
var ErrApiFailedFindYearlyTransferReceiverAmountByCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly transfer receiver amount by card", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindYearlyTransferSenderAmount returns an API error response when fetching the yearly transfer sender amount fails.

```go
var ErrApiFailedFindYearlyTransferSenderAmount = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly transfer sender amount", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindYearlyTransferSenderAmountByCard represents error when failing to fetch yearly transfer amount by sender card.

```go
var ErrApiFailedFindYearlyTransferSenderAmountByCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly transfer sender amount by card", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindYearlyWithdrawAmount returns an API error response when fetching the yearly withdrawal amount fails.

```go
var ErrApiFailedFindYearlyWithdrawAmount = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly withdraw amount", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindYearlyWithdrawAmountByCard returns an API error response when fetching the yearly withdrawal amount by card fails.

```go
var ErrApiFailedFindYearlyWithdrawAmountByCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly withdraw amount by card", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedRestoreAllCard represents error when failing to restore all trashed cards.

```go
var ErrApiFailedRestoreAllCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to restore all cards", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedRestoreCard represents error when failing to restore a trashed card.

```go
var ErrApiFailedRestoreCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to restore card", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedTrashCard represents error when failing to move a card to trash.

```go
var ErrApiFailedTrashCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to trash card", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedUpdateCard represents error when failing to update a card.

```go
var ErrApiFailedUpdateCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to update card", http.StatusInternalServerError)
}
```

**Var:**

ErrApiInvalidCardID returns a 400 Bad Request error when the card ID is invalid.

```go
var ErrApiInvalidCardID = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Invalid card ID", http.StatusBadRequest)
}
```

**Var:**

ErrApiInvalidCardNumber returns a 400 Bad Request error when the card number is invalid.

```go
var ErrApiInvalidCardNumber = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Invalid card number", http.StatusBadRequest)
}
```

**Var:**

ErrApiInvalidMonth returns a 400 Bad Request error when the month value is invalid.

```go
var ErrApiInvalidMonth = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Invalid month value", http.StatusBadRequest)
}
```

**Var:**

ErrApiInvalidUserID returns a 400 Bad Request error when the user ID is invalid.

```go
var ErrApiInvalidUserID = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Invalid user ID", http.StatusBadRequest)
}
```

**Var:**

ErrApiInvalidYear returns a 400 Bad Request error when the year value is invalid.

```go
var ErrApiInvalidYear = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Invalid year value", http.StatusBadRequest)
}
```

**Var:**

ErrApiValidateCreateCard represents validation error for create card request.

```go
var ErrApiValidateCreateCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "validation failed: invalid create Card request", http.StatusBadRequest)
}
```

**Var:**

ErrApiValidateUpdateCard represents validation error for update card request.

```go
var ErrApiValidateUpdateCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "validation failed: invalid update Card request", http.StatusBadRequest)
}
```

**Var:**

ErrCardAlreadyExists is returned when a card already exists.

```go
var ErrCardAlreadyExists = errors.New("card already exists")
```

**Var:**

ErrCardNotFound is returned when a card is not found.

```go
var ErrCardNotFound = errors.New("card not found")
```

**Var:**

ErrCardNotFoundRes is an error response when a requested card was not found.
Its HTTP status code is 404.

```go
var ErrCardNotFoundRes = response.NewErrorResponse("Card not found", http.StatusNotFound)
```

**Var:**

ErrCreateCardFailed is returned when creating a new card fails.

```go
var ErrCreateCardFailed = errors.New("failed to create card")
```

**Var:**

ErrDeleteAllCardsPermanentFailed is returned when permanently deleting all cards fails.

```go
var ErrDeleteAllCardsPermanentFailed = errors.New("failed to delete all cards permanently")
```

**Var:**

ErrDeleteCardPermanentFailed is returned when permanently deleting a card fails.

```go
var ErrDeleteCardPermanentFailed = errors.New("failed to delete card permanently")
```

**Var:**

ErrFailedCreateCard is returned when creating a new Card record fails.
Its HTTP status code is 500.

```go
var ErrFailedCreateCard = response.NewErrorResponse("Failed to create Card", http.StatusInternalServerError)
```

**Var:**

ErrFailedDashboardCard is an error response when retrieving card dashboard
fails. Its HTTP status code is 500.

```go
var ErrFailedDashboardCard = response.NewErrorResponse("Failed to get Card dashboard", http.StatusInternalServerError)
```

**Var:**

ErrFailedDashboardCardNumber is an error response when retrieving card
dashboard by card number fails. Its HTTP status code is 500.

```go
var ErrFailedDashboardCardNumber = response.NewErrorResponse("Failed to get Card dashboard by card number", http.StatusInternalServerError)
```

**Var:**

ErrFailedDeleteAllCards is returned when permanently deleting all Cards fails.
Its HTTP status code is 500.

```go
var ErrFailedDeleteAllCards = response.NewErrorResponse("Failed to delete all Cards permanently", http.StatusInternalServerError)
```

**Var:**

ErrFailedDeleteCard is returned when permanently deleting a Card fails.
Its HTTP status code is 500.

```go
var ErrFailedDeleteCard = response.NewErrorResponse("Failed to delete Card permanently", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindActiveCards is an error response when retrieving active card
records fails. Its HTTP status code is 500.

```go
var ErrFailedFindActiveCards = response.NewErrorResponse("Failed to fetch active Cards", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindAllCards is an error response when retrieving all card records
fails. Its HTTP status code is 500.

```go
var ErrFailedFindAllCards = response.NewErrorResponse("Failed to fetch Cards", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindByCardNumber is an error response when finding a card by its card
number fails. Its HTTP status code is 500.

```go
var ErrFailedFindByCardNumber = response.NewErrorResponse("Failed to find Card by Card Number", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindById is an error response when finding a card by its ID fails.
Its HTTP status code is 500.

```go
var ErrFailedFindById = response.NewErrorResponse("Failed to find Card by ID", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindByUserID is an error response when finding a card by its user ID
fails. Its HTTP status code is 500.

```go
var ErrFailedFindByUserID = response.NewErrorResponse("Failed to find Card by User ID", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindMonthlyBalance indicates a failure in retrieving the monthly balance.

```go
var ErrFailedFindMonthlyBalance = response.NewErrorResponse("Failed to get monthly balance", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindMonthlyBalanceByCard returns an error response when retrieving
monthly balance by card number fails. Its HTTP status code is 500.

```go
var ErrFailedFindMonthlyBalanceByCard = response.NewErrorResponse("Failed to get monthly balance by card", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindMonthlyTopupAmount indicates a failure in retrieving the monthly top-up amount.

```go
var ErrFailedFindMonthlyTopupAmount = response.NewErrorResponse("Failed to get monthly topup amount", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindMonthlyTopupAmountByCard returns an error response when retrieving
monthly top-up amount by card number fails. Its HTTP status code is 500.

```go
var ErrFailedFindMonthlyTopupAmountByCard = response.NewErrorResponse("Failed to get monthly topup amount by card", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindMonthlyTransactionAmount indicates a failure in retrieving the monthly transaction amount.

```go
var ErrFailedFindMonthlyTransactionAmount = response.NewErrorResponse("Failed to get monthly transaction amount", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindMonthlyTransactionAmountByCard returns an error response when retrieving
monthly transaction amount by card number fails. Its HTTP status code is 500.

```go
var ErrFailedFindMonthlyTransactionAmountByCard = response.NewErrorResponse("Failed to get monthly transaction amount by card", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindMonthlyTransferAmountByReceiver returns an error response when retrieving
monthly transfer amount by receiver card number fails. Its HTTP status code is 500.

```go
var ErrFailedFindMonthlyTransferAmountByReceiver = response.NewErrorResponse("Failed to get monthly transfer amount by receiver", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindMonthlyTransferAmountBySender returns an error response when retrieving
monthly transfer amount by sender card number fails. Its HTTP status code is 500.

```go
var ErrFailedFindMonthlyTransferAmountBySender = response.NewErrorResponse("Failed to get monthly transfer amount by sender", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindMonthlyTransferAmountReceiver indicates a failure in retrieving the monthly transfer amount by receiver.

```go
var ErrFailedFindMonthlyTransferAmountReceiver = response.NewErrorResponse("Failed to get monthly transfer amount by receiver", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindMonthlyTransferAmountSender indicates a failure in retrieving the monthly transfer amount by sender.

```go
var ErrFailedFindMonthlyTransferAmountSender = response.NewErrorResponse("Failed to get monthly transfer amount by sender", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindMonthlyWithdrawAmount indicates a failure in retrieving the monthly withdraw amount.

```go
var ErrFailedFindMonthlyWithdrawAmount = response.NewErrorResponse("Failed to get monthly withdraw amount", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindMonthlyWithdrawAmountByCard returns an error response when retrieving
monthly withdraw amount by card number fails. Its HTTP status code is 500.

```go
var ErrFailedFindMonthlyWithdrawAmountByCard = response.NewErrorResponse("Failed to get monthly withdraw amount by card", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindTotalBalanceByCard is an error response when retrieving total
balance by card fails. Its HTTP status code is 500.

```go
var ErrFailedFindTotalBalanceByCard = response.NewErrorResponse("Failed to Find total balance by card", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindTotalBalances is an error response when retrieving total balances
fails. Its HTTP status code is 500.

```go
var ErrFailedFindTotalBalances = response.NewErrorResponse("Failed to Find total balances", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindTotalTopAmount is an error response when retrieving total topup
amount fails. Its HTTP status code is 500.

```go
var ErrFailedFindTotalTopAmount = response.NewErrorResponse("Failed to Find total topup amount", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindTotalTopupAmountByCard is an error response when retrieving total
topup amount by card fails. Its HTTP status code is 500.

```go
var ErrFailedFindTotalTopupAmountByCard = response.NewErrorResponse("Failed to Find total topup amount by card", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindTotalTransactionAmount is an error response when retrieving total
transaction amount fails. Its HTTP status code is 500.

```go
var ErrFailedFindTotalTransactionAmount = response.NewErrorResponse("Failed to Find total transaction amount", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindTotalTransactionAmountByCard is an error response when
retrieving total transaction amount by card fails. Its HTTP status code is
500.

```go
var ErrFailedFindTotalTransactionAmountByCard = response.NewErrorResponse("Failed to Find total transaction amount by card", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindTotalTransferAmount is an error response when retrieving total
transfer amount fails. Its HTTP status code is 500.

```go
var ErrFailedFindTotalTransferAmount = response.NewErrorResponse("Failed to Find total transfer amount", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindTotalTransferAmountByReceiver is an error response when
retrieving total transfer amount by receiver fails. Its HTTP status code is
500.

```go
var ErrFailedFindTotalTransferAmountByReceiver = response.NewErrorResponse("Failed to Find total transfer amount by receiver", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindTotalTransferAmountBySender is an error response when
retrieving total transfer amount by sender fails. Its HTTP status code is
500.

```go
var ErrFailedFindTotalTransferAmountBySender = response.NewErrorResponse("Failed to Find total transfer amount by sender", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindTotalWithdrawAmount is an error response when retrieving total
withdraw amount fails. Its HTTP status code is 500.

```go
var ErrFailedFindTotalWithdrawAmount = response.NewErrorResponse("Failed to Find total withdraw amount", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindTotalWithdrawAmountByCard is an error response when retrieving
total withdraw amount by card fails. Its HTTP status code is 500.

```go
var ErrFailedFindTotalWithdrawAmountByCard = response.NewErrorResponse("Failed to Find total withdraw amount by card", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindTrashedCards is an error response when retrieving trashed card
records fails. Its HTTP status code is 500.

```go
var ErrFailedFindTrashedCards = response.NewErrorResponse("Failed to fetch trashed Cards", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindYearlyBalance indicates a failure in retrieving the yearly balance.

```go
var ErrFailedFindYearlyBalance = response.NewErrorResponse("Failed to get yearly balance", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindYearlyBalanceByCard returns an error response when retrieving
yearly balance by card number fails. Its HTTP status code is 500.

```go
var ErrFailedFindYearlyBalanceByCard = response.NewErrorResponse("Failed to get yearly balance by card", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindYearlyTopupAmount indicates a failure in retrieving the yearly top-up amount.

```go
var ErrFailedFindYearlyTopupAmount = response.NewErrorResponse("Failed to get yearly topup amount", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindYearlyTopupAmountByCard returns an error response when retrieving
yearly top-up amount by card number fails. Its HTTP status code is 500.

```go
var ErrFailedFindYearlyTopupAmountByCard = response.NewErrorResponse("Failed to get yearly topup amount by card", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindYearlyTransactionAmount indicates a failure in retrieving the yearly transaction amount.

```go
var ErrFailedFindYearlyTransactionAmount = response.NewErrorResponse("Failed to get yearly transaction amount", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindYearlyTransactionAmountByCard returns an error response when retrieving
yearly transaction amount by card number fails. Its HTTP status code is 500.

```go
var ErrFailedFindYearlyTransactionAmountByCard = response.NewErrorResponse("Failed to get yearly transaction amount by card", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindYearlyTransferAmountByReceiver returns an error response when retrieving
yearly transfer amount by receiver card number fails. Its HTTP status code is 500.

```go
var ErrFailedFindYearlyTransferAmountByReceiver = response.NewErrorResponse("Failed to get yearly transfer amount by receiver", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindYearlyTransferAmountBySender returns an error response when retrieving
yearly transfer amount by sender card number fails. Its HTTP status code is 500.

```go
var ErrFailedFindYearlyTransferAmountBySender = response.NewErrorResponse("Failed to get yearly transfer amount by sender", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindYearlyTransferAmountReceiver indicates a failure in retrieving the yearly transfer amount by receiver.

```go
var ErrFailedFindYearlyTransferAmountReceiver = response.NewErrorResponse("Failed to get yearly transfer amount by receiver", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindYearlyTransferAmountSender indicates a failure in retrieving the yearly transfer amount by sender.

```go
var ErrFailedFindYearlyTransferAmountSender = response.NewErrorResponse("Failed to get yearly transfer amount by sender", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindYearlyWithdrawAmount indicates a failure in retrieving the yearly withdraw amount.

```go
var ErrFailedFindYearlyWithdrawAmount = response.NewErrorResponse("Failed to get yearly withdraw amount", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindYearlyWithdrawAmountByCard returns an error response when retrieving
yearly withdraw amount by card number fails. Its HTTP status code is 500.

```go
var ErrFailedFindYearlyWithdrawAmountByCard = response.NewErrorResponse("Failed to get yearly withdraw amount by card", http.StatusInternalServerError)
```

**Var:**

ErrFailedRestoreAllCards is returned when restoring all trashed Cards fails.
Its HTTP status code is 500.

```go
var ErrFailedRestoreAllCards = response.NewErrorResponse("Failed to restore all Cards", http.StatusInternalServerError)
```

**Var:**

ErrFailedRestoreCard is returned when restoring a trashed Card fails.
Its HTTP status code is 500.

```go
var ErrFailedRestoreCard = response.NewErrorResponse("Failed to restore Card", http.StatusInternalServerError)
```

**Var:**

ErrFailedTrashCard is returned when moving a Card to trash fails.
Its HTTP status code is 500.

```go
var ErrFailedTrashCard = response.NewErrorResponse("Failed to trash Card", http.StatusInternalServerError)
```

**Var:**

ErrFailedUpdateCard is returned when updating an existing Card record fails.
Its HTTP status code is 500.

```go
var ErrFailedUpdateCard = response.NewErrorResponse("Failed to update Card", http.StatusInternalServerError)
```

**Var:**

ErrFindActiveCardsFailed is returned when fetching active card records fails.

```go
var ErrFindActiveCardsFailed = errors.New("failed to find active cards")
```

**Var:**

ErrFindAllCardsFailed is returned when fetching all card records fails.

```go
var ErrFindAllCardsFailed = errors.New("failed to find all cards")
```

**Var:**

ErrFindCardByCardNumberFailed is returned when fetching a card by card number fails.

```go
var ErrFindCardByCardNumberFailed = errors.New("failed to find card by card number")
```

**Var:**

ErrFindCardByIdFailed is returned when fetching a card by its ID fails.

```go
var ErrFindCardByIdFailed = errors.New("failed to find card by ID")
```

**Var:**

ErrFindCardByUserIdFailed is returned when fetching a card by user ID fails.

```go
var ErrFindCardByUserIdFailed = errors.New("failed to find card by user ID")
```

**Var:**

ErrFindTrashedCardsFailed is returned when fetching trashed card records fails.

```go
var ErrFindTrashedCardsFailed = errors.New("failed to find trashed cards")
```

**Var:**

ErrGetMonthlyBalanceByCardFailed is returned when fetching monthly balances by card fails.

```go
var ErrGetMonthlyBalanceByCardFailed = errors.New("failed to get monthly balance by card")
```

**Var:**

ErrGetMonthlyBalanceFailed is returned when fetching monthly balances fails.

```go
var ErrGetMonthlyBalanceFailed = errors.New("failed to get monthly balance")
```

**Var:**

ErrGetMonthlyTopupAmountByCardFailed is returned when fetching monthly top-up amounts by card fails.

```go
var ErrGetMonthlyTopupAmountByCardFailed = errors.New("failed to get monthly topup amount by card")
```

**Var:**

ErrGetMonthlyTopupAmountFailed is returned when fetching monthly top-up amounts fails.

```go
var ErrGetMonthlyTopupAmountFailed = errors.New("failed to get monthly topup amount")
```

**Var:**

ErrGetMonthlyTransactionAmountByCardFailed is returned when fetching monthly transaction amounts by card fails.

```go
var ErrGetMonthlyTransactionAmountByCardFailed = errors.New("failed to get monthly transaction amount by card")
```

**Var:**

ErrGetMonthlyTransactionAmountFailed is returned when fetching monthly transaction amounts fails.

```go
var ErrGetMonthlyTransactionAmountFailed = errors.New("failed to get monthly transaction amount")
```

**Var:**

ErrGetMonthlyTransferAmountByReceiverFailed is returned when fetching monthly transfer amount by receiver fails.

```go
var ErrGetMonthlyTransferAmountByReceiverFailed = errors.New("failed to get monthly transfer amount by receiver")
```

**Var:**

ErrGetMonthlyTransferAmountBySenderFailed is returned when fetching monthly transfer amount by sender fails.

```go
var ErrGetMonthlyTransferAmountBySenderFailed = errors.New("failed to get monthly transfer amount by sender")
```

**Var:**

ErrGetMonthlyTransferAmountReceiverFailed is returned when fetching monthly transfer amounts by receiver fails.

```go
var ErrGetMonthlyTransferAmountReceiverFailed = errors.New("failed to get monthly transfer amount by receiver")
```

**Var:**

ErrGetMonthlyTransferAmountSenderFailed is returned when fetching monthly transfer amounts by sender fails.

```go
var ErrGetMonthlyTransferAmountSenderFailed = errors.New("failed to get monthly transfer amount by sender")
```

**Var:**

ErrGetMonthlyWithdrawAmountByCardFailed is returned when fetching monthly withdrawal amounts by card fails.

```go
var ErrGetMonthlyWithdrawAmountByCardFailed = errors.New("failed to get monthly withdraw amount by card")
```

**Var:**

ErrGetMonthlyWithdrawAmountFailed is returned when fetching monthly withdrawal amounts fails.

```go
var ErrGetMonthlyWithdrawAmountFailed = errors.New("failed to get monthly withdraw amount")
```

**Var:**

ErrGetTotalBalanceByCardFailed is returned when fetching the total balance by card fails.

```go
var ErrGetTotalBalanceByCardFailed = errors.New("failed to get total balance by card")
```

**Var:**

ErrGetTotalBalancesFailed is returned when fetching total balances fails.

```go
var ErrGetTotalBalancesFailed = errors.New("failed to get total balances")
```

**Var:**

ErrGetTotalTopAmountFailed is returned when fetching the total top-up amount fails.

```go
var ErrGetTotalTopAmountFailed = errors.New("failed to get total topup amount")
```

**Var:**

ErrGetTotalTopupAmountByCardFailed is returned when fetching the total top-up amount by card fails.

```go
var ErrGetTotalTopupAmountByCardFailed = errors.New("failed to get total topup amount by card")
```

**Var:**

ErrGetTotalTransactionAmountByCardFailed is returned when fetching the total transaction amount by card fails.

```go
var ErrGetTotalTransactionAmountByCardFailed = errors.New("failed to get total transaction amount by card")
```

**Var:**

ErrGetTotalTransactionAmountFailed is returned when fetching the total transaction amount fails.

```go
var ErrGetTotalTransactionAmountFailed = errors.New("failed to get total transaction amount")
```

**Var:**

ErrGetTotalTransferAmountByReceiverFailed is returned when fetching the total transfer amount by receiver fails.

```go
var ErrGetTotalTransferAmountByReceiverFailed = errors.New("failed to get total transfer amount by receiver")
```

**Var:**

ErrGetTotalTransferAmountBySenderFailed is returned when fetching the total transfer amount by sender fails.

```go
var ErrGetTotalTransferAmountBySenderFailed = errors.New("failed to get total transfer amount by sender")
```

**Var:**

ErrGetTotalTransferAmountFailed is returned when fetching the total transfer amount fails.

```go
var ErrGetTotalTransferAmountFailed = errors.New("failed to get total transfer amount")
```

**Var:**

ErrGetTotalWithdrawAmountByCardFailed is returned when fetching the total withdrawal amount by card fails.

```go
var ErrGetTotalWithdrawAmountByCardFailed = errors.New("failed to get total withdraw amount by card")
```

**Var:**

ErrGetTotalWithdrawAmountFailed is returned when fetching the total withdrawal amount fails.

```go
var ErrGetTotalWithdrawAmountFailed = errors.New("failed to get total withdraw amount")
```

**Var:**

ErrGetYearlyBalanceByCardFailed is returned when fetching yearly balances by card fails.

```go
var ErrGetYearlyBalanceByCardFailed = errors.New("failed to get yearly balance by card")
```

**Var:**

ErrGetYearlyBalanceFailed is returned when fetching yearly balances fails.

```go
var ErrGetYearlyBalanceFailed = errors.New("failed to get yearly balance")
```

**Var:**

ErrGetYearlyTopupAmountByCardFailed is returned when fetching yearly top-up amounts by card fails.

```go
var ErrGetYearlyTopupAmountByCardFailed = errors.New("failed to get yearly topup amount by card")
```

**Var:**

ErrGetYearlyTopupAmountFailed is returned when fetching yearly top-up amounts fails.

```go
var ErrGetYearlyTopupAmountFailed = errors.New("failed to get yearly topup amount")
```

**Var:**

ErrGetYearlyTransactionAmountByCardFailed is returned when fetching yearly transaction amounts by card fails.

```go
var ErrGetYearlyTransactionAmountByCardFailed = errors.New("failed to get yearly transaction amount by card")
```

**Var:**

ErrGetYearlyTransactionAmountFailed is returned when fetching yearly transaction amounts fails.

```go
var ErrGetYearlyTransactionAmountFailed = errors.New("failed to get yearly transaction amount")
```

**Var:**

ErrGetYearlyTransferAmountByReceiverFailed is returned when fetching yearly transfer amount by receiver fails.

```go
var ErrGetYearlyTransferAmountByReceiverFailed = errors.New("failed to get yearly transfer amount by receiver")
```

**Var:**

ErrGetYearlyTransferAmountBySenderFailed is returned when fetching yearly transfer amount by sender fails.

```go
var ErrGetYearlyTransferAmountBySenderFailed = errors.New("failed to get yearly transfer amount by sender")
```

**Var:**

ErrGetYearlyTransferAmountReceiverFailed is returned when fetching yearly transfer amounts by receiver fails.

```go
var ErrGetYearlyTransferAmountReceiverFailed = errors.New("failed to get yearly transfer amount by receiver")
```

**Var:**

ErrGetYearlyTransferAmountSenderFailed is returned when fetching yearly transfer amounts by sender fails.

```go
var ErrGetYearlyTransferAmountSenderFailed = errors.New("failed to get yearly transfer amount by sender")
```

**Var:**

ErrGetYearlyWithdrawAmountByCardFailed is returned when fetching yearly withdrawal amounts by card fails.

```go
var ErrGetYearlyWithdrawAmountByCardFailed = errors.New("failed to get yearly withdraw amount by card")
```

**Var:**

ErrGetYearlyWithdrawAmountFailed is returned when fetching yearly withdrawal amounts fails.

```go
var ErrGetYearlyWithdrawAmountFailed = errors.New("failed to get yearly withdraw amount")
```

**Var:**

ErrInvalidCardId is returned when the card ID is invalid.

```go
var ErrInvalidCardId = errors.New("invalid card ID")
```

**Var:**

ErrInvalidCardNumber is returned when the card number is invalid.

```go
var ErrInvalidCardNumber = errors.New("invalid card number")
```

**Var:**

ErrInvalidCardRequest is returned when the card request data is invalid.

```go
var ErrInvalidCardRequest = errors.New("invalid card request data")
```

**Var:**

ErrInvalidUserId is returned when the user ID is invalid.

```go
var ErrInvalidUserId = errors.New("invalid user ID")
```

**Var:**

ErrRestoreAllCardsFailed is returned when restoring all trashed cards fails.

```go
var ErrRestoreAllCardsFailed = errors.New("failed to restore all cards")
```

**Var:**

ErrRestoreCardFailed is returned when restoring a trashed card fails.

```go
var ErrRestoreCardFailed = errors.New("failed to restore card")
```

**Var:**

ErrTrashCardFailed is returned when trashing a card fails.

```go
var ErrTrashCardFailed = errors.New("failed to trash card")
```

**Var:**

ErrUpdateCardFailed is returned when updating a card fails.

```go
var ErrUpdateCardFailed = errors.New("failed to update card")
```

