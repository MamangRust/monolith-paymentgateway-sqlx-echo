package transactionstatsservice

import (
	"context"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/errorhandler"
	transaction_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/transaction_errors/service"
	"github.com/MamangRust/monolith-payment-gateway-shared/observability"
	mencache "github.com/MamangRust/monolith-payment-gateway-transaction/redis/stats"
	repository "github.com/MamangRust/monolith-payment-gateway-transaction/repository/stats"
	"go.opentelemetry.io/otel/attribute"
	"go.uber.org/zap"
)

type transactionStatsMethodServiceDeps struct {
	Cache mencache.TransactionStatsMethodCache

	Repository repository.TransactionStatsMethodRepository

	Logger logger.LoggerInterface

	Observability observability.TraceLoggerObservability
}

type transactionStatsMethodService struct {
	cache mencache.TransactionStatsMethodCache

	repository repository.TransactionStatsMethodRepository

	logger logger.LoggerInterface

	observability observability.TraceLoggerObservability
}

func NewTransactionStatsMethodService(params *transactionStatsMethodServiceDeps) TransactionStatsMethodService {
	return &transactionStatsMethodService{
		cache:         params.Cache,
		repository:    params.Repository,
		logger:        params.Logger,
		observability: params.Observability,
	}
}

func (s *transactionStatsMethodService) FindMonthlyPaymentMethods(ctx context.Context, year int) ([]*db.GetMonthlyPaymentMethodsRow, error) {
	const method = "FindMonthlyPaymentMethods"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", year))

	defer func() {
		end(status)
	}()

	if dbRows, found := s.cache.GetMonthlyPaymentMethodsCache(ctx, year); found {
		logSuccess("Successfully fetched monthly payment methods (from cache)", zap.Int("year", year))
		return dbRows, nil
	}

	dbRows, err := s.repository.GetMonthlyPaymentMethods(ctx, year)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetMonthlyPaymentMethodsRow](
			s.logger,
			transaction_errors.ErrFailedFindMonthlyPaymentMethods,
			method,
			span,

			zap.Int("year", year),
		)
	}

	s.cache.SetMonthlyPaymentMethodsCache(ctx, year, dbRows)

	logSuccess("Successfully fetched monthly payment methods (from DB)", zap.Int("year", year))

	return dbRows, nil
}

func (s *transactionStatsMethodService) FindYearlyPaymentMethods(ctx context.Context, year int) ([]*db.GetYearlyPaymentMethodsRow, error) {
	const method = "FindYearlyPaymentMethods"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", year))

	defer func() {
		end(status)
	}()

	if dbRows, found := s.cache.GetYearlyPaymentMethodsCache(ctx, year); found {
		logSuccess("Successfully fetched yearly payment methods (from cache)", zap.Int("year", year))
		return dbRows, nil
	}

	dbRows, err := s.repository.GetYearlyPaymentMethods(ctx, year)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetYearlyPaymentMethodsRow](
			s.logger,
			transaction_errors.ErrFailedFindYearlyPaymentMethods,
			method,
			span,

			zap.Int("year", year),
		)
	}

	s.cache.SetYearlyPaymentMethodsCache(ctx, year, dbRows)

	logSuccess("Successfully fetched yearly payment methods (from DB)", zap.Int("year", year))

	return dbRows, nil
}
