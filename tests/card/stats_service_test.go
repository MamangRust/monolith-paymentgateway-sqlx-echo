package card_test

import (
	"context"
	"testing"
	"time"

	"github.com/MamangRust/monolith-payment-gateway-card/repository"
	"github.com/MamangRust/monolith-payment-gateway-card/service"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/cache"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/observability"
	tests "github.com/MamangRust/monolith-payment-gateway-test"
	"github.com/jmoiron/sqlx"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/suite"
	sdklog "go.opentelemetry.io/otel/sdk/log"
)

type CardStatsServiceTestSuite struct {
	suite.Suite
	ts          *tests.TestSuite
	dbPool      *sqlx.DB
	cardService service.Service
	cardNumber1 string
	cardNumber2 string
	testYear    int
}

func (s *CardStatsServiceTestSuite) SetupSuite() {
	ts, err := tests.SetupTestSuite()
	s.Require().NoError(err)
	s.ts = ts

	pool, err := sqlx.ConnectContext(s.ts.Ctx, "pgx", s.ts.DBURL)
	s.Require().NoError(err)
	s.dbPool = pool

	opts, err := redis.ParseURL(s.ts.RedisURL)
	s.Require().NoError(err)
	redisClient := redis.NewClient(opts)

	queries := pool
	repos := repository.NewRepositories(queries, placeholderUserQueryClient())

	logger.ResetInstance()
	lp := sdklog.NewLoggerProvider()
	log, _ := logger.NewLogger("test", lp)
	cacheMetrics, _ := observability.NewCacheMetrics("test")
	cacheStore := cache.NewCacheStore(redisClient, log, cacheMetrics)

	s.cardService = service.NewService(&service.Deps{
		Repositories: repos,
		Logger:       log,
		Cache:        cacheStore,
		Kafka:        nil,
	})

	s.testYear = time.Now().Year()

	// Seed Users and Cards using direct SQL (same as repository test)
	ctx := context.Background()
	s.cardNumber1 = "2222333344445555"
	s.cardNumber2 = "6666777788889999"

	// Create a user first
	var userID int32
	err = s.dbPool.QueryRowContext(ctx, "INSERT INTO users (firstname, lastname, email, password, verification_code, is_verified) VALUES ('Service', 'Stats', 'service_stats@example.com', 'pass', '123456', true) RETURNING user_id").Scan(&userID)
	s.Require().NoError(err)

	// Create Cards
	_, err = s.dbPool.ExecContext(ctx, "INSERT INTO cards (user_id, card_number, card_type, cvv, card_provider, expire_date) VALUES ($1, $2, 'debit', '123', 'visa', '2030-01-01')", userID, s.cardNumber1)
	s.Require().NoError(err)
	_, err = s.dbPool.ExecContext(ctx, "INSERT INTO cards (user_id, card_number, card_type, cvv, card_provider, expire_date) VALUES ($1, $2, 'credit', '456', 'mastercard', '2030-01-01')", userID, s.cardNumber2)
	s.Require().NoError(err)

	s.seedHistoricalData()
}

func (s *CardStatsServiceTestSuite) seedHistoricalData() {
	// Seed Saldos
	s.insertSaldo(s.cardNumber1, 1000, time.Date(s.testYear, 1, 15, 10, 0, 0, 0, time.UTC))
	s.insertSaldo(s.cardNumber1, 2000, time.Date(s.testYear, 2, 15, 10, 0, 0, 0, time.UTC))
	s.insertSaldo(s.cardNumber2, 3000, time.Date(s.testYear, 2, 20, 10, 0, 0, 0, time.UTC))

	// Seed Topups
	s.insertTopup(s.cardNumber1, 500, time.Date(s.testYear, 1, 10, 10, 0, 0, 0, time.UTC))
	s.insertTopup(s.cardNumber1, 1000, time.Date(s.testYear, 2, 10, 10, 0, 0, 0, time.UTC))
	s.insertTopup(s.cardNumber2, 1500, time.Date(s.testYear, 2, 12, 10, 0, 0, 0, time.UTC))

	// Seed Withdraws
	s.insertWithdraw(s.cardNumber1, 100, time.Date(s.testYear, 1, 20, 10, 0, 0, 0, time.UTC))
	s.insertWithdraw(s.cardNumber1, 200, time.Date(s.testYear, 2, 20, 10, 0, 0, 0, time.UTC))
	s.insertWithdraw(s.cardNumber2, 300, time.Date(s.testYear, 2, 25, 10, 0, 0, 0, time.UTC))

	// Seed Transfers
	s.insertTransfer(s.cardNumber1, s.cardNumber2, 50, time.Date(s.testYear, 1, 25, 10, 0, 0, 0, time.UTC))
	s.insertTransfer(s.cardNumber1, s.cardNumber2, 100, time.Date(s.testYear, 2, 28, 10, 0, 0, 0, time.UTC))
}

func (s *CardStatsServiceTestSuite) insertSaldo(cardNumber string, amount int, t time.Time) {
	// Saldo keeps a single active snapshot per card (idx_saldos_card_number_active);
	// superseded snapshots are soft-deleted so the stats history still counts them.
	_, err := s.dbPool.ExecContext(context.Background(),
		"UPDATE saldos SET deleted_at = current_timestamp WHERE card_number = $1 AND deleted_at IS NULL",
		cardNumber)
	s.Require().NoError(err)
	_, err = s.dbPool.ExecContext(context.Background(),
		"INSERT INTO saldos (card_number, total_balance, created_at) VALUES ($1, $2, $3)",
		cardNumber, amount, t)
	s.Require().NoError(err)
}

func (s *CardStatsServiceTestSuite) insertTopup(cardNumber string, amount int, t time.Time) {
	_, err := s.dbPool.ExecContext(context.Background(),
		"INSERT INTO topups (card_number, topup_amount, topup_time, topup_method, status) VALUES ($1, $2, $3, 'bank_transfer', 'success')",
		cardNumber, amount, t)
	s.Require().NoError(err)
}

func (s *CardStatsServiceTestSuite) insertWithdraw(cardNumber string, amount int, t time.Time) {
	_, err := s.dbPool.ExecContext(context.Background(),
		"INSERT INTO withdraws (card_number, withdraw_amount, withdraw_time, status) VALUES ($1, $2, $3, 'success')",
		cardNumber, amount, t)
	s.Require().NoError(err)
}

func (s *CardStatsServiceTestSuite) insertTransfer(from, to string, amount int, t time.Time) {
	_, err := s.dbPool.ExecContext(context.Background(),
		"INSERT INTO transfers (transfer_from, transfer_to, transfer_amount, transfer_time, status) VALUES ($1, $2, $3, $4, 'success')",
		from, to, amount, t)
	s.Require().NoError(err)
}

func (s *CardStatsServiceTestSuite) TearDownSuite() {
	if s.dbPool != nil {
		s.dbPool.Close()
	}
	if s.ts != nil {
		s.ts.Teardown()
	}
}

// --- Balance Service Tests ---

func (s *CardStatsServiceTestSuite) TestBalanceService() {
	ctx := context.Background()

	// Global Monthly
	res, err := s.cardService.FindMonthlyBalance(ctx, s.testYear)
	s.NoError(err)
	s.NotEmpty(res)
	s.Equal(int32(1000), res[0].TotalBalance) // Jan
	s.Equal(int32(5000), res[1].TotalBalance) // Feb

	// By Card Monthly
	req1 := &requests.MonthYearCardNumberCard{CardNumber: s.cardNumber1, Year: s.testYear}
	res1, err := s.cardService.FindMonthlyBalancesByCardNumber(ctx, req1)
	s.NoError(err)
	s.Equal(int32(1000), res1[0].TotalBalance) // Jan Card 1
	s.Equal(int32(2000), res1[1].TotalBalance) // Feb Card 1

	// Yearly
	yRes, err := s.cardService.FindYearlyBalance(ctx, s.testYear)
	s.NoError(err)
	s.NotEmpty(yRes)
}

// --- Topup Service Tests ---

func (s *CardStatsServiceTestSuite) TestTopupService() {
	ctx := context.Background()

	// Global Monthly
	res, err := s.cardService.FindMonthlyTopupAmount(ctx, s.testYear)
	s.NoError(err)
	s.Equal(int32(500), res[0].TotalTopupAmount)
	s.Equal(int32(2500), res[1].TotalTopupAmount)

	// By Card Monthly
	req1 := &requests.MonthYearCardNumberCard{CardNumber: s.cardNumber1, Year: s.testYear}
	res1, err := s.cardService.FindMonthlyTopupAmountByCardNumber(ctx, req1)
	s.NoError(err)
	s.Equal(int32(500), res1[0].TotalTopupAmount)
	s.Equal(int32(1000), res1[1].TotalTopupAmount)
}

// --- Withdraw Service Tests ---

func (s *CardStatsServiceTestSuite) TestWithdrawService() {
	ctx := context.Background()

	// Global Monthly
	res, err := s.cardService.FindMonthlyWithdrawAmount(ctx, s.testYear)
	s.NoError(err)
	s.Equal(int32(100), res[0].TotalWithdrawAmount)
	s.Equal(int32(500), res[1].TotalWithdrawAmount)

	// By Card Monthly
	req1 := &requests.MonthYearCardNumberCard{CardNumber: s.cardNumber1, Year: s.testYear}
	res1, err := s.cardService.FindMonthlyWithdrawAmountByCardNumber(ctx, req1)
	s.NoError(err)
	s.Equal(int32(100), res1[0].TotalWithdrawAmount)
	s.Equal(int32(200), res1[1].TotalWithdrawAmount)
}

// --- Transfer Service Tests ---

func (s *CardStatsServiceTestSuite) TestTransferService() {
	ctx := context.Background()

	// Global Monthly Sender
	resS, err := s.cardService.FindMonthlyTransferAmountSender(ctx, s.testYear)
	s.NoError(err)
	s.Equal(int32(50), resS[0].TotalSentAmount)
	s.Equal(int32(100), resS[1].TotalSentAmount)

	// By Card Monthly Sender
	req1 := &requests.MonthYearCardNumberCard{CardNumber: s.cardNumber1, Year: s.testYear}
	res1S, err := s.cardService.FindMonthlyTransferAmountBySender(ctx, req1)
	s.NoError(err)
	s.Equal(int32(50), res1S[0].TotalSentAmount)
	s.Equal(int32(100), res1S[1].TotalSentAmount)

	// By Card Monthly Receiver
	req2 := &requests.MonthYearCardNumberCard{CardNumber: s.cardNumber2, Year: s.testYear}
	res2R, err := s.cardService.FindMonthlyTransferAmountByReceiver(ctx, req2)
	s.NoError(err)
	s.Equal(int32(50), res2R[0].TotalReceivedAmount)
	s.Equal(int32(100), res2R[1].TotalReceivedAmount)
}

func TestCardStatsServiceSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	suite.Run(t, new(CardStatsServiceTestSuite))
}
