package card_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	sharederrors "github.com/MamangRust/monolith-payment-gateway-shared/errors"

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

type CardPaymentServiceTestSuite struct {
	suite.Suite
	ts          *tests.TestSuite
	cardService service.Service
	dbPool      *sqlx.DB
	userID      int
	cardNumber  string
	identity    *tests.IdentityStack
}

func (s *CardPaymentServiceTestSuite) SetupSuite() {
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
		FirstName: "Payment",
		LastName:  "Tester",
		Email:     fmt.Sprintf("payment.tester.%d.%d@example.com", time.Now().UnixNano(), time.Now().UnixNano()%10000),
		Password:  "password123",
	})
	s.Require().NoError(err)
	s.userID = int(user.UserID)

	// Seed Card
	cardReq := &requests.CreateCardRequest{
		UserID:       s.userID,
		CardType:     "debit",
		ExpireDate:   time.Now().AddDate(2, 0, 0),
		CVV:          "123",
		CardProvider: "VISA",
	}
	card, err := s.cardService.CreateCard(context.Background(), cardReq)
	s.Require().NoError(err)

	// Get the full card details to extract the card number
	cardDetails, err := s.cardService.FindById(context.Background(), int(card.CardID))
	s.Require().NoError(err)
	s.cardNumber = cardDetails.CardNumber
}

func (s *CardPaymentServiceTestSuite) TearDownSuite() {
	s.identity.Close()
	s.dbPool.Close()
	s.ts.Teardown()
}

func (s *CardPaymentServiceTestSuite) Test1_PostPayment() {
	ctx := context.Background()
	req := &requests.PostPaymentRequest{
		ReferenceID:    fmt.Sprintf("ref-%d", time.Now().UnixNano()),
		CardNumber:     s.cardNumber,
		Amount:         100000,
		PaymentChannel: "bank_transfer",
		BillingID:      nil,
	}

	res, err := s.cardService.PostPayment(ctx, req)
	s.NoError(err)
	s.NotNil(res)
	s.Equal(s.cardNumber, res.CardNumber)
	s.Equal(int64(100000), res.Amount)
}

func (s *CardPaymentServiceTestSuite) Test2_IdempotentReplayAndPayloadConflict() {
	ctx := context.Background()
	referenceID := fmt.Sprintf("payment-replay-%d", time.Now().UnixNano())
	firstRequest := &requests.PostPaymentRequest{
		ReferenceID:    referenceID,
		CardNumber:     s.cardNumber,
		Amount:         25000,
		PaymentChannel: "bank_transfer",
	}

	first, err := s.cardService.PostPayment(ctx, firstRequest)
	s.Require().NoError(err)

	second, err := s.cardService.PostPayment(ctx, firstRequest)
	s.Require().NoError(err)
	s.Equal(first.PaymentID, second.PaymentID)

	conflictingRequest := *firstRequest
	conflictingRequest.Amount = firstRequest.Amount + 1
	_, err = s.cardService.PostPayment(ctx, &conflictingRequest)
	s.Require().Error(err)
	appErr, ok := err.(*sharederrors.AppError)
	s.Require().True(ok)
	s.Equal(sharederrors.ErrorTypeConflict, appErr.Type)

	var count int
	err = s.dbPool.QueryRowContext(ctx, "SELECT COUNT(*) FROM card_payments WHERE reference_id = $1", referenceID).Scan(&count)
	s.Require().NoError(err)
	s.Equal(1, count)
}

func (s *CardPaymentServiceTestSuite) Test3_GetPaymentHistory() {
	ctx := context.Background()

	res, err := s.cardService.GetPaymentHistory(ctx, s.cardNumber, 1, 10)
	s.NoError(err)
	s.NotNil(res)
	s.GreaterOrEqual(len(res), 1)
}

func (s *CardPaymentServiceTestSuite) Test4_CountPayments() {
	ctx := context.Background()

	total, err := s.cardService.CountPayments(ctx, s.cardNumber)
	s.NoError(err)
	s.GreaterOrEqual(total, 1)
}

func TestCardPaymentServiceSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	suite.Run(t, new(CardPaymentServiceTestSuite))
}
