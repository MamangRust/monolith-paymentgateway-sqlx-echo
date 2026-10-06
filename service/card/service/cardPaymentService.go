package service

import (
	"context"

	"github.com/MamangRust/monolith-payment-gateway-card/repository"
	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	statemachine "github.com/MamangRust/monolith-payment-gateway-shared/domain/status"
	sharederrorhandler "github.com/MamangRust/monolith-payment-gateway-shared/errorhandler"
	sharedErrors "github.com/MamangRust/monolith-payment-gateway-shared/errors"
	card_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/card_errors/service"
	"github.com/MamangRust/monolith-payment-gateway-shared/observability"
	"go.opentelemetry.io/otel/attribute"
	"go.uber.org/zap"
)

// CardPaymentServiceDeps defines dependencies for cardPaymentService.
type CardPaymentServiceDeps struct {
	CardPaymentRepository  repository.CardPaymentRepository
	BillingCycleRepository repository.BillingCycleRepository
	CardCommandRepository  repository.CardCommandRepository
	Logger                 logger.LoggerInterface
	Observability          observability.TraceLoggerObservability
}

// cardPaymentService implements CardPaymentService.
type cardPaymentService struct {
	cardPaymentRepository  repository.CardPaymentRepository
	billingCycleRepository repository.BillingCycleRepository
	cardCommandRepository  repository.CardCommandRepository
	logger                 logger.LoggerInterface
	observability          observability.TraceLoggerObservability
}

func NewCardPaymentService(deps *CardPaymentServiceDeps) CardPaymentService {
	return &cardPaymentService{
		cardPaymentRepository:  deps.CardPaymentRepository,
		billingCycleRepository: deps.BillingCycleRepository,
		cardCommandRepository:  deps.CardCommandRepository,
		logger:                 deps.Logger,
		observability:          deps.Observability,
	}
}

func (s *cardPaymentService) PostPayment(ctx context.Context, request *requests.PostPaymentRequest) (*db.CardPayment, error) {
	const method = "PostPayment"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.String("card_number", observability.MaskIdentifier(request.CardNumber)),
		attribute.Int64("amount", request.Amount),
		attribute.String("payment_channel", request.PaymentChannel))

	defer func() {
		end(status)
	}()

	// Fast-path an already persisted payment before validating the billing
	// cycle. A legitimate retry happens after the first request has already
	// moved unpaid -> paid, so validating the cycle first would reject replay.
	if request.ReferenceID != "" {
		existing, lookupErr := s.cardPaymentRepository.GetPaymentByReferenceID(ctx, request.ReferenceID)
		if lookupErr != nil {
			status = "error"
			return sharederrorhandler.HandleError[*db.CardPayment](s.logger, lookupErr, method, span,
				zap.String("reference_id", observability.MaskIdentifier(request.ReferenceID)))
		}
		if existing != nil {
			if err := validatePaymentReplayRequest(existing, request); err != nil {
				status = "error"
				return sharederrorhandler.HandleError[*db.CardPayment](s.logger, err, method, span)
			}
			span.SetAttributes(attribute.Bool("idempotent_replay", true))
			logSuccess("Idempotent payment replay returned existing payment", zap.Int32("payment_id", existing.PaymentID))
			return existing, nil
		}
	}

	// A new billing payment is a lifecycle transition, not a blind insert.
	// The repository repeats the expected-state predicate atomically to close
	// the race between this read and the payment write.
	expectedBillingStatus := ""
	targetBillingStatus := ""
	if request.BillingID != nil {
		cycle, cycleErr := s.billingCycleRepository.GetBillingCycleByID(ctx, *request.BillingID)
		if cycleErr != nil {
			status = "error"
			return sharederrorhandler.HandleError[*db.CardPayment](s.logger, cycleErr, method, span,
				zap.Int("billing_id", *request.BillingID))
		}
		if transitionErr := statemachine.ValidateTransition(statemachine.DomainCardBilling, cycle.Status, statemachine.Paid); transitionErr != nil {
			// A concurrent first request may have completed the transition after
			// the fast-path lookup. Resolve that race as a replay before returning
			// the state-machine error to the caller.
			if request.ReferenceID != "" {
				existing, replayErr := s.cardPaymentRepository.GetPaymentByReferenceID(ctx, request.ReferenceID)
				if replayErr != nil {
					status = "error"
					return sharederrorhandler.HandleError[*db.CardPayment](s.logger, replayErr, method, span,
						zap.String("reference_id", observability.MaskIdentifier(request.ReferenceID)))
				}
				if existing != nil {
					if validationErr := validatePaymentReplayRequest(existing, request); validationErr != nil {
						status = "error"
						return sharederrorhandler.HandleError[*db.CardPayment](s.logger, validationErr, method, span)
					}
					span.SetAttributes(attribute.Bool("idempotent_replay", true))
					logSuccess("Idempotent payment replay returned existing payment", zap.Int32("payment_id", existing.PaymentID))
					return existing, nil
				}
			}
			status = "error"
			return sharederrorhandler.HandleError[*db.CardPayment](s.logger, transitionErr, method, span,
				zap.Int("billing_id", *request.BillingID), zap.String("current_status", cycle.Status))
		}
		expectedBillingStatus = cycle.Status
		targetBillingStatus = statemachine.Paid
	}

	result, err := s.cardPaymentRepository.PostPaymentIdempotent(ctx, request, expectedBillingStatus, targetBillingStatus)
	if err != nil {
		status = "error"
		return sharederrorhandler.HandleError[*db.CardPayment](
			s.logger,
			err,
			method,
			span,
			zap.String("card_number", observability.MaskIdentifier(request.CardNumber)),
			zap.Int64("amount", request.Amount),
		)
	}

	res := result.Payment
	if result.Replayed {
		span.SetAttributes(attribute.Bool("idempotent_replay", true))
		logSuccess("Idempotent payment replay returned existing payment", zap.Int32("payment_id", res.PaymentID))
		return res, nil
	}

	logSuccess("Successfully posted payment",
		zap.Int32("payment_id", res.PaymentID),
		zap.String("card_number", observability.MaskIdentifier(request.CardNumber)),
		zap.Int64("amount", res.Amount),
		zap.String("channel", request.PaymentChannel),
	)

	return res, nil
}

func validatePaymentReplayRequest(existing *db.CardPayment, request *requests.PostPaymentRequest) error {
	if existing == nil || existing.CardNumber != request.CardNumber || existing.Amount != request.Amount || existing.PaymentChannel != request.PaymentChannel {
		return sharedErrors.ErrConflict.WithMessage("card payment reference already used with different payload")
	}
	if request.BillingID == nil {
		if existing.BillingID != nil {
			return sharedErrors.ErrConflict.WithMessage("card payment reference already used with different billing cycle")
		}
		return nil
	}
	if existing.BillingID == nil || *existing.BillingID != int32(*request.BillingID) {
		return sharedErrors.ErrConflict.WithMessage("card payment reference already used with different billing cycle")
	}
	return nil
}

func (s *cardPaymentService) GetPaymentHistory(ctx context.Context, cardNumber string, page, pageSize int) ([]*db.GetCardPaymentsByCardNumberRow, error) {
	const method = "GetPaymentHistory"

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

	res, err := s.cardPaymentRepository.GetPaymentHistory(ctx, cardNumber, page, pageSize)
	if err != nil {
		status = "error"
		_, _, err = sharederrorhandler.HandlerErrorPagination[[]*db.GetCardPaymentsByCardNumberRow](
			s.logger,
			card_errors.ErrFailedFindByCardNumber,
			method,
			span,
			zap.String("card_number", observability.MaskIdentifier(cardNumber)),
			zap.Int("page", page),
			zap.Int("page_size", pageSize),
		)
		return nil, err
	}

	logSuccess("Successfully retrieved payment history",
		zap.String("card_number", observability.MaskIdentifier(cardNumber)),
		zap.Int("count", len(res)),
	)

	return res, nil
}

func (s *cardPaymentService) CountPayments(ctx context.Context, cardNumber string) (int, error) {
	const method = "CountPayments"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.String("card_number", observability.MaskIdentifier(cardNumber)))

	defer func() {
		end(status)
	}()

	total, err := s.cardPaymentRepository.CountPayments(ctx, cardNumber)
	if err != nil {
		status = "error"
		return sharederrorhandler.HandleError[int](
			s.logger,
			card_errors.ErrFailedFindByCardNumber,
			method,
			span,
			zap.String("card_number", observability.MaskIdentifier(cardNumber)),
		)
	}

	logSuccess("Successfully counted payments",
		zap.Int("total", total),
		zap.String("card_number", observability.MaskIdentifier(cardNumber)),
	)

	return total, nil
}
