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

type MerchantStatsMethodDeps struct {
	Cache mencache.MerchantStatsMethodCache

	Repository repository.MerchantStatsMethodRepository

	Logger logger.LoggerInterface

	Observability observability.TraceLoggerObservability
}

type merchantStatsMethodService struct {
	cache mencache.MerchantStatsMethodCache

	repository repository.MerchantStatsMethodRepository

	logger logger.LoggerInterface

	observability observability.TraceLoggerObservability
}

func NewMerchantStatsMethodService(params *MerchantStatsMethodDeps) MerchantStatsMethodService {
	return &merchantStatsMethodService{
		cache:         params.Cache,
		repository:    params.Repository,
		logger:        params.Logger,
		observability: params.Observability,
	}
}


func (s *merchantStatsMethodService) FindMonthlyPaymentMethodsMerchant(ctx context.Context, year int) ([]*db.GetMonthlyPaymentMethodsMerchantRow, error) {
	const method = "FindMonthlyPaymentMethodsMerchant"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", year))

	defer func() {
		end(status)
	}()

	if cachedMerchant, found := s.cache.GetMonthlyPaymentMethodsMerchantCache(ctx, year); found {
		logSuccess("Successfully fetched merchant from cache", zap.Int("year", year))
		return cachedMerchant, nil
	}

	s.logger.Debug("Cache miss for monthly payment methods for merchant, fetching from DB", zap.Int("year", year))

	dbRows, err := s.repository.GetMonthlyPaymentMethodsMerchant(ctx, year)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetMonthlyPaymentMethodsMerchantRow](
			s.logger,
			merchant_errors.ErrFailedFindMonthlyPaymentMethodsMerchant,
			method,
			span,

			zap.Int("year", year),
		)
	}

	s.cache.SetMonthlyPaymentMethodsMerchantCache(ctx, year, dbRows)

	logSuccess("Successfully found monthly payment methods for merchant (from DB)", zap.Int("year", year))
	return dbRows, nil
}

func (s *merchantStatsMethodService) FindYearlyPaymentMethodMerchant(ctx context.Context, year int) ([]*db.GetYearlyPaymentMethodMerchantRow, error) {
	const method = "FindYearlyPaymentMethodMerchant"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", year))

	defer func() {
		end(status)
	}()

	if cachedMerchant, found := s.cache.GetYearlyPaymentMethodMerchantCache(ctx, year); found {
		logSuccess("Successfully fetched merchant from cache", zap.Int("year", year))
		return cachedMerchant, nil
	}

	s.logger.Debug("Cache miss for yearly payment methods for merchant, fetching from DB", zap.Int("year", year))

	dbRows, err := s.repository.GetYearlyPaymentMethodMerchant(ctx, year)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetYearlyPaymentMethodMerchantRow](
			s.logger,
			merchant_errors.ErrFailedFindYearlyPaymentMethodMerchant,
			method,
			span,

			zap.Int("year", year),
		)
	}

	s.cache.SetYearlyPaymentMethodMerchantCache(ctx, year, dbRows)

	logSuccess("Successfully found yearly payment methods for merchant (from DB)", zap.Int("year", year))
	return dbRows, nil
}
