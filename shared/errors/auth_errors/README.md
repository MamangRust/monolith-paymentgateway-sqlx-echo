# 📦 Package `auth_errors`

**Source Path:** `shared/errors/auth_errors`

## 🏷️ Variables

**Var:**

ErrApiForgotPassword returns an API error response when forgot password request fails.

```go
var ErrApiForgotPassword = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "failed to forgot password", http.StatusBadRequest)
}
```

**Var:**

ErrApiGetMe returns an API error response when unauthenticated user tries to access their info.

```go
var ErrApiGetMe = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "get user info failed: unauthenticated", http.StatusUnauthorized)
}
```

**Var:**

ErrApiLogin returns an API error response when login fails due to invalid arguments.

```go
var ErrApiLogin = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "login failed: invalid argument provided", http.StatusBadRequest)
}
```

**Var:**

ErrApiRefreshToken returns an API error response when access token is invalid during refresh token flow.

```go
var ErrApiRefreshToken = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "refresh-token failed: invalid access token", http.StatusBadRequest)
}
```

**Var:**

ErrApiRegister returns an API error response when register request fails due to invalid arguments.

```go
var ErrApiRegister = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "register failed: invalid argument", http.StatusBadRequest)
}
```

**Var:**

ErrApiResetPassword returns an API error response when reset password request fails.

```go
var ErrApiResetPassword = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "failed to reset password", http.StatusBadRequest)
}
```

**Var:**

ErrApiVerifyCode returns an API error response when code verification fails.

```go
var ErrApiVerifyCode = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "failed to verify code", http.StatusBadRequest)
}
```

**Var:**

ErrBindForgotPassword returns an API error response for binding failure on forgot password request payload.

```go
var ErrBindForgotPassword = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "binding failed: invalid forgot password request payload", http.StatusBadRequest)
}
```

**Var:**

ErrBindLogin returns an API error response for binding failure on login request payload.

```go
var ErrBindLogin = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "binding failed: invalid login request payload", http.StatusBadRequest)
}
```

**Var:**

ErrBindRefreshToken returns an API error response for binding failure on refresh token request payload.

```go
var ErrBindRefreshToken = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "binding failed: invalid refresh token request payload", http.StatusBadRequest)
}
```

**Var:**

ErrBindRegister returns an API error response for binding failure on register request payload.

```go
var ErrBindRegister = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "binding failed: invalid register request payload", http.StatusBadRequest)
}
```

**Var:**

ErrBindResetPassword returns an API error response for binding failure on reset password request payload.

```go
var ErrBindResetPassword = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "binding failed: invalid reset password request payload", http.StatusBadRequest)
}
```

```go
var ErrGrpcGetMe = response.NewGrpcError("error", "get user info failed: unauthenticated", http.StatusUnauthorized)
```

```go
var ErrGrpcLogin = response.NewGrpcError("error", "login failed: invalid argument provided", http.StatusBadRequest)
```

```go
var ErrGrpcRegisterToken = response.NewGrpcError("error", "register failed: invalid argument", http.StatusBadRequest)
```

**Var:**

ErrInvalidAccessToken returns an API error response for an invalid access token.

```go
var ErrInvalidAccessToken = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "invalid access token", http.StatusInternalServerError)
}
```

**Var:**

ErrInvalidLogin returns an API error response for incorrect email or password.

```go
var ErrInvalidLogin = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "invalid email or password", http.StatusUnauthorized)
}
```

**Var:**

ErrValidateForgotPassword returns an API error response for invalid forgot password request validation.

```go
var ErrValidateForgotPassword = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "validation failed: invalid forgot-password request", http.StatusBadRequest)
}
```

**Var:**

ErrValidateLogin returns an API error response for invalid login request validation.

```go
var ErrValidateLogin = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "validation failed: invalid login request", http.StatusBadRequest)
}
```

**Var:**

ErrValidateRefreshToken returns an API error response for invalid refresh token request validation.

```go
var ErrValidateRefreshToken = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "validation failed: invalid refresh-token request", http.StatusBadRequest)
}
```

**Var:**

ErrValidateRegister returns an API error response for invalid register request validation.

```go
var ErrValidateRegister = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "validation failed: invalid register request", http.StatusBadRequest)
}
```

**Var:**

ErrValidateResetPassword returns an API error response for invalid reset password request validation.

```go
var ErrValidateResetPassword = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "validation failed: invalid reset-password request", http.StatusBadRequest)
}
```

