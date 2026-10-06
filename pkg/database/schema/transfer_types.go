// Code migrated from sqlc-generated sources; domain types kept for repository layer.
// source: transfer.sql

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
	"time"
)

type CreateTransferParams struct {
	TransferFrom   string    `db:"transfer_from" json:"transfer_from"`
	TransferTo     string    `db:"transfer_to" json:"transfer_to"`
	TransferAmount int32     `db:"transfer_amount" json:"transfer_amount"`
	TransferTime   time.Time `db:"transfer_time" json:"transfer_time"`
	Status         string    `db:"status" json:"status"`
}

type CreateTransferRow struct {
	TransferID     int32            `db:"transfer_id" json:"transfer_id"`
	TransferNo     pgtype.UUID      `db:"transfer_no" json:"transfer_no"`
	TransferFrom   string           `db:"transfer_from" json:"transfer_from"`
	TransferTo     string           `db:"transfer_to" json:"transfer_to"`
	TransferAmount int32            `db:"transfer_amount" json:"transfer_amount"`
	TransferTime   time.Time        `db:"transfer_time" json:"transfer_time"`
	Status         string           `db:"status" json:"status"`
	CreatedAt      pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt      pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type CreateTransferAtomicParams struct {
	TransferFrom   string `db:"transfer_from" json:"transfer_from"`
	TransferTo     string `db:"transfer_to" json:"transfer_to"`
	TransferAmount int32  `db:"transfer_amount" json:"transfer_amount"`
	IdempotencyKey string `db:"idempotency_key" json:"idempotency_key"`
}

type CreateTransferAtomicRow struct {
	TransferID     int32            `db:"transfer_id" json:"transfer_id"`
	TransferNo     pgtype.UUID      `db:"transfer_no" json:"transfer_no"`
	TransferFrom   string           `db:"transfer_from" json:"transfer_from"`
	TransferTo     string           `db:"transfer_to" json:"transfer_to"`
	TransferAmount int32            `db:"transfer_amount" json:"transfer_amount"`
	TransferTime   time.Time        `db:"transfer_time" json:"transfer_time"`
	Status         string           `db:"status" json:"status"`
	CreatedAt      pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt      pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type FindAllTransfersByCardNumberAsReceiverRow struct {
	TransferID     int32            `db:"transfer_id" json:"transfer_id"`
	TransferFrom   string           `db:"transfer_from" json:"transfer_from"`
	TransferTo     string           `db:"transfer_to" json:"transfer_to"`
	TransferAmount int32            `db:"transfer_amount" json:"transfer_amount"`
	TransferTime   time.Time        `db:"transfer_time" json:"transfer_time"`
	CreatedAt      pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt      pgtype.Timestamp `db:"updated_at" json:"updated_at"`
	DeletedAt      pgtype.Timestamp `db:"deleted_at" json:"deleted_at"`
}

type FindAllTransfersByCardNumberAsSenderRow struct {
	TransferID     int32            `db:"transfer_id" json:"transfer_id"`
	TransferFrom   string           `db:"transfer_from" json:"transfer_from"`
	TransferTo     string           `db:"transfer_to" json:"transfer_to"`
	TransferAmount int32            `db:"transfer_amount" json:"transfer_amount"`
	TransferTime   time.Time        `db:"transfer_time" json:"transfer_time"`
	CreatedAt      pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt      pgtype.Timestamp `db:"updated_at" json:"updated_at"`
	DeletedAt      pgtype.Timestamp `db:"deleted_at" json:"deleted_at"`
}

type GetActiveTransfersParams struct {
	Column1 string `db:"column1" json:"column_1"`
	Limit   int32  `db:"limit" json:"limit"`
	Offset  int32  `db:"offset" json:"offset"`
}

type GetActiveTransfersRow struct {
	TransferID     int32            `db:"transfer_id" json:"transfer_id"`
	TransferNo     pgtype.UUID      `db:"transfer_no" json:"transfer_no"`
	TransferFrom   string           `db:"transfer_from" json:"transfer_from"`
	TransferTo     string           `db:"transfer_to" json:"transfer_to"`
	TransferAmount int32            `db:"transfer_amount" json:"transfer_amount"`
	TransferTime   time.Time        `db:"transfer_time" json:"transfer_time"`
	Status         string           `db:"status" json:"status"`
	CreatedAt      pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt      pgtype.Timestamp `db:"updated_at" json:"updated_at"`
	DeletedAt      pgtype.Timestamp `db:"deleted_at" json:"deleted_at"`
	TotalCount     int64            `db:"total_count" json:"total_count"`
}

type GetMonthTransferStatusFailedParams struct {
	Column1 time.Time `db:"column1" json:"column_1"`
	Column2 time.Time `db:"column2" json:"column_2"`
	Column3 time.Time `db:"column3" json:"column_3"`
	Column4 time.Time `db:"column4" json:"column_4"`
}

type GetMonthTransferStatusFailedRow struct {
	Year        string `db:"year" json:"year"`
	Month       string `db:"month" json:"month"`
	TotalFailed int64  `db:"total_failed" json:"total_failed"`
	TotalAmount int32  `db:"total_amount" json:"total_amount"`
}

type GetMonthTransferStatusFailedCardNumberParams struct {
	TransferFrom string    `db:"transfer_from" json:"transfer_from"`
	Column2      time.Time `db:"column2" json:"column_2"`
	Column3      time.Time `db:"column3" json:"column_3"`
	Column4      time.Time `db:"column4" json:"column_4"`
	Column5      time.Time `db:"column5" json:"column_5"`
}

type GetMonthTransferStatusFailedCardNumberRow struct {
	Year        string `db:"year" json:"year"`
	Month       string `db:"month" json:"month"`
	TotalFailed int64  `db:"total_failed" json:"total_failed"`
	TotalAmount int32  `db:"total_amount" json:"total_amount"`
}

type GetMonthTransferStatusSuccessParams struct {
	Column1 time.Time `db:"column1" json:"column_1"`
	Column2 time.Time `db:"column2" json:"column_2"`
	Column3 time.Time `db:"column3" json:"column_3"`
	Column4 time.Time `db:"column4" json:"column_4"`
}

type GetMonthTransferStatusSuccessRow struct {
	Year         string `db:"year" json:"year"`
	Month        string `db:"month" json:"month"`
	TotalSuccess int64  `db:"total_success" json:"total_success"`
	TotalAmount  int32  `db:"total_amount" json:"total_amount"`
}

type GetMonthTransferStatusSuccessCardNumberParams struct {
	TransferFrom string    `db:"transfer_from" json:"transfer_from"`
	Column2      time.Time `db:"column2" json:"column_2"`
	Column3      time.Time `db:"column3" json:"column_3"`
	Column4      time.Time `db:"column4" json:"column_4"`
	Column5      time.Time `db:"column5" json:"column_5"`
}

type GetMonthTransferStatusSuccessCardNumberRow struct {
	Year         string `db:"year" json:"year"`
	Month        string `db:"month" json:"month"`
	TotalSuccess int64  `db:"total_success" json:"total_success"`
	TotalAmount  int32  `db:"total_amount" json:"total_amount"`
}

type GetMonthlyTransferAmountsRow struct {
	Month               string `db:"month" json:"month"`
	TotalTransferAmount int32  `db:"total_transfer_amount" json:"total_transfer_amount"`
}

type GetMonthlyTransferAmountsByReceiverCardNumberParams struct {
	TransferTo string    `db:"transfer_to" json:"transfer_to"`
	Column2    time.Time `db:"column2" json:"column_2"`
}

type GetMonthlyTransferAmountsByReceiverCardNumberRow struct {
	Month               string `db:"month" json:"month"`
	TotalTransferAmount int32  `db:"total_transfer_amount" json:"total_transfer_amount"`
}

type GetMonthlyTransferAmountsBySenderCardNumberParams struct {
	TransferFrom string    `db:"transfer_from" json:"transfer_from"`
	Column2      time.Time `db:"column2" json:"column_2"`
}

type GetMonthlyTransferAmountsBySenderCardNumberRow struct {
	Month               string `db:"month" json:"month"`
	TotalTransferAmount int32  `db:"total_transfer_amount" json:"total_transfer_amount"`
}

type GetTransferByIDRow struct {
	TransferID     int32            `db:"transfer_id" json:"transfer_id"`
	TransferNo     pgtype.UUID      `db:"transfer_no" json:"transfer_no"`
	TransferFrom   string           `db:"transfer_from" json:"transfer_from"`
	TransferTo     string           `db:"transfer_to" json:"transfer_to"`
	TransferAmount int32            `db:"transfer_amount" json:"transfer_amount"`
	TransferTime   time.Time        `db:"transfer_time" json:"transfer_time"`
	Status         string           `db:"status" json:"status"`
	CreatedAt      pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt      pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type GetTransferByIdempotencyKeyRow struct {
	TransferID     int32            `db:"transfer_id" json:"transfer_id"`
	TransferNo     pgtype.UUID      `db:"transfer_no" json:"transfer_no"`
	TransferFrom   string           `db:"transfer_from" json:"transfer_from"`
	TransferTo     string           `db:"transfer_to" json:"transfer_to"`
	TransferAmount int32            `db:"transfer_amount" json:"transfer_amount"`
	TransferTime   time.Time        `db:"transfer_time" json:"transfer_time"`
	Status         string           `db:"status" json:"status"`
	CreatedAt      pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt      pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type GetTransfersParams struct {
	Column1 string `db:"column1" json:"column_1"`
	Limit   int32  `db:"limit" json:"limit"`
	Offset  int32  `db:"offset" json:"offset"`
}

type GetTransfersRow struct {
	TransferID     int32            `db:"transfer_id" json:"transfer_id"`
	TransferNo     pgtype.UUID      `db:"transfer_no" json:"transfer_no"`
	TransferFrom   string           `db:"transfer_from" json:"transfer_from"`
	TransferTo     string           `db:"transfer_to" json:"transfer_to"`
	TransferAmount int32            `db:"transfer_amount" json:"transfer_amount"`
	TransferTime   time.Time        `db:"transfer_time" json:"transfer_time"`
	Status         string           `db:"status" json:"status"`
	CreatedAt      pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt      pgtype.Timestamp `db:"updated_at" json:"updated_at"`
	TotalCount     int64            `db:"total_count" json:"total_count"`
}

type GetTransfersByCardNumberRow struct {
	TransferID     int32            `db:"transfer_id" json:"transfer_id"`
	TransferNo     pgtype.UUID      `db:"transfer_no" json:"transfer_no"`
	TransferFrom   string           `db:"transfer_from" json:"transfer_from"`
	TransferTo     string           `db:"transfer_to" json:"transfer_to"`
	TransferAmount int32            `db:"transfer_amount" json:"transfer_amount"`
	TransferTime   time.Time        `db:"transfer_time" json:"transfer_time"`
	Status         string           `db:"status" json:"status"`
	CreatedAt      pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt      pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type GetTransfersByDestinationCardRow struct {
	TransferID     int32            `db:"transfer_id" json:"transfer_id"`
	TransferNo     pgtype.UUID      `db:"transfer_no" json:"transfer_no"`
	TransferFrom   string           `db:"transfer_from" json:"transfer_from"`
	TransferTo     string           `db:"transfer_to" json:"transfer_to"`
	TransferAmount int32            `db:"transfer_amount" json:"transfer_amount"`
	TransferTime   time.Time        `db:"transfer_time" json:"transfer_time"`
	Status         string           `db:"status" json:"status"`
	CreatedAt      pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt      pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type GetTransfersBySourceCardRow struct {
	TransferID     int32            `db:"transfer_id" json:"transfer_id"`
	TransferNo     pgtype.UUID      `db:"transfer_no" json:"transfer_no"`
	TransferFrom   string           `db:"transfer_from" json:"transfer_from"`
	TransferTo     string           `db:"transfer_to" json:"transfer_to"`
	TransferAmount int32            `db:"transfer_amount" json:"transfer_amount"`
	TransferTime   time.Time        `db:"transfer_time" json:"transfer_time"`
	Status         string           `db:"status" json:"status"`
	CreatedAt      pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt      pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type GetTrashedTransfersParams struct {
	Column1 string `db:"column1" json:"column_1"`
	Limit   int32  `db:"limit" json:"limit"`
	Offset  int32  `db:"offset" json:"offset"`
}

type GetTrashedTransfersRow struct {
	TransferID     int32            `db:"transfer_id" json:"transfer_id"`
	TransferNo     pgtype.UUID      `db:"transfer_no" json:"transfer_no"`
	TransferFrom   string           `db:"transfer_from" json:"transfer_from"`
	TransferTo     string           `db:"transfer_to" json:"transfer_to"`
	TransferAmount int32            `db:"transfer_amount" json:"transfer_amount"`
	TransferTime   time.Time        `db:"transfer_time" json:"transfer_time"`
	Status         string           `db:"status" json:"status"`
	CreatedAt      pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt      pgtype.Timestamp `db:"updated_at" json:"updated_at"`
	DeletedAt      pgtype.Timestamp `db:"deleted_at" json:"deleted_at"`
	TotalCount     int64            `db:"total_count" json:"total_count"`
}

type GetYearlyTransferAmountsRow struct {
	Year                pgtype.Numeric `db:"year" json:"year"`
	TotalTransferAmount int64          `db:"total_transfer_amount" json:"total_transfer_amount"`
}

type GetYearlyTransferAmountsByReceiverCardNumberParams struct {
	TransferTo string      `db:"transfer_to" json:"transfer_to"`
	Column2    interface{} `db:"column2" json:"column_2"`
}

type GetYearlyTransferAmountsByReceiverCardNumberRow struct {
	Year                pgtype.Numeric `db:"year" json:"year"`
	TotalTransferAmount int64          `db:"total_transfer_amount" json:"total_transfer_amount"`
}

type GetYearlyTransferAmountsBySenderCardNumberParams struct {
	TransferFrom string      `db:"transfer_from" json:"transfer_from"`
	Column2      interface{} `db:"column2" json:"column_2"`
}

type GetYearlyTransferAmountsBySenderCardNumberRow struct {
	Year                pgtype.Numeric `db:"year" json:"year"`
	TotalTransferAmount int64          `db:"total_transfer_amount" json:"total_transfer_amount"`
}

type GetYearlyTransferStatusFailedRow struct {
	Year        string `db:"year" json:"year"`
	TotalFailed int32  `db:"total_failed" json:"total_failed"`
	TotalAmount int32  `db:"total_amount" json:"total_amount"`
}

type GetYearlyTransferStatusFailedCardNumberParams struct {
	TransferFrom string `db:"transfer_from" json:"transfer_from"`
	Column2      int32  `db:"column2" json:"column_2"`
}

type GetYearlyTransferStatusFailedCardNumberRow struct {
	Year        string `db:"year" json:"year"`
	TotalFailed int32  `db:"total_failed" json:"total_failed"`
	TotalAmount int32  `db:"total_amount" json:"total_amount"`
}

type GetYearlyTransferStatusSuccessRow struct {
	Year         string `db:"year" json:"year"`
	TotalSuccess int32  `db:"total_success" json:"total_success"`
	TotalAmount  int32  `db:"total_amount" json:"total_amount"`
}

type GetYearlyTransferStatusSuccessCardNumberParams struct {
	TransferFrom string `db:"transfer_from" json:"transfer_from"`
	Column2      int32  `db:"column2" json:"column_2"`
}

type GetYearlyTransferStatusSuccessCardNumberRow struct {
	Year         string `db:"year" json:"year"`
	TotalSuccess int32  `db:"total_success" json:"total_success"`
	TotalAmount  int32  `db:"total_amount" json:"total_amount"`
}

type RestoreTransferRow struct {
	TransferID     int32            `db:"transfer_id" json:"transfer_id"`
	TransferNo     pgtype.UUID      `db:"transfer_no" json:"transfer_no"`
	TransferFrom   string           `db:"transfer_from" json:"transfer_from"`
	TransferTo     string           `db:"transfer_to" json:"transfer_to"`
	TransferAmount int32            `db:"transfer_amount" json:"transfer_amount"`
	TransferTime   time.Time        `db:"transfer_time" json:"transfer_time"`
	Status         string           `db:"status" json:"status"`
	CreatedAt      pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt      pgtype.Timestamp `db:"updated_at" json:"updated_at"`
	DeletedAt      pgtype.Timestamp `db:"deleted_at" json:"deleted_at"`
}

type TrashTransferRow struct {
	TransferID     int32            `db:"transfer_id" json:"transfer_id"`
	TransferNo     pgtype.UUID      `db:"transfer_no" json:"transfer_no"`
	TransferFrom   string           `db:"transfer_from" json:"transfer_from"`
	TransferTo     string           `db:"transfer_to" json:"transfer_to"`
	TransferAmount int32            `db:"transfer_amount" json:"transfer_amount"`
	TransferTime   time.Time        `db:"transfer_time" json:"transfer_time"`
	Status         string           `db:"status" json:"status"`
	CreatedAt      pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt      pgtype.Timestamp `db:"updated_at" json:"updated_at"`
	DeletedAt      pgtype.Timestamp `db:"deleted_at" json:"deleted_at"`
}

type UpdateTransferParams struct {
	TransferID     int32     `db:"transfer_id" json:"transfer_id"`
	TransferFrom   string    `db:"transfer_from" json:"transfer_from"`
	TransferTo     string    `db:"transfer_to" json:"transfer_to"`
	TransferAmount int32     `db:"transfer_amount" json:"transfer_amount"`
	TransferTime   time.Time `db:"transfer_time" json:"transfer_time"`
}

type UpdateTransferRow struct {
	TransferID     int32            `db:"transfer_id" json:"transfer_id"`
	TransferNo     pgtype.UUID      `db:"transfer_no" json:"transfer_no"`
	TransferFrom   string           `db:"transfer_from" json:"transfer_from"`
	TransferTo     string           `db:"transfer_to" json:"transfer_to"`
	TransferAmount int32            `db:"transfer_amount" json:"transfer_amount"`
	TransferTime   time.Time        `db:"transfer_time" json:"transfer_time"`
	Status         string           `db:"status" json:"status"`
	CreatedAt      pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt      pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type UpdateTransferAmountParams struct {
	TransferID     int32 `db:"transfer_id" json:"transfer_id"`
	TransferAmount int32 `db:"transfer_amount" json:"transfer_amount"`
}

type UpdateTransferAmountRow struct {
	TransferID     int32            `db:"transfer_id" json:"transfer_id"`
	TransferNo     pgtype.UUID      `db:"transfer_no" json:"transfer_no"`
	TransferFrom   string           `db:"transfer_from" json:"transfer_from"`
	TransferTo     string           `db:"transfer_to" json:"transfer_to"`
	TransferAmount int32            `db:"transfer_amount" json:"transfer_amount"`
	TransferTime   time.Time        `db:"transfer_time" json:"transfer_time"`
	Status         string           `db:"status" json:"status"`
	CreatedAt      pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt      pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type UpdateTransferAtomicParams struct {
	TransferID     int32     `db:"transfer_id" json:"transfer_id"`
	TransferFrom   string    `db:"transfer_from" json:"transfer_from"`
	TransferTo     string    `db:"transfer_to" json:"transfer_to"`
	TransferAmount int32     `db:"transfer_amount" json:"transfer_amount"`
	TransferTime   time.Time `db:"transfer_time" json:"transfer_time"`
}

type UpdateTransferAtomicRow struct {
	TransferID     int32            `db:"transfer_id" json:"transfer_id"`
	TransferNo     pgtype.UUID      `db:"transfer_no" json:"transfer_no"`
	TransferFrom   string           `db:"transfer_from" json:"transfer_from"`
	TransferTo     string           `db:"transfer_to" json:"transfer_to"`
	TransferAmount int32            `db:"transfer_amount" json:"transfer_amount"`
	TransferTime   time.Time        `db:"transfer_time" json:"transfer_time"`
	Status         string           `db:"status" json:"status"`
	CreatedAt      pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt      pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type UpdateTransferSettlementDeltaParams struct {
	CardNumber     string `db:"card_number" json:"card_number"`
	CardNumber_2   string `db:"card_number_2" json:"card_number_2"`
	TotalBalance   int32  `db:"total_balance" json:"total_balance"`
	TotalBalance_2 int32  `db:"total_balance_2" json:"total_balance_2"`
}

type UpdateTransferStatusParams struct {
	TransferID int32  `db:"transfer_id" json:"transfer_id"`
	Status     string `db:"status" json:"status"`
}

type UpdateTransferStatusRow struct {
	TransferID     int32            `db:"transfer_id" json:"transfer_id"`
	TransferNo     pgtype.UUID      `db:"transfer_no" json:"transfer_no"`
	TransferFrom   string           `db:"transfer_from" json:"transfer_from"`
	TransferTo     string           `db:"transfer_to" json:"transfer_to"`
	TransferAmount int32            `db:"transfer_amount" json:"transfer_amount"`
	TransferTime   time.Time        `db:"transfer_time" json:"transfer_time"`
	Status         string           `db:"status" json:"status"`
	CreatedAt      pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt      pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}
