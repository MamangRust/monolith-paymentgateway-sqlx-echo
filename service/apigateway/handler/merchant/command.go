package merchanthandler

import (
	"fmt"
	"net/http"
	"strconv"

	merchant_cache "github.com/MamangRust/monolith-payment-gateway-apigateway/redis/api/merchant"
	pb "github.com/MamangRust/monolith-payment-gateway-pb/merchant"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	errors "github.com/MamangRust/monolith-payment-gateway-shared/errors"
	apimapper "github.com/MamangRust/monolith-payment-gateway-shared/mapper/merchant"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"google.golang.org/protobuf/types/known/emptypb"
)

type merchantCommandHandleApi struct {
	client pb.MerchantCommandServiceClient

	logger logger.LoggerInterface

	mapper apimapper.MerchantCommandResponseMapper

	cache merchant_cache.MerchantMencache

	apiHandler errors.ApiHandler
}

type merchantCommandHandleDeps struct {
	client pb.MerchantCommandServiceClient

	router *echo.Echo

	logger logger.LoggerInterface

	mapper apimapper.MerchantCommandResponseMapper

	cache merchant_cache.MerchantMencache

	apiHandler errors.ApiHandler
}

func NewMerchantCommandHandleApi(params *merchantCommandHandleDeps) *merchantCommandHandleApi {

	merchantHandler := &merchantCommandHandleApi{
		client:     params.client,
		logger:     params.logger,
		mapper:     params.mapper,
		cache:      params.cache,
		apiHandler: params.apiHandler,
	}

	routerMerchant := params.router.Group("/api/merchant-command")

	routerMerchant.POST("/create", params.apiHandler.Handle("create-merchant", merchantHandler.Create))
	routerMerchant.POST("/update/:id", params.apiHandler.Handle("update-merchant", merchantHandler.Update))
	routerMerchant.POST("/update-status/:id", params.apiHandler.Handle("update-merchant-status", merchantHandler.UpdateStatus))

	routerMerchant.POST("/trashed/:id", params.apiHandler.Handle("trash-merchant", merchantHandler.TrashedMerchant))
	routerMerchant.POST("/restore/:id", params.apiHandler.Handle("restore-merchant", merchantHandler.RestoreMerchant))
	routerMerchant.DELETE("/permanent/:id", params.apiHandler.Handle("delete-merchant-permanent", merchantHandler.Delete))

	routerMerchant.POST("/restore/all", params.apiHandler.Handle("restore-all-merchants", merchantHandler.RestoreAllMerchant))
	routerMerchant.POST("/permanent/all", params.apiHandler.Handle("delete-all-merchants-permanent", merchantHandler.DeleteAllMerchantPermanent))

	return merchantHandler
}

// Create godoc
// @Summary Create a new merchant
// @Tags Merchant Command
// @Security Bearer
// @Description Create a new merchant with the given name and user ID
// @Accept json
// @Produce json
// @Param body body requests.CreateMerchantRequest true "Create merchant request"
// @Success 200 {object} response.ApiResponseMerchant "Created merchant"
// @Failure 400 {object} response.ErrorResponse "Bad request or validation error"
// @Failure 500 {object} response.ErrorResponse "Failed to create merchant"
// @Router /api/merchant-command/create [post]
func (h *merchantCommandHandleApi) Create(c echo.Context) error {
	var body requests.CreateMerchantRequest

	if err := c.Bind(&body); err != nil {
		return errors.NewBadRequestError("Invalid request")
	}

	if err := body.Validate(); err != nil {
		validations := h.parseValidationErrors(err)
		return errors.NewValidationError(validations)
	}

	ctx := c.Request().Context()

	req := &pb.CreateMerchantRequest{
		Name:   body.Name,
		UserId: int32(body.UserID),
	}

	res, err := h.client.CreateMerchant(ctx, req)

	if err != nil {
		return errors.ParseGrpcError(err)
	}

	so := h.mapper.ToApiResponseMerchant(res)

	return c.JSON(http.StatusOK, so)
}

// Update godoc
// @Summary Update a merchant
// @Tags Merchant Command
// @Security Bearer
// @Description Update a merchant with the given ID
// @Accept json
// @Produce json
// @Param body body requests.UpdateMerchantRequest true "Update merchant request"
// @Success 200 {object} response.ApiResponseMerchant "Updated merchant"
// @Failure 400 {object} response.ErrorResponse "Bad request or validation error"
// @Failure 500 {object} response.ErrorResponse "Failed to update merchant"
// @Router /api/merchant-command/update/{id} [post]
func (h *merchantCommandHandleApi) Update(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return errors.NewBadRequestError("id is required")
	}

	var body requests.UpdateMerchantRequest

	if err := c.Bind(&body); err != nil {
		return errors.NewBadRequestError("Invalid request")
	}

	if err := body.Validate(); err != nil {
		validations := h.parseValidationErrors(err)
		return errors.NewValidationError(validations)
	}

	ctx := c.Request().Context()
	req := &pb.UpdateMerchantRequest{
		MerchantId: int32(id),
		Name:       body.Name,
		UserId:     int32(body.UserID),
		Status:     body.Status,
	}

	res, err := h.client.UpdateMerchant(ctx, req)

	if err != nil {
		return errors.ParseGrpcError(err)
	}

	so := h.mapper.ToApiResponseMerchant(res)

	return c.JSON(http.StatusOK, so)
}

// UpdateStatus godoc
// @Summary Update a merchant status
// @Tags Merchant Command
// @Security Bearer
// @Description Update the status of a merchant with the given ID
// @Accept json
// @Produce json
// @Param id path int true "Merchant ID"
// @Param body body requests.UpdateMerchantStatusRequest true "Update merchant status request"
// @Success 200 {object} response.ApiResponseMerchant "Updated merchant"
// @Failure 400 {object} response.ErrorResponse "Bad request or validation error"
// @Failure 500 {object} response.ErrorResponse "Failed to update merchant status"
// @Router /api/merchant-command/update-status/{id} [post]
func (h *merchantCommandHandleApi) UpdateStatus(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return errors.NewBadRequestError("id is required")
	}

	var body requests.UpdateMerchantStatusRequest

	if err := c.Bind(&body); err != nil {
		return errors.NewBadRequestError("Invalid request")
	}

	if err := body.Validate(); err != nil {
		validations := h.parseValidationErrors(err)
		return errors.NewValidationError(validations)
	}

	ctx := c.Request().Context()
	req := &pb.UpdateMerchantStatusRequest{
		MerchantId: int32(id),
		Status:     body.Status,
	}

	res, err := h.client.UpdateMerchantStatus(ctx, req)

	if err != nil {
		return errors.ParseGrpcError(err)
	}

	so := h.mapper.ToApiResponseMerchant(res)

	return c.JSON(http.StatusOK, so)
}

// TrashedMerchant godoc
// @Summary Trashed a merchant
// @Tags Merchant Command
// @Security Bearer
// @Description Trashed a merchant by its ID
// @Accept json
// @Produce json
// @Param id path int true "Merchant ID"
// @Success 200 {object} response.ApiResponseMerchant "Trashed merchant"
// @Failure 400 {object} response.ErrorResponse "Bad request or invalid ID"
// @Failure 500 {object} response.ErrorResponse "Failed to trashed merchant"
// @Router /api/merchant-command/trashed/{id} [post]
func (h *merchantCommandHandleApi) TrashedMerchant(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		return errors.NewBadRequestError("id is required")
	}

	ctx := c.Request().Context()

	res, err := h.client.TrashedMerchant(ctx, &pb.FindByIdMerchantRequest{
		MerchantId: int32(idInt),
	})

	if err != nil {
		return errors.ParseGrpcError(err)
	}

	so := h.mapper.ToApiResponseMerchantDeleteAt(res)

	return c.JSON(http.StatusOK, so)
}

// RestoreMerchant godoc
// @Summary Restore a merchant
// @Tags Merchant Command
// @Security Bearer
// @Description Restore a merchant by its ID
// @Accept json
// @Produce json
// @Param id path int true "Merchant ID"
// @Success 200 {object} response.ApiResponseMerchant "Restored merchant"
// @Failure 400 {object} response.ErrorResponse "Bad request or invalid ID"
// @Failure 500 {object} response.ErrorResponse "Failed to restore merchant"
// @Router /api/merchant-command/restore/{id} [post]
func (h *merchantCommandHandleApi) RestoreMerchant(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		return errors.NewBadRequestError("id is required")
	}

	ctx := c.Request().Context()

	res, err := h.client.RestoreMerchant(ctx, &pb.FindByIdMerchantRequest{
		MerchantId: int32(idInt),
	})

	if err != nil {
		return errors.ParseGrpcError(err)
	}

	so := h.mapper.ToApiResponseMerchantDeleteAt(res)

	return c.JSON(http.StatusOK, so)
}

// Delete godoc
// @Summary Delete a merchant permanently
// @Tags Merchant Command
// @Security Bearer
// @Description Delete a merchant by its ID permanently
// @Accept json
// @Produce json
// @Param id path int true "Merchant ID"
// @Success 200 {object} response.ApiResponseMerchantDelete "Deleted merchant"
// @Failure 400 {object} response.ErrorResponse "Bad request or invalid ID"
// @Failure 500 {object} response.ErrorResponse "Failed to delete merchant"
// @Router /api/merchant-command/{id} [delete]
func (h *merchantCommandHandleApi) Delete(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		return errors.NewBadRequestError("id is required")
	}

	ctx := c.Request().Context()

	res, err := h.client.DeleteMerchantPermanent(ctx, &pb.FindByIdMerchantRequest{
		MerchantId: int32(idInt),
	})

	if err != nil {
		return errors.ParseGrpcError(err)
	}

	so := h.mapper.ToApiResponseMerchantDelete(res)

	return c.JSON(http.StatusOK, so)
}

// RestoreAllMerchant godoc.
// @Summary Restore all merchant records
// @Tags Merchant Command
// @Security Bearer
// @Description Restore all merchant records that were previously deleted.
// @Accept json
// @Produce json
// @Success 200 {object} response.ApiResponseMerchantAll "Successfully restored all merchant records"
// @Failure 500 {object} response.ErrorResponse "Failed to restore all merchant records"
// @Router /api/merchant-command/restore/all [post]
func (h *merchantCommandHandleApi) RestoreAllMerchant(c echo.Context) error {
	ctx := c.Request().Context()

	res, err := h.client.RestoreAllMerchant(ctx, &emptypb.Empty{})

	if err != nil {
		return errors.ParseGrpcError(err)
	}

	h.logger.Debug("Successfully restored all merchant")

	so := h.mapper.ToApiResponseMerchantAll(res)

	return c.JSON(http.StatusOK, so)
}

// DeleteAllMerchantPermanent godoc.
// @Summary Permanently delete all merchant records
// @Tags Merchant Command
// @Security Bearer
// @Description Permanently delete all merchant records from the database.
// @Accept json
// @Produce json
// @Success 200 {object} response.ApiResponseMerchantAll "Successfully deleted all merchant records permanently"
// @Failure 500 {object} response.ErrorResponse "Failed to permanently delete all merchant records"
// @Router /api/merchant-command/permanent/all [post]
func (h *merchantCommandHandleApi) DeleteAllMerchantPermanent(c echo.Context) error {
	ctx := c.Request().Context()

	res, err := h.client.DeleteAllMerchantPermanent(ctx, &emptypb.Empty{})

	if err != nil {
		return errors.ParseGrpcError(err)
	}

	h.logger.Debug("Successfully deleted all merchant permanently")

	so := h.mapper.ToApiResponseMerchantAll(res)

	return c.JSON(http.StatusOK, so)
}

func (h *merchantCommandHandleApi) parseValidationErrors(err error) []errors.ValidationError {
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

func (h *merchantCommandHandleApi) getValidationMessage(fe validator.FieldError) string {
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
