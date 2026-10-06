package topupstatsservice

import (
	"context"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/errorhandler"
	topup_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/topup_errors/service"
	"github.com/MamangRust/monolith-payment-gateway-shared/observability"
	mencache "github.com/MamangRust/monolith-payment-gateway-topup/redis/stats"
	repository "github.com/MamangRust/monolith-payment-gateway-topup/repository/stats"
	"go.opentelemetry.io/otel/attribute"
	"go.uber.org/zap"
)

type topupStatsAmountDeps struct {
	Cache mencache.TopupStatsAmountCache

	Repository repository.TopupStatsAmountRepository

	Logger logger.LoggerInterface

	Observability observability.TraceLoggerObservability
}

type topupStatsAmountService struct {
	cache mencache.TopupStatsAmountCache

	repository repository.TopupStatsAmountRepository

	logger logger.LoggerInterface

	observability observability.TraceLoggerObservability
}

func NewTopupStatsAmountService(params *topupStatsAmountDeps) TopupStatsAmountService {

	return &topupStatsAmountService{
		cache:         params.Cache,
		repository:    params.Repository,
		logger:        params.Logger,
		observability: params.Observability,
	}
}

func (s *topupStatsAmountService) FindMonthlyTopupAmounts(ctx context.Context, year int) ([]*db.GetMonthlyTopupAmountsRow, error) {
	const method = "FindMonthlyTopupAmounts"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", year))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetMonthlyTopupAmountsCache(ctx, year); found {
		logSuccess("Successfully fetched monthly topup amounts from cache", zap.Int("year", year))
		return data, nil
	}

	dbRows, err := s.repository.GetMonthlyTopupAmounts(ctx, year)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetMonthlyTopupAmountsRow](
			s.logger,
			topup_errors.ErrFailedFindMonthlyTopupAmounts,
			method,
			span,

			zap.Int("year", year),
		)
	}

	s.cache.SetMonthlyTopupAmountsCache(ctx, year, dbRows)

	logSuccess("Successfully fetched monthly topup amounts (from DB)", zap.Int("year", year))

	return dbRows, nil
}

func (s *topupStatsAmountService) FindYearlyTopupAmounts(ctx context.Context, year int) ([]*db.GetYearlyTopupAmountsRow, error) {
	const method = "FindYearlyTopupAmounts"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", year))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetYearlyTopupAmountsCache(ctx, year); found {
		logSuccess("Successfully fetched yearly topup amounts from cache", zap.Int("year", year))
		return data, nil
	}

	dbRows, err := s.repository.GetYearlyTopupAmounts(ctx, year)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetYearlyTopupAmountsRow](
			s.logger,
			topup_errors.ErrFailedFindYearlyTopupAmounts,
			method,
			span,

			zap.Int("year", year),
		)
	}

	s.cache.SetYearlyTopupAmountsCache(ctx, year, dbRows)

	logSuccess("Successfully fetched yearly topup amounts (from DB)", zap.Int("year", year))

	return dbRows, nil
}
