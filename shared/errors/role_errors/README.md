# 📦 Package `role_errors`

**Source Path:** `./shared/errors/role_errors`

## 🏷️ Variables

**Var:**

ErrApiBindCreateRole returns an API error response for failed Role creation request binding.

```go
var ErrApiBindCreateRole = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "bind failed: invalid create Role request", http.StatusBadRequest)
}
```

**Var:**

ErrApiBindUpdateRole returns an API error response for failed Role update request binding.

```go
var ErrApiBindUpdateRole = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "bind failed: invalid update Role request", http.StatusBadRequest)
}
```

**Var:**

ErrApiFailedCreateRole returns an API error response when Role creation fails.

```go
var ErrApiFailedCreateRole = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "failed to create Role", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedDeleteAll returns an API error response when permanently deleting all Roles fails.

```go
var ErrApiFailedDeleteAll = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "failed to delete all Roles permanently", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedDeletePermanent returns an API error response when permanently deleting a Role fails.

```go
var ErrApiFailedDeletePermanent = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "failed to delete Role permanently", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindActive returns an API error response when fetching active Roles fails.

```go
var ErrApiFailedFindActive = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "failed to fetch active Roles", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindAll returns an API error response when fetching all Roles fails.

```go
var ErrApiFailedFindAll = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "failed to fetch Roles", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindTrashed returns an API error response when fetching trashed Roles fails.

```go
var ErrApiFailedFindTrashed = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "failed to fetch trashed Roles", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedRestoreAll returns an API error response when restoring all trashed Roles fails.

```go
var ErrApiFailedRestoreAll = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "failed to restore all Roles", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedRestoreRole returns an API error response when restoring a trashed Role fails.

```go
var ErrApiFailedRestoreRole = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "failed to restore Role", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedTrashedRole returns an API error response when soft-deleting (trashing) a Role fails.

```go
var ErrApiFailedTrashedRole = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "failed to move Role to trash", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedUpdateRole returns an API error response when Role update fails.

```go
var ErrApiFailedUpdateRole = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "failed to update Role", http.StatusInternalServerError)
}
```

**Var:**

ErrApiRoleInvalidId returns an API error response for an invalid Role ID.

```go
var ErrApiRoleInvalidId = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "invalid Role id", http.StatusNotFound)
}
```

**Var:**

ErrApiRoleNotFound returns an API error response when the requested Role is not found.

```go
var ErrApiRoleNotFound = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Role not found", http.StatusNotFound)
}
```

**Var:**

ErrApiValidateCreateRole returns an API error response for invalid Role creation request validation.

```go
var ErrApiValidateCreateRole = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "validation failed: invalid create Role request", http.StatusBadRequest)
}
```

**Var:**

ErrApiValidateUpdateRole returns an API error response for invalid Role update request validation.

```go
var ErrApiValidateUpdateRole = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "validation failed: invalid update Role request", http.StatusBadRequest)
}
```

**Var:**

ErrCreateRole is returned when creating a new Role fails.

```go
var ErrCreateRole = errors.New("failed to create Role")
```

**Var:**

ErrDeleteAllRoles is returned when permanently deleting all Roles fails.

```go
var ErrDeleteAllRoles = errors.New("failed to permanently delete all Roles")
```

**Var:**

ErrDeleteRolePermanent is returned when permanently deleting a Role fails.

```go
var ErrDeleteRolePermanent = errors.New("failed to permanently delete Role")
```

**Var:**

ErrFailedCreateRole is returned when there is a failure in creating a role.

```go
var ErrFailedCreateRole = response.NewErrorResponse("Failed to create Role", http.StatusInternalServerError)
```

**Var:**

ErrFailedDeleteAll is returned when there is a failure in permanently deleting all roles.

```go
var ErrFailedDeleteAll = response.NewErrorResponse("Failed to delete all Roles permanently", http.StatusInternalServerError)
```

**Var:**

ErrFailedDeletePermanent is returned when there is a failure in permanently deleting a role.

```go
var ErrFailedDeletePermanent = response.NewErrorResponse("Failed to delete Role permanently", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindActive is returned when there is a failure in fetching active roles.

```go
var ErrFailedFindActive = response.NewErrorResponse("Failed to fetch active Roles", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindAll is returned when there is a failure in fetching all roles.

```go
var ErrFailedFindAll = response.NewErrorResponse("Failed to fetch Roles", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindTrashed is returned when there is a failure in fetching trashed roles.

```go
var ErrFailedFindTrashed = response.NewErrorResponse("Failed to fetch trashed Roles", http.StatusInternalServerError)
```

**Var:**

ErrFailedRestoreAll is returned when there is a failure in restoring all trashed roles.

```go
var ErrFailedRestoreAll = response.NewErrorResponse("Failed to restore all Roles", http.StatusInternalServerError)
```

**Var:**

ErrFailedRestoreRole is returned when there is a failure in restoring a trashed role.

```go
var ErrFailedRestoreRole = response.NewErrorResponse("Failed to restore Role", http.StatusInternalServerError)
```

**Var:**

ErrFailedTrashedRole is returned when there is a failure in moving a role to trash.

```go
var ErrFailedTrashedRole = response.NewErrorResponse("Failed to move Role to trash", http.StatusInternalServerError)
```

**Var:**

ErrFailedUpdateRole is returned when there is a failure in updating a role.

```go
var ErrFailedUpdateRole = response.NewErrorResponse("Failed to update Role", http.StatusInternalServerError)
```

**Var:**

ErrFindActiveRoles is returned when retrieving all active Roles fails.

```go
var ErrFindActiveRoles = errors.New("failed to find active Roles")
```

**Var:**

ErrFindAllRoles is returned when retrieving all Roles from the database fails.

```go
var ErrFindAllRoles = errors.New("failed to find all Roles")
```

**Var:**

ErrFindTrashedRoles is returned when retrieving trashed (soft-deleted) Roles fails.

```go
var ErrFindTrashedRoles = errors.New("failed to find trashed Roles")
```

**Var:**

ErrInvalidRoleId returns an API error response for an invalid Role ID format.

```go
var ErrInvalidRoleId = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "invalid Role id", http.StatusBadRequest)
}
```

**Var:**

ErrRestoreAllRoles is returned when restoring all trashed Roles fails.

```go
var ErrRestoreAllRoles = errors.New("failed to restore all Roles")
```

**Var:**

ErrRestoreRole is returned when restoring a trashed Role fails.

```go
var ErrRestoreRole = errors.New("failed to restore Role from trash")
```

**Var:**

ErrRoleConflict indicates a conflict where a Role already exists.

```go
var ErrRoleConflict = errors.New("failed Role already exists")
```

**Var:**

ErrRoleNotFound indicates that the requested Role was not found in the database.

```go
var ErrRoleNotFound = errors.New("role not found")
```

**Var:**

ErrRoleNotFoundRes is returned when the requested role is not found.

```go
var ErrRoleNotFoundRes = response.NewErrorResponse("Role not found", http.StatusNotFound)
```

**Var:**

ErrTrashedRole is returned when moving a Role to trash (soft-delete) fails.

```go
var ErrTrashedRole = errors.New("failed to move Role to trash")
```

**Var:**

ErrUpdateRole is returned when updating an existing Role fails.

```go
var ErrUpdateRole = errors.New("failed to update Role")
```

