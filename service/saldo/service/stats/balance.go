package saldostatsservice

import (
	"context"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	mencache "github.com/MamangRust/monolith-payment-gateway-saldo/redis/stats"
	repository "github.com/MamangRust/monolith-payment-gateway-saldo/repository/stats"
	"github.com/MamangRust/monolith-payment-gateway-shared/errorhandler"
	saldo_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/saldo_errors/service"
	"github.com/MamangRust/monolith-payment-gateway-shared/observability"
	"go.opentelemetry.io/otel/attribute"
	"go.uber.org/zap"
)

type saldoStatsBalanceDeps struct {
	Cache mencache.SaldoStatsBalanceCache

	Repository repository.SaldoStatsBalanceRepository

	Logger logger.LoggerInterface

	Observability observability.TraceLoggerObservability
}

type saldoStatsBalanceService struct {
	cache mencache.SaldoStatsBalanceCache

	repository repository.SaldoStatsBalanceRepository

	logger logger.LoggerInterface

	observability observability.TraceLoggerObservability
}

func NewSaldoStatsBalanceService(params *saldoStatsBalanceDeps) SaldoStatsBalanceService {
	return &saldoStatsBalanceService{
		cache:         params.Cache,
		repository:    params.Repository,
		logger:        params.Logger,
		observability: params.Observability,
	}
}

func (s *saldoStatsBalanceService) FindMonthlySaldoBalances(ctx context.Context, year int) ([]*db.GetMonthlySaldoBalancesRow, error) {
	const method = "FindMonthlySaldoBalances"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", year))

	defer func() {
		end(status)
	}()

	if cache, found := s.cache.GetMonthlySaldoBalanceCache(ctx, year); found {
		logSuccess("Successfully fetched monthly saldo balances from cache", zap.Int("year", year))
		return cache, nil
	}

	s.logger.Debug("Cache miss for monthly saldo balances, fetching from DB", zap.Int("year", year))

	dbRows, err := s.repository.GetMonthlySaldoBalances(ctx, year)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetMonthlySaldoBalancesRow](
			s.logger,
			saldo_errors.ErrFailedFindMonthlySaldoBalances,
			method,
			span,

			zap.Int("year", year),
		)
	}

	s.logger.Debug("Setting cache for monthly saldo balances", zap.Int("year", year))

	s.cache.SetMonthlySaldoBalanceCache(ctx, year, dbRows)

	logSuccess("Successfully fetched monthly saldo balances (from DB)", zap.Int("year", year))

	return dbRows, nil
}

func (s *saldoStatsBalanceService) FindYearlySaldoBalances(ctx context.Context, year int) ([]*db.GetYearlySaldoBalancesRow, error) {
	const method = "FindYearlySaldoBalances"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", year))

	defer func() {
		end(status)
	}()

	if cache, found := s.cache.GetYearlySaldoBalanceCache(ctx, year); found {
		logSuccess("Successfully fetched yearly saldo balances from cache", zap.Int("year", year))
		return cache, nil
	}

	s.logger.Debug("Cache miss for yearly saldo balances, fetching from DB", zap.Int("year", year))

	dbRows, err := s.repository.GetYearlySaldoBalances(ctx, year)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetYearlySaldoBalancesRow](
			s.logger,
			saldo_errors.ErrFailedFindYearlySaldoBalances,
			method,
			span,

			zap.Int("year", year),
		)
	}

	s.logger.Debug("Setting cache for yearly saldo balances", zap.Int("year", year))
	s.cache.SetYearlySaldoBalanceCache(ctx, year, dbRows)

	logSuccess("Successfully fetched yearly saldo balances (from DB)", zap.Int("year", year))

	return dbRows, nil
}
