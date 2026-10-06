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

type withdrawStatsStatusHandleApi struct {
	client pb.WithdrawStatsStatusServiceClient

	logger logger.LoggerInterface

	mapper apimapper.WithdrawStatsStatusResponseMapper

	cache withdraw_cache.WithdrawMencache

	apiHandler errors.ApiHandler
}

type withdrawStatsStatusHandleDeps struct {
	client pb.WithdrawStatsStatusServiceClient

	router *echo.Echo

	logger logger.LoggerInterface

	mapper apimapper.WithdrawStatsStatusResponseMapper

	cache withdraw_cache.WithdrawMencache

	apiHandler errors.ApiHandler
}

func NewWithdrawStatsStatusHandleApi(params *withdrawStatsStatusHandleDeps) *withdrawStatsStatusHandleApi {
	withdrawStatsStatusHandleApi := &withdrawStatsStatusHandleApi{
		client:     params.client,
		logger:     params.logger,
		mapper:     params.mapper,
		cache:      params.cache,
		apiHandler: params.apiHandler,
	}

	routerWithdraw := params.router.Group("/api/withdraw-stats-status")

	routerWithdraw.GET("/monthly-success", params.apiHandler.Handle("find-monthly-withdraw-status-success", withdrawStatsStatusHandleApi.FindMonthlyWithdrawStatusSuccess))
	routerWithdraw.GET("/yearly-success", params.apiHandler.Handle("find-yearly-withdraw-status-success", withdrawStatsStatusHandleApi.FindYearlyWithdrawStatusSuccess))
	routerWithdraw.GET("/monthly-failed", params.apiHandler.Handle("find-monthly-withdraw-status-failed", withdrawStatsStatusHandleApi.FindMonthlyWithdrawStatusFailed))
	routerWithdraw.GET("/yearly-failed", params.apiHandler.Handle("find-yearly-withdraw-status-failed", withdrawStatsStatusHandleApi.FindYearlyWithdrawStatusFailed))

	routerWithdraw.GET("/monthly-success-by-card", params.apiHandler.Handle("find-monthly-withdraw-status-success-by-card", withdrawStatsStatusHandleApi.FindMonthlyWithdrawStatusSuccessByCardNumber))
	routerWithdraw.GET("/yearly-success-by-card", params.apiHandler.Handle("find-yearly-withdraw-status-success-by-card", withdrawStatsStatusHandleApi.FindYearlyWithdrawStatusSuccessByCardNumber))
	routerWithdraw.GET("/monthly-failed-by-card", params.apiHandler.Handle("find-monthly-withdraw-status-failed-by-card", withdrawStatsStatusHandleApi.FindMonthlyWithdrawStatusFailedByCardNumber))
	routerWithdraw.GET("/yearly-failed-by-card", params.apiHandler.Handle("find-yearly-withdraw-status-failed-by-card", withdrawStatsStatusHandleApi.FindYearlyWithdrawStatusFailedByCardNumber))

	return withdrawStatsStatusHandleApi
}

// FindMonthlyWithdrawStatusSuccess retrieves the monthly withdraw status for successful transactions.
// @Summary Get monthly withdraw status for successful transactions
// @Tags Withdraw Stats Withdraw
// @Security Bearer
// @Description Retrieve the monthly withdraw status for successful transactions by year and month.
// @Accept json
// @Produce json
// @Param year query int true "Year"
// @Param month query int true "Month"
// @Success 200 {object} response.ApiResponseWithdrawMonthStatusSuccess "Monthly withdraw status for successful transactions"
// @Failure 400 {object} response.ErrorResponse "Invalid year or month"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve monthly withdraw status for successful transactions"
// @Router /api/withdraw-stats-status/monthly-success [get]
func (h *withdrawStatsStatusHandleApi) FindMonthlyWithdrawStatusSuccess(c echo.Context) error {
	yearStr := c.QueryParam("year")
	monthStr := c.QueryParam("month")

	year, err := strconv.Atoi(yearStr)
	if err != nil || year <= 0 {
		return errors.NewBadRequestError("invalid year parameter")
	}

	month, err := strconv.Atoi(monthStr)
	if err != nil || month <= 0 || month > 12 {
		return errors.NewBadRequestError("invalid month parameter")
	}

	ctx := c.Request().Context()

	reqCache := &requests.MonthStatusWithdraw{
		Year:  year,
		Month: month,
	}

	cachedData, found := h.cache.GetCachedMonthWithdrawStatusSuccessCache(ctx, reqCache)
	if found {
		return c.JSON(http.StatusOK, cachedData)
	}

	res, err := h.client.FindMonthlyWithdrawStatusSuccess(ctx, &pbwithdraw.FindMonthlyWithdrawStatus{
		Year:  int32(year),
		Month: int32(month),
	})
	if err != nil {
		h.logger.Debug("Failed to retrieve monthly withdraw status success", zap.Error(err))
		return errors.ParseGrpcError(err)
	}

	apiResponse := h.mapper.ToApiResponseWithdrawMonthStatusSuccess(res)
	h.cache.SetCachedMonthWithdrawStatusSuccessCache(ctx, reqCache, apiResponse)

	return c.JSON(http.StatusOK, apiResponse)
}

// FindYearlyWithdrawStatusSuccess retrieves the yearly withdraw status for successful transactions.
// @Summary Get yearly withdraw status for successful transactions
// @Tags Withdraw Stats Withdraw
// @Security Bearer
// @Description Retrieve the yearly withdraw status for successful transactions by year.
// @Accept json
// @Produce json
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponseWithdrawYearStatusSuccess "Yearly withdraw status for successful transactions"
// @Failure 400 {object} response.ErrorResponse "Invalid year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve yearly withdraw status for successful transactions"
// @Router /api/withdraw-stats-status/yearly-success [get]
func (h *withdrawStatsStatusHandleApi) FindYearlyWithdrawStatusSuccess(c echo.Context) error {
	yearStr := c.QueryParam("year")

	year, err := strconv.Atoi(yearStr)
	if err != nil || year <= 0 {
		return errors.NewBadRequestError("invalid year parameter")
	}

	ctx := c.Request().Context()

	cachedData, found := h.cache.GetCachedYearlyWithdrawStatusSuccessCache(ctx, year)
	if found {
		return c.JSON(http.StatusOK, cachedData)
	}

	res, err := h.client.FindYearlyWithdrawStatusSuccess(ctx, &pbwithdraw.FindYearWithdrawStatus{
		Year: int32(year),
	})
	if err != nil {
		h.logger.Debug("Failed to retrieve yearly withdraw status success", zap.Error(err))
		return errors.ParseGrpcError(err)
	}

	apiResponse := h.mapper.ToApiResponseWithdrawYearStatusSuccess(res)
	h.cache.SetCachedYearlyWithdrawStatusSuccessCache(ctx, year, apiResponse)

	return c.JSON(http.StatusOK, apiResponse)
}

// FindMonthlyWithdrawStatusFailed retrieves the monthly withdraw status for failed transactions.
// @Summary Get monthly withdraw status for failed transactions
// @Tags Withdraw Stats Withdraw
// @Security Bearer
// @Description Retrieve the monthly withdraw status for failed transactions by year and month.
// @Accept json
// @Produce json
// @Param year query int true "Year"
// @Param month query int true "Month"
// @Success 200 {object} response.ApiResponseWithdrawMonthStatusFailed "Monthly withdraw status for failed transactions"
// @Failure 400 {object} response.ErrorResponse "Invalid year or month"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve monthly withdraw status for failed transactions"
// @Router /api/withdraw-stats-status/monthly-failed [get]
func (h *withdrawStatsStatusHandleApi) FindMonthlyWithdrawStatusFailed(c echo.Context) error {
	yearStr := c.QueryParam("year")
	monthStr := c.QueryParam("month")

	year, err := strconv.Atoi(yearStr)
	if err != nil || year <= 0 {
		return errors.NewBadRequestError("invalid year parameter")
	}

	month, err := strconv.Atoi(monthStr)
	if err != nil || month <= 0 || month > 12 {
		return errors.NewBadRequestError("invalid month parameter")
	}

	ctx := c.Request().Context()

	reqCache := &requests.MonthStatusWithdraw{
		Year:  year,
		Month: month,
	}

	cachedData, found := h.cache.GetCachedMonthWithdrawStatusFailedCache(ctx, reqCache)
	if found {
		return c.JSON(http.StatusOK, cachedData)
	}

	res, err := h.client.FindMonthlyWithdrawStatusFailed(ctx, &pbwithdraw.FindMonthlyWithdrawStatus{
		Year:  int32(year),
		Month: int32(month),
	})
	if err != nil {
		h.logger.Debug("Failed to retrieve monthly withdraw status failed", zap.Error(err))
		return errors.ParseGrpcError(err)
	}

	apiResponse := h.mapper.ToApiResponseWithdrawMonthStatusFailed(res)
	h.cache.SetCachedMonthWithdrawStatusFailedCache(ctx, reqCache, apiResponse)

	return c.JSON(http.StatusOK, apiResponse)
}

// FindYearlyWithdrawStatusFailed retrieves the yearly withdraw status for failed transactions.
// @Summary Get yearly withdraw status for failed transactions
// @Tags Withdraw Stats Withdraw
// @Security Bearer
// @Description Retrieve the yearly withdraw status for failed transactions by year.
// @Accept json
// @Produce json
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponseWithdrawYearStatusSuccess "Yearly withdraw status for failed transactions"
// @Failure 400 {object} response.ErrorResponse "Invalid year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve yearly withdraw status for failed transactions"
// @Router /api/withdraw-stats-status/yearly-failed [get]
func (h *withdrawStatsStatusHandleApi) FindYearlyWithdrawStatusFailed(c echo.Context) error {
	yearStr := c.QueryParam("year")

	year, err := strconv.Atoi(yearStr)
	if err != nil || year <= 0 {
		return errors.NewBadRequestError("invalid year parameter")
	}

	ctx := c.Request().Context()

	cachedData, found := h.cache.GetCachedYearlyWithdrawStatusFailedCache(ctx, year)
	if found {
		return c.JSON(http.StatusOK, cachedData)
	}

	res, err := h.client.FindYearlyWithdrawStatusFailed(ctx, &pbwithdraw.FindYearWithdrawStatus{
		Year: int32(year),
	})
	if err != nil {
		h.logger.Debug("Failed to retrieve yearly withdraw status failed", zap.Error(err))
		return errors.ParseGrpcError(err)
	}

	apiResponse := h.mapper.ToApiResponseWithdrawYearStatusFailed(res)
	h.cache.SetCachedYearlyWithdrawStatusFailedCache(ctx, year, apiResponse)

	return c.JSON(http.StatusOK, apiResponse)
}

// FindMonthlyWithdrawStatusSuccessByCardNumber retrieves the monthly withdraw status for successful transactions.
// @Summary Get monthly withdraw status for successful transactions
// @Tags Withdraw Stats Withdraw
// @Security Bearer
// @Description Retrieve the monthly withdraw status for successful transactions by year and month.
// @Accept json
// @Produce json
// @Param year query int true "Year"
// @Param month query int true "Month"
// @Param card_number query string true "Card Number"
// @Success 200 {object} response.ApiResponseWithdrawMonthStatusSuccess "Monthly withdraw status for successful transactions"
// @Failure 400 {object} response.ErrorResponse "Invalid year or month"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve monthly withdraw status for successful transactions"
// @Router /api/withdraw-stats-status/monthly-success-by-card [get]
func (h *withdrawStatsStatusHandleApi) FindMonthlyWithdrawStatusSuccessByCardNumber(c echo.Context) error {
	yearStr := c.QueryParam("year")
	monthStr := c.QueryParam("month")
	cardNumber := c.QueryParam("card_number")

	if cardNumber == "" {
		return errors.NewBadRequestError("invalid card_number parameter")
	}

	year, err := strconv.Atoi(yearStr)
	if err != nil || year <= 0 {
		return errors.NewBadRequestError("invalid year parameter")
	}

	month, err := strconv.Atoi(monthStr)
	if err != nil || month <= 0 || month > 12 {
		return errors.NewBadRequestError("invalid month parameter")
	}

	ctx := c.Request().Context()

	reqCache := &requests.MonthStatusWithdrawCardNumber{
		CardNumber: cardNumber,
		Year:       year,
		Month:      month,
	}

	cachedData, found := h.cache.GetCachedMonthWithdrawStatusSuccessByCardNumber(ctx, reqCache)
	if found {
		return c.JSON(http.StatusOK, cachedData)
	}

	res, err := h.client.FindMonthlyWithdrawStatusSuccessCardNumber(ctx, &pbwithdraw.FindMonthlyWithdrawStatusCardNumber{
		Year:       int32(year),
		Month:      int32(month),
		CardNumber: cardNumber,
	})
	if err != nil {
		h.logger.Debug("Failed to retrieve monthly withdraw status success", zap.Error(err))
		return errors.ParseGrpcError(err)
	}

	apiResponse := h.mapper.ToApiResponseWithdrawMonthStatusSuccess(res)
	h.cache.SetCachedMonthWithdrawStatusSuccessByCardNumber(ctx, reqCache, apiResponse)

	return c.JSON(http.StatusOK, apiResponse)
}

// FindYearlyWithdrawStatusSuccessByCardNumber retrieves the yearly withdraw status for successful transactions.
// @Summary Get yearly withdraw status for successful transactions
// @Tags Withdraw Stats Withdraw
// @Security Bearer
// @Description Retrieve the yearly withdraw status for successful transactions by year.
// @Accept json
// @Produce json
// @Param year query int true "Year"
// @Param card_number query string true "Card Number"
// @Success 200 {object} response.ApiResponseWithdrawYearStatusSuccess "Yearly withdraw status for successful transactions"
// @Failure 400 {object} response.ErrorResponse "Invalid year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve yearly withdraw status for successful transactions"
// @Router /api/withdraw-stats-status/yearly-success-by-card-number [get]
func (h *withdrawStatsStatusHandleApi) FindYearlyWithdrawStatusSuccessByCardNumber(c echo.Context) error {
	yearStr := c.QueryParam("year")
	cardNumber := c.QueryParam("card_number")

	if cardNumber == "" {
		return errors.NewBadRequestError("invalid card_number parameter")
	}

	year, err := strconv.Atoi(yearStr)
	if err != nil || year <= 0 {
		return errors.NewBadRequestError("invalid year parameter")
	}

	ctx := c.Request().Context()

	reqCache := &requests.YearStatusWithdrawCardNumber{
		CardNumber: cardNumber,
		Year:       year,
	}

	cachedData, found := h.cache.GetCachedYearlyWithdrawStatusSuccessByCardNumber(ctx, reqCache)
	if found {
		return c.JSON(http.StatusOK, cachedData)
	}

	res, err := h.client.FindYearlyWithdrawStatusSuccessCardNumber(ctx, &pbwithdraw.FindYearWithdrawStatusCardNumber{
		CardNumber: cardNumber,
		Year:       int32(year),
	})
	if err != nil {
		h.logger.Debug("Failed to retrieve yearly withdraw status success", zap.Error(err))
		return errors.ParseGrpcError(err)
	}

	apiResponse := h.mapper.ToApiResponseWithdrawYearStatusSuccess(res)
	h.cache.SetCachedYearlyWithdrawStatusSuccessByCardNumber(ctx, reqCache, apiResponse)

	return c.JSON(http.StatusOK, apiResponse)
}

// FindMonthlyWithdrawStatusFailedByCardNumber retrieves the monthly withdraw status for failed transactions.
// @Summary Get monthly withdraw status for failed transactions
// @Tags Withdraw Stats Withdraw
// @Security Bearer
// @Description Retrieve the monthly withdraw status for failed transactions by year and month.
// @Accept json
// @Produce json
// @Param year query int true "Year"
// @Param month query int true "Month"
// @Param card_number query string true "Card Number"
// @Success 200 {object} response.ApiResponseWithdrawMonthStatusFailed "Monthly withdraw status for failed transactions"
// @Failure 400 {object} response.ErrorResponse "Invalid year or month"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve monthly withdraw status for failed transactions"
// @Router /api/withdraw-stats-status/monthly-failed-by-card [get]
func (h *withdrawStatsStatusHandleApi) FindMonthlyWithdrawStatusFailedByCardNumber(c echo.Context) error {
	yearStr := c.QueryParam("year")
	monthStr := c.QueryParam("month")
	cardNumber := c.QueryParam("card_number")

	if cardNumber == "" {
		return errors.NewBadRequestError("invalid card_number parameter")
	}

	year, err := strconv.Atoi(yearStr)
	if err != nil || year <= 0 {
		return errors.NewBadRequestError("invalid year parameter")
	}

	month, err := strconv.Atoi(monthStr)
	if err != nil || month <= 0 || month > 12 {
		return errors.NewBadRequestError("invalid month parameter")
	}

	ctx := c.Request().Context()

	reqCache := &requests.MonthStatusWithdrawCardNumber{
		CardNumber: cardNumber,
		Year:       year,
		Month:      month,
	}

	cachedData, found := h.cache.GetCachedMonthWithdrawStatusFailedByCardNumber(ctx, reqCache)
	if found {
		return c.JSON(http.StatusOK, cachedData)
	}

	res, err := h.client.FindMonthlyWithdrawStatusFailedCardNumber(ctx, &pbwithdraw.FindMonthlyWithdrawStatusCardNumber{
		Year:       int32(year),
		Month:      int32(month),
		CardNumber: cardNumber,
	})
	if err != nil {
		h.logger.Debug("Failed to retrieve monthly withdraw status failed", zap.Error(err))
		return errors.ParseGrpcError(err)
	}

	apiResponse := h.mapper.ToApiResponseWithdrawMonthStatusFailed(res)
	h.cache.SetCachedMonthWithdrawStatusFailedByCardNumber(ctx, reqCache, apiResponse)

	return c.JSON(http.StatusOK, apiResponse)
}

// FindYearlyWithdrawStatusFailedByCardNumber retrieves the yearly withdraw status for failed transactions.
// @Summary Get yearly withdraw status for failed transactions
// @Tags Withdraw Stats Withdraw
// @Security Bearer
// @Description Retrieve the yearly withdraw status for failed transactions by year.
// @Accept json
// @Produce json
// @Param year query int true "Year"
// @Param card_number query string true "Card Number"
// @Success 200 {object} response.ApiResponseWithdrawYearStatusSuccess "Yearly withdraw status for failed transactions"
// @Failure 400 {object} response.ErrorResponse "Invalid year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve yearly withdraw status for failed transactions"
// @Router /api/withdraw-stats-status/yearly-failed-by-card [get]
func (h *withdrawStatsStatusHandleApi) FindYearlyWithdrawStatusFailedByCardNumber(c echo.Context) error {
	yearStr := c.QueryParam("year")
	cardNumber := c.QueryParam("card_number")

	if cardNumber == "" {
		return errors.NewBadRequestError("invalid card_number parameter")
	}

	year, err := strconv.Atoi(yearStr)
	if err != nil || year <= 0 {
		return errors.NewBadRequestError("invalid year parameter")
	}

	ctx := c.Request().Context()

	reqCache := &requests.YearStatusWithdrawCardNumber{
		CardNumber: cardNumber,
		Year:       year,
	}

	cachedData, found := h.cache.GetCachedYearlyWithdrawStatusFailedByCardNumber(ctx, reqCache)
	if found {
		return c.JSON(http.StatusOK, cachedData)
	}

	res, err := h.client.FindYearlyWithdrawStatusFailedCardNumber(ctx, &pbwithdraw.FindYearWithdrawStatusCardNumber{
		CardNumber: cardNumber,
		Year:       int32(year),
	})
	if err != nil {
		h.logger.Debug("Failed to retrieve yearly withdraw status failed", zap.Error(err))
		return errors.ParseGrpcError(err)
	}

	apiResponse := h.mapper.ToApiResponseWithdrawYearStatusFailed(res)
	h.cache.SetCachedYearlyWithdrawStatusFailedByCardNumber(ctx, reqCache, apiResponse)

	return c.JSON(http.StatusOK, apiResponse)
}
