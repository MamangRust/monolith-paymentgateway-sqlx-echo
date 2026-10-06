package requests

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

// MonthYearCardNumber represents parameters for operations requiring card number and year.
// Used for card-specific annual data retrieval operations.
type MonthYearCardNumber struct {
	CardNumber string `json:"card_number" validate:"required,min=1"` // Card number to query (minimum 1 character)
	Year       int    `json:"year" validate:"required"`              // Year for the query (YYYY format)
}

// MonthStatusTransferCardNumber represents parameters for retrieving monthly transfer status by card.
// Used to get monthly transfer statistics for specific cards.
type MonthStatusTransferCardNumber struct {
	CardNumber string `json:"card_number" validate:"required,min=1"` // Card number to analyze
	Year       int    `json:"year" validate:"required"`              // Year for analysis (YYYY format)
	Month      int    `json:"month" validate:"required"`             // Month for analysis (1-12)
}

// YearStatusTransferCardNumber represents parameters for retrieving annual transfer status by card.
// Used to get yearly transfer statistics for specific cards.
type YearStatusTransferCardNumber struct {
	CardNumber string `json:"card_number" validate:"required,min=1"` // Card number to analyze
	Year       int    `json:"year" validate:"required"`              // Year for analysis (YYYY format)
}

// MonthStatusTransfer represents parameters for retrieving monthly transfer status.
// Used for general monthly transfer statistics.
type MonthStatusTransfer struct {
	Year  int `json:"year" validate:"required"`  // Year for analysis (YYYY format)
	Month int `json:"month" validate:"required"` // Month for analysis (1-12)
}

// FindAllTransfers represents parameters for searching and paginating transfer records.
// Used in transfer administration interfaces.
type FindAllTransfers struct {
	Search   string `json:"search" validate:"required"`         // Search term to filter transfers
	Page     int    `json:"page" validate:"min=1"`              // Page number for pagination (1-based index)
	PageSize int    `json:"page_size" validate:"min=1,max=100"` // Number of items per page (1-100)
}

// CreateTransferRequest represents the payload for initiating a new transfer.
// Used when moving funds between accounts/cards.
type CreateTransferRequest struct {
	TransferFrom   string `json:"transfer_from" validate:"required"`             // Source account/card number
	TransferTo     string `json:"transfer_to" validate:"required,min=1"`         // Destination account/card number (minimum 1 character)
	TransferAmount int    `json:"transfer_amount" validate:"required,min=50000"` // Amount to transfer (minimum 50,000 in smallest unit)
	IdempotencyKey string `json:"idempotency_key" validate:"omitempty,max=64"`   // Client-supplied key for safe retries (max 64 chars)
}

// UpdateTransferRequest represents the payload for modifying a transfer record.
// Used to update existing transfer details.
type UpdateTransferRequest struct {
	TransferID     *int   `json:"transfer_id"`                                   // ID of transfer to update
	TransferFrom   string `json:"transfer_from" validate:"required"`             // Updated source account/card
	TransferTo     string `json:"transfer_to" validate:"required,min=1"`         // Updated destination account/card
	TransferAmount int    `json:"transfer_amount" validate:"required,min=50000"` // Updated transfer amount (minimum 50,000)
}

// UpdateTransferAmountRequest represents the payload for adjusting a transfer amount.
// Used specifically for amount corrections.
type UpdateTransferAmountRequest struct {
	TransferID     int `json:"transfer_id" validate:"required,min=1"`    // ID of transfer record
	TransferAmount int `json:"transfer_amount" validate:"required,gt=0"` // New amount (must be greater than 0)
}

// UpdateTransferStatus represents the payload for changing a transfer status.
// Used to update processing status of transfers.
type UpdateTransferStatus struct {
	TransferID int    `json:"transfer_id" validate:"required,min=1"` // ID of transfer to update
	Status     string `json:"status" validate:"required"`            // New status (e.g., "completed", "failed")
}

// Validate performs comprehensive validation of CreateTransferRequest fields.
// Ensures:
// - Minimum transfer amount (50,000)
// - Valid account/card numbers
// - Required fields are present
// Returns:
//   - error: if validation fails with specific validation messages
func (r *CreateTransferRequest) Validate() error {
	validate := validator.New()
	if err := validate.Struct(r); err != nil {
		return err
	}

	if r.TransferAmount < 50000 {
		return errors.New("transfer amount must be at least 50,000")
	}

	return nil
}

// Validate performs comprehensive validation of UpdateTransferRequest fields.
// Ensures:
// - Valid transfer ID
// - Minimum transfer amount (50,000)
// - Valid account/card numbers
// Returns:
//   - error: if validation fails with specific validation messages
func (r *UpdateTransferRequest) Validate() error {
	validate := validator.New()
	if err := validate.Struct(r); err != nil {
		return err
	}

	if *r.TransferID <= 0 {
		return errors.New("transfer ID must be a positive integer")
	}

	if r.TransferAmount < 50000 {
		return errors.New("transfer amount must be at least 50,000")
	}

	return nil
}

// Validate performs validation of UpdateTransferAmountRequest fields.
// Ensures:
// - Valid transfer ID
// - Positive transfer amount
// Returns:
//   - error: if validation fails
func (r *UpdateTransferAmountRequest) Validate() error {
	validate := validator.New()
	if err := validate.Struct(r); err != nil {
		return err
	}

	if r.TransferID <= 0 {
		return errors.New("transfer ID must be a positive integer")
	}

	if r.TransferAmount <= 0 {
		return errors.New("transfer amount must be greater than zero")
	}

	return nil
}

// Validate performs basic validation of UpdateTransferStatus fields.
// Ensures required fields are present.
// Returns:
//   - error: if validation fails
func (r *UpdateTransferStatus) Validate() error {
	validate := validator.New()
	if err := validate.Struct(r); err != nil {
		return err
	}
	return nil
}
