# 📦 Package `response`

**Source Path:** `./shared/domain/response`

## 🧩 Types

### `APIResponsePagination`

APIResponsePagination is a generic paginated response wrapper.
Combines data payload with pagination metadata.

```go
type APIResponsePagination struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data T `json:"data"`
	Meta PaginationMeta `json:"pagination"`
}
```

### `ApiResponse`

ApiResponse is a generic response wrapper for successful API responses.
The type parameter T allows this to be used with any data type.

```go
type ApiResponse struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data T `json:"data"`
}
```

### `ApiResponseCard`

ApiResponseCard is the standard response format for single card requests.

```go
type ApiResponseCard struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data *CardResponse `json:"data"`
}
```

### `ApiResponseCardAll`

ApiResponseCardAll is the response format for bulk card operations.

```go
type ApiResponseCardAll struct {
	Status string `json:"status"`
	Message string `json:"message"`
}
```

### `ApiResponseCardDelete`

ApiResponseCardDelete is the response format for card deletion operations.

```go
type ApiResponseCardDelete struct {
	Status string `json:"status"`
	Message string `json:"message"`
}
```

### `ApiResponseDashboardCard`

ApiResponseDashboardCard is the response format for dashboard card statistics.

```go
type ApiResponseDashboardCard struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data *DashboardCard `json:"data"`
}
```

### `ApiResponseDashboardCardNumber`

ApiResponseDashboardCardNumber is the response format for card-specific dashboard stats.

```go
type ApiResponseDashboardCardNumber struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data *DashboardCardCardNumber `json:"data"`
}
```

### `ApiResponseForgotPassword`

ApiResponseForgotPassword is the response format for password reset initiation.
Used when requesting a password reset link.

```go
type ApiResponseForgotPassword struct {
	Status string `json:"status"`
	Message string `json:"message"`
}
```

### `ApiResponseGetMe`

ApiResponseGetMe is the response format for current user profile requests.
Contains the authenticated user's information.

```go
type ApiResponseGetMe struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data *UserResponse `json:"data"`
}
```

### `ApiResponseLogin`

ApiResponseLogin is the response format for successful login attempts.
Contains the authentication tokens.

```go
type ApiResponseLogin struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data *TokenResponse `json:"data"`
}
```

### `ApiResponseMerchant`

ApiResponseMerchant wraps a single merchant response.

```go
type ApiResponseMerchant struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data *MerchantResponse `json:"data"`
}
```

### `ApiResponseMerchantAll`

ApiResponseMerchantAll wraps bulk merchant operation responses.

```go
type ApiResponseMerchantAll struct {
	Status string `json:"status"`
	Message string `json:"message"`
}
```

### `ApiResponseMerchantDelete`

ApiResponseMerchantDelete wraps merchant deletion responses.

```go
type ApiResponseMerchantDelete struct {
	Status string `json:"status"`
	Message string `json:"message"`
}
```

### `ApiResponseMerchantDocument`

ApiResponseMerchantDocument is the standard response format for single document requests.

```go
type ApiResponseMerchantDocument struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data *MerchantDocumentResponse `json:"data"`
}
```

### `ApiResponseMerchantDocumentAll`

ApiResponseMerchantDocumentAll is the response format for bulk document operations.

```go
type ApiResponseMerchantDocumentAll struct {
	Status string `json:"status"`
	Message string `json:"message"`
}
```

### `ApiResponseMerchantDocumentDelete`

ApiResponseMerchantDocumentDelete is the response format for document deletion operations.

```go
type ApiResponseMerchantDocumentDelete struct {
	Status string `json:"status"`
	Message string `json:"message"`
}
```

### `ApiResponseMerchantMonthlyAmount`

ApiResponseMerchantMonthlyAmount wraps monthly transaction amounts.

```go
type ApiResponseMerchantMonthlyAmount struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data []*MerchantResponseMonthlyAmount `json:"data"`
}
```

### `ApiResponseMerchantMonthlyPaymentMethod`

ApiResponseMerchantMonthlyPaymentMethod wraps monthly payment method statistics.

```go
type ApiResponseMerchantMonthlyPaymentMethod struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data []*MerchantResponseMonthlyPaymentMethod `json:"data"`
}
```

### `ApiResponseMerchantMonthlyTotalAmount`

ApiResponseMerchantMonthlyTotalAmount wraps detailed monthly transaction totals.

```go
type ApiResponseMerchantMonthlyTotalAmount struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data []*MerchantResponseMonthlyTotalAmount `json:"data"`
}
```

### `ApiResponseMerchantYearlyAmount`

ApiResponseMerchantYearlyAmount wraps yearly transaction amounts.

```go
type ApiResponseMerchantYearlyAmount struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data []*MerchantResponseYearlyAmount `json:"data"`
}
```

### `ApiResponseMerchantYearlyPaymentMethod`

ApiResponseMerchantYearlyPaymentMethod wraps yearly payment method statistics.

```go
type ApiResponseMerchantYearlyPaymentMethod struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data []*MerchantResponseYearlyPaymentMethod `json:"data"`
}
```

### `ApiResponseMerchantYearlyTotalAmount`

ApiResponseMerchantYearlyTotalAmount wraps comprehensive yearly transaction totals.

```go
type ApiResponseMerchantYearlyTotalAmount struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data []*MerchantResponseYearlyTotalAmount `json:"data"`
}
```

### `ApiResponseMonthSaldoBalances`

ApiResponseMonthSaldoBalances wraps monthly balance trend responses.

```go
type ApiResponseMonthSaldoBalances struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data []*SaldoMonthBalanceResponse `json:"data"`
}
```

### `ApiResponseMonthTotalSaldo`

ApiResponseMonthTotalSaldo wraps monthly total balance responses.

```go
type ApiResponseMonthTotalSaldo struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data []*SaldoMonthTotalBalanceResponse `json:"data"`
}
```

### `ApiResponseMonthlyAmount`

ApiResponseMonthlyAmount is the response format for monthly transaction amounts.

```go
type ApiResponseMonthlyAmount struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data []*CardResponseMonthAmount `json:"data"`
}
```

### `ApiResponseMonthlyBalance`

ApiResponseMonthlyBalance is the response format for monthly balance data.

```go
type ApiResponseMonthlyBalance struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data []*CardResponseMonthBalance `json:"data"`
}
```

### `ApiResponsePaginationCard`

ApiResponsePaginationCard is the paginated response for card lists.

```go
type ApiResponsePaginationCard struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data []*CardResponse `json:"data"`
	Pagination *PaginationMeta `json:"pagination"`
}
```

### `ApiResponsePaginationCardDeleteAt`

ApiResponsePaginationCardDeleteAt is the paginated response including deleted cards.

```go
type ApiResponsePaginationCardDeleteAt struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data []*CardResponseDeleteAt `json:"data"`
	Pagination *PaginationMeta `json:"pagination"`
}
```

### `ApiResponsePaginationMerchant`

ApiResponsePaginationMerchant wraps paginated merchant lists.

```go
type ApiResponsePaginationMerchant struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data []*MerchantResponse `json:"data"`
	Pagination *PaginationMeta `json:"pagination"`
}
```

### `ApiResponsePaginationMerchantDeleteAt`

ApiResponsePaginationMerchantDeleteAt wraps paginated lists including deleted merchants.

```go
type ApiResponsePaginationMerchantDeleteAt struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data []*MerchantResponseDeleteAt `json:"data"`
	Pagination *PaginationMeta `json:"pagination"`
}
```

### `ApiResponsePaginationMerchantDocument`

ApiResponsePaginationMerchantDocument is the paginated response for document lists.

```go
type ApiResponsePaginationMerchantDocument struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data []*MerchantDocumentResponse `json:"data"`
	Pagination *PaginationMeta `json:"pagination"`
}
```

### `ApiResponsePaginationMerchantDocumentDeleteAt`

ApiResponsePaginationMerchantDocumentDeleteAt is the paginated response including deleted documents.

```go
type ApiResponsePaginationMerchantDocumentDeleteAt struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data []*MerchantDocumentResponseDeleteAt `json:"data"`
	Pagination *PaginationMeta `json:"pagination"`
}
```

### `ApiResponsePaginationMerchantTransaction`

ApiResponsePaginationMerchantTransaction wraps paginated merchant transaction lists.

```go
type ApiResponsePaginationMerchantTransaction struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data []*MerchantTransactionResponse `json:"data"`
	Pagination *PaginationMeta `json:"pagination"`
}
```

### `ApiResponsePaginationRole`

ApiResponsePaginationRole is the paginated response for role lists.

```go
type ApiResponsePaginationRole struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data []*RoleResponse `json:"data"`
	Pagination *PaginationMeta `json:"pagination"`
}
```

### `ApiResponsePaginationRoleDeleteAt`

ApiResponsePaginationRoleDeleteAt is the paginated response including deleted roles.

```go
type ApiResponsePaginationRoleDeleteAt struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data []*RoleResponseDeleteAt `json:"data"`
	Pagination *PaginationMeta `json:"pagination"`
}
```

### `ApiResponsePaginationSaldo`

ApiResponsePaginationSaldo is the paginated response for saldo lists.

```go
type ApiResponsePaginationSaldo struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data []*SaldoResponse `json:"data"`
	Pagination *PaginationMeta `json:"pagination"`
}
```

### `ApiResponsePaginationSaldoDeleteAt`

ApiResponsePaginationSaldoDeleteAt is the paginated response including deleted saldo records.

```go
type ApiResponsePaginationSaldoDeleteAt struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data []*SaldoResponseDeleteAt `json:"data"`
	Pagination *PaginationMeta `json:"pagination"`
}
```

### `ApiResponsePaginationTopup`

ApiResponsePaginationTopup is the paginated API response for top-ups

```go
type ApiResponsePaginationTopup struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data []*TopupResponse `json:"data"`
	Pagination *PaginationMeta `json:"pagination"`
}
```

### `ApiResponsePaginationTopupDeleteAt`

ApiResponsePaginationTopupDeleteAt is the paginated API response for top-ups with delete info

```go
type ApiResponsePaginationTopupDeleteAt struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data []*TopupResponseDeleteAt `json:"data"`
	Pagination *PaginationMeta `json:"pagination"`
}
```

### `ApiResponsePaginationTransaction`

ApiResponsePaginationTransaction is paginated API response for transactions

```go
type ApiResponsePaginationTransaction struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data []*TransactionResponse `json:"data"`
	Pagination *PaginationMeta `json:"pagination"`
}
```

### `ApiResponsePaginationTransactionDeleteAt`

ApiResponsePaginationTransactionDeleteAt is paginated API response for soft-deleted transactions

```go
type ApiResponsePaginationTransactionDeleteAt struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data []*TransactionResponseDeleteAt `json:"data"`
	Pagination *PaginationMeta `json:"pagination"`
}
```

### `ApiResponsePaginationTransfer`

ApiResponsePaginationTransfer is paginated API response for transfers

```go
type ApiResponsePaginationTransfer struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data []*TransferResponse `json:"data"`
	Pagination *PaginationMeta `json:"pagination"`
}
```

### `ApiResponsePaginationTransferDeleteAt`

ApiResponsePaginationTransferDeleteAt is paginated API response for soft-deleted transfers

```go
type ApiResponsePaginationTransferDeleteAt struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data []*TransferResponseDeleteAt `json:"data"`
	Pagination *PaginationMeta `json:"pagination"`
}
```

### `ApiResponsePaginationUser`

ApiResponsePaginationUser is a paginated API response for active users

```go
type ApiResponsePaginationUser struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data []*UserResponse `json:"data"`
	Pagination PaginationMeta `json:"pagination"`
}
```

### `ApiResponsePaginationUserDeleteAt`

ApiResponsePaginationUserDeleteAt is a paginated API response for soft-deleted users

```go
type ApiResponsePaginationUserDeleteAt struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data []*UserResponseDeleteAt `json:"data"`
	Pagination PaginationMeta `json:"pagination"`
}
```

### `ApiResponsePaginationWithdraw`

ApiResponsePaginationWithdraw is paginated API response for withdrawals

```go
type ApiResponsePaginationWithdraw struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data []*WithdrawResponse `json:"data"`
	Pagination PaginationMeta `json:"pagination"`
}
```

### `ApiResponsePaginationWithdrawDeleteAt`

ApiResponsePaginationWithdrawDeleteAt is paginated API response for soft-deleted withdrawals

```go
type ApiResponsePaginationWithdrawDeleteAt struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data []*WithdrawResponseDeleteAt `json:"data"`
	Pagination PaginationMeta `json:"pagination"`
}
```

### `ApiResponseRefreshToken`

ApiResponseRefreshToken is the response format for token refresh operations.
Contains new authentication tokens.

```go
type ApiResponseRefreshToken struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data *TokenResponse `json:"data"`
}
```

### `ApiResponseRegister`

ApiResponseRegister is the response format for new user registration.
Contains the created user's information.

```go
type ApiResponseRegister struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data *UserResponse `json:"data"`
}
```

### `ApiResponseResetPassword`

ApiResponseResetPassword is the response format for password reset completion.
Used after submitting a new password.

```go
type ApiResponseResetPassword struct {
	Status string `json:"status"`
	Message string `json:"message"`
}
```

### `ApiResponseRole`

ApiResponseRole is the standard response format for single role requests.

```go
type ApiResponseRole struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data *RoleResponse `json:"data"`
}
```

### `ApiResponseRoleAll`

ApiResponseRoleAll is the response format for bulk role operations.

```go
type ApiResponseRoleAll struct {
	Status string `json:"status"`
	Message string `json:"message"`
}
```

### `ApiResponseRoleDelete`

ApiResponseRoleDelete is the response format for role deletion operations.

```go
type ApiResponseRoleDelete struct {
	Status string `json:"status"`
	Message string `json:"message"`
}
```

### `ApiResponseSaldo`

ApiResponseSaldo is the standard response format for single saldo requests.

```go
type ApiResponseSaldo struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data SaldoResponse `json:"data"`
}
```

### `ApiResponseSaldoAll`

ApiResponseSaldoAll is the response format for bulk saldo operations.

```go
type ApiResponseSaldoAll struct {
	Status string `json:"status"`
	Message string `json:"message"`
}
```

### `ApiResponseSaldoDelete`

ApiResponseSaldoDelete is the response format for saldo deletion operations.

```go
type ApiResponseSaldoDelete struct {
	Status string `json:"status"`
	Message string `json:"message"`
}
```

### `ApiResponseTopup`

ApiResponseTopup is the API response for a single top-up transaction

```go
type ApiResponseTopup struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data *TopupResponse `json:"data"`
}
```

### `ApiResponseTopupAll`

ApiResponseTopupAll is a generic API response for top-up operations

```go
type ApiResponseTopupAll struct {
	Status string `json:"status"`
	Message string `json:"message"`
}
```

### `ApiResponseTopupDelete`

ApiResponseTopupDelete is the API response for a delete operation

```go
type ApiResponseTopupDelete struct {
	Status string `json:"status"`
	Message string `json:"message"`
}
```

### `ApiResponseTopupDeleteAt`

ApiResponseTopupDeleteAt is the API response for a single top-up including delete info

```go
type ApiResponseTopupDeleteAt struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data *TopupResponseDeleteAt `json:"data"`
}
```

### `ApiResponseTopupMonthAmount`

ApiResponseTopupMonthAmount is the API response for monthly top-up amount stats

```go
type ApiResponseTopupMonthAmount struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data []*TopupMonthAmountResponse `json:"data"`
}
```

### `ApiResponseTopupMonthMethod`

ApiResponseTopupMonthMethod is the API response for monthly top-up method stats

```go
type ApiResponseTopupMonthMethod struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data []*TopupMonthMethodResponse `json:"data"`
}
```

### `ApiResponseTopupMonthStatusFailed`

ApiResponseTopupMonthStatusFailed is the API response for monthly failed top-up stats

```go
type ApiResponseTopupMonthStatusFailed struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data []*TopupResponseMonthStatusFailed `json:"data"`
}
```

### `ApiResponseTopupMonthStatusSuccess`

ApiResponseTopupMonthStatusSuccess is the API response for monthly successful top-up stats

```go
type ApiResponseTopupMonthStatusSuccess struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data []*TopupResponseMonthStatusSuccess `json:"data"`
}
```

### `ApiResponseTopupYearAmount`

ApiResponseTopupYearAmount is the API response for yearly top-up amount stats

```go
type ApiResponseTopupYearAmount struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data []*TopupYearlyAmountResponse `json:"data"`
}
```

### `ApiResponseTopupYearMethod`

ApiResponseTopupYearMethod is the API response for yearly top-up method stats

```go
type ApiResponseTopupYearMethod struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data []*TopupYearlyMethodResponse `json:"data"`
}
```

### `ApiResponseTopupYearStatusFailed`

ApiResponseTopupYearStatusFailed is the API response for yearly failed top-up stats

```go
type ApiResponseTopupYearStatusFailed struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data []*TopupResponseYearStatusFailed `json:"data"`
}
```

### `ApiResponseTopupYearStatusSuccess`

ApiResponseTopupYearStatusSuccess is the API response for yearly successful top-up stats

```go
type ApiResponseTopupYearStatusSuccess struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data []*TopupResponseYearStatusSuccess `json:"data"`
}
```

### `ApiResponseTransaction`

ApiResponseTransaction is API response for a single transaction

```go
type ApiResponseTransaction struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data *TransactionResponse `json:"data"`
}
```

### `ApiResponseTransactionAll`

ApiResponseTransactionAll is generic API response for transaction operations

```go
type ApiResponseTransactionAll struct {
	Status string `json:"status"`
	Message string `json:"message"`
}
```

### `ApiResponseTransactionDelete`

ApiResponseTransactionDelete is API response for transaction deletion

```go
type ApiResponseTransactionDelete struct {
	Status string `json:"status"`
	Message string `json:"message"`
}
```

### `ApiResponseTransactionMonthAmount`

ApiResponseTransactionMonthAmount is API response for monthly amount metrics

```go
type ApiResponseTransactionMonthAmount struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data []*TransactionMonthAmountResponse `json:"data"`
}
```

### `ApiResponseTransactionMonthMethod`

ApiResponseTransactionMonthMethod is API response for monthly method metrics

```go
type ApiResponseTransactionMonthMethod struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data []*TransactionMonthMethodResponse `json:"data"`
}
```

### `ApiResponseTransactionMonthStatusFailed`

ApiResponseTransactionMonthStatusFailed is API response for monthly failed metrics

```go
type ApiResponseTransactionMonthStatusFailed struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data []*TransactionResponseMonthStatusFailed `json:"data"`
}
```

### `ApiResponseTransactionMonthStatusSuccess`

ApiResponseTransactionMonthStatusSuccess is API response for monthly success metrics

```go
type ApiResponseTransactionMonthStatusSuccess struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data []*TransactionResponseMonthStatusSuccess `json:"data"`
}
```

### `ApiResponseTransactionYearAmount`

ApiResponseTransactionYearAmount is API response for yearly amount metrics

```go
type ApiResponseTransactionYearAmount struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data []*TransactionYearlyAmountResponse `json:"data"`
}
```

### `ApiResponseTransactionYearMethod`

ApiResponseTransactionYearMethod is API response for yearly method metrics

```go
type ApiResponseTransactionYearMethod struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data []*TransactionYearMethodResponse `json:"data"`
}
```

### `ApiResponseTransactionYearStatusFailed`

ApiResponseTransactionYearStatusFailed is API response for yearly failed metrics

```go
type ApiResponseTransactionYearStatusFailed struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data []*TransactionResponseYearStatusFailed `json:"data"`
}
```

### `ApiResponseTransactionYearStatusSuccess`

ApiResponseTransactionYearStatusSuccess is API response for yearly success metrics

```go
type ApiResponseTransactionYearStatusSuccess struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data []*TransactionResponseYearStatusSuccess `json:"data"`
}
```

### `ApiResponseTransactions`

ApiResponseTransactions is API response for multiple transactions

```go
type ApiResponseTransactions struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data []*TransactionResponse `json:"data"`
}
```

### `ApiResponseTransfer`

ApiResponseTransfer is API response for a single transfer

```go
type ApiResponseTransfer struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data *TransferResponse `json:"data"`
}
```

### `ApiResponseTransferAll`

ApiResponseTransferAll is generic API response for transfer operations

```go
type ApiResponseTransferAll struct {
	Status string `json:"status"`
	Message string `json:"message"`
}
```

### `ApiResponseTransferDelete`

ApiResponseTransferDelete is API response for transfer deletion

```go
type ApiResponseTransferDelete struct {
	Status string `json:"status"`
	Message string `json:"message"`
}
```

### `ApiResponseTransferMonthAmount`

ApiResponseTransferMonthAmount is API response for monthly amount metrics

```go
type ApiResponseTransferMonthAmount struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data []*TransferMonthAmountResponse `json:"data"`
}
```

### `ApiResponseTransferMonthStatusFailed`

ApiResponseTransferMonthStatusFailed is API response for monthly failed metrics

```go
type ApiResponseTransferMonthStatusFailed struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data []*TransferResponseMonthStatusFailed `json:"data"`
}
```

### `ApiResponseTransferMonthStatusSuccess`

ApiResponseTransferMonthStatusSuccess is API response for monthly success metrics

```go
type ApiResponseTransferMonthStatusSuccess struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data []*TransferResponseMonthStatusSuccess `json:"data"`
}
```

### `ApiResponseTransferYearAmount`

ApiResponseTransferYearAmount is API response for yearly amount metrics

```go
type ApiResponseTransferYearAmount struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data []*TransferYearAmountResponse `json:"data"`
}
```

### `ApiResponseTransferYearStatusFailed`

ApiResponseTransferYearStatusFailed is API response for yearly failed metrics

```go
type ApiResponseTransferYearStatusFailed struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data []*TransferResponseYearStatusFailed `json:"data"`
}
```

### `ApiResponseTransferYearStatusSuccess`

ApiResponseTransferYearStatusSuccess is API response for yearly success metrics

```go
type ApiResponseTransferYearStatusSuccess struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data []*TransferResponseYearStatusSuccess `json:"data"`
}
```

### `ApiResponseTransfers`

ApiResponseTransfers is API response for multiple transfers

```go
type ApiResponseTransfers struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data []*TransferResponse `json:"data"`
}
```

### `ApiResponseUser`

ApiResponseUser is the API response for a single user

```go
type ApiResponseUser struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data *UserResponse `json:"data"`
}
```

### `ApiResponseUserAll`

ApiResponseUserAll is a generic API response for user operations

```go
type ApiResponseUserAll struct {
	Status string `json:"status"`
	Message string `json:"message"`
}
```

### `ApiResponseUserDelete`

ApiResponseUserDelete is the API response for user deletion

```go
type ApiResponseUserDelete struct {
	Status string `json:"status"`
	Message string `json:"message"`
}
```

### `ApiResponseUserDeleteAt`

ApiResponseUserDeleteAt is the API response for a single user including deletion info

```go
type ApiResponseUserDeleteAt struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data *UserResponseDeleteAt `json:"data"`
}
```

### `ApiResponseVerifyCode`

ApiResponseVerifyCode is the response format for verification code operations.
Used in email/SMS verification flows.

```go
type ApiResponseVerifyCode struct {
	Status string `json:"status"`
	Message string `json:"message"`
}
```

### `ApiResponseWithdraw`

ApiResponseWithdraw is API response for a single withdrawal record

```go
type ApiResponseWithdraw struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data *WithdrawResponse `json:"data"`
}
```

### `ApiResponseWithdrawAll`

ApiResponseWithdrawAll is generic API response for withdrawal operations

```go
type ApiResponseWithdrawAll struct {
	Status string `json:"status"`
	Message string `json:"message"`
}
```

### `ApiResponseWithdrawDelete`

ApiResponseWithdrawDelete is API response for withdrawal deletion

```go
type ApiResponseWithdrawDelete struct {
	Status string `json:"status"`
	Message string `json:"message"`
}
```

### `ApiResponseWithdrawMonthAmount`

ApiResponseWithdrawMonthAmount is API response for monthly withdrawal amounts

```go
type ApiResponseWithdrawMonthAmount struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data []*WithdrawMonthlyAmountResponse `json:"data"`
}
```

### `ApiResponseWithdrawMonthStatusFailed`

ApiResponseWithdrawMonthStatusFailed is API response for monthly failed metrics

```go
type ApiResponseWithdrawMonthStatusFailed struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data []*WithdrawResponseMonthStatusFailed `json:"data"`
}
```

### `ApiResponseWithdrawMonthStatusSuccess`

ApiResponseWithdrawMonthStatusSuccess is API response for monthly success metrics

```go
type ApiResponseWithdrawMonthStatusSuccess struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data []*WithdrawResponseMonthStatusSuccess `json:"data"`
}
```

### `ApiResponseWithdrawYearAmount`

ApiResponseWithdrawYearAmount is API response for yearly withdrawal amounts

```go
type ApiResponseWithdrawYearAmount struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data []*WithdrawYearlyAmountResponse `json:"data"`
}
```

### `ApiResponseWithdrawYearStatusFailed`

ApiResponseWithdrawYearStatusFailed is API response for yearly failed metrics

```go
type ApiResponseWithdrawYearStatusFailed struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data []*WithdrawResponseYearStatusFailed `json:"data"`
}
```

### `ApiResponseWithdrawYearStatusSuccess`

ApiResponseWithdrawYearStatusSuccess is API response for yearly success metrics

```go
type ApiResponseWithdrawYearStatusSuccess struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data []*WithdrawResponseYearStatusSuccess `json:"data"`
}
```

### `ApiResponseYearSaldoBalances`

ApiResponseYearSaldoBalances wraps yearly balance trend responses.

```go
type ApiResponseYearSaldoBalances struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data []*SaldoYearBalanceResponse `json:"data"`
}
```

### `ApiResponseYearTotalSaldo`

ApiResponseYearTotalSaldo wraps yearly total balance responses.

```go
type ApiResponseYearTotalSaldo struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data []*SaldoYearTotalBalanceResponse `json:"data"`
}
```

### `ApiResponseYearlyAmount`

ApiResponseYearlyAmount is the response format for yearly transaction amounts.

```go
type ApiResponseYearlyAmount struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data []*CardResponseYearAmount `json:"data"`
}
```

### `ApiResponseYearlyBalance`

ApiResponseYearlyBalance is the response format for yearly balance data.

```go
type ApiResponseYearlyBalance struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data []*CardResponseYearlyBalance `json:"data"`
}
```

### `ApiResponsesMerchant`

ApiResponsesMerchant wraps multiple merchant responses.

```go
type ApiResponsesMerchant struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data []*MerchantResponse `json:"data"`
}
```

### `ApiResponsesMerchantDocument`

ApiResponsesMerchantDocument is the response format for multiple document listings.
Used when returning arrays of document data without pagination.

```go
type ApiResponsesMerchantDocument struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data []*MerchantDocumentResponse `json:"data"`
}
```

### `ApiResponsesRole`

ApiResponsesRole is the response format for multiple role listings (non-paginated).

```go
type ApiResponsesRole struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data []*RoleResponse `json:"data"`
}
```

### `ApiResponsesSaldo`

ApiResponsesSaldo is the response format for multiple saldo records (non-paginated).

```go
type ApiResponsesSaldo struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data []*SaldoResponse `json:"data"`
}
```

### `ApiResponsesTopup`

ApiResponsesTopup is the API response for multiple top-up transactions

```go
type ApiResponsesTopup struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data []*TopupResponse `json:"data"`
}
```

### `ApiResponsesUser`

ApiResponsesUser is the API response for multiple users

```go
type ApiResponsesUser struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data []*UserResponse `json:"data"`
}
```

### `ApiResponsesWithdraw`

ApiResponsesWithdraw is API response for multiple withdrawal records

```go
type ApiResponsesWithdraw struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data []*WithdrawResponse `json:"data"`
}
```

### `CardResponse`

CardResponse represents the basic card information returned in API responses.
Used for non-sensitive card data display in most endpoints.

```go
type CardResponse struct {
	ID int `json:"id"`
	UserID int `json:"user_id"`
	CardNumber string `json:"card_number"`
	CardType string `json:"card_type"`
	ExpireDate string `json:"expire_date"`
	CVV string `json:"cvv"`
	CardProvider string `json:"card_provider"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
```

### `CardResponseDeleteAt`

CardResponseDeleteAt extends CardResponse with deletion information.
Used in administrative interfaces showing soft-deleted cards.

```go
type CardResponseDeleteAt struct {
	ID int `json:"id"`
	UserID int `json:"user_id"`
	CardNumber string `json:"card_number"`
	CardType string `json:"card_type"`
	ExpireDate string `json:"expire_date"`
	CVV string `json:"cvv"`
	CardProvider string `json:"card_provider"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt *string `json:"deleted_at"`
}
```

### `CardResponseMonthAmount`

CardResponseMonthAmount represents monthly transaction amounts.
Shows activity volume rather than balance.

```go
type CardResponseMonthAmount struct {
	Month string `json:"month"`
	TotalAmount int64 `json:"total_amount"`
}
```

### `CardResponseMonthBalance`

CardResponseMonthBalance represents monthly balance information.
Used for balance trend analysis and reporting.

```go
type CardResponseMonthBalance struct {
	Month string `json:"month"`
	TotalBalance int64 `json:"total_balance"`
}
```

### `CardResponseYearAmount`

CardResponseYearAmount represents annual transaction amounts.
Shows yearly activity volume.

```go
type CardResponseYearAmount struct {
	Year string `json:"year"`
	TotalAmount int64 `json:"total_amount"`
}
```

### `CardResponseYearlyBalance`

CardResponseYearlyBalance represents annual balance information.
Used for yearly financial summaries.

```go
type CardResponseYearlyBalance struct {
	Year string `json:"year"`
	TotalBalance int64 `json:"total_balance"`
}
```

### `DashboardCard`

DashboardCard represents aggregated card statistics for dashboard display.
Provides overview of financial activity across all cards.

```go
type DashboardCard struct {
	TotalBalance *int64 `json:"total_balance"`
	TotalTopup *int64 `json:"total_topup"`
	TotalWithdraw *int64 `json:"total_withdraw"`
	TotalTransaction *int64 `json:"total_transaction"`
	TotalTransfer *int64 `json:"total_transfer"`
}
```

### `DashboardCardCardNumber`

DashboardCardCardNumber provides detailed stats for a specific card.
Shows both sending and receiving transfer amounts separately.

```go
type DashboardCardCardNumber struct {
	TotalBalance *int64 `json:"total_balance"`
	TotalTopup *int64 `json:"total_topup"`
	TotalWithdraw *int64 `json:"total_withdraw"`
	TotalTransaction *int64 `json:"total_transaction"`
	TotalTransferSend *int64 `json:"total_transfer_send"`
	TotalTransferReceiver *int64 `json:"total_transfer_receiver"`
}
```

### `ErrorResponse`

ErrorResponse represents standardized error responses from the API.
Used to maintain consistent error reporting across endpoints.

```go
type ErrorResponse struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Code int `json:"code"`
}
```

### `MerchantDocumentResponse`

MerchantDocumentResponse represents the basic merchant document information returned in API responses.
Contains all relevant document metadata without deletion information.

```go
type MerchantDocumentResponse struct {
	ID int `json:"id"`
	MerchantID int `json:"merchant_id"`
	DocumentType string `json:"document_type"`
	DocumentURL string `json:"document_url"`
	Status string `json:"status"`
	Note string `json:"note"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
```

### `MerchantDocumentResponseDeleteAt`

MerchantDocumentResponseDeleteAt extends MerchantDocumentResponse with deletion information.
Used in administrative interfaces showing soft-deleted documents.

```go
type MerchantDocumentResponseDeleteAt struct {
	ID int `json:"id"`
	MerchantID int `json:"merchant_id"`
	DocumentType string `json:"document_type"`
	DocumentURL string `json:"document_url"`
	Status string `json:"status"`
	Note string `json:"note"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt *string `json:"deleted_at"`
}
```

### `MerchantResponse`

MerchantResponse represents the basic merchant information returned in API responses.
Contains core merchant data without deletion information.

```go
type MerchantResponse struct {
	ID int `json:"id"`
	Name string `json:"name"`
	UserID int `json:"user_id"`
	ApiKey string `json:"api_key"`
	Status string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
```

### `MerchantResponseDeleteAt`

MerchantResponseDeleteAt extends MerchantResponse with deletion information.
Used in administrative interfaces showing soft-deleted merchants.

```go
type MerchantResponseDeleteAt struct {
	ID int `json:"id"`
	Name string `json:"name"`
	UserID int `json:"user_id"`
	ApiKey string `json:"api_key"`
	Status string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt *string `json:"deleted_at"`
}
```

### `MerchantResponseMonthlyAmount`

MerchantResponseMonthlyAmount represents monthly transaction totals.
Shows aggregated transaction volume by month.

```go
type MerchantResponseMonthlyAmount struct {
	Month string `json:"month"`
	TotalAmount int `json:"total_amount"`
}
```

### `MerchantResponseMonthlyPaymentMethod`

MerchantResponseMonthlyPaymentMethod represents monthly transaction totals by payment method.
Used for analyzing payment method trends by month.

```go
type MerchantResponseMonthlyPaymentMethod struct {
	Month string `json:"month"`
	PaymentMethod string `json:"payment_method"`
	TotalAmount int `json:"total_amount"`
}
```

### `MerchantResponseMonthlyTotalAmount`

MerchantResponseMonthlyTotalAmount represents detailed monthly transaction totals.
Includes both year and month for precise period identification.

```go
type MerchantResponseMonthlyTotalAmount struct {
	Year string `json:"year"`
	Month string `json:"month"`
	TotalAmount int `json:"total_amount"`
}
```

### `MerchantResponsePayload`

MerchantResponsePayload represents the base payload structure for merchant API responses.
Used for validation and correlation purposes in merchant operations.

```go
type MerchantResponsePayload struct {
	CorrelationID string `json:"correlation_id"`
	Valid bool `json:"valid"`
	MerchantID int64 `json:"merchant_id,omitempty"`
}
```

### `MerchantResponseYearlyAmount`

MerchantResponseYearlyAmount represents annual transaction totals.
Shows aggregated transaction volume by year.

```go
type MerchantResponseYearlyAmount struct {
	Year string `json:"year"`
	TotalAmount int `json:"total_amount"`
}
```

### `MerchantResponseYearlyPaymentMethod`

MerchantResponseYearlyPaymentMethod represents annual transaction totals by payment method.
Used for analyzing yearly payment method trends.

```go
type MerchantResponseYearlyPaymentMethod struct {
	Year string `json:"year"`
	PaymentMethod string `json:"payment_method"`
	TotalAmount int `json:"total_amount"`
}
```

### `MerchantResponseYearlyTotalAmount`

MerchantResponseYearlyTotalAmount represents comprehensive annual transaction totals.

```go
type MerchantResponseYearlyTotalAmount struct {
	Year string `json:"year"`
	TotalAmount int `json:"total_amount"`
}
```

### `MerchantTransactionResponse`

MerchantTransactionResponse represents a merchant transaction record.
Contains details of transactions processed by the merchant.

```go
type MerchantTransactionResponse struct {
	ID int `json:"id"`
	CardNumber string `json:"card_number"`
	Amount int32 `json:"amount"`
	PaymentMethod string `json:"payment_method"`
	MerchantID int32 `json:"merchant_id"`
	MerchantName string `json:"merchant_name"`
	TransactionTime string `json:"transaction_time"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
```

### `PaginationMeta`

PaginationMeta contains metadata for paginated responses.
Used to provide navigation information in list endpoints.

```go
type PaginationMeta struct {
	CurrentPage int `json:"current_page"`
	PageSize int `json:"page_size"`
	TotalPages int `json:"total_pages"`
	TotalRecords int `json:"total_records"`
}
```

### `RefreshTokenResponse`

RefreshTokenResponse represents a refresh token record in API responses.
Contains the token details along with timestamps for tracking.

```go
type RefreshTokenResponse struct {
	UserID int `json:"user_id"`
	Token string `json:"token"`
	ExpiredAt string `json:"expired_at"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
```

### `RoleResponse`

RoleResponse represents basic role information in API responses.
Contains core role data without deletion information.

```go
type RoleResponse struct {
	ID int `json:"id"`
	Name string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
```

### `RoleResponseDeleteAt`

RoleResponseDeleteAt extends RoleResponse with deletion information.
Used in administrative interfaces showing soft-deleted roles.

```go
type RoleResponseDeleteAt struct {
	ID int `json:"id"`
	Name string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt *string `json:"deleted_at"`
}
```

### `RoleResponsePayload`

RoleResponsePayload represents the base payload structure for role validation responses.
Used to verify role assignments and permissions.

```go
type RoleResponsePayload struct {
	CorrelationID string `json:"correlation_id"`
	Valid bool `json:"valid"`
	RoleNames []string `json:"role_names"`
}
```

### `SaldoMonthBalanceResponse`

SaldoMonthBalanceResponse represents monthly balance information.
Used for balance trend analysis by month.

```go
type SaldoMonthBalanceResponse struct {
	Month string `json:"month"`
	TotalBalance int `json:"total_balance"`
}
```

### `SaldoMonthTotalBalanceResponse`

SaldoMonthTotalBalanceResponse represents monthly balance totals including year.
Used for detailed monthly balance reporting.

```go
type SaldoMonthTotalBalanceResponse struct {
	Month string `json:"month"`
	Year string `json:"year"`
	TotalBalance int `json:"total_balance"`
}
```

### `SaldoResponse`

SaldoResponse represents the basic saldo (balance) information in API responses.
Contains core balance data without deletion information.

```go
type SaldoResponse struct {
	ID int `json:"id"`
	CardNumber string `json:"card_number"`
	TotalBalance int `json:"total_balance"`
	WithdrawAmount int `json:"withdraw_amount"`
	WithdrawTime string `json:"withdraw_time"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
```

### `SaldoResponseDeleteAt`

SaldoResponseDeleteAt extends SaldoResponse with deletion information.
Used in administrative interfaces showing soft-deleted balance records.

```go
type SaldoResponseDeleteAt struct {
	ID int `json:"id"`
	CardNumber string `json:"card_number"`
	TotalBalance int `json:"total_balance"`
	WithdrawAmount int `json:"withdraw_amount"`
	WithdrawTime string `json:"withdraw_time"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt *string `json:"deleted_at"`
}
```

### `SaldoYearBalanceResponse`

SaldoYearBalanceResponse represents annual balance information.
Used for long-term balance trend analysis.

```go
type SaldoYearBalanceResponse struct {
	Year string `json:"year"`
	TotalBalance int `json:"total_balance"`
}
```

### `SaldoYearTotalBalanceResponse`

SaldoYearTotalBalanceResponse represents annual balance totals.
Used for yearly financial summaries.

```go
type SaldoYearTotalBalanceResponse struct {
	Year string `json:"year"`
	TotalBalance int `json:"total_balance"`
}
```

### `TokenResponse`

TokenResponse represents a pair of authentication tokens.
Used in login and token refresh operations.

```go
type TokenResponse struct {
	AccessToken string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
```

### `TopupMonthAmountResponse`

TopupMonthAmountResponse represents monthly top-up amount statistics

```go
type TopupMonthAmountResponse struct {
	Month string `json:"month"`
	TotalAmount int `json:"total_amount"`
}
```

### `TopupMonthMethodResponse`

TopupMonthMethodResponse represents monthly top-up statistics by payment method

```go
type TopupMonthMethodResponse struct {
	Month string `json:"month"`
	TopupMethod string `json:"topup_method"`
	TotalTopups int `json:"total_topups"`
	TotalAmount int `json:"total_amount"`
}
```

### `TopupResponse`

TopupResponse represents a successful top-up transaction response

```go
type TopupResponse struct {
	ID int `json:"id"`
	CardNumber string `json:"card_number"`
	TopupNo string `json:"topup_no"`
	TopupAmount int `json:"topup_amount"`
	TopupMethod string `json:"topup_method"`
	TopupTime string `json:"topup_time"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
```

### `TopupResponseDeleteAt`

TopupResponseDeleteAt extends TopupResponse with soft delete information

```go
type TopupResponseDeleteAt struct {
	DeletedAt *string `json:"deleted_At"`
	ID int `json:"id"`
	CardNumber string `json:"card_number"`
	TopupNo string `json:"topup_no"`
	TopupAmount int `json:"topup_amount"`
	TopupMethod string `json:"topup_method"`
	TopupTime string `json:"topup_time"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
```

### `TopupResponseMonthStatusFailed`

TopupResponseMonthStatusFailed represents monthly failed top-up statistics

```go
type TopupResponseMonthStatusFailed struct {
	Year string `json:"year"`
	Month string `json:"month"`
	TotalAmount int `json:"total_amount"`
	TotalFailed int `json:"total_failed"`
}
```

### `TopupResponseMonthStatusSuccess`

TopupResponseMonthStatusSuccess represents monthly successful top-up statistics

```go
type TopupResponseMonthStatusSuccess struct {
	Year string `json:"year"`
	Month string `json:"month"`
	TotalAmount int `json:"total_amount"`
	TotalSuccess int `json:"total_success"`
}
```

### `TopupResponseYearStatusFailed`

TopupResponseYearStatusFailed represents yearly failed top-up statistics

```go
type TopupResponseYearStatusFailed struct {
	Year string `json:"year"`
	TotalAmount int `json:"total_amount"`
	TotalFailed int `json:"total_failed"`
}
```

### `TopupResponseYearStatusSuccess`

TopupResponseYearStatusSuccess represents yearly successful top-up statistics

```go
type TopupResponseYearStatusSuccess struct {
	Year string `json:"year"`
	TotalAmount int `json:"total_amount"`
	TotalSuccess int `json:"total_success"`
}
```

### `TopupYearlyAmountResponse`

TopupYearlyAmountResponse represents yearly top-up amount statistics

```go
type TopupYearlyAmountResponse struct {
	Year string `json:"year"`
	TotalAmount int `json:"total_amount"`
}
```

### `TopupYearlyMethodResponse`

TopupYearlyMethodResponse represents yearly top-up statistics by payment method

```go
type TopupYearlyMethodResponse struct {
	Year string `json:"year"`
	TopupMethod string `json:"topup_method"`
	TotalTopups int `json:"total_topups"`
	TotalAmount int `json:"total_amount"`
}
```

### `TransactionMonthAmountResponse`

TransactionMonthAmountResponse represents monthly transaction amount metrics

```go
type TransactionMonthAmountResponse struct {
	Month string `json:"month"`
	TotalAmount int `json:"total_amount"`
}
```

### `TransactionMonthMethodResponse`

TransactionMonthMethodResponse represents monthly transaction metrics by payment method

```go
type TransactionMonthMethodResponse struct {
	Month string `json:"month"`
	PaymentMethod string `json:"payment_method"`
	TotalTransactions int `json:"total_transactions"`
	TotalAmount int `json:"total_amount"`
}
```

### `TransactionResponse`

TransactionResponse represents a transaction record

```go
type TransactionResponse struct {
	ID int `json:"id"`
	TransactionNo string `json:"transaction_no"`
	CardNumber string `json:"card_number"`
	Amount int `json:"amount"`
	PaymentMethod string `json:"payment_method"`
	MerchantID int `json:"merchant_id"`
	TransactionTime string `json:"transaction_time"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
```

### `TransactionResponseDeleteAt`

TransactionResponseDeleteAt extends TransactionResponse with soft delete capability

```go
type TransactionResponseDeleteAt struct {
	DeletedAt *string `json:"deleted_at"`
	ID int `json:"id"`
	TransactionNo string `json:"transaction_no"`
	CardNumber string `json:"card_number"`
	Amount int `json:"amount"`
	PaymentMethod string `json:"payment_method"`
	MerchantID int `json:"merchant_id"`
	TransactionTime string `json:"transaction_time"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
```

### `TransactionResponseMonthStatusFailed`

TransactionResponseMonthStatusFailed represents monthly failed transaction metrics

```go
type TransactionResponseMonthStatusFailed struct {
	Year string `json:"year"`
	Month string `json:"month"`
	TotalAmount int `json:"total_amount"`
	TotalFailed int `json:"total_failed"`
}
```

### `TransactionResponseMonthStatusSuccess`

TransactionResponseMonthStatusSuccess represents monthly successful transaction metrics

```go
type TransactionResponseMonthStatusSuccess struct {
	Year string `json:"year"`
	Month string `json:"month"`
	TotalAmount int `json:"total_amount"`
	TotalSuccess int `json:"total_success"`
}
```

### `TransactionResponseYearStatusFailed`

TransactionResponseYearStatusFailed represents yearly failed transaction metrics

```go
type TransactionResponseYearStatusFailed struct {
	Year string `json:"year"`
	TotalAmount int `json:"total_amount"`
	TotalFailed int `json:"total_failed"`
}
```

### `TransactionResponseYearStatusSuccess`

TransactionResponseYearStatusSuccess represents yearly successful transaction metrics

```go
type TransactionResponseYearStatusSuccess struct {
	Year string `json:"year"`
	TotalAmount int `json:"total_amount"`
	TotalSuccess int `json:"total_success"`
}
```

### `TransactionYearMethodResponse`

TransactionYearMethodResponse represents yearly transaction metrics by payment method

```go
type TransactionYearMethodResponse struct {
	Year string `json:"year"`
	PaymentMethod string `json:"payment_method"`
	TotalTransactions int `json:"total_transactions"`
	TotalAmount int `json:"total_amount"`
}
```

### `TransactionYearlyAmountResponse`

TransactionYearlyAmountResponse represents yearly transaction amount metrics

```go
type TransactionYearlyAmountResponse struct {
	Year string `json:"year"`
	TotalAmount int `json:"total_amount"`
}
```

### `TransferMonthAmountResponse`

TransferMonthAmountResponse represents monthly transfer amount metrics

```go
type TransferMonthAmountResponse struct {
	Month string `json:"month"`
	TotalAmount int `json:"total_amount"`
}
```

### `TransferResponse`

TransferResponse represents a fund transfer between accounts

```go
type TransferResponse struct {
	ID int `json:"id"`
	TransferNo string `json:"transfer_no"`
	TransferFrom string `json:"transfer_from"`
	TransferTo string `json:"transfer_to"`
	TransferAmount int `json:"transfer_amount"`
	TransferTime string `json:"transfer_time"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
```

### `TransferResponseDeleteAt`

TransferResponseDeleteAt extends TransferResponse with soft delete capability

```go
type TransferResponseDeleteAt struct {
	DeletedAt *string `json:"deleted_at"`
	ID int `json:"id"`
	TransferNo string `json:"transfer_no"`
	TransferFrom string `json:"transfer_from"`
	TransferTo string `json:"transfer_to"`
	TransferAmount int `json:"transfer_amount"`
	TransferTime string `json:"transfer_time"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
```

### `TransferResponseMonthStatusFailed`

TransferResponseMonthStatusFailed represents monthly failed transfer metrics

```go
type TransferResponseMonthStatusFailed struct {
	Year string `json:"year"`
	Month string `json:"month"`
	TotalAmount int `json:"total_amount"`
	TotalFailed int `json:"total_failed"`
}
```

### `TransferResponseMonthStatusSuccess`

TransferResponseMonthStatusSuccess represents monthly successful transfer metrics

```go
type TransferResponseMonthStatusSuccess struct {
	Year string `json:"year"`
	Month string `json:"month"`
	TotalAmount int `json:"total_amount"`
	TotalSuccess int `json:"total_success"`
}
```

### `TransferResponseYearStatusFailed`

TransferResponseYearStatusFailed represents yearly failed transfer metrics

```go
type TransferResponseYearStatusFailed struct {
	Year string `json:"year"`
	TotalAmount int `json:"total_amount"`
	TotalFailed int `json:"total_failed"`
}
```

### `TransferResponseYearStatusSuccess`

TransferResponseYearStatusSuccess represents yearly successful transfer metrics

```go
type TransferResponseYearStatusSuccess struct {
	Year string `json:"year"`
	TotalAmount int `json:"total_amount"`
	TotalSuccess int `json:"total_success"`
}
```

### `TransferYearAmountResponse`

TransferYearAmountResponse represents yearly transfer amount metrics

```go
type TransferYearAmountResponse struct {
	Year string `json:"year"`
	TotalAmount int `json:"total_amount"`
}
```

### `UserResponse`

UserResponse represents a user account in the system

```go
type UserResponse struct {
	ID int `json:"id"`
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
	Email string `json:"email"`
	IsVerified bool `json:"is_verified"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
```

### `UserResponseDeleteAt`

UserResponseDeleteAt extends UserResponse with soft delete capability

```go
type UserResponseDeleteAt struct {
	DeletedAt *string `json:"deleted_at"`
	ID int `json:"id"`
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
	Email string `json:"email"`
	IsVerified bool `json:"is_verified"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
```

### `WithdrawMonthlyAmountResponse`

WithdrawMonthlyAmountResponse represents monthly withdrawal amount totals

```go
type WithdrawMonthlyAmountResponse struct {
	Month string `json:"month"`
	TotalAmount int `json:"total_amount"`
}
```

### `WithdrawResponse`

WithdrawResponse represents a withdrawal transaction record

```go
type WithdrawResponse struct {
	ID int `json:"id"`
	WithdrawNo string `json:"withdraw_no"`
	CardNumber string `json:"card_number"`
	WithdrawAmount int `json:"withdraw_amount"`
	WithdrawTime string `json:"withdraw_time"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
```

### `WithdrawResponseDeleteAt`

WithdrawResponseDeleteAt extends WithdrawResponse with soft-delete capability

```go
type WithdrawResponseDeleteAt struct {
	DeletedAt *string `json:"deleted_at"`
	ID int `json:"id"`
	WithdrawNo string `json:"withdraw_no"`
	CardNumber string `json:"card_number"`
	WithdrawAmount int `json:"withdraw_amount"`
	WithdrawTime string `json:"withdraw_time"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
```

### `WithdrawResponseMonthStatusFailed`

WithdrawResponseMonthStatusFailed represents monthly failed withdrawal metrics

```go
type WithdrawResponseMonthStatusFailed struct {
	Year string `json:"year"`
	Month string `json:"month"`
	TotalAmount int `json:"total_amount"`
	TotalFailed int `json:"total_failed"`
}
```

### `WithdrawResponseMonthStatusSuccess`

WithdrawResponseMonthStatusSuccess represents monthly successful withdrawal metrics

```go
type WithdrawResponseMonthStatusSuccess struct {
	Year string `json:"year"`
	Month string `json:"month"`
	TotalAmount int `json:"total_amount"`
	TotalSuccess int `json:"total_success"`
}
```

### `WithdrawResponseYearStatusFailed`

WithdrawResponseYearStatusFailed represents yearly failed withdrawal metrics

```go
type WithdrawResponseYearStatusFailed struct {
	Year string `json:"year"`
	TotalAmount int `json:"total_amount"`
	TotalFailed int `json:"total_failed"`
}
```

### `WithdrawResponseYearStatusSuccess`

WithdrawResponseYearStatusSuccess represents yearly successful withdrawal metrics

```go
type WithdrawResponseYearStatusSuccess struct {
	Year string `json:"year"`
	TotalAmount int `json:"total_amount"`
	TotalSuccess int `json:"total_success"`
}
```

### `WithdrawYearlyAmountResponse`

WithdrawYearlyAmountResponse represents yearly withdrawal amount totals

```go
type WithdrawYearlyAmountResponse struct {
	Year string `json:"year"`
	TotalAmount int `json:"total_amount"`
}
```

## 🚀 Functions

### `NewApiErrorResponse`

NewApiErrorResponse creates and sends a JSON error response using Echo framework.
This is a convenience wrapper for returning standardized error responses in HTTP handlers.

Parameters:
  - c: Echo context
  - statusText: Status category ("error", "fail", etc.)
  - message: Human-readable error description
  - code: HTTP status code

Returns:
  - error: Echo error that will trigger the JSON response

```go
func NewApiErrorResponse(c echo.Context, statusText string, message string, code int) error
```

### `NewGrpcError`

NewGrpcError creates a new gRPC error with standardized format.
This is used to maintain consistent error responses across gRPC services.

Parameters:
  - statusText: Status category ("error", "fail", etc.)
  - message: Human-readable error description
  - code: gRPC status code

Returns:
  - error: gRPC status error with embedded error details

```go
func NewGrpcError(statusText string, message string, code int) error
```

### `ToGrpcErrorFromErrorResponse`

ToGrpcErrorFromErrorResponse converts an ErrorResponse to a gRPC error.
This enables consistent error handling between HTTP and gRPC interfaces.

Parameters:
  - err: ErrorResponse to convert

Returns:
  - error: gRPC status error with embedded error details
  - nil: if input error is nil

```go
func ToGrpcErrorFromErrorResponse(err *ErrorResponse) error
```