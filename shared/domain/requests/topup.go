package requests

import (
	"errors"

	methodtopup "github.com/MamangRust/monolith-payment-gateway-pkg/method_topup"
	"github.com/go-playground/validator/v10"
)

// MonthTopupStatus represents parameters for retrieving top-up status by month.
// Used to query top-up statistics for a specific month and year.
type MonthTopupStatus struct {
	Year  int `json:"year" validate:"required"`  // Year for the query (YYYY format)
	Month int `json:"month" validate:"required"` // Month for the query (1-12)
}

// MonthTopupStatusCardNumber represents parameters for retrieving card-specific top-up status by month.
// Used to query monthly top-up statistics for a specific card.
type MonthTopupStatusCardNumber struct {
	CardNumber string `json:"card_number" validate:"required,min=1"` // Card number to query
	Year       int    `json:"year" validate:"required"`              // Year for the query (YYYY format)
	Month      int    `json:"month" validate:"required"`             // Month for the query (1-12)
}

// YearTopupStatusCardNumber represents parameters for retrieving card-specific top-up status by year.
// Used to query annual top-up statistics for a specific card.
type YearTopupStatusCardNumber struct {
	CardNumber string `json:"card_number" validate:"required,min=1"` // Card number to query
	Year       int    `json:"year" validate:"required"`              // Year for the query (YYYY format)
}

// YearMonthMethod represents parameters for retrieving top-up methods by year and card.
// Used to analyze payment method trends for specific cards.
type YearMonthMethod struct {
	CardNumber string `json:"card_number" validate:"required,min=1"` // Card number to query
	Year       int    `json:"year" validate:"required"`              // Year for the query (YYYY format)
}

// FindAllTopups represents parameters for searching and paginating top-up records.
// Used in top-up administration and reporting interfaces.
type FindAllTopups struct {
	Search   string `json:"search" validate:"required"`         // Search term to filter records
	Page     int    `json:"page" validate:"min=1"`              // Page number for pagination (1-based index)
	PageSize int    `json:"page_size" validate:"min=1,max=100"` // Number of items per page (1-100)
}

// FindAllTopupsByCardNumber represents parameters for searching top-ups by card number.
// Used to retrieve top-up history for specific cards.
type FindAllTopupsByCardNumber struct {
	CardNumber string `json:"card_number" validate:"required,min=1"` // Specific card number to filter by
	Search     string `json:"search" validate:"required"`            // Additional search filter
	Page       int    `json:"page" validate:"min=1"`                 // Page number for pagination
	PageSize   int    `json:"page_size" validate:"min=1,max=100"`    // Number of items per page
}

// CreateTopupRequest represents the payload for creating a new top-up transaction.
// Used when adding funds to a card/account.
type CreateTopupRequest struct {
	CardNumber     string `json:"card_number" validate:"required,min=1"`       // Card number receiving funds
	TopupAmount    int    `json:"topup_amount" validate:"required,min=50000"`  // Amount to add (minimum 50,000 in smallest unit)
	TopupMethod    string `json:"topup_method" validate:"required"`            // Payment method used (e.g., "bank_transfer")
	IdempotencyKey string `json:"idempotency_key" validate:"omitempty,max=64"` // Client-supplied key for safe retries (max 64 chars)
}

// UpdateTopupRequest represents the payload for updating a top-up transaction.
// Used to modify existing top-up records.
type UpdateTopupRequest struct {
	CardNumber  string `json:"card_number" validate:"required,min=1"`      // Card number associated with top-up
	TopupID     *int   `json:"topup_id"`                                   // ID of top-up record to update
	TopupAmount int    `json:"topup_amount" validate:"required,min=50000"` // Updated amount (minimum 50,000)
	TopupMethod string `json:"topup_method" validate:"required"`           // Updated payment method
}

// UpdateTopupAmount represents the payload for adjusting a top-up amount.
// Used specifically for amount corrections.
type UpdateTopupAmount struct {
	TopupID     int `json:"topup_id" validate:"required,min=1"`         // ID of top-up record
	TopupAmount int `json:"topup_amount" validate:"required,min=50000"` // New amount (minimum 50,000)
}

// UpdateTopupStatus represents the payload for changing a top-up status.
// Used to update processing status of top-up transactions.
type UpdateTopupStatus struct {
	TopupID int    `json:"topup_id" validate:"required,min=1"` // ID of top-up record
	Status  string `json:"status" validate:"required"`         // New status (e.g., "completed", "failed")
}

// Validate performs comprehensive validation of CreateTopupRequest fields.
// Ensures:
// - Minimum top-up amount (50,000)
// - Valid payment method
// - Required fields are present
// Returns:
//   - error: if validation fails with specific validation messages
func (r *CreateTopupRequest) Validate() error {
	validate := validator.New()
	if err := validate.Struct(r); err != nil {
		return err
	}

	if r.TopupAmount < 50000 {
		return errors.New("topup amount must be greater than or equal to 50000")
	}

	if r.TopupMethod == "" {
		return errors.New("top-up method is required")
	}

	if !methodtopup.PaymentMethodValidator(r.TopupMethod) {
		return errors.New("topup method not found")
	}

	return nil
}

// Validate performs comprehensive validation of UpdateTopupRequest fields.
// Ensures:
// - Valid top-up ID
// - Minimum top-up amount (50,000)
// - Valid payment method
// Returns:
//   - error: if validation fails with specific validation messages
func (r *UpdateTopupRequest) Validate() error {
	validate := validator.New()
	if err := validate.Struct(r); err != nil {
		return err
	}

	if r.TopupID == nil || *r.TopupID <= 0 {
		return errors.New("top-up ID must be a positive integer")
	}

	if r.TopupAmount < 50000 {
		return errors.New("topup amount must be greater than or equal to 50000")
	}

	if r.TopupMethod == "" {
		return errors.New("top-up method is required")
	}

	if !methodtopup.PaymentMethodValidator(r.TopupMethod) {
		return errors.New("topup method not found")
	}

	return nil
}

// Validate performs validation of UpdateTopupAmount fields.
// Ensures:
// - Valid top-up ID
// - Minimum top-up amount (50,000)
// Returns:
//   - error: if validation fails
func (r *UpdateTopupAmount) Validate() error {
	validate := validator.New()
	if err := validate.Struct(r); err != nil {
		return err
	}

	if r.TopupID <= 0 {
		return errors.New("top-up ID must be a positive integer")
	}

	if r.TopupAmount < 50000 {
		return errors.New("topup amount must be greater than or equal to 50000")
	}

	return nil
}

// Validate performs basic validation of UpdateTopupStatus fields.
// Ensures required fields are present.
// Returns:
//   - error: if validation fails
func (r *UpdateTopupStatus) Validate() error {
	validate := validator.New()
	if err := validate.Struct(r); err != nil {
		return err
	}
	return nil
}
