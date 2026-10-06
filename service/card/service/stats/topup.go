package cardstatsservice

import (
	"context"

	cardstatsmencache "github.com/MamangRust/monolith-payment-gateway-card/redis/stats"
	repository "github.com/MamangRust/monolith-payment-gateway-card/repository/stats"
	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	sharederrorhandler "github.com/MamangRust/monolith-payment-gateway-shared/errorhandler"
	card_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/card_errors/service"
	"github.com/MamangRust/monolith-payment-gateway-shared/observability"
	"go.opentelemetry.io/otel/attribute"
	"go.uber.org/zap"
)

type cardStatsTopupService struct {
	cache cardstatsmencache.CardStatsTopupCache

	repository repository.CardStatsTopupRepository

	logger logger.LoggerInterface

	observability observability.TraceLoggerObservability
}

type cardStatsTopupServiceDeps struct {
	Cache cardstatsmencache.CardStatsTopupCache

	Repository repository.CardStatsTopupRepository

	Logger logger.LoggerInterface

	Observability observability.TraceLoggerObservability
}

func NewCardStatsTopupService(params *cardStatsTopupServiceDeps) CardStatsTopupService {
	return &cardStatsTopupService{
		cache:         params.Cache,
		repository:    params.Repository,
		logger:        params.Logger,
		observability: params.Observability,
	}
}

func (s *cardStatsTopupService) FindMonthlyTopupAmount(ctx context.Context, year int) ([]*db.GetMonthlyTopupAmountRow, error) {
	const method = "FindMonthlyTopupAmount"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", year))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetMonthlyTopupCache(ctx, year); found {
		logSuccess("Monthly topup amount cache hit", zap.Int("year", year))
		return data, nil
	}

	res, err := s.repository.GetMonthlyTopupAmount(ctx, year)
	if err != nil {
		status = "error"
		return sharederrorhandler.HandleError[[]*db.GetMonthlyTopupAmountRow](
			s.logger,
			card_errors.ErrFailedFindMonthlyTopupAmount,
			method,
			span,
			zap.Int("year", year),
		)
	}

	s.cache.SetMonthlyTopupCache(ctx, year, res)

	logSuccess("Monthly topup amount retrieved successfully",
		zap.Int("year", year),
		zap.Int("result_count", len(res)),
	)

	return res, nil
}

func (s *cardStatsTopupService) FindYearlyTopupAmount(ctx context.Context, year int) ([]*db.GetYearlyTopupAmountRow, error) {
	const method = "FindYearlyTopupAmount"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", year))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetYearlyTopupCache(ctx, year); found {
		logSuccess("Yearly topup amount cache hit", zap.Int("year", year))
		return data, nil
	}

	res, err := s.repository.GetYearlyTopupAmount(ctx, year)
	if err != nil {
		status = "error"
		return sharederrorhandler.HandleError[[]*db.GetYearlyTopupAmountRow](
			s.logger,
			card_errors.ErrFailedFindYearlyTopupAmount,
			method,
			span,
			zap.Int("year", year),
		)
	}

	s.cache.SetYearlyTopupCache(ctx, year, res)

	logSuccess("Yearly topup amount retrieved successfully",
		zap.Int("year", year),
		zap.Int("result_count", len(res)),
	)

	return res, nil
}
