# 📦 Package `rupiah`

**Source Path:** `pkg/rupiah`

## 🚀 Functions

### `RupiahFormat`

RupiahFormat formats a given number string into a string
with a 'Rp. ' prefix and thousand separator.

If the given number string is invalid, it returns "Rp 0".

```go
func RupiahFormat(digit string) string
```

