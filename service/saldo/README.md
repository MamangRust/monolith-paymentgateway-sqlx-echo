# Payment Gateway Saldo Service

## Overview

`SaldoService` is a monolith responsible for **balance management**, including balance retrieval, statistics, and soft-delete operations. It also listens to events such as **saldo creation** triggered by the Card Service. Communication is done via **gRPC**, and event consumption is done via **Kafka** from the `card-service`.

---


### 🔄 Service Architecture

`SaldoService` is divided into several specialized services:

- **Command Service**: Manages saldo creation, update, trashing, and deletion.
- **Query Service**: Handles saldo retrieval and lookup by filters or metadata.
- **Statistics Service**: Provides historical and aggregated balance statistics.

All services use:

- **Kafka** to consume saldo creation events from Card Service.
- **Prometheus** metrics for request counting and duration tracking.
- **OpenTelemetry Tracing** for distributed system observability.
- **Zap Logger** for structured and contextual logging.

---


### 📌 Available RPC Methods

### 📘 Query Service Methods

| Method             | Description                                          |
| ------------------ | ---------------------------------------------------- |
| `FindAllSaldo`     | Retrieves all saldo records with pagination.         |
| `FindByIdSaldo`    | Retrieves saldo by its unique ID.                    |
| `FindByCardNumber` | Retrieves saldo data by card number.                 |
| `FindByActive`     | Retrieves saldo records that are currently active.   |
| `FindByTrashed`    | Retrieves saldo records that have been soft-deleted. |


### ✏️ Command Service Methods
| Method                    | Description                                    |
| ------------------------- | ---------------------------------------------- |
| `CreateSaldo`             | Creates a new saldo record.                    |
| `UpdateSaldo`             | Updates an existing saldo record.              |
| `TrashedSaldo`            | Soft-deletes a saldo record.                   |
| `RestoreSaldo`            | Restores a previously soft-deleted saldo.      |
| `DeleteSaldoPermanent`    | Permanently deletes a saldo record.            |
| `RestoreAllSaldo`         | Restores all soft-deleted saldo records.       |
| `DeleteAllSaldoPermanent` | Permanently deletes all trashed saldo records. |


### 📊 Statistics Service
| Method                         | Description                                   |
| ------------------------------ | --------------------------------------------- |
| `FindMonthlyTotalSaldoBalance` | Returns total saldo balance grouped by month. |
| `FindYearTotalSaldoBalance`    | Returns total saldo balance for a year.       |
| `FindMonthlySaldoBalances`     | Returns saldo data for each month of a year.  |
| `FindYearlySaldoBalances`      | Returns saldo data grouped by year.           |



### 📨 Kafka Integration

`SaldoService` listens for saldo creation events from the Card Service. When a card is created, the Card Service sends an event to the following Kafka topic:

| Kafka Topic | Purpose | Consumer Group |
|-------------|---------|----------------|
| `saldo-service-topic-create-saldo` | Triggers saldo creation when a new card is issued | `saldo-service-group` |

---

### 📊 Monitoring & Observability

#### Prometheus Metrics

Each RPC method is instrumented with:

- **Request Counter**: Tracks total number of requests by method and status.
- **Request Duration Histogram**: Tracks execution duration for each RPC method.

**Sample Metrics**:
- `saldo_command_service_request_total`
- `saldo_command_service_request_duration_seconds`
- `saldo_query_service_request_total`
- `saldo_query_service_request_duration_seconds`
- `saldo_statistics_service_request_total`
- `saldo_statistics_service_request_duration_seconds`

---

#### 🔍 Tracing with OpenTelemetry
Each internal service is assigned a tracer:

- `saldo-command-service`
- `saldo-query-service`
- `saldo-statistics-service`
---