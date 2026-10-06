package requests

import (
	"fmt"
	"time"

	methodtopup "github.com/MamangRust/monolith-payment-gateway-pkg/method_topup"
	"github.com/go-playground/validator/v10"
)

// MonthYearPaymentMethod represents parameters for retrieving payment method statistics by card and year.
// Used to analyze payment method usage patterns for specific cards.
type MonthYearPaymentMethod struct {
	CardNumber string `json:"card_number" validate:"required,min=1"` // Card number to analyze
	Year       int    `json:"year" validate:"required"`              // Year for analysis (YYYY format)
}

// MonthStatusTransaction represents parameters for retrieving transaction status by month.
// Used to get monthly transaction statistics.
type MonthStatusTransaction struct {
	Year  int `json:"year" validate:"required"`  // Year for analysis (YYYY format)
	Month int `json:"month" validate:"required"` // Month for analysis (1-12)
}

// YearStatusTransactionCardNumber represents parameters for retrieving card-specific annual transaction status.
// Used to get yearly transaction statistics for specific cards.
type YearStatusTransactionCardNumber struct {
	CardNumber string `json:"card_number" validate:"required,min=1"` // Card number to analyze
	Year       int    `json:"year" validate:"required"`              // Year for analysis (YYYY format)
}

// MonthStatusTransactionCardNumber represents parameters for retrieving card-specific monthly transaction status.
// Used to get monthly transaction statistics for specific cards.
type MonthStatusTransactionCardNumber struct {
	CardNumber string `json:"card_number" validate:"required,min=1"` // Card number to analyze
	Year       int    `json:"year" validate:"required"`              // Year for analysis (YYYY format)
	Month      int    `json:"month" validate:"required"`             // Month for analysis (1-12)
}

// FindAllTransactions represents parameters for searching and paginating transaction records.
// Used in transaction administration interfaces.
type FindAllTransactions struct {
	Search   string `json:"search" validate:"required"`         // Search term to filter transactions
	Page     int    `json:"page" validate:"min=1"`              // Page number for pagination (1-based index)
	PageSize int    `json:"page_size" validate:"min=1,max=100"` // Number of items per page (1-100)
}

// FindAllTransactionCardNumber represents parameters for searching transactions by card number.
// Used to retrieve transaction history for specific cards.
type FindAllTransactionCardNumber struct {
	CardNumber string `json:"card_number" validate:"required,min=1"` // Specific card number to filter by
	Search     string `json:"search" validate:"required"`            // Additional search filter
	Page       int    `json:"page" validate:"min=1"`                 // Page number for pagination
	PageSize   int    `json:"page_size" validate:"min=1,max=100"`    // Number of items per page
}

// CreateTransactionRequest represents the payload for creating a new transaction.
// Used when recording payment transactions.
type CreateTransactionRequest struct {
	CardNumber      string    `json:"card_number" validate:"required,min=1"` // Card number used in transaction
	Amount          int       `json:"amount" validate:"required,min=50000"`  // Transaction amount (minimum 50,000 in smallest unit)
	PaymentMethod   string    `json:"payment_method" validate:"required"`    // Payment method used (e.g., "credit", "debit")
	MerchantID      *int      `json:"merchant_id" validate:"required,min=1"` // ID of merchant receiving payment
	TransactionTime time.Time `json:"transaction_time" validate:"required"`  // Timestamp of transaction
	IdempotencyKey  string    `json:"idempotency_key" validate:"omitempty,max=64"` // Client-supplied key for safe retries (max 64 chars)
}

// UpdateTransactionRequest represents the payload for updating a transaction record.
// Used to modify existing transaction details.
type UpdateTransactionRequest struct {
	TransactionID   *int      `json:"transaction_id"`                        // ID of transaction to update
	CardNumber      string    `json:"card_number" validate:"required,min=1"` // Updated card number
	Amount          int       `json:"amount" validate:"required,min=50000"`  // Updated amount (minimum 50,000)
	PaymentMethod   string    `json:"payment_method" validate:"required"`    // Updated payment method
	MerchantID      *int      `json:"merchant_id" validate:"required,min=1"` // Updated merchant ID
	TransactionTime time.Time `json:"transaction_time" validate:"required"`  // Updated transaction timestamp
}

// UpdateTransactionStatus represents the payload for changing a transaction status.
// Used to update processing status of transactions.
type UpdateTransactionStatus struct {
	TransactionID int    `json:"transaction_id" validate:"required,min=1"` // ID of transaction to update
	Status        string `json:"status" validate:"required"`               // New status (e.g., "completed", "failed")
}

// Validate performs comprehensive validation of CreateTransactionRequest fields.
// Ensures:
// - Valid payment method
// - Minimum transaction amount (50,000)
// - Required fields are present
// Returns:
//   - error: if validation fails with specific validation messages
func (r *CreateTransactionRequest) Validate() error {
	validate := validator.New()
	err := validate.Struct(r)

	if !methodtopup.PaymentMethodValidator(r.PaymentMethod) {
		return fmt.Errorf("payment method not found")
	}

	if err != nil {
		return err
	}

	return nil
}

// Validate performs comprehensive validation of UpdateTransactionRequest fields.
// Ensures:
// - Valid payment method
// - Minimum transaction amount (50,000)
// - Required fields are present
// Returns:
//   - error: if validation fails with specific validation messages
func (r *UpdateTransactionRequest) Validate() error {
	validate := validator.New()
	err := validate.Struct(r)

	if !methodtopup.PaymentMethodValidator(r.PaymentMethod) {
		return fmt.Errorf("payment method not found")
	}

	if err != nil {
		return err
	}

	return nil
}

// Validate performs basic validation of UpdateTransactionStatus fields.
// Ensures required fields are present.
// Returns:
//   - error: if validation fails
func (r *UpdateTransactionStatus) Validate() error {
	validate := validator.New()
	if err := validate.Struct(r); err != nil {
		return err
	}
	return nil
}

// UpdateFraudScore represents the payload for updating the fraud score of a transaction.
type UpdateFraudScore struct {
	TransactionID int `json:"transaction_id" validate:"required,min=1"` // ID of transaction
	FraudScore    int `json:"fraud_score" validate:"min=0,max=100"`     // Fraud score (0-100)
}
