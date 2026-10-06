# Payment Gateway User Service

## Overview

`UserService` is a monolith responsible for **user account management** including creation, updates, soft deletion, restoration, and permanent deletion. It also supports user queries based on active or trashed status.

Communication is done via **gRPC**, and observability is built-in through **Prometheus metrics**, **OpenTelemetry tracing**, and structured logging using **Zap**.

---


### 🔄 Service Architecture

The `UserService` is split into two primary services for separation of concerns and scalability:

- **userCommandService**
  Handles all **write** operations such as `Create`, `Update`, `TrashedUser`, `RestoreUser`, and `DeleteUserPermanent`.

- **userQueryService**
  Handles all **read** operations such as `FindAll`, `FindById`, `FindByActive`, and `FindByTrashed`.

All services integrate with:
- **Prometheus metrics** to track request counts and durations for observability.
- **OpenTelemetry tracing** to provide distributed tracing for requests, helping with debugging and performance tuning.
- **Structured logging** with Zap for detailed error and operational logs.


### 📌 Available RPC Methods

#### 📘 Query Service Methods

| Method          | Description                                  |
| --------------- | -------------------------------------------- |
| `FindAll`       | Fetches all users with pagination.           |
| `FindById`      | Retrieves a user by ID.                      |
| `FindByActive`  | Retrieves users that are currently active.   |
| `FindByTrashed` | Retrieves users that have been soft-deleted. |

#### ✏️ Command Service Methods

| Method                   | Description                                 |
| ------------------------ | ------------------------------------------- |
| `Create`                 | Creates a new user.                         |
| `Update`                 | Updates user information.                   |
| `TrashedUser`            | Soft-deletes a user.                        |
| `RestoreUser`            | Restores a soft-deleted user.               |
| `DeleteUserPermanent`    | Permanently deletes a user from the system. |
| `RestoreAllUser`         | Restores all soft-deleted users.            |
| `DeleteAllUserPermanent` | Permanently deletes all soft-deleted users. |


---

### 📊 Monitoring & Observability

#### Prometheus Metrics

Each RPC method is instrumented with:

- **Request Counter**: total number of requests by method and status.
- **Request Duration Histogram**: duration of each RPC method execution.

**Metric examples**:
- `user_query_service_request_total`
- `user_query_service_request_duration_seconds`
- `user_command_service_requests_total`
- `user_command_service_request_duration_seconds`

---

#### 🔍 Tracing with OpenTelemetry

Each service is assigned a tracer for distributed tracing:
- `user-command-service`
- `user-query-service`