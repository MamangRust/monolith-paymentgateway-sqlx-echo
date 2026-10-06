package saldostatsservice

import (
	"context"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	mencache "github.com/MamangRust/monolith-payment-gateway-saldo/redis/stats"
	repository "github.com/MamangRust/monolith-payment-gateway-saldo/repository/stats"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/errorhandler"
	saldo_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/saldo_errors/service"
	"github.com/MamangRust/monolith-payment-gateway-shared/observability"
	"go.opentelemetry.io/otel/attribute"
	"go.uber.org/zap"
)

type saldoStatsTotalBalanceDeps struct {
	Cache mencache.SaldoStatsTotalCache

	Repository repository.SaldoStatsTotalSaldoRepository

	Logger logger.LoggerInterface

	Observability observability.TraceLoggerObservability
}

type saldoStatsTotalBalanceService struct {
	cache mencache.SaldoStatsTotalCache

	repository repository.SaldoStatsTotalSaldoRepository

	logger logger.LoggerInterface

	observability observability.TraceLoggerObservability
}

func NewSaldoStatsTotalBalanceService(params *saldoStatsTotalBalanceDeps) SaldoStatsTotalBalanceService {

	return &saldoStatsTotalBalanceService{
		cache:         params.Cache,
		repository:    params.Repository,
		logger:        params.Logger,
		observability: params.Observability,
	}
}

func (s *saldoStatsTotalBalanceService) FindMonthlyTotalSaldoBalance(ctx context.Context, req *requests.MonthTotalSaldoBalance) ([]*db.GetMonthlyTotalSaldoBalanceRow, error) {
	const method = "FindMonthlyTotalSaldoBalance"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", req.Year),
		attribute.Int("month", req.Month))

	defer func() {
		end(status)
	}()

	if cache, found := s.cache.GetMonthlyTotalSaldoBalanceCache(ctx, req); found {
		logSuccess("Successfully fetched monthly total saldo balance from cache", zap.Int("year", req.Year), zap.Int("month", req.Month))
		return cache, nil
	}

	dbRows, err := s.repository.GetMonthlyTotalSaldoBalance(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetMonthlyTotalSaldoBalanceRow](
			s.logger,
			saldo_errors.ErrFailedFindMonthlyTotalSaldoBalance,
			method,
			span,

			zap.Int("year", req.Year),
			zap.Int("month", req.Month),
		)
	}

	s.logger.Debug("Setting cache for monthly total saldo balance", zap.Int("year", req.Year), zap.Int("month", req.Month))
	s.cache.SetMonthlyTotalSaldoCache(ctx, req, dbRows)

	logSuccess("Successfully fetched monthly total saldo balance (from DB)", zap.Int("year", req.Year), zap.Int("month", req.Month))

	return dbRows, nil
}

func (s *saldoStatsTotalBalanceService) FindYearTotalSaldoBalance(ctx context.Context, year int) ([]*db.GetYearlyTotalSaldoBalancesRow, error) {
	const method = "FindYearTotalSaldoBalance"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", year))

	defer func() {
		end(status)
	}()

	if cache, found := s.cache.GetYearTotalSaldoBalanceCache(ctx, year); found {
		logSuccess("Successfully fetched yearly total saldo balance from cache", zap.Int("year", year))
		return cache, nil
	}

	s.logger.Debug("Cache miss for yearly total saldo balance, fetching from DB", zap.Int("year", year))

	dbRows, err := s.repository.GetYearTotalSaldoBalance(ctx, year)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetYearlyTotalSaldoBalancesRow](
			s.logger,
			saldo_errors.ErrFailedFindYearTotalSaldoBalance,
			method,
			span,

			zap.Int("year", year),
		)
	}

	s.logger.Debug("Setting cache for yearly total saldo balance", zap.Int("year", year))

	s.cache.SetYearTotalSaldoBalanceCache(ctx, year, dbRows)

	logSuccess("Successfully fetched yearly total saldo balance (from DB)", zap.Int("year", year))

	return dbRows, nil
}
