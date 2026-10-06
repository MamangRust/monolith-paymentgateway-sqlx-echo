package repository

import (
	"context"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	merchant_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/merchant_errors/repository"
	"github.com/jmoiron/sqlx"
)

const findAllTransactions = `SELECT
    t.transaction_id,
    t.card_number,
    t.amount,
    t.payment_method,
    t.merchant_id,
    m.name AS merchant_name,
    t.transaction_time,
    t.created_at,
    t.updated_at,
    t.deleted_at,
    COUNT(*) OVER () AS total_count
FROM transactions t
    JOIN merchants m ON t.merchant_id = m.merchant_id
WHERE
    t.deleted_at IS NULL
    AND (
        $1::TEXT IS NULL
        OR t.card_number ILIKE '%' || $1 || '%'
        OR t.payment_method ILIKE '%' || $1 || '%'
    )
ORDER BY t.transaction_time DESC
LIMIT $2
OFFSET
    $3`

const findAllTransactionsByMerchant = `SELECT
    t.transaction_id,
    t.card_number,
    t.amount,
    t.payment_method,
    t.merchant_id,
    m.name AS merchant_name,
    t.transaction_time,
    t.created_at,
    t.updated_at,
    t.deleted_at,
    COUNT(*) OVER () AS total_count
FROM transactions t
    JOIN merchants m ON t.merchant_id = m.merchant_id
WHERE
    t.deleted_at IS NULL
    AND t.merchant_id = $1
    AND (
        $2::TEXT IS NULL
        OR t.card_number ILIKE '%' || $2 || '%'
        OR t.payment_method ILIKE '%' || $2 || '%'
    )
ORDER BY t.transaction_time DESC
LIMIT $3
OFFSET
    $4`

const findAllTransactionsByApikey = `SELECT
    t.transaction_id,
    t.card_number,
    t.amount,
    t.payment_method,
    t.merchant_id,
    m.name AS merchant_name,
    t.transaction_time,
    t.created_at,
    t.updated_at,
    t.deleted_at,
    COUNT(*) OVER () AS total_count
FROM transactions t
    JOIN merchants m ON t.merchant_id = m.merchant_id
WHERE
    t.deleted_at IS NULL
    AND m.api_key = $1
    AND (
        $2::TEXT IS NULL
        OR t.card_number ILIKE '%' || $2 || '%'
        OR t.payment_method ILIKE '%' || $2 || '%'
    )
ORDER BY t.transaction_time DESC
LIMIT $3
OFFSET
    $4`

type merchantTransactionRepository struct {
	db *sqlx.DB
}

func NewMerchantTransactionRepository(db *sqlx.DB) MerchantTransactionRepository {
	return &merchantTransactionRepository{
		db: db,
	}
}

func (r *merchantTransactionRepository) FindAllTransactions(ctx context.Context, req *requests.FindAllMerchantTransactions) ([]*db.FindAllTransactionsRow, error) {
	offset := (req.Page - 1) * req.PageSize

	var transactions []*db.FindAllTransactionsRow
	err := r.db.SelectContext(ctx, &transactions, findAllTransactions,
		req.Search,
		int32(req.PageSize),
		int32(offset),
	)

	if err != nil {
		return nil, merchant_errors.ErrFindAllTransactionsFailed.WithInternal(err)
	}

	return transactions, nil
}

func (r *merchantTransactionRepository) FindAllTransactionsByMerchant(ctx context.Context, req *requests.FindAllMerchantTransactionsById) ([]*db.FindAllTransactionsByMerchantRow, error) {
	offset := (req.Page - 1) * req.PageSize

	var transactions []*db.FindAllTransactionsByMerchantRow
	err := r.db.SelectContext(ctx, &transactions, findAllTransactionsByMerchant,
		int32(req.MerchantID),
		req.Search,
		int32(req.PageSize),
		int32(offset),
	)

	if err != nil {
		return nil, merchant_errors.ErrFindAllTransactionsByMerchantFailed.WithInternal(err)
	}

	return transactions, nil
}

func (r *merchantTransactionRepository) FindAllTransactionsByApikey(ctx context.Context, req *requests.FindAllMerchantTransactionsByApiKey) ([]*db.FindAllTransactionsByApikeyRow, error) {
	offset := (req.Page - 1) * req.PageSize

	var transactions []*db.FindAllTransactionsByApikeyRow
	err := r.db.SelectContext(ctx, &transactions, findAllTransactionsByApikey,
		req.ApiKey,
		req.Search,
		int32(req.PageSize),
		int32(offset),
	)

	if err != nil {
		return nil, merchant_errors.ErrFindAllTransactionsByApiKeyFailed.WithInternal(err)
	}

	return transactions, nil
}
