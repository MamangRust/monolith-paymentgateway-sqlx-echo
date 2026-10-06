# 📦 Package `requests`

**Source Path:** `./shared/domain/requests`

## 🧩 Types

### `AuthRequest`

AuthRequest represents the payload for user authentication (login).
Used when a user attempts to log in to the system.

```go
type AuthRequest struct {
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}
```

#### Methods

##### `Validate`

Validate validates the AuthRequest struct using the validator.v10 package.

Returns an error if any of the validation rules fail.

```go
func (r *AuthRequest) Validate() error
```

### `CreateCardRequest`

CreateCardRequest represents the payload for creating a new payment card.
Contains all necessary information to register a payment card in the system.

```go
type CreateCardRequest struct {
	UserID int `json:"user_id"`
	CardType string `json:"card_type" validate:"required"`
	ExpireDate time.Time `json:"expire_date" validate:"required"`
	CVV string `json:"cvv" validate:"required"`
	CardProvider string `json:"card_provider" validate:"required"`
}
```

#### Methods

##### `Validate`

Validate performs validation of CreateCardRequest fields beyond basic struct validation.
Ensures card type and provider are valid values.
Returns:
  - error: if validation fails, describing the specific validation error

```go
func (r *CreateCardRequest) Validate() error
```

### `CreateMerchantDocumentRequest`

CreateMerchantDocumentRequest represents the payload for uploading a new merchant document.
Contains all necessary information to register a document for a merchant.

```go
type CreateMerchantDocumentRequest struct {
	MerchantID int `json:"merchant_id" validate:"required,min=1"`
	DocumentType string `json:"document_type" validate:"required"`
	DocumentUrl string `json:"document_url" validate:"required"`
}
```

#### Methods

##### `Validate`

Validate performs basic validation of CreateMerchantDocumentRequest fields using struct tags.
Returns:
  - error: if validation fails according to struct tag rules

```go
func (r *CreateMerchantDocumentRequest) Validate() error
```

### `CreateMerchantRequest`

CreateMerchantRequest represents the payload for registering a new merchant.
Used in merchant onboarding processes.

```go
type CreateMerchantRequest struct {
	Name string `json:"name" validate:"required"`
	UserID int `json:"user_id" validate:"required,min=1"`
}
```

#### Methods

##### `Validate`

Validate performs validation of CreateMerchantRequest fields.
Ensures all required fields are present and valid.
Returns:
  - error: if validation fails

```go
func (r CreateMerchantRequest) Validate() error
```

### `CreateRefreshToken`

CreateRefreshToken represents the payload for generating a new refresh token.
Used when creating persistent authentication tokens for users.

```go
type CreateRefreshToken struct {
	UserId int `json:"user_id" validate:"required,min=1"`
	Token string `json:"token" validate:"required,min=1"`
	ExpiresAt string `json:"expires_at" validate:"required,min=1"`
}
```

#### Methods

##### `Validate`

Validate performs structural validation of CreateRefreshToken fields.
Checks that all required fields meet minimum requirements.
Returns:
  - error: if validation fails according to struct tag rules

```go
func (r *CreateRefreshToken) Validate() error
```

### `CreateResetPasswordRequest`

CreateResetPasswordRequest represents the payload for actually resetting a user's password.
Used when submitting a new password during the reset flow.

```go
type CreateResetPasswordRequest struct {
	ResetToken string `json:"reset_token" validate:"required"`
	Password string `json:"password" validate:"required,min=6"`
	ConfirmPassword string `json:"confirm_password" validate:"required,min=6"`
}
```

#### Methods

##### `Validate`

Validate performs validation of CreateResetPasswordRequest fields.
Ensures the reset token is present and passwords meet complexity requirements.
Returns:
  - error: if validation fails, describing the validation failure

```go
func (r *CreateResetPasswordRequest) Validate() error
```

### `CreateResetTokenRequest`

CreateResetTokenRequest represents the payload for generating a password reset token.
Used when initiating a password reset process for a user.

```go
type CreateResetTokenRequest struct {
	UserID int `json:"user_id" validate:"required"`
	ResetToken string `json:"reset_token" validate:"required"`
	ExpiredAt string `json:"expired_at" validate:"required"`
}
```

### `CreateRoleRequest`

CreateRoleRequest represents the payload for creating a new role.
Used when defining new access control roles in the system.

```go
type CreateRoleRequest struct {
	Name string `json:"name" validate:"required"`
}
```

#### Methods

##### `Validate`

Validate performs validation of CreateRoleRequest fields.
Ensures the role name is provided and meets requirements.
Returns:
  - error: if validation fails, describing the validation failure

```go
func (r *CreateRoleRequest) Validate() error
```

### `CreateSaldoRequest`

CreateSaldoRequest represents the payload for initializing a new saldo record.
Used when creating a new balance entry for a card.

```go
type CreateSaldoRequest struct {
	CardNumber string `json:"card_number" validate:"required"`
	TotalBalance int `json:"total_balance" validate:"required"`
}
```

#### Methods

##### `Validate`

Validate performs basic validation of CreateSaldoRequest fields.
Ensures required fields are present and properly formatted.
Returns:
  - error: if validation fails

```go
func (r *CreateSaldoRequest) Validate() error
```

### `CreateTopupRequest`

CreateTopupRequest represents the payload for creating a new top-up transaction.
Used when adding funds to a card/account.

```go
type CreateTopupRequest struct {
	CardNumber string `json:"card_number" validate:"required,min=1"`
	TopupAmount int `json:"topup_amount" validate:"required,min=50000"`
	TopupMethod string `json:"topup_method" validate:"required"`
}
```

#### Methods

##### `Validate`

Validate performs comprehensive validation of CreateTopupRequest fields.
Ensures:
- Minimum top-up amount (50,000)
- Valid payment method
- Required fields are present
Returns:
  - error: if validation fails with specific validation messages

```go
func (r *CreateTopupRequest) Validate() error
```

### `CreateTransactionRequest`

CreateTransactionRequest represents the payload for creating a new transaction.
Used when recording payment transactions.

```go
type CreateTransactionRequest struct {
	CardNumber string `json:"card_number" validate:"required,min=1"`
	Amount int `json:"amount" validate:"required,min=50000"`
	PaymentMethod string `json:"payment_method" validate:"required"`
	MerchantID *int `json:"merchant_id" validate:"required,min=1"`
	TransactionTime time.Time `json:"transaction_time" validate:"required"`
}
```

#### Methods

##### `Validate`

Validate performs comprehensive validation of CreateTransactionRequest fields.
Ensures:
- Valid payment method
- Minimum transaction amount (50,000)
- Required fields are present
Returns:
  - error: if validation fails with specific validation messages

```go
func (r *CreateTransactionRequest) Validate() error
```

### `CreateTransferRequest`

CreateTransferRequest represents the payload for initiating a new transfer.
Used when moving funds between accounts/cards.

```go
type CreateTransferRequest struct {
	TransferFrom string `json:"transfer_from" validate:"required"`
	TransferTo string `json:"transfer_to" validate:"required,min=1"`
	TransferAmount int `json:"transfer_amount" validate:"required,min=50000"`
}
```

#### Methods

##### `Validate`

Validate performs comprehensive validation of CreateTransferRequest fields.
Ensures:
- Minimum transfer amount (50,000)
- Valid account/card numbers
- Required fields are present
Returns:
  - error: if validation fails with specific validation messages

```go
func (r *CreateTransferRequest) Validate() error
```

### `CreateUserRequest`

CreateUserRequest represents the payload for registering a new user.
Contains all required information for user account creation.

```go
type CreateUserRequest struct {
	FirstName string `json:"firstname" validate:"required,alpha"`
	LastName string `json:"lastname" validate:"required,alpha"`
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=Password"`
}
```

#### Methods

##### `Validate`

Validate performs structural validation of CreateUserRequest fields.
Ensures all required fields are present and meet format requirements.
Returns:
  - error: if validation fails according to struct tag rules

```go
func (r *CreateUserRequest) Validate() error
```

### `CreateUserRoleRequest`

CreateUserRoleRequest represents the payload for assigning a role to a user.
Used when granting specific permissions or access levels to users.

```go
type CreateUserRoleRequest struct {
	UserId int `json:"user_id" validate:"required"`
	RoleId int `json:"role_id" validate:"required"`
}
```

### `CreateWithdrawRequest`

CreateWithdrawRequest represents the payload for creating a new withdrawal.
Used when processing cash withdrawals from cards/accounts.

```go
type CreateWithdrawRequest struct {
	CardNumber string `json:"card_number" validate:"required,min=1"`
	WithdrawAmount int `json:"withdraw_amount" validate:"required,min=50000"`
	WithdrawTime time.Time `json:"withdraw_time" validate:"required"`
}
```

#### Methods

##### `Validate`

Validate performs comprehensive validation of CreateWithdrawRequest fields.
Ensures:
- Minimum withdrawal amount (50,000)
- Withdrawal time is not in the future
- Required fields are present
Returns:
  - error: if validation fails with specific validation messages

```go
func (r *CreateWithdrawRequest) Validate() error
```

### `FindAllCards`

FindAllCards represents the request parameters for searching and paginating card records.
Used in card listing operations with search functionality.

```go
type FindAllCards struct {
	Search string `json:"search" validate:"required"`
	Page int `json:"page" validate:"min=1"`
	PageSize int `json:"page_size" validate:"min=1,max=100"`
}
```

### `FindAllMerchantDocuments`

FindAllMerchantDocuments represents the request parameters for searching and paginating merchant documents.
Used in document listing operations with search functionality.

```go
type FindAllMerchantDocuments struct {
	Search string `json:"search" validate:"required"`
	Page int `json:"page" validate:"min=1"`
	PageSize int `json:"page_size" validate:"min=1,max=100"`
}
```

### `FindAllMerchantTransactions`

FindAllMerchantTransactions represents parameters for searching merchant transactions.
Used in global transaction reporting.

```go
type FindAllMerchantTransactions struct {
	Search string `json:"search" validate:"required"`
	Page int `json:"page" validate:"min=1"`
	PageSize int `json:"page_size" validate:"min=1,max=100"`
}
```

### `FindAllMerchantTransactionsByApiKey`

FindAllMerchantTransactionsByApiKey represents parameters for searching transactions by API key.
Used in merchant-facing transaction reporting.

```go
type FindAllMerchantTransactionsByApiKey struct {
	ApiKey string `json:"api_key" validate:"required"`
	Search string `json:"search" validate:"required"`
	Page int `json:"page" validate:"min=1"`
	PageSize int `json:"page_size" validate:"min=1,max=100"`
}
```

### `FindAllMerchantTransactionsById`

FindAllMerchantTransactionsById represents parameters for searching transactions by merchant ID.
Used for merchant-specific transaction reporting.

```go
type FindAllMerchantTransactionsById struct {
	MerchantID int `json:"merchant_id" validate:"required,min=1"`
	Search string `json:"search" validate:"required"`
	Page int `json:"page" validate:"min=1"`
	PageSize int `json:"page_size" validate:"min=1,max=100"`
}
```

### `FindAllMerchants`

FindAllMerchants represents parameters for searching and paginating merchant records.
Used in merchant administration interfaces.

```go
type FindAllMerchants struct {
	Search string `json:"search" validate:"required"`
	Page int `json:"page" validate:"min=1"`
	PageSize int `json:"page_size" validate:"min=1,max=100"`
}
```

### `FindAllRoles`

FindAllRoles represents parameters for searching and paginating role records.
Used in role administration and listing operations.

```go
type FindAllRoles struct {
	Search string `json:"search" validate:"required"`
	Page int `json:"page" validate:"min=1"`
	PageSize int `json:"page_size" validate:"min=1,max=100"`
}
```

### `FindAllSaldos`

FindAllSaldos represents parameters for searching and paginating saldo records.
Used to retrieve and filter saldo/balance information with pagination support.

```go
type FindAllSaldos struct {
	Search string `json:"search" validate:"required"`
	Page int `json:"page" validate:"min=1"`
	PageSize int `json:"page_size" validate:"min=1,max=100"`
}
```

### `FindAllTopups`

FindAllTopups represents parameters for searching and paginating top-up records.
Used in top-up administration and reporting interfaces.

```go
type FindAllTopups struct {
	Search string `json:"search" validate:"required"`
	Page int `json:"page" validate:"min=1"`
	PageSize int `json:"page_size" validate:"min=1,max=100"`
}
```

### `FindAllTopupsByCardNumber`

FindAllTopupsByCardNumber represents parameters for searching top-ups by card number.
Used to retrieve top-up history for specific cards.

```go
type FindAllTopupsByCardNumber struct {
	CardNumber string `json:"card_number" validate:"required,min=1"`
	Search string `json:"search" validate:"required"`
	Page int `json:"page" validate:"min=1"`
	PageSize int `json:"page_size" validate:"min=1,max=100"`
}
```

### `FindAllTransactionCardNumber`

FindAllTransactionCardNumber represents parameters for searching transactions by card number.
Used to retrieve transaction history for specific cards.

```go
type FindAllTransactionCardNumber struct {
	CardNumber string `json:"card_number" validate:"required,min=1"`
	Search string `json:"search" validate:"required"`
	Page int `json:"page" validate:"min=1"`
	PageSize int `json:"page_size" validate:"min=1,max=100"`
}
```

### `FindAllTransactions`

FindAllTransactions represents parameters for searching and paginating transaction records.
Used in transaction administration interfaces.

```go
type FindAllTransactions struct {
	Search string `json:"search" validate:"required"`
	Page int `json:"page" validate:"min=1"`
	PageSize int `json:"page_size" validate:"min=1,max=100"`
}
```

### `FindAllTransfers`

FindAllTransfers represents parameters for searching and paginating transfer records.
Used in transfer administration interfaces.

```go
type FindAllTransfers struct {
	Search string `json:"search" validate:"required"`
	Page int `json:"page" validate:"min=1"`
	PageSize int `json:"page_size" validate:"min=1,max=100"`
}
```

### `FindAllUsers`

FindAllUsers represents parameters for searching and paginating user records.
Used in user administration interfaces to list and filter users.

```go
type FindAllUsers struct {
	Search string `json:"search" validate:"required"`
	Page int `json:"page" validate:"min=1"`
	PageSize int `json:"page_size" validate:"min=1,max=100"`
}
```

### `FindAllWithdrawCardNumber`

FindAllWithdrawCardNumber represents parameters for searching withdrawals by card number.
Used to retrieve withdrawal history for specific cards.

```go
type FindAllWithdrawCardNumber struct {
	CardNumber string `json:"card_number" validate:"required,min=1"`
	Search string `json:"search" validate:"required"`
	Page int `json:"page" validate:"min=1"`
	PageSize int `json:"page_size" validate:"min=1,max=100"`
}
```

### `FindAllWithdraws`

FindAllWithdraws represents parameters for searching and paginating withdrawal records.
Used in withdrawal administration interfaces.

```go
type FindAllWithdraws struct {
	Search string `json:"search" validate:"required"`
	Page int `json:"page" validate:"min=1"`
	PageSize int `json:"page_size" validate:"min=1,max=100"`
}
```

### `ForgotPasswordRequest`

ForgotPasswordRequest represents the initial forgot password request payload.
Used when a user initiates the password reset process via email.

```go
type ForgotPasswordRequest struct {
	Email string `json:"email" validate:"required,email"`
}
```

#### Methods

##### `Validate`

Validate performs validation of ForgotPasswordRequest fields.
Ensures the email field is present and properly formatted.
Returns:
  - error: if validation fails, describing the validation failure

```go
func (r *ForgotPasswordRequest) Validate() error
```

### `JWTToken`

```go
type JWTToken struct {
	Token string `json:"token"`
}
```

### `MerchantRequestPayload`

MerchantRequestPayload represents the base payload structure for merchant API requests.
Used as a foundation for authenticated merchant operations.

```go
type MerchantRequestPayload struct {
	ApiKey string `json:"api_key"`
	CorrelationID string `json:"correlation_id"`
	ReplyTopic string `json:"reply_topic"`
}
```

### `MonthStatusTransaction`

MonthStatusTransaction represents parameters for retrieving transaction status by month.
Used to get monthly transaction statistics.

```go
type MonthStatusTransaction struct {
	Year int `json:"year" validate:"required"`
	Month int `json:"month" validate:"required"`
}
```

### `MonthStatusTransactionCardNumber`

MonthStatusTransactionCardNumber represents parameters for retrieving card-specific monthly transaction status.
Used to get monthly transaction statistics for specific cards.

```go
type MonthStatusTransactionCardNumber struct {
	CardNumber string `json:"card_number" validate:"required,min=1"`
	Year int `json:"year" validate:"required"`
	Month int `json:"month" validate:"required"`
}
```

### `MonthStatusTransfer`

MonthStatusTransfer represents parameters for retrieving monthly transfer status.
Used for general monthly transfer statistics.

```go
type MonthStatusTransfer struct {
	Year int `json:"year" validate:"required"`
	Month int `json:"month" validate:"required"`
}
```

### `MonthStatusTransferCardNumber`

MonthStatusTransferCardNumber represents parameters for retrieving monthly transfer status by card.
Used to get monthly transfer statistics for specific cards.

```go
type MonthStatusTransferCardNumber struct {
	CardNumber string `json:"card_number" validate:"required,min=1"`
	Year int `json:"year" validate:"required"`
	Month int `json:"month" validate:"required"`
}
```

### `MonthStatusWithdraw`

MonthStatusWithdraw represents parameters for retrieving monthly withdrawal status.
Used to get general monthly withdrawal statistics.

```go
type MonthStatusWithdraw struct {
	Year int `json:"year" validate:"required"`
	Month int `json:"month" validate:"required"`
}
```

### `MonthStatusWithdrawCardNumber`

MonthStatusWithdrawCardNumber represents parameters for retrieving card-specific monthly withdrawal status.
Used to get monthly withdrawal statistics for specific cards.

```go
type MonthStatusWithdrawCardNumber struct {
	CardNumber string `json:"card_number" validate:"required,min=1"`
	Year int `json:"year" validate:"required"`
	Month int `json:"month" validate:"required"`
}
```

### `MonthTopupStatus`

MonthTopupStatus represents parameters for retrieving top-up status by month.
Used to query top-up statistics for a specific month and year.

```go
type MonthTopupStatus struct {
	Year int `json:"year" validate:"required"`
	Month int `json:"month" validate:"required"`
}
```

### `MonthTopupStatusCardNumber`

MonthTopupStatusCardNumber represents parameters for retrieving card-specific top-up status by month.
Used to query monthly top-up statistics for a specific card.

```go
type MonthTopupStatusCardNumber struct {
	CardNumber string `json:"card_number" validate:"required,min=1"`
	Year int `json:"year" validate:"required"`
	Month int `json:"month" validate:"required"`
}
```

### `MonthTotalSaldoBalance`

MonthTotalSaldoBalance represents parameters for retrieving monthly balance summaries.
Used to fetch aggregated balance information for specific month/year periods.

```go
type MonthTotalSaldoBalance struct {
	Year int `json:"year" validate:"required"`
	Month int `json:"month" validate:"required"`
}
```

### `MonthYearAmountApiKey`

MonthYearAmountApiKey represents a request for transaction amount statistics by API key and year.
Used to retrieve financial summaries for a specific merchant.

```go
type MonthYearAmountApiKey struct {
	Apikey string `json:"api_key" validate:"required,min=1"`
	Year int `json:"year" validate:"required"`
}
```

### `MonthYearAmountMerchant`

MonthYearAmountMerchant represents a request for transaction amount statistics by merchant ID and year.
Used internally for financial reporting.

```go
type MonthYearAmountMerchant struct {
	MerchantID int `json:"merchant_id" validate:"required,min=1"`
	Year int `json:"year" validate:"required"`
}
```

### `MonthYearCardNumber`

MonthYearCardNumber represents parameters for operations requiring card number and year.
Used for card-specific annual data retrieval operations.

```go
type MonthYearCardNumber struct {
	CardNumber string `json:"card_number" validate:"required,min=1"`
	Year int `json:"year" validate:"required"`
}
```

### `MonthYearCardNumberCard`

MonthYearCardNumberCard represents a request structure for operations requiring
card number and year information, typically used for card lookups or expiration checks.

```go
type MonthYearCardNumberCard struct {
	CardNumber string `json:"card_number" validate:"required"`
	Year int `json:"year" validate:"required"`
}
```

### `MonthYearPaymentMethod`

MonthYearPaymentMethod represents parameters for retrieving payment method statistics by card and year.
Used to analyze payment method usage patterns for specific cards.

```go
type MonthYearPaymentMethod struct {
	CardNumber string `json:"card_number" validate:"required,min=1"`
	Year int `json:"year" validate:"required"`
}
```

### `MonthYearPaymentMethodApiKey`

MonthYearPaymentMethodApiKey represents a request for payment method statistics by API key and year.
Used to retrieve aggregated payment method data for a specific merchant.

```go
type MonthYearPaymentMethodApiKey struct {
	Apikey string `json:"api_key" validate:"required,min=1"`
	Year int `json:"year" validate:"required"`
}
```

### `MonthYearPaymentMethodMerchant`

MonthYearPaymentMethodMerchant represents a request for payment method statistics by merchant ID and year.
Used internally for merchant analytics.

```go
type MonthYearPaymentMethodMerchant struct {
	MerchantID int `json:"merchant_id" validate:"required,min=1"`
	Year int `json:"year" validate:"required"`
}
```

### `MonthYearTotalAmountApiKey`

MonthYearTotalAmountApiKey represents a request for total transaction amounts by API key and year.
Used for annual financial reporting per merchant.

```go
type MonthYearTotalAmountApiKey struct {
	Apikey string `json:"api_key" validate:"required,min=1"`
	Year int `json:"year" validate:"required"`
}
```

### `MonthYearTotalAmountMerchant`

MonthYearTotalAmountMerchant represents a request for total transaction amounts by merchant ID and year.
Used for comprehensive merchant financial analysis.

```go
type MonthYearTotalAmountMerchant struct {
	MerchantID int `json:"merchant_id" validate:"required,min=1"`
	Year int `json:"year" validate:"required"`
}
```

### `RefreshTokenRequest`

RefreshTokenRequest represents the payload for refresh token operations.
Used when exchanging a refresh token for new access credentials.

```go
type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required,min=1"`
}
```

#### Methods

##### `Validate`

Validate performs structural validation of RefreshTokenRequest fields.
Checks that the refresh token is present and meets minimum requirements.
Returns:
  - error: if validation fails according to struct tag rules

```go
func (r *RefreshTokenRequest) Validate() error
```

### `RegisterRequest`

RegisterRequest represents the payload for user registration.
Contains all necessary information to create a new user account.

```go
type RegisterRequest struct {
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	ConfirmPassword string `json:"confirm_password" validate:"required,min=6"`
	VerifiedCode string `json:"verified_code"`
	IsVerified bool `json:"is_verified"`
}
```

#### Methods

##### `Validate`

Validate validates the RegisterRequest struct using the validator.v10 package.

Returns an error if any of the validation rules fail.

```go
func (r *RegisterRequest) Validate() error
```

### `RemoveUserRoleRequest`

RemoveUserRoleRequest represents the payload for revoking a role from a user.
Used when removing specific permissions or access levels from users.

```go
type RemoveUserRoleRequest struct {
	UserId int `json:"user_id" validate:"required"`
	RoleId int `json:"role_id" validate:"required"`
}
```

### `RoleRequestPayload`

RoleRequestPayload represents the base payload structure for role management requests.
Contains common fields used across role-related operations.

```go
type RoleRequestPayload struct {
	UserID int `json:"user_id"`
	CorrelationID string `json:"correlation_id"`
	ReplyTopic string `json:"reply_topic"`
}
```

### `UpdateCardRequest`

UpdateCardRequest represents the payload for updating an existing payment card.
Contains all modifiable fields for a payment card record.

```go
type UpdateCardRequest struct {
	CardID int `json:"card_id" validate:"required,min=1"`
	UserID int `json:"user_id" validate:"required,min=1"`
	CardType string `json:"card_type" validate:"required"`
	ExpireDate time.Time `json:"expire_date" validate:"required"`
	CVV string `json:"cvv" validate:"required"`
	CardProvider string `json:"card_provider" validate:"required"`
}
```

#### Methods

##### `Validate`

Validate performs validation of UpdateCardRequest fields beyond basic struct validation.
Ensures card type and provider are valid values and IDs are positive.
Returns:
  - error: if validation fails, describing the specific validation error

```go
func (r *UpdateCardRequest) Validate() error
```

### `UpdateMerchantDocumentRequest`

UpdateMerchantDocumentRequest represents the payload for updating an existing merchant document.
Contains all modifiable fields for a merchant document record.

```go
type UpdateMerchantDocumentRequest struct {
	DocumentID *int `json:"document_id"`
	MerchantID int `json:"merchant_id" validate:"required,min=1"`
	DocumentType string `json:"document_type" validate:"required"`
	DocumentUrl string `json:"document_url" validate:"required"`
	Status string `json:"status" validate:"required"`
	Note string `json:"note" validate:"required"`
}
```

#### Methods

##### `Validate`

Validate performs basic validation of UpdateMerchantDocumentRequest fields using struct tags.
Returns:
  - error: if validation fails according to struct tag rules

```go
func (r *UpdateMerchantDocumentRequest) Validate() error
```

### `UpdateMerchantDocumentStatusRequest`

UpdateMerchantDocumentStatusRequest represents the payload for changing a document's review status.
Used specifically for status updates during document review processes.

```go
type UpdateMerchantDocumentStatusRequest struct {
	DocumentID *int `json:"document_id"`
	MerchantID int `json:"merchant_id" validate:"required,min=1"`
	Status string `json:"status" validate:"required"`
	Note string `json:"note" validate:"required"`
}
```

#### Methods

##### `Validate`

Validate performs basic validation of UpdateMerchantDocumentStatusRequest fields using struct tags.
Returns:
  - error: if validation fails according to struct tag rules

```go
func (r *UpdateMerchantDocumentStatusRequest) Validate() error
```

### `UpdateMerchantRequest`

UpdateMerchantRequest represents the payload for updating merchant details.
Used in merchant profile management.

```go
type UpdateMerchantRequest struct {
	MerchantID *int `json:"merchant_id"`
	Name string `json:"name" validate:"required"`
	UserID int `json:"user_id" validate:"required,min=1"`
	Status string `json:"status" validate:"required"`
}
```

#### Methods

##### `Validate`

Validate performs validation of UpdateMerchantRequest fields.
Ensures all required fields are present and valid.
Returns:
  - error: if validation fails

```go
func (r UpdateMerchantRequest) Validate() error
```

### `UpdateMerchantStatusRequest`

UpdateMerchantStatusRequest represents the payload for changing merchant status.
Used specifically for merchant account status management.

```go
type UpdateMerchantStatusRequest struct {
	MerchantID *int `json:"merchant_id"`
	Status string `json:"status" validate:"required"`
}
```

#### Methods

##### `Validate`

Validate performs validation of UpdateMerchantStatusRequest fields.
Ensures all required fields are present and valid.
Returns:
  - error: if validation fails

```go
func (r UpdateMerchantStatusRequest) Validate() error
```

### `UpdateRefreshToken`

UpdateRefreshToken represents the payload for updating an existing refresh token.
Used when rotating or extending refresh token validity.

```go
type UpdateRefreshToken struct {
	UserId int `json:"user_id" validate:"required,min=1"`
	Token string `json:"token" validate:"required,min=1"`
	ExpiresAt string `json:"expires_at" validate:"required,min=1"`
}
```

#### Methods

##### `Validate`

Validate performs structural validation of UpdateRefreshToken fields.
Checks that all required fields meet minimum requirements.
Returns:
  - error: if validation fails according to struct tag rules

```go
func (r *UpdateRefreshToken) Validate() error
```

### `UpdateRoleRequest`

UpdateRoleRequest represents the payload for modifying an existing role.
Used when renaming or updating role definitions.

```go
type UpdateRoleRequest struct {
	ID *int `json:"id"`
	Name string `json:"name" validate:"required"`
}
```

#### Methods

##### `Validate`

Validate performs validation of UpdateRoleRequest fields.
Ensures the role ID (when provided) and name meet requirements.
Returns:
  - error: if validation fails, describing the validation failure

```go
func (r *UpdateRoleRequest) Validate() error
```

### `UpdateSaldoBalance`

UpdateSaldoBalance represents the payload for balance adjustment operations.
Used specifically for updating card balance amounts.

```go
type UpdateSaldoBalance struct {
	CardNumber string `json:"card_number" validate:"required,min=1"`
	TotalBalance int `json:"total_balance" validate:"required,min=50000"`
}
```

#### Methods

##### `Validate`

Validate performs validation of UpdateSaldoBalance fields.
Ensures balance meets minimum requirements and card number is valid.
Returns:
  - error: if validation fails

```go
func (r *UpdateSaldoBalance) Validate() error
```

### `UpdateSaldoRequest`

UpdateSaldoRequest represents the payload for modifying an existing saldo record.
Used for comprehensive updates to balance information.

```go
type UpdateSaldoRequest struct {
	SaldoID *int `json:"saldo_id"`
	CardNumber string `json:"card_number" validate:"required"`
	TotalBalance int `json:"total_balance" validate:"required"`
}
```

#### Methods

##### `Validate`

Validate performs basic validation of UpdateSaldoRequest fields.
Ensures required fields are present and properly formatted.
Returns:
  - error: if validation fails

```go
func (r *UpdateSaldoRequest) Validate() error
```

### `UpdateSaldoWithdraw`

UpdateSaldoWithdraw represents the payload for withdrawal operations.
Used when processing balance withdrawals from cards.

```go
type UpdateSaldoWithdraw struct {
	CardNumber string `json:"card_number" validate:"required,min=1"`
	TotalBalance int `json:"total_balance" validate:"required,min=50000"`
	WithdrawAmount *int `json:"withdraw_amount" validate:"omitempty,gte=0"`
	WithdrawTime *time.Time `json:"withdraw_time" validate:"omitempty"`
}
```

#### Methods

##### `Validate`

Validate performs comprehensive validation of UpdateSaldoWithdraw fields.
Ensures:
- Withdrawal amount and time are either both provided or both omitted
- Withdrawal amount doesn't exceed available balance
- All required fields meet minimum requirements
Returns:
  - error: if validation fails with specific validation messages

```go
func (r *UpdateSaldoWithdraw) Validate() error
```

### `UpdateTopupAmount`

UpdateTopupAmount represents the payload for adjusting a top-up amount.
Used specifically for amount corrections.

```go
type UpdateTopupAmount struct {
	TopupID int `json:"topup_id" validate:"required,min=1"`
	TopupAmount int `json:"topup_amount" validate:"required,min=50000"`
}
```

#### Methods

##### `Validate`

Validate performs validation of UpdateTopupAmount fields.
Ensures:
- Valid top-up ID
- Minimum top-up amount (50,000)
Returns:
  - error: if validation fails

```go
func (r *UpdateTopupAmount) Validate() error
```

### `UpdateTopupRequest`

UpdateTopupRequest represents the payload for updating a top-up transaction.
Used to modify existing top-up records.

```go
type UpdateTopupRequest struct {
	CardNumber string `json:"card_number" validate:"required,min=1"`
	TopupID *int `json:"topup_id"`
	TopupAmount int `json:"topup_amount" validate:"required,min=50000"`
	TopupMethod string `json:"topup_method" validate:"required"`
}
```

#### Methods

##### `Validate`

Validate performs comprehensive validation of UpdateTopupRequest fields.
Ensures:
- Valid top-up ID
- Minimum top-up amount (50,000)
- Valid payment method
Returns:
  - error: if validation fails with specific validation messages

```go
func (r *UpdateTopupRequest) Validate() error
```

### `UpdateTopupStatus`

UpdateTopupStatus represents the payload for changing a top-up status.
Used to update processing status of top-up transactions.

```go
type UpdateTopupStatus struct {
	TopupID int `json:"topup_id" validate:"required,min=1"`
	Status string `json:"status" validate:"required"`
}
```

#### Methods

##### `Validate`

Validate performs basic validation of UpdateTopupStatus fields.
Ensures required fields are present.
Returns:
  - error: if validation fails

```go
func (r *UpdateTopupStatus) Validate() error
```

### `UpdateTransactionRequest`

UpdateTransactionRequest represents the payload for updating a transaction record.
Used to modify existing transaction details.

```go
type UpdateTransactionRequest struct {
	TransactionID *int `json:"transaction_id"`
	CardNumber string `json:"card_number" validate:"required,min=1"`
	Amount int `json:"amount" validate:"required,min=50000"`
	PaymentMethod string `json:"payment_method" validate:"required"`
	MerchantID *int `json:"merchant_id" validate:"required,min=1"`
	TransactionTime time.Time `json:"transaction_time" validate:"required"`
}
```

#### Methods

##### `Validate`

Validate performs comprehensive validation of UpdateTransactionRequest fields.
Ensures:
- Valid payment method
- Minimum transaction amount (50,000)
- Required fields are present
Returns:
  - error: if validation fails with specific validation messages

```go
func (r *UpdateTransactionRequest) Validate() error
```

### `UpdateTransactionStatus`

UpdateTransactionStatus represents the payload for changing a transaction status.
Used to update processing status of transactions.

```go
type UpdateTransactionStatus struct {
	TransactionID int `json:"transaction_id" validate:"required,min=1"`
	Status string `json:"status" validate:"required"`
}
```

#### Methods

##### `Validate`

Validate performs basic validation of UpdateTransactionStatus fields.
Ensures required fields are present.
Returns:
  - error: if validation fails

```go
func (r *UpdateTransactionStatus) Validate() error
```

### `UpdateTransferAmountRequest`

UpdateTransferAmountRequest represents the payload for adjusting a transfer amount.
Used specifically for amount corrections.

```go
type UpdateTransferAmountRequest struct {
	TransferID int `json:"transfer_id" validate:"required,min=1"`
	TransferAmount int `json:"transfer_amount" validate:"required,gt=0"`
}
```

#### Methods

##### `Validate`

Validate performs validation of UpdateTransferAmountRequest fields.
Ensures:
- Valid transfer ID
- Positive transfer amount
Returns:
  - error: if validation fails

```go
func (r *UpdateTransferAmountRequest) Validate() error
```

### `UpdateTransferRequest`

UpdateTransferRequest represents the payload for modifying a transfer record.
Used to update existing transfer details.

```go
type UpdateTransferRequest struct {
	TransferID *int `json:"transfer_id"`
	TransferFrom string `json:"transfer_from" validate:"required"`
	TransferTo string `json:"transfer_to" validate:"required,min=1"`
	TransferAmount int `json:"transfer_amount" validate:"required,min=50000"`
}
```

#### Methods

##### `Validate`

Validate performs comprehensive validation of UpdateTransferRequest fields.
Ensures:
- Valid transfer ID
- Minimum transfer amount (50,000)
- Valid account/card numbers
Returns:
  - error: if validation fails with specific validation messages

```go
func (r *UpdateTransferRequest) Validate() error
```

### `UpdateTransferStatus`

UpdateTransferStatus represents the payload for changing a transfer status.
Used to update processing status of transfers.

```go
type UpdateTransferStatus struct {
	TransferID int `json:"transfer_id" validate:"required,min=1"`
	Status string `json:"status" validate:"required"`
}
```

#### Methods

##### `Validate`

Validate performs basic validation of UpdateTransferStatus fields.
Ensures required fields are present.
Returns:
  - error: if validation fails

```go
func (r *UpdateTransferStatus) Validate() error
```

### `UpdateUserRequest`

UpdateUserRequest represents the payload for modifying an existing user's information.
Used when updating user profile details or credentials.

```go
type UpdateUserRequest struct {
	UserID *int `json:"user_id"`
	FirstName string `json:"firstname" validate:"required,alpha"`
	LastName string `json:"lastname" validate:"required,alpha"`
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=Password"`
}
```

#### Methods

##### `Validate`

Validate performs structural validation of UpdateUserRequest fields.
Ensures all required fields are present and meet format requirements.
Returns:
  - error: if validation fails according to struct tag rules

```go
func (r *UpdateUserRequest) Validate() error
```

### `UpdateWithdrawRequest`

UpdateWithdrawRequest represents the payload for modifying a withdrawal record.
Used to update existing withdrawal details.

```go
type UpdateWithdrawRequest struct {
	CardNumber string `json:"card_number" validate:"required,min=1"`
	WithdrawID *int `json:"withdraw_id"`
	WithdrawAmount int `json:"withdraw_amount" validate:"required,min=50000"`
	WithdrawTime time.Time `json:"withdraw_time" validate:"required"`
}
```

#### Methods

##### `Validate`

Validate performs comprehensive validation of UpdateWithdrawRequest fields.
Ensures:
- Minimum withdrawal amount (50,000)
- Withdrawal time is not in the future
- Required fields are present
Returns:
  - error: if validation fails with specific validation messages

```go
func (r *UpdateWithdrawRequest) Validate() error
```

### `UpdateWithdrawStatus`

UpdateWithdrawStatus represents the payload for changing a withdrawal status.
Used to update processing status of withdrawals.

```go
type UpdateWithdrawStatus struct {
	WithdrawID int `json:"withdraw_id" validate:"required,min=1"`
	Status string `json:"status" validate:"required"`
}
```

#### Methods

##### `Validate`

Validate performs basic validation of UpdateWithdrawStatus fields.
Ensures required fields are present.
Returns:
  - error: if validation fails

```go
func (r *UpdateWithdrawStatus) Validate() error
```

### `YearMonthCardNumber`

YearMonthCardNumber represents parameters for operations requiring card number and year.
Used for card-specific annual data retrieval operations.

```go
type YearMonthCardNumber struct {
	CardNumber string `json:"card_number" validate:"required,min=1"`
	Year int `json:"year" validate:"required"`
}
```

### `YearMonthMethod`

YearMonthMethod represents parameters for retrieving top-up methods by year and card.
Used to analyze payment method trends for specific cards.

```go
type YearMonthMethod struct {
	CardNumber string `json:"card_number" validate:"required,min=1"`
	Year int `json:"year" validate:"required"`
}
```

### `YearStatusTransactionCardNumber`

YearStatusTransactionCardNumber represents parameters for retrieving card-specific annual transaction status.
Used to get yearly transaction statistics for specific cards.

```go
type YearStatusTransactionCardNumber struct {
	CardNumber string `json:"card_number" validate:"required,min=1"`
	Year int `json:"year" validate:"required"`
}
```

### `YearStatusTransferCardNumber`

YearStatusTransferCardNumber represents parameters for retrieving annual transfer status by card.
Used to get yearly transfer statistics for specific cards.

```go
type YearStatusTransferCardNumber struct {
	CardNumber string `json:"card_number" validate:"required,min=1"`
	Year int `json:"year" validate:"required"`
}
```

### `YearStatusWithdrawCardNumber`

YearStatusWithdrawCardNumber represents parameters for retrieving card-specific annual withdrawal status.
Used to get yearly withdrawal statistics for specific cards.

```go
type YearStatusWithdrawCardNumber struct {
	CardNumber string `json:"card_number" validate:"required,min=1"`
	Year int `json:"year" validate:"required"`
}
```

### `YearTopupStatusCardNumber`

YearTopupStatusCardNumber represents parameters for retrieving card-specific top-up status by year.
Used to query annual top-up statistics for a specific card.

```go
type YearTopupStatusCardNumber struct {
	CardNumber string `json:"card_number" validate:"required,min=1"`
	Year int `json:"year" validate:"required"`
}
```