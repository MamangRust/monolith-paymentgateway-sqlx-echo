// Code migrated from sqlc-generated sources; domain types kept for repository layer.
// source: role.sql

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type GetActiveRolesParams struct {
	Column1 string `db:"column1" json:"column_1"`
	Limit   int32  `db:"limit" json:"limit"`
	Offset  int32  `db:"offset" json:"offset"`
}

type GetActiveRolesRow struct {
	RoleID     int32            `db:"role_id" json:"role_id"`
	RoleName   string           `db:"role_name" json:"role_name"`
	CreatedAt  pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt  pgtype.Timestamp `db:"updated_at" json:"updated_at"`
	DeletedAt  pgtype.Timestamp `db:"deleted_at" json:"deleted_at"`
	TotalCount int64            `db:"total_count" json:"total_count"`
}

type GetRolesParams struct {
	Column1 string `db:"column1" json:"column_1"`
	Limit   int32  `db:"limit" json:"limit"`
	Offset  int32  `db:"offset" json:"offset"`
}

type GetRolesRow struct {
	RoleID     int32            `db:"role_id" json:"role_id"`
	RoleName   string           `db:"role_name" json:"role_name"`
	CreatedAt  pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt  pgtype.Timestamp `db:"updated_at" json:"updated_at"`
	TotalCount int64            `db:"total_count" json:"total_count"`
}

type GetTrashedRolesParams struct {
	Column1 string `db:"column1" json:"column_1"`
	Limit   int32  `db:"limit" json:"limit"`
	Offset  int32  `db:"offset" json:"offset"`
}

type GetTrashedRolesRow struct {
	RoleID     int32            `db:"role_id" json:"role_id"`
	RoleName   string           `db:"role_name" json:"role_name"`
	CreatedAt  pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt  pgtype.Timestamp `db:"updated_at" json:"updated_at"`
	DeletedAt  pgtype.Timestamp `db:"deleted_at" json:"deleted_at"`
	TotalCount int64            `db:"total_count" json:"total_count"`
}

type UpdateRoleParams struct {
	RoleID   int32  `db:"role_id" json:"role_id"`
	RoleName string `db:"role_name" json:"role_name"`
}
