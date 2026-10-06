# Payment Gateway Topup Service

## Overview

`TopupService` is a monolith responsible for handling balance top-ups. It supports monthly and yearly statistics based on method, amount, status, and card number. This service is built with strong observability using Prometheus, OpenTelemetry, and Zap Logger.


### 🔄 Service Architecture

TopupService is separated into several components for scalability and responsibility separation:
  - **topupCommandService**
    Handles write operations such as CreateTopup, UpdateTopup, TrashedTopup, RestoreTopup, and DeleteTopupPermanent.
  - **topupQueryService**
    Handles read operations such as FindAllTopup, FindByIdTopup, FindByCardNumberTopup, and top-up statistics.
  - **topupStatisticByCardService**
    Provides top-up statistic reports based on card number.
  - **topupStatisticService**
    Provides general top-up statistics (for all cards).

----

### 📌 Available RPC Methods

#### 📘 Query Service Methods

| Method                                      | Description                                              |
| ------------------------------------------- | -------------------------------------------------------- |
| `FindAllTopup`                              | Retrieve all top-up data with pagination.                |
| `FindAllTopupByCardNumber`                  | Retrieve all top-up data filtered by card number.        |
| `FindByIdTopup`                             | Get top-up details by ID.                                |
| `FindByCardNumberTopup`                     | Get top-up details by card number.                       |
| `FindByActive`                              | Get all active (non-deleted) top-up records.             |
| `FindByTrashed`                             | Get all soft-deleted top-up records.                     |
| `FindMonthlyTopupStatusSuccess`             | Monthly statistics of successful top-ups for all cards.  |
| `FindYearlyTopupStatusSuccess`              | Yearly statistics of successful top-ups for all cards.   |
| `FindMonthlyTopupStatusFailed`              | Monthly statistics of failed top-ups for all cards.      |
| `FindYearlyTopupStatusFailed`               | Yearly statistics of failed top-ups for all cards.       |
| `FindMonthlyTopupStatusSuccessByCardNumber` | Monthly statistics of successful top-ups by card number. |
| `FindYearlyTopupStatusSuccessByCardNumber`  | Yearly statistics of successful top-ups by card number.  |
| `FindMonthlyTopupStatusFailedByCardNumber`  | Monthly statistics of failed top-ups by card number.     |
| `FindYearlyTopupStatusFailedByCardNumber`   | Yearly statistics of failed top-ups by card number.      |
| `FindMonthlyTopupMethods`                   | Monthly statistics of top-up methods for all cards.      |
| `FindYearlyTopupMethods`                    | Yearly statistics of top-up methods for all cards.       |
| `FindMonthlyTopupAmounts`                   | Monthly statistics of top-up amounts for all cards.      |
| `FindYearlyTopupAmounts`                    | Yearly statistics of top-up amounts for all cards.       |
| `FindMonthlyTopupMethodsByCardNumber`       | Monthly statistics of top-up methods by card number.     |
| `FindYearlyTopupMethodsByCardNumber`        | Yearly statistics of top-up methods by card number.      |
| `FindMonthlyTopupAmountsByCardNumber`       | Monthly statistics of top-up amounts by card number.     |
| `FindYearlyTopupAmountsByCardNumber`        | Yearly statistics of top-up amounts by card number.      |

#### ✏️ Command Service Methods

| Method                    | Description                                         |
| ------------------------- | --------------------------------------------------- |
| `CreateTopup`             | Create a new top-up record.                         |
| `UpdateTopup`             | Update an existing top-up record.                   |
| `TrashedTopup`            | Soft-delete a top-up record.                        |
| `RestoreTopup`            | Restore a soft-deleted top-up record.               |
| `DeleteTopupPermanent`    | Permanently delete a top-up record.                 |
| `RestoreAllTopup`         | Restore all soft-deleted top-up records.            |
| `DeleteAllTopupPermanent` | Permanently delete all soft-deleted top-up records. |



### 📊 Monitoring & Observability

#### 📈 Prometheus Metrics

Each RPC method includes:
  - **Request Counter**: Total number of requests, labeled by method and status.
  - **Request Duration Histogram**: Measures execution time for each method.

Metric examples:
  - `topup_query_service_request_total`
  - `topup_query_service_request_duration_seconds`
  - `topup_command_service_request_total`
  - `topup_command_service_request_duration_seconds`
  - `topup_statistic_service_request_total`
  - `topup_statistic_service_request_duration_seconds`
  - `topup_statistic_by_card_service_request_total`
  - `topup_statistic_by_card_service_request_duration_seconds`


#### 🔍 OpenTelemetry Tracing

Each service is instrumented with a dedicated tracer:
  - `topup-command-service`
  - `topup-query-service`
  - `topup-statistic-service`
  - `topup-statistic-by-card-service`

----


### 📬 Kafka Integration

The TopupService integrates with Kafka to publish domain events related to top-up transactions. One such event is the top-up creation, which triggers downstream processes like email notifications.

These messages are published to specific Kafka topics and consumed by the email-service, which is responsible for sending out transactional emails such as successful top-up confirmations.

| Kafka Topic                                            | Purpose                                                               | Consumer Group        |
| ------------------------------------------------------ | --------------------------------------------------------------------- | --------------------- |
| `email-service-topic-topup-created`                 | Notify user when a new topup is successfully created               | `email-service-group` |

## Status Lifecycle & Idempotency

Topup mematuhi state machine kanonik (sumber kebenaran:
[`shared/domain/status/README.md`](../../shared/domain/status/README.md)):

| From | To (diizinkan) |
|---|---|
| `pending` | `success`, `failed` |
| `failed` | `pending`, `success` |
| `success` | `success`, `failed` (kompensasi sistem) |

- Transisi invalid ditolak **400** sebelum delta saldo diterapkan (mis.
  `success → pending` tidak pernah diizinkan).
- Flow `UpdateTopup` memvalidasi `status existing → success` via
  `status.ValidateTransition`; `markTopupAsFailed` adalah kompensasi sistem.
- Create idempotent: kirim `idempotency_key` (`≤ 64` char, format
  `topup:create:user:<id>:<uuid>`) → replay mengembalikan response asli tanpa
  efek ganda (balancing + card metadata tidak diulang). Key kosong = non
  idempotent.
