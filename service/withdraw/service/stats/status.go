package withdrawstatsservice

import (
	"context"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/errorhandler"
	withdraw_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/withdraw_errors/service"
	"github.com/MamangRust/monolith-payment-gateway-shared/observability"
	mencache "github.com/MamangRust/monolith-payment-gateway-withdraw/redis/stats"
	repository "github.com/MamangRust/monolith-payment-gateway-withdraw/repository/stats"
	"go.opentelemetry.io/otel/attribute"
	"go.uber.org/zap"
)

type WithdrawStatsStatusDeps struct {
	Cache mencache.WithdrawStatsCache

	Repository repository.WithdrawStatsStatusRepository

	Logger logger.LoggerInterface

	Observability observability.TraceLoggerObservability
}

type withdrawStatsStatusService struct {
	cache mencache.WithdrawStatsCache

	repository repository.WithdrawStatsStatusRepository

	logger logger.LoggerInterface

	observability observability.TraceLoggerObservability
}


func NewWithdrawStatsStatusService(params *WithdrawStatsStatusDeps) WithdrawStatsStatusService {
	return &withdrawStatsStatusService{
		cache:         params.Cache,
		repository:    params.Repository,
		logger:        params.Logger,
		observability: params.Observability,
	}
}

func (s *withdrawStatsStatusService) FindMonthWithdrawStatusSuccess(ctx context.Context, req *requests.MonthStatusWithdraw) ([]*db.GetMonthWithdrawStatusSuccessRow, error) {
	const method = "FindMonthWithdrawStatusSuccess"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", req.Year),
		attribute.Int("month", req.Month))

	defer func() {
		end(status)
	}()

	s.logger.Debug("Checking cache for monthly withdraw status success", zap.Int("year", req.Year), zap.Int("month", req.Month))

	if dbRows, found := s.cache.GetCachedMonthWithdrawStatusSuccessCache(ctx, req); found {
		s.logger.Info("Cache hit for monthly withdraw status success", zap.Int("year", req.Year), zap.Int("month", req.Month))
		status = "ok"
		logSuccess("Successfully fetched monthly withdraw status success (from cache)", zap.Int("year", req.Year), zap.Int("month", req.Month))
		return dbRows, nil
	}

	dbRows, err := s.repository.GetMonthWithdrawStatusSuccess(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetMonthWithdrawStatusSuccessRow](
			s.logger,
			withdraw_errors.ErrFailedFindMonthWithdrawStatusSuccess,
			method,
			span,

			zap.Int("year", req.Year),
			zap.Int("month", req.Month),
		)
	}

	s.cache.SetCachedMonthWithdrawStatusSuccessCache(ctx, req, dbRows)

	logSuccess("Successfully fetched monthly withdraw status success (from DB)", zap.Int("year", req.Year), zap.Int("month", req.Month))

	return dbRows, nil
}

func (s *withdrawStatsStatusService) FindYearlyWithdrawStatusSuccess(ctx context.Context, year int) ([]*db.GetYearlyWithdrawStatusSuccessRow, error) {
	const method = "FindYearlyWithdrawStatusSuccess"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", year))

	defer func() {
		end(status)
	}()

	if dbRows, found := s.cache.GetCachedYearlyWithdrawStatusSuccessCache(ctx, year); found {
		logSuccess("Successfully fetched yearly withdraw status success (from cache)", zap.Int("year", year))
		return dbRows, nil
	}

	dbRows, err := s.repository.GetYearlyWithdrawStatusSuccess(ctx, year)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetYearlyWithdrawStatusSuccessRow](
			s.logger,
			withdraw_errors.ErrFailedFindYearWithdrawStatusSuccess,
			method,
			span,

			zap.Int("year", year),
		)
	}

	s.cache.SetCachedYearlyWithdrawStatusSuccessCache(ctx, year, dbRows)

	logSuccess("Successfully fetched yearly withdraw status success (from DB)", zap.Int("year", year))

	return dbRows, nil
}

func (s *withdrawStatsStatusService) FindMonthWithdrawStatusFailed(ctx context.Context, req *requests.MonthStatusWithdraw) ([]*db.GetMonthWithdrawStatusFailedRow, error) {
	const method = "FindMonthWithdrawStatusFailed"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", req.Year),
		attribute.Int("month", req.Month))

	defer func() {
		end(status)
	}()

	if dbRows, found := s.cache.GetCachedMonthWithdrawStatusFailedCache(ctx, req); found {
		logSuccess("Successfully fetched monthly withdraw status failed (from cache)", zap.Int("year", req.Year), zap.Int("month", req.Month))
		return dbRows, nil
	}

	dbRows, err := s.repository.GetMonthWithdrawStatusFailed(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetMonthWithdrawStatusFailedRow](
			s.logger,
			withdraw_errors.ErrFailedFindMonthWithdrawStatusFailed,
			method,
			span,

			zap.Int("year", req.Year),
			zap.Int("month", req.Month),
		)
	}

	s.cache.SetCachedMonthWithdrawStatusFailedCache(ctx, req, dbRows)

	logSuccess("Successfully fetched monthly withdraw status failed (from DB)", zap.Int("year", req.Year), zap.Int("month", req.Month))

	return dbRows, nil
}

func (s *withdrawStatsStatusService) FindYearlyWithdrawStatusFailed(ctx context.Context, year int) ([]*db.GetYearlyWithdrawStatusFailedRow, error) {
	const method = "FindYearlyWithdrawStatusFailed"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", year))

	defer func() {
		end(status)
	}()

	if dbRows, found := s.cache.GetCachedYearlyWithdrawStatusFailedCache(ctx, year); found {
		logSuccess("Successfully fetched yearly withdraw status failed (from cache)", zap.Int("year", year))
		return dbRows, nil
	}

	dbRows, err := s.repository.GetYearlyWithdrawStatusFailed(ctx, year)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetYearlyWithdrawStatusFailedRow](
			s.logger,
			withdraw_errors.ErrFailedFindYearWithdrawStatusFailed,
			method,
			span,

			zap.Int("year", year),
		)
	}

	s.cache.SetCachedYearlyWithdrawStatusFailedCache(ctx, year, dbRows)

	logSuccess("Successfully fetched yearly withdraw status failed (from DB)", zap.Int("year", year))

	return dbRows, nil
}
