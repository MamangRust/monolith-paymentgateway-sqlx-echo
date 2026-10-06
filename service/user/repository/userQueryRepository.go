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

// userQueryRepository implements UserQueryRepository.
type userQueryRepository struct {
	db *sqlx.DB
}

// NewUserQueryRepository creates a new UserQueryRepository.
func NewUserQueryRepository(db *sqlx.DB) UserQueryRepository {
	return &userQueryRepository{
		db: db,
	}
}

const getUsersWithPagination = `SELECT
    user_id,
    firstname,
    lastname,
    email,
    password,
    created_at,
    updated_at,
    COUNT(*) OVER () AS total_count
FROM users
WHERE
    deleted_at IS NULL
    AND (
        $1::TEXT IS NULL
        OR firstname ILIKE '%' || $1 || '%'
        OR lastname ILIKE '%' || $1 || '%'
        OR email ILIKE '%' || $1 || '%'
    )
ORDER BY created_at DESC
LIMIT $2
OFFSET
    $3`

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

const getActiveUsersWithPagination = `SELECT
    user_id,
    firstname,
    lastname,
    email,
    password,
    created_at,
    updated_at,
    deleted_at,
    COUNT(*) OVER () AS total_count
FROM users
WHERE
    deleted_at IS NULL
    AND (
        $1::TEXT IS NULL
        OR firstname ILIKE '%' || $1 || '%'
        OR lastname ILIKE '%' || $1 || '%'
        OR email ILIKE '%' || $1 || '%'
    )
ORDER BY created_at DESC
LIMIT $2
OFFSET
    $3`

const getTrashedUsersWithPagination = `SELECT
    user_id,
    firstname,
    lastname,
    email,
    password,
    created_at,
    updated_at,
    deleted_at,
    COUNT(*) OVER () AS total_count
FROM users
WHERE
    deleted_at IS NOT NULL
    AND (
        $1::TEXT IS NULL
        OR firstname ILIKE '%' || $1 || '%'
        OR lastname ILIKE '%' || $1 || '%'
        OR email ILIKE '%' || $1 || '%'
    )
ORDER BY created_at DESC
LIMIT $2
OFFSET
    $3`

const getUserByEmail = `SELECT user_id, email
FROM users
WHERE
    email = $1
    AND deleted_at IS NULL`

func (r *userQueryRepository) FindAllUsers(ctx context.Context, req *requests.FindAllUsers) ([]*db.GetUsersWithPaginationRow, error) {
	offset := (req.Page - 1) * req.PageSize

	var users []*db.GetUsersWithPaginationRow
	err := r.db.SelectContext(ctx, &users, getUsersWithPagination,
		req.Search,
		int32(req.PageSize),
		int32(offset),
	)

	if err != nil {
		return nil, user_errors.ErrFindAllUsers.WithInternal(err)
	}

	return users, nil
}

func (r *userQueryRepository) FindById(ctx context.Context, user_id int) (*db.GetUserByIDRow, error) {
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

func (r *userQueryRepository) FindByActive(ctx context.Context, req *requests.FindAllUsers) ([]*db.GetActiveUsersWithPaginationRow, error) {
	offset := (req.Page - 1) * req.PageSize

	var users []*db.GetActiveUsersWithPaginationRow
	err := r.db.SelectContext(ctx, &users, getActiveUsersWithPagination,
		req.Search,
		int32(req.PageSize),
		int32(offset),
	)

	if err != nil {
		return nil, user_errors.ErrFindActiveUsers.WithInternal(err)
	}

	return users, nil
}

func (r *userQueryRepository) FindByTrashed(ctx context.Context, req *requests.FindAllUsers) ([]*db.GetTrashedUsersWithPaginationRow, error) {
	offset := (req.Page - 1) * req.PageSize

	var users []*db.GetTrashedUsersWithPaginationRow
	err := r.db.SelectContext(ctx, &users, getTrashedUsersWithPagination,
		req.Search,
		int32(req.PageSize),
		int32(offset),
	)

	if err != nil {
		return nil, user_errors.ErrFindTrashedUsers.WithInternal(err)
	}

	return users, nil
}

func (r *userQueryRepository) FindByEmail(ctx context.Context, email string) (*db.GetUserByEmailRow, error) {
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
