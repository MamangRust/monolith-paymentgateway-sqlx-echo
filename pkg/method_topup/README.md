# 📦 Package `methodtopup`

**Source Path:** `./pkg/method_topup`

## 🚀 Functions

### `PaymentMethodValidator`

PaymentMethodValidator will validate the payment method string with list of known payment methods.

It will check if the given paymentMethod is in the list of known payment methods.
If the payment method is in the list, it will return true, otherwise it will return false.

```go
func PaymentMethodValidator(paymentMethod string) bool
```

