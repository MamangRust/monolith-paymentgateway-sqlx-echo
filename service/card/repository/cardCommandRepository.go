package repository

import (
	"context"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-pkg/randomvcc"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	sharedErrors "github.com/MamangRust/monolith-payment-gateway-shared/errors"
	card_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/card_errors/repository"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jmoiron/sqlx"
)

const createCard = `WITH inserted_card AS (
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
FROM inserted_card`

const updateCard = `UPDATE cards
SET
    card_type = $2,
    expire_date = $3,
    cvv = $4,
    card_provider = $5,
    updated_at = current_timestamp
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
    created_at,
    updated_at`

const trashCard = `UPDATE cards
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
    deleted_at`

const restoreCard = `UPDATE cards
SET
    deleted_at = NULL
WHERE
    card_id = $1
    AND deleted_at IS NOT NULL
RETURNING
    card_id,
    user_id,
    card_number,
    card_type,
    expire_date,
    cvv,
    card_provider,
    created_at,
    updated_at,
    deleted_at`

const deleteCardPermanently = `DELETE FROM cards WHERE card_id = $1 AND deleted_at IS NOT NULL`

const restoreAllCards = `UPDATE cards
SET
    deleted_at = NULL
WHERE
    deleted_at IS NOT NULL`

const deleteAllPermanentCards = `DELETE FROM cards WHERE deleted_at IS NOT NULL`

const updateCardStatus = `UPDATE cards
SET
    status = $2,
    updated_at = CURRENT_TIMESTAMP
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
    updated_at`

const updateCreditLimit = `WITH updated_card AS (
    UPDATE cards
    SET credit_limit = $2, event_version = event_version + 1, updated_at = CURRENT_TIMESTAMP
    WHERE card_id = $1 AND deleted_at IS NULL
    RETURNING card_id, user_id, card_number, card_type, expire_date, cvv, card_provider, status, credit_limit, outstanding_balance, reward_points, created_at, updated_at, event_version
), queued_event AS (
    INSERT INTO outbox_events (event_key, topic, message_key, payload)
    SELECT
        'card-limit-changed:' || card_id::text || ':' || event_version::text,
        'card.limit.changed',
        card_number,
        jsonb_build_object('card_id', card_id, 'card_number', card_number, 'credit_limit', credit_limit, 'event_version', event_version)
    FROM updated_card
    ON CONFLICT (event_key) DO NOTHING
    RETURNING event_key
)
SELECT card_id, user_id, card_number, card_type, expire_date, cvv, card_provider, status, credit_limit, outstanding_balance, reward_points, created_at, updated_at
FROM updated_card`

const redeemRewardPoints = `UPDATE cards
SET
    reward_points = GREATEST(0, reward_points - $2),
    updated_at = CURRENT_TIMESTAMP
WHERE
    card_id = $1
    AND deleted_at IS NULL
    AND reward_points >= $2
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
    updated_at`

const processBillingCycles = `WITH inserted_cycles AS (
    INSERT INTO billing_cycles (
        card_number,
        cycle_start,
        cycle_end,
        amount_due,
        due_date,
        status
    )
    SELECT
        c.card_number,
        date_trunc('month', CURRENT_DATE),
        date_trunc('month', CURRENT_DATE) + interval '1 month' - interval '1 day',
        c.outstanding_balance,
        CURRENT_DATE,
        'unpaid'
    FROM cards c
    WHERE c.card_type = 'credit'
      AND c.status = 'active'
      AND c.deleted_at IS NULL
      AND c.outstanding_balance > 0
      AND $1 BETWEEN 1 AND 31
      AND EXTRACT(DAY FROM CURRENT_DATE) = $1
    ON CONFLICT (card_number, cycle_start, cycle_end) DO NOTHING
    RETURNING billing_id, card_number, cycle_start, cycle_end, amount_due, due_date, status, created_at, updated_at
), queued_events AS (
    INSERT INTO outbox_events (event_key, topic, message_key, payload)
    SELECT
        'card-statement:' || billing_id::text,
        'card.statement.generated',
        billing_id::text,
        jsonb_build_object(
            'billing_id', billing_id,
            'card_number', card_number,
            'billing_cycle_day', $1,
            'cycle_start', cycle_start,
            'cycle_end', cycle_end,
            'amount_due', amount_due
        )
    FROM inserted_cycles
    ON CONFLICT (event_key) DO NOTHING
    RETURNING event_key
)
SELECT
    billing_id,
    card_number,
    cycle_start,
    cycle_end,
    amount_due,
    due_date,
    status,
    created_at,
    updated_at
FROM inserted_cycles`

type cardCommandRepository struct {
	db *sqlx.DB
}

func NewCardCommandRepository(db *sqlx.DB) CardCommandRepository {
	return &cardCommandRepository{
		db: db,
	}
}

func (r *cardCommandRepository) CreateCard(ctx context.Context, request *requests.CreateCardRequest) (*db.CreateCardRow, error) {
	number, err := randomvcc.RandomCardNumber()
	if err != nil {
		return nil, sharedErrors.ErrInternal.WithInternal(err)
	}

	expireDate := pgtype.Date{
		Time:  request.ExpireDate,
		Valid: true,
	}

	var res db.CreateCardRow
	err = r.db.GetContext(ctx, &res, createCard,
		int32(request.UserID),
		number,
		request.CardType,
		expireDate,
		request.CVV,
		request.CardProvider,
	)
	if err != nil {
		return nil, sharedErrors.ErrConstraintOrFailed(err, "Card", "create card")
	}

	return &res, nil
}

func (r *cardCommandRepository) UpdateCard(ctx context.Context, request *requests.UpdateCardRequest) (*db.UpdateCardRow, error) {
	expireDate := pgtype.Date{
		Time:  request.ExpireDate,
		Valid: true,
	}

	var res db.UpdateCardRow
	err := r.db.GetContext(ctx, &res, updateCard,
		int32(request.CardID),
		request.CardType,
		expireDate,
		request.CVV,
		request.CardProvider,
	)
	if err != nil {
		return nil, sharedErrors.ErrNoRowsOrFailed(err, "Card", "update card")
	}

	return &res, nil
}

func (r *cardCommandRepository) TrashedCard(ctx context.Context, card_id int) (*db.TrashCardRow, error) {
	var res db.TrashCardRow
	err := r.db.GetContext(ctx, &res, trashCard, int32(card_id))

	if err != nil {
		return nil, sharedErrors.ErrNoRowsOrFailed(err, "Card", "trash card")
	}

	return &res, nil
}

func (r *cardCommandRepository) RestoreCard(ctx context.Context, card_id int) (*db.RestoreCardRow, error) {
	var res db.RestoreCardRow
	err := r.db.GetContext(ctx, &res, restoreCard, int32(card_id))

	if err != nil {
		return nil, sharedErrors.ErrNoRowsOrFailed(err, "Card", "restore card")
	}

	return &res, nil
}

func (r *cardCommandRepository) DeleteCardPermanent(ctx context.Context, card_id int) (bool, error) {
	if _, err := r.db.ExecContext(ctx, deleteCardPermanently, int32(card_id)); err != nil {
		return false, sharedErrors.ErrNoRowsOrFailed(err, "Card", "delete card permanently")
	}

	return true, nil
}

func (r *cardCommandRepository) RestoreAllCard(ctx context.Context) (bool, error) {
	if _, err := r.db.ExecContext(ctx, restoreAllCards); err != nil {
		return false, card_errors.ErrRestoreAllCardsFailed.WithInternal(err)
	}

	return true, nil
}

func (r *cardCommandRepository) DeleteAllCardPermanent(ctx context.Context) (bool, error) {
	if _, err := r.db.ExecContext(ctx, deleteAllPermanentCards); err != nil {
		return false, card_errors.ErrDeleteAllCardsPermanentFailed.WithInternal(err)
	}

	return true, nil
}

func (r *cardCommandRepository) ToggleCardStatus(ctx context.Context, request *requests.ToggleCardStatusRequest) (*db.UpdateCardStatusRow, error) {
	// Determine new status: if currently active, suspend; if suspended, activate
	var card db.GetCardByIDRow
	err := r.db.GetContext(ctx, &card, getCardByID, int32(request.CardID))
	if err != nil {
		return nil, sharedErrors.ErrNoRowsOrFailed(err, "Card", "toggle card status")
	}

	newStatus := "active"
	if card.Status == "active" {
		newStatus = "suspended"
	}

	var res db.UpdateCardStatusRow
	if err := r.db.GetContext(ctx, &res, updateCardStatus, int32(request.CardID), newStatus); err != nil {
		return nil, sharedErrors.ErrNoRowsOrFailed(err, "Card", "toggle card status")
	}

	return &res, nil
}

func (r *cardCommandRepository) UpdateCreditLimit(ctx context.Context, request *requests.UpdateCreditLimitRequest) (*db.UpdateCreditLimitRow, error) {
	var res db.UpdateCreditLimitRow
	if err := r.db.GetContext(ctx, &res, updateCreditLimit, int32(request.CardID), int32(request.CreditLimit)); err != nil {
		return nil, sharedErrors.ErrNoRowsOrFailed(err, "Card", "update credit limit")
	}

	return &res, nil
}

func (r *cardCommandRepository) RedeemPoints(ctx context.Context, request *requests.RedeemPointsRequest) (*db.RedeemRewardPointsRow, error) {
	var res db.RedeemRewardPointsRow
	if err := r.db.GetContext(ctx, &res, redeemRewardPoints, int32(request.CardID), int32(request.Points)); err != nil {
		return nil, sharedErrors.ErrNoRowsOrFailed(err, "Card", "redeem reward points")
	}

	return &res, nil
}

func (r *cardCommandRepository) ProcessBillingCycles(ctx context.Context, billingCycleDay int) ([]*db.BillingCycle, error) {
	var cycles []*db.ProcessBillingCyclesRow
	if err := r.db.SelectContext(ctx, &cycles, processBillingCycles, int32(billingCycleDay)); err != nil {
		return nil, card_errors.ErrInsertBillingCycleFailed.WithInternal(err)
	}

	result := make([]*db.BillingCycle, 0, len(cycles))
	for _, cycle := range cycles {
		result = append(result, &db.BillingCycle{
			BillingID:  cycle.BillingID,
			CardNumber: cycle.CardNumber,
			CycleStart: cycle.CycleStart,
			CycleEnd:   cycle.CycleEnd,
			AmountDue:  cycle.AmountDue,
			DueDate:    cycle.DueDate,
			Status:     cycle.Status,
			CreatedAt:  cycle.CreatedAt,
			UpdatedAt:  cycle.UpdatedAt,
		})
	}
	return result, nil
}
