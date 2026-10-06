# Payment Gateway Transaction Service

## Overview

``TransactionService`` is a monolith responsible for managing balance top-up transactions. It supports monthly and yearly statistics based on method, amount, status, and card number. This service is built with strong observability using Prometheus, OpenTelemetry, and Zap Logger.


### 🔄 Service Architecture
`TransactionService` is divided into multiple components to ensure scalability and clear separation of concerns:
  - **transactionCommandService**
    Handles write operations: CreateTransaction, UpdateTransaction, TrashedTransaction, RestoreTransaction, and DeleteTransactionPermanent.
  - **transactionQueryService**
    Handles read operations: FindAllTransaction, FindByIdTransaction, FindByCardNumberTransaction, and statistical queries.
  - **transactionStatisticService**
    Provides general transaction statistics for all cards.
  - **transactionStatisticByCardService**
    Provides transaction statistics grouped by card number.


### 📌 Available RPC Methods

#### 📘 Query Service Methods
| Method                           | Description                                    |
| -------------------------------- | ---------------------------------------------- |
| `FindAllTransaction`             | Retrieve all transactions with pagination.     |
| `FindAllTransactionByCardNumber` | Retrieve transactions filtered by card number. |
| `FindByIdTransaction`            | Get transaction details by ID.                 |
| `FindByCardNumberTransaction`    | Get transaction details by card number.        |
| `FindByActive`                   | Get all active (non-deleted) transactions.     |
| `FindByTrashed`                  | Get all soft-deleted transactions.             |


#### 📊 Statistics Service Methods

| Method                                | Description                                                |
| ------------------------------------- | ---------------------------------------------------------- |
| `FindMonthlyTransactionStatusSuccess` | Monthly statistics of successful transactions (all cards). |
| `FindYearlyTransactionStatusSuccess`  | Yearly statistics of successful transactions (all cards).  |
| `FindMonthlyTransactionStatusFailed`  | Monthly statistics of failed transactions (all cards).     |
| `FindYearlyTransactionStatusFailed`   | Yearly statistics of failed transactions (all cards).      |
| `FindMonthlyTransactionStatusSuccessByCardNumber` | Monthly successful transactions by card number. |
| `FindYearlyTransactionStatusSuccessByCardNumber`  | Yearly successful transactions by card number.  |
| `FindMonthlyTransactionStatusFailedByCardNumber`  | Monthly failed transactions by card number.     |
| `FindYearlyTransactionStatusFailedByCardNumber`   | Yearly failed transactions by card number.      |
| `FindMonthlyTransactionMethods`             | Monthly transaction method statistics (all cards). |
| `FindYearlyTransactionMethods`              | Yearly transaction method statistics (all cards).  |
| `FindMonthlyTransactionMethodsByCardNumber` | Monthly method stats by card number.               |
| `FindYearlyTransactionMethodsByCardNumber`  | Yearly method stats by card number.                |
| `FindMonthlyTransactionAmounts`             | Monthly transaction amount stats (all cards). |
| `FindYearlyTransactionAmounts`              | Yearly transaction amount stats (all cards).  |
| `FindMonthlyTransactionAmountsByCardNumber` | Monthly amount stats by card number.          |
| `FindYearlyTransactionAmountsByCardNumber`  | Yearly amount stats by card number.           |

#### 📝 Command Service Methods

| Method                          | Description                                       |
| ------------------------------- | ------------------------------------------------- |
| `CreateTransaction`             | Create a new transaction record.                  |
| `UpdateTransaction`             | Update an existing transaction record.            |
| `TrashedTransaction`            | Soft-delete a transaction record.                 |
| `RestoreTransaction`            | Restore a soft-deleted transaction.               |
| `DeleteTransactionPermanent`    | Permanently delete a transaction.                 |
| `RestoreAllTransaction`         | Restore all soft-deleted transactions.            |
| `DeleteAllTransactionPermanent` | Permanently delete all soft-deleted transactions. |

### 📊 Monitoring & Observability

#### 📈 Prometheus Metrics

Each service exposes the following metrics:
  - **Request Counter**: Total number of requests, labeled by method and status.
  - **Request Duration Histogram**: Measures execution time for each method.

Metric examples:
  - `transaction_query_service_request_total`
  - `transaction_query_service_request_duration_seconds`
  - `transaction_command_service_request_total`
  - `transaction_command_service_request_duration_seconds`
  - `transaction_statistic_service_request_total`
  - `transaction_statistic_service_request_duration_seconds`
  - `transaction_statistic_by_card_service_request_total`
  - `transaction_statistic_by_card_service_request_duration_seconds`

#### 🔍 OpenTelemetry Tracing

All services are instrumented with their own tracer:
  - `transaction-command-service`
  - `transaction-query-service`
  - `transaction-statistic-service`
  - `transaction-statistic-by-card-service`


### 📬 Kafka Integration

TransactionService publishes events related to transaction lifecycle (e.g., creation) to Kafka topics. These are consumed by the email-service to send out notification emails.

| Kafka Topic                               | Purpose                                                | Consumer Group        |
| ----------------------------------------- | ------------------------------------------------------ | --------------------- |
| `email-service-topic-transaction-created` | Notify users when a transaction (t is successful | `email-service-group` |

## Status Lifecycle & Idempotency

Transaction mematuhi state machine kanonik (sumber kebenaran:
[`shared/domain/status/README.md`](../../shared/domain/status/README.md)):

| From | To (diizinkan) |
|---|---|
| `pending` | `authorized`, `voided`, `success`, `failed` |
| `authorized` | `captured`, `voided`, `failed` |
| `captured` | `refunded`, `failed` |
| `success` | `success`, `failed` (kompensasi sistem) |
| `failed` | `pending`, `success` |

- `Capture`/`Void`/`Refund`/`Authorize` memvalidasi transisi via
  `status.ValidateTransition` → **400** dengan pesan jelas saat status saat ini
  tidak mengizinkan operasi (mis. capture dari non-`authorized`, refund dari
  non-`captured`).
- Flow `Update` memvalidasi `status existing → success` sebelum delta saldo;
  `markTransactionAsFailed` adalah kompensasi sistem.
- Create idempotent: kirim `idempotency_key` (`≤ 64` char, format
  `transaction:create:merchant:<id>:<uuid>`) → replay mengembalikan response
  asli tanpa double settlement; email & cache invalidation tidak diulang saat
  replay. Key kosong = non idempotent.
