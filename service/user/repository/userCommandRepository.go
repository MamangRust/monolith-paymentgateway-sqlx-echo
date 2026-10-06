package repository

import (
	"context"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	sharedErrors "github.com/MamangRust/monolith-payment-gateway-shared/errors"
	user_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/user_errors/repository"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// userCommandRepository is a struct that implements the UserCommandRepository interface
type userCommandRepository struct {
	db *sqlx.DB
}

// NewUserCommandRepository creates a new instance of userCommandRepository.
func NewUserCommandRepository(db *sqlx.DB) UserCommandRepository {
	return &userCommandRepository{
		db: db,
	}
}

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

const updateUser = `UPDATE users
SET
    firstname = $2,
    lastname = $3,
    email = $4,
    password = $5,
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

const trashUser = `UPDATE users
SET
    deleted_at = current_timestamp,
    updated_at = current_timestamp
WHERE
    user_id = $1
    AND deleted_at IS NULL
RETURNING
    user_id,
    firstname,
    lastname,
    email,
    created_at,
    updated_at,
    deleted_at`

const restoreUser = `UPDATE users
SET
    deleted_at = NULL,
    updated_at = current_timestamp
WHERE
    user_id = $1
    AND deleted_at IS NOT NULL
RETURNING
    user_id,
    firstname,
    lastname,
    email,
    created_at,
    updated_at,
    deleted_at`

const deleteUserPermanently = `DELETE FROM users WHERE user_id = $1 AND deleted_at IS NOT NULL`

const restoreAllUsers = `UPDATE users
SET
    deleted_at = NULL
WHERE
    deleted_at IS NOT NULL`

const deleteAllPermanentUsers = `DELETE FROM users WHERE deleted_at IS NOT NULL`

// CreateUser inserts a new user record into the database.
func (r *userCommandRepository) CreateUser(ctx context.Context, request *requests.CreateUserRequest) (*db.CreateUserRow, error) {
	verified := false
	verifyCode := uuid.New().String()

	var user db.CreateUserRow
	err := r.db.GetContext(ctx, &user, createUser,
		request.FirstName,
		request.LastName,
		request.Email,
		request.Password,
		verifyCode,
		&verified,
	)

	if err != nil {
		return nil, sharedErrors.ErrConstraintOrFailed(err, "User", "create user")
	}

	return &user, nil
}

// UpdateUser updates an existing user record in the database.
func (r *userCommandRepository) UpdateUser(ctx context.Context, request *requests.UpdateUserRequest) (*db.UpdateUserRow, error) {
	var user db.UpdateUserRow
	err := r.db.GetContext(ctx, &user, updateUser,
		int32(*request.UserID),
		request.FirstName,
		request.LastName,
		request.Email,
		request.Password,
	)

	if err != nil {
		return nil, sharedErrors.ErrNoRowsOrFailed(err, "User", "update user")
	}

	return &user, nil
}

// TrashedUser soft-deletes a user by marking it as trashed.
func (r *userCommandRepository) TrashedUser(ctx context.Context, user_id int) (*db.TrashUserRow, error) {
	var user db.TrashUserRow
	err := r.db.GetContext(ctx, &user, trashUser, int32(user_id))

	if err != nil {
		return nil, sharedErrors.ErrNoRowsOrFailed(err, "User", "trash user")
	}

	return &user, nil
}

// RestoreUser restores a soft-deleted (trashed) user.
func (r *userCommandRepository) RestoreUser(ctx context.Context, user_id int) (*db.RestoreUserRow, error) {
	var user db.RestoreUserRow
	err := r.db.GetContext(ctx, &user, restoreUser, int32(user_id))

	if err != nil {
		return nil, sharedErrors.ErrNoRowsOrFailed(err, "User", "restore user")
	}

	return &user, nil
}

// DeleteUserPermanent permanently deletes a user from the database.
func (r *userCommandRepository) DeleteUserPermanent(ctx context.Context, user_id int) (bool, error) {
	if _, err := r.db.ExecContext(ctx, deleteUserPermanently, int32(user_id)); err != nil {
		return false, sharedErrors.ErrNoRowsOrFailed(err, "User", "delete user")
	}

	return true, nil
}

// RestoreAllUser restores all soft-deleted users in the database.
func (r *userCommandRepository) RestoreAllUser(ctx context.Context) (bool, error) {
	if _, err := r.db.ExecContext(ctx, restoreAllUsers); err != nil {
		return false, user_errors.ErrRestoreAllUsers.WithInternal(err)
	}

	return true, nil
}

// DeleteAllUserPermanent permanently deletes all trashed users from the database.
func (r *userCommandRepository) DeleteAllUserPermanent(ctx context.Context) (bool, error) {
	if _, err := r.db.ExecContext(ctx, deleteAllPermanentUsers); err != nil {
		return false, user_errors.ErrDeleteAllUsers.WithInternal(err)
	}
	return true, nil
}
