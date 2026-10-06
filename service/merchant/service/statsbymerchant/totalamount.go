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

type MerchantStatsByMerchantTotalAmountDeps struct {
	Cache mencache.MerchantStatsTotalAmountByMerchantCache

	Repository repository.MerchantStatsTotalAmountByMerchantRepository

	Logger logger.LoggerInterface

	Observability observability.TraceLoggerObservability
}

type merchantStatsByMerchantTotalAmountService struct {
	cache mencache.MerchantStatsTotalAmountByMerchantCache

	repository repository.MerchantStatsTotalAmountByMerchantRepository

	logger logger.LoggerInterface

	observability observability.TraceLoggerObservability
}


func NewMerchantStatsByMerchantTotalAmountService(params *MerchantStatsByMerchantTotalAmountDeps) *merchantStatsByMerchantTotalAmountService {
	return &merchantStatsByMerchantTotalAmountService{
		cache:         params.Cache,
		repository:    params.Repository,
		logger:        params.Logger,
		observability: params.Observability,
	}
}

func (s *merchantStatsByMerchantTotalAmountService) FindMonthlyTotalAmountByMerchants(ctx context.Context, req *requests.MonthYearTotalAmountMerchant) ([]*db.GetMonthlyTotalAmountByMerchantRow, error) {
	const method = "FindMonthlyTotalAmountByMerchants"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", req.Year),
		attribute.Int("merchant_id", req.MerchantID))

	defer func() {
		end(status)
	}()

	if cachedMerchant, found := s.cache.GetMonthlyTotalAmountByMerchantsCache(ctx, req); found {
		logSuccess("Successfully fetched merchant from cache", zap.Int("merchant.id", req.MerchantID), zap.Int("year", req.Year))
		return cachedMerchant, nil
	}

	dbRows, err := s.repository.GetMonthlyTotalAmountByMerchants(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetMonthlyTotalAmountByMerchantRow](
			s.logger,
			merchant_errors.ErrFailedFindMonthlyTotalAmountByMerchants,
			method,
			span,

			zap.Int("merchant_id", req.MerchantID),
			zap.Int("year", req.Year),
		)
	}

	s.cache.SetMonthlyTotalAmountByMerchantsCache(ctx, req, dbRows)

	logSuccess("Successfully found monthly total amount by merchant (from DB)", zap.Int("merchant_id", req.MerchantID), zap.Int("year", req.Year))
	return dbRows, nil
}

func (s *merchantStatsByMerchantTotalAmountService) FindYearlyTotalAmountByMerchants(ctx context.Context, req *requests.MonthYearTotalAmountMerchant) ([]*db.GetYearlyTotalAmountByMerchantRow, error) {
	const method = "FindYearlyTotalAmountByMerchants"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", req.Year),
		attribute.Int("merchant_id", req.MerchantID))

	defer func() {
		end(status)
	}()

	if cachedMerchant, found := s.cache.GetYearlyTotalAmountByMerchantsCache(ctx, req); found {
		logSuccess("Successfully fetched merchant from cache", zap.Int("merchant.id", req.MerchantID), zap.Int("year", req.Year))
		return cachedMerchant, nil
	}

	dbRows, err := s.repository.GetYearlyTotalAmountByMerchants(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetYearlyTotalAmountByMerchantRow](
			s.logger,
			merchant_errors.ErrFailedFindYearlyTotalAmountByMerchants,
			method,
			span,

			zap.Int("merchant_id", req.MerchantID),
			zap.Int("year", req.Year),
		)
	}

	s.cache.SetYearlyTotalAmountByMerchantsCache(ctx, req, dbRows)

	logSuccess("Successfully found yearly total amount by merchant (from DB)", zap.Int("merchant_id", req.MerchantID), zap.Int("year", req.Year))
	return dbRows, nil
}
