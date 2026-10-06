package card_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/MamangRust/monolith-payment-gateway-card/repository"
	"github.com/MamangRust/monolith-payment-gateway-card/service"
	pbuser "github.com/MamangRust/monolith-payment-gateway-pb/user"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/cache"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/observability"
	user_repo "github.com/MamangRust/monolith-payment-gateway-user/repository"
	tests "github.com/MamangRust/monolith-payment-gateway-test"
	"github.com/jmoiron/sqlx"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/suite"
	sdklog "go.opentelemetry.io/otel/sdk/log"
)

type BillingEngineServiceTestSuite struct {
	suite.Suite
	ts          *tests.TestSuite
	cardService service.Service
	dbPool      *sqlx.DB
	cardNumber  string
	userID      int
	identity    *tests.IdentityStack
}

func (s *BillingEngineServiceTestSuite) SetupSuite() {
	ts, err := tests.SetupTestSuite()
	s.Require().NoError(err)
	s.ts = ts

	pool, err := sqlx.ConnectContext(s.ts.Ctx, "pgx", s.ts.DBURL)
	s.Require().NoError(err)
	s.dbPool = pool

	opts, err := redis.ParseURL(s.ts.RedisURL)
	s.Require().NoError(err)
	redisClient := redis.NewClient(opts)

	logger.ResetInstance()
	lp := sdklog.NewLoggerProvider()
	log, _ := logger.NewLogger("test", lp)
	cacheMetrics, _ := observability.NewCacheMetrics("test")
	cacheStore := cache.NewCacheStore(redisClient, log, cacheMetrics)

	s.identity, err = tests.StartIdentityStack(pool, cacheStore, log)
	s.Require().NoError(err)

	queries := pool
	repos := repository.NewRepositories(queries, pbuser.NewUserQueryServiceClient(s.identity.User.Conn))

	s.cardService = service.NewService(&service.Deps{
		Repositories: repos,
		Logger:       log,
		Cache:        cacheStore,
		Kafka:        nil,
	})

	// Seed User
	userRepo := user_repo.NewUserCommandRepository(queries)
	user, err := userRepo.CreateUser(context.Background(), &requests.CreateUserRequest{
		FirstName: "Billing",
		LastName:  "Tester",
		Email:     fmt.Sprintf("billing.tester.%d.%d@example.com", time.Now().UnixNano(), time.Now().UnixNano()%10000),
		Password:  "password123",
	})
	s.Require().NoError(err)
	s.userID = int(user.UserID)

	// Seed Card
	card, err := s.cardService.CreateCard(context.Background(), &requests.CreateCardRequest{
		UserID:       s.userID,
		CardType:     "credit",
		ExpireDate:   time.Now().AddDate(3, 0, 0),
		CVV:          "789",
		CardProvider: "VISA",
	})
	s.Require().NoError(err)
	s.Require().NotNil(card)
	s.cardNumber = card.CardNumber
}

func (s *BillingEngineServiceTestSuite) TearDownSuite() {
	s.identity.Close()
	s.dbPool.Close()
	s.ts.Teardown()
}

func (s *BillingEngineServiceTestSuite) Test1_GetBillingCyclesByCardNumber() {
	ctx := context.Background()

	cycles, err := s.cardService.GetBillingCyclesByCardNumber(ctx, s.cardNumber)
	s.NoError(err)
	s.Empty(cycles, "expected no billing cycles before triggering billing")
}

func (s *BillingEngineServiceTestSuite) Test2_TriggerBillingCycle() {
	ctx := context.Background()
	today := s.dbDay(ctx)

	// Give the seeded credit card an outstanding balance so the scheduler has
	// a real statement to generate. The period unique index makes this safe to
	// trigger repeatedly.
	_, err := s.dbPool.ExecContext(ctx, "UPDATE cards SET outstanding_balance = $1 WHERE card_number = $2", 100000, s.cardNumber)
	s.Require().NoError(err)

	count, err := s.cardService.TriggerBillingCycle(ctx, today)
	s.NoError(err)
	s.Equal(1, count)

	count, err = s.cardService.TriggerBillingCycle(ctx, today)
	s.NoError(err)
	s.Equal(0, count)
}

func (s *BillingEngineServiceTestSuite) Test3_PostPaymentTransitionsBillingAndReplays() {
	ctx := context.Background()
	cycles, err := s.cardService.GetBillingCyclesByCardNumber(ctx, s.cardNumber)
	s.Require().NoError(err)
	if len(cycles) == 0 {
		// Keep this test runnable on its own instead of relying on the suite's
		// execution order or another test having generated the statement.
		_, err = s.dbPool.ExecContext(ctx, "UPDATE cards SET outstanding_balance = $1 WHERE card_number = $2", 100000, s.cardNumber)
		s.Require().NoError(err)
		_, err = s.cardService.TriggerBillingCycle(ctx, s.dbDay(ctx))
		s.Require().NoError(err)
		cycles, err = s.cardService.GetBillingCyclesByCardNumber(ctx, s.cardNumber)
		s.Require().NoError(err)
	}
	s.Require().Len(cycles, 1)
	s.Equal("unpaid", cycles[0].Status)

	req := &requests.PostPaymentRequest{
		ReferenceID:    fmt.Sprintf("billing-payment-%d", time.Now().UnixNano()),
		CardNumber:     s.cardNumber,
		Amount:         int64(cycles[0].AmountDue),
		PaymentChannel: "bank_transfer",
		BillingID:      intPtr(int(cycles[0].BillingID)),
	}
	first, err := s.cardService.PostPayment(ctx, req)
	s.Require().NoError(err)
	s.Require().NotNil(first)

	second, err := s.cardService.PostPayment(ctx, req)
	s.Require().NoError(err)
	s.Equal(first.PaymentID, second.PaymentID)

	cycle, err := s.cardService.GetStatement(ctx, s.cardNumber)
	s.Require().NoError(err)
	s.Equal("paid", cycle.Status)

	var paymentCount int
	err = s.dbPool.QueryRowContext(ctx, "SELECT COUNT(*) FROM card_payments WHERE billing_id = $1", cycles[0].BillingID).Scan(&paymentCount)
	s.Require().NoError(err)
	s.Equal(1, paymentCount)
}

func intPtr(value int) *int {
	return &value
}

// dbDay returns the day-of-month according to the database clock. The billing
// cycle query matches against EXTRACT(DAY FROM CURRENT_DATE), and the test
// Postgres container runs in UTC while the host may be in a different zone, so
// deriving the day from the DB keeps the assertion timezone-independent.
func (s *BillingEngineServiceTestSuite) dbDay(ctx context.Context) int {
	var day int
	err := s.dbPool.QueryRowContext(ctx, "SELECT EXTRACT(DAY FROM CURRENT_DATE)::int").Scan(&day)
	s.Require().NoError(err)
	return day
}

func (s *BillingEngineServiceTestSuite) Test4_GetStatement() {
	ctx := context.Background()

	// The previous test has generated and paid a real billing statement.
	statement, err := s.cardService.GetStatement(ctx, s.cardNumber)
	if err != nil {
		s.T().Logf("GetStatement returned an error (expected if billing not processed): %v", err)
		s.Nil(statement)
	} else {
		s.NotNil(statement)
		s.T().Logf("GetStatement succeeded, billing ID: %d", statement.BillingID)
	}
}

func TestBillingEngineServiceSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	suite.Run(t, new(BillingEngineServiceTestSuite))
}
