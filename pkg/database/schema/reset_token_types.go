// Code migrated from sqlc-generated sources; domain types kept for repository layer.
// source: reset_token.sql

package db

import (
	"time"
)

type CreateResetTokenParams struct {
	UserID     int64     `db:"user_id" json:"user_id"`
	Token      string    `db:"token" json:"token"`
	ExpiryDate time.Time `db:"expiry_date" json:"expiry_date"`
}

type CreateResetTokenRow struct {
	UserID     int64     `db:"user_id" json:"user_id"`
	Token      string    `db:"token" json:"token"`
	ExpiryDate time.Time `db:"expiry_date" json:"expiry_date"`
}

type GetResetTokenRow struct {
	UserID     int64     `db:"user_id" json:"user_id"`
	Token      string    `db:"token" json:"token"`
	ExpiryDate time.Time `db:"expiry_date" json:"expiry_date"`
}
