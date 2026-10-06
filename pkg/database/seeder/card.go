package seeder

import (
	"context"
	"fmt"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-pkg/date"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-pkg/randomvcc"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

const seedCreateCard = `WITH inserted_card AS (
    INSERT INTO cards (
        user_id, card_number, card_type, expire_date, cvv, card_provider, created_at, updated_at
    )
    VALUES ($1, $2, $3, $4, $5, $6, current_timestamp, current_timestamp)
    RETURNING card_id, user_id, card_number, card_type, expire_date, cvv, card_provider, created_at, updated_at
), queued_event AS (
    INSERT INTO outbox_events (event_key, topic, message_key, payload)
    SELECT
        'card-create-saldo:' || card_id::text,
        'saldo-service-topic-create-saldo',
        card_id::text,
        jsonb_build_object('card_number', card_number, 'total_balance', 0)
    FROM inserted_card
    ON CONFLICT (event_key) DO NOTHING
    RETURNING event_key
)
SELECT card_id, user_id, card_number, card_type, expire_date, cvv, card_provider, created_at, updated_at
FROM inserted_card
`

const seedTrashCard = `UPDATE cards
SET
    deleted_at = current_timestamp
WHERE
    card_id = $1
    AND deleted_at IS NULL
RETURNING
    card_id,
    user_id,
    card_number,
    card_type,
    expire_date,
    cvv,
    card_provider,
    status,
    credit_limit,
    outstanding_balance,
    reward_points,
    created_at,
    updated_at,
    deleted_at
`

const seedGetCardByUserID = `SELECT
    card_id,
    user_id,
    card_number,
    card_type,
    expire_date,
    cvv,
    card_provider,
    status,
    credit_limit,
    outstanding_balance,
    reward_points,
    created_at,
    updated_at
FROM cards
WHERE
    user_id = $1
    AND deleted_at IS NULL
LIMIT 1
`

// cardSeeder is a struct that represents a seeder for the cards table.
type cardSeeder struct {
	db     *sqlx.DB
	ctx    context.Context
	logger logger.LoggerInterface
}

// NewCardSeeder creates a new instance of the cardSeeder, which is
// responsible for populating the cards table with fake data.
func NewCardSeeder(db *sqlx.DB, ctx context.Context, logger logger.LoggerInterface) *cardSeeder {
	return &cardSeeder{
		db:     db,
		ctx:    ctx,
		logger: logger,
	}
}

// Seed populates the cards table with fake data.
//
// It generates 10 cards, 5 of which are active and 5 of which are trashed.
// The cards are seeded with random card numbers, card types, expire dates, and
// CVV numbers. The card provider is randomly selected from a list of providers.
//
// If any errors occur during the seeding process, an error is returned.
func (r *cardSeeder) Seed() error {
	cardTypes := []string{"credit", "debit"}
	cardProviders := []string{"mandiri", "bni", "bri"}

	totalCards := 10
	activeCards := 5
	trashedCards := 5

	cardNumbers := make([]string, totalCards)
	for i := 0; i < totalCards; i++ {
		cardNumber, err := randomvcc.RandomCardNumber()
		if err != nil {
			r.logger.Error("failed to generate card number", zap.Int("index", i), zap.Error(err))
			return fmt.Errorf("failed to generate card number: %w", err)
		}
		cardNumbers[i] = cardNumber
	}

	for i := 0; i < totalCards; i++ {
		expireDate := pgtype.Date{
			Time:  date.GenerateExpireDate(),
			Valid: true,
		}

		var card db.CreateCardRow
		err := r.db.GetContext(r.ctx, &card, seedCreateCard,
			int32(i+1),
			cardNumbers[i],
			cardTypes[i%len(cardTypes)],
			expireDate,
			fmt.Sprintf("%03d", i%1000),
			cardProviders[i%len(cardProviders)],
		)
		if err != nil {
			r.logger.Error("failed to seed card", zap.Int("card", i+1), zap.Error(err))
			return fmt.Errorf("failed to seed card %d: %w", i+1, err)
		}

		if i >= activeCards {
			var trashed db.TrashCardRow
			err = r.db.GetContext(r.ctx, &trashed, seedTrashCard, card.CardID)
			if err != nil {
				r.logger.Error("failed to trash card", zap.Int("card", i+1), zap.Error(err))
				return fmt.Errorf("failed to trash card %d: %w", i+1, err)
			}
		}
	}

	r.logger.Info("card seeded successfully", zap.Int("totalCards", totalCards), zap.Int("activeCards", activeCards), zap.Int("trashedCards", trashedCards))
	return nil
}
