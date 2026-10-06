package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/IBM/sarama"
	"github.com/MamangRust/monolith-payment-gateway-card/service"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"go.uber.org/zap"
)

type BillingTriggerHandler struct {
	billingService service.BillingEngineService
	logger         logger.LoggerInterface
}

func NewBillingTriggerHandler(
	billingService service.BillingEngineService,
	logger logger.LoggerInterface,
) *BillingTriggerHandler {
	return &BillingTriggerHandler{
		billingService: billingService,
		logger:         logger,
	}
}

// Setup is called at the beginning of a new consumer group session.
func (h *BillingTriggerHandler) Setup(session sarama.ConsumerGroupSession) error {
	h.logger.Info("billing trigger handler setup")
	return nil
}

// Cleanup is called at the end of a consumer group session.
func (h *BillingTriggerHandler) Cleanup(session sarama.ConsumerGroupSession) error {
	h.logger.Info("billing trigger handler cleanup")
	return nil
}

// ConsumeClaim processes messages from the "card.billing.trigger" topic.
// Payload expects: { "billing_cycle_day": 25 }
func (h *BillingTriggerHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		// Scope the deadline to one message. A single slow/failing message must
		// not leave all later messages using an expired context.
		ctx, cancel := context.WithTimeout(session.Context(), 30*time.Second)
		var payload struct {
			BillingCycleDay int `json:"billing_cycle_day"`
		}
		if err := json.Unmarshal(msg.Value, &payload); err != nil {
			cancel()
			h.logger.Error("failed to unmarshal billing trigger payload", zap.Error(err))
			// Malformed input is non-retryable and is safe to acknowledge.
			session.MarkMessage(msg, "")
			continue
		}

		if payload.BillingCycleDay <= 0 || payload.BillingCycleDay > 31 {
			cancel()
			h.logger.Warn("invalid billing_cycle_day", zap.Int("day", payload.BillingCycleDay))
			// Invalid business input is non-retryable and is safe to acknowledge.
			session.MarkMessage(msg, "")
			continue
		}

		affected, err := h.billingService.TriggerBillingCycle(ctx, payload.BillingCycleDay)
		cancel()
		if err != nil {
			h.logger.Error("billing trigger failed", zap.Int("day", payload.BillingCycleDay), zap.Error(err))
			// Do not mark infrastructure/service errors: Sarama will redeliver.
			return fmt.Errorf("billing trigger error: %w", err)
		}

		h.logger.Info("billing cycle triggered", zap.Int("day", payload.BillingCycleDay), zap.Int("affected", affected))
		session.MarkMessage(msg, "")
	}

	return nil
}
