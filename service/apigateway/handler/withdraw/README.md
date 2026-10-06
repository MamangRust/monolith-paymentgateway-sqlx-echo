# 📦 Package `withdrawhandler`

**Source Path:** `apigateway/internal/handler/withdraw`

## 🧩 Types

### `DepsWithdraw`

```go
type DepsWithdraw struct {
	Client *grpc.ClientConn
	E *echo.Echo
	Logger logger.LoggerInterface
}
```

### `withdrawCommandHandleApi`

```go
type withdrawCommandHandleApi struct {
	client pb.WithdrawCommandServiceClient
	logger logger.LoggerInterface
	mapper apimapper.WithdrawCommandResponseMapper
	trace trace.Tracer
	requestCounter *prometheus.CounterVec
	requestDuration *prometheus.HistogramVec
}
```

#### Methods

##### `Create`

@Summary Create a new withdraw
@Tags Withdraw Command
@Security Bearer
@Description Create a new withdraw record with the provided details.
@Accept json
@Produce json
@Param CreateWithdrawRequest body requests.CreateWithdrawRequest true "Create Withdraw Request"
@Success 200 {object} response.ApiResponseWithdraw "Successfully created withdraw record"
@Failure 400 {object} response.ErrorResponse "Bad Request: Invalid request body or validation error"
@Failure 500 {object} response.ErrorResponse "Failed to create withdraw"
@Router /api/withdraw-command/create [post]

```go
func (h *withdrawCommandHandleApi) Create(c echo.Context) error
```

##### `DeleteAllWithdrawPermanent`

@Summary Permanently delete a withdraw by ID
@Tags Withdraw Command
@Security Bearer
@Description Permanently delete a withdraw by its ID
@Accept json
@Produce json
@Success 200 {object} response.ApiResponseWithdrawAll "Successfully deleted withdraw permanently"
@Failure 400 {object} response.ErrorResponse "Bad Request: Invalid ID"
@Failure 500 {object} response.ErrorResponse "Failed to delete withdraw permanently:"
@Router /api/withdraw-command/permanent/all [post]

```go
func (h *withdrawCommandHandleApi) DeleteAllWithdrawPermanent(c echo.Context) error
```

##### `DeleteWithdrawPermanent`

@Summary Permanently delete a withdraw by ID
@Tags Withdraw Command
@Security Bearer
@Description Permanently delete a withdraw by its ID
@Accept json
@Produce json
@Param id path int true "Withdraw ID"
@Success 200 {object} response.ApiResponseWithdrawDelete "Successfully deleted withdraw permanently"
@Failure 400 {object} response.ErrorResponse "Bad Request: Invalid ID"
@Failure 500 {object} response.ErrorResponse "Failed to delete withdraw permanently:"
@Router /api/withdraw-command/permanent/{id} [delete]

```go
func (h *withdrawCommandHandleApi) DeleteWithdrawPermanent(c echo.Context) error
```

##### `RestoreAllWithdraw`

@Summary Restore a withdraw all
@Tags Withdraw Command
@Security Bearer
@Description Restore a withdraw all
@Accept json
@Produce json
@Success 200 {object} response.ApiResponseWithdrawAll "Withdraw data"
@Failure 400 {object} response.ErrorResponse "Bad Request: Invalid ID"
@Failure 500 {object} response.ErrorResponse "Failed to restore withdraw"
@Router /api/withdraw-command/restore/all [post]

```go
func (h *withdrawCommandHandleApi) RestoreAllWithdraw(c echo.Context) error
```

##### `RestoreWithdraw`

@Summary Restore a withdraw by ID
@Tags Withdraw Command
@Security Bearer
@Description Restore a withdraw by its ID
@Accept json
@Produce json
@Param id path int true "Withdraw ID"
@Success 200 {object} response.ApiResponseWithdraw "Withdraw data"
@Failure 400 {object} response.ErrorResponse "Bad Request: Invalid ID"
@Failure 500 {object} response.ErrorResponse "Failed to restore withdraw"
@Router /api/withdraw-command/restore/{id} [post]

```go
func (h *withdrawCommandHandleApi) RestoreWithdraw(c echo.Context) error
```

##### `TrashWithdraw`

@Summary Trash a withdraw by ID
@Tags Withdraw Command
@Security Bearer
@Description Trash a withdraw using its ID
@Accept json
@Produce json
@Param id path int true "Withdraw ID"
@Success 200 {object} response.ApiResponseWithdraw "Withdaw data"
@Failure 400 {object} response.ErrorResponse "Bad Request: Invalid ID"
@Failure 500 {object} response.ErrorResponse "Failed to trash withdraw"
@Router /api/withdraw-command/trashed/{id} [post]

```go
func (h *withdrawCommandHandleApi) TrashWithdraw(c echo.Context) error
```

##### `Update`

@Summary Update an existing withdraw
@Tags Withdraw Command
@Security Bearer
@Description Update an existing withdraw record with the provided details.
@Accept json
@Produce json
@Param id path int true "Withdraw ID"
@Param UpdateWithdrawRequest body requests.UpdateWithdrawRequest true "Update Withdraw Request"
@Success 200 {object} response.ApiResponseWithdraw "Successfully updated withdraw record"
@Failure 400 {object} response.ErrorResponse "Bad Request: Invalid request body or validation error"
@Failure 500 {object} response.ErrorResponse "Failed to update withdraw"
@Router /api/withdraw-command/update/{id} [post]

```go
func (h *withdrawCommandHandleApi) Update(c echo.Context) error
```

##### `recordMetrics`

```go
func (s *withdrawCommandHandleApi) recordMetrics(method string, status string, start time.Time)
```

##### `startTracingAndLogging`

```go
func (s *withdrawCommandHandleApi) startTracingAndLogging(ctx context.Context, method string, attrs ...attribute.KeyValue) (end func(), logSuccess func(string, ...zap.Field), logError func(string, error, ...zap.Field))
```

### `withdrawCommandHandleParams`

```go
type withdrawCommandHandleParams struct {
	client pb.WithdrawCommandServiceClient
	router *echo.Echo
	logger logger.LoggerInterface
	mapper apimapper.WithdrawCommandResponseMapper
}
```

### `withdrawQueryHandleApi`

```go
type withdrawQueryHandleApi struct {
	client pb.WithdrawQueryServiceClient
	logger logger.LoggerInterface
	mapper apimapper.WithdrawQueryResponseMapper
	trace trace.Tracer
	requestCounter *prometheus.CounterVec
	requestDuration *prometheus.HistogramVec
}
```

#### Methods

##### `FindAll`

@Summary Find all withdraw records
@Tags Withdraw Query
@Security Bearer
@Description Retrieve a list of all withdraw records with pagination and search
@Accept json
@Produce json
@Param page query int false "Page number" default(1)
@Param page_size query int false "Page size" default(10)
@Param search query string false "Search query"
@Success 200 {object} response.ApiResponsePaginationWithdraw "List of withdraw records"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve withdraw data"
@Router /api/withdraw-query [get]

```go
func (h *withdrawQueryHandleApi) FindAll(c echo.Context) error
```

##### `FindAllByCardNumber`

@Summary Find all withdraw records by card number
@Tags Withdraw Query
@Security Bearer
@Description Retrieve a list of withdraw records for a specific card number with pagination and search
@Accept json
@Produce json
@Param card_number path string true "Card Number"
@Param page query int false "Page number" default(1)
@Param page_size query int false "Page size" default(10)
@Param search query string false "Search query"
@Success 200 {object} response.ApiResponsePaginationWithdraw "List of withdraw records"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve withdraw data"
@Router /api/withdraw-query/card-number/{card_number} [get]

```go
func (h *withdrawQueryHandleApi) FindAllByCardNumber(c echo.Context) error
```

##### `FindByActive`

@Summary Retrieve all active withdraw data
@Tags Withdraw Query
@Security Bearer
@Description Retrieve a list of all active withdraw data
@Accept json
@Produce json
@Success 200 {object} response.ApiResponsesWithdraw "List of withdraw data"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve withdraw data"
@Router /api/withdraw-query-query/active [get]

```go
func (h *withdrawQueryHandleApi) FindByActive(c echo.Context) error
```

##### `FindByCardNumber`

@Summary Find a withdraw by card number
@Tags Withdraw
@Security Bearer
@Description Retrieve a withdraw record using its card number
@Accept json
@Produce json
@Param card_number query string true "Card number"
@Success 200 {object} response.ApiResponsesWithdraw "Withdraw data"
@Failure 400 {object} response.ErrorResponse "Bad Request: Invalid card number"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve withdraw data"
@Router /api/withdraw-query/card/{card_number} [get]

```go
func (h *withdrawQueryHandleApi) FindByCardNumber(c echo.Context) error
```

##### `FindById`

@Summary Find a withdraw by ID
@Tags Withdraw Query
@Security Bearer
@Description Retrieve a withdraw record using its ID
@Accept json
@Produce json
@Param id path int true "Withdraw ID"
@Success 200 {object} response.ApiResponseWithdraw "Withdraw data"
@Failure 400 {object} response.ErrorResponse "Bad Request: Invalid ID"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve withdraw data"
@Router /api/withdraw-query/{id} [get]

```go
func (h *withdrawQueryHandleApi) FindById(c echo.Context) error
```

##### `FindByTrashed`

@Summary Retrieve trashed withdraw data
@Tags Withdraw Query
@Security Bearer
@Description Retrieve a list of trashed withdraw data
@Accept json
@Produce json
@Success 200 {object} response.ApiResponsesWithdraw "List of trashed withdraw data"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve withdraw data"
@Router /api/withdraw-query-query/trashed [get]

```go
func (h *withdrawQueryHandleApi) FindByTrashed(c echo.Context) error
```

##### `recordMetrics`

```go
func (s *withdrawQueryHandleApi) recordMetrics(method string, status string, start time.Time)
```

##### `startTracingAndLogging`

```go
func (s *withdrawQueryHandleApi) startTracingAndLogging(ctx context.Context, method string, attrs ...attribute.KeyValue) (end func(), logSuccess func(string, ...zap.Field), logError func(string, error, ...zap.Field))
```

### `withdrawQueryHandleParams`

```go
type withdrawQueryHandleParams struct {
	client pb.WithdrawQueryServiceClient
	router *echo.Echo
	logger logger.LoggerInterface
	mapper apimapper.WithdrawQueryResponseMapper
}
```

### `withdrawStatsAmountHandleApi`

```go
type withdrawStatsAmountHandleApi struct {
	client pb.WithdrawStatsAmountServiceClient
	logger logger.LoggerInterface
	mapper apimapper.WithdrawStatsAmountResponseMapper
	trace trace.Tracer
	requestCounter *prometheus.CounterVec
	requestDuration *prometheus.HistogramVec
}
```

#### Methods

##### `FindMonthlyWithdraws`

FindMonthlyWithdraws retrieves the monthly withdraws for a specific year.
@Summary Get monthly withdraws
@Tags Withdraw Stats Amount
@Security Bearer
@Description Retrieve the monthly withdraws for a specific year.
@Accept json
@Produce json
@Param year query int true "Year"
@Success 200 {object} response.ApiResponseWithdrawMonthAmount "Monthly withdraws"
@Failure 400 {object} response.ErrorResponse "Invalid year parameter"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve monthly withdraws"
@Router /api/withdraw-stats-amount/monthly [get]

```go
func (h *withdrawStatsAmountHandleApi) FindMonthlyWithdraws(c echo.Context) error
```

##### `FindMonthlyWithdrawsByCardNumber`

FindMonthlyWithdrawsByCardNumber retrieves the monthly withdraws for a specific card number and year.
@Summary Get monthly withdraws by card number
@Tags Withdraw Stats Amount
@Security Bearer
@Description Retrieve the monthly withdraws for a specific card number and year.
@Accept json
@Produce json
@Param card_number query string true "Card Number"
@Param year query int true "Year"
@Success 200 {object} response.ApiResponseWithdrawMonthAmount "Monthly withdraws by card number"
@Failure 400 {object} response.ErrorResponse "Invalid card number or year parameter"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve monthly withdraws by card number"
@Router /api/withdraw-stats-amount/monthly-by-card [get]

```go
func (h *withdrawStatsAmountHandleApi) FindMonthlyWithdrawsByCardNumber(c echo.Context) error
```

##### `FindYearlyWithdraws`

FindYearlyWithdraws retrieves the yearly withdraws for a specific year.
@Summary Get yearly withdraws
@Tags Withdraw Stats Amount
@Security Bearer
@Description Retrieve the yearly withdraws for a specific year.
@Accept json
@Produce json
@Param year query int true "Year"
@Success 200 {object} response.ApiResponseWithdrawYearAmount "Yearly withdraws"
@Failure 400 {object} response.ErrorResponse "Invalid year parameter"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve yearly withdraws"
@Router /api/withdraw-stats-amount/yearly [get]

```go
func (h *withdrawStatsAmountHandleApi) FindYearlyWithdraws(c echo.Context) error
```

##### `FindYearlyWithdrawsByCardNumber`

FindYearlyWithdrawsByCardNumber retrieves the yearly withdraws for a specific card number and year.
@Summary Get yearly withdraws by card number
@Tags Withdraw Stats Amount
@Security Bearer
@Description Retrieve the yearly withdraws for a specific card number and year.
@Accept json
@Produce json
@Param card_number query string true "Card Number"
@Param year query int true "Year"
@Success 200 {object} response.ApiResponseWithdrawYearAmount "Yearly withdraws by card number"
@Failure 400 {object} response.ErrorResponse "Invalid card number or year parameter"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve yearly withdraws by card number"
@Router /api/withdraw-stats-amount/yearly-by-card [get]

```go
func (h *withdrawStatsAmountHandleApi) FindYearlyWithdrawsByCardNumber(c echo.Context) error
```

##### `recordMetrics`

```go
func (s *withdrawStatsAmountHandleApi) recordMetrics(method string, status string, start time.Time)
```

##### `startTracingAndLogging`

```go
func (s *withdrawStatsAmountHandleApi) startTracingAndLogging(ctx context.Context, method string, attrs ...attribute.KeyValue) (end func(), logSuccess func(string, ...zap.Field), logError func(string, error, ...zap.Field))
```

### `withdrawStatsAmountHandleParams`

```go
type withdrawStatsAmountHandleParams struct {
	client pb.WithdrawStatsAmountServiceClient
	router *echo.Echo
	logger logger.LoggerInterface
	mapper apimapper.WithdrawStatsAmountResponseMapper
}
```

### `withdrawStatsStatusHandleApi`

```go
type withdrawStatsStatusHandleApi struct {
	client pb.WithdrawStatsStatusClient
	logger logger.LoggerInterface
	mapper apimapper.WithdrawStatsStatusResponseMapper
	trace trace.Tracer
	requestCounter *prometheus.CounterVec
	requestDuration *prometheus.HistogramVec
}
```

#### Methods

##### `FindMonthlyWithdrawStatusFailed`

FindMonthlyWithdrawStatusFailed retrieves the monthly withdraw status for failed transactions.
@Summary Get monthly withdraw status for failed transactions
@Tags Withdraw Stats Withdraw
@Security Bearer
@Description Retrieve the monthly withdraw status for failed transactions by year and month.
@Accept json
@Produce json
@Param year query int true "Year"
@Param month query int true "Month"
@Success 200 {object} response.ApiResponseWithdrawMonthStatusFailed "Monthly withdraw status for failed transactions"
@Failure 400 {object} response.ErrorResponse "Invalid year or month"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve monthly withdraw status for failed transactions"
@Router /api/withdraw-stats-status/monthly-failed [get]

```go
func (h *withdrawStatsStatusHandleApi) FindMonthlyWithdrawStatusFailed(c echo.Context) error
```

##### `FindMonthlyWithdrawStatusFailedByCardNumber`

FindMonthlyWithdrawStatusFailedByCardNumber retrieves the monthly withdraw status for failed transactions.
@Summary Get monthly withdraw status for failed transactions
@Tags Withdraw Stats Withdraw
@Security Bearer
@Description Retrieve the monthly withdraw status for failed transactions by year and month.
@Accept json
@Produce json
@Param year query int true "Year"
@Param month query int true "Month"
@Param card_number query string true "Card Number"
@Success 200 {object} response.ApiResponseWithdrawMonthStatusFailed "Monthly withdraw status for failed transactions"
@Failure 400 {object} response.ErrorResponse "Invalid year or month"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve monthly withdraw status for failed transactions"
@Router /api/withdraw-stats-status/monthly-failed-by-card [get]

```go
func (h *withdrawStatsStatusHandleApi) FindMonthlyWithdrawStatusFailedByCardNumber(c echo.Context) error
```

##### `FindMonthlyWithdrawStatusSuccess`

FindMonthlyWithdrawStatusSuccess retrieves the monthly withdraw status for successful transactions.
@Summary Get monthly withdraw status for successful transactions
@Tags Withdraw Stats Withdraw
@Security Bearer
@Description Retrieve the monthly withdraw status for successful transactions by year and month.
@Accept json
@Produce json
@Param year query int true "Year"
@Param month query int true "Month"
@Success 200 {object} response.ApiResponseWithdrawMonthStatusSuccess "Monthly withdraw status for successful transactions"
@Failure 400 {object} response.ErrorResponse "Invalid year or month"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve monthly withdraw status for successful transactions"
@Router /api/withdraw-stats-status/monthly-success [get]

```go
func (h *withdrawStatsStatusHandleApi) FindMonthlyWithdrawStatusSuccess(c echo.Context) error
```

##### `FindMonthlyWithdrawStatusSuccessByCardNumber`

FindMonthlyWithdrawStatusSuccessByCardNumber retrieves the monthly withdraw status for successful transactions.
@Summary Get monthly withdraw status for successful transactions
@Tags Withdraw Stats Withdraw
@Security Bearer
@Description Retrieve the monthly withdraw status for successful transactions by year and month.
@Accept json
@Produce json
@Param year query int true "Year"
@Param month query int true "Month"
@Param card_number query string true "Card Number"
@Success 200 {object} response.ApiResponseWithdrawMonthStatusSuccess "Monthly withdraw status for successful transactions"
@Failure 400 {object} response.ErrorResponse "Invalid year or month"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve monthly withdraw status for successful transactions"
@Router /api/withdraw-stats-status/monthly-success-by-card [get]

```go
func (h *withdrawStatsStatusHandleApi) FindMonthlyWithdrawStatusSuccessByCardNumber(c echo.Context) error
```

##### `FindYearlyWithdrawStatusFailed`

FindYearlyWithdrawStatusFailed retrieves the yearly withdraw status for failed transactions.
@Summary Get yearly withdraw status for failed transactions
@Tags Withdraw Stats Withdraw
@Security Bearer
@Description Retrieve the yearly withdraw status for failed transactions by year.
@Accept json
@Produce json
@Param year query int true "Year"
@Success 200 {object} response.ApiResponseWithdrawYearStatusSuccess "Yearly withdraw status for failed transactions"
@Failure 400 {object} response.ErrorResponse "Invalid year"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve yearly withdraw status for failed transactions"
@Router /api/withdraw-stats-status/yearly-failed [get]

```go
func (h *withdrawStatsStatusHandleApi) FindYearlyWithdrawStatusFailed(c echo.Context) error
```

##### `FindYearlyWithdrawStatusFailedByCardNumber`

FindYearlyWithdrawStatusFailedByCardNumber retrieves the yearly withdraw status for failed transactions.
@Summary Get yearly withdraw status for failed transactions
@Tags Withdraw Stats Withdraw
@Security Bearer
@Description Retrieve the yearly withdraw status for failed transactions by year.
@Accept json
@Produce json
@Param year query int true "Year"
@Param card_number query string true "Card Number"
@Success 200 {object} response.ApiResponseWithdrawYearStatusSuccess "Yearly withdraw status for failed transactions"
@Failure 400 {object} response.ErrorResponse "Invalid year"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve yearly withdraw status for failed transactions"
@Router /api/withdraw-stats-status/yearly-failed-by-card [get]

```go
func (h *withdrawStatsStatusHandleApi) FindYearlyWithdrawStatusFailedByCardNumber(c echo.Context) error
```

##### `FindYearlyWithdrawStatusSuccess`

FindYearlyWithdrawStatusSuccess retrieves the yearly withdraw status for successful transactions.
@Summary Get yearly withdraw status for successful transactions
@Tags Withdraw Stats Withdraw
@Security Bearer
@Description Retrieve the yearly withdraw status for successful transactions by year.
@Accept json
@Produce json
@Param year query int true "Year"
@Success 200 {object} response.ApiResponseWithdrawYearStatusSuccess "Yearly withdraw status for successful transactions"
@Failure 400 {object} response.ErrorResponse "Invalid year"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve yearly withdraw status for successful transactions"
@Router /api/withdraw-stats-status/yearly-success [get]

```go
func (h *withdrawStatsStatusHandleApi) FindYearlyWithdrawStatusSuccess(c echo.Context) error
```

##### `FindYearlyWithdrawStatusSuccessByCardNumber`

FindYearlyWithdrawStatusSuccessByCardNumber retrieves the yearly withdraw status for successful transactions.
@Summary Get yearly withdraw status for successful transactions
@Tags Withdraw Stats Withdraw
@Security Bearer
@Description Retrieve the yearly withdraw status for successful transactions by year.
@Accept json
@Produce json
@Param year query int true "Year"
@Param card_number query string true "Card Number"
@Success 200 {object} response.ApiResponseWithdrawYearStatusSuccess "Yearly withdraw status for successful transactions"
@Failure 400 {object} response.ErrorResponse "Invalid year"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve yearly withdraw status for successful transactions"
@Router /api/withdraw-stats-status/yearly-success-by-card-number [get]

```go
func (h *withdrawStatsStatusHandleApi) FindYearlyWithdrawStatusSuccessByCardNumber(c echo.Context) error
```

##### `recordMetrics`

```go
func (s *withdrawStatsStatusHandleApi) recordMetrics(method string, status string, start time.Time)
```

##### `startTracingAndLogging`

```go
func (s *withdrawStatsStatusHandleApi) startTracingAndLogging(ctx context.Context, method string, attrs ...attribute.KeyValue) (end func(), logSuccess func(string, ...zap.Field), logError func(string, error, ...zap.Field))
```

### `withdrawStatsStatusHandleParams`

```go
type withdrawStatsStatusHandleParams struct {
	client pb.WithdrawStatsStatusClient
	router *echo.Echo
	logger logger.LoggerInterface
	mapper apimapper.WithdrawStatsStatusResponseMapper
}
```

## 🚀 Functions

### `RegisterWithdrawHandler`

```go
func RegisterWithdrawHandler(deps *DepsWithdraw)
```

### `setupWithdrawCommandHandler`

```go
func setupWithdrawCommandHandler(deps *DepsWithdraw, mapper apimapper.WithdrawCommandResponseMapper) func()
```

### `setupWithdrawQueryHandler`

```go
func setupWithdrawQueryHandler(deps *DepsWithdraw, mapper apimapper.WithdrawQueryResponseMapper) func()
```

### `setupWithdrawStatsAmountHandler`

```go
func setupWithdrawStatsAmountHandler(deps *DepsWithdraw, mapper apimapper.WithdrawStatsAmountResponseMapper) func()
```

### `setupWithdrawStatsStatusHandler`

```go
func setupWithdrawStatsStatusHandler(deps *DepsWithdraw, mapper apimapper.WithdrawStatsStatusResponseMapper) func()
```

