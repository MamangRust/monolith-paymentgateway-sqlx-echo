package response

// TransferResponse represents a fund transfer between accounts
type TransferResponse struct {
	ID             int    `json:"id"`              // Unique transfer identifier
	TransferNo     string `json:"transfer_no"`     // Unique transfer reference number
	TransferFrom   string `json:"transfer_from"`   // Source account number
	TransferTo     string `json:"transfer_to"`     // Destination account number
	TransferAmount int    `json:"transfer_amount"` // Amount transferred
	TransferTime   string `json:"transfer_time"`   // Timestamp when transfer was executed
	CreatedAt      string `json:"created_at"`      // When the record was created
	UpdatedAt      string `json:"updated_at"`      // When the record was last updated
}

// TransferResponseDeleteAt extends TransferResponse with soft delete capability
type TransferResponseDeleteAt struct {
	ID             int     `json:"id"`
	TransferNo     string  `json:"transfer_no"`
	TransferFrom   string  `json:"transfer_from"`
	TransferTo     string  `json:"transfer_to"`
	TransferAmount int     `json:"transfer_amount"`
	TransferTime   string  `json:"transfer_time"`
	CreatedAt      string  `json:"created_at"`
	UpdatedAt      string  `json:"updated_at"`
	DeletedAt      *string `json:"deleted_at"` // Timestamp when record was soft deleted (nil if active)
	// Embedded TransferResponse fields
}

// TransferResponseMonthStatusSuccess represents monthly successful transfer metrics
type TransferResponseMonthStatusSuccess struct {
	Year         string `json:"year"`          // Year of the metrics
	Month        string `json:"month"`         // Month of the metrics
	TotalAmount  int    `json:"total_amount"`  // Total amount successfully transferred
	TotalSuccess int    `json:"total_success"` // Count of successful transfers
}

// TransferResponseYearStatusSuccess represents yearly successful transfer metrics
type TransferResponseYearStatusSuccess struct {
	Year         string `json:"year"`          // Year of the metrics
	TotalAmount  int    `json:"total_amount"`  // Total amount successfully transferred
	TotalSuccess int    `json:"total_success"` // Count of successful transfers
}

// TransferResponseMonthStatusFailed represents monthly failed transfer metrics
type TransferResponseMonthStatusFailed struct {
	Year        string `json:"year"`         // Year of the metrics
	Month       string `json:"month"`        // Month of the metrics
	TotalAmount int    `json:"total_amount"` // Total amount of failed transfers
	TotalFailed int    `json:"total_failed"` // Count of failed transfers
}

// TransferResponseYearStatusFailed represents yearly failed transfer metrics
type TransferResponseYearStatusFailed struct {
	Year        string `json:"year"`         // Year of the metrics
	TotalAmount int    `json:"total_amount"` // Total amount of failed transfers
	TotalFailed int    `json:"total_failed"` // Count of failed transfers
}

// TransferMonthAmountResponse represents monthly transfer amount metrics
type TransferMonthAmountResponse struct {
	Month       string `json:"month"`        // Month of the metrics
	TotalAmount int    `json:"total_amount"` // Total amount transferred
}

// TransferYearAmountResponse represents yearly transfer amount metrics
type TransferYearAmountResponse struct {
	Year        string `json:"year"`         // Year of the metrics
	TotalAmount int    `json:"total_amount"` // Total amount transferred
}

// ApiResponseTransferMonthStatusSuccess is API response for monthly success metrics
type ApiResponseTransferMonthStatusSuccess struct {
	Status  string                                `json:"status"`  // Response status
	Message string                                `json:"message"` // Response message
	Data    []*TransferResponseMonthStatusSuccess `json:"data"`    // Array of monthly success data
}

// ApiResponseTransferYearStatusSuccess is API response for yearly success metrics
type ApiResponseTransferYearStatusSuccess struct {
	Status  string                               `json:"status"`  // Response status
	Message string                               `json:"message"` // Response message
	Data    []*TransferResponseYearStatusSuccess `json:"data"`    // Array of yearly success data
}

// ApiResponseTransferMonthStatusFailed is API response for monthly failed metrics
type ApiResponseTransferMonthStatusFailed struct {
	Status  string                               `json:"status"`  // Response status
	Message string                               `json:"message"` // Response message
	Data    []*TransferResponseMonthStatusFailed `json:"data"`    // Array of monthly failed data
}

// ApiResponseTransferYearStatusFailed is API response for yearly failed metrics
type ApiResponseTransferYearStatusFailed struct {
	Status  string                              `json:"status"`  // Response status
	Message string                              `json:"message"` // Response message
	Data    []*TransferResponseYearStatusFailed `json:"data"`    // Array of yearly failed data
}

// ApiResponseTransferMonthAmount is API response for monthly amount metrics
type ApiResponseTransferMonthAmount struct {
	Status  string                         `json:"status"`  // Response status
	Message string                         `json:"message"` // Response message
	Data    []*TransferMonthAmountResponse `json:"data"`    // Array of monthly amount data
}

// ApiResponseTransferYearAmount is API response for yearly amount metrics
type ApiResponseTransferYearAmount struct {
	Status  string                        `json:"status"`  // Response status
	Message string                        `json:"message"` // Response message
	Data    []*TransferYearAmountResponse `json:"data"`    // Array of yearly amount data
}

// ApiResponseTransfer is API response for a single transfer
type ApiResponseTransfer struct {
	Status  string            `json:"status"`  // Response status
	Message string            `json:"message"` // Response message
	Data    *TransferResponse `json:"data"`    // Single transfer data
}

// ApiResponseTransferDeleteAt is API response for soft-deleted transfer
type ApiResponseTransferDeleteAt struct {
	Status  string                    `json:"status"`  // Response status
	Message string                    `json:"message"` // Response message
	Data    *TransferResponseDeleteAt `json:"data"`
}

// ApiResponseTransfers is API response for multiple transfers
type ApiResponseTransfers struct {
	Status  string              `json:"status"`  // Response status
	Message string              `json:"message"` // Response message
	Data    []*TransferResponse `json:"data"`    // Array of transfer data
}

// ApiResponseTransferDelete is API response for transfer deletion
type ApiResponseTransferDelete struct {
	Status  string `json:"status"`  // Response status
	Message string `json:"message"` // Response message
}

// ApiResponseTransferAll is generic API response for transfer operations
type ApiResponseTransferAll struct {
	Status  string `json:"status"`  // Response status
	Message string `json:"message"` // Response message
}

// ApiResponsePaginationTransfer is paginated API response for transfers
type ApiResponsePaginationTransfer struct {
	Status     string              `json:"status"`     // Response status
	Message    string              `json:"message"`    // Response message
	Data       []*TransferResponse `json:"data"`       // Array of transfer data
	Pagination *PaginationMeta     `json:"pagination"` // Pagination metadata
}

// ApiResponsePaginationTransferDeleteAt is paginated API response for soft-deleted transfers
type ApiResponsePaginationTransferDeleteAt struct {
	Status     string                      `json:"status"`     // Response status
	Message    string                      `json:"message"`    // Response message
	Data       []*TransferResponseDeleteAt `json:"data"`       // Array of transfer data with delete info
	Pagination *PaginationMeta             `json:"pagination"` // Pagination metadata
}
