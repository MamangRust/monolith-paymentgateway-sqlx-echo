package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	sharedErrors "github.com/MamangRust/monolith-payment-gateway-shared/errors"
	transaction_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/transaction_errors/repository"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jmoiron/sqlx"
)

const authorizeTransactionAtomic = `WITH locked_card AS MATERIALIZED (
    SELECT c.card_id, c.card_number, c.card_type, c.outstanding_balance, c.credit_limit
    FROM cards AS c
    WHERE c.card_number = $1
      AND c.deleted_at IS NULL
      AND c.status = 'active'
    FOR UPDATE
), locked_saldo AS MATERIALIZED (
    SELECT card_number, total_balance
    FROM saldos
    WHERE card_number = $1
      AND deleted_at IS NULL
    FOR UPDATE
), debit_hold AS (
    UPDATE saldos s
    SET total_balance = s.total_balance - $2,
        updated_at = current_timestamp
    FROM locked_card c, locked_saldo ls
    WHERE c.card_type <> 'credit'
      AND $2 > 0
      AND ls.total_balance >= $2
      AND s.card_number = ls.card_number
    RETURNING s.card_number
), credit_hold AS (
    UPDATE cards c
    SET outstanding_balance = c.outstanding_balance + $2,
        updated_at = current_timestamp
    FROM locked_card lc
    WHERE c.card_id = lc.card_id
      AND lc.card_type = 'credit'
      AND $2 > 0
      AND (lc.credit_limit <= 0 OR lc.outstanding_balance + $2 <= lc.credit_limit)
    RETURNING c.card_number
), authorized_transaction AS (
    INSERT INTO transactions (
        card_number,
        amount,
        payment_method,
        merchant_id,
        transaction_time,
        idempotency_key,
        status,
        created_at,
        updated_at
    )
    SELECT $1, $2, $3, $4, $5, $6, 'authorized', current_timestamp, current_timestamp
    WHERE EXISTS (SELECT 1 FROM debit_hold)
       OR EXISTS (SELECT 1 FROM credit_hold)
    RETURNING
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
), queued_fraud_event AS (
    INSERT INTO outbox_events (event_key, topic, message_key, payload)
    SELECT
        'transaction-fraud:' || at.transaction_id::text,
        'transaction-fraud-check',
        at.transaction_id::text,
        jsonb_build_object(
            'transaction_id', at.transaction_id,
            'card_number', at.card_number,
            'amount', at.amount,
            'card_type', lc.card_type
        )
    FROM authorized_transaction at
    JOIN locked_card lc ON lc.card_number = at.card_number
    ON CONFLICT (event_key) DO NOTHING
    RETURNING event_key
)
SELECT
    at.transaction_id,
    at.transaction_no,
    at.card_number,
    at.amount,
    at.payment_method,
    at.merchant_id,
    at.transaction_time,
    at.status,
    at.created_at,
    at.updated_at
FROM authorized_transaction at`

const getTransactionByIdempotencyKey = `SELECT
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
WHERE idempotency_key = $1
  AND idempotency_key <> ''`

const createTransactionAtomic = `WITH active_recipient AS (
    SELECT c.card_number, u.email
    FROM cards c
    JOIN users u ON u.user_id = c.user_id AND u.deleted_at IS NULL
    WHERE c.card_number = $1 AND c.deleted_at IS NULL
), locked_saldos AS MATERIALIZED (
    SELECT s.card_number, s.total_balance
    FROM saldos s
    WHERE s.card_number IN ($1, $2)
      AND s.deleted_at IS NULL
    ORDER BY s.card_number
    FOR UPDATE
), eligible AS (
    SELECT 1
    FROM saldos AS user_saldo
    JOIN saldos AS merchant_saldo ON merchant_saldo.card_number = $2
    WHERE user_saldo.card_number = $1
      AND user_saldo.deleted_at IS NULL
      AND merchant_saldo.deleted_at IS NULL
      AND user_saldo.total_balance >= $3
      AND EXISTS (SELECT 1 FROM locked_saldos)
      AND EXISTS (SELECT 1 FROM active_recipient)
      AND $3 > 0
      AND $1 <> $2
), settled_saldos AS (
    UPDATE saldos s
    SET total_balance = s.total_balance + CASE
            WHEN s.card_number = $1 THEN -$3
            ELSE $3
        END,
        updated_at = current_timestamp
    WHERE s.card_number IN ($1, $2)
      AND s.deleted_at IS NULL
      AND EXISTS (SELECT 1 FROM eligible)
    RETURNING s.card_number
), inserted_transaction AS (
    INSERT INTO transactions (
        card_number,
        amount,
        payment_method,
        merchant_id,
        transaction_time,
        idempotency_key,
        status,
        created_at,
        updated_at
    )
    SELECT $1, $3, $4, $5, $6, $7, 'success', current_timestamp, current_timestamp
    WHERE (SELECT COUNT(*) FROM settled_saldos) = 2
      AND EXISTS (SELECT 1 FROM active_recipient)
    RETURNING
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
), queued_event AS (
    INSERT INTO outbox_events (event_key, topic, message_key, payload)
    SELECT
        'transaction-email:' || it.transaction_id::text,
        'email-service-topic-transaction-create',
        it.transaction_id::text,
        jsonb_build_object(
            'event_id', gen_random_uuid()::text,
            'schema_version', 1,
            'event_type', 'transaction.created',
            'occurred_at', now(),
            'email', COALESCE(u.email, ''),
            'subject', 'Transaction Successful - SanEdge',
            'body', 'Your transaction of ' || it.amount::text || ' has been processed successfully.'
        )
    FROM inserted_transaction it
    JOIN active_recipient u ON u.card_number = it.card_number
    ON CONFLICT (event_key) DO NOTHING
    RETURNING event_key
), queued_cache_event AS (
    INSERT INTO outbox_events (event_key, topic, message_key, payload)
    SELECT
        'transaction-cache-invalidate:' || it.transaction_id::text,
        'saldo-service-topic-invalidate-cache',
        it.transaction_id::text,
        jsonb_build_object(
            'event', 'saldo.cache.invalidate',
            'card_numbers', jsonb_build_array($1, $2)
        )
    FROM inserted_transaction it
    ON CONFLICT (event_key) DO NOTHING
    RETURNING event_key
)
SELECT
    it.transaction_id,
    it.transaction_no,
    it.card_number,
    it.amount,
    it.payment_method,
    it.merchant_id,
    it.transaction_time,
    it.status,
    it.created_at,
    it.updated_at
FROM inserted_transaction it`

const createTransaction = `INSERT INTO
    transactions (
        card_number,
        amount,
        payment_method,
        merchant_id,
        transaction_time,
        created_at,
        updated_at
    )
VALUES (
        $1,
        $2,
        $3,
        $4,
        $5,
        current_timestamp,
        current_timestamp
    )
RETURNING
    transaction_id,
    transaction_no,
    card_number,
    amount,
    payment_method,
    merchant_id,
    transaction_time,
    status,
    created_at,
    updated_at`

const updateTransaction = `UPDATE transactions
SET
    card_number = $2,
    amount = $3,
    payment_method = $4,
    merchant_id = $5,
    transaction_time = $6,
    updated_at = current_timestamp
WHERE
    transaction_id = $1
    AND deleted_at IS NULL
RETURNING
    transaction_id,
    transaction_no,
    card_number,
    amount,
    payment_method,
    merchant_id,
    transaction_time,
    status,
    created_at,
    updated_at`

const updateTransactionAtomic = `WITH locked_transaction AS MATERIALIZED (
    SELECT t.transaction_id, t.transaction_no, t.card_number, t.amount, t.payment_method, t.merchant_id, t.transaction_time, t.status, t.created_at, t.updated_at
    FROM transactions AS t
    WHERE t.transaction_id = $1
      AND t.status IN ('pending', 'failed', 'success', 'compensation_required')
      AND t.deleted_at IS NULL
    FOR UPDATE
), locked_saldos AS MATERIALIZED (
    SELECT s.card_number, s.total_balance
    FROM saldos s
    WHERE s.card_number IN ((SELECT card_number FROM locked_transaction), $2)
      AND s.deleted_at IS NULL
    ORDER BY s.card_number
    FOR UPDATE
), eligible AS (
    SELECT 1
    FROM locked_transaction lt
    JOIN saldos AS customer_saldo ON customer_saldo.card_number = lt.card_number
    JOIN saldos AS merchant_saldo ON merchant_saldo.card_number = $2
    WHERE lt.card_number <> $2
      AND customer_saldo.deleted_at IS NULL
      AND merchant_saldo.deleted_at IS NULL
      AND EXISTS (SELECT 1 FROM locked_saldos)
      AND $3 > 0
      AND customer_saldo.total_balance + lt.amount - $3 >= 0
), settled_saldos AS (
    UPDATE saldos s
    SET total_balance = s.total_balance + CASE
            WHEN s.card_number = lt.card_number THEN lt.amount - $3
            ELSE $3 - lt.amount
        END,
        updated_at = current_timestamp
    FROM locked_transaction lt
    WHERE s.card_number IN (lt.card_number, $2)
      AND s.deleted_at IS NULL
      AND EXISTS (SELECT 1 FROM eligible)
    RETURNING s.card_number
), updated_transaction AS (
    UPDATE transactions t
    SET card_number = $4,
        amount = $3,
        payment_method = $5,
        merchant_id = $6,
        transaction_time = $7,
        status = 'success',
        updated_at = current_timestamp
    FROM locked_transaction lt
    WHERE t.transaction_id = lt.transaction_id
      AND (SELECT COUNT(*) FROM settled_saldos) = 2
    RETURNING
        t.transaction_id,
        t.transaction_no,
        t.card_number,
        t.amount,
        t.payment_method,
        t.merchant_id,
        t.transaction_time,
        t.status,
        t.created_at,
        t.updated_at
)
SELECT
    ut.transaction_id,
    ut.transaction_no,
    ut.card_number,
    ut.amount,
    ut.payment_method,
    ut.merchant_id,
    ut.transaction_time,
    ut.status,
    ut.created_at,
    ut.updated_at
FROM updated_transaction ut`

const updateTransactionStatus = `UPDATE transactions
SET
    status = $2,
    updated_at = current_timestamp
WHERE
    transaction_id = $1
    AND deleted_at IS NULL
    AND (
        ($2 = 'failed' AND status IN ('pending', 'failed', 'compensation_required'))
        OR ($2 = 'compensation_required' AND status IN ('pending', 'failed', 'success', 'authorized', 'captured', 'compensation_required'))
        OR ($2 = 'success' AND status IN ('pending', 'failed', 'success', 'compensation_required'))
        OR ($2 = 'pending' AND status IN ('failed', 'compensation_required'))
        OR ($2 = 'flagged_fraud' AND status IN ('success', 'authorized', 'captured'))
    )
RETURNING
    transaction_id,
    transaction_no,
    card_number,
    amount,
    payment_method,
    merchant_id,
    transaction_time,
    status,
    created_at,
    updated_at`

const captureTransactionAtomic = `WITH locked_transaction AS MATERIALIZED (
    SELECT t.transaction_id, t.transaction_no, t.card_number, t.amount, t.payment_method, t.merchant_id, t.transaction_time, t.status, t.created_at, t.updated_at
    FROM transactions AS t
    JOIN cards AS customer_card
      ON customer_card.card_number = t.card_number
     AND customer_card.deleted_at IS NULL
    WHERE t.transaction_id = $1
      AND t.status = 'authorized'
      AND t.deleted_at IS NULL
    FOR UPDATE OF t
), settled_merchant AS (
    UPDATE saldos s
    SET total_balance = s.total_balance + lt.amount,
        updated_at = current_timestamp
    FROM locked_transaction lt
    WHERE s.card_number = $2
      AND s.deleted_at IS NULL
    RETURNING s.card_number
), reward_ledger AS (
    INSERT INTO card_rewards (
        card_number, txn_id, amount, mcc, points_earned, expires_at, redeemed
    )
    SELECT
        lt.card_number,
        lt.transaction_id::text,
        lt.amount,
        '',
        GREATEST((lt.amount / 10000)::integer, 1),
        current_timestamp + interval '1 year',
        FALSE
    FROM locked_transaction lt
    WHERE EXISTS (SELECT 1 FROM settled_merchant)
    ON CONFLICT (txn_id) WHERE txn_id <> '' DO NOTHING
    RETURNING card_number, points_earned
), earned_rewards AS (
    UPDATE cards c
    SET reward_points = c.reward_points + rl.points_earned,
        updated_at = current_timestamp
    FROM reward_ledger rl
    WHERE c.card_number = rl.card_number
      AND c.deleted_at IS NULL
    RETURNING c.card_number
), captured_transaction AS (

    UPDATE transactions t
    SET status = 'captured',
        updated_at = current_timestamp
    FROM locked_transaction lt
    WHERE t.transaction_id = lt.transaction_id
      AND EXISTS (SELECT 1 FROM settled_merchant)
      AND EXISTS (SELECT 1 FROM earned_rewards)
    RETURNING
        t.transaction_id,
        t.transaction_no,
        t.card_number,
        t.amount,
        t.payment_method,
        t.merchant_id,
        t.transaction_time,
        t.status,
        t.created_at,
        t.updated_at
)
SELECT
    ct.transaction_id,
    ct.transaction_no,
    ct.card_number,
    ct.amount,
    ct.payment_method,
    ct.merchant_id,
    ct.transaction_time,
    ct.status,
    ct.created_at,
    ct.updated_at
FROM captured_transaction ct`

const voidTransactionAtomic = `WITH locked_transaction AS MATERIALIZED (
    SELECT t.transaction_id, t.transaction_no, t.card_number, t.amount, t.payment_method, t.merchant_id, t.transaction_time, t.status, t.created_at, t.updated_at, c.card_type
    FROM transactions t
    JOIN cards c ON c.card_number = t.card_number AND c.deleted_at IS NULL
    WHERE t.transaction_id = $1
      AND t.status IN ('pending', 'authorized')
      AND t.deleted_at IS NULL
    FOR UPDATE OF t
), released_hold AS (
    UPDATE saldos s
    SET total_balance = s.total_balance + lt.amount,
        updated_at = current_timestamp
    FROM locked_transaction lt
    WHERE lt.card_type = 'debit'
      AND s.card_number = lt.card_number
      AND s.deleted_at IS NULL
    RETURNING s.card_number
), released_credit_hold AS (
    UPDATE cards c
    SET outstanding_balance = c.outstanding_balance - lt.amount,
        updated_at = current_timestamp
    FROM locked_transaction lt
    WHERE lt.card_type = 'credit'
      AND c.card_number = lt.card_number
      AND c.deleted_at IS NULL
      AND c.outstanding_balance >= lt.amount
    RETURNING c.card_number
), voided_transaction AS (
    UPDATE transactions t
    SET status = 'voided',
        updated_at = current_timestamp
    FROM locked_transaction lt
    WHERE t.transaction_id = lt.transaction_id
      AND ((lt.card_type = 'credit' AND EXISTS (SELECT 1 FROM released_credit_hold))
           OR (lt.card_type <> 'credit' AND EXISTS (SELECT 1 FROM released_hold)))
    RETURNING
        t.transaction_id,
        t.transaction_no,
        t.card_number,
        t.amount,
        t.payment_method,
        t.merchant_id,
        t.transaction_time,
        t.status,
        t.created_at,
        t.updated_at
)
SELECT
    vt.transaction_id,
    vt.transaction_no,
    vt.card_number,
    vt.amount,
    vt.payment_method,
    vt.merchant_id,
    vt.transaction_time,
    vt.status,
    vt.created_at,
    vt.updated_at
FROM voided_transaction vt`

const refundTransactionAtomic = `WITH locked_transaction AS MATERIALIZED (
    SELECT t.transaction_id, t.transaction_no, t.card_number, t.amount, t.payment_method, t.merchant_id, t.transaction_time, t.status, t.created_at, t.updated_at, c.card_type
    FROM transactions t
    JOIN cards c ON c.card_number = t.card_number AND c.deleted_at IS NULL
    WHERE t.transaction_id = $1
      AND t.status = 'captured'
      AND t.deleted_at IS NULL
    FOR UPDATE OF t
), locked_saldos AS MATERIALIZED (
    SELECT s.card_number, s.total_balance
    FROM saldos s
    WHERE s.card_number IN ((SELECT card_number FROM locked_transaction), $2)
      AND s.deleted_at IS NULL
    ORDER BY s.card_number
    FOR UPDATE
), eligible AS (
    SELECT 1
    FROM locked_transaction lt
    LEFT JOIN locked_saldos customer_saldo ON customer_saldo.card_number = lt.card_number
    JOIN locked_saldos merchant_saldo ON merchant_saldo.card_number = $2
    WHERE lt.card_number <> $2
      AND merchant_saldo.total_balance >= lt.amount
      AND (lt.card_type = 'credit' OR customer_saldo.card_number IS NOT NULL)
), customer_refunded AS (
    UPDATE saldos s
    SET total_balance = s.total_balance + lt.amount,
        updated_at = current_timestamp
    FROM locked_transaction lt
    WHERE lt.card_type = 'debit'
      AND s.card_number = lt.card_number
      AND s.deleted_at IS NULL
      AND EXISTS (SELECT 1 FROM eligible)
    RETURNING s.card_number
), outstanding_reduced AS (
    UPDATE cards c
    SET outstanding_balance = GREATEST(0, c.outstanding_balance - lt.amount),
        updated_at = current_timestamp
    FROM locked_transaction lt
    WHERE lt.card_type = 'credit'
      AND c.card_number = lt.card_number
      AND c.deleted_at IS NULL
      AND EXISTS (SELECT 1 FROM eligible)
    RETURNING c.card_number
), merchant_debited AS (
    UPDATE saldos s
    SET total_balance = s.total_balance - lt.amount,
        updated_at = current_timestamp
    FROM locked_transaction lt
    WHERE s.card_number = $2
      AND s.deleted_at IS NULL
      AND s.total_balance - lt.amount >= 0
      AND EXISTS (SELECT 1 FROM eligible)
    RETURNING s.card_number
), refunded_transaction AS (
    UPDATE transactions t
    SET status = 'refunded',
        updated_at = current_timestamp
    FROM locked_transaction lt
    WHERE t.transaction_id = lt.transaction_id
      AND EXISTS (SELECT 1 FROM merchant_debited)
      AND (EXISTS (SELECT 1 FROM customer_refunded) OR EXISTS (SELECT 1 FROM outstanding_reduced))
    RETURNING
        t.transaction_id,
        t.transaction_no,
        t.card_number,
        t.amount,
        t.payment_method,
        t.merchant_id,
        t.transaction_time,
        t.status,
        t.created_at,
        t.updated_at
)
SELECT
    rt.transaction_id,
    rt.transaction_no,
    rt.card_number,
    rt.amount,
    rt.payment_method,
    rt.merchant_id,
    rt.transaction_time,
    rt.status,
    rt.created_at,
    rt.updated_at
FROM refunded_transaction rt`

const trashTransaction = `UPDATE transactions
SET
    deleted_at = current_timestamp
WHERE
    transaction_id = $1
    AND deleted_at IS NULL
RETURNING
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
    deleted_at`

const restoreTransaction = `UPDATE transactions
SET
    deleted_at = NULL
WHERE
    transaction_id = $1
    AND deleted_at IS NOT NULL
RETURNING
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
    deleted_at`

const deleteTransactionPermanently = `DELETE FROM transactions
WHERE
    transaction_id = $1
    AND deleted_at IS NOT NULL`

const restoreAllTransactions = `UPDATE transactions
SET
    deleted_at = NULL
WHERE
    deleted_at IS NOT NULL`

const deleteAllPermanentTransactions = `DELETE FROM transactions WHERE deleted_at IS NOT NULL`

const updateSaldoBalanceDelta = `UPDATE saldos
SET
    total_balance = total_balance + $2,
    updated_at = current_timestamp
WHERE
    card_number = $1
    AND deleted_at IS NULL
    AND total_balance + $2 >= 0
RETURNING
    saldo_id,
    card_number,
    total_balance,
    withdraw_amount,
    withdraw_time,
    created_at,
    updated_at`

type transactionCommandRepository struct{ db *sqlx.DB }

type AuthorizeTransactionResult struct {
	Row      *db.AuthorizeTransactionAtomicRow
	Replayed bool
}

type TransactionAtomicResult struct {
	Row      *db.CreateTransactionAtomicRow
	Replayed bool
}

func NewTransactionCommandRepository(db *sqlx.DB) TransactionCommandRepository {
	return &transactionCommandRepository{db: db}
}

func (r *transactionCommandRepository) AuthorizeTransactionAtomic(ctx context.Context, params db.AuthorizeTransactionAtomicParams) (*db.AuthorizeTransactionAtomicRow, error) {
	var row db.AuthorizeTransactionAtomicRow
	err := r.db.GetContext(ctx, &row, authorizeTransactionAtomic,
		params.CardNumber,
		params.Amount,
		params.PaymentMethod,
		params.MerchantID,
		params.TransactionTime,
		params.IdempotencyKey,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sharedErrors.NewBadRequestError("card is inactive, balance is insufficient, or credit limit is exceeded")
		}
		return nil, sharedErrors.ErrNoRowsOrFailed(err, "Transaction", "authorize transaction atomically")
	}
	return &row, nil
}

func (r *transactionCommandRepository) AuthorizeTransactionAtomicIdempotent(ctx context.Context, params db.AuthorizeTransactionAtomicParams) (*AuthorizeTransactionResult, error) {
	if params.IdempotencyKey != "" {
		existing, err := r.getTransactionByIdempotencyKey(ctx, params.IdempotencyKey)
		if err == nil {
			if validationErr := validateAuthorizeReplay(existing, params); validationErr != nil {
				return nil, validationErr
			}
			return &AuthorizeTransactionResult{Row: toAuthorizeTransactionRow(existing), Replayed: true}, nil
		}
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
	}

	row, err := r.AuthorizeTransactionAtomic(ctx, params)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" && params.IdempotencyKey != "" {
			existing, getErr := r.getTransactionByIdempotencyKey(ctx, params.IdempotencyKey)
			if getErr == nil {
				if validationErr := validateAuthorizeReplay(existing, params); validationErr != nil {
					return nil, validationErr
				}
				return &AuthorizeTransactionResult{Row: toAuthorizeTransactionRow(existing), Replayed: true}, nil
			}
		}
		return nil, err
	}
	return &AuthorizeTransactionResult{Row: row}, nil
}

func (r *transactionCommandRepository) getTransactionByIdempotencyKey(ctx context.Context, idempotencyKey string) (*db.GetTransactionByIdempotencyKeyRow, error) {
	var row db.GetTransactionByIdempotencyKeyRow
	err := r.db.GetContext(ctx, &row, getTransactionByIdempotencyKey, idempotencyKey)
	if err != nil {
		return nil, err
	}
	return &row, nil
}

func validateAuthorizeReplay(existing *db.GetTransactionByIdempotencyKeyRow, params db.AuthorizeTransactionAtomicParams) error {
	if existing == nil || existing.Status != "authorized" || existing.CardNumber != params.CardNumber || existing.Amount != params.Amount || existing.PaymentMethod != params.PaymentMethod || existing.MerchantID != params.MerchantID || !sameWallClock(existing.TransactionTime, params.TransactionTime) {
		return sharedErrors.ErrConflict.WithMessage("transaction authorization idempotency key already used with a different operation or payload")
	}
	return nil
}

// sameWallClock reports whether two timestamps represent the same wall-clock
// value. PostgreSQL TIMESTAMP (without time zone) columns store wall-clock
// time, and pgx reads them back as UTC, so instant equality (time.Time.Equal)
// is unreliable across locations; comparing the wall-clock fields is. Note:
// a replay from a different timezone sending the same instant yields a
// different wall-clock string and therefore conflicts by design.
func sameWallClock(a, b time.Time) bool {
	const layout = "2006-01-02 15:04:05.999999"
	return a.Truncate(time.Microsecond).Format(layout) == b.Truncate(time.Microsecond).Format(layout)
}

func toAuthorizeTransactionRow(r *db.GetTransactionByIdempotencyKeyRow) *db.AuthorizeTransactionAtomicRow {
	if r == nil {
		return nil
	}
	return &db.AuthorizeTransactionAtomicRow{TransactionID: r.TransactionID, TransactionNo: r.TransactionNo, CardNumber: r.CardNumber, Amount: r.Amount, PaymentMethod: r.PaymentMethod, MerchantID: r.MerchantID, TransactionTime: r.TransactionTime, Status: r.Status, CreatedAt: r.CreatedAt, UpdatedAt: r.UpdatedAt}
}

func (r *transactionCommandRepository) CreateTransactionAtomic(ctx context.Context, params db.CreateTransactionAtomicParams) (*db.CreateTransactionAtomicRow, error) {
	var row db.CreateTransactionAtomicRow
	err := r.db.GetContext(ctx, &row, createTransactionAtomic,
		params.CardNumber,
		params.CardNumber2,
		params.Amount,
		params.PaymentMethod,
		params.MerchantID,
		params.TransactionTime,
		params.IdempotencyKey,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sharedErrors.NewBadRequestError("insufficient balance or invalid settlement accounts")
		}
		return nil, sharedErrors.ErrFailed("create transaction atomically").WithInternal(err)
	}
	return &row, nil
}

func (r *transactionCommandRepository) CreateTransactionAtomicIdempotent(ctx context.Context, params db.CreateTransactionAtomicParams) (*TransactionAtomicResult, error) {
	if params.IdempotencyKey != "" {
		existing, err := r.getTransactionByIdempotencyKey(ctx, params.IdempotencyKey)
		if err == nil {
			if validationErr := validateTransactionReplay(existing, params); validationErr != nil {
				return nil, validationErr
			}
			return &TransactionAtomicResult{Row: toTransactionAtomicRow(existing), Replayed: true}, nil
		}
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
	}

	res, err := r.CreateTransactionAtomic(ctx, params)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" && params.IdempotencyKey != "" {
			existing, gerr := r.getTransactionByIdempotencyKey(ctx, params.IdempotencyKey)
			if gerr == nil {
				if validationErr := validateTransactionReplay(existing, params); validationErr != nil {
					return nil, validationErr
				}
				return &TransactionAtomicResult{Row: toTransactionAtomicRow(existing), Replayed: true}, nil
			}
		}
		return nil, err
	}
	return &TransactionAtomicResult{Row: res}, nil
}

func validateTransactionReplay(existing *db.GetTransactionByIdempotencyKeyRow, params db.CreateTransactionAtomicParams) error {
	if existing == nil || existing.Status != "success" || existing.CardNumber != params.CardNumber || existing.Amount != params.Amount || existing.PaymentMethod != params.PaymentMethod || existing.MerchantID != params.MerchantID || !sameWallClock(existing.TransactionTime, params.TransactionTime) {
		return sharedErrors.ErrConflict.WithMessage("transaction idempotency key already used with a different operation or payload")
	}
	return nil
}

func toTransactionAtomicRow(r *db.GetTransactionByIdempotencyKeyRow) *db.CreateTransactionAtomicRow {
	if r == nil {
		return nil
	}
	return &db.CreateTransactionAtomicRow{TransactionID: r.TransactionID, TransactionNo: r.TransactionNo, CardNumber: r.CardNumber, Amount: r.Amount, PaymentMethod: r.PaymentMethod, MerchantID: r.MerchantID, TransactionTime: r.TransactionTime, Status: r.Status, CreatedAt: r.CreatedAt, UpdatedAt: r.UpdatedAt}
}

func (r *transactionCommandRepository) UpdateSaldoBalanceDelta(ctx context.Context, cardNumber string, delta int) error {
	var row db.UpdateSaldoBalanceDeltaRow
	err := r.db.GetContext(ctx, &row, updateSaldoBalanceDelta, cardNumber, int32(delta))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return sharedErrors.NewBadRequestError("insufficient balance after rollback")
		}
		return sharedErrors.ErrNoRowsOrFailed(err, "Transaction", "update transaction balance")
	}
	return nil
}

func (r *transactionCommandRepository) CreateTransaction(ctx context.Context, request *requests.CreateTransactionRequest) (*db.CreateTransactionRow, error) {
	var row db.CreateTransactionRow
	err := r.db.GetContext(ctx, &row, createTransaction,
		request.CardNumber,
		int32(request.Amount),
		request.PaymentMethod,
		int32(*request.MerchantID),
		request.TransactionTime,
	)
	if err != nil {
		return nil, transaction_errors.ErrCreateTransactionFailed.WithInternal(err)
	}
	return &row, nil
}

func (r *transactionCommandRepository) UpdateTransaction(ctx context.Context, request *requests.UpdateTransactionRequest) (*db.UpdateTransactionRow, error) {
	var row db.UpdateTransactionRow
	err := r.db.GetContext(ctx, &row, updateTransaction,
		int32(*request.TransactionID),
		request.CardNumber,
		int32(request.Amount),
		request.PaymentMethod,
		int32(*request.MerchantID),
		request.TransactionTime,
	)
	if err != nil {
		return nil, sharedErrors.ErrNoRowsOrFailed(err, "Transaction", "update transaction")
	}
	return &row, nil
}

func (r *transactionCommandRepository) UpdateTransactionStatus(ctx context.Context, request *requests.UpdateTransactionStatus) (*db.UpdateTransactionStatusRow, error) {
	var row db.UpdateTransactionStatusRow
	err := r.db.GetContext(ctx, &row, updateTransactionStatus,
		int32(request.TransactionID),
		request.Status,
	)
	if err != nil {
		return nil, sharedErrors.ErrNoRowsOrFailed(err, "Transaction", "update transaction status")
	}
	return &row, nil
}

func (r *transactionCommandRepository) UpdateTransactionAtomic(ctx context.Context, params db.UpdateTransactionAtomicParams) (*db.UpdateTransactionAtomicRow, error) {
	var row db.UpdateTransactionAtomicRow
	err := r.db.GetContext(ctx, &row, updateTransactionAtomic,
		params.TransactionID,
		params.MerchantCardNumber,
		params.Amount,
		params.CardNumber,
		params.PaymentMethod,
		params.MerchantID,
		params.TransactionTime,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sharedErrors.NewBadRequestError("insufficient balance for transaction update")
		}
		return nil, sharedErrors.ErrNoRowsOrFailed(err, "Transaction", "update transaction balance")
	}
	return &row, nil
}

func (r *transactionCommandRepository) CaptureTransactionAtomic(ctx context.Context, params db.CaptureTransactionAtomicParams) (*db.CaptureTransactionAtomicRow, error) {
	var row db.CaptureTransactionAtomicRow
	err := r.db.GetContext(ctx, &row, captureTransactionAtomic,
		params.TransactionID,
		params.MerchantCardNumber,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sharedErrors.NewBadRequestError("transaction is not authorized or merchant settlement failed")
		}
		return nil, sharedErrors.ErrNoRowsOrFailed(err, "Transaction", "capture transaction")
	}
	return &row, nil
}

func (r *transactionCommandRepository) VoidTransactionAtomic(ctx context.Context, transactionID int32) (*db.VoidTransactionAtomicRow, error) {
	var row db.VoidTransactionAtomicRow
	err := r.db.GetContext(ctx, &row, voidTransactionAtomic, transactionID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sharedErrors.NewBadRequestError("transaction is not pending/authorized or hold release failed")
		}
		return nil, sharedErrors.ErrNoRowsOrFailed(err, "Transaction", "void transaction")
	}
	return &row, nil
}

func (r *transactionCommandRepository) RefundTransactionAtomic(ctx context.Context, params db.RefundTransactionAtomicParams) (*db.RefundTransactionAtomicRow, error) {
	var row db.RefundTransactionAtomicRow
	err := r.db.GetContext(ctx, &row, refundTransactionAtomic,
		params.TransactionID,
		params.MerchantCardNumber,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sharedErrors.NewBadRequestError("transaction is not captured or merchant balance is insufficient")
		}
		return nil, sharedErrors.ErrNoRowsOrFailed(err, "Transaction", "refund transaction")
	}
	return &row, nil
}

func (r *transactionCommandRepository) TrashedTransaction(ctx context.Context, id int) (*db.TrashTransactionRow, error) {
	var row db.TrashTransactionRow
	err := r.db.GetContext(ctx, &row, trashTransaction, int32(id))
	if err != nil {
		return nil, sharedErrors.ErrNoRowsOrFailed(err, "Transaction", "trash transaction")
	}
	return &row, nil
}
func (r *transactionCommandRepository) RestoreTransaction(ctx context.Context, id int) (*db.RestoreTransactionRow, error) {
	var row db.RestoreTransactionRow
	err := r.db.GetContext(ctx, &row, restoreTransaction, int32(id))
	if err != nil {
		return nil, sharedErrors.ErrNoRowsOrFailed(err, "Transaction", "restore transaction")
	}
	return &row, nil
}
func (r *transactionCommandRepository) DeleteTransactionPermanent(ctx context.Context, id int) (bool, error) {
	if _, err := r.db.ExecContext(ctx, deleteTransactionPermanently, int32(id)); err != nil {
		return false, sharedErrors.ErrNoRowsOrFailed(err, "Transaction", "delete transaction permanently")
	}
	return true, nil
}
func (r *transactionCommandRepository) RestoreAllTransaction(ctx context.Context) (bool, error) {
	if _, err := r.db.ExecContext(ctx, restoreAllTransactions); err != nil {
		return false, transaction_errors.ErrRestoreAllTransactionsFailed.WithInternal(err)
	}
	return true, nil
}
func (r *transactionCommandRepository) DeleteAllTransactionPermanent(ctx context.Context) (bool, error) {
	if _, err := r.db.ExecContext(ctx, deleteAllPermanentTransactions); err != nil {
		return false, transaction_errors.ErrDeleteAllTransactionsPermanentFailed.WithInternal(err)
	}
	return true, nil
}
