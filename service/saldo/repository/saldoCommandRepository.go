package repository

import (
	"context"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	sharedErrors "github.com/MamangRust/monolith-payment-gateway-shared/errors"
	saldo_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/saldo_errors/repository"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jmoiron/sqlx"
)

type saldoCommandRepository struct {
	db *sqlx.DB
}

func NewSaldoCommandRepository(db *sqlx.DB) SaldoCommandRepository {
	return &saldoCommandRepository{
		db: db,
	}
}

const createSaldo = `INSERT INTO
    saldos (
        card_number,
        total_balance,
        created_at,
        updated_at
    )
VALUES (
        $1,
        $2,
        current_timestamp,
        current_timestamp
    )
ON CONFLICT (card_number) WHERE deleted_at IS NULL
DO UPDATE SET
    total_balance = EXCLUDED.total_balance,
    updated_at = current_timestamp
RETURNING
    saldo_id,
    card_number,
    total_balance,
    withdraw_amount,
    withdraw_time,
    created_at,
    updated_at`

const createSaldoIfNotExists = `INSERT INTO saldos (
    card_number,
    total_balance,
    created_at,
    updated_at
)
VALUES ($1, $2, current_timestamp, current_timestamp)
ON CONFLICT (card_number) WHERE deleted_at IS NULL
DO NOTHING`

const updateSaldo = `UPDATE saldos
SET
    card_number = $2,
    total_balance = $3,
    updated_at = current_timestamp
WHERE
    saldo_id = $1
    AND deleted_at IS NULL
RETURNING
    saldo_id,
    card_number,
    total_balance,
    withdraw_amount,
    withdraw_time,
    created_at,
    updated_at`

const updateSaldoBalance = `UPDATE saldos
SET
    total_balance = $2,
    updated_at = current_timestamp
WHERE
    card_number = $1
    AND deleted_at IS NULL
RETURNING
    saldo_id,
    card_number,
    total_balance,
    withdraw_amount,
    withdraw_time,
    created_at,
    updated_at`

const updateSaldoWithdraw = `UPDATE saldos
SET
    withdraw_amount = $2,
    total_balance = total_balance - $2,
    withdraw_time = $3,
    updated_at = current_timestamp
WHERE
    card_number = $1
    AND deleted_at IS NULL
    AND total_balance >= $2
RETURNING
    saldo_id,
    card_number,
    total_balance,
    withdraw_amount,
    withdraw_time,
    created_at,
    updated_at`

const trashSaldo = `UPDATE saldos
SET
    deleted_at = current_timestamp
WHERE
    saldo_id = $1
    AND deleted_at IS NULL
RETURNING
    saldo_id,
    card_number,
    total_balance,
    withdraw_amount,
    withdraw_time,
    created_at,
    updated_at,
    deleted_at`

const restoreSaldo = `UPDATE saldos
SET
    deleted_at = NULL
WHERE
    saldo_id = $1
    AND deleted_at IS NOT NULL
RETURNING
    saldo_id,
    card_number,
    total_balance,
    withdraw_amount,
    withdraw_time,
    created_at,
    updated_at,
    deleted_at`

const deleteSaldoPermanently = `DELETE FROM saldos WHERE saldo_id = $1 AND deleted_at IS NOT NULL`

const restoreAllSaldos = `UPDATE saldos
SET
    deleted_at = NULL
WHERE
    deleted_at IS NOT NULL`

const deleteAllPermanentSaldos = `DELETE FROM saldos WHERE deleted_at IS NOT NULL`

func (r *saldoCommandRepository) CreateSaldo(ctx context.Context, request *requests.CreateSaldoRequest) (*db.CreateSaldoRow, error) {
	var saldo db.CreateSaldoRow
	err := r.db.GetContext(ctx, &saldo, createSaldo,
		request.CardNumber,
		int32(request.TotalBalance),
	)

	if err != nil {
		return nil, sharedErrors.ErrConstraintOrFailed(err, "Saldo", "create saldo")
	}

	return &saldo, nil
}

func (r *saldoCommandRepository) CreateSaldoIfNotExists(ctx context.Context, request *requests.CreateSaldoRequest) error {
	if _, err := r.db.ExecContext(ctx, createSaldoIfNotExists,
		request.CardNumber,
		int32(request.TotalBalance),
	); err != nil {
		return sharedErrors.ErrConstraintOrFailed(err, "Saldo", "create saldo")
	}
	return nil
}

func (r *saldoCommandRepository) UpdateSaldo(ctx context.Context, request *requests.UpdateSaldoRequest) (*db.UpdateSaldoRow, error) {
	var saldo db.UpdateSaldoRow
	err := r.db.GetContext(ctx, &saldo, updateSaldo,
		int32(*request.SaldoID),
		request.CardNumber,
		int32(request.TotalBalance),
	)

	if err != nil {
		return nil, sharedErrors.ErrNoRowsOrFailed(err, "Saldo", "update saldo")
	}

	return &saldo, nil
}

func (r *saldoCommandRepository) UpdateSaldoBalance(ctx context.Context, request *requests.UpdateSaldoBalance) (*db.UpdateSaldoBalanceRow, error) {
	var saldo db.UpdateSaldoBalanceRow
	err := r.db.GetContext(ctx, &saldo, updateSaldoBalance,
		request.CardNumber,
		int32(request.TotalBalance),
	)

	if err != nil {
		return nil, sharedErrors.ErrNoRowsOrFailed(err, "Saldo", "update saldo balance")
	}

	return &saldo, nil
}

func (r *saldoCommandRepository) UpdateSaldoWithdraw(ctx context.Context, request *requests.UpdateSaldoWithdraw) (*db.UpdateSaldoWithdrawRow, error) {
	var withdrawAmount *int32
	if request.WithdrawAmount != nil {
		amount := int32(*request.WithdrawAmount)
		withdrawAmount = &amount
	}

	var withdrawTime pgtype.Timestamp
	if request.WithdrawTime != nil {
		withdrawTime = pgtype.Timestamp{
			Time:  *request.WithdrawTime,
			Valid: true,
		}
	}

	var saldo db.UpdateSaldoWithdrawRow
	err := r.db.GetContext(ctx, &saldo, updateSaldoWithdraw,
		request.CardNumber,
		withdrawAmount,
		withdrawTime,
	)

	if err != nil {
		return nil, sharedErrors.ErrNoRowsOrFailed(err, "Saldo", "update saldo withdraw")
	}

	return &saldo, nil
}

func (r *saldoCommandRepository) TrashedSaldo(ctx context.Context, saldo_id int) (*db.Saldo, error) {
	var saldo db.Saldo
	err := r.db.GetContext(ctx, &saldo, trashSaldo, int32(saldo_id))
	if err != nil {
		return nil, sharedErrors.ErrNoRowsOrFailed(err, "Saldo", "trash saldo")
	}
	return &saldo, nil
}

func (r *saldoCommandRepository) RestoreSaldo(ctx context.Context, saldo_id int) (*db.Saldo, error) {
	var saldo db.Saldo
	err := r.db.GetContext(ctx, &saldo, restoreSaldo, int32(saldo_id))
	if err != nil {
		return nil, sharedErrors.ErrNoRowsOrFailed(err, "Saldo", "restore saldo")
	}
	return &saldo, nil
}

func (r *saldoCommandRepository) DeleteSaldoPermanent(ctx context.Context, saldo_id int) (bool, error) {
	if _, err := r.db.ExecContext(ctx, deleteSaldoPermanently, int32(saldo_id)); err != nil {
		return false, sharedErrors.ErrNoRowsOrFailed(err, "Saldo", "delete saldo")
	}
	return true, nil
}

func (r *saldoCommandRepository) RestoreAllSaldo(ctx context.Context) (bool, error) {
	if _, err := r.db.ExecContext(ctx, restoreAllSaldos); err != nil {
		return false, saldo_errors.ErrRestoreAllSaldosFailed.WithInternal(err)
	}

	return true, nil
}

func (r *saldoCommandRepository) DeleteAllSaldoPermanent(ctx context.Context) (bool, error) {
	if _, err := r.db.ExecContext(ctx, deleteAllPermanentSaldos); err != nil {
		return false, saldo_errors.ErrDeleteAllSaldosPermanentFailed.WithInternal(err)
	}

	return true, nil
}
