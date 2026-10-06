# 📦 Package `merchantdocument_errors`

**Source Path:** `shared/errors/merchant_document_errors`

## 🏷️ Variables

**Var:**

ErrApiBindCreateMerchantDocument is returned when binding the create request payload fails.

```go
var ErrApiBindCreateMerchantDocument = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "bind failed: invalid create merchant document request", http.StatusBadRequest)
}
```

**Var:**

ErrApiBindUpdateMerchantDocument is returned when binding the update request payload fails.

```go
var ErrApiBindUpdateMerchantDocument = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "bind failed: invalid update merchant document request", http.StatusBadRequest)
}
```

**Var:**

ErrApiBindUpdateMerchantDocumentStatus is returned when binding the update status request payload fails.

```go
var ErrApiBindUpdateMerchantDocumentStatus = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "bind failed: invalid update merchant document status request", http.StatusBadRequest)
}
```

```go
var ErrApiFailedCreateMerchantDocument = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to create merchant document", http.StatusInternalServerError)
}
```

```go
var ErrApiFailedDeleteAllMerchantDocumentsPermanent = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to permanently delete all merchant documents", http.StatusInternalServerError)
}
```

```go
var ErrApiFailedDeleteMerchantDocumentPermanent = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to permanently delete merchant document", http.StatusInternalServerError)
}
```

```go
var ErrApiFailedFindAllActiveMerchantDocuments = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch all active merchant documents", http.StatusInternalServerError)
}
```

```go
var ErrApiFailedFindAllMerchantDocuments = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch all merchant documents", http.StatusInternalServerError)
}
```

```go
var ErrApiFailedFindAllTrashedMerchantDocuments = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch all trashed merchant documents", http.StatusInternalServerError)
}
```

```go
var ErrApiFailedFindByIdMerchantDocument = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch merchant document by ID", http.StatusInternalServerError)
}
```

```go
var ErrApiFailedRestoreAllMerchantDocuments = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to restore all merchant documents", http.StatusInternalServerError)
}
```

```go
var ErrApiFailedRestoreMerchantDocument = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to restore merchant document", http.StatusInternalServerError)
}
```

```go
var ErrApiFailedTrashMerchantDocument = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to trash merchant document", http.StatusInternalServerError)
}
```

```go
var ErrApiFailedUpdateMerchantDocument = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to update merchant document", http.StatusInternalServerError)
}
```

```go
var ErrApiFailedUpdateMerchantDocumentStatus = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to update merchant document status", http.StatusInternalServerError)
}
```

```go
var ErrApiInvalidMerchantDocumentID = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Invalid merchant document ID", http.StatusBadRequest)
}
```

```go
var ErrApiValidateCreateMerchantDocument = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "validation failed: invalid create merchant document request", http.StatusBadRequest)
}
```

```go
var ErrApiValidateUpdateMerchantDocument = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "validation failed: invalid update merchant document request", http.StatusBadRequest)
}
```

```go
var ErrApiValidateUpdateMerchantDocumentStatus = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "validation failed: invalid update merchant document status request", http.StatusBadRequest)
}
```

**Var:**

ErrCreateMerchantDocumentFailed is returned when failing to create a new merchant document.

```go
var ErrCreateMerchantDocumentFailed = errors.New("failed to create merchant document")
```

**Var:**

ErrDeleteAllMerchantDocumentsPermanentFailed is returned when failing to permanently delete all merchant documents.

```go
var ErrDeleteAllMerchantDocumentsPermanentFailed = errors.New("failed to permanently delete all merchant documents")
```

**Var:**

ErrDeleteMerchantDocumentPermanentFailed is returned when failing to permanently delete a merchant document.

```go
var ErrDeleteMerchantDocumentPermanentFailed = errors.New("failed to permanently delete merchant document")
```

**Var:**

ErrFailedCreateMerchantDocument is returned when failing to create a new merchant document.

```go
var ErrFailedCreateMerchantDocument = response.NewErrorResponse("Failed to create Merchant Document", http.StatusInternalServerError)
```

**Var:**

ErrFailedDeleteAllMerchantDocuments is returned when failing to permanently delete all merchant documents.

```go
var ErrFailedDeleteAllMerchantDocuments = response.NewErrorResponse("Failed to delete all Merchant Documents permanently", http.StatusInternalServerError)
```

**Var:**

ErrFailedDeleteMerchantDocument is returned when failing to permanently delete a merchant document.

```go
var ErrFailedDeleteMerchantDocument = response.NewErrorResponse("Failed to delete Merchant Document permanently", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindActiveMerchantDocuments is returned when failing to fetch active merchant documents.

```go
var ErrFailedFindActiveMerchantDocuments = response.NewErrorResponse("Failed to fetch active Merchant Documents", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindAllMerchantDocuments is returned when failing to fetch all merchant documents.

```go
var ErrFailedFindAllMerchantDocuments = response.NewErrorResponse("Failed to fetch Merchant Documents", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindMerchantDocumentById is returned when failing to find a merchant document by ID.

```go
var ErrFailedFindMerchantDocumentById = response.NewErrorResponse("Failed to find Merchant Document by ID", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindTrashedMerchantDocuments is returned when failing to fetch trashed merchant documents.

```go
var ErrFailedFindTrashedMerchantDocuments = response.NewErrorResponse("Failed to fetch trashed Merchant Documents", http.StatusInternalServerError)
```

**Var:**

ErrFailedRestoreAllMerchantDocuments is returned when failing to restore all trashed merchant documents.

```go
var ErrFailedRestoreAllMerchantDocuments = response.NewErrorResponse("Failed to restore all Merchant Documents", http.StatusInternalServerError)
```

**Var:**

ErrFailedRestoreMerchantDocument is returned when failing to restore a trashed merchant document.

```go
var ErrFailedRestoreMerchantDocument = response.NewErrorResponse("Failed to restore Merchant Document", http.StatusInternalServerError)
```

**Var:**

ErrFailedTrashMerchantDocument is returned when failing to move a merchant document to trash.

```go
var ErrFailedTrashMerchantDocument = response.NewErrorResponse("Failed to trash Merchant Document", http.StatusInternalServerError)
```

**Var:**

ErrFailedUpdateMerchantDocument is returned when failing to update an existing merchant document.

```go
var ErrFailedUpdateMerchantDocument = response.NewErrorResponse("Failed to update Merchant Document", http.StatusInternalServerError)
```

**Var:**

ErrFindActiveMerchantDocumentsFailed is returned when failing to retrieve active merchant documents.

```go
var ErrFindActiveMerchantDocumentsFailed = errors.New("failed to find active merchant documents")
```

**Var:**

ErrFindAllMerchantDocumentsFailed is returned when failing to retrieve all merchant documents.

```go
var ErrFindAllMerchantDocumentsFailed = errors.New("failed to find all merchant documents")
```

**Var:**

ErrFindMerchantDocumentByIdFailed is returned when failing to retrieve a merchant document by ID.

```go
var ErrFindMerchantDocumentByIdFailed = errors.New("failed to find merchant document by ID")
```

**Var:**

ErrFindTrashedMerchantDocumentsFailed is returned when failing to retrieve trashed merchant documents.

```go
var ErrFindTrashedMerchantDocumentsFailed = errors.New("failed to find trashed merchant documents")
```

**Var:**

ErrMerchantDocumentNotFoundRes is returned when the requested merchant document cannot be found.

```go
var ErrMerchantDocumentNotFoundRes = response.NewErrorResponse("Merchant Document not found", http.StatusNotFound)
```

**Var:**

ErrRestoreAllMerchantDocumentsFailed is returned when failing to restore all trashed merchant documents.

```go
var ErrRestoreAllMerchantDocumentsFailed = errors.New("failed to restore all merchant documents")
```

**Var:**

ErrRestoreMerchantDocumentFailed is returned when failing to restore a trashed merchant document.

```go
var ErrRestoreMerchantDocumentFailed = errors.New("failed to restore merchant document")
```

**Var:**

ErrTrashedMerchantDocumentFailed is returned when failing to move a merchant document to trash.

```go
var ErrTrashedMerchantDocumentFailed = errors.New("failed to soft-delete (trash) merchant document")
```

**Var:**

ErrUpdateMerchantDocumentFailed is returned when failing to update an existing merchant document.

```go
var ErrUpdateMerchantDocumentFailed = errors.New("failed to update merchant document")
```

**Var:**

ErrUpdateMerchantDocumentStatusFailed is returned when failing to update the status of a merchant document.

```go
var ErrUpdateMerchantDocumentStatusFailed = errors.New("failed to update merchant document status")
```

