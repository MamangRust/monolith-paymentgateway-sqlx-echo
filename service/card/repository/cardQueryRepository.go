package repository

import (
	"context"
	"database/sql"
	"errors"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	sharedErrors "github.com/MamangRust/monolith-payment-gateway-shared/errors"
	card_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/card_errors/repository"
	"github.com/jmoiron/sqlx"
)

const getCards = `SELECT
    card_id,
    user_id,
    card_number,
    card_type,
    expire_date,
    cvv,
    card_provider,
    created_at,
    updated_at,
    COUNT(*) OVER () AS total_count
FROM cards
WHERE
    deleted_at IS NULL
    AND (
        $1::TEXT IS NULL
        OR card_number ILIKE '%' || $1 || '%'
        OR card_type ILIKE '%' || $1 || '%'
        OR card_provider ILIKE '%' || $1 || '%'
    )
ORDER BY card_id
LIMIT $2
OFFSET
    $3`

const getCardByID = `SELECT
    card_id,
    user_id,
    card_number,
    card_type,
    expire_date,
    cvv,
    card_provider,
    status,
    credit_limit,
    outstanding_balance,
    reward_points,
    created_at,
    updated_at
FROM cards
WHERE
    card_id = $1
    AND deleted_at IS NULL`

const getUserEmailByCardNumber = `SELECT c.card_id, u.email, c.user_id, c.card_number, c.card_type, c.expire_date, c.cvv, c.card_provider, c.created_at, c.updated_at
FROM cards c
    JOIN users u ON u.user_id = c.user_id
WHERE
    c.card_number = $1
    AND c.deleted_at IS NULL
    AND u.deleted_at IS NULL`

const getActiveCardsWithCount = `SELECT
    card_id,
    user_id,
    card_number,
    card_type,
    expire_date,
    cvv,
    card_provider,
    created_at,
    updated_at,
    deleted_at,
    COUNT(*) OVER () AS total_count
FROM cards
WHERE
    deleted_at IS NULL
    AND (
        $1::TEXT IS NULL
        OR card_number ILIKE '%' || $1 || '%'
        OR card_type ILIKE '%' || $1 || '%'
        OR card_provider ILIKE '%' || $1 || '%'
    )
ORDER BY card_id
LIMIT $2
OFFSET
    $3`

const getTrashedCardsWithCount = `SELECT
    card_id,
    user_id,
    card_number,
    card_type,
    expire_date,
    cvv,
    card_provider,
    created_at,
    updated_at,
    deleted_at,
    COUNT(*) OVER () AS total_count
FROM cards
WHERE
    deleted_at IS NOT NULL
    AND (
        $1::TEXT IS NULL
        OR card_number ILIKE '%' || $1 || '%'
        OR card_type ILIKE '%' || $1 || '%'
        OR card_provider ILIKE '%' || $1 || '%'
    )
ORDER BY card_id
LIMIT $2
OFFSET
    $3`

const getCardByUserID = `SELECT
    card_id,
    user_id,
    card_number,
    card_type,
    expire_date,
    cvv,
    card_provider,
    status,
    credit_limit,
    outstanding_balance,
    reward_points,
    created_at,
    updated_at
FROM cards
WHERE
    user_id = $1
    AND deleted_at IS NULL
LIMIT 1`

const getCardByCardNumber = `SELECT
    card_id,
    user_id,
    card_number,
    card_type,
    expire_date,
    cvv,
    card_provider,
    status,
    credit_limit,
    outstanding_balance,
    reward_points,
    created_at,
    updated_at
FROM cards
WHERE
    card_number = $1
    AND deleted_at IS NULL`

type cardQueryRepository struct {
	db *sqlx.DB
}

func NewCardQueryRepository(db *sqlx.DB) CardQueryRepository {
	return &cardQueryRepository{
		db: db,
	}
}

func (r *cardQueryRepository) FindAllCards(ctx context.Context, req *requests.FindAllCards) ([]*db.GetCardsRow, error) {
	page, pageSize := normalizePagination(req.Page, req.PageSize)
	offset := (page - 1) * pageSize

	var cards []*db.GetCardsRow
	err := r.db.SelectContext(ctx, &cards, getCards,
		req.Search,
		int32(pageSize),
		int32(offset),
	)

	if err != nil {
		return nil, card_errors.ErrFindAllCardsFailed.WithInternal(err)
	}

	return cards, nil
}

func (r *cardQueryRepository) FindByActive(ctx context.Context, req *requests.FindAllCards) ([]*db.GetActiveCardsWithCountRow, error) {
	page, pageSize := normalizePagination(req.Page, req.PageSize)
	offset := (page - 1) * pageSize

	var res []*db.GetActiveCardsWithCountRow
	err := r.db.SelectContext(ctx, &res, getActiveCardsWithCount,
		req.Search,
		int32(pageSize),
		int32(offset),
	)

	if err != nil {
		return nil, card_errors.ErrFindActiveCardsFailed.WithInternal(err)
	}

	return res, nil
}

func (r *cardQueryRepository) FindByTrashed(ctx context.Context, req *requests.FindAllCards) ([]*db.GetTrashedCardsWithCountRow, error) {
	page, pageSize := normalizePagination(req.Page, req.PageSize)
	offset := (page - 1) * pageSize

	var res []*db.GetTrashedCardsWithCountRow
	err := r.db.SelectContext(ctx, &res, getTrashedCardsWithCount,
		req.Search,
		int32(pageSize),
		int32(offset),
	)

	if err != nil {
		return nil, card_errors.ErrFindTrashedCardsFailed.WithInternal(err)
	}

	return res, nil
}

func normalizePagination(page, pageSize int) (int, int) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	return page, pageSize
}

func (r *cardQueryRepository) FindById(ctx context.Context, card_id int) (*db.GetCardByIDRow, error) {
	if card_id <= 0 {
		return nil, sharedErrors.NewBadRequestError("card ID must be greater than zero")
	}

	var res db.GetCardByIDRow
	err := r.db.GetContext(ctx, &res, getCardByID, int32(card_id))

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, card_errors.ErrCardNotFound.WithInternal(err)
		}
		return nil, sharedErrors.ErrInternal.WithInternal(err)
	}

	return &res, nil
}

func (r *cardQueryRepository) FindCardByUserId(ctx context.Context, user_id int) (*db.GetCardByUserIDRow, error) {
	var res db.GetCardByUserIDRow
	err := r.db.GetContext(ctx, &res, getCardByUserID, int32(user_id))

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, card_errors.ErrCardNotFound.WithInternal(err)
		}
		return nil, sharedErrors.ErrInternal.WithInternal(err)
	}

	return &res, nil
}

func (r *cardQueryRepository) FindCardByCardNumber(ctx context.Context, card_number string) (*db.GetCardByCardNumberRow, error) {
	var res db.GetCardByCardNumberRow
	err := r.db.GetContext(ctx, &res, getCardByCardNumber, card_number)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, card_errors.ErrCardNotFound.WithInternal(err)
		}
		return nil, sharedErrors.ErrInternal.WithInternal(err)
	}

	return &res, nil
}

func (r *cardQueryRepository) FindUserCardByCardNumber(ctx context.Context, card_number string) (*db.GetUserEmailByCardNumberRow, error) {
	var res db.GetUserEmailByCardNumberRow
	err := r.db.GetContext(ctx, &res, getUserEmailByCardNumber, card_number)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, card_errors.ErrCardNotFound.WithInternal(err)
		}
		return nil, sharedErrors.ErrInternal.WithInternal(err)
	}

	return &res, nil
}
