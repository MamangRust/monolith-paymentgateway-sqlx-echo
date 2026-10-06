package cardhandler

import (
	"net/http"

	card_cache "github.com/MamangRust/monolith-payment-gateway-apigateway/redis/api/card"
	pb "github.com/MamangRust/monolith-payment-gateway-pb/card"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
	apimapper "github.com/MamangRust/monolith-payment-gateway-shared/mapper/card"
	"github.com/labstack/echo/v4"
	"google.golang.org/protobuf/types/known/emptypb"
)

type cardDashboardHandleApi struct {
	card pb.CardDashboardServiceClient

	logger logger.LoggerInterface

	mapper apimapper.CardDashboardResponseMapper

	cache card_cache.CardMencache

	apiHandler errors.ApiHandler
}

type cardDashboardHandleApiDeps struct {
	client pb.CardDashboardServiceClient

	router *echo.Echo

	logger logger.LoggerInterface

	cache card_cache.CardMencache

	apiHandler errors.ApiHandler

	mapper apimapper.CardDashboardResponseMapper
}

func NewCardDashboardHandleApi(
	params *cardDashboardHandleApiDeps,
) *cardDashboardHandleApi {

	cardHandler := &cardDashboardHandleApi{
		card:       params.client,
		logger:     params.logger,
		mapper:     params.mapper,
		cache:      params.cache,
		apiHandler: params.apiHandler,
	}

	routerCard := params.router.Group("/api/card-dashboard")

	routerCard.GET("/dashboard", params.apiHandler.Handle("dashboard-card", cardHandler.DashboardCard))
	routerCard.GET("/dashboard/:cardNumber", params.apiHandler.Handle("dashboard-card-by-card-number", cardHandler.DashboardCardCardNumber))

	return cardHandler
}

// DashboardCard godoc
// @Summary Get dashboard card data
// @Description Retrieve dashboard card data
// @Tags Card Dashboard
// @Security Bearer
// @Produce json
// @Success 200 {object} response.ApiResponseDashboardCard
// @Failure 500 {object} response.ErrorResponse
// @Router /api/card-dashboard [get]
func (h *cardDashboardHandleApi) DashboardCard(c echo.Context) error {
	ctx := c.Request().Context()

	cachedData, found := h.cache.GetDashboardCardCache(ctx)
	if found {
		return c.JSON(http.StatusOK, cachedData)
	}

	res, err := h.card.DashboardCard(ctx, &emptypb.Empty{})
	if err != nil {
		return errors.ParseGrpcError(err)
	}

	apiResponse := h.mapper.ToApiResponseDashboardCard(res)
	h.cache.SetDashboardCardCache(ctx, apiResponse)

	return c.JSON(http.StatusOK, apiResponse)
}

// DashboardCardCardNumber godoc
// @Summary Get dashboard card data by card number
// @Description Retrieve dashboard card data for a specific card number
// @Tags Card Dashboard
// @Security Bearer
// @Produce json
// @Param cardNumber path string true "Card Number"
// @Success 200 {object} response.ApiResponseDashboardCardNumber
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /api/card-dashboard/{cardNumber} [get]
func (h *cardDashboardHandleApi) DashboardCardCardNumber(c echo.Context) error {
	ctx := c.Request().Context()

	cardNumber := c.Param("cardNumber")
	if cardNumber == "" {
		return errors.NewBadRequestError("cardNumber is required")
	}

	cachedData, found := h.cache.GetDashboardCardCardNumberCache(ctx, cardNumber)
	if found {
		return c.JSON(http.StatusOK, cachedData)
	}

	req := &pb.FindByCardNumberRequest{
		CardNumber: cardNumber,
	}

	res, err := h.card.DashboardCardNumber(ctx, req)
	if err != nil {
		return errors.ParseGrpcError(err)
	}

	apiResponse := h.mapper.ToApiResponseDashboardCardCardNumber(res)
	h.cache.SetDashboardCardCardNumberCache(ctx, cardNumber, apiResponse)

	return c.JSON(http.StatusOK, apiResponse)
}
