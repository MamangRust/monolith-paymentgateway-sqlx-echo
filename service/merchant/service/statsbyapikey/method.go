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

type MerchantStatsMethodByApiKeyDeps struct {
	Cache mencache.MerchantStatsMethodByApiKeyCache

	Repository repository.MerchantStatsMethodByApiKeyRepository

	Logger logger.LoggerInterface

	Observability observability.TraceLoggerObservability
}

type merchantStatsMethodByApiKeyService struct {
	cache mencache.MerchantStatsMethodByApiKeyCache

	repository repository.MerchantStatsMethodByApiKeyRepository

	logger logger.LoggerInterface

	observability observability.TraceLoggerObservability
}

func NewMerchantStatsMethodByApiKeyService(params *MerchantStatsMethodByApiKeyDeps) MerchantStatsByApiKeyMethodService {
	return &merchantStatsMethodByApiKeyService{
		cache:         params.Cache,
		repository:    params.Repository,
		logger:        params.Logger,
		observability: params.Observability,
	}
}

func (s *merchantStatsMethodByApiKeyService) FindMonthlyPaymentMethodByApikey(ctx context.Context, req *requests.MonthYearPaymentMethodApiKey) ([]*db.GetMonthlyPaymentMethodByApikeyRow, error) {
	const method = "FindMonthlyPaymentMethodByApikeys"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.String("api_key", observability.MaskIdentifier(req.Apikey)),
		attribute.Int("year", req.Year))

	defer func() {
		end(status)
	}()

	if cachedMerchant, found := s.cache.GetMonthlyPaymentMethodByApikeysCache(ctx, req); found {
		logSuccess("Successfully fetched merchant from cache", zap.String("api_key", observability.MaskIdentifier(req.Apikey)), zap.Int("year", req.Year))

		return cachedMerchant, nil
	}

	dbRows, err := s.repository.GetMonthlyPaymentMethodByApikey(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetMonthlyPaymentMethodByApikeyRow](
			s.logger,
			merchant_errors.ErrFailedFindMonthlyPaymentMethodByApikeys,
			method,
			span,

			zap.String("api_key", observability.MaskIdentifier(req.Apikey)),
			zap.Int("year", req.Year),
		)
	}

	s.cache.SetMonthlyPaymentMethodByApikeysCache(ctx, req, dbRows)

	logSuccess("Successfully found monthly payment methods by API key (from DB)", zap.String("api_key", observability.MaskIdentifier(req.Apikey)), zap.Int("year", req.Year))
	return dbRows, nil
}

func (s *merchantStatsMethodByApiKeyService) FindYearlyPaymentMethodByApikey(ctx context.Context, req *requests.MonthYearPaymentMethodApiKey) ([]*db.GetYearlyPaymentMethodByApikeyRow, error) {
	const method = "FindYearlyPaymentMethodByApikeys"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.String("api_key", observability.MaskIdentifier(req.Apikey)),
		attribute.Int("year", req.Year))

	defer func() {
		end(status)
	}()

	if cachedMerchant, found := s.cache.GetYearlyPaymentMethodByApikeysCache(ctx, req); found {
		logSuccess("Successfully fetched merchant from cache", zap.String("api_key", observability.MaskIdentifier(req.Apikey)), zap.Int("year", req.Year))
		return cachedMerchant, nil
	}

	dbRows, err := s.repository.GetYearlyPaymentMethodByApikey(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetYearlyPaymentMethodByApikeyRow](
			s.logger,
			merchant_errors.ErrFailedFindYearlyPaymentMethodByApikeys,
			method,
			span,

			zap.String("api_key", observability.MaskIdentifier(req.Apikey)),
			zap.Int("year", req.Year),
		)
	}

	s.cache.SetYearlyPaymentMethodByApikeysCache(ctx, req, dbRows)

	logSuccess("Successfully found yearly payment methods by API key (from DB)", zap.String("api_key", observability.MaskIdentifier(req.Apikey)), zap.Int("year", req.Year))
	return dbRows, nil
}
