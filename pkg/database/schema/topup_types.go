// Code migrated from sqlc-generated sources; domain types kept for repository layer.
// source: topup.sql

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
	"time"
)

type CreateTopupParams struct {
	CardNumber  string    `db:"card_number" json:"card_number"`
	TopupAmount int32     `db:"topup_amount" json:"topup_amount"`
	TopupMethod string    `db:"topup_method" json:"topup_method"`
	TopupTime   time.Time `db:"topup_time" json:"topup_time"`
}

type CreateTopupRow struct {
	TopupID     int32            `db:"topup_id" json:"topup_id"`
	TopupNo     pgtype.UUID      `db:"topup_no" json:"topup_no"`
	CardNumber  string           `db:"card_number" json:"card_number"`
	TopupAmount int32            `db:"topup_amount" json:"topup_amount"`
	TopupMethod string           `db:"topup_method" json:"topup_method"`
	TopupTime   time.Time        `db:"topup_time" json:"topup_time"`
	Status      string           `db:"status" json:"status"`
	CreatedAt   pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt   pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type CreateTopupAtomicParams struct {
	CardNumber     string `db:"card_number" json:"card_number"`
	TopupAmount    int32  `db:"topup_amount" json:"topup_amount"`
	TopupMethod    string `db:"topup_method" json:"topup_method"`
	IdempotencyKey string `db:"idempotency_key" json:"idempotency_key"`
}

type CreateTopupAtomicRow struct {
	TopupID     int32            `db:"topup_id" json:"topup_id"`
	TopupNo     pgtype.UUID      `db:"topup_no" json:"topup_no"`
	CardNumber  string           `db:"card_number" json:"card_number"`
	TopupAmount int32            `db:"topup_amount" json:"topup_amount"`
	TopupMethod string           `db:"topup_method" json:"topup_method"`
	TopupTime   time.Time        `db:"topup_time" json:"topup_time"`
	Status      string           `db:"status" json:"status"`
	CreatedAt   pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt   pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type GetActiveTopupsParams struct {
	Column1 string `db:"column1" json:"column_1"`
	Limit   int32  `db:"limit" json:"limit"`
	Offset  int32  `db:"offset" json:"offset"`
}

type GetActiveTopupsRow struct {
	TopupID     int32            `db:"topup_id" json:"topup_id"`
	TopupNo     pgtype.UUID      `db:"topup_no" json:"topup_no"`
	CardNumber  string           `db:"card_number" json:"card_number"`
	TopupAmount int32            `db:"topup_amount" json:"topup_amount"`
	TopupMethod string           `db:"topup_method" json:"topup_method"`
	TopupTime   time.Time        `db:"topup_time" json:"topup_time"`
	Status      string           `db:"status" json:"status"`
	CreatedAt   pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt   pgtype.Timestamp `db:"updated_at" json:"updated_at"`
	DeletedAt   pgtype.Timestamp `db:"deleted_at" json:"deleted_at"`
	TotalCount  int64            `db:"total_count" json:"total_count"`
}

type GetMonthTopupStatusFailedParams struct {
	Column1 time.Time `db:"column1" json:"column_1"`
	Column2 time.Time `db:"column2" json:"column_2"`
	Column3 time.Time `db:"column3" json:"column_3"`
	Column4 time.Time `db:"column4" json:"column_4"`
}

type GetMonthTopupStatusFailedRow struct {
	Year        string `db:"year" json:"year"`
	Month       string `db:"month" json:"month"`
	TotalFailed int64  `db:"total_failed" json:"total_failed"`
	TotalAmount int32  `db:"total_amount" json:"total_amount"`
}

type GetMonthTopupStatusFailedCardNumberParams struct {
	CardNumber string    `db:"card_number" json:"card_number"`
	Column2    time.Time `db:"column2" json:"column_2"`
	Column3    time.Time `db:"column3" json:"column_3"`
	Column4    time.Time `db:"column4" json:"column_4"`
	Column5    time.Time `db:"column5" json:"column_5"`
}

type GetMonthTopupStatusFailedCardNumberRow struct {
	Year        string `db:"year" json:"year"`
	Month       string `db:"month" json:"month"`
	TotalFailed int64  `db:"total_failed" json:"total_failed"`
	TotalAmount int32  `db:"total_amount" json:"total_amount"`
}

type GetMonthTopupStatusSuccessParams struct {
	Column1 time.Time `db:"column1" json:"column_1"`
	Column2 time.Time `db:"column2" json:"column_2"`
	Column3 time.Time `db:"column3" json:"column_3"`
	Column4 time.Time `db:"column4" json:"column_4"`
}

type GetMonthTopupStatusSuccessRow struct {
	Year         string `db:"year" json:"year"`
	Month        string `db:"month" json:"month"`
	TotalSuccess int64  `db:"total_success" json:"total_success"`
	TotalAmount  int32  `db:"total_amount" json:"total_amount"`
}

type GetMonthTopupStatusSuccessCardNumberParams struct {
	CardNumber string    `db:"card_number" json:"card_number"`
	Column2    time.Time `db:"column2" json:"column_2"`
	Column3    time.Time `db:"column3" json:"column_3"`
	Column4    time.Time `db:"column4" json:"column_4"`
	Column5    time.Time `db:"column5" json:"column_5"`
}

type GetMonthTopupStatusSuccessCardNumberRow struct {
	Year         string `db:"year" json:"year"`
	Month        string `db:"month" json:"month"`
	TotalSuccess int64  `db:"total_success" json:"total_success"`
	TotalAmount  int32  `db:"total_amount" json:"total_amount"`
}

type GetMonthlyTopupAmountsRow struct {
	Month       string `db:"month" json:"month"`
	TotalAmount int32  `db:"total_amount" json:"total_amount"`
}

type GetMonthlyTopupAmountsByCardNumberParams struct {
	CardNumber string    `db:"card_number" json:"card_number"`
	Column2    time.Time `db:"column2" json:"column_2"`
}

type GetMonthlyTopupAmountsByCardNumberRow struct {
	Month       string `db:"month" json:"month"`
	TotalAmount int32  `db:"total_amount" json:"total_amount"`
}

type GetMonthlyTopupMethodsRow struct {
	Month       string `db:"month" json:"month"`
	TopupMethod string `db:"topup_method" json:"topup_method"`
	TotalTopups int32  `db:"total_topups" json:"total_topups"`
	TotalAmount int32  `db:"total_amount" json:"total_amount"`
}

type GetMonthlyTopupMethodsByCardNumberParams struct {
	CardNumber string    `db:"card_number" json:"card_number"`
	Column2    time.Time `db:"column2" json:"column_2"`
}

type GetMonthlyTopupMethodsByCardNumberRow struct {
	Month       string `db:"month" json:"month"`
	TopupMethod string `db:"topup_method" json:"topup_method"`
	TotalTopups int32  `db:"total_topups" json:"total_topups"`
	TotalAmount int32  `db:"total_amount" json:"total_amount"`
}

type GetTopupByIDRow struct {
	TopupID     int32            `db:"topup_id" json:"topup_id"`
	TopupNo     pgtype.UUID      `db:"topup_no" json:"topup_no"`
	CardNumber  string           `db:"card_number" json:"card_number"`
	TopupAmount int32            `db:"topup_amount" json:"topup_amount"`
	TopupMethod string           `db:"topup_method" json:"topup_method"`
	TopupTime   time.Time        `db:"topup_time" json:"topup_time"`
	Status      string           `db:"status" json:"status"`
	CreatedAt   pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt   pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type GetTopupByIdempotencyKeyRow struct {
	TopupID     int32            `db:"topup_id" json:"topup_id"`
	TopupNo     pgtype.UUID      `db:"topup_no" json:"topup_no"`
	CardNumber  string           `db:"card_number" json:"card_number"`
	TopupAmount int32            `db:"topup_amount" json:"topup_amount"`
	TopupMethod string           `db:"topup_method" json:"topup_method"`
	TopupTime   time.Time        `db:"topup_time" json:"topup_time"`
	Status      string           `db:"status" json:"status"`
	CreatedAt   pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt   pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type GetTopupsParams struct {
	Column1 string `db:"column1" json:"column_1"`
	Limit   int32  `db:"limit" json:"limit"`
	Offset  int32  `db:"offset" json:"offset"`
}

type GetTopupsRow struct {
	TopupID     int32            `db:"topup_id" json:"topup_id"`
	TopupNo     pgtype.UUID      `db:"topup_no" json:"topup_no"`
	CardNumber  string           `db:"card_number" json:"card_number"`
	TopupAmount int32            `db:"topup_amount" json:"topup_amount"`
	TopupMethod string           `db:"topup_method" json:"topup_method"`
	TopupTime   time.Time        `db:"topup_time" json:"topup_time"`
	Status      string           `db:"status" json:"status"`
	CreatedAt   pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt   pgtype.Timestamp `db:"updated_at" json:"updated_at"`
	TotalCount  int64            `db:"total_count" json:"total_count"`
}

type GetTopupsByCardNumberParams struct {
	CardNumber string `db:"card_number" json:"card_number"`
	Column2    string `db:"column2" json:"column_2"`
	Limit      int32  `db:"limit" json:"limit"`
	Offset     int32  `db:"offset" json:"offset"`
}

type GetTopupsByCardNumberRow struct {
	TopupID     int32            `db:"topup_id" json:"topup_id"`
	TopupNo     pgtype.UUID      `db:"topup_no" json:"topup_no"`
	CardNumber  string           `db:"card_number" json:"card_number"`
	TopupAmount int32            `db:"topup_amount" json:"topup_amount"`
	TopupMethod string           `db:"topup_method" json:"topup_method"`
	TopupTime   time.Time        `db:"topup_time" json:"topup_time"`
	Status      string           `db:"status" json:"status"`
	CreatedAt   pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt   pgtype.Timestamp `db:"updated_at" json:"updated_at"`
	TotalCount  int64            `db:"total_count" json:"total_count"`
}

type GetTrashedTopupsParams struct {
	Column1 string `db:"column1" json:"column_1"`
	Limit   int32  `db:"limit" json:"limit"`
	Offset  int32  `db:"offset" json:"offset"`
}

type GetTrashedTopupsRow struct {
	TopupID     int32            `db:"topup_id" json:"topup_id"`
	TopupNo     pgtype.UUID      `db:"topup_no" json:"topup_no"`
	CardNumber  string           `db:"card_number" json:"card_number"`
	TopupAmount int32            `db:"topup_amount" json:"topup_amount"`
	TopupMethod string           `db:"topup_method" json:"topup_method"`
	TopupTime   time.Time        `db:"topup_time" json:"topup_time"`
	Status      string           `db:"status" json:"status"`
	CreatedAt   pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt   pgtype.Timestamp `db:"updated_at" json:"updated_at"`
	DeletedAt   pgtype.Timestamp `db:"deleted_at" json:"deleted_at"`
	TotalCount  int64            `db:"total_count" json:"total_count"`
}

type GetYearlyTopupAmountsRow struct {
	Year        pgtype.Numeric `db:"year" json:"year"`
	TotalAmount int64          `db:"total_amount" json:"total_amount"`
}

type GetYearlyTopupAmountsByCardNumberParams struct {
	CardNumber string      `db:"card_number" json:"card_number"`
	Column2    interface{} `db:"column2" json:"column_2"`
}

type GetYearlyTopupAmountsByCardNumberRow struct {
	Year        pgtype.Numeric `db:"year" json:"year"`
	TotalAmount int64          `db:"total_amount" json:"total_amount"`
}

type GetYearlyTopupMethodsRow struct {
	Year        pgtype.Numeric `db:"year" json:"year"`
	TopupMethod string         `db:"topup_method" json:"topup_method"`
	TotalTopups int64          `db:"total_topups" json:"total_topups"`
	TotalAmount int64          `db:"total_amount" json:"total_amount"`
}

type GetYearlyTopupMethodsByCardNumberParams struct {
	CardNumber string      `db:"card_number" json:"card_number"`
	Column2    interface{} `db:"column2" json:"column_2"`
}

type GetYearlyTopupMethodsByCardNumberRow struct {
	Year        pgtype.Numeric `db:"year" json:"year"`
	TopupMethod string         `db:"topup_method" json:"topup_method"`
	TotalTopups int64          `db:"total_topups" json:"total_topups"`
	TotalAmount int64          `db:"total_amount" json:"total_amount"`
}

type GetYearlyTopupStatusFailedRow struct {
	Year        string `db:"year" json:"year"`
	TotalFailed int32  `db:"total_failed" json:"total_failed"`
	TotalAmount int32  `db:"total_amount" json:"total_amount"`
}

type GetYearlyTopupStatusFailedCardNumberParams struct {
	CardNumber string `db:"card_number" json:"card_number"`
	Column2    int32  `db:"column2" json:"column_2"`
}

type GetYearlyTopupStatusFailedCardNumberRow struct {
	Year        string `db:"year" json:"year"`
	TotalFailed int32  `db:"total_failed" json:"total_failed"`
	TotalAmount int32  `db:"total_amount" json:"total_amount"`
}

type GetYearlyTopupStatusSuccessRow struct {
	Year         string `db:"year" json:"year"`
	TotalSuccess int32  `db:"total_success" json:"total_success"`
	TotalAmount  int32  `db:"total_amount" json:"total_amount"`
}

type GetYearlyTopupStatusSuccessCardNumberParams struct {
	CardNumber string `db:"card_number" json:"card_number"`
	Column2    int32  `db:"column2" json:"column_2"`
}

type GetYearlyTopupStatusSuccessCardNumberRow struct {
	Year         string `db:"year" json:"year"`
	TotalSuccess int32  `db:"total_success" json:"total_success"`
	TotalAmount  int32  `db:"total_amount" json:"total_amount"`
}

type RestoreTopupRow struct {
	TopupID     int32            `db:"topup_id" json:"topup_id"`
	TopupNo     pgtype.UUID      `db:"topup_no" json:"topup_no"`
	CardNumber  string           `db:"card_number" json:"card_number"`
	TopupAmount int32            `db:"topup_amount" json:"topup_amount"`
	TopupMethod string           `db:"topup_method" json:"topup_method"`
	TopupTime   time.Time        `db:"topup_time" json:"topup_time"`
	Status      string           `db:"status" json:"status"`
	CreatedAt   pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt   pgtype.Timestamp `db:"updated_at" json:"updated_at"`
	DeletedAt   pgtype.Timestamp `db:"deleted_at" json:"deleted_at"`
}

type TrashTopupRow struct {
	TopupID     int32            `db:"topup_id" json:"topup_id"`
	TopupNo     pgtype.UUID      `db:"topup_no" json:"topup_no"`
	CardNumber  string           `db:"card_number" json:"card_number"`
	TopupAmount int32            `db:"topup_amount" json:"topup_amount"`
	TopupMethod string           `db:"topup_method" json:"topup_method"`
	TopupTime   time.Time        `db:"topup_time" json:"topup_time"`
	Status      string           `db:"status" json:"status"`
	CreatedAt   pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt   pgtype.Timestamp `db:"updated_at" json:"updated_at"`
	DeletedAt   pgtype.Timestamp `db:"deleted_at" json:"deleted_at"`
}

type UpdateTopupParams struct {
	TopupID     int32     `db:"topup_id" json:"topup_id"`
	CardNumber  string    `db:"card_number" json:"card_number"`
	TopupAmount int32     `db:"topup_amount" json:"topup_amount"`
	TopupMethod string    `db:"topup_method" json:"topup_method"`
	TopupTime   time.Time `db:"topup_time" json:"topup_time"`
}

type UpdateTopupRow struct {
	TopupID     int32            `db:"topup_id" json:"topup_id"`
	TopupNo     pgtype.UUID      `db:"topup_no" json:"topup_no"`
	CardNumber  string           `db:"card_number" json:"card_number"`
	TopupAmount int32            `db:"topup_amount" json:"topup_amount"`
	TopupMethod string           `db:"topup_method" json:"topup_method"`
	TopupTime   time.Time        `db:"topup_time" json:"topup_time"`
	Status      string           `db:"status" json:"status"`
	CreatedAt   pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt   pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type UpdateTopupAmountParams struct {
	TopupID     int32 `db:"topup_id" json:"topup_id"`
	TopupAmount int32 `db:"topup_amount" json:"topup_amount"`
}

type UpdateTopupAmountRow struct {
	TopupID     int32            `db:"topup_id" json:"topup_id"`
	TopupNo     pgtype.UUID      `db:"topup_no" json:"topup_no"`
	CardNumber  string           `db:"card_number" json:"card_number"`
	TopupAmount int32            `db:"topup_amount" json:"topup_amount"`
	TopupMethod string           `db:"topup_method" json:"topup_method"`
	TopupTime   time.Time        `db:"topup_time" json:"topup_time"`
	Status      string           `db:"status" json:"status"`
	CreatedAt   pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt   pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type UpdateTopupAtomicParams struct {
	TopupID     int32  `db:"topup_id" json:"topup_id"`
	TopupAmount int32  `db:"topup_amount" json:"topup_amount"`
	CardNumber  string `db:"card_number" json:"card_number"`
	TopupMethod string `db:"topup_method" json:"topup_method"`
}

type UpdateTopupAtomicRow struct {
	TopupID     int32            `db:"topup_id" json:"topup_id"`
	TopupNo     pgtype.UUID      `db:"topup_no" json:"topup_no"`
	CardNumber  string           `db:"card_number" json:"card_number"`
	TopupAmount int32            `db:"topup_amount" json:"topup_amount"`
	TopupMethod string           `db:"topup_method" json:"topup_method"`
	TopupTime   time.Time        `db:"topup_time" json:"topup_time"`
	Status      string           `db:"status" json:"status"`
	CreatedAt   pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt   pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type UpdateTopupStatusParams struct {
	TopupID int32  `db:"topup_id" json:"topup_id"`
	Status  string `db:"status" json:"status"`
}

type UpdateTopupStatusRow struct {
	TopupID     int32            `db:"topup_id" json:"topup_id"`
	TopupNo     pgtype.UUID      `db:"topup_no" json:"topup_no"`
	CardNumber  string           `db:"card_number" json:"card_number"`
	TopupAmount int32            `db:"topup_amount" json:"topup_amount"`
	TopupMethod string           `db:"topup_method" json:"topup_method"`
	TopupTime   time.Time        `db:"topup_time" json:"topup_time"`
	Status      string           `db:"status" json:"status"`
	CreatedAt   pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt   pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}
