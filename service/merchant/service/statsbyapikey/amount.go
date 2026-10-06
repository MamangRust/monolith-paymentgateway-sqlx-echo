package merchantstatsbyapikeyservice

import (
	"context"

	mencache "github.com/MamangRust/monolith-payment-gateway-merchant/redis/statsbyapikey"
	repository "github.com/MamangRust/monolith-payment-gateway-merchant/repository/statsbyapikey"
	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/errorhandler"
	merchant_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/merchant_errors/service"
	"github.com/MamangRust/monolith-payment-gateway-shared/observability"
	"go.opentelemetry.io/otel/attribute"
	"go.uber.org/zap"
)

type MerchantStatsAmountByApiKeyDeps struct {
	Cache mencache.MerchantStatsAmountByApiKeyCache

	Repository repository.MerchantStatsAmountByApiKeyRepository

	Logger logger.LoggerInterface

	Observability observability.TraceLoggerObservability
}

type merchantStatsAmountByApiKeyService struct {
	cache mencache.MerchantStatsAmountByApiKeyCache

	repository repository.MerchantStatsAmountByApiKeyRepository

	logger logger.LoggerInterface

	observability observability.TraceLoggerObservability
}

func NewMerchantStatsAmountByApiKeyService(params *MerchantStatsAmountByApiKeyDeps) MerchantStatsByApiKeyAmountService {
	return &merchantStatsAmountByApiKeyService{
		cache:         params.Cache,
		repository:    params.Repository,
		logger:        params.Logger,
		observability: params.Observability,
	}
}

func (s *merchantStatsAmountByApiKeyService) FindMonthlyAmountByApikey(ctx context.Context, req *requests.MonthYearAmountApiKey) ([]*db.GetMonthlyAmountByApikeyRow, error) {
	const method = "FindMonthlyAmountByApikeys"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.String("api_key", observability.MaskIdentifier(req.Apikey)),
		attribute.Int("year", req.Year))

	defer func() {
		end(status)
	}()

	if cachedMerchant, found := s.cache.GetMonthlyAmountByApikeysCache(ctx, req); found {
		logSuccess("Successfully fetched merchant from cache", zap.String("api_key", observability.MaskIdentifier(req.Apikey)), zap.Int("year", req.Year))

		return cachedMerchant, nil
	}

	dbRows, err := s.repository.GetMonthlyAmountByApikey(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetMonthlyAmountByApikeyRow](
			s.logger,
			merchant_errors.ErrFailedFindMonthlyAmountByApikeys,
			method,
			span,

			zap.String("api_key", observability.MaskIdentifier(req.Apikey)),
			zap.Int("year", req.Year),
		)
	}

	s.cache.SetMonthlyAmountByApikeysCache(ctx, req, dbRows)

	logSuccess("Successfully found monthly amount by API key (from DB)", zap.String("api_key", observability.MaskIdentifier(req.Apikey)), zap.Int("year", req.Year))
	return dbRows, nil
}

func (s *merchantStatsAmountByApiKeyService) FindYearlyAmountByApikey(ctx context.Context, req *requests.MonthYearAmountApiKey) ([]*db.GetYearlyAmountByApikeyRow, error) {
	const method = "FindYearlyAmountByApikeys"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.String("api_key", observability.MaskIdentifier(req.Apikey)),
		attribute.Int("year", req.Year))

	defer func() {
		end(status)
	}()

	if cachedMerchant, found := s.cache.GetYearlyAmountByApikeysCache(ctx, req); found {
		logSuccess("Successfully fetched merchant from cache", zap.String("api_key", observability.MaskIdentifier(req.Apikey)), zap.Int("year", req.Year))
		return cachedMerchant, nil
	}

	dbRows, err := s.repository.GetYearlyAmountByApikey(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetYearlyAmountByApikeyRow](
			s.logger,
			merchant_errors.ErrFailedFindYearlyAmountByApikeys,
			method,
			span,

			zap.String("api_key", observability.MaskIdentifier(req.Apikey)),
			zap.Int("year", req.Year),
		)
	}

	s.cache.SetYearlyAmountByApikeysCache(ctx, req, dbRows)

	logSuccess("Successfully found yearly amount by API key (from DB)", zap.String("api_key", observability.MaskIdentifier(req.Apikey)), zap.Int("year", req.Year))
	return dbRows, nil
}
