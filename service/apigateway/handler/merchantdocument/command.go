package merchantdocumenthandler

import (
	"fmt"
	"net/http"
	"strconv"

	merchantdocument_cache "github.com/MamangRust/monolith-payment-gateway-apigateway/redis/api/merchantdocument"
	pb "github.com/MamangRust/monolith-payment-gateway-pb/merchant_document"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
	merchantdocumentapimapper "github.com/MamangRust/monolith-payment-gateway-shared/mapper/merchantdocument"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"google.golang.org/protobuf/types/known/emptypb"
)

type merchantCommandDocumentHandleApi struct {
	merchantDocument pb.MerchantDocumentCommandServiceClient

	logger logger.LoggerInterface

	mapper merchantdocumentapimapper.MerchantDocumentCommandResponseMapper

	cache merchantdocument_cache.MerchantDocumentQueryCache

	apiHandler errors.ApiHandler
}

type merchantCommandDocumentHandleDeps struct {
	client pb.MerchantDocumentCommandServiceClient

	router *echo.Echo

	logger logger.LoggerInterface

	mapper merchantdocumentapimapper.MerchantDocumentCommandResponseMapper

	cache merchantdocument_cache.MerchantDocumentQueryCache

	apiHandler errors.ApiHandler
}

func NewMerchantCommandDocumentHandler(params *merchantCommandDocumentHandleDeps) {

	merchantDocumentHandler := &merchantCommandDocumentHandleApi{
		merchantDocument: params.client,
		logger:           params.logger,
		mapper:           params.mapper,
		cache:            params.cache,
		apiHandler:       params.apiHandler,
	}

	routerMerchantDocument := params.router.Group("/api/merchant-document-command")

	routerMerchantDocument.POST("/create", merchantDocumentHandler.Create)
	routerMerchantDocument.POST("/update/:id", merchantDocumentHandler.Update)
	routerMerchantDocument.POST("/update-status/:id", merchantDocumentHandler.UpdateStatus)

	routerMerchantDocument.POST("/trashed/:id", merchantDocumentHandler.TrashedDocument)
	routerMerchantDocument.POST("/restore/:id", merchantDocumentHandler.RestoreDocument)
	routerMerchantDocument.DELETE("/permanent/:id", merchantDocumentHandler.Delete)

	routerMerchantDocument.POST("/restore/all", merchantDocumentHandler.RestoreAllDocuments)
	routerMerchantDocument.POST("/permanent/all", merchantDocumentHandler.DeleteAllDocumentsPermanent)
}

// Create godoc
// @Summary Create a new merchant document
// @Tags Merchant Document Command
// @Security Bearer
// @Description Create a new document for a merchant
// @Accept json
// @Produce json
// @Param body body requests.CreateMerchantDocumentRequest true "Create merchant document request"
// @Success 200 {object} response.ApiResponseMerchantDocument "Created document"
// @Failure 400 {object} response.ErrorResponse "Bad request or validation error"
// @Failure 500 {object} response.ErrorResponse "Failed to create document"
// @Router /api/merchant-document-command/create [post]
func (h *merchantCommandDocumentHandleApi) Create(c echo.Context) error {
	var body requests.CreateMerchantDocumentRequest

	if err := c.Bind(&body); err != nil {
		return errors.NewBadRequestError("Invalid request format").WithInternal(err)
	}

	if err := body.Validate(); err != nil {
		validations := h.parseValidationErrors(err)
		return errors.NewValidationError(validations)
	}

	reqPb := &pb.CreateMerchantDocumentRequest{
		MerchantId:   int32(body.MerchantID),
		DocumentType: body.DocumentType,
		DocumentUrl:  body.DocumentUrl,
	}

	ctx := c.Request().Context()

	res, err := h.merchantDocument.Create(ctx, reqPb)
	if err != nil {
		return errors.ParseGrpcError(err)
	}

	apiResponse := h.mapper.ToApiResponseMerchantDocument(res)

	return c.JSON(http.StatusOK, apiResponse)
}

// Update godoc
// @Summary Update a merchant document
// @Tags Merchant Document Command
// @Security Bearer
// @Description Update a merchant document with the given ID
// @Accept json
// @Produce json
// @Param id path int true "Document ID"
// @Param body body requests.UpdateMerchantDocumentRequest true "Update merchant document request"
// @Success 200 {object} response.ApiResponseMerchantDocument "Updated document"
// @Failure 400 {object} response.ErrorResponse "Bad request or validation error"
// @Failure 500 {object} response.ErrorResponse "Failed to update document"
// @Router /api/merchant-document-command/update/{id} [post]
func (h *merchantCommandDocumentHandleApi) Update(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return errors.NewBadRequestError("Invalid ID format").WithInternal(err)
	}

	var body requests.UpdateMerchantDocumentRequest
	if err := c.Bind(&body); err != nil {
		return errors.NewBadRequestError("Invalid request format").WithInternal(err)
	}

	if err := body.Validate(); err != nil {
		validations := h.parseValidationErrors(err)
		return errors.NewValidationError(validations)
	}

	reqPb := &pb.UpdateMerchantDocumentRequest{
		DocumentId:   int32(id),
		MerchantId:   int32(body.MerchantID),
		DocumentType: body.DocumentType,
		DocumentUrl:  body.DocumentUrl,
		Status:       body.Status,
		Note:         body.Note,
	}

	ctx := c.Request().Context()

	res, err := h.merchantDocument.Update(ctx, reqPb)
	if err != nil {
		return errors.ParseGrpcError(err)
	}

	apiResponse := h.mapper.ToApiResponseMerchantDocument(res)

	return c.JSON(http.StatusOK, apiResponse)
}

// UpdateStatus godoc
// @Summary Update merchant document status
// @Tags Merchant Document Command
// @Security Bearer
// @Description Update the status of a merchant document
// @Accept json
// @Produce json
// @Param id path int true "Document ID"
// @Param body body requests.UpdateMerchantDocumentStatusRequest true "Update status request"
// @Success 200 {object} response.ApiResponseMerchantDocument "Updated document"
// @Failure 400 {object} response.ErrorResponse "Bad request or validation error"
// @Failure 500 {object} response.ErrorResponse "Failed to update document status"
// @Router /api/merchants-documents/update-status/{id} [post]
func (h *merchantCommandDocumentHandleApi) UpdateStatus(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return errors.NewBadRequestError("Invalid ID format").WithInternal(err)
	}

	var body requests.UpdateMerchantDocumentStatusRequest
	if err := c.Bind(&body); err != nil {
		return errors.NewBadRequestError("Invalid request format").WithInternal(err)
	}

	if err := body.Validate(); err != nil {
		validations := h.parseValidationErrors(err)
		return errors.NewValidationError(validations)
	}

	reqPb := &pb.UpdateMerchantDocumentStatusRequest{
		DocumentId: int32(id),
		MerchantId: int32(body.MerchantID),
		Status:     body.Status,
		Note:       body.Note,
	}

	ctx := c.Request().Context()

	res, err := h.merchantDocument.UpdateStatus(ctx, reqPb)
	if err != nil {
		return errors.ParseGrpcError(err)
	}

	apiResponse := h.mapper.ToApiResponseMerchantDocument(res)

	return c.JSON(http.StatusOK, apiResponse)
}

// TrashedDocument godoc
// @Summary Trashed a merchant document
// @Tags Merchant Document Command
// @Security Bearer
// @Description Trashed a merchant document by its ID
// @Accept json
// @Produce json
// @Param id path int true "Document ID"
// @Success 200 {object} response.ApiResponseMerchantDocument "Trashed document"
// @Failure 400 {object} response.ErrorResponse "Bad request or invalid ID"
// @Failure 500 {object} response.ErrorResponse "Failed to trashed document"
// @Router /api/merchant-document-command/trashed/{id} [post]
func (h *merchantCommandDocumentHandleApi) TrashedDocument(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return errors.NewBadRequestError("Invalid ID format").WithInternal(err)
	}

	reqPb := &pb.FindMerchantDocumentByIdRequest{
		DocumentId: int32(id),
	}

	ctx := c.Request().Context()

	res, err := h.merchantDocument.Trashed(ctx, reqPb)
	if err != nil {
		return errors.ParseGrpcError(err)
	}

	apiResponse := h.mapper.ToApiResponseMerchantDocumentDeleteAt(res)

	return c.JSON(http.StatusOK, apiResponse)
}

// RestoreDocument godoc
// @Summary Restore a merchant document
// @Tags Merchant Document Command
// @Security Bearer
// @Description Restore a merchant document by its ID
// @Accept json
// @Produce json
// @Param id path int true "Document ID"
// @Success 200 {object} response.ApiResponseMerchantDocument "Restored document"
// @Failure 400 {object} response.ErrorResponse "Bad request or invalid ID"
// @Failure 500 {object} response.ErrorResponse "Failed to restore document"
// @Router /api/merchant-document-command/restore/{id} [post]
func (h *merchantCommandDocumentHandleApi) RestoreDocument(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return errors.NewBadRequestError("Invalid ID format").WithInternal(err)
	}

	reqPb := &pb.FindMerchantDocumentByIdRequest{
		DocumentId: int32(id),
	}

	ctx := c.Request().Context()

	res, err := h.merchantDocument.Restore(ctx, reqPb)
	if err != nil {
		return errors.ParseGrpcError(err)
	}

	apiResponse := h.mapper.ToApiResponseMerchantDocumentDeleteAt(res)

	return c.JSON(http.StatusOK, apiResponse)
}

// Delete godoc
// @Summary Delete a merchant document
// @Tags Merchant Document Command
// @Security Bearer
// @Description Delete a merchant document by its ID
// @Accept json
// @Produce json
// @Param id path int true "Document ID"
// @Success 200 {object} response.ApiResponseMerchantDocumentDelete "Deleted document"
// @Failure 400 {object} response.ErrorResponse "Invalid ID"
// @Failure 500 {object} response.ErrorResponse "Failed to delete document"
// @Router /api/merchant-document-command/permanent/{id} [delete]
func (h *merchantCommandDocumentHandleApi) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return errors.NewBadRequestError("Invalid ID format").WithInternal(err)
	}

	reqPb := &pb.FindMerchantDocumentByIdRequest{
		DocumentId: int32(id),
	}

	ctx := c.Request().Context()

	res, err := h.merchantDocument.DeletePermanent(ctx, reqPb)
	if err != nil {
		return errors.ParseGrpcError(err)
	}

	apiResponse := h.mapper.ToApiResponseMerchantDocumentDelete(res)

	return c.JSON(http.StatusOK, apiResponse)
}

// RestoreAllDocuments godoc
// @Summary Restore all merchant documents
// @Tags Merchant Document Command
// @Security Bearer
// @Description Restore all merchant documents that were previously deleted
// @Accept json
// @Produce json
// @Success 200 {object} response.ApiResponseMerchantDocumentAll "Successfully restored all documents"
// @Failure 500 {object} response.ErrorResponse "Failed to restore all documents"
// @Router /api/merchant-document-command/restore/all [post]
func (h *merchantCommandDocumentHandleApi) RestoreAllDocuments(c echo.Context) error {
	ctx := c.Request().Context()

	res, err := h.merchantDocument.RestoreAll(ctx, &emptypb.Empty{})
	if err != nil {
		return errors.ParseGrpcError(err)
	}

	apiResponse := h.mapper.ToApiResponseMerchantDocumentAll(res)

	return c.JSON(http.StatusOK, apiResponse)
}

// DeleteAllDocumentsPermanent godoc
// @Summary Permanently delete all merchant documents
// @Tags Merchant Document Command
// @Security Bearer
// @Description Permanently delete all merchant documents from the database
// @Accept json
// @Produce json
// @Success 200 {object} response.ApiResponseMerchantDocumentAll "Successfully deleted all documents permanently"
// @Failure 500 {object} response.ErrorResponse "Failed to permanently delete all documents"
// @Router /api/merchant-document-command/permanent/all [post]
func (h *merchantCommandDocumentHandleApi) DeleteAllDocumentsPermanent(c echo.Context) error {
	ctx := c.Request().Context()

	res, err := h.merchantDocument.DeleteAllPermanent(ctx, &emptypb.Empty{})
	if err != nil {
		return errors.ParseGrpcError(err)
	}

	apiResponse := h.mapper.ToApiResponseMerchantDocumentAll(res)

	return c.JSON(http.StatusOK, apiResponse)
}

func (h *merchantCommandDocumentHandleApi) parseValidationErrors(err error) []errors.ValidationError {
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

func (h *merchantCommandDocumentHandleApi) getValidationMessage(fe validator.FieldError) string {
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
