package service

import (
	"context"

	mencache "github.com/MamangRust/monolith-payment-gateway-card/redis"
	"github.com/MamangRust/monolith-payment-gateway-card/repository"
	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	sharederrorhandler "github.com/MamangRust/monolith-payment-gateway-shared/errorhandler"
	card_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/card_errors/service"
	"github.com/MamangRust/monolith-payment-gateway-shared/observability"
	"go.opentelemetry.io/otel/attribute"
	"go.uber.org/zap"
)

// CardQueryServiceDeps defines dependencies for cardQueryService.
type CardQueryServiceDeps struct {
	Cache               mencache.CardQueryCache
	CardQueryRepository repository.CardQueryRepository
	Logger              logger.LoggerInterface
	Observability       observability.TraceLoggerObservability
}

// cardQueryService implements CardQueryService.
type cardQueryService struct {
	cache               mencache.CardQueryCache
	cardQueryRepository repository.CardQueryRepository
	logger              logger.LoggerInterface
	observability       observability.TraceLoggerObservability
}

// NewCardQueryService initializes a new instance of cardQueryService with the provided parameters.
//
// It sets up Prometheus metrics for counting and measuring the duration of card query requests.
//
// Parameters:
// - params: A pointer to CardQueryServiceDeps containing the necessary dependencies.
//
// Returns:
// - A pointer to a newly created cardQueryService.
func NewCardQueryService(
	params *CardQueryServiceDeps,
) CardQueryService {
	return &cardQueryService{
		cardQueryRepository: params.CardQueryRepository,
		logger:              params.Logger,
		observability:       params.Observability,
		cache:               params.Cache,
	}
}

func (s *cardQueryService) FindAll(ctx context.Context, req *requests.FindAllCards) ([]*db.GetCardsRow, *int, error) {
	const method = "FindAll"

	page := req.Page
	pageSize := req.PageSize

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	normalizedReq := *req
	normalizedReq.Page = page
	normalizedReq.PageSize = pageSize

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("page", page),
		attribute.Int("pageSize", pageSize),
		attribute.String("search", req.Search))

	defer func() {
		end(status)
	}()

	if data, total, found := s.cache.GetFindAllCache(ctx, &normalizedReq); found {
		logSuccess("Successfully fetched card records from cache", zap.Int("totalRecords", *total), zap.Int("page", page), zap.Int("pageSize", pageSize))
		return data, total, nil
	}

	cards, err := s.cardQueryRepository.FindAllCards(ctx, &normalizedReq)
	if err != nil {
		status = "error"
		return sharederrorhandler.HandlerErrorPagination[[]*db.GetCardsRow](
			s.logger,
			card_errors.ErrFailedFindAllCards,
			method,
			span,
			zap.Int("page", req.Page),
			zap.Int("pageSize", req.PageSize),
			zap.String("search", req.Search),
		)
	}

	var totalCount int

	if len(cards) > 0 {
		totalCount = int(cards[0].TotalCount)
	} else {
		totalCount = 0
	}

	s.cache.SetFindAllCache(ctx, &normalizedReq, cards, &totalCount)

	logSuccess("Successfully fetched card records",
		zap.Int("totalRecords", totalCount),

		zap.Int("page", page),
		zap.Int("pageSize", pageSize))

	return cards, &totalCount, nil
}

func (s *cardQueryService) FindByActive(ctx context.Context, req *requests.FindAllCards) ([]*db.GetActiveCardsWithCountRow, *int, error) {
	const method = "FindByActive"

	page := req.Page
	pageSize := req.PageSize
	search := req.Search

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	normalizedReq := *req
	normalizedReq.Page = page
	normalizedReq.PageSize = pageSize

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("page", page),
		attribute.Int("pageSize", pageSize),
		attribute.String("search", search))

	defer func() {
		end(status)
	}()

	if data, total, found := s.cache.GetByActiveCache(ctx, &normalizedReq); found {
		logSuccess("Successfully fetched active card records from cache", zap.Int("totalRecords", *total), zap.Int("page", page), zap.Int("pageSize", pageSize))
		return data, total, nil
	}

	res, err := s.cardQueryRepository.FindByActive(ctx, &normalizedReq)
	if err != nil {
		status = "error"
		return sharederrorhandler.HandlerErrorPagination[[]*db.GetActiveCardsWithCountRow](
			s.logger,
			card_errors.ErrFailedFindActiveCards,
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

	s.cache.SetByActiveCache(ctx, &normalizedReq, res, &totalCount)

	logSuccess("Successfully fetched active card records",
		zap.Int("totalRecords", totalCount),
		zap.Int("page", page),
		zap.Int("pageSize", pageSize))

	return res, &totalCount, nil
}

func (s *cardQueryService) FindByTrashed(ctx context.Context, req *requests.FindAllCards) ([]*db.GetTrashedCardsWithCountRow, *int, error) {
	const method = "FindByTrashed"

	page := req.Page
	pageSize := req.PageSize
	search := req.Search

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	normalizedReq := *req
	normalizedReq.Page = page
	normalizedReq.PageSize = pageSize

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("page", page),
		attribute.Int("pageSize", pageSize),
		attribute.String("search", search))

	defer func() {
		end(status)
	}()

	if data, total, found := s.cache.GetByTrashedCache(ctx, &normalizedReq); found {
		logSuccess("Successfully fetched trashed card records from cache", zap.Int("totalRecords", *total), zap.Int("page", page), zap.Int("pageSize", pageSize))
		return data, total, nil
	}

	res, err := s.cardQueryRepository.FindByTrashed(ctx, &normalizedReq)
	if err != nil {
		status = "error"
		return sharederrorhandler.HandlerErrorPagination[[]*db.GetTrashedCardsWithCountRow](
			s.logger,
			card_errors.ErrFailedFindTrashedCards,
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

	s.cache.SetByTrashedCache(ctx, &normalizedReq, res, &totalCount)

	logSuccess("Successfully fetched trashed card records",
		zap.Int("totalRecords", totalCount),
		zap.Int("page", page),
		zap.Int("pageSize", pageSize))

	return res, &totalCount, nil
}

func (s *cardQueryService) FindById(ctx context.Context, card_id int) (*db.GetCardByIDRow, error) {
	const method = "FindById"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("card_id", card_id))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetByIdCache(ctx, card_id); found {
		logSuccess("Successfully fetched card from cache", zap.Int("card.id", card_id))
		return data, nil
	}

	res, err := s.cardQueryRepository.FindById(ctx, card_id)
	if err != nil {
		status = "error"
		return sharederrorhandler.HandleError[*db.GetCardByIDRow](
			s.logger,
			err,
			method,
			span,

			zap.Int("card_id", card_id),
		)
	}

	s.cache.SetByIdCache(ctx, card_id, res)

	logSuccess("Successfully fetched card", zap.Int("card_id", card_id))

	return res, nil
}

func (s *cardQueryService) FindByCardNumber(ctx context.Context, card_number string) (*db.GetCardByCardNumberRow, error) {
	const method = "FindByCardNumber"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.String("card_number", observability.MaskIdentifier(card_number)))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetByCardNumberCache(ctx, card_number); found {
		logSuccess("Successfully fetched card record by card number from cache", zap.String("card_number", observability.MaskIdentifier(card_number)))
		return data, nil
	}

	res, err := s.cardQueryRepository.FindCardByCardNumber(ctx, card_number)
	if err != nil {
		status = "error"
		return sharederrorhandler.HandleError[*db.GetCardByCardNumberRow](
			s.logger,
			card_errors.ErrCardNotFoundRes,
			method,
			span,
			zap.String("card_number", observability.MaskIdentifier(card_number)),
		)
	}

	s.cache.SetByCardNumberCache(ctx, card_number, res)

	logSuccess("Successfully fetched card record by card number", zap.String("card_number", observability.MaskIdentifier(card_number)))

	return res, nil
}

func (s *cardQueryService) FindByUserID(ctx context.Context, user_id int) (*db.GetCardByUserIDRow, error) {
	const method = "FindByUserID"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("user_id", user_id))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetByUserIDCache(ctx, user_id); found {
		logSuccess("Successfully fetched card records by user ID from cache", zap.Int("user.id", user_id))
		return data, nil
	}

	res, err := s.cardQueryRepository.FindCardByUserId(ctx, user_id)
	if err != nil {
		status = "error"
		return sharederrorhandler.HandleError[*db.GetCardByUserIDRow](
			s.logger,
			card_errors.ErrFailedFindByUserID,
			method,
			span,

			zap.Int("user_id", user_id),
		)
	}

	s.cache.SetByUserIDCache(ctx, user_id, res)

	logSuccess("Successfully fetched card records by user ID", zap.Int("user_id", user_id))

	return res, nil
}

func (s *cardQueryService) FindUserCardByCardNumber(ctx context.Context, card_number string) (*db.GetUserEmailByCardNumberRow, error) {
	const method = "FindUserCardByCardNumber"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.String("card_number", observability.MaskIdentifier(card_number)))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetUserCardByCardNumberCache(ctx, card_number); found {
		logSuccess("Successfully fetched card records by user ID from cache", zap.String("card_number", observability.MaskIdentifier(card_number)))
		return data, nil
	}

	res, err := s.cardQueryRepository.FindUserCardByCardNumber(ctx, card_number)
	if err != nil {
		status = "error"
		return sharederrorhandler.HandleError[*db.GetUserEmailByCardNumberRow](
			s.logger,
			card_errors.ErrFailedFindByUserID,
			method,
			span,
			zap.String("card_number", observability.MaskIdentifier(card_number)),
		)
	}

	s.cache.SetUserCardByCardNumberCache(ctx, card_number, res)

	logSuccess("Successfully fetched card records by user ID", zap.String("card_number", observability.MaskIdentifier(card_number)))

	return res, nil
}

// normalizePagination normalizes pagination page and pageSize arguments.
//
// If page or pageSize is less than or equal to 0, it is set to the default value of 1 or 10, respectively.
//
// Parameters:
//   - page: The input page number.
//   - pageSize: The input page size.
//
// Returns:
//   - The normalized page number.
//   - The normalized page size.
func (s *cardQueryService) normalizePagination(page, pageSize int) (int, int) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	return page, pageSize
}
