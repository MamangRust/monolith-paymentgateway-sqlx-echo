// Code migrated from sqlc-generated sources; domain types kept for repository layer.
// source: transaction.sql

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
	"time"
)

type AuthorizeTransactionAtomicParams struct {
	CardNumber      string    `db:"card_number" json:"card_number"`
	Amount          int32     `db:"amount" json:"amount"`
	PaymentMethod   string    `db:"payment_method" json:"payment_method"`
	MerchantID      int32     `db:"merchant_id" json:"merchant_id"`
	TransactionTime time.Time `db:"transaction_time" json:"transaction_time"`
	IdempotencyKey  string    `db:"idempotency_key" json:"idempotency_key"`
}

type AuthorizeTransactionAtomicRow struct {
	TransactionID   int32            `db:"transaction_id" json:"transaction_id"`
	TransactionNo   pgtype.UUID      `db:"transaction_no" json:"transaction_no"`
	CardNumber      string           `db:"card_number" json:"card_number"`
	Amount          int32            `db:"amount" json:"amount"`
	PaymentMethod   string           `db:"payment_method" json:"payment_method"`
	MerchantID      int32            `db:"merchant_id" json:"merchant_id"`
	TransactionTime time.Time        `db:"transaction_time" json:"transaction_time"`
	Status          string           `db:"status" json:"status"`
	CreatedAt       pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt       pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type CaptureTransactionAtomicParams struct {
	TransactionID      int32  `db:"transaction_id" json:"transaction_id"`
	MerchantCardNumber string `db:"merchant_card_number" json:"merchant_card_number"`
}

type CaptureTransactionAtomicRow struct {
	TransactionID   int32            `db:"transaction_id" json:"transaction_id"`
	TransactionNo   pgtype.UUID      `db:"transaction_no" json:"transaction_no"`
	CardNumber      string           `db:"card_number" json:"card_number"`
	Amount          int32            `db:"amount" json:"amount"`
	PaymentMethod   string           `db:"payment_method" json:"payment_method"`
	MerchantID      int32            `db:"merchant_id" json:"merchant_id"`
	TransactionTime time.Time        `db:"transaction_time" json:"transaction_time"`
	Status          string           `db:"status" json:"status"`
	CreatedAt       pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt       pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type CreateTransactionParams struct {
	CardNumber      string    `db:"card_number" json:"card_number"`
	Amount          int32     `db:"amount" json:"amount"`
	PaymentMethod   string    `db:"payment_method" json:"payment_method"`
	MerchantID      int32     `db:"merchant_id" json:"merchant_id"`
	TransactionTime time.Time `db:"transaction_time" json:"transaction_time"`
}

type CreateTransactionRow struct {
	TransactionID   int32            `db:"transaction_id" json:"transaction_id"`
	TransactionNo   pgtype.UUID      `db:"transaction_no" json:"transaction_no"`
	CardNumber      string           `db:"card_number" json:"card_number"`
	Amount          int32            `db:"amount" json:"amount"`
	PaymentMethod   string           `db:"payment_method" json:"payment_method"`
	MerchantID      int32            `db:"merchant_id" json:"merchant_id"`
	TransactionTime time.Time        `db:"transaction_time" json:"transaction_time"`
	Status          string           `db:"status" json:"status"`
	CreatedAt       pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt       pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type CreateTransactionAtomicParams struct {
	CardNumber      string    `db:"card_number" json:"card_number"`
	CardNumber2     string    `db:"card_number2" json:"card_number_2"`
	Amount          int32     `db:"amount" json:"amount"`
	PaymentMethod   string    `db:"payment_method" json:"payment_method"`
	MerchantID      int32     `db:"merchant_id" json:"merchant_id"`
	TransactionTime time.Time `db:"transaction_time" json:"transaction_time"`
	IdempotencyKey  string    `db:"idempotency_key" json:"idempotency_key"`
}

type CreateTransactionAtomicRow struct {
	TransactionID   int32            `db:"transaction_id" json:"transaction_id"`
	TransactionNo   pgtype.UUID      `db:"transaction_no" json:"transaction_no"`
	CardNumber      string           `db:"card_number" json:"card_number"`
	Amount          int32            `db:"amount" json:"amount"`
	PaymentMethod   string           `db:"payment_method" json:"payment_method"`
	MerchantID      int32            `db:"merchant_id" json:"merchant_id"`
	TransactionTime time.Time        `db:"transaction_time" json:"transaction_time"`
	Status          string           `db:"status" json:"status"`
	CreatedAt       pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt       pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type GetActiveTransactionsParams struct {
	Column1 string `db:"column1" json:"column_1"`
	Limit   int32  `db:"limit" json:"limit"`
	Offset  int32  `db:"offset" json:"offset"`
}

type GetActiveTransactionsRow struct {
	TransactionID   int32            `db:"transaction_id" json:"transaction_id"`
	TransactionNo   pgtype.UUID      `db:"transaction_no" json:"transaction_no"`
	CardNumber      string           `db:"card_number" json:"card_number"`
	Amount          int32            `db:"amount" json:"amount"`
	PaymentMethod   string           `db:"payment_method" json:"payment_method"`
	MerchantID      int32            `db:"merchant_id" json:"merchant_id"`
	TransactionTime time.Time        `db:"transaction_time" json:"transaction_time"`
	Status          string           `db:"status" json:"status"`
	CreatedAt       pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt       pgtype.Timestamp `db:"updated_at" json:"updated_at"`
	DeletedAt       pgtype.Timestamp `db:"deleted_at" json:"deleted_at"`
	TotalCount      int64            `db:"total_count" json:"total_count"`
}

type GetMonthTransactionStatusFailedParams struct {
	Column1 time.Time `db:"column1" json:"column_1"`
	Column2 time.Time `db:"column2" json:"column_2"`
	Column3 time.Time `db:"column3" json:"column_3"`
	Column4 time.Time `db:"column4" json:"column_4"`
}

type GetMonthTransactionStatusFailedRow struct {
	Year        string `db:"year" json:"year"`
	Month       string `db:"month" json:"month"`
	TotalFailed int64  `db:"total_failed" json:"total_failed"`
	TotalAmount int32  `db:"total_amount" json:"total_amount"`
}

type GetMonthTransactionStatusFailedCardNumberParams struct {
	CardNumber string    `db:"card_number" json:"card_number"`
	Column2    time.Time `db:"column2" json:"column_2"`
	Column3    time.Time `db:"column3" json:"column_3"`
	Column4    time.Time `db:"column4" json:"column_4"`
	Column5    time.Time `db:"column5" json:"column_5"`
}

type GetMonthTransactionStatusFailedCardNumberRow struct {
	Year        string `db:"year" json:"year"`
	Month       string `db:"month" json:"month"`
	TotalFailed int64  `db:"total_failed" json:"total_failed"`
	TotalAmount int32  `db:"total_amount" json:"total_amount"`
}

type GetMonthTransactionStatusSuccessParams struct {
	Column1 time.Time `db:"column1" json:"column_1"`
	Column2 time.Time `db:"column2" json:"column_2"`
	Column3 time.Time `db:"column3" json:"column_3"`
	Column4 time.Time `db:"column4" json:"column_4"`
}

type GetMonthTransactionStatusSuccessRow struct {
	Year         string `db:"year" json:"year"`
	Month        string `db:"month" json:"month"`
	TotalSuccess int64  `db:"total_success" json:"total_success"`
	TotalAmount  int32  `db:"total_amount" json:"total_amount"`
}

type GetMonthTransactionStatusSuccessCardNumberParams struct {
	CardNumber string    `db:"card_number" json:"card_number"`
	Column2    time.Time `db:"column2" json:"column_2"`
	Column3    time.Time `db:"column3" json:"column_3"`
	Column4    time.Time `db:"column4" json:"column_4"`
	Column5    time.Time `db:"column5" json:"column_5"`
}

type GetMonthTransactionStatusSuccessCardNumberRow struct {
	Year         string `db:"year" json:"year"`
	Month        string `db:"month" json:"month"`
	TotalSuccess int64  `db:"total_success" json:"total_success"`
	TotalAmount  int32  `db:"total_amount" json:"total_amount"`
}

type GetMonthlyAmountsRow struct {
	Month       string `db:"month" json:"month"`
	TotalAmount int32  `db:"total_amount" json:"total_amount"`
}

type GetMonthlyAmountsByCardNumberParams struct {
	CardNumber string    `db:"card_number" json:"card_number"`
	Column2    time.Time `db:"column2" json:"column_2"`
}

type GetMonthlyAmountsByCardNumberRow struct {
	Month       string `db:"month" json:"month"`
	TotalAmount int32  `db:"total_amount" json:"total_amount"`
}

type GetMonthlyPaymentMethodsRow struct {
	Month             string `db:"month" json:"month"`
	PaymentMethod     string `db:"payment_method" json:"payment_method"`
	TotalTransactions int32  `db:"total_transactions" json:"total_transactions"`
	TotalAmount       int32  `db:"total_amount" json:"total_amount"`
}

type GetMonthlyPaymentMethodsByCardNumberParams struct {
	CardNumber string    `db:"card_number" json:"card_number"`
	Column2    time.Time `db:"column2" json:"column_2"`
}

type GetMonthlyPaymentMethodsByCardNumberRow struct {
	Month             string `db:"month" json:"month"`
	PaymentMethod     string `db:"payment_method" json:"payment_method"`
	TotalTransactions int32  `db:"total_transactions" json:"total_transactions"`
	TotalAmount       int32  `db:"total_amount" json:"total_amount"`
}

type GetTransactionByCardNumberParams struct {
	CardNumber string `db:"card_number" json:"card_number"`
	Column2    string `db:"column2" json:"column_2"`
	Limit      int32  `db:"limit" json:"limit"`
	Offset     int32  `db:"offset" json:"offset"`
}

type GetTransactionByCardNumberRow struct {
	TransactionID   int32            `db:"transaction_id" json:"transaction_id"`
	TransactionNo   pgtype.UUID      `db:"transaction_no" json:"transaction_no"`
	CardNumber      string           `db:"card_number" json:"card_number"`
	Amount          int32            `db:"amount" json:"amount"`
	PaymentMethod   string           `db:"payment_method" json:"payment_method"`
	MerchantID      int32            `db:"merchant_id" json:"merchant_id"`
	TransactionTime time.Time        `db:"transaction_time" json:"transaction_time"`
	Status          string           `db:"status" json:"status"`
	CreatedAt       pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt       pgtype.Timestamp `db:"updated_at" json:"updated_at"`
	DeletedAt       pgtype.Timestamp `db:"deleted_at" json:"deleted_at"`
	FraudScore      int32            `db:"fraud_score" json:"fraud_score"`
	IdempotencyKey  string           `db:"idempotency_key" json:"idempotency_key"`
	TotalCount      int64            `db:"total_count" json:"total_count"`
}

type GetTransactionByIDRow struct {
	TransactionID   int32            `db:"transaction_id" json:"transaction_id"`
	TransactionNo   pgtype.UUID      `db:"transaction_no" json:"transaction_no"`
	CardNumber      string           `db:"card_number" json:"card_number"`
	Amount          int32            `db:"amount" json:"amount"`
	PaymentMethod   string           `db:"payment_method" json:"payment_method"`
	MerchantID      int32            `db:"merchant_id" json:"merchant_id"`
	TransactionTime time.Time        `db:"transaction_time" json:"transaction_time"`
	Status          string           `db:"status" json:"status"`
	CreatedAt       pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt       pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type GetTransactionByIDWithFraudScoreRow struct {
	TransactionID   int32            `db:"transaction_id" json:"transaction_id"`
	TransactionNo   pgtype.UUID      `db:"transaction_no" json:"transaction_no"`
	CardNumber      string           `db:"card_number" json:"card_number"`
	Amount          int32            `db:"amount" json:"amount"`
	PaymentMethod   string           `db:"payment_method" json:"payment_method"`
	MerchantID      int32            `db:"merchant_id" json:"merchant_id"`
	TransactionTime time.Time        `db:"transaction_time" json:"transaction_time"`
	Status          string           `db:"status" json:"status"`
	FraudScore      int32            `db:"fraud_score" json:"fraud_score"`
	CreatedAt       pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt       pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type GetTransactionByIdempotencyKeyRow struct {
	TransactionID   int32            `db:"transaction_id" json:"transaction_id"`
	TransactionNo   pgtype.UUID      `db:"transaction_no" json:"transaction_no"`
	CardNumber      string           `db:"card_number" json:"card_number"`
	Amount          int32            `db:"amount" json:"amount"`
	PaymentMethod   string           `db:"payment_method" json:"payment_method"`
	MerchantID      int32            `db:"merchant_id" json:"merchant_id"`
	TransactionTime time.Time        `db:"transaction_time" json:"transaction_time"`
	Status          string           `db:"status" json:"status"`
	CreatedAt       pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt       pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type GetTransactionVelocityParams struct {
	CardNumber string           `db:"card_number" json:"card_number"`
	CreatedAt  pgtype.Timestamp `db:"created_at" json:"created_at"`
}

type GetTransactionsParams struct {
	Column1 string `db:"column1" json:"column_1"`
	Limit   int32  `db:"limit" json:"limit"`
	Offset  int32  `db:"offset" json:"offset"`
}

type GetTransactionsRow struct {
	TransactionID   int32            `db:"transaction_id" json:"transaction_id"`
	TransactionNo   pgtype.UUID      `db:"transaction_no" json:"transaction_no"`
	CardNumber      string           `db:"card_number" json:"card_number"`
	Amount          int32            `db:"amount" json:"amount"`
	PaymentMethod   string           `db:"payment_method" json:"payment_method"`
	MerchantID      int32            `db:"merchant_id" json:"merchant_id"`
	TransactionTime time.Time        `db:"transaction_time" json:"transaction_time"`
	Status          string           `db:"status" json:"status"`
	CreatedAt       pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt       pgtype.Timestamp `db:"updated_at" json:"updated_at"`
	TotalCount      int64            `db:"total_count" json:"total_count"`
}

type GetTransactionsByCardNumberParams struct {
	CardNumber string `db:"card_number" json:"card_number"`
	Column2    string `db:"column2" json:"column_2"`
	Limit      int32  `db:"limit" json:"limit"`
	Offset     int32  `db:"offset" json:"offset"`
}

type GetTransactionsByCardNumberRow struct {
	TransactionID   int32            `db:"transaction_id" json:"transaction_id"`
	TransactionNo   pgtype.UUID      `db:"transaction_no" json:"transaction_no"`
	CardNumber      string           `db:"card_number" json:"card_number"`
	Amount          int32            `db:"amount" json:"amount"`
	PaymentMethod   string           `db:"payment_method" json:"payment_method"`
	MerchantID      int32            `db:"merchant_id" json:"merchant_id"`
	TransactionTime time.Time        `db:"transaction_time" json:"transaction_time"`
	Status          string           `db:"status" json:"status"`
	CreatedAt       pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt       pgtype.Timestamp `db:"updated_at" json:"updated_at"`
	TotalCount      int64            `db:"total_count" json:"total_count"`
}

type GetTransactionsByMerchantIDRow struct {
	TransactionID   int32            `db:"transaction_id" json:"transaction_id"`
	TransactionNo   pgtype.UUID      `db:"transaction_no" json:"transaction_no"`
	CardNumber      string           `db:"card_number" json:"card_number"`
	Amount          int32            `db:"amount" json:"amount"`
	PaymentMethod   string           `db:"payment_method" json:"payment_method"`
	MerchantID      int32            `db:"merchant_id" json:"merchant_id"`
	TransactionTime time.Time        `db:"transaction_time" json:"transaction_time"`
	Status          string           `db:"status" json:"status"`
	CreatedAt       pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt       pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type GetTrashedTransactionsParams struct {
	Column1 string `db:"column1" json:"column_1"`
	Limit   int32  `db:"limit" json:"limit"`
	Offset  int32  `db:"offset" json:"offset"`
}

type GetTrashedTransactionsRow struct {
	TransactionID   int32            `db:"transaction_id" json:"transaction_id"`
	TransactionNo   pgtype.UUID      `db:"transaction_no" json:"transaction_no"`
	CardNumber      string           `db:"card_number" json:"card_number"`
	Amount          int32            `db:"amount" json:"amount"`
	PaymentMethod   string           `db:"payment_method" json:"payment_method"`
	MerchantID      int32            `db:"merchant_id" json:"merchant_id"`
	TransactionTime time.Time        `db:"transaction_time" json:"transaction_time"`
	Status          string           `db:"status" json:"status"`
	CreatedAt       pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt       pgtype.Timestamp `db:"updated_at" json:"updated_at"`
	DeletedAt       pgtype.Timestamp `db:"deleted_at" json:"deleted_at"`
	TotalCount      int64            `db:"total_count" json:"total_count"`
}

type GetYearlyAmountsRow struct {
	Year        pgtype.Numeric `db:"year" json:"year"`
	TotalAmount int64          `db:"total_amount" json:"total_amount"`
}

type GetYearlyAmountsByCardNumberParams struct {
	CardNumber string      `db:"card_number" json:"card_number"`
	Column2    interface{} `db:"column2" json:"column_2"`
}

type GetYearlyAmountsByCardNumberRow struct {
	Year        pgtype.Numeric `db:"year" json:"year"`
	TotalAmount int64          `db:"total_amount" json:"total_amount"`
}

type GetYearlyPaymentMethodsRow struct {
	Year              pgtype.Numeric `db:"year" json:"year"`
	PaymentMethod     string         `db:"payment_method" json:"payment_method"`
	TotalTransactions int64          `db:"total_transactions" json:"total_transactions"`
	TotalAmount       int64          `db:"total_amount" json:"total_amount"`
}

type GetYearlyPaymentMethodsByCardNumberParams struct {
	CardNumber string      `db:"card_number" json:"card_number"`
	Column2    interface{} `db:"column2" json:"column_2"`
}

type GetYearlyPaymentMethodsByCardNumberRow struct {
	Year              pgtype.Numeric `db:"year" json:"year"`
	PaymentMethod     string         `db:"payment_method" json:"payment_method"`
	TotalTransactions int64          `db:"total_transactions" json:"total_transactions"`
	TotalAmount       int64          `db:"total_amount" json:"total_amount"`
}

type GetYearlyTransactionStatusFailedRow struct {
	Year        string `db:"year" json:"year"`
	TotalFailed int32  `db:"total_failed" json:"total_failed"`
	TotalAmount int32  `db:"total_amount" json:"total_amount"`
}

type GetYearlyTransactionStatusFailedCardNumberParams struct {
	CardNumber string `db:"card_number" json:"card_number"`
	Column2    int32  `db:"column2" json:"column_2"`
}

type GetYearlyTransactionStatusFailedCardNumberRow struct {
	Year        string `db:"year" json:"year"`
	TotalFailed int32  `db:"total_failed" json:"total_failed"`
	TotalAmount int32  `db:"total_amount" json:"total_amount"`
}

type GetYearlyTransactionStatusSuccessRow struct {
	Year         string `db:"year" json:"year"`
	TotalSuccess int32  `db:"total_success" json:"total_success"`
	TotalAmount  int32  `db:"total_amount" json:"total_amount"`
}

type GetYearlyTransactionStatusSuccessCardNumberParams struct {
	CardNumber string `db:"card_number" json:"card_number"`
	Column2    int32  `db:"column2" json:"column_2"`
}

type GetYearlyTransactionStatusSuccessCardNumberRow struct {
	Year         string `db:"year" json:"year"`
	TotalSuccess int32  `db:"total_success" json:"total_success"`
	TotalAmount  int32  `db:"total_amount" json:"total_amount"`
}

type RefundTransactionAtomicParams struct {
	TransactionID      int32  `db:"transaction_id" json:"transaction_id"`
	MerchantCardNumber string `db:"merchant_card_number" json:"merchant_card_number"`
}

type RefundTransactionAtomicRow struct {
	TransactionID   int32            `db:"transaction_id" json:"transaction_id"`
	TransactionNo   pgtype.UUID      `db:"transaction_no" json:"transaction_no"`
	CardNumber      string           `db:"card_number" json:"card_number"`
	Amount          int32            `db:"amount" json:"amount"`
	PaymentMethod   string           `db:"payment_method" json:"payment_method"`
	MerchantID      int32            `db:"merchant_id" json:"merchant_id"`
	TransactionTime time.Time        `db:"transaction_time" json:"transaction_time"`
	Status          string           `db:"status" json:"status"`
	CreatedAt       pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt       pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type RestoreTransactionRow struct {
	TransactionID   int32            `db:"transaction_id" json:"transaction_id"`
	TransactionNo   pgtype.UUID      `db:"transaction_no" json:"transaction_no"`
	CardNumber      string           `db:"card_number" json:"card_number"`
	Amount          int32            `db:"amount" json:"amount"`
	PaymentMethod   string           `db:"payment_method" json:"payment_method"`
	MerchantID      int32            `db:"merchant_id" json:"merchant_id"`
	TransactionTime time.Time        `db:"transaction_time" json:"transaction_time"`
	Status          string           `db:"status" json:"status"`
	CreatedAt       pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt       pgtype.Timestamp `db:"updated_at" json:"updated_at"`
	DeletedAt       pgtype.Timestamp `db:"deleted_at" json:"deleted_at"`
}

type TrashTransactionRow struct {
	TransactionID   int32            `db:"transaction_id" json:"transaction_id"`
	TransactionNo   pgtype.UUID      `db:"transaction_no" json:"transaction_no"`
	CardNumber      string           `db:"card_number" json:"card_number"`
	Amount          int32            `db:"amount" json:"amount"`
	PaymentMethod   string           `db:"payment_method" json:"payment_method"`
	MerchantID      int32            `db:"merchant_id" json:"merchant_id"`
	TransactionTime time.Time        `db:"transaction_time" json:"transaction_time"`
	Status          string           `db:"status" json:"status"`
	CreatedAt       pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt       pgtype.Timestamp `db:"updated_at" json:"updated_at"`
	DeletedAt       pgtype.Timestamp `db:"deleted_at" json:"deleted_at"`
}

type UpdateTransactionParams struct {
	TransactionID   int32     `db:"transaction_id" json:"transaction_id"`
	CardNumber      string    `db:"card_number" json:"card_number"`
	Amount          int32     `db:"amount" json:"amount"`
	PaymentMethod   string    `db:"payment_method" json:"payment_method"`
	MerchantID      int32     `db:"merchant_id" json:"merchant_id"`
	TransactionTime time.Time `db:"transaction_time" json:"transaction_time"`
}

type UpdateTransactionRow struct {
	TransactionID   int32            `db:"transaction_id" json:"transaction_id"`
	TransactionNo   pgtype.UUID      `db:"transaction_no" json:"transaction_no"`
	CardNumber      string           `db:"card_number" json:"card_number"`
	Amount          int32            `db:"amount" json:"amount"`
	PaymentMethod   string           `db:"payment_method" json:"payment_method"`
	MerchantID      int32            `db:"merchant_id" json:"merchant_id"`
	TransactionTime time.Time        `db:"transaction_time" json:"transaction_time"`
	Status          string           `db:"status" json:"status"`
	CreatedAt       pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt       pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type UpdateTransactionAtomicParams struct {
	TransactionID      int32       `db:"transaction_id" json:"transaction_id"`
	MerchantCardNumber string      `db:"merchant_card_number" json:"merchant_card_number"`
	Amount             interface{} `db:"amount" json:"amount"`
	CardNumber         string      `db:"card_number" json:"card_number"`
	PaymentMethod      string      `db:"payment_method" json:"payment_method"`
	MerchantID         int32       `db:"merchant_id" json:"merchant_id"`
	TransactionTime    time.Time   `db:"transaction_time" json:"transaction_time"`
}

type UpdateTransactionAtomicRow struct {
	TransactionID   int32            `db:"transaction_id" json:"transaction_id"`
	TransactionNo   pgtype.UUID      `db:"transaction_no" json:"transaction_no"`
	CardNumber      string           `db:"card_number" json:"card_number"`
	Amount          int32            `db:"amount" json:"amount"`
	PaymentMethod   string           `db:"payment_method" json:"payment_method"`
	MerchantID      int32            `db:"merchant_id" json:"merchant_id"`
	TransactionTime time.Time        `db:"transaction_time" json:"transaction_time"`
	Status          string           `db:"status" json:"status"`
	CreatedAt       pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt       pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type UpdateTransactionFraudScoreParams struct {
	TransactionID int32 `db:"transaction_id" json:"transaction_id"`
	FraudScore    int32 `db:"fraud_score" json:"fraud_score"`
}

type UpdateTransactionFraudScoreRow struct {
	TransactionID   int32            `db:"transaction_id" json:"transaction_id"`
	TransactionNo   pgtype.UUID      `db:"transaction_no" json:"transaction_no"`
	CardNumber      string           `db:"card_number" json:"card_number"`
	Amount          int32            `db:"amount" json:"amount"`
	PaymentMethod   string           `db:"payment_method" json:"payment_method"`
	MerchantID      int32            `db:"merchant_id" json:"merchant_id"`
	TransactionTime time.Time        `db:"transaction_time" json:"transaction_time"`
	Status          string           `db:"status" json:"status"`
	FraudScore      int32            `db:"fraud_score" json:"fraud_score"`
	CreatedAt       pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt       pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type UpdateTransactionStatusParams struct {
	TransactionID int32  `db:"transaction_id" json:"transaction_id"`
	Status        string `db:"status" json:"status"`
}

type UpdateTransactionStatusRow struct {
	TransactionID   int32            `db:"transaction_id" json:"transaction_id"`
	TransactionNo   pgtype.UUID      `db:"transaction_no" json:"transaction_no"`
	CardNumber      string           `db:"card_number" json:"card_number"`
	Amount          int32            `db:"amount" json:"amount"`
	PaymentMethod   string           `db:"payment_method" json:"payment_method"`
	MerchantID      int32            `db:"merchant_id" json:"merchant_id"`
	TransactionTime time.Time        `db:"transaction_time" json:"transaction_time"`
	Status          string           `db:"status" json:"status"`
	CreatedAt       pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt       pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type UpdateTransactionStatusDirectParams struct {
	TransactionID int32  `db:"transaction_id" json:"transaction_id"`
	Status        string `db:"status" json:"status"`
}

type UpdateTransactionStatusDirectRow struct {
	TransactionID   int32            `db:"transaction_id" json:"transaction_id"`
	TransactionNo   pgtype.UUID      `db:"transaction_no" json:"transaction_no"`
	CardNumber      string           `db:"card_number" json:"card_number"`
	Amount          int32            `db:"amount" json:"amount"`
	PaymentMethod   string           `db:"payment_method" json:"payment_method"`
	MerchantID      int32            `db:"merchant_id" json:"merchant_id"`
	TransactionTime time.Time        `db:"transaction_time" json:"transaction_time"`
	Status          string           `db:"status" json:"status"`
	FraudScore      int32            `db:"fraud_score" json:"fraud_score"`
	CreatedAt       pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt       pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type VoidTransactionAtomicRow struct {
	TransactionID   int32            `db:"transaction_id" json:"transaction_id"`
	TransactionNo   pgtype.UUID      `db:"transaction_no" json:"transaction_no"`
	CardNumber      string           `db:"card_number" json:"card_number"`
	Amount          int32            `db:"amount" json:"amount"`
	PaymentMethod   string           `db:"payment_method" json:"payment_method"`
	MerchantID      int32            `db:"merchant_id" json:"merchant_id"`
	TransactionTime time.Time        `db:"transaction_time" json:"transaction_time"`
	Status          string           `db:"status" json:"status"`
	CreatedAt       pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt       pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}
