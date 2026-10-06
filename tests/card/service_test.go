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

type CardServiceTestSuite struct {
	suite.Suite
	ts          *tests.TestSuite
	cardService service.Service
	dbPool      *sqlx.DB
	cardID      int
	userID      int
	identity    *tests.IdentityStack
}

func (s *CardServiceTestSuite) SetupSuite() {
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
		FirstName: "Card",
		LastName:  "Tester",
		Email:     fmt.Sprintf("card.tester.%d.%d@example.com", time.Now().UnixNano(), time.Now().UnixNano()%10000),
		Password:  "password123",
	})
	s.Require().NoError(err)
	s.userID = int(user.UserID)
}

func (s *CardServiceTestSuite) TearDownSuite() {
	s.identity.Close()
	s.dbPool.Close()
	s.ts.Teardown()
}

func (s *CardServiceTestSuite) Test1_CreateCard() {
	ctx := context.Background()
	req := &requests.CreateCardRequest{
		UserID:       s.userID,
		CardType:     "debit",
		ExpireDate:   time.Now().AddDate(2, 0, 0),
		CVV:          "123",
		CardProvider: "VISA",
	}

	res, err := s.cardService.CreateCard(ctx, req)
	s.NoError(err)
	s.NotNil(res)
	s.Equal(int32(s.userID), res.UserID)
	s.cardID = int(res.CardID)
}

func (s *CardServiceTestSuite) Test2_FindById() {
	s.Require().NotZero(s.cardID)
	ctx := context.Background()

	res, err := s.cardService.FindById(ctx, s.cardID)
	s.NoError(err)
	s.NotNil(res)
	s.Equal(int32(s.cardID), res.CardID)
}

func (s *CardServiceTestSuite) Test3_FindAll() {
	ctx := context.Background()
	req := &requests.FindAllCards{
		Page:     1,
		PageSize: 10,
		Search:   "",
	}

	res, total, err := s.cardService.FindAll(ctx, req)
	s.NoError(err)
	s.NotNil(res)
	s.NotZero(*total)
}

func (s *CardServiceTestSuite) Test4_FindByActive() {
	ctx := context.Background()
	req := &requests.FindAllCards{
		Page:     1,
		PageSize: 10,
		Search:   "",
	}

	res, total, err := s.cardService.FindByActive(ctx, req)
	s.NoError(err)
	s.NotNil(res)
	s.NotZero(*total)
}

func (s *CardServiceTestSuite) Test5_UpdateCard() {
	s.Require().NotZero(s.cardID)
	ctx := context.Background()
	req := &requests.UpdateCardRequest{
		CardID:       s.cardID,
		UserID:       s.userID,
		CardType:     "credit",
		ExpireDate:   time.Now().AddDate(3, 0, 0),
		CVV:          "456",
		CardProvider: "MASTERCARD",
	}

	res, err := s.cardService.UpdateCard(ctx, req)
	s.NoError(err)
	s.NotNil(res)
	s.Equal("credit", res.CardType)
}

func (s *CardServiceTestSuite) Test6_TrashAndRestore() {
	s.Require().NotZero(s.cardID)
	ctx := context.Background()

	// Trash
	res, err := s.cardService.TrashedCard(ctx, s.cardID)
	s.NoError(err)
	s.NotNil(res)

	// Verify in trashed
	req := &requests.FindAllCards{
		Page:     1,
		PageSize: 10,
	}
	trashed, total, err := s.cardService.FindByTrashed(ctx, req)
	s.NoError(err)
	s.NotNil(trashed)
	s.NotZero(*total)

	// Restore
	restoredRes, err := s.cardService.RestoreCard(ctx, s.cardID)
	s.NoError(err)
	s.NotNil(restoredRes)
}

func (s *CardServiceTestSuite) Test7_BulkOperations() {
	ctx := context.Background()

	// Restore All
	ok, err := s.cardService.RestoreAllCard(ctx)
	s.NoError(err)
	s.True(ok)

	// Delete All Permanent
	ok, err = s.cardService.DeleteAllCardPermanent(ctx)
	s.NoError(err)
	s.True(ok)
}

func (s *CardServiceTestSuite) Test8_DeletePermanent() {
	ctx := context.Background()

	// Create another one to delete
	req := &requests.CreateCardRequest{
		UserID:       s.userID,
		CardType:     "debit",
		ExpireDate:   time.Now().AddDate(1, 0, 0),
		CVV:          "999",
		CardProvider: "BCA",
	}
	res, err := s.cardService.CreateCard(ctx, req)
	s.NoError(err)
	s.NotNil(res)

	ok, err := s.cardService.DeleteCardPermanent(ctx, int(res.CardID))
	s.NoError(err)
	s.True(ok)
}

func (s *CardServiceTestSuite) TestZZ_NotFound404() {
	ctx := context.Background()
	_, err := s.cardService.FindById(ctx, 999999999)
	s.Require().Error(err)
	s.Contains(err.Error(), "not found")
}

func TestCardServiceSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	suite.Run(t, new(CardServiceTestSuite))
}
