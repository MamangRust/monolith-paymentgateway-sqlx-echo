# Payment Gateway Transfer Service

## Overview


`TransferService` is a monolith responsible for managing card-to-card transfer transactions. It supports monthly and yearly statistics based on method, amount, status, and card numbers (sender/receiver). This service is built with strong observability using Prometheus, OpenTelemetry, and Zap Logger.


### 🔄 Service Architecture

TransferService is divided into multiple components to ensure scalability and clear separation of concerns:
  - **transferCommandService**
    Handles write operations such as CreateTransfer, UpdateTransfer, TrashedTransfer, RestoreTransfer, and DeleteTransferPermanent.
  - **transferQueryService**
    Handles read operations such as FindAllTransfer, FindByIdTransfer, FindTransferByTransferFrom, FindTransferByTransferTo, and soft-deletion filters.
  - **transferStatisticService**
    Provides general transfer statistics for all cards.
  - **transferStatisticByCardService**
    Provides transfer statistics grouped by sender and receiver card numbers.

----

### 📌 Available RPC Methods

#### 📘 Query Service Methods
| Method                       | Description                                               |
| ---------------------------- | --------------------------------------------------------- |
| `FindAllTransfer`            | Retrieve all transfer records with pagination.            |
| `FindByIdTransfer`           | Retrieve a specific transfer by ID.                       |
| `FindTransferByTransferFrom` | Retrieve transfers initiated from a specific card number. |
| `FindTransferByTransferTo`   | Retrieve transfers received by a specific card number.    |
| `FindByActiveTransfer`       | Get all active (non-deleted) transfers.                   |
| `FindByTrashedTransfer`      | Get all soft-deleted transfers.                           |


#### 📊 Statistics Service Methods
| Method                                           | Description                                           |
| ------------------------------------------------ | ----------------------------------------------------- |
| `FindMonthlyTransferStatusSuccess`               | Monthly success transfer stats for all cards.         |
| `FindYearlyTransferStatusSuccess`                | Yearly success transfer stats for all cards.          |
| `FindMonthlyTransferStatusFailed`                | Monthly failed transfer stats for all cards.          |
| `FindYearlyTransferStatusFailed`                 | Yearly failed transfer stats for all cards.           |
| `FindMonthlyTransferStatusSuccessByCardNumber`   | Monthly success stats by sender/receiver card number. |
| `FindYearlyTransferStatusSuccessByCardNumber`    | Yearly success stats by sender/receiver card number.  |
| `FindMonthlyTransferStatusFailedByCardNumber`    | Monthly failed stats by sender/receiver card number.  |
| `FindYearlyTransferStatusFailedByCardNumber`     | Yearly failed stats by sender/receiver card number.   |
| `FindMonthlyTransferAmounts`                     | Monthly aggregated transfer amounts (all cards).      |
| `FindYearlyTransferAmounts`                      | Yearly aggregated transfer amounts (all cards).       |
| `FindMonthlyTransferAmountsBySenderCardNumber`   | Monthly amount stats by sender card.                  |
| `FindMonthlyTransferAmountsByReceiverCardNumber` | Monthly amount stats by receiver card.                |
| `FindYearlyTransferAmountsBySenderCardNumber`    | Yearly amount stats by sender card.                   |
| `FindYearlyTransferAmountsByReceiverCardNumber`  | Yearly amount stats by receiver card.                 |


#### 📝 Command Service Methods
| Method                       | Description                                    |
| ---------------------------- | ---------------------------------------------- |
| `CreateTransfer`             | Create a new transfer record.                  |
| `UpdateTransfer`             | Update an existing transfer record.            |
| `TrashedTransfer`            | Soft-delete a transfer record.                 |
| `RestoreTransfer`            | Restore a soft-deleted transfer.               |
| `DeleteTransferPermanent`    | Permanently delete a transfer.                 |
| `RestoreAllTransfer`         | Restore all soft-deleted transfers.            |
| `DeleteAllTransferPermanent` | Permanently delete all soft-deleted transfers. |


----

### 📊 Monitoring & Observability

#### 📈 Prometheus Metrics

Each service exposes the following metrics:
  - **Request Counter**: Counts all requests per method and status.
  - **Request Duration** Histogram: Measures the latency for each method.

Metric examples:
  - `transfer_query_service_request_total`
  - `transfer_query_service_request_duration_seconds`
  - `transfer_command_service_request_total`
  - `transfer_command_service_request_duration_seconds`
  - `transfer_statistic_service_request_total`
  - `transfer_statistic_service_request_duration_seconds`
  - `transfer_statistic_by_card_service_request_total`
  - `transfer_statistic_by_card_service_request_duration_seconds`


#### 🔍 OpenTelemetry Tracing
All services are instrumented with their own tracer:
  - `transfer-command-service`
  - `transfer-query-service`
  - `transfer-statistic-service`
  - `transfer-statistic-by-card-service`


----

### 📬 Kafka Integration

`TransferService` publishes events related to transfer lifecycle (e.g., after successful transfer) to Kafka topics. These events are consumed by other services like the email-service to send notifications.

| Kafka Topic                            | Purpose                                               | Consumer Group        |
| -------------------------------------- | ----------------------------------------------------- | --------------------- |
| `email-service-topic-transfer-created` | Notify users when a transfer is successfully created. | `email-service-group` |

## Status Lifecycle & Idempotency

Transfer mematuhi state machine kanonik (sumber kebenaran:
[`shared/domain/status/README.md`](../../shared/domain/status/README.md)):

| From | To (diizinkan) |
|---|---|
| `pending` | `success`, `failed` |
| `failed` | `pending`, `success` |
| `success` | `success`, `failed` (kompensasi sistem) |

- Transisi invalid ditolak **400** sebelum settlement delta (sender/receiver)
  dipindahkan (mis. `success → pending` tidak pernah diizinkan).
- Flow `UpdateTransaction` memvalidasi `status existing → success` via
  `status.ValidateTransition`; `markTransferAsFailed` adalah kompensasi
  sistem.
- Create idempotent: kirim `idempotency_key` (`≤ 64` char, format
  `transfer:create:user:<id>:<uuid>`) → replay mengembalikan response asli
  tanpa double settlement; email tidak dikirim ulang saat replay. Key kosong =
  non idempotent.