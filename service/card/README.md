# Payment Gateway Card Service

## Overview

`CardService` is a monolith responsible for card management and balance tracking within the system. It supports a wide range of operations, including card creation, updates, deletions (soft and hard), as well as balance/statistics reporting based on various dimensions such as monthly, yearly, and card number specific views.

This service communicates via gRPC, tracks observability via Prometheus and OpenTelemetry, and integrates with Kafka to publish balance creation events to the saldo-service.


### 🔄 Service Architecture

`CardService` is split into several subservices for clean separation of concerns:

- **CommandService**: Handles create/update/delete and produces Kafka events.
- **QueryService**: Handles all data retrieval operations.
- **DashboardService**: Aggregates metrics for dashboards.
- **StatisticService**: Provides time-based and card-based statistics.

All services are instrumented with:
- **Kafka** to produce and consume balance creation events.
- **Prometheus metrics** to track request counts and durations for observability.
- **OpenTelemetry tracing** to provide distributed tracing for requests, helping with debugging and performance tuning.
- **Structured logging** with Zap for detailed error and operational logs.

----

### 📌 Available RPC Methods

#### 📘 Query Operations

| Method                    | Description                               |
| ------------------------- | ----------------------------------------- |
| `FindAllCard`             | Fetch all cards with pagination.          |
| `FindByIdCard`            | Fetch a card by its ID.                   |
| `FindByUserIdCard`        | Fetch cards belonging to a specific user. |
| `FindByActiveCard`        | Get all active cards.                     |
| `FindByTrashedCard`       | Get soft-deleted cards.                   |
| `FindByCardNumber`        | Fetch a card using its card number.       |
| `FindUserCardByCardNumber`| Fetch user information and card details.  |

#### ✏️ Command Service Methods

| Method                   | Description                                     |
| ------------------------ | ----------------------------------------------- |
| `CreateCard`             | Create a new card and produce a Kafka event.    |
| `UpdateCard`             | Update card details.                            |
| `TrashedCard`            | Soft-delete a card.                             |
| `RestoreCard`            | Restore a soft-deleted card.                    |
| `DeleteCardPermanent`    | Permanently delete a card.                      |
| `RestoreAllCard`         | Restore all soft-deleted cards.                 |
| `DeleteAllCardPermanent` | Permanently delete all soft-deleted cards.      |
| `ToggleCardStatus`       | Toggle active/suspended card status.            |
| `UpdateCreditLimit`      | Update the credit limit for a card.             |
| `RedeemPoints`           | Redeem card reward points.                      |
| `ProcessBillingCycles`   | Process monthly billing cycles.                 |

#### 📊 Statistics Service Methods

| Method                                                                                           | Description                                        |
| ------------------------------------------------------------------------------------------------ | -------------------------------------------------- |
| `FindMonthlyBalance` / `FindYearlyBalance`                                                       | Monthly/Yearly balance summaries.                  |
| `FindMonthlyTopupAmount` / `FindYearlyTopupAmount`                                               | Monthly/Yearly top-up stats.                       |
| `FindMonthlyWithdrawAmount` / `FindYearlyWithdrawAmount`                                         | Monthly/Yearly withdrawal stats.                   |
| `FindMonthlyTransactionAmount` / `FindYearlyTransactionAmount`                                   | Monthly/Yearly transaction stats.                  |
| `FindMonthlyTransferSenderAmount` / `FindYearlyTransferSenderAmount`                             | Monthly/Yearly transfer (sent).                    |
| `FindMonthlyTransferReceiverAmount` / `FindYearlyTransferReceiverAmount`                         | Monthly/Yearly transfer (received).                |
| `FindMonthlyBalanceByCardNumber` / `FindYearlyBalanceByCardNumber`                               | Balance stats filtered by card number.             |
| `FindMonthlyTopupAmountByCardNumber` / `FindYearlyTopupAmountByCardNumber`                       | Top-up stats filtered by card number.              |
| `FindMonthlyWithdrawAmountByCardNumber` / `FindYearlyWithdrawAmountByCardNumber`                 | Withdrawal stats filtered by card number.          |
| `FindMonthlyTransactionAmountByCardNumber` / `FindYearlyTransactionAmountByCardNumber`           | Transaction stats filtered by card number.         |
| `FindMonthlyTransferSenderAmountByCardNumber` / `FindYearlyTransferSenderAmountByCardNumber`     | Transfer (sent) stats filtered by card number.     |
| `FindMonthlyTransferReceiverAmountByCardNumber` / `FindYearlyTransferReceiverAmountByCardNumber` | Transfer (received) stats filtered by card number. |
| `FindDailyTopupAmount` / `FindDailyTopupAmountByCardNumber`                                       | Daily top-up stats globally or by card number.     |
| `FindDailyWithdrawAmount` / `FindDailyWithdrawAmountByCardNumber`                                 | Daily withdrawal stats globally or by card number. |
| `FindDailyTransactionAmount` / `FindDailyTransactionAmountByCardNumber`                           | Daily transaction stats globally or by card number.|
| `FindDailyTransferSenderAmount` / `FindDailyTransferSenderAmountByCardNumber`                     | Daily transfer (sent) stats globally/by card number.|
| `FindDailyTransferReceiverAmount` / `FindDailyTransferReceiverAmountByCardNumber`                 | Daily transfer (received) stats globally/by card.  |


#### 📊 Dashboard Service Methods

| Method                | Description                             |
| --------------------- | --------------------------------------- |
| `DashboardCard`       | Fetch dashboard overview for all cards. |
| `DashboardCardNumber` | Fetch dashboard data by card number.    |


---

### 📊 Monitoring & Observability

#### Prometheus Metrics

Each RPC method is instrumented with:
- **Request Counter**: total number of requests by method and status.
- **Request Duration Histogram**: duration of each RPC method execution.

**Metric examples**:
- `card_command_service_requests_total`
- `card_command_service_request_duration_seconds`
- `card_dashboard_request_count`
- `card_dashboard_request_duration_seconds`
- `card_query_service_requests_total`
- `card_query_service_request_duration_seconds`
- `card_statistic_service_requests_total`
- `card_statistic_service_request_duration_seconds`
- `card_statistic_bycard_service_requests_total`
- `card_statistic_bycard_service_request_duration_seconds`

#### 🔍 Tracing with OpenTelemetry

Each service is assigned a tracer for distributed tracing:
- `card-query-service`
- `card-command-service`
- `card-dashboard-service`
- `card-statistic-service`
- `card-statistic-bycard-service`

----

### 📬 Kafka Integration

When a card is created, CardService emits an event to the saldo-service-topic-create-saldo topic.

| Kafka Topic                        | Description                                              | Consumer Group        |
| ---------------------------------- | -------------------------------------------------------- | --------------------- |
| `saldo-service-topic-create-saldo` | Sent when a new card is created to trigger balance setup | `saldo-service-group` |
