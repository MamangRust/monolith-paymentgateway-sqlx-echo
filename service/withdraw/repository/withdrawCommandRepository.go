package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	sharedErrors "github.com/MamangRust/monolith-payment-gateway-shared/errors"
	withdraw_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/withdraw_errors/repository"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jmoiron/sqlx"
)

const defaultDailyWithdrawLimit int64 = 10_000_000

const createWithdraw = `INSERT INTO
    withdraws (
        card_number,
        withdraw_amount,
        withdraw_time,
        created_at,
        updated_at
    )
VALUES (
        $1,
        $2,
        $3,
        $4,
        current_timestamp
    )
RETURNING
    withdraw_id,
    withdraw_no,
    card_number,
    withdraw_amount,
    withdraw_time,
    status,
    created_at,
    updated_at`

const createWithdrawAtomic = `WITH active_recipient AS (
    SELECT c.card_number, u.email
    FROM cards c
    JOIN users u ON u.user_id = c.user_id AND u.deleted_at IS NULL
    WHERE c.card_number = $1 AND c.deleted_at IS NULL
), locked_saldo AS MATERIALIZED (
    SELECT s.card_number, s.total_balance
    FROM saldos s
    WHERE s.card_number = $1
      AND s.deleted_at IS NULL
    FOR UPDATE
), daily_withdrawals AS (
    SELECT COALESCE(SUM(w.withdraw_amount), 0)::BIGINT AS total_amount
    FROM withdraws w
    WHERE w.card_number = $1
      AND w.status = 'success'
      AND w.deleted_at IS NULL
      AND w.withdraw_time >= date_trunc('day', $2::timestamp)
      AND w.withdraw_time < date_trunc('day', $2::timestamp) + interval '1 day'
), eligible AS (
    SELECT 1
    FROM saldos AS saldo
    WHERE saldo.card_number = $1
      AND saldo.deleted_at IS NULL
      AND saldo.total_balance >= $3
      AND EXISTS (SELECT 1 FROM locked_saldo)
      AND EXISTS (SELECT 1 FROM active_recipient)
      AND COALESCE((
          SELECT COALESCE(SUM(daily_w.withdraw_amount), 0)::BIGINT
          FROM withdraws AS daily_w
          WHERE daily_w.card_number = $1
            AND daily_w.status = 'success'
            AND daily_w.deleted_at IS NULL
            AND daily_w.withdraw_time >= date_trunc('day', $2::timestamp)
            AND daily_w.withdraw_time < date_trunc('day', $2::timestamp) + interval '1 day'
      ), 0)::BIGINT + $3 <= $4
), debited_saldo AS (
    UPDATE saldos s
    SET total_balance = s.total_balance - $3,
        withdraw_amount = $3,
        withdraw_time = $2,
        updated_at = current_timestamp
    WHERE s.card_number = $1
      AND s.deleted_at IS NULL
      AND EXISTS (SELECT 1 FROM eligible)
    RETURNING s.card_number
), inserted_withdraw AS (
    INSERT INTO withdraws (
        card_number,
        withdraw_amount,
        withdraw_time,
        idempotency_key,
        status,
        created_at,
        updated_at
    )
    SELECT $1, $3, $2, $5, 'success', current_timestamp, current_timestamp
    WHERE (SELECT COUNT(*) FROM debited_saldo) = 1
      AND EXISTS (SELECT 1 FROM active_recipient)
    RETURNING
        withdraw_id,
        withdraw_no,
        card_number,
        withdraw_amount,
        withdraw_time,
        status,
        created_at,
        updated_at
), queued_event AS (
    INSERT INTO outbox_events (event_key, topic, message_key, payload)
    SELECT
        'withdraw-email:' || iw.withdraw_id::text,
        'email-service-topic-withdraw-create',
        iw.withdraw_id::text,
        jsonb_build_object(
            'event_id', gen_random_uuid()::text,
            'schema_version', 1,
            'event_type', 'withdraw.created',
            'occurred_at', now(),
            'email', COALESCE(u.email, ''),
            'subject', 'Withdraw Successful - SanEdge',
            'body', 'Your withdrawal of ' || iw.withdraw_amount::text || ' has been processed successfully.'
        )
    FROM inserted_withdraw iw
    JOIN active_recipient u ON u.card_number = iw.card_number
    ON CONFLICT (event_key) DO NOTHING
    RETURNING event_key
)
SELECT
    iw.withdraw_id,
    iw.withdraw_no,
    iw.card_number,
    iw.withdraw_amount,
    iw.withdraw_time,
    iw.status,
    iw.created_at,
    iw.updated_at
FROM inserted_withdraw iw`

const getWithdrawByIdempotencyKey = `SELECT
    withdraw_id,
    withdraw_no,
    card_number,
    withdraw_amount,
    withdraw_time,
    status,
    created_at,
    updated_at
FROM withdraws
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

const updateWithdraw = `UPDATE withdraws
SET
    card_number = $2,
    withdraw_amount = $3,
    withdraw_time = $4,
    updated_at = current_timestamp
WHERE
    withdraw_id = $1
    AND deleted_at IS NULL
RETURNING
    withdraw_id,
    withdraw_no,
    card_number,
    withdraw_amount,
    withdraw_time,
    status,
    created_at,
    updated_at`

const updateWithdrawAtomic = `WITH locked_withdraw AS MATERIALIZED (
    SELECT w.withdraw_id, w.withdraw_amount
    FROM withdraws w
    WHERE w.withdraw_id = $1
      AND status IN ('pending', 'failed', 'success', 'compensation_required')
      AND deleted_at IS NULL
    FOR UPDATE
), settled_saldo AS (
    UPDATE saldos s
    SET total_balance = s.total_balance + (
            (
                SELECT w0.withdraw_amount
                FROM withdraws AS w0
                WHERE w0.withdraw_id = $1
            ) - $2
        ),
        updated_at = current_timestamp
    WHERE s.card_number = $3
      AND s.deleted_at IS NULL
      AND EXISTS (SELECT 1 FROM locked_withdraw)
      AND s.total_balance + (
            (
                SELECT w0.withdraw_amount
                FROM withdraws AS w0
                WHERE w0.withdraw_id = $1
            ) - $2
        ) >= 0
    RETURNING s.card_number
), updated_withdraw AS (
    UPDATE withdraws t
    SET card_number = $3,
        withdraw_amount = $2,
        withdraw_time = $4,
        status = 'success',
        updated_at = current_timestamp
    WHERE t.withdraw_id = $1
      AND EXISTS (SELECT 1 FROM locked_withdraw)
      AND EXISTS (SELECT 1 FROM settled_saldo)
    RETURNING
        t.withdraw_id,
        t.withdraw_no,
        t.card_number,
        t.withdraw_amount,
        t.withdraw_time,
        t.status,
        t.created_at,
        t.updated_at
)
SELECT
    ut.withdraw_id,
    ut.withdraw_no,
    ut.card_number,
    ut.withdraw_amount,
    ut.withdraw_time,
    ut.status,
    ut.created_at,
    ut.updated_at
FROM updated_withdraw ut`

const updateWithdrawStatus = `UPDATE withdraws
SET
    status = $2,
    updated_at = current_timestamp
WHERE
    withdraw_id = $1
    AND deleted_at IS NULL
    AND (
        ($2 = 'failed' AND status IN ('pending', 'failed', 'compensation_required'))
        OR ($2 = 'compensation_required' AND status IN ('pending', 'failed', 'success', 'compensation_required'))
        OR ($2 = 'success' AND status IN ('pending', 'failed', 'success', 'compensation_required'))
        OR ($2 = 'pending' AND status IN ('failed', 'compensation_required'))
    )
RETURNING
    withdraw_id,
    withdraw_no,
    card_number,
    withdraw_amount,
    withdraw_time,
    status,
    created_at,
    updated_at`

const trashWithdraw = `UPDATE withdraws
SET
    deleted_at = current_timestamp
WHERE
    withdraw_id = $1
    AND deleted_at IS NULL
RETURNING
    withdraw_id,
    withdraw_no,
    card_number,
    withdraw_amount,
    withdraw_time,
    status,
    created_at,
    updated_at,
    deleted_at`

const restoreWithdraw = `UPDATE withdraws
SET
    deleted_at = NULL
WHERE
    withdraw_id = $1
    AND deleted_at IS NOT NULL
RETURNING
    withdraw_id,
    withdraw_no,
    card_number,
    withdraw_amount,
    withdraw_time,
    status,
    created_at,
    updated_at,
    deleted_at`

const deleteWithdrawPermanently = `DELETE FROM withdraws
WHERE
    withdraw_id = $1
    AND deleted_at IS NOT NULL`

const restoreAllWithdraws = `UPDATE withdraws
SET
    deleted_at = NULL
WHERE
    deleted_at IS NOT NULL`

const deleteAllPermanentWithdraws = `DELETE FROM withdraws WHERE deleted_at IS NOT NULL`

type withdrawCommandRepository struct {
	db         *sqlx.DB
	dailyLimit int64
}

func NewWithdrawCommandRepository(db *sqlx.DB, dailyLimit ...int64) WithdrawCommandRepository {
	limit := defaultDailyWithdrawLimit
	if len(dailyLimit) > 0 && dailyLimit[0] > 0 {
		limit = dailyLimit[0]
	}
	return &withdrawCommandRepository{
		db:         db,
		dailyLimit: limit,
	}
}

func (r *withdrawCommandRepository) CreateWithdraw(ctx context.Context, request *requests.CreateWithdrawRequest) (*db.CreateWithdrawRow, error) {
	var withdraw db.CreateWithdrawRow
	err := r.db.GetContext(ctx, &withdraw, createWithdraw,
		request.CardNumber,
		int32(request.WithdrawAmount),
		request.WithdrawTime,
		pgtype.Timestamp{},
	)

	if err != nil {
		return nil, sharedErrors.ErrConstraintOrFailed(err, "Withdraw", "create withdraw")
	}

	return &withdraw, nil
}

// WithdrawAtomicResult wraps the created row and whether the request was a
// replay of an already persisted idempotency key. Services use Replayed to
// skip side-effects (email) on replays.
type WithdrawAtomicResult struct {
	Row      *db.CreateWithdrawAtomicRow
	Replayed bool
}

func (r *withdrawCommandRepository) CreateWithdrawAtomic(ctx context.Context, request *requests.CreateWithdrawRequest) (*WithdrawAtomicResult, error) {
	if request.IdempotencyKey != "" {
		existing, err := r.getWithdrawByIdempotencyKey(ctx, request.IdempotencyKey)
		if err == nil {
			if validationErr := validateWithdrawReplay(existing, request); validationErr != nil {
				return nil, validationErr
			}
			return &WithdrawAtomicResult{Row: toWithdrawAtomicRow(existing), Replayed: true}, nil
		}
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
	}

	var withdraw db.CreateWithdrawAtomicRow
	err := r.db.GetContext(ctx, &withdraw, createWithdrawAtomic,
		request.CardNumber,
		request.WithdrawTime,
		int32(request.WithdrawAmount),
		int32(r.dailyLimit),
		request.IdempotencyKey,
	)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" && request.IdempotencyKey != "" {
			// Concurrent duplicate: the whole statement rolled back (no balance
			// change), return the existing record.
			existing, gerr := r.getWithdrawByIdempotencyKey(ctx, request.IdempotencyKey)
			if gerr == nil {
				if validationErr := validateWithdrawReplay(existing, request); validationErr != nil {
					return nil, validationErr
				}
				return &WithdrawAtomicResult{Row: toWithdrawAtomicRow(existing), Replayed: true}, nil
			}
		}
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sharedErrors.NewBadRequestError("insufficient balance")
		}
		return nil, sharedErrors.ErrConstraintOrFailed(err, "Withdraw", "create withdraw")
	}
	return &WithdrawAtomicResult{Row: &withdraw}, nil
}

func (r *withdrawCommandRepository) getWithdrawByIdempotencyKey(ctx context.Context, idempotencyKey string) (*db.GetWithdrawByIdempotencyKeyRow, error) {
	var withdraw db.GetWithdrawByIdempotencyKeyRow
	err := r.db.GetContext(ctx, &withdraw, getWithdrawByIdempotencyKey, idempotencyKey)
	if err != nil {
		return nil, err
	}
	return &withdraw, nil
}

func validateWithdrawReplay(existing *db.GetWithdrawByIdempotencyKeyRow, request *requests.CreateWithdrawRequest) error {
	if existing == nil || existing.CardNumber != request.CardNumber || existing.WithdrawAmount != int32(request.WithdrawAmount) || !sameWallClock(existing.WithdrawTime, request.WithdrawTime) {
		return sharedErrors.ErrConflict.WithMessage("withdraw idempotency key already used with different payload")
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

func toWithdrawAtomicRow(r *db.GetWithdrawByIdempotencyKeyRow) *db.CreateWithdrawAtomicRow {
	if r == nil {
		return nil
	}
	return &db.CreateWithdrawAtomicRow{
		WithdrawID:     r.WithdrawID,
		WithdrawNo:     r.WithdrawNo,
		CardNumber:     r.CardNumber,
		WithdrawAmount: r.WithdrawAmount,
		WithdrawTime:   r.WithdrawTime,
		Status:         r.Status,
		CreatedAt:      r.CreatedAt,
		UpdatedAt:      r.UpdatedAt,
	}
}

// UpdateSaldoBalanceDelta applies an atomic signed delta to a card balance.
// It returns an error (400) when the resulting balance would be negative.
func (r *withdrawCommandRepository) UpdateSaldoBalanceDelta(ctx context.Context, cardNumber string, delta int) error {
	var saldo db.UpdateSaldoBalanceDeltaRow
	err := r.db.GetContext(ctx, &saldo, updateSaldoBalanceDelta,
		cardNumber,
		int32(delta),
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return sharedErrors.NewBadRequestError("insufficient balance for update")
		}
		return sharedErrors.ErrNoRowsOrFailed(err, "Withdraw", "update withdraw balance")
	}
	return nil
}

// UpdateWithdrawAtomic updates a withdrawal and settles its card balance in one
// SQL statement. The delta is computed from the locked pre-update amount inside
// the statement, so concurrent updates can never read a stale amount (no lost
// update, no negative double-apply). Returns 400 when the resulting balance
// would be negative or the card has no saldo.
func (r *withdrawCommandRepository) UpdateWithdrawAtomic(ctx context.Context, request *requests.UpdateWithdrawRequest) (*db.UpdateWithdrawAtomicRow, error) {
	var withdraw db.UpdateWithdrawAtomicRow
	err := r.db.GetContext(ctx, &withdraw, updateWithdrawAtomic,
		int32(*request.WithdrawID),
		int32(request.WithdrawAmount),
		request.CardNumber,
		request.WithdrawTime,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sharedErrors.NewBadRequestError("insufficient balance for withdraw update")
		}
		return nil, sharedErrors.ErrNoRowsOrFailed(err, "Withdraw", "update withdraw balance")
	}
	return &withdraw, nil
}

func (r *withdrawCommandRepository) UpdateWithdraw(ctx context.Context, request *requests.UpdateWithdrawRequest) (*db.UpdateWithdrawRow, error) {
	var withdraw db.UpdateWithdrawRow
	err := r.db.GetContext(ctx, &withdraw, updateWithdraw,
		int32(*request.WithdrawID),
		request.CardNumber,
		int32(request.WithdrawAmount),
		request.WithdrawTime,
	)

	if err != nil {
		return nil, sharedErrors.ErrNoRowsOrFailed(err, "Withdraw", "update withdraw")
	}

	return &withdraw, nil
}

func (r *withdrawCommandRepository) UpdateWithdrawStatus(ctx context.Context, request *requests.UpdateWithdrawStatus) (*db.UpdateWithdrawStatusRow, error) {
	var withdraw db.UpdateWithdrawStatusRow
	err := r.db.GetContext(ctx, &withdraw, updateWithdrawStatus,
		int32(request.WithdrawID),
		request.Status,
	)

	if err != nil {
		return nil, sharedErrors.ErrNoRowsOrFailed(err, "Withdraw", "update withdraw status")
	}

	return &withdraw, nil
}

func (r *withdrawCommandRepository) TrashedWithdraw(ctx context.Context, withdraw_id int) (*db.Withdraw, error) {
	var row db.TrashWithdrawRow
	err := r.db.GetContext(ctx, &row, trashWithdraw, int32(withdraw_id))

	if err != nil {
		return nil, sharedErrors.ErrNoRowsOrFailed(err, "Withdraw", "trash withdraw")
	}

	return &db.Withdraw{
		WithdrawID:     row.WithdrawID,
		WithdrawNo:     row.WithdrawNo,
		CardNumber:     row.CardNumber,
		WithdrawAmount: row.WithdrawAmount,
		WithdrawTime:   row.WithdrawTime,
		Status:         row.Status,
		CreatedAt:      row.CreatedAt,
		UpdatedAt:      row.UpdatedAt,
		DeletedAt:      row.DeletedAt,
	}, nil
}

func (r *withdrawCommandRepository) RestoreWithdraw(ctx context.Context, withdraw_id int) (*db.Withdraw, error) {
	var row db.RestoreWithdrawRow
	err := r.db.GetContext(ctx, &row, restoreWithdraw, int32(withdraw_id))

	if err != nil {
		return nil, sharedErrors.ErrNoRowsOrFailed(err, "Withdraw", "restore withdraw")
	}

	return &db.Withdraw{
		WithdrawID:     row.WithdrawID,
		WithdrawNo:     row.WithdrawNo,
		CardNumber:     row.CardNumber,
		WithdrawAmount: row.WithdrawAmount,
		WithdrawTime:   row.WithdrawTime,
		Status:         row.Status,
		CreatedAt:      row.CreatedAt,
		UpdatedAt:      row.UpdatedAt,
		DeletedAt:      row.DeletedAt,
	}, nil
}

func (r *withdrawCommandRepository) DeleteWithdrawPermanent(ctx context.Context, withdraw_id int) (bool, error) {
	if _, err := r.db.ExecContext(ctx, deleteWithdrawPermanently, int32(withdraw_id)); err != nil {
		return false, sharedErrors.ErrNoRowsOrFailed(err, "Withdraw", "delete withdraw permanently")
	}

	return true, nil
}

func (r *withdrawCommandRepository) RestoreAllWithdraw(ctx context.Context) (bool, error) {
	if _, err := r.db.ExecContext(ctx, restoreAllWithdraws); err != nil {
		return false, withdraw_errors.ErrRestoreAllWithdrawsFailed.WithInternal(err)
	}

	return true, nil
}

func (r *withdrawCommandRepository) DeleteAllWithdrawPermanent(ctx context.Context) (bool, error) {
	if _, err := r.db.ExecContext(ctx, deleteAllPermanentWithdraws); err != nil {
		return false, withdraw_errors.ErrDeleteAllWithdrawsPermanentFailed.WithInternal(err)
	}

	return true, nil
}
