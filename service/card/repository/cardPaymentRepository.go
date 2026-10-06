package repository

import (
	"context"
	"database/sql"
	"errors"
	"math"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	sharedErrors "github.com/MamangRust/monolith-payment-gateway-shared/errors"
	card_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/card_errors/repository"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

const getCardPaymentByReferenceID = `SELECT payment_id, payment_uuid, card_number, billing_id, amount,
       payment_channel, reference_id, status, created_at, updated_at
FROM card_payments
WHERE reference_id = $1
  AND reference_id <> ''
LIMIT 1`

const getCardPaymentsByCardNumber = `SELECT payment_id, payment_uuid, card_number, billing_id, amount, payment_channel, reference_id, status, created_at, updated_at,
    COUNT(*) OVER() AS total_count
FROM card_payments
WHERE card_number = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3`

const countCardPayments = `SELECT COUNT(*)::INT AS total_count
FROM card_payments
WHERE card_number = $1`

const getBillingCycleByIDForUpdate = `SELECT
    billing_id,
    card_number,
    cycle_start,
    cycle_end,
    amount_due,
    due_date,
    status,
    created_at,
    updated_at
FROM billing_cycles
WHERE billing_id = $1
FOR UPDATE`

const updateBillingCycleStatus = `UPDATE billing_cycles
SET
    status = $2,
    updated_at = CURRENT_TIMESTAMP
WHERE
    billing_id = $1
    AND status = $3
RETURNING
    billing_id,
    card_number,
    cycle_start,
    cycle_end,
    amount_due,
    due_date,
    status,
    created_at,
    updated_at`

const insertCardPayment = `WITH inserted_payment AS (
    INSERT INTO card_payments (
        payment_uuid, card_number, billing_id, amount, payment_channel, reference_id, status
    ) VALUES ($1, $2, $3, $4, $5, $6, 'completed')
    RETURNING payment_id, payment_uuid, card_number, billing_id, amount, payment_channel, reference_id, status, created_at, updated_at
), queued_event AS (
    INSERT INTO outbox_events (event_key, topic, message_key, payload)
    SELECT
        'card-payment-posted:' || payment_id::text,
        'card.payment.posted',
        payment_id::text,
        jsonb_build_object('card_number', card_number, 'payment_id', payment_id, 'amount', amount, 'channel', payment_channel)
    FROM inserted_payment
    ON CONFLICT (event_key) DO NOTHING
    RETURNING event_key
)
SELECT payment_id, payment_uuid, card_number, billing_id, amount, payment_channel, reference_id, status, created_at, updated_at
FROM inserted_payment`

type cardPaymentRepository struct {
	db *sqlx.DB
}

func NewCardPaymentRepository(db *sqlx.DB) CardPaymentRepository {
	return &cardPaymentRepository{
		db: db,
	}
}

func (r *cardPaymentRepository) PostPaymentIdempotent(ctx context.Context, req *requests.PostPaymentRequest, expectedBillingStatus, targetBillingStatus string) (*CardPaymentResult, error) {
	if req.ReferenceID != "" {
		existing, err := r.getPaymentByReferenceID(ctx, r.db, req.ReferenceID)
		if err == nil {
			if validationErr := validatePaymentReplay(existing, req); validationErr != nil {
				return nil, validationErr
			}
			return &CardPaymentResult{Payment: existing, Replayed: true}, nil
		}
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, card_errors.ErrGetCardPaymentFailed.WithInternal(err)
		}
	}

	var billingID *int32
	if req.BillingID != nil {
		billingID = new(int32)
		*billingID = int32(*req.BillingID)
	}

	// Keep billing state change and payment insert in one explicit transaction.
	// The billing row lock makes the expected-state check safe against two
	// concurrent payments for the same cycle.
	tx, err := r.db.Beginx()
	if err != nil {
		return nil, card_errors.ErrInsertCardPaymentFailed.WithInternal(err)
	}
	defer func() { _ = tx.Rollback() }()

	if billingID != nil {
		// Lock the billing row before reading/checking its state.
		var cycle db.BillingCycle
		lockErr := tx.GetContext(ctx, &cycle, getBillingCycleByIDForUpdate, *billingID)
		if lockErr != nil {
			return nil, card_errors.ErrGetBillingCyclesFailed.WithInternal(lockErr)
		}
		if cycle.Status != expectedBillingStatus {
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				return nil, card_errors.ErrGetBillingCyclesFailed.WithInternal(rollbackErr)
			}
			if req.ReferenceID != "" {
				if existing, lookupErr := r.getPaymentByReferenceID(ctx, r.db, req.ReferenceID); lookupErr == nil {
					if validationErr := validatePaymentReplay(existing, req); validationErr != nil {
						return nil, validationErr
					}
					return &CardPaymentResult{Payment: existing, Replayed: true}, nil
				}
			}
			return nil, sharedErrors.NewBadRequestError("billing cycle is not in a payable state")
		}
		if req.Amount != int64(cycle.AmountDue) {
			return nil, sharedErrors.NewBadRequestError("payment amount must equal billing amount due")
		}
		var updated db.BillingCycle
		if updateErr := tx.GetContext(ctx, &updated, updateBillingCycleStatus, *billingID, targetBillingStatus, expectedBillingStatus); updateErr != nil {
			return nil, sharedErrors.NewBadRequestError("billing cycle is not in a payable state")
		}
	}

	var res db.InsertCardPaymentRow
	insertErr := tx.GetContext(ctx, &res, insertCardPayment,
		uuid.New().String(),
		req.CardNumber,
		billingID,
		req.Amount,
		req.PaymentChannel,
		req.ReferenceID,
	)
	if insertErr != nil {
		if isUniqueViolation(insertErr) && req.ReferenceID != "" {
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				return nil, card_errors.ErrInsertCardPaymentFailed.WithInternal(rollbackErr)
			}
			existing, lookupErr := r.getPaymentByReferenceID(ctx, r.db, req.ReferenceID)
			if lookupErr == nil {
				if validationErr := validatePaymentReplay(existing, req); validationErr != nil {
					return nil, validationErr
				}
				return &CardPaymentResult{Payment: existing, Replayed: true}, nil
			}
		}
		return nil, card_errors.ErrInsertCardPaymentFailed.WithInternal(insertErr)
	}
	if err := tx.Commit(); err != nil {
		return nil, card_errors.ErrInsertCardPaymentFailed.WithInternal(err)
	}
	return &CardPaymentResult{Payment: &db.CardPayment{
		PaymentID:      res.PaymentID,
		PaymentUuid:    res.PaymentUuid,
		CardNumber:     res.CardNumber,
		BillingID:      res.BillingID,
		Amount:         res.Amount,
		PaymentChannel: res.PaymentChannel,
		ReferenceID:    res.ReferenceID,
		Status:         res.Status,
		CreatedAt:      res.CreatedAt,
		UpdatedAt:      res.UpdatedAt,
	}}, nil
}

func isUniqueViolation(err error) bool {
	var stateErr interface{ SQLState() string }
	return errors.As(err, &stateErr) && stateErr.SQLState() == "23505"
}

func validatePaymentReplay(existing *db.CardPayment, req *requests.PostPaymentRequest) error {
	if existing == nil || existing.CardNumber != req.CardNumber || existing.Amount != req.Amount || existing.PaymentChannel != req.PaymentChannel {
		return sharedErrors.ErrConflict.WithMessage("card payment reference already used with different payload")
	}
	if req.BillingID == nil {
		if existing.BillingID != nil {
			return sharedErrors.ErrConflict.WithMessage("card payment reference already used with different billing cycle")
		}
		return nil
	}
	if existing.BillingID == nil || *existing.BillingID != int32(*req.BillingID) {
		return sharedErrors.ErrConflict.WithMessage("card payment reference already used with different billing cycle")
	}
	return nil
}

func (r *cardPaymentRepository) getPaymentByReferenceID(ctx context.Context, q sqlx.QueryerContext, referenceID string) (*db.CardPayment, error) {
	var res db.CardPayment
	if err := sqlx.GetContext(ctx, q, &res, getCardPaymentByReferenceID, referenceID); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r *cardPaymentRepository) GetPaymentByReferenceID(ctx context.Context, referenceID string) (*db.CardPayment, error) {
	res, err := r.getPaymentByReferenceID(ctx, r.db, referenceID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, card_errors.ErrGetCardPaymentFailed.WithInternal(err)
	}
	return res, nil
}

func (r *cardPaymentRepository) PostPayment(ctx context.Context, req *requests.PostPaymentRequest) (*db.CardPayment, error) {
	expectedBillingStatus, targetBillingStatus := "", ""
	if req.BillingID != nil {
		expectedBillingStatus = "unpaid"
		targetBillingStatus = "paid"
	}
	result, err := r.PostPaymentIdempotent(ctx, req, expectedBillingStatus, targetBillingStatus)
	if err != nil {
		return nil, err
	}
	return result.Payment, nil
}

func (r *cardPaymentRepository) GetPaymentHistory(ctx context.Context, cardNumber string, page, pageSize int) ([]*db.GetCardPaymentsByCardNumberRow, error) {
	offset := (page - 1) * pageSize

	var res []*db.GetCardPaymentsByCardNumberRow
	err := r.db.SelectContext(ctx, &res, getCardPaymentsByCardNumber,
		cardNumber,
		int32(pageSize),
		int32(math.Max(0, float64(offset))),
	)
	if err != nil {
		return nil, card_errors.ErrGetCardPaymentFailed.WithInternal(err)
	}

	return res, nil
}

func (r *cardPaymentRepository) CountPayments(ctx context.Context, cardNumber string) (int, error) {
	var count int32
	if err := r.db.GetContext(ctx, &count, countCardPayments, cardNumber); err != nil {
		return 0, card_errors.ErrGetCardPaymentFailed.WithInternal(err)
	}

	return int(count), nil
}
