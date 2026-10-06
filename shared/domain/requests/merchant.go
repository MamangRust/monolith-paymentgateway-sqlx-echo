package requests

import "github.com/go-playground/validator/v10"

// MerchantRequestPayload represents the base payload structure for merchant API requests.
// Used as a foundation for authenticated merchant operations.
type MerchantRequestPayload struct {
	ApiKey        string `json:"api_key"`        // Merchant's authentication API key
	CorrelationID string `json:"correlation_id"` // Unique ID for request tracing
	ReplyTopic    string `json:"reply_topic"`    // Topic name for asynchronous responses
}

// MonthYearPaymentMethodApiKey represents a request for payment method statistics by API key and year.
// Used to retrieve aggregated payment method data for a specific merchant.
type MonthYearPaymentMethodApiKey struct {
	Apikey string `json:"api_key" validate:"required,min=1"` // Merchant's authentication API key
	Year   int    `json:"year" validate:"required"`          // Year for data aggregation (YYYY format)
}

// MonthYearAmountApiKey represents a request for transaction amount statistics by API key and year.
// Used to retrieve financial summaries for a specific merchant.
type MonthYearAmountApiKey struct {
	Apikey string `json:"api_key" validate:"required,min=1"` // Merchant's authentication API key
	Year   int    `json:"year" validate:"required"`          // Year for data aggregation (YYYY format)
}

// MonthYearTotalAmountApiKey represents a request for total transaction amounts by API key and year.
// Used for annual financial reporting per merchant.
type MonthYearTotalAmountApiKey struct {
	Apikey string `json:"api_key" validate:"required,min=1"` // Merchant's authentication API key
	Year   int    `json:"year" validate:"required"`          // Year for data aggregation (YYYY format)
}

// MonthYearPaymentMethodMerchant represents a request for payment method statistics by merchant ID and year.
// Used internally for merchant analytics.
type MonthYearPaymentMethodMerchant struct {
	MerchantID int `json:"merchant_id" validate:"required,min=1"` // Unique merchant identifier
	Year       int `json:"year" validate:"required"`              // Year for data aggregation (YYYY format)
}

// MonthYearAmountMerchant represents a request for transaction amount statistics by merchant ID and year.
// Used internally for financial reporting.
type MonthYearAmountMerchant struct {
	MerchantID int `json:"merchant_id" validate:"required,min=1"` // Unique merchant identifier
	Year       int `json:"year" validate:"required"`              // Year for data aggregation (YYYY format)
}

// MonthYearTotalAmountMerchant represents a request for total transaction amounts by merchant ID and year.
// Used for comprehensive merchant financial analysis.
type MonthYearTotalAmountMerchant struct {
	MerchantID int `json:"merchant_id" validate:"required,min=1"` // Unique merchant identifier
	Year       int `json:"year" validate:"required"`              // Year for data aggregation (YYYY format)
}

// FindAllMerchants represents parameters for searching and paginating merchant records.
// Used in merchant administration interfaces.
type FindAllMerchants struct {
	Search   string `json:"search" validate:"required"`         // Search term to filter merchants
	Page     int    `json:"page" validate:"min=1"`              // Page number for pagination
	PageSize int    `json:"page_size" validate:"min=1,max=100"` // Number of items per page
}

// FindAllMerchantTransactions represents parameters for searching merchant transactions.
// Used in global transaction reporting.
type FindAllMerchantTransactions struct {
	Search   string `json:"search" validate:"required"`         // Search term to filter transactions
	Page     int    `json:"page" validate:"min=1"`              // Page number for pagination
	PageSize int    `json:"page_size" validate:"min=1,max=100"` // Number of items per page
}

// FindAllMerchantTransactionsById represents parameters for searching transactions by merchant ID.
// Used for merchant-specific transaction reporting.
type FindAllMerchantTransactionsById struct {
	MerchantID int    `json:"merchant_id" validate:"required,min=1"` // Specific merchant to filter by
	Search     string `json:"search" validate:"required"`            // Additional search filter
	Page       int    `json:"page" validate:"min=1"`                 // Page number for pagination
	PageSize   int    `json:"page_size" validate:"min=1,max=100"`    // Number of items per page
}

// FindAllMerchantTransactionsByApiKey represents parameters for searching transactions by API key.
// Used in merchant-facing transaction reporting.
type FindAllMerchantTransactionsByApiKey struct {
	ApiKey   string `json:"api_key" validate:"required"`        // Merchant's authentication API key
	Search   string `json:"search" validate:"required"`         // Search term to filter transactions
	Page     int    `json:"page" validate:"min=1"`              // Page number for pagination
	PageSize int    `json:"page_size" validate:"min=1,max=100"` // Number of items per page
}

// CreateMerchantRequest represents the payload for registering a new merchant.
// Used in merchant onboarding processes.
type CreateMerchantRequest struct {
	Name   string `json:"name" validate:"required"`          // Legal name of the merchant
	UserID int    `json:"user_id" validate:"required,min=1"` // ID of the user creating the merchant
}

// UpdateMerchantRequest represents the payload for updating merchant details.
// Used in merchant profile management.
type UpdateMerchantRequest struct {
	MerchantID *int   `json:"merchant_id"`                       // ID of merchant to update (optional in some flows)
	Name       string `json:"name" validate:"required"`          // Updated merchant name
	UserID     int    `json:"user_id" validate:"required,min=1"` // ID of user performing update
	Status     string `json:"status" validate:"required"`        // New status (e.g., "active", "suspended")
}

// UpdateMerchantStatusRequest represents the payload for changing merchant status.
// Used specifically for merchant account status management.
type UpdateMerchantStatusRequest struct {
	MerchantID *int   `json:"merchant_id"`                // ID of merchant to update (optional in some flows)
	Status     string `json:"status" validate:"required"` // New status (e.g., "approved", "rejected")
}

// Validate performs validation of CreateMerchantRequest fields.
// Ensures all required fields are present and valid.
// Returns:
//   - error: if validation fails
func (r CreateMerchantRequest) Validate() error {
	validate := validator.New()

	if err := validate.Struct(r); err != nil {
		return err
	}

	return nil
}

// Validate performs validation of UpdateMerchantRequest fields.
// Ensures all required fields are present and valid.
// Returns:
//   - error: if validation fails
func (r UpdateMerchantRequest) Validate() error {
	validate := validator.New()

	if err := validate.Struct(r); err != nil {
		return err
	}

	return nil
}

// Validate performs validation of UpdateMerchantStatusRequest fields.
// Ensures all required fields are present and valid.
// Returns:
//   - error: if validation fails
func (r UpdateMerchantStatusRequest) Validate() error {
	validate := validator.New()

	if err := validate.Struct(r); err != nil {
		return err
	}

	return nil
}
