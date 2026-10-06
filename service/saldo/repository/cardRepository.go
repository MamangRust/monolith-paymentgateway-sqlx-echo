package repository

import (
	"context"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	card_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/card_errors/repository"
	"github.com/jmoiron/sqlx"
)

type cardRepository struct {
	db *sqlx.DB
}

func NewCardRepository(db *sqlx.DB) CardRepository {
	return &cardRepository{
		db: db,
	}
}

const getCardByCardNumber = `SELECT
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
    card_number = $1
    AND deleted_at IS NULL`

func (r *cardRepository) FindCardByCardNumber(ctx context.Context, card_number string) (*db.GetCardByCardNumberRow, error) {
	var card db.GetCardByCardNumberRow
	err := r.db.GetContext(ctx, &card, getCardByCardNumber, card_number)

	if err != nil {
		return nil, card_errors.ErrFindCardByCardNumberFailed
	}

	return &card, nil
}
