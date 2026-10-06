package repository

import (
	"context"
	"database/sql"
	"errors"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	sharedErrors "github.com/MamangRust/monolith-payment-gateway-shared/errors"
	role_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/role_errors/repository"
	"github.com/jmoiron/sqlx"
)

const getRoles = `SELECT
    role_id,
    role_name,
    created_at,
    updated_at,
    COUNT(*) OVER () AS total_count
FROM roles
WHERE
    $1::TEXT IS NULL
    OR role_name ILIKE '%' || $1 || '%'
ORDER BY created_at ASC
LIMIT $2
OFFSET
    $3`

const getRole = `SELECT
    role_id,
    role_name,
    created_at,
    updated_at,
    deleted_at
FROM roles
WHERE
    role_id = $1`

const getRoleByName = `SELECT
    role_id,
    role_name,
    created_at,
    updated_at,
    deleted_at
FROM roles
WHERE
    role_name = $1`

const getUserRoles = `SELECT r.role_id, r.role_name, r.created_at, r.updated_at, r.deleted_at
FROM roles r
    JOIN user_roles ur ON ur.role_id = r.role_id
WHERE
    ur.user_id = $1
ORDER BY r.created_at ASC`

const getActiveRoles = `SELECT
    role_id,
    role_name,
    created_at,
    updated_at,
    deleted_at,
    COUNT(*) OVER () AS total_count
FROM roles
WHERE
    deleted_at IS NULL
    AND (
        $1::TEXT IS NULL
        OR role_name ILIKE '%' || $1 || '%'
    )
ORDER BY created_at ASC
LIMIT $2
OFFSET
    $3`

const getTrashedRoles = `SELECT
    role_id,
    role_name,
    created_at,
    updated_at,
    deleted_at,
    COUNT(*) OVER () AS total_count
FROM roles
WHERE
    deleted_at IS NOT NULL
    AND (
        $1::TEXT IS NULL
        OR role_name ILIKE '%' || $1 || '%'
    )
ORDER BY deleted_at DESC
LIMIT $2
OFFSET
    $3`

// roleQueryRepository is a struct that implements the RoleQueryRepository interface
type roleQueryRepository struct {
	db *sqlx.DB
}

func NewRoleQueryRepository(db *sqlx.DB) RoleQueryRepository {
	return &roleQueryRepository{
		db: db,
	}
}

func (r *roleQueryRepository) FindAllRoles(ctx context.Context, req *requests.FindAllRoles) ([]*db.GetRolesRow, error) {
	offset := (req.Page - 1) * req.PageSize

	var roles []*db.GetRolesRow
	err := r.db.SelectContext(ctx, &roles, getRoles,
		req.Search,
		int32(req.PageSize),
		int32(offset),
	)

	if err != nil {
		return nil, role_errors.ErrFindAllRoles.WithInternal(err)
	}

	return roles, nil
}

func (r *roleQueryRepository) FindById(ctx context.Context, id int) (*db.Role, error) {
	if id <= 0 {
		return nil, sharedErrors.NewBadRequestError("role ID must be greater than zero")
	}

	var role db.Role
	err := r.db.GetContext(ctx, &role, getRole, int32(id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, role_errors.ErrRoleNotFound.WithInternal(err)
		}
		return nil, sharedErrors.ErrInternal.WithInternal(err)
	}
	return &role, nil
}

func (r *roleQueryRepository) FindByName(ctx context.Context, name string) (*db.Role, error) {
	var role db.Role
	err := r.db.GetContext(ctx, &role, getRoleByName, name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, role_errors.ErrRoleNotFound.WithInternal(err)
		}

		return nil, sharedErrors.ErrInternal.WithInternal(err)
	}
	return &role, nil
}

func (r *roleQueryRepository) FindByUserId(ctx context.Context, user_id int) ([]*db.Role, error) {
	var roles []*db.Role
	err := r.db.SelectContext(ctx, &roles, getUserRoles, int32(user_id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, role_errors.ErrRoleNotFound.WithInternal(err)
		}

		return nil, sharedErrors.ErrInternal.WithInternal(err)
	}
	return roles, nil
}

func (r *roleQueryRepository) FindByActiveRole(ctx context.Context, req *requests.FindAllRoles) ([]*db.GetActiveRolesRow, error) {
	offset := (req.Page - 1) * req.PageSize

	var roles []*db.GetActiveRolesRow
	err := r.db.SelectContext(ctx, &roles, getActiveRoles,
		req.Search,
		int32(req.PageSize),
		int32(offset),
	)

	if err != nil {
		return nil, role_errors.ErrFindActiveRoles.WithInternal(err)
	}

	return roles, nil
}

func (r *roleQueryRepository) FindByTrashedRole(ctx context.Context, req *requests.FindAllRoles) ([]*db.GetTrashedRolesRow, error) {
	offset := (req.Page - 1) * req.PageSize

	var roles []*db.GetTrashedRolesRow
	err := r.db.SelectContext(ctx, &roles, getTrashedRoles,
		req.Search,
		int32(req.PageSize),
		int32(offset),
	)

	if err != nil {
		return nil, role_errors.ErrFindTrashedRoles.WithInternal(err)
	}

	return roles, nil
}
