# 📦 Package `mencache`

**Source Path:** `service/auth/internal/redis`

## 🏷️ Variables

```go
var (
	keyIdentityRefreshToken	= "identity:refresh_token:%s"
	keyIdentityUserInfo	= "identity:user_info:%s"
)
```

```go
var (
	keyPasswordResetToken	= "password_reset:token:%s"
	keyVerifyCode		= "register:verify_code:%s"
)
```

```go
var keylogin = "auth:login:%s"
```

## 🧩 Types

### `CacheStore`

```go
type CacheStore struct {
	ctx context.Context
	redis *redis.Client
	logger logger.LoggerInterface
}
```

### `Deps`

```go
type Deps struct {
	Ctx context.Context
	Redis *redis.Client
	Logger logger.LoggerInterface
}
```

### `IdentityCache`

IdentityCache defines the interface for identity-related caching operations.
It provides methods to manage refresh tokens and user information in cache.

```go
type IdentityCache interface {
	SetRefreshToken func(token string, expiration time.Duration)
	GetRefreshToken func(token string) (string, bool)
	DeleteRefreshToken func(token string)
	SetCachedUserInfo func(user *response.UserResponse, expiration time.Duration)
	GetCachedUserInfo func(userId string) (*response.UserResponse, bool)
	DeleteCachedUserInfo func(userId string)
}
```

### `LoginCache`

LoginCache defines the interface for login-related caching operations.
It provides methods to manage cached login sessions and tokens.

```go
type LoginCache interface {
	SetCachedLogin func(email string, data *response.TokenResponse, expiration time.Duration)
	GetCachedLogin func(email string) (*response.TokenResponse, bool)
}
```

### `Mencache`

```go
type Mencache struct {
	IdentityCache IdentityCache
	LoginCache LoginCache
	PasswordResetCache PasswordResetCache
	RegisterCache RegisterCache
}
```

### `PasswordResetCache`

PasswordResetCache defines the interface for password reset caching operations.
It provides methods to manage password reset tokens and verification codes.

```go
type PasswordResetCache interface {
	SetResetTokenCache func(token string, userID int, expiration time.Duration)
	GetResetTokenCache func(token string) (int, bool)
	DeleteResetTokenCache func(token string)
	DeleteVerificationCodeCache func(email string)
}
```

### `RegisterCache`

RegisterCache defines the interface for registration-related caching operations.
It provides methods to manage verification codes during user registration.

```go
type RegisterCache interface {
	SetVerificationCodeCache func(email string, code string, expiration time.Duration)
}
```

### `identityCache`

```go
type identityCache struct {
	store *CacheStore
}
```

#### Methods

##### `DeleteCachedUserInfo`

DeleteCachedUserInfo removes user information from cache.
Parameters:
  - userId: The user ID string to remove from cache

```go
func (c *identityCache) DeleteCachedUserInfo(userId string)
```

##### `DeleteRefreshToken`

DeleteRefreshToken removes a refresh token from the cache.
Parameters:
  - token: The refresh token string to remove

```go
func (c *identityCache) DeleteRefreshToken(token string)
```

##### `GetCachedUserInfo`

GetCachedUserInfo retrieves user information from cache.
Parameters:
  - userId: The user ID string to look up

Returns:
  - *UserResponse: Pointer to cached user data if found
  - bool: True if user exists in cache, false otherwise

```go
func (c *identityCache) GetCachedUserInfo(userId string) (*response.UserResponse, bool)
```

##### `GetRefreshToken`

GetRefreshToken retrieves a refresh token from the cache.
Parameters:
  - token: The refresh token string to look up

Returns:
  - string: The stored token value if found
  - bool: True if token exists in cache, false otherwise

```go
func (c *identityCache) GetRefreshToken(token string) (string, bool)
```

##### `SetCachedUserInfo`

SetCachedUserInfo stores user information in cache with expiration duration.
Parameters:
  - user: The UserResponse containing user details to cache
  - expiration: Duration until the user info expires from cache

```go
func (c *identityCache) SetCachedUserInfo(user *response.UserResponse, expiration time.Duration)
```

##### `SetRefreshToken`

SetRefreshToken stores a refresh token in the cache with the specified expiration duration.
It formats the cache key using the token and sets the token into the cache.
Parameters:
  - token: The refresh token string to be cached.
  - expiration: The duration until the token expires and is removed from the cache.

```go
func (c *identityCache) SetRefreshToken(token string, expiration time.Duration)
```

### `loginCache`

```go
type loginCache struct {
	store *CacheStore
}
```

#### Methods

##### `GetCachedLogin`

GetCachedLogin retrieves the cached login token response associated with the given email.
It formats the cache key using the email and attempts to fetch the token data from the cache.
Returns a pointer to the TokenResponse and a boolean indicating whether the data was found in the cache.

```go
func (s *loginCache) GetCachedLogin(email string) (*response.TokenResponse, bool)
```

##### `SetCachedLogin`

SetCachedLogin stores the provided login token response in the cache with an expiration duration.
Parameters:
  - email: The user's email address used as the cache key.
  - data: A pointer to the TokenResponse containing login tokens to be cached.
  - expiration: The duration until the cached data expires and is removed from the cache.

```go
func (s *loginCache) SetCachedLogin(email string, data *response.TokenResponse, expiration time.Duration)
```

### `passwordResetCache`

```go
type passwordResetCache struct {
	store *CacheStore
}
```

#### Methods

##### `DeleteResetTokenCache`

DeleteResetTokenCache removes a password reset token from the cache.
Parameters:
  - token: The reset token string to remove

```go
func (c *passwordResetCache) DeleteResetTokenCache(token string)
```

##### `DeleteVerificationCodeCache`

DeleteVerificationCodeCache removes a verification code from cache.
Parameters:
  - email: The user's email address whose code should be removed

```go
func (c *passwordResetCache) DeleteVerificationCodeCache(email string)
```

##### `GetResetTokenCache`

GetResetTokenCache retrieves a user ID associated with a reset token from cache.
Parameters:
  - token: The reset token string to look up

Returns:
  - int: The associated user ID if found
  - bool: True if token exists in cache, false otherwise

```go
func (c *passwordResetCache) GetResetTokenCache(token string) (int, bool)
```

##### `SetResetTokenCache`

SetResetTokenCache stores a password reset token in cache with expiration duration.
Parameters:
  - token: The reset token string to store
  - userID: The associated user ID to cache with the token
  - expiration: Duration until the token expires from cache

```go
func (c *passwordResetCache) SetResetTokenCache(token string, userID int, expiration time.Duration)
```

### `registerCache`

```go
type registerCache struct {
	store *CacheStore
}
```

#### Methods

##### `SetVerificationCodeCache`

SetVerificationCodeCache stores a verification code in cache with expiration duration.
Parameters:
  - email: The user's email address as cache key
  - code: The verification code string to store
  - expiration: Duration until the code expires from cache

```go
func (c *registerCache) SetVerificationCodeCache(email string, code string, expiration time.Duration)
```

## 🚀 Functions

### `DeleteFromCache`

DeleteFromCache removes the entry associated with the specified key from the Redis cache.
It accepts a CacheStore and a key string as parameters.
If the deletion process encounters an error, it logs the error message along with the cache key.

```go
func DeleteFromCache(store *CacheStore, key string)
```

### `GetFromCache`

GetFromCache attempts to retrieve a cached item from Redis by the specified key.
It returns a pointer to the item of type T and a boolean indicating whether the
item was found in the cache. If the item is not found or an error occurs during
retrieval or unmarshalling, it logs an error and returns nil and false.

```go
func GetFromCache[T any](store *CacheStore, key string) (*T, bool)
```

### `SetToCache`

SetToCache stores the given data in Redis cache under the specified key with an expiration duration.
It accepts a CacheStore, a key string, the data to be cached, and the expiration time.
The data is marshaled into JSON format before being set in the cache.
If marshaling fails, an error is logged and the function returns.
If setting the data in the cache fails, an error is logged; otherwise, a debug log indicates success.

```go
func SetToCache[T any](store *CacheStore, key string, data *T, expiration time.Duration)
```

