package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	sharedErrors "github.com/MamangRust/monolith-payment-gateway-shared/errors"
	topup_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/topup_errors/repository"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jmoiron/sqlx"
)

const createTopup = `INSERT INTO
    topups (
        card_number,
        topup_amount,
        topup_method,
        topup_time,
        created_at,
        updated_at
    )
VALUES (
        $1,
        $2,
        $3,
        $4,
        current_timestamp,
        current_timestamp
    )
RETURNING
    topup_id,
    topup_no,
    card_number,
    topup_amount,
    topup_method,
    topup_time,
    status,
    created_at,
    updated_at`

const createTopupAtomic = `WITH active_recipient AS (
    SELECT c.card_number, u.email
    FROM cards c
    JOIN users u ON u.user_id = c.user_id AND u.deleted_at IS NULL
    WHERE c.card_number = $1 AND c.deleted_at IS NULL
), locked_saldo AS MATERIALIZED (
    SELECT s.card_number
    FROM saldos s
    WHERE s.card_number = $1
      AND s.deleted_at IS NULL
    FOR UPDATE
), credited_saldo AS (
    UPDATE saldos s
    SET total_balance = s.total_balance + $2,
        updated_at = current_timestamp
    WHERE s.card_number = $1
      AND s.deleted_at IS NULL
      AND $2 > 0
      AND EXISTS (SELECT 1 FROM locked_saldo)
      AND EXISTS (SELECT 1 FROM active_recipient)
    RETURNING s.card_number
), inserted_topup AS (
    INSERT INTO topups (
        card_number,
        topup_amount,
        topup_method,
        topup_time,
        idempotency_key,
        status,
        created_at,
        updated_at
    )
    SELECT $1, $2, $3, current_timestamp, $4, 'success', current_timestamp, current_timestamp
    WHERE (SELECT COUNT(*) FROM credited_saldo) = 1
      AND EXISTS (SELECT 1 FROM active_recipient)
    RETURNING
        topup_id,
        topup_no,
        card_number,
        topup_amount,
        topup_method,
        topup_time,
        status,
        created_at,
        updated_at
), queued_event AS (
    INSERT INTO outbox_events (event_key, topic, message_key, payload)
    SELECT
        'topup-email:' || it.topup_id::text,
        'email-service-topic-topup-create',
        it.topup_id::text,
        jsonb_build_object(
            'event_id', gen_random_uuid()::text,
            'schema_version', 1,
            'event_type', 'topup.created',
            'occurred_at', now(),
            'email', COALESCE(u.email, ''),
            'subject', 'Topup Successful - SanEdge',
            'body', 'Your topup of ' || it.topup_amount::text || ' has been processed successfully.'
        )
    FROM inserted_topup it
    JOIN active_recipient u ON u.card_number = it.card_number
    ON CONFLICT (event_key) DO NOTHING
    RETURNING event_key
)
SELECT it.topup_id, it.topup_no, it.card_number, it.topup_amount,
       it.topup_method, it.topup_time, it.status, it.created_at, it.updated_at
FROM inserted_topup it`

const getTopupByIdempotencyKey = `SELECT
    topup_id,
    topup_no,
    card_number,
    topup_amount,
    topup_method,
    topup_time,
    status,
    created_at,
    updated_at
FROM topups
WHERE idempotency_key = $1
  AND idempotency_key <> ''`

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

const updateTopup = `UPDATE topups
SET
    card_number = $2,
    topup_amount = $3,
    topup_method = $4,
    topup_time = $5,
    updated_at = current_timestamp
WHERE
    topup_id = $1
    AND deleted_at IS NULL
RETURNING
    topup_id,
    topup_no,
    card_number,
    topup_amount,
    topup_method,
    topup_time,
    status,
    created_at,
    updated_at`

const updateTopupAtomic = `WITH locked_topup AS MATERIALIZED (
    SELECT t.topup_id, t.topup_amount
    FROM topups t
    WHERE t.topup_id = $1
      AND status IN ('pending', 'failed', 'success', 'compensation_required')
      AND deleted_at IS NULL
    FOR UPDATE
), settled_saldo AS (
    UPDATE saldos s
    SET total_balance = s.total_balance + (
            $2 - (
                SELECT t0.topup_amount
                FROM topups AS t0
                WHERE t0.topup_id = $1
            )
        ),
        updated_at = current_timestamp
    WHERE s.card_number = $3
      AND s.deleted_at IS NULL
      AND EXISTS (SELECT 1 FROM locked_topup)
      AND s.total_balance + (
            $2 - (
                SELECT t0.topup_amount
                FROM topups AS t0
                WHERE t0.topup_id = $1
            )
        ) >= 0
    RETURNING s.card_number
), updated_topup AS (
    UPDATE topups t
    SET card_number = $3,
        topup_amount = $2,
        topup_method = $4,
        topup_time = current_timestamp,
        status = 'success',
        updated_at = current_timestamp
    WHERE t.topup_id = $1
      AND EXISTS (SELECT 1 FROM locked_topup)
      AND EXISTS (SELECT 1 FROM settled_saldo)
    RETURNING
        t.topup_id,
        t.topup_no,
        t.card_number,
        t.topup_amount,
        t.topup_method,
        t.topup_time,
        t.status,
        t.created_at,
        t.updated_at
)
SELECT
    ut.topup_id,
    ut.topup_no,
    ut.card_number,
    ut.topup_amount,
    ut.topup_method,
    ut.topup_time,
    ut.status,
    ut.created_at,
    ut.updated_at
FROM updated_topup ut`

const updateTopupAmount = `UPDATE topups
SET
    topup_amount = $2,
    updated_at = current_timestamp
WHERE
    topup_id = $1
    AND deleted_at IS NULL
RETURNING
    topup_id,
    topup_no,
    card_number,
    topup_amount,
    topup_method,
    topup_time,
    status,
    created_at,
    updated_at`

const updateTopupStatus = `UPDATE topups
SET
    status = $2,
    updated_at = current_timestamp
WHERE
    topup_id = $1
    AND deleted_at IS NULL
    AND (
        ($2 = 'failed' AND status IN ('pending', 'failed', 'compensation_required'))
        OR ($2 = 'compensation_required' AND status IN ('pending', 'failed', 'success', 'compensation_required'))
        OR ($2 = 'success' AND status IN ('pending', 'failed', 'success', 'compensation_required'))
        OR ($2 = 'pending' AND status IN ('failed', 'compensation_required'))
    )
RETURNING
    topup_id,
    topup_no,
    card_number,
    topup_amount,
    topup_method,
    topup_time,
    status,
    created_at,
    updated_at`

const trashTopup = `UPDATE topups
SET
    deleted_at = current_timestamp
WHERE
    topup_id = $1
    AND deleted_at IS NULL
RETURNING
    topup_id,
    topup_no,
    card_number,
    topup_amount,
    topup_method,
    topup_time,
    status,
    created_at,
    updated_at,
    deleted_at`

const restoreTopup = `UPDATE topups
SET
    deleted_at = NULL
WHERE
    topup_id = $1
    AND deleted_at IS NOT NULL
RETURNING
    topup_id,
    topup_no,
    card_number,
    topup_amount,
    topup_method,
    topup_time,
    status,
    created_at,
    updated_at,
    deleted_at`

const deleteTopupPermanently = `DELETE FROM topups WHERE topup_id = $1`

const restoreAllTopups = `UPDATE topups
SET
    deleted_at = NULL
WHERE
    deleted_at IS NOT NULL`

const deleteAllPermanentTopups = `DELETE FROM topups WHERE deleted_at IS NOT NULL`

type topupCommandRepository struct {
	db *sqlx.DB
}

func NewTopupCommandRepository(db *sqlx.DB) TopupCommandRepository {
	return &topupCommandRepository{
		db: db,
	}
}

func (r *topupCommandRepository) CreateTopup(ctx context.Context, request *requests.CreateTopupRequest) (*db.CreateTopupRow, error) {
	var topup db.CreateTopupRow
	err := r.db.GetContext(ctx, &topup, createTopup,
		request.CardNumber,
		int32(request.TopupAmount),
		request.TopupMethod,
		time.Time{},
	)

	if err != nil {
		return nil, sharedErrors.ErrConstraintOrFailed(err, "Topup", "create topup")
	}

	return &topup, nil
}

// TopupAtomicResult wraps the created row and whether the request was a
// replay of an already persisted idempotency key. Services use Replayed to
// skip side-effects (email, metadata updates) on replays.
type TopupAtomicResult struct {
	Row      *db.CreateTopupAtomicRow
	Replayed bool
}

func (r *topupCommandRepository) CreateTopupAtomic(ctx context.Context, request *requests.CreateTopupRequest) (*TopupAtomicResult, error) {
	if request.IdempotencyKey != "" {
		existing, err := r.getTopupByIdempotencyKey(ctx, request.IdempotencyKey)
		if err == nil {
			if validationErr := validateTopupReplay(existing, request); validationErr != nil {
				return nil, validationErr
			}
			return &TopupAtomicResult{Row: toTopupAtomicRow(existing), Replayed: true}, nil
		}
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
	}

	var topup db.CreateTopupAtomicRow
	err := r.db.GetContext(ctx, &topup, createTopupAtomic,
		request.CardNumber,
		int32(request.TopupAmount),
		request.TopupMethod,
		request.IdempotencyKey,
	)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" && request.IdempotencyKey != "" {
			// Concurrent duplicate: the whole statement rolled back (no balance
			// change), return the existing record.
			existing, gerr := r.getTopupByIdempotencyKey(ctx, request.IdempotencyKey)
			if gerr == nil {
				if validationErr := validateTopupReplay(existing, request); validationErr != nil {
					return nil, validationErr
				}
				return &TopupAtomicResult{Row: toTopupAtomicRow(existing), Replayed: true}, nil
			}
		}
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sharedErrors.NewBadRequestError("saldo not found for topup")
		}
		return nil, sharedErrors.ErrConstraintOrFailed(err, "Topup", "create topup")
	}
	return &TopupAtomicResult{Row: &topup}, nil
}

func (r *topupCommandRepository) getTopupByIdempotencyKey(ctx context.Context, idempotencyKey string) (*db.GetTopupByIdempotencyKeyRow, error) {
	var topup db.GetTopupByIdempotencyKeyRow
	err := r.db.GetContext(ctx, &topup, getTopupByIdempotencyKey, idempotencyKey)
	if err != nil {
		return nil, err
	}
	return &topup, nil
}

func validateTopupReplay(existing *db.GetTopupByIdempotencyKeyRow, request *requests.CreateTopupRequest) error {
	if existing == nil || existing.CardNumber != request.CardNumber || existing.TopupAmount != int32(request.TopupAmount) || existing.TopupMethod != request.TopupMethod {
		return sharedErrors.ErrConflict.WithMessage("topup idempotency key already used with different payload")
	}
	return nil
}

func toTopupAtomicRow(r *db.GetTopupByIdempotencyKeyRow) *db.CreateTopupAtomicRow {
	if r == nil {
		return nil
	}
	return &db.CreateTopupAtomicRow{
		TopupID:     r.TopupID,
		TopupNo:     r.TopupNo,
		CardNumber:  r.CardNumber,
		TopupAmount: r.TopupAmount,
		TopupMethod: r.TopupMethod,
		TopupTime:   r.TopupTime,
		Status:      r.Status,
		CreatedAt:   r.CreatedAt,
		UpdatedAt:   r.UpdatedAt,
	}
}

// UpdateSaldoBalanceDelta applies an atomic signed delta to a card balance.
// It returns an error (400) when the resulting balance would be negative.
func (r *topupCommandRepository) UpdateSaldoBalanceDelta(ctx context.Context, cardNumber string, delta int) error {
	var saldo db.UpdateSaldoBalanceDeltaRow
	err := r.db.GetContext(ctx, &saldo, updateSaldoBalanceDelta,
		cardNumber,
		int32(delta),
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return sharedErrors.NewBadRequestError("insufficient balance for topup update")
		}
		return sharedErrors.ErrNoRowsOrFailed(err, "Topup", "update topup balance")
	}
	return nil
}

// UpdateTopupAtomic updates a topup and settles its card balance in one SQL
// statement. The delta is computed from the locked pre-update amount inside
// the statement, so concurrent updates can never read a stale amount (no lost
// update, no negative double-apply). Returns 400 when the resulting balance
// would be negative or the card has no saldo.
func (r *topupCommandRepository) UpdateTopupAtomic(ctx context.Context, request *requests.UpdateTopupRequest) (*db.UpdateTopupAtomicRow, error) {
	var topup db.UpdateTopupAtomicRow
	err := r.db.GetContext(ctx, &topup, updateTopupAtomic,
		int32(*request.TopupID),
		int32(request.TopupAmount),
		request.CardNumber,
		request.TopupMethod,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sharedErrors.NewBadRequestError("insufficient balance for topup update")
		}
		return nil, sharedErrors.ErrNoRowsOrFailed(err, "Topup", "update topup balance")
	}
	return &topup, nil
}

func (r *topupCommandRepository) UpdateTopup(ctx context.Context, request *requests.UpdateTopupRequest) (*db.UpdateTopupRow, error) {
	var topup db.UpdateTopupRow
	err := r.db.GetContext(ctx, &topup, updateTopup,
		int32(*request.TopupID),
		request.CardNumber,
		int32(request.TopupAmount),
		request.TopupMethod,
		time.Time{},
	)

	if err != nil {
		return nil, sharedErrors.ErrNoRowsOrFailed(err, "Topup", "update topup")
	}

	return &topup, nil
}

func (r *topupCommandRepository) UpdateTopupAmount(ctx context.Context, request *requests.UpdateTopupAmount) (*db.UpdateTopupAmountRow, error) {
	var topup db.UpdateTopupAmountRow
	err := r.db.GetContext(ctx, &topup, updateTopupAmount,
		int32(request.TopupID),
		int32(request.TopupAmount),
	)

	if err != nil {
		return nil, sharedErrors.ErrNoRowsOrFailed(err, "Topup", "update topup amount")
	}

	return &topup, nil
}

func (r *topupCommandRepository) UpdateTopupStatus(ctx context.Context, request *requests.UpdateTopupStatus) (*db.UpdateTopupStatusRow, error) {
	var topup db.UpdateTopupStatusRow
	err := r.db.GetContext(ctx, &topup, updateTopupStatus,
		int32(request.TopupID),
		request.Status,
	)

	if err != nil {
		return nil, sharedErrors.ErrNoRowsOrFailed(err, "Topup", "update topup status")
	}

	return &topup, nil
}

func (r *topupCommandRepository) TrashedTopup(ctx context.Context, topup_id int) (*db.Topup, error) {
	var row db.TrashTopupRow
	err := r.db.GetContext(ctx, &row, trashTopup, int32(topup_id))
	if err != nil {
		return nil, sharedErrors.ErrNoRowsOrFailed(err, "Topup", "trash topup")
	}
	return &db.Topup{
		TopupID:     row.TopupID,
		TopupNo:     row.TopupNo,
		CardNumber:  row.CardNumber,
		TopupAmount: row.TopupAmount,
		TopupMethod: row.TopupMethod,
		TopupTime:   row.TopupTime,
		Status:      row.Status,
		CreatedAt:   row.CreatedAt,
		UpdatedAt:   row.UpdatedAt,
		DeletedAt:   row.DeletedAt,
	}, nil
}

func (r *topupCommandRepository) RestoreTopup(ctx context.Context, topup_id int) (*db.Topup, error) {
	var row db.RestoreTopupRow
	err := r.db.GetContext(ctx, &row, restoreTopup, int32(topup_id))
	if err != nil {
		return nil, sharedErrors.ErrNoRowsOrFailed(err, "Topup", "restore topup")
	}
	return &db.Topup{
		TopupID:     row.TopupID,
		TopupNo:     row.TopupNo,
		CardNumber:  row.CardNumber,
		TopupAmount: row.TopupAmount,
		TopupMethod: row.TopupMethod,
		TopupTime:   row.TopupTime,
		Status:      row.Status,
		CreatedAt:   row.CreatedAt,
		UpdatedAt:   row.UpdatedAt,
		DeletedAt:   row.DeletedAt,
	}, nil
}

func (r *topupCommandRepository) DeleteTopupPermanent(ctx context.Context, topup_id int) (bool, error) {
	if _, err := r.db.ExecContext(ctx, deleteTopupPermanently, int32(topup_id)); err != nil {
		return false, sharedErrors.ErrNoRowsOrFailed(err, "Topup", "delete topup permanently")
	}
	return true, nil
}

func (r *topupCommandRepository) RestoreAllTopup(ctx context.Context) (bool, error) {
	if _, err := r.db.ExecContext(ctx, restoreAllTopups); err != nil {
		return false, topup_errors.ErrRestoreAllTopupFailed.WithInternal(err)
	}

	return true, nil
}

func (r *topupCommandRepository) DeleteAllTopupPermanent(ctx context.Context) (bool, error) {
	if _, err := r.db.ExecContext(ctx, deleteAllPermanentTopups); err != nil {
		return false, topup_errors.ErrDeleteAllTopupPermanentFailed.WithInternal(err)
	}

	return true, nil
}
