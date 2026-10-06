package seeder

import (
	"context"
	"fmt"
	"time"

	"github.com/MamangRust/monolith-payment-gateway-pkg/hash"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/jmoiron/sqlx"
)

// Deps is a struct that contains the dependencies for the seeder
type Deps struct {
	DB     *sqlx.DB
	Hash   hash.HashPassword
	Ctx    context.Context
	Logger logger.LoggerInterface
}

// Seeder is a struct that contains all the seeders
type Seeder struct {
	User        *userSeeder
	Role        *roleSeeder
	Saldo       *saldoSeeder
	Topup       *topupSeeder
	Withdraw    *withdrawSeeder
	Transfer    *transferSeeder
	Merchant    *merchantSeeder
	Card        *cardSeeder
	Transaction *transactionSeeder
}

// NewSeeder initializes and returns the Seeder.
//
// It takes a Deps struct which contains the required dependencies
// to initialize the Seeder.
//
// The returned Seeder contains the following:
// - User: the user seeder
// - Role: the role seeder
// - Saldo: the saldo seeder
// - Topup: the topup seeder
// - Withdraw: the withdraw seeder
// - Transfer: the transfer seeder
// - Merchant: the merchant seeder
// - Card: the card seeder
// - Transaction: the transaction seeder
func NewSeeder(deps Deps) *Seeder {
	return &Seeder{
		User:        NewUserSeeder(deps.DB, deps.Ctx, deps.Hash, deps.Logger),
		Role:        NewRoleSeeder(deps.DB, deps.Ctx, deps.Logger),
		Saldo:       NewSaldoSeeder(deps.DB, deps.Ctx, deps.Logger),
		Topup:       NewTopupSeeder(deps.DB, deps.Ctx, deps.Logger),
		Withdraw:    NewWithdrawSeeder(deps.DB, deps.Ctx, deps.Logger),
		Transfer:    NewTransferSeeder(deps.DB, deps.Ctx, deps.Logger),
		Merchant:    NewMerchantSeeder(deps.DB, deps.Ctx, deps.Logger),
		Card:        NewCardSeeder(deps.DB, deps.Ctx, deps.Logger),
		Transaction: NewTransactionSeeder(deps.DB, deps.Ctx, deps.Logger),
	}
}

// Run runs all the seeders in sequence with a delay of 30 seconds
// between each seeder.
//
// It calls the Seed method of each seeder in the following order:
//
// 1. User
// 2. Role
// 3. Card
// 4. Saldo
// 5. Topup
// 6. Withdraw
// 7. Transfer
// 8. Merchant
// 9. Transaction
//
// If any of the seeders fail, the function returns an error.
//
// Returns:
// an error if any of the seeders fail, otherwise nil
func (s *Seeder) Run() error {
	if err := s.seedWithDelay("roles", s.Role.Seed); err != nil {
		return err
	}

	if err := s.seedWithDelay("users", s.User.Seed); err != nil {
		return err
	}

	if err := s.seedWithDelay("cards", s.Card.Seed); err != nil {
		return err
	}

	if err := s.seedWithDelay("saldo", s.Saldo.Seed); err != nil {
		return err
	}

	if err := s.seedWithDelay("topups", s.Topup.Seed); err != nil {
		return err
	}

	if err := s.seedWithDelay("withdrawals", s.Withdraw.Seed); err != nil {
		return err
	}

	if err := s.seedWithDelay("transfers", s.Transfer.Seed); err != nil {
		return err
	}

	if err := s.seedWithDelay("merchants", s.Merchant.Seed); err != nil {
		return err
	}

	if err := s.seedWithDelay("transactions", s.Transaction.Seed); err != nil {
		return err
	}

	return nil
}

// seedWithDelay is a helper function that seeds the given entity with a delay of 30 seconds.
//
// It takes the name of the entity and a seed function as parameters.
// The seed function is called, and if it returns an error, the function returns an error with a message
// that includes the name of the entity.
// After the seed function has been called, the function sleeps for 30 seconds to allow the next seeder
// to run without interfering with the previous one.
//
// Args:
// entityName - the name of the entity being seeded
// seedFunc - the seed function to be called
//
// Returns:
// an error if the seed function fails, otherwise nil
func (s *Seeder) seedWithDelay(entityName string, seedFunc func() error) error {
	if err := seedFunc(); err != nil {
		return fmt.Errorf("failed to seed %s: %w", entityName, err)
	}

	time.Sleep(30 * time.Second)
	return nil
}
