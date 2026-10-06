// Code migrated from sqlc-generated sources; domain types kept for repository layer.
// source: refresh_token.sql

package db

import (
	"time"
)

type CreateRefreshTokenParams struct {
	UserID     int32     `db:"user_id" json:"user_id"`
	Token      string    `db:"token" json:"token"`
	Expiration time.Time `db:"expiration" json:"expiration"`
}

type UpdateRefreshTokenByUserIdParams struct {
	UserID     int32     `db:"user_id" json:"user_id"`
	Token      string    `db:"token" json:"token"`
	Expiration time.Time `db:"expiration" json:"expiration"`
}
