package transactionhandler

import (
	"net/http"
	"strconv"

	transaction_cache "github.com/MamangRust/monolith-payment-gateway-apigateway/redis/api/transaction"
	pb "github.com/MamangRust/monolith-payment-gateway-pb/transaction"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
	apimapper "github.com/MamangRust/monolith-payment-gateway-shared/mapper/transaction"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type transactionQueryHandleApi struct {
	client pb.TransactionQueryServiceClient

	logger logger.LoggerInterface

	mapper apimapper.TransactionQueryResponseMapper

	cache transaction_cache.TransactionMencache

	apiHandler errors.ApiHandler
}

type transactionQueryHandleDeps struct {
	client pb.TransactionQueryServiceClient

	router *echo.Echo

	logger logger.LoggerInterface

	mapper apimapper.TransactionQueryResponseMapper

	cache transaction_cache.TransactionMencache

	apiHandler errors.ApiHandler
}

func NewTransactionQueryHandleApi(params *transactionQueryHandleDeps) *transactionQueryHandleApi {

	transactionQueryHandleApi := &transactionQueryHandleApi{
		client:     params.client,
		logger:     params.logger,
		mapper:     params.mapper,
		cache:      params.cache,
		apiHandler: params.apiHandler,
	}

	routerTransaction := params.router.Group("/api/transaction-query")

	routerTransaction.GET("", transactionQueryHandleApi.FindAll)
	routerTransaction.GET("/card-number/:card_number", transactionQueryHandleApi.FindAllTransactionByCardNumber)
	routerTransaction.GET("/:id", transactionQueryHandleApi.FindById)
	routerTransaction.GET("/merchant/:merchant_id", transactionQueryHandleApi.FindByTransactionMerchantId)
	routerTransaction.GET("/active", transactionQueryHandleApi.FindByActiveTransaction)
	routerTransaction.GET("/trashed", transactionQueryHandleApi.FindByTrashedTransaction)

	return transactionQueryHandleApi
}

// @Summary Find all
// @Tags Transaction Query
// @Security Bearer
// @Description Retrieve a list of all transactions
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Number of items per page" default(10)
// @Param search query string false "Search query"
// @Success 200 {object} response.ApiResponsePaginationTransaction "List of transactions"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve transaction data"
// @Router /api/transaction-query [get]
func (h *transactionQueryHandleApi) FindAll(c echo.Context) error {
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

	reqCache := &requests.FindAllTransactions{
		Page:     page,
		PageSize: pageSize,
		Search:   search,
	}

	cachedData, found := h.cache.GetCachedTransactionsCache(ctx, reqCache)
	if found {
		return c.JSON(http.StatusOK, cachedData)
	}

	reqGrpc := &pb.FindAllTransactionRequest{
		Page:     int32(page),
		PageSize: int32(pageSize),
		Search:   search,
	}

	res, err := h.client.FindAllTransaction(ctx, reqGrpc)
	if err != nil {
		h.logger.Debug("Failed to retrieve transaction data", zap.Error(err))
		return errors.ParseGrpcError(err)
	}

	apiResponse := h.mapper.ToApiResponsePaginationTransaction(res)
	h.cache.SetCachedTransactionsCache(ctx, reqCache, apiResponse)

	return c.JSON(http.StatusOK, apiResponse)
}

// @Summary Find all transactions by card number
// @Tags Transaction Query
// @Security Bearer
// @Description Retrieve a list of transactions for a specific card number
// @Accept json
// @Produce json
// @Param card_number path string true "Card Number"
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Number of items per page" default(10)
// @Param search query string false "Search query"
// @Success 200 {object} response.ApiResponsePaginationTransaction "List of transactions"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve transaction data"
// @Router /api/transaction-query/card-number/{card_number} [get]
func (h *transactionQueryHandleApi) FindAllTransactionByCardNumber(c echo.Context) error {
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

	reqCache := &requests.FindAllTransactionCardNumber{
		CardNumber: cardNumber,
		Page:       page,
		PageSize:   pageSize,
		Search:     search,
	}

	cachedData, found := h.cache.GetCachedTransactionByCardNumberCache(ctx, reqCache)
	if found {
		return c.JSON(http.StatusOK, cachedData)
	}

	reqGrpc := &pb.FindAllTransactionCardNumberRequest{
		CardNumber: cardNumber,
		Page:       int32(page),
		PageSize:   int32(pageSize),
		Search:     search,
	}

	res, err := h.client.FindAllTransactionByCardNumber(ctx, reqGrpc)
	if err != nil {
		h.logger.Debug("Failed to retrieve transaction data", zap.Error(err))
		return errors.ParseGrpcError(err)
	}

	apiResponse := h.mapper.ToApiResponsePaginationTransaction(res)
	h.cache.SetCachedTransactionByCardNumberCache(ctx, reqCache, apiResponse)

	return c.JSON(http.StatusOK, apiResponse)
}

// @Summary Find a transaction by ID
// @Tags Transaction Query
// @Security Bearer
// @Description Retrieve a transaction record using its ID
// @Accept json
// @Produce json
// @Param id path string true "Transaction ID"
// @Success 200 {object} response.ApiResponseTransaction "Transaction data"
// @Failure 400 {object} response.ErrorResponse "Bad Request: Invalid ID"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve transaction data"
// @Router /api/transaction-query/{id} [get]
func (h *transactionQueryHandleApi) FindById(c echo.Context) error {
	idStr := c.Param("id")
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		h.logger.Debug("Invalid transaction ID", zap.Error(err))
		return errors.NewBadRequestError("invalid id parameter")
	}

	ctx := c.Request().Context()

	cachedData, found := h.cache.GetCachedTransactionCache(ctx, idInt)
	if found {
		return c.JSON(http.StatusOK, cachedData)
	}

	res, err := h.client.FindByIdTransaction(ctx, &pb.FindByIdTransactionRequest{
		TransactionId: int32(idInt),
	})
	if err != nil {
		h.logger.Debug("Failed to retrieve transaction data", zap.Error(err))
		return errors.ParseGrpcError(err)
	}

	apiResponse := h.mapper.ToApiResponseTransaction(res)
	h.cache.SetCachedTransactionCache(ctx, apiResponse)

	return c.JSON(http.StatusOK, apiResponse)
}

// @Summary Find transactions by merchant ID
// @Tags Transaction
// @Security Bearer
// @Description Retrieve a list of transactions using the merchant ID
// @Accept json
// @Produce json
// @Param merchant_id query string true "Merchant ID"
// @Success 200 {object} response.ApiResponseTransactions "Transaction data"
// @Failure 400 {object} response.ErrorResponse "Bad Request: Invalid ID"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve transaction data"
// @Router /api/transactions-query/merchant/{merchant_id} [get]
func (h *transactionQueryHandleApi) FindByTransactionMerchantId(c echo.Context) error {
	merchantIdStr := c.Param("merchant_id")
	merchantIdInt, err := strconv.Atoi(merchantIdStr)
	if err != nil {
		return errors.NewBadRequestError("invalid merchant_id parameter")
	}

	ctx := c.Request().Context()

	cachedData, found := h.cache.GetCachedTransactionByMerchantIdCache(ctx, merchantIdInt)
	if found {
		return c.JSON(http.StatusOK, cachedData)
	}

	reqGrpc := &pb.FindTransactionByMerchantIdRequest{
		MerchantId: int32(merchantIdInt),
	}

	res, err := h.client.FindTransactionByMerchantId(ctx, reqGrpc)
	if err != nil {
		h.logger.Debug("Failed to retrieve transaction data", zap.Error(err))
		return errors.ParseGrpcError(err)
	}

	apiResponse := h.mapper.ToApiResponseTransactions(res)
	h.cache.SetCachedTransactionByMerchantIdCache(ctx, merchantIdInt, apiResponse)

	return c.JSON(http.StatusOK, apiResponse)
}

// @Summary Find active transactions
// @Tags Transaction Query
// @Security Bearer
// @Description Retrieve a list of active transactions
// @Accept json
// @Produce json
// @Param page query int false "Page number (default: 1)"
// @Param page_size query int false "Number of items per page (default: 10)"
// @Param search query string false "Search keyword"
// @Success 200 {object} response.ApiResponseTransactions "List of active transactions"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve transaction data"
// @Router /api/transaction-query/active [get]
func (h *transactionQueryHandleApi) FindByActiveTransaction(c echo.Context) error {
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

	reqCache := &requests.FindAllTransactions{
		Page:     page,
		PageSize: pageSize,
		Search:   search,
	}

	cachedData, found := h.cache.GetCachedTransactionActiveCache(ctx, reqCache)
	if found {
		return c.JSON(http.StatusOK, cachedData)
	}

	reqGrpc := &pb.FindAllTransactionRequest{
		Page:     int32(page),
		PageSize: int32(pageSize),
		Search:   search,
	}

	res, err := h.client.FindByActiveTransaction(ctx, reqGrpc)
	if err != nil {
		h.logger.Debug("Failed to retrieve transaction data", zap.Error(err))
		return errors.ParseGrpcError(err)
	}

	apiResponse := h.mapper.ToApiResponsePaginationTransactionDeleteAt(res)
	h.cache.SetCachedTransactionActiveCache(ctx, reqCache, apiResponse)

	return c.JSON(http.StatusOK, apiResponse)
}

// @Summary Retrieve trashed transactions
// @Tags Transaction Query
// @Security Bearer
// @Description Retrieve a list of trashed transactions
// @Accept json
// @Produce json
// @Param page query int false "Page number (default: 1)"
// @Param page_size query int false "Number of items per page (default: 10)"
// @Param search query string false "Search keyword"
// @Success 200 {object} response.ApiResponseTransactions "List of trashed transactions"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve transaction data"
// @Router /api/transaction-query/trashed [get]
func (h *transactionQueryHandleApi) FindByTrashedTransaction(c echo.Context) error {
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

	reqCache := &requests.FindAllTransactions{
		Page:     page,
		PageSize: pageSize,
		Search:   search,
	}

	cachedData, found := h.cache.GetCachedTransactionTrashedCache(ctx, reqCache)
	if found {
		return c.JSON(http.StatusOK, cachedData)
	}

	reqGrpc := &pb.FindAllTransactionRequest{
		Page:     int32(page),
		PageSize: int32(pageSize),
		Search:   search,
	}

	res, err := h.client.FindByTrashedTransaction(ctx, reqGrpc)
	if err != nil {
		h.logger.Debug("Failed to retrieve transaction data", zap.Error(err))
		return errors.ParseGrpcError(err)
	}

	apiResponse := h.mapper.ToApiResponsePaginationTransactionDeleteAt(res)
	h.cache.SetCachedTransactionTrashedCache(ctx, reqCache, apiResponse)

	return c.JSON(http.StatusOK, apiResponse)
}
