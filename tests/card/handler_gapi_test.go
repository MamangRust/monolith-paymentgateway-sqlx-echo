package card_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/MamangRust/monolith-payment-gateway-card/handler"
	"github.com/MamangRust/monolith-payment-gateway-card/repository"
	"github.com/MamangRust/monolith-payment-gateway-card/service"
	pb "github.com/MamangRust/monolith-payment-gateway-pb/card"
	pbrole "github.com/MamangRust/monolith-payment-gateway-pb/role"
	pbuser "github.com/MamangRust/monolith-payment-gateway-pb/user"
	"github.com/MamangRust/monolith-payment-gateway-pkg/hash"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/cache"
	"github.com/MamangRust/monolith-payment-gateway-shared/observability"
	user_handler "github.com/MamangRust/monolith-payment-gateway-user/handler"
	user_repository "github.com/MamangRust/monolith-payment-gateway-user/repository"
	user_service "github.com/MamangRust/monolith-payment-gateway-user/service"
	tests "github.com/MamangRust/monolith-payment-gateway-test"
	"github.com/jmoiron/sqlx"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/suite"
	sdklog "go.opentelemetry.io/otel/sdk/log"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type CardGapiTestSuite struct {
	suite.Suite
	ts       *tests.TestSuite
	dbPool   *sqlx.DB
	cardH    handler.Handler
	userH    user_handler.Handler
	userID   int32
	cardID   int32
	identity *tests.IdentityStack
}

func (s *CardGapiTestSuite) SetupSuite() {
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
	hasher := hash.NewHashingPassword()
	cacheMetrics, _ := observability.NewCacheMetrics("test")
	cacheStore := cache.NewCacheStore(redisClient, log, cacheMetrics)

	s.identity, err = tests.StartIdentityStack(pool, cacheStore, log)
	s.Require().NoError(err)

	queries := pool
	repos := repository.NewRepositories(queries, pbuser.NewUserQueryServiceClient(s.identity.User.Conn))
	userRepos := user_repository.NewRepositories(
		queries,
		pbrole.NewRoleServiceClient(s.identity.Role.Conn),
		pbrole.NewRoleCommandServiceClient(s.identity.Role.Conn),
	)

	cardSvc := service.NewService(&service.Deps{
		Cache:        cacheStore,
		Repositories: repos,
		Logger:       log,
		Kafka:        nil,
	})

	userSvc := user_service.NewService(&user_service.Deps{
		Cache:           cacheStore,
		Repositories:    userRepos,
		Hash:            hasher,
		Logger:          log,
		RoleAdapter:     s.identity.Role.RoleAdapter,
		UserRoleAdapter: s.identity.Role.UserRoleAdapter,
	})

	s.cardH = handler.NewHandler(cardSvc)
	s.userH = user_handler.NewHandler(userSvc)

	// Create a user for testing
	ctx := context.Background()
	userRes, err := s.userH.Create(ctx, &pbuser.CreateUserRequest{
		Firstname:       "Card",
		Lastname:        "User",
		Email:           fmt.Sprintf("card.user.%d.%d@example.com", time.Now().UnixNano(), time.Now().UnixNano()%10000),
		Password:        "Password123!",
		ConfirmPassword: "Password123!",
	})
	s.Require().NoError(err)
	s.userID = userRes.Data.Id
}

func (s *CardGapiTestSuite) TearDownSuite() {
	if s.identity != nil {
		s.identity.Close()
	}
	if s.dbPool != nil {
		s.dbPool.Close()
	}
	if s.ts != nil {
		s.ts.Teardown()
	}
}

func (s *CardGapiTestSuite) Test1_CardLifecycle() {
	ctx := context.Background()

	// Create
	createReq := &pb.CreateCardRequest{
		UserId:       s.userID,
		CardType:     "debit",
		ExpireDate:   timestamppb.New(time.Now().AddDate(5, 0, 0)),
		Cvv:          "123",
		CardProvider: "visa",
	}
	res, err := s.cardH.CreateCard(ctx, createReq)
	s.NoError(err)
	s.NotNil(res)
	s.cardID = res.Data.Id
	s.Equal("debit", res.Data.CardType)

	// FindById
	findReq := &pb.FindByIdCardRequest{CardId: s.cardID}
	resF, err := s.cardH.FindByIdCard(ctx, findReq)
	s.NoError(err)
	s.Equal("debit", resF.Data.CardType)

	// Update
	updateReq := &pb.UpdateCardRequest{
		CardId:       s.cardID,
		UserId:       s.userID,
		CardType:     "credit",
		ExpireDate:   createReq.ExpireDate,
		Cvv:          "456",
		CardProvider: "mastercard",
	}
	resU, err := s.cardH.UpdateCard(ctx, updateReq)
	s.NoError(err)
	s.Equal("credit", resU.Data.CardType)
}

func (s *CardGapiTestSuite) Test2_QueryOperations() {
	ctx := context.Background()
	s.Require().NotZero(s.cardID)

	// FindAll
	allReq := &pb.FindAllCardRequest{Page: 1, PageSize: 10}
	resA, err := s.cardH.FindAllCard(ctx, allReq)
	s.NoError(err)
	s.GreaterOrEqual(resA.PaginationMeta.TotalRecords, int32(1))

	// FindByActive
	resAc, err := s.cardH.FindByActiveCard(ctx, allReq)
	s.NoError(err)
	s.GreaterOrEqual(resAc.PaginationMeta.TotalRecords, int32(1))
}

func (s *CardGapiTestSuite) Test3_TrashAndRestore() {
	ctx := context.Background()
	s.Require().NotZero(s.cardID)

	// Trash
	resT, err := s.cardH.TrashedCard(ctx, &pb.FindByIdCardRequest{CardId: s.cardID})
	s.NoError(err)
	s.NotNil(resT)

	// FindByTrashed
	resTL, err := s.cardH.FindByTrashedCard(ctx, &pb.FindAllCardRequest{Page: 1, PageSize: 10})
	s.NoError(err)
	s.GreaterOrEqual(resTL.PaginationMeta.TotalRecords, int32(1))

	// Restore
	resR, err := s.cardH.RestoreCard(ctx, &pb.FindByIdCardRequest{CardId: s.cardID})
	s.NoError(err)
	s.NotNil(resR)
}

func (s *CardGapiTestSuite) Test4_BulkOperations() {
	ctx := context.Background()

	// Restore All
	resR, err := s.cardH.RestoreAllCard(ctx, &emptypb.Empty{})
	s.NoError(err)
	s.Equal("success", resR.Status)

	// Delete All Permanent
	resD, err := s.cardH.DeleteAllCardPermanent(ctx, &emptypb.Empty{})
	s.NoError(err)
	s.Equal("success", resD.Status)
}

func TestCardGapiSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	suite.Run(t, new(CardGapiTestSuite))
}
