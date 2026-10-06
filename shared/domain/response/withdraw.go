package response

// WithdrawResponse represents a withdrawal transaction record
type WithdrawResponse struct {
	ID             int    `json:"id"`              // Unique withdrawal transaction identifier
	WithdrawNo     string `json:"withdraw_no"`     // Unique withdrawal reference number
	CardNumber     string `json:"card_number"`     // Card number used for the withdrawal
	WithdrawAmount int    `json:"withdraw_amount"` // Amount withdrawn (in smallest currency unit)
	WithdrawTime   string `json:"withdraw_time"`   // Timestamp when withdrawal was processed
	CreatedAt      string `json:"created_at"`      // When the withdrawal record was created
	UpdatedAt      string `json:"updated_at"`      // When the withdrawal record was last updated
}

// WithdrawResponseDeleteAt extends WithdrawResponse with soft-delete capability
type WithdrawResponseDeleteAt struct {
	DeletedAt *string `json:"deleted_at"` // Timestamp when record was soft deleted (nil if active)
	// Embedded WithdrawResponse fields
	ID             int    `json:"id"`
	WithdrawNo     string `json:"withdraw_no"`
	CardNumber     string `json:"card_number"`
	WithdrawAmount int    `json:"withdraw_amount"`
	WithdrawTime   string `json:"withdraw_time"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}

// WithdrawResponseMonthStatusSuccess represents monthly successful withdrawal metrics
type WithdrawResponseMonthStatusSuccess struct {
	Year         string `json:"year"`          // Year of the metrics (YYYY format)
	Month        string `json:"month"`         // Month of the metrics (MM format)
	TotalAmount  int    `json:"total_amount"`  // Total amount successfully withdrawn
	TotalSuccess int    `json:"total_success"` // Count of successful withdrawals
}

// WithdrawResponseYearStatusSuccess represents yearly successful withdrawal metrics
type WithdrawResponseYearStatusSuccess struct {
	Year         string `json:"year"`          // Year of the metrics (YYYY format)
	TotalAmount  int    `json:"total_amount"`  // Total amount successfully withdrawn
	TotalSuccess int    `json:"total_success"` // Count of successful withdrawals
}

// WithdrawResponseMonthStatusFailed represents monthly failed withdrawal metrics
type WithdrawResponseMonthStatusFailed struct {
	Year        string `json:"year"`         // Year of the metrics (YYYY format)
	Month       string `json:"month"`        // Month of the metrics (MM format)
	TotalAmount int    `json:"total_amount"` // Total amount of failed withdrawals
	TotalFailed int    `json:"total_failed"` // Count of failed withdrawals
}

// WithdrawResponseYearStatusFailed represents yearly failed withdrawal metrics
type WithdrawResponseYearStatusFailed struct {
	Year        string `json:"year"`         // Year of the metrics (YYYY format)
	TotalAmount int    `json:"total_amount"` // Total amount of failed withdrawals
	TotalFailed int    `json:"total_failed"` // Count of failed withdrawals
}

// WithdrawMonthlyAmountResponse represents monthly withdrawal amount totals
type WithdrawMonthlyAmountResponse struct {
	Month       string `json:"month"`        // Month of the metrics (MM format)
	TotalAmount int    `json:"total_amount"` // Total amount withdrawn
}

// WithdrawYearlyAmountResponse represents yearly withdrawal amount totals
type WithdrawYearlyAmountResponse struct {
	Year        string `json:"year"`         // Year of the metrics (YYYY format)
	TotalAmount int    `json:"total_amount"` // Total amount withdrawn
}

// ApiResponseWithdrawMonthStatusSuccess is API response for monthly success metrics
type ApiResponseWithdrawMonthStatusSuccess struct {
	Status  string                                `json:"status"`  // Response status ("success" or "error")
	Message string                                `json:"message"` // Descriptive response message
	Data    []*WithdrawResponseMonthStatusSuccess `json:"data"`    // Array of monthly success metrics
}

// ApiResponseWithdrawYearStatusSuccess is API response for yearly success metrics
type ApiResponseWithdrawYearStatusSuccess struct {
	Status  string                               `json:"status"`  // Response status
	Message string                               `json:"message"` // Response message
	Data    []*WithdrawResponseYearStatusSuccess `json:"data"`    // Array of yearly success metrics
}

// ApiResponseWithdrawMonthStatusFailed is API response for monthly failed metrics
type ApiResponseWithdrawMonthStatusFailed struct {
	Status  string                               `json:"status"`  // Response status
	Message string                               `json:"message"` // Response message
	Data    []*WithdrawResponseMonthStatusFailed `json:"data"`    // Array of monthly failed metrics
}

// ApiResponseWithdrawYearStatusFailed is API response for yearly failed metrics
type ApiResponseWithdrawYearStatusFailed struct {
	Status  string                              `json:"status"`  // Response status
	Message string                              `json:"message"` // Response message
	Data    []*WithdrawResponseYearStatusFailed `json:"data"`    // Array of yearly failed metrics
}

// ApiResponseWithdrawMonthAmount is API response for monthly withdrawal amounts
type ApiResponseWithdrawMonthAmount struct {
	Status  string                           `json:"status"`  // Response status
	Message string                           `json:"message"` // Response message
	Data    []*WithdrawMonthlyAmountResponse `json:"data"`    // Array of monthly amount totals
}

// ApiResponseWithdrawYearAmount is API response for yearly withdrawal amounts
type ApiResponseWithdrawYearAmount struct {
	Status  string                          `json:"status"`  // Response status
	Message string                          `json:"message"` // Response message
	Data    []*WithdrawYearlyAmountResponse `json:"data"`    // Array of yearly amount totals
}

// ApiResponsesWithdraw is API response for multiple withdrawal records
type ApiResponsesWithdraw struct {
	Status  string              `json:"status"`  // Response status
	Message string              `json:"message"` // Response message
	Data    []*WithdrawResponse `json:"data"`    // Array of withdrawal records
}

// ApiResponseWithdraw is API response for a single withdrawal record
type ApiResponseWithdraw struct {
	Status  string            `json:"status"`  // Response status
	Message string            `json:"message"` // Response message
	Data    *WithdrawResponse `json:"data"`    // Single withdrawal record
}

type ApiResponseWithdrawDeleteAt struct {
	Status  string                    `json:"status"`  // Response status
	Message string                    `json:"message"` // Response message
	Data    *WithdrawResponseDeleteAt `json:"data"`
}

// ApiResponseWithdrawDelete is API response for withdrawal deletion
type ApiResponseWithdrawDelete struct {
	Status  string `json:"status"`  // Response status
	Message string `json:"message"` // Operation result message
}

// ApiResponseWithdrawAll is generic API response for withdrawal operations
type ApiResponseWithdrawAll struct {
	Status  string `json:"status"`  // Response status
	Message string `json:"message"` // Operation message
}

// ApiResponsePaginationWithdraw is paginated API response for withdrawals
type ApiResponsePaginationWithdraw struct {
	Status     string              `json:"status"`     // Response status
	Message    string              `json:"message"`    // Response message
	Data       []*WithdrawResponse `json:"data"`       // Array of withdrawal records
	Pagination *PaginationMeta     `json:"pagination"` // Pagination metadata
}

// ApiResponsePaginationWithdrawDeleteAt is paginated API response for soft-deleted withdrawals
type ApiResponsePaginationWithdrawDeleteAt struct {
	Status     string                      `json:"status"`     // Response status
	Message    string                      `json:"message"`    // Response message
	Data       []*WithdrawResponseDeleteAt `json:"data"`       // Array of soft-deleted withdrawals
	Pagination *PaginationMeta             `json:"pagination"` // Pagination metadata
}
