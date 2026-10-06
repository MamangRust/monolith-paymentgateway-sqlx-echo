package repository

import (
	"context"
	"time"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	card_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/card_errors/repository"
	"github.com/jmoiron/sqlx"
)

const approveAuthTransaction = `UPDATE card_auth_transactions
SET status = 'approved', updated_at = CURRENT_TIMESTAMP
WHERE txn_id = $1 AND status = 'pending'
RETURNING auth_id, txn_id, card_number, merchant_id, amount, currency, mcc, pos_entry_mode, status, idempotency_key, risk_score, created_at, updated_at`

const declineAuthTransaction = `UPDATE card_auth_transactions
SET status = 'declined', updated_at = CURRENT_TIMESTAMP
WHERE txn_id = $1 AND status = 'pending'
RETURNING auth_id, txn_id, card_number, merchant_id, amount, currency, mcc, pos_entry_mode, status, idempotency_key, risk_score, created_at, updated_at`

const reverseAuthTransaction = `UPDATE card_auth_transactions
SET status = 'reversed', updated_at = CURRENT_TIMESTAMP
WHERE txn_id = $1 AND status = 'approved'
RETURNING auth_id, txn_id, card_number, merchant_id, amount, currency, mcc, pos_entry_mode, status, idempotency_key, risk_score, created_at, updated_at`

const insertAuthTransaction = `WITH inserted_auth AS (
    INSERT INTO card_auth_transactions (
        txn_id, card_number, merchant_id, amount, currency, mcc, pos_entry_mode, status, idempotency_key, risk_score
    ) VALUES ($1, $2, $3, $4, $5, $6, $7, 'pending', $8, 0)
    RETURNING auth_id, txn_id, card_number, merchant_id, amount, currency, mcc, pos_entry_mode, status, idempotency_key, risk_score, created_at, updated_at
), queued_event AS (
    INSERT INTO outbox_events (event_key, topic, message_key, payload)
    SELECT
        'card-txn-created:' || txn_id,
        'card.txn.created',
        txn_id,
        jsonb_build_object('txn_id', txn_id, 'card_number', card_number, 'amount', amount, 'mcc', mcc, 'merchant_id', merchant_id, 'status', status)
    FROM inserted_auth
    ON CONFLICT (event_key) DO NOTHING
    RETURNING event_key
)
SELECT auth_id, txn_id, card_number, merchant_id, amount, currency, mcc, pos_entry_mode, status, idempotency_key, risk_score, created_at, updated_at
FROM inserted_auth`

const getAuthTransactionByIdempotencyKey = `SELECT auth_id, txn_id, card_number, merchant_id, amount, currency, mcc, pos_entry_mode, status, idempotency_key, risk_score, created_at, updated_at
FROM card_auth_transactions
WHERE idempotency_key = $1
LIMIT 1`

const getAuthTransactionByTxnId = `SELECT auth_id, txn_id, card_number, merchant_id, amount, currency, mcc, pos_entry_mode, status, idempotency_key, risk_score, created_at, updated_at
FROM card_auth_transactions
WHERE txn_id = $1`

const getAuthTransactionsByCardNumber = `SELECT auth_id, txn_id, card_number, merchant_id, amount, currency, mcc, pos_entry_mode, status, idempotency_key, risk_score, created_at, updated_at,
    COUNT(*) OVER() AS total_count
FROM card_auth_transactions
WHERE card_number = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3`

const countRecentAuthByCardNumber = `SELECT COUNT(*)::INT AS total_count
FROM card_auth_transactions
WHERE card_number = $1 AND created_at >= $2`

const updateAuthTransactionRiskScore = `WITH updated_auth AS (
    UPDATE card_auth_transactions
    SET risk_score = $2, updated_at = CURRENT_TIMESTAMP
    WHERE txn_id = $1
    RETURNING txn_id, card_number
), blocked_card AS (
    UPDATE cards c
    SET status = 'blocked', updated_at = CURRENT_TIMESTAMP
    FROM updated_auth ua
    WHERE c.card_number = ua.card_number
      AND $2 >= 70
      AND c.deleted_at IS NULL
    RETURNING c.card_number
), queued_event AS (
    INSERT INTO outbox_events (event_key, topic, message_key, payload)
    SELECT
        'card-fraud-alert:' || ua.txn_id,
        'card.fraud.alert',
        ua.txn_id,
        jsonb_build_object(
            'txn_id', ua.txn_id,
            'card_number', ua.card_number,
            'risk_score', $2,
            'reason', CASE WHEN $2 >= 70 THEN 'HIGH_RISK' ELSE 'MEDIUM_RISK' END,
            'action', CASE WHEN $2 >= 70 THEN 'BLOCKED' ELSE 'FLAGGED' END
        )
    FROM updated_auth ua
    WHERE $2 >= 30
    ON CONFLICT (event_key) DO NOTHING
    RETURNING event_key
)
SELECT 1`

type cardAuthTransactionRepository struct {
	db *sqlx.DB
}

func NewCardAuthTransactionRepository(q *sqlx.DB) CardAuthTransactionRepository {
	return &cardAuthTransactionRepository{db: q}
}

func (r *cardAuthTransactionRepository) InsertPending(ctx context.Context, req *requests.AuthorizeCardRequest) (*db.CardAuthTransaction, error) {
	var res db.InsertAuthTransactionRow
	err := r.db.GetContext(ctx, &res, insertAuthTransaction,
		req.IdempotencyKey+"-txn",
		req.CardNumber,
		int32(req.MerchantID),
		req.Amount,
		req.Currency,
		req.Mcc,
		req.PosEntryMode,
		req.IdempotencyKey,
	)
	if err != nil {
		return nil, card_errors.ErrInsertAuthTransactionFailed.WithInternal(err)
	}
	return &db.CardAuthTransaction{
		AuthID:         res.AuthID,
		TxnID:          res.TxnID,
		CardNumber:     res.CardNumber,
		MerchantID:     res.MerchantID,
		Amount:         res.Amount,
		Currency:       res.Currency,
		Mcc:            res.Mcc,
		PosEntryMode:   res.PosEntryMode,
		Status:         res.Status,
		IdempotencyKey: res.IdempotencyKey,
		RiskScore:      res.RiskScore,
		CreatedAt:      res.CreatedAt,
		UpdatedAt:      res.UpdatedAt,
	}, nil
}

func (r *cardAuthTransactionRepository) Approve(ctx context.Context, txnID string) (*db.CardAuthTransaction, error) {
	var res db.CardAuthTransaction
	if err := r.db.GetContext(ctx, &res, approveAuthTransaction, txnID); err != nil {
		return nil, card_errors.ErrApproveAuthTransactionFailed.WithInternal(err)
	}
	return &res, nil
}

func (r *cardAuthTransactionRepository) Decline(ctx context.Context, txnID string) (*db.CardAuthTransaction, error) {
	var res db.CardAuthTransaction
	if err := r.db.GetContext(ctx, &res, declineAuthTransaction, txnID); err != nil {
		return nil, card_errors.ErrDeclineAuthTransactionFailed.WithInternal(err)
	}
	return &res, nil
}

func (r *cardAuthTransactionRepository) Reverse(ctx context.Context, txnID string) (*db.CardAuthTransaction, error) {
	var res db.CardAuthTransaction
	if err := r.db.GetContext(ctx, &res, reverseAuthTransaction, txnID); err != nil {
		return nil, card_errors.ErrReverseAuthTransactionFailed.WithInternal(err)
	}
	return &res, nil
}

func (r *cardAuthTransactionRepository) FindByIdempotencyKey(ctx context.Context, key string) (*db.CardAuthTransaction, error) {
	var res db.CardAuthTransaction
	if err := r.db.GetContext(ctx, &res, getAuthTransactionByIdempotencyKey, key); err != nil {
		return nil, card_errors.ErrGetAuthTransactionFailed.WithInternal(err)
	}
	return &res, nil
}

func (r *cardAuthTransactionRepository) FindByTxnID(ctx context.Context, txnID string) (*db.CardAuthTransaction, error) {
	var res db.CardAuthTransaction
	if err := r.db.GetContext(ctx, &res, getAuthTransactionByTxnId, txnID); err != nil {
		return nil, card_errors.ErrGetAuthTransactionFailed.WithInternal(err)
	}
	return &res, nil
}

func (r *cardAuthTransactionRepository) FindByCardNumber(ctx context.Context, cardNumber string, page, pageSize int) ([]*db.GetAuthTransactionsByCardNumberRow, error) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize
	var res []*db.GetAuthTransactionsByCardNumberRow
	err := r.db.SelectContext(ctx, &res, getAuthTransactionsByCardNumber,
		cardNumber,
		int32(pageSize),
		int32(offset),
	)
	if err != nil {
		return nil, card_errors.ErrGetAuthTransactionFailed.WithInternal(err)
	}
	return res, nil
}

func (r *cardAuthTransactionRepository) CountRecentByCardNumber(ctx context.Context, cardNumber string, since time.Time) (int, error) {
	var count int32
	err := r.db.GetContext(ctx, &count, countRecentAuthByCardNumber, cardNumber, since)
	if err != nil {
		return 0, card_errors.ErrGetAuthTransactionFailed.WithInternal(err)
	}
	return int(count), nil
}

func (r *cardAuthTransactionRepository) UpdateRiskScore(ctx context.Context, txnID string, score int) error {
	if _, err := r.db.ExecContext(ctx, updateAuthTransactionRiskScore, txnID, int32(score)); err != nil {
		return card_errors.ErrUpdateRiskScoreFailed.WithInternal(err)
	}
	return nil
}
