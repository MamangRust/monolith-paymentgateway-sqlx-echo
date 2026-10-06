package repository

import (
	"context"
	"database/sql"
	"errors"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	sharedErrors "github.com/MamangRust/monolith-payment-gateway-shared/errors"
	saldo_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/saldo_errors/repository"
	"github.com/jmoiron/sqlx"
)

type saldoQueryRepository struct {
	db *sqlx.DB
}

func NewSaldoQueryRepository(db *sqlx.DB) SaldoQueryRepository {
	return &saldoQueryRepository{
		db: db,
	}
}

const getSaldos = `SELECT
    saldo_id,
    card_number,
    total_balance,
    withdraw_amount,
    withdraw_time,
    created_at,
    updated_at,
    COUNT(*) OVER () AS total_count
FROM saldos
WHERE
    deleted_at IS NULL
    AND (
        $1::TEXT IS NULL
        OR card_number ILIKE '%' || $1 || '%'
    )
ORDER BY saldo_id
LIMIT $2
OFFSET
    $3`

const getActiveSaldos = `SELECT
    saldo_id,
    card_number,
    total_balance,
    withdraw_amount,
    withdraw_time,
    created_at,
    updated_at,
    deleted_at,
    COUNT(*) OVER () AS total_count
FROM saldos
WHERE
    deleted_at IS NULL
    AND (
        $1::TEXT IS NULL
        OR card_number ILIKE '%' || $1 || '%'
    )
ORDER BY saldo_id
LIMIT $2
OFFSET
    $3`

const getTrashedSaldos = `SELECT
    saldo_id,
    card_number,
    total_balance,
    withdraw_amount,
    withdraw_time,
    created_at,
    updated_at,
    deleted_at,
    COUNT(*) OVER () AS total_count
FROM saldos
WHERE
    deleted_at IS NOT NULL
    AND (
        $1::TEXT IS NULL
        OR card_number ILIKE '%' || $1 || '%'
    )
ORDER BY saldo_id
LIMIT $2
OFFSET
    $3`

const getSaldoByCardNumber = `SELECT saldo_id, card_number, total_balance, withdraw_amount, withdraw_time, created_at, updated_at, deleted_at
FROM saldos
WHERE card_number = $1
  AND deleted_at IS NULL
ORDER BY saldo_id DESC
LIMIT 1`

const getSaldoByID = `SELECT
    saldo_id,
    card_number,
    total_balance,
    withdraw_amount,
    withdraw_time,
    created_at,
    updated_at
FROM saldos
WHERE
    saldo_id = $1
    AND deleted_at IS NULL`

func (r *saldoQueryRepository) FindAllSaldos(ctx context.Context, req *requests.FindAllSaldos) ([]*db.GetSaldosRow, error) {
	offset := (req.Page - 1) * req.PageSize

	var saldos []*db.GetSaldosRow
	err := r.db.SelectContext(ctx, &saldos, getSaldos,
		req.Search,
		int32(req.PageSize),
		int32(offset),
	)

	if err != nil {
		return nil, saldo_errors.ErrFindAllSaldosFailed.WithInternal(err)
	}

	return saldos, nil
}

func (r *saldoQueryRepository) FindByActive(ctx context.Context, req *requests.FindAllSaldos) ([]*db.GetActiveSaldosRow, error) {
	offset := (req.Page - 1) * req.PageSize

	var saldos []*db.GetActiveSaldosRow
	err := r.db.SelectContext(ctx, &saldos, getActiveSaldos,
		req.Search,
		int32(req.PageSize),
		int32(offset),
	)

	if err != nil {
		return nil, saldo_errors.ErrFindActiveSaldosFailed.WithInternal(err)
	}

	return saldos, nil
}

func (r *saldoQueryRepository) FindByTrashed(ctx context.Context, req *requests.FindAllSaldos) ([]*db.GetTrashedSaldosRow, error) {
	offset := (req.Page - 1) * req.PageSize

	var saldos []*db.GetTrashedSaldosRow
	err := r.db.SelectContext(ctx, &saldos, getTrashedSaldos,
		req.Search,
		int32(req.PageSize),
		int32(offset),
	)

	if err != nil {
		return nil, saldo_errors.ErrFindTrashedSaldosFailed.WithInternal(err)
	}

	return saldos, nil
}

func (r *saldoQueryRepository) FindByCardNumber(ctx context.Context, card_number string) (*db.Saldo, error) {
	var saldo db.Saldo
	err := r.db.GetContext(ctx, &saldo, getSaldoByCardNumber, card_number)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, saldo_errors.ErrSaldoNotFound.WithInternal(err)
		}
		return nil, sharedErrors.ErrInternal.WithInternal(err)
	}

	return &saldo, nil
}

func (r *saldoQueryRepository) FindById(ctx context.Context, saldo_id int) (*db.GetSaldoByIDRow, error) {
	if saldo_id <= 0 {
		return nil, sharedErrors.NewBadRequestError("saldo ID must be greater than zero")
	}

	var saldo db.GetSaldoByIDRow
	err := r.db.GetContext(ctx, &saldo, getSaldoByID, int32(saldo_id))

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, saldo_errors.ErrSaldoNotFound.WithInternal(err)
		}
		return nil, sharedErrors.ErrInternal.WithInternal(err)
	}

	return &saldo, nil
}
