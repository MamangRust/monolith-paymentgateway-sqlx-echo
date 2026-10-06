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

type MerchantStatsTotalAmountByApiKeyDeps struct {
	Cache mencache.MerchantStatsTotalAmountByApiKeyCache

	Repository repository.MerchantStatsTotalAmountByApiKeyRepository

	Logger logger.LoggerInterface

	Observability observability.TraceLoggerObservability
}

type merchantStatsTotalAmountByApiKeyService struct {
	cache mencache.MerchantStatsTotalAmountByApiKeyCache

	repository repository.MerchantStatsTotalAmountByApiKeyRepository

	logger logger.LoggerInterface

	observability observability.TraceLoggerObservability
}

func NewMerchantStatsTotalAmountByApiKeyService(params *MerchantStatsTotalAmountByApiKeyDeps) MerchantStatsByApiKeyTotalAmountService {
	return &merchantStatsTotalAmountByApiKeyService{
		cache:         params.Cache,
		repository:    params.Repository,
		logger:        params.Logger,
		observability: params.Observability,
	}
}

func (s *merchantStatsTotalAmountByApiKeyService) FindMonthlyTotalAmountByApikey(ctx context.Context, req *requests.MonthYearTotalAmountApiKey) ([]*db.GetMonthlyTotalAmountByApikeyRow, error) {
	const method = "FindMonthlyTotalAmountByApikeys"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.String("api_key", observability.MaskIdentifier(req.Apikey)),
		attribute.Int("year", req.Year))

	defer func() {
		end(status)
	}()

	if cachedMerchant, found := s.cache.GetMonthlyTotalAmountByApikeysCache(ctx, req); found {
		logSuccess("Successfully fetched merchant from cache", zap.String("api_key", observability.MaskIdentifier(req.Apikey)), zap.Int("year", req.Year))
		return cachedMerchant, nil
	}

	dbRows, err := s.repository.GetMonthlyTotalAmountByApikey(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetMonthlyTotalAmountByApikeyRow](
			s.logger,
			merchant_errors.ErrFailedFindMonthlyTotalAmountByApikeys,
			method,
			span,

			zap.String("api_key", observability.MaskIdentifier(req.Apikey)),
			zap.Int("year", req.Year),
		)
	}

	s.cache.SetMonthlyTotalAmountByApikeysCache(ctx, req, dbRows)

	logSuccess("Successfully found monthly total amount by API key (from DB)", zap.String("api_key", observability.MaskIdentifier(req.Apikey)), zap.Int("year", req.Year))
	return dbRows, nil
}

func (s *merchantStatsTotalAmountByApiKeyService) FindYearlyTotalAmountByApikey(ctx context.Context, req *requests.MonthYearTotalAmountApiKey) ([]*db.GetYearlyTotalAmountByApikeyRow, error) {
	const method = "FindYearlyTotalAmountByApikeys"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.String("api_key", observability.MaskIdentifier(req.Apikey)),
		attribute.Int("year", req.Year))

	defer func() {
		end(status)
	}()

	if cachedMerchant, found := s.cache.GetYearlyTotalAmountByApikeysCache(ctx, req); found {
		logSuccess("Successfully fetched merchant from cache", zap.String("api_key", observability.MaskIdentifier(req.Apikey)), zap.Int("year", req.Year))
		return cachedMerchant, nil
	}

	dbRows, err := s.repository.GetYearlyTotalAmountByApikey(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetYearlyTotalAmountByApikeyRow](
			s.logger,
			merchant_errors.ErrFailedFindYearlyTotalAmountByApikeys,
			method,
			span,

			zap.String("api_key", observability.MaskIdentifier(req.Apikey)),
			zap.Int("year", req.Year),
		)
	}

	s.cache.SetYearlyTotalAmountByApikeysCache(ctx, req, dbRows)

	logSuccess("Successfully found yearly total amount by API key (from DB)", zap.String("api_key", observability.MaskIdentifier(req.Apikey)), zap.Int("year", req.Year))
	return dbRows, nil
}
