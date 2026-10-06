package kafka

import (
	"context"
	"encoding/json"
	"time"

	"github.com/IBM/sarama"
	"github.com/MamangRust/monolith-payment-gateway-card/repository"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"go.uber.org/zap"
)

// FraudScoringConsumer implements sarama.ConsumerGroupHandler to consume
// card transaction creation events and compute a real-time fraud risk score.
type FraudScoringConsumer struct {
	authTxnRepo repository.CardAuthTransactionRepository
	logger      logger.LoggerInterface
}

// NewFraudScoringConsumer creates a new FraudScoringConsumer with the provided
// dependencies.
func NewFraudScoringConsumer(
	authTxnRepo repository.CardAuthTransactionRepository,
	logger logger.LoggerInterface,
) *FraudScoringConsumer {
	return &FraudScoringConsumer{
		authTxnRepo: authTxnRepo,
		logger:      logger,
	}
}

// Setup is called when a new consumer group session begins.
func (f *FraudScoringConsumer) Setup(session sarama.ConsumerGroupSession) error {
	f.logger.Info("fraud scoring consumer setup")
	return nil
}

// Cleanup is called when a consumer group session ends.
func (f *FraudScoringConsumer) Cleanup(session sarama.ConsumerGroupSession) error {
	f.logger.Info("fraud scoring consumer cleanup")
	return nil
}

// ConsumeClaim processes messages from the consumer group claim. For each message
// it parses the transaction payload, computes a risk score, persists the score, and
// publishes a fraud alert when the score is elevated.
func (f *FraudScoringConsumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		var payload struct {
			TxnID      string `json:"txn_id"`
			CardNumber string `json:"card_number"`
			Amount     int64  `json:"amount"`
			MCC        string `json:"mcc"`
			Status     string `json:"status"`
		}
		if err := json.Unmarshal(msg.Value, &payload); err != nil {
			f.logger.Error("failed to unmarshal transaction payload", zap.Error(err))
			session.MarkMessage(msg, "")
			continue
		}

		score := computeRiskScore(payload.Amount, payload.MCC)

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

		if err := f.authTxnRepo.UpdateRiskScore(ctx, payload.TxnID, score); err != nil {
			cancel()
			f.logger.Error("failed to update risk score; leaving message uncommitted for retry",
				zap.String("txn_id", payload.TxnID),
				zap.Int("score", score),
				zap.Error(err),
			)
			return err
		}

		cancel()

		if score < 30 {
			session.MarkMessage(msg, "")
			continue
		}

		session.MarkMessage(msg, "")
	}

	return nil
}

// computeRiskScore evaluates a transaction's risk level based on its amount and
// merchant category code (MCC). Returns a score from 0 to 100.
//
// Amount thresholds:
//   - > 10,000,000: +30 points
//   - > 5,000,000:  +15 points
//   - > 1,000,000:  +5 points
//
// MCC risk categories:
//   - High-risk ("7995" gambling): +40 points
//   - High-risk ("6051" crypto/non-financial institution): +35 points
//   - Medium-risk ("4829" money transfer): +20 points
//   - Elevated-risk ("5813" bars, "5933" pawn shops): +15 points
func computeRiskScore(amount int64, mcc string) int {
	score := 0

	switch {
	case amount > 10000000:
		score += 30
	case amount > 5000000:
		score += 15
	case amount > 1000000:
		score += 5
	}

	switch mcc {
	case "7995":
		score += 40
	case "6051":
		score += 35
	case "4829":
		score += 20
	case "5813", "5933":
		score += 15
	}

	if score > 100 {
		return 100
	}

	return score
}

// Ensure FraudScoringConsumer implements sarama.ConsumerGroupHandler.
var _ sarama.ConsumerGroupHandler = (*FraudScoringConsumer)(nil)
