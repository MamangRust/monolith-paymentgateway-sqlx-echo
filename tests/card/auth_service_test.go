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

type CardAuthServiceTestSuite struct {
	suite.Suite
	ts          *tests.TestSuite
	cardService service.Service
	dbPool      *sqlx.DB
	cardNumber  string
	userID      int
	txnID       string
	identity    *tests.IdentityStack
}

func (s *CardAuthServiceTestSuite) SetupSuite() {
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
		FirstName: "Auth",
		LastName:  "Tester",
		Email:     fmt.Sprintf("auth.tester.%d.%d@example.com", time.Now().UnixNano(), time.Now().UnixNano()%10000),
		Password:  "password123",
	})
	s.Require().NoError(err)
	s.userID = int(user.UserID)

	// Seed Card
	card, err := s.cardService.CreateCard(context.Background(), &requests.CreateCardRequest{
		UserID:       s.userID,
		CardType:     "debit",
		ExpireDate:   time.Now().AddDate(2, 0, 0),
		CVV:          "123",
		CardProvider: "VISA",
	})
	s.Require().NoError(err)
	s.Require().NotNil(card)
	s.cardNumber = card.CardNumber
}

func (s *CardAuthServiceTestSuite) TearDownSuite() {
	s.identity.Close()
	s.dbPool.Close()
	s.ts.Teardown()
}

func (s *CardAuthServiceTestSuite) Test1_Authorize() {
	s.Require().NotEmpty(s.cardNumber)
	ctx := context.Background()

	req := &requests.AuthorizeCardRequest{
		CardNumber:     s.cardNumber,
		MerchantID:     1,
		Amount:         50000,
		Currency:       "USD",
		PosEntryMode:   "05",
		Mcc:            "5812",
		IdempotencyKey: fmt.Sprintf("auth-idem-%d", time.Now().UnixNano()),
	}

	res, err := s.cardService.Authorize(ctx, req)
	s.NoError(err)
	s.NotNil(res)
	s.Equal(s.cardNumber, res.CardNumber)
	s.Equal("pending", res.Status)
	s.NotEmpty(res.TxnID)
	s.txnID = res.TxnID
}

func (s *CardAuthServiceTestSuite) Test2_GetAuthTransaction() {
	s.Require().NotEmpty(s.txnID)
	ctx := context.Background()

	res, err := s.cardService.GetAuthTransaction(ctx, s.txnID)
	s.NoError(err)
	s.NotNil(res)
	s.Equal(s.txnID, res.TxnID)
	s.Equal(s.cardNumber, res.CardNumber)
}

func (s *CardAuthServiceTestSuite) Test3_Reverse() {
	s.Require().NotEmpty(s.txnID)
	s.Require().NotEmpty(s.cardNumber)
	ctx := context.Background()

	// The auth state machine only permits reversing approved transactions:
	// pending -> approved -> reversed. Authorize leaves the transaction
	// "pending", so approve it before attempting the reversal.
	_, err := s.dbPool.ExecContext(ctx, "UPDATE card_auth_transactions SET status = 'approved', updated_at = current_timestamp WHERE txn_id = $1", s.txnID)
	s.Require().NoError(err)

	req := &requests.ReverseTransactionRequest{
		TxnID:          s.txnID,
		CardNumber:     s.cardNumber,
		Amount:         50000,
		IdempotencyKey: fmt.Sprintf("rev-idem-%d", time.Now().UnixNano()),
	}

	res, err := s.cardService.Reverse(ctx, req)
	s.NoError(err)
	s.NotNil(res)
	s.Equal(s.txnID, res.TxnID)
	s.Equal("reversed", res.Status)
}

func (s *CardAuthServiceTestSuite) Test4_Authorize_Idempotency() {
	s.Require().NotEmpty(s.cardNumber)
	ctx := context.Background()

	idempotencyKey := fmt.Sprintf("idem-test-%d", time.Now().UnixNano())
	req := &requests.AuthorizeCardRequest{
		CardNumber:     s.cardNumber,
		MerchantID:     2,
		Amount:         25000,
		Currency:       "USD",
		PosEntryMode:   "07",
		Mcc:            "5411",
		IdempotencyKey: idempotencyKey,
	}

	first, err := s.cardService.Authorize(ctx, req)
	s.NoError(err)
	s.NotNil(first)

	second, err := s.cardService.Authorize(ctx, req)
	s.NoError(err)
	s.NotNil(second)

	s.Equal(first.TxnID, second.TxnID)
	s.Equal(first.Amount, second.Amount)
}

func TestCardAuthServiceSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	suite.Run(t, new(CardAuthServiceTestSuite))
}
