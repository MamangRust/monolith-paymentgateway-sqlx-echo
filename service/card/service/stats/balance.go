package cardstatsservice

import (
	"context"

	cardstatsmencache "github.com/MamangRust/monolith-payment-gateway-card/redis/stats"
	repository "github.com/MamangRust/monolith-payment-gateway-card/repository/stats"
	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	sharederrorhandler "github.com/MamangRust/monolith-payment-gateway-shared/errorhandler"
	card_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/card_errors/service"
	"go.opentelemetry.io/otel/attribute"
	"go.uber.org/zap"

	"github.com/MamangRust/monolith-payment-gateway-shared/observability"
)

type cardStatsBalanceService struct {
	cache cardstatsmencache.CardStatsBalanceCache

	repository repository.CardStatsBalanceRepository

	logger logger.LoggerInterface

	observability observability.TraceLoggerObservability
}

type cardStatsBalanceServiceDeps struct {
	Cache cardstatsmencache.CardStatsBalanceCache

	Repository repository.CardStatsBalanceRepository

	Logger logger.LoggerInterface

	Observability observability.TraceLoggerObservability
}

func NewCardStatsBalanceService(params *cardStatsBalanceServiceDeps) CardStatsBalanceService {
	return &cardStatsBalanceService{
		cache:         params.Cache,
		repository:    params.Repository,
		logger:        params.Logger,
		observability: params.Observability,
	}
}

func (s *cardStatsBalanceService) FindMonthlyBalance(ctx context.Context, year int) ([]*db.GetMonthlyBalancesRow, error) {
	const method = "FindMonthlyBalance"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", year))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetMonthlyBalanceCache(ctx, year); found {
		logSuccess("Monthly balance cache hit", zap.Int("year", year))
		return data, nil
	}

	res, err := s.repository.GetMonthlyBalance(ctx, year)
	if err != nil {
		status = "error"
		return sharederrorhandler.HandleError[[]*db.GetMonthlyBalancesRow](
			s.logger,
			card_errors.ErrFailedFindMonthlyBalance,
			method,
			span,
			zap.Int("year", year),
		)
	}

	s.cache.SetMonthlyBalanceCache(ctx, year, res)

	logSuccess("Monthly balance retrieved successfully",
		zap.Int("year", year),
		zap.Int("result_count", len(res)),
	)

	return res, nil
}

func (s *cardStatsBalanceService) FindYearlyBalance(ctx context.Context, year int) ([]*db.GetYearlyBalancesRow, error) {
	const method = "FindYearlyBalance"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", year))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetYearlyBalanceCache(ctx, year); found {
		logSuccess("Yearly balance cache hit", zap.Int("year", year))
		return data, nil
	}

	res, err := s.repository.GetYearlyBalance(ctx, year)
	if err != nil {
		status = "error"
		return sharederrorhandler.HandleError[[]*db.GetYearlyBalancesRow](
			s.logger,
			card_errors.ErrFailedFindYearlyBalance,
			method,
			span,
			zap.Int("year", year),
		)
	}

	s.cache.SetYearlyBalanceCache(ctx, year, res)

	logSuccess("Yearly balance retrieved successfully",
		zap.Int("year", year),
		zap.Int("result_count", len(res)),
	)

	return res, nil
}
