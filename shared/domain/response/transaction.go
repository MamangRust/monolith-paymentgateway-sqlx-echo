package response

// TransactionResponse represents a transaction record
type TransactionResponse struct {
	ID              int    `json:"id"`               // Unique transaction identifier
	TransactionNo   string `json:"transaction_no"`   // Unique transaction reference number
	CardNumber      string `json:"card_number"`      // Card number used for the transaction
	Amount          int    `json:"amount"`           // Transaction amount
	PaymentMethod   string `json:"payment_method"`   // Payment method used (e.g., "credit", "debit")
	MerchantID      int    `json:"merchant_id"`      // ID of the merchant where transaction occurred
	TransactionTime string `json:"transaction_time"` // Timestamp of the transaction
	CreatedAt       string `json:"created_at"`       // When the record was created
	UpdatedAt       string `json:"updated_at"`       // When the record was last updated
}

// TransactionResponseDeleteAt extends TransactionResponse with soft delete capability
type TransactionResponseDeleteAt struct {
	DeletedAt *string `json:"deleted_at"` // Timestamp when record was soft deleted (nil if active)
	// Embedded TransactionResponse fields
	ID              int    `json:"id"`
	TransactionNo   string `json:"transaction_no"`
	CardNumber      string `json:"card_number"`
	Amount          int    `json:"amount"`
	PaymentMethod   string `json:"payment_method"`
	MerchantID      int    `json:"merchant_id"`
	TransactionTime string `json:"transaction_time"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
}

// TransactionResponseMonthStatusSuccess represents monthly successful transaction metrics
type TransactionResponseMonthStatusSuccess struct {
	Year         string `json:"year"`          // Year of the metrics
	Month        string `json:"month"`         // Month of the metrics
	TotalAmount  int    `json:"total_amount"`  // Total amount of successful transactions
	TotalSuccess int    `json:"total_success"` // Count of successful transactions
}

// TransactionResponseYearStatusSuccess represents yearly successful transaction metrics
type TransactionResponseYearStatusSuccess struct {
	Year         string `json:"year"`          // Year of the metrics
	TotalAmount  int    `json:"total_amount"`  // Total amount of successful transactions
	TotalSuccess int    `json:"total_success"` // Count of successful transactions
}

// TransactionResponseMonthStatusFailed represents monthly failed transaction metrics
type TransactionResponseMonthStatusFailed struct {
	Year        string `json:"year"`         // Year of the metrics
	Month       string `json:"month"`        // Month of the metrics
	TotalAmount int    `json:"total_amount"` // Total amount of failed transactions
	TotalFailed int    `json:"total_failed"` // Count of failed transactions
}

// TransactionResponseYearStatusFailed represents yearly failed transaction metrics
type TransactionResponseYearStatusFailed struct {
	Year        string `json:"year"`         // Year of the metrics
	TotalAmount int    `json:"total_amount"` // Total amount of failed transactions
	TotalFailed int    `json:"total_failed"` // Count of failed transactions
}

// TransactionMonthMethodResponse represents monthly transaction metrics by payment method
type TransactionMonthMethodResponse struct {
	Month             string `json:"month"`              // Month of the metrics
	PaymentMethod     string `json:"payment_method"`     // Payment method used
	TotalTransactions int    `json:"total_transactions"` // Number of transactions
	TotalAmount       int    `json:"total_amount"`       // Total amount processed
}

// TransactionYearMethodResponse represents yearly transaction metrics by payment method
type TransactionYearMethodResponse struct {
	Year              string `json:"year"`               // Year of the metrics
	PaymentMethod     string `json:"payment_method"`     // Payment method used
	TotalTransactions int    `json:"total_transactions"` // Number of transactions
	TotalAmount       int    `json:"total_amount"`       // Total amount processed
}

// TransactionMonthAmountResponse represents monthly transaction amount metrics
type TransactionMonthAmountResponse struct {
	Month       string `json:"month"`        // Month of the metrics
	TotalAmount int    `json:"total_amount"` // Total amount processed
}

// TransactionYearlyAmountResponse represents yearly transaction amount metrics
type TransactionYearlyAmountResponse struct {
	Year        string `json:"year"`         // Year of the metrics
	TotalAmount int    `json:"total_amount"` // Total amount processed
}

// ApiResponseTransactionMonthStatusSuccess is API response for monthly success metrics
type ApiResponseTransactionMonthStatusSuccess struct {
	Status  string                                   `json:"status"`  // Response status
	Message string                                   `json:"message"` // Response message
	Data    []*TransactionResponseMonthStatusSuccess `json:"data"`    // Array of monthly success data
}

// ApiResponseTransactionYearStatusSuccess is API response for yearly success metrics
type ApiResponseTransactionYearStatusSuccess struct {
	Status  string                                  `json:"status"`  // Response status
	Message string                                  `json:"message"` // Response message
	Data    []*TransactionResponseYearStatusSuccess `json:"data"`    // Array of yearly success data
}

// ApiResponseTransactionMonthStatusFailed is API response for monthly failed metrics
type ApiResponseTransactionMonthStatusFailed struct {
	Status  string                                  `json:"status"`  // Response status
	Message string                                  `json:"message"` // Response message
	Data    []*TransactionResponseMonthStatusFailed `json:"data"`    // Array of monthly failed data
}

// ApiResponseTransactionYearStatusFailed is API response for yearly failed metrics
type ApiResponseTransactionYearStatusFailed struct {
	Status  string                                 `json:"status"`  // Response status
	Message string                                 `json:"message"` // Response message
	Data    []*TransactionResponseYearStatusFailed `json:"data"`    // Array of yearly failed data
}

// ApiResponseTransactionMonthMethod is API response for monthly method metrics
type ApiResponseTransactionMonthMethod struct {
	Status  string                            `json:"status"`  // Response status
	Message string                            `json:"message"` // Response message
	Data    []*TransactionMonthMethodResponse `json:"data"`    // Array of monthly method data
}

// ApiResponseTransactionYearMethod is API response for yearly method metrics
type ApiResponseTransactionYearMethod struct {
	Status  string                           `json:"status"`  // Response status
	Message string                           `json:"message"` // Response message
	Data    []*TransactionYearMethodResponse `json:"data"`    // Array of yearly method data
}

// ApiResponseTransactionMonthAmount is API response for monthly amount metrics
type ApiResponseTransactionMonthAmount struct {
	Status  string                            `json:"status"`  // Response status
	Message string                            `json:"message"` // Response message
	Data    []*TransactionMonthAmountResponse `json:"data"`    // Array of monthly amount data
}

// ApiResponseTransactionYearAmount is API response for yearly amount metrics
type ApiResponseTransactionYearAmount struct {
	Status  string                             `json:"status"`  // Response status
	Message string                             `json:"message"` // Response message
	Data    []*TransactionYearlyAmountResponse `json:"data"`    // Array of yearly amount data
}

// ApiResponseTransaction is API response for a single transaction
type ApiResponseTransaction struct {
	Status  string               `json:"status"`  // Response status
	Message string               `json:"message"` // Response message
	Data    *TransactionResponse `json:"data"`    // Single transaction data
}

type ApiResponseTransactionDeleteAt struct {
	Status  string                       `json:"status"`
	Message string                       `json:"message"`
	Data    *TransactionResponseDeleteAt `json:"data"`
}

// ApiResponseTransactions is API response for multiple transactions
type ApiResponseTransactions struct {
	Status  string                 `json:"status"`  // Response status
	Message string                 `json:"message"` // Response message
	Data    []*TransactionResponse `json:"data"`    // Array of transaction data
}

// ApiResponseTransactionDelete is API response for transaction deletion
type ApiResponseTransactionDelete struct {
	Status  string `json:"status"`  // Response status
	Message string `json:"message"` // Response message
}

// ApiResponseTransactionAll is generic API response for transaction operations
type ApiResponseTransactionAll struct {
	Status  string `json:"status"`  // Response status
	Message string `json:"message"` // Response message
}

// ApiResponsePaginationTransaction is paginated API response for transactions
type ApiResponsePaginationTransaction struct {
	Status     string                 `json:"status"`     // Response status
	Message    string                 `json:"message"`    // Response message
	Data       []*TransactionResponse `json:"data"`       // Array of transaction data
	Pagination *PaginationMeta        `json:"pagination"` // Pagination metadata
}

// ApiResponsePaginationTransactionDeleteAt is paginated API response for soft-deleted transactions
type ApiResponsePaginationTransactionDeleteAt struct {
	Status     string                         `json:"status"`     // Response status
	Message    string                         `json:"message"`    // Response message
	Data       []*TransactionResponseDeleteAt `json:"data"`       // Array of transaction data with delete info
	Pagination *PaginationMeta                `json:"pagination"` // Pagination metadata
}
