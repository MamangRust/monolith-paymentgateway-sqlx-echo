package transactionhandler

import (
	"net/http"
	"strconv"

	transaction_cache "github.com/MamangRust/monolith-payment-gateway-apigateway/redis/api/transaction"
	pbtransaction "github.com/MamangRust/monolith-payment-gateway-pb/transaction"
	pb "github.com/MamangRust/monolith-payment-gateway-pb/transaction/stats"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
	apimapper "github.com/MamangRust/monolith-payment-gateway-shared/mapper/transaction"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type transactionStatsAmountHandleApi struct {
	client pb.TransactionStatsAmountServiceClient

	logger logger.LoggerInterface

	mapper apimapper.TransactionStatsAmountResponseMapper

	cache transaction_cache.TransactionMencache

	apiHandler errors.ApiHandler
}

type transactionStatsAmountHandleDeps struct {
	client pb.TransactionStatsAmountServiceClient

	router *echo.Echo

	logger logger.LoggerInterface

	mapper apimapper.TransactionStatsAmountResponseMapper

	cache transaction_cache.TransactionMencache

	apiHandler errors.ApiHandler
}

func NewTransactionStatsAmountHandleApi(params *transactionStatsAmountHandleDeps) *transactionStatsAmountHandleApi {

	transactionStatsAmountHandleApi := &transactionStatsAmountHandleApi{
		client:     params.client,
		logger:     params.logger,
		mapper:     params.mapper,
		cache:      params.cache,
		apiHandler: params.apiHandler,
	}

	routerTransaction := params.router.Group("/api/transaction-stats-amount")

	routerTransaction.GET("/monthly-amounts-by-card", params.apiHandler.Handle("find-monthly-amounts-by-card", transactionStatsAmountHandleApi.FindMonthlyAmountsByCardNumber))
	routerTransaction.GET("/yearly-amounts-by-card", params.apiHandler.Handle("find-yearly-amounts-by-card", transactionStatsAmountHandleApi.FindYearlyAmountsByCardNumber))
	routerTransaction.GET("/monthly-amounts", params.apiHandler.Handle("find-monthly-amounts", transactionStatsAmountHandleApi.FindMonthlyAmounts))
	routerTransaction.GET("/yearly-amounts", params.apiHandler.Handle("find-yearly-amounts", transactionStatsAmountHandleApi.FindYearlyAmounts))

	return transactionStatsAmountHandleApi
}

// FindMonthlyAmounts retrieves the monthly transaction amounts for a specific year.
// @Summary Get monthly transaction amounts
// @Tags Transaction Stats Amount
// @Security Bearer
// @Description Retrieve the monthly transaction amounts for a specific year.
// @Accept json
// @Produce json
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponseTransactionMonthAmount "Monthly transaction amounts"
// @Failure 400 {object} response.ErrorResponse "Invalid year parameter"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve monthly transaction amounts"
// @Router /api/transaction-stats-amount/monthly-amounts [get]
func (h *transactionStatsAmountHandleApi) FindMonthlyAmounts(c echo.Context) error {
	yearStr := c.QueryParam("year")
	year, err := strconv.Atoi(yearStr)
	if err != nil || year <= 0 {
		return errors.NewBadRequestError("invalid year parameter")
	}

	ctx := c.Request().Context()

	cachedData, found := h.cache.GetMonthlyAmountsCache(ctx, year)
	if found {
		return c.JSON(http.StatusOK, cachedData)
	}

	res, err := h.client.FindMonthlyAmounts(ctx, &pbtransaction.FindYearTransactionStatus{
		Year: int32(year),
	})
	if err != nil {
		h.logger.Debug("Failed to retrieve monthly amounts", zap.Error(err))
		return errors.ParseGrpcError(err)
	}

	apiResponse := h.mapper.ToApiResponseTransactionMonthAmount(res)
	h.cache.SetMonthlyAmountsCache(ctx, year, apiResponse)

	return c.JSON(http.StatusOK, apiResponse)
}

// FindYearlyAmounts retrieves the yearly transaction amounts for a specific year.
// @Summary Get yearly transaction amounts
// @Tags Transaction Stats Amount
// @Security Bearer
// @Description Retrieve the yearly transaction amounts for a specific year.
// @Accept json
// @Produce json
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponseTransactionYearAmount "Yearly transaction amounts"
// @Failure 400 {object} response.ErrorResponse "Invalid year parameter"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve yearly transaction amounts"
// @Router /api/transaction-stats-amount/yearly-amounts [get]
func (h *transactionStatsAmountHandleApi) FindYearlyAmounts(c echo.Context) error {
	yearStr := c.QueryParam("year")
	year, err := strconv.Atoi(yearStr)
	if err != nil || year <= 0 {
		return errors.NewBadRequestError("invalid year parameter")
	}

	ctx := c.Request().Context()

	cachedData, found := h.cache.GetYearlyAmountsCache(ctx, year)
	if found {
		return c.JSON(http.StatusOK, cachedData)
	}

	res, err := h.client.FindYearlyAmounts(ctx, &pbtransaction.FindYearTransactionStatus{
		Year: int32(year),
	})
	if err != nil {
		h.logger.Debug("Failed to retrieve yearly amounts", zap.Error(err))
		return errors.ParseGrpcError(err)
	}

	apiResponse := h.mapper.ToApiResponseTransactionYearAmount(res)
	h.cache.SetYearlyAmountsCache(ctx, year, apiResponse)

	return c.JSON(http.StatusOK, apiResponse)
}

// FindMonthlyAmountsByCardNumber retrieves the monthly transaction amounts for a specific card number and year.
// @Summary Get monthly transaction amounts by card number
// @Tags Transaction Stats Amount
// @Security Bearer
// @Description Retrieve the monthly transaction amounts for a specific card number and year.
// @Accept json
// @Produce json
// @Param card_number query string true "Card Number"
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponseTransactionMonthAmount "Monthly transaction amounts by card number"
// @Failure 400 {object} response.ErrorResponse "Invalid card number or year parameter"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve monthly transaction amounts by card number"
// @Router /api/transaction-stats-amount/monthly-amounts-by-card [get]
func (h *transactionStatsAmountHandleApi) FindMonthlyAmountsByCardNumber(c echo.Context) error {
	yearStr := c.QueryParam("year")
	cardNumber := c.QueryParam("card_number")

	year, err := strconv.Atoi(yearStr)
	if err != nil || year <= 0 {
		return errors.NewBadRequestError("invalid year parameter")
	}

	if cardNumber == "" {
		return errors.NewBadRequestError("card_number is required")
	}

	ctx := c.Request().Context()

	reqCache := &requests.MonthYearPaymentMethod{
		CardNumber: cardNumber,
		Year:       year,
	}

	cachedData, found := h.cache.GetMonthlyAmountsByCardCache(ctx, reqCache)
	if found {
		return c.JSON(http.StatusOK, cachedData)
	}

	res, err := h.client.FindMonthlyAmountsByCardNumber(ctx, &pbtransaction.FindByYearCardNumberTransactionRequest{
		CardNumber: cardNumber,
		Year:       int32(year),
	})
	if err != nil {
		h.logger.Debug("Failed to retrieve monthly amounts by card number", zap.Error(err))
		return errors.ParseGrpcError(err)
	}

	apiResponse := h.mapper.ToApiResponseTransactionMonthAmount(res)
	h.cache.SetMonthlyAmountsByCardCache(ctx, reqCache, apiResponse)

	return c.JSON(http.StatusOK, apiResponse)
}

// FindYearlyAmountsByCardNumber retrieves the yearly transaction amounts for a specific card number and year.
// @Summary Get yearly transaction amounts by card number
// @Tags Transaction Stats Amount
// @Security Bearer
// @Description Retrieve the yearly transaction amounts for a specific card number and year.
// @Accept json
// @Produce json
// @Param card_number query string true "Card Number"
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponseTransactionYearAmount "Yearly transaction amounts by card number"
// @Failure 400 {object} response.ErrorResponse "Invalid card number or year parameter"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve yearly transaction amounts by card number"
// @Router /api/transaction-stats-amount/yearly-amounts-by-card [get]
func (h *transactionStatsAmountHandleApi) FindYearlyAmountsByCardNumber(c echo.Context) error {
	yearStr := c.QueryParam("year")
	cardNumber := c.QueryParam("card_number")

	year, err := strconv.Atoi(yearStr)
	if err != nil || year <= 0 {
		return errors.NewBadRequestError("invalid year parameter")
	}

	if cardNumber == "" {
		return errors.NewBadRequestError("card_number is required")
	}

	ctx := c.Request().Context()

	reqCache := &requests.MonthYearPaymentMethod{
		CardNumber: cardNumber,
		Year:       year,
	}

	cachedData, found := h.cache.GetYearlyAmountsByCardCache(ctx, reqCache)
	if found {
		return c.JSON(http.StatusOK, cachedData)
	}

	res, err := h.client.FindYearlyAmountsByCardNumber(ctx, &pbtransaction.FindByYearCardNumberTransactionRequest{
		CardNumber: cardNumber,
		Year:       int32(year),
	})
	if err != nil {
		h.logger.Debug("Failed to retrieve yearly amounts by card number", zap.Error(err))
		return errors.ParseGrpcError(err)
	}

	apiResponse := h.mapper.ToApiResponseTransactionYearAmount(res)
	h.cache.SetYearlyAmountsByCardCache(ctx, reqCache, apiResponse)

	return c.JSON(http.StatusOK, apiResponse)
}
