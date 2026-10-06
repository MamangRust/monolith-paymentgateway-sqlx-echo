package withdraw_test

import (
	"context"
	"testing"
	"time"

	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/cache"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/observability"
	withdraw_cache "github.com/MamangRust/monolith-payment-gateway-withdraw/redis/stats"
	withdraw_cache_bycard "github.com/MamangRust/monolith-payment-gateway-withdraw/redis/statsbycard"
	withdraw_repo "github.com/MamangRust/monolith-payment-gateway-withdraw/repository/stats"
	withdraw_repo_bycard "github.com/MamangRust/monolith-payment-gateway-withdraw/repository/statsbycard"
	withdraw_service "github.com/MamangRust/monolith-payment-gateway-withdraw/service/stats"
	withdraw_service_bycard "github.com/MamangRust/monolith-payment-gateway-withdraw/service/statsbycard"
	tests "github.com/MamangRust/monolith-payment-gateway-test"
	"github.com/jmoiron/sqlx"
	sdklog "go.opentelemetry.io/otel/sdk/log"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/suite"
)

type WithdrawStatsServiceTestSuite struct {
	suite.Suite
	ts        *tests.TestSuite
	dbPool    *sqlx.DB
	svc       withdraw_service_bycard.WithdrawStatsByCardStatusService
	svcAmount withdraw_service_bycard.WithdrawStatsByCardAmountService
	testYear  int
}

func (s *WithdrawStatsServiceTestSuite) SetupSuite() {
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

	// Logger & Observability
	logger.ResetInstance()
	lp := sdklog.NewLoggerProvider()
	log, _ := logger.NewLogger("test", lp)
	obs, _ := observability.NewObservability("test", log)
	cacheMetrics, _ := observability.NewCacheMetrics("test")
	cacheStore := cache.NewCacheStore(redisClient, log, cacheMetrics)

	// Caches
	s_cache := withdraw_cache.NewWithdrawStatsCache(cacheStore)
	c_cache := withdraw_cache_bycard.NewWithdrawStatsByCardCache(cacheStore)

	// Repositories
	s_repo := withdraw_repo.NewWithdrawStatsStatusRepository(queries)
	sa_repo := withdraw_repo.NewWithdrawStatsAmountRepository(queries)
	c_repo := withdraw_repo_bycard.NewWithdrawStatsByCardStatusRepository(queries)
	ca_repo := withdraw_repo_bycard.NewWithdrawStatsByCardAmountRepository(queries)

	// Services
	s.svc = withdraw_service_bycard.NewWithdrawStatsByCardStatusService(&withdraw_service_bycard.WithdrawStatsByCardStatusDeps{
		Cache:         c_cache,
		Logger:        log,
		Repository:    c_repo,
		Observability: obs,
	})
	s.svcAmount = withdraw_service_bycard.NewWithdrawStatsByCardAmountService(&withdraw_service_bycard.WithdrawStatsByCardAmountDeps{
		Cache:         c_cache,
		Logger:        log,
		Repository:    ca_repo,
		Observability: obs,
	})

	_ = withdraw_service.NewWithdrawStatsStatusService(&withdraw_service.WithdrawStatsStatusDeps{
		Cache:         s_cache,
		Logger:        log,
		Repository:    s_repo,
		Observability: obs,
	})
	_ = withdraw_service.NewWithdrawStatsAmountService(&withdraw_service.WithdrawStatsAmountDeps{
		Cache:         s_cache,
		Logger:        log,
		Repository:    sa_repo,
		Observability: obs,
	})

	s.testYear = time.Now().Year()

	// Seed Data
	ctx := context.Background()
	var userID int32
	err = s.dbPool.QueryRowContext(ctx, "INSERT INTO users (firstname, lastname, email, password, verification_code, is_verified) VALUES ('WithdrawSvc', 'Stats', 'withdraw_svc_stats@example.com', 'pass', '123', true) RETURNING user_id").Scan(&userID)
	s.Require().NoError(err)

	cardNumber1 := "5555444433332222"
	_, err = s.dbPool.ExecContext(ctx, "INSERT INTO cards (user_id, card_number, card_type, cvv, card_provider, expire_date) VALUES ($1, $2, 'debit', '123', 'visa', '2030-01-01')", userID, cardNumber1)
	s.Require().NoError(err)

	// Jan Success (2000)
	_, err = s.dbPool.ExecContext(ctx, "INSERT INTO withdraws (card_number, withdraw_amount, withdraw_time, status) VALUES ($1, $2, $3, 'success')", cardNumber1, 2000, time.Date(s.testYear, 1, 10, 10, 0, 0, 0, time.UTC))
	s.Require().NoError(err)
}

func (s *WithdrawStatsServiceTestSuite) TearDownSuite() {
	if s.dbPool != nil {
		s.dbPool.Close()
	}
	s.ts.Teardown()
}

func (s *WithdrawStatsServiceTestSuite) TestByCardService() {
	ctx := context.Background()
	cardNumber1 := "5555444433332222"
	reqStatus := &requests.MonthStatusWithdrawCardNumber{Year: s.testYear, Month: 1, CardNumber: cardNumber1}
	reqYearMonth := &requests.YearMonthCardNumber{Year: s.testYear, CardNumber: cardNumber1}

	resS, err := s.svc.FindMonthWithdrawStatusSuccessByCardNumber(ctx, reqStatus)
	s.NoError(err)
	s.NotEmpty(resS)

	resA, err := s.svcAmount.FindMonthlyWithdrawsByCardNumber(ctx, reqYearMonth)
	s.NoError(err)
	s.NotEmpty(resA)

	resY, err := s.svcAmount.FindYearlyWithdrawsByCardNumber(ctx, reqYearMonth)
	s.NoError(err)
	s.NotEmpty(resY)

}

func TestWithdrawStatsServiceSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	suite.Run(t, new(WithdrawStatsServiceTestSuite))
}
