package seeder

import (
	"context"
	"fmt"
	"math/rand"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

const seedCreateSaldo = `INSERT INTO
    saldos (
        card_number,
        total_balance,
        created_at,
        updated_at
    )
VALUES (
        $1,
        $2,
        current_timestamp,
        current_timestamp
    )
ON CONFLICT (card_number) WHERE deleted_at IS NULL
DO UPDATE SET
    total_balance = EXCLUDED.total_balance,
    updated_at = current_timestamp
RETURNING
    saldo_id,
    card_number,
    total_balance,
    withdraw_amount,
    withdraw_time,
    created_at,
    updated_at
`

const seedTrashSaldo = `UPDATE saldos
SET
    deleted_at = current_timestamp
WHERE
    saldo_id = $1
    AND deleted_at IS NULL
RETURNING
    saldo_id,
    card_number,
    total_balance,
    withdraw_amount,
    withdraw_time,
    created_at,
    updated_at,
    deleted_at
`

// saldoSeeder is a struct that represents a seeder for the saldos table.
type saldoSeeder struct {
	db     *sqlx.DB
	ctx    context.Context
	logger logger.LoggerInterface
}

// NewSaldoSeeder creates a new instance of the saldoSeeder, which is
// responsible for populating the saldos table with fake data.
//
// Args:
// db: a pointer to the database pool
// ctx: a context.Context object
// logger: a logger.LoggerInterface object
//
// Returns:
// a pointer to the saldoSeeder struct
func NewSaldoSeeder(db *sqlx.DB, ctx context.Context, logger logger.LoggerInterface) *saldoSeeder {
	return &saldoSeeder{
		db:     db,
		ctx:    ctx,
		logger: logger,
	}
}

// Seed populates the saldos table with fake data.
//
// It generates 10 saldos, 5 of which are active and 5 of which are trashed.
// The saldos are seeded with random totalBalance values, and the cardNumber
// is randomly selected from the cards table.
//
// If any errors occur during the seeding process, an error is returned.
//
// Returns:
// an error if any of the saldos fail to be created, otherwise nil
func (r *saldoSeeder) Seed() error {
	totalSaldos := 10
	activeSaldos := 5
	trashedSaldos := 5

	var cards []db.GetCardByUserIDRow
	for i := 1; i <= totalSaldos; i++ {
		var card db.GetCardByUserIDRow
		err := r.db.GetContext(r.ctx, &card, seedGetCardByUserID, int32(i))
		if err != nil {
			r.logger.Error("failed to get card for user", zap.Int("userID", i), zap.Error(err))
			return fmt.Errorf("failed to get card for user %d: %w", i, err)
		}
		cards = append(cards, card)
	}

	if len(cards) < totalSaldos {
		r.logger.Error("not enough cards to seed saldo", zap.Int("required", totalSaldos), zap.Int("available", len(cards)))
		return fmt.Errorf("not enough cards to seed saldo: required %d, got %d", totalSaldos, len(cards))
	}

	for i, card := range cards {
		request := db.CreateSaldoParams{
			CardNumber:   card.CardNumber,
			TotalBalance: int32(rand.Intn(9_000_000) + 1_000_000),
		}

		var saldo db.CreateSaldoRow
		err := r.db.GetContext(r.ctx, &saldo, seedCreateSaldo,
			request.CardNumber,
			request.TotalBalance,
		)
		if err != nil {
			r.logger.Error("failed to seed saldo", zap.Int("index", i), zap.String("card", card.CardNumber), zap.Error(err))
			return fmt.Errorf("failed to seed saldo for card %s: %w", card.CardNumber, err)
		}

		if i >= activeSaldos {
			var trashed db.Saldo
			err = r.db.GetContext(r.ctx, &trashed, seedTrashSaldo, saldo.SaldoID)
			if err != nil {
				r.logger.Error("failed to trash saldo", zap.Int("index", i), zap.String("card", card.CardNumber), zap.Error(err))
				return fmt.Errorf("failed to trash saldo %d for card %s: %w", i+1, card.CardNumber, err)
			}
		}
	}

	r.logger.Info("saldo seeded successfully",
		zap.Int("totalSaldos", totalSaldos),
		zap.Int("activeSaldos", activeSaldos),
		zap.Int("trashedSaldos", trashedSaldos))

	return nil
}
