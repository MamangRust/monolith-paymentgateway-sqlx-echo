# 📦 Package `cardhandler`

**Source Path:** `apigateway/internal/handler/card`

## 🧩 Types

### `DepsCard`

```go
type DepsCard struct {
	Client *grpc.ClientConn
	E *echo.Echo
	Logger logger.LoggerInterface
}
```

### `cardCommandHandleApi`

```go
type cardCommandHandleApi struct {
	card pb.CardCommandServiceClient
	logger logger.LoggerInterface
	mapper apimapper.CardCommandResponseMapper
	trace trace.Tracer
	requestCounter *prometheus.CounterVec
	requestDuration *prometheus.HistogramVec
}
```

#### Methods

##### `CreateCard`

@Security Bearer
@Summary Create a new card
@Tags Card-Command
@Description Create a new card for a user
@Accept json
@Produce json
@Param CreateCardRequest body requests.CreateCardRequest true "Create card request"
@Success 200 {object} response.ApiResponseCard "Created card"
@Failure 400 {object} response.ErrorResponse "Bad request or validation error"
@Failure 500 {object} response.ErrorResponse "Failed to create card"
@Router /api/card/create [post]

```go
func (h *cardCommandHandleApi) CreateCard(c echo.Context) error
```

##### `DeleteAllCardPermanent`

@Security Bearer.
@Summary Permanently delete all card records
@Tags Card-Command
@Description Permanently delete all card records from the database.
@Accept json
@Produce json
@Success 200 {object} response.ApiResponseCardAll "Successfully deleted all card records permanently"
@Failure 500 {object} response.ErrorResponse "Failed to permanently delete all card records"
@Router /api/card/permanent/all [post]

```go
func (h *cardCommandHandleApi) DeleteAllCardPermanent(c echo.Context) error
```

##### `DeleteCardPermanent`

@Security Bearer
@Summary Delete a card permanently
@Tags Card-Command
@Description Delete a card by its ID permanently
@Accept json
@Produce json
@Param id path int true "Card ID"
@Success 200 {object} response.ApiResponseCardDelete "Deleted card"
@Failure 400 {object} response.ErrorResponse "Bad request or invalid ID"
@Failure 500 {object} response.ErrorResponse "Failed to delete card"
@Router /api/card/permanent/{id} [delete]

```go
func (h *cardCommandHandleApi) DeleteCardPermanent(c echo.Context) error
```

##### `RestoreAllCard`

@Security Bearer
@Summary Restore all card records
@Tags Card-Command
@Description Restore all card records that were previously deleted.
@Accept json
@Produce json
@Success 200 {object} response.ApiResponseCardAll "Successfully restored all card records"
@Failure 500 {object} response.ErrorResponse "Failed to restore all card records"
@Router /api/card/restore/all [post]

```go
func (h *cardCommandHandleApi) RestoreAllCard(c echo.Context) error
```

##### `RestoreCard`

@Security Bearer
@Summary Restore a card
@Tags Card-Command
@Description Restore a card by its ID
@Accept json
@Produce json
@Param id path int true "Card ID"
@Success 200 {object} response.ApiResponseCard "Restored card"
@Failure 400 {object} response.ErrorResponse "Bad request or invalid ID"
@Failure 500 {object} response.ErrorResponse "Failed to restore card"
@Router /api/card/restore/{id} [post]

```go
func (h *cardCommandHandleApi) RestoreCard(c echo.Context) error
```

##### `TrashedCard`

@Security Bearer
@Summary Trashed a card
@Tags Card-Command
@Description Trashed a card by its ID
@Accept json
@Produce json
@Param id path int true "Card ID"
@Success 200 {object} response.ApiResponseCard "Trashed card"
@Failure 400 {object} response.ErrorResponse "Bad request or invalid ID"
@Failure 500 {object} response.ErrorResponse "Failed to trashed card"
@Router /api/card/trashed/{id} [post]

```go
func (h *cardCommandHandleApi) TrashedCard(c echo.Context) error
```

##### `UpdateCard`

@Security Bearer
@Summary Update a card
@Tags Card-Command
@Description Update a card for a user
@Accept json
@Produce json
@Param id path int true "Card ID"
@Param UpdateCardRequest body requests.UpdateCardRequest true "Update card request"
@Success 200 {object} response.ApiResponseCard "Updated card"
@Failure 400 {object} response.ErrorResponse "Bad request or validation error"
@Failure 500 {object} response.ErrorResponse "Failed to update card"
@Router /api/card/update/{id} [post]

```go
func (h *cardCommandHandleApi) UpdateCard(c echo.Context) error
```

##### `recordMetrics`

```go
func (s *cardCommandHandleApi) recordMetrics(method string, status string, start time.Time)
```

##### `startTracingAndLogging`

```go
func (s *cardCommandHandleApi) startTracingAndLogging(ctx context.Context, method string, attrs ...attribute.KeyValue) (end func(), logSuccess func(string, ...zap.Field), logError func(string, error, ...zap.Field))
```

### `cardCommandHandleApiParams`

cardHandleParams contains the dependencies required to initialize the card handler.

This struct is typically passed to a constructor function to set up routes
and initialize the `cardCommandHandleParams`.

```go
type cardCommandHandleApiParams struct {
	client pb.CardCommandServiceClient
	router *echo.Echo
	logger logger.LoggerInterface
	mapper apimapper.CardCommandResponseMapper
}
```

### `cardDashboardHandleApi`

```go
type cardDashboardHandleApi struct {
	card pb.CardDashboardServiceClient
	logger logger.LoggerInterface
	mapper apimapper.CardDashboardResponseMapper
	trace trace.Tracer
	requestCounter *prometheus.CounterVec
	requestDuration *prometheus.HistogramVec
}
```

#### Methods

##### `DashboardCard`

DashboardCard godoc
@Summary Get dashboard card data
@Description Retrieve dashboard card data
@Tags Card-Dashboard
@Security Bearer
@Produce json
@Success 200 {object} response.ApiResponseDashboardCard
@Failure 500 {object} response.ErrorResponse
@Router /api/card/dashboard [get]

```go
func (h *cardDashboardHandleApi) DashboardCard(c echo.Context) error
```

##### `DashboardCardCardNumber`

DashboardCardCardNumber godoc
@Summary Get dashboard card data by card number
@Description Retrieve dashboard card data for a specific card number
@Tags Card-Dashboard
@Security Bearer
@Produce json
@Param cardNumber path string true "Card Number"
@Success 200 {object} response.ApiResponseDashboardCardNumber
@Failure 400 {object} response.ErrorResponse
@Failure 500 {object} response.ErrorResponse
@Router /api/card/dashboard/{cardNumber} [get]

```go
func (h *cardDashboardHandleApi) DashboardCardCardNumber(c echo.Context) error
```

##### `recordMetrics`

```go
func (s *cardDashboardHandleApi) recordMetrics(method string, status string, start time.Time)
```

##### `startTracingAndLogging`

```go
func (s *cardDashboardHandleApi) startTracingAndLogging(ctx context.Context, method string, attrs ...attribute.KeyValue) (end func(), logSuccess func(string, ...zap.Field), logError func(string, error, ...zap.Field))
```

### `cardDashboardHandleApiParams`

```go
type cardDashboardHandleApiParams struct {
	client pb.CardDashboardServiceClient
	router *echo.Echo
	logger logger.LoggerInterface
	mapper apimapper.CardDashboardResponseMapper
}
```

### `cardQueryHandleApi`

```go
type cardQueryHandleApi struct {
	card pb.CardQueryServiceClient
	logger logger.LoggerInterface
	mapper apimapper.CardQueryResponseMapper
	trace trace.Tracer
	requestCounter *prometheus.CounterVec
	requestDuration *prometheus.HistogramVec
}
```

#### Methods

##### `FindAll`

FindAll godoc
@Summary Retrieve all cards
@Tags Card-Query
@Security Bearer
@Description Retrieve all cards with pagination
@Accept json
@Produce json
@Param page query int false "Page number"
@Param page_size query int false "Number of data per page"
@Param search query string false "Search keyword"
@Success 200 {object} response.ApiResponsePaginationCard "Card data"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve card data"
@Router /api/card [get]

```go
func (h *cardQueryHandleApi) FindAll(c echo.Context) error
```

##### `FindByActive`

@Security Bearer
@Summary Retrieve active card by Saldo ID
@Tags Card-Query
@Description Retrieve an active card associated with a Saldo ID
@Accept json
@Produce json
@Success 200 {object} response.ApiResponsePaginationCardDeleteAt "Card data"
@Failure 400 {object} response.ErrorResponse "Invalid Saldo ID"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve card record"
@Router /api/card/active [get]

```go
func (h *cardQueryHandleApi) FindByActive(c echo.Context) error
```

##### `FindByCardNumber`

@Security Bearer
@Summary Retrieve card by card number
@Tags Card-Query
@Description Retrieve a card by its card number
@Accept json
@Produce json
@Param card_number path string true "Card number"
@Success 200 {object} response.ApiResponseCard "Card data"
@Failure 400 {object} response.ErrorResponse "Failed to fetch card record"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve card record"
@Router /api/card/{card_number} [get]

```go
func (h *cardQueryHandleApi) FindByCardNumber(c echo.Context) error
```

##### `FindById`

FindById godoc
@Summary Retrieve card by ID
@Tags Card-Query
@Security Bearer
@Description Retrieve a card by its ID
@Accept json
@Produce json
@Param id path int true "Card ID"
@Success 200 {object} response.ApiResponseCard "Card data"
@Failure 400 {object} response.ErrorResponse "Invalid card ID"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve card record"
@Router /api/card/{id} [get]

```go
func (h *cardQueryHandleApi) FindById(c echo.Context) error
```

##### `FindByTrashed`

@Summary Retrieve trashed cards
@Tags Card-Query
@Security Bearer
@Description Retrieve a list of trashed cards
@Accept json
@Produce json
@Param page query int false "Page number (default: 1)"
@Param page_size query int false "Number of items per page (default: 10)"
@Param search query string false "Search keyword"
@Success 200 {object} response.ApiResponsePaginationCardDeleteAt "Card data"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve card record"
@Router /api/card/trashed [get]

```go
func (h *cardQueryHandleApi) FindByTrashed(c echo.Context) error
```

##### `FindByUserID`

FindByUserID godoc
@Summary Retrieve cards by user ID
@Tags Card-Query
@Security Bearer
@Description Retrieve a list of cards associated with a user by their ID
@Accept json
@Produce json
@Success 200 {object} response.ApiResponseCard "Card data"
@Failure 400 {object} response.ErrorResponse "Invalid user ID"
@Failure 500 {object} response.ErrorResponse "Failed to retrieve card record"
@Router /api/card/user [get]

```go
func (h *cardQueryHandleApi) FindByUserID(c echo.Context) error
```

##### `recordMetrics`

```go
func (s *cardQueryHandleApi) recordMetrics(method string, status string, start time.Time)
```

##### `startTracingAndLogging`

```go
func (s *cardQueryHandleApi) startTracingAndLogging(ctx context.Context, method string, attrs ...attribute.KeyValue) (end func(), logSuccess func(string, ...zap.Field), logError func(string, error, ...zap.Field))
```

### `cardQueryHandleApiParams`

```go
type cardQueryHandleApiParams struct {
	client pb.CardQueryServiceClient
	router *echo.Echo
	logger logger.LoggerInterface
	mapper apimapper.CardQueryResponseMapper
}
```

### `cardStatsBalanceHandleApi`

```go
type cardStatsBalanceHandleApi struct {
	card pb.CardStatsBalanceServiceClient
	logger logger.LoggerInterface
	mapper apimapper.CardStatsBalanceResponseMapper
	trace trace.Tracer
	requestCounter *prometheus.CounterVec
	requestDuration *prometheus.HistogramVec
}
```

#### Methods

##### `FindMonthlyBalance`

FindMonthlyBalance godoc
@Summary Get monthly balance data
@Description Retrieve monthly balance data for a specific year
@Tags Card-Stats-Balance
@Security Bearer
@Produce json
@Param year query int true "Year"
@Success 200 {object} response.ApiResponseMonthlyBalance
@Failure 400 {object} response.ErrorResponse
@Failure 500 {object} response.ErrorResponse
@Router /api/card/monthly-balance [get]

```go
func (h *cardStatsBalanceHandleApi) FindMonthlyBalance(c echo.Context) error
```

##### `FindMonthlyBalanceByCardNumber`

FindMonthlyBalanceByCardNumber godoc
@Summary Get monthly balance data by card number
@Description Retrieve monthly balance data for a specific year and card number
@Tags Card-Stats-Balance
@Security Bearer
@Produce json
@Param year query int true "Year"
@Param card_number query string true "Card Number"
@Success 200 {object} response.ApiResponseMonthlyBalance
@Failure 400 {object} response.ErrorResponse
@Failure 500 {object} response.ErrorResponse
@Router /api/card/monthly-balance-by-card [get]

```go
func (h *cardStatsBalanceHandleApi) FindMonthlyBalanceByCardNumber(c echo.Context) error
```

##### `FindYearlyBalance`

FindYearlyBalance godoc
@Summary Get yearly balance data
@Description Retrieve yearly balance data for a specific year
@Tags Card-Stats-Balance
@Security Bearer
@Produce json
@Param year query int true "Year"
@Success 200 {object} response.ApiResponseYearlyBalance
@Failure 400 {object} response.ErrorResponse
@Failure 500 {object} response.ErrorResponse
@Router /api/card/yearly-balance [get]

```go
func (h *cardStatsBalanceHandleApi) FindYearlyBalance(c echo.Context) error
```

##### `FindYearlyBalanceByCardNumber`

FindYearlyBalanceByCardNumber godoc
@Summary Get yearly balance data by card number
@Description Retrieve yearly balance data for a specific year and card number
@Tags Card-Stats-Balance
@Security Bearer
@Produce json
@Param year query int true "Year"
@Param card_number query string true "Card Number"
@Success 200 {object} response.ApiResponseYearlyBalance
@Failure 400 {object} response.ErrorResponse
@Failure 500 {object} response.ErrorResponse
@Router /api/card/yearly-balance-by-card [get]

```go
func (h *cardStatsBalanceHandleApi) FindYearlyBalanceByCardNumber(c echo.Context) error
```

##### `recordMetrics`

recordMetrics records a Prometheus metric for the given method and status.
It increments a counter and records the duration since the provided start time.

```go
func (s *cardStatsBalanceHandleApi) recordMetrics(method string, status string, start time.Time)
```

##### `startTracingAndLogging`

startTracingAndLogging starts a tracing span and returns functions to log the outcome of the call.
The returned functions are logSuccess and logError, which log the outcome of the call to the trace span.
The returned end function records the metrics and ends the trace span.

```go
func (s *cardStatsBalanceHandleApi) startTracingAndLogging(ctx context.Context, method string, attrs ...attribute.KeyValue) (end func(), logSuccess func(string, ...zap.Field), logError func(string, error, ...zap.Field))
```

### `cardStatsBalanceHandleApiParams`

```go
type cardStatsBalanceHandleApiParams struct {
	client pb.CardStatsBalanceServiceClient
	router *echo.Echo
	logger logger.LoggerInterface
	mapper apimapper.CardStatsBalanceResponseMapper
}
```

### `cardStatsTopupHandleApi`

```go
type cardStatsTopupHandleApi struct {
	card pb.CardStatsTopupServiceClient
	logger logger.LoggerInterface
	mapper apimapper.CardStatsAmountResponseMapper
	trace trace.Tracer
	requestCounter *prometheus.CounterVec
	requestDuration *prometheus.HistogramVec
}
```

#### Methods

##### `FindMonthlyTopupAmount`

FindMonthlyTopupAmount godoc
@Summary Get monthly topup amount data
@Description Retrieve monthly topup amount data for a specific year
@Tags Card-Stats-Topup
@Security Bearer
@Produce json
@Param year query int true "Year"
@Success 200 {object} response.ApiResponseMonthlyAmount
@Failure 400 {object} response.ErrorResponse
@Failure 500 {object} response.ErrorResponse
@Router /api/card/monthly-topup-amount [get]

```go
func (h *cardStatsTopupHandleApi) FindMonthlyTopupAmount(c echo.Context) error
```

##### `FindMonthlyTopupAmountByCardNumber`

FindMonthlyTopupAmountByCardNumber godoc
@Summary Get monthly topup amount data by card number
@Description Retrieve monthly topup amount data for a specific year and card number
@Tags Card-Stats-Topup
@Security Bearer
@Produce json
@Param year query int true "Year"
@Param card_number query string true "Card Number"
@Success 200 {object} response.ApiResponseMonthlyAmount
@Failure 400 {object} response.ErrorResponse
@Failure 500 {object} response.ErrorResponse
@Router /api/card/monthly-topup-amount-by-card [get]

```go
func (h *cardStatsTopupHandleApi) FindMonthlyTopupAmountByCardNumber(c echo.Context) error
```

##### `FindYearlyTopupAmount`

FindYearlyTopupAmount godoc
@Summary Get yearly topup amount data
@Description Retrieve yearly topup amount data for a specific year
@Tags Card-Stats-Topup
@Security Bearer
@Produce json
@Param year query int true "Year"
@Success 200 {object} response.ApiResponseYearlyAmount
@Failure 400 {object} response.ErrorResponse
@Failure 500 {object} response.ErrorResponse
@Router /api/topup/yearly-topup-amount [get]

```go
func (h *cardStatsTopupHandleApi) FindYearlyTopupAmount(c echo.Context) error
```

##### `FindYearlyTopupAmountByCardNumber`

FindYearlyTopupAmountByCardNumber godoc
@Summary Get yearly topup amount data by card number
@Description Retrieve yearly topup amount data for a specific year and card number
@Tags Card-Stats-Topup
@Security Bearer
@Produce json
@Param year query int true "Year"
@Param card_number query string true "Card Number"
@Success 200 {object} response.ApiResponseYearlyAmount
@Failure 400 {object} response.ErrorResponse
@Failure 500 {object} response.ErrorResponse
@Router /api/card/yearly-topup-amount-by-card [get]

```go
func (h *cardStatsTopupHandleApi) FindYearlyTopupAmountByCardNumber(c echo.Context) error
```

##### `recordMetrics`

recordMetrics records Prometheus metrics for the given method and status.
It increments the request counter and observes the request duration
for the given method and status, using the provided start time.

```go
func (s *cardStatsTopupHandleApi) recordMetrics(method string, status string, start time.Time)
```

##### `startTracingAndLogging`

startTracingAndLogging starts a tracing span and returns functions to log the outcome of the call.
The returned functions are logSuccess and logError, which log the outcome of the call to the trace span.
The returned end function records the metrics and ends the trace span.

```go
func (s *cardStatsTopupHandleApi) startTracingAndLogging(ctx context.Context, method string, attrs ...attribute.KeyValue) (end func(), logSuccess func(string, ...zap.Field), logError func(string, error, ...zap.Field))
```

### `cardStatsTopupHandleApiParams`

```go
type cardStatsTopupHandleApiParams struct {
	client pb.CardStatsTopupServiceClient
	router *echo.Echo
	logger logger.LoggerInterface
	mapper apimapper.CardStatsAmountResponseMapper
}
```

### `cardStatsTransactionHandleApi`

```go
type cardStatsTransactionHandleApi struct {
	card pb.CardStatsTransactonServiceClient
	logger logger.LoggerInterface
	mapper apimapper.CardStatsAmountResponseMapper
	trace trace.Tracer
	requestCounter *prometheus.CounterVec
	requestDuration *prometheus.HistogramVec
}
```

#### Methods

##### `FindMonthlyTransactionAmount`

FindMonthlyTransactionAmount godoc
@Summary Get monthly transaction amount data
@Description Retrieve monthly transaction amount data for a specific year
@Tags Card-Stats-Transaction
@Security Bearer
@Produce json
@Param year query int true "Year"
@Success 200 {object} response.ApiResponseMonthlyAmount
@Failure 400 {object} response.ErrorResponse
@Failure 500 {object} response.ErrorResponse
@Router /api/card-stats-transaction/monthly-transaction-amount [get]

```go
func (h *cardStatsTransactionHandleApi) FindMonthlyTransactionAmount(c echo.Context) error
```

##### `FindMonthlyTransactionAmountByCardNumber`

FindMonthlyTransactionAmountByCardNumber godoc
@Summary Get monthly transaction amount data by card number
@Description Retrieve monthly transaction amount data for a specific year and card number
@Tags Card-Stats-Transaction
@Security Bearer
@Produce json
@Param year query int true "Year"
@Param card_number query string true "Card Number"
@Success 200 {object} response.ApiResponseMonthlyAmount
@Failure 400 {object} response.ErrorResponse
@Failure 500 {object} response.ErrorResponse
@Router /api/card-stats-transaction/monthly-transaction-amount-by-card [get]

```go
func (h *cardStatsTransactionHandleApi) FindMonthlyTransactionAmountByCardNumber(c echo.Context) error
```

##### `FindYearlyTransactionAmount`

FindYearlyTransactionAmount godoc
@Summary Get yearly transaction amount data
@Description Retrieve yearly transaction amount data for a specific year
@Tags Card-Stats-Transaction
@Security Bearer
@Produce json
@Param year query int true "Year"
@Success 200 {object} response.ApiResponseYearlyAmount
@Failure 400 {object} response.ErrorResponse
@Failure 500 {object} response.ErrorResponse
@Router /api/card-stats-transaction/yearly-transaction-amount [get]

```go
func (h *cardStatsTransactionHandleApi) FindYearlyTransactionAmount(c echo.Context) error
```

##### `FindYearlyTransactionAmountByCardNumber`

FindYearlyTransactionAmountByCardNumber godoc
@Summary Get yearly transaction amount data by card number
@Description Retrieve yearly transaction amount data for a specific year and card number
@Tags Card-Stats-Transaction
@Security Bearer
@Produce json
@Param year query int true "Year"
@Param card_number query string true "Card Number"
@Success 200 {object} response.ApiResponseYearlyAmount
@Failure 400 {object} response.ErrorResponse
@Failure 500 {object} response.ErrorResponse
@Router /api/card-stats-transaction/yearly-transaction-amount-by-card [get]

```go
func (h *cardStatsTransactionHandleApi) FindYearlyTransactionAmountByCardNumber(c echo.Context) error
```

##### `recordMetrics`

recordMetrics records a Prometheus metric for the given method and status.
It increments a counter and records the duration since the provided start time.

```go
func (s *cardStatsTransactionHandleApi) recordMetrics(method string, status string, start time.Time)
```

##### `startTracingAndLogging`

startTracingAndLogging initializes tracing and logging for a given method within the context.
It returns three functions: `end` to conclude the tracing and log metrics,
`logSuccess` to log successful events, and `logError` to log errors.

Parameters:
- ctx: The context in which the tracing and logging occur.
- method: The name of the method being traced and logged.
- attrs: Optional key-value attributes to be set on the trace span.

```go
func (s *cardStatsTransactionHandleApi) startTracingAndLogging(ctx context.Context, method string, attrs ...attribute.KeyValue) (end func(), logSuccess func(string, ...zap.Field), logError func(string, error, ...zap.Field))
```

### `cardStatsTransactionHandleApiParams`

```go
type cardStatsTransactionHandleApiParams struct {
	client pb.CardStatsTransactonServiceClient
	router *echo.Echo
	logger logger.LoggerInterface
	mapper apimapper.CardStatsAmountResponseMapper
}
```

### `cardStatsTransferHandleApi`

```go
type cardStatsTransferHandleApi struct {
	card pb.CardStatsTransferServiceClient
	logger logger.LoggerInterface
	mapper apimapper.CardStatsAmountResponseMapper
	trace trace.Tracer
	requestCounter *prometheus.CounterVec
	requestDuration *prometheus.HistogramVec
}
```

#### Methods

##### `FindMonthlyTransferReceiverAmount`

FindMonthlyTransferReceiverAmount godoc
@Summary Get monthly transfer receiver amount data
@Description Retrieve monthly transfer receiver amount data for a specific year
@Tags Card-Stats-Transfer
@Security Bearer
@Produce json
@Param year query int true "Year"
@Success 200 {object} response.ApiResponseMonthlyAmount
@Failure 400 {object} response.ErrorResponse
@Failure 500 {object} response.ErrorResponse
@Router /api/card-stats-transfer/monthly-transfer-receiver-amount [get]

```go
func (h *cardStatsTransferHandleApi) FindMonthlyTransferReceiverAmount(c echo.Context) error
```

##### `FindMonthlyTransferReceiverAmountByCardNumber`

FindMonthlyTransferReceiverAmountByCardNumber godoc
@Summary Get monthly transfer receiver amount by card number
@Description Retrieve the total amount received by a specific card number in a given year, broken down by month
@Tags Card-Stats-Transfer
@Security Bearer
@Accept json
@Produce json
@Param year query int true "Year for which the data is requested"
@Param card_number query string true "Card number for which the data is requested"
@Success 200 {object} response.ApiResponseYearlyAmount
@Failure 400 {object} response.ErrorResponse
@Failure 500 {object} response.ErrorResponse
@Router /api/card-stats-transfer/monthly-transfer-receiver-amount-by-card [get]

```go
func (h *cardStatsTransferHandleApi) FindMonthlyTransferReceiverAmountByCardNumber(c echo.Context) error
```

##### `FindMonthlyTransferSenderAmount`

FindMonthlyTransferSenderAmount godoc
@Summary Get monthly transfer sender amount data
@Description Retrieve monthly transfer sender amount data for a specific year
@Tags Card-Stats-Transfer
@Security Bearer
@Produce json
@Param year query int true "Year"
@Success 200 {object} response.ApiResponseMonthlyAmount
@Failure 400 {object} response.ErrorResponse
@Failure 500 {object} response.ErrorResponse
@Router /api/card-stats-transfer/monthly-transfer-sender-amount [get]

```go
func (h *cardStatsTransferHandleApi) FindMonthlyTransferSenderAmount(c echo.Context) error
```

##### `FindMonthlyTransferSenderAmountByCardNumber`

FindMonthlyTransferSenderAmountByCardNumber godoc
@Summary Get monthly transfer sender amount data by card number
@Description Retrieve monthly transfer sender amount data for a specific year and card number
@Tags Card-Stats-Transfer
@Security Bearer
@Produce json
@Param year query int true "Year"
@Param card_number query string true "Card Number"
@Success 200 {object} response.ApiResponseMonthlyAmount
@Failure 400 {object} response.ErrorResponse
@Failure 500 {object} response.ErrorResponse
@Router /api/card-stats-transfer/monthly-transfer-sender-amount-by-card [get]

```go
func (h *cardStatsTransferHandleApi) FindMonthlyTransferSenderAmountByCardNumber(c echo.Context) error
```

##### `FindYearlyTransferReceiverAmount`

FindYearlyTransferReceiverAmount godoc
@Summary Get yearly transfer receiver amount data
@Description Retrieve yearly transfer receiver amount data for a specific year
@Tags Card-Stats-Transfer
@Security Bearer
@Produce json
@Param year query int true "Year"
@Success 200 {object} response.ApiResponseYearlyAmount
@Failure 400 {object} response.ErrorResponse
@Failure 500 {object} response.ErrorResponse
@Router /api/card-stats-transfer/yearly-transfer-receiver-amount [get]

```go
func (h *cardStatsTransferHandleApi) FindYearlyTransferReceiverAmount(c echo.Context) error
```

##### `FindYearlyTransferReceiverAmountByCardNumber`

FindYearlyTransferReceiverAmountByCardNumber godoc
@Summary Get yearly transfer receiver amount by card number
@Description Retrieve the total amount received by a specific card number in a given year
@Tags Card-Stats-Transfer
@Security Bearer
@Accept json
@Produce json
@Param year query int true "Year for which the data is requested"
@Param card_number query string true "Card number for which the data is requested"
@Success 200 {object} response.ApiResponseYearlyAmount
@Failure 400 {object} response.ErrorResponse
@Failure 500 {object} response.ErrorResponse
@Router /api/card-stats-transfer/yearly-transfer-receiver-amount-by-card [get]

```go
func (h *cardStatsTransferHandleApi) FindYearlyTransferReceiverAmountByCardNumber(c echo.Context) error
```

##### `FindYearlyTransferSenderAmount`

FindYearlyTransferSenderAmount godoc
@Summary Get yearly transfer sender amount data
@Description Retrieve yearly transfer sender amount data for a specific year
@Tags Card-Stats-Transfer
@Security Bearer
@Produce json
@Param year query int true "Year"
@Success 200 {object} response.ApiResponseYearlyAmount
@Failure 400 {object} response.ErrorResponse
@Failure 500 {object} response.ErrorResponse
@Router /api/card-stats-transfer/yearly-transfer-sender-amount [get]

```go
func (h *cardStatsTransferHandleApi) FindYearlyTransferSenderAmount(c echo.Context) error
```

##### `FindYearlyTransferSenderAmountByCardNumber`

FindYearlyTransferSenderAmountByCardNumber godoc
@Summary Get yearly transfer sender amount by card number
@Description Retrieve the total amount sent by a specific card number in a given year
@Tags Card-Stats-Transfer
@Security Bearer
@Accept json
@Produce json
@Param year query int true "Year for which the data is requested"
@Param card_number query string true "Card number for which the data is requested"
@Success 200 {object} response.ApiResponseYearlyAmount
@Failure 400 {object} response.ErrorResponse
@Failure 500 {object} response.ErrorResponse
@Router /api/card-stats-transfer/yearly-transfer-sender-amount-by-card [get]

```go
func (h *cardStatsTransferHandleApi) FindYearlyTransferSenderAmountByCardNumber(c echo.Context) error
```

##### `recordMetrics`

recordMetrics records Prometheus metrics for the specified method and status.
It increments the request counter for the provided method and status,
and observes the request duration by calculating the time elapsed since the provided start time.

```go
func (s *cardStatsTransferHandleApi) recordMetrics(method string, status string, start time.Time)
```

##### `startTracingAndLogging`

startTracingAndLogging returns three functions: end, logSuccess and logError.
The end function must be called to end the tracing and logging.
The logSuccess and logError functions can be used to log success and error messages.
If an error is passed to logError, it will be recorded as an error event on the span.

```go
func (s *cardStatsTransferHandleApi) startTracingAndLogging(ctx context.Context, method string, attrs ...attribute.KeyValue) (end func(), logSuccess func(string, ...zap.Field), logError func(string, error, ...zap.Field))
```

### `cardStatsTransferHandleApiParams`

```go
type cardStatsTransferHandleApiParams struct {
	client pb.CardStatsTransferServiceClient
	router *echo.Echo
	logger logger.LoggerInterface
	mapper apimapper.CardStatsAmountResponseMapper
}
```

### `cardStatsWithdrawHandleApi`

```go
type cardStatsWithdrawHandleApi struct {
	card pb.CardStatsWithdrawServiceClient
	logger logger.LoggerInterface
	mapper apimapper.CardStatsAmountResponseMapper
	trace trace.Tracer
	requestCounter *prometheus.CounterVec
	requestDuration *prometheus.HistogramVec
}
```

#### Methods

##### `FindMonthlyWithdrawAmount`

FindMonthlyWithdrawAmount
godoc
@Summary Get monthly withdraw amount data
@Description Retrieve monthly withdraw amount data for a specific year
@Tags Card-Stats-Withdraw
@Security Bearer
@Produce json
@Param year query int true "Year"
@Success 200 {object} response.ApiResponseMonthlyAmount
@Failure 400 {object} response.ErrorResponse
@Failure 500 {object} response.ErrorResponse
@Router /api/card-stats-withdraw/monthly-withdraw-amount [get]

```go
func (h *cardStatsWithdrawHandleApi) FindMonthlyWithdrawAmount(c echo.Context) error
```

##### `FindMonthlyWithdrawAmountByCardNumber`

FindMonthlyWithdrawAmountByCardNumber godoc
@Summary Get monthly withdraw amount data by card number
@Description Retrieve monthly withdraw amount data for a specific year and card number
@Tags Card-Stats-Withdraw
@Security Bearer
@Produce json
@Param year query int true "Year"
@Param card_number query string true "Card Number"
@Success 200 {object} response.ApiResponseMonthlyAmount
@Failure 400 {object} response.ErrorResponse
@Failure 500 {object} response.ErrorResponse
@Router /api/card-stats-withdraw/monthly-withdraw-amount-by-card [get]

```go
func (h *cardStatsWithdrawHandleApi) FindMonthlyWithdrawAmountByCardNumber(c echo.Context) error
```

##### `FindYearlyWithdrawAmount`

FindYearlyWithdrawAmount godoc
@Summary Get yearly withdraw amount data
@Description Retrieve yearly withdraw amount data for a specific year
@Tags Card-Stats-Withdraw
@Security Bearer
@Produce json
@Param year query int true "Year"
@Success 200 {object} response.ApiResponseYearlyAmount
@Failure 400 {object} response.ErrorResponse
@Failure 500 {object} response.ErrorResponse
@Router /api/card-stats-withdraw/yearly-withdraw-amount [get]

```go
func (h *cardStatsWithdrawHandleApi) FindYearlyWithdrawAmount(c echo.Context) error
```

##### `FindYearlyWithdrawAmountByCardNumber`

FindYearlyWithdrawAmountByCardNumber godoc
@Summary Get yearly withdraw amount data by card number
@Description Retrieve yearly withdraw amount data for a specific year and card number
@Tags Card-Stats-Withdraw
@Security Bearer
@Produce json
@Param year query int true "Year"
@Param card_number query string true "Card Number"
@Success 200 {object} response.ApiResponseYearlyAmount
@Failure 400 {object} response.ErrorResponse
@Failure 500 {object} response.ErrorResponse
@Router /api/card-stats-withdraw/yearly-withdraw-amount-by-card [get]

```go
func (h *cardStatsWithdrawHandleApi) FindYearlyWithdrawAmountByCardNumber(c echo.Context) error
```

##### `recordMetrics`

recordMetrics records Prometheus metrics for the specified method and status.
It increments the request counter for the provided method and status,
and observes the request duration by calculating the time elapsed since the provided start time.

```go
func (s *cardStatsWithdrawHandleApi) recordMetrics(method string, status string, start time.Time)
```

##### `startTracingAndLogging`

startTracingAndLogging starts a tracing span and returns functions to log the outcome of the call.
The returned functions are logSuccess and logError, which log the outcome of the call to the trace span.
The returned end function records the metrics and ends the trace span.

```go
func (s *cardStatsWithdrawHandleApi) startTracingAndLogging(ctx context.Context, method string, attrs ...attribute.KeyValue) (end func(), logSuccess func(string, ...zap.Field), logError func(string, error, ...zap.Field))
```

### `cardStatsWithdrawHandleApiParams`

```go
type cardStatsWithdrawHandleApiParams struct {
	client pb.CardStatsWithdrawServiceClient
	router *echo.Echo
	logger logger.LoggerInterface
	mapper apimapper.CardStatsAmountResponseMapper
}
```

## 🚀 Functions

### `RegisterCardHandler`

NewCardHandler initializes handlers for various card-related operations.

This function sets up multiple handlers for card operations including query,
command, dashboard, and statistics (balance, transaction, top-up, withdrawal, transfer).
It takes a DepsCard struct which contains the necessary dependencies such as
gRPC client connection, Echo router, and logger. Each handler is initialized
with the corresponding response mapper and added to a slice of handler functions,
which are executed sequentially to set up the routes.

```go
func RegisterCardHandler(deps *DepsCard)
```

### `setupCardCommandHandler`

setupCardCommandHandler sets up the handler for the card command service.

It creates a new instance of the CardCommandHandleApi and registers the
handler with the Echo router. It takes a pointer to DepsCard and a
mapper for CardResponse. It returns a function that can be executed to
set up the handler.

```go
func setupCardCommandHandler(deps *DepsCard, mapper apimapper.CardCommandResponseMapper) func()
```

### `setupCardDashboardHandler`

setupCardDashboardHandler sets up the handler for the card dashboard service.

It creates a new instance of the CardDashboardHandleApi and registers the
handler with the Echo router. It takes a pointer to DepsCard and a
mapper for CardResponse. It returns a function that can be executed to
set up the handler.

```go
func setupCardDashboardHandler(deps *DepsCard, mapper apimapper.CardDashboardResponseMapper) func()
```

### `setupCardQueryHandler`

setupCardQueryHandler sets up the handler for the card query service.

It creates a new instance of the CardQueryHandleApi and registers the
handler with the Echo router. It takes a pointer to DepsCard and a
mapper for CardResponse. It returns a function that can be executed to
set up the handler.

```go
func setupCardQueryHandler(deps *DepsCard, mapper apimapper.CardQueryResponseMapper) func()
```

### `setupCardStatsBalanceHandler`

setupCardStatsBalanceHandler sets up the handler for the card statistics balance service.

It creates a new instance of the CardStatsBalanceHandleApi and registers the
handler with the Echo router. It takes a pointer to DepsCard and a
mapper for CardStatsBalanceResponse. It returns a function that can be executed to
set up the handler.

```go
func setupCardStatsBalanceHandler(deps *DepsCard, mapper apimapper.CardStatsBalanceResponseMapper) func()
```

### `setupCardStatsTopupHandler`

setupCardStatsTopupHandler sets up the handler for the card statistics top-up service.

It creates a new instance of the CardStatsTopupHandleApi and registers the
handler with the Echo router. It takes a pointer to DepsCard and a
mapper for CardStatsAmountResponse. It returns a function that can be executed to
set up the handler.

```go
func setupCardStatsTopupHandler(deps *DepsCard, mapper apimapper.CardStatsAmountResponseMapper) func()
```

### `setupCardStatsTransactionHandler`

setupCardStatsTransactionHandler sets up the handler for the card statistics transaction service.

It creates a new instance of the CardStatsTransactionHandleApi and registers the
handler with the Echo router. It takes a pointer to DepsCard and a
mapper for CardStatsAmountResponse. It returns a function that can be executed to
set up the handler.

```go
func setupCardStatsTransactionHandler(deps *DepsCard, mapper apimapper.CardStatsAmountResponseMapper) func()
```

### `setupCardStatsTransferHandler`

setupCardStatsTransferHandler sets up the handler for the card statistics transfer service.

It creates a new instance of the CardStatsTransferHandleApi and registers the
handler with the Echo router. It takes a pointer to DepsCard and a
mapper for CardStatsAmountResponse. It returns a function that can be executed to
set up the handler.

```go
func setupCardStatsTransferHandler(deps *DepsCard, mapper apimapper.CardStatsAmountResponseMapper) func()
```

### `setupCardStatsWithdrawHandler`

setupCardStatsWithdrawHandler sets up the handler for the card statistics withdraw service.

It creates a new instance of the CardStatsWithdrawHandleApi and registers the
handler with the Echo router. It takes a pointer to DepsCard and a
mapper for CardStatsAmountResponse. It returns a function that can be executed to
set up the handler.

```go
func setupCardStatsWithdrawHandler(deps *DepsCard, mapper apimapper.CardStatsAmountResponseMapper) func()
```

