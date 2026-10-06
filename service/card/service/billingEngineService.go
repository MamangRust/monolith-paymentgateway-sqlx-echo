package service

import (
	"context"

	"github.com/MamangRust/monolith-payment-gateway-card/repository"
	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-pkg/kafka"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	sharederrorhandler "github.com/MamangRust/monolith-payment-gateway-shared/errorhandler"
	card_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/card_errors/service"
	"github.com/MamangRust/monolith-payment-gateway-shared/observability"
	"go.opentelemetry.io/otel/attribute"
	"go.uber.org/zap"
)

// BillingEngineServiceDeps defines dependencies for billingEngineService.
type BillingEngineServiceDeps struct {
	CardCommandRepository  repository.CardCommandRepository
	BillingCycleRepository repository.BillingCycleRepository
	CardQueryRepository    repository.CardQueryRepository
	Kafka                  *kafka.Kafka
	Logger                 logger.LoggerInterface
	Observability          observability.TraceLoggerObservability
}

// billingEngineService implements BillingEngineService.
type billingEngineService struct {
	cardCommandRepository  repository.CardCommandRepository
	billingCycleRepository repository.BillingCycleRepository
	cardQueryRepository    repository.CardQueryRepository
	kafka                  *kafka.Kafka
	logger                 logger.LoggerInterface
	observability          observability.TraceLoggerObservability
}

// NewBillingEngineService creates a new BillingEngineService instance.
func NewBillingEngineService(deps *BillingEngineServiceDeps) BillingEngineService {
	return &billingEngineService{
		cardCommandRepository:  deps.CardCommandRepository,
		billingCycleRepository: deps.BillingCycleRepository,
		cardQueryRepository:    deps.CardQueryRepository,
		kafka:                  deps.Kafka,
		logger:                 deps.Logger,
		observability:          deps.Observability,
	}
}

// TriggerBillingCycle triggers the billing cycle for a given billing cycle day.
func (s *billingEngineService) TriggerBillingCycle(ctx context.Context, billingCycleDay int) (int, error) {
	const method = "TriggerBillingCycle"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("billing_cycle_day", billingCycleDay))

	defer func() {
		end(status)
	}()

	cycles, err := s.cardCommandRepository.ProcessBillingCycles(ctx, billingCycleDay)
	if err != nil {
		status = "error"
		return sharederrorhandler.HandleError[int](
			s.logger,
			card_errors.ErrFailedProcessBillingCycles,
			method,
			span,
			zap.Int("billing_cycle_day", billingCycleDay),
		)
	}

	// Statement events are enqueued by the same SQL statement as the billing
	// cycle. The card outbox relay owns Kafka delivery and retry; publishing
	// here would create duplicates and would lose events on a scheduler retry.

	logSuccess("Successfully triggered billing cycle", zap.Int("billing_cycle_day", billingCycleDay), zap.Int("affected", len(cycles)))

	return len(cycles), nil
}

// GetStatement retrieves the most recent billing statement for a card.
func (s *billingEngineService) GetStatement(ctx context.Context, cardNumber string) (*db.BillingCycle, error) {
	const method = "GetStatement"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.String("card_number", observability.MaskIdentifier(cardNumber)))

	defer func() {
		end(status)
	}()

	cycles, err := s.billingCycleRepository.GetBillingCyclesByCardNumber(ctx, cardNumber)
	if err != nil {
		status = "error"
		return sharederrorhandler.HandleError[*db.BillingCycle](
			s.logger,
			err,
			method,
			span,
			zap.String("card_number", observability.MaskIdentifier(cardNumber)),
		)
	}

	if len(cycles) == 0 {
		status = "error"
		return sharederrorhandler.HandleError[*db.BillingCycle](
			s.logger,
			card_errors.ErrCardNotFoundRes,
			method,
			span,
			zap.String("card_number", observability.MaskIdentifier(cardNumber)),
		)
	}

	logSuccess("Successfully retrieved statement", zap.String("card_number", observability.MaskIdentifier(cardNumber)), zap.Int("billing_id", int(cycles[0].BillingID)))

	return cycles[0], nil
}

// GetStatementsByCard retrieves paginated billing statements for a card.
func (s *billingEngineService) GetStatementsByCard(ctx context.Context, cardNumber string, page, pageSize int) ([]*db.BillingCycle, error) {
	const method = "GetStatementsByCard"

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.String("card_number", observability.MaskIdentifier(cardNumber)),
		attribute.Int("page", page),
		attribute.Int("page_size", pageSize))

	defer func() {
		end(status)
	}()

	cycles, err := s.billingCycleRepository.GetBillingCyclesByCardNumber(ctx, cardNumber)
	if err != nil {
		status = "error"
		return sharederrorhandler.HandleError[[]*db.BillingCycle](
			s.logger,
			err,
			method,
			span,
			zap.String("card_number", observability.MaskIdentifier(cardNumber)),
		)
	}

	startIdx := (page - 1) * pageSize
	if startIdx < 0 {
		startIdx = 0
	}

	if startIdx >= len(cycles) {
		logSuccess("Successfully retrieved billing statements (empty result)", zap.String("card_number", observability.MaskIdentifier(cardNumber)), zap.Int("total", len(cycles)))
		return []*db.BillingCycle{}, nil
	}

	endIdx := startIdx + pageSize
	if endIdx > len(cycles) {
		endIdx = len(cycles)
	}

	logSuccess("Successfully retrieved billing statements", zap.String("card_number", observability.MaskIdentifier(cardNumber)), zap.Int("count", len(cycles[startIdx:endIdx])))

	return cycles[startIdx:endIdx], nil
}

// GetBillingCyclesByCardNumber retrieves all billing cycles for a card.
func (s *billingEngineService) GetBillingCyclesByCardNumber(ctx context.Context, cardNumber string) ([]*db.BillingCycle, error) {
	const method = "GetBillingCyclesByCardNumber"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.String("card_number", observability.MaskIdentifier(cardNumber)))

	defer func() {
		end(status)
	}()

	cycles, err := s.billingCycleRepository.GetBillingCyclesByCardNumber(ctx, cardNumber)
	if err != nil {
		status = "error"
		return sharederrorhandler.HandleError[[]*db.BillingCycle](
			s.logger,
			err,
			method,
			span,
			zap.String("card_number", observability.MaskIdentifier(cardNumber)),
		)
	}

	logSuccess("Successfully retrieved billing cycles", zap.String("card_number", observability.MaskIdentifier(cardNumber)), zap.Int("count", len(cycles)))

	return cycles, nil
}
