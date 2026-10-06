package transferstatsservice

import (
	"context"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/errorhandler"
	transfer_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/transfer_errors/service"
	"github.com/MamangRust/monolith-payment-gateway-shared/observability"
	mencache "github.com/MamangRust/monolith-payment-gateway-transfer/redis/stats"
	repository "github.com/MamangRust/monolith-payment-gateway-transfer/repository/stats"
	"go.opentelemetry.io/otel/attribute"
	"go.uber.org/zap"
)

type TransferStatsAmountDeps struct {
	Cache mencache.TransferStatsAmountCache

	Repository repository.TransferStatsAmountRepository

	Logger logger.LoggerInterface

	Observability observability.TraceLoggerObservability
}

type transferStatsAmountService struct {
	cache mencache.TransferStatsAmountCache

	repository repository.TransferStatsAmountRepository

	logger logger.LoggerInterface

	observability observability.TraceLoggerObservability
}

func NewTransferStatsAmountService(params *TransferStatsAmountDeps) TransferStatsAmountService {
	return &transferStatsAmountService{
		cache:         params.Cache,
		repository:    params.Repository,
		logger:        params.Logger,
		observability: params.Observability,
	}
}


func (s *transferStatsAmountService) FindMonthlyTransferAmounts(ctx context.Context, year int) ([]*db.GetMonthlyTransferAmountsRow, error) {
	const method = "FindMonthlyTransferAmounts"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", year))

	defer func() {
		end(status)
	}()

	if dbRows, found := s.cache.GetCachedMonthTransferAmounts(ctx, year); found {
		logSuccess("Successfully fetched monthly transfer amounts (from cache)", zap.Int("year", year))
		return dbRows, nil
	}

	dbRows, err := s.repository.GetMonthlyTransferAmounts(ctx, year)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetMonthlyTransferAmountsRow](
			s.logger,
			transfer_errors.ErrFailedFindMonthlyTransferAmounts,
			method,
			span,

			zap.Int("year", year),
		)
	}

	s.cache.SetCachedMonthTransferAmounts(ctx, year, dbRows)

	logSuccess("Successfully fetched monthly transfer amounts (from DB)", zap.Int("year", year))

	return dbRows, nil
}

func (s *transferStatsAmountService) FindYearlyTransferAmounts(ctx context.Context, year int) ([]*db.GetYearlyTransferAmountsRow, error) {
	const method = "FindYearlyTransferAmounts"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", year))

	defer func() {
		end(status)
	}()

	if dbRows, found := s.cache.GetCachedYearlyTransferAmounts(ctx, year); found {
		logSuccess("Successfully fetched yearly transfer amounts (from cache)", zap.Int("year", year))
		return dbRows, nil
	}

	dbRows, err := s.repository.GetYearlyTransferAmounts(ctx, year)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetYearlyTransferAmountsRow](
			s.logger,
			transfer_errors.ErrFailedFindYearlyTransferAmounts,
			method,
			span,

			zap.Int("year", year),
		)
	}

	s.cache.SetCachedYearlyTransferAmounts(ctx, year, dbRows)

	logSuccess("Successfully fetched yearly transfer amounts (from DB)", zap.Int("year", year))

	return dbRows, nil
}
