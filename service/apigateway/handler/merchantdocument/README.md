# 📦 Package `merchantdocumenthandler`

**Source Path:** `apigateway/internal/handler/merchantdocument`

## 🧩 Types

### `DepsMerchantDocument`

```go
type DepsMerchantDocument struct {
	Client *grpc.ClientConn
	E *echo.Echo
	Logger logger.LoggerInterface
}
```

### `merchantCommandDocumentHandleApi`

```go
type merchantCommandDocumentHandleApi struct {
	merchantDocument pb.MerchantDocumentCommandServiceClient
	logger logger.LoggerInterface
	mapper apimapper.MerchantDocumentCommandResponseMapper
	trace trace.Tracer
	requestCounter *prometheus.CounterVec
	requestDuration *prometheus.HistogramVec
}
```

#### Methods

##### `Create`

Create godoc
@Summary Create a new merchant document
@Tags Merchant Document Command
@Security Bearer
@Description Create a new document for a merchant
@Accept json
@Produce json
@Param body body requests.CreateMerchantDocumentRequest true "Create merchant document request"
@Success 200 {object} response.ApiResponseMerchantDocument "Created document"
@Failure 400 {object} response.ErrorResponse "Bad request or validation error"
@Failure 500 {object} response.ErrorResponse "Failed to create document"
@Router /api/merchant-document-command/create [post]

```go
func (h *merchantCommandDocumentHandleApi) Create(c echo.Context) error
```

##### `Delete`

Delete godoc
@Summary Delete a merchant document
@Tags Merchant Document Command
@Security Bearer
@Description Delete a merchant document by its ID
@Accept json
@Produce json
@Param id path int true "Document ID"
@Success 200 {object} response.ApiResponseMerchantDocumentDelete "Deleted document"
@Failure 400 {object} response.ErrorResponse "Invalid ID"
@Failure 500 {object} response.ErrorResponse "Failed to delete document"
@Router /api/merchant-document-command/permanent/{id} [delete]

```go
func (h *merchantCommandDocumentHandleApi) Delete(c echo.Context) error
```

##### `DeleteAllDocumentsPermanent`

DeleteAllDocumentsPermanent godoc
@Summary Permanently delete all merchant documents
@Tags Merchant Document Command
@Security Bearer
@Description Permanently delete all merchant documents from the database
@Accept json
@Produce json
@Success 200 {object} response.ApiResponseMerchantDocumentAll "Successfully deleted all documents permanently"
@Failure 500 {object} response.ErrorResponse "Failed to permanently delete all documents"
@Router /api/merchant-document-command/permanent/all [post]

```go
func (h *merchantCommandDocumentHandleApi) DeleteAllDocumentsPermanent(c echo.Context) error
```

##### `RestoreAllDocuments`

RestoreAllDocuments godoc
@Summary Restore all merchant documents
@Tags Merchant Document Command
@Security Bearer
@Description Restore all merchant documents that were previously deleted
@Accept json
@Produce json
@Success 200 {object} response.ApiResponseMerchantDocumentAll "Successfully restored all documents"
@Failure 500 {object} response.ErrorResponse "Failed to restore all documents"
@Router /api/merchant-document-command/restore/all [post]

```go
func (h *merchantCommandDocumentHandleApi) RestoreAllDocuments(c echo.Context) error
```

##### `RestoreDocument`

RestoreDocument godoc
@Summary Restore a merchant document
@Tags Merchant Document Command
@Security Bearer
@Description Restore a merchant document by its ID
@Accept json
@Produce json
@Param id path int true "Document ID"
@Success 200 {object} response.ApiResponseMerchantDocument "Restored document"
@Failure 400 {object} response.ErrorResponse "Bad request or invalid ID"
@Failure 500 {object} response.ErrorResponse "Failed to restore document"
@Router /api/merchant-document-command/restore/{id} [post]

```go
func (h *merchantCommandDocumentHandleApi) RestoreDocument(c echo.Context) error
```

##### `TrashedDocument`

TrashedDocument godoc
@Summary Trashed a merchant document
@Tags Merchant Document Command
@Security Bearer
@Description Trashed a merchant document by its ID
@Accept json
@Produce json
@Param id path int true "Document ID"
@Success 200 {object} response.ApiResponseMerchantDocument "Trashed document"
@Failure 400 {object} response.ErrorResponse "Bad request or invalid ID"
@Failure 500 {object} response.ErrorResponse "Failed to trashed document"
@Router /api/merchant-document-command/trashed/{id} [post]

```go
func (h *merchantCommandDocumentHandleApi) TrashedDocument(c echo.Context) error
```

##### `Update`

Update godoc
@Summary Update a merchant document
@Tags Merchant Document Command
@Security Bearer
@Description Update a merchant document with the given ID
@Accept json
@Produce json
@Param id path int true "Document ID"
@Param body body requests.UpdateMerchantDocumentRequest true "Update merchant document request"
@Success 200 {object} response.ApiResponseMerchantDocument "Updated document"
@Failure 400 {object} response.ErrorResponse "Bad request or validation error"
@Failure 500 {object} response.ErrorResponse "Failed to update document"
@Router /api/merchant-document-command/update/{id} [post]

```go
func (h *merchantCommandDocumentHandleApi) Update(c echo.Context) error
```

##### `UpdateStatus`

UpdateStatus godoc
@Summary Update merchant document status
@Tags Merchant Document Command
@Security Bearer
@Description Update the status of a merchant document
@Accept json
@Produce json
@Param id path int true "Document ID"
@Param body body requests.UpdateMerchantDocumentStatusRequest true "Update status request"
@Success 200 {object} response.ApiResponseMerchantDocument "Updated document"
@Failure 400 {object} response.ErrorResponse "Bad request or validation error"
@Failure 500 {object} response.ErrorResponse "Failed to update document status"
@Router /api/merchants-documents/update-status/{id} [post]

```go
func (h *merchantCommandDocumentHandleApi) UpdateStatus(c echo.Context) error
```

##### `recordMetrics`

recordMetrics records a Prometheus metric for the given method and status.
It increments a counter and records the duration since the provided start time.

```go
func (s *merchantCommandDocumentHandleApi) recordMetrics(method string, status string, start time.Time)
```

##### `startTracingAndLogging`

startTracingAndLogging starts tracing for a method, logs that the method has started,
and returns a span, a function to end the span, the initial status of the span, and
a function to log a success message.

```go
func (s *merchantCommandDocumentHandleApi) startTracingAndLogging(ctx context.Context, method string, attrs ...attribute.KeyValue) (end func(), logSuccess func(string, ...zap.Field), logError func(string, error, ...zap.Field))
```

### `merchantCommandDocumentHandleParams`

merchantDocumentHandleParams defines the parameters required to initialize
the merchant document handler and register its HTTP routes.

```go
type merchantCommandDocumentHandleParams struct {
	client pb.MerchantDocumentCommandServiceClient
	router *echo.Echo
	logger logger.LoggerInterface
	mapper apimapper.MerchantDocumentCommandResponseMapper
}
```

### `merchantDocumentQueryDocumentHandleParams`

merchantDocumentHandleParams defines the parameters required to initialize
the merchant document handler and register its HTTP routes.

```go
type merchantDocumentQueryDocumentHandleParams struct {
	client pb.MerchantDocumentServiceClient
	router *echo.Echo
	logger logger.LoggerInterface
	mapper apimapper.MerchantDocumentQueryResponseMapper
}
```

### `merchantQueryDocumentHandleApi`

```go
type merchantQueryDocumentHandleApi struct {
	merchantDocument pb.MerchantDocumentServiceClient
	logger logger.LoggerInterface
	mapper apimapper.MerchantDocumentQueryResponseMapper
	trace trace.Tracer
	requestCounter *prometheus.CounterVec
	requestDuration *prometheus.HistogramVec
}
```

#### Methods

##### `FindAll`

FindAll godoc
@Summary Find all merchant documents
@Tags Merchant Document Query
@Security Bearer
@Description Retrieve a list of all merchant documents
@Accept json
@Produce json
@Param page query int false "Page number" default(1)
@Param page_size query int false "Number of items per page" default(10)
@Param search query string false "Search query"
@Success 200 {object} response.ApiResponsePaginationMerchantDocument "List of merchant documents"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve merchant document data"
@Router /api/merchant-document-query [get]

```go
func (h *merchantQueryDocumentHandleApi) FindAll(c echo.Context) error
```

##### `FindAllActive`

FindAllActive godoc
@Summary Find all active merchant documents
@Tags Merchant Document Query
@Security Bearer
@Description Retrieve a list of all active merchant documents
@Accept json
@Produce json
@Param page query int false "Page number" default(1)
@Param page_size query int false "Number of items per page" default(10)
@Param search query string false "Search query"
@Success 200 {object} response.ApiResponsePaginationMerchantDocument "List of active merchant documents"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve active merchant documents"
@Router /api/merchant-document-query/active [get]

```go
func (h *merchantQueryDocumentHandleApi) FindAllActive(c echo.Context) error
```

##### `FindAllTrashed`

FindAllTrashed godoc
@Summary Find all trashed merchant documents
@Tags Merchant Document Query
@Security Bearer
@Description Retrieve a list of all trashed merchant documents
@Accept json
@Produce json
@Param page query int false "Page number" default(1)
@Param page_size query int false "Number of items per page" default(10)
@Param search query string false "Search query"
@Success 200 {object} response.ApiResponsePaginationMerchantDocumentDeleteAt "List of trashed merchant documents"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve trashed merchant documents"
@Router /api/merchant-document-query/trashed [get]

```go
func (h *merchantQueryDocumentHandleApi) FindAllTrashed(c echo.Context) error
```

##### `FindById`

FindById godoc
@Summary Get merchant document by ID
@Tags Merchant Document Query
@Security Bearer
@Description Get a merchant document by its ID
@Accept json
@Produce json
@Param id path int true "Document ID"
@Success 200 {object} response.ApiResponseMerchantDocument "Document details"
@Failure 400 {object} response.ErrorResponse "Invalid ID"
@Failure 500 {object} response.ErrorResponse "Failed to get document"
@Router /api/merchant-document-query/{id} [get]

```go
func (h *merchantQueryDocumentHandleApi) FindById(c echo.Context) error
```

##### `recordMetrics`

```go
func (s *merchantQueryDocumentHandleApi) recordMetrics(method string, status string, start time.Time)
```

##### `startTracingAndLogging`

```go
func (s *merchantQueryDocumentHandleApi) startTracingAndLogging(ctx context.Context, method string, attrs ...attribute.KeyValue) (end func(), logSuccess func(string, ...zap.Field), logError func(string, error, ...zap.Field))
```

## 🚀 Functions

### `NewMerchantCommandDocumentHandler`

```go
func NewMerchantCommandDocumentHandler(params *merchantCommandDocumentHandleParams)
```

### `RegisterMerchantDocumentHandler`

```go
func RegisterMerchantDocumentHandler(deps *DepsMerchantDocument)
```

### `setupMerchantDocumentCommandHandler`

```go
func setupMerchantDocumentCommandHandler(deps *DepsMerchantDocument, mapper apimapper.MerchantDocumentCommandResponseMapper) func()
```

### `setupMerchantDocumentQueryHandler`

```go
func setupMerchantDocumentQueryHandler(deps *DepsMerchantDocument, mapper apimapper.MerchantDocumentQueryResponseMapper) func()
```

