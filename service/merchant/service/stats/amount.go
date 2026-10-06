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

type MerchantStatsAmountDeps struct {
	Cache mencache.MerchantStatsAmountCache

	Repository repository.MerchantStatsAmountRepository

	Logger logger.LoggerInterface

	Observability observability.TraceLoggerObservability
}

type merchantStatsAmountService struct {
	cache mencache.MerchantStatsAmountCache

	repository repository.MerchantStatsAmountRepository

	logger logger.LoggerInterface

	observability observability.TraceLoggerObservability
}

func NewMerchantStatsAmountService(params *MerchantStatsAmountDeps) MerchantStatsAmountService {
	return &merchantStatsAmountService{
		cache:         params.Cache,
		repository:    params.Repository,
		logger:        params.Logger,
		observability: params.Observability,
	}
}


func (s *merchantStatsAmountService) FindMonthlyAmountMerchant(ctx context.Context, year int) ([]*db.GetMonthlyAmountMerchantRow, error) {
	const method = "FindMonthlyAmountMerchant"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", year))

	defer func() {
		end(status)
	}()

	if cachedMerchant, found := s.cache.GetMonthlyAmountMerchantCache(ctx, year); found {
		logSuccess("Successfully fetched merchant from cache", zap.Int("year", year))
		return cachedMerchant, nil
	}

	s.logger.Debug("Cache miss for monthly amount for merchant, fetching from DB", zap.Int("year", year))

	dbRows, err := s.repository.GetMonthlyAmountMerchant(ctx, year)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetMonthlyAmountMerchantRow](
			s.logger,
			merchant_errors.ErrFailedFindMonthlyAmountMerchant,
			method,
			span,

			zap.Int("year", year),
		)
	}

	s.cache.SetMonthlyAmountMerchantCache(ctx, year, dbRows)

	logSuccess("Successfully found monthly amount for merchant (from DB)", zap.Int("year", year))
	return dbRows, nil
}

func (s *merchantStatsAmountService) FindYearlyAmountMerchant(ctx context.Context, year int) ([]*db.GetYearlyAmountMerchantRow, error) {
	const method = "FindYearlyAmountMerchant"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", year))

	defer func() {
		end(status)
	}()

	if cachedMerchant, found := s.cache.GetYearlyAmountMerchantCache(ctx, year); found {
		logSuccess("Successfully fetched merchant from cache", zap.Int("year", year))
		return cachedMerchant, nil
	}

	dbRows, err := s.repository.GetYearlyAmountMerchant(ctx, year)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetYearlyAmountMerchantRow](
			s.logger,
			merchant_errors.ErrFailedFindYearlyAmountMerchant,
			method,
			span,

			zap.Int("year", year),
		)
	}

	s.cache.SetYearlyAmountMerchantCache(ctx, year, dbRows)

	logSuccess("Successfully found yearly amount for merchant (from DB)", zap.Int("year", year))
	return dbRows, nil
}
