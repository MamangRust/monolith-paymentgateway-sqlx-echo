package topupstatsservice

import (
	"context"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/errorhandler"
	topup_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/topup_errors/service"
	"github.com/MamangRust/monolith-payment-gateway-shared/observability"
	cache "github.com/MamangRust/monolith-payment-gateway-topup/redis/stats"
	repository "github.com/MamangRust/monolith-payment-gateway-topup/repository/stats"
	"go.opentelemetry.io/otel/attribute"
	"go.uber.org/zap"
)

type topupStatsStatusDeps struct {
	Cache cache.TopupStatsStatusCache

	Repository repository.TopupStatsStatusRepository

	Logger logger.LoggerInterface

	Observability observability.TraceLoggerObservability
}

type topupStatsStatusService struct {
	cache cache.TopupStatsStatusCache

	repository repository.TopupStatsStatusRepository

	logger logger.LoggerInterface

	observability observability.TraceLoggerObservability
}

func NewTopupStatsStatusService(params *topupStatsStatusDeps) TopupStatsStatusService {

	return &topupStatsStatusService{
		cache:         params.Cache,
		repository:    params.Repository,
		logger:        params.Logger,
		observability: params.Observability,
	}
}

func (s *topupStatsStatusService) FindMonthTopupStatusSuccess(ctx context.Context, req *requests.MonthTopupStatus) ([]*db.GetMonthTopupStatusSuccessRow, error) {
	const method = "FindMonthTopupStatusSuccess"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", req.Year),
		attribute.Int("month", req.Month))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetMonthTopupStatusSuccessCache(ctx, req); found {
		logSuccess("Successfully fetched monthly topup status success from cache", zap.Int("year", req.Year), zap.Int("month", req.Month))
		return data, nil
	}

	s.logger.Debug("Cache miss for monthly topup status success, fetching from DB", zap.Int("year", req.Year), zap.Int("month", req.Month))

	dbRows, err := s.repository.GetMonthTopupStatusSuccess(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetMonthTopupStatusSuccessRow](
			s.logger,
			topup_errors.ErrFailedFindMonthTopupStatusSuccess,
			method,
			span,

			zap.Int("year", req.Year),
			zap.Int("month", req.Month),
		)
	}

	s.cache.SetMonthTopupStatusSuccessCache(ctx, req, dbRows)

	logSuccess("Successfully fetched monthly topup status success (from DB)", zap.Int("year", req.Year), zap.Int("month", req.Month))

	return dbRows, nil
}

func (s *topupStatsStatusService) FindYearlyTopupStatusSuccess(ctx context.Context, year int) ([]*db.GetYearlyTopupStatusSuccessRow, error) {
	const method = "FindYearlyTopupStatusSuccess"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", year))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetYearlyTopupStatusSuccessCache(ctx, year); found {
		logSuccess("Successfully fetched yearly topup status success from cache", zap.Int("year", year))
		return data, nil
	}

	dbRows, err := s.repository.GetYearlyTopupStatusSuccess(ctx, year)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetYearlyTopupStatusSuccessRow](
			s.logger,
			topup_errors.ErrFailedFindYearlyTopupStatusSuccess,
			method,
			span,

			zap.Int("year", year),
		)
	}

	s.cache.SetYearlyTopupStatusSuccessCache(ctx, year, dbRows)

	logSuccess("Successfully fetched yearly topup status success (from DB)", zap.Int("year", year))

	return dbRows, nil
}

func (s *topupStatsStatusService) FindMonthTopupStatusFailed(ctx context.Context, req *requests.MonthTopupStatus) ([]*db.GetMonthTopupStatusFailedRow, error) {
	const method = "FindMonthTopupStatusFailed"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", req.Year),
		attribute.Int("month", req.Month))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetMonthTopupStatusFailedCache(ctx, req); found {
		logSuccess("Successfully fetched monthly topup status Failed from cache", zap.Int("year", req.Year), zap.Int("month", req.Month))
		return data, nil
	}

	dbRows, err := s.repository.GetMonthTopupStatusFailed(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetMonthTopupStatusFailedRow](
			s.logger,
			topup_errors.ErrFailedFindMonthTopupStatusFailed,
			method,
			span,

			zap.Int("year", req.Year),
			zap.Int("month", req.Month),
		)
	}

	s.cache.SetMonthTopupStatusFailedCache(ctx, req, dbRows)

	logSuccess("Successfully fetched monthly topup status failed (from DB)", zap.Int("year", req.Year), zap.Int("month", req.Month))

	return dbRows, nil
}

func (s *topupStatsStatusService) FindYearlyTopupStatusFailed(ctx context.Context, year int) ([]*db.GetYearlyTopupStatusFailedRow, error) {
	const method = "FindYearlyTopupStatusFailed"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", year))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetYearlyTopupStatusFailedCache(ctx, year); found {
		logSuccess("Successfully fetched yearly topup status Failed from cache", zap.Int("year", year))
		return data, nil
	}

	dbRows, err := s.repository.GetYearlyTopupStatusFailed(ctx, year)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetYearlyTopupStatusFailedRow](
			s.logger,
			topup_errors.ErrFailedFindYearlyTopupStatusFailed,
			method,
			span,

			zap.Int("year", year),
		)
	}

	s.cache.SetYearlyTopupStatusFailedCache(ctx, year, dbRows)

	logSuccess("Successfully fetched yearly topup status failed (from DB)", zap.Int("year", year))

	return dbRows, nil
}
