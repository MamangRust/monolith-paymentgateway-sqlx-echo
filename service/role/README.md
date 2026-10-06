# Payment Gateway Role Service

## Overview

`RoleService` is a monolith responsible for managing **user roles** and **role-related operations** within the system. It supports essential functionalities such as role creation, updating, soft deletion (trashed), restoration, and permanent deletion. The service provides both command and query RPC methods to handle data modifications and retrieval efficiently. Communication happens over **gRPC**.


### 🔄 Service Architecture

The RoleService separates command and query responsibilities into two dedicated service components for better scalability and maintainability:

- **roleCommandService:** Handles all write operations such as create, update, trash, restore, and delete.
- **roleQueryService:** Handles read operations including fetching roles by various criteria.

Both services integrate with:
- **Prometheus metrics** to monitor total request counts and request duration histograms for each RPC method.
- **OpenTelemetry tracing** to capture detailed traces for observability, performance tuning, and troubleshooting.
- **Structured logging** with Zap for capturing detailed logs and errors.

---

### 📌 Available RPC Methods

#### 📘 Query Service Methods

| Method          | Description                                 |
| --------------- | ------------------------------------------- |
| `FindAllRole`   | Retrieve all roles with pagination.         |
| `FindByIdRole`  | Get detailed information of a role by ID.   |
| `FindByActive`  | List all active roles (non-deleted).        |
| `FindByTrashed` | List all soft-deleted (trashed) roles.      |
| `FindByUserId`  | Retrieve roles assigned to a specific user. |

#### ✏️ Command Service Methods

| Method                   | Description                                |
| ------------------------ | ------------------------------------------ |
| `CreateRole`             | Create a new role.                         |
| `UpdateRole`             | Update an existing role.                   |
| `TrashedRole`            | Soft-delete (trash) a role.                |
| `RestoreRole`            | Restore a previously trashed role.         |
| `DeleteRolePermanent`    | Permanently delete a role from the system. |
| `RestoreAllRole`         | Restore all trashed roles.                 |
| `DeleteAllRolePermanent` | Permanently delete all trashed roles.      |
---

---

### 📊 Monitoring with Prometheus

Each RPC method is instrumented with:

- **Request Counter**: total number of requests by method and status.
- **Request Duration Histogram**: duration of each RPC method execution.

**Metric examples**:
- `role_query_service_request_total`
- `role_query_service_request_duration_seconds`
- `role_command_service_requests_total`
- `role_command_service_request_duration_seconds`

### 🔍 Tracing with OpenTelemetry

Each service is assigned a tracer for distributed tracing:
- `role-command-service`
- `role-query-service`
