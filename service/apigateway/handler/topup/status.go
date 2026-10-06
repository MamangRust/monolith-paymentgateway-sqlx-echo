package topuphandler

import (
	"net/http"
	"strconv"

	topup_cache "github.com/MamangRust/monolith-payment-gateway-apigateway/redis/api/topup"
	pbtopup "github.com/MamangRust/monolith-payment-gateway-pb/topup"
	pb "github.com/MamangRust/monolith-payment-gateway-pb/topup/stats"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
	apimapper "github.com/MamangRust/monolith-payment-gateway-shared/mapper/topup"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type topupStatsStatusHandleApi struct {
	client pb.TopupStatsStatusServiceClient

	logger logger.LoggerInterface

	mapper apimapper.TopupStatsStatusResponseMapper

	cache topup_cache.TopupMencach

	apiHandler errors.ApiHandler
}

type topupStatsStatusHandleDeps struct {
	client pb.TopupStatsStatusServiceClient

	router *echo.Echo

	logger logger.LoggerInterface

	mapper apimapper.TopupStatsStatusResponseMapper

	cache topup_cache.TopupMencach

	apiHandler errors.ApiHandler
}

func NewTopupStatsStatusHandleApi(params *topupStatsStatusHandleDeps) *topupStatsStatusHandleApi {

	topupHandler := &topupStatsStatusHandleApi{
		client:     params.client,
		logger:     params.logger,
		mapper:     params.mapper,
		cache:      params.cache,
		apiHandler: params.apiHandler,
	}

	routerTopup := params.router.Group("/api/topup-stats-status")

	routerTopup.GET("/monthly-success", params.apiHandler.Handle("find-monthly-topup-status-success", topupHandler.FindMonthlyTopupStatusSuccess))
	routerTopup.GET("/yearly-success", params.apiHandler.Handle("find-yearly-topup-status-success", topupHandler.FindYearlyTopupStatusSuccess))
	routerTopup.GET("/monthly-failed", params.apiHandler.Handle("find-monthly-topup-status-failed", topupHandler.FindMonthlyTopupStatusFailed))
	routerTopup.GET("/yearly-failed", params.apiHandler.Handle("find-yearly-topup-status-failed", topupHandler.FindYearlyTopupStatusFailed))

	routerTopup.GET("/monthly-success-by-card", params.apiHandler.Handle("find-monthly-topup-status-success-by-card", topupHandler.FindMonthlyTopupStatusSuccessByCardNumber))
	routerTopup.GET("/yearly-success-by-card", params.apiHandler.Handle("find-yearly-topup-status-success-by-card", topupHandler.FindYearlyTopupStatusSuccessByCardNumber))
	routerTopup.GET("/monthly-failed-by-card", params.apiHandler.Handle("find-monthly-topup-status-failed-by-card", topupHandler.FindMonthlyTopupStatusFailedByCardNumber))
	routerTopup.GET("/yearly-failed-by-card", params.apiHandler.Handle("find-yearly-topup-status-failed-by-card", topupHandler.FindYearlyTopupStatusFailedByCardNumber))

	return topupHandler
}

// FindMonthlyTopupStatusSuccess retrieves the monthly top-up status for successful transactions.
// @Summary Get monthly top-up status for successful transactions
// @Tags Topup Stats Status
// @Security Bearer
// @Description Retrieve the monthly top-up status for successful transactions by year and month.
// @Accept json
// @Produce json
// @Param year query int true "Year"
// @Param month query int true "Month"
// @Success 200 {object} response.ApiResponseTopupMonthStatusSuccess "Monthly top-up status for successful transactions"
// @Failure 400 {object} response.ErrorResponse "Invalid year or month"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve monthly top-up status for successful transactions"
// @Router /api/topup-stats-status/monthly-success [get]
func (h *topupStatsStatusHandleApi) FindMonthlyTopupStatusSuccess(c echo.Context) error {
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

	reqCache := &requests.MonthTopupStatus{
		Year:  year,
		Month: month,
	}

	cachedData, found := h.cache.GetMonthTopupStatusSuccessCache(ctx, reqCache)
	if found {
		return c.JSON(http.StatusOK, cachedData)
	}

	res, err := h.client.FindMonthlyTopupStatusSuccess(ctx, &pbtopup.FindMonthlyTopupStatus{
		Year:  int32(year),
		Month: int32(month),
	})
	if err != nil {
		h.logger.Debug("Failed to retrieve monthly topup status success", zap.Error(err))
		return errors.ParseGrpcError(err)
	}

	apiResponse := h.mapper.ToApiResponseTopupMonthStatusSuccess(res)
	h.cache.SetMonthTopupStatusSuccessCache(ctx, reqCache, apiResponse)

	return c.JSON(http.StatusOK, apiResponse)
}

// FindYearlyTopupStatusSuccess retrieves the yearly top-up status for successful transactions.
// @Summary Get yearly top-up status for successful transactions
// @Tags Topup Stats Status
// @Security Bearer
// @Description Retrieve the yearly top-up status for successful transactions by year.
// @Accept json
// @Produce json
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponseTopupYearStatusSuccess "Yearly top-up status for successful transactions"
// @Failure 400 {object} response.ErrorResponse "Invalid year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve yearly top-up status for successful transactions"
// @Router /api/topup-stats-status/yearly-success [get]
func (h *topupStatsStatusHandleApi) FindYearlyTopupStatusSuccess(c echo.Context) error {
	yearStr := c.QueryParam("year")

	year, err := strconv.Atoi(yearStr)
	if err != nil || year <= 0 {
		return errors.NewBadRequestError("invalid year parameter")
	}

	ctx := c.Request().Context()

	cachedData, found := h.cache.GetYearlyTopupStatusSuccessCache(ctx, year)
	if found {
		return c.JSON(http.StatusOK, cachedData)
	}

	res, err := h.client.FindYearlyTopupStatusSuccess(ctx, &pbtopup.FindYearTopupStatus{
		Year: int32(year),
	})
	if err != nil {
		h.logger.Debug("Failed to retrieve yearly topup status success", zap.Error(err))
		return errors.ParseGrpcError(err)
	}

	apiResponse := h.mapper.ToApiResponseTopupYearStatusSuccess(res)
	h.cache.SetYearlyTopupStatusSuccessCache(ctx, year, apiResponse)

	return c.JSON(http.StatusOK, apiResponse)
}

// FindMonthlyTopupStatusFailed retrieves the monthly top-up status for failed transactions.
// @Summary Get monthly top-up status for failed transactions
// @Tags Topup Stats Status
// @Security Bearer
// @Description Retrieve the monthly top-up status for failed transactions by year and month.
// @Accept json
// @Produce json
// @Param year query int true "Year"
// @Param month query int true "Month"
// @Success 200 {object} response.ApiResponseTopupMonthStatusFailed "Monthly top-up status for failed transactions"
// @Failure 400 {object} response.ErrorResponse "Invalid year or month"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve monthly top-up status for failed transactions"
// @Router /api/topup-stats-status/monthly-failed [get]
func (h *topupStatsStatusHandleApi) FindMonthlyTopupStatusFailed(c echo.Context) error {
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

	reqCache := &requests.MonthTopupStatus{
		Year:  year,
		Month: month,
	}

	cachedData, found := h.cache.GetMonthTopupStatusFailedCache(ctx, reqCache)
	if found {
		return c.JSON(http.StatusOK, cachedData)
	}

	res, err := h.client.FindMonthlyTopupStatusFailed(ctx, &pbtopup.FindMonthlyTopupStatus{
		Year:  int32(year),
		Month: int32(month),
	})
	if err != nil {
		h.logger.Debug("Failed to retrieve monthly topup status failed", zap.Error(err))
		return errors.ParseGrpcError(err)
	}

	apiResponse := h.mapper.ToApiResponseTopupMonthStatusFailed(res)
	h.cache.SetMonthTopupStatusFailedCache(ctx, reqCache, apiResponse)

	return c.JSON(http.StatusOK, apiResponse)
}

// FindYearlyTopupStatusFailed retrieves the yearly top-up status for failed transactions.
// @Summary Get yearly top-up status for failed transactions
// @Tags Topup Stats Status
// @Security Bearer
// @Description Retrieve the yearly top-up status for failed transactions by year.
// @Accept json
// @Produce json
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponseTopupYearStatusFailed "Yearly top-up status for failed transactions"
// @Failure 400 {object} response.ErrorResponse "Invalid year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve yearly top-up status for failed transactions"
// @Router /api/topup-stats-status/yearly-failed [get]
func (h *topupStatsStatusHandleApi) FindYearlyTopupStatusFailed(c echo.Context) error {
	yearStr := c.QueryParam("year")

	year, err := strconv.Atoi(yearStr)
	if err != nil || year <= 0 {
		return errors.NewBadRequestError("invalid year parameter")
	}

	ctx := c.Request().Context()

	cachedData, found := h.cache.GetYearlyTopupStatusFailedCache(ctx, year)
	if found {
		return c.JSON(http.StatusOK, cachedData)
	}

	res, err := h.client.FindYearlyTopupStatusFailed(ctx, &pbtopup.FindYearTopupStatus{
		Year: int32(year),
	})
	if err != nil {
		h.logger.Debug("Failed to retrieve yearly topup status failed", zap.Error(err))
		return errors.ParseGrpcError(err)
	}

	apiResponse := h.mapper.ToApiResponseTopupYearStatusFailed(res)
	h.cache.SetYearlyTopupStatusFailedCache(ctx, year, apiResponse)

	return c.JSON(http.StatusOK, apiResponse)
}

// FindMonthlyTopupStatusSuccess retrieves the monthly top-up status for successful transactions.
// @Summary Get monthly top-up status for successful transactions
// @Tags Topup Stats Status
// @Security Bearer
// @Description Retrieve the monthly top-up status for successful transactions by year and month.
// @Accept json
// @Produce json
// @Param year query int true "Year"
// @Param month query int true "Month"
// @Param card_number query string true "Card Number"
// @Success 200 {object} response.ApiResponseTopupMonthStatusSuccess "Monthly top-up status for successful transactions"
// @Failure 400 {object} response.ErrorResponse "Invalid year or month"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve monthly top-up status for successful transactions"
// @Router /api/topup-stats-status/monthly-success [get]
func (h *topupStatsStatusHandleApi) FindMonthlyTopupStatusSuccessByCardNumber(c echo.Context) error {
	yearStr := c.QueryParam("year")
	monthStr := c.QueryParam("month")
	cardNumber := c.QueryParam("card_number")

	year, err := strconv.Atoi(yearStr)
	if err != nil || year <= 0 {
		return errors.NewBadRequestError("invalid year parameter")
	}

	month, err := strconv.Atoi(monthStr)
	if err != nil || month <= 0 || month > 12 {
		return errors.NewBadRequestError("invalid month parameter")
	}

	if cardNumber == "" {
		return errors.NewBadRequestError("card_number is required")
	}

	ctx := c.Request().Context()

	reqCache := &requests.MonthTopupStatusCardNumber{
		CardNumber: cardNumber,
		Year:       year,
		Month:      month,
	}

	cachedData, found := h.cache.GetMonthTopupStatusSuccessByCardNumberCache(ctx, reqCache)
	if found {
		return c.JSON(http.StatusOK, cachedData)
	}

	res, err := h.client.FindMonthlyTopupStatusSuccessByCardNumber(ctx, &pbtopup.FindMonthlyTopupStatusCardNumber{
		Year:       int32(year),
		Month:      int32(month),
		CardNumber: cardNumber,
	})
	if err != nil {
		h.logger.Debug("Failed to retrieve monthly topup status success", zap.Error(err))
		return errors.ParseGrpcError(err)
	}

	apiResponse := h.mapper.ToApiResponseTopupMonthStatusSuccess(res)
	h.cache.SetMonthTopupStatusSuccessByCardNumberCache(ctx, reqCache, apiResponse)

	return c.JSON(http.StatusOK, apiResponse)
}

// FindYearlyTopupStatusSuccess retrieves the yearly top-up status for successful transactions.
// @Summary Get yearly top-up status for successful transactions
// @Tags Topup Stats Status
// @Security Bearer
// @Description Retrieve the yearly top-up status for successful transactions by year.
// @Accept json
// @Produce json
// @Param year query int true "Year"
// @Param card_number query string true "Card Number"
// @Success 200 {object} response.ApiResponseTopupYearStatusSuccess "Yearly top-up status for successful transactions"
// @Failure 400 {object} response.ErrorResponse "Invalid year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve yearly top-up status for successful transactions"
// @Router /api/topup-stats-status/yearly-success [get]
func (h *topupStatsStatusHandleApi) FindYearlyTopupStatusSuccessByCardNumber(c echo.Context) error {
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

	reqCache := &requests.YearTopupStatusCardNumber{
		CardNumber: cardNumber,
		Year:       year,
	}

	cachedData, found := h.cache.GetYearlyTopupStatusSuccessByCardNumberCache(ctx, reqCache)
	if found {
		return c.JSON(http.StatusOK, cachedData)
	}

	res, err := h.client.FindYearlyTopupStatusSuccessByCardNumber(ctx, &pbtopup.FindYearTopupStatusCardNumber{
		Year:       int32(year),
		CardNumber: cardNumber,
	})
	if err != nil {
		h.logger.Debug("Failed to retrieve yearly topup status success", zap.Error(err))
		return errors.ParseGrpcError(err)
	}

	apiResponse := h.mapper.ToApiResponseTopupYearStatusSuccess(res)
	h.cache.SetYearlyTopupStatusSuccessByCardNumberCache(ctx, reqCache, apiResponse)

	return c.JSON(http.StatusOK, apiResponse)
}

// FindMonthlyTopupStatusFailed retrieves the monthly top-up status for failed transactions.
// @Summary Get monthly top-up status for failed transactions
// @Tags Topup Stats Status
// @Security Bearer
// @Description Retrieve the monthly top-up status for failed transactions by year and month.
// @Accept json
// @Produce json
// @Param year query int true "Year"
// @Param month query int true "Month"
// @Param card_number query string true "Card Number"
// @Success 200 {object} response.ApiResponseTopupMonthStatusFailed "Monthly top-up status for failed transactions"
// @Failure 400 {object} response.ErrorResponse "Invalid year or month"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve monthly top-up status for failed transactions"
// @Router /api/topup-stats-status/monthly-failed [get]
func (h *topupStatsStatusHandleApi) FindMonthlyTopupStatusFailedByCardNumber(c echo.Context) error {
	yearStr := c.QueryParam("year")
	monthStr := c.QueryParam("month")
	cardNumber := c.QueryParam("card_number")

	year, err := strconv.Atoi(yearStr)
	if err != nil || year <= 0 {
		return errors.NewBadRequestError("invalid year parameter")
	}

	month, err := strconv.Atoi(monthStr)
	if err != nil || month <= 0 || month > 12 {
		return errors.NewBadRequestError("invalid month parameter")
	}

	if cardNumber == "" {
		return errors.NewBadRequestError("card_number is required")
	}

	ctx := c.Request().Context()

	reqCache := &requests.MonthTopupStatusCardNumber{
		CardNumber: cardNumber,
		Year:       year,
		Month:      month,
	}

	cachedData, found := h.cache.GetMonthTopupStatusFailedByCardNumberCache(ctx, reqCache)
	if found {
		return c.JSON(http.StatusOK, cachedData)
	}

	res, err := h.client.FindMonthlyTopupStatusFailedByCardNumber(ctx, &pbtopup.FindMonthlyTopupStatusCardNumber{
		Year:       int32(year),
		Month:      int32(month),
		CardNumber: cardNumber,
	})
	if err != nil {
		h.logger.Debug("Failed to retrieve monthly topup status failed", zap.Error(err))
		return errors.ParseGrpcError(err)
	}

	apiResponse := h.mapper.ToApiResponseTopupMonthStatusFailed(res)
	h.cache.SetMonthTopupStatusFailedByCardNumberCache(ctx, reqCache, apiResponse)

	return c.JSON(http.StatusOK, apiResponse)
}

// FindYearlyTopupStatusFailedByCardNumber retrieves the yearly top-up status for failed transactions.
// @Summary Get yearly top-up status for failed transactions
// @Tags Topup Stats Status
// @Security Bearer
// @Description Retrieve the yearly top-up status for failed transactions by year.
// @Accept json
// @Produce json
// @Param year query int true "Year"
// @Param card_number query string true "Card Number"
// @Success 200 {object} response.ApiResponseTopupYearStatusFailed "Yearly top-up status for failed transactions"
// @Failure 400 {object} response.ErrorResponse "Invalid year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve yearly top-up status for failed transactions"
// @Router /api/topup-stats-status/yearly-failed [get]
func (h *topupStatsStatusHandleApi) FindYearlyTopupStatusFailedByCardNumber(c echo.Context) error {
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

	reqCache := &requests.YearTopupStatusCardNumber{
		CardNumber: cardNumber,
		Year:       year,
	}

	cachedData, found := h.cache.GetYearlyTopupStatusFailedByCardNumberCache(ctx, reqCache)
	if found {
		return c.JSON(http.StatusOK, cachedData)
	}

	res, err := h.client.FindYearlyTopupStatusFailedByCardNumber(ctx, &pbtopup.FindYearTopupStatusCardNumber{
		Year:       int32(year),
		CardNumber: cardNumber,
	})
	if err != nil {
		h.logger.Debug("Failed to retrieve yearly topup status failed", zap.Error(err))
		return errors.ParseGrpcError(err)
	}

	apiResponse := h.mapper.ToApiResponseTopupYearStatusFailed(res)
	h.cache.SetYearlyTopupStatusFailedByCardNumberCache(ctx, reqCache, apiResponse)

	return c.JSON(http.StatusOK, apiResponse)
}
