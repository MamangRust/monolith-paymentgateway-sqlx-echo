package repository

import (
	"context"
	"database/sql"
	"errors"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	sharedErrors "github.com/MamangRust/monolith-payment-gateway-shared/errors"
	merchant_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/merchant_errors/repository"
	"github.com/jmoiron/sqlx"
)

const getMerchants = `SELECT
    merchant_id,
    merchant_no,
    name,
    api_key,
    user_id,
    status,
    created_at,
    updated_at,
    COUNT(*) OVER () AS total_count
FROM merchants
WHERE
    deleted_at IS NULL
    AND (
        $1::TEXT IS NULL
        OR name ILIKE '%' || $1 || '%'
        OR api_key ILIKE '%' || $1 || '%'
        OR status ILIKE '%' || $1 || '%'
    )
ORDER BY merchant_id
LIMIT $2
OFFSET
    $3`

const getActiveMerchants = `SELECT
    merchant_id,
    merchant_no,
    name,
    api_key,
    user_id,
    status,
    created_at,
    updated_at,
    deleted_at,
    COUNT(*) OVER () AS total_count
FROM merchants
WHERE
    deleted_at IS NULL
    AND (
        $1::TEXT IS NULL
        OR name ILIKE '%' || $1 || '%'
        OR api_key ILIKE '%' || $1 || '%'
        OR status ILIKE '%' || $1 || '%'
    )
ORDER BY merchant_id
LIMIT $2
OFFSET
    $3`

const getTrashedMerchants = `SELECT
    merchant_id,
    merchant_no,
    name,
    api_key,
    user_id,
    status,
    created_at,
    updated_at,
    deleted_at,
    COUNT(*) OVER () AS total_count
FROM merchants
WHERE
    deleted_at IS NOT NULL
    AND (
        $1::TEXT IS NULL
        OR name ILIKE '%' || $1 || '%'
        OR api_key ILIKE '%' || $1 || '%'
        OR status ILIKE '%' || $1 || '%'
    )
ORDER BY merchant_id
LIMIT $2
OFFSET
    $3`

const getMerchantByID = `SELECT
    merchant_id,
    merchant_no,
    name,
    api_key,
    user_id,
    status,
    created_at,
    updated_at
FROM merchants
WHERE
    merchant_id = $1
    AND deleted_at IS NULL`

const getMerchantByApiKey = `SELECT
    merchant_id,
    merchant_no,
    name,
    api_key,
    user_id,
    status,
    created_at,
    updated_at
FROM merchants
WHERE
    api_key = $1
    AND deleted_at IS NULL`

const getMerchantByName = `SELECT
    merchant_id,
    merchant_no,
    name,
    api_key,
    user_id,
    status,
    created_at,
    updated_at
FROM merchants
WHERE
    name = $1
    AND deleted_at IS NULL`

const getMerchantsByUserID = `SELECT
    merchant_id,
    merchant_no,
    name,
    api_key,
    user_id,
    status,
    created_at,
    updated_at
FROM merchants
WHERE
    user_id = $1
    AND deleted_at IS NULL`

type merchantQueryRepository struct {
	db *sqlx.DB
}

func NewMerchantQueryRepository(db *sqlx.DB) MerchantQueryRepository {
	return &merchantQueryRepository{
		db: db,
	}
}

func (r *merchantQueryRepository) FindAllMerchants(ctx context.Context, req *requests.FindAllMerchants) ([]*db.GetMerchantsRow, error) {
	offset := (req.Page - 1) * req.PageSize

	var merchants []*db.GetMerchantsRow
	err := r.db.SelectContext(ctx, &merchants, getMerchants,
		req.Search,
		int32(req.PageSize),
		int32(offset),
	)

	if err != nil {
		return nil, merchant_errors.ErrFindAllMerchantsFailed.WithInternal(err)
	}

	return merchants, nil
}

func (r *merchantQueryRepository) FindByActive(ctx context.Context, req *requests.FindAllMerchants) ([]*db.GetActiveMerchantsRow, error) {
	offset := (req.Page - 1) * req.PageSize

	var merchants []*db.GetActiveMerchantsRow
	err := r.db.SelectContext(ctx, &merchants, getActiveMerchants,
		req.Search,
		int32(req.PageSize),
		int32(offset),
	)

	if err != nil {
		return nil, merchant_errors.ErrFindActiveMerchantsFailed.WithInternal(err)
	}

	return merchants, nil
}

func (r *merchantQueryRepository) FindByTrashed(ctx context.Context, req *requests.FindAllMerchants) ([]*db.GetTrashedMerchantsRow, error) {
	offset := (req.Page - 1) * req.PageSize

	var merchants []*db.GetTrashedMerchantsRow
	err := r.db.SelectContext(ctx, &merchants, getTrashedMerchants,
		req.Search,
		int32(req.PageSize),
		int32(offset),
	)

	if err != nil {
		return nil, merchant_errors.ErrFindTrashedMerchantsFailed.WithInternal(err)
	}

	return merchants, nil
}

func (r *merchantQueryRepository) FindByMerchantId(ctx context.Context, merchant_id int) (*db.GetMerchantByIDRow, error) {
	if merchant_id <= 0 {
		return nil, sharedErrors.NewBadRequestError("merchant ID must be greater than zero")
	}

	var merchant db.GetMerchantByIDRow
	err := r.db.GetContext(ctx, &merchant, getMerchantByID, int32(merchant_id))

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, merchant_errors.ErrMerchantNotFound.WithInternal(err)
		}
		return nil, sharedErrors.ErrInternal.WithInternal(err)
	}

	return &merchant, nil
}

func (r *merchantQueryRepository) FindByApiKey(ctx context.Context, api_key string) (*db.GetMerchantByApiKeyRow, error) {
	var merchant db.GetMerchantByApiKeyRow
	err := r.db.GetContext(ctx, &merchant, getMerchantByApiKey, api_key)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, merchant_errors.ErrMerchantNotFound.WithInternal(err)
		}
		return nil, sharedErrors.ErrInternal.WithInternal(err)
	}

	return &merchant, nil
}

func (r *merchantQueryRepository) FindByName(ctx context.Context, name string) (*db.GetMerchantByNameRow, error) {
	var merchant db.GetMerchantByNameRow
	err := r.db.GetContext(ctx, &merchant, getMerchantByName, name)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, merchant_errors.ErrMerchantNotFound.WithInternal(err)
		}
		return nil, sharedErrors.ErrInternal.WithInternal(err)
	}

	return &merchant, nil
}

func (r *merchantQueryRepository) FindByMerchantUserId(ctx context.Context, user_id int) ([]*db.GetMerchantsByUserIDRow, error) {
	var merchants []*db.GetMerchantsByUserIDRow
	err := r.db.SelectContext(ctx, &merchants, getMerchantsByUserID, int32(user_id))

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, merchant_errors.ErrMerchantNotFound.WithInternal(err)
		}
		return nil, sharedErrors.ErrInternal.WithInternal(err)
	}

	return merchants, nil
}
