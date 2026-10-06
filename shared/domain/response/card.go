package response

// CardResponse represents the basic card information returned in API responses.
// Used for non-sensitive card data display in most endpoints.
type CardResponse struct {
	ID           int    `json:"id"`            // Unique card identifier
	UserID       int    `json:"user_id"`       // ID of the card owner
	CardNumber   string `json:"card_number"`   // Masked card number (e.g., "4242********4242")
	CardType     string `json:"card_type"`     // Type of card ("credit" or "debit")
	ExpireDate   string `json:"expire_date"`   // Card expiration date (MM/YY format)
	CVV          string `json:"cvv"`           // Masked CVV (typically "***")
	CardProvider string `json:"card_provider"` // Card issuer (e.g., "Visa", "MasterCard")
	CreatedAt    string `json:"created_at"`    // Timestamp when card was added (RFC3339 format)
	UpdatedAt    string `json:"updated_at"`    // Timestamp when card was last updated
}

// CardResponseDeleteAt extends CardResponse with deletion information.
// Used in administrative interfaces showing soft-deleted cards.
type CardResponseDeleteAt struct {
	ID           int     `json:"id"`            // Unique card identifier
	UserID       int     `json:"user_id"`       // ID of the card owner
	CardNumber   string  `json:"card_number"`   // Masked card number
	CardType     string  `json:"card_type"`     // Type of card
	ExpireDate   string  `json:"expire_date"`   // Card expiration date
	CVV          string  `json:"cvv"`           // Masked CVV
	CardProvider string  `json:"card_provider"` // Card issuer
	CreatedAt    string  `json:"created_at"`    // Creation timestamp
	UpdatedAt    string  `json:"updated_at"`    // Last update timestamp
	DeletedAt    *string `json:"deleted_at"`    // Deletion timestamp (nil if not deleted)
}

// DashboardCard represents aggregated card statistics for dashboard display.
// Provides overview of financial activity across all cards.
type DashboardCard struct {
	TotalBalance     *int64 `json:"total_balance"`     // Sum of all card balances (in smallest currency unit)
	TotalTopup       *int64 `json:"total_topup"`       // Total top-up amount this period
	TotalWithdraw    *int64 `json:"total_withdraw"`    // Total withdrawals this period
	TotalTransaction *int64 `json:"total_transaction"` // Total transaction volume
	TotalTransfer    *int64 `json:"total_transfer"`    // Total transfer amount
}

// DashboardCardCardNumber provides detailed stats for a specific card.
// Shows both sending and receiving transfer amounts separately.
type DashboardCardCardNumber struct {
	TotalBalance          *int64 `json:"total_balance"`           // Current card balance
	TotalTopup            *int64 `json:"total_topup"`             // Total top-ups to this card
	TotalWithdraw         *int64 `json:"total_withdraw"`          // Total withdrawals from this card
	TotalTransaction      *int64 `json:"total_transaction"`       // Total transaction amount
	TotalTransferSend     *int64 `json:"total_transfer_send"`     // Total sent via transfers
	TotalTransferReceiver *int64 `json:"total_transfer_receiver"` // Total received via transfers
}

// CardResponseMonthBalance represents monthly balance information.
// Used for balance trend analysis and reporting.
type CardResponseMonthBalance struct {
	Month        string `json:"month"`         // Month in "YYYY-MM" format
	TotalBalance int64  `json:"total_balance"` // Balance at month end (in smallest currency unit)
}

// CardResponseYearlyBalance represents annual balance information.
// Used for yearly financial summaries.
type CardResponseYearlyBalance struct {
	Year         string `json:"year"`          // Year in "YYYY" format
	TotalBalance int64  `json:"total_balance"` // Balance at year end
}

// CardResponseMonthAmount represents monthly transaction amounts.
// Shows activity volume rather than balance.
type CardResponseMonthAmount struct {
	Month       string `json:"month"`        // Month in "YYYY-MM" format
	TotalAmount int64  `json:"total_amount"` // Total transaction amount for the month
}

// CardResponseYearAmount represents annual transaction amounts.
// Shows yearly activity volume.
type CardResponseYearAmount struct {
	Year        string `json:"year"`         // Year in "YYYY" format
	TotalAmount int64  `json:"total_amount"` // Total transaction amount for the year
}

// ApiResponseCard is the standard response format for single card requests.
type ApiResponseCard struct {
	Status  string        `json:"status"`  // Response status ("success" or "error")
	Message string        `json:"message"` // Descriptive message
	Data    *CardResponse `json:"data"`    // Card data payload
}

type ApiResponseCardDeleteAt struct {
	Status  string                `json:"status"`  // Response status ("success" or "error")
	Message string                `json:"message"` // Descriptive message
	Data    *CardResponseDeleteAt `json:"data"`    // Card data payload
}

// ApiResponseCardDelete is the response format for card deletion operations.
type ApiResponseCardDelete struct {
	Status  string `json:"status"`  // Operation status
	Message string `json:"message"` // Result message
}

// ApiResponseCardAll is the response format for bulk card operations.
type ApiResponseCardAll struct {
	Status  string `json:"status"`  // Operation status
	Message string `json:"message"` // Result message
}

// ApiResponsePaginationCard is the paginated response for card lists.
type ApiResponsePaginationCard struct {
	Status     string          `json:"status"`     // Response status
	Message    string          `json:"message"`    // Descriptive message
	Data       []*CardResponse `json:"data"`       // Array of card data
	Pagination *PaginationMeta `json:"pagination"` // Pagination metadata
}

// ApiResponsePaginationCardDeleteAt is the paginated response including deleted cards.
type ApiResponsePaginationCardDeleteAt struct {
	Status     string                  `json:"status"`     // Response status
	Message    string                  `json:"message"`    // Descriptive message
	Data       []*CardResponseDeleteAt `json:"data"`       // Array of card data (with deletion info)
	Pagination *PaginationMeta         `json:"pagination"` // Pagination metadata
}

// ApiResponseMonthlyBalance is the response format for monthly balance data.
type ApiResponseMonthlyBalance struct {
	Status  string                      `json:"status"`  // Response status
	Message string                      `json:"message"` // Descriptive message
	Data    []*CardResponseMonthBalance `json:"data"`    // Array of monthly balances
}

// ApiResponseYearlyBalance is the response format for yearly balance data.
type ApiResponseYearlyBalance struct {
	Status  string                       `json:"status"`  // Response status
	Message string                       `json:"message"` // Descriptive message
	Data    []*CardResponseYearlyBalance `json:"data"`    // Array of yearly balances
}

// ApiResponseMonthlyAmount is the response format for monthly transaction amounts.
type ApiResponseMonthlyAmount struct {
	Status  string                     `json:"status"`  // Response status
	Message string                     `json:"message"` // Descriptive message
	Data    []*CardResponseMonthAmount `json:"data"`    // Array of monthly amounts
}

// ApiResponseYearlyAmount is the response format for yearly transaction amounts.
type ApiResponseYearlyAmount struct {
	Status  string                    `json:"status"`  // Response status
	Message string                    `json:"message"` // Descriptive message
	Data    []*CardResponseYearAmount `json:"data"`    // Array of yearly amounts
}

// ApiResponseDashboardCard is the response format for dashboard card statistics.
type ApiResponseDashboardCard struct {
	Status  string         `json:"status"`  // Response status
	Message string         `json:"message"` // Descriptive message
	Data    *DashboardCard `json:"data"`    // Dashboard statistics
}

// ApiResponseDashboardCardNumber is the response format for card-specific dashboard stats.
type ApiResponseDashboardCardNumber struct {
	Status  string                   `json:"status"`  // Response status
	Message string                   `json:"message"` // Descriptive message
	Data    *DashboardCardCardNumber `json:"data"`    // Card-specific statistics
}

// CardAuthTransactionResponse represents an authorization transaction.
type CardAuthTransactionResponse struct {
	AuthID         int    `json:"auth_id"`
	TxnID          string `json:"txn_id"`
	CardNumber     string `json:"card_number"`
	MerchantID     int    `json:"merchant_id"`
	Amount         int64  `json:"amount"`
	Currency       string `json:"currency"`
	Mcc            string `json:"mcc"`
	PosEntryMode   string `json:"pos_entry_mode"`
	Status         string `json:"status"`
	IdempotencyKey string `json:"idempotency_key"`
	RiskScore      int    `json:"risk_score"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}

// ApiResponseCardAuthTransaction is the response for a single auth transaction.
type ApiResponseCardAuthTransaction struct {
	Status  string                       `json:"status"`
	Message string                       `json:"message"`
	Data    *CardAuthTransactionResponse `json:"data"`
}

// BillingCycleResponse represents a billing cycle record.
type BillingCycleResponse struct {
	BillingID  int    `json:"billing_id"`
	CardNumber string `json:"card_number"`
	CycleStart string `json:"cycle_start"`
	CycleEnd   string `json:"cycle_end"`
	AmountDue  int32  `json:"amount_due"`
	DueDate    string `json:"due_date"`
	Status     string `json:"status"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

// ApiResponseBillingCycle is the response for a single billing cycle.
type ApiResponseBillingCycle struct {
	Status  string                `json:"status"`
	Message string                `json:"message"`
	Data    *BillingCycleResponse `json:"data"`
}

// ApiResponseBillingCycles is the response for a list of billing cycles.
type ApiResponseBillingCycles struct {
	Status  string                  `json:"status"`
	Message string                  `json:"message"`
	Data    []*BillingCycleResponse `json:"data"`
}

// CardPaymentResponse represents a card payment record.
type CardPaymentResponse struct {
	PaymentID      int    `json:"payment_id"`
	PaymentUUID    string `json:"payment_uuid"`
	CardNumber     string `json:"card_number"`
	BillingID      *int32 `json:"billing_id"`
	Amount         int64  `json:"amount"`
	PaymentChannel string `json:"payment_channel"`
	ReferenceID    string `json:"reference_id"`
	Status         string `json:"status"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}

// ApiResponseCardPayment is the response for a single card payment.
type ApiResponseCardPayment struct {
	Status  string               `json:"status"`
	Message string               `json:"message"`
	Data    *CardPaymentResponse `json:"data"`
}

// ApiResponseCardPayments is the response for a list of card payments.
type ApiResponseCardPayments struct {
	Status  string                 `json:"status"`
	Message string                 `json:"message"`
	Data    []*CardPaymentResponse `json:"data"`
}

// CardRewardResponse represents a card reward record.
type CardRewardResponse struct {
	RewardID     int    `json:"reward_id"`
	CardNumber   string `json:"card_number"`
	TxnID        string `json:"txn_id"`
	Amount       int64  `json:"amount"`
	Mcc          string `json:"mcc"`
	PointsEarned int32  `json:"points_earned"`
	ExpiresAt    string `json:"expires_at"`
	Redeemed     bool   `json:"redeemed"`
	CreatedAt    string `json:"created_at"`
}

// ApiResponseCardReward is the response for a single reward operation.
type ApiResponseCardReward struct {
	Status  string              `json:"status"`
	Message string              `json:"message"`
	Data    *CardRewardResponse `json:"data"`
}

// ApiResponseCardRewards is the response for a list of rewards.
type ApiResponseCardRewards struct {
	Status  string                `json:"status"`
	Message string                `json:"message"`
	Data    []*CardRewardResponse `json:"data"`
}

// ApiResponseRewardBalance is the response for a reward balance query.
type ApiResponseRewardBalance struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Balance int64  `json:"balance"`
}
