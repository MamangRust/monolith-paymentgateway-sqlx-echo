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

type transactionStatsAmountServiceDeps struct {
	Cache mencache.TransactionStatsAmountCache

	Repository repository.TransactionStatsAmountRepository

	Logger logger.LoggerInterface

	Observability observability.TraceLoggerObservability
}

type transactionStatsAmountService struct {
	cache mencache.TransactionStatsAmountCache

	repository repository.TransactionStatsAmountRepository

	logger logger.LoggerInterface

	observability observability.TraceLoggerObservability
}

func NewTransactionStatsAmountService(params *transactionStatsAmountServiceDeps) TransactionStatsAmountService {
	return &transactionStatsAmountService{
		cache:         params.Cache,
		repository:    params.Repository,
		logger:        params.Logger,
		observability: params.Observability,
	}
}

func (s *transactionStatsAmountService) FindMonthlyAmounts(ctx context.Context, year int) ([]*db.GetMonthlyAmountsRow, error) {
	const method = "FindMonthlyAmounts"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", year))

	defer func() {
		end(status)
	}()

	if dbRows, found := s.cache.GetMonthlyAmountsCache(ctx, year); found {
		s.logger.Info("Cache hit for monthly amounts", zap.Int("year", year))
		status = "ok"
		logSuccess("Successfully fetched monthly amounts (from cache)", zap.Int("year", year))
		return dbRows, nil
	}

	dbRows, err := s.repository.GetMonthlyAmounts(ctx, year)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetMonthlyAmountsRow](
			s.logger,
			transaction_errors.ErrFailedFindMonthlyAmounts,
			method,
			span,

			zap.Int("year", year),
		)
	}

	s.cache.SetMonthlyAmountsCache(ctx, year, dbRows)

	logSuccess("Successfully fetched monthly amounts (from DB)", zap.Int("year", year))

	return dbRows, nil
}

func (s *transactionStatsAmountService) FindYearlyAmounts(ctx context.Context, year int) ([]*db.GetYearlyAmountsRow, error) {
	const method = "FindYearlyAmounts"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", year))

	defer func() {
		end(status)
	}()

	if dbRows, found := s.cache.GetYearlyAmountsCache(ctx, year); found {
		s.logger.Info("Cache hit for yearly amounts", zap.Int("year", year))
		status = "ok"
		logSuccess("Successfully fetched yearly amounts (from cache)", zap.Int("year", year))
		return dbRows, nil
	}

	dbRows, err := s.repository.GetYearlyAmounts(ctx, year)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetYearlyAmountsRow](
			s.logger,
			transaction_errors.ErrFailedFindYearlyAmounts,
			method,
			span,

			zap.Int("year", year),
		)
	}

	s.cache.SetYearlyAmountsCache(ctx, year, dbRows)

	logSuccess("Successfully fetched yearly amounts (from DB)", zap.Int("year", year))

	return dbRows, nil
}
