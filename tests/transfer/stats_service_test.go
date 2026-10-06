package transfer_test

import (
	"context"
	"testing"
	"time"

	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/cache"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-transfer/repository"
	"github.com/MamangRust/monolith-payment-gateway-transfer/service"
	tests "github.com/MamangRust/monolith-payment-gateway-test"
	"github.com/jmoiron/sqlx"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type TransferStatsServiceTestSuite struct {
	suite.Suite
	ts          *tests.TestSuite
	dbPool      *sqlx.DB
	svc         service.Service
	deps        *grpcDeps
	userID      int32
	cardNumber1 string
	cardNumber2 string
	testYear    int
	testMonth   int
}

func (s *TransferStatsServiceTestSuite) SetupSuite() {
	ts, err := tests.SetupTestSuite()
	s.Require().NoError(err)
	s.ts = ts

	pool, err := sqlx.ConnectContext(s.ts.Ctx, "pgx", s.ts.DBURL)
	s.Require().NoError(err)
	s.dbPool = pool
	queries := pool

	zapLog := zap.NewNop()
	myLogger := &logger.Logger{Log: zapLog}

	redisOption, err := redis.ParseURL(s.ts.RedisURL)
	s.Require().NoError(err)
	redisClient := redis.NewClient(redisOption)
	cacheStore := cache.NewCacheStore(redisClient, myLogger, &dummyCacheMetrics{})

	deps, err := startGRPCDeps(queries, cacheStore, myLogger)
	s.Require().NoError(err)
	s.deps = deps

	repos := repository.NewRepositories(
		queries,
		deps.cardQueryClient,
		deps.cardCommandClient,
		deps.saldoQueryClient,
		deps.saldoCommandClient,
	)

	s.svc = service.NewService(&service.Deps{
		Kafka:        nil,
		Repositories: repos,
		Logger:       myLogger,
		Cache:        cacheStore,
	})

	s.testYear = time.Now().Year()
	s.testMonth = int(time.Now().Month())

	ctx := context.Background()
	err = s.dbPool.QueryRowContext(ctx, "INSERT INTO users (firstname, lastname, email, password, verification_code, is_verified) VALUES ('TransferService', 'Stats', 'transfer_svc_stats@example.com', 'pass', '123', true) RETURNING user_id").Scan(&s.userID)
	s.Require().NoError(err)

	s.cardNumber1 = "1234567890123456"
	s.cardNumber2 = "6543210987654321"
	_, err = s.dbPool.ExecContext(ctx, "INSERT INTO cards (user_id, card_number, card_type, cvv, card_provider, expire_date) VALUES ($1, $2, 'debit', '123', 'visa', '2030-01-01')", s.userID, s.cardNumber1)
	s.Require().NoError(err)
	_, err = s.dbPool.ExecContext(ctx, "INSERT INTO cards (user_id, card_number, card_type, cvv, card_provider, expire_date) VALUES ($1, $2, 'debit', '123', 'mastercard', '2030-01-01')", s.userID, s.cardNumber2)
	s.Require().NoError(err)

	// Seed Transfers
	_, err = s.dbPool.ExecContext(ctx, "INSERT INTO transfers (transfer_from, transfer_to, transfer_amount, transfer_time, status) VALUES ($1, $2, $3, $4, 'success')",
		s.cardNumber1, s.cardNumber2, 100000, time.Date(s.testYear, time.Month(s.testMonth), 10, 10, 0, 0, 0, time.UTC))
	s.Require().NoError(err)

	_, err = s.dbPool.ExecContext(ctx, "INSERT INTO transfers (transfer_from, transfer_to, transfer_amount, transfer_time, status) VALUES ($1, $2, $3, $4, 'failed')",
		s.cardNumber1, s.cardNumber2, 50000, time.Date(s.testYear, time.Month(s.testMonth), 11, 10, 0, 0, 0, time.UTC))
	s.Require().NoError(err)
}

func (s *TransferStatsServiceTestSuite) TearDownSuite() {
	if s.deps != nil {
		s.deps.Close()
	}
	if s.dbPool != nil {
		s.dbPool.Close()
	}
	s.ts.Teardown()
}

func (s *TransferStatsServiceTestSuite) TestTransferStatsService() {
	ctx := context.Background()

	// Global Monthly Success
	reqMonth := &requests.MonthStatusTransfer{Year: s.testYear, Month: s.testMonth}
	resSuccess, err := s.svc.FindMonthTransferStatusSuccess(ctx, reqMonth)
	s.NoError(err)
	s.NotEmpty(resSuccess)

	// Global Monthly Failed
	resFailed, err := s.svc.FindMonthTransferStatusFailed(ctx, reqMonth)
	s.NoError(err)
	s.NotEmpty(resFailed)

	// Yearly Success
	resYearSuccess, err := s.svc.FindYearlyTransferStatusSuccess(ctx, s.testYear)
	s.NoError(err)
	s.NotEmpty(resYearSuccess)

	// Card Specific Monthly Success
	reqCard := &requests.MonthStatusTransferCardNumber{
		Year:       s.testYear,
		Month:      s.testMonth,
		CardNumber: s.cardNumber1,
	}
	resCardSuccess, err := s.svc.FindMonthTransferStatusSuccessByCardNumber(ctx, reqCard)
	s.NoError(err)
	s.NotEmpty(resCardSuccess)

}

func TestTransferStatsServiceSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	suite.Run(t, new(TransferStatsServiceTestSuite))
}
