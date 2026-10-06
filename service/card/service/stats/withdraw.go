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

type cardStatsWithdrawService struct {
	cache cardstatsmencache.CardStatsWithdrawCache

	repository repository.CardStatsWithdrawRepository

	logger logger.LoggerInterface

	observability observability.TraceLoggerObservability
}

type cardStatsWithdrawServiceDeps struct {
	Cache cardstatsmencache.CardStatsWithdrawCache

	Repository repository.CardStatsWithdrawRepository

	Logger logger.LoggerInterface

	Observability observability.TraceLoggerObservability
}

func NewCardStatsWithdrawService(params *cardStatsWithdrawServiceDeps) CardStatsWithdrawService {

	return &cardStatsWithdrawService{
		cache:         params.Cache,
		repository:    params.Repository,
		logger:        params.Logger,
		observability: params.Observability,
	}
}

func (s *cardStatsWithdrawService) FindMonthlyWithdrawAmount(ctx context.Context, year int) ([]*db.GetMonthlyWithdrawAmountRow, error) {
	const method = "FindMonthlyWithdrawAmount"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", year))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetMonthlyWithdrawCache(ctx, year); found {
		logSuccess("Monthly withdraw amount cache hit", zap.Int("year", year))
		return data, nil
	}

	res, err := s.repository.GetMonthlyWithdrawAmount(ctx, year)
	if err != nil {
		status = "error"
		return sharederrorhandler.HandleError[[]*db.GetMonthlyWithdrawAmountRow](
			s.logger,
			card_errors.ErrFailedFindMonthlyWithdrawAmount,
			method,
			span,
			zap.Int("year", year),
		)
	}

	s.cache.SetMonthlyWithdrawCache(ctx, year, res)

	logSuccess("Monthly withdraw amount retrieved successfully",
		zap.Int("year", year),
		zap.Int("result_count", len(res)),
	)

	return res, nil
}

func (s *cardStatsWithdrawService) FindYearlyWithdrawAmount(ctx context.Context, year int) ([]*db.GetYearlyWithdrawAmountRow, error) {
	const method = "FindYearlyWithdrawAmount"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", year))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetYearlyWithdrawCache(ctx, year); found {
		logSuccess("Yearly withdraw amount cache hit", zap.Int("year", year))
		return data, nil
	}

	res, err := s.repository.GetYearlyWithdrawAmount(ctx, year)
	if err != nil {
		status = "error"
		return sharederrorhandler.HandleError[[]*db.GetYearlyWithdrawAmountRow](
			s.logger,
			card_errors.ErrFailedFindYearlyWithdrawAmount,
			method,
			span,
			zap.Int("year", year),
		)
	}

	s.cache.SetYearlyWithdrawCache(ctx, year, res)

	logSuccess("Yearly withdraw amount retrieved successfully",
		zap.Int("year", year),
		zap.Int("result_count", len(res)),
	)

	return res, nil
}
