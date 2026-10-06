package response

// SaldoResponse represents the basic saldo (balance) information in API responses.
// Contains core balance data without deletion information.
type SaldoResponse struct {
	ID             int    `json:"id"`              // Unique balance record identifier
	CardNumber     string `json:"card_number"`     // Masked card number (e.g., "4242******4242")
	TotalBalance   int    `json:"total_balance"`   // Current balance amount (in smallest currency unit)
	WithdrawAmount int    `json:"withdraw_amount"` // Last withdrawal amount (in smallest currency unit)
	WithdrawTime   string `json:"withdraw_time"`   // Timestamp of last withdrawal (RFC3339 format)
	CreatedAt      string `json:"created_at"`      // Timestamp when record was created
	UpdatedAt      string `json:"updated_at"`      // Timestamp when record was last updated
}

// SaldoResponseDeleteAt extends SaldoResponse with deletion information.
// Used in administrative interfaces showing soft-deleted balance records.
type SaldoResponseDeleteAt struct {
	ID             int     `json:"id"`              // Unique balance record identifier
	CardNumber     string  `json:"card_number"`     // Masked card number
	TotalBalance   int     `json:"total_balance"`   // Balance amount before deletion
	WithdrawAmount int     `json:"withdraw_amount"` // Last withdrawal amount
	WithdrawTime   string  `json:"withdraw_time"`   // Timestamp of last withdrawal
	CreatedAt      string  `json:"created_at"`      // Original creation timestamp
	UpdatedAt      string  `json:"updated_at"`      // Last update timestamp
	DeletedAt      *string `json:"deleted_at"`      // Deletion timestamp (nil if not deleted)
}

// SaldoMonthTotalBalanceResponse represents monthly balance totals including year.
// Used for detailed monthly balance reporting.
type SaldoMonthTotalBalanceResponse struct {
	Month        string `json:"month"`         // Month in "MM" format
	Year         string `json:"year"`          // Year in "YYYY" format
	TotalBalance int    `json:"total_balance"` // Aggregate balance for the month
}

// SaldoYearTotalBalanceResponse represents annual balance totals.
// Used for yearly financial summaries.
type SaldoYearTotalBalanceResponse struct {
	Year         string `json:"year"`          // Year in "YYYY" format
	TotalBalance int    `json:"total_balance"` // Aggregate balance for the year
}

// SaldoMonthBalanceResponse represents monthly balance information.
// Used for balance trend analysis by month.
type SaldoMonthBalanceResponse struct {
	Month        string `json:"month"`         // Month in "YYYY-MM" format
	TotalBalance int    `json:"total_balance"` // Balance amount for the month
}

// SaldoYearBalanceResponse represents annual balance information.
// Used for long-term balance trend analysis.
type SaldoYearBalanceResponse struct {
	Year         string `json:"year"`          // Year in "YYYY" format
	TotalBalance int    `json:"total_balance"` // Balance amount for the year
}

// API Response Wrappers:

// ApiResponseSaldo is the standard response format for single saldo requests.
type ApiResponseSaldo struct {
	Status  string         `json:"status"`  // Response status ("success" or "error")
	Message string         `json:"message"` // Descriptive message
	Data    *SaldoResponse `json:"data"`    // Saldo data payload
}

type ApiResponseSaldoDeleteAt struct {
	Status  string                 `json:"status"`  // Response status ("success" or "error")
	Message string                 `json:"message"` // Descriptive message
	Data    *SaldoResponseDeleteAt `json:"data"`    // Saldo data payload
}

// ApiResponsesSaldo is the response format for multiple saldo records (non-paginated).
type ApiResponsesSaldo struct {
	Status  string           `json:"status"`  // Response status
	Message string           `json:"message"` // Descriptive message
	Data    []*SaldoResponse `json:"data"`    // Array of saldo data
}

// ApiResponseSaldoDelete is the response format for saldo deletion operations.
type ApiResponseSaldoDelete struct {
	Status  string `json:"status"`  // Operation status
	Message string `json:"message"` // Result message
}

// ApiResponseSaldoAll is the response format for bulk saldo operations.
type ApiResponseSaldoAll struct {
	Status  string `json:"status"`  // Operation status
	Message string `json:"message"` // Result message
}

// ApiResponseMonthTotalSaldo wraps monthly total balance responses.
type ApiResponseMonthTotalSaldo struct {
	Status  string                            `json:"status"`  // Response status
	Message string                            `json:"message"` // Descriptive message
	Data    []*SaldoMonthTotalBalanceResponse `json:"data"`    // Array of monthly total balances
}

// ApiResponseYearTotalSaldo wraps yearly total balance responses.
type ApiResponseYearTotalSaldo struct {
	Status  string                           `json:"status"`  // Response status
	Message string                           `json:"message"` // Descriptive message
	Data    []*SaldoYearTotalBalanceResponse `json:"data"`    // Array of yearly total balances
}

// ApiResponseMonthSaldoBalances wraps monthly balance trend responses.
type ApiResponseMonthSaldoBalances struct {
	Status  string                       `json:"status"`  // Response status
	Message string                       `json:"message"` // Descriptive message
	Data    []*SaldoMonthBalanceResponse `json:"data"`    // Array of monthly balances
}

// ApiResponseYearSaldoBalances wraps yearly balance trend responses.
type ApiResponseYearSaldoBalances struct {
	Status  string                      `json:"status"`  // Response status
	Message string                      `json:"message"` // Descriptive message
	Data    []*SaldoYearBalanceResponse `json:"data"`    // Array of yearly balances
}

// ApiResponsePaginationSaldo is the paginated response for saldo lists.
type ApiResponsePaginationSaldo struct {
	Status     string           `json:"status"`     // Response status
	Message    string           `json:"message"`    // Descriptive message
	Data       []*SaldoResponse `json:"data"`       // Array of saldo data
	Pagination *PaginationMeta  `json:"pagination"` // Pagination metadata
}

// ApiResponsePaginationSaldoDeleteAt is the paginated response including deleted saldo records.
type ApiResponsePaginationSaldoDeleteAt struct {
	Status     string                   `json:"status"`     // Response status
	Message    string                   `json:"message"`    // Descriptive message
	Data       []*SaldoResponseDeleteAt `json:"data"`       // Array of saldo data (with deletion info)
	Pagination *PaginationMeta          `json:"pagination"` // Pagination metadata
}
