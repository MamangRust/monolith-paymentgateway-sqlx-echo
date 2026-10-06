package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/IBM/sarama"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-saldo/service"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"go.uber.org/zap"
)

// saldoKafkaHandler is a struct that implements the sarama.ConsumerGroupHandler interface
type saldoKafkaHandler struct {
	logger       logger.LoggerInterface
	saldoService service.SaldoCommandService
	ctx          context.Context
}

// NewSaldoKafkaHandler creates a new Kafka consumer group handler for processing saldo-related Kafka messages.
//
// It takes a saldo command service and a logger as parameters.
// The handler is responsible for consuming messages from Kafka topics related to saldo operations.
// It implements the sarama.ConsumerGroupHandler interface to manage consumer group lifecycle events.
func NewSaldoKafkaHandler(saldoService service.SaldoCommandService, logger logger.LoggerInterface, ctx context.Context) sarama.ConsumerGroupHandler {
	return &saldoKafkaHandler{
		saldoService: saldoService,
		logger:       logger,
		ctx:          ctx,
	}
}

// Setup is called when the consumer group is first initialized.
func (s *saldoKafkaHandler) Setup(session sarama.ConsumerGroupSession) error {
	s.logger.Info("saldo kafka handler setup")
	return nil
}

// Cleanup is called when the consumer group is closed.
func (s *saldoKafkaHandler) Cleanup(session sarama.ConsumerGroupSession) error {
	s.logger.Info("saldo kafka handler cleanup")
	return nil
}

// ConsumeClaim processes incoming Kafka messages from the specified consumer group claim.
// It unmarshals each message into a payload map and creates a new saldo request.
// If a valid saldo request is found, it calls CreateSaldoIfNotExists on the saldo
// service. The operation is atomic and never overwrites an existing API-created
// balance with the default balance from the card-created event.
// Each message is marked as processed in the consumer group session.
func (s *saldoKafkaHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	s.logger.Info("saldo kafka handler consume claim")

	for msg := range claim.Messages() {
		ctx, cancel := context.WithTimeout(s.ctx, 20*time.Second)

		var payload map[string]interface{}

		if err := json.Unmarshal(msg.Value, &payload); err != nil {
			cancel()
			// Malformed input cannot be repaired by retrying.
			s.logger.Error("invalid saldo Kafka payload", zap.Error(err))
			session.MarkMessage(msg, "")
			continue
		}

		if event, ok := payload["event"].(string); ok && event == "saldo.cache.invalidate" {
			s.saldoService.InvalidateSaldoCache(ctx)
			cancel()
			session.MarkMessage(msg, "")
			continue
		}

		cardNumber, ok := payload["card_number"].(string)
		if !ok || cardNumber == "" {
			cancel()
			s.logger.Error("saldo payload card_number missing or invalid")
			session.MarkMessage(msg, "")
			continue
		}

		totalBalanceFloat, ok := payload["total_balance"].(float64)
		if !ok {
			cancel()
			s.logger.Error("saldo payload total_balance missing or invalid")
			session.MarkMessage(msg, "")
			continue
		}

		totalBalance := int(totalBalanceFloat)

		errRes := s.saldoService.CreateSaldoIfNotExists(ctx, &requests.CreateSaldoRequest{
			CardNumber:   cardNumber,
			TotalBalance: int(totalBalance),
		})

		if errRes != nil {
			cancel()
			s.logger.Error("saldo service failed; leaving Kafka message uncommitted", zap.Error(errRes))
			return fmt.Errorf("saldo service error: %w", errRes)
		}

		cancel()
		session.MarkMessage(msg, "")
	}

	s.logger.Info("saldo kafka handler consume claim success", zap.Bool("success", true))

	return nil
}
