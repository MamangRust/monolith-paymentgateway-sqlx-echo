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

const seedCreateTopup = `INSERT INTO
    topups (
        card_number,
        topup_amount,
        topup_method,
        topup_time,
        created_at,
        updated_at
    )
VALUES (
        $1,
        $2,
        $3,
        $4,
        current_timestamp,
        current_timestamp
    )
RETURNING
    topup_id,
    topup_no,
    card_number,
    topup_amount,
    topup_method,
    topup_time,
    status,
    created_at,
    updated_at
`

const seedUpdateTopupStatus = `UPDATE topups
SET
    status = $2,
    updated_at = current_timestamp
WHERE
    topup_id = $1
    AND deleted_at IS NULL
    AND (
        ($2 = 'failed' AND status IN ('pending', 'failed', 'compensation_required'))
        OR ($2 = 'compensation_required' AND status IN ('pending', 'failed', 'success', 'compensation_required'))
        OR ($2 = 'success' AND status IN ('pending', 'failed', 'success', 'compensation_required'))
        OR ($2 = 'pending' AND status IN ('failed', 'compensation_required'))
    )
RETURNING
    topup_id,
    topup_no,
    card_number,
    topup_amount,
    topup_method,
    topup_time,
    status,
    created_at,
    updated_at
`

const seedTrashTopup = `UPDATE topups
SET
    deleted_at = current_timestamp
WHERE
    topup_id = $1
    AND deleted_at IS NULL
RETURNING
    topup_id,
    topup_no,
    card_number,
    topup_amount,
    topup_method,
    topup_time,
    status,
    created_at,
    updated_at,
    deleted_at
`

// topupSeeder is a struct that represents a seeder for the topups table.
type topupSeeder struct {
	db     *sqlx.DB
	ctx    context.Context
	logger logger.LoggerInterface
}

// NewTopupSeeder creates a new instance of the topupSeeder, which is
// responsible for populating the topups table with fake data.
//
// Args:
// db: a pointer to the database pool
// ctx: a context.Context object
// logger: a logger.LoggerInterface object
//
// Returns:
// a pointer to the topupSeeder struct
func NewTopupSeeder(db *sqlx.DB, ctx context.Context, logger logger.LoggerInterface) *topupSeeder {
	return &topupSeeder{
		db:     db,
		ctx:    ctx,
		logger: logger,
	}
}

// Seed populates the topups table with fake data.
//
// It creates a total of 10 topups, with 5 of them being active and 5 of them being
// trashed. The topups are seeded with a random card number, topup amount, topup method,
// and topup time. The status of the topup is randomly set to one of the following:
// pending, success, or failed.
//
// If any errors occur during the seeding process, an error is returned.
//
// Returns:
// an error if any of the topups fail to be created, otherwise nil
func (r *topupSeeder) Seed() error {
	totalTopups := 10
	activeTopups := 5
	trashedTopups := 5

	var cards []db.GetCardByUserIDRow

	for i := 1; i <= totalTopups; i++ {
		var card db.GetCardByUserIDRow
		err := r.db.GetContext(r.ctx, &card, seedGetCardByUserID, int32(i))
		if err != nil {
			r.logger.Error("failed to get card for user", zap.Int("userID", i), zap.Error(err))
			return fmt.Errorf("failed to get card for user %d: %w", i, err)
		}
		cards = append(cards, card)
	}

	if len(cards) < totalTopups {
		r.logger.Error("not enough cards found to seed topups", zap.Int("found", len(cards)))
		return fmt.Errorf("not enough cards to seed topups")
	}

	topupMethods := []string{"Bank Alpha", "Bank Beta", "Bank Gamma"}
	statusOptions := []string{"pending", "success", "failed"}

	months := make([]time.Time, 12)
	currentYear := time.Now().Year()
	for i := 0; i < 12; i++ {
		months[i] = time.Date(currentYear, time.Month(i+1), 1, 0, 0, 0, 0, time.UTC)
	}

	for i := 0; i < totalTopups; i++ {
		card := cards[i]
		cardNumber := card.CardNumber

		monthIndex := i % 12
		topupTime := months[monthIndex].Add(time.Duration(rand.Intn(28)) * 24 * time.Hour)

		request := db.CreateTopupParams{
			CardNumber:  cardNumber,
			TopupAmount: int32(rand.Intn(10000000) + 1000000),
			TopupMethod: topupMethods[rand.Intn(len(topupMethods))],
			TopupTime:   topupTime,
		}

		var topup db.CreateTopupRow
		err := r.db.GetContext(r.ctx, &topup, seedCreateTopup,
			request.CardNumber,
			request.TopupAmount,
			request.TopupMethod,
			request.TopupTime,
		)
		if err != nil {
			r.logger.Error("failed to seed topup", zap.String("card", cardNumber), zap.Error(err))
			return fmt.Errorf("failed to seed topup for card %s: %w", cardNumber, err)
		}

		status := statusOptions[rand.Intn(len(statusOptions))]
		var updated db.UpdateTopupStatusRow
		err = r.db.GetContext(r.ctx, &updated, seedUpdateTopupStatus, topup.TopupID, status)
		if err != nil {
			r.logger.Error("failed to update topup status", zap.Int("topupID", int(topup.TopupID)), zap.Error(err))
			return fmt.Errorf("failed to update status for topup %d: %w", topup.TopupID, err)
		}

		if i >= activeTopups {
			var trashed db.TrashTopupRow
			err = r.db.GetContext(r.ctx, &trashed, seedTrashTopup, topup.TopupID)
			if err != nil {
				r.logger.Error("failed to trash topup", zap.Int("topup", i+1), zap.String("card", cardNumber), zap.Error(err))
				return fmt.Errorf("failed to trash topup %d for card %s: %w", i+1, cardNumber, err)
			}
		}
	}

	r.logger.Info("topup seeded successfully", zap.Int("totalTopups", totalTopups), zap.Int("activeTopups", activeTopups), zap.Int("trashedTopups", trashedTopups))
	return nil
}
