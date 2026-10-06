# Mapper Layer

This package provides mappers to convert data between various layers such as:
- Database records
- API responses
- Proto (gRPC) messages
- Service domain logic

Each subdirectory is organized by conversion type and grouped by feature/domain (e.g., Auth, User, Topup).

## Structure

- `record/`: Maps from database models to internal record representations.
- `response/`: Maps from records to API/service responses.
  - `api/`: Used in controller/handler layer for HTTP/gRPC responses.
  - `service/`: Used within services to transform internal models into service outputs.
- `proto/`: Maps between proto (gRPC) messages and domain models.

All mappers implement interfaces for maintainability and ease of testing.
