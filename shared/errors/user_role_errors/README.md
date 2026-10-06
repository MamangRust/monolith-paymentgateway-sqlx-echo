# 📦 Package `userrole_errors`

**Source Path:** `shared/errors/user_role_errors`

## 🏷️ Variables

**Var:**

ErrAssignRoleToUser is an error that is returned from the repository layer
when an error occurs while trying to assign a role to the user.

```go
var ErrAssignRoleToUser = errors.New("failed to assign role to user")
```

**Var:**

ErrFailedAssignRoleToUser is an error that is returned from the service layer
when an error occurs while trying to assign a role to the user.

```go
var ErrFailedAssignRoleToUser = response.NewErrorResponse("Failed to assign role to user", http.StatusInternalServerError)
```

**Var:**

ErrFailedRemoveRole is an error that is returned from the service layer
when an error occurs while trying to remove a role from the user.

```go
var ErrFailedRemoveRole = response.NewErrorResponse("Failed to remove role from user", http.StatusInternalServerError)
```

**Var:**

ErrRemoveRole is an error that is returned from the repository layer
when an error occurs while trying to remove a role from the user.

```go
var ErrRemoveRole = errors.New("failed to remove role from user")
```