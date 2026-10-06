package withdrawhandler

import (
	"net/http"
	"strconv"

	withdraw_cache "github.com/MamangRust/monolith-payment-gateway-apigateway/redis/api/withdraw"
	pbwithdraw "github.com/MamangRust/monolith-payment-gateway-pb/withdraw"
	pb "github.com/MamangRust/monolith-payment-gateway-pb/withdraw/stats"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
	apimapper "github.com/MamangRust/monolith-payment-gateway-shared/mapper/withdraw"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type withdrawStatsAmountHandleApi struct {
	client pb.WithdrawStatsAmountServiceClient

	logger logger.LoggerInterface

	mapper apimapper.WithdrawStatsAmountResponseMapper

	cache withdraw_cache.WithdrawMencache

	apiHandler errors.ApiHandler
}

type withdrawStatsAmountHandleDeps struct {
	client pb.WithdrawStatsAmountServiceClient

	router *echo.Echo

	logger logger.LoggerInterface

	mapper apimapper.WithdrawStatsAmountResponseMapper

	cache withdraw_cache.WithdrawMencache

	apiHandler errors.ApiHandler
}

func NewWithdrawStatsAmountHandleApi(params *withdrawStatsAmountHandleDeps) *withdrawStatsAmountHandleApi {

	withdrawStatsAmountHandleApi := &withdrawStatsAmountHandleApi{
		client:     params.client,
		logger:     params.logger,
		mapper:     params.mapper,
		cache:      params.cache,
		apiHandler: params.apiHandler,
	}

	routerWithdraw := params.router.Group("/api/withdraw-stats-amount")

	routerWithdraw.GET("/monthly-amount", params.apiHandler.Handle("find-monthly-withdraw-amounts", withdrawStatsAmountHandleApi.FindMonthlyWithdraws))
	routerWithdraw.GET("/yearly-amount", params.apiHandler.Handle("find-yearly-withdraw-amounts", withdrawStatsAmountHandleApi.FindYearlyWithdraws))

	routerWithdraw.GET("/monthly-amount-card", params.apiHandler.Handle("find-monthly-withdraw-amounts-by-card", withdrawStatsAmountHandleApi.FindMonthlyWithdrawsByCardNumber))
	routerWithdraw.GET("/yearly-amount-card", params.apiHandler.Handle("find-yearly-withdraw-amounts-by-card", withdrawStatsAmountHandleApi.FindYearlyWithdrawsByCardNumber))

	return withdrawStatsAmountHandleApi
}

// FindMonthlyWithdraws retrieves the monthly withdraws for a specific year.
// @Summary Get monthly withdraws
// @Tags Withdraw Stats Amount
// @Security Bearer
// @Description Retrieve the monthly withdraws for a specific year.
// @Accept json
// @Produce json
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponseWithdrawMonthAmount "Monthly withdraws"
// @Failure 400 {object} response.ErrorResponse "Invalid year parameter"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve monthly withdraws"
// @Router /api/withdraw-stats-amount/monthly [get]
func (h *withdrawStatsAmountHandleApi) FindMonthlyWithdraws(c echo.Context) error {
	yearStr := c.QueryParam("year")
	year, err := strconv.Atoi(yearStr)
	if err != nil || year <= 0 {
		return errors.NewBadRequestError("invalid year parameter")
	}

	ctx := c.Request().Context()

	cachedData, found := h.cache.GetCachedMonthlyWithdraws(ctx, year)
	if found {
		return c.JSON(http.StatusOK, cachedData)
	}

	res, err := h.client.FindMonthlyWithdraws(ctx, &pbwithdraw.FindYearWithdrawStatus{
		Year: int32(year),
	})
	if err != nil {
		h.logger.Debug("Failed to retrieve monthly withdraws", zap.Error(err))
		return errors.ParseGrpcError(err)
	}

	apiResponse := h.mapper.ToApiResponseWithdrawMonthAmount(res)
	h.cache.SetCachedMonthlyWithdraws(ctx, year, apiResponse)

	return c.JSON(http.StatusOK, apiResponse)
}

// FindYearlyWithdraws retrieves the yearly withdraws for a specific year.
// @Summary Get yearly withdraws
// @Tags Withdraw Stats Amount
// @Security Bearer
// @Description Retrieve the yearly withdraws for a specific year.
// @Accept json
// @Produce json
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponseWithdrawYearAmount "Yearly withdraws"
// @Failure 400 {object} response.ErrorResponse "Invalid year parameter"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve yearly withdraws"
// @Router /api/withdraw-stats-amount/yearly [get]
func (h *withdrawStatsAmountHandleApi) FindYearlyWithdraws(c echo.Context) error {
	yearStr := c.QueryParam("year")
	year, err := strconv.Atoi(yearStr)
	if err != nil || year <= 0 {
		return errors.NewBadRequestError("invalid year parameter")
	}

	ctx := c.Request().Context()

	cachedData, found := h.cache.GetCachedYearlyWithdraws(ctx, year)
	if found {
		return c.JSON(http.StatusOK, cachedData)
	}

	res, err := h.client.FindYearlyWithdraws(ctx, &pbwithdraw.FindYearWithdrawStatus{
		Year: int32(year),
	})
	if err != nil {
		h.logger.Debug("Failed to retrieve yearly withdraws", zap.Error(err))
		return errors.ParseGrpcError(err)
	}

	apiResponse := h.mapper.ToApiResponseWithdrawYearAmount(res)
	h.cache.SetCachedYearlyWithdraws(ctx, year, apiResponse)

	return c.JSON(http.StatusOK, apiResponse)
}

// FindMonthlyWithdrawsByCardNumber retrieves the monthly withdraws for a specific card number and year.
// @Summary Get monthly withdraws by card number
// @Tags Withdraw Stats Amount
// @Security Bearer
// @Description Retrieve the monthly withdraws for a specific card number and year.
// @Accept json
// @Produce json
// @Param card_number query string true "Card Number"
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponseWithdrawMonthAmount "Monthly withdraws by card number"
// @Failure 400 {object} response.ErrorResponse "Invalid card number or year parameter"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve monthly withdraws by card number"
// @Router /api/withdraw-stats-amount/monthly-by-card [get]
func (h *withdrawStatsAmountHandleApi) FindMonthlyWithdrawsByCardNumber(c echo.Context) error {
	cardNumber := c.QueryParam("card_number")
	yearStr := c.QueryParam("year")

	if cardNumber == "" {
		return errors.NewBadRequestError("invalid card_number parameter")
	}

	year, err := strconv.Atoi(yearStr)
	if err != nil || year <= 0 {
		return errors.NewBadRequestError("invalid year parameter")
	}

	ctx := c.Request().Context()

	reqCache := &requests.YearMonthCardNumber{
		CardNumber: cardNumber,
		Year:       year,
	}

	cachedData, found := h.cache.GetCachedMonthlyWithdrawsByCardNumber(ctx, reqCache)
	if found {
		return c.JSON(http.StatusOK, cachedData)
	}

	res, err := h.client.FindMonthlyWithdrawsByCardNumber(ctx, &pbwithdraw.FindYearWithdrawCardNumber{
		CardNumber: cardNumber,
		Year:       int32(year),
	})
	if err != nil {
		h.logger.Debug("Failed to retrieve monthly withdraws by card number", zap.Error(err))
		return errors.ParseGrpcError(err)
	}

	apiResponse := h.mapper.ToApiResponseWithdrawMonthAmount(res)
	h.cache.SetCachedMonthlyWithdrawsByCardNumber(ctx, reqCache, apiResponse)

	return c.JSON(http.StatusOK, apiResponse)
}

// FindYearlyWithdrawsByCardNumber retrieves the yearly withdraws for a specific card number and year.
// @Summary Get yearly withdraws by card number
// @Tags Withdraw Stats Amount
// @Security Bearer
// @Description Retrieve the yearly withdraws for a specific card number and year.
// @Accept json
// @Produce json
// @Param card_number query string true "Card Number"
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponseWithdrawYearAmount "Yearly withdraws by card number"
// @Failure 400 {object} response.ErrorResponse "Invalid card number or year parameter"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve yearly withdraws by card number"
// @Router /api/withdraw-stats-amount/yearly-by-card [get]
func (h *withdrawStatsAmountHandleApi) FindYearlyWithdrawsByCardNumber(c echo.Context) error {
	cardNumber := c.QueryParam("card_number")
	yearStr := c.QueryParam("year")

	if cardNumber == "" {
		return errors.NewBadRequestError("invalid card_number parameter")
	}

	year, err := strconv.Atoi(yearStr)
	if err != nil || year <= 0 {
		return errors.NewBadRequestError("invalid year parameter")
	}

	ctx := c.Request().Context()

	reqCache := &requests.YearMonthCardNumber{
		CardNumber: cardNumber,
		Year:       year,
	}

	cachedData, found := h.cache.GetCachedYearlyWithdrawsByCardNumber(ctx, reqCache)
	if found {
		return c.JSON(http.StatusOK, cachedData)
	}

	res, err := h.client.FindYearlyWithdrawsByCardNumber(ctx, &pbwithdraw.FindYearWithdrawCardNumber{
		CardNumber: cardNumber,
		Year:       int32(year),
	})
	if err != nil {
		h.logger.Debug("Failed to retrieve yearly withdraws by card number", zap.Error(err))
		return errors.ParseGrpcError(err)
	}

	apiResponse := h.mapper.ToApiResponseWithdrawYearAmount(res)
	h.cache.SetCachedYearlyWithdrawsByCardNumber(ctx, reqCache, apiResponse)

	return c.JSON(http.StatusOK, apiResponse)
}
