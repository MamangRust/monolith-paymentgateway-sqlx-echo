package withdrawstatsservice

import (
	"context"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/errorhandler"
	withdraw_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/withdraw_errors/service"
	"github.com/MamangRust/monolith-payment-gateway-shared/observability"
	mencache "github.com/MamangRust/monolith-payment-gateway-withdraw/redis/stats"
	repository "github.com/MamangRust/monolith-payment-gateway-withdraw/repository/stats"
	"go.opentelemetry.io/otel/attribute"
	"go.uber.org/zap"
)

type WithdrawStatsAmountDeps struct {
	Cache mencache.WithdrawStatsAmountCache

	Repository repository.WithdrawStatsAmountRepository

	Logger logger.LoggerInterface

	Observability observability.TraceLoggerObservability
}

type withdrawStatsAmountService struct {
	cache mencache.WithdrawStatsAmountCache

	repository repository.WithdrawStatsAmountRepository

	logger logger.LoggerInterface

	observability observability.TraceLoggerObservability
}

func NewWithdrawStatsAmountService(params *WithdrawStatsAmountDeps) WithdrawStatsAmountService {
	return &withdrawStatsAmountService{
		cache:         params.Cache,
		repository:    params.Repository,
		logger:        params.Logger,
		observability: params.Observability,
	}
}

func (s *withdrawStatsAmountService) FindMonthlyWithdraws(ctx context.Context, year int) ([]*db.GetMonthlyWithdrawsRow, error) {
	const method = "FindMonthlyWithdraws"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", year))

	defer func() {
		end(status)
	}()

	if dbRows, found := s.cache.GetCachedMonthlyWithdraws(ctx, year); found {
		logSuccess("Successfully fetched monthly withdraws (from cache)", zap.Int("year", year))
		return dbRows, nil
	}

	dbRows, err := s.repository.GetMonthlyWithdraws(ctx, year)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetMonthlyWithdrawsRow](
			s.logger,
			withdraw_errors.ErrFailedFindMonthlyWithdraws,
			method,
			span,

			zap.Int("year", year),
		)
	}

	s.cache.SetCachedMonthlyWithdraws(ctx, year, dbRows)

	logSuccess("Successfully fetched monthly withdraws (from DB)", zap.Int("year", year))

	return dbRows, nil
}

func (s *withdrawStatsAmountService) FindYearlyWithdraws(ctx context.Context, year int) ([]*db.GetYearlyWithdrawsRow, error) {
	const method = "FindYearlyWithdraws"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", year))

	defer func() {
		end(status)
	}()

	if dbRows, found := s.cache.GetCachedYearlyWithdraws(ctx, year); found {
		logSuccess("Successfully fetched yearly withdraws (from cache)", zap.Int("year", year))
		return dbRows, nil
	}

	dbRows, err := s.repository.GetYearlyWithdraws(ctx, year)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetYearlyWithdrawsRow](
			s.logger,
			withdraw_errors.ErrFailedFindYearlyWithdraws,
			method,
			span,

			zap.Int("year", year),
		)
	}

	s.cache.SetCachedYearlyWithdraws(ctx, year, dbRows)

	logSuccess("Successfully fetched yearly withdraws (from DB)", zap.Int("year", year))

	return dbRows, nil
}
