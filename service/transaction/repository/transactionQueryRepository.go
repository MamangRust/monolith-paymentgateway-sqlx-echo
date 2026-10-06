package repository

import (
	"context"
	"errors"
	"database/sql"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	sharedErrors "github.com/MamangRust/monolith-payment-gateway-shared/errors"
	transaction_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/transaction_errors/repository"
	"github.com/jmoiron/sqlx"
)

const getTransactions = `SELECT
    transaction_id,
    transaction_no,
    card_number,
    amount,
    payment_method,
    merchant_id,
    transaction_time,
    status,
    created_at,
    updated_at,
    COUNT(*) OVER () AS total_count
FROM transactions
WHERE
    deleted_at IS NULL
    AND (
        $1::TEXT IS NULL
        OR card_number ILIKE '%' || $1 || '%'
        OR payment_method ILIKE '%' || $1 || '%'
        OR status ILIKE '%' || $1 || '%'
    )
ORDER BY transaction_time DESC
LIMIT $2
OFFSET
    $3`

const getActiveTransactions = `SELECT
    transaction_id,
    transaction_no,
    card_number,
    amount,
    payment_method,
    merchant_id,
    transaction_time,
    status,
    created_at,
    updated_at,
    deleted_at,
    COUNT(*) OVER () AS total_count
FROM transactions
WHERE
    deleted_at IS NULL
    AND (
        $1::TEXT IS NULL
        OR card_number ILIKE '%' || $1 || '%'
        OR payment_method ILIKE '%' || $1 || '%'
    )
ORDER BY transaction_time DESC
LIMIT $2
OFFSET
    $3`

const getTrashedTransactions = `SELECT
    transaction_id,
    transaction_no,
    card_number,
    amount,
    payment_method,
    merchant_id,
    transaction_time,
    status,
    created_at,
    updated_at,
    deleted_at,
    COUNT(*) OVER () AS total_count
FROM transactions
WHERE
    deleted_at IS NOT NULL
    AND (
        $1::TEXT IS NULL
        OR card_number ILIKE '%' || $1 || '%'
        OR payment_method ILIKE '%' || $1 || '%'
    )
ORDER BY transaction_time DESC
LIMIT $2
OFFSET
    $3`

const getTransactionsByCardNumber = `SELECT
    transaction_id,
    transaction_no,
    card_number,
    amount,
    payment_method,
    merchant_id,
    transaction_time,
    status,
    created_at,
    updated_at,
    COUNT(*) OVER () AS total_count
FROM transactions
WHERE
    deleted_at IS NULL
    AND card_number = $1
    AND (
        $2::TEXT IS NULL
        OR payment_method ILIKE '%' || $2 || '%'
        OR status ILIKE '%' || $2 || '%'
    )
ORDER BY transaction_time DESC
LIMIT $3
OFFSET
    $4`

const getTransactionByID = `SELECT
    transaction_id,
    transaction_no,
    card_number,
    amount,
    payment_method,
    merchant_id,
    transaction_time,
    status,
    created_at,
    updated_at
FROM transactions
WHERE
    transaction_id = $1
    AND deleted_at IS NULL`

const getTransactionsByMerchantID = `SELECT
    transaction_id,
    transaction_no,
    card_number,
    amount,
    payment_method,
    merchant_id,
    transaction_time,
    status,
    created_at,
    updated_at
FROM transactions
WHERE
    merchant_id = $1
    AND deleted_at IS NULL
ORDER BY transaction_time DESC`

type transactionQueryRepository struct {
	db *sqlx.DB
}

func NewTransactionQueryRepository(db *sqlx.DB) TransactionQueryRepository {
	return &transactionQueryRepository{
		db: db,
	}
}

func (r *transactionQueryRepository) FindAllTransactions(ctx context.Context, req *requests.FindAllTransactions) ([]*db.GetTransactionsRow, error) {
	offset := (req.Page - 1) * req.PageSize

	var transactions []*db.GetTransactionsRow
	err := r.db.SelectContext(ctx, &transactions, getTransactions,
		req.Search,
		int32(req.PageSize),
		int32(offset),
	)

	if err != nil {
		return nil, transaction_errors.ErrFindAllTransactionsFailed.WithInternal(err)
	}

	return transactions, nil
}

func (r *transactionQueryRepository) FindAllTransactionByCardNumber(ctx context.Context, req *requests.FindAllTransactionCardNumber) ([]*db.GetTransactionsByCardNumberRow, error) {
	offset := (req.Page - 1) * req.PageSize

	var transactions []*db.GetTransactionsByCardNumberRow
	err := r.db.SelectContext(ctx, &transactions, getTransactionsByCardNumber,
		req.CardNumber,
		req.Search,
		int32(req.PageSize),
		int32(offset),
	)

	if err != nil {
		return nil, transaction_errors.ErrFindTransactionsByCardNumberFailed.WithInternal(err)
	}

	return transactions, nil
}

func (r *transactionQueryRepository) FindByActive(ctx context.Context, req *requests.FindAllTransactions) ([]*db.GetActiveTransactionsRow, error) {
	offset := (req.Page - 1) * req.PageSize

	var transactions []*db.GetActiveTransactionsRow
	err := r.db.SelectContext(ctx, &transactions, getActiveTransactions,
		req.Search,
		int32(req.PageSize),
		int32(offset),
	)

	if err != nil {
		return nil, transaction_errors.ErrFindActiveTransactionsFailed.WithInternal(err)
	}

	return transactions, nil
}

func (r *transactionQueryRepository) FindByTrashed(ctx context.Context, req *requests.FindAllTransactions) ([]*db.GetTrashedTransactionsRow, error) {
	offset := (req.Page - 1) * req.PageSize

	var transactions []*db.GetTrashedTransactionsRow
	err := r.db.SelectContext(ctx, &transactions, getTrashedTransactions,
		req.Search,
		int32(req.PageSize),
		int32(offset),
	)

	if err != nil {
		return nil, transaction_errors.ErrFindTrashedTransactionsFailed.WithInternal(err)
	}

	return transactions, nil
}

func (r *transactionQueryRepository) FindById(ctx context.Context, transaction_id int) (*db.GetTransactionByIDRow, error) {
	if transaction_id <= 0 {
		return nil, sharedErrors.NewBadRequestError("transaction ID must be greater than zero")
	}

	var transaction db.GetTransactionByIDRow
	err := r.db.GetContext(ctx, &transaction, getTransactionByID, int32(transaction_id))

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, transaction_errors.ErrTransactionNotFound.WithInternal(err)
		}
		return nil, sharedErrors.ErrInternal.WithInternal(err)
	}

	return &transaction, nil
}

func (r *transactionQueryRepository) FindTransactionByMerchantId(ctx context.Context, merchant_id int) ([]*db.GetTransactionsByMerchantIDRow, error) {
	var transactions []*db.GetTransactionsByMerchantIDRow
	err := r.db.SelectContext(ctx, &transactions, getTransactionsByMerchantID, int32(merchant_id))

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, transaction_errors.ErrFindTransactionByMerchantIdFailed.WithInternal(err)
		}
		return nil, sharedErrors.ErrInternal.WithInternal(err)
	}

	return transactions, nil
}
