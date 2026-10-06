// Code migrated from sqlc-generated sources; domain types kept for repository layer.
// source: saldo.sql

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
	"time"
)

type CreateSaldoParams struct {
	CardNumber   string `db:"card_number" json:"card_number"`
	TotalBalance int32  `db:"total_balance" json:"total_balance"`
}

type CreateSaldoRow struct {
	SaldoID        int32            `db:"saldo_id" json:"saldo_id"`
	CardNumber     string           `db:"card_number" json:"card_number"`
	TotalBalance   int32            `db:"total_balance" json:"total_balance"`
	WithdrawAmount *int32           `db:"withdraw_amount" json:"withdraw_amount"`
	WithdrawTime   pgtype.Timestamp `db:"withdraw_time" json:"withdraw_time"`
	CreatedAt      pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt      pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type CreateSaldoIfNotExistsParams struct {
	CardNumber   string `db:"card_number" json:"card_number"`
	TotalBalance int32  `db:"total_balance" json:"total_balance"`
}

type GetActiveSaldosParams struct {
	Column1 string `db:"column1" json:"column_1"`
	Limit   int32  `db:"limit" json:"limit"`
	Offset  int32  `db:"offset" json:"offset"`
}

type GetActiveSaldosRow struct {
	SaldoID        int32            `db:"saldo_id" json:"saldo_id"`
	CardNumber     string           `db:"card_number" json:"card_number"`
	TotalBalance   int32            `db:"total_balance" json:"total_balance"`
	WithdrawAmount *int32           `db:"withdraw_amount" json:"withdraw_amount"`
	WithdrawTime   pgtype.Timestamp `db:"withdraw_time" json:"withdraw_time"`
	CreatedAt      pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt      pgtype.Timestamp `db:"updated_at" json:"updated_at"`
	DeletedAt      pgtype.Timestamp `db:"deleted_at" json:"deleted_at"`
	TotalCount     int64            `db:"total_count" json:"total_count"`
}

type GetMonthlySaldoBalancesRow struct {
	Month        string `db:"month" json:"month"`
	TotalBalance int32  `db:"total_balance" json:"total_balance"`
}

type GetMonthlyTotalSaldoBalanceParams struct {
	Column1 time.Time `db:"column1" json:"column_1"`
	Column2 time.Time `db:"column2" json:"column_2"`
	Column3 time.Time `db:"column3" json:"column_3"`
	Column4 time.Time `db:"column4" json:"column_4"`
}

type GetMonthlyTotalSaldoBalanceRow struct {
	Year         string `db:"year" json:"year"`
	Month        string `db:"month" json:"month"`
	TotalBalance int32  `db:"total_balance" json:"total_balance"`
}

type GetSaldoByIDRow struct {
	SaldoID        int32            `db:"saldo_id" json:"saldo_id"`
	CardNumber     string           `db:"card_number" json:"card_number"`
	TotalBalance   int32            `db:"total_balance" json:"total_balance"`
	WithdrawAmount *int32           `db:"withdraw_amount" json:"withdraw_amount"`
	WithdrawTime   pgtype.Timestamp `db:"withdraw_time" json:"withdraw_time"`
	CreatedAt      pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt      pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type GetSaldosParams struct {
	Column1 string `db:"column1" json:"column_1"`
	Limit   int32  `db:"limit" json:"limit"`
	Offset  int32  `db:"offset" json:"offset"`
}

type GetSaldosRow struct {
	SaldoID        int32            `db:"saldo_id" json:"saldo_id"`
	CardNumber     string           `db:"card_number" json:"card_number"`
	TotalBalance   int32            `db:"total_balance" json:"total_balance"`
	WithdrawAmount *int32           `db:"withdraw_amount" json:"withdraw_amount"`
	WithdrawTime   pgtype.Timestamp `db:"withdraw_time" json:"withdraw_time"`
	CreatedAt      pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt      pgtype.Timestamp `db:"updated_at" json:"updated_at"`
	TotalCount     int64            `db:"total_count" json:"total_count"`
}

type GetTrashedSaldosParams struct {
	Column1 string `db:"column1" json:"column_1"`
	Limit   int32  `db:"limit" json:"limit"`
	Offset  int32  `db:"offset" json:"offset"`
}

type GetTrashedSaldosRow struct {
	SaldoID        int32            `db:"saldo_id" json:"saldo_id"`
	CardNumber     string           `db:"card_number" json:"card_number"`
	TotalBalance   int32            `db:"total_balance" json:"total_balance"`
	WithdrawAmount *int32           `db:"withdraw_amount" json:"withdraw_amount"`
	WithdrawTime   pgtype.Timestamp `db:"withdraw_time" json:"withdraw_time"`
	CreatedAt      pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt      pgtype.Timestamp `db:"updated_at" json:"updated_at"`
	DeletedAt      pgtype.Timestamp `db:"deleted_at" json:"deleted_at"`
	TotalCount     int64            `db:"total_count" json:"total_count"`
}

type GetYearlySaldoBalancesRow struct {
	YYear        string `db:"year" json:"y_year"`
	TotalBalance int64  `db:"total_balance" json:"total_balance"`
}

type GetYearlyTotalSaldoBalancesRow struct {
	Year         string `db:"year" json:"year"`
	TotalBalance int32  `db:"total_balance" json:"total_balance"`
}

type UpdateSaldoParams struct {
	SaldoID      int32  `db:"saldo_id" json:"saldo_id"`
	CardNumber   string `db:"card_number" json:"card_number"`
	TotalBalance int32  `db:"total_balance" json:"total_balance"`
}

type UpdateSaldoRow struct {
	SaldoID        int32            `db:"saldo_id" json:"saldo_id"`
	CardNumber     string           `db:"card_number" json:"card_number"`
	TotalBalance   int32            `db:"total_balance" json:"total_balance"`
	WithdrawAmount *int32           `db:"withdraw_amount" json:"withdraw_amount"`
	WithdrawTime   pgtype.Timestamp `db:"withdraw_time" json:"withdraw_time"`
	CreatedAt      pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt      pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type UpdateSaldoBalanceParams struct {
	CardNumber   string `db:"card_number" json:"card_number"`
	TotalBalance int32  `db:"total_balance" json:"total_balance"`
}

type UpdateSaldoBalanceRow struct {
	SaldoID        int32            `db:"saldo_id" json:"saldo_id"`
	CardNumber     string           `db:"card_number" json:"card_number"`
	TotalBalance   int32            `db:"total_balance" json:"total_balance"`
	WithdrawAmount *int32           `db:"withdraw_amount" json:"withdraw_amount"`
	WithdrawTime   pgtype.Timestamp `db:"withdraw_time" json:"withdraw_time"`
	CreatedAt      pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt      pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type UpdateSaldoBalanceDeltaParams struct {
	CardNumber   string `db:"card_number" json:"card_number"`
	TotalBalance int32  `db:"total_balance" json:"total_balance"`
}

type UpdateSaldoBalanceDeltaRow struct {
	SaldoID        int32            `db:"saldo_id" json:"saldo_id"`
	CardNumber     string           `db:"card_number" json:"card_number"`
	TotalBalance   int32            `db:"total_balance" json:"total_balance"`
	WithdrawAmount *int32           `db:"withdraw_amount" json:"withdraw_amount"`
	WithdrawTime   pgtype.Timestamp `db:"withdraw_time" json:"withdraw_time"`
	CreatedAt      pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt      pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type UpdateSaldoWithdrawParams struct {
	CardNumber     string           `db:"card_number" json:"card_number"`
	WithdrawAmount *int32           `db:"withdraw_amount" json:"withdraw_amount"`
	WithdrawTime   pgtype.Timestamp `db:"withdraw_time" json:"withdraw_time"`
}

type UpdateSaldoWithdrawRow struct {
	SaldoID        int32            `db:"saldo_id" json:"saldo_id"`
	CardNumber     string           `db:"card_number" json:"card_number"`
	TotalBalance   int32            `db:"total_balance" json:"total_balance"`
	WithdrawAmount *int32           `db:"withdraw_amount" json:"withdraw_amount"`
	WithdrawTime   pgtype.Timestamp `db:"withdraw_time" json:"withdraw_time"`
	CreatedAt      pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt      pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}
