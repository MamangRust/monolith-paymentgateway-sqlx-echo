package service

import (
	"context"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/errorhandler"
	transfer_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/transfer_errors/service"
	"github.com/MamangRust/monolith-payment-gateway-shared/observability"

	mencache "github.com/MamangRust/monolith-payment-gateway-transfer/redis"
	"github.com/MamangRust/monolith-payment-gateway-transfer/repository"
	"go.opentelemetry.io/otel/attribute"
	"go.uber.org/zap"
)

// transferQueryDeps defines dependencies for transferQueryService.
type TransferQueryDeps struct {
	Cache         mencache.TransferQueryCache
	Repository    repository.TransferQueryRepository
	Logger        logger.LoggerInterface
	Observability observability.TraceLoggerObservability
}

// transferQueryService handles read-only transfer queries.
type transferQueryService struct {
	cache                   mencache.TransferQueryCache
	transferQueryRepository repository.TransferQueryRepository
	logger                  logger.LoggerInterface
	observability           observability.TraceLoggerObservability
}

func NewTransferQueryService(
	params *TransferQueryDeps,
) TransferQueryService {
	return &transferQueryService{
		cache:                   params.Cache,
		transferQueryRepository: params.Repository,
		logger:                  params.Logger,
		observability:           params.Observability,
	}
}

func (s *transferQueryService) FindAll(ctx context.Context, req *requests.FindAllTransfers) ([]*db.GetTransfersRow, *int, error) {
	const method = "FindAll"

	page, pageSize := s.normalizePagination(req.Page, req.PageSize)
	search := req.Search

	if page <= 0 {
		page = 1
	}

	if pageSize <= 0 {
		pageSize = 10
	}

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("page", page),
		attribute.Int("pageSize", pageSize),
		attribute.String("search", search))

	defer func() {
		end(status)
	}()

	if data, total, found := s.cache.GetCachedTransfersCache(ctx, req); found {
		logSuccess("Successfully retrieved all transfer records from cache", zap.Int("totalRecords", *total), zap.Int("page", page), zap.Int("pageSize", pageSize))
		return data, total, nil
	}

	transfers, err := s.transferQueryRepository.FindAll(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandlerErrorPagination[[]*db.GetTransfersRow](
			s.logger,
			transfer_errors.ErrFailedFindAllTransfers,
			method,
			span,

			zap.Int("page", page),
			zap.Int("pageSize", pageSize),
			zap.String("search", search),
		)
	}

	var totalCount int

	if len(transfers) > 0 {
		totalCount = int(transfers[0].TotalCount)
	} else {
		totalCount = 0
	}

	s.cache.SetCachedTransfersCache(ctx, req, transfers, &totalCount)

	logSuccess("Successfully fetched transfer",
		zap.Int("totalRecords", totalCount),
		zap.Int("page", page),
		zap.Int("pageSize", pageSize))

	return transfers, &totalCount, nil
}

func (s *transferQueryService) FindById(ctx context.Context, transferId int) (*db.GetTransferByIDRow, error) {
	const method = "FindById"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("transfer_id", transferId))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetCachedTransferCache(ctx, transferId); found {
		logSuccess("Successfully fetched transfer from cache", zap.Int("transfer.id", transferId))
		return data, nil
	}

	transfer, err := s.transferQueryRepository.FindById(ctx, transferId)

	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.GetTransferByIDRow](
			s.logger,
			transfer_errors.ErrTransferNotFound,
			method,
			span,

			zap.Int("transfer_id", transferId),
		)
	}
	s.cache.SetCachedTransferCache(ctx, transfer)

	logSuccess("Successfully fetched transfer", zap.Int("transfer_id", transferId))

	return transfer, nil
}

func (s *transferQueryService) FindByActive(ctx context.Context, req *requests.FindAllTransfers) ([]*db.GetActiveTransfersRow, *int, error) {
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

	if data, total, found := s.cache.GetCachedTransferActiveCache(ctx, req); found {
		logSuccess("Successfully retrieved active transfer records from cache", zap.Int("totalRecords", *total), zap.Int("page", page), zap.Int("pageSize", pageSize))
		return data, total, nil
	}

	transfers, err := s.transferQueryRepository.FindByActive(ctx, req)

	if err != nil {
		status = "error"
		return errorhandler.HandlerErrorPagination[[]*db.GetActiveTransfersRow](
			s.logger,
			transfer_errors.ErrFailedFindActiveTransfers,
			method,
			span,

			zap.Int("page", page),
			zap.Int("pageSize", pageSize),
			zap.String("search", search),
		)
	}

	var totalCount int

	if len(transfers) > 0 {
		totalCount = int(transfers[0].TotalCount)
	} else {
		totalCount = 0
	}

	s.cache.SetCachedTransferActiveCache(ctx, req, transfers, &totalCount)

	logSuccess("Successfully fetched active transfer",
		zap.Int("totalRecords", totalCount),
		zap.Int("page", page),
		zap.Int("pageSize", pageSize))

	return transfers, &totalCount, nil
}

func (s *transferQueryService) FindByTrashed(ctx context.Context, req *requests.FindAllTransfers) ([]*db.GetTrashedTransfersRow, *int, error) {
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

	if data, total, found := s.cache.GetCachedTransferTrashedCache(ctx, req); found {
		logSuccess("Successfully retrieved trashed transfer records from cache", zap.Int("totalRecords", *total), zap.Int("page", page), zap.Int("pageSize", pageSize))
		return data, total, nil
	}

	transfers, err := s.transferQueryRepository.FindByTrashed(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandlerErrorPagination[[]*db.GetTrashedTransfersRow](
			s.logger,
			transfer_errors.ErrFailedFindTrashedTransfers,
			method,
			span,

			zap.Int("page", page),
			zap.Int("pageSize", pageSize),
			zap.String("search", search),
		)
	}

	var totalCount int

	if len(transfers) > 0 {
		totalCount = int(transfers[0].TotalCount)
	} else {
		totalCount = 0
	}

	s.cache.SetCachedTransferTrashedCache(ctx, req, transfers, &totalCount)

	logSuccess("Successfully fetched trashed transfer",
		zap.Int("totalRecords", totalCount),
		zap.Int("page", page),
		zap.Int("pageSize", pageSize))

	return transfers, &totalCount, nil
}

func (s *transferQueryService) FindTransferByTransferFrom(ctx context.Context, transfer_from string) ([]*db.GetTransfersBySourceCardRow, error) {
	const method = "FindTransferByTransferFrom"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.String("transfer_from", transfer_from))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetCachedTransferByFrom(ctx, transfer_from); found {
		logSuccess("Successfully fetched transfer from cache", zap.String("transfer_from", transfer_from))
		return data, nil
	}

	res, err := s.transferQueryRepository.FindTransferByTransferFrom(ctx, transfer_from)

	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetTransfersBySourceCardRow](
			s.logger,
			transfer_errors.ErrTransferNotFound,
			method,
			span,

			zap.String("transfer_from", transfer_from),
		)
	}

	s.cache.SetCachedTransferByFrom(ctx, transfer_from, res)

	logSuccess("Successfully fetched transfer record by transfer_from", zap.String("transfer_from", transfer_from))

	return res, nil
}

func (s *transferQueryService) FindTransferByTransferTo(ctx context.Context, transfer_to string) ([]*db.GetTransfersByDestinationCardRow, error) {
	const method = "FindTransferByTransferTo"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.String("transfer_to", transfer_to))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetCachedTransferByTo(ctx, transfer_to); found {
		logSuccess("Successfully fetched transfer from cache", zap.String("transfer_to", transfer_to))
		return data, nil
	}

	res, err := s.transferQueryRepository.FindTransferByTransferTo(ctx, transfer_to)

	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetTransfersByDestinationCardRow](
			s.logger,
			transfer_errors.ErrTransferNotFound,
			method,
			span,

			zap.String("transfer_to", transfer_to),
		)
	}

	s.cache.SetCachedTransferByTo(ctx, transfer_to, res)

	logSuccess("Successfully fetched transfer record by transfer_to", zap.String("transfer_to", transfer_to))

	return res, nil
}

func (s *transferQueryService) normalizePagination(page, pageSize int) (int, int) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	return page, pageSize
}
