package seeder

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

const seedCreateWithdraw = `INSERT INTO
    withdraws (
        card_number,
        withdraw_amount,
        withdraw_time,
        created_at,
        updated_at
    )
VALUES (
        $1,
        $2,
        $3,
        $4,
        current_timestamp
    )
RETURNING
    withdraw_id,
    withdraw_no,
    card_number,
    withdraw_amount,
    withdraw_time,
    status,
    created_at,
    updated_at
`

const seedUpdateWithdrawStatus = `UPDATE withdraws
SET
    status = $2,
    updated_at = current_timestamp
WHERE
    withdraw_id = $1
    AND deleted_at IS NULL
    AND (
        ($2 = 'failed' AND status IN ('pending', 'failed', 'compensation_required'))
        OR ($2 = 'compensation_required' AND status IN ('pending', 'failed', 'success', 'compensation_required'))
        OR ($2 = 'success' AND status IN ('pending', 'failed', 'success', 'compensation_required'))
        OR ($2 = 'pending' AND status IN ('failed', 'compensation_required'))
    )
RETURNING
    withdraw_id,
    withdraw_no,
    card_number,
    withdraw_amount,
    withdraw_time,
    status,
    created_at,
    updated_at
`

const seedTrashWithdraw = `UPDATE withdraws
SET
    deleted_at = current_timestamp
WHERE
    withdraw_id = $1
    AND deleted_at IS NULL
RETURNING
    withdraw_id,
    withdraw_no,
    card_number,
    withdraw_amount,
    withdraw_time,
    status,
    created_at,
    updated_at,
    deleted_at
`

// withdrawSeeder is a struct that represents a seeder for the withdraws table.
type withdrawSeeder struct {
	db     *sqlx.DB
	ctx    context.Context
	logger logger.LoggerInterface
}

// NewWithdrawSeeder creates a new instance of the withdrawSeeder, which is
// responsible for populating the withdraws table with fake data.
//
// Args:
// db: a pointer to the database pool
// ctx: a context.Context object
// logger: a logger.LoggerInterface object
//
// Returns:
// a pointer to the withdrawSeeder struct
func NewWithdrawSeeder(db *sqlx.DB, ctx context.Context, logger logger.LoggerInterface) *withdrawSeeder {
	return &withdrawSeeder{
		db:     db,
		ctx:    ctx,
		logger: logger,
	}
}

// Seed populates the withdraws table with fake data.
//
// It generates a total of 10 withdraws, with 5 of them being active and 5 of them being
// trashed. The withdraws are seeded with random card numbers, withdraw amounts, and
// withdraw times. The status of each withdraw is randomly set to one of the following:
// pending, success, or failed. The card numbers are randomly selected from the cards
// table.
//
// If any errors occur during the seeding process, an error is returned.
//
// Returns:
// an error if any of the withdraws fail to be created or updated, otherwise nil
func (r *withdrawSeeder) Seed() error {
	total := 10
	active := 5
	trashed := total - active

	var cards []db.GetCardByUserIDRow
	for i := 1; i <= total; i++ {
		var card db.GetCardByUserIDRow
		err := r.db.GetContext(r.ctx, &card, seedGetCardByUserID, int32(i))
		if err != nil {
			r.logger.Debug("failed to get card for user", zap.Int("userID", i), zap.Error(err))
			continue
		}
		cards = append(cards, card)
	}

	if len(cards) < total {
		r.logger.Error("not enough cards for withdraw seeding", zap.Int("required", total), zap.Int("available", len(cards)))
		return fmt.Errorf("not enough cards: required %d, got %d", total, len(cards))
	}

	statusOptions := []string{"pending", "success", "failed"}

	months := make([]time.Time, 12)
	currentYear := time.Now().Year()
	for i := 0; i < 12; i++ {
		months[i] = time.Date(currentYear, time.Month(i+1), 1, 0, 0, 0, 0, time.UTC)
	}

	for i := 0; i < total; i++ {
		card := cards[i]
		status := statusOptions[rand.Intn(len(statusOptions))]

		monthIndex := i % 12
		withdrawTime := months[monthIndex].Add(time.Duration(rand.Intn(28)) * 24 * time.Hour)

		req := db.CreateWithdrawParams{
			CardNumber:     card.CardNumber,
			WithdrawAmount: int32(rand.Intn(1000000) + 50000),
			WithdrawTime:   withdrawTime,
		}

		var withdraw db.CreateWithdrawRow
		err := r.db.GetContext(r.ctx, &withdraw, seedCreateWithdraw,
			req.CardNumber,
			req.WithdrawAmount,
			req.WithdrawTime,
			req.CreatedAt,
		)
		if err != nil {
			r.logger.Error("failed to seed withdraw", zap.Int("index", i), zap.Error(err))
			return fmt.Errorf("failed to create withdraw %d: %w", i, err)
		}

		var updated db.UpdateWithdrawStatusRow
		err = r.db.GetContext(r.ctx, &updated, seedUpdateWithdrawStatus, withdraw.WithdrawID, status)
		if err != nil {
			r.logger.Error("failed to update withdraw status", zap.Int("withdraw.id", int(withdraw.WithdrawID)), zap.Error(err))
			return fmt.Errorf("failed to update status: %w", err)
		}

		if i >= active {
			var trashed db.TrashWithdrawRow
			err = r.db.GetContext(r.ctx, &trashed, seedTrashWithdraw, withdraw.WithdrawID)
			if err != nil {
				r.logger.Error("failed to trash withdraw", zap.Int("withdraw.id", int(withdraw.WithdrawID)), zap.Error(err))
				return fmt.Errorf("failed to trash withdraw %d: %w", i, err)
			}
		}
	}

	r.logger.Info("withdraw seeding completed",
		zap.Int("total", total),
		zap.Int("active", active),
		zap.Int("trashed", trashed))

	return nil
}
