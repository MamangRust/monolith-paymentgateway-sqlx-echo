package saldohandler

import (
	"net/http"
	"strconv"

	saldo_cache "github.com/MamangRust/monolith-payment-gateway-apigateway/redis/api/saldo"
	pbsaldo "github.com/MamangRust/monolith-payment-gateway-pb/saldo"
	pb "github.com/MamangRust/monolith-payment-gateway-pb/saldo/stats"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
	apimapper "github.com/MamangRust/monolith-payment-gateway-shared/mapper/saldo"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type saldoStatsBalanceHandleApi struct {
	saldo pb.SaldoStatsBalanceServiceClient

	logger logger.LoggerInterface

	mapper apimapper.SaldoStatsBalanceResponseMapper

	cache saldo_cache.SaldoMencache

	apiHandler errors.ApiHandler
}

type saldoStatsBalanceHandleDeps struct {
	client pb.SaldoStatsBalanceServiceClient

	router *echo.Echo

	logger logger.LoggerInterface

	mapper apimapper.SaldoStatsBalanceResponseMapper

	cache saldo_cache.SaldoMencache

	apiHandler errors.ApiHandler
}

func NewSaldoStatsBalanceHandleApi(params *saldoStatsBalanceHandleDeps) *saldoStatsBalanceHandleApi {

	saldoHandler := &saldoStatsBalanceHandleApi{
		saldo:      params.client,
		logger:     params.logger,
		mapper:     params.mapper,
		cache:      params.cache,
		apiHandler: params.apiHandler,
	}

	routerSaldo := params.router.Group("/api/saldo-stats-balance")

	routerSaldo.GET("/monthly-balances", params.apiHandler.Handle("find-monthly-saldo-balances", saldoHandler.FindMonthlySaldoBalances))
	routerSaldo.GET("/yearly-balances", params.apiHandler.Handle("find-yearly-saldo-balances", saldoHandler.FindYearlySaldoBalances))

	return saldoHandler
}

// FindMonthlySaldoBalances retrieves monthly saldo balances for a specific year.
// @Summary Get monthly saldo balances
// @Tags Saldo Stats Balance
// @Security Bearer
// @Description Retrieve monthly saldo balances for a specific year.
// @Accept json
// @Produce json
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponseMonthSaldoBalances "Monthly saldo balances"
// @Failure 400 {object} response.ErrorResponse "Invalid year parameter"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve monthly saldo balances"
// @Router /api/saldo-stats-balances/monthly-balances [get]
func (h *saldoStatsBalanceHandleApi) FindMonthlySaldoBalances(c echo.Context) error {
	yearStr := c.QueryParam("year")
	year, err := strconv.Atoi(yearStr)
	if err != nil || year <= 0 {
		return errors.NewBadRequestError("invalid year parameter")
	}

	ctx := c.Request().Context()

	cachedData, found := h.cache.GetMonthlySaldoBalanceCache(ctx, year)
	if found {
		return c.JSON(http.StatusOK, cachedData)
	}

	res, err := h.saldo.FindMonthlySaldoBalances(ctx, &pbsaldo.FindYearlySaldo{
		Year: int32(year),
	})
	if err != nil {
		h.logger.Debug("Failed to retrieve monthly saldo balances", zap.Error(err))
		return errors.ParseGrpcError(err)
	}

	apiResponse := h.mapper.ToApiResponseMonthSaldoBalances(res)
	h.cache.SetMonthlySaldoBalanceCache(ctx, year, apiResponse)

	return c.JSON(http.StatusOK, apiResponse)
}

// FindYearlySaldoBalances retrieves yearly saldo balances for a specific year.
// @Summary Get yearly saldo balances
// @Tags Saldo Stats Balance
// @Security Bearer
// @Description Retrieve yearly saldo balances for a specific year.
// @Accept json
// @Produce json
// @Param year query int true "Year"
// @Success 200 {object} response.ApiResponseYearSaldoBalances "Yearly saldo balances"
// @Failure 400 {object} response.ErrorResponse "Invalid year parameter"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve yearly saldo balances"
// @Router /api/saldo-stats-balance/yearly-balances [get]
func (h *saldoStatsBalanceHandleApi) FindYearlySaldoBalances(c echo.Context) error {
	yearStr := c.QueryParam("year")
	year, err := strconv.Atoi(yearStr)
	if err != nil || year <= 0 {
		return errors.NewBadRequestError("invalid year parameter")
	}

	ctx := c.Request().Context()

	cachedData, found := h.cache.GetYearlySaldoBalanceCache(ctx, year)
	if found {
		return c.JSON(http.StatusOK, cachedData)
	}

	res, err := h.saldo.FindYearlySaldoBalances(ctx, &pbsaldo.FindYearlySaldo{
		Year: int32(year),
	})
	if err != nil {
		h.logger.Debug("Failed to retrieve yearly saldo balances", zap.Error(err))
		return errors.ParseGrpcError(err)
	}

	apiResponse := h.mapper.ToApiResponseYearSaldoBalances(res)
	h.cache.SetYearlySaldoBalanceCache(ctx, year, apiResponse)

	return c.JSON(http.StatusOK, apiResponse)
}
