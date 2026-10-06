package repository

import (
	"context"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	sharedErrors "github.com/MamangRust/monolith-payment-gateway-shared/errors"
	role_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/role_errors/repository"
	"github.com/jmoiron/sqlx"
)

const createRole = `INSERT INTO
    roles (
        role_name,
        created_at,
        updated_at
    )
VALUES (
        $1,
        current_timestamp,
        current_timestamp
    )
RETURNING
    role_id,
    role_name,
    created_at,
    updated_at,
    deleted_at`

const updateRole = `UPDATE roles
SET
    role_name = $2,
    updated_at = current_timestamp
WHERE
    role_id = $1
RETURNING
    role_id,
    role_name,
    created_at,
    updated_at,
    deleted_at`

const trashRole = `UPDATE roles
SET
    deleted_at = current_timestamp
WHERE
    role_id = $1
RETURNING
    *`

const restoreRole = `UPDATE roles
SET
    deleted_at = NULL
WHERE
    role_id = $1
RETURNING
    *`

const deletePermanentRole = `DELETE FROM roles
WHERE
    role_id = $1
    AND deleted_at IS NOT NULL
RETURNING
    *`

const restoreAllRoles = `UPDATE roles
SET
    deleted_at = NULL
WHERE
    deleted_at IS NOT NULL`

const deleteAllPermanentRoles = `DELETE FROM roles WHERE deleted_at IS NOT NULL`

const assignRoleToUser = `INSERT INTO
    user_roles (
        user_id,
        role_id,
        created_at,
        updated_at
    )
VALUES (
        $1,
        $2,
        current_timestamp,
        current_timestamp
    )
RETURNING
    user_role_id,
    user_id,
    role_id,
    created_at,
    updated_at,
    deleted_at`

// roleCommandRepository is a struct that implements the RoleCommandRepository interface
type roleCommandRepository struct {
	db *sqlx.DB
}

// NewRoleCommandRepository creates a new RoleCommandRepository instance with the provided
// database connection. This repository is responsible for
// executing command operations related to role records in the database.
//
// Parameters:
//   - db: A pointer to the sqlx.DB object for executing database queries.
//
// Returns:
//   - A pointer to the newly created roleCommandRepository instance.
func NewRoleCommandRepository(db *sqlx.DB) RoleCommandRepository {
	return &roleCommandRepository{
		db: db,
	}
}

func (r *roleCommandRepository) CreateRole(ctx context.Context, req *requests.CreateRoleRequest) (*db.Role, error) {
	var role db.Role
	err := r.db.GetContext(ctx, &role, createRole, req.Name)

	if err != nil {
		return nil, sharedErrors.ErrConstraintOrFailed(err, "Role", "create role")
	}

	return &role, nil
}

func (r *roleCommandRepository) UpdateRole(ctx context.Context, req *requests.UpdateRoleRequest) (*db.Role, error) {
	var role db.Role
	err := r.db.GetContext(ctx, &role, updateRole,
		int32(*req.ID),
		req.Name,
	)

	if err != nil {
		return nil, sharedErrors.ErrNoRowsOrFailed(err, "Role", "update role")
	}

	return &role, nil
}

func (r *roleCommandRepository) TrashedRole(ctx context.Context, id int) (*db.Role, error) {
	var role db.Role
	err := r.db.GetContext(ctx, &role, trashRole, int32(id))
	if err != nil {
		return nil, sharedErrors.ErrNoRowsOrFailed(err, "Role", "trash role")
	}
	return &role, nil
}

func (r *roleCommandRepository) RestoreRole(ctx context.Context, id int) (*db.Role, error) {
	var role db.Role
	err := r.db.GetContext(ctx, &role, restoreRole, int32(id))
	if err != nil {
		return nil, sharedErrors.ErrNoRowsOrFailed(err, "Role", "restore role")
	}
	return &role, nil
}

func (r *roleCommandRepository) DeleteRolePermanent(ctx context.Context, role_id int) (bool, error) {
	if _, err := r.db.ExecContext(ctx, deletePermanentRole, int32(role_id)); err != nil {
		return false, sharedErrors.ErrNoRowsOrFailed(err, "Role", "delete role")
	}
	return true, nil
}

func (r *roleCommandRepository) AssignRoleToUser(ctx context.Context, req *requests.CreateUserRoleRequest) (*db.UserRole, error) {
	var userRole db.UserRole
	err := r.db.GetContext(ctx, &userRole, assignRoleToUser,
		int32(req.UserId),
		int32(req.RoleId),
	)
	if err != nil {
		return nil, sharedErrors.ErrConstraintOrFailed(err, "UserRole", "assign role to user")
	}
	return &userRole, nil
}

func (r *roleCommandRepository) RestoreAllRole(ctx context.Context) (bool, error) {
	if _, err := r.db.ExecContext(ctx, restoreAllRoles); err != nil {
		return false, role_errors.ErrRestoreAllRoles.WithInternal(err)
	}

	return true, nil
}

func (r *roleCommandRepository) DeleteAllRolePermanent(ctx context.Context) (bool, error) {
	if _, err := r.db.ExecContext(ctx, deleteAllPermanentRoles); err != nil {
		return false, role_errors.ErrDeleteAllRoles.WithInternal(err)
	}

	return true, nil
}
