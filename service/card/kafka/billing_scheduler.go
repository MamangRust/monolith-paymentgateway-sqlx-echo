package kafka

import (
	"context"
	"time"

	"github.com/MamangRust/monolith-payment-gateway-card/service"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"go.uber.org/zap"
)

type BillingScheduler struct {
	billingService service.BillingEngineService
	logger         logger.LoggerInterface
}

func NewBillingScheduler(
	billingService service.BillingEngineService,
	logger logger.LoggerInterface,
) *BillingScheduler {
	return &BillingScheduler{
		billingService: billingService,
		logger:         logger,
	}
}

// StartPeriodic runs a periodic ticker that checks every hour if today matches the billing cycle day.
func (b *BillingScheduler) StartPeriodic(ctx context.Context) {
	ticker := time.NewTicker(1 * time.Hour)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			today := time.Now().Day()
			b.logger.Info("billing scheduler: checking billing cycle", zap.Int("day", today))
			affected, err := b.billingService.TriggerBillingCycle(ctx, today)
			if err != nil {
				b.logger.Error("billing scheduler: trigger failed", zap.Error(err))
			} else {
				b.logger.Info("billing scheduler: trigger completed", zap.Int("affected", affected))
			}
		}
	}
}
