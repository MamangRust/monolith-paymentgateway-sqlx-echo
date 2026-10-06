# 📦 Package `randomstring`

**Source Path:** `pkg/random_string`

## 🔢 Constants

```go
const characters = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
```

## 🚀 Functions

### `GenerateRandomString`

GenerateRandomString generates a random string with the given length
using the characters: "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789".
If the length is less than 1, it will return an empty string.
If the length is greater than the number of characters, it will return an error.

```go
func GenerateRandomString(length int) (string, error)
```

