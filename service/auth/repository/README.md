# 📦 Package `repository`

**Source Path:** `service/auth/internal/repository`

## 🧩 Types

### `Deps`

```go
type Deps struct {
	DB *db.Queries
	Ctx context.Context
	MapperRecord *recordmapper.RecordMapper
}
```

### `RefreshTokenRepository`

RefreshTokenRepository defines operations for refresh token management

```go
type RefreshTokenRepository interface {
	FindByToken func(token string) (*record.RefreshTokenRecord, error)
	FindByUserId func(user_id int) (*record.RefreshTokenRecord, error)
	CreateRefreshToken func(req *requests.CreateRefreshToken) (*record.RefreshTokenRecord, error)
	UpdateRefreshToken func(req *requests.UpdateRefreshToken) (*record.RefreshTokenRecord, error)
	DeleteRefreshToken func(token string) (error)
	DeleteRefreshTokenByUserId func(user_id int) (error)
}
```

### `Repositories`

```go
type Repositories struct {
	User UserRepository
	RefreshToken RefreshTokenRepository
	UserRole UserRoleRepository
	Role RoleRepository
	ResetToken ResetTokenRepository
}
```

### `ResetTokenRepository`

ResetTokenRepository defines operations for password reset token management

```go
type ResetTokenRepository interface {
	FindByToken func(token string) (*record.ResetTokenRecord, error)
	CreateResetToken func(req *requests.CreateResetTokenRequest) (*record.ResetTokenRecord, error)
	DeleteResetToken func(user_id int) (error)
}
```

### `RoleRepository`

RoleRepository defines operations for role management

```go
type RoleRepository interface {
	FindById func(role_id int) (*record.RoleRecord, error)
	FindByName func(name string) (*record.RoleRecord, error)
}
```

### `UserRepository`

UserRepository defines operations for user data persistence

```go
type UserRepository interface {
	FindByEmail func(email string) (*record.UserRecord, error)
	FindByEmailAndVerify func(email string) (*record.UserRecord, error)
	FindById func(id int) (*record.UserRecord, error)
	CreateUser func(request *requests.RegisterRequest) (*record.UserRecord, error)
	UpdateUserIsVerified func(user_id int, is_verified bool) (*record.UserRecord, error)
	UpdateUserPassword func(user_id int, password string) (*record.UserRecord, error)
	FindByVerificationCode func(verification_code string) (*record.UserRecord, error)
}
```

### `UserRoleRepository`

UserRoleRepository defines operations for user role assignments

```go
type UserRoleRepository interface {
	AssignRoleToUser func(req *requests.CreateUserRoleRequest) (*record.UserRoleRecord, error)
	RemoveRoleFromUser func(req *requests.RemoveUserRoleRequest) (error)
}
```

### `refreshTokenRepository`

```go
type refreshTokenRepository struct {
	db *db.Queries
	ctx context.Context
	mapping recordmapper.RefreshTokenRecordMapping
}
```

#### Methods

##### `CreateRefreshToken`

CreateRefreshToken generates a new refresh token for a given user ID.
It takes a CreateRefreshToken struct with user ID, token and expiration time
and returns the created refresh token record if successful, or an error if the
token creation fails.
Returns:
- *record.RefreshTokenRecord: the created refresh token record
- error: an error if the token creation fails

```go
func (r *refreshTokenRepository) CreateRefreshToken(req *requests.CreateRefreshToken) (*record.RefreshTokenRecord, error)
```

##### `DeleteRefreshToken`

```go
func (r *refreshTokenRepository) DeleteRefreshToken(token string) error
```

##### `DeleteRefreshTokenByUserId`

DeleteRefreshTokenByUserId removes a refresh token from the database using the given user ID.
It returns an error if the deletion fails, or nil if the deletion is successful.
Returns:
- error: an error if the token deletion fails, otherwise nil.

```go
func (r *refreshTokenRepository) DeleteRefreshTokenByUserId(user_id int) error
```

##### `FindByToken`

FindByToken retrieves a refresh token record from the database using the given token string.
It converts the database result into a RefreshTokenRecord object.
Returns:
- *record.RefreshTokenRecord: the refresh token record if found.
- error: an error if the token is not found or if there is a database error.

```go
func (r *refreshTokenRepository) FindByToken(token string) (*record.RefreshTokenRecord, error)
```

##### `FindByUserId`

FindByUserId retrieves a refresh token record from the database using the given user ID.
It converts the database result into a RefreshTokenRecord object.
Returns:
- *record.RefreshTokenRecord: the refresh token record if found.
- error: an error if the token is not found or if there is a database error.

```go
func (r *refreshTokenRepository) FindByUserId(user_id int) (*record.RefreshTokenRecord, error)
```

##### `UpdateRefreshToken`

UpdateRefreshToken updates a refresh token record in the database using the given user ID.
It takes an UpdateRefreshToken struct with user ID, token and expiration time
and returns the updated refresh token record if successful, or an error if the
token update fails.
Returns:
- *record.RefreshTokenRecord: the updated refresh token record
- error: an error if the token update fails

```go
func (r *refreshTokenRepository) UpdateRefreshToken(req *requests.UpdateRefreshToken) (*record.RefreshTokenRecord, error)
```

### `resetTokenRepository`

```go
type resetTokenRepository struct {
	db *db.Queries
	ctx context.Context
	mapping recordmapper.ResetTokenRecordMapping
}
```

#### Methods

##### `CreateResetToken`

CreateResetToken generates and persists a new reset token
Returns:
  - *record.ResetTokenRecord: Created token record
  - error: Error if creation fails

```go
func (r *resetTokenRepository) CreateResetToken(req *requests.CreateResetTokenRequest) (*record.ResetTokenRecord, error)
```

##### `DeleteResetToken`

DeleteResetToken removes all reset tokens for a user
Returns:
  - error: Error if deletion fails

```go
func (r *resetTokenRepository) DeleteResetToken(user_id int) error
```

##### `FindByToken`

FindByToken retrieves a reset token by its value
Returns:
  - *record.ResetTokenRecord: Token record if found
  - error: Error if operation fails (ErrTokenNotFound when token invalid)

```go
func (r *resetTokenRepository) FindByToken(code string) (*record.ResetTokenRecord, error)
```

### `roleRepository`

```go
type roleRepository struct {
	db *db.Queries
	ctx context.Context
	mapping recordmapper.RoleRecordMapping
}
```

#### Methods

##### `FindById`

FindById retrieves a role by its unique identifier from the database.

Args:
id: An integer representing the unique ID of the role to retrieve.

Returns:
*record.RoleRecord: A pointer to the RoleRecord if found.
error: An error if the role is not found or if there is a database access issue.

```go
func (r *roleRepository) FindById(id int) (*record.RoleRecord, error)
```

##### `FindByName`

FindByName retrieves a role by its name from the database.

Args:
name: The name of the role to retrieve.

Returns:
*record.RoleRecord: A pointer to the RoleRecord if found.
error: An error if the role is not found or if there is a database access issue.

```go
func (r *roleRepository) FindByName(name string) (*record.RoleRecord, error)
```

### `userRepository`

```go
type userRepository struct {
	db *db.Queries
	ctx context.Context
	mapping recordmapper.UserRecordMapping
}
```

#### Methods

##### `CreateUser`

CreateUser inserts a new user record into the database.

It accepts a RegisterRequest containing user details such as FirstName, LastName, Email,
Password, VerificationCode, and IsVerified status.

It returns a UserRecord if the creation is successful, or an error if it fails.
The error will be ErrCreateUser if there is a database error during the creation process.

```go
func (r *userRepository) CreateUser(request *requests.RegisterRequest) (*record.UserRecord, error)
```

##### `FindByEmail`

FindByEmail retrieves a user by their email address

It takes in a string email as the email address for the user.

It returns a *record.UserRecord if the user is found, or an error if the
operation fails.

In the case of an error, it will return a user_errors.ErrUserNotFound if the
user is not found in the database, or a user_errors.ErrUserNotFoundRes if
there is a database error.

```go
func (r *userRepository) FindByEmail(email string) (*record.UserRecord, error)
```

##### `FindByEmailAndVerify`

FindByEmailAndVerify retrieves a user by their email address with verification check

It takes in a string email as the email address for the user.

It returns a *record.UserRecord if the user is found and verified, or an error if the
operation fails.

In the case of an error, it will return a user_errors.ErrUserNotFound if the user is not
found in the database, or a user_errors.ErrUserNotFoundRes if there is a database error.

```go
func (r *userRepository) FindByEmailAndVerify(email string) (*record.UserRecord, error)
```

##### `FindById`

FindById retrieves a user by their unique identifier

It takes in an integer user_id as the unique identifier for the user.

It returns a *record.UserRecord if the user is found, or an error if the
operation fails.

In the case of an error, it will return a user_errors.ErrUserNotFound if the
user is not found in the database, or a user_errors.ErrUserNotFoundRes if
there is a database error.

```go
func (r *userRepository) FindById(user_id int) (*record.UserRecord, error)
```

##### `FindByVerificationCode`

FindByVerificationCode retrieves a user by their verification code

It takes in a string verification_code as the verification code for the user.

It returns a *record.UserRecord if the user is found, or an error if the
operation fails.

In the case of an error, it will return a user_errors.ErrUserNotFound if the
user is not found in the database, or a user_errors.ErrUserNotFoundRes if
there is a database error.

```go
func (r *userRepository) FindByVerificationCode(verification_code string) (*record.UserRecord, error)
```

##### `UpdateUserIsVerified`

UpdateUserIsVerified updates a user's verification status

It takes in an integer user_id to identify the user, and a boolean
is_verified to update the user's verification status.

It returns a *record.UserRecord if the update is successful, or an error
if the operation fails.

In the case of an error, it will return an ErrUpdateUserVerificationCode
if the user is not found in the database, or an ErrUpdateUserVerificationCodeRes
if there is a database error.

```go
func (r *userRepository) UpdateUserIsVerified(user_id int, is_verified bool) (*record.UserRecord, error)
```

##### `UpdateUserPassword`

UpdateUserPassword updates a user's password

It takes in an integer user_id to identify the user, and a string
password to update the user's password.

It returns a *record.UserRecord if the update is successful, or an error
if the operation fails.

In the case of an error, it will return an ErrUpdateUserPassword if the
user is not found in the database, or an ErrUpdateUserPasswordRes if
there is a database error.

```go
func (r *userRepository) UpdateUserPassword(user_id int, password string) (*record.UserRecord, error)
```

### `userRoleRepository`

```go
type userRoleRepository struct {
	db *db.Queries
	ctx context.Context
	mapping recordmapper.UserRoleRecordMapping
}
```

#### Methods

##### `AssignRoleToUser`

AssignRoleToUser assigns a role to a user.

Args:
req: a pointer to a CreateUserRoleRequest object

Returns:
a pointer to a UserRoleRecord object containing the assigned role information
an error object if the database operation fails

```go
func (r *userRoleRepository) AssignRoleToUser(req *requests.CreateUserRoleRequest) (*record.UserRoleRecord, error)
```

##### `RemoveRoleFromUser`

RemoveRoleFromUser revokes a role from a user

Args:
req: a pointer to a RemoveUserRoleRequest object

Returns:
an error object if the database operation fails

```go
func (r *userRoleRepository) RemoveRoleFromUser(req *requests.RemoveUserRoleRequest) error
```

