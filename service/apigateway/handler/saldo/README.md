# 📦 Package `saldohandler`

**Source Path:** `apigateway/internal/handler/saldo`

## 🧩 Types

### `DepsSaldo`

```go
type DepsSaldo struct {
	Client *grpc.ClientConn
	E *echo.Echo
	Logger logger.LoggerInterface
}
```

### `saldoCommandHandleApi`

```go
type saldoCommandHandleApi struct {
	saldo pb.SaldoCommandServiceClient
	logger logger.LoggerInterface
	mapper apimapper.SaldoCommandResponseMapper
	trace trace.Tracer
	requestCounter *prometheus.CounterVec
	requestDuration *prometheus.HistogramVec
}
```

#### Methods

##### `Create`

@Summary Create a new saldo
@Tags Saldo-Command
@Security Bearer
@Description Create a new saldo record with the provided card number and total balance.
@Accept json
@Produce json
@Param CreateSaldoRequest body requests.CreateSaldoRequest true "Create Saldo Request"
@Success 200 {object} response.ApiResponseSaldo "Successfully created saldo record"
@Failure 400 {object} response.ErrorResponse "Bad Request: Invalid request body or validation error"
@Failure 500 {object} response.ErrorResponse "Failed to create saldo"
@Router /api/saldo-command/create [post]

```go
func (h *saldoCommandHandleApi) Create(c echo.Context) error
```

##### `Delete`

@Summary Permanently delete a saldo
@Tags Saldo-Command
@Security Bearer
@Description Permanently delete an existing saldo record by its ID.
@Accept json
@Produce json
@Param id path int true "Saldo ID"
@Success 200 {object} response.ApiResponseSaldoDelete "Successfully deleted saldo record"
@Failure 400 {object} response.ErrorResponse "Bad Request: Invalid ID"
@Failure 500 {object} response.ErrorResponse "Failed to delete saldo"
@Router /api/saldo-command/permanent/{id} [delete]

```go
func (h *saldoCommandHandleApi) Delete(c echo.Context) error
```

##### `DeleteAllSaldoPermanent`

@Summary Permanently delete all saldo records
@Tags Saldo-Command
@Security Bearer
@Description Permanently delete all saldo records from the database.
@Accept json
@Produce json
@Success 200 {object} response.ApiResponseSaldoAll "Successfully deleted all saldo records permanently"
@Failure 500 {object} response.ErrorResponse "Failed to permanently delete all saldo records"
@Router /api/saldo-command/permanent/all [post]

```go
func (h *saldoCommandHandleApi) DeleteAllSaldoPermanent(c echo.Context) error
```

##### `RestoreAllSaldo`

RestoreAllSaldo restores all saldo records.
@Summary Restore all saldo records
@Tags Saldo-Command
@Security Bearer
@Description Restore all saldo records that were previously deleted.
@Accept json
@Produce json
@Success 200 {object} response.ApiResponseSaldoAll "Successfully restored all saldo records"
@Failure 500 {object} response.ErrorResponse "Failed to restore all saldo records"
@Router /api/saldo-command/restore/all [post]

```go
func (h *saldoCommandHandleApi) RestoreAllSaldo(c echo.Context) error
```

##### `RestoreSaldo`

@Summary Restore a trashed saldo
@Tags Saldo-Command
@Security Bearer
@Description Restore an existing saldo record from the trash by its ID.
@Accept json
@Produce json
@Param id path int true "Saldo ID"
@Success 200 {object} response.ApiResponseSaldo "Successfully restored saldo record"
@Failure 400 {object} response.ErrorResponse "Bad Request: Invalid ID"
@Failure 500 {object} response.ErrorResponse "Failed to restore saldo"
@Router /api/saldo-command/restore/{id} [post]

```go
func (h *saldoCommandHandleApi) RestoreSaldo(c echo.Context) error
```

##### `TrashSaldo`

@Summary Soft delete a saldo
@Tags Saldo-Command
@Security Bearer
@Description Soft delete an existing saldo record by its ID.
@Accept json
@Produce json
@Param id path int true "Saldo ID"
@Success 200 {object} response.ApiResponseSaldo "Successfully trashed saldo record"
@Failure 400 {object} response.ErrorResponse "Bad Request: Invalid ID"
@Failure 500 {object} response.ErrorResponse "Failed to trashed saldo"
@Router /api/saldo-command/trashed/{id} [post]

```go
func (h *saldoCommandHandleApi) TrashSaldo(c echo.Context) error
```

##### `Update`

@Summary Update an existing saldo
@Tags Saldo-Command
@Security Bearer
@Description Update an existing saldo record with the provided card number and total balance.
@Accept json
@Produce json
@Param id path int true "Saldo ID"
@Param UpdateSaldoRequest body requests.UpdateSaldoRequest true "Update Saldo Request"
@Success 200 {object} response.ApiResponseSaldo "Successfully updated saldo record"
@Failure 400 {object} response.ErrorResponse "Bad Request: Invalid request body or validation error"
@Failure 500 {object} response.ErrorResponse "Failed to update saldo"
@Router /api/saldo-command/update/{id} [post]

```go
func (h *saldoCommandHandleApi) Update(c echo.Context) error
```

##### `recordMetrics`

```go
func (s *saldoCommandHandleApi) recordMetrics(method string, status string, start time.Time)
```

##### `startTracingAndLogging`

```go
func (s *saldoCommandHandleApi) startTracingAndLogging(ctx context.Context, method string, attrs ...attribute.KeyValue) (end func(), logSuccess func(string, ...zap.Field), logError func(string, error, ...zap.Field))
```

### `saldoCommandHandleParams`

```go
type saldoCommandHandleParams struct {
	client pb.SaldoCommandServiceClient
	router *echo.Echo
	logger logger.LoggerInterface
	mapper apimapper.SaldoCommandResponseMapper
}
```

### `saldoQueryHandleApi`

```go
type saldoQueryHandleApi struct {
	saldo pb.SaldoQueryServiceClient
	logger logger.LoggerInterface
	mapper apimapper.SaldoQueryResponseMapper
	trace trace.Tracer
	requestCounter *prometheus.CounterVec
	requestDuration *prometheus.HistogramVec
}
```

#### Methods

##### `FindAll`

@Summary Find all saldo data
@Tags Saldo-Query
@Security Bearer
@Description Retrieve a list of all saldo data with pagination and search
@Accept json
@Produce json
@Param page query int false "Page number" default(1)
@Param page_size query int false "Page size" default(10)
@Param search query string false "Search query"
@Success 200 {object} response.ApiResponsePaginationSaldo "List of saldo data"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve saldo data"
@Router /api/saldo-query [get]

```go
func (h *saldoQueryHandleApi) FindAll(c echo.Context) error
```

##### `FindByActive`

@Summary Retrieve all active saldo data
@Tags Saldo-Query
@Security Bearer
@Description Retrieve a list of all active saldo data
@Accept json
@Produce json
@Param page query int false "Page number" default(1)
@Param page_size query int false "Page size" default(10)
@Param search query string false "Search query"
@Success 200 {object} response.ApiResponsesSaldo "List of saldo data"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve saldo data"
@Router /api/saldo-query/active [get]

```go
func (h *saldoQueryHandleApi) FindByActive(c echo.Context) error
```

##### `FindByCardNumber`

@Summary Find a saldo by card number
@Tags Saldo-Query
@Security Bearer
@Description Retrieve a saldo by its card number
@Accept json
@Produce json
@Param card_number path string true "Card number"
@Success 200 {object} response.ApiResponseSaldo "Saldo data"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve saldo data"
@Router /api/saldo-query/card_number/{card_number} [get]

```go
func (h *saldoQueryHandleApi) FindByCardNumber(c echo.Context) error
```

##### `FindById`

@Summary Find a saldo by ID
@Tags Saldo-Query
@Security Bearer
@Description Retrieve a saldo by its ID
@Accept json
@Produce json
@Param id path int true "Saldo ID"
@Success 200 {object} response.ApiResponseSaldo "Saldo data"
@Failure 400 {object} response.ErrorResponse "Invalid saldo ID"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve saldo data"
@Router /api/saldo-query/{id} [get]

```go
func (h *saldoQueryHandleApi) FindById(c echo.Context) error
```

##### `FindByTrashed`

@Summary Retrieve trashed saldo data
@Tags Saldo-Query
@Security Bearer
@Description Retrieve a list of all trashed saldo data
@Accept json
@Produce json
@Param page query int false "Page number" default(1)
@Param page_size query int false "Page size" default(10)
@Param search query string false "Search query"
@Success 200 {object} response.ApiResponsesSaldo "List of trashed saldo data"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve saldo data"
@Router /api/saldo-query/trashed [get]

```go
func (h *saldoQueryHandleApi) FindByTrashed(c echo.Context) error
```

##### `recordMetrics`

```go
func (s *saldoQueryHandleApi) recordMetrics(method string, status string, start time.Time)
```

##### `startTracingAndLogging`

```go
func (s *saldoQueryHandleApi) startTracingAndLogging(ctx context.Context, method string, attrs ...attribute.KeyValue) (end func(), logSuccess func(string, ...zap.Field), logError func(string, error, ...zap.Field))
```

### `saldoQueryHandleParams`

```go
type saldoQueryHandleParams struct {
	client pb.SaldoQueryServiceClient
	router *echo.Echo
	logger logger.LoggerInterface
	mapper apimapper.SaldoQueryResponseMapper
}
```

### `saldoStatsBalanceHandleApi`

```go
type saldoStatsBalanceHandleApi struct {
	saldo pb.SaldoStatsBalanceServiceClient
	logger logger.LoggerInterface
	mapper apimapper.SaldoStatsBalanceResponseMapper
	trace trace.Tracer
	requestCounter *prometheus.CounterVec
	requestDuration *prometheus.HistogramVec
}
```

#### Methods

##### `FindMonthlySaldoBalances`

FindMonthlySaldoBalances retrieves monthly saldo balances for a specific year.
@Summary Get monthly saldo balances
@Tags Saldo-Stats-Balance
@Security Bearer
@Description Retrieve monthly saldo balances for a specific year.
@Accept json
@Produce json
@Param year query int true "Year"
@Success 200 {object} response.ApiResponseMonthSaldoBalances "Monthly saldo balances"
@Failure 400 {object} response.ErrorResponse "Invalid year parameter"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve monthly saldo balances"
@Router /api/saldo-stats-balance/monthly-balances [get]

```go
func (h *saldoStatsBalanceHandleApi) FindMonthlySaldoBalances(c echo.Context) error
```

##### `FindYearlySaldoBalances`

FindYearlySaldoBalances retrieves yearly saldo balances for a specific year.
@Summary Get yearly saldo balances
@Tags Saldo-Stats-Balance
@Security Bearer
@Description Retrieve yearly saldo balances for a specific year.
@Accept json
@Produce json
@Param year query int true "Year"
@Success 200 {object} response.ApiResponseYearSaldoBalances "Yearly saldo balances"
@Failure 400 {object} response.ErrorResponse "Invalid year parameter"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve yearly saldo balances"
@Router /api/saldo-stats-balance/yearly-balances [get]

```go
func (h *saldoStatsBalanceHandleApi) FindYearlySaldoBalances(c echo.Context) error
```

##### `recordMetrics`

```go
func (s *saldoStatsBalanceHandleApi) recordMetrics(method string, status string, start time.Time)
```

##### `startTracingAndLogging`

```go
func (s *saldoStatsBalanceHandleApi) startTracingAndLogging(ctx context.Context, method string, attrs ...attribute.KeyValue) (end func(), logSuccess func(string, ...zap.Field), logError func(string, error, ...zap.Field))
```

### `saldoStatsBalanceHandleParams`

```go
type saldoStatsBalanceHandleParams struct {
	client pb.SaldoStatsBalanceServiceClient
	router *echo.Echo
	logger logger.LoggerInterface
	mapper apimapper.SaldoStatsBalanceResponseMapper
}
```

### `saldoTotalBalanceHandleApi`

```go
type saldoTotalBalanceHandleApi struct {
	saldo pb.SaldoStatsTotalBalanceClient
	logger logger.LoggerInterface
	mapper apimapper.SaldoStatsTotalResponseMapper
	trace trace.Tracer
	requestCounter *prometheus.CounterVec
	requestDuration *prometheus.HistogramVec
}
```

#### Methods

##### `FindMonthlyTotalSaldoBalance`

FindMonthlyTotalSaldoBalance retrieves the total saldo balance for a specific month and year.
@Summary Get monthly total saldo balance
@Tags Saldo-Stats-Total-Balance
@Security Bearer
@Description Retrieve the total saldo balance for a specific month and year.
@Accept json
@Produce json
@Param year query int true "Year"
@Param month query int true "Month"
@Success 200 {object} response.ApiResponseMonthTotalSaldo "Monthly total saldo balance"
@Failure 400 {object} response.ErrorResponse "Invalid year or month parameter"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve monthly total saldo balance"
@Router /api/saldo-stats-total-balance/monthly-total-balance [get]

```go
func (h *saldoTotalBalanceHandleApi) FindMonthlyTotalSaldoBalance(c echo.Context) error
```

##### `FindYearTotalSaldoBalance`

FindYearTotalSaldoBalance retrieves the total saldo balance for a specific year.
@Summary Get yearly total saldo balance
@Tags Saldo-Stats-Total-Balance
@Security Bearer
@Description Retrieve the total saldo balance for a specific year.
@Accept json
@Produce json
@Param year query int true "Year"
@Success 200 {object} response.ApiResponseYearTotalSaldo "Yearly total saldo balance"
@Failure 400 {object} response.ErrorResponse "Invalid year parameter"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve yearly total saldo balance"
@Router /api/saldo-stats-total-balance/yearly-total-balance [get]

```go
func (h *saldoTotalBalanceHandleApi) FindYearTotalSaldoBalance(c echo.Context) error
```

##### `recordMetrics`

```go
func (s *saldoTotalBalanceHandleApi) recordMetrics(method string, status string, start time.Time)
```

##### `startTracingAndLogging`

```go
func (s *saldoTotalBalanceHandleApi) startTracingAndLogging(ctx context.Context, method string, attrs ...attribute.KeyValue) (end func(), logSuccess func(string, ...zap.Field), logError func(string, error, ...zap.Field))
```

### `saldoTotalBalanceHandleParams`

```go
type saldoTotalBalanceHandleParams struct {
	client pb.SaldoStatsTotalBalanceClient
	router *echo.Echo
	logger logger.LoggerInterface
	mapper apimapper.SaldoStatsTotalResponseMapper
}
```

## 🚀 Functions

### `RegisterSaldoHandler`

```go
func RegisterSaldoHandler(deps *DepsSaldo)
```

### `setupSaldoCommandHandler`

```go
func setupSaldoCommandHandler(deps *DepsSaldo, mapper apimapper.SaldoCommandResponseMapper) func()
```

### `setupSaldoQueryHandler`

```go
func setupSaldoQueryHandler(deps *DepsSaldo, mapper apimapper.SaldoQueryResponseMapper) func()
```

### `setupSaldoStatsBalanceHandler`

```go
func setupSaldoStatsBalanceHandler(deps *DepsSaldo, mapper apimapper.SaldoStatsBalanceResponseMapper) func()
```

### `setupStatsSaldoTotalBalanceHandler`

```go
func setupStatsSaldoTotalBalanceHandler(deps *DepsSaldo, mapper apimapper.SaldoStatsTotalResponseMapper) func()
```

