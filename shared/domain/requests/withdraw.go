package requests

import (
	"errors"
	"time"

	"github.com/go-playground/validator/v10"
)

// YearMonthCardNumber represents parameters for operations requiring card number and year.
// Used for card-specific annual data retrieval operations.
type YearMonthCardNumber struct {
	CardNumber string `json:"card_number" validate:"required,min=1"` // Card number to query (minimum 1 character)
	Year       int    `json:"year" validate:"required"`              // Year for the query (YYYY format)
}

// MonthStatusWithdraw represents parameters for retrieving monthly withdrawal status.
// Used to get general monthly withdrawal statistics.
type MonthStatusWithdraw struct {
	Year  int `json:"year" validate:"required"`  // Year for analysis (YYYY format)
	Month int `json:"month" validate:"required"` // Month for analysis (1-12)
}

// MonthStatusWithdrawCardNumber represents parameters for retrieving card-specific monthly withdrawal status.
// Used to get monthly withdrawal statistics for specific cards.
type MonthStatusWithdrawCardNumber struct {
	CardNumber string `json:"card_number" validate:"required,min=1"` // Card number to analyze
	Year       int    `json:"year" validate:"required"`              // Year for analysis (YYYY format)
	Month      int    `json:"month" validate:"required"`             // Month for analysis (1-12)
}

// YearStatusWithdrawCardNumber represents parameters for retrieving card-specific annual withdrawal status.
// Used to get yearly withdrawal statistics for specific cards.
type YearStatusWithdrawCardNumber struct {
	CardNumber string `json:"card_number" validate:"required,min=1"` // Card number to analyze
	Year       int    `json:"year" validate:"required"`              // Year for analysis (YYYY format)
}

// FindAllWithdraws represents parameters for searching and paginating withdrawal records.
// Used in withdrawal administration interfaces.
type FindAllWithdraws struct {
	Search   string `json:"search" validate:"required"`         // Search term to filter withdrawals
	Page     int    `json:"page" validate:"min=1"`              // Page number for pagination (1-based index)
	PageSize int    `json:"page_size" validate:"min=1,max=100"` // Number of items per page (1-100)
}

// FindAllWithdrawCardNumber represents parameters for searching withdrawals by card number.
// Used to retrieve withdrawal history for specific cards.
type FindAllWithdrawCardNumber struct {
	CardNumber string `json:"card_number" validate:"required,min=1"` // Specific card number to filter by
	Search     string `json:"search" validate:"required"`            // Additional search filter
	Page       int    `json:"page" validate:"min=1"`                 // Page number for pagination
	PageSize   int    `json:"page_size" validate:"min=1,max=100"`    // Number of items per page
}

// CreateWithdrawRequest represents the payload for creating a new withdrawal.
// Used when processing cash withdrawals from cards/accounts.
type CreateWithdrawRequest struct {
	CardNumber     string    `json:"card_number" validate:"required,min=1"`         // Card/account number for withdrawal
	WithdrawAmount int       `json:"withdraw_amount" validate:"required,min=50000"` // Amount to withdraw (minimum 50,000 in smallest unit)
	WithdrawTime   time.Time `json:"withdraw_time" validate:"required"`             // Timestamp of withdrawal
	IdempotencyKey string    `json:"idempotency_key" validate:"omitempty,max=64"`   // Client-supplied key for safe retries (max 64 chars)
}

// UpdateWithdrawRequest represents the payload for modifying a withdrawal record.
// Used to update existing withdrawal details.
type UpdateWithdrawRequest struct {
	CardNumber     string    `json:"card_number" validate:"required,min=1"`         // Card/account number
	WithdrawID     *int      `json:"withdraw_id"`                                   // ID of withdrawal to update
	WithdrawAmount int       `json:"withdraw_amount" validate:"required,min=50000"` // Updated withdrawal amount
	WithdrawTime   time.Time `json:"withdraw_time" validate:"required"`             // Updated withdrawal timestamp
}

// UpdateWithdrawStatus represents the payload for changing a withdrawal status.
// Used to update processing status of withdrawals.
type UpdateWithdrawStatus struct {
	WithdrawID int    `json:"withdraw_id" validate:"required,min=1"` // ID of withdrawal to update
	Status     string `json:"status" validate:"required"`            // New status (e.g., "completed", "failed")
}

// Validate performs comprehensive validation of CreateWithdrawRequest fields.
// Ensures:
// - Minimum withdrawal amount (50,000)
// - Withdrawal time is not in the future
// - Required fields are present
// Returns:
//   - error: if validation fails with specific validation messages
func (r *CreateWithdrawRequest) Validate() error {
	validate := validator.New()
	if err := validate.Struct(r); err != nil {
		return err
	}

	if r.WithdrawAmount < 50000 {
		return errors.New("withdraw amount must be at least 50,000")
	}

	if r.WithdrawTime.After(time.Now()) {
		return errors.New("withdraw time cannot be in the future")
	}

	return nil
}

// Validate performs comprehensive validation of UpdateWithdrawRequest fields.
// Ensures:
// - Minimum withdrawal amount (50,000)
// - Withdrawal time is not in the future
// - Required fields are present
// Returns:
//   - error: if validation fails with specific validation messages
func (r *UpdateWithdrawRequest) Validate() error {
	validate := validator.New()
	if err := validate.Struct(r); err != nil {
		return err
	}

	if r.WithdrawAmount < 50000 {
		return errors.New("withdraw amount must be at least 50,000")
	}

	if r.WithdrawTime.After(time.Now()) {
		return errors.New("withdraw time cannot be in the future")
	}

	return nil
}

// Validate performs basic validation of UpdateWithdrawStatus fields.
// Ensures required fields are present.
// Returns:
//   - error: if validation fails
func (r *UpdateWithdrawStatus) Validate() error {
	validate := validator.New()
	if err := validate.Struct(r); err != nil {
		return err
	}
	return nil
}
