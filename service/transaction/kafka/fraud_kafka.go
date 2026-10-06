package kafka

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"time"

	"github.com/IBM/sarama"
	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/observability"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

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

const updateTransactionStatusDirect = `UPDATE transactions
SET
    status = $2,
    updated_at = CURRENT_TIMESTAMP
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
    fraud_score,
    created_at,
    updated_at`

const updateTransactionFraudScore = `UPDATE transactions
SET
    fraud_score = $2,
    updated_at = CURRENT_TIMESTAMP
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
    fraud_score,
    created_at,
    updated_at`

// fraudKafkaHandler processes fraud check events from the transaction-fraud-check topic.
type fraudKafkaHandler struct {
	logger logger.LoggerInterface
	db     *sqlx.DB
}

// NewFraudKafkaHandler initializes a new fraudKafkaHandler with direct database access.
func NewFraudKafkaHandler(
	logger logger.LoggerInterface,
	db *sqlx.DB,
) *fraudKafkaHandler {
	return &fraudKafkaHandler{
		logger: logger,
		db:     db,
	}
}

func (s *fraudKafkaHandler) Setup(session sarama.ConsumerGroupSession) error {
	s.logger.Info("fraud kafka handler setup")
	return nil
}

func (s *fraudKafkaHandler) Cleanup(session sarama.ConsumerGroupSession) error {
	s.logger.Info("fraud kafka handler cleanup")
	return nil
}

// FraudCheckPayload represents the Kafka message payload for fraud checking.
type FraudCheckPayload struct {
	TransactionID int    `json:"transaction_id"`
	CardNumber    string `json:"card_number"`
	Amount        int32  `json:"amount"`
	CardType      string `json:"card_type"`
}

func (s *fraudKafkaHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	for msg := range claim.Messages() {
		var payload FraudCheckPayload
		if err := json.Unmarshal(msg.Value, &payload); err != nil {
			s.logger.Error("failed to unmarshal fraud check payload", zap.Error(err))
			session.MarkMessage(msg, "")
			continue
		}

		score, err := s.evaluateFraudRules(ctx, &payload)
		if err != nil {
			s.logger.Error("failed to evaluate fraud rules; leaving message uncommitted for retry", zap.Error(err), zap.Int("transaction_id", payload.TransactionID))
			return err
		}

		// Update fraud score on the transaction
		if _, err := updateTransactionRow(ctx, s.db, updateTransactionFraudScore, payload.TransactionID, int32(score)); err != nil {
			s.logger.Error("failed to update fraud score; leaving message uncommitted for retry", zap.Error(err), zap.Int("transaction_id", payload.TransactionID))
			return err
		}

		s.logger.Info("fraud score evaluated",
			zap.Int("transaction_id", payload.TransactionID),
			zap.Int("score", score),
			zap.String("card_number", observability.MaskIdentifier(payload.CardNumber)),
		)

		// If score >= 80, suspend the card and flag transaction as fraud
		if score >= 80 {
			s.logger.Warn("high fraud score detected, triggering card suspension",
				zap.Int("transaction_id", payload.TransactionID),
				zap.Int("score", score),
				zap.String("card_number", observability.MaskIdentifier(payload.CardNumber)),
			)

			// Flag transaction as flagged_fraud
			if _, err := updateTransactionRow(ctx, s.db, updateTransactionStatusDirect, payload.TransactionID, "flagged_fraud"); err != nil {
				s.logger.Error("failed to flag transaction as fraud; leaving message uncommitted for retry", zap.Error(err), zap.Int("transaction_id", payload.TransactionID))
				return err
			}
		}

		session.MarkMessage(msg, "")
	}

	return nil
}

func (s *fraudKafkaHandler) evaluateFraudRules(ctx context.Context, payload *FraudCheckPayload) (int, error) {
	score := 0

	// Rule 1: Large transaction > 10,000,000 IDR
	if payload.Amount > 10_000_000 {
		score += 30
		s.logger.Debug("fraud rule triggered: large amount", zap.Int32("amount", payload.Amount))
	}

	// Rule 2: Check card status - if suspicious, increase score
	var card db.GetCardByCardNumberRow
	err := s.db.GetContext(ctx, &card, getCardByCardNumber, payload.CardNumber)
	if err == nil {
		if card.Status == "suspended" {
			score += 100
			s.logger.Debug("fraud rule triggered: suspended card attempted transaction")
		}
	} else {
		s.logger.Warn("could not fetch card for fraud evaluation; retrying message", zap.String("card_number", observability.MaskIdentifier(payload.CardNumber)), zap.Error(err))
		return 0, err
	}

	return score, nil
}

// updateTransactionRow runs one of the transaction status/fraud-score update
// queries and returns the updated transaction row. sql.ErrNoRows is surfaced
// to the caller so the Kafka message stays uncommitted for retry.
func updateTransactionRow(ctx context.Context, dbConn *sqlx.DB, query string, transactionID int, arg any) (*db.UpdateTransactionStatusDirectRow, error) {
	var row db.UpdateTransactionStatusDirectRow
	if err := dbConn.GetContext(ctx, &row, query, int32(transactionID), arg); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
		return nil, err
	}
	return &row, nil
}

// StartFraudConsumer starts the Kafka consumer for fraud checking with a
// background context for backwards compatibility. Prefer
// StartFraudConsumerWithContext for services that own a shutdown context.
func StartFraudConsumer(
	consumerGroup sarama.ConsumerGroup,
	logger logger.LoggerInterface,
	db *sqlx.DB,
) error {
	return StartFraudConsumerWithContext(context.Background(), consumerGroup, logger, db)
}

// StartFraudConsumerWithContext starts the fraud consumer until ctx is
// cancelled. Closing the consumer group also terminates the error/rebalance
// loop instead of leaking a goroutine after service shutdown.
func StartFraudConsumerWithContext(
	ctx context.Context,
	consumerGroup sarama.ConsumerGroup,
	logger logger.LoggerInterface,
	db *sqlx.DB,
) error {
	handler := NewFraudKafkaHandler(logger, db)
	topics := []string{"transaction-fraud-check"}

	go func() {
		for {
			if ctx.Err() != nil {
				return
			}
			err := consumerGroup.Consume(ctx, topics, handler)
			if ctx.Err() != nil {
				return
			}
			if err != nil {
				logger.Error("error from fraud consumer", zap.Error(err))
				continue
			}
		}
	}()

	go func() {
		<-ctx.Done()
		if err := consumerGroup.Close(); err != nil {
			logger.Error("failed to close fraud consumer", zap.Error(err))
		}
	}()

	logger.Info("fraud consumer started", zap.Strings("topics", topics))
	return nil
}
