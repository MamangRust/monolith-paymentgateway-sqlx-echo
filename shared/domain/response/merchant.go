package response

// MerchantResponsePayload represents the base payload structure for merchant API responses.
// Used for validation and correlation purposes in merchant operations.
type MerchantResponsePayload struct {
	CorrelationID string `json:"correlation_id"`        // Unique ID for request tracing
	Valid         bool   `json:"valid"`                 // Validation status of the merchant
	MerchantID    int64  `json:"merchant_id,omitempty"` // ID of the merchant (omitted when empty)
}

// MerchantResponse represents the basic merchant information returned in API responses.
// Contains core merchant data without deletion information.
type MerchantResponse struct {
	ID        int    `json:"id"`         // Unique merchant identifier
	Name      string `json:"name"`       // Legal name of the merchant
	UserID    int    `json:"user_id"`    // ID of the user who owns/administers this merchant
	ApiKey    string `json:"api_key"`    // Merchant's API key (typically masked in responses)
	Status    string `json:"status"`     // Current status ("active", "suspended", etc.)
	CreatedAt string `json:"created_at"` // Timestamp when merchant was created (RFC3339 format)
	UpdatedAt string `json:"updated_at"` // Timestamp when merchant was last updated
}

// MerchantResponseDeleteAt extends MerchantResponse with deletion information.
// Used in administrative interfaces showing soft-deleted merchants.
type MerchantResponseDeleteAt struct {
	ID        int     `json:"id"`         // Unique merchant identifier
	Name      string  `json:"name"`       // Legal name of the merchant
	UserID    int     `json:"user_id"`    // ID of the administrator user
	ApiKey    string  `json:"api_key"`    // Merchant's API key
	Status    string  `json:"status"`     // Status before deletion
	CreatedAt string  `json:"created_at"` // Original creation timestamp
	UpdatedAt string  `json:"updated_at"` // Last update timestamp
	DeletedAt *string `json:"deleted_at"` // Deletion timestamp (nil if not deleted)
}

// MerchantTransactionResponse represents a merchant transaction record.
// Contains details of transactions processed by the merchant.
type MerchantTransactionResponse struct {
	ID              int    `json:"id"`               // Unique transaction identifier
	CardNumber      string `json:"card_number"`      // Masked card number used in transaction
	Amount          int32  `json:"amount"`           // Transaction amount (in smallest currency unit)
	PaymentMethod   string `json:"payment_method"`   // Method used ("credit", "debit", etc.)
	MerchantID      int32  `json:"merchant_id"`      // ID of the merchant
	MerchantName    string `json:"merchant_name"`    // Name of merchant at time of transaction
	TransactionTime string `json:"transaction_time"` // Timestamp of transaction (RFC3339 format)
	CreatedAt       string `json:"created_at"`       // Record creation timestamp
	UpdatedAt       string `json:"updated_at"`       // Last update timestamp
}

// MerchantResponseMonthlyPaymentMethod represents monthly transaction totals by payment method.
// Used for analyzing payment method trends by month.
type MerchantResponseMonthlyPaymentMethod struct {
	Month         string `json:"month"`          // Month in "YYYY-MM" format
	PaymentMethod string `json:"payment_method"` // Type of payment method
	TotalAmount   int    `json:"total_amount"`   // Total amount processed by this method
}

// MerchantResponseYearlyPaymentMethod represents annual transaction totals by payment method.
// Used for analyzing yearly payment method trends.
type MerchantResponseYearlyPaymentMethod struct {
	Year          string `json:"year"`           // Year in "YYYY" format
	PaymentMethod string `json:"payment_method"` // Type of payment method
	TotalAmount   int    `json:"total_amount"`   // Total amount processed by this method
}

// MerchantResponseMonthlyAmount represents monthly transaction totals.
// Shows aggregated transaction volume by month.
type MerchantResponseMonthlyAmount struct {
	Month       string `json:"month"`        // Month in "YYYY-MM" format
	TotalAmount int    `json:"total_amount"` // Total transaction amount for the month
}

// MerchantResponseYearlyAmount represents annual transaction totals.
// Shows aggregated transaction volume by year.
type MerchantResponseYearlyAmount struct {
	Year        string `json:"year"`         // Year in "YYYY" format
	TotalAmount int    `json:"total_amount"` // Total transaction amount for the year
}

// MerchantResponseMonthlyTotalAmount represents detailed monthly transaction totals.
// Includes both year and month for precise period identification.
type MerchantResponseMonthlyTotalAmount struct {
	Year        string `json:"year"`         // Year in "YYYY" format
	Month       string `json:"month"`        // Month in "MM" format
	TotalAmount int    `json:"total_amount"` // Total transaction amount
}

// MerchantResponseYearlyTotalAmount represents comprehensive annual transaction totals.
type MerchantResponseYearlyTotalAmount struct {
	Year        string `json:"year"`         // Year in "YYYY" format
	TotalAmount int    `json:"total_amount"` // Total transaction amount
}

// API Response Wrappers:

// ApiResponseMerchantMonthlyPaymentMethod wraps monthly payment method statistics.
type ApiResponseMerchantMonthlyPaymentMethod struct {
	Status  string                                  `json:"status"`  // Response status
	Message string                                  `json:"message"` // Descriptive message
	Data    []*MerchantResponseMonthlyPaymentMethod `json:"data"`    // Array of monthly payment method data
}

// ApiResponseMerchantYearlyPaymentMethod wraps yearly payment method statistics.
type ApiResponseMerchantYearlyPaymentMethod struct {
	Status  string                                 `json:"status"`  // Response status
	Message string                                 `json:"message"` // Descriptive message
	Data    []*MerchantResponseYearlyPaymentMethod `json:"data"`    // Array of yearly payment method data
}

// ApiResponseMerchantMonthlyAmount wraps monthly transaction amounts.
type ApiResponseMerchantMonthlyAmount struct {
	Status  string                           `json:"status"`  // Response status
	Message string                           `json:"message"` // Descriptive message
	Data    []*MerchantResponseMonthlyAmount `json:"data"`    // Array of monthly amounts
}

// ApiResponseMerchantYearlyAmount wraps yearly transaction amounts.
type ApiResponseMerchantYearlyAmount struct {
	Status  string                          `json:"status"`  // Response status
	Message string                          `json:"message"` // Descriptive message
	Data    []*MerchantResponseYearlyAmount `json:"data"`    // Array of yearly amounts
}

// ApiResponseMerchantMonthlyTotalAmount wraps detailed monthly transaction totals.
type ApiResponseMerchantMonthlyTotalAmount struct {
	Status  string                                `json:"status"`  // Response status
	Message string                                `json:"message"` // Descriptive message
	Data    []*MerchantResponseMonthlyTotalAmount `json:"data"`    // Array of detailed monthly totals
}

// ApiResponseMerchantYearlyTotalAmount wraps comprehensive yearly transaction totals.
type ApiResponseMerchantYearlyTotalAmount struct {
	Status  string                               `json:"status"`  // Response status
	Message string                               `json:"message"` // Descriptive message
	Data    []*MerchantResponseYearlyTotalAmount `json:"data"`    // Array of yearly totals
}

// ApiResponsesMerchant wraps multiple merchant responses.
type ApiResponsesMerchant struct {
	Status  string              `json:"status"`  // Response status
	Message string              `json:"message"` // Descriptive message
	Data    []*MerchantResponse `json:"data"`    // Array of merchant data
}

// ApiResponseMerchant wraps a single merchant response.
type ApiResponseMerchant struct {
	Status  string            `json:"status"`  // Response status
	Message string            `json:"message"` // Descriptive message
	Data    *MerchantResponse `json:"data"`    // Single merchant data
}

type ApiResponseMerchantDeleteAt struct {
	Status  string                    `json:"status"`  // Response status
	Message string                    `json:"message"` // Descriptive message
	Data    *MerchantResponseDeleteAt `json:"data"`
}

// ApiResponseMerchantDelete wraps merchant deletion responses.
type ApiResponseMerchantDelete struct {
	Status  string `json:"status"`  // Operation status
	Message string `json:"message"` // Result message
}

// ApiResponseMerchantAll wraps bulk merchant operation responses.
type ApiResponseMerchantAll struct {
	Status  string `json:"status"`  // Operation status
	Message string `json:"message"` // Result message
}

// ApiResponsePaginationMerchant wraps paginated merchant lists.
type ApiResponsePaginationMerchant struct {
	Status     string              `json:"status"`     // Response status
	Message    string              `json:"message"`    // Descriptive message
	Data       []*MerchantResponse `json:"data"`       // Array of merchant data
	Pagination *PaginationMeta     `json:"pagination"` // Pagination metadata
}

// ApiResponsePaginationMerchantDeleteAt wraps paginated lists including deleted merchants.
type ApiResponsePaginationMerchantDeleteAt struct {
	Status     string                      `json:"status"`     // Response status
	Message    string                      `json:"message"`    // Descriptive message
	Data       []*MerchantResponseDeleteAt `json:"data"`       // Array of merchant data (with deletion info)
	Pagination *PaginationMeta             `json:"pagination"` // Pagination metadata
}

// ApiResponsePaginationMerchantTransaction wraps paginated merchant transaction lists.
type ApiResponsePaginationMerchantTransaction struct {
	Status     string                         `json:"status"`     // Response status
	Message    string                         `json:"message"`    // Descriptive message
	Data       []*MerchantTransactionResponse `json:"data"`       // Array of transaction data
	Pagination *PaginationMeta                `json:"pagination"` // Pagination metadata
}
