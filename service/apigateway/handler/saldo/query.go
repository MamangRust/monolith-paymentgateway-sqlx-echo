package saldohandler

import (
	"net/http"
	"strconv"

	saldo_cache "github.com/MamangRust/monolith-payment-gateway-apigateway/redis/api/saldo"
	pbhelper "github.com/MamangRust/monolith-payment-gateway-pb/card"
	pb "github.com/MamangRust/monolith-payment-gateway-pb/saldo"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
	apimapper "github.com/MamangRust/monolith-payment-gateway-shared/mapper/saldo"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type saldoQueryHandleApi struct {
	saldo pb.SaldoQueryServiceClient

	logger logger.LoggerInterface

	mapper apimapper.SaldoQueryResponseMapper

	cache saldo_cache.SaldoMencache

	apiHandler errors.ApiHandler
}

type saldoQueryHandleDeps struct {
	client pb.SaldoQueryServiceClient

	router *echo.Echo

	logger logger.LoggerInterface

	mapper apimapper.SaldoQueryResponseMapper

	cache saldo_cache.SaldoMencache

	apiHandler errors.ApiHandler
}

func NewSaldoQueryHandleApi(params *saldoQueryHandleDeps) *saldoQueryHandleApi {

	saldoHandler := &saldoQueryHandleApi{
		saldo:      params.client,
		logger:     params.logger,
		mapper:     params.mapper,
		cache:      params.cache,
		apiHandler: params.apiHandler,
	}

	routerSaldo := params.router.Group("/api/saldo-query")

	routerSaldo.GET("", params.apiHandler.Handle("find-all-saldos", saldoHandler.FindAll))
	routerSaldo.GET("/:id", params.apiHandler.Handle("find-saldo-by-id", saldoHandler.FindById))
	routerSaldo.GET("/active", params.apiHandler.Handle("find-active-saldos", saldoHandler.FindByActive))
	routerSaldo.GET("/trashed", params.apiHandler.Handle("find-trashed-saldos", saldoHandler.FindByTrashed))
	routerSaldo.GET("/card_number/:card_number", params.apiHandler.Handle("find-saldo-by-card-number", saldoHandler.FindByCardNumber))

	return saldoHandler
}

// @Summary Find all saldo data
// @Tags Saldo Query
// @Security Bearer
// @Description Retrieve a list of all saldo data with pagination and search
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Page size" default(10)
// @Param search query string false "Search query"
// @Success 200 {object} response.ApiResponsePaginationSaldo "List of saldo data"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve saldo data"
// @Router /api/saldo-query [get]
func (h *saldoQueryHandleApi) FindAll(c echo.Context) error {
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

	reqCache := &requests.FindAllSaldos{
		Page:     page,
		PageSize: pageSize,
		Search:   search,
	}

	cachedData, found := h.cache.GetCachedSaldos(ctx, reqCache)
	if found {
		return c.JSON(http.StatusOK, cachedData)
	}

	reqGrpc := &pb.FindAllSaldoRequest{
		Page:     int32(page),
		PageSize: int32(pageSize),
		Search:   search,
	}

	res, err := h.saldo.FindAllSaldo(ctx, reqGrpc)
	if err != nil {
		h.logger.Debug("Failed to retrieve saldo data", zap.Error(err))
		return errors.ParseGrpcError(err)
	}

	apiResponse := h.mapper.ToApiResponsePaginationSaldo(res)
	h.cache.SetCachedSaldos(ctx, reqCache, apiResponse)

	return c.JSON(http.StatusOK, apiResponse)
}

// @Summary Find a saldo by ID
// @Tags Saldo Query
// @Security Bearer
// @Description Retrieve a saldo by its ID
// @Accept json
// @Produce json
// @Param id path int true "Saldo ID"
// @Success 200 {object} response.ApiResponseSaldo "Saldo data"
// @Failure 400 {object} response.ErrorResponse "Invalid saldo ID"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve saldo data"
// @Router /api/saldo-query/{id} [get]
func (h *saldoQueryHandleApi) FindById(c echo.Context) error {
	idStr := c.Param("id")
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		h.logger.Debug("Invalid saldo ID", zap.Error(err))
		return errors.NewBadRequestError("invalid id parameter")
	}

	ctx := c.Request().Context()

	cachedData, found := h.cache.GetCachedSaldoById(ctx, idInt)
	if found {
		return c.JSON(http.StatusOK, cachedData)
	}

	reqGrpc := &pb.FindByIdSaldoRequest{
		SaldoId: int32(idInt),
	}

	res, err := h.saldo.FindByIdSaldo(ctx, reqGrpc)
	if err != nil {
		h.logger.Debug("Failed to retrieve saldo data", zap.Error(err))
		return errors.ParseGrpcError(err)
	}

	apiResponse := h.mapper.ToApiResponseSaldo(res)
	h.cache.SetCachedSaldoById(ctx, apiResponse.Data.ID, apiResponse)

	return c.JSON(http.StatusOK, apiResponse)
}

// @Summary Find a saldo by card number
// @Tags Saldo Query
// @Security Bearer
// @Description Retrieve a saldo by its card number
// @Accept json
// @Produce json
// @Param card_number path string true "Card number"
// @Success 200 {object} response.ApiResponseSaldo "Saldo data"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve saldo data"
// @Router /api/saldo-query/card_number/{card_number} [get]
func (h *saldoQueryHandleApi) FindByCardNumber(c echo.Context) error {
	cardNumber := c.Param("card_number")
	if cardNumber == "" {
		return errors.NewBadRequestError("card_number is required")
	}

	ctx := c.Request().Context()

	cachedData, found := h.cache.GetCachedSaldoByCardNumber(ctx, cardNumber)
	if found {
		return c.JSON(http.StatusOK, cachedData)
	}

	reqGrpc := &pbhelper.FindByCardNumberRequest{
		CardNumber: cardNumber,
	}

	res, err := h.saldo.FindByCardNumber(ctx, reqGrpc)
	if err != nil {
		h.logger.Debug("Failed to retrieve saldo data", zap.Error(err))
		return errors.ParseGrpcError(err)
	}

	apiResponse := h.mapper.ToApiResponseSaldo(res)
	h.cache.SetCachedSaldoByCardNumber(ctx, cardNumber, apiResponse)

	return c.JSON(http.StatusOK, apiResponse)
}

// @Summary Retrieve all active saldo data
// @Tags Saldo Query
// @Security Bearer
// @Description Retrieve a list of all active saldo data
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Page size" default(10)
// @Param search query string false "Search query"
// @Success 200 {object} response.ApiResponsesSaldo "List of saldo data"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve saldo data"
// @Router /api/saldo-query/active [get]
func (h *saldoQueryHandleApi) FindByActive(c echo.Context) error {
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

	reqCache := &requests.FindAllSaldos{
		Page:     page,
		PageSize: pageSize,
		Search:   search,
	}

	cachedData, found := h.cache.GetCachedSaldoByActive(ctx, reqCache)
	if found {
		return c.JSON(http.StatusOK, cachedData)
	}

	reqGrpc := &pb.FindAllSaldoRequest{
		Page:     int32(page),
		PageSize: int32(pageSize),
		Search:   search,
	}

	res, err := h.saldo.FindByActive(ctx, reqGrpc)
	if err != nil {
		h.logger.Debug("Failed to retrieve saldo data", zap.Error(err))
		return errors.ParseGrpcError(err)
	}

	apiResponse := h.mapper.ToApiResponsePaginationSaldoDeleteAt(res)
	h.cache.SetCachedSaldoByActive(ctx, reqCache, apiResponse)

	return c.JSON(http.StatusOK, apiResponse)
}

// @Summary Retrieve trashed saldo data
// @Tags Saldo Query
// @Security Bearer
// @Description Retrieve a list of all trashed saldo data
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Page size" default(10)
// @Param search query string false "Search query"
// @Success 200 {object} response.ApiResponsesSaldo "List of trashed saldo data"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve saldo data"
// @Router /api/saldo-query/trashed [get]
func (h *saldoQueryHandleApi) FindByTrashed(c echo.Context) error {
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

	reqCache := &requests.FindAllSaldos{
		Page:     page,
		PageSize: pageSize,
		Search:   search,
	}

	cachedData, found := h.cache.GetCachedSaldoByTrashed(ctx, reqCache)
	if found {
		return c.JSON(http.StatusOK, cachedData)
	}

	reqGrpc := &pb.FindAllSaldoRequest{
		Page:     int32(page),
		PageSize: int32(pageSize),
		Search:   search,
	}

	res, err := h.saldo.FindByTrashed(ctx, reqGrpc)
	if err != nil {
		h.logger.Debug("Failed to retrieve saldo data", zap.Error(err))
		return errors.ParseGrpcError(err)
	}

	apiResponse := h.mapper.ToApiResponsePaginationSaldoDeleteAt(res)
	h.cache.SetCachedSaldoByTrashed(ctx, reqCache, apiResponse)

	return c.JSON(http.StatusOK, apiResponse)
}
