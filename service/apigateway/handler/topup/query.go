package topuphandler

import (
	"net/http"
	"strconv"

	topup_cache "github.com/MamangRust/monolith-payment-gateway-apigateway/redis/api/topup"
	pb "github.com/MamangRust/monolith-payment-gateway-pb/topup"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
	apimapper "github.com/MamangRust/monolith-payment-gateway-shared/mapper/topup"

	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type topupQueryHandleApi struct {
	client pb.TopupQueryServiceClient

	logger logger.LoggerInterface

	mapper apimapper.TopupQueryResponseMapper

	cache topup_cache.TopupMencach

	apiHandler errors.ApiHandler
}

type topupQueryHandleDeps struct {
	client pb.TopupQueryServiceClient

	router *echo.Echo

	logger logger.LoggerInterface

	mapper apimapper.TopupQueryResponseMapper

	cache topup_cache.TopupMencach

	apiHandler errors.ApiHandler
}

func NewTopupQueryHandleApi(params *topupQueryHandleDeps) *topupQueryHandleApi {

	topupHandler := &topupQueryHandleApi{
		client:     params.client,
		logger:     params.logger,
		mapper:     params.mapper,
		cache:      params.cache,
		apiHandler: params.apiHandler,
	}

	routerTopup := params.router.Group("/api/topup-query")

	routerTopup.GET("", params.apiHandler.Handle("find-all-topups", topupHandler.FindAll))
	routerTopup.GET("/card-number/:card_number", params.apiHandler.Handle("find-all-topups-by-card-number", topupHandler.FindAllByCardNumber))
	routerTopup.GET("/:id", params.apiHandler.Handle("find-topup-by-id", topupHandler.FindById))
	routerTopup.GET("/active", params.apiHandler.Handle("find-active-topups", topupHandler.FindByActive))
	routerTopup.GET("/trashed", params.apiHandler.Handle("find-trashed-topups", topupHandler.FindByTrashed))

	return topupHandler
}

// @Tags Topup Query
// @Security Bearer
// @Description Retrieve a list of all topup data with pagination and search
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Page size" default(10)
// @Param search query string false "Search query"
// @Success 200 {object} response.ApiResponsePaginationTopup "List of topup data"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve topup data"
// @Router /api/topup-query [get]
func (h *topupQueryHandleApi) FindAll(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page <= 0 {
		page = 1
	}

	pageSize, err := strconv.Atoi(c.QueryParam("page_size"))
	if err != nil || pageSize <= 0 {
		pageSize = 10
	}

	search := c.QueryParam("search")

	ctx := c.Request().Context()

	reqCache := &requests.FindAllTopups{
		Page:     page,
		PageSize: pageSize,
		Search:   search,
	}

	cachedData, found := h.cache.GetCachedTopupsCache(ctx, reqCache)
	if found {
		return c.JSON(http.StatusOK, cachedData)
	}

	reqGrpc := &pb.FindAllTopupRequest{
		Page:     int32(page),
		PageSize: int32(pageSize),
		Search:   search,
	}

	res, err := h.client.FindAllTopup(ctx, reqGrpc)
	if err != nil {
		h.logger.Debug("Failed to retrieve topup data", zap.Error(err))
		return errors.ParseGrpcError(err)
	}

	apiResponse := h.mapper.ToApiResponsePaginationTopup(res)
	h.cache.SetCachedTopupsCache(ctx, reqCache, apiResponse)

	return c.JSON(http.StatusOK, apiResponse)
}

// @Summary Find all topup by card number
// @Tags Topup Query
// @Security Bearer
// @Description Retrieve a list of transactions for a specific card number
// @Accept json
// @Produce json
// @Param card_number path string true "Card Number"
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Number of items per page" default(10)
// @Param search query string false "Search query"
// @Success 200 {object} response.ApiResponsePaginationTopup "List of topups"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve topups data"
// @Router /api/topup-query/card-number/{card_number} [get]
func (h *topupQueryHandleApi) FindAllByCardNumber(c echo.Context) error {
	cardNumber := c.Param("card_number")
	if cardNumber == "" {
		return errors.NewBadRequestError("card_number is required")
	}

	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page <= 0 {
		page = 1
	}

	pageSize, err := strconv.Atoi(c.QueryParam("page_size"))
	if err != nil || pageSize <= 0 {
		pageSize = 10
	}

	search := c.QueryParam("search")

	ctx := c.Request().Context()

	reqCache := &requests.FindAllTopupsByCardNumber{
		CardNumber: cardNumber,
		Page:       page,
		PageSize:   pageSize,
		Search:     search,
	}

	cachedData, found := h.cache.GetCacheTopupByCardCache(ctx, reqCache)
	if found {
		return c.JSON(http.StatusOK, cachedData)
	}

	reqGrpc := &pb.FindAllTopupByCardNumberRequest{
		CardNumber: cardNumber,
		Page:       int32(page),
		PageSize:   int32(pageSize),
		Search:     search,
	}

	res, err := h.client.FindAllTopupByCardNumber(ctx, reqGrpc)
	if err != nil {
		h.logger.Debug("Failed to retrieve topup data", zap.Error(err))
		return errors.ParseGrpcError(err)
	}

	apiResponse := h.mapper.ToApiResponsePaginationTopup(res)
	h.cache.SetCacheTopupByCardCache(ctx, reqCache, apiResponse)

	return c.JSON(http.StatusOK, apiResponse)
}

// @Summary Find a topup by ID
// @Tags Topup Query
// @Security Bearer
// @Description Retrieve a topup record using its ID
// @Accept json
// @Produce json
// @Param id path string true "Topup ID"
// @Success 200 {object} response.ApiResponseTopup "Topup data"
// @Failure 400 {object} response.ErrorResponse "Bad Request: Invalid ID"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve topup data"
// @Router /api/topup-query/{id} [get]
func (h *topupQueryHandleApi) FindById(c echo.Context) error {
	idStr := c.Param("id")
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		return errors.NewBadRequestError("invalid id parameter")
	}

	ctx := c.Request().Context()

	cachedData, found := h.cache.GetCachedTopupCache(ctx, idInt)
	if found {
		return c.JSON(http.StatusOK, cachedData)
	}

	res, err := h.client.FindByIdTopup(ctx, &pb.FindByIdTopupRequest{
		TopupId: int32(idInt),
	})
	if err != nil {
		h.logger.Debug("Failed to retrieve topup data", zap.Error(err))
		return errors.ParseGrpcError(err)
	}

	apiResponse := h.mapper.ToApiResponseTopup(res)
	h.cache.SetCachedTopupCache(ctx, apiResponse)

	return c.JSON(http.StatusOK, apiResponse)
}

// @Summary Find active topups
// @Tags Topup Query
// @Security Bearer
// @Description Retrieve a list of active topup records
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Page size" default(10)
// @Param search query string false "Search query"
// @Success 200 {object} response.ApiResponsesTopup "Active topup data"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve topup data"
// @Router /api/topup-query/active [get]
func (h *topupQueryHandleApi) FindByActive(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page <= 0 {
		page = 1
	}

	pageSize, err := strconv.Atoi(c.QueryParam("page_size"))
	if err != nil || pageSize <= 0 {
		pageSize = 10
	}

	search := c.QueryParam("search")

	ctx := c.Request().Context()

	reqCache := &requests.FindAllTopups{
		Page:     page,
		PageSize: pageSize,
		Search:   search,
	}

	cachedData, found := h.cache.GetCachedTopupActiveCache(ctx, reqCache)
	if found {
		return c.JSON(http.StatusOK, cachedData)
	}

	reqGrpc := &pb.FindAllTopupRequest{
		Page:     int32(page),
		PageSize: int32(pageSize),
		Search:   search,
	}

	res, err := h.client.FindByActive(ctx, reqGrpc)
	if err != nil {
		h.logger.Debug("Failed to retrieve topup data", zap.Error(err))
		return errors.ParseGrpcError(err)
	}

	apiResponse := h.mapper.ToApiResponsePaginationTopupDeleteAt(res)
	h.cache.SetCachedTopupActiveCache(ctx, reqCache, apiResponse)

	return c.JSON(http.StatusOK, apiResponse)
}

// @Summary Retrieve trashed topups
// @Tags Topup Query
// @Security Bearer
// @Description Retrieve a list of trashed topup records
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Page size" default(10)
// @Param search query string false "Search query"
// @Success 200 {object} response.ApiResponsesTopup "List of trashed topup data"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve topup data"
// @Router /api/topup-query/trashed [get]
func (h *topupQueryHandleApi) FindByTrashed(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page <= 0 {
		page = 1
	}

	pageSize, err := strconv.Atoi(c.QueryParam("page_size"))
	if err != nil || pageSize <= 0 {
		pageSize = 10
	}

	search := c.QueryParam("search")

	ctx := c.Request().Context()

	reqCache := &requests.FindAllTopups{
		Page:     page,
		PageSize: pageSize,
		Search:   search,
	}

	cachedData, found := h.cache.GetCachedTopupTrashedCache(ctx, reqCache)
	if found {
		return c.JSON(http.StatusOK, cachedData)
	}

	reqGrpc := &pb.FindAllTopupRequest{
		Page:     int32(page),
		PageSize: int32(pageSize),
		Search:   search,
	}

	res, err := h.client.FindByTrashed(ctx, reqGrpc)
	if err != nil {
		h.logger.Debug("Failed to retrieve topup data", zap.Error(err))
		return errors.ParseGrpcError(err)
	}

	apiResponse := h.mapper.ToApiResponsePaginationTopupDeleteAt(res)
	h.cache.SetCachedTopupTrashedCache(ctx, reqCache, apiResponse)

	return c.JSON(http.StatusOK, apiResponse)
}
