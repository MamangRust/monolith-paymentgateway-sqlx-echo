package topup_test

import (
	"context"
	"testing"
	"time"

	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	stats_repo "github.com/MamangRust/monolith-payment-gateway-topup/repository/stats"
	statsbycard_repo "github.com/MamangRust/monolith-payment-gateway-topup/repository/statsbycard"
	tests "github.com/MamangRust/monolith-payment-gateway-test"
	"github.com/jmoiron/sqlx"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/stretchr/testify/suite"
)

type TopupStatsRepositoryTestSuite struct {
	suite.Suite
	ts          *tests.TestSuite
	dbPool      *sqlx.DB
	repo        stats_repo.TopupStatsRepository
	statsbycard statsbycard_repo.TopupStatsByCardRepository
	userID      int32
	cardNumber1 string
	cardNumber2 string
	testYear    int
}

func (s *TopupStatsRepositoryTestSuite) SetupSuite() {
	ts, err := tests.SetupTestSuite()
	s.Require().NoError(err)
	s.ts = ts

	pool, err := sqlx.ConnectContext(s.ts.Ctx, "pgx", s.ts.DBURL)
	s.Require().NoError(err)
	s.dbPool = pool

	queries := pool
	s.repo = stats_repo.NewTopupStatsRepository(queries)
	s.statsbycard = statsbycard_repo.NewTopupStatsByCardRepository(queries)

	s.testYear = time.Now().Year()

	// Seed Data
	ctx := context.Background()
	err = s.dbPool.QueryRowContext(ctx, "INSERT INTO users (firstname, lastname, email, password, verification_code, is_verified) VALUES ('Topup', 'Stats', 'topup_stats@example.com', 'pass', '123', true) RETURNING user_id").Scan(&s.userID)
	s.Require().NoError(err)

	s.cardNumber1 = "1111222233334444"
	s.cardNumber2 = "5555666677778888"

	_, err = s.dbPool.ExecContext(ctx, "INSERT INTO cards (user_id, card_number, card_type, cvv, card_provider, expire_date) VALUES ($1, $2, 'debit', '123', 'visa', '2030-01-01')", s.userID, s.cardNumber1)
	s.Require().NoError(err)
	_, err = s.dbPool.ExecContext(ctx, "INSERT INTO cards (user_id, card_number, card_type, cvv, card_provider, expire_date) VALUES ($1, $2, 'debit', '123', 'mastercard', '2030-01-01')", s.userID, s.cardNumber2)
	s.Require().NoError(err)

	// Seed Topups for trailing 12 months and 5 years
	now := time.Now().UTC()

	// Topups in current month
	_, err = s.dbPool.ExecContext(ctx, "INSERT INTO topups (card_number, topup_amount, topup_method, topup_time, status) VALUES ($1, $2, $3, $4, 'success')", s.cardNumber1, 500, "bank_transfer", now.AddDate(0, 0, -1))
	s.Require().NoError(err)

	// Topups last month
	_, err = s.dbPool.ExecContext(ctx, "INSERT INTO topups (card_number, topup_amount, topup_method, topup_time, status) VALUES ($1, $2, $3, $4, 'success')", s.cardNumber1, 300, "e-wallet", now.AddDate(0, -1, -5))
	s.Require().NoError(err)

	// Topups 13 months ago (should not show in trailing 12)
	_, err = s.dbPool.ExecContext(ctx, "INSERT INTO topups (card_number, topup_amount, topup_method, topup_time, status) VALUES ($1, $2, $3, $4, 'success')", s.cardNumber1, 1000, "bank_transfer", now.AddDate(0, -13, 0))
	s.Require().NoError(err)

	// Topups for yearly stats (5 years rolling)
	_, err = s.dbPool.ExecContext(ctx, "INSERT INTO topups (card_number, topup_amount, topup_method, topup_time, status) VALUES ($1, $2, $3, $4, 'success')", s.cardNumber2, 1000, "bank_transfer", now)
	s.Require().NoError(err)
	_, err = s.dbPool.ExecContext(ctx, "INSERT INTO topups (card_number, topup_amount, topup_method, topup_time, status) VALUES ($1, $2, $3, $4, 'success')", s.cardNumber2, 2000, "bank_transfer", now.AddDate(-1, 0, 0))
	s.Require().NoError(err)
	_, err = s.dbPool.ExecContext(ctx, "INSERT INTO topups (card_number, topup_amount, topup_method, topup_time, status) VALUES ($1, $2, $3, $4, 'success')", s.cardNumber2, 3000, "bank_transfer", now.AddDate(-6, 0, 0)) // Should not show in 5 years
	s.Require().NoError(err)
}

func (s *TopupStatsRepositoryTestSuite) TearDownSuite() {
	if s.dbPool != nil {
		s.dbPool.Close()
	}
	s.ts.Teardown()
}

func (s *TopupStatsRepositoryTestSuite) TestGlobalStats() {
	ctx := context.Background()

	// Monthly Amount
	res, err := s.repo.GetMonthlyTopupAmounts(ctx, s.testYear)
	s.NoError(err)
	s.NotEmpty(res)

	// Yearly Amount
	yearlyRes, err := s.repo.GetYearlyTopupAmounts(ctx, s.testYear)
	s.NoError(err)
	s.NotEmpty(yearlyRes)

	// Monthly Methods (Trailing 12 months)
	methRes, err := s.repo.GetMonthlyTopupMethods(ctx, s.testYear)
	s.NoError(err)
	s.NotEmpty(methRes)

	// Status Stats
	// Global Monthly
	reqMonth := &requests.MonthTopupStatus{Year: s.testYear, Month: 1}
	resS, err := s.repo.GetMonthTopupStatusSuccess(ctx, reqMonth)
	s.NoError(err)
	s.NotEmpty(resS)

	resF, err := s.repo.GetMonthTopupStatusFailed(ctx, reqMonth)
	s.NoError(err)
	s.NotEmpty(resF)

}

func (s *TopupStatsRepositoryTestSuite) TestCardStats() {
	ctx := context.Background()

	reqMethod := &requests.YearMonthMethod{
		CardNumber: s.cardNumber1,
		Year:       s.testYear,
	}

	// Card 1 Monthly Amounts
	res1, err := s.statsbycard.GetMonthlyTopupAmountsByCardNumber(ctx, reqMethod)
	s.NoError(err)
	s.NotEmpty(res1)

	// Card 1 Monthly Methods
	methRes1, err := s.statsbycard.GetMonthlyTopupMethodsByCardNumber(ctx, reqMethod)
	s.NoError(err)
	s.NotEmpty(methRes1)

	// Card 1 Status success
	now := time.Now().UTC()
	reqStatus := &requests.MonthTopupStatusCardNumber{
		CardNumber: s.cardNumber1,
		Year:       s.testYear,
		Month:      int(now.Month()),
	}
	statusRes1, err := s.statsbycard.GetMonthTopupStatusSuccessByCardNumber(ctx, reqStatus)
	s.NoError(err)
	s.NotEmpty(statusRes1)

}

func TestTopupStatsRepositorySuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	suite.Run(t, new(TopupStatsRepositoryTestSuite))
}
