# 📦 Package `seeder`

**Source Path:** `pkg/database/seeder`

## 🧩 Types

### `Deps`

Deps is a struct that contains the dependencies for the seeder

```go
type Deps struct {
	DB *db.Queries
	Hash hash.HashPassword
	Ctx context.Context
	Logger logger.LoggerInterface
}
```

### `Seeder`

Seeder is a struct that contains all the seeders

```go
type Seeder struct {
	User *userSeeder
	Role *roleSeeder
	Saldo *saldoSeeder
	Topup *topupSeeder
	Withdraw *withdrawSeeder
	Transfer *transferSeeder
	Merchant *merchantSeeder
	Card *cardSeeder
	Transaction *transactionSeeder
}
```

#### Methods

##### `Run`

Run runs all the seeders in sequence with a delay of 30 seconds
between each seeder.

It calls the Seed method of each seeder in the following order:

1. User
2. Role
3. Card
4. Saldo
5. Topup
6. Withdraw
7. Transfer
8. Merchant
9. Transaction

If any of the seeders fail, the function returns an error.

Returns:
an error if any of the seeders fail, otherwise nil

```go
func (s *Seeder) Run() error
```

##### `seedWithDelay`

seedWithDelay is a helper function that seeds the given entity with a delay of 30 seconds.

It takes the name of the entity and a seed function as parameters.
The seed function is called, and if it returns an error, the function returns an error with a message
that includes the name of the entity.
After the seed function has been called, the function sleeps for 30 seconds to allow the next seeder
to run without interfering with the previous one.

Args:
entityName - the name of the entity being seeded
seedFunc - the seed function to be called

Returns:
an error if the seed function fails, otherwise nil

```go
func (s *Seeder) seedWithDelay(entityName string, seedFunc func() error) error
```

### `cardSeeder`

cardSeeder is a struct that represents a seeder for the cards table.

```go
type cardSeeder struct {
	db *db.Queries
	ctx context.Context
	logger logger.LoggerInterface
}
```

#### Methods

##### `Seed`

Seed populates the cards table with fake data.

It generates 10 cards, 5 of which are active and 5 of which are trashed.
The cards are seeded with random card numbers, card types, expire dates, and
CVV numbers. The card provider is randomly selected from a list of providers.

If any errors occur during the seeding process, an error is returned.

```go
func (r *cardSeeder) Seed() error
```

### `merchantSeeder`

merchantSeeder is a struct that represents a seeder for the merchants table.

```go
type merchantSeeder struct {
	db *db.Queries
	ctx context.Context
	logger logger.LoggerInterface
}
```

#### Methods

##### `Seed`

Seed populates the merchants table with fake data.

It creates a total of 10 merchants, with 5 of them being active and 5 of them being
trashed. The active merchants have a status of "active", while the trashed merchants
have a status of "deactive". The merchant names are generated randomly from the
combination of the adjectives and nouns provided. The API key is generated using
the GenerateApiKey function from the api-key package.

```go
func (r *merchantSeeder) Seed() error
```

### `roleSeeder`

roleSeeder is a struct that represents a seeder for the roles table.

```go
type roleSeeder struct {
	db *db.Queries
	ctx context.Context
	logger logger.LoggerInterface
}
```

#### Methods

##### `Seed`

Seed populates the roles table with fake data.

It creates a total of 10 roles, with names randomly selected from the
following list:

- Super Admin
- Admin
- Merchant Admin
- Merchant Operator
- Finance
- Compliance
- Auditor
- Support
- Viewer
- User

If any of the roles fail to be created, the function returns an error.

Returns:
an error if any of the roles fail to be created, otherwise nil

```go
func (r *roleSeeder) Seed() error
```

### `saldoSeeder`

saldoSeeder is a struct that represents a seeder for the saldos table.

```go
type saldoSeeder struct {
	db *db.Queries
	ctx context.Context
	logger logger.LoggerInterface
}
```

#### Methods

##### `Seed`

Seed populates the saldos table with fake data.

It generates 10 saldos, 5 of which are active and 5 of which are trashed.
The saldos are seeded with random totalBalance values, and the cardNumber
is randomly selected from the cards table.

If any errors occur during the seeding process, an error is returned.

Returns:
an error if any of the saldos fail to be created, otherwise nil

```go
func (r *saldoSeeder) Seed() error
```

### `topupSeeder`

topupSeeder is a struct that represents a seeder for the topups table.

```go
type topupSeeder struct {
	db *db.Queries
	ctx context.Context
	logger logger.LoggerInterface
}
```

#### Methods

##### `Seed`

Seed populates the topups table with fake data.

It creates a total of 10 topups, with 5 of them being active and 5 of them being
trashed. The topups are seeded with a random card number, topup amount, topup method,
and topup time. The status of the topup is randomly set to one of the following:
pending, success, or failed.

If any errors occur during the seeding process, an error is returned.

Returns:
an error if any of the topups fail to be created, otherwise nil

```go
func (r *topupSeeder) Seed() error
```

### `transactionSeeder`

transactionSeeder is a struct that represents a seeder for the transactions table.

```go
type transactionSeeder struct {
	db *db.Queries
	ctx context.Context
	logger logger.LoggerInterface
}
```

#### Methods

##### `Seed`

Seed populates the transactions table with fake data.

It generates a total of 10 transactions, with 5 of them being active and 5 of them being
trashed. The transactions are seeded with random payment methods, statuses, and transaction
times. The card numbers are randomly selected from the cards table. The merchant IDs are
randomly selected from the merchants table.

If any of the transactions fail to be created, the function returns an error.

Returns:
an error if any of the transactions fail to be created, otherwise nil

```go
func (r *transactionSeeder) Seed() error
```

### `transferSeeder`

transferSeeder is a struct that represents a seeder for the transfers table.

```go
type transferSeeder struct {
	db *db.Queries
	ctx context.Context
	logger logger.LoggerInterface
}
```

#### Methods

##### `Seed`

Seed populates the transfers table with fake data.

It creates a total of 10 transfers, with 5 of them being active and 5 of them being
trashed. The active transfers have a status of "pending", "success", or "failed", while
the trashed transfers have a status of "trashed". The transfers are seeded with random
transfer from and to, transfer time, and transfer amount. The transfer time is randomly
selected from the first day of each month of the current year.

If any of the transfers fail to be created, the function returns an error.

Returns:
an error if any of the transfers fail to be created, otherwise nil

```go
func (r *transferSeeder) Seed() error
```

### `userSeeder`

userSeeder is a struct that represents a seeder for the users table.

```go
type userSeeder struct {
	db *db.Queries
	hash hash.HashPassword
	ctx context.Context
	logger logger.LoggerInterface
}
```

#### Methods

##### `Seed`

Seed populates the users table with fake data.

It creates a total of 15 users, with 10 of them being active and 5 of them being
trashed. The active users have a status of "active", while the trashed users have a
status of "deactive". The user names are generated randomly from the combination of
the adjectives and nouns provided. The API key is generated using the GenerateApiKey
function from the api-key package.

If any errors occur during the seeding process, an error is returned.

Returns:
an error if any of the users fail to be created, otherwise nil

```go
func (r *userSeeder) Seed() error
```

### `withdrawSeeder`

withdrawSeeder is a struct that represents a seeder for the withdraws table.

```go
type withdrawSeeder struct {
	db *db.Queries
	ctx context.Context
	logger logger.LoggerInterface
}
```

#### Methods

##### `Seed`

Seed populates the withdraws table with fake data.

It generates a total of 10 withdraws, with 5 of them being active and 5 of them being
trashed. The withdraws are seeded with random card numbers, withdraw amounts, and
withdraw times. The status of each withdraw is randomly set to one of the following:
pending, success, or failed. The card numbers are randomly selected from the cards
table.

If any errors occur during the seeding process, an error is returned.

Returns:
an error if any of the withdraws fail to be created or updated, otherwise nil

```go
func (r *withdrawSeeder) Seed() error
```

