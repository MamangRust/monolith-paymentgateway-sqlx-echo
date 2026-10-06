package repository

import (
	"context"
	"database/sql"
	"errors"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	sharedErrors "github.com/MamangRust/monolith-payment-gateway-shared/errors"
	withdraw_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/withdraw_errors/repository"
	"github.com/jmoiron/sqlx"
)

const getWithdraws = `SELECT
    withdraw_id,
    withdraw_no,
    card_number,
    withdraw_amount,
    withdraw_time,
    status,
    created_at,
    updated_at,
    COUNT(*) OVER () AS total_count
FROM withdraws
WHERE
    deleted_at IS NULL
    AND (
        $1::TEXT IS NULL
        OR card_number ILIKE '%' || $1 || '%'
        OR withdraw_amount::TEXT ILIKE '%' || $1 || '%'
        OR withdraw_time::TEXT ILIKE '%' || $1 || '%'
        OR status ILIKE '%' || $1 || '%'
    )
ORDER BY withdraw_time DESC
LIMIT $2
OFFSET
    $3`

const getActiveWithdraws = `SELECT
    withdraw_id,
    withdraw_no,
    card_number,
    withdraw_amount,
    withdraw_time,
    status,
    created_at,
    updated_at,
    deleted_at,
    COUNT(*) OVER () AS total_count
FROM withdraws
WHERE
    deleted_at IS NULL
    AND (
        $1::TEXT IS NULL
        OR card_number ILIKE '%' || $1 || '%'
        OR withdraw_amount::TEXT ILIKE '%' || $1 || '%'
        OR withdraw_time::TEXT ILIKE '%' || $1 || '%'
        OR status ILIKE '%' || $1 || '%'
    )
ORDER BY withdraw_time DESC
LIMIT $2
OFFSET
    $3`

const getTrashedWithdraws = `SELECT
    withdraw_id,
    withdraw_no,
    card_number,
    withdraw_amount,
    withdraw_time,
    status,
    created_at,
    updated_at,
    deleted_at,
    COUNT(*) OVER () AS total_count
FROM withdraws
WHERE
    deleted_at IS NOT NULL
    AND (
        $1::TEXT IS NULL
        OR card_number ILIKE '%' || $1 || '%'
        OR withdraw_amount::TEXT ILIKE '%' || $1 || '%'
        OR withdraw_time::TEXT ILIKE '%' || $1 || '%'
        OR status ILIKE '%' || $1 || '%'
    )
ORDER BY withdraw_time DESC
LIMIT $2
OFFSET
    $3`

const getWithdrawByID = `SELECT
    withdraw_id,
    withdraw_no,
    card_number,
    withdraw_amount,
    withdraw_time,
    status,
    created_at,
    updated_at
FROM withdraws
WHERE
    withdraw_id = $1
    AND deleted_at IS NULL`

const getWithdrawsByCardNumber = `SELECT
    withdraw_id,
    withdraw_no,
    card_number,
    withdraw_amount,
    withdraw_time,
    status,
    created_at,
    updated_at,
    COUNT(*) OVER () AS total_count
FROM withdraws
WHERE
    deleted_at IS NULL
    AND card_number = $1
    AND (
        $2::TEXT IS NULL
        OR CAST(withdraw_amount AS TEXT) ILIKE '%' || $2 || '%'
        OR TO_CHAR(
            withdraw_time,
            'YYYY-MM-DD HH24:MI:SS'
        ) ILIKE '%' || $2 || '%'
        OR status ILIKE '%' || $2 || '%'
    )
ORDER BY withdraw_time DESC
LIMIT $3
OFFSET
    $4`

type withdrawQueryRepository struct {
	db *sqlx.DB
}

func NewWithdrawQueryRepository(db *sqlx.DB) WithdrawQueryRepository {
	return &withdrawQueryRepository{
		db: db,
	}
}

func (r *withdrawQueryRepository) FindAll(ctx context.Context, req *requests.FindAllWithdraws) ([]*db.GetWithdrawsRow, error) {
	offset := (req.Page - 1) * req.PageSize

	var withdraws []*db.GetWithdrawsRow
	err := r.db.SelectContext(ctx, &withdraws, getWithdraws,
		req.Search,
		int32(req.PageSize),
		int32(offset),
	)

	if err != nil {
		return nil, withdraw_errors.ErrFindAllWithdrawsFailed.WithInternal(err)
	}

	return withdraws, nil

}

func (r *withdrawQueryRepository) FindByActive(ctx context.Context, req *requests.FindAllWithdraws) ([]*db.GetActiveWithdrawsRow, error) {
	offset := (req.Page - 1) * req.PageSize

	var withdraws []*db.GetActiveWithdrawsRow
	err := r.db.SelectContext(ctx, &withdraws, getActiveWithdraws,
		req.Search,
		int32(req.PageSize),
		int32(offset),
	)

	if err != nil {
		return nil, withdraw_errors.ErrFindActiveWithdrawsFailed.WithInternal(err)
	}

	return withdraws, nil
}

func (r *withdrawQueryRepository) FindByTrashed(ctx context.Context, req *requests.FindAllWithdraws) ([]*db.GetTrashedWithdrawsRow, error) {
	offset := (req.Page - 1) * req.PageSize

	var withdraws []*db.GetTrashedWithdrawsRow
	err := r.db.SelectContext(ctx, &withdraws, getTrashedWithdraws,
		req.Search,
		int32(req.PageSize),
		int32(offset),
	)

	if err != nil {
		return nil, withdraw_errors.ErrFindTrashedWithdrawsFailed.WithInternal(err)
	}

	return withdraws, nil
}

func (r *withdrawQueryRepository) FindAllByCardNumber(ctx context.Context, req *requests.FindAllWithdrawCardNumber) ([]*db.GetWithdrawsByCardNumberRow, error) {
	offset := (req.Page - 1) * req.PageSize

	var withdraws []*db.GetWithdrawsByCardNumberRow
	err := r.db.SelectContext(ctx, &withdraws, getWithdrawsByCardNumber,
		req.CardNumber,
		req.Search,
		int32(req.PageSize),
		int32(offset),
	)

	if err != nil {
		return nil, withdraw_errors.ErrFindWithdrawsByCardNumberFailed.WithInternal(err)
	}

	return withdraws, nil

}

func (r *withdrawQueryRepository) FindById(ctx context.Context, id int) (*db.GetWithdrawByIDRow, error) {
	if id <= 0 {
		return nil, sharedErrors.NewBadRequestError("withdraw ID must be greater than zero")
	}

	var withdraw db.GetWithdrawByIDRow
	err := r.db.GetContext(ctx, &withdraw, getWithdrawByID, int32(id))

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, withdraw_errors.ErrWithdrawNotFound.WithInternal(err)
		}
		return nil, sharedErrors.ErrInternal.WithInternal(err)
	}

	return &withdraw, nil
}
