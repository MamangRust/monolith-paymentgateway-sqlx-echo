# 📦 Package `topuphandler`

**Source Path:** `apigateway/internal/handler/topup`

## 🧩 Types

### `DepsTopup`

```go
type DepsTopup struct {
	Client *grpc.ClientConn
	E *echo.Echo
	Logger logger.LoggerInterface
}
```

### `topupCommandHandleApi`

```go
type topupCommandHandleApi struct {
	client pb.TopupCommandServiceClient
	logger logger.LoggerInterface
	mapper apimapper.TopupCommandResponseMapper
	trace trace.Tracer
	requestCounter *prometheus.CounterVec
	requestDuration *prometheus.HistogramVec
}
```

#### Methods

##### `Create`

@Summary Create topup
@Tags Topup Command
@Security Bearer
@Description Create a new topup record
@Accept json
@Produce json
@Param CreateTopupRequest body requests.CreateTopupRequest true "Create topup request"
@Success 200 {object} response.ApiResponseTopup "Created topup data"
@Failure 400 {object} response.ErrorResponse "Bad Request: "
@Failure 500 {object} response.ErrorResponse "Failed to create topup: "
@Router /api/topup-command/create [post]

```go
func (h *topupCommandHandleApi) Create(c echo.Context) error
```

##### `DeleteAllTopupPermanent`

@Summary Permanently delete all topup records
@Tags Topup Command
@Security Bearer
@Description Permanently delete all topup records from the database.
@Accept json
@Produce json
@Success 200 {object} response.ApiResponseTopupAll "Successfully deleted all topup records permanently"
@Failure 500 {object} response.ErrorResponse "Failed to permanently delete all topup records"
@Router /api/topup-command/permanent/all [post]

```go
func (h *topupCommandHandleApi) DeleteAllTopupPermanent(c echo.Context) error
```

##### `DeleteTopupPermanent`

@Summary Permanently delete a topup
@Tags Topup Command
@Security Bearer
@Description Permanently delete a topup record by its ID.
@Accept json
@Produce json
@Param id path int true "Topup ID"
@Success 200 {object} response.ApiResponseTopupDelete "Successfully deleted topup record permanently"
@Failure 400 {object} response.ErrorResponse "Bad Request: Invalid ID"
@Failure 500 {object} response.ErrorResponse "Failed to delete topup:"
@Router /api/topup-command/permanent/{id} [delete]

```go
func (h *topupCommandHandleApi) DeleteTopupPermanent(c echo.Context) error
```

##### `RestoreAllTopup`

@Summary Restore all topup records
@Tags Topup Command
@Security Bearer
@Description Restore all topup records that were previously deleted.
@Accept json
@Produce json
@Success 200 {object} response.ApiResponseTopupAll "Successfully restored all topup records"
@Failure 500 {object} response.ErrorResponse "Failed to restore all topup records"
@Router /api/topup-command/restore/all [post]

```go
func (h *topupCommandHandleApi) RestoreAllTopup(c echo.Context) error
```

##### `RestoreTopup`

@Summary Restore a trashed topup
@Tags Topup Command
@Security Bearer
@Description Restore a trashed topup record by its ID.
@Accept json
@Produce json
@Param id path int true "Topup ID"
@Success 200 {object} response.ApiResponseTopup "Successfully restored topup record"
@Failure 400 {object} response.ErrorResponse "Bad Request: Invalid ID"
@Failure 500 {object} response.ErrorResponse "Failed to restore topup:"
@Router /api/topup-command/restore/{id} [post]

```go
func (h *topupCommandHandleApi) RestoreTopup(c echo.Context) error
```

##### `TrashTopup`

@Summary Trash a topup
@Tags Topup Command
@Security Bearer
@Description Trash a topup record by its ID.
@Accept json
@Produce json
@Param id path int true "Topup ID"
@Success 200 {object} response.ApiResponseTopup "Successfully trashed topup record"
@Failure 400 {object} response.ErrorResponse "Bad Request: Invalid ID"
@Failure 500 {object} response.ErrorResponse "Failed to trash topup:"
@Router /api/topup-command/trash/{id} [post]

```go
func (h *topupCommandHandleApi) TrashTopup(c echo.Context) error
```

##### `Update`

@Summary Update topup
@Tags Topup Command
@Security Bearer
@Description Update an existing topup record with the provided details
@Accept json
@Produce json
@Param id path int true "Topup ID"
@Param UpdateTopupRequest body requests.UpdateTopupRequest true "Update topup request"
@Success 200 {object} response.ApiResponseTopup "Updated topup data"
@Failure 400 {object} response.ErrorResponse "Bad Request: Invalid input data"
@Failure 500 {object} response.ErrorResponse "Failed to update topup: "
@Router /api/topup-command/update/{id} [post]

```go
func (h *topupCommandHandleApi) Update(c echo.Context) error
```

##### `recordMetrics`

```go
func (s *topupCommandHandleApi) recordMetrics(method string, status string, start time.Time)
```

##### `startTracingAndLogging`

```go
func (s *topupCommandHandleApi) startTracingAndLogging(ctx context.Context, method string, attrs ...attribute.KeyValue) (end func(), logSuccess func(string, ...zap.Field), logError func(string, error, ...zap.Field))
```

### `topupCommandHandleParams`

```go
type topupCommandHandleParams struct {
	client pb.TopupCommandServiceClient
	router *echo.Echo
	logger logger.LoggerInterface
	mapper apimapper.TopupCommandResponseMapper
}
```

### `topupQueryHandleApi`

```go
type topupQueryHandleApi struct {
	client pb.TopupQueryServiceClient
	logger logger.LoggerInterface
	mapper apimapper.TopupQueryResponseMapper
	trace trace.Tracer
	requestCounter *prometheus.CounterVec
	requestDuration *prometheus.HistogramVec
}
```

#### Methods

##### `FindAll`

@Summary Retrieve a list of all topup data
@Tags Topup Query
@Security Bearer
@Description Retrieve a list of all topup data with pagination and search
@Accept json
@Produce json
@Param page query int false "Page number" default(1)
@Param page_size query int false "Page size" default(10)
@Param search query string false "Search query"
@Success 200 {object} response.ApiResponsePaginationTopup "List of topup data"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve topup data"
@Router /api/topup-query [get]

```go
func (h topupQueryHandleApi) FindAll(c echo.Context) error
```

##### `FindAllByCardNumber`

@Summary Find all topup by card number
@Tags Transaction
@Security Bearer
@Description Retrieve a list of transactions for a specific card number
@Accept json
@Produce json
@Param card_number path string true "Card Number"
@Param page query int false "Page number" default(1)
@Param page_size query int false "Number of items per page" default(10)
@Param search query string false "Search query"
@Success 200 {object} response.ApiResponsePaginationTopup "List of topups"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve topups data"
@Router /api/topup-query/card-number/{card_number} [get]

```go
func (h *topupQueryHandleApi) FindAllByCardNumber(c echo.Context) error
```

##### `FindByActive`

@Summary Find active topups
@Tags Topup Query
@Security Bearer
@Description Retrieve a list of active topup records
@Accept json
@Produce json
@Param page query int false "Page number" default(1)
@Param page_size query int false "Page size" default(10)
@Param search query string false "Search query"
@Success 200 {object} response.ApiResponsesTopup "Active topup data"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve topup data"
@Router /api/topup-query/active [get]

```go
func (h *topupQueryHandleApi) FindByActive(c echo.Context) error
```

##### `FindById`

@Summary Find a topup by ID
@Tags Topup Query
@Security Bearer
@Description Retrieve a topup record using its ID
@Accept json
@Produce json
@Param id path string true "Topup ID"
@Success 200 {object} response.ApiResponseTopup "Topup data"
@Failure 400 {object} response.ErrorResponse "Bad Request: Invalid ID"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve topup data"
@Router /api/topup-query/{id} [get]

```go
func (h topupQueryHandleApi) FindById(c echo.Context) error
```

##### `FindByTrashed`

@Summary Retrieve trashed topups
@Tags Topup Query
@Security Bearer
@Description Retrieve a list of trashed topup records
@Accept json
@Produce json
@Param page query int false "Page number" default(1)
@Param page_size query int false "Page size" default(10)
@Param search query string false "Search query"
@Success 200 {object} response.ApiResponsesTopup "List of trashed topup data"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve topup data"
@Router /api/topup-query/trashed [get]

```go
func (h *topupQueryHandleApi) FindByTrashed(c echo.Context) error
```

##### `recordMetrics`

recordMetrics records a Prometheus metric for the given method and status.
It increments a counter and records the duration since the provided start time.

```go
func (s *topupQueryHandleApi) recordMetrics(method string, status string, start time.Time)
```

##### `startTracingAndLogging`

startTracingAndLogging starts a tracing span and returns functions to log the outcome of the call.
The returned functions are logSuccess and logError, which log the outcome of the call to the trace span.
The returned end function records the metrics and ends the trace span.

```go
func (s *topupQueryHandleApi) startTracingAndLogging(ctx context.Context, method string, attrs ...attribute.KeyValue) (end func(), logSuccess func(string, ...zap.Field), logError func(string, error, ...zap.Field))
```

### `topupQueryHandleParams`

```go
type topupQueryHandleParams struct {
	client pb.TopupQueryServiceClient
	router *echo.Echo
	logger logger.LoggerInterface
	mapper apimapper.TopupQueryResponseMapper
}
```

### `topupStatsAmountHandleApi`

```go
type topupStatsAmountHandleApi struct {
	client pb.TopupStatsAmountServiceClient
	logger logger.LoggerInterface
	mapper apimapper.TopupStatsAmountResponseMapper
	trace trace.Tracer
	requestCounter *prometheus.CounterVec
	requestDuration *prometheus.HistogramVec
}
```

#### Methods

##### `FindMonthlyTopupAmounts`

FindMonthlyTopupAmounts retrieves the monthly top-up amounts for a specific year.
@Summary Get monthly top-up amounts
@Tags Topup Amount
@Security Bearer
@Description Retrieve the monthly top-up amounts for a specific year.
@Accept json
@Produce json
@Param year query int true "Year"
@Success 200 {object} response.ApiResponseTopupMonthAmount "Monthly top-up amounts"
@Failure 400 {object} response.ErrorResponse "Invalid year parameter"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve monthly top-up amounts"
@Router /api/topup/monthly-amounts [get]

```go
func (h *topupStatsAmountHandleApi) FindMonthlyTopupAmounts(c echo.Context) error
```

##### `FindMonthlyTopupAmountsByCardNumber`

FindMonthlyTopupAmountsByCardNumber retrieves the monthly top-up amounts for a specific card number and year.
@Summary Get monthly top-up amounts by card number
@Tags Topup Amount
@Security Bearer
@Description Retrieve the monthly top-up amounts for a specific card number and year.
@Accept json
@Produce json
@Param card_number query string true "Card Number"
@Param year query int true "Year"
@Success 200 {object} response.ApiResponseTopupMonthAmount "Monthly top-up amounts by card number"
@Failure 400 {object} response.ErrorResponse "Invalid card number or year parameter"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve monthly top-up amounts by card number"
@Router /api/topup-stats-amount/monthly-amounts-by-card [get]

```go
func (h *topupStatsAmountHandleApi) FindMonthlyTopupAmountsByCardNumber(c echo.Context) error
```

##### `FindYearlyTopupAmounts`

FindYearlyTopupAmounts retrieves the yearly top-up amounts for a specific year.
@Summary Get yearly top-up amounts
@Tags Topup Amount
@Security Bearer
@Description Retrieve the yearly top-up amounts for a specific year.
@Accept json
@Produce json
@Param year query int true "Year"
@Success 200 {object} response.ApiResponseTopupYearAmount "Yearly top-up amounts"
@Failure 400 {object} response.ErrorResponse "Invalid year parameter"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve yearly top-up amounts"
@Router /api/topup-stats-amount/yearly-amounts [get]

```go
func (h *topupStatsAmountHandleApi) FindYearlyTopupAmounts(c echo.Context) error
```

##### `FindYearlyTopupAmountsByCardNumber`

FindYearlyTopupAmountsByCardNumber retrieves the yearly top-up amounts for a specific card number and year.
@Summary Get yearly top-up amounts by card number
@Tags Topup Amount
@Security Bearer
@Description Retrieve the yearly top-up amounts for a specific card number and year.
@Accept json
@Produce json
@Param card_number query string true "Card Number"
@Param year query int true "Year"
@Success 200 {object} response.ApiResponseTopupYearAmount "Yearly top-up amounts by card number"
@Failure 400 {object} response.ErrorResponse "Invalid card number or year parameter"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve yearly top-up amounts by card number"
@Router /api/topup-stats-amount/yearly-amounts-by-card [get]

```go
func (h *topupStatsAmountHandleApi) FindYearlyTopupAmountsByCardNumber(c echo.Context) error
```

##### `recordMetrics`

```go
func (s *topupStatsAmountHandleApi) recordMetrics(method string, status string, start time.Time)
```

##### `startTracingAndLogging`

```go
func (s *topupStatsAmountHandleApi) startTracingAndLogging(ctx context.Context, method string, attrs ...attribute.KeyValue) (end func(), logSuccess func(string, ...zap.Field), logError func(string, error, ...zap.Field))
```

### `topupStatsAmountHandleParams`

```go
type topupStatsAmountHandleParams struct {
	client pb.TopupStatsAmountServiceClient
	router *echo.Echo
	logger logger.LoggerInterface
	mapper apimapper.TopupStatsAmountResponseMapper
}
```

### `topupStatsMethodHandleApi`

```go
type topupStatsMethodHandleApi struct {
	client pb.TopupStatsMethodServiceClient
	logger logger.LoggerInterface
	mapper apimapper.TopupStatsMethodResponseMapper
	trace trace.Tracer
	requestCounter *prometheus.CounterVec
	requestDuration *prometheus.HistogramVec
}
```

#### Methods

##### `FindMonthlyTopupMethods`

FindMonthlyTopupMethods retrieves the monthly top-up methods for a specific year.
@Summary Get monthly top-up methods
@Tags Topup Method
@Security Bearer
@Description Retrieve the monthly top-up methods for a specific year.
@Accept json
@Produce json
@Param year query int true "Year"
@Success 200 {object} response.ApiResponseTopupMonthMethod "Monthly top-up methods"
@Failure 400 {object} response.ErrorResponse "Invalid year parameter"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve monthly top-up methods"
@Router /api/topup-stats-method/monthly-methods [get]

```go
func (h *topupStatsMethodHandleApi) FindMonthlyTopupMethods(c echo.Context) error
```

##### `FindMonthlyTopupMethodsByCardNumber`

FindMonthlyTopupMethodsByCardNumber retrieves the monthly top-up methods for a specific card number and year.
@Summary Get monthly top-up methods by card number
@Tags Topup Method
@Security Bearer
@Description Retrieve the monthly top-up methods for a specific card number and year.
@Accept json
@Produce json
@Param card_number query string true "Card Number"
@Param year query int true "Year"
@Success 200 {object} response.ApiResponseTopupMonthMethod "Monthly top-up methods by card number"
@Failure 400 {object} response.ErrorResponse "Invalid card number or year parameter"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve monthly top-up methods by card number"
@Router /api/topup-stats-method/monthly-methods-by-card [get]

```go
func (h *topupStatsMethodHandleApi) FindMonthlyTopupMethodsByCardNumber(c echo.Context) error
```

##### `FindYearlyTopupMethods`

FindYearlyTopupMethods retrieves the yearly top-up methods for a specific year.
@Summary Get yearly top-up methods
@Tags Topup Method
@Security Bearer
@Description Retrieve the yearly top-up methods for a specific year.
@Accept json
@Produce json
@Param year query int true "Year"
@Success 200 {object} response.ApiResponseTopupYearMethod "Yearly top-up methods"
@Failure 400 {object} response.ErrorResponse "Invalid year parameter"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve yearly top-up methods"
@Router /api/topup-stats-method/yearly-methods [get]

```go
func (h *topupStatsMethodHandleApi) FindYearlyTopupMethods(c echo.Context) error
```

##### `FindYearlyTopupMethodsByCardNumber`

FindYearlyTopupMethodsByCardNumber retrieves the yearly top-up methods for a specific card number and year.
@Summary Get yearly top-up methods by card number
@Tags Topup Method
@Security Bearer
@Description Retrieve the yearly top-up methods for a specific card number and year.
@Accept json
@Produce json
@Param card_number query string true "Card Number"
@Param year query int true "Year"
@Success 200 {object} response.ApiResponseTopupYearMethod "Yearly top-up methods by card number"
@Failure 400 {object} response.ErrorResponse "Invalid card number or year parameter"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve yearly top-up methods by card number"
@Router /api/topup-stats-method/yearly-methods-by-card [get]

```go
func (h *topupStatsMethodHandleApi) FindYearlyTopupMethodsByCardNumber(c echo.Context) error
```

##### `recordMetrics`

```go
func (s *topupStatsMethodHandleApi) recordMetrics(method string, status string, start time.Time)
```

##### `startTracingAndLogging`

```go
func (s *topupStatsMethodHandleApi) startTracingAndLogging(ctx context.Context, method string, attrs ...attribute.KeyValue) (end func(), logSuccess func(string, ...zap.Field), logError func(string, error, ...zap.Field))
```

### `topupStatsMethodHandleParams`

```go
type topupStatsMethodHandleParams struct {
	client pb.TopupStatsMethodServiceClient
	router *echo.Echo
	logger logger.LoggerInterface
	mapper apimapper.TopupStatsMethodResponseMapper
}
```

### `topupStatsStatusHandleApi`

```go
type topupStatsStatusHandleApi struct {
	client pb.TopupStatsStatusServiceClient
	logger logger.LoggerInterface
	mapper apimapper.TopupStatsStatusResponseMapper
	trace trace.Tracer
	requestCounter *prometheus.CounterVec
	requestDuration *prometheus.HistogramVec
}
```

#### Methods

##### `FindMonthlyTopupStatusFailed`

FindMonthlyTopupStatusFailed retrieves the monthly top-up status for failed transactions.
@Summary Get monthly top-up status for failed transactions
@Tags Topup Stats Status
@Security Bearer
@Description Retrieve the monthly top-up status for failed transactions by year and month.
@Accept json
@Produce json
@Param year query int true "Year"
@Param month query int true "Month"
@Success 200 {object} response.ApiResponseTopupMonthStatusFailed "Monthly top-up status for failed transactions"
@Failure 400 {object} response.ErrorResponse "Invalid year or month"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve monthly top-up status for failed transactions"
@Router /api/topup-stats-status/monthly-failed [get]

```go
func (h *topupStatsStatusHandleApi) FindMonthlyTopupStatusFailed(c echo.Context) error
```

##### `FindMonthlyTopupStatusFailedByCardNumber`

FindMonthlyTopupStatusFailed retrieves the monthly top-up status for failed transactions.
@Summary Get monthly top-up status for failed transactions
@Tags Topup Stats Status
@Security Bearer
@Description Retrieve the monthly top-up status for failed transactions by year and month.
@Accept json
@Produce json
@Param year query int true "Year"
@Param month query int true "Month"
@Param card_number query string true "Card Number"
@Success 200 {object} response.ApiResponseTopupMonthStatusFailed "Monthly top-up status for failed transactions"
@Failure 400 {object} response.ErrorResponse "Invalid year or month"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve monthly top-up status for failed transactions"
@Router /api/topup-stats-status/monthly-failed-by-card [get]

```go
func (h *topupStatsStatusHandleApi) FindMonthlyTopupStatusFailedByCardNumber(c echo.Context) error
```

##### `FindMonthlyTopupStatusSuccess`

FindMonthlyTopupStatusSuccess retrieves the monthly top-up status for successful transactions.
@Summary Get monthly top-up status for successful transactions
@Tags Topup Stats Status
@Security Bearer
@Description Retrieve the monthly top-up status for successful transactions by year and month.
@Accept json
@Produce json
@Param year query int true "Year"
@Param month query int true "Month"
@Success 200 {object} response.ApiResponseTopupMonthStatusSuccess "Monthly top-up status for successful transactions"
@Failure 400 {object} response.ErrorResponse "Invalid year or month"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve monthly top-up status for successful transactions"
@Router /api/topup-stats-status/monthly-success [get]

```go
func (h *topupStatsStatusHandleApi) FindMonthlyTopupStatusSuccess(c echo.Context) error
```

##### `FindMonthlyTopupStatusSuccessByCardNumber`

FindMonthlyTopupStatusSuccess retrieves the monthly top-up status for successful transactions.
@Summary Get monthly top-up status for successful transactions
@Tags Topup Stats Status
@Security Bearer
@Description Retrieve the monthly top-up status for successful transactions by year and month.
@Accept json
@Produce json
@Param year query int true "Year"
@Param month query int true "Month"
@Param card_number query string true "Card Number"
@Success 200 {object} response.ApiResponseTopupMonthStatusSuccess "Monthly top-up status for successful transactions"
@Failure 400 {object} response.ErrorResponse "Invalid year or month"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve monthly top-up status for successful transactions"
@Router /api/topup-stats-status/monthly-success-by-card [get]

```go
func (h *topupStatsStatusHandleApi) FindMonthlyTopupStatusSuccessByCardNumber(c echo.Context) error
```

##### `FindYearlyTopupStatusFailed`

FindYearlyTopupStatusFailed retrieves the yearly top-up status for failed transactions.
@Summary Get yearly top-up status for failed transactions
@Tags Topup Stats Status
@Security Bearer
@Description Retrieve the yearly top-up status for failed transactions by year.
@Accept json
@Produce json
@Param year query int true "Year"
@Success 200 {object} response.ApiResponseTopupYearStatusFailed "Yearly top-up status for failed transactions"
@Failure 400 {object} response.ErrorResponse "Invalid year"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve yearly top-up status for failed transactions"
@Router /api/topup-stats-status/yearly-failed [get]

```go
func (h *topupStatsStatusHandleApi) FindYearlyTopupStatusFailed(c echo.Context) error
```

##### `FindYearlyTopupStatusFailedByCardNumber`

FindYearlyTopupStatusFailedByCardNumber retrieves the yearly top-up status for failed transactions.
@Summary Get yearly top-up status for failed transactions
@Tags Topup Stats Status
@Security Bearer
@Description Retrieve the yearly top-up status for failed transactions by year.
@Accept json
@Produce json
@Param year query int true "Year"
@Param card_number query string true "Card Number"
@Success 200 {object} response.ApiResponseTopupYearStatusFailed "Yearly top-up status for failed transactions"
@Failure 400 {object} response.ErrorResponse "Invalid year"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve yearly top-up status for failed transactions"
@Router /api/topup-stats-status/yearly-failed-by-card [get]

```go
func (h *topupStatsStatusHandleApi) FindYearlyTopupStatusFailedByCardNumber(c echo.Context) error
```

##### `FindYearlyTopupStatusSuccess`

FindYearlyTopupStatusSuccess retrieves the yearly top-up status for successful transactions.
@Summary Get yearly top-up status for successful transactions
@Tags Topup Stats Status
@Security Bearer
@Description Retrieve the yearly top-up status for successful transactions by year.
@Accept json
@Produce json
@Param year query int true "Year"
@Success 200 {object} response.ApiResponseTopupYearStatusSuccess "Yearly top-up status for successful transactions"
@Failure 400 {object} response.ErrorResponse "Invalid year"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve yearly top-up status for successful transactions"
@Router /api/topup-stats-status/yearly-success [get]

```go
func (h *topupStatsStatusHandleApi) FindYearlyTopupStatusSuccess(c echo.Context) error
```

##### `FindYearlyTopupStatusSuccessByCardNumber`

FindYearlyTopupStatusSuccess retrieves the yearly top-up status for successful transactions.
@Summary Get yearly top-up status for successful transactions
@Tags Topup Stats Status
@Security Bearer
@Description Retrieve the yearly top-up status for successful transactions by year.
@Accept json
@Produce json
@Param year query int true "Year"
@Param card_number query string true "Card Number"
@Success 200 {object} response.ApiResponseTopupYearStatusSuccess "Yearly top-up status for successful transactions"
@Failure 400 {object} response.ErrorResponse "Invalid year"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve yearly top-up status for successful transactions"
@Router /api/topup-stats-status/yearly-success-by-card [get]

```go
func (h *topupStatsStatusHandleApi) FindYearlyTopupStatusSuccessByCardNumber(c echo.Context) error
```

##### `recordMetrics`

```go
func (s *topupStatsStatusHandleApi) recordMetrics(method string, status string, start time.Time)
```

##### `startTracingAndLogging`

```go
func (s *topupStatsStatusHandleApi) startTracingAndLogging(ctx context.Context, method string, attrs ...attribute.KeyValue) (end func(), logSuccess func(string, ...zap.Field), logError func(string, error, ...zap.Field))
```

### `topupStatsStatusHandleParams`

```go
type topupStatsStatusHandleParams struct {
	client pb.TopupStatsStatusServiceClient
	router *echo.Echo
	logger logger.LoggerInterface
	mapper apimapper.TopupStatsStatusResponseMapper
}
```

## 🚀 Functions

### `RegisterTopupHandler`

```go
func RegisterTopupHandler(deps *DepsTopup)
```

### `setupTopupCommandHandler`

```go
func setupTopupCommandHandler(deps *DepsTopup, mapper apimapper.TopupCommandResponseMapper) func()
```

### `setupTopupQueryHandler`

```go
func setupTopupQueryHandler(deps *DepsTopup, mapper apimapper.TopupQueryResponseMapper) func()
```

### `setupTopupStatsAmountHandler`

```go
func setupTopupStatsAmountHandler(deps *DepsTopup, mapper apimapper.TopupStatsAmountResponseMapper) func()
```

### `setupTopupStatsMethodHandler`

```go
func setupTopupStatsMethodHandler(deps *DepsTopup, mapper apimapper.TopupStatsMethodResponseMapper) func()
```

### `setupTopupStatsStatusHandler`

```go
func setupTopupStatsStatusHandler(deps *DepsTopup, mapper apimapper.TopupStatsStatusResponseMapper) func()
```

