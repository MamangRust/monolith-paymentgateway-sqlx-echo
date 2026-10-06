package transferstatsservice

import (
	"context"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/errorhandler"
	transfer_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/transfer_errors/service"
	"github.com/MamangRust/monolith-payment-gateway-shared/observability"
	mencache "github.com/MamangRust/monolith-payment-gateway-transfer/redis/stats"
	repository "github.com/MamangRust/monolith-payment-gateway-transfer/repository/stats"
	"go.opentelemetry.io/otel/attribute"
	"go.uber.org/zap"
)

type TransferStatsStatusDeps struct {
	Cache mencache.TransferStatsStatusCache

	Repository repository.TransferStatsStatusRepository

	Logger logger.LoggerInterface

	Observability observability.TraceLoggerObservability
}

type transferStatsStatusService struct {
	cache mencache.TransferStatsStatusCache

	repository repository.TransferStatsStatusRepository

	logger logger.LoggerInterface

	observability observability.TraceLoggerObservability
}

func NewTransferStatsStatusService(params *TransferStatsStatusDeps) TransferStatsStatusService {
	return &transferStatsStatusService{
		cache:         params.Cache,
		repository:    params.Repository,
		logger:        params.Logger,
		observability: params.Observability,
	}
}


func (s *transferStatsStatusService) FindMonthTransferStatusSuccess(ctx context.Context, req *requests.MonthStatusTransfer) ([]*db.GetMonthTransferStatusSuccessRow, error) {
	const method = "FindMonthTransferStatusSuccess"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", req.Year),
		attribute.Int("month", req.Month))

	defer func() {
		end(status)
	}()

	if dbRows, found := s.cache.GetCachedMonthTransferStatusSuccess(ctx, req); found {
		logSuccess("Successfully fetched monthly transfer status success (from cache)", zap.Int("year", req.Year), zap.Int("month", req.Month))
		return dbRows, nil
	}

	dbRows, err := s.repository.GetMonthTransferStatusSuccess(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetMonthTransferStatusSuccessRow](
			s.logger,
			transfer_errors.ErrFailedFindMonthTransferStatusSuccess,
			method,
			span,

			zap.Int("year", req.Year),
			zap.Int("month", req.Month),
		)
	}

	s.cache.SetCachedMonthTransferStatusSuccess(ctx, req, dbRows)

	logSuccess("Successfully fetched monthly transfer status success (from DB)", zap.Int("year", req.Year), zap.Int("month", req.Month))

	return dbRows, nil
}

func (s *transferStatsStatusService) FindYearlyTransferStatusSuccess(ctx context.Context, year int) ([]*db.GetYearlyTransferStatusSuccessRow, error) {
	const method = "FindYearlyTransferStatusSuccess"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", year))

	defer func() {
		end(status)
	}()

	if dbRows, found := s.cache.GetCachedYearlyTransferStatusSuccess(ctx, year); found {
		logSuccess("Successfully fetched yearly transfer status success (from cache)", zap.Int("year", year))
		return dbRows, nil
	}

	dbRows, err := s.repository.GetYearlyTransferStatusSuccess(ctx, year)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetYearlyTransferStatusSuccessRow](
			s.logger,
			transfer_errors.ErrFailedFindYearTransferStatusSuccess,
			method,
			span,

			zap.Int("year", year),
		)
	}

	s.cache.SetCachedYearlyTransferStatusSuccess(ctx, year, dbRows)

	logSuccess("Successfully fetched yearly transfer status success (from DB)", zap.Int("year", year))

	return dbRows, nil
}

func (s *transferStatsStatusService) FindMonthTransferStatusFailed(ctx context.Context, req *requests.MonthStatusTransfer) ([]*db.GetMonthTransferStatusFailedRow, error) {
	const method = "FindMonthTransferStatusFailed"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", req.Year),
		attribute.Int("month", req.Month))

	defer func() {
		end(status)
	}()

	if dbRows, found := s.cache.GetCachedMonthTransferStatusFailed(ctx, req); found {
		logSuccess("Successfully fetched monthly transfer status failed (from cache)", zap.Int("year", req.Year), zap.Int("month", req.Month))
		return dbRows, nil
	}

	dbRows, err := s.repository.GetMonthTransferStatusFailed(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetMonthTransferStatusFailedRow](
			s.logger,
			transfer_errors.ErrFailedFindMonthTransferStatusFailed,
			method,
			span,

			zap.Int("year", req.Year),
			zap.Int("month", req.Month),
		)
	}

	s.cache.SetCachedMonthTransferStatusFailed(ctx, req, dbRows)

	logSuccess("Successfully fetched monthly transfer status failed (from DB)", zap.Int("year", req.Year), zap.Int("month", req.Month))

	return dbRows, nil
}

func (s *transferStatsStatusService) FindYearlyTransferStatusFailed(ctx context.Context, year int) ([]*db.GetYearlyTransferStatusFailedRow, error) {
	const method = "FindYearlyTransferStatusFailed"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", year))

	defer func() {
		end(status)
	}()

	if dbRows, found := s.cache.GetCachedYearlyTransferStatusFailed(ctx, year); found {
		logSuccess("Successfully fetched yearly transfer status failed (from cache)", zap.Int("year", year))
		return dbRows, nil
	}

	dbRows, err := s.repository.GetYearlyTransferStatusFailed(ctx, year)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetYearlyTransferStatusFailedRow](
			s.logger,
			transfer_errors.ErrFailedFindYearTransferStatusFailed,
			method,
			span,

			zap.Int("year", year),
		)
	}

	s.cache.SetCachedYearlyTransferStatusFailed(ctx, year, dbRows)

	logSuccess("Successfully fetched yearly transfer status failed (from DB)", zap.Int("year", year))

	return dbRows, nil
}
