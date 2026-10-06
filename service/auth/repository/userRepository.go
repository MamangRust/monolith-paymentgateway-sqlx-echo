package repository

import (
	"context"
	"database/sql"
	"errors"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	sharedErrors "github.com/MamangRust/monolith-payment-gateway-shared/errors"
	user_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/user_errors/repository"
	"github.com/jmoiron/sqlx"
)

// userRepository is a struct that represents a user repository
type userRepository struct {
	db *sqlx.DB
}

// NewUserRepository returns a new instance of userRepository.
func NewUserRepository(db *sqlx.DB) *userRepository {
	return &userRepository{
		db: db,
	}
}

const getUserByID = `SELECT
    user_id,
    firstname,
    lastname,
    email,
    password,
    created_at,
    updated_at
FROM users
WHERE
    user_id = $1
    AND deleted_at IS NULL`

const getUserByEmail = `SELECT user_id, email
FROM users
WHERE
    email = $1
    AND deleted_at IS NULL`

const getUserByEmailAndVerified = `SELECT
    user_id,
    firstname,
    lastname,
    email,
    password,
    created_at,
    updated_at
FROM users
WHERE
    email = $1
    AND is_verified = true
    AND deleted_at IS NULL`

const getUserByVerificationCode = `SELECT
    user_id,
    firstname,
    lastname,
    email,
    password,
    created_at,
    updated_at
FROM users
WHERE
    verification_code = $1`

const createUser = `WITH inserted_user AS (
    INSERT INTO users (
        firstname,
        lastname,
        email,
        password,
        verification_code,
        is_verified,
        created_at,
        updated_at
    )
    VALUES ($1, $2, $3, $4, $5, $6, current_timestamp, current_timestamp)
    RETURNING user_id, firstname, lastname, email, password, verification_code, created_at, updated_at
), queued_event AS (
    INSERT INTO outbox_events (event_key, topic, message_key, payload)
    SELECT
        'auth-register:' || user_id::text,
        'email-service-topic-auth-register',
        user_id::text,
        jsonb_build_object(
            'event_id', gen_random_uuid()::text,
            'schema_version', 1,
            'event_type', 'auth.register',
            'occurred_at', now(),
            'email', email,
            'subject', 'Welcome to SanEdge',
            'body', 'Your account has been successfully created. Verify your account here: https://sanedge.example.com/login?verify_code=' || verification_code
        )
    FROM inserted_user
    ON CONFLICT (event_key) DO NOTHING
    RETURNING event_key
)
SELECT user_id, firstname, lastname, email, password, created_at, updated_at
FROM inserted_user`

const updateUserIsVerified = `WITH updated_user AS (
    UPDATE users
    SET is_verified = $2, updated_at = current_timestamp
    WHERE user_id = $1 AND deleted_at IS NULL
    RETURNING user_id, firstname, lastname, email, password, created_at, updated_at
), queued_event AS (
    INSERT INTO outbox_events (event_key, topic, message_key, payload)
    SELECT
        'auth-verify:' || user_id::text,
        'email-service-topic-auth-verify-code-success',
        user_id::text,
        jsonb_build_object(
            'event_id', gen_random_uuid()::text,
            'schema_version', 1,
            'event_type', 'auth.verify_code_success',
            'occurred_at', now(),
            'email', email,
            'subject', 'Verification Success',
            'body', 'Your account has been successfully verified.'
        )
    FROM updated_user
    WHERE $2 = true
    ON CONFLICT (event_key) DO NOTHING
    RETURNING event_key
)
SELECT user_id, firstname, lastname, email, password, created_at, updated_at
FROM updated_user`

const updateUserPassword = `UPDATE users
SET
    password = $2,
    updated_at = current_timestamp
WHERE
    user_id = $1
    AND deleted_at IS NULL
RETURNING
    user_id,
    firstname,
    lastname,
    email,
    password,
    created_at,
    updated_at`

// FindById retrieves a user by their unique ID.
func (r *userRepository) FindById(ctx context.Context, user_id int) (*db.GetUserByIDRow, error) {
	if user_id <= 0 {
		return nil, sharedErrors.NewBadRequestError("user ID must be greater than zero")
	}

	var user db.GetUserByIDRow
	err := r.db.GetContext(ctx, &user, getUserByID, int32(user_id))

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, user_errors.ErrUserNotFound.WithInternal(err)
		}

		return nil, sharedErrors.ErrInternal.WithInternal(err)
	}

	return &user, nil
}

// FindByEmail retrieves a user by their email address.
func (r *userRepository) FindByEmail(ctx context.Context, email string) (*db.GetUserByEmailRow, error) {
	var user db.GetUserByEmailRow
	err := r.db.GetContext(ctx, &user, getUserByEmail, email)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, user_errors.ErrUserNotFound.WithInternal(err)
		}

		return nil, sharedErrors.ErrInternal.WithInternal(err)
	}

	return &user, nil
}

// FindByEmailAndVerify retrieves a verified user by their email address.
func (r *userRepository) FindByEmailAndVerify(ctx context.Context, email string) (*db.GetUserByEmailAndVerifiedRow, error) {
	var user db.GetUserByEmailAndVerifiedRow
	err := r.db.GetContext(ctx, &user, getUserByEmailAndVerified, email)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, user_errors.ErrUserNotFound.WithInternal(err)
		}

		return nil, sharedErrors.ErrInternal.WithInternal(err)
	}

	return &user, nil
}

// FindByVerificationCode retrieves a user by their verification code.
func (r *userRepository) FindByVerificationCode(ctx context.Context, verification_code string) (*db.GetUserByVerificationCodeRow, error) {
	var user db.GetUserByVerificationCodeRow
	err := r.db.GetContext(ctx, &user, getUserByVerificationCode, verification_code)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, sharedErrors.ErrInternal.WithInternal(err)
	}

	return &user, nil
}

// CreateUser inserts a new user into the database.
func (r *userRepository) CreateUser(ctx context.Context, request *requests.RegisterRequest) (*db.CreateUserRow, error) {
	isVerified := request.IsVerified

	var user db.CreateUserRow
	err := r.db.GetContext(ctx, &user, createUser,
		request.FirstName,
		request.LastName,
		request.Email,
		request.Password,
		request.VerifiedCode,
		&isVerified,
	)

	if err != nil {
		return nil, user_errors.ErrCreateUser.WithInternal(err)
	}

	return &user, nil
}

// UpdateUserIsVerified updates the verification status of a user.
func (r *userRepository) UpdateUserIsVerified(ctx context.Context, user_id int, is_verified bool) (*db.UpdateUserIsVerifiedRow, error) {
	var user db.UpdateUserIsVerifiedRow
	err := r.db.GetContext(ctx, &user, updateUserIsVerified,
		int32(user_id),
		&is_verified,
	)

	if err != nil {
		return nil, sharedErrors.ErrNoRowsOrFailed(err, "User", "update user verification")
	}

	return &user, nil
}

// UpdateUserPassword updates a user's password.
func (r *userRepository) UpdateUserPassword(ctx context.Context, user_id int, password string) (*db.UpdateUserPasswordRow, error) {
	var user db.UpdateUserPasswordRow
	err := r.db.GetContext(ctx, &user, updateUserPassword,
		int32(user_id),
		password,
	)

	if err != nil {
		return nil, sharedErrors.ErrNoRowsOrFailed(err, "User", "update user password")
	}

	return &user, nil
}
