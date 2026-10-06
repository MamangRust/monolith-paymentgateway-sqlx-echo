package service

import (
	"context"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	mencache "github.com/MamangRust/monolith-payment-gateway-saldo/redis"
	"github.com/MamangRust/monolith-payment-gateway-saldo/repository"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/errorhandler"
	saldo_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/saldo_errors/service"
	"github.com/MamangRust/monolith-payment-gateway-shared/observability"
	"go.opentelemetry.io/otel/attribute"
	"go.uber.org/zap"
)

type SaldoQueryParams struct {
	Cache         mencache.SaldoQueryCache
	Repository    repository.SaldoQueryRepository
	Logger        logger.LoggerInterface
	Observability observability.TraceLoggerObservability
}

type saldoQueryService struct {
	mencache             mencache.SaldoQueryCache
	saldoQueryRepository repository.SaldoQueryRepository
	logger               logger.LoggerInterface
	observability        observability.TraceLoggerObservability
}

func NewSaldoQueryService(
	params *SaldoQueryParams,
) SaldoQueryService {
	return &saldoQueryService{
		mencache:             params.Cache,
		saldoQueryRepository: params.Repository,
		logger:               params.Logger,
		observability:        params.Observability,
	}
}

func (s *saldoQueryService) FindAll(ctx context.Context, req *requests.FindAllSaldos) ([]*db.GetSaldosRow, *int, error) {
	const method = "FindAll"

	page, pageSize := s.normalizePagination(req.Page, req.PageSize)
	search := req.Search

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("page", page),
		attribute.Int("pageSize", pageSize),
		attribute.String("search", search))

	defer func() {
		end(status)
	}()

	res, err := s.saldoQueryRepository.FindAllSaldos(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandlerErrorPagination[[]*db.GetSaldosRow](
			s.logger,
			saldo_errors.ErrFailedFindAllSaldos,
			method,
			span,

			zap.Int("page", page),
			zap.Int("pageSize", pageSize),
			zap.String("search", search),
		)
	}

	var totalCount int

	if len(res) > 0 {
		totalCount = int(res[0].TotalCount)
	} else {
		totalCount = 0
	}

	logSuccess("Successfully fetched saldo",
		zap.Int("totalRecords", totalCount),
		zap.Int("page", req.Page),
		zap.Int("pageSize", req.PageSize))

	return res, &totalCount, nil
}

func (s *saldoQueryService) FindByActive(ctx context.Context, req *requests.FindAllSaldos) ([]*db.GetActiveSaldosRow, *int, error) {
	const method = "FindByActive"

	page, pageSize := s.normalizePagination(req.Page, req.PageSize)
	search := req.Search

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("page", page),
		attribute.Int("pageSize", pageSize),
		attribute.String("search", search))

	defer func() {
		end(status)
	}()

	res, err := s.saldoQueryRepository.FindByActive(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandlerErrorPagination[[]*db.GetActiveSaldosRow](
			s.logger,
			saldo_errors.ErrFailedFindActiveSaldos,
			method,
			span,

			zap.Int("page", page),
			zap.Int("page_size", pageSize),
			zap.String("search", search),
		)
	}

	var totalCount int

	if len(res) > 0 {
		totalCount = int(res[0].TotalCount)
	} else {
		totalCount = 0
	}

	logSuccess("Successfully fetched active saldo",
		zap.Int("totalRecords", totalCount),
		zap.Int("page", page),
		zap.Int("pageSize", pageSize))

	return res, &totalCount, nil
}

func (s *saldoQueryService) FindByTrashed(ctx context.Context, req *requests.FindAllSaldos) ([]*db.GetTrashedSaldosRow, *int, error) {
	const method = "FindByTrashed"

	page, pageSize := s.normalizePagination(req.Page, req.PageSize)
	search := req.Search

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("page", page),
		attribute.Int("pageSize", pageSize),
		attribute.String("search", search))

	defer func() {
		end(status)
	}()

	res, err := s.saldoQueryRepository.FindByTrashed(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandlerErrorPagination[[]*db.GetTrashedSaldosRow](
			s.logger,
			saldo_errors.ErrFailedFindTrashedSaldos,
			method,
			span,

			zap.Int("page", page),
			zap.Int("page_size", pageSize),
			zap.String("search", search),
		)
	}

	var totalCount int

	if len(res) > 0 {
		totalCount = int(res[0].TotalCount)
	} else {
		totalCount = 0
	}

	logSuccess("Successfully fetched trashed saldo",
		zap.Int("totalRecords", totalCount),
		zap.Int("page", req.Page),
		zap.Int("pageSize", req.PageSize))

	return res, &totalCount, nil
}

func (s *saldoQueryService) FindById(ctx context.Context, saldo_id int) (*db.GetSaldoByIDRow, error) {
	const method = "FindById"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("saldo_id", saldo_id))

	defer func() {
		end(status)
	}()

	res, err := s.saldoQueryRepository.FindById(ctx, saldo_id)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.GetSaldoByIDRow](
			s.logger,
			err,
			method,
			span,

			zap.Int("saldo_id", saldo_id),
		)
	}

	logSuccess("Successfully fetched saldo", zap.Int("saldo_id", saldo_id))

	return res, nil
}

func (s *saldoQueryService) FindByCardNumber(ctx context.Context, card_number string) (*db.Saldo, error) {
	const method = "FindByCardNumber"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.String("card_number", observability.MaskIdentifier(card_number)))

	defer func() {
		end(status)
	}()

	res, err := s.saldoQueryRepository.FindByCardNumber(ctx, card_number)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.Saldo](
			s.logger,
			err,
			method,
			span,

			zap.String("card_number", observability.MaskIdentifier(card_number)),
		)
	}

	logSuccess("Successfully fetched saldo by card number", zap.String("card_number", observability.MaskIdentifier(card_number)))

	return res, nil
}

func (s *saldoQueryService) normalizePagination(page, pageSize int) (int, int) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	return page, pageSize
}
