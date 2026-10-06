# 📦 Package `auth`

**Source Path:** `pkg/auth`

## 🏷️ Variables

```go
var ErrTokenExpired = errors.New("token expired")
```

## 🧩 Types

### `Manager`

```go
type Manager struct {
	secretKey []byte
}
```

#### Methods

##### `GenerateToken`

GenerateToken generates a new JWT token for the given user ID and audience.

The token is valid for 12 hours from the time it is generated.
The subject claim is set to the given user ID.
The audience claim is set to the given audience.
The token is signed with the secret key set on the Manager during initialization.

If the token cannot be generated, an error is returned.

```go
func (m *Manager) GenerateToken(userId int, audience string) (string, error)
```

##### `ValidateToken`

ValidateToken validates a JWT token and returns the user ID string if the validation is successful.
If the token is invalid or expired, an error is returned.
The error is wrapped with jwt.ErrTokenExpired if the token is expired.
The error is wrapped with jwt.ErrTokenExpired if the token is expired.

```go
func (m *Manager) ValidateToken(accessToken string) (string, error)
```

### `TokenManager`

```go
type TokenManager interface {
	GenerateToken func(userId int, audience string) (string, error)
	ValidateToken func(tokenString string) (string, error)
}
```

