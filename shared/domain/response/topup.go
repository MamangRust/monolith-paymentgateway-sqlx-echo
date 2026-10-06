package response

// TopupResponse represents a successful top-up transaction response
type TopupResponse struct {
	ID          int    `json:"id"`           // Unique identifier for the top-up
	CardNumber  string `json:"card_number"`  // Card number associated with the top-up
	TopupNo     string `json:"topup_no"`     // Unique top-up reference number
	TopupAmount int    `json:"topup_amount"` // Amount of the top-up
	TopupMethod string `json:"topup_method"` // Payment method used for the top-up
	TopupTime   string `json:"topup_time"`   // Time when the top-up was processed
	CreatedAt   string `json:"created_at"`   // Timestamp when record was created
	UpdatedAt   string `json:"updated_at"`   // Timestamp when record was last updated
}

// TopupResponseDeleteAt extends TopupResponse with soft delete information
type TopupResponseDeleteAt struct {
	DeletedAt *string `json:"deleted_At"` // Timestamp when record was soft deleted (nil if not deleted)
	// Embedded TopupResponse fields
	ID          int    `json:"id"`
	CardNumber  string `json:"card_number"`
	TopupNo     string `json:"topup_no"`
	TopupAmount int    `json:"topup_amount"`
	TopupMethod string `json:"topup_method"`
	TopupTime   string `json:"topup_time"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

// TopupResponseMonthStatusSuccess represents monthly successful top-up statistics
type TopupResponseMonthStatusSuccess struct {
	Year         string `json:"year"`          // Year of the statistics
	Month        string `json:"month"`         // Month of the statistics
	TotalAmount  int    `json:"total_amount"`  // Total amount of successful top-ups
	TotalSuccess int    `json:"total_success"` // Count of successful top-ups
}

// TopupResponseYearStatusSuccess represents yearly successful top-up statistics
type TopupResponseYearStatusSuccess struct {
	Year         string `json:"year"`          // Year of the statistics
	TotalAmount  int    `json:"total_amount"`  // Total amount of successful top-ups
	TotalSuccess int    `json:"total_success"` // Count of successful top-ups
}

// TopupResponseMonthStatusFailed represents monthly failed top-up statistics
type TopupResponseMonthStatusFailed struct {
	Year        string `json:"year"`         // Year of the statistics
	Month       string `json:"month"`        // Month of the statistics
	TotalAmount int    `json:"total_amount"` // Total amount of failed top-ups
	TotalFailed int    `json:"total_failed"` // Count of failed top-ups
}

// TopupResponseYearStatusFailed represents yearly failed top-up statistics
type TopupResponseYearStatusFailed struct {
	Year        string `json:"year"`         // Year of the statistics
	TotalAmount int    `json:"total_amount"` // Total amount of failed top-ups
	TotalFailed int    `json:"total_failed"` // Count of failed top-ups
}

// TopupMonthMethodResponse represents monthly top-up statistics by payment method
type TopupMonthMethodResponse struct {
	Month       string `json:"month"`        // Month of the statistics
	TopupMethod string `json:"topup_method"` // Payment method used
	TotalTopups int    `json:"total_topups"` // Number of top-ups using this method
	TotalAmount int    `json:"total_amount"` // Total amount processed with this method
}

// TopupYearlyMethodResponse represents yearly top-up statistics by payment method
type TopupYearlyMethodResponse struct {
	Year        string `json:"year"`         // Year of the statistics
	TopupMethod string `json:"topup_method"` // Payment method used
	TotalTopups int    `json:"total_topups"` // Number of top-ups using this method
	TotalAmount int    `json:"total_amount"` // Total amount processed with this method
}

// TopupMonthAmountResponse represents monthly top-up amount statistics
type TopupMonthAmountResponse struct {
	Month       string `json:"month"`        // Month of the statistics
	TotalAmount int    `json:"total_amount"` // Total amount processed
}

// TopupYearlyAmountResponse represents yearly top-up amount statistics
type TopupYearlyAmountResponse struct {
	Year        string `json:"year"`         // Year of the statistics
	TotalAmount int    `json:"total_amount"` // Total amount processed
}

// ApiResponseTopupMonthStatusSuccess is the API response for monthly successful top-up stats
type ApiResponseTopupMonthStatusSuccess struct {
	Status  string                             `json:"status"`  // Response status (e.g., "success")
	Message string                             `json:"message"` // Descriptive message
	Data    []*TopupResponseMonthStatusSuccess `json:"data"`    // Array of monthly success data
}

// ApiResponseTopupYearStatusSuccess is the API response for yearly successful top-up stats
type ApiResponseTopupYearStatusSuccess struct {
	Status  string                            `json:"status"`  // Response status
	Message string                            `json:"message"` // Descriptive message
	Data    []*TopupResponseYearStatusSuccess `json:"data"`    // Array of yearly success data
}

// ApiResponseTopupMonthStatusFailed is the API response for monthly failed top-up stats
type ApiResponseTopupMonthStatusFailed struct {
	Status  string                            `json:"status"`  // Response status
	Message string                            `json:"message"` // Descriptive message
	Data    []*TopupResponseMonthStatusFailed `json:"data"`    // Array of monthly failed data
}

// ApiResponseTopupYearStatusFailed is the API response for yearly failed top-up stats
type ApiResponseTopupYearStatusFailed struct {
	Status  string                           `json:"status"`  // Response status
	Message string                           `json:"message"` // Descriptive message
	Data    []*TopupResponseYearStatusFailed `json:"data"`    // Array of yearly failed data
}

// ApiResponseTopupMonthMethod is the API response for monthly top-up method stats
type ApiResponseTopupMonthMethod struct {
	Status  string                      `json:"status"`  // Response status
	Message string                      `json:"message"` // Descriptive message
	Data    []*TopupMonthMethodResponse `json:"data"`    // Array of monthly method data
}

// ApiResponseTopupYearMethod is the API response for yearly top-up method stats
type ApiResponseTopupYearMethod struct {
	Status  string                       `json:"status"`  // Response status
	Message string                       `json:"message"` // Descriptive message
	Data    []*TopupYearlyMethodResponse `json:"data"`    // Array of yearly method data
}

// ApiResponseTopupMonthAmount is the API response for monthly top-up amount stats
type ApiResponseTopupMonthAmount struct {
	Status  string                      `json:"status"`  // Response status
	Message string                      `json:"message"` // Descriptive message
	Data    []*TopupMonthAmountResponse `json:"data"`    // Array of monthly amount data
}

// ApiResponseTopupYearAmount is the API response for yearly top-up amount stats
type ApiResponseTopupYearAmount struct {
	Status  string                       `json:"status"`  // Response status
	Message string                       `json:"message"` // Descriptive message
	Data    []*TopupYearlyAmountResponse `json:"data"`    // Array of yearly amount data
}

// ApiResponseTopup is the API response for a single top-up transaction
type ApiResponseTopup struct {
	Status  string         `json:"status"`  // Response status
	Message string         `json:"message"` // Descriptive message
	Data    *TopupResponse `json:"data"`    // Single top-up data
}

// ApiResponseTopupDeleteAt is the API response for a single top-up including delete info
type ApiResponseTopupDeleteAt struct {
	Status  string                 `json:"status"`  // Response status
	Message string                 `json:"message"` // Descriptive message
	Data    *TopupResponseDeleteAt `json:"data"`    // Single top-up data with delete info
}

// ApiResponsesTopup is the API response for multiple top-up transactions
type ApiResponsesTopup struct {
	Status  string           `json:"status"`  // Response status
	Message string           `json:"message"` // Descriptive message
	Data    []*TopupResponse `json:"data"`    // Array of top-up data
}

// ApiResponsePaginationTopup is the paginated API response for top-ups
type ApiResponsePaginationTopup struct {
	Status     string           `json:"status"`     // Response status
	Message    string           `json:"message"`    // Descriptive message
	Data       []*TopupResponse `json:"data"`       // Array of top-up data
	Pagination *PaginationMeta  `json:"pagination"` // Pagination metadata
}

// ApiResponsePaginationTopupDeleteAt is the paginated API response for top-ups with delete info
type ApiResponsePaginationTopupDeleteAt struct {
	Status     string                   `json:"status"`     // Response status
	Message    string                   `json:"message"`    // Descriptive message
	Data       []*TopupResponseDeleteAt `json:"data"`       // Array of top-up data with delete info
	Pagination *PaginationMeta          `json:"pagination"` // Pagination metadata
}

// ApiResponseTopupDelete is the API response for a delete operation
type ApiResponseTopupDelete struct {
	Status  string `json:"status"`  // Response status
	Message string `json:"message"` // Descriptive message
}

// ApiResponseTopupAll is a generic API response for top-up operations
type ApiResponseTopupAll struct {
	Status  string `json:"status"`  // Response status
	Message string `json:"message"` // Descriptive message
}
