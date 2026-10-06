package withdraw_test

import (
	"context"
	"testing"
	"time"

	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	withdrawstatsrepository "github.com/MamangRust/monolith-payment-gateway-withdraw/repository/stats"
	withdrawstatsbycardrepository "github.com/MamangRust/monolith-payment-gateway-withdraw/repository/statsbycard"
	tests "github.com/MamangRust/monolith-payment-gateway-test"
	"github.com/jmoiron/sqlx"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/stretchr/testify/suite"
)

type WithdrawStatsRepositoryTestSuite struct {
	suite.Suite
	ts               *tests.TestSuite
	dbPool           *sqlx.DB
	repo             withdrawstatsrepository.WithdrawStatsStatusRepository
	repoAmount       withdrawstatsrepository.WithdrawStatsAmountRepository
	repoByCard       withdrawstatsbycardrepository.WithdrawStatsByCardStatusRepository
	repoAmountByCard withdrawstatsbycardrepository.WithdrawStatsByCardAmountRepository
	testYear         int
}

func (s *WithdrawStatsRepositoryTestSuite) SetupSuite() {
	ts, err := tests.SetupTestSuite()
	s.Require().NoError(err)
	s.ts = ts

	pool, err := sqlx.ConnectContext(s.ts.Ctx, "pgx", s.ts.DBURL)
	s.Require().NoError(err)
	s.dbPool = pool

	queries := pool
	s.repo = withdrawstatsrepository.NewWithdrawStatsStatusRepository(queries)
	s.repoAmount = withdrawstatsrepository.NewWithdrawStatsAmountRepository(queries)
	s.repoByCard = withdrawstatsbycardrepository.NewWithdrawStatsByCardStatusRepository(queries)
	s.repoAmountByCard = withdrawstatsbycardrepository.NewWithdrawStatsByCardAmountRepository(queries)
	s.testYear = time.Now().Year()

	ctx := context.Background()

	var userID int32
	err = s.dbPool.QueryRowContext(ctx, "INSERT INTO users (firstname, lastname, email, password, verification_code, is_verified) VALUES ('WithdrawRepo', 'Stats', 'withdraw_repo_stats@example.com', 'pass', '123', true) RETURNING user_id").Scan(&userID)
	s.Require().NoError(err)

	cardNumber1 := "1111222233334444"
	_, err = s.dbPool.ExecContext(ctx, "INSERT INTO cards (user_id, card_number, card_type, cvv, card_provider, expire_date) VALUES ($1, $2, 'debit', '123', 'visa', '2030-01-01')", userID, cardNumber1)
	s.Require().NoError(err)

	// Jan Success (1000)
	_, err = s.dbPool.ExecContext(ctx, "INSERT INTO withdraws (card_number, withdraw_amount, withdraw_time, status) VALUES ($1, $2, $3, 'success')", cardNumber1, 1000, time.Date(s.testYear, 1, 10, 10, 0, 0, 0, time.UTC))
	s.Require().NoError(err)
	// Feb Failed (500)
	_, err = s.dbPool.ExecContext(ctx, "INSERT INTO withdraws (card_number, withdraw_amount, withdraw_time, status) VALUES ($1, $2, $3, 'failed')", cardNumber1, 500, time.Date(s.testYear, 2, 15, 10, 0, 0, 0, time.UTC))
	s.Require().NoError(err)
}

func (s *WithdrawStatsRepositoryTestSuite) TearDownSuite() {
	if s.dbPool != nil {
		s.dbPool.Close()
	}
	s.ts.Teardown()
}

func (s *WithdrawStatsRepositoryTestSuite) TestGlobalStats() {
	ctx := context.Background()

	reqMonth := &requests.MonthStatusWithdraw{Year: s.testYear, Month: 1}
	resS, err := s.repo.GetMonthWithdrawStatusSuccess(ctx, reqMonth)
	s.NoError(err)
	s.NotEmpty(resS)

	resF, err := s.repo.GetMonthWithdrawStatusFailed(ctx, reqMonth)
	s.NoError(err)
	s.NotEmpty(resF)

	resA, err := s.repoAmount.GetMonthlyWithdraws(ctx, s.testYear)
	s.NoError(err)
	s.NotEmpty(resA)

}

func (s *WithdrawStatsRepositoryTestSuite) TestByCardStats() {
	ctx := context.Background()
	cardNumber1 := "1111222233334444"

	reqMonth := &requests.MonthStatusWithdrawCardNumber{Year: s.testYear, Month: 1, CardNumber: cardNumber1}
	reqYearMonth := &requests.YearMonthCardNumber{Year: s.testYear, CardNumber: cardNumber1}

	resS, err := s.repoByCard.GetMonthWithdrawStatusSuccessByCardNumber(ctx, reqMonth)
	s.NoError(err)
	s.NotEmpty(resS)

	resF, err := s.repoByCard.GetMonthWithdrawStatusFailedByCardNumber(ctx, reqMonth)
	s.NoError(err)
	s.NotEmpty(resF)

	resA, err := s.repoAmountByCard.GetMonthlyWithdrawsByCardNumber(ctx, reqYearMonth)
	s.NoError(err)
	s.NotEmpty(resA)

	resY, err := s.repoAmountByCard.GetYearlyWithdrawsByCardNumber(ctx, reqYearMonth)
	s.NoError(err)
	s.NotEmpty(resY)

}

func TestWithdrawStatsRepositorySuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	suite.Run(t, new(WithdrawStatsRepositoryTestSuite))
}
