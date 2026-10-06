package transactionstatsbycardservice

import (
	"context"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/errorhandler"
	transaction_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/transaction_errors/service"
	"github.com/MamangRust/monolith-payment-gateway-shared/observability"
	mencache "github.com/MamangRust/monolith-payment-gateway-transaction/redis/statsbycard"
	repository "github.com/MamangRust/monolith-payment-gateway-transaction/repository/statsbycard"
	"go.opentelemetry.io/otel/attribute"
	"go.uber.org/zap"
)

type transactionStatsByCardMethodServiceDeps struct {
	Cache mencache.TransactionStatsByCardMethodCache

	Repository repository.TransactionStatsByCardMethodRepository

	Logger logger.LoggerInterface

	Observability observability.TraceLoggerObservability
}

type transactionStatsByCardMethodService struct {
	cache mencache.TransactionStatsByCardMethodCache

	repository repository.TransactionStatsByCardMethodRepository

	logger logger.LoggerInterface

	observability observability.TraceLoggerObservability
}

func NewTransactionStatsByCardMethodService(params *transactionStatsByCardMethodServiceDeps) TransactionStatsByCardMethodService {
	return &transactionStatsByCardMethodService{
		cache:         params.Cache,
		repository:    params.Repository,
		logger:        params.Logger,
		observability: params.Observability,
	}
}

func (s *transactionStatsByCardMethodService) FindMonthlyPaymentMethodsByCardNumber(ctx context.Context, req *requests.MonthYearPaymentMethod) ([]*db.GetMonthlyPaymentMethodsByCardNumberRow, error) {
	const method = "FindMonthlyPaymentMethodsByCardNumber"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.String("card_number", observability.MaskIdentifier(req.CardNumber)),
		attribute.Int("year", req.Year))

	defer func() {
		end(status)
	}()

	if dbRows, found := s.cache.GetMonthlyPaymentMethodsByCardCache(ctx, req); found {
		logSuccess("Successfully fetched monthly payment methods by card number (from cache)",
			zap.String("card_number", observability.MaskIdentifier(req.CardNumber)),
			zap.Int("year", req.Year))
		return dbRows, nil
	}

	dbRows, err := s.repository.GetMonthlyPaymentMethodsByCardNumber(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetMonthlyPaymentMethodsByCardNumberRow](
			s.logger,
			transaction_errors.ErrFailedFindMonthlyPaymentMethodsByCard,
			method,
			span,

			zap.String("card_number", observability.MaskIdentifier(req.CardNumber)),
			zap.Int("year", req.Year),
		)
	}

	s.cache.SetMonthlyPaymentMethodsByCardCache(ctx, req, dbRows)

	logSuccess("Successfully fetched monthly payment methods by card number (from DB)",
		zap.String("card_number", observability.MaskIdentifier(req.CardNumber)),
		zap.Int("year", req.Year))

	return dbRows, nil
}

func (s *transactionStatsByCardMethodService) FindYearlyPaymentMethodsByCardNumber(ctx context.Context, req *requests.MonthYearPaymentMethod) ([]*db.GetYearlyPaymentMethodsByCardNumberRow, error) {
	const method = "FindYearlyPaymentMethodsByCardNumber"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.String("card_number", observability.MaskIdentifier(req.CardNumber)),
		attribute.Int("year", req.Year))

	defer func() {
		end(status)
	}()

	if dbRows, found := s.cache.GetYearlyPaymentMethodsByCardCache(ctx, req); found {
		logSuccess("Successfully fetched yearly payment methods by card number (from cache)",
			zap.String("card_number", observability.MaskIdentifier(req.CardNumber)),
			zap.Int("year", req.Year))
		return dbRows, nil
	}

	dbRows, err := s.repository.GetYearlyPaymentMethodsByCardNumber(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetYearlyPaymentMethodsByCardNumberRow](
			s.logger,
			transaction_errors.ErrFailedFindYearlyPaymentMethodsByCard,
			method,
			span,

			zap.String("card_number", observability.MaskIdentifier(req.CardNumber)),
			zap.Int("year", req.Year),
		)
	}

	s.cache.SetYearlyPaymentMethodsByCardCache(ctx, req, dbRows)

	logSuccess("Successfully fetched yearly payment methods by card number (from DB)",
		zap.String("card_number", observability.MaskIdentifier(req.CardNumber)),
		zap.Int("year", req.Year))

	return dbRows, nil
}
