package withdrawhandler

import (
	"fmt"
	"net/http"
	"strconv"

	withdraw_cache "github.com/MamangRust/monolith-payment-gateway-apigateway/redis/api/withdraw"
	pb "github.com/MamangRust/monolith-payment-gateway-pb/withdraw"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
	apimapper "github.com/MamangRust/monolith-payment-gateway-shared/mapper/withdraw"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type withdrawCommandHandleApi struct {
	client pb.WithdrawCommandServiceClient

	logger logger.LoggerInterface

	mapper apimapper.WithdrawCommandResponseMapper

	cache withdraw_cache.WithdrawMencache

	apiHandler errors.ApiHandler
}

type withdrawCommandHandleDeps struct {
	client pb.WithdrawCommandServiceClient

	router *echo.Echo

	logger logger.LoggerInterface

	mapper apimapper.WithdrawCommandResponseMapper

	cache withdraw_cache.WithdrawMencache

	apiHandler errors.ApiHandler
}

func NewWithdrawCommandHandleApi(params *withdrawCommandHandleDeps) *withdrawCommandHandleApi {

	withdrawCommandHandleApi := &withdrawCommandHandleApi{
		client:     params.client,
		logger:     params.logger,
		mapper:     params.mapper,
		cache:      params.cache,
		apiHandler: params.apiHandler,
	}

	routerWithdraw := params.router.Group("/api/withdraw-command")

	routerWithdraw.POST("/create", params.apiHandler.Handle("create-withdraw", withdrawCommandHandleApi.Create))
	routerWithdraw.POST("/update/:id", params.apiHandler.Handle("update-withdraw", withdrawCommandHandleApi.Update))

	routerWithdraw.POST("/trashed/:id", params.apiHandler.Handle("trash-withdraw", withdrawCommandHandleApi.TrashWithdraw))
	routerWithdraw.POST("/restore/:id", params.apiHandler.Handle("restore-withdraw", withdrawCommandHandleApi.RestoreWithdraw))
	routerWithdraw.DELETE("/permanent/:id", params.apiHandler.Handle("delete-withdraw-permanent", withdrawCommandHandleApi.DeleteWithdrawPermanent))

	routerWithdraw.POST("/restore/all", params.apiHandler.Handle("restore-all-withdraws", withdrawCommandHandleApi.RestoreAllWithdraw))
	routerWithdraw.POST("/permanent/all", params.apiHandler.Handle("delete-all-withdraws-permanent", withdrawCommandHandleApi.DeleteAllWithdrawPermanent))

	return withdrawCommandHandleApi
}

// @Summary Create a new withdraw
// @Tags Withdraw Command
// @Security Bearer
// @Description Create a new withdraw record with the provided details.
// @Accept json
// @Produce json
// @Param CreateWithdrawRequest body requests.CreateWithdrawRequest true "Create Withdraw Request"
// @Success 200 {object} response.ApiResponseWithdraw "Successfully created withdraw record"
// @Failure 400 {object} response.ErrorResponse "Bad Request: Invalid request body or validation error"
// @Failure 500 {object} response.ErrorResponse "Failed to create withdraw"
// @Router /api/withdraw-command/create [post]
func (h *withdrawCommandHandleApi) Create(c echo.Context) error {
	var body requests.CreateWithdrawRequest

	if err := c.Bind(&body); err != nil {
		return errors.NewBadRequestError("Invalid request")
	}

	if err := body.Validate(); err != nil {
		validations := h.parseValidationErrors(err)
		return errors.NewValidationError(validations)
	}

	ctx := c.Request().Context()

	res, err := h.client.CreateWithdraw(ctx, &pb.CreateWithdrawRequest{
		CardNumber:     body.CardNumber,
		WithdrawAmount: int32(body.WithdrawAmount),
		WithdrawTime:   timestamppb.New(body.WithdrawTime),
		IdempotencyKey: body.IdempotencyKey,
	})

	if err != nil {
		h.logger.Debug("Failed to create withdraw", zap.Error(err))
		return errors.ParseGrpcError(err)
	}

	so := h.mapper.ToApiResponseWithdraw(res)
	h.cache.SetCachedWithdrawCache(ctx, so)

	return c.JSON(http.StatusOK, so)
}

// @Summary Update an existing withdraw
// @Tags Withdraw Command
// @Security Bearer
// @Description Update an existing withdraw record with the provided details.
// @Accept json
// @Produce json
// @Param id path int true "Withdraw ID"
// @Param UpdateWithdrawRequest body requests.UpdateWithdrawRequest true "Update Withdraw Request"
// @Success 200 {object} response.ApiResponseWithdraw "Successfully updated withdraw record"
// @Failure 400 {object} response.ErrorResponse "Bad Request: Invalid request body or validation error"
// @Failure 500 {object} response.ErrorResponse "Failed to update withdraw"
// @Router /api/withdraw-command/update/{id} [post]
func (h *withdrawCommandHandleApi) Update(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return errors.NewBadRequestError("id is required")
	}

	var body requests.UpdateWithdrawRequest

	if err := c.Bind(&body); err != nil {
		return errors.NewBadRequestError("Invalid request")
	}

	if err := body.Validate(); err != nil {
		validations := h.parseValidationErrors(err)
		return errors.NewValidationError(validations)
	}

	ctx := c.Request().Context()

	res, err := h.client.UpdateWithdraw(ctx, &pb.UpdateWithdrawRequest{
		WithdrawId:     int32(id),
		CardNumber:     body.CardNumber,
		WithdrawAmount: int32(body.WithdrawAmount),
		WithdrawTime:   timestamppb.New(body.WithdrawTime),
	})

	if err != nil {
		h.logger.Debug("Failed to update withdraw", zap.Error(err))
		return errors.ParseGrpcError(err)
	}

	so := h.mapper.ToApiResponseWithdraw(res)

	h.cache.DeleteCachedWithdrawCache(ctx, id)
	h.cache.SetCachedWithdrawCache(ctx, so)

	return c.JSON(http.StatusOK, so)
}

// @Summary Trash a withdraw by ID
// @Tags Withdraw Command
// @Security Bearer
// @Description Trash a withdraw using its ID
// @Accept json
// @Produce json
// @Param id path int true "Withdraw ID"
// @Success 200 {object} response.ApiResponseWithdraw "Withdaw data"
// @Failure 400 {object} response.ErrorResponse "Bad Request: Invalid ID"
// @Failure 500 {object} response.ErrorResponse "Failed to trash withdraw"
// @Router /api/withdraw-command/trashed/{id} [post]
func (h *withdrawCommandHandleApi) TrashWithdraw(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return errors.NewBadRequestError("id is required")
	}

	ctx := c.Request().Context()

	res, err := h.client.TrashedWithdraw(ctx, &pb.FindByIdWithdrawRequest{
		WithdrawId: int32(id),
	})

	if err != nil {
		h.logger.Debug("Failed to trash withdraw", zap.Error(err))
		return errors.ParseGrpcError(err)
	}

	so := h.mapper.ToApiResponseWithdrawDeleteAt(res)

	h.cache.DeleteCachedWithdrawCache(ctx, id)

	return c.JSON(http.StatusOK, so)
}

// @Summary Restore a withdraw by ID
// @Tags Withdraw Command
// @Security Bearer
// @Description Restore a withdraw by its ID
// @Accept json
// @Produce json
// @Param id path int true "Withdraw ID"
// @Success 200 {object} response.ApiResponseWithdraw "Withdraw data"
// @Failure 400 {object} response.ErrorResponse "Bad Request: Invalid ID"
// @Failure 500 {object} response.ErrorResponse "Failed to restore withdraw"
// @Router /api/withdraw-command/restore/{id} [post]
func (h *withdrawCommandHandleApi) RestoreWithdraw(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return errors.NewBadRequestError("id is required")
	}

	ctx := c.Request().Context()

	res, err := h.client.RestoreWithdraw(ctx, &pb.FindByIdWithdrawRequest{
		WithdrawId: int32(id),
	})

	if err != nil {
		h.logger.Debug("Failed to restore withdraw", zap.Error(err))
		return errors.ParseGrpcError(err)
	}

	so := h.mapper.ToApiResponseWithdrawDeleteAt(res)

	h.cache.DeleteCachedWithdrawCache(ctx, id)

	return c.JSON(http.StatusOK, so)
}

// @Summary Permanently delete a withdraw by ID
// @Tags Withdraw Command
// @Security Bearer
// @Description Permanently delete a withdraw by its ID
// @Accept json
// @Produce json
// @Param id path int true "Withdraw ID"
// @Success 200 {object} response.ApiResponseWithdrawDelete "Successfully deleted withdraw permanently"
// @Failure 400 {object} response.ErrorResponse "Bad Request: Invalid ID"
// @Failure 500 {object} response.ErrorResponse "Failed to delete withdraw permanently:"
// @Router /api/withdraw-command/permanent/{id} [delete]
func (h *withdrawCommandHandleApi) DeleteWithdrawPermanent(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return errors.NewBadRequestError("id is required")
	}

	ctx := c.Request().Context()

	res, err := h.client.DeleteWithdrawPermanent(ctx, &pb.FindByIdWithdrawRequest{
		WithdrawId: int32(id),
	})

	if err != nil {
		h.logger.Debug("Failed to delete withdraw permanent", zap.Error(err))
		return errors.ParseGrpcError(err)
	}

	so := h.mapper.ToApiResponseWithdrawDelete(res)

	h.cache.DeleteCachedWithdrawCache(ctx, id)

	return c.JSON(http.StatusOK, so)
}

// @Summary Restore a withdraw all
// @Tags Withdraw Command
// @Security Bearer
// @Description Restore a withdraw all
// @Accept json
// @Produce json
// @Success 200 {object} response.ApiResponseWithdrawAll "Withdraw data"
// @Failure 400 {object} response.ErrorResponse "Bad Request: Invalid ID"
// @Failure 500 {object} response.ErrorResponse "Failed to restore withdraw"
// @Router /api/withdraw-command/restore/all [post]
func (h *withdrawCommandHandleApi) RestoreAllWithdraw(c echo.Context) error {
	ctx := c.Request().Context()

	res, err := h.client.RestoreAllWithdraw(ctx, &emptypb.Empty{})

	if err != nil {
		h.logger.Debug("Failed to restore all withdraw", zap.Error(err))
		return errors.ParseGrpcError(err)
	}

	h.logger.Debug("Successfully restored all withdraw")

	so := h.mapper.ToApiResponseWithdrawAll(res)

	return c.JSON(http.StatusOK, so)
}

// @Summary Permanently delete a withdraw by ID
// @Tags Withdraw Command
// @Security Bearer
// @Description Permanently delete a withdraw by its ID
// @Accept json
// @Produce json
// @Success 200 {object} response.ApiResponseWithdrawAll "Successfully deleted withdraw permanently"
// @Failure 400 {object} response.ErrorResponse "Bad Request: Invalid ID"
// @Failure 500 {object} response.ErrorResponse "Failed to delete withdraw permanently:"
// @Router /api/withdraw-command/permanent/all [post]
func (h *withdrawCommandHandleApi) DeleteAllWithdrawPermanent(c echo.Context) error {
	ctx := c.Request().Context()

	res, err := h.client.DeleteAllWithdrawPermanent(ctx, &emptypb.Empty{})

	if err != nil {
		h.logger.Debug("Failed to delete all withdraw permanent", zap.Error(err))
		return errors.ParseGrpcError(err)
	}

	h.logger.Debug("Successfully deleted all withdraw permanently")

	so := h.mapper.ToApiResponseWithdrawAll(res)

	return c.JSON(http.StatusOK, so)
}

func (h *withdrawCommandHandleApi) parseValidationErrors(err error) []errors.ValidationError {
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

func (h *withdrawCommandHandleApi) getValidationMessage(fe validator.FieldError) string {
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
