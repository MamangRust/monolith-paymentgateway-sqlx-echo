package repository

import (
	"context"
	"time"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	sharedErrors "github.com/MamangRust/monolith-payment-gateway-shared/errors"
	refresh_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/refresh_token_errors/repository"
	"github.com/jmoiron/sqlx"
)

// resetTokenRepository is a struct that implements the ResetTokenRepository interface
type resetTokenRepository struct {
	db *sqlx.DB
}

// NewResetTokenRepository creates a new instance of resetTokenRepository.
func NewResetTokenRepository(db *sqlx.DB) *resetTokenRepository {
	return &resetTokenRepository{
		db: db,
	}
}

const getResetToken = `SELECT user_id, token, expiry_date
FROM reset_tokens
WHERE
    token = $1`

const createResetToken = `WITH inserted_token AS (
    INSERT INTO reset_tokens (user_id, token, expiry_date)
    VALUES ($1, $2, $3)
    RETURNING user_id, token, expiry_date
), queued_event AS (
    INSERT INTO outbox_events (event_key, topic, message_key, payload)
    SELECT
        'auth-forgot-password:' || it.user_id::text || ':' || it.token,
        'email-service-topic-auth-forgot-password',
        it.user_id::text,
        jsonb_build_object(
            'event_id', gen_random_uuid()::text,
            'schema_version', 1,
            'event_type', 'auth.forgot_password',
            'occurred_at', now(),
            'email', u.email,
            'subject', 'Password Reset Request',
            'body', 'Click the reset link to change your password: https://sanedge.example.com/reset-password?token=' || it.token
        )
    FROM inserted_token it
    JOIN users u ON u.user_id = it.user_id AND u.deleted_at IS NULL
    ON CONFLICT (event_key) DO NOTHING
    RETURNING event_key
)
SELECT user_id, token, expiry_date
FROM inserted_token`

const deleteResetToken = `DELETE FROM reset_tokens WHERE user_id = $1`

// FindByToken retrieves a reset token record by token string.
func (r *resetTokenRepository) FindByToken(ctx context.Context, code string) (*db.GetResetTokenRow, error) {
	var token db.GetResetTokenRow
	if err := r.db.GetContext(ctx, &token, getResetToken, code); err != nil {
		return nil, sharedErrors.ErrNoRowsOrFailed(err, "Reset Token", "find reset token")
	}
	return &token, nil
}

// CreateResetToken inserts a new reset token into the database.
func (r *resetTokenRepository) CreateResetToken(ctx context.Context, req *requests.CreateResetTokenRequest) (*db.CreateResetTokenRow, error) {
	expiryDate, err := time.Parse("2006-01-02 15:04:05", req.ExpiredAt)
	if err != nil {
		return nil, sharedErrors.ErrInternal.WithInternal(err)
	}
	var token db.CreateResetTokenRow
	if err := r.db.GetContext(ctx, &token, createResetToken,
		int64(req.UserID),
		req.ResetToken,
		expiryDate,
	); err != nil {
		return nil, refresh_errors.ErrCreateRefreshToken.WithInternal(err)
	}
	return &token, nil
}

// DeleteResetToken removes the reset token associated with the given user ID.
func (r *resetTokenRepository) DeleteResetToken(ctx context.Context, user_id int) error {
	if _, err := r.db.ExecContext(ctx, deleteResetToken, int64(user_id)); err != nil {
		return refresh_errors.ErrDeleteByUserID.WithInternal(err)
	}
	return nil
}
