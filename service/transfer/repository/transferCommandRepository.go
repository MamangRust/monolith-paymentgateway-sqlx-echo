package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	sharedErrors "github.com/MamangRust/monolith-payment-gateway-shared/errors"
	transfer_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/transfer_errors/repository"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jmoiron/sqlx"
)

const createTransfer = `INSERT INTO
    transfers (
        transfer_from,
        transfer_to,
        transfer_amount,
        transfer_time,
        status,
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
    transfer_id,
    transfer_no,
    transfer_from,
    transfer_to,
    transfer_amount,
    transfer_time,
    status,
    created_at,
    updated_at`

const getTransferByIdempotencyKey = `SELECT
    transfer_id,
    transfer_no,
    transfer_from,
    transfer_to,
    transfer_amount,
    transfer_time,
    status,
    created_at,
    updated_at
FROM transfers
WHERE idempotency_key = $1
  AND idempotency_key <> ''`

const createTransferAtomic = `WITH active_recipient AS (
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
    FROM saldos AS sender_saldo
    JOIN saldos AS receiver_saldo ON receiver_saldo.card_number = $2
    WHERE sender_saldo.card_number = $1
      AND sender_saldo.deleted_at IS NULL
      AND receiver_saldo.deleted_at IS NULL
      AND sender_saldo.total_balance >= $3
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
), inserted_transfer AS (
    INSERT INTO transfers (
        transfer_from,
        transfer_to,
        transfer_amount,
        transfer_time,
        idempotency_key,
        status,
        created_at,
        updated_at
    )
    SELECT $1, $2, $3, current_timestamp, $4, 'success', current_timestamp, current_timestamp
    WHERE (SELECT COUNT(*) FROM settled_saldos) = 2
      AND EXISTS (SELECT 1 FROM active_recipient)
    RETURNING
        transfer_id,
        transfer_no,
        transfer_from,
        transfer_to,
        transfer_amount,
        transfer_time,
        status,
        created_at,
        updated_at
), queued_event AS (
    INSERT INTO outbox_events (event_key, topic, message_key, payload)
    SELECT
        'transfer-email:' || it.transfer_id::text,
        'email-service-topic-transfer-create',
        it.transfer_id::text,
        jsonb_build_object(
            'event_id', gen_random_uuid()::text,
            'schema_version', 1,
            'event_type', 'transfer.created',
            'occurred_at', now(),
            'email', COALESCE(u.email, ''),
            'subject', 'Transfer Successful - SanEdge',
            'body', 'Your transfer of ' || it.transfer_amount::text || ' has been processed successfully.'
        )
    FROM inserted_transfer it
    JOIN active_recipient u ON u.card_number = it.transfer_from
    ON CONFLICT (event_key) DO NOTHING
    RETURNING event_key
)
SELECT
    it.transfer_id,
    it.transfer_no,
    it.transfer_from,
    it.transfer_to,
    it.transfer_amount,
    it.transfer_time,
    it.status,
    it.created_at,
    it.updated_at
FROM inserted_transfer it`

const updateTransferSettlementDelta = `WITH locked AS MATERIALIZED (
    SELECT s.card_number
    FROM saldos s
    WHERE s.card_number IN ($1, $2)
      AND s.deleted_at IS NULL
    ORDER BY s.card_number
    FOR UPDATE
), eligible AS (
    SELECT 1
    FROM locked
    JOIN saldos sender ON sender.card_number = $1 AND sender.deleted_at IS NULL
    WHERE (SELECT COUNT(*) FROM locked) = 2
      AND sender.total_balance + $3 >= 0
), updated AS (
    UPDATE saldos s
    SET total_balance = s.total_balance + CASE
            WHEN s.card_number = $1 THEN $3
            ELSE $4
        END,
        updated_at = current_timestamp
    WHERE s.card_number IN ($1, $2)
      AND s.deleted_at IS NULL
      AND EXISTS (SELECT 1 FROM eligible)
    RETURNING s.card_number
)
SELECT COUNT(*)::int AS updated_count FROM updated`

const updateTransferAtomic = `WITH locked_transfer AS MATERIALIZED (
    SELECT transfers.transfer_id, transfers.transfer_from, transfers.transfer_to, transfers.transfer_amount
    FROM transfers
    WHERE transfers.transfer_id = $1
      AND transfers.deleted_at IS NULL
      AND transfers.status IN ('pending', 'failed', 'success', 'compensation_required')
    FOR UPDATE
), locked_saldos AS MATERIALIZED (
    SELECT s.card_number, s.total_balance
    FROM saldos s, locked_transfer t
    WHERE s.card_number IN (t.transfer_from, t.transfer_to)
      AND s.deleted_at IS NULL
    ORDER BY s.card_number
    FOR UPDATE
), eligible AS (
    SELECT 1
    FROM locked_transfer t
    JOIN saldos AS sender_saldo ON sender_saldo.card_number = t.transfer_from
    JOIN saldos AS receiver_saldo ON receiver_saldo.card_number = t.transfer_to
    WHERE t.transfer_from = $2
      AND t.transfer_to = $3
      AND t.transfer_from <> t.transfer_to
      AND sender_saldo.deleted_at IS NULL
      AND receiver_saldo.deleted_at IS NULL
      AND EXISTS (SELECT 1 FROM locked_saldos)
      AND sender_saldo.total_balance + t.transfer_amount - $4 >= 0
      AND $4 > 0
), settled_saldos AS (
    UPDATE saldos s
    SET total_balance = s.total_balance + CASE
        WHEN s.card_number = t.transfer_from THEN t.transfer_amount - $4
        ELSE $4 - t.transfer_amount
    END,
    updated_at = current_timestamp
    FROM locked_transfer t
    WHERE s.card_number IN (t.transfer_from, t.transfer_to)
      AND s.deleted_at IS NULL
      AND EXISTS (SELECT 1 FROM eligible)
    RETURNING s.card_number
), updated_transfer AS (
    UPDATE transfers t
    SET transfer_from = $2,
        transfer_to = $3,
        transfer_amount = $4,
        transfer_time = $5,
        status = 'success',
        updated_at = current_timestamp
    WHERE t.transfer_id = (SELECT transfer_id FROM locked_transfer)
      AND (SELECT COUNT(*) FROM settled_saldos) = 2
    RETURNING transfer_id, transfer_no, transfer_from, transfer_to, transfer_amount, transfer_time, status, created_at, updated_at
)
SELECT transfer_id, transfer_no, transfer_from, transfer_to, transfer_amount, transfer_time, status, created_at, updated_at
FROM updated_transfer`

const updateTransfer = `UPDATE transfers
SET
    transfer_from = $2,
    transfer_to = $3,
    transfer_amount = $4,
    transfer_time = $5,
    updated_at = current_timestamp
WHERE
    transfer_id = $1
    AND deleted_at IS NULL
RETURNING
    transfer_id,
    transfer_no,
    transfer_from,
    transfer_to,
    transfer_amount,
    transfer_time,
    status,
    created_at,
    updated_at`

const updateTransferAmount = `UPDATE transfers
SET
    transfer_amount = $2,
    transfer_time = current_timestamp,
    updated_at = current_timestamp
WHERE
    transfer_id = $1
    AND deleted_at IS NULL
RETURNING
    transfer_id,
    transfer_no,
    transfer_from,
    transfer_to,
    transfer_amount,
    transfer_time,
    status,
    created_at,
    updated_at`

const updateTransferStatus = `UPDATE transfers
SET
    status = $2,
    updated_at = current_timestamp
WHERE
    transfer_id = $1
    AND deleted_at IS NULL
    AND (
        ($2 = 'failed' AND status IN ('pending', 'failed', 'compensation_required'))
        OR ($2 = 'compensation_required' AND status IN ('pending', 'failed', 'success', 'compensation_required'))
        OR ($2 = 'success' AND status IN ('pending', 'failed', 'success', 'compensation_required'))
        OR ($2 = 'pending' AND status IN ('failed', 'compensation_required'))
    )
RETURNING
    transfer_id,
    transfer_no,
    transfer_from,
    transfer_to,
    transfer_amount,
    transfer_time,
    status,
    created_at,
    updated_at`

const trashTransfer = `UPDATE transfers
SET
    deleted_at = current_timestamp
WHERE
    transfer_id = $1
    AND deleted_at IS NULL
RETURNING
    transfer_id,
    transfer_no,
    transfer_from,
    transfer_to,
    transfer_amount,
    transfer_time,
    status,
    created_at,
    updated_at,
    deleted_at`

const restoreTransfer = `UPDATE transfers
SET
    deleted_at = NULL
WHERE
    transfer_id = $1
    AND deleted_at IS NOT NULL
RETURNING
    transfer_id,
    transfer_no,
    transfer_from,
    transfer_to,
    transfer_amount,
    transfer_time,
    status,
    created_at,
    updated_at,
    deleted_at`

const deleteTransferPermanently = `DELETE FROM transfers
WHERE
    transfer_id = $1
    AND deleted_at IS NOT NULL`

const restoreAllTransfers = `UPDATE transfers
SET
    deleted_at = NULL
WHERE
    deleted_at IS NOT NULL`

const deleteAllPermanentTransfers = `DELETE FROM transfers WHERE deleted_at IS NOT NULL`

type transferCommandRepository struct {
	db *sqlx.DB
}

func NewTransferCommandRepository(db *sqlx.DB) TransferCommandRepository {
	return &transferCommandRepository{
		db: db,
	}
}

func (r *transferCommandRepository) CreateTransfer(ctx context.Context, request *requests.CreateTransferRequest) (*db.CreateTransferRow, error) {
	var row db.CreateTransferRow
	err := r.db.GetContext(ctx, &row, createTransfer,
		request.TransferFrom,
		request.TransferTo,
		int32(request.TransferAmount),
		time.Now(),
		"success",
	)

	if err != nil {
		return nil, sharedErrors.ErrConstraintOrFailed(err, "Transfer", "create transfer")
	}

	return &row, nil
}

// TransferAtomicResult wraps the created row and whether the request was a
// replay of an already persisted idempotency key. Services use Replayed to
// skip side-effects (email, status updates) on replays.
type TransferAtomicResult struct {
	Row      *db.CreateTransferAtomicRow
	Replayed bool
}

func (r *transferCommandRepository) CreateTransferAtomic(ctx context.Context, request *requests.CreateTransferRequest) (*TransferAtomicResult, error) {
	if request.IdempotencyKey != "" {
		existing, err := r.getTransferByIdempotencyKey(ctx, request.IdempotencyKey)
		if err == nil {
			if validationErr := validateTransferReplay(existing, request); validationErr != nil {
				return nil, validationErr
			}
			return &TransferAtomicResult{Row: toTransferAtomicRow(existing), Replayed: true}, nil
		}
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
	}

	var row db.CreateTransferAtomicRow
	err := r.db.GetContext(ctx, &row, createTransferAtomic,
		request.TransferFrom,
		request.TransferTo,
		int32(request.TransferAmount),
		request.IdempotencyKey,
	)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" && request.IdempotencyKey != "" {
			// Concurrent duplicate: the whole statement rolled back (no balance
			// change), return the existing record.
			existing, gerr := r.getTransferByIdempotencyKey(ctx, request.IdempotencyKey)
			if gerr == nil {
				if validationErr := validateTransferReplay(existing, request); validationErr != nil {
					return nil, validationErr
				}
				return &TransferAtomicResult{Row: toTransferAtomicRow(existing), Replayed: true}, nil
			}
		}
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sharedErrors.NewBadRequestError("insufficient balance for transfer")
		}
		return nil, sharedErrors.ErrConstraintOrFailed(err, "Transfer", "create transfer")
	}
	return &TransferAtomicResult{Row: &row}, nil
}

func (r *transferCommandRepository) getTransferByIdempotencyKey(ctx context.Context, idempotencyKey string) (*db.GetTransferByIdempotencyKeyRow, error) {
	var row db.GetTransferByIdempotencyKeyRow
	err := r.db.GetContext(ctx, &row, getTransferByIdempotencyKey, idempotencyKey)
	if err != nil {
		return nil, err
	}
	return &row, nil
}

func validateTransferReplay(existing *db.GetTransferByIdempotencyKeyRow, request *requests.CreateTransferRequest) error {
	if existing == nil || existing.TransferFrom != request.TransferFrom || existing.TransferTo != request.TransferTo || existing.TransferAmount != int32(request.TransferAmount) {
		return sharedErrors.ErrConflict.WithMessage("transfer idempotency key already used with different payload")
	}
	return nil
}

func toTransferAtomicRow(r *db.GetTransferByIdempotencyKeyRow) *db.CreateTransferAtomicRow {
	if r == nil {
		return nil
	}
	return &db.CreateTransferAtomicRow{
		TransferID:     r.TransferID,
		TransferNo:     r.TransferNo,
		TransferFrom:   r.TransferFrom,
		TransferTo:     r.TransferTo,
		TransferAmount: r.TransferAmount,
		TransferTime:   r.TransferTime,
		Status:         r.Status,
		CreatedAt:      r.CreatedAt,
		UpdatedAt:      r.UpdatedAt,
	}
}

// UpdateTransferSettlementDelta applies sender/receiver deltas atomically for
// transfer amount updates. Returns a 400 when either account is missing or the
// sender's resulting balance would be negative.
func (r *transferCommandRepository) UpdateTransferAtomic(ctx context.Context, params db.UpdateTransferAtomicParams) (*db.UpdateTransferAtomicRow, error) {
	var row db.UpdateTransferAtomicRow
	err := r.db.GetContext(ctx, &row, updateTransferAtomic,
		params.TransferID,
		params.TransferFrom,
		params.TransferTo,
		params.TransferAmount,
		params.TransferTime,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sharedErrors.NewBadRequestError("transfer update is invalid or sender balance is insufficient")
		}
		return nil, sharedErrors.ErrNoRowsOrFailed(err, "Transfer", "update transfer atomically")
	}
	return &row, nil
}

func (r *transferCommandRepository) UpdateTransferSettlementDelta(ctx context.Context, senderCard, receiverCard string, senderDelta, receiverDelta int) error {
	var updatedCount int32
	err := r.db.GetContext(ctx, &updatedCount, updateTransferSettlementDelta,
		senderCard,
		receiverCard,
		int32(senderDelta),
		int32(receiverDelta),
	)
	if err != nil {
		return sharedErrors.ErrNoRowsOrFailed(err, "Transfer", "update transfer settlement")
	}
	if updatedCount != 2 {
		return sharedErrors.NewBadRequestError("insufficient balance for transfer update")
	}
	return nil
}

func (r *transferCommandRepository) UpdateTransfer(ctx context.Context, request *requests.UpdateTransferRequest) (*db.UpdateTransferRow, error) {
	var row db.UpdateTransferRow
	err := r.db.GetContext(ctx, &row, updateTransfer,
		int32(*request.TransferID),
		request.TransferFrom,
		request.TransferTo,
		int32(request.TransferAmount),
		time.Now(),
	)

	if err != nil {
		return nil, sharedErrors.ErrNoRowsOrFailed(err, "Transfer", "update transfer")
	}

	return &row, nil

}

func (r *transferCommandRepository) UpdateTransferAmount(ctx context.Context, request *requests.UpdateTransferAmountRequest) (*db.UpdateTransferAmountRow, error) {
	var row db.UpdateTransferAmountRow
	err := r.db.GetContext(ctx, &row, updateTransferAmount,
		int32(request.TransferID),
		int32(request.TransferAmount),
	)

	if err != nil {
		return nil, sharedErrors.ErrNoRowsOrFailed(err, "Transfer", "update transfer amount")
	}

	return &row, nil
}

func (r *transferCommandRepository) UpdateTransferStatus(ctx context.Context, request *requests.UpdateTransferStatus) (*db.UpdateTransferStatusRow, error) {
	var row db.UpdateTransferStatusRow
	err := r.db.GetContext(ctx, &row, updateTransferStatus,
		int32(request.TransferID),
		request.Status,
	)

	if err != nil {
		return nil, sharedErrors.ErrNoRowsOrFailed(err, "Transfer", "update transfer status")
	}

	return &row, nil
}

func (r *transferCommandRepository) TrashedTransfer(ctx context.Context, transfer_id int) (*db.Transfer, error) {
	var res db.TrashTransferRow
	err := r.db.GetContext(ctx, &res, trashTransfer, int32(transfer_id))

	if err != nil {
		return nil, sharedErrors.ErrNoRowsOrFailed(err, "Transfer", "trash transfer")
	}
	return &db.Transfer{
		TransferID:     res.TransferID,
		TransferNo:     res.TransferNo,
		TransferFrom:   res.TransferFrom,
		TransferTo:     res.TransferTo,
		TransferAmount: res.TransferAmount,
		TransferTime:   res.TransferTime,
		Status:         res.Status,
		CreatedAt:      res.CreatedAt,
		UpdatedAt:      res.UpdatedAt,
		DeletedAt:      res.DeletedAt,
	}, nil
}

func (r *transferCommandRepository) RestoreTransfer(ctx context.Context, transfer_id int) (*db.Transfer, error) {
	var res db.RestoreTransferRow
	err := r.db.GetContext(ctx, &res, restoreTransfer, int32(transfer_id))

	if err != nil {
		return nil, sharedErrors.ErrNoRowsOrFailed(err, "Transfer", "restore transfer")
	}
	return &db.Transfer{
		TransferID:     res.TransferID,
		TransferNo:     res.TransferNo,
		TransferFrom:   res.TransferFrom,
		TransferTo:     res.TransferTo,
		TransferAmount: res.TransferAmount,
		TransferTime:   res.TransferTime,
		Status:         res.Status,
		CreatedAt:      res.CreatedAt,
		UpdatedAt:      res.UpdatedAt,
		DeletedAt:      res.DeletedAt,
	}, nil
}

func (r *transferCommandRepository) DeleteTransferPermanent(ctx context.Context, transfer_id int) (bool, error) {
	if _, err := r.db.ExecContext(ctx, deleteTransferPermanently, int32(transfer_id)); err != nil {
		return false, sharedErrors.ErrNoRowsOrFailed(err, "Transfer", "delete transfer permanently")
	}
	return true, nil
}

func (r *transferCommandRepository) RestoreAllTransfer(ctx context.Context) (bool, error) {
	if _, err := r.db.ExecContext(ctx, restoreAllTransfers); err != nil {
		return false, transfer_errors.ErrRestoreAllTransfersFailed.WithInternal(err)
	}

	return true, nil
}

func (r *transferCommandRepository) DeleteAllTransferPermanent(ctx context.Context) (bool, error) {
	if _, err := r.db.ExecContext(ctx, deleteAllPermanentTransfers); err != nil {
		return false, transfer_errors.ErrDeleteAllTransfersPermanentFailed.WithInternal(err)
	}

	return true, nil
}
