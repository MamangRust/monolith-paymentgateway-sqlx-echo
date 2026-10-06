# 📦 Package `user_errors`

**Source Path:** `shared/errors/user_errors`

## 🏷️ Variables

**Var:**

ErrApiBindCreateUser is an error response for bind failure: invalid create user request.

```go
var ErrApiBindCreateUser = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "bind failed: invalid create User request", http.StatusBadRequest)
}
```

**Var:**

ErrApiBindUpdateUser is an error response for bind failure: invalid update user request.

```go
var ErrApiBindUpdateUser = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "bind failed: invalid update User request", http.StatusBadRequest)
}
```

**Var:**

ErrApiFailedCreateUser is an error response for failing to create a user.

```go
var ErrApiFailedCreateUser = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "failed to create User", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedDeleteAll is an error response for failed to permanently delete all users.

```go
var ErrApiFailedDeleteAll = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "failed to delete all Users permanently", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedDeletePermanent is an error response for failed to delete user permanently.

```go
var ErrApiFailedDeletePermanent = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "failed to delete User permanently", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindActive is an error response for failing to fetch active users.

```go
var ErrApiFailedFindActive = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "failed to fetch active Users", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindAll is an error response for failing to fetch all users.

```go
var ErrApiFailedFindAll = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "failed to fetch Users", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedFindTrashed is an error response for failing to fetch trashed users.

```go
var ErrApiFailedFindTrashed = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "failed to fetch trashed Users", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedRestoreAll is an error response for failed to restore all users.

```go
var ErrApiFailedRestoreAll = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "failed to restore all Users", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedRestoreUser is an error response for failed to restore user.

```go
var ErrApiFailedRestoreUser = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "failed to restore User", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedTrashedUser is an error response for failed to move user to trash.

```go
var ErrApiFailedTrashedUser = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "failed to move User to trash", http.StatusInternalServerError)
}
```

**Var:**

ErrApiFailedUpdateUser is an error response for failing to update a user.

```go
var ErrApiFailedUpdateUser = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "failed to update User", http.StatusInternalServerError)
}
```

**Var:**

ErrApiUserInvalidId is an error response for an invalid user ID.

```go
var ErrApiUserInvalidId = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "invalid User id", http.StatusNotFound)
}
```

**Var:**

ErrApiUserNotFound is an error response for when the user is not found.

```go
var ErrApiUserNotFound = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "User not found", http.StatusNotFound)
}
```

**Var:**

ErrApiValidateCreateUser is an error response for validation failure in creating a user.

```go
var ErrApiValidateCreateUser = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "validation failed: invalid create User request", http.StatusBadRequest)
}
```

**Var:**

ErrApiValidateUpdateUser is an error response for validation failure in updating a user.

```go
var ErrApiValidateUpdateUser = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "validation failed: invalid update User request", http.StatusBadRequest)
}
```

**Var:**

ErrCreateUser is returned when creating a user fails.

```go
var ErrCreateUser = errors.New("failed to create user")
```

**Var:**

ErrDeleteAllUsers is returned when permanently deleting all users fails.

```go
var ErrDeleteAllUsers = errors.New("failed to permanently delete all users")
```

**Var:**

ErrDeleteUserPermanent is returned when permanently deleting a user fails.

```go
var ErrDeleteUserPermanent = errors.New("failed to permanently delete user")
```

**Var:**

ErrFailedCreateUser is returned when creating a user fails.

```go
var ErrFailedCreateUser = response.NewErrorResponse("Failed to create user", http.StatusInternalServerError)
```

**Var:**

ErrFailedDeleteAll is returned when permanently deleting all users fails.

```go
var ErrFailedDeleteAll = response.NewErrorResponse("Failed to delete all users permanently", http.StatusInternalServerError)
```

**Var:**

ErrFailedDeletePermanent is returned when permanently deleting a user fails.

```go
var ErrFailedDeletePermanent = response.NewErrorResponse("Failed to delete user permanently", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindActive is returned when fetching active users fails.

```go
var ErrFailedFindActive = response.NewErrorResponse("Failed to fetch active users", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindAll is returned when fetching users fails.

```go
var ErrFailedFindAll = response.NewErrorResponse("Failed to fetch users", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindTrashed is returned when fetching trashed users fails.

```go
var ErrFailedFindTrashed = response.NewErrorResponse("Failed to fetch trashed users", http.StatusInternalServerError)
```

**Var:**

ErrFailedPasswordNoMatch is returned when passwords do not match.

```go
var ErrFailedPasswordNoMatch = response.NewErrorResponse("Failed password not match", http.StatusBadRequest)
```

**Var:**

ErrFailedRestoreAll is returned when restoring all users fails.

```go
var ErrFailedRestoreAll = response.NewErrorResponse("Failed to restore all users", http.StatusInternalServerError)
```

**Var:**

ErrFailedRestoreUser is returned when restoring a user fails.

```go
var ErrFailedRestoreUser = response.NewErrorResponse("Failed to restore user", http.StatusInternalServerError)
```

**Var:**

ErrFailedSendEmail is returned when sending an email fails.

```go
var ErrFailedSendEmail = response.NewErrorResponse("Failed to send email", http.StatusInternalServerError)
```

**Var:**

ErrFailedTrashedUser is returned when moving a user to trash fails.

```go
var ErrFailedTrashedUser = response.NewErrorResponse("Failed to move user to trash", http.StatusInternalServerError)
```

**Var:**

ErrFailedUpdateUser is returned when updating a user fails.

```go
var ErrFailedUpdateUser = response.NewErrorResponse("Failed to update user", http.StatusInternalServerError)
```

**Var:**

ErrFindActiveUsers is returned when fetching active users fails.

```go
var ErrFindActiveUsers = errors.New("failed to find active users")
```

**Var:**

ErrFindAllUsers is returned when fetching all users fails.

```go
var ErrFindAllUsers = errors.New("failed to find all users")
```

**Var:**

ErrFindTrashedUsers is returned when fetching trashed users fails.

```go
var ErrFindTrashedUsers = errors.New("failed to find trashed users")
```

**Var:**

ErrInternalServerError is a generic internal server error.

```go
var ErrInternalServerError = response.NewErrorResponse("Internal server error", http.StatusInternalServerError)
```

**Var:**

ErrInvalidUserId is an error response for an invalid user ID.

```go
var ErrInvalidUserId = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "invalid User id", http.StatusBadRequest)
}
```

**Var:**

ErrRestoreAllUsers is returned when restoring all users fails.

```go
var ErrRestoreAllUsers = errors.New("failed to restore all users")
```

**Var:**

ErrRestoreUser is returned when restoring a user from trash fails.

```go
var ErrRestoreUser = errors.New("failed to restore user from trash")
```

**Var:**

ErrTrashedUser is returned when moving a user to trash fails.

```go
var ErrTrashedUser = errors.New("failed to move user to trash")
```

**Var:**

ErrUpdateUser is returned when updating a user fails.

```go
var ErrUpdateUser = errors.New("failed to update user")
```

**Var:**

ErrUpdateUserPassword is returned when updating a user password fails.

```go
var ErrUpdateUserPassword = errors.New("failed to update user password")
```

**Var:**

ErrUpdateUserVerificationCode is returned when updating a user verification code fails.

```go
var ErrUpdateUserVerificationCode = errors.New("failed to update user verification code")
```

**Var:**

ErrUserConflict is returned when a user with the same email address already exists.

```go
var ErrUserConflict = errors.New("failed user already exists")
```

**Var:**

ErrUserEmailAlready is returned when a user email already exists.

```go
var ErrUserEmailAlready = response.NewErrorResponse("User email already exists", http.StatusBadRequest)
```

**Var:**

ErrUserNotFound is returned when a user is not found.

```go
var ErrUserNotFound = errors.New("user not found")
```

**Var:**

ErrUserNotFoundRes is returned when a user is not found.

```go
var ErrUserNotFoundRes = response.NewErrorResponse("User not found", http.StatusNotFound)
```

**Var:**

ErrUserPassword is returned when there is an invalid password.

```go
var ErrUserPassword = response.NewErrorResponse("Failed invalid password", http.StatusBadRequest)
```