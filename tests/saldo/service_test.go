package saldo_test

import (
	"context"
	"fmt"
	"net"
	"testing"
	"time"

	card_handler "github.com/MamangRust/monolith-payment-gateway-card/handler"
	card_repo "github.com/MamangRust/monolith-payment-gateway-card/repository"
	card_service "github.com/MamangRust/monolith-payment-gateway-card/service"
	pbcard "github.com/MamangRust/monolith-payment-gateway-pb/card"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-saldo/repository"
	"github.com/MamangRust/monolith-payment-gateway-saldo/service"
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
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type SaldoServiceTestSuite struct {
	suite.Suite
	ts           *tests.TestSuite
	saldoService service.Service
	dbPool       *sqlx.DB
	cardServer   *grpc.Server
	saldoID      int
	cardNumber   string
	userID       int
}

func (s *SaldoServiceTestSuite) SetupSuite() {
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

	logger.ResetInstance()
	lp := sdklog.NewLoggerProvider()
	log, _ := logger.NewLogger("test", lp)
	cacheMetrics, _ := observability.NewCacheMetrics("test")
	cacheStore := cache.NewCacheStore(redisClient, log, cacheMetrics)

	// The saldo service validates cards through the card adapter, so expose a
	// minimal card service over an in-process gRPC server. grpc.NewClient dials
	// lazily; the service is registered before it is first used.
	lis, err := net.Listen("tcp", "localhost:0")
	s.Require().NoError(err)

	cardConn, err := grpc.NewClient(lis.Addr().String(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	s.Require().NoError(err)

	cardSvc := card_service.NewService(&card_service.Deps{
		Cache:        cacheStore,
		Repositories: card_repo.NewRepositories(queries, nil),
		Logger:       log,
		Kafka:        nil,
	})
	cardH := card_handler.NewHandler(cardSvc)

	s.cardServer = grpc.NewServer()
	pbcard.RegisterCardQueryServiceServer(s.cardServer, cardH)
	pbcard.RegisterCardCommandServiceServer(s.cardServer, cardH)

	go func() {
		_ = s.cardServer.Serve(lis)
	}()

	repos := repository.NewRepositories(
		queries,
		pbcard.NewCardQueryServiceClient(cardConn),
		pbcard.NewCardCommandServiceClient(cardConn),
	)

	s.saldoService = service.NewService(&service.Deps{
		Repositories: repos,
		Logger:       log,
		Cache:        cacheStore,
	})

	// Seed User
	userRepo := user_repo.NewUserCommandRepository(queries)
	user, err := userRepo.CreateUser(context.Background(), &requests.CreateUserRequest{
		FirstName: "Saldo",
		LastName:  "Tester",
		Email:     fmt.Sprintf("saldo.tester.%d@example.com", time.Now().UnixNano()),
		Password:  "password123",
	})
	s.Require().NoError(err)
	s.userID = int(user.UserID)

	// Seed Card
	cardRepo := card_repo.NewCardCommandRepository(queries)
	card, err := cardRepo.CreateCard(context.Background(), &requests.CreateCardRequest{
		UserID:       s.userID,
		CardType:     "debit",
		ExpireDate:   time.Now().AddDate(5, 0, 0),
		CVV:          "123",
		CardProvider: "visa",
	})
	s.Require().NoError(err)
	s.cardNumber = card.CardNumber
}

func (s *SaldoServiceTestSuite) TearDownSuite() {
	if s.cardServer != nil {
		s.cardServer.Stop()
	}
	s.dbPool.Close()
	s.ts.Teardown()
}

func (s *SaldoServiceTestSuite) Test1_SaldoLifecycle() {
	ctx := context.Background()

	// Create Saldo
	createReq := &requests.CreateSaldoRequest{
		CardNumber:   s.cardNumber,
		TotalBalance: 1000000,
	}
	res, err := s.saldoService.CreateSaldo(ctx, createReq)
	s.NoError(err)
	s.NotNil(res)
	s.saldoID = int(res.SaldoID)

	// FindById
	found, err := s.saldoService.FindById(ctx, s.saldoID)
	s.NoError(err)
	s.NotNil(found)
	s.Equal(int32(s.saldoID), found.SaldoID)

	// FindByCardNumber
	foundByCard, err := s.saldoService.FindByCardNumber(ctx, s.cardNumber)
	s.NoError(err)
	s.NotNil(foundByCard)
	s.Equal(s.cardNumber, foundByCard.CardNumber)

	// Update Saldo
	updateReq := &requests.UpdateSaldoRequest{
		SaldoID:      &s.saldoID,
		CardNumber:   s.cardNumber,
		TotalBalance: 2000000,
	}
	updated, err := s.saldoService.UpdateSaldo(ctx, updateReq)
	s.NoError(err)
	s.NotNil(updated)
	s.Equal(int32(2000000), updated.TotalBalance)

	// Update Withdraw
	withdrawAmount := 500000
	now := time.Now()
	withdrawReq := &requests.UpdateSaldoWithdraw{
		CardNumber:     s.cardNumber,
		TotalBalance:   2000000,
		WithdrawAmount: &withdrawAmount,
		WithdrawTime:   &now,
	}
	withdrawn, err := s.saldoService.UpdateSaldoWithdraw(ctx, withdrawReq)
	s.NoError(err)
	s.NotNil(withdrawn)
}

func (s *SaldoServiceTestSuite) Test2_QueryOperations() {
	ctx := context.Background()

	// FindAll
	all, total, err := s.saldoService.FindAll(ctx, &requests.FindAllSaldos{
		Page:     1,
		PageSize: 10,
	})
	s.NoError(err)
	s.NotNil(all)
	s.GreaterOrEqual(*total, 1)

	// FindByActive
	active, totalActive, err := s.saldoService.FindByActive(ctx, &requests.FindAllSaldos{
		Page:     1,
		PageSize: 10,
	})
	s.NoError(err)
	s.NotNil(active)
	s.GreaterOrEqual(*totalActive, 1)
}

func (s *SaldoServiceTestSuite) Test3_TrashAndRestore() {
	ctx := context.Background()
	s.Require().NotZero(s.saldoID)

	// Trash
	trashed, err := s.saldoService.TrashSaldo(ctx, s.saldoID)
	s.NoError(err)
	s.NotNil(trashed)

	// FindByTrashed
	trashedList, totalTrashed, err := s.saldoService.FindByTrashed(ctx, &requests.FindAllSaldos{
		Page:     1,
		PageSize: 10,
	})
	s.NoError(err)
	s.NotNil(trashedList)
	s.GreaterOrEqual(*totalTrashed, 1)

	// Restore
	restored, err := s.saldoService.RestoreSaldo(ctx, s.saldoID)
	s.NoError(err)
	s.NotNil(restored)
}

func (s *SaldoServiceTestSuite) Test4_BulkOperations() {
	ctx := context.Background()

	// Restore All
	ok, err := s.saldoService.RestoreAllSaldo(ctx)
	s.NoError(err)
	s.True(ok)

	// Delete All Permanent
	ok, err = s.saldoService.DeleteAllSaldoPermanent(ctx)
	s.NoError(err)
	s.True(ok)
}

func (s *SaldoServiceTestSuite) TestZZ_NotFound404() {
	ctx := context.Background()
	_, err := s.saldoService.FindById(ctx, 999999999)
	s.Require().Error(err)
	s.Contains(err.Error(), "not found")
}

func TestSaldoServiceSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	suite.Run(t, new(SaldoServiceTestSuite))
}
