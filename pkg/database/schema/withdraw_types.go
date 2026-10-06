// Code migrated from sqlc-generated sources; domain types kept for repository layer.
// source: withdraw.sql

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
	"time"
)

type CreateWithdrawParams struct {
	CardNumber     string           `db:"card_number" json:"card_number"`
	WithdrawAmount int32            `db:"withdraw_amount" json:"withdraw_amount"`
	WithdrawTime   time.Time        `db:"withdraw_time" json:"withdraw_time"`
	CreatedAt      pgtype.Timestamp `db:"created_at" json:"created_at"`
}

type CreateWithdrawRow struct {
	WithdrawID     int32            `db:"withdraw_id" json:"withdraw_id"`
	WithdrawNo     pgtype.UUID      `db:"withdraw_no" json:"withdraw_no"`
	CardNumber     string           `db:"card_number" json:"card_number"`
	WithdrawAmount int32            `db:"withdraw_amount" json:"withdraw_amount"`
	WithdrawTime   time.Time        `db:"withdraw_time" json:"withdraw_time"`
	Status         string           `db:"status" json:"status"`
	CreatedAt      pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt      pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type CreateWithdrawAtomicParams struct {
	CardNumber     string    `db:"card_number" json:"card_number"`
	WithdrawTime   time.Time `db:"withdraw_time" json:"withdraw_time"`
	WithdrawAmount int32     `db:"withdraw_amount" json:"withdraw_amount"`
	DailyLimit     int32     `db:"daily_limit" json:"daily_limit"`
	IdempotencyKey string    `db:"idempotency_key" json:"idempotency_key"`
}

type CreateWithdrawAtomicRow struct {
	WithdrawID     int32            `db:"withdraw_id" json:"withdraw_id"`
	WithdrawNo     pgtype.UUID      `db:"withdraw_no" json:"withdraw_no"`
	CardNumber     string           `db:"card_number" json:"card_number"`
	WithdrawAmount int32            `db:"withdraw_amount" json:"withdraw_amount"`
	WithdrawTime   time.Time        `db:"withdraw_time" json:"withdraw_time"`
	Status         string           `db:"status" json:"status"`
	CreatedAt      pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt      pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type FindAllWithdrawsByCardNumberRow struct {
	WithdrawID     int32            `db:"withdraw_id" json:"withdraw_id"`
	CardNumber     string           `db:"card_number" json:"card_number"`
	WithdrawAmount int32            `db:"withdraw_amount" json:"withdraw_amount"`
	WithdrawTime   time.Time        `db:"withdraw_time" json:"withdraw_time"`
	CreatedAt      pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt      pgtype.Timestamp `db:"updated_at" json:"updated_at"`
	DeletedAt      pgtype.Timestamp `db:"deleted_at" json:"deleted_at"`
}

type GetActiveWithdrawsParams struct {
	Column1 string `db:"column1" json:"column_1"`
	Limit   int32  `db:"limit" json:"limit"`
	Offset  int32  `db:"offset" json:"offset"`
}

type GetActiveWithdrawsRow struct {
	WithdrawID     int32            `db:"withdraw_id" json:"withdraw_id"`
	WithdrawNo     pgtype.UUID      `db:"withdraw_no" json:"withdraw_no"`
	CardNumber     string           `db:"card_number" json:"card_number"`
	WithdrawAmount int32            `db:"withdraw_amount" json:"withdraw_amount"`
	WithdrawTime   time.Time        `db:"withdraw_time" json:"withdraw_time"`
	Status         string           `db:"status" json:"status"`
	CreatedAt      pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt      pgtype.Timestamp `db:"updated_at" json:"updated_at"`
	DeletedAt      pgtype.Timestamp `db:"deleted_at" json:"deleted_at"`
	TotalCount     int64            `db:"total_count" json:"total_count"`
}

type GetMonthWithdrawStatusFailedParams struct {
	Column1 time.Time `db:"column1" json:"column_1"`
	Column2 time.Time `db:"column2" json:"column_2"`
	Column3 time.Time `db:"column3" json:"column_3"`
	Column4 time.Time `db:"column4" json:"column_4"`
}

type GetMonthWithdrawStatusFailedRow struct {
	Year        string `db:"year" json:"year"`
	Month       string `db:"month" json:"month"`
	TotalFailed int64  `db:"total_failed" json:"total_failed"`
	TotalAmount int32  `db:"total_amount" json:"total_amount"`
}

type GetMonthWithdrawStatusFailedCardNumberParams struct {
	CardNumber string    `db:"card_number" json:"card_number"`
	Column2    time.Time `db:"column2" json:"column_2"`
	Column3    time.Time `db:"column3" json:"column_3"`
	Column4    time.Time `db:"column4" json:"column_4"`
	Column5    time.Time `db:"column5" json:"column_5"`
}

type GetMonthWithdrawStatusFailedCardNumberRow struct {
	Year        string `db:"year" json:"year"`
	Month       string `db:"month" json:"month"`
	TotalFailed int64  `db:"total_failed" json:"total_failed"`
	TotalAmount int32  `db:"total_amount" json:"total_amount"`
}

type GetMonthWithdrawStatusSuccessParams struct {
	Column1 time.Time `db:"column1" json:"column_1"`
	Column2 time.Time `db:"column2" json:"column_2"`
	Column3 time.Time `db:"column3" json:"column_3"`
	Column4 time.Time `db:"column4" json:"column_4"`
}

type GetMonthWithdrawStatusSuccessRow struct {
	Year         string `db:"year" json:"year"`
	Month        string `db:"month" json:"month"`
	TotalSuccess int64  `db:"total_success" json:"total_success"`
	TotalAmount  int32  `db:"total_amount" json:"total_amount"`
}

type GetMonthWithdrawStatusSuccessCardNumberParams struct {
	CardNumber string    `db:"card_number" json:"card_number"`
	Column2    time.Time `db:"column2" json:"column_2"`
	Column3    time.Time `db:"column3" json:"column_3"`
	Column4    time.Time `db:"column4" json:"column_4"`
	Column5    time.Time `db:"column5" json:"column_5"`
}

type GetMonthWithdrawStatusSuccessCardNumberRow struct {
	Year         string `db:"year" json:"year"`
	Month        string `db:"month" json:"month"`
	TotalSuccess int64  `db:"total_success" json:"total_success"`
	TotalAmount  int32  `db:"total_amount" json:"total_amount"`
}

type GetMonthlyWithdrawsRow struct {
	Month               string `db:"month" json:"month"`
	TotalWithdrawAmount int32  `db:"total_withdraw_amount" json:"total_withdraw_amount"`
}

type GetMonthlyWithdrawsByCardNumberParams struct {
	CardNumber string    `db:"card_number" json:"card_number"`
	Column2    time.Time `db:"column2" json:"column_2"`
}

type GetMonthlyWithdrawsByCardNumberRow struct {
	Month               string `db:"month" json:"month"`
	TotalWithdrawAmount int32  `db:"total_withdraw_amount" json:"total_withdraw_amount"`
}

type GetTrashedWithdrawsParams struct {
	Column1 string `db:"column1" json:"column_1"`
	Limit   int32  `db:"limit" json:"limit"`
	Offset  int32  `db:"offset" json:"offset"`
}

type GetTrashedWithdrawsRow struct {
	WithdrawID     int32            `db:"withdraw_id" json:"withdraw_id"`
	WithdrawNo     pgtype.UUID      `db:"withdraw_no" json:"withdraw_no"`
	CardNumber     string           `db:"card_number" json:"card_number"`
	WithdrawAmount int32            `db:"withdraw_amount" json:"withdraw_amount"`
	WithdrawTime   time.Time        `db:"withdraw_time" json:"withdraw_time"`
	Status         string           `db:"status" json:"status"`
	CreatedAt      pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt      pgtype.Timestamp `db:"updated_at" json:"updated_at"`
	DeletedAt      pgtype.Timestamp `db:"deleted_at" json:"deleted_at"`
	TotalCount     int64            `db:"total_count" json:"total_count"`
}

type GetWithdrawByIDRow struct {
	WithdrawID     int32            `db:"withdraw_id" json:"withdraw_id"`
	WithdrawNo     pgtype.UUID      `db:"withdraw_no" json:"withdraw_no"`
	CardNumber     string           `db:"card_number" json:"card_number"`
	WithdrawAmount int32            `db:"withdraw_amount" json:"withdraw_amount"`
	WithdrawTime   time.Time        `db:"withdraw_time" json:"withdraw_time"`
	Status         string           `db:"status" json:"status"`
	CreatedAt      pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt      pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type GetWithdrawByIdempotencyKeyRow struct {
	WithdrawID     int32            `db:"withdraw_id" json:"withdraw_id"`
	WithdrawNo     pgtype.UUID      `db:"withdraw_no" json:"withdraw_no"`
	CardNumber     string           `db:"card_number" json:"card_number"`
	WithdrawAmount int32            `db:"withdraw_amount" json:"withdraw_amount"`
	WithdrawTime   time.Time        `db:"withdraw_time" json:"withdraw_time"`
	Status         string           `db:"status" json:"status"`
	CreatedAt      pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt      pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type GetWithdrawsParams struct {
	Column1 string `db:"column1" json:"column_1"`
	Limit   int32  `db:"limit" json:"limit"`
	Offset  int32  `db:"offset" json:"offset"`
}

type GetWithdrawsRow struct {
	WithdrawID     int32            `db:"withdraw_id" json:"withdraw_id"`
	WithdrawNo     pgtype.UUID      `db:"withdraw_no" json:"withdraw_no"`
	CardNumber     string           `db:"card_number" json:"card_number"`
	WithdrawAmount int32            `db:"withdraw_amount" json:"withdraw_amount"`
	WithdrawTime   time.Time        `db:"withdraw_time" json:"withdraw_time"`
	Status         string           `db:"status" json:"status"`
	CreatedAt      pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt      pgtype.Timestamp `db:"updated_at" json:"updated_at"`
	TotalCount     int64            `db:"total_count" json:"total_count"`
}

type GetWithdrawsByCardNumberParams struct {
	CardNumber string `db:"card_number" json:"card_number"`
	Column2    string `db:"column2" json:"column_2"`
	Limit      int32  `db:"limit" json:"limit"`
	Offset     int32  `db:"offset" json:"offset"`
}

type GetWithdrawsByCardNumberRow struct {
	WithdrawID     int32            `db:"withdraw_id" json:"withdraw_id"`
	WithdrawNo     pgtype.UUID      `db:"withdraw_no" json:"withdraw_no"`
	CardNumber     string           `db:"card_number" json:"card_number"`
	WithdrawAmount int32            `db:"withdraw_amount" json:"withdraw_amount"`
	WithdrawTime   time.Time        `db:"withdraw_time" json:"withdraw_time"`
	Status         string           `db:"status" json:"status"`
	CreatedAt      pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt      pgtype.Timestamp `db:"updated_at" json:"updated_at"`
	TotalCount     int64            `db:"total_count" json:"total_count"`
}

type GetYearlyWithdrawStatusFailedRow struct {
	Year        string `db:"year" json:"year"`
	TotalFailed int32  `db:"total_failed" json:"total_failed"`
	TotalAmount int32  `db:"total_amount" json:"total_amount"`
}

type GetYearlyWithdrawStatusFailedCardNumberParams struct {
	CardNumber string `db:"card_number" json:"card_number"`
	Column2    int32  `db:"column2" json:"column_2"`
}

type GetYearlyWithdrawStatusFailedCardNumberRow struct {
	Year        string `db:"year" json:"year"`
	TotalFailed int32  `db:"total_failed" json:"total_failed"`
	TotalAmount int32  `db:"total_amount" json:"total_amount"`
}

type GetYearlyWithdrawStatusSuccessRow struct {
	Year         string `db:"year" json:"year"`
	TotalSuccess int32  `db:"total_success" json:"total_success"`
	TotalAmount  int32  `db:"total_amount" json:"total_amount"`
}

type GetYearlyWithdrawStatusSuccessCardNumberParams struct {
	CardNumber string `db:"card_number" json:"card_number"`
	Column2    int32  `db:"column2" json:"column_2"`
}

type GetYearlyWithdrawStatusSuccessCardNumberRow struct {
	Year         string `db:"year" json:"year"`
	TotalSuccess int32  `db:"total_success" json:"total_success"`
	TotalAmount  int32  `db:"total_amount" json:"total_amount"`
}

type GetYearlyWithdrawsRow struct {
	Year                pgtype.Numeric `db:"year" json:"year"`
	TotalWithdrawAmount int64          `db:"total_withdraw_amount" json:"total_withdraw_amount"`
}

type GetYearlyWithdrawsByCardNumberParams struct {
	CardNumber string      `db:"card_number" json:"card_number"`
	Column2    interface{} `db:"column2" json:"column_2"`
}

type GetYearlyWithdrawsByCardNumberRow struct {
	Year                pgtype.Numeric `db:"year" json:"year"`
	TotalWithdrawAmount int64          `db:"total_withdraw_amount" json:"total_withdraw_amount"`
}

type RestoreWithdrawRow struct {
	WithdrawID     int32            `db:"withdraw_id" json:"withdraw_id"`
	WithdrawNo     pgtype.UUID      `db:"withdraw_no" json:"withdraw_no"`
	CardNumber     string           `db:"card_number" json:"card_number"`
	WithdrawAmount int32            `db:"withdraw_amount" json:"withdraw_amount"`
	WithdrawTime   time.Time        `db:"withdraw_time" json:"withdraw_time"`
	Status         string           `db:"status" json:"status"`
	CreatedAt      pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt      pgtype.Timestamp `db:"updated_at" json:"updated_at"`
	DeletedAt      pgtype.Timestamp `db:"deleted_at" json:"deleted_at"`
}

type TrashWithdrawRow struct {
	WithdrawID     int32            `db:"withdraw_id" json:"withdraw_id"`
	WithdrawNo     pgtype.UUID      `db:"withdraw_no" json:"withdraw_no"`
	CardNumber     string           `db:"card_number" json:"card_number"`
	WithdrawAmount int32            `db:"withdraw_amount" json:"withdraw_amount"`
	WithdrawTime   time.Time        `db:"withdraw_time" json:"withdraw_time"`
	Status         string           `db:"status" json:"status"`
	CreatedAt      pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt      pgtype.Timestamp `db:"updated_at" json:"updated_at"`
	DeletedAt      pgtype.Timestamp `db:"deleted_at" json:"deleted_at"`
}

type UpdateWithdrawParams struct {
	WithdrawID     int32     `db:"withdraw_id" json:"withdraw_id"`
	CardNumber     string    `db:"card_number" json:"card_number"`
	WithdrawAmount int32     `db:"withdraw_amount" json:"withdraw_amount"`
	WithdrawTime   time.Time `db:"withdraw_time" json:"withdraw_time"`
}

type UpdateWithdrawRow struct {
	WithdrawID     int32            `db:"withdraw_id" json:"withdraw_id"`
	WithdrawNo     pgtype.UUID      `db:"withdraw_no" json:"withdraw_no"`
	CardNumber     string           `db:"card_number" json:"card_number"`
	WithdrawAmount int32            `db:"withdraw_amount" json:"withdraw_amount"`
	WithdrawTime   time.Time        `db:"withdraw_time" json:"withdraw_time"`
	Status         string           `db:"status" json:"status"`
	CreatedAt      pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt      pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type UpdateWithdrawAtomicParams struct {
	WithdrawID     int32     `db:"withdraw_id" json:"withdraw_id"`
	WithdrawAmount int32     `db:"withdraw_amount" json:"withdraw_amount"`
	CardNumber     string    `db:"card_number" json:"card_number"`
	WithdrawTime   time.Time `db:"withdraw_time" json:"withdraw_time"`
}

type UpdateWithdrawAtomicRow struct {
	WithdrawID     int32            `db:"withdraw_id" json:"withdraw_id"`
	WithdrawNo     pgtype.UUID      `db:"withdraw_no" json:"withdraw_no"`
	CardNumber     string           `db:"card_number" json:"card_number"`
	WithdrawAmount int32            `db:"withdraw_amount" json:"withdraw_amount"`
	WithdrawTime   time.Time        `db:"withdraw_time" json:"withdraw_time"`
	Status         string           `db:"status" json:"status"`
	CreatedAt      pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt      pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type UpdateWithdrawStatusParams struct {
	WithdrawID int32  `db:"withdraw_id" json:"withdraw_id"`
	Status     string `db:"status" json:"status"`
}

type UpdateWithdrawStatusRow struct {
	WithdrawID     int32            `db:"withdraw_id" json:"withdraw_id"`
	WithdrawNo     pgtype.UUID      `db:"withdraw_no" json:"withdraw_no"`
	CardNumber     string           `db:"card_number" json:"card_number"`
	WithdrawAmount int32            `db:"withdraw_amount" json:"withdraw_amount"`
	WithdrawTime   time.Time        `db:"withdraw_time" json:"withdraw_time"`
	Status         string           `db:"status" json:"status"`
	CreatedAt      pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt      pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}
