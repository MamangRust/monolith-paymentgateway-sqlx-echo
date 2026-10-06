package saldohandler

import (
	"fmt"
	"net/http"
	"strconv"

	saldo_cache "github.com/MamangRust/monolith-payment-gateway-apigateway/redis/api/saldo"
	pb "github.com/MamangRust/monolith-payment-gateway-pb/saldo"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
	apimapper "github.com/MamangRust/monolith-payment-gateway-shared/mapper/saldo"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"
)

type saldoCommandHandleApi struct {
	saldo pb.SaldoCommandServiceClient

	logger logger.LoggerInterface

	mapper apimapper.SaldoCommandResponseMapper

	cache saldo_cache.SaldoMencache

	apiHandler errors.ApiHandler
}

type saldoCommandHandleDeps struct {
	client pb.SaldoCommandServiceClient

	router *echo.Echo

	logger logger.LoggerInterface

	mapper apimapper.SaldoCommandResponseMapper

	cache saldo_cache.SaldoMencache

	apiHandler errors.ApiHandler
}

func NewSaldoCommandHandleApi(params *saldoCommandHandleDeps) *saldoCommandHandleApi {

	saldoHandler := &saldoCommandHandleApi{
		saldo:      params.client,
		logger:     params.logger,
		mapper:     params.mapper,
		cache:      params.cache,
		apiHandler: params.apiHandler,
	}

	routerSaldo := params.router.Group("/api/saldo-command")

	routerSaldo.POST("/create", params.apiHandler.Handle("create-saldo", saldoHandler.Create))
	routerSaldo.POST("/update/:id", params.apiHandler.Handle("update-saldo", saldoHandler.Update))
	routerSaldo.POST("/trashed/:id", params.apiHandler.Handle("trash-saldo", saldoHandler.TrashSaldo))
	routerSaldo.POST("/restore/:id", params.apiHandler.Handle("restore-saldo", saldoHandler.RestoreSaldo))
	routerSaldo.DELETE("/permanent/:id", params.apiHandler.Handle("delete-saldo-permanent", saldoHandler.Delete))

	routerSaldo.POST("/restore/all", params.apiHandler.Handle("restore-all-saldos", saldoHandler.RestoreAllSaldo))
	routerSaldo.POST("/permanent/all", params.apiHandler.Handle("delete-all-saldos-permanent", saldoHandler.DeleteAllSaldoPermanent))

	return saldoHandler
}

// @Summary Create a new saldo
// @Tags Saldo Command
// @Security Bearer
// @Description Create a new saldo record with the provided card number and total balance.
// @Accept json
// @Produce json
// @Param CreateSaldoRequest body requests.CreateSaldoRequest true "Create Saldo Request"
// @Success 200 {object} response.ApiResponseSaldo "Successfully created saldo record"
// @Failure 400 {object} response.ErrorResponse "Bad Request: Invalid request body or validation error"
// @Failure 500 {object} response.ErrorResponse "Failed to create saldo"
// @Router /api/saldo-command/create [post]
func (h *saldoCommandHandleApi) Create(c echo.Context) error {
	var body requests.CreateSaldoRequest

	if err := c.Bind(&body); err != nil {
		return errors.NewBadRequestError("Invalid request")
	}

	if err := body.Validate(); err != nil {
		validations := h.parseValidationErrors(err)
		return errors.NewValidationError(validations)
	}

	ctx := c.Request().Context()

	res, err := h.saldo.CreateSaldo(ctx, &pb.CreateSaldoRequest{
		CardNumber:   body.CardNumber,
		TotalBalance: int32(body.TotalBalance),
	})

	if err != nil {
		return errors.ParseGrpcError(err)
	}

	so := h.mapper.ToApiResponseSaldo(res)

	h.cache.SetCachedSaldoById(ctx, so.Data.ID, so)

	return c.JSON(http.StatusOK, so)
}

// @Summary Update an existing saldo
// @Tags Saldo Command
// @Security Bearer
// @Description Update an existing saldo record with the provided card number and total balance.
// @Accept json
// @Produce json
// @Param id path int true "Saldo ID"
// @Param UpdateSaldoRequest body requests.UpdateSaldoRequest true "Update Saldo Request"
// @Success 200 {object} response.ApiResponseSaldo "Successfully updated saldo record"
// @Failure 400 {object} response.ErrorResponse "Bad Request: Invalid request body or validation error"
// @Failure 500 {object} response.ErrorResponse "Failed to update saldo"
// @Router /api/saldo-command/update/{id} [post]
func (h *saldoCommandHandleApi) Update(c echo.Context) error {
	idint, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return errors.NewBadRequestError("id is required")
	}

	var body requests.UpdateSaldoRequest

	if err := c.Bind(&body); err != nil {
		return errors.NewBadRequestError("Invalid request")
	}

	if err := body.Validate(); err != nil {
		validations := h.parseValidationErrors(err)
		return errors.NewValidationError(validations)
	}

	ctx := c.Request().Context()

	res, err := h.saldo.UpdateSaldo(ctx, &pb.UpdateSaldoRequest{
		SaldoId:      int32(idint),
		CardNumber:   body.CardNumber,
		TotalBalance: int32(body.TotalBalance),
	})

	if err != nil {
		h.logger.Debug("Failed to update saldo", zap.Error(err))
		return errors.ParseGrpcError(err)
	}

	so := h.mapper.ToApiResponseSaldo(res)

	h.cache.DeleteSaldoCache(ctx, idint)

	h.cache.SetCachedSaldoById(ctx, so.Data.ID, so)

	return c.JSON(http.StatusOK, so)
}

// @Summary Soft delete a saldo
// @Tags Saldo Command
// @Security Bearer
// @Description Soft delete an existing saldo record by its ID.
// @Accept json
// @Produce json
// @Param id path int true "Saldo ID"
// @Success 200 {object} response.ApiResponseSaldo "Successfully trashed saldo record"
// @Failure 400 {object} response.ErrorResponse "Bad Request: Invalid ID"
// @Failure 500 {object} response.ErrorResponse "Failed to trashed saldo"
// @Router /api/saldo-command/trashed/{id} [post]
func (h *saldoCommandHandleApi) TrashSaldo(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		h.logger.Debug("Bad Request: Invalid ID", zap.Error(err))
		return errors.NewBadRequestError("id is required")
	}

	ctx := c.Request().Context()

	res, err := h.saldo.TrashedSaldo(ctx, &pb.FindByIdSaldoRequest{
		SaldoId: int32(idInt),
	})

	if err != nil {
		h.logger.Debug("Failed to trashed saldo", zap.Error(err))
		return errors.ParseGrpcError(err)
	}

	so := h.mapper.ToApiResponseSaldoDeleteAt(res)

	h.cache.DeleteSaldoCache(ctx, idInt)

	return c.JSON(http.StatusOK, so)
}

// @Summary Restore a trashed saldo
// @Tags Saldo Command
// @Security Bearer
// @Description Restore an existing saldo record from the trash by its ID.
// @Accept json
// @Produce json
// @Param id path int true "Saldo ID"
// @Success 200 {object} response.ApiResponseSaldo "Successfully restored saldo record"
// @Failure 400 {object} response.ErrorResponse "Bad Request: Invalid ID"
// @Failure 500 {object} response.ErrorResponse "Failed to restore saldo"
// @Router /api/saldo-command/restore/{id} [post]
func (h *saldoCommandHandleApi) RestoreSaldo(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		h.logger.Debug("Bad Request: Invalid ID", zap.Error(err))
		return errors.NewBadRequestError("id is required")
	}

	ctx := c.Request().Context()

	res, err := h.saldo.RestoreSaldo(ctx, &pb.FindByIdSaldoRequest{
		SaldoId: int32(idInt),
	})

	if err != nil {
		h.logger.Debug("Failed to restore saldo", zap.Error(err))
		return errors.ParseGrpcError(err)
	}

	so := h.mapper.ToApiResponseSaldoDeleteAt(res)

	h.cache.DeleteSaldoCache(ctx, idInt)

	return c.JSON(http.StatusOK, so)
}

// @Summary Permanently delete a saldo
// @Tags Saldo Command
// @Security Bearer
// @Description Permanently delete an existing saldo record by its ID.
// @Accept json
// @Produce json
// @Param id path int true "Saldo ID"
// @Success 200 {object} response.ApiResponseSaldoDelete "Successfully deleted saldo record"
// @Failure 400 {object} response.ErrorResponse "Bad Request: Invalid ID"
// @Failure 500 {object} response.ErrorResponse "Failed to delete saldo"
// @Router /api/saldo-command/permanent/{id} [delete]
func (h *saldoCommandHandleApi) Delete(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		h.logger.Debug("Bad Request: Invalid ID", zap.Error(err))
		return errors.NewBadRequestError("id is required")
	}

	ctx := c.Request().Context()

	res, err := h.saldo.DeleteSaldoPermanent(ctx, &pb.FindByIdSaldoRequest{
		SaldoId: int32(idInt),
	})

	if err != nil {
		h.logger.Debug("Failed to delete saldo", zap.Error(err))
		return errors.ParseGrpcError(err)
	}

	so := h.mapper.ToApiResponseSaldoDelete(res)

	h.cache.DeleteSaldoCache(ctx, idInt)

	return c.JSON(http.StatusOK, so)
}

// RestoreAllSaldo restores all saldo records.
// @Summary Restore all saldo records
// @Tags Saldo Command
// @Security Bearer
// @Description Restore all saldo records that were previously deleted.
// @Accept json
// @Produce json
// @Success 200 {object} response.ApiResponseSaldoAll "Successfully restored all saldo records"
// @Failure 500 {object} response.ErrorResponse "Failed to restore all saldo records"
// @Router /api/saldo-command/restore/all [post]
func (h *saldoCommandHandleApi) RestoreAllSaldo(c echo.Context) error {
	ctx := c.Request().Context()

	res, err := h.saldo.RestoreAllSaldo(ctx, &emptypb.Empty{})

	if err != nil {
		h.logger.Error("Failed to restore all saldo", zap.Error(err))
		return errors.ParseGrpcError(err)
	}

	h.logger.Debug("Successfully restored all saldo")

	so := h.mapper.ToApiResponseSaldoAll(res)

	return c.JSON(http.StatusOK, so)
}

// @Summary Permanently delete all saldo records
// @Tags Saldo Command
// @Security Bearer
// @Description Permanently delete all saldo records from the database.
// @Accept json
// @Produce json
// @Success 200 {object} response.ApiResponseSaldoAll "Successfully deleted all saldo records permanently"
// @Failure 500 {object} response.ErrorResponse "Failed to permanently delete all saldo records"
// @Router /api/saldo-command/permanent/all [post]
func (h *saldoCommandHandleApi) DeleteAllSaldoPermanent(c echo.Context) error {
	ctx := c.Request().Context()

	res, err := h.saldo.DeleteAllSaldoPermanent(ctx, &emptypb.Empty{})

	if err != nil {
		h.logger.Error("Failed to permanently delete all saldo", zap.Error(err))

		return errors.ParseGrpcError(err)
	}

	h.logger.Debug("Successfully deleted all saldo permanently")

	so := h.mapper.ToApiResponseSaldoAll(res)

	return c.JSON(http.StatusOK, so)
}

func (h *saldoCommandHandleApi) parseValidationErrors(err error) []errors.ValidationError {
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

func (h *saldoCommandHandleApi) getValidationMessage(fe validator.FieldError) string {
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
