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

const seedGetMerchants = `SELECT
    merchant_id,
    merchant_no,
    name,
    api_key,
    user_id,
    status,
    created_at,
    updated_at,
    COUNT(*) OVER () AS total_count
FROM merchants
WHERE
    deleted_at IS NULL
    AND (
        $1::TEXT IS NULL
        OR name ILIKE '%' || $1 || '%'
        OR api_key ILIKE '%' || $1 || '%'
        OR status ILIKE '%' || $1 || '%'
    )
ORDER BY merchant_id
LIMIT $2
OFFSET
    $3
`

const seedCreateTransaction = `INSERT INTO
    transactions (
        card_number,
        amount,
        payment_method,
        merchant_id,
        transaction_time,
        created_at,
        updated_at
    )
VALUES (
        $1,
        $2,
        $3,
        $4,
        $5,
        current_timestamp,
        current_timestamp
    )
RETURNING
    transaction_id,
    transaction_no,
    card_number,
    amount,
    payment_method,
    merchant_id,
    transaction_time,
    status,
    created_at,
    updated_at
`

const seedUpdateTransactionStatus = `UPDATE transactions
SET
    status = $2,
    updated_at = current_timestamp
WHERE
    transaction_id = $1
    AND deleted_at IS NULL
    AND (
        ($2 = 'failed' AND status IN ('pending', 'failed', 'compensation_required'))
        OR ($2 = 'compensation_required' AND status IN ('pending', 'failed', 'success', 'authorized', 'captured', 'compensation_required'))
        OR ($2 = 'success' AND status IN ('pending', 'failed', 'success', 'compensation_required'))
        OR ($2 = 'pending' AND status IN ('failed', 'compensation_required'))
        OR ($2 = 'flagged_fraud' AND status IN ('success', 'authorized', 'captured'))
    )
RETURNING
    transaction_id,
    transaction_no,
    card_number,
    amount,
    payment_method,
    merchant_id,
    transaction_time,
    status,
    created_at,
    updated_at
`

const seedTrashTransaction = `UPDATE transactions
SET
    deleted_at = current_timestamp
WHERE
    transaction_id = $1
    AND deleted_at IS NULL
RETURNING
    transaction_id,
    transaction_no,
    card_number,
    amount,
    payment_method,
    merchant_id,
    transaction_time,
    status,
    created_at,
    updated_at,
    deleted_at
`

// transactionSeeder is a struct that represents a seeder for the transactions table.
type transactionSeeder struct {
	db     *sqlx.DB
	ctx    context.Context
	logger logger.LoggerInterface
}

// NewTransactionSeeder creates a new instance of the transactionSeeder, which is
// responsible for populating the transactions table with fake data.
//
// Args:
// db: a pointer to the database pool
// ctx: a context.Context object
// logger: a logger.LoggerInterface object
//
// Returns:
// a pointer to the transactionSeeder struct
func NewTransactionSeeder(db *sqlx.DB, ctx context.Context, logger logger.LoggerInterface) *transactionSeeder {
	return &transactionSeeder{
		db:     db,
		ctx:    ctx,
		logger: logger,
	}
}

// Seed populates the transactions table with fake data.
//
// It generates a total of 10 transactions, with 5 of them being active and 5 of them being
// trashed. The transactions are seeded with random payment methods, statuses, and transaction
// times. The card numbers are randomly selected from the cards table. The merchant IDs are
// randomly selected from the merchants table.
//
// If any of the transactions fail to be created, the function returns an error.
//
// Returns:
// an error if any of the transactions fail to be created, otherwise nil
func (r *transactionSeeder) Seed() error {
	total := 10
	active := 5
	trashed := 5

	paymentMethods := []string{"Bank Alpha", "Bank Beta", "Bank Gamma"}
	statusOptions := []string{"pending", "success", "failed"}

	var cards []db.GetCardByUserIDRow
	for i := 1; i <= total; i++ {
		var card db.GetCardByUserIDRow
		err := r.db.GetContext(r.ctx, &card, seedGetCardByUserID, int32(i))
		if err != nil {
			r.logger.Error("failed to get card for user", zap.Int("userID", i), zap.Error(err))
			return fmt.Errorf("failed to get card for user %d: %w", i, err)
		}
		cards = append(cards, card)
	}

	if len(cards) < total {
		r.logger.Error("not enough cards for transaction seeding", zap.Int("required", total), zap.Int("available", len(cards)))
		return fmt.Errorf("not enough cards for transaction seeding: required %d, got %d", total, len(cards))
	}

	var merchants []*db.GetMerchantsRow
	err := r.db.SelectContext(r.ctx, &merchants, seedGetMerchants,
		"",
		int32(total),
		0,
	)
	if err != nil {
		r.logger.Error("failed to get merchant list", zap.Error(err))
		return fmt.Errorf("failed to get merchant list: %w", err)
	}

	if len(merchants) < total {
		r.logger.Error("not enough merchants for transaction seeding", zap.Int("required", total), zap.Int("available", len(merchants)))
		return fmt.Errorf("not enough merchants: required %d, got %d", total, len(merchants))
	}

	months := make([]time.Time, 12)
	currentYear := time.Now().Year()
	for i := 0; i < 12; i++ {
		months[i] = time.Date(currentYear, time.Month(i+1), 1, 0, 0, 0, 0, time.UTC)
	}

	for i := 0; i < total; i++ {
		card := cards[i]
		merchant := merchants[i%len(merchants)]
		paymentMethod := paymentMethods[i%len(paymentMethods)]
		status := statusOptions[i%len(statusOptions)]

		monthIndex := i % 12
		transactionTime := months[monthIndex].Add(time.Duration(rand.Intn(28)) * 24 * time.Hour)

		request := db.CreateTransactionParams{
			CardNumber:      card.CardNumber,
			Amount:          int32(rand.Intn(1000000-50000) + 50000),
			PaymentMethod:   paymentMethod,
			MerchantID:      merchant.MerchantID,
			TransactionTime: transactionTime,
		}

		var transaction db.CreateTransactionRow
		err = r.db.GetContext(r.ctx, &transaction, seedCreateTransaction,
			request.CardNumber,
			request.Amount,
			request.PaymentMethod,
			request.MerchantID,
			request.TransactionTime,
		)
		if err != nil {
			r.logger.Error("failed to seed transaction", zap.Int("index", i), zap.Error(err))
			return fmt.Errorf("failed to seed transaction %d: %w", i, err)
		}

		var updated db.UpdateTransactionStatusRow
		err = r.db.GetContext(r.ctx, &updated, seedUpdateTransactionStatus, transaction.TransactionID, status)
		if err != nil {
			r.logger.Error("failed to update transaction status", zap.Int("transactionID", int(transaction.TransactionID)), zap.String("status", status), zap.Error(err))
			return fmt.Errorf("failed to update status for transaction ID %d: %w", transaction.TransactionID, err)
		}

		if i >= active {
			var trashed db.TrashTransactionRow
			err = r.db.GetContext(r.ctx, &trashed, seedTrashTransaction, transaction.TransactionID)
			if err != nil {
				r.logger.Error("failed to trash transaction", zap.Int("index", i), zap.Error(err))
				return fmt.Errorf("failed to trash transaction %d: %w", i, err)
			}
		}
	}

	r.logger.Info("transaction seeded successfully", zap.Int("total", total), zap.Int("active", active), zap.Int("trashed", trashed))
	return nil
}
