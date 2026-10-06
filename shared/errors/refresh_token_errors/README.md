# 📦 Package `refreshtoken_errors`

**Source Path:** `shared/errors/refresh_token_errors`

## 🏷️ Variables

**Var:**

ErrCreateRefreshToken indicates that an error occurred while creating a new refresh token.

```go
var ErrCreateRefreshToken = errors.New("failed to create refresh token")
```

**Var:**

ErrDeleteByUserID indicates a failure when attempting to delete a refresh token using the user ID.

```go
var ErrDeleteByUserID = errors.New("failed to delete refresh token by user ID")
```

**Var:**

ErrDeleteRefreshToken is returned when deleting a refresh token fails.

```go
var ErrDeleteRefreshToken = errors.New("failed to delete refresh token")
```

**Var:**

ErrFailedCreateAccess is returned when the creation of an access token fails.

```go
var ErrFailedCreateAccess = response.NewErrorResponse("Failed to create access token", http.StatusInternalServerError)
```

**Var:**

ErrFailedCreateRefresh is returned when the creation of a refresh token fails.

```go
var ErrFailedCreateRefresh = response.NewErrorResponse("Failed to create refresh token", http.StatusInternalServerError)
```

**Var:**

ErrFailedCreateRefreshToken is returned when refresh token creation fails.

```go
var ErrFailedCreateRefreshToken = response.NewErrorResponse("Failed to create refresh token", http.StatusInternalServerError)
```

**Var:**

ErrFailedDeleteByUserID is returned when deletion of a refresh token by user ID fails.

```go
var ErrFailedDeleteByUserID = response.NewErrorResponse("Failed to delete refresh token by user ID", http.StatusInternalServerError)
```

**Var:**

ErrFailedDeleteRefreshToken is returned when refresh token deletion fails.

```go
var ErrFailedDeleteRefreshToken = response.NewErrorResponse("Failed to delete refresh token", http.StatusInternalServerError)
```

**Var:**

ErrFailedExpire occurs when expiring a refresh token fails.

```go
var ErrFailedExpire = response.NewErrorResponse("Failed to find refresh token by token", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindByToken indicates failure when searching for a refresh token by its token value.

```go
var ErrFailedFindByToken = response.NewErrorResponse("Failed to find refresh token by token", http.StatusInternalServerError)
```

**Var:**

ErrFailedFindByUserID indicates failure when searching for a refresh token by user ID.

```go
var ErrFailedFindByUserID = response.NewErrorResponse("Failed to find refresh token by user ID", http.StatusInternalServerError)
```

**Var:**

ErrFailedInValidToken is returned when an access token is invalid.

```go
var ErrFailedInValidToken = response.NewErrorResponse("Failed to invalid access token", http.StatusInternalServerError)
```

**Var:**

ErrFailedInValidUserId is returned when a user ID is invalid.

```go
var ErrFailedInValidUserId = response.NewErrorResponse("Failed to invalid user id", http.StatusInternalServerError)
```

**Var:**

ErrFailedParseExpirationDate indicates failure when parsing the expiration date of a token.

```go
var ErrFailedParseExpirationDate = response.NewErrorResponse("Failed to parse expiration date", http.StatusBadRequest)
```

**Var:**

ErrFailedUpdateRefreshToken is returned when refresh token update fails.

```go
var ErrFailedUpdateRefreshToken = response.NewErrorResponse("Failed to update refresh token", http.StatusInternalServerError)
```

**Var:**

ErrFindByToken is returned when a lookup for the refresh token by token value fails.

```go
var ErrFindByToken = errors.New("failed to find refresh token by token")
```

**Var:**

ErrFindByUserID is returned when a lookup for the refresh token by user ID fails.

```go
var ErrFindByUserID = errors.New("failed to find refresh token by user ID")
```

**Var:**

ErrParseDate is returned when parsing the expiration date of a token fails.

```go
var ErrParseDate = errors.New("failed to parse expiration date")
```

**Var:**

ErrRefreshTokenNotFound indicates that the refresh token was not found.

```go
var ErrRefreshTokenNotFound = response.NewErrorResponse("Refresh token not found", http.StatusNotFound)
```

**Var:**

ErrTokenNotFound indicates that the refresh token could not be found.

```go
var ErrTokenNotFound = errors.New("refresh token not found")
```

**Var:**

ErrUpdateRefreshToken is returned when the refresh token update process fails.

```go
var ErrUpdateRefreshToken = errors.New("failed to update refresh token")
```

