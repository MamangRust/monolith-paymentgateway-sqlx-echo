package topup_test

import (
	"context"
	"testing"
	"time"

	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/cache"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-topup/repository"
	"github.com/MamangRust/monolith-payment-gateway-topup/service"
	tests "github.com/MamangRust/monolith-payment-gateway-test"
	"github.com/jmoiron/sqlx"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type TopupStatsServiceTestSuite struct {
	suite.Suite
	ts          *tests.TestSuite
	dbPool      *sqlx.DB
	svc         service.Service
	userID      int32
	cardNumber1 string
	cardNumber2 string
	testYear    int
}

func (s *TopupStatsServiceTestSuite) SetupSuite() {
	ts, err := tests.SetupTestSuite()
	s.Require().NoError(err)
	s.ts = ts

	pool, err := sqlx.ConnectContext(s.ts.Ctx, "pgx", s.ts.DBURL)
	s.Require().NoError(err)
	s.dbPool = pool
	queries := pool
	// Stats methods never reach the card/saldo adapters; their clients are unused.
	repos := repository.NewRepositories(queries, nil, nil, nil, nil)

	// Setup Logger
	zapLog := zap.NewNop()
	myLogger := &logger.Logger{Log: zapLog}

	// Setup Redis & Cache
	redisOption, err := redis.ParseURL(s.ts.RedisURL)
	s.Require().NoError(err)
	redisClient := redis.NewClient(redisOption)
	cacheStore := cache.NewCacheStore(redisClient, myLogger, &dummyCacheMetrics{})

	s.svc = service.NewService(&service.Deps{
		Kafka:        nil, // Not used in stats
		Repositories: repos,
		Logger:       myLogger,
		Cache:        cacheStore,
	})

	s.testYear = time.Now().Year()

	// Seed Data
	ctx := context.Background()
	err = s.dbPool.QueryRowContext(ctx, "INSERT INTO users (firstname, lastname, email, password, verification_code, is_verified) VALUES ('TopupService', 'Stats', 'topup_svc_stats@example.com', 'pass', '123', true) RETURNING user_id").Scan(&s.userID)
	s.Require().NoError(err)

	s.cardNumber1 = "2222333344445555"
	s.cardNumber2 = "6666777788889999"

	_, err = s.dbPool.ExecContext(ctx, "INSERT INTO cards (user_id, card_number, card_type, cvv, card_provider, expire_date) VALUES ($1, $2, 'debit', '123', 'visa', '2030-01-01')", s.userID, s.cardNumber1)
	s.Require().NoError(err)
	_, err = s.dbPool.ExecContext(ctx, "INSERT INTO cards (user_id, card_number, card_type, cvv, card_provider, expire_date) VALUES ($1, $2, 'debit', '123', 'mastercard', '2030-01-01')", s.userID, s.cardNumber2)
	s.Require().NoError(err)

	// Seed Topups
	_, err = s.dbPool.ExecContext(ctx, "INSERT INTO topups (card_number, topup_amount, topup_method, topup_time, status) VALUES ($1, $2, $3, $4, 'success')", s.cardNumber1, 1000, "bank_transfer", time.Date(s.testYear, 1, 10, 10, 0, 0, 0, time.UTC))
	s.Require().NoError(err)
	_, err = s.dbPool.ExecContext(ctx, "INSERT INTO topups (card_number, topup_amount, topup_method, topup_time, status) VALUES ($1, $2, $3, $4, 'success')", s.cardNumber2, 2000, "bank_transfer", time.Date(s.testYear, 1, 15, 10, 0, 0, 0, time.UTC))
	s.Require().NoError(err)
}

func (s *TopupStatsServiceTestSuite) TearDownSuite() {
	if s.dbPool != nil {
		s.dbPool.Close()
	}
	s.ts.Teardown()
}

func (s *TopupStatsServiceTestSuite) TestTopupStatsService() {
	ctx := context.Background()
	// Global Monthly
	res, err := s.svc.FindMonthlyTopupAmounts(ctx, s.testYear)
	s.NoError(err)
	s.NotEmpty(res)

	// Card Specific Monthly
	reqMethod := &requests.YearMonthMethod{
		Year:       s.testYear,
		CardNumber: s.cardNumber1,
	}
	resCard, err := s.svc.FindMonthlyTopupAmountsByCardNumber(ctx, reqMethod)
	s.NoError(err)
	s.NotEmpty(resCard)

}

func TestTopupStatsServiceSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	suite.Run(t, new(TopupStatsServiceTestSuite))
}
