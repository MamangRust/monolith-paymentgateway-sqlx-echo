// Code migrated from sqlc-generated sources; domain types kept for repository layer.
// source: user.sql

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type CreateUserParams struct {
	Firstname        string `db:"firstname" json:"firstname"`
	Lastname         string `db:"lastname" json:"lastname"`
	Email            string `db:"email" json:"email"`
	Password         string `db:"password" json:"password"`
	VerificationCode string `db:"verification_code" json:"verification_code"`
	IsVerified       *bool  `db:"is_verified" json:"is_verified"`
}

type CreateUserRow struct {
	UserID    int32            `db:"user_id" json:"user_id"`
	Firstname string           `db:"firstname" json:"firstname"`
	Lastname  string           `db:"lastname" json:"lastname"`
	Email     string           `db:"email" json:"email"`
	Password  string           `db:"password" json:"password"`
	CreatedAt pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type GetActiveUsersWithPaginationParams struct {
	Column1 string `db:"column1" json:"column_1"`
	Limit   int32  `db:"limit" json:"limit"`
	Offset  int32  `db:"offset" json:"offset"`
}

type GetActiveUsersWithPaginationRow struct {
	UserID     int32            `db:"user_id" json:"user_id"`
	Firstname  string           `db:"firstname" json:"firstname"`
	Lastname   string           `db:"lastname" json:"lastname"`
	Email      string           `db:"email" json:"email"`
	Password   string           `db:"password" json:"password"`
	CreatedAt  pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt  pgtype.Timestamp `db:"updated_at" json:"updated_at"`
	DeletedAt  pgtype.Timestamp `db:"deleted_at" json:"deleted_at"`
	TotalCount int64            `db:"total_count" json:"total_count"`
}

type GetTrashedUsersWithPaginationParams struct {
	Column1 string `db:"column1" json:"column_1"`
	Limit   int32  `db:"limit" json:"limit"`
	Offset  int32  `db:"offset" json:"offset"`
}

type GetTrashedUsersWithPaginationRow struct {
	UserID     int32            `db:"user_id" json:"user_id"`
	Firstname  string           `db:"firstname" json:"firstname"`
	Lastname   string           `db:"lastname" json:"lastname"`
	Email      string           `db:"email" json:"email"`
	Password   string           `db:"password" json:"password"`
	CreatedAt  pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt  pgtype.Timestamp `db:"updated_at" json:"updated_at"`
	DeletedAt  pgtype.Timestamp `db:"deleted_at" json:"deleted_at"`
	TotalCount int64            `db:"total_count" json:"total_count"`
}

type GetUserByEmailRow struct {
	UserID int32  `db:"user_id" json:"user_id"`
	Email  string `db:"email" json:"email"`
}

type GetUserByEmailAndVerifiedRow struct {
	UserID    int32            `db:"user_id" json:"user_id"`
	Firstname string           `db:"firstname" json:"firstname"`
	Lastname  string           `db:"lastname" json:"lastname"`
	Email     string           `db:"email" json:"email"`
	Password  string           `db:"password" json:"password"`
	CreatedAt pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type GetUserByEmailWithPasswordRow struct {
	UserID   int32  `db:"user_id" json:"user_id"`
	Email    string `db:"email" json:"email"`
	Password string `db:"password" json:"password"`
}

type GetUserByIDRow struct {
	UserID    int32            `db:"user_id" json:"user_id"`
	Firstname string           `db:"firstname" json:"firstname"`
	Lastname  string           `db:"lastname" json:"lastname"`
	Email     string           `db:"email" json:"email"`
	Password  string           `db:"password" json:"password"`
	CreatedAt pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type GetUserByVerificationCodeRow struct {
	UserID    int32            `db:"user_id" json:"user_id"`
	Firstname string           `db:"firstname" json:"firstname"`
	Lastname  string           `db:"lastname" json:"lastname"`
	Email     string           `db:"email" json:"email"`
	Password  string           `db:"password" json:"password"`
	CreatedAt pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type GetUsersWithPaginationParams struct {
	Column1 string `db:"column1" json:"column_1"`
	Limit   int32  `db:"limit" json:"limit"`
	Offset  int32  `db:"offset" json:"offset"`
}

type GetUsersWithPaginationRow struct {
	UserID     int32            `db:"user_id" json:"user_id"`
	Firstname  string           `db:"firstname" json:"firstname"`
	Lastname   string           `db:"lastname" json:"lastname"`
	Email      string           `db:"email" json:"email"`
	Password   string           `db:"password" json:"password"`
	CreatedAt  pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt  pgtype.Timestamp `db:"updated_at" json:"updated_at"`
	TotalCount int64            `db:"total_count" json:"total_count"`
}

type RestoreUserRow struct {
	UserID    int32            `db:"user_id" json:"user_id"`
	Firstname string           `db:"firstname" json:"firstname"`
	Lastname  string           `db:"lastname" json:"lastname"`
	Email     string           `db:"email" json:"email"`
	CreatedAt pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt pgtype.Timestamp `db:"updated_at" json:"updated_at"`
	DeletedAt pgtype.Timestamp `db:"deleted_at" json:"deleted_at"`
}

type TrashUserRow struct {
	UserID    int32            `db:"user_id" json:"user_id"`
	Firstname string           `db:"firstname" json:"firstname"`
	Lastname  string           `db:"lastname" json:"lastname"`
	Email     string           `db:"email" json:"email"`
	CreatedAt pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt pgtype.Timestamp `db:"updated_at" json:"updated_at"`
	DeletedAt pgtype.Timestamp `db:"deleted_at" json:"deleted_at"`
}

type UpdateUserParams struct {
	UserID    int32  `db:"user_id" json:"user_id"`
	Firstname string `db:"firstname" json:"firstname"`
	Lastname  string `db:"lastname" json:"lastname"`
	Email     string `db:"email" json:"email"`
	Password  string `db:"password" json:"password"`
}

type UpdateUserRow struct {
	UserID    int32            `db:"user_id" json:"user_id"`
	Firstname string           `db:"firstname" json:"firstname"`
	Lastname  string           `db:"lastname" json:"lastname"`
	Email     string           `db:"email" json:"email"`
	Password  string           `db:"password" json:"password"`
	CreatedAt pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type UpdateUserIsVerifiedParams struct {
	UserID     int32 `db:"user_id" json:"user_id"`
	IsVerified *bool `db:"is_verified" json:"is_verified"`
}

type UpdateUserIsVerifiedRow struct {
	UserID    int32            `db:"user_id" json:"user_id"`
	Firstname string           `db:"firstname" json:"firstname"`
	Lastname  string           `db:"lastname" json:"lastname"`
	Email     string           `db:"email" json:"email"`
	Password  string           `db:"password" json:"password"`
	CreatedAt pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type UpdateUserPasswordParams struct {
	UserID   int32  `db:"user_id" json:"user_id"`
	Password string `db:"password" json:"password"`
}

type UpdateUserPasswordRow struct {
	UserID    int32            `db:"user_id" json:"user_id"`
	Firstname string           `db:"firstname" json:"firstname"`
	Lastname  string           `db:"lastname" json:"lastname"`
	Email     string           `db:"email" json:"email"`
	Password  string           `db:"password" json:"password"`
	CreatedAt pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}
