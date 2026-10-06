package merchantstatsservice

import (
	"context"

	mencache "github.com/MamangRust/monolith-payment-gateway-merchant/redis/stats"
	repository "github.com/MamangRust/monolith-payment-gateway-merchant/repository/stats"
	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/errorhandler"
	merchant_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/merchant_errors/service"
	"github.com/MamangRust/monolith-payment-gateway-shared/observability"
	"go.opentelemetry.io/otel/attribute"
	"go.uber.org/zap"
)

type MerchantStatsTotalAmountDeps struct {
	Cache mencache.MerchantStatsTotalAmountCache

	Repository repository.MerchantStatsTotalAmountRepository

	Logger logger.LoggerInterface

	Observability observability.TraceLoggerObservability
}

type merchantStatsTotalAmountService struct {
	cache mencache.MerchantStatsTotalAmountCache

	repository repository.MerchantStatsTotalAmountRepository

	logger logger.LoggerInterface

	observability observability.TraceLoggerObservability
}

func NewMerchantStatsTotalAmountService(params *MerchantStatsTotalAmountDeps) MerchantStatsTotalAmountService {
	return &merchantStatsTotalAmountService{
		cache:         params.Cache,
		repository:    params.Repository,
		logger:        params.Logger,
		observability: params.Observability,
	}
}


func (s *merchantStatsTotalAmountService) FindMonthlyTotalAmountMerchant(ctx context.Context, year int) ([]*db.GetMonthlyTotalAmountMerchantRow, error) {
	const method = "FindMonthlyTotalAmountMerchant"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", year))

	defer func() {
		end(status)
	}()

	if cachedMerchant, found := s.cache.GetMonthlyTotalAmountMerchantCache(ctx, year); found {
		logSuccess("Successfully fetched merchant from cache", zap.Int("year", year))
		return cachedMerchant, nil
	}

	s.logger.Debug("Cache miss for monthly total amount for merchant, fetching from DB", zap.Int("year", year))

	dbRows, err := s.repository.GetMonthlyTotalAmountMerchant(ctx, year)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetMonthlyTotalAmountMerchantRow](
			s.logger,
			merchant_errors.ErrFailedFindMonthlyTotalAmountMerchant,
			method,
			span,

			zap.Int("year", year),
		)
	}

	s.cache.SetMonthlyTotalAmountMerchantCache(ctx, year, dbRows)

	logSuccess("Successfully found monthly total amount for merchant (from DB)", zap.Int("year", year))
	return dbRows, nil
}

func (s *merchantStatsTotalAmountService) FindYearlyTotalAmountMerchant(ctx context.Context, year int) ([]*db.GetYearlyTotalAmountMerchantRow, error) {
	const method = "FindYearlyTotalAmountMerchant"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", year))

	defer func() {
		end(status)
	}()

	if cachedMerchant, found := s.cache.GetYearlyTotalAmountMerchantCache(ctx, year); found {
		logSuccess("Successfully fetched merchant from cache", zap.Int("year", year))
		return cachedMerchant, nil
	}

	dbRows, err := s.repository.GetYearlyTotalAmountMerchant(ctx, year)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetYearlyTotalAmountMerchantRow](
			s.logger,
			merchant_errors.ErrFailedFindYearlyTotalAmountMerchant,
			method,
			span,

			zap.Int("year", year),
		)
	}

	s.cache.SetYearlyTotalAmountMerchantCache(ctx, year, dbRows)

	logSuccess("Successfully found yearly total amount for merchant (from DB)", zap.Int("year", year))
	return dbRows, nil
}
