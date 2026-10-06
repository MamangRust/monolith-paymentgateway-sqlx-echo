package merchantdocumenthandler

import (
	"net/http"
	"strconv"

	merchantdocument_cache "github.com/MamangRust/monolith-payment-gateway-apigateway/redis/api/merchantdocument"
	pb "github.com/MamangRust/monolith-payment-gateway-pb/merchant_document"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
	merchantdocumentapimapper "github.com/MamangRust/monolith-payment-gateway-shared/mapper/merchantdocument"
	"github.com/labstack/echo/v4"
)

type merchantQueryDocumentHandleApi struct {
	merchantDocument pb.MerchantDocumentQueryServiceClient

	logger logger.LoggerInterface

	mapper merchantdocumentapimapper.MerchantDocumentQueryResponseMapper

	cache merchantdocument_cache.MerchantDocumentQueryCache

	apiHandler errors.ApiHandler
}

type merchantDocumentQueryDocumentHandleDeps struct {
	client pb.MerchantDocumentQueryServiceClient

	router *echo.Echo

	logger logger.LoggerInterface

	cache merchantdocument_cache.MerchantDocumentQueryCache

	apiHandler errors.ApiHandler

	mapper merchantdocumentapimapper.MerchantDocumentQueryResponseMapper
}

func NewMerchantQueryDocumentHandler(params *merchantDocumentQueryDocumentHandleDeps) *merchantQueryDocumentHandleApi {

	merchantDocumentHandler := &merchantQueryDocumentHandleApi{
		merchantDocument: params.client,
		logger:           params.logger,
		mapper:           params.mapper,
		cache:            params.cache,
		apiHandler:       params.apiHandler,
	}

	routerMerchantDocument := params.router.Group("/api/merchant-document-query")

	routerMerchantDocument.GET("", merchantDocumentHandler.FindAll)
	routerMerchantDocument.GET("/:id", merchantDocumentHandler.FindById)
	routerMerchantDocument.GET("/active", merchantDocumentHandler.FindAllActive)
	routerMerchantDocument.GET("/trashed", merchantDocumentHandler.FindAllTrashed)

	return merchantDocumentHandler
}

// FindAll godoc
// @Summary Find all merchant documents
// @Tags Merchant Document Query
// @Security Bearer
// @Description Retrieve a list of all merchant documents
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Number of items per page" default(10)
// @Param search query string false "Search query"
// @Success 200 {object} response.ApiResponsePaginationMerchantDocument "List of merchant documents"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve merchant document data"
// @Router /api/merchant-document-query [get]
func (h *merchantQueryDocumentHandleApi) FindAll(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	if page <= 0 {
		page = 1
	}

	pageSize, _ := strconv.Atoi(c.QueryParam("page_size"))
	if pageSize <= 0 {
		pageSize = 10
	}

	search := c.QueryParam("search")

	req := &requests.FindAllMerchantDocuments{
		Page:     page,
		PageSize: pageSize,
		Search:   search,
	}

	ctx := c.Request().Context()

	cachedData, found := h.cache.GetCachedMerchants(ctx, req)
	if found {
		return c.JSON(http.StatusOK, cachedData)
	}

	grpcReq := &pb.FindAllMerchantDocumentsRequest{
		Page:     int32(page),
		PageSize: int32(pageSize),
		Search:   search,
	}

	res, err := h.merchantDocument.FindAll(ctx, grpcReq)
	if err != nil {
		return errors.ParseGrpcError(err)
	}

	apiResponse := h.mapper.ToApiResponsePaginationMerchantDocument(res)

	h.cache.SetCachedMerchants(ctx, req, apiResponse)

	return c.JSON(http.StatusOK, apiResponse)
}

// FindById godoc
// @Summary Get merchant document by ID
// @Tags Merchant Document Query
// @Security Bearer
// @Description Get a merchant document by its ID
// @Accept json
// @Produce json
// @Param id path int true "Document ID"
// @Success 200 {object} response.ApiResponseMerchantDocument "Document details"
// @Failure 400 {object} response.ErrorResponse "Invalid ID"
// @Failure 500 {object} response.ErrorResponse "Failed to get document"
// @Router /api/merchant-document-query/{id} [get]
func (h *merchantQueryDocumentHandleApi) FindById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return errors.NewBadRequestError("Invalid ID format").WithInternal(err)
	}

	ctx := c.Request().Context()

	cachedData, found := h.cache.GetCachedMerchant(ctx, id)
	if found {
		return c.JSON(http.StatusOK, cachedData)
	}

	res, err := h.merchantDocument.FindById(ctx, &pb.FindMerchantDocumentByIdRequest{
		DocumentId: int32(id),
	})
	if err != nil {
		return errors.ParseGrpcError(err)
	}

	apiResponse := h.mapper.ToApiResponseMerchantDocument(res)

	h.cache.SetCachedMerchant(ctx, apiResponse)

	return c.JSON(http.StatusOK, apiResponse)
}

// FindAllActive godoc
// @Summary Find all active merchant documents
// @Tags Merchant Document Query
// @Security Bearer
// @Description Retrieve a list of all active merchant documents
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Number of items per page" default(10)
// @Param search query string false "Search query"
// @Success 200 {object} response.ApiResponsePaginationMerchantDocument "List of active merchant documents"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve active merchant documents"
// @Router /api/merchant-document-query/active [get]
func (h *merchantQueryDocumentHandleApi) FindAllActive(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	if page <= 0 {
		page = 1
	}

	pageSize, _ := strconv.Atoi(c.QueryParam("page_size"))
	if pageSize <= 0 {
		pageSize = 10
	}

	search := c.QueryParam("search")

	req := &requests.FindAllMerchantDocuments{
		Page:     page,
		PageSize: pageSize,
		Search:   search,
	}

	ctx := c.Request().Context()

	cachedData, found := h.cache.GetCachedMerchantActive(ctx, req)
	if found {
		return c.JSON(http.StatusOK, cachedData)
	}

	grpcReq := &pb.FindAllMerchantDocumentsRequest{
		Page:     int32(page),
		PageSize: int32(pageSize),
		Search:   search,
	}

	res, err := h.merchantDocument.FindAllActive(ctx, grpcReq)
	if err != nil {
		return errors.ParseGrpcError(err)
	}

	apiResponse := h.mapper.ToApiResponsePaginationMerchantDocumentDeleteAt(res)

	h.cache.SetCachedMerchantActive(ctx, req, apiResponse)

	return c.JSON(http.StatusOK, apiResponse)
}

// FindAllTrashed godoc
// @Summary Find all trashed merchant documents
// @Tags Merchant Document Query
// @Security Bearer
// @Description Retrieve a list of all trashed merchant documents
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Number of items per page" default(10)
// @Param search query string false "Search query"
// @Success 200 {object} response.ApiResponsePaginationMerchantDocumentDeleteAt "List of trashed merchant documents"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve trashed merchant documents"
// @Router /api/merchant-document-query/trashed [get]
func (h *merchantQueryDocumentHandleApi) FindAllTrashed(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	if page <= 0 {
		page = 1
	}

	pageSize, _ := strconv.Atoi(c.QueryParam("page_size"))
	if pageSize <= 0 {
		pageSize = 10
	}

	search := c.QueryParam("search")

	req := &requests.FindAllMerchantDocuments{
		Page:     page,
		PageSize: pageSize,
		Search:   search,
	}

	ctx := c.Request().Context()

	cachedData, found := h.cache.GetCachedMerchantTrashed(ctx, req)
	if found {
		return c.JSON(http.StatusOK, cachedData)
	}

	grpcReq := &pb.FindAllMerchantDocumentsRequest{
		Page:     int32(page),
		PageSize: int32(pageSize),
		Search:   search,
	}

	res, err := h.merchantDocument.FindAllTrashed(ctx, grpcReq)
	if err != nil {
		return errors.ParseGrpcError(err)
	}

	apiResponse := h.mapper.ToApiResponsePaginationMerchantDocumentDeleteAt(res)

	h.cache.SetCachedMerchantTrashed(ctx, req, apiResponse)

	return c.JSON(http.StatusOK, apiResponse)
}
