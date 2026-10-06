package cardhandler

import (
	"net/http"
	"strconv"

	card_cache "github.com/MamangRust/monolith-payment-gateway-apigateway/redis/api/card"
	pbcard "github.com/MamangRust/monolith-payment-gateway-pb/card"
	pb "github.com/MamangRust/monolith-payment-gateway-pb/card/stats"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
	apimapper "github.com/MamangRust/monolith-payment-gateway-shared/mapper/card"
	"github.com/labstack/echo/v4"
)

type cardStatsWithdrawHandleApi struct {
	card       pb.CardStatsWithdrawServiceClient
	apiHandler errors.ApiHandler
	cache      card_cache.CardMencache
	logger     logger.LoggerInterface
	mapper     apimapper.CardStatsAmountResponseMapper
}

type cardStatsWithdrawHandleApiDeps struct {
	client     pb.CardStatsWithdrawServiceClient
	router     *echo.Echo
	apiHandler errors.ApiHandler
	cache      card_cache.CardMencache
	logger     logger.LoggerInterface
	mapper     apimapper.CardStatsAmountResponseMapper
}

func NewCardStatsWithdrawHandleApi(
	params *cardStatsWithdrawHandleApiDeps,
) *cardStatsWithdrawHandleApi {

	cardHandler := &cardStatsWithdrawHandleApi{
		card:       params.client,
		logger:     params.logger,
		mapper:     params.mapper,
		cache:      params.cache,
		apiHandler: params.apiHandler,
	}

	routerCard := params.router.Group("/api/card-stats-withdraw")

	routerCard.GET("/monthly-withdraw-amount", params.apiHandler.Handle("find-monthly-withdraw-amount", cardHandler.FindMonthlyWithdrawAmount))
	routerCard.GET("/yearly-withdraw-amount", params.apiHandler.Handle("find-yearly-withdraw-amount", cardHandler.FindYearlyWithdrawAmount))
	routerCard.GET("/monthly-withdraw-amount-by-card", params.apiHandler.Handle("find-monthly-withdraw-amount-by-card", cardHandler.FindMonthlyWithdrawAmountByCardNumber))
	routerCard.GET("/yearly-withdraw-amount-by-card", params.apiHandler.Handle("find-yearly-withdraw-amount-by-card", cardHandler.FindYearlyWithdrawAmountByCardNumber))

	return cardHandler
}

// FindMonthlyWithdrawAmount godoc
// @Summary Get monthly withdraw amount data
// @Description Retrieve monthly withdraw amount data for a specific year
// @Tags Card Stats Withdraw
// @Security Bearer
// @Produce json
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponseMonthlyAmount
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /api/card-stats-withdraw/monthly-withdraw-amount [get]
func (h *cardStatsWithdrawHandleApi) FindMonthlyWithdrawAmount(c echo.Context) error {
	yearStr := c.QueryParam("year")
	year, err := strconv.Atoi(yearStr)
	if err != nil || year <= 0 {
		return errors.NewBadRequestError("year is required and must be a positive integer")
	}

	ctx := c.Request().Context()

	cachedData, found := h.cache.GetMonthlyWithdrawCache(ctx, year)
	if found {
		return c.JSON(http.StatusOK, cachedData)
	}

	reqGrpc := &pbcard.FindYearAmount{
		Year: int32(year),
	}

	res, err := h.card.FindMonthlyWithdrawAmount(ctx, reqGrpc)
	if err != nil {
		return errors.ParseGrpcError(err)
	}

	apiResponse := h.mapper.ToApiResponseMonthlyAmounts(res)
	h.cache.SetMonthlyWithdrawCache(ctx, year, apiResponse)

	return c.JSON(http.StatusOK, apiResponse)
}

// FindYearlyWithdrawAmount godoc
// @Summary Get yearly withdraw amount data
// @Description Retrieve yearly withdraw amount data for a specific year
// @Tags Card Stats Withdraw
// @Security Bearer
// @Produce json
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponseYearlyAmount
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /api/card-stats-withdraw/yearly-withdraw-amount [get]
func (h *cardStatsWithdrawHandleApi) FindYearlyWithdrawAmount(c echo.Context) error {
	yearStr := c.QueryParam("year")
	year, err := strconv.Atoi(yearStr)
	if err != nil || year <= 0 {
		return errors.NewBadRequestError("year is required and must be a positive integer")
	}

	ctx := c.Request().Context()

	cachedData, found := h.cache.GetYearlyWithdrawCache(ctx, year)
	if found {
		return c.JSON(http.StatusOK, cachedData)
	}

	reqGrpc := &pbcard.FindYearAmount{
		Year: int32(year),
	}

	res, err := h.card.FindYearlyWithdrawAmount(ctx, reqGrpc)
	if err != nil {
		return errors.ParseGrpcError(err)
	}

	apiResponse := h.mapper.ToApiResponseYearlyAmounts(res)
	h.cache.SetYearlyWithdrawCache(ctx, year, apiResponse)

	return c.JSON(http.StatusOK, apiResponse)
}

// FindMonthlyWithdrawAmountByCardNumber godoc
// @Summary Get monthly withdraw amount data by card number
// @Description Retrieve monthly withdraw amount data for a specific year and card number
// @Tags Card Stats Withdraw
// @Security Bearer
// @Produce json
// @Param year query int true "Year"
// @Param card_number query string true "Card Number"
// @Success 200 {object} response.ApiResponseMonthlyAmount
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /api/card-stats-withdraw/monthly-withdraw-amount-by-card [get]
func (h *cardStatsWithdrawHandleApi) FindMonthlyWithdrawAmountByCardNumber(c echo.Context) error {
	yearStr := c.QueryParam("year")
	year, err := strconv.Atoi(yearStr)
	if err != nil || year <= 0 {
		return errors.NewBadRequestError("year is required and must be a positive integer")
	}

	cardNumber := c.QueryParam("card_number")
	if cardNumber == "" {
		return errors.NewBadRequestError("card_number is required")
	}

	ctx := c.Request().Context()

	reqCache := &requests.MonthYearCardNumberCard{
		Year:       year,
		CardNumber: cardNumber,
	}

	cachedData, found := h.cache.GetMonthlyWithdrawByNumberCache(ctx, reqCache)
	if found {
		return c.JSON(http.StatusOK, cachedData)
	}

	reqGrpc := &pbcard.FindYearAmountCardNumber{
		Year:       int32(year),
		CardNumber: cardNumber,
	}

	res, err := h.card.FindMonthlyWithdrawAmountByCardNumber(ctx, reqGrpc)
	if err != nil {
		return errors.ParseGrpcError(err)
	}

	apiResponse := h.mapper.ToApiResponseMonthlyAmounts(res)
	h.cache.SetMonthlyWithdrawByNumberCache(ctx, reqCache, apiResponse)

	return c.JSON(http.StatusOK, apiResponse)
}

// FindYearlyWithdrawAmountByCardNumber godoc
// @Summary Get yearly withdraw amount data by card number
// @Description Retrieve yearly withdraw amount data for a specific year and card number
// @Tags Card Stats Withdraw
// @Security Bearer
// @Produce json
// @Param year query int true "Year"
// @Param card_number query string true "Card Number"
// @Success 200 {object} response.ApiResponseYearlyAmount
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /api/card-stats-withdraw/yearly-withdraw-amount-by-card [get]
func (h *cardStatsWithdrawHandleApi) FindYearlyWithdrawAmountByCardNumber(c echo.Context) error {
	yearStr := c.QueryParam("year")
	year, err := strconv.Atoi(yearStr)
	if err != nil || year <= 0 {
		return errors.NewBadRequestError("year is required and must be a positive integer")
	}

	cardNumber := c.QueryParam("card_number")
	if cardNumber == "" {
		return errors.NewBadRequestError("card_number is required")
	}

	ctx := c.Request().Context()

	reqCache := &requests.MonthYearCardNumberCard{
		Year:       year,
		CardNumber: cardNumber,
	}

	cachedData, found := h.cache.GetYearlyWithdrawByNumberCache(ctx, reqCache)
	if found {
		return c.JSON(http.StatusOK, cachedData)
	}

	reqGrpc := &pbcard.FindYearAmountCardNumber{
		Year:       int32(year),
		CardNumber: cardNumber,
	}

	res, err := h.card.FindYearlyWithdrawAmountByCardNumber(ctx, reqGrpc)
	if err != nil {
		return errors.ParseGrpcError(err)
	}

	apiResponse := h.mapper.ToApiResponseYearlyAmounts(res)
	h.cache.SetYearlyWithdrawByNumberCache(ctx, reqCache, apiResponse)

	return c.JSON(http.StatusOK, apiResponse)
}
