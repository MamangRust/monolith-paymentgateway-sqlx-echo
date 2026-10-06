# Payment Gateway Merchant Service

## Overview

The **Merchant Service** is a comprehensive monolith designed to manage merchant accounts, documents, transactions, and analytics within a payment processing ecosystem. It provides both query and command operations through a CQRS (Command Query Responsibility Segregation) pattern and integrates seamlessly with the Auth Service for authentication and authorization.


### 🔄 Service Architecture

The MerchantService is organized into multiple specialized service interfaces to separate responsibilities and improve scalability and maintainability:

  - **MerchantCommandService**
    Handles merchant lifecycle operations such as creation, update, soft delete, restore, and permanent deletion.
    Publishes domain events (e.g., email-service-topic-merchant-created, email-service-topic-merchant-update-status) to Kafka for asynchronous downstream processing.
    Also produces messages to Kafka topics consumed by the email-service-group to notify users when a merchant is approved or rejected.

  - **MerchantQueryService**
    Focuses on data retrieval operations for merchant entities, including paginated listings, filtering (e.g., by status), and fetching individual merchant details. Optimized for performance and scalability.

  - **MerchantDocumentCommandService**
    Manages merchant document operations such as uploading, updating, soft deleting, restoring, and permanently deleting documents.
    Publishes domain events (e.g., email-service-topic-merchant-documents-created, email-service-topic-merchant-documents-update-status) to Kafka, including events consumed by the email-service-group to notify users of document approval status.

  - **MerchantDocumentQueryService**
    Handles retrieval of merchant documents, supports filtering by status (active/trashed), and provides document metadata such as file type, upload date, and current status.

  - **MerchantTransactionService**
    Provides access to merchant-related transaction history and supports querying, filtering, and analytics related to transaction flow and grouping.

  - **MerchantStatisticsService**
    Composed of multiple services responsible for delivering analytics and statistics:
      - **MerchantStatisService**: General analytics across all merchants.
      - **MerchantStatisByApiKeyService**: Metrics and insights scoped by API key.
      - **MerchantStatisByMerchantService**: Analytics scoped per individual merchant entity.

All services are tightly integrated with:
  - **Kafka** — Used for producing and consuming domain events such as email-service-topic-merchant-created, email-service-topic-merchant-update-status, email-service-topic-merchant-documents-created, email-service-topic-merchant-documents-update-status, and others.
    Key Kafka topics are consumed by systems like email-service-group to drive asynchronous workflows, including email notifications for merchant and document approvals.

  - **Prometheus metrics** — Integrated across all gRPC endpoints to track request counts, durations, and error rates, enabling system observability.

  - **OpenTelemetry tracing** — Provides end-to-end distributed tracing across services and infrastructure, allowing fine-grained visibility for debugging and performance optimization.

  - **Structured logging with Zap** — Consistent and structured log outputs across all services to support effective monitoring and operational insights.


## 📌 Available RPC Methods

### MerchantService RPCs

#### Query Operations
| Method | Description | Request Type | Response Type |
|--------|-------------|--------------|---------------|
| `FindAllMerchant` | Retrieve all merchants with pagination | `FindAllMerchantRequest` | `ApiResponsePaginationMerchant` |
| `FindByIdMerchant` | Find merchant by ID | `FindByIdMerchantRequest` | `ApiResponseMerchant` |
| `FindByApiKey` | Find merchant by API key | `FindByApiKeyRequest` | `ApiResponseMerchant` |
| `FindByMerchantUserId` | Find merchants by user ID | `FindByMerchantUserIdRequest` | `ApiResponsesMerchant` |
| `FindByActive` | Get active merchants only | `FindAllMerchantRequest` | `ApiResponsePaginationMerchantDeleteAt` |
| `FindByTrashed` | Get soft-deleted merchants | `FindAllMerchantRequest` | `ApiResponsePaginationMerchantDeleteAt` |

#### Command Operations
| Method | Description | Request Type | Response Type |
|--------|-------------|--------------|---------------|
| `CreateMerchant` | Create new merchant account | `CreateMerchantRequest` | `ApiResponseMerchant` |
| `UpdateMerchant` | Update merchant information | `UpdateMerchantRequest` | `ApiResponseMerchant` |
| `UpdateMerchantStatus` | Update merchant status | `UpdateMerchantStatusRequest` | `ApiResponseMerchant` |
| `TrashedMerchant` | Soft delete merchant | `FindByIdMerchantRequest` | `ApiResponseMerchant` |
| `RestoreMerchant` | Restore soft-deleted merchant | `FindByIdMerchantRequest` | `ApiResponseMerchant` |
| `DeleteMerchantPermanent` | Permanently delete merchant | `FindByIdMerchantRequest` | `ApiResponseMerchantDelete` |
| `RestoreAllMerchant` | Restore all soft-deleted merchants | `google.protobuf.Empty` | `ApiResponseMerchantAll` |
| `DeleteAllMerchantPermanent` | Permanently delete all merchants | `google.protobuf.Empty` | `ApiResponseMerchantAll` |

#### Transaction Analytics
| Method | Description | Request Type | Response Type |
|--------|-------------|--------------|---------------|
| `FindAllTransactionMerchant` | Get all merchant transactions | `FindAllMerchantRequest` | `ApiResponsePaginationMerchantTransaction` |
| `FindAllTransactionByMerchant` | Get transactions for specific merchant | `FindAllMerchantTransaction` | `ApiResponsePaginationMerchantTransaction` |
| `FindAllTransactionByApikey` | Get transactions by API key | `FindAllMerchantApikey` | `ApiResponsePaginationMerchantTransaction` |

#### Payment Method Analytics
| Method | Description | Request Type | Response Type |
|--------|-------------|--------------|---------------|
| `FindMonthlyPaymentMethodsMerchant` | Monthly payment method statistics | `FindYearMerchant` | `ApiResponseMerchantMonthlyPaymentMethod` |
| `FindYearlyPaymentMethodMerchant` | Yearly payment method statistics | `FindYearMerchant` | `ApiResponseMerchantYearlyPaymentMethod` |
| `FindMonthlyPaymentMethodByMerchants` | Monthly stats by merchant ID | `FindYearMerchantById` | `ApiResponseMerchantMonthlyPaymentMethod` |
| `FindYearlyPaymentMethodByMerchants` | Yearly stats by merchant ID | `FindYearMerchantById` | `ApiResponseMerchantYearlyPaymentMethod` |
| `FindMonthlyPaymentMethodByApikey` | Monthly stats by API key | `FindYearMerchantByApikey` | `ApiResponseMerchantMonthlyPaymentMethod` |
| `FindYearlyPaymentMethodByApikey` | Yearly stats by API key | `FindYearMerchantByApikey` | `ApiResponseMerchantYearlyPaymentMethod` |

#### Amount Analytics
| Method | Description | Request Type | Response Type |
|--------|-------------|--------------|---------------|
| `FindMonthlyAmountMerchant` | Monthly amount statistics | `FindYearMerchant` | `ApiResponseMerchantMonthlyAmount` |
| `FindYearlyAmountMerchant` | Yearly amount statistics | `FindYearMerchant` | `ApiResponseMerchantYearlyAmount` |
| `FindMonthlyTotalAmountMerchant` | Monthly total amount | `FindYearMerchant` | `ApiResponseMerchantMonthlyTotalAmount` |
| `FindYearlyTotalAmountMerchant` | Yearly total amount | `FindYearMerchant` | `ApiResponseMerchantYearlyTotalAmount` |

### MerchantDocumentService RPCs

#### Query Operations
| Method | Description | Request Type | Response Type |
|--------|-------------|--------------|---------------|
| `FindAll` | Get all merchant documents | `FindAllMerchantDocumentsRequest` | `ApiResponsePaginationMerchantDocument` |
| `FindAllActive` | Get active documents only | `FindAllMerchantDocumentsRequest` | `ApiResponsePaginationMerchantDocument` |
| `FindAllTrashed` | Get soft-deleted documents | `FindAllMerchantDocumentsRequest` | `ApiResponsePaginationMerchantDocumentAt` |
| `FindById` | Find document by ID | `FindMerchantDocumentByIdRequest` | `ApiResponseMerchantDocument` |

#### Command Operations
| Method | Description | Request Type | Response Type |
|--------|-------------|--------------|---------------|
| `Create` | Upload new merchant document | `CreateMerchantDocumentRequest` | `ApiResponseMerchantDocument` |
| `Update` | Update document information | `UpdateMerchantDocumentRequest` | `ApiResponseMerchantDocument` |
| `UpdateStatus` | Update document status | `UpdateMerchantDocumentStatusRequest` | `ApiResponseMerchantDocument` |
| `Trashed` | Soft delete document | `TrashedMerchantDocumentRequest` | `ApiResponseMerchantDocument` |
| `Restore` | Restore soft-deleted document | `RestoreMerchantDocumentRequest` | `ApiResponseMerchantDocument` |
| `DeletePermanent` | Permanently delete document | `DeleteMerchantDocumentPermanentRequest` | `ApiResponseMerchantDocumentDelete` |
| `RestoreAll` | Restore all soft-deleted documents | `google.protobuf.Empty` | `ApiResponseMerchantDocumentAll` |
| `DeleteAllPermanent` | Permanently delete all documents | `google.protobuf.Empty` | `ApiResponseMerchantDocumentAll` |


### 📊 Monitoring & Observability

#### Prometheus Metrics

Each service component exposes standardized metrics:

#### Request Metrics
- **Counter**: `{service_name}_requests_total` (labels: method, status)
- **Histogram**: `{service_name}_request_duration_seconds` (labels: method)

#### Service-Specific Metrics
- `merchant_command_service_requests_total`
- `merchant_command_service_request_duration_seconds`
- `merchant_query_service_requests_total`
- `merchant_query_service_request_duration_seconds`
- `merchant_document_command_request_count`
- `merchant_document_command_request_duration_seconds`
- `merchant_document_query_request_count`
- `merchant_document_query_request_duration_seconds`
- `merchant_transaction_service_requests_total`
- `merchant_transaction_service_request_duration_seconds`
- `merchant_statistic_service_request_total`
- `merchant_statistic_service_request_duration_seconds`
- `merchant_statis_by_apikey_service_requests_total`
- `merchant_statis_by_apikey_service_request_duration_seconds`
- `merchant_statis_by_merchant_service_requests_total`
- `merchant_statis_by_merchant_service_request_duration_seconds`

#### 🔍 Tracing with OpenTelemetry

Distributed tracing is implemented with dedicated tracers:
- `merchant-command-service`
- `merchant-query-service`
- `merchant-document-command-service`
- `merchant-document-query-service`
- `merchant-transaction-service`
- `merchant-statistic-service`
- `merchant-statis-by-apikey-service`
- `merchant-statis-by-merchant-service`

----

### 📬 Kafka Integration

The MerchantService uses Kafka to publish domain events related to merchant and document lifecycle operations. These messages are consumed by the email-service, which is responsible for sending out email notifications—such as merchant approval, document status updates, or creation confirmations.

| Kafka Topic                                            | Purpose                                                               | Consumer Group        |
| ------------------------------------------------------ | --------------------------------------------------------------------- | --------------------- |
| `email-service-topic-merchant-created`                 | Notify user when a new merchant is successfully created               | `email-service-group` |
| `email-service-topic-merchant-update-status`           | Notify user when merchant status is updated (e.g., approved)          | `email-service-group` |
| `email-service-topic-merchant-documents-created`       | Notify user when a merchant document is uploaded                      | `email-service-group` |
| `email-service-topic-merchant-documents-update-status` | Notify user when document status is updated (e.g., approved/rejected) | `email-service-group` |
