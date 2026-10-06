package transaction_test

import (
	"context"
	"testing"
	"time"

	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	repository "github.com/MamangRust/monolith-payment-gateway-transaction/repository/stats"
	statsbycard_repository "github.com/MamangRust/monolith-payment-gateway-transaction/repository/statsbycard"
	tests "github.com/MamangRust/monolith-payment-gateway-test"
	"github.com/jmoiron/sqlx"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/stretchr/testify/suite"
)

type TransactionStatsRepositoryTestSuite struct {
	suite.Suite
	ts         *tests.TestSuite
	dbPool     *sqlx.DB
	repo       repository.TransactionStatsStatusRepository
	repoByCard statsbycard_repository.TransactionStatsByCardStatusRepository
	testYear   int
	testMonth  int
}

func (s *TransactionStatsRepositoryTestSuite) SetupSuite() {
	ts, err := tests.SetupTestSuite()
	s.Require().NoError(err)
	s.ts = ts

	pool, err := sqlx.ConnectContext(s.ts.Ctx, "pgx", s.ts.DBURL)
	s.Require().NoError(err)
	s.dbPool = pool

	queries := pool
	s.repo = repository.NewTransactionStatsStatusRepository(queries)
	s.repoByCard = statsbycard_repository.NewTransactionStatsByCardStatusRepository(queries)
	s.testYear = time.Now().Year()
	s.testMonth = int(time.Now().Month())
}

func (s *TransactionStatsRepositoryTestSuite) TearDownSuite() {
	if s.dbPool != nil {
		s.dbPool.Close()
	}
	s.ts.Teardown()
}

func (s *TransactionStatsRepositoryTestSuite) TestTransactionStatusStats() {
	ctx := context.Background()

	// Seed data
	var userID int32
	err := s.dbPool.QueryRowContext(ctx, "INSERT INTO users (firstname, lastname, email, password, verification_code, is_verified) VALUES ('Transaction', 'Stats', 'transaction_stats@example.com', 'pass', '123', true) RETURNING user_id").Scan(&userID)
	s.Require().NoError(err)

	cardNumber := "1111222233334444"
	_, err = s.dbPool.ExecContext(ctx, "INSERT INTO cards (user_id, card_number, card_type, cvv, card_provider, expire_date) VALUES ($1, $2, 'debit', '123', 'visa', '2030-01-01')", userID, cardNumber)
	s.Require().NoError(err)

	var merchantID int32
	err = s.dbPool.QueryRowContext(ctx, "INSERT INTO merchants (name, api_key, user_id, status) VALUES ('Merchant Stats', 'test_api_key', $1, 'active') RETURNING merchant_id", userID).Scan(&merchantID)
	s.Require().NoError(err)

	// Successful transaction
	_, err = s.dbPool.ExecContext(ctx, "INSERT INTO transactions (card_number, amount, payment_method, merchant_id, transaction_time, status) VALUES ($1, $2, 'credit_card', $3, $4, 'success')",
		cardNumber, 100000, merchantID, time.Date(s.testYear, time.Month(s.testMonth), 10, 10, 0, 0, 0, time.UTC))
	s.Require().NoError(err)

	// Failed transaction
	_, err = s.dbPool.ExecContext(ctx, "INSERT INTO transactions (card_number, amount, payment_method, merchant_id, transaction_time, status) VALUES ($1, $2, 'credit_card', $3, $4, 'failed')",
		cardNumber, 50000, merchantID, time.Date(s.testYear, time.Month(s.testMonth), 11, 10, 0, 0, 0, time.UTC))
	s.Require().NoError(err)

	// Global Monthly
	reqMonth := &requests.MonthStatusTransaction{Year: s.testYear, Month: s.testMonth}
	resSuccess, err := s.repo.GetMonthTransactionStatusSuccess(ctx, reqMonth)
	s.NoError(err)
	s.NotEmpty(resSuccess)

	resFailed, err := s.repo.GetMonthTransactionStatusFailed(ctx, reqMonth)
	s.NoError(err)
	s.NotEmpty(resFailed)

	// Global Yearly
	resYearSuccess, err := s.repo.GetYearlyTransactionStatusSuccess(ctx, s.testYear)
	s.NoError(err)
	s.NotEmpty(resYearSuccess)

	resYearFailed, err := s.repo.GetYearlyTransactionStatusFailed(ctx, s.testYear)
	s.NoError(err)
	s.NotEmpty(resYearFailed)

	// Card Monthly
	reqCard := &requests.MonthStatusTransactionCardNumber{
		Year:       s.testYear,
		Month:      s.testMonth,
		CardNumber: cardNumber,
	}
	resCardSuccess, err := s.repoByCard.GetMonthTransactionStatusSuccessByCardNumber(ctx, reqCard)
	s.NoError(err)
	s.NotEmpty(resCardSuccess)

	resCardFailed, err := s.repoByCard.GetMonthTransactionStatusFailedByCardNumber(ctx, reqCard)
	s.NoError(err)
	s.NotEmpty(resCardFailed)

}

func TestTransactionStatsRepositorySuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	suite.Run(t, new(TransactionStatsRepositoryTestSuite))
}
