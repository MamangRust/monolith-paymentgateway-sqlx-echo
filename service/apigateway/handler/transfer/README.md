# 📦 Package `transferhandler`

**Source Path:** `apigateway/internal/handler/transfer`

## 🧩 Types

### `DepsTransfer`

```go
type DepsTransfer struct {
	client *grpc.ClientConn
	E *echo.Echo
	logger logger.LoggerInterface
}
```

### `transferCommandHandleApi`

```go
type transferCommandHandleApi struct {
	client pb.TransferCommandServiceClient
	logger logger.LoggerInterface
	mapper apimapper.TransferCommandResponseMapper
	trace trace.Tracer
	requestCounter *prometheus.CounterVec
	requestDuration *prometheus.HistogramVec
}
```

#### Methods

##### `CreateTransfer`

@Summary Create a transfer
@Tags Transfer Command
@Security Bearer
@Description Create a new transfer record
@Accept json
@Produce json
@Param body body requests.CreateTransferRequest true "Transfer request"
@Success 200 {object} response.ApiResponseTransfer "Transfer data"
@Failure 400 {object} response.ErrorResponse "Validation Error"
@Failure 500 {object} response.ErrorResponse "Failed to create transfer"
@Router /api/transfer-command/create [post]

```go
func (h *transferCommandHandleApi) CreateTransfer(c echo.Context) error
```

##### `DeleteAllTransferPermanent`

@Summary Permanently delete a transfer
@Tags Transfer Command
@Security Bearer
@Description Permanently delete a transfer record all.
@Accept json
@Produce json
@Param id path int true "Transfer ID"
@Success 200 {object} response.ApiResponseTransferAll "Successfully deleted transfer all"
@Failure 400 {object} response.ErrorResponse "Bad Request: Invalid ID"
@Failure 500 {object} response.ErrorResponse "Failed to delete transfer:"
@Router /api/transfer-command/permanent/all [post]

```go
func (h *transferCommandHandleApi) DeleteAllTransferPermanent(c echo.Context) error
```

##### `DeleteTransferPermanent`

@Summary Permanently delete a transfer
@Tags Transfer Command
@Security Bearer
@Description Permanently delete a transfer record by its ID.
@Accept json
@Produce json
@Param id path int true "Transfer ID"
@Success 200 {object} response.ApiResponseTransferDelete "Successfully deleted transfer record permanently"
@Failure 400 {object} response.ErrorResponse "Bad Request: Invalid ID"
@Failure 500 {object} response.ErrorResponse "Failed to delete transfer:"
@Router /api/transfer-command/permanent/{id} [delete]

```go
func (h *transferCommandHandleApi) DeleteTransferPermanent(c echo.Context) error
```

##### `RestoreAllTransfer`

@Summary Restore a trashed transfer
@Tags Transfer Command
@Security Bearer
@Description Restore a trashed transfer all
@Accept json
@Produce json
@Success 200 {object} response.ApiResponseTransferAll "Successfully restored transfer record"
@Failure 400 {object} response.ErrorResponse "Bad Request: Invalid ID"
@Failure 500 {object} response.ErrorResponse "Failed to restore transfer:"
@Router /api/transfer-command/restore/all [post]

```go
func (h *transferCommandHandleApi) RestoreAllTransfer(c echo.Context) error
```

##### `RestoreTransfer`

@Summary Restore a trashed transfer
@Tags Transfer Command
@Security Bearer
@Description Restore a trashed transfer record by its ID.
@Accept json
@Produce json
@Param id path int true "Transfer ID"
@Success 200 {object} response.ApiResponseTransfer "Successfully restored transfer record"
@Failure 400 {object} response.ErrorResponse "Bad Request: Invalid ID"
@Failure 500 {object} response.ErrorResponse "Failed to restore transfer:"
@Router /api/transfer-command/restore/{id} [post]

```go
func (h *transferCommandHandleApi) RestoreTransfer(c echo.Context) error
```

##### `TrashTransfer`

@Summary Soft delete a transfer
@Tags Transfer Command
@Security Bearer
@Description Soft delete a transfer record by its ID.
@Accept json
@Produce json
@Param id path int true "Transfer ID"
@Success 200 {object} response.ApiResponseTransfer "Successfully trashed transfer record"
@Failure 400 {object} response.ErrorResponse "Bad Request: Invalid ID"
@Failure 500 {object} response.ErrorResponse "Failed to trashed transfer"
@Router /api/transfer-command/trash/{id} [post]

```go
func (h *transferCommandHandleApi) TrashTransfer(c echo.Context) error
```

##### `UpdateTransfer`

@Summary Update a transfer
@Tags Transfer Command
@Security Bearer
@Description Update an existing transfer record
@Accept json
@Produce json
@Param id path int true "Transfer ID"
@Param body body requests.UpdateTransferRequest true "Transfer request"
@Success 200 {object} response.ApiResponseTransfer "Transfer data"
@Failure 400 {object} response.ErrorResponse "Validation Error"
@Failure 500 {object} response.ErrorResponse "Failed to update transfer"
@Router /api/transfer-command/update/{id} [post]

```go
func (h *transferCommandHandleApi) UpdateTransfer(c echo.Context) error
```

##### `recordMetrics`

```go
func (s *transferCommandHandleApi) recordMetrics(method string, status string, start time.Time)
```

##### `startTracingAndLogging`

```go
func (s *transferCommandHandleApi) startTracingAndLogging(ctx context.Context, method string, attrs ...attribute.KeyValue) (end func(), logSuccess func(string, ...zap.Field), logError func(string, error, ...zap.Field))
```

### `transferCommandHandleParams`

```go
type transferCommandHandleParams struct {
	client pb.TransferCommandServiceClient
	router *echo.Echo
	logger logger.LoggerInterface
	mapper apimapper.TransferCommandResponseMapper
}
```

### `transferQueryHandleApi`

```go
type transferQueryHandleApi struct {
	client pb.TransferQueryServiceClient
	logger logger.LoggerInterface
	mapper apimapper.TransferQueryResponseMapper
	trace trace.Tracer
	requestCounter *prometheus.CounterVec
	requestDuration *prometheus.HistogramVec
}
```

#### Methods

##### `FindAll`

@Summary Find all transfer records
@Tags Transfer Query
@Security Bearer
@Description Retrieve a list of all transfer records with pagination
@Accept json
@Produce json
@Param page query int false "Page number" default(1)
@Param page_size query int false "Number of items per page" default(10)
@Param search query string false "Search query"
@Success 200 {object} response.ApiResponsePaginationTransfer "List of transfer records"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve transfer data"
@Router /api/transfer-query [get]

```go
func (h *transferQueryHandleApi) FindAll(c echo.Context) error
```

##### `FindByActiveTransfer`

@Summary Find active transfers
@Tags Transfer Query
@Security Bearer
@Description Retrieve a list of active transfer records
@Accept json
@Produce json
@Param page query int false "Page number (default: 1)"
@Param page_size query int false "Number of items per page (default: 10)"
@Param search query string false "Search keyword"
@Success 200 {object} response.ApiResponseTransfers "Active transfer data"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve transfer data"
@Router /api/transfer-query/active [get]

```go
func (h *transferQueryHandleApi) FindByActiveTransfer(c echo.Context) error
```

##### `FindById`

@Summary Find a transfer by ID
@Tags Transfer Query
@Security Bearer
@Description Retrieve a transfer record using its ID
@Accept json
@Produce json
@Param id path string true "Transfer ID"
@Success 200 {object} response.ApiResponseTransfer "Transfer data"
@Failure 400 {object} response.ErrorResponse "Bad Request: Invalid ID"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve transfer data"
@Router /api/transfer-query/{id} [get]

```go
func (h *transferQueryHandleApi) FindById(c echo.Context) error
```

##### `FindByTransferByTransferFrom`

@Summary Find transfers by transfer_from
@Tags Transfer Query
@Security Bearer
@Description Retrieve a list of transfer records using the transfer_from parameter
@Accept json
@Produce json
@Param transfer_from path string true "Transfer From"
@Success 200 {object} response.ApiResponseTransfers "Transfer data"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve transfer data"
@Router /api/transfer-query/transfer_from/{transfer_from} [get]

```go
func (h *transferQueryHandleApi) FindByTransferByTransferFrom(c echo.Context) error
```

##### `FindByTransferByTransferTo`

@Summary Find transfers by transfer_to
@Tags Transfer Query
@Security Bearer
@Description Retrieve a list of transfer records using the transfer_to parameter
@Accept json
@Produce json
@Param transfer_to path string true "Transfer To"
@Success 200 {object} response.ApiResponseTransfers "Transfer data"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve transfer data"
@Router /api/transfer-query/transfer_to/{transfer_to} [get]

```go
func (h *transferQueryHandleApi) FindByTransferByTransferTo(c echo.Context) error
```

##### `FindByTrashedTransfer`

@Summary Retrieve trashed transfers
@Tags Transfer Query
@Security Bearer
@Description Retrieve a list of trashed transfer records
@Accept json
@Produce json
@Param page query int false "Page number (default: 1)"
@Param page_size query int false "Number of items per page (default: 10)"
@Param search query string false "Search keyword"
@Success 200 {object} response.ApiResponseTransfers "List of trashed transfer records"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve transfer data"
@Router /api/transfer-query/trashed [get]

```go
func (h *transferQueryHandleApi) FindByTrashedTransfer(c echo.Context) error
```

##### `recordMetrics`

```go
func (s *transferQueryHandleApi) recordMetrics(method string, status string, start time.Time)
```

##### `startTracingAndLogging`

```go
func (s *transferQueryHandleApi) startTracingAndLogging(ctx context.Context, method string, attrs ...attribute.KeyValue) (end func(), logSuccess func(string, ...zap.Field), logError func(string, error, ...zap.Field))
```

### `transferQueryHandleParams`

```go
type transferQueryHandleParams struct {
	client pb.TransferQueryServiceClient
	router *echo.Echo
	logger logger.LoggerInterface
	mapper apimapper.TransferQueryResponseMapper
}
```

### `transferStatsAmountHandleApi`

```go
type transferStatsAmountHandleApi struct {
	client pb.TransferStatsAmountServiceClient
	logger logger.LoggerInterface
	mapper apimapper.TransferStatsAmountResponseMapper
	trace trace.Tracer
	requestCounter *prometheus.CounterVec
	requestDuration *prometheus.HistogramVec
}
```

#### Methods

##### `FindMonthlyTransferAmounts`

FindMonthlyTransferAmounts retrieves the monthly transfer amounts for a specific year.
@Summary Get monthly transfer amounts
@Tags Transfer Stats Amount
@Security Bearer
@Description Retrieve the monthly transfer amounts for a specific year.
@Accept json
@Produce json
@Param year query int true "Year"
@Success 200 {object} response.ApiResponseTransferMonthAmount "Monthly transfer amounts"
@Failure 400 {object} response.ErrorResponse "Invalid year parameter"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve monthly transfer amounts"
@Router /api/transfer-stats-amount/monthly-amounts [get]

```go
func (h *transferStatsAmountHandleApi) FindMonthlyTransferAmounts(c echo.Context) error
```

##### `FindMonthlyTransferAmountsByReceiverCardNumber`

FindMonthlyTransferAmountsByReceiverCardNumber retrieves the monthly transfer amounts for a specific receiver card number and year.
@Summary Get monthly transfer amounts by receiver card number
@Tags Transfer Stats Amount
@Security Bearer
@Description Retrieve the monthly transfer amounts for a specific receiver card number and year.
@Accept json
@Produce json
@Param card_number query string true "Receiver Card Number"
@Param year query int true "Year"
@Success 200 {object} response.ApiResponseTransferMonthAmount "Monthly transfer amounts by receiver card number"
@Failure 400 {object} response.ErrorResponse "Invalid card number or year parameter"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve monthly transfer amounts by receiver card number"
@Router /api/transfer-stats-amount/monthly-amounts-by-receiver-card [get]

```go
func (h *transferStatsAmountHandleApi) FindMonthlyTransferAmountsByReceiverCardNumber(c echo.Context) error
```

##### `FindMonthlyTransferAmountsBySenderCardNumber`

FindMonthlyTransferAmountsBySenderCardNumber retrieves the monthly transfer amounts for a specific sender card number and year.
@Summary Get monthly transfer amounts by sender card number
@Tags Transfer Stats Amount
@Security Bearer
@Description Retrieve the monthly transfer amounts for a specific sender card number and year.
@Accept json
@Produce json
@Param card_number query string true "Sender Card Number"
@Param year query int true "Year"
@Success 200 {object} response.ApiResponseTransferMonthAmount "Monthly transfer amounts by sender card number"
@Failure 400 {object} response.ErrorResponse "Invalid card number or year parameter"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve monthly transfer amounts by sender card number"
@Router /api/transfer-stats-amount/monthly-amounts-by-sender-card [get]

```go
func (h *transferStatsAmountHandleApi) FindMonthlyTransferAmountsBySenderCardNumber(c echo.Context) error
```

##### `FindYearlyTransferAmounts`

FindYearlyTransferAmounts retrieves the yearly transfer amounts for a specific year.
@Summary Get yearly transfer amounts
@Tags Transfer Stats Amount
@Security Bearer
@Description Retrieve the yearly transfer amounts for a specific year.
@Accept json
@Produce json
@Param year query int true "Year"
@Success 200 {object} response.ApiResponseTransferYearAmount "Yearly transfer amounts"
@Failure 400 {object} response.ErrorResponse "Invalid year parameter"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve yearly transfer amounts"
@Router /api/transfer-stats-amount/yearly-amounts [get]

```go
func (h *transferStatsAmountHandleApi) FindYearlyTransferAmounts(c echo.Context) error
```

##### `FindYearlyTransferAmountsByReceiverCardNumber`

FindYearlyTransferAmountsByReceiverCardNumber retrieves the yearly transfer amounts for a specific receiver card number and year.
@Summary Get yearly transfer amounts by receiver card number
@Tags Transfer Stats Amount
@Security Bearer
@Description Retrieve the yearly transfer amounts for a specific receiver card number and year.
@Accept json
@Produce json
@Param card_number query string true "Receiver Card Number"
@Param year query int true "Year"
@Success 200 {object} response.ApiResponseTransferYearAmount "Yearly transfer amounts by receiver card number"
@Failure 400 {object} response.ErrorResponse "Invalid card number or year parameter"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve yearly transfer amounts by receiver card number"
@Router /api/transfer-stats-amount/yearly-amounts-by-receiver-card [get]

```go
func (h *transferStatsAmountHandleApi) FindYearlyTransferAmountsByReceiverCardNumber(c echo.Context) error
```

##### `FindYearlyTransferAmountsBySenderCardNumber`

FindYearlyTransferAmountsBySenderCardNumber retrieves the yearly transfer amounts for a specific sender card number and year.
@Summary Get yearly transfer amounts by sender card number
@Tags Transfer Stats Amount
@Security Bearer
@Description Retrieve the yearly transfer amounts for a specific sender card number and year.
@Accept json
@Produce json
@Param card_number query string true "Sender Card Number"
@Param year query int true "Year"
@Success 200 {object} response.ApiResponseTransferYearAmount "Yearly transfer amounts by sender card number"
@Failure 400 {object} response.ErrorResponse "Invalid card number or year parameter"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve yearly transfer amounts by sender card number"
@Router /api/transfer-stats-amount/yearly-amounts-by-sender-card [get]

```go
func (h *transferStatsAmountHandleApi) FindYearlyTransferAmountsBySenderCardNumber(c echo.Context) error
```

##### `recordMetrics`

```go
func (s *transferStatsAmountHandleApi) recordMetrics(method string, status string, start time.Time)
```

##### `startTracingAndLogging`

```go
func (s *transferStatsAmountHandleApi) startTracingAndLogging(ctx context.Context, method string, attrs ...attribute.KeyValue) (end func(), logSuccess func(string, ...zap.Field), logError func(string, error, ...zap.Field))
```

### `transferStatsAmountHandleParams`

```go
type transferStatsAmountHandleParams struct {
	client pb.TransferStatsAmountServiceClient
	router *echo.Echo
	logger logger.LoggerInterface
	mapper apimapper.TransferStatsAmountResponseMapper
}
```

### `transferStatsStatusHandleApi`

```go
type transferStatsStatusHandleApi struct {
	client pb.TransferStatsStatusServiceClient
	logger logger.LoggerInterface
	mapper apimapper.TransferStatsStatusResponseMapper
	trace trace.Tracer
	requestCounter *prometheus.CounterVec
	requestDuration *prometheus.HistogramVec
}
```

#### Methods

##### `FindMonthlyTransferStatusFailed`

FindMonthlyTransferStatusFailed retrieves the monthly transfer status for failed transactions.
@Summary Get monthly transfer status for failed transactions
@Tags Transfer Stats Status
@Security Bearer
@Description Retrieve the monthly transfer status for failed transactions by year and month.
@Accept json
@Produce json
@Param year query int true "Year"
@Param month query int true "Month"
@Success 200 {object} response.ApiResponseTransferMonthStatusFailed "Monthly transfer status for failed transactions"
@Failure 400 {object} response.ErrorResponse "Invalid year or month"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve monthly transfer status for failed transactions"
@Router /api/transfer-stats-status/monthly-failed [get]

```go
func (h *transferStatsStatusHandleApi) FindMonthlyTransferStatusFailed(c echo.Context) error
```

##### `FindMonthlyTransferStatusFailedByCardNumber`

FindMonthlyTransferStatusFailedByCardNumber retrieves the monthly transfer status for failed transactions.
@Summary Get monthly transfer status for failed transactions
@Tags Transfer Stats Status
@Security Bearer
@Description Retrieve the monthly transfer status for failed transactions by year and month.
@Accept json
@Produce json
@Param year query int true "Year"
@Param month query int true "Month"
@Param card_number query string true "Card Number"
@Success 200 {object} response.ApiResponseTransferMonthStatusFailed "Monthly transfer status for failed transactions"
@Failure 400 {object} response.ErrorResponse "Invalid year or month"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve monthly transfer status for failed transactions"
@Router /api/transfer-stats-status/monthly-failed-by-card [get]

```go
func (h *transferStatsStatusHandleApi) FindMonthlyTransferStatusFailedByCardNumber(c echo.Context) error
```

##### `FindMonthlyTransferStatusSuccess`

FindMonthlyTransferStatusSuccess retrieves the monthly transfer status for successful transactions.
@Summary Get monthly transfer status for successful transactions
@Tags Transfer Stats Status
@Security Bearer
@Description Retrieve the monthly transfer status for successful transactions by year and month.
@Accept json
@Produce json
@Param year query int true "Year"
@Param month query int true "Month"
@Success 200 {object} response.ApiResponseTransferMonthStatusSuccess "Monthly transfer status for successful transactions"
@Failure 400 {object} response.ErrorResponse "Invalid year or month"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve monthly transfer status for successful transactions"
@Router /api/transfer-stats-status/monthly-success [get]

```go
func (h *transferStatsStatusHandleApi) FindMonthlyTransferStatusSuccess(c echo.Context) error
```

##### `FindMonthlyTransferStatusSuccessByCardNumber`

FindMonthlyTransferStatusSuccessByCardNumber retrieves the monthly transfer status for successful transactions.
@Summary Get monthly transfer status for successful transactions
@Tags Transfer Stats Status
@Security Bearer
@Description Retrieve the monthly transfer status for successful transactions by year and month.
@Accept json
@Produce json
@Param year query int true "Year"
@Param month query int true "Month"
@Success 200 {object} response.ApiResponseTransferMonthStatusSuccess "Monthly transfer status for successful transactions"
@Failure 400 {object} response.ErrorResponse "Invalid year or month"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve monthly transfer status for successful transactions"
@Router /api/transfer-stats-status/monthly-success-by-card [get]

```go
func (h *transferStatsStatusHandleApi) FindMonthlyTransferStatusSuccessByCardNumber(c echo.Context) error
```

##### `FindYearlyTransferStatusFailed`

FindYearlyTransferStatusFailed retrieves the yearly transfer status for failed transactions.
@Summary Get yearly transfer status for failed transactions
@Tags Transfer Stats Status
@Security Bearer
@Description Retrieve the yearly transfer status for failed transactions by year.
@Accept json
@Produce json
@Param year query int true "Year"
@Success 200 {object} response.ApiResponseTransferYearStatusFailed "Yearly transfer status for failed transactions"
@Failure 400 {object} response.ErrorResponse "Invalid year"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve yearly transfer status for failed transactions"
@Router /api/transfer-stats-status/yearly-failed [get]

```go
func (h *transferStatsStatusHandleApi) FindYearlyTransferStatusFailed(c echo.Context) error
```

##### `FindYearlyTransferStatusFailedByCardNumber`

FindYearlyTransferStatusFailedByCardNumber retrieves the yearly transfer status for failed transactions.
@Summary Get yearly transfer status for failed transactions
@Tags Transfer Stats Status
@Security Bearer
@Description Retrieve the yearly transfer status for failed transactions by year.
@Accept json
@Produce json
@Param year query int true "Year"
@Param card_number query string true "Card Number"
@Success 200 {object} response.ApiResponseTransferYearStatusFailed "Yearly transfer status for failed transactions"
@Failure 400 {object} response.ErrorResponse "Invalid year"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve yearly transfer status for failed transactions"
@Router /api/transfer-stats-status/yearly-failed-by-card [get]

```go
func (h *transferStatsStatusHandleApi) FindYearlyTransferStatusFailedByCardNumber(c echo.Context) error
```

##### `FindYearlyTransferStatusSuccess`

FindYearlyTransferStatusSuccess retrieves the yearly transfer status for successful transactions.
@Summary Get yearly transfer status for successful transactions
@Tags Transfer Stats Status
@Security Bearer
@Description Retrieve the yearly transfer status for successful transactions by year.
@Accept json
@Produce json
@Param year query int true "Year"
@Success 200 {object} response.ApiResponseTransferYearStatusSuccess "Yearly transfer status for successful transactions"
@Failure 400 {object} response.ErrorResponse "Invalid year"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve yearly transfer status for successful transactions"
@Router /api/transfer-stats-status/yearly-success [get]

```go
func (h *transferStatsStatusHandleApi) FindYearlyTransferStatusSuccess(c echo.Context) error
```

##### `FindYearlyTransferStatusSuccessByCardNumber`

FindYearlyTransferStatusSuccessByCardNumber retrieves the yearly transfer status for successful transactions.
@Summary Get yearly transfer status for successful transactions
@Tags Transfer Stats Status
@Security Bearer
@Description Retrieve the yearly transfer status for successful transactions by year.
@Accept json
@Produce json
@Param year query int true "Year"
@Param card_number query string true "Card Number"
@Success 200 {object} response.ApiResponseTransferYearStatusSuccess "Yearly transfer status for successful transactions"
@Failure 400 {object} response.ErrorResponse "Invalid year"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve yearly transfer status for successful transactions"
@Router /api/transfer-stats-status/yearly-success-by-card [get]

```go
func (h *transferStatsStatusHandleApi) FindYearlyTransferStatusSuccessByCardNumber(c echo.Context) error
```

##### `recordMetrics`

```go
func (s *transferStatsStatusHandleApi) recordMetrics(method string, status string, start time.Time)
```

##### `startTracingAndLogging`

```go
func (s *transferStatsStatusHandleApi) startTracingAndLogging(ctx context.Context, method string, attrs ...attribute.KeyValue) (end func(), logSuccess func(string, ...zap.Field), logError func(string, error, ...zap.Field))
```

### `transferStatsStatusHandleParams`

```go
type transferStatsStatusHandleParams struct {
	client pb.TransferStatsStatusServiceClient
	router *echo.Echo
	logger logger.LoggerInterface
	mapper apimapper.TransferStatsStatusResponseMapper
}
```

## 🚀 Functions

### `RegisterTransferHandler`

```go
func RegisterTransferHandler(deps *DepsTransfer)
```

### `setupTransferCommandHandler`

```go
func setupTransferCommandHandler(deps *DepsTransfer, mapper apimapper.TransferCommandResponseMapper) func()
```

### `setupTransferQueryHandler`

```go
func setupTransferQueryHandler(deps *DepsTransfer, mapper apimapper.TransferQueryResponseMapper) func()
```

### `setupTransferStatsAmountHandler`

```go
func setupTransferStatsAmountHandler(deps *DepsTransfer, mapper apimapper.TransferStatsAmountResponseMapper) func()
```

### `setupTransferStatsStatusHandler`

```go
func setupTransferStatsStatusHandler(deps *DepsTransfer, mapper apimapper.TransferStatsStatusResponseMapper) func()
```

