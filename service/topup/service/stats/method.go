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

type topupStatsMethodDeps struct {
	Cache mencache.TopupStatsMethodCache

	Repository repository.TopupStatsMethodRepository

	Logger logger.LoggerInterface

	Observability observability.TraceLoggerObservability
}

type topupStatsMethodService struct {
	cache mencache.TopupStatsMethodCache

	repository repository.TopupStatsMethodRepository

	logger logger.LoggerInterface

	observability observability.TraceLoggerObservability
}

func NewTopupStatsMethodService(params *topupStatsMethodDeps) TopupStatsMethodService {

	return &topupStatsMethodService{
		cache:         params.Cache,
		repository:    params.Repository,
		logger:        params.Logger,
		observability: params.Observability,
	}
}

func (s *topupStatsMethodService) FindMonthlyTopupMethods(ctx context.Context, year int) ([]*db.GetMonthlyTopupMethodsRow, error) {
	const method = "FindMonthlyTopupMethods"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", year))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetMonthlyTopupMethodsCache(ctx, year); found {
		logSuccess("Successfully fetched monthly topup methods from cache", zap.Int("year", year))
		return data, nil
	}

	dbRows, err := s.repository.GetMonthlyTopupMethods(ctx, year)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetMonthlyTopupMethodsRow](
			s.logger,
			topup_errors.ErrFailedFindMonthlyTopupMethods,
			method,
			span,

			zap.Int("year", year),
		)
	}

	s.cache.SetMonthlyTopupMethodsCache(ctx, year, dbRows)

	logSuccess("Successfully fetched monthly topup methods (from DB)", zap.Int("year", year))

	return dbRows, nil
}

func (s *topupStatsMethodService) FindYearlyTopupMethods(ctx context.Context, year int) ([]*db.GetYearlyTopupMethodsRow, error) {
	const method = "FindYearlyTopupMethods"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", year))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetYearlyTopupMethodsCache(ctx, year); found {
		logSuccess("Successfully fetched yearly topup methods from cache", zap.Int("year", year))
		return data, nil
	}

	s.logger.Debug("Cache miss for yearly topup methods, fetching from DB", zap.Int("year", year))

	dbRows, err := s.repository.GetYearlyTopupMethods(ctx, year)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetYearlyTopupMethodsRow](
			s.logger,
			topup_errors.ErrFailedFindYearlyTopupMethods,
			method,
			span,

			zap.Int("year", year),
		)
	}

	s.cache.SetYearlyTopupMethodsCache(ctx, year, dbRows)

	logSuccess("Successfully fetched yearly topup methods (from DB)", zap.Int("year", year))

	return dbRows, nil
}
