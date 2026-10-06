# 📦 Package `transactionhandler`

**Source Path:** `apigateway/internal/handler/transaction`

## 🧩 Types

### `DepsTransaction`

```go
type DepsTransaction struct {
	Client *grpc.ClientConn
	E *echo.Echo
	Kafka *kafka.Kafka
	Logger logger.LoggerInterface
}
```

### `transactionCommandHandleApi`

```go
type transactionCommandHandleApi struct {
	kafka *kafka.Kafka
	client pb.TransactionCommandServiceClient
	logger logger.LoggerInterface
	mapper apimapper.TransactionCommandResponseMapper
	trace trace.Tracer
	requestCounter *prometheus.CounterVec
	requestDuration *prometheus.HistogramVec
}
```

#### Methods

##### `Create`

@Summary Create a new transaction
@Tags Transaction Command
@Security Bearer
@Description Create a new transaction record with the provided details.
@Accept json
@Produce json
@Param CreateTransactionRequest body requests.CreateTransactionRequest true "Create Transaction Request"
@Success 200 {object} response.ApiResponseTransaction "Successfully created transaction record"
@Failure 400 {object} response.ErrorResponse "Bad Request: Invalid request body or validation error"
@Failure 500 {object} response.ErrorResponse "Failed to create transaction"
@Router /api/transaction-command/create [post]

```go
func (h *transactionCommandHandleApi) Create(c echo.Context) error
```

##### `DeleteAllTransactionPermanent`

@Summary Permanently delete a transaction
@Tags Transaction Command
@Security Bearer
@Description Permanently delete a transaction all.
@Accept json
@Produce json
@Success 200 {object} response.ApiResponseTransactionAll "Successfully deleted transaction record"
@Failure 400 {object} response.ErrorResponse "Bad Request: Invalid ID"
@Failure 500 {object} response.ErrorResponse "Failed to delete transaction:"
@Router /api/transaction-command/delete/all [post]

```go
func (h *transactionCommandHandleApi) DeleteAllTransactionPermanent(c echo.Context) error
```

##### `DeletePermanent`

@Summary Permanently delete a transaction
@Tags Transaction Command
@Security Bearer
@Description Permanently delete a transaction record by its ID.
@Accept json
@Produce json
@Param id path int true "Transaction ID"
@Success 200 {object} response.ApiResponseTransactionDelete "Successfully deleted transaction record"
@Failure 400 {object} response.ErrorResponse "Bad Request: Invalid ID"
@Failure 500 {object} response.ErrorResponse "Failed to delete transaction:"
@Router /api/transaction-command/permanent/{id} [delete]

```go
func (h *transactionCommandHandleApi) DeletePermanent(c echo.Context) error
```

##### `RestoreAllTransaction`

@Summary Restore a trashed transaction
@Tags Transaction Command
@Security Bearer
@Description Restore a trashed transaction all.
@Accept json
@Produce json
@Success 200 {object} response.ApiResponseTransactionAll "Successfully restored transaction record"
@Failure 400 {object} response.ErrorResponse "Bad Request: Invalid ID"
@Failure 500 {object} response.ErrorResponse "Failed to restore transaction:"
@Router /api/transaction-command/restore/all [post]

```go
func (h *transactionCommandHandleApi) RestoreAllTransaction(c echo.Context) error
```

##### `RestoreTransaction`

@Summary Restore a trashed transaction
@Tags Transaction Command
@Security Bearer
@Description Restore a trashed transaction record by its ID.
@Accept json
@Produce json
@Param id path int true "Transaction ID"
@Success 200 {object} response.ApiResponseTransaction "Successfully restored transaction record"
@Failure 400 {object} response.ErrorResponse "Bad Request: Invalid ID"
@Failure 500 {object} response.ErrorResponse "Failed to restore transaction:"
@Router /api/transaction-command/restore/{id} [post]

```go
func (h *transactionCommandHandleApi) RestoreTransaction(c echo.Context) error
```

##### `TrashedTransaction`

@Summary Trash a transaction
@Tags Transaction Command
@Security Bearer
@Description Trash a transaction record by its ID.
@Accept json
@Produce json
@Param id path int true "Transaction ID"
@Success 200 {object} response.ApiResponseTransaction "Successfully trashed transaction record"
@Failure 400 {object} response.ErrorResponse "Bad Request: Invalid ID"
@Failure 500 {object} response.ErrorResponse "Failed to trashed transaction"
@Router /api/transaction-command/trashed/{id} [post]

```go
func (h *transactionCommandHandleApi) TrashedTransaction(c echo.Context) error
```

##### `Update`

@Summary Update a transaction
@Tags Transaction Command
@Security Bearer
@Description Update an existing transaction record using its ID
@Accept json
@Produce json
@Param transaction body requests.UpdateTransactionRequest true "Transaction data"
@Success 200 {object} response.ApiResponseTransaction "Updated transaction data"
@Failure 400 {object} response.ErrorResponse "Bad Request: Invalid request body or validation error"
@Failure 500 {object} response.ErrorResponse "Failed to update transaction"
@Router /api/transaction-command/update [post]

```go
func (h *transactionCommandHandleApi) Update(c echo.Context) error
```

##### `recordMetrics`

```go
func (s *transactionCommandHandleApi) recordMetrics(method string, status string, start time.Time)
```

##### `startTracingAndLogging`

```go
func (s *transactionCommandHandleApi) startTracingAndLogging(ctx context.Context, method string, attrs ...attribute.KeyValue) (end func(), logSuccess func(string, ...zap.Field), logError func(string, error, ...zap.Field))
```

### `transactionCommandHandleParams`

```go
type transactionCommandHandleParams struct {
	kafka *kafka.Kafka
	client pb.TransactionCommandServiceClient
	router *echo.Echo
	logger logger.LoggerInterface
	mapper apimapper.TransactionCommandResponseMapper
}
```

### `transactionQueryHandleApi`

```go
type transactionQueryHandleApi struct {
	client pb.TransactionQueryServiceClient
	logger logger.LoggerInterface
	mapper apimapper.TransactionQueryResponseMapper
	trace trace.Tracer
	requestCounter *prometheus.CounterVec
	requestDuration *prometheus.HistogramVec
}
```

#### Methods

##### `FindAll`

@Summary Find all
@Tags Transaction Query
@Security Bearer
@Description Retrieve a list of all transactions
@Accept json
@Produce json
@Param page query int false "Page number" default(1)
@Param page_size query int false "Number of items per page" default(10)
@Param search query string false "Search query"
@Success 200 {object} response.ApiResponsePaginationTransaction "List of transactions"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve transaction data"
@Router /api/transaction-query [get]

```go
func (h *transactionQueryHandleApi) FindAll(c echo.Context) error
```

##### `FindAllTransactionByCardNumber`

@Summary Find all transactions by card number
@Tags Transaction Query
@Security Bearer
@Description Retrieve a list of transactions for a specific card number
@Accept json
@Produce json
@Param card_number path string true "Card Number"
@Param page query int false "Page number" default(1)
@Param page_size query int false "Number of items per page" default(10)
@Param search query string false "Search query"
@Success 200 {object} response.ApiResponsePaginationTransaction "List of transactions"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve transaction data"
@Router /api/transaction-query/card-number/{card_number} [get]

```go
func (h *transactionQueryHandleApi) FindAllTransactionByCardNumber(c echo.Context) error
```

##### `FindByActiveTransaction`

@Summary Find active transactions
@Tags Transaction Query
@Security Bearer
@Description Retrieve a list of active transactions
@Accept json
@Produce json
@Param page query int false "Page number (default: 1)"
@Param page_size query int false "Number of items per page (default: 10)"
@Param search query string false "Search keyword"
@Success 200 {object} response.ApiResponseTransactions "List of active transactions"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve transaction data"
@Router /api/transaction-query/active [get]

```go
func (h *transactionQueryHandleApi) FindByActiveTransaction(c echo.Context) error
```

##### `FindById`

@Summary Find a transaction by ID
@Tags Transaction Query
@Security Bearer
@Description Retrieve a transaction record using its ID
@Accept json
@Produce json
@Param id path string true "Transaction ID"
@Success 200 {object} response.ApiResponseTransaction "Transaction data"
@Failure 400 {object} response.ErrorResponse "Bad Request: Invalid ID"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve transaction data"
@Router /api/transaction-query/{id} [get]

```go
func (h *transactionQueryHandleApi) FindById(c echo.Context) error
```

##### `FindByTransactionMerchantId`

@Summary Find transactions by merchant ID
@Tags Transaction Query
@Security Bearer
@Description Retrieve a list of transactions using the merchant ID
@Accept json
@Produce json
@Param merchant_id query string true "Merchant ID"
@Success 200 {object} response.ApiResponseTransactions "Transaction data"
@Failure 400 {object} response.ErrorResponse "Bad Request: Invalid ID"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve transaction data"
@Router /api/transaction-query/merchant/{merchant_id} [get]

```go
func (h *transactionQueryHandleApi) FindByTransactionMerchantId(c echo.Context) error
```

##### `FindByTrashedTransaction`

@Summary Retrieve trashed transactions
@Tags Transaction Query
@Security Bearer
@Description Retrieve a list of trashed transactions
@Accept json
@Produce json
@Param page query int false "Page number (default: 1)"
@Param page_size query int false "Number of items per page (default: 10)"
@Param search query string false "Search keyword"
@Success 200 {object} response.ApiResponseTransactions "List of trashed transactions"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve transaction data"
@Router /api/transaction-query/trashed [get]

```go
func (h *transactionQueryHandleApi) FindByTrashedTransaction(c echo.Context) error
```

##### `recordMetrics`

```go
func (s *transactionQueryHandleApi) recordMetrics(method string, status string, start time.Time)
```

##### `startTracingAndLogging`

```go
func (s *transactionQueryHandleApi) startTracingAndLogging(ctx context.Context, method string, attrs ...attribute.KeyValue) (end func(), logSuccess func(string, ...zap.Field), logError func(string, error, ...zap.Field))
```

### `transactionQueryHandleParams`

```go
type transactionQueryHandleParams struct {
	client pb.TransactionQueryServiceClient
	router *echo.Echo
	logger logger.LoggerInterface
	mapper apimapper.TransactionQueryResponseMapper
}
```

### `transactionStatsAmountHandleApi`

```go
type transactionStatsAmountHandleApi struct {
	client pb.TransactionsStatsAmountServiceClient
	logger logger.LoggerInterface
	mapper apimapper.TransactionStatsAmountResponseMapper
	trace trace.Tracer
	requestCounter *prometheus.CounterVec
	requestDuration *prometheus.HistogramVec
}
```

#### Methods

##### `FindMonthlyAmounts`

FindMonthlyAmounts retrieves the monthly transaction amounts for a specific year.
@Summary Get monthly transaction amounts
@Tags Transaction Stats Amount
@Security Bearer
@Description Retrieve the monthly transaction amounts for a specific year.
@Accept json
@Produce json
@Param year query int true "Year"
@Success 200 {object} response.ApiResponseTransactionMonthAmount "Monthly transaction amounts"
@Failure 400 {object} response.ErrorResponse "Invalid year parameter"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve monthly transaction amounts"
@Router /api/transaction-stats-amount/monthly-amounts [get]

```go
func (h *transactionStatsAmountHandleApi) FindMonthlyAmounts(c echo.Context) error
```

##### `FindMonthlyAmountsByCardNumber`

FindMonthlyAmountsByCardNumber retrieves the monthly transaction amounts for a specific card number and year.
@Summary Get monthly transaction amounts by card number
@Tags Transaction Stats Amount
@Security Bearer
@Description Retrieve the monthly transaction amounts for a specific card number and year.
@Accept json
@Produce json
@Param card_number query string true "Card Number"
@Param year query int true "Year"
@Success 200 {object} response.ApiResponseTransactionMonthAmount "Monthly transaction amounts by card number"
@Failure 400 {object} response.ErrorResponse "Invalid card number or year parameter"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve monthly transaction amounts by card number"
@Router /api/transaction-stats-amount/monthly-amounts-by-card [get]

```go
func (h *transactionStatsAmountHandleApi) FindMonthlyAmountsByCardNumber(c echo.Context) error
```

##### `FindYearlyAmounts`

FindYearlyAmounts retrieves the yearly transaction amounts for a specific year.
@Summary Get yearly transaction amounts
@Tags Transaction Stats Amount
@Security Bearer
@Description Retrieve the yearly transaction amounts for a specific year.
@Accept json
@Produce json
@Param year query int true "Year"
@Success 200 {object} response.ApiResponseTransactionYearAmount "Yearly transaction amounts"
@Failure 400 {object} response.ErrorResponse "Invalid year parameter"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve yearly transaction amounts"
@Router /api/transaction-stats-amount/yearly-amounts [get]

```go
func (h *transactionStatsAmountHandleApi) FindYearlyAmounts(c echo.Context) error
```

##### `FindYearlyAmountsByCardNumber`

FindYearlyAmountsByCardNumber retrieves the yearly transaction amounts for a specific card number and year.
@Summary Get yearly transaction amounts by card number
@Tags Transaction Stats Amount
@Security Bearer
@Description Retrieve the yearly transaction amounts for a specific card number and year.
@Accept json
@Produce json
@Param card_number query string true "Card Number"
@Param year query int true "Year"
@Success 200 {object} response.ApiResponseTransactionYearAmount "Yearly transaction amounts by card number"
@Failure 400 {object} response.ErrorResponse "Invalid card number or year parameter"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve yearly transaction amounts by card number"
@Router /api/transaction-stats-amount/yearly-amounts-by-card [get]

```go
func (h *transactionStatsAmountHandleApi) FindYearlyAmountsByCardNumber(c echo.Context) error
```

##### `recordMetrics`

```go
func (s *transactionStatsAmountHandleApi) recordMetrics(method string, status string, start time.Time)
```

##### `startTracingAndLogging`

```go
func (s *transactionStatsAmountHandleApi) startTracingAndLogging(ctx context.Context, method string, attrs ...attribute.KeyValue) (end func(), logSuccess func(string, ...zap.Field), logError func(string, error, ...zap.Field))
```

### `transactionStatsAmountHandleParams`

```go
type transactionStatsAmountHandleParams struct {
	client pb.TransactionsStatsAmountServiceClient
	router *echo.Echo
	logger logger.LoggerInterface
	mapper apimapper.TransactionStatsAmountResponseMapper
}
```

### `transactionStatsMethodHandleApi`

```go
type transactionStatsMethodHandleApi struct {
	client pb.TransactionStatsMethodServiceClient
	logger logger.LoggerInterface
	mapper apimapper.TransactionStatsMethodResponseMapper
	trace trace.Tracer
	requestCounter *prometheus.CounterVec
	requestDuration *prometheus.HistogramVec
}
```

#### Methods

##### `FindMonthlyPaymentMethods`

FindMonthlyPaymentMethods retrieves the monthly payment methods for transactions.
@Summary Get monthly payment methods
@Tags Transaction Stats Method
@Security Bearer
@Description Retrieve the monthly payment methods for transactions by year.
@Accept json
@Produce json
@Param year query int true "Year"
@Success 200 {object} response.ApiResponseTransactionMonthMethod "Monthly payment methods"
@Failure 400 {object} response.ErrorResponse "Invalid year parameter"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve monthly payment methods"
@Router /api/transaction-stats-method/monthly-payment-methods [get]

```go
func (h *transactionStatsMethodHandleApi) FindMonthlyPaymentMethods(c echo.Context) error
```

##### `FindMonthlyPaymentMethodsByCardNumber`

FindMonthlyPaymentMethodsByCardNumber retrieves the monthly payment methods for transactions by card number and year.
@Summary Get monthly payment methods by card number
@Tags Transaction Stats Method
@Security Bearer
@Description Retrieve the monthly payment methods for transactions by card number and year.
@Accept json
@Produce json
@Param card_number query string true "Card Number"
@Param year query int true "Year"
@Success 200 {object} response.ApiResponseTransactionMonthMethod "Monthly payment methods by card number"
@Failure 400 {object} response.ErrorResponse "Invalid card number or year parameter"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve monthly payment methods by card number"
@Router /api/transaction-stats-method/monthly-payment-methods-by-card [get]

```go
func (h *transactionStatsMethodHandleApi) FindMonthlyPaymentMethodsByCardNumber(c echo.Context) error
```

##### `FindYearlyPaymentMethods`

FindYearlyPaymentMethods retrieves the yearly payment methods for transactions.
@Summary Get yearly payment methods
@Tags Transaction Stats Method
@Security Bearer
@Description Retrieve the yearly payment methods for transactions by year.
@Accept json
@Produce json
@Param year query int true "Year"
@Success 200 {object} response.ApiResponseTransactionYearMethod "Yearly payment methods"
@Failure 400 {object} response.ErrorResponse "Invalid year parameter"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve yearly payment methods"
@Router /api/transaction-stats-method/yearly-payment-methods [get]

```go
func (h *transactionStatsMethodHandleApi) FindYearlyPaymentMethods(c echo.Context) error
```

##### `FindYearlyPaymentMethodsByCardNumber`

FindYearlyPaymentMethodsByCardNumber retrieves the yearly payment methods for transactions by card number and year.
@Summary Get yearly payment methods by card number
@Tags Transaction Stats Method
@Security Bearer
@Description Retrieve the yearly payment methods for transactions by card number and year.
@Accept json
@Produce json
@Param card_number query string true "Card Number"
@Param year query int true "Year"
@Success 200 {object} response.ApiResponseTransactionYearMethod "Yearly payment methods by card number"
@Failure 400 {object} response.ErrorResponse "Invalid card number or year parameter"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve yearly payment methods by card number"
@Router /api/transaction-stats-method/yearly-payment-methods-by-card [get]

```go
func (h *transactionStatsMethodHandleApi) FindYearlyPaymentMethodsByCardNumber(c echo.Context) error
```

##### `recordMetrics`

```go
func (s *transactionStatsMethodHandleApi) recordMetrics(method string, status string, start time.Time)
```

##### `startTracingAndLogging`

```go
func (s *transactionStatsMethodHandleApi) startTracingAndLogging(ctx context.Context, method string, attrs ...attribute.KeyValue) (end func(), logSuccess func(string, ...zap.Field), logError func(string, error, ...zap.Field))
```

### `transactionStatsMethodHandleParams`

```go
type transactionStatsMethodHandleParams struct {
	client pb.TransactionStatsMethodServiceClient
	router *echo.Echo
	logger logger.LoggerInterface
	mapper apimapper.TransactionStatsMethodResponseMapper
}
```

### `transactionStatsStatusHandleApi`

```go
type transactionStatsStatusHandleApi struct {
	client pb.TransactionStatsStatusServiceClient
	logger logger.LoggerInterface
	mapper apimapper.TransactionStatsStatusResponseMapper
	trace trace.Tracer
	requestCounter *prometheus.CounterVec
	requestDuration *prometheus.HistogramVec
}
```

#### Methods

##### `FindMonthlyTransactionStatusFailed`

FindMonthlyTransactionStatusFailed retrieves the monthly transaction status for failed transactions.
@Summary Get monthly transaction status for failed transactions
@Tags Transaction Stats Status
@Security Bearer
@Description Retrieve the monthly transaction status for failed transactions by year and month.
@Accept json
@Produce json
@Param year query int true "Year"
@Param month query int true "Month"
@Success 200 {object} response.ApiResponseTransactionMonthStatusFailed "Monthly transaction status for failed transactions"
@Failure 400 {object} response.ErrorResponse "Invalid year or month"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve monthly transaction status for failed transactions"
@Router /api/transaction-stats-status/monthly-failed [get]

```go
func (h *transactionStatsStatusHandleApi) FindMonthlyTransactionStatusFailed(c echo.Context) error
```

##### `FindMonthlyTransactionStatusFailedByCardNumber`

FindMonthlyTransactionStatusFailed retrieves the monthly transaction status for failed transactions.
@Summary Get monthly transaction status for failed transactions
@Tags Transaction Stats Status
@Security Bearer
@Description Retrieve the monthly transaction status for failed transactions by year and month.
@Accept json
@Produce json
@Param year query int true "Year"
@Param month query int true "Month"
@Param cardNumber query string true "Card Number"
@Success 200 {object} response.ApiResponseTransactionMonthStatusFailed "Monthly transaction status for failed transactions"
@Failure 400 {object} response.ErrorResponse "Invalid year or month"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve monthly transaction status for failed transactions"
@Router /api/transaction-stats-status/monthly-failed-by-card [get]

```go
func (h *transactionStatsStatusHandleApi) FindMonthlyTransactionStatusFailedByCardNumber(c echo.Context) error
```

##### `FindMonthlyTransactionStatusSuccess`

FindMonthlyTransactionStatusSuccess retrieves the monthly transaction status for successful transactions.
@Summary Get monthly transaction status for successful transactions
@Tags Transaction Stats Status
@Security Bearer
@Description Retrieve the monthly transaction status for successful transactions by year and month.
@Accept json
@Produce json
@Param year query int true "Year"
@Param month query int true "Month"
@Success 200 {object} response.ApiResponseTransactionMonthStatusSuccess "Monthly transaction status for successful transactions"
@Failure 400 {object} response.ErrorResponse "Invalid year or month"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve monthly transaction status for successful transactions"
@Router /api/transaction-stats-status/monthly-success [get]

```go
func (h *transactionStatsStatusHandleApi) FindMonthlyTransactionStatusSuccess(c echo.Context) error
```

##### `FindMonthlyTransactionStatusSuccessByCardNumber`

FindMonthlyTransactionStatusSuccess retrieves the monthly transaction status for successful transactions.
@Summary Get monthly transaction status for successful transactions
@Tags Transaction Stats Status
@Security Bearer
@Description Retrieve the monthly transaction status for successful transactions by year and month.
@Accept json
@Produce json
@Param year query int true "Year"
@Param month query int true "Month"
@Param card_number query string true "Card Number"
@Success 200 {object} response.ApiResponseTransactionMonthStatusSuccess "Monthly transaction status for successful transactions"
@Failure 400 {object} response.ErrorResponse "Invalid year or month"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve monthly transaction status for successful transactions"
@Router /api/transaction-stats-status/monthly-success-by-card [get]

```go
func (h *transactionStatsStatusHandleApi) FindMonthlyTransactionStatusSuccessByCardNumber(c echo.Context) error
```

##### `FindYearlyTransactionStatusFailed`

FindYearlyTransactionStatusFailed retrieves the yearly transaction status for failed transactions.
@Summary Get yearly transaction status for failed transactions
@Tags Transaction Stats Status
@Security Bearer
@Description Retrieve the yearly transaction status for failed transactions by year.
@Accept json
@Produce json
@Param year query int true "Year"
@Success 200 {object} response.ApiResponseTransactionYearStatusFailed "Yearly transaction status for failed transactions"
@Failure 400 {object} response.ErrorResponse "Invalid year"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve yearly transaction status for failed transactions"
@Router /api/transaction-stats-status/yearly-failed [get]

```go
func (h *transactionStatsStatusHandleApi) FindYearlyTransactionStatusFailed(c echo.Context) error
```

##### `FindYearlyTransactionStatusFailedByCardNumber`

FindYearlyTransactionStatusFailedByCardNumber retrieves the yearly transaction status for failed transactions.
@Summary Get yearly transaction status for failed transactions
@Tags Transaction Stats Status
@Security Bearer
@Description Retrieve the yearly transaction status for failed transactions by year.
@Accept json
@Produce json
@Param card_number query string true "Card Number"
@Param year query int true "Year"
@Success 200 {object} response.ApiResponseTransactionYearStatusFailed "Yearly transaction status for failed transactions"
@Failure 400 {object} response.ErrorResponse "Invalid year"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve yearly transaction status for failed transactions"
@Router /api/transaction-stats-status/yearly-failed-by-card [get]

```go
func (h *transactionStatsStatusHandleApi) FindYearlyTransactionStatusFailedByCardNumber(c echo.Context) error
```

##### `FindYearlyTransactionStatusSuccess`

FindYearlyTransactionStatusSuccess retrieves the yearly transaction status for successful transactions.
@Summary Get yearly transaction status for successful transactions
@Tags Transaction Stats Status
@Security Bearer
@Description Retrieve the yearly transaction status for successful transactions by year.
@Accept json
@Produce json
@Param year query int true "Year"
@Success 200 {object} response.ApiResponseTransactionYearStatusSuccess "Yearly transaction status for successful transactions"
@Failure 400 {object} response.ErrorResponse "Invalid year"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve yearly transaction status for successful transactions"
@Router /api/transaction-stats-status/yearly-success [get]

```go
func (h *transactionStatsStatusHandleApi) FindYearlyTransactionStatusSuccess(c echo.Context) error
```

##### `FindYearlyTransactionStatusSuccessByCardNumber`

FindYearlyTransactionStatusSuccess retrieves the yearly transaction status for successful transactions.
@Summary Get yearly transaction status for successful transactions
@Tags Transaction Stats Status
@Security Bearer
@Description Retrieve the yearly transaction status for successful transactions by year.
@Accept json
@Produce json
@Param year query int true "Year"
@Param cardNumber query string true "Card Number"
@Success 200 {object} response.ApiResponseTransactionYearStatusSuccess "Yearly transaction status for successful transactions"
@Failure 400 {object} response.ErrorResponse "Invalid year"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve yearly transaction status for successful transactions"
@Router /api/transaction-stats-status/yearly-success-by-card [get]

```go
func (h *transactionStatsStatusHandleApi) FindYearlyTransactionStatusSuccessByCardNumber(c echo.Context) error
```

##### `recordMetrics`

```go
func (s *transactionStatsStatusHandleApi) recordMetrics(method string, status string, start time.Time)
```

##### `startTracingAndLogging`

```go
func (s *transactionStatsStatusHandleApi) startTracingAndLogging(ctx context.Context, method string, attrs ...attribute.KeyValue) (end func(), logSuccess func(string, ...zap.Field), logError func(string, error, ...zap.Field))
```

### `transactionStatsStatusHandleParams`

```go
type transactionStatsStatusHandleParams struct {
	client pb.TransactionStatsStatusServiceClient
	router *echo.Echo
	logger logger.LoggerInterface
	mapper apimapper.TransactionStatsStatusResponseMapper
}
```

## 🚀 Functions

### `RegisterTransactionHandler`

```go
func RegisterTransactionHandler(deps *DepsTransaction)
```

### `setupTransactionCommandHandler`

```go
func setupTransactionCommandHandler(deps *DepsTransaction, mapper apimapper.TransactionCommandResponseMapper) func()
```

### `setupTransactionQueryHandler`

```go
func setupTransactionQueryHandler(deps *DepsTransaction, mapper apimapper.TransactionQueryResponseMapper) func()
```

### `setupTransactionStatsAmountHandler`

```go
func setupTransactionStatsAmountHandler(deps *DepsTransaction, mapper apimapper.TransactionStatsAmountResponseMapper) func()
```

### `setupTransactionStatsMethodHandler`

```go
func setupTransactionStatsMethodHandler(deps *DepsTransaction, mapper apimapper.TransactionStatsMethodResponseMapper) func()
```

### `setupTransactionStatsStatusHandler`

```go
func setupTransactionStatsStatusHandler(deps *DepsTransaction, mapper apimapper.TransactionStatsStatusResponseMapper) func()
```

