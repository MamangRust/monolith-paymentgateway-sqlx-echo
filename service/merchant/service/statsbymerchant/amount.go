package merchantstatsbymerchantservice

import (
	"context"

	mencache "github.com/MamangRust/monolith-payment-gateway-merchant/redis/statsbymerchant"
	repository "github.com/MamangRust/monolith-payment-gateway-merchant/repository/statsbymerchant"
	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/errorhandler"
	merchant_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/merchant_errors/service"
	"github.com/MamangRust/monolith-payment-gateway-shared/observability"
	"go.opentelemetry.io/otel/attribute"
	"go.uber.org/zap"
)

type MerchantStatsByMerchantAmountServiceDeps struct {
	Cache mencache.MerchantStatsAmountByMerchantCache

	Repository repository.MerchantStatsAmountByMerchantRepository

	Logger logger.LoggerInterface

	Observability observability.TraceLoggerObservability
}

type merchantStatsByMerchantAmountService struct {
	cache mencache.MerchantStatsAmountByMerchantCache

	repository repository.MerchantStatsAmountByMerchantRepository

	logger logger.LoggerInterface

	observability observability.TraceLoggerObservability
}


func NewMerchantStatsByMerchantAmountService(params *MerchantStatsByMerchantAmountServiceDeps) MerchantStatsByMerchantAmountService {

	return &merchantStatsByMerchantAmountService{
		cache:         params.Cache,
		repository:    params.Repository,
		logger:        params.Logger,
		observability: params.Observability,
	}
}


func (s *merchantStatsByMerchantAmountService) FindMonthlyAmountByMerchants(ctx context.Context, req *requests.MonthYearAmountMerchant) ([]*db.GetMonthlyAmountByMerchantsRow, error) {
	const method = "FindMonthlyAmountByMerchants"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", req.Year),
		attribute.Int("merchant_id", req.MerchantID))

	defer func() {
		end(status)
	}()

	if cachedMerchant, found := s.cache.GetMonthlyAmountByMerchantsCache(ctx, req); found {
		logSuccess("Successfully fetched merchant from cache", zap.Int("merchant.id", req.MerchantID))

		return cachedMerchant, nil
	}

	dbRows, err := s.repository.GetMonthlyAmountByMerchants(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetMonthlyAmountByMerchantsRow](
			s.logger,
			merchant_errors.ErrFailedFindMonthlyAmountByMerchants,
			method,
			span,

			zap.Int("merchant_id", req.MerchantID),
			zap.Int("year", req.Year),
		)
	}

	s.cache.SetMonthlyAmountByMerchantsCache(ctx, req, dbRows)

	logSuccess("Successfully found monthly amount by merchant (from DB)", zap.Int("merchant_id", req.MerchantID), zap.Int("year", req.Year))
	return dbRows, nil
}

func (s *merchantStatsByMerchantAmountService) FindYearlyAmountByMerchants(ctx context.Context, req *requests.MonthYearAmountMerchant) ([]*db.GetYearlyAmountByMerchantsRow, error) {
	const method = "FindYearlyAmountByMerchants"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", req.Year),
		attribute.Int("merchant_id", req.MerchantID))

	defer func() {
		end(status)
	}()

	if cachedMerchant, found := s.cache.GetYearlyAmountByMerchantsCache(ctx, req); found {
		logSuccess("Successfully fetched merchant from cache", zap.Int("merchant.id", req.MerchantID))

		return cachedMerchant, nil
	}

	dbRows, err := s.repository.GetYearlyAmountByMerchants(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetYearlyAmountByMerchantsRow](
			s.logger,
			merchant_errors.ErrFailedFindYearlyAmountByMerchants,
			method,
			span,

			zap.Int("merchant_id", req.MerchantID),
			zap.Int("year", req.Year),
		)
	}

	s.cache.SetYearlyAmountByMerchantsCache(ctx, req, dbRows)

	logSuccess("Successfully found yearly amount by merchant (from DB)", zap.Int("merchant_id", req.MerchantID), zap.Int("year", req.Year))
	return dbRows, nil
}
