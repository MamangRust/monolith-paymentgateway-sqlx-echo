package withdraw_test

import (
	"context"
	card_repo "github.com/MamangRust/monolith-payment-gateway-card/repository"
	pb "github.com/MamangRust/monolith-payment-gateway-pb/withdraw"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	saldo_repo "github.com/MamangRust/monolith-payment-gateway-saldo/repository"
	"github.com/MamangRust/monolith-payment-gateway-shared/cache"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/observability"
	user_repo "github.com/MamangRust/monolith-payment-gateway-user/repository"
	"github.com/MamangRust/monolith-payment-gateway-withdraw/handler"
	"github.com/MamangRust/monolith-payment-gateway-withdraw/repository"
	"github.com/MamangRust/monolith-payment-gateway-withdraw/service"
	"github.com/MamangRust/monolith-payment-gateway-test"
	"net"
	"testing"
	"time"

	"github.com/jmoiron/sqlx"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/suite"
	sdklog "go.opentelemetry.io/otel/sdk/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type WithdrawGapiTestSuite struct {
	suite.Suite
	ts            *tests.TestSuite
	dbPool        *sqlx.DB
	redisClient   *redis.Client
	grpcServer    *grpc.Server
	commandClient pb.WithdrawCommandServiceClient
	queryClient   pb.WithdrawQueryServiceClient
	conn          *grpc.ClientConn
	repos         repository.Repositories
	userRepo      user_repo.UserCommandRepository
	cardRepo      card_repo.CardCommandRepository
	saldoRepo     saldo_repo.Repositories
	deps          *grpcDeps

	cardNumber string
	withdrawID int32
}

func (s *WithdrawGapiTestSuite) SetupSuite() {
	ts, err := tests.SetupTestSuite()
	s.Require().NoError(err)
	s.ts = ts

	pool, err := sqlx.ConnectContext(s.ts.Ctx, "pgx", s.ts.DBURL)
	s.Require().NoError(err)
	s.dbPool = pool

	opts, err := redis.ParseURL(s.ts.RedisURL)
	s.Require().NoError(err)
	s.redisClient = redis.NewClient(opts)

	queries := pool

	logger.ResetInstance()
	lp := sdklog.NewLoggerProvider()
	log, _ := logger.NewLogger("test", lp)
	cacheMetrics, _ := observability.NewCacheMetrics("test")
	cacheStore := cache.NewCacheStore(s.redisClient, log, cacheMetrics)

	deps, err := startGRPCDeps(queries, cacheStore, log)
	s.Require().NoError(err)
	s.deps = deps

	// Repositories for seeding and service dependencies
	s.userRepo = user_repo.NewUserCommandRepository(queries)
	s.cardRepo = deps.cardRepo.CardCommand
	s.saldoRepo = deps.saldoRepo

	s.repos = repository.NewRepositories(
		queries,
		deps.cardQueryClient,
		deps.cardCommandClient,
		deps.saldoQueryClient,
		deps.saldoCommandClient,
		testDailyWithdrawLimit,
	)

	withdrawService := service.NewService(&service.Deps{
		Kafka:        nil,
		Repositories: s.repos,
		Logger:       log,
		Cache:        cacheStore,
	})

	// Seed User, Card, Saldo
	user, _ := s.userRepo.CreateUser(context.Background(), &requests.CreateUserRequest{
		FirstName: "Withdraw", LastName: "Gapi", Email: "withdraw.gapi@test.com", Password: "password123",
	})
	card, _ := s.cardRepo.CreateCard(context.Background(), &requests.CreateCardRequest{
		UserID: int(user.UserID), CardType: "debit", ExpireDate: time.Now().AddDate(1, 0, 0), CVV: "999", CardProvider: "visa",
	})
	s.cardNumber = card.CardNumber
	s.saldoRepo.CreateSaldo(context.Background(), &requests.CreateSaldoRequest{
		CardNumber: s.cardNumber, TotalBalance: 1000000,
	})

	// Start gRPC Server
	withdrawHandler := handler.NewHandler(withdrawService)
	server := grpc.NewServer()
	pb.RegisterWithdrawCommandServiceServer(server, withdrawHandler)
	pb.RegisterWithdrawQueryServiceServer(server, withdrawHandler)
	s.grpcServer = server

	lis, err := net.Listen("tcp", ":0")
	s.Require().NoError(err)
	go func() { _ = server.Serve(lis) }()

	conn, err := grpc.NewClient(lis.Addr().String(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	s.Require().NoError(err)
	s.conn = conn
	s.commandClient = pb.NewWithdrawCommandServiceClient(conn)
	s.queryClient = pb.NewWithdrawQueryServiceClient(conn)
}

func (s *WithdrawGapiTestSuite) TearDownSuite() {
	if s.deps != nil {
		s.deps.Close()
	}
	if s.conn != nil {
		s.conn.Close()
	}
	if s.grpcServer != nil {
		s.grpcServer.Stop()
	}
	if s.redisClient != nil {
		s.redisClient.Close()
	}
	if s.dbPool != nil {
		s.dbPool.Close()
	}
	if s.ts != nil {
		s.ts.Teardown()
	}
}

func (s *WithdrawGapiTestSuite) Test1_Create() {
	ctx := context.Background()

	createReq := &pb.CreateWithdrawRequest{
		CardNumber:     s.cardNumber,
		WithdrawAmount: 100000,
		WithdrawTime:   timestamppb.New(time.Now()),
	}
	res, err := s.commandClient.CreateWithdraw(ctx, createReq)
	s.NoError(err)
	s.Equal(int32(100000), res.Data.WithdrawAmount)

	s.withdrawID = res.Data.WithdrawId

	// Verify balance
	saldo, _ := s.saldoRepo.FindByCardNumber(ctx, s.cardNumber)
	s.Equal(int32(900000), saldo.TotalBalance)
}

func (s *WithdrawGapiTestSuite) Test2_FindById() {
	s.Require().NotZero(s.withdrawID)

	ctx := context.Background()
	found, err := s.queryClient.FindByIdWithdraw(ctx, &pb.FindByIdWithdrawRequest{WithdrawId: s.withdrawID})
	s.NoError(err)
	s.Equal(s.withdrawID, found.Data.WithdrawId)
}

func (s *WithdrawGapiTestSuite) Test3_Update() {
	s.Require().NotZero(s.withdrawID)

	ctx := context.Background()
	updateReq := &pb.UpdateWithdrawRequest{
		WithdrawId:     s.withdrawID,
		CardNumber:     s.cardNumber,
		WithdrawAmount: 150000,
		WithdrawTime:   timestamppb.New(time.Now()),
	}
	updated, err := s.commandClient.UpdateWithdraw(ctx, updateReq)
	s.NoError(err)
	s.Equal(int32(150000), updated.Data.WithdrawAmount)

	// Verify adjusted balance (900k - 50k = 850k)
	saldo, _ := s.saldoRepo.FindByCardNumber(ctx, s.cardNumber)
	s.Equal(int32(850000), saldo.TotalBalance)
}

func (s *WithdrawGapiTestSuite) Test4_Trashed() {
	s.Require().NotZero(s.withdrawID)

	ctx := context.Background()
	_, err := s.commandClient.TrashedWithdraw(ctx, &pb.FindByIdWithdrawRequest{WithdrawId: s.withdrawID})
	s.NoError(err)
}

func (s *WithdrawGapiTestSuite) Test5_Restore() {
	s.Require().NotZero(s.withdrawID)

	ctx := context.Background()
	_, err := s.commandClient.RestoreWithdraw(ctx, &pb.FindByIdWithdrawRequest{WithdrawId: s.withdrawID})
	s.NoError(err)
}

func (s *WithdrawGapiTestSuite) Test6_PermanentDelete() {
	s.Require().NotZero(s.withdrawID)

	ctx := context.Background()
	_, err := s.commandClient.DeleteWithdrawPermanent(ctx, &pb.FindByIdWithdrawRequest{WithdrawId: s.withdrawID})
	s.NoError(err)
}

func TestWithdrawGapiSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	suite.Run(t, new(WithdrawGapiTestSuite))
}
