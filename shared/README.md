# Shared Package

The `shared/` directory contains reusable components, domain models, and utilities shared across all services in the system.

This package helps maintain consistency, reduce duplication, and enforce domain contracts between monolith.

## Structure Overview

### 1. `domain/`
Defines all domain-related types:
- `record/`: Core domain data structures (used internally).
- `requests/`: API request models (with validation support).
- `response/`: API response models (including success and error formats).

### 2. `errors/`
Centralized error definitions grouped by domain and layer:
- `api_<module>_errors.go`: For handler/controller-level errors.
- `repo_<module>_errors.go`: For repository/database-level errors.
- `service_<module>_errors.go`: For service logic-level errors.

Each subdirectory contains a `README.md` describing its error strategy.

### 3. `mapper/`
Provides mapping utilities for transforming data between layers:
- `record/`: DB models → domain records.
- `response/api/`: Domain records → API response.
- `response/service/`: Domain records → internal service response.
- `proto/`: Domain ↔ gRPC proto messages.

### 4. `grpc_error.go`
Contains common gRPC error translation and handling logic.

## Purpose

This shared module is designed to:
- Promote DRY (Don't Repeat Yourself) principles
- Standardize data transformation and error handling
- Serve as a foundation layer for services like `auth`, `user`, `transaction`, etc.

All changes in this package should be made with caution and proper review, as it affects multiple parts of the system.
