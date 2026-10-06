# 📦 Package `hash`

**Source Path:** `pkg/hash`

## 🏷️ Variables

```go
var (
	ErrInvalidCredentials = errors.New("invalid credentials")
)
```

## 🧩 Types

### `HashPassword`

```go
type HashPassword interface {
	HashPassword func(password string) (string, error)
	ComparePassword func(hashPassword string, password string) (error)
}
```

### `Hashing`

```go
type Hashing struct {
}
```

#### Methods

##### `ComparePassword`

ComparePassword takes a hashed password and a plaintext password and
returns an error if the passwords do not match.

Parameters:
  - hashPassword: The hashed password to compare against (string)
  - password: The plaintext password to compare to the hashed password (string)

Returns:
  - error: nil if the passwords match, otherwise an error (error)

```go
func (h Hashing) ComparePassword(hashPassword string, password string) error
```

##### `HashPassword`

HashPassword takes a plaintext password and returns a hashed version of the password.
The hashed password is a string that can be safely stored in a database or other
secure storage.

Parameters:
  - password: The plaintext password to be hashed (string)

Returns:
  - string: The hashed version of the password (string)
  - error: Any error encountered while hashing the password (error)

```go
func (h Hashing) HashPassword(password string) (string, error)
```

