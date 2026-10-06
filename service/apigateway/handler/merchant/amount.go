package merchanthandler

import (
	"net/http"
	"strconv"

	merchant_cache "github.com/MamangRust/monolith-payment-gateway-apigateway/redis/api/merchant"
	pbmerchant "github.com/MamangRust/monolith-payment-gateway-pb/merchant"
	pb "github.com/MamangRust/monolith-payment-gateway-pb/merchant/stats"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	errors "github.com/MamangRust/monolith-payment-gateway-shared/errors"
	apimapper "github.com/MamangRust/monolith-payment-gateway-shared/mapper/merchant"
	"github.com/labstack/echo/v4"
)

type merchantStatsAmountHandleApi struct {
	client pb.MerchantStatsAmountServiceClient

	logger logger.LoggerInterface

	cache merchant_cache.MerchantMencache

	apiHandler errors.ApiHandler

	mapper apimapper.MerchantStatsAmountResponseMapper
}

type merchantStatsAmountHandleDeps struct {
	client pb.MerchantStatsAmountServiceClient

	router *echo.Echo

	logger logger.LoggerInterface

	cache merchant_cache.MerchantMencache

	apiHandler errors.ApiHandler

	mapper apimapper.MerchantStatsAmountResponseMapper
}

func NewMerchantStatsAmountHandleApi(params *merchantStatsAmountHandleDeps) *merchantStatsAmountHandleApi {

	merchantHandler := &merchantStatsAmountHandleApi{
		client:     params.client,
		logger:     params.logger,
		mapper:     params.mapper,
		cache:      params.cache,
		apiHandler: params.apiHandler,
	}

	routerMerchant := params.router.Group("/api/merchant-stats-amount")

	routerMerchant.GET("/monthly-amount", params.apiHandler.Handle("find-monthly-amount", merchantHandler.FindMonthlyAmountMerchant))
	routerMerchant.GET("/yearly-amount", params.apiHandler.Handle("find-yearly-amount", merchantHandler.FindYearlyAmountMerchant))

	routerMerchant.GET("/monthly-amount-by-merchant", params.apiHandler.Handle("find-monthly-amount-by-merchant", merchantHandler.FindMonthlyAmountByMerchants))
	routerMerchant.GET("/yearly-amount-by-merchant", params.apiHandler.Handle("find-yearly-amount-by-merchant", merchantHandler.FindYearlyAmountByMerchants))

	routerMerchant.GET("/monthly-amount-by-apikey", params.apiHandler.Handle("find-monthly-amount-by-apikey", merchantHandler.FindMonthlyAmountByApikeys))
	routerMerchant.GET("/yearly-amount-by-apikey", params.apiHandler.Handle("find-yearly-amount-by-apikey", merchantHandler.FindYearlyAmountByApikeys))

	return merchantHandler
}

// FindMonthlyAmountMerchant godoc
// @Summary Find monthly transaction amounts for a merchant
// @Tags Merchant Stats Amount
// @Security Bearer
// @Description Retrieve monthly transaction amounts for a merchant by year.
// @Accept json
// @Produce json
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponseMerchantMonthlyAmount "Monthly transaction amounts"
// @Failure 400 {object} response.ErrorResponse "Invalid year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve monthly transaction amounts"
// @Router /api/merchant-stats-amount/monthly-amount [get]
func (h *merchantStatsAmountHandleApi) FindMonthlyAmountMerchant(c echo.Context) error {
	yearStr := c.QueryParam("year")
	year, err := strconv.Atoi(yearStr)
	if err != nil || year <= 0 {
		return errors.NewBadRequestError("year is required and must be a positive integer")
	}

	ctx := c.Request().Context()

	cachedData, found := h.cache.GetMonthlyAmountMerchantCache(ctx, year)
	if found {
		return c.JSON(http.StatusOK, cachedData)
	}

	req := &pbmerchant.FindYearMerchant{
		Year: int32(year),
	}

	res, err := h.client.FindMonthlyAmountMerchant(ctx, req)
	if err != nil {
		return errors.ParseGrpcError(err)
	}

	apiResponse := h.mapper.ToApiResponseMonthlyAmounts(res)
	h.cache.SetMonthlyAmountMerchantCache(ctx, year, apiResponse)

	return c.JSON(http.StatusOK, apiResponse)
}

// FindYearlyAmountMerchant godoc.
// @Summary Find yearly transaction amounts for a merchant
// @Tags Merchant Stats Amount
// @Security Bearer
// @Description Retrieve yearly transaction amounts for a merchant by year.
// @Accept json
// @Produce json
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponseYearlyAmount "Yearly transaction amounts"
// @Failure 400 {object} response.ErrorResponse "Invalid year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve yearly transaction amounts"
// @Router /api/merchant-stats-amount/yearly-amount [get]
func (h *merchantStatsAmountHandleApi) FindYearlyAmountMerchant(c echo.Context) error {
	yearStr := c.QueryParam("year")
	year, err := strconv.Atoi(yearStr)
	if err != nil || year <= 0 {
		return errors.NewBadRequestError("year is required and must be a positive integer")
	}

	ctx := c.Request().Context()

	cachedData, found := h.cache.GetYearlyAmountMerchantCache(ctx, year)
	if found {
		return c.JSON(http.StatusOK, cachedData)
	}

	req := &pbmerchant.FindYearMerchant{
		Year: int32(year),
	}

	res, err := h.client.FindYearlyAmountMerchant(ctx, req)
	if err != nil {
		return errors.ParseGrpcError(err)
	}

	apiResponse := h.mapper.ToApiResponseYearlyAmounts(res)
	h.cache.SetYearlyAmountMerchantCache(ctx, year, apiResponse)

	return c.JSON(http.StatusOK, apiResponse)
}

// FindMonthlyAmountByMerchants godoc.
// @Summary Find monthly transaction amounts for a specific merchant
// @Tags Merchant Stats Amount
// @Security Bearer
// @Description Retrieve monthly transaction amounts for a specific merchant by year.
// @Accept json
// @Produce json
// @Param merchant_id query int true "Merchant ID"
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponseMerchantMonthlyAmount "Monthly transaction amounts"
// @Failure 400 {object} response.ErrorResponse "Invalid merchant ID or year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve monthly transaction amounts"
// @Router /api/merchant-stats-amount/monthly-amount-by-merchant [get]
func (h *merchantStatsAmountHandleApi) FindMonthlyAmountByMerchants(c echo.Context) error {
	merchantIDStr := c.QueryParam("merchant_id")
	yearStr := c.QueryParam("year")

	merchantID, err := strconv.Atoi(merchantIDStr)
	if err != nil || merchantID <= 0 {
		return errors.NewBadRequestError("merchant_id is required and must be a positive integer")
	}

	year, err := strconv.Atoi(yearStr)
	if err != nil || year <= 0 {
		return errors.NewBadRequestError("year is required and must be a positive integer")
	}

	ctx := c.Request().Context()

	reqCache := &requests.MonthYearAmountMerchant{
		MerchantID: merchantID,
		Year:       year,
	}

	cachedData, found := h.cache.GetMonthlyAmountByMerchantsCache(ctx, reqCache)
	if found {
		return c.JSON(http.StatusOK, cachedData)
	}

	reqGrpc := &pbmerchant.FindYearMerchantById{
		MerchantId: int32(merchantID),
		Year:       int32(year),
	}

	res, err := h.client.FindMonthlyAmountByMerchants(ctx, reqGrpc)
	if err != nil {
		return errors.ParseGrpcError(err)
	}

	apiResponse := h.mapper.ToApiResponseMonthlyAmounts(res)
	h.cache.SetMonthlyAmountByMerchantsCache(ctx, reqCache, apiResponse)

	return c.JSON(http.StatusOK, apiResponse)
}

// FindYearlyAmountByMerchants godoc.
// @Summary Find yearly transaction amounts for a specific merchant
// @Tags Merchant Stats Amount
// @Security Bearer
// @Description Retrieve yearly transaction amounts for a specific merchant by year.
// @Accept json
// @Produce json
// @Param merchant_id query int true "Merchant ID"
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponseMerchantYearlyAmount "Yearly transaction amounts"
// @Failure 400 {object} response.ErrorResponse "Invalid merchant ID or year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve yearly transaction amounts"
// @Router /api/merchant-stats-amount/yearly-amount-by-merchant [get]
func (h *merchantStatsAmountHandleApi) FindYearlyAmountByMerchants(c echo.Context) error {
	merchantIDStr := c.QueryParam("merchant_id")
	yearStr := c.QueryParam("year")

	merchantID, err := strconv.Atoi(merchantIDStr)
	if err != nil || merchantID <= 0 {
		return errors.NewBadRequestError("merchant_id is required and must be a positive integer")
	}

	year, err := strconv.Atoi(yearStr)
	if err != nil || year <= 0 {
		return errors.NewBadRequestError("year is required and must be a positive integer")
	}

	ctx := c.Request().Context()

	reqCache := &requests.MonthYearAmountMerchant{
		MerchantID: merchantID,
		Year:       year,
	}

	cachedData, found := h.cache.GetYearlyAmountByMerchantsCache(ctx, reqCache)
	if found {
		return c.JSON(http.StatusOK, cachedData)
	}

	reqGrpc := &pbmerchant.FindYearMerchantById{
		MerchantId: int32(merchantID),
		Year:       int32(year),
	}

	res, err := h.client.FindYearlyAmountByMerchants(ctx, reqGrpc)
	if err != nil {
		return errors.ParseGrpcError(err)
	}

	apiResponse := h.mapper.ToApiResponseYearlyAmounts(res)
	h.cache.SetYearlyAmountByMerchantsCache(ctx, reqCache, apiResponse)

	return c.JSON(http.StatusOK, apiResponse)
}

// FindMonthlyAmountByApikeys godoc.
// @Summary Find monthly transaction amounts for a specific merchant
// @Tags Merchant Stats Amount
// @Security Bearer
// @Description Retrieve monthly transaction amounts for a specific merchant by year.
// @Accept json
// @Produce json
// @Param merchant_id query int true "Merchant ID"
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponseMerchantMonthlyAmount "Monthly transaction amounts"
// @Failure 400 {object} response.ErrorResponse "Invalid merchant ID or year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve monthly transaction amounts"
// @Router /api/merchant-stats-amount/monthly-amount-by-apikey [get]
func (h *merchantStatsAmountHandleApi) FindMonthlyAmountByApikeys(c echo.Context) error {
	api_key := c.QueryParam("api_key")
	yearStr := c.QueryParam("year")

	if api_key == "" {
		return errors.NewBadRequestError("api_key is required")
	}

	year, err := strconv.Atoi(yearStr)
	if err != nil || year <= 0 {
		return errors.NewBadRequestError("year is required and must be a positive integer")
	}

	ctx := c.Request().Context()

	reqCache := &requests.MonthYearAmountApiKey{
		Apikey: api_key,
		Year:   year,
	}

	cachedData, found := h.cache.GetMonthlyAmountByApikeysCache(ctx, reqCache)
	if found {
		return c.JSON(http.StatusOK, cachedData)
	}

	reqGrpc := &pbmerchant.FindYearMerchantByApikey{
		ApiKey: api_key,
		Year:   int32(year),
	}

	res, err := h.client.FindMonthlyAmountByApikey(ctx, reqGrpc)
	if err != nil {
		return errors.ParseGrpcError(err)
	}

	apiResponse := h.mapper.ToApiResponseMonthlyAmounts(res)
	h.cache.SetMonthlyAmountByApikeysCache(ctx, reqCache, apiResponse)

	return c.JSON(http.StatusOK, apiResponse)
}

// FindYearlyAmountByApikeys godoc.
// @Summary Find yearly transaction amounts for a specific merchant
// @Tags Merchant Stats Amount
// @Security Bearer
// @Description Retrieve yearly transaction amounts for a specific merchant by year.
// @Accept json
// @Produce json
// @Param merchant_id query int true "Merchant ID"
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponseMerchantYearlyAmount "Yearly transaction amounts"
// @Failure 400 {object} response.ErrorResponse "Invalid merchant ID or year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve yearly transaction amounts"
// @Router /api/merchant-stats-amount/yearly-amount-by-apikey [get]
func (h *merchantStatsAmountHandleApi) FindYearlyAmountByApikeys(c echo.Context) error {
	api_key := c.QueryParam("api_key")
	yearStr := c.QueryParam("year")

	if api_key == "" {
		return errors.NewBadRequestError("api_key is required")
	}

	year, err := strconv.Atoi(yearStr)
	if err != nil || year <= 0 {
		return errors.NewBadRequestError("year is required and must be a positive integer")
	}

	ctx := c.Request().Context()

	reqCache := &requests.MonthYearAmountApiKey{
		Apikey: api_key,
		Year:   year,
	}

	cachedData, found := h.cache.GetYearlyAmountByApikeysCache(ctx, reqCache)
	if found {
		return c.JSON(http.StatusOK, cachedData)
	}

	reqGrpc := &pbmerchant.FindYearMerchantByApikey{
		ApiKey: api_key,
		Year:   int32(year),
	}

	res, err := h.client.FindYearlyAmountByApikey(ctx, reqGrpc)
	if err != nil {
		return errors.ParseGrpcError(err)
	}

	apiResponse := h.mapper.ToApiResponseYearlyAmounts(res)
	h.cache.SetYearlyAmountByApikeysCache(ctx, reqCache, apiResponse)

	return c.JSON(http.StatusOK, apiResponse)
}
