# Shared Error Definitions

This directory contains centralized error definitions for all domains in the project. 
It is structured by feature/module to maintain clarity and separation of concerns.

## Structure

Each module has its own folder with the following convention:
- `api_<module>_errors.go` — for HTTP/API-related errors.
- `repo_<module>_errors.go` — for repository/database-related errors.
- `service_<module>_errors.go` — for service/business logic-related errors.
- `README.md` — brief explanation per module.

This structure allows better error management and decouples concerns between layers (API, Service, Repository).

## Modules

- `auth_errors`
- `card_errors`
- `merchant_errors`
- `merchant_document_errors`
- `refresh_token_errors`
- `role_errors`
- `saldo_errors`
- `topup_errors`
- `transaction_errors`
- `transfer_errors`
- `user_errors`
- `user_role_errors`
- `withdraw_errors`
