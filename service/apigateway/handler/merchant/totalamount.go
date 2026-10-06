package merchanthandler

import (
	"net/http"
	"strconv"

	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	errors "github.com/MamangRust/monolith-payment-gateway-shared/errors"

	merchant_cache "github.com/MamangRust/monolith-payment-gateway-apigateway/redis/api/merchant"
	pbmerchant "github.com/MamangRust/monolith-payment-gateway-pb/merchant"
	pb "github.com/MamangRust/monolith-payment-gateway-pb/merchant/stats"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	apimapper "github.com/MamangRust/monolith-payment-gateway-shared/mapper/merchant"
	"github.com/labstack/echo/v4"
)

type merchantStatsTotalAmountHandleApi struct {
	client pb.MerchantStatsTotalAmountServiceClient

	logger logger.LoggerInterface

	mapper apimapper.MerchantStatsTotalAmountResponseMapper

	cache merchant_cache.MerchantMencache

	apiHandler errors.ApiHandler
}

type merchantStatsTotalAmountHandleDeps struct {
	client pb.MerchantStatsTotalAmountServiceClient

	router *echo.Echo

	logger logger.LoggerInterface

	mapper apimapper.MerchantStatsTotalAmountResponseMapper

	cache merchant_cache.MerchantMencache

	apiHandler errors.ApiHandler
}

func NewMerchantStatsTotalAmountHandleApi(params *merchantStatsTotalAmountHandleDeps) *merchantStatsTotalAmountHandleApi {

	merchantHandler := &merchantStatsTotalAmountHandleApi{
		client:     params.client,
		logger:     params.logger,
		mapper:     params.mapper,
		apiHandler: params.apiHandler,
		cache:      params.cache,
	}

	routerMerchant := params.router.Group("/api/merchant-stats-totalamount")

	routerMerchant.GET("/monthly-total-amount", params.apiHandler.Handle("find-monthly-total-amount", merchantHandler.FindMonthlyTotalAmountMerchant))
	routerMerchant.GET("/yearly-total-amount", params.apiHandler.Handle("find-yearly-total-amount", merchantHandler.FindYearlyTotalAmountMerchant))

	routerMerchant.GET("/monthly-totalamount-by-merchant", params.apiHandler.Handle("find-monthly-total-amount-by-merchant", merchantHandler.FindMonthlyTotalAmountByMerchants))
	routerMerchant.GET("/yearly-totalamount-by-merchant", params.apiHandler.Handle("find-yearly-total-amount-by-merchant", merchantHandler.FindYearlyTotalAmountByMerchants))

	routerMerchant.GET("/monthly-totalamount-by-apikey", params.apiHandler.Handle("find-monthly-total-amount-by-apikey", merchantHandler.FindMonthlyTotalAmountByApikeys))
	routerMerchant.GET("/yearly-totalamount-by-apikey", params.apiHandler.Handle("find-yearly-total-amount-by-apikey", merchantHandler.FindYearlyTotalAmountByApikeys))

	return merchantHandler
}

// FindMonthlyAmountMerchant godoc
// @Summary Find monthly transaction amounts for a merchant
// @Tags Merchant Stats Total Amount
// @Security Bearer
// @Description Retrieve monthly transaction amounts for a merchant by year.
// @Accept json
// @Produce json
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponseMerchantMonthlyAmount "Monthly transaction amounts"
// @Failure 400 {object} response.ErrorResponse "Invalid year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve monthly transaction amounts"
// @Router /api/merchant-stats-totalamount/monthly-total-amount [get]
func (h *merchantStatsTotalAmountHandleApi) FindMonthlyTotalAmountMerchant(c echo.Context) error {
	yearStr := c.QueryParam("year")
	year, err := strconv.Atoi(yearStr)
	if err != nil || year <= 0 {
		return errors.NewBadRequestError("year is required and must be a positive integer")
	}

	ctx := c.Request().Context()

	cachedData, found := h.cache.GetMonthlyTotalAmountMerchantCache(ctx, year)
	if found {
		return c.JSON(http.StatusOK, cachedData)
	}

	req := &pbmerchant.FindYearMerchant{
		Year: int32(year),
	}

	res, err := h.client.FindMonthlyTotalAmountMerchant(ctx, req)
	if err != nil {
		return errors.ParseGrpcError(err)
	}

	apiResponse := h.mapper.ToApiResponseMonthlyTotalAmounts(res)
	h.cache.SetMonthlyTotalAmountMerchantCache(ctx, year, apiResponse)

	return c.JSON(http.StatusOK, apiResponse)
}

// FindYearlyAmountMerchant godoc.
// @Summary Find yearly transaction amounts for a merchant
// @Tags Merchant Stats Total Amount
// @Security Bearer
// @Description Retrieve yearly transaction amounts for a merchant by year.
// @Accept json
// @Produce json
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponseYearlyAmount "Yearly transaction amounts"
// @Failure 400 {object} response.ErrorResponse "Invalid year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve yearly transaction amounts"
// @Router /api/merchant-stats-totalamount/yearly-total-amount [get]
func (h *merchantStatsTotalAmountHandleApi) FindYearlyTotalAmountMerchant(c echo.Context) error {
	yearStr := c.QueryParam("year")
	year, err := strconv.Atoi(yearStr)
	if err != nil || year <= 0 {
		return errors.NewBadRequestError("year is required and must be a positive integer")
	}

	ctx := c.Request().Context()

	cachedData, found := h.cache.GetYearlyTotalAmountMerchantCache(ctx, year)
	if found {
		return c.JSON(http.StatusOK, cachedData)
	}

	req := &pbmerchant.FindYearMerchant{
		Year: int32(year),
	}

	res, err := h.client.FindYearlyTotalAmountMerchant(ctx, req)
	if err != nil {
		return errors.ParseGrpcError(err)
	}

	apiResponse := h.mapper.ToApiResponseYearlyTotalAmounts(res)
	h.cache.SetYearlyTotalAmountMerchantCache(ctx, year, apiResponse)

	return c.JSON(http.StatusOK, apiResponse)
}

// FindMonthlyAmountByMerchants godoc.
// @Summary Find monthly transaction amounts for a specific merchant
// @Tags Merchant Stats Total Amount
// @Security Bearer
// @Description Retrieve monthly transaction amounts for a specific merchant by year.
// @Accept json
// @Produce json
// @Param merchant_id query int true "Merchant ID"
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponseMerchantMonthlyAmount "Monthly transaction amounts"
// @Failure 400 {object} response.ErrorResponse "Invalid merchant ID or year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve monthly transaction amounts"
// @Router /api/merchant-stats-totalamount/monthly-totalamount-by-merchant [get]
func (h *merchantStatsTotalAmountHandleApi) FindMonthlyTotalAmountByMerchants(c echo.Context) error {
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

	reqCache := &requests.MonthYearTotalAmountMerchant{
		MerchantID: merchantID,
		Year:       year,
	}

	cachedData, found := h.cache.GetMonthlyTotalAmountByMerchantsCache(ctx, reqCache)
	if found {
		return c.JSON(http.StatusOK, cachedData)
	}

	reqGrpc := &pbmerchant.FindYearMerchantById{
		MerchantId: int32(merchantID),
		Year:       int32(year),
	}

	res, err := h.client.FindMonthlyTotalAmountByMerchants(ctx, reqGrpc)
	if err != nil {
		return errors.ParseGrpcError(err)
	}

	apiResponse := h.mapper.ToApiResponseMonthlyTotalAmounts(res)
	h.cache.SetMonthlyTotalAmountByMerchantsCache(ctx, reqCache, apiResponse)

	return c.JSON(http.StatusOK, apiResponse)
}

// FindYearlyAmountByMerchants godoc.
// @Summary Find yearly transaction amounts for a specific merchant
// @Tags Merchant Stats Total Amount
// @Security Bearer
// @Description Retrieve yearly transaction amounts for a specific merchant by year.
// @Accept json
// @Produce json
// @Param merchant_id query int true "Merchant ID"
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponseMerchantYearlyAmount "Yearly transaction amounts"
// @Failure 400 {object} response.ErrorResponse "Invalid merchant ID or year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve yearly transaction amounts"
// @Router /api/merchant-stats-totalamount/yearly-totalamount-by-merchant [get]
func (h *merchantStatsTotalAmountHandleApi) FindYearlyTotalAmountByMerchants(c echo.Context) error {
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

	reqCache := &requests.MonthYearTotalAmountMerchant{
		MerchantID: merchantID,
		Year:       year,
	}

	cachedData, found := h.cache.GetYearlyTotalAmountByMerchantsCache(ctx, reqCache)
	if found {
		return c.JSON(http.StatusOK, cachedData)
	}

	reqGrpc := &pbmerchant.FindYearMerchantById{
		MerchantId: int32(merchantID),
		Year:       int32(year),
	}

	res, err := h.client.FindYearlyTotalAmountByMerchants(ctx, reqGrpc)
	if err != nil {
		return errors.ParseGrpcError(err)
	}

	apiResponse := h.mapper.ToApiResponseYearlyTotalAmounts(res)
	h.cache.SetYearlyTotalAmountByMerchantsCache(ctx, reqCache, apiResponse)

	return c.JSON(http.StatusOK, apiResponse)
}

// FindMonthlyAmountByApikeys godoc.
// @Summary Find monthly transaction amounts for a specific merchant
// @Tags Merchant Stats Total Amount
// @Security Bearer
// @Description Retrieve monthly transaction amounts for a specific merchant by year.
// @Accept json
// @Produce json
// @Param merchant_id query int true "Merchant ID"
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponseMerchantMonthlyAmount "Monthly transaction amounts"
// @Failure 400 {object} response.ErrorResponse "Invalid merchant ID or year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve monthly transaction amounts"
// @Router /api/merchant-stats-totalamount/monthly-totalamount-by-apikey [get]
func (h *merchantStatsTotalAmountHandleApi) FindMonthlyTotalAmountByApikeys(c echo.Context) error {
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

	reqCache := &requests.MonthYearTotalAmountApiKey{
		Apikey: api_key,
		Year:   year,
	}

	cachedData, found := h.cache.GetMonthlyTotalAmountByApikeysCache(ctx, reqCache)
	if found {
		return c.JSON(http.StatusOK, cachedData)
	}

	reqGrpc := &pbmerchant.FindYearMerchantByApikey{
		ApiKey: api_key,
		Year:   int32(year),
	}

	res, err := h.client.FindMonthlyTotalAmountByApikey(ctx, reqGrpc)
	if err != nil {
		return errors.ParseGrpcError(err)
	}

	apiResponse := h.mapper.ToApiResponseMonthlyTotalAmounts(res)
	h.cache.SetMonthlyTotalAmountByApikeysCache(ctx, reqCache, apiResponse)

	return c.JSON(http.StatusOK, apiResponse)
}

// FindYearlyAmountByApikeys godoc.
// @Summary Find yearly transaction amounts for a specific merchant
// @Tags Merchant Stats Total Amount
// @Security Bearer
// @Description Retrieve yearly transaction amounts for a specific merchant by year.
// @Accept json
// @Produce json
// @Param merchant_id query int true "Merchant ID"
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponseMerchantYearlyAmount "Yearly transaction amounts"
// @Failure 400 {object} response.ErrorResponse "Invalid merchant ID or year"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve yearly transaction amounts"
// @Router /api/merchant-stats-totalamount/yearly-totalamount-by-apikey [get]
func (h *merchantStatsTotalAmountHandleApi) FindYearlyTotalAmountByApikeys(c echo.Context) error {
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

	reqCache := &requests.MonthYearTotalAmountApiKey{
		Apikey: api_key,
		Year:   year,
	}

	cachedData, found := h.cache.GetYearlyTotalAmountByApikeysCache(ctx, reqCache)
	if found {
		return c.JSON(http.StatusOK, cachedData)
	}

	reqGrpc := &pbmerchant.FindYearMerchantByApikey{
		ApiKey: api_key,
		Year:   int32(year),
	}

	res, err := h.client.FindYearlyTotalAmountByApikey(ctx, reqGrpc)
	if err != nil {
		return errors.ParseGrpcError(err)
	}

	apiResponse := h.mapper.ToApiResponseYearlyTotalAmounts(res)
	h.cache.SetYearlyTotalAmountByApikeysCache(ctx, reqCache, apiResponse)

	return c.JSON(http.StatusOK, apiResponse)
}
