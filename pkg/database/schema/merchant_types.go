// Code migrated from sqlc-generated sources; domain types kept for repository layer.
// source: merchant.sql

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
	"time"
)

type CreateMerchantParams struct {
	Name   string `db:"name" json:"name"`
	ApiKey string `db:"api_key" json:"api_key"`
	UserID int32  `db:"user_id" json:"user_id"`
	Status string `db:"status" json:"status"`
}

type CreateMerchantRow struct {
	MerchantID int32            `db:"merchant_id" json:"merchant_id"`
	MerchantNo pgtype.UUID      `db:"merchant_no" json:"merchant_no"`
	Name       string           `db:"name" json:"name"`
	ApiKey     string           `db:"api_key" json:"api_key"`
	UserID     int32            `db:"user_id" json:"user_id"`
	Status     string           `db:"status" json:"status"`
	CreatedAt  pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt  pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type FindAllTransactionsParams struct {
	Column1 string `db:"column1" json:"column_1"`
	Limit   int32  `db:"limit" json:"limit"`
	Offset  int32  `db:"offset" json:"offset"`
}

type FindAllTransactionsRow struct {
	TransactionID   int32            `db:"transaction_id" json:"transaction_id"`
	CardNumber      string           `db:"card_number" json:"card_number"`
	Amount          int32            `db:"amount" json:"amount"`
	PaymentMethod   string           `db:"payment_method" json:"payment_method"`
	MerchantID      int32            `db:"merchant_id" json:"merchant_id"`
	MerchantName    string           `db:"merchant_name" json:"merchant_name"`
	TransactionTime time.Time        `db:"transaction_time" json:"transaction_time"`
	CreatedAt       pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt       pgtype.Timestamp `db:"updated_at" json:"updated_at"`
	DeletedAt       pgtype.Timestamp `db:"deleted_at" json:"deleted_at"`
	TotalCount      int64            `db:"total_count" json:"total_count"`
}

type FindAllTransactionsByApikeyParams struct {
	ApiKey  string `db:"api_key" json:"api_key"`
	Column2 string `db:"column2" json:"column_2"`
	Limit   int32  `db:"limit" json:"limit"`
	Offset  int32  `db:"offset" json:"offset"`
}

type FindAllTransactionsByApikeyRow struct {
	TransactionID   int32            `db:"transaction_id" json:"transaction_id"`
	CardNumber      string           `db:"card_number" json:"card_number"`
	Amount          int32            `db:"amount" json:"amount"`
	PaymentMethod   string           `db:"payment_method" json:"payment_method"`
	MerchantID      int32            `db:"merchant_id" json:"merchant_id"`
	MerchantName    string           `db:"merchant_name" json:"merchant_name"`
	TransactionTime time.Time        `db:"transaction_time" json:"transaction_time"`
	CreatedAt       pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt       pgtype.Timestamp `db:"updated_at" json:"updated_at"`
	DeletedAt       pgtype.Timestamp `db:"deleted_at" json:"deleted_at"`
	TotalCount      int64            `db:"total_count" json:"total_count"`
}

type FindAllTransactionsByMerchantParams struct {
	MerchantID int32  `db:"merchant_id" json:"merchant_id"`
	Column2    string `db:"column2" json:"column_2"`
	Limit      int32  `db:"limit" json:"limit"`
	Offset     int32  `db:"offset" json:"offset"`
}

type FindAllTransactionsByMerchantRow struct {
	TransactionID   int32            `db:"transaction_id" json:"transaction_id"`
	CardNumber      string           `db:"card_number" json:"card_number"`
	Amount          int32            `db:"amount" json:"amount"`
	PaymentMethod   string           `db:"payment_method" json:"payment_method"`
	MerchantID      int32            `db:"merchant_id" json:"merchant_id"`
	MerchantName    string           `db:"merchant_name" json:"merchant_name"`
	TransactionTime time.Time        `db:"transaction_time" json:"transaction_time"`
	CreatedAt       pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt       pgtype.Timestamp `db:"updated_at" json:"updated_at"`
	DeletedAt       pgtype.Timestamp `db:"deleted_at" json:"deleted_at"`
	TotalCount      int64            `db:"total_count" json:"total_count"`
}

type GetActiveMerchantsParams struct {
	Column1 string `db:"column1" json:"column_1"`
	Limit   int32  `db:"limit" json:"limit"`
	Offset  int32  `db:"offset" json:"offset"`
}

type GetActiveMerchantsRow struct {
	MerchantID int32            `db:"merchant_id" json:"merchant_id"`
	MerchantNo pgtype.UUID      `db:"merchant_no" json:"merchant_no"`
	Name       string           `db:"name" json:"name"`
	ApiKey     string           `db:"api_key" json:"api_key"`
	UserID     int32            `db:"user_id" json:"user_id"`
	Status     string           `db:"status" json:"status"`
	CreatedAt  pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt  pgtype.Timestamp `db:"updated_at" json:"updated_at"`
	DeletedAt  pgtype.Timestamp `db:"deleted_at" json:"deleted_at"`
	TotalCount int64            `db:"total_count" json:"total_count"`
}

type GetMerchantByApiKeyRow struct {
	MerchantID int32            `db:"merchant_id" json:"merchant_id"`
	MerchantNo pgtype.UUID      `db:"merchant_no" json:"merchant_no"`
	Name       string           `db:"name" json:"name"`
	ApiKey     string           `db:"api_key" json:"api_key"`
	UserID     int32            `db:"user_id" json:"user_id"`
	Status     string           `db:"status" json:"status"`
	CreatedAt  pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt  pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type GetMerchantByIDRow struct {
	MerchantID int32            `db:"merchant_id" json:"merchant_id"`
	MerchantNo pgtype.UUID      `db:"merchant_no" json:"merchant_no"`
	Name       string           `db:"name" json:"name"`
	ApiKey     string           `db:"api_key" json:"api_key"`
	UserID     int32            `db:"user_id" json:"user_id"`
	Status     string           `db:"status" json:"status"`
	CreatedAt  pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt  pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type GetMerchantByNameRow struct {
	MerchantID int32            `db:"merchant_id" json:"merchant_id"`
	MerchantNo pgtype.UUID      `db:"merchant_no" json:"merchant_no"`
	Name       string           `db:"name" json:"name"`
	ApiKey     string           `db:"api_key" json:"api_key"`
	UserID     int32            `db:"user_id" json:"user_id"`
	Status     string           `db:"status" json:"status"`
	CreatedAt  pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt  pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type GetMerchantsParams struct {
	Column1 string `db:"column1" json:"column_1"`
	Limit   int32  `db:"limit" json:"limit"`
	Offset  int32  `db:"offset" json:"offset"`
}

type GetMerchantsRow struct {
	MerchantID int32            `db:"merchant_id" json:"merchant_id"`
	MerchantNo pgtype.UUID      `db:"merchant_no" json:"merchant_no"`
	Name       string           `db:"name" json:"name"`
	ApiKey     string           `db:"api_key" json:"api_key"`
	UserID     int32            `db:"user_id" json:"user_id"`
	Status     string           `db:"status" json:"status"`
	CreatedAt  pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt  pgtype.Timestamp `db:"updated_at" json:"updated_at"`
	TotalCount int64            `db:"total_count" json:"total_count"`
}

type GetMerchantsByUserIDRow struct {
	MerchantID int32            `db:"merchant_id" json:"merchant_id"`
	MerchantNo pgtype.UUID      `db:"merchant_no" json:"merchant_no"`
	Name       string           `db:"name" json:"name"`
	ApiKey     string           `db:"api_key" json:"api_key"`
	UserID     int32            `db:"user_id" json:"user_id"`
	Status     string           `db:"status" json:"status"`
	CreatedAt  pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt  pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type GetMonthlyAmountByApikeyParams struct {
	Column1 time.Time `db:"column1" json:"column_1"`
	ApiKey  string    `db:"api_key" json:"api_key"`
}

type GetMonthlyAmountByApikeyRow struct {
	Month       string `db:"month" json:"month"`
	TotalAmount int32  `db:"total_amount" json:"total_amount"`
}

type GetMonthlyAmountByMerchantsParams struct {
	Column1    time.Time `db:"column1" json:"column_1"`
	MerchantID int32     `db:"merchant_id" json:"merchant_id"`
}

type GetMonthlyAmountByMerchantsRow struct {
	Month       string `db:"month" json:"month"`
	TotalAmount int32  `db:"total_amount" json:"total_amount"`
}

type GetMonthlyAmountMerchantRow struct {
	Month       string `db:"month" json:"month"`
	TotalAmount int32  `db:"total_amount" json:"total_amount"`
}

type GetMonthlyPaymentMethodByApikeyParams struct {
	Column1 time.Time `db:"column1" json:"column_1"`
	ApiKey  string    `db:"api_key" json:"api_key"`
}

type GetMonthlyPaymentMethodByApikeyRow struct {
	Month         string `db:"month" json:"month"`
	PaymentMethod string `db:"payment_method" json:"payment_method"`
	TotalAmount   int32  `db:"total_amount" json:"total_amount"`
}

type GetMonthlyPaymentMethodByMerchantsParams struct {
	Column1    time.Time `db:"column1" json:"column_1"`
	MerchantID int32     `db:"merchant_id" json:"merchant_id"`
}

type GetMonthlyPaymentMethodByMerchantsRow struct {
	Month         string `db:"month" json:"month"`
	PaymentMethod string `db:"payment_method" json:"payment_method"`
	TotalAmount   int32  `db:"total_amount" json:"total_amount"`
}

type GetMonthlyPaymentMethodsMerchantRow struct {
	Month         string `db:"month" json:"month"`
	PaymentMethod string `db:"payment_method" json:"payment_method"`
	TotalAmount   int32  `db:"total_amount" json:"total_amount"`
}

type GetMonthlyTotalAmountByApikeyParams struct {
	Column1 time.Time `db:"column1" json:"column_1"`
	ApiKey  string    `db:"api_key" json:"api_key"`
}

type GetMonthlyTotalAmountByApikeyRow struct {
	Year        string `db:"year" json:"year"`
	Month       string `db:"month" json:"month"`
	TotalAmount int32  `db:"total_amount" json:"total_amount"`
}

type GetMonthlyTotalAmountByMerchantParams struct {
	Column1 time.Time `db:"column1" json:"column_1"`
	Column2 int32     `db:"column2" json:"column_2"`
}

type GetMonthlyTotalAmountByMerchantRow struct {
	Year        string `db:"year" json:"year"`
	Month       string `db:"month" json:"month"`
	TotalAmount int32  `db:"total_amount" json:"total_amount"`
}

type GetMonthlyTotalAmountMerchantRow struct {
	Year        string `db:"year" json:"year"`
	Month       string `db:"month" json:"month"`
	TotalAmount int32  `db:"total_amount" json:"total_amount"`
}

type GetTrashedMerchantsParams struct {
	Column1 string `db:"column1" json:"column_1"`
	Limit   int32  `db:"limit" json:"limit"`
	Offset  int32  `db:"offset" json:"offset"`
}

type GetTrashedMerchantsRow struct {
	MerchantID int32            `db:"merchant_id" json:"merchant_id"`
	MerchantNo pgtype.UUID      `db:"merchant_no" json:"merchant_no"`
	Name       string           `db:"name" json:"name"`
	ApiKey     string           `db:"api_key" json:"api_key"`
	UserID     int32            `db:"user_id" json:"user_id"`
	Status     string           `db:"status" json:"status"`
	CreatedAt  pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt  pgtype.Timestamp `db:"updated_at" json:"updated_at"`
	DeletedAt  pgtype.Timestamp `db:"deleted_at" json:"deleted_at"`
	TotalCount int64            `db:"total_count" json:"total_count"`
}

type GetYearlyAmountByApikeyParams struct {
	ApiKey  string      `db:"api_key" json:"api_key"`
	Column2 interface{} `db:"column2" json:"column_2"`
}

type GetYearlyAmountByApikeyRow struct {
	Year        pgtype.Numeric `db:"year" json:"year"`
	TotalAmount int64          `db:"total_amount" json:"total_amount"`
}

type GetYearlyAmountByMerchantsParams struct {
	MerchantID int32       `db:"merchant_id" json:"merchant_id"`
	Column2    interface{} `db:"column2" json:"column_2"`
}

type GetYearlyAmountByMerchantsRow struct {
	Year        pgtype.Numeric `db:"year" json:"year"`
	TotalAmount int64          `db:"total_amount" json:"total_amount"`
}

type GetYearlyAmountMerchantRow struct {
	Year        pgtype.Numeric `db:"year" json:"year"`
	TotalAmount int64          `db:"total_amount" json:"total_amount"`
}

type GetYearlyPaymentMethodByApikeyParams struct {
	ApiKey  string      `db:"api_key" json:"api_key"`
	Column2 interface{} `db:"column2" json:"column_2"`
}

type GetYearlyPaymentMethodByApikeyRow struct {
	Year          pgtype.Numeric `db:"year" json:"year"`
	PaymentMethod string         `db:"payment_method" json:"payment_method"`
	TotalAmount   int64          `db:"total_amount" json:"total_amount"`
}

type GetYearlyPaymentMethodByMerchantsParams struct {
	Column1    interface{} `db:"column1" json:"column_1"`
	MerchantID int32       `db:"merchant_id" json:"merchant_id"`
}

type GetYearlyPaymentMethodByMerchantsRow struct {
	Year          pgtype.Numeric `db:"year" json:"year"`
	PaymentMethod string         `db:"payment_method" json:"payment_method"`
	TotalAmount   int64          `db:"total_amount" json:"total_amount"`
}

type GetYearlyPaymentMethodMerchantRow struct {
	Year          pgtype.Numeric `db:"year" json:"year"`
	PaymentMethod string         `db:"payment_method" json:"payment_method"`
	TotalAmount   int64          `db:"total_amount" json:"total_amount"`
}

type GetYearlyTotalAmountByApikeyParams struct {
	Column1 int32  `db:"column1" json:"column_1"`
	ApiKey  string `db:"api_key" json:"api_key"`
}

type GetYearlyTotalAmountByApikeyRow struct {
	Year        string `db:"year" json:"year"`
	TotalAmount int32  `db:"total_amount" json:"total_amount"`
}

type GetYearlyTotalAmountByMerchantParams struct {
	Column1 int32 `db:"column1" json:"column_1"`
	Column2 int32 `db:"column2" json:"column_2"`
}

type GetYearlyTotalAmountByMerchantRow struct {
	Year        string `db:"year" json:"year"`
	TotalAmount int32  `db:"total_amount" json:"total_amount"`
}

type GetYearlyTotalAmountMerchantRow struct {
	Year        string `db:"year" json:"year"`
	TotalAmount int32  `db:"total_amount" json:"total_amount"`
}

type RestoreMerchantRow struct {
	MerchantID int32            `db:"merchant_id" json:"merchant_id"`
	MerchantNo pgtype.UUID      `db:"merchant_no" json:"merchant_no"`
	Name       string           `db:"name" json:"name"`
	ApiKey     string           `db:"api_key" json:"api_key"`
	UserID     int32            `db:"user_id" json:"user_id"`
	Status     string           `db:"status" json:"status"`
	CreatedAt  pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt  pgtype.Timestamp `db:"updated_at" json:"updated_at"`
	DeletedAt  pgtype.Timestamp `db:"deleted_at" json:"deleted_at"`
}

type TrashMerchantRow struct {
	MerchantID int32            `db:"merchant_id" json:"merchant_id"`
	MerchantNo pgtype.UUID      `db:"merchant_no" json:"merchant_no"`
	Name       string           `db:"name" json:"name"`
	ApiKey     string           `db:"api_key" json:"api_key"`
	UserID     int32            `db:"user_id" json:"user_id"`
	Status     string           `db:"status" json:"status"`
	CreatedAt  pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt  pgtype.Timestamp `db:"updated_at" json:"updated_at"`
	DeletedAt  pgtype.Timestamp `db:"deleted_at" json:"deleted_at"`
}

type UpdateMerchantParams struct {
	MerchantID int32  `db:"merchant_id" json:"merchant_id"`
	Name       string `db:"name" json:"name"`
	UserID     int32  `db:"user_id" json:"user_id"`
	Status     string `db:"status" json:"status"`
}

type UpdateMerchantRow struct {
	MerchantID int32            `db:"merchant_id" json:"merchant_id"`
	MerchantNo pgtype.UUID      `db:"merchant_no" json:"merchant_no"`
	Name       string           `db:"name" json:"name"`
	ApiKey     string           `db:"api_key" json:"api_key"`
	UserID     int32            `db:"user_id" json:"user_id"`
	Status     string           `db:"status" json:"status"`
	CreatedAt  pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt  pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type UpdateMerchantStatusParams struct {
	MerchantID int32  `db:"merchant_id" json:"merchant_id"`
	Status     string `db:"status" json:"status"`
}

type UpdateMerchantStatusRow struct {
	MerchantID int32            `db:"merchant_id" json:"merchant_id"`
	MerchantNo pgtype.UUID      `db:"merchant_no" json:"merchant_no"`
	Name       string           `db:"name" json:"name"`
	ApiKey     string           `db:"api_key" json:"api_key"`
	UserID     int32            `db:"user_id" json:"user_id"`
	Status     string           `db:"status" json:"status"`
	CreatedAt  pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt  pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}
