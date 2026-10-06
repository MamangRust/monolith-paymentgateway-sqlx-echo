package transaction_test

import (
	"context"
	"net"
	"testing"
	"time"

	card_repo "github.com/MamangRust/monolith-payment-gateway-card/repository"
	merchant_repo "github.com/MamangRust/monolith-payment-gateway-merchant/repository"
	pb "github.com/MamangRust/monolith-payment-gateway-pb/transaction"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	saldo_repo "github.com/MamangRust/monolith-payment-gateway-saldo/repository"
	"github.com/MamangRust/monolith-payment-gateway-shared/cache"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/observability"
	"github.com/MamangRust/monolith-payment-gateway-transaction/handler"
	"github.com/MamangRust/monolith-payment-gateway-transaction/repository"
	"github.com/MamangRust/monolith-payment-gateway-transaction/service"
	user_repo "github.com/MamangRust/monolith-payment-gateway-user/repository"
	tests "github.com/MamangRust/monolith-payment-gateway-test"

	"github.com/jmoiron/sqlx"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/suite"
	sdklog "go.opentelemetry.io/otel/sdk/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type TransactionGapiTestSuite struct {
	suite.Suite
	ts            *tests.TestSuite
	dbPool        *sqlx.DB
	redisClient   *redis.Client
	grpcServer    *grpc.Server
	commandClient pb.TransactionCommandServiceClient
	queryClient   pb.TransactionQueryServiceClient
	conn          *grpc.ClientConn
	deps          *grpcDeps

	// Repositories for seeding
	userRepo     user_repo.UserCommandRepository
	cardRepo     card_repo.Repositories
	saldoRepo    saldo_repo.Repositories
	merchantRepo merchant_repo.Repositories

	customerCardNumber string
	merchantID         int32
	merchantApiKey     string
	transactionID      int32
}

func (s *TransactionGapiTestSuite) SetupSuite() {
	ts, err := tests.SetupTestSuite()
	s.Require().NoError(err)
	s.ts = ts

	pool, err := sqlx.ConnectContext(s.ts.Ctx, "pgx", s.ts.DBURL)
	s.Require().NoError(err)
	s.dbPool = pool

	queries := pool
	s.userRepo = user_repo.NewUserCommandRepository(queries)

	opts, err := redis.ParseURL(s.ts.RedisURL)
	s.Require().NoError(err)
	s.redisClient = redis.NewClient(opts)

	logger.ResetInstance()
	lp := sdklog.NewLoggerProvider()
	log, _ := logger.NewLogger("test", lp)
	_ = log
	cacheMetrics, _ := observability.NewCacheMetrics("test")
	cacheStore := cache.NewCacheStore(s.redisClient, log, cacheMetrics)

	deps, err := startGRPCDeps(queries, cacheStore, log)
	s.Require().NoError(err)
	s.deps = deps
	s.cardRepo = *deps.cardRepo
	s.saldoRepo = deps.saldoRepo
	s.merchantRepo = deps.merchantRepo

	transactionRepos := repository.NewRepositories(
		queries,
		deps.cardQueryClient,
		deps.cardCommandClient,
		deps.saldoQueryClient,
		deps.saldoCommandClient,
		deps.merchantQueryClient,
	)
	transactionService := service.NewService(&service.Deps{
		Kafka:        nil,
		Repositories: transactionRepos,
		Logger:       log,
		Cache:        cacheStore,
	})

	transactionHandler := handler.NewHandler(transactionService)
	server := grpc.NewServer()
	pb.RegisterTransactionCommandServiceServer(server, transactionHandler)
	pb.RegisterTransactionQueryServiceServer(server, transactionHandler)
	s.grpcServer = server

	lis, err := net.Listen("tcp", ":0")
	s.Require().NoError(err)
	go func() { _ = server.Serve(lis) }()

	conn, err := grpc.NewClient(lis.Addr().String(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	s.Require().NoError(err)
	s.conn = conn
	s.commandClient = pb.NewTransactionCommandServiceClient(conn)
	s.queryClient = pb.NewTransactionQueryServiceClient(conn)

	// Seed Customer User, Card, Saldo
	user, err := s.userRepo.CreateUser(context.Background(), &requests.CreateUserRequest{
		FirstName: "Trans", LastName: "Gapi", Email: "trans.gapi@test.com", Password: "password123",
	})
	s.Require().NoError(err)

	card, err := s.cardRepo.CardCommand.CreateCard(context.Background(), &requests.CreateCardRequest{
		UserID: int(user.UserID), CardType: "debit", ExpireDate: time.Now().AddDate(1, 0, 0), CVV: "123", CardProvider: "visa",
	})
	s.Require().NoError(err)
	s.customerCardNumber = card.CardNumber

	_, err = s.saldoRepo.CreateSaldo(context.Background(), &requests.CreateSaldoRequest{
		CardNumber: s.customerCardNumber, TotalBalance: 1000000,
	})
	s.Require().NoError(err)

	// Seed a separate Merchant Owner user with its own card + saldo. The atomic
	// settlement SQL rejects identical customer/merchant accounts ($1 <> $3), so the
	// merchant must NOT belong to the payer's user (merchant card = GetCardByUserID).
	merchantOwner, err := s.userRepo.CreateUser(context.Background(), &requests.CreateUserRequest{
		FirstName: "Merchant", LastName: "Owner", Email: "merchant.owner.gapi@test.com", Password: "password123",
	})
	s.Require().NoError(err)

	merchantCard, err := s.cardRepo.CardCommand.CreateCard(context.Background(), &requests.CreateCardRequest{
		UserID: int(merchantOwner.UserID), CardType: "debit", ExpireDate: time.Now().AddDate(1, 0, 0), CVV: "123", CardProvider: "visa",
	})
	s.Require().NoError(err)

	_, err = s.saldoRepo.CreateSaldo(context.Background(), &requests.CreateSaldoRequest{
		CardNumber: merchantCard.CardNumber, TotalBalance: 1000000,
	})
	s.Require().NoError(err)

	merchant, err := s.merchantRepo.CreateMerchant(context.Background(), &requests.CreateMerchantRequest{
		UserID: int(merchantOwner.UserID), Name: "Gapi Merchant",
	})
	s.Require().NoError(err)
	s.merchantID = merchant.MerchantID
	s.merchantApiKey = merchant.ApiKey
}

func (s *TransactionGapiTestSuite) TearDownSuite() {
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

func (s *TransactionGapiTestSuite) Test1_CreateTransaction() {
	ctx := context.Background()
	createReq := &pb.CreateTransactionRequest{
		ApiKey:          s.merchantApiKey,
		CardNumber:      s.customerCardNumber,
		Amount:          100000,
		PaymentMethod:   "visa",
		MerchantId:      s.merchantID,
		TransactionTime: timestamppb.New(time.Now()),
	}
	res, err := s.commandClient.CreateTransaction(ctx, createReq)
	s.Require().NoError(err)
	s.Equal("success", res.Status)
	s.transactionID = res.Data.Id
}

func (s *TransactionGapiTestSuite) Test2_FindTransactionById() {
	ctx := context.Background()
	s.Require().NotZero(s.transactionID)
	res, err := s.queryClient.FindByIdTransaction(ctx, &pb.FindByIdTransactionRequest{TransactionId: s.transactionID})
	s.Require().NoError(err)
	s.Equal(s.transactionID, res.Data.Id)
}

func (s *TransactionGapiTestSuite) Test3_UpdateTransaction() {
	ctx := context.Background()
	s.Require().NotZero(s.transactionID)
	updateReq := &pb.UpdateTransactionRequest{
		ApiKey:          s.merchantApiKey,
		TransactionId:   s.transactionID,
		CardNumber:      s.customerCardNumber,
		Amount:          200000,
		PaymentMethod:   "visa",
		MerchantId:      s.merchantID,
		TransactionTime: timestamppb.New(time.Now()),
	}
	res, err := s.commandClient.UpdateTransaction(ctx, updateReq)
	s.Require().NoError(err)
	s.Equal(int32(200000), res.Data.Amount)
}

func (s *TransactionGapiTestSuite) Test4_TrashedTransaction() {
	ctx := context.Background()
	s.Require().NotZero(s.transactionID)
	res, err := s.commandClient.TrashedTransaction(ctx, &pb.FindByIdTransactionRequest{TransactionId: s.transactionID})
	s.Require().NoError(err)
	s.Equal("success", res.Status)
}

func (s *TransactionGapiTestSuite) Test5_RestoreTransaction() {
	ctx := context.Background()
	s.Require().NotZero(s.transactionID)
	res, err := s.commandClient.RestoreTransaction(ctx, &pb.FindByIdTransactionRequest{TransactionId: s.transactionID})
	s.Require().NoError(err)
	s.Equal("success", res.Status)
}

func (s *TransactionGapiTestSuite) Test6_PermanentDeleteTransaction() {
	ctx := context.Background()
	s.Require().NotZero(s.transactionID)
	res, err := s.commandClient.DeleteTransactionPermanent(ctx, &pb.FindByIdTransactionRequest{TransactionId: s.transactionID})
	s.Require().NoError(err)
	s.Equal("success", res.Status)
}

func TestTransactionGapiSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	suite.Run(t, new(TransactionGapiTestSuite))
}
