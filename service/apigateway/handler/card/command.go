package cardhandler

import (
	"fmt"
	"net/http"
	"strconv"

	card_cache "github.com/MamangRust/monolith-payment-gateway-apigateway/redis/api/card"
	errors "github.com/MamangRust/monolith-payment-gateway-shared/errors"
	"github.com/go-playground/validator/v10"

	pb "github.com/MamangRust/monolith-payment-gateway-pb/card"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	apimapper "github.com/MamangRust/monolith-payment-gateway-shared/mapper/card"
	"github.com/labstack/echo/v4"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type cardCommandHandleApi struct {
	card pb.CardCommandServiceClient

	logger logger.LoggerInterface
	mapper apimapper.CardCommandResponseMapper

	cache card_cache.CardMencache

	apiHandler errors.ApiHandler
}

type cardCommandHandleApiDeps struct {
	client pb.CardCommandServiceClient

	router *echo.Echo

	logger logger.LoggerInterface

	cache card_cache.CardMencache

	apiHandler errors.ApiHandler

	mapper apimapper.CardCommandResponseMapper
}

func NewCardCommandHandleApi(params *cardCommandHandleApiDeps) *cardCommandHandleApi {
	cardHandler := &cardCommandHandleApi{
		card:       params.client,
		logger:     params.logger,
		mapper:     params.mapper,
		cache:      params.cache,
		apiHandler: params.apiHandler,
	}

	routerCard := params.router.Group("/api/card-command")

	routerCard.POST("/create", params.apiHandler.Handle("create-card", cardHandler.CreateCard))
	routerCard.POST("/update/:id", params.apiHandler.Handle("update-card", cardHandler.UpdateCard))
	routerCard.POST("/trashed/:id", params.apiHandler.Handle("trashed-card", cardHandler.TrashedCard))
	routerCard.POST("/restore/:id", params.apiHandler.Handle("restore-card", cardHandler.RestoreCard))
	routerCard.DELETE("/permanent/:id", params.apiHandler.Handle("delete-card-permanent", cardHandler.DeleteCardPermanent))
	routerCard.POST("/restore/all", params.apiHandler.Handle("restore-all-card", cardHandler.RestoreAllCard))
	routerCard.POST("/permanent/all", params.apiHandler.Handle("delete-all-card-permanent", cardHandler.DeleteAllCardPermanent))

	routerCard.POST("/toggle-status", params.apiHandler.Handle("toggle-card-status", cardHandler.ToggleCardStatus))
	routerCard.POST("/update-credit-limit/:id", params.apiHandler.Handle("update-credit-limit", cardHandler.UpdateCreditLimit))
	routerCard.POST("/redeem-points/:id", params.apiHandler.Handle("redeem-points", cardHandler.RedeemPoints))
	routerCard.POST("/process-billing", params.apiHandler.Handle("process-billing-cycles", cardHandler.ProcessBillingCycles))

	return cardHandler
}

// @Security Bearer
// @Summary Create a new card
// @Tags Card Command
// @Description Create a new card for a user
// @Accept json
// @Produce json
// @Param CreateCardRequest body requests.CreateCardRequest true "Create card request"
// @Success 200 {object} response.ApiResponseCard "Created card"
// @Failure 400 {object} response.ErrorResponse "Bad request or validation error"
// @Failure 500 {object} response.ErrorResponse "Failed to create card"
// @Router /api/card-command/create [post]
func (h *cardCommandHandleApi) CreateCard(c echo.Context) error {
	var body requests.CreateCardRequest

	if err := c.Bind(&body); err != nil {
		return errors.NewBadRequestError("Invalid request")
	}

	if err := body.Validate(); err != nil {
		validations := h.parseValidationErrors(err)
		return errors.NewValidationError(validations)
	}

	ctx := c.Request().Context()

	req := &pb.CreateCardRequest{
		UserId:       int32(body.UserID),
		CardType:     body.CardType,
		ExpireDate:   timestamppb.New(body.ExpireDate),
		Cvv:          body.CVV,
		CardProvider: body.CardProvider,
	}

	res, err := h.card.CreateCard(ctx, req)

	if err != nil {
		return errors.ParseGrpcError(err)
	}

	so := h.mapper.ToApiResponseCard(res)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Update a card
// @Tags Card Command
// @Description Update a card for a user
// @Accept json
// @Produce json
// @Param id path int true "Card ID"
// @Param UpdateCardRequest body requests.UpdateCardRequest true "Update card request"
// @Success 200 {object} response.ApiResponseCard "Updated card"
// @Failure 400 {object} response.ErrorResponse "Bad request or validation error"
// @Failure 500 {object} response.ErrorResponse "Failed to update card"
// @Router /api/card-command/update/{id} [post]
func (h *cardCommandHandleApi) UpdateCard(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		return errors.NewBadRequestError("id is required")
	}

	var body requests.UpdateCardRequest

	if err := c.Bind(&body); err != nil {
		return errors.NewBadRequestError("Invalid request")
	}

	if err := body.Validate(); err != nil {
		validations := h.parseValidationErrors(err)
		return errors.NewValidationError(validations)
	}

	ctx := c.Request().Context()

	req := &pb.UpdateCardRequest{
		CardId:       int32(idInt),
		UserId:       int32(body.UserID),
		CardType:     body.CardType,
		ExpireDate:   timestamppb.New(body.ExpireDate),
		Cvv:          body.CVV,
		CardProvider: body.CardProvider,
	}

	res, err := h.card.UpdateCard(ctx, req)

	if err != nil {
		return errors.ParseGrpcError(err)
	}

	so := h.mapper.ToApiResponseCard(res)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Trashed a card
// @Tags Card Command
// @Description Trashed a card by its ID
// @Accept json
// @Produce json
// @Param id path int true "Card ID"
// @Success 200 {object} response.ApiResponseCard "Trashed card"
// @Failure 400 {object} response.ErrorResponse "Bad request or invalid ID"
// @Failure 500 {object} response.ErrorResponse "Failed to trashed card"
// @Router /api/card-command/trashed/{id} [post]
func (h *cardCommandHandleApi) TrashedCard(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		return errors.NewBadRequestError("id is required")
	}

	ctx := c.Request().Context()

	req := &pb.FindByIdCardRequest{
		CardId: int32(idInt),
	}

	res, err := h.card.TrashedCard(ctx, req)

	if err != nil {
		return errors.ParseGrpcError(err)
	}

	so := h.mapper.ToApiResponseCardDeleteAt(res)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Restore a card
// @Tags Card Command
// @Description Restore a card by its ID
// @Accept json
// @Produce json
// @Param id path int true "Card ID"
// @Success 200 {object} response.ApiResponseCard "Restored card"
// @Failure 400 {object} response.ErrorResponse "Bad request or invalid ID"
// @Failure 500 {object} response.ErrorResponse "Failed to restore card"
// @Router /api/card-command/restore/{id} [post]
func (h *cardCommandHandleApi) RestoreCard(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		return errors.NewBadRequestError("id is required")
	}

	ctx := c.Request().Context()

	req := &pb.FindByIdCardRequest{
		CardId: int32(idInt),
	}

	res, err := h.card.RestoreCard(ctx, req)

	if err != nil {
		return errors.ParseGrpcError(err)
	}

	so := h.mapper.ToApiResponseCardDeleteAt(res)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Delete a card permanently
// @Tags Card Command
// @Description Delete a card by its ID permanently
// @Accept json
// @Produce json
// @Param id path int true "Card ID"
// @Success 200 {object} response.ApiResponseCardDelete "Deleted card"
// @Failure 400 {object} response.ErrorResponse "Bad request or invalid ID"
// @Failure 500 {object} response.ErrorResponse "Failed to delete card"
// @Router /api/card-command/permanent/{id} [delete]
func (h *cardCommandHandleApi) DeleteCardPermanent(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		return errors.NewBadRequestError("id is required")
	}

	ctx := c.Request().Context()

	req := &pb.FindByIdCardRequest{
		CardId: int32(idInt),
	}

	res, err := h.card.DeleteCardPermanent(ctx, req)

	if err != nil {
		return errors.ParseGrpcError(err)
	}

	so := h.mapper.ToApiResponseCardDelete(res)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Restore all card records
// @Tags Card Command
// @Description Restore all card records that were previously deleted.
// @Accept json
// @Produce json
// @Success 200 {object} response.ApiResponseCardAll "Successfully restored all card records"
// @Failure 500 {object} response.ErrorResponse "Failed to restore all card records"
// @Router /api/card-command/restore/all [post]
func (h *cardCommandHandleApi) RestoreAllCard(c echo.Context) error {
	ctx := c.Request().Context()

	res, err := h.card.RestoreAllCard(ctx, &emptypb.Empty{})
	if err != nil {
		return errors.ParseGrpcError(err)
	}

	h.logger.Debug("Successfully restored all cards")

	so := h.mapper.ToApiResponseCardAll(res)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer.
// @Summary Permanently delete all card records
// @Tags Card Command
// @Description Permanently delete all card records from the database.
// @Accept json
// @Produce json
// @Success 200 {object} response.ApiResponseCardAll "Successfully deleted all card records permanently"
// @Failure 500 {object} response.ErrorResponse "Failed to permanently delete all card records"
// @Router /api/card-command/permanent/all [post]
func (h *cardCommandHandleApi) DeleteAllCardPermanent(c echo.Context) error {
	ctx := c.Request().Context()

	res, err := h.card.DeleteAllCardPermanent(ctx, &emptypb.Empty{})

	if err != nil {
		return errors.ParseGrpcError(err)
	}

	h.logger.Debug("Successfully deleted all cards permanently")

	so := h.mapper.ToApiResponseCardAll(res)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Toggle card status
// @Tags Card Command
// @Description Toggle a card's status between active and suspended
// @Accept json
// @Produce json
// @Param ToggleCardStatusRequest body requests.ToggleCardStatusRequest true "Toggle card status request"
// @Success 200 {object} response.ApiResponseCard "Toggled card status"
// @Failure 400 {object} response.ErrorResponse "Bad request or validation error"
// @Failure 500 {object} response.ErrorResponse "Failed to toggle card status"
// @Router /api/card-command/toggle-status [post]
func (h *cardCommandHandleApi) ToggleCardStatus(c echo.Context) error {
	var body requests.ToggleCardStatusRequest

	if err := c.Bind(&body); err != nil {
		return errors.NewBadRequestError("Invalid request")
	}

	if err := body.Validate(); err != nil {
		validations := h.parseValidationErrors(err)
		return errors.NewValidationError(validations)
	}

	ctx := c.Request().Context()

	req := &pb.ToggleCardStatusRequest{
		CardId: int32(body.CardID),
	}

	res, err := h.card.ToggleCardStatus(ctx, req)

	if err != nil {
		return errors.ParseGrpcError(err)
	}

	so := h.mapper.ToApiResponseCard(res)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Update credit limit
// @Tags Card Command
// @Description Update the credit limit for a card
// @Accept json
// @Produce json
// @Param id path int true "Card ID"
// @Param UpdateCreditLimitRequest body requests.UpdateCreditLimitRequest true "Update credit limit request"
// @Success 200 {object} response.ApiResponseCard "Updated credit limit"
// @Failure 400 {object} response.ErrorResponse "Bad request or validation error"
// @Failure 500 {object} response.ErrorResponse "Failed to update credit limit"
// @Router /api/card-command/update-credit-limit/{id} [post]
func (h *cardCommandHandleApi) UpdateCreditLimit(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		return errors.NewBadRequestError("id is required")
	}

	var body requests.UpdateCreditLimitRequest

	if err := c.Bind(&body); err != nil {
		return errors.NewBadRequestError("Invalid request")
	}

	if err := body.Validate(); err != nil {
		validations := h.parseValidationErrors(err)
		return errors.NewValidationError(validations)
	}

	ctx := c.Request().Context()

	req := &pb.UpdateCreditLimitRequest{
		CardId:      int32(idInt),
		CreditLimit: int32(body.CreditLimit),
	}

	res, err := h.card.UpdateCreditLimit(ctx, req)

	if err != nil {
		return errors.ParseGrpcError(err)
	}

	so := h.mapper.ToApiResponseCard(res)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Redeem reward points
// @Tags Card Command
// @Description Redeem reward points from a card
// @Accept json
// @Produce json
// @Param id path int true "Card ID"
// @Param RedeemPointsRequest body requests.RedeemPointsRequest true "Redeem points request"
// @Success 200 {object} response.ApiResponseCard "Redeemed reward points"
// @Failure 400 {object} response.ErrorResponse "Bad request or validation error"
// @Failure 500 {object} response.ErrorResponse "Failed to redeem reward points"
// @Router /api/card-command/redeem-points/{id} [post]
func (h *cardCommandHandleApi) RedeemPoints(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		return errors.NewBadRequestError("id is required")
	}

	var body requests.RedeemPointsRequest

	if err := c.Bind(&body); err != nil {
		return errors.NewBadRequestError("Invalid request")
	}

	if err := body.Validate(); err != nil {
		validations := h.parseValidationErrors(err)
		return errors.NewValidationError(validations)
	}

	ctx := c.Request().Context()

	req := &pb.RedeemPointsRequest{
		CardId: int32(idInt),
		Points: int32(body.Points),
	}

	res, err := h.card.RedeemPoints(ctx, req)

	if err != nil {
		return errors.ParseGrpcError(err)
	}

	so := h.mapper.ToApiResponseCard(res)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Process billing cycles
// @Tags Card Command
// @Description Process billing cycles for all credit cards
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{} "Billing cycles processed"
// @Failure 500 {object} response.ErrorResponse "Failed to process billing cycles"
// @Router /api/card-command/process-billing [post]
func (h *cardCommandHandleApi) ProcessBillingCycles(c echo.Context) error {
	ctx := c.Request().Context()

	_, err := h.card.ProcessBillingCycles(ctx, &emptypb.Empty{})

	if err != nil {
		return errors.ParseGrpcError(err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "Successfully processed billing cycles",
	})
}

func (h *cardCommandHandleApi) parseValidationErrors(err error) []errors.ValidationError {
	var validationErrs []errors.ValidationError

	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, fe := range ve {
			validationErrs = append(validationErrs, errors.ValidationError{
				Field:   fe.Field(),
				Message: h.getValidationMessage(fe),
			})
		}
		return validationErrs
	}

	return []errors.ValidationError{
		{
			Field:   "general",
			Message: err.Error(),
		},
	}
}

func (h *cardCommandHandleApi) getValidationMessage(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email format"
	case "min":
		return fmt.Sprintf("Must be at least %s", fe.Param())
	case "max":
		return fmt.Sprintf("Must be at most %s", fe.Param())
	case "gte":
		return fmt.Sprintf("Must be greater than or equal to %s", fe.Param())
	case "lte":
		return fmt.Sprintf("Must be less than or equal to %s", fe.Param())
	case "oneof":
		return fmt.Sprintf("Must be one of: %s", fe.Param())
	default:
		return fmt.Sprintf("Validation failed on '%s' tag", fe.Tag())
	}
}
