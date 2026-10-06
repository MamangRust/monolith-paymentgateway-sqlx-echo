package topupstatsbycardservice

import (
	"context"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/errorhandler"
	topup_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/topup_errors/service"
	"github.com/MamangRust/monolith-payment-gateway-shared/observability"
	cache "github.com/MamangRust/monolith-payment-gateway-topup/redis/statsbycard"
	repository "github.com/MamangRust/monolith-payment-gateway-topup/repository/statsbycard"
	"go.opentelemetry.io/otel/attribute"
	"go.uber.org/zap"
)

type topupStatsByCardStatusDeps struct {
	Cache cache.TopupStatsStatusByCardCache

	Repository repository.TopupStatsByCardStatusRepository

	Logger logger.LoggerInterface

	Observability observability.TraceLoggerObservability
}

type topupStatsByCardStatusService struct {
	cache cache.TopupStatsStatusByCardCache

	repository repository.TopupStatsByCardStatusRepository

	logger logger.LoggerInterface

	observability observability.TraceLoggerObservability
}

func NewTopupStatsByCardStatusService(params *topupStatsByCardStatusDeps) TopupStatsByCardStatusService {
	return &topupStatsByCardStatusService{
		cache:         params.Cache,
		repository:    params.Repository,
		logger:        params.Logger,
		observability: params.Observability,
	}
}

func (s *topupStatsByCardStatusService) FindMonthTopupStatusSuccessByCardNumber(ctx context.Context, req *requests.MonthTopupStatusCardNumber) ([]*db.GetMonthTopupStatusSuccessCardNumberRow, error) {
	const method = "FindMonthTopupStatusSuccessByCardNumber"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.String("card_number", observability.MaskIdentifier(req.CardNumber)),
		attribute.Int("year", req.Year),
		attribute.Int("month", req.Month))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetMonthTopupStatusSuccessByCardNumberCache(ctx, req); found {
		logSuccess("Successfully fetched monthly topup status success", zap.Int("year", req.Year), zap.Int("month", req.Month), zap.String("card_number", observability.MaskIdentifier(req.CardNumber)))
		return data, nil
	}

	dbRows, err := s.repository.GetMonthTopupStatusSuccessByCardNumber(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetMonthTopupStatusSuccessCardNumberRow](
			s.logger,
			topup_errors.ErrFailedFindMonthTopupStatusSuccessByCard,
			method,
			span,

			zap.String("card_number", observability.MaskIdentifier(req.CardNumber)),
			zap.Int("year", req.Year),
			zap.Int("month", req.Month),
		)
	}

	s.cache.SetMonthTopupStatusSuccessByCardNumberCache(ctx, req, dbRows)

	logSuccess("Successfully fetched monthly topup status success by card number (from DB)",
		zap.String("card_number", observability.MaskIdentifier(req.CardNumber)),
		zap.Int("year", req.Year),
		zap.Int("month", req.Month))

	return dbRows, nil
}

func (s *topupStatsByCardStatusService) FindYearlyTopupStatusSuccessByCardNumber(ctx context.Context, req *requests.YearTopupStatusCardNumber) ([]*db.GetYearlyTopupStatusSuccessCardNumberRow, error) {
	const method = "FindYearlyTopupStatusSuccessByCardNumber"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.String("card_number", observability.MaskIdentifier(req.CardNumber)),
		attribute.Int("year", req.Year))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetYearlyTopupStatusSuccessByCardNumberCache(ctx, req); found {
		logSuccess("Successfully fetched yearly topup status success", zap.Int("year", req.Year), zap.String("card_number", observability.MaskIdentifier(req.CardNumber)))
		return data, nil
	}

	dbRows, err := s.repository.GetYearlyTopupStatusSuccessByCardNumber(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetYearlyTopupStatusSuccessCardNumberRow](
			s.logger,
			topup_errors.ErrFailedFindYearlyTopupStatusSuccessByCard,
			method,
			span,

			zap.String("card_number", observability.MaskIdentifier(req.CardNumber)),
			zap.Int("year", req.Year),
		)
	}

	s.logger.Debug("Setting cache for yearly topup status success by card number",
		zap.String("card_number", observability.MaskIdentifier(req.CardNumber)),
		zap.Int("year", req.Year))

	s.cache.SetYearlyTopupStatusSuccessByCardNumberCache(ctx, req, dbRows)

	logSuccess("Successfully fetched yearly topup status success by card number (from DB)",
		zap.String("card_number", observability.MaskIdentifier(req.CardNumber)),
		zap.Int("year", req.Year))

	return dbRows, nil
}

func (s *topupStatsByCardStatusService) FindMonthTopupStatusFailedByCardNumber(ctx context.Context, req *requests.MonthTopupStatusCardNumber) ([]*db.GetMonthTopupStatusFailedCardNumberRow, error) {
	const method = "FindMonthTopupStatusFailedByCardNumber"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.String("card_number", observability.MaskIdentifier(req.CardNumber)),
		attribute.Int("year", req.Year),
		attribute.Int("month", req.Month))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetMonthTopupStatusFailedByCardNumberCache(ctx, req); found {
		logSuccess("Successfully fetched monthly topup status Failed", zap.Int("year", req.Year), zap.Int("month", req.Month), zap.String("card_number", observability.MaskIdentifier(req.CardNumber)))
		return data, nil
	}

	s.logger.Debug("Cache miss for monthly topup status failed by card number, fetching from DB",
		zap.String("card_number", observability.MaskIdentifier(req.CardNumber)),
		zap.Int("year", req.Year),
		zap.Int("month", req.Month))

	dbRows, err := s.repository.GetMonthTopupStatusFailedByCardNumber(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetMonthTopupStatusFailedCardNumberRow](
			s.logger,
			topup_errors.ErrFailedFindMonthTopupStatusFailedByCard,
			method,
			span,

			zap.String("card_number", observability.MaskIdentifier(req.CardNumber)),
			zap.Int("year", req.Year),
			zap.Int("month", req.Month),
		)
	}

	s.cache.SetMonthTopupStatusFailedByCardNumberCache(ctx, req, dbRows)

	logSuccess("Successfully fetched monthly topup status failed by card number (from DB)",
		zap.String("card_number", observability.MaskIdentifier(req.CardNumber)),
		zap.Int("year", req.Year),
		zap.Int("month", req.Month))

	return dbRows, nil
}

func (s *topupStatsByCardStatusService) FindYearlyTopupStatusFailedByCardNumber(ctx context.Context, req *requests.YearTopupStatusCardNumber) ([]*db.GetYearlyTopupStatusFailedCardNumberRow, error) {
	const method = "FindYearlyTopupStatusFailedByCardNumber"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.String("card_number", observability.MaskIdentifier(req.CardNumber)),
		attribute.Int("year", req.Year))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetYearlyTopupStatusFailedByCardNumberCache(ctx, req); found {
		logSuccess("Successfully fetched yearly topup status Failed", zap.Int("year", req.Year), zap.String("card_number", observability.MaskIdentifier(req.CardNumber)))
		return data, nil
	}

	s.logger.Debug("Cache miss for yearly topup status failed by card number, fetching from DB",
		zap.String("card_number", observability.MaskIdentifier(req.CardNumber)),
		zap.Int("year", req.Year))

	dbRows, err := s.repository.GetYearlyTopupStatusFailedByCardNumber(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetYearlyTopupStatusFailedCardNumberRow](
			s.logger,
			topup_errors.ErrFailedFindYearlyTopupStatusFailedByCard,
			method,
			span,

			zap.String("card_number", observability.MaskIdentifier(req.CardNumber)),
			zap.Int("year", req.Year),
		)
	}

	s.cache.SetYearlyTopupStatusFailedByCardNumberCache(ctx, req, dbRows)

	logSuccess("Successfully fetched yearly topup status failed by card number (from DB)",
		zap.String("card_number", observability.MaskIdentifier(req.CardNumber)),
		zap.Int("year", req.Year))

	return dbRows, nil
}
