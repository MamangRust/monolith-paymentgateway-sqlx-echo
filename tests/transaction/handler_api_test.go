package transaction_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	api_transaction "github.com/MamangRust/monolith-payment-gateway-apigateway/handler/transaction"
	mencache "github.com/MamangRust/monolith-payment-gateway-apigateway/redis"
	card_repo "github.com/MamangRust/monolith-payment-gateway-card/repository"
	merchant_repo "github.com/MamangRust/monolith-payment-gateway-merchant/repository"
	pb_merchant "github.com/MamangRust/monolith-payment-gateway-pb/merchant"
	pb "github.com/MamangRust/monolith-payment-gateway-pb/transaction"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	saldo_repo "github.com/MamangRust/monolith-payment-gateway-saldo/repository"
	"github.com/MamangRust/monolith-payment-gateway-shared/cache"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
	"github.com/MamangRust/monolith-payment-gateway-shared/observability"
	tests "github.com/MamangRust/monolith-payment-gateway-test"
	"github.com/MamangRust/monolith-payment-gateway-transaction/handler"
	"github.com/MamangRust/monolith-payment-gateway-transaction/repository"
	"github.com/MamangRust/monolith-payment-gateway-transaction/service"
	user_repo "github.com/MamangRust/monolith-payment-gateway-user/repository"
	"github.com/jmoiron/sqlx"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/suite"
	sdklog "go.opentelemetry.io/otel/sdk/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type TransactionHandlerTestSuite struct {
	suite.Suite
	ts             *tests.TestSuite
	dbPool         *sqlx.DB
	redisClient    *redis.Client
	grpcServer     *grpc.Server
	commandClient  pb.TransactionCommandServiceClient
	queryClient    pb.TransactionQueryServiceClient
	merchantClient pb_merchant.MerchantCommandServiceClient
	conn           *grpc.ClientConn
	router         *echo.Echo
	deps           *grpcDeps

	// Repositories for seeding
	userRepo     user_repo.UserCommandRepository
	cardRepo     card_repo.Repositories
	saldoRepo    saldo_repo.Repositories
	merchantRepo merchant_repo.Repositories

	customerCardNumber string
	merchantApiKey     string
	merchantID         int
	merchantCardNumber string
	transactionID      int
}

func (s *TransactionHandlerTestSuite) SetupSuite() {
	ts, err := tests.SetupTestSuite()
	s.Require().NoError(err)
	s.ts = ts

	pool, err := sqlx.ConnectContext(s.ts.Ctx, "pgx", s.ts.DBURL)
	s.Require().NoError(err)
	s.dbPool = pool

	queries := pool

	opts, err := redis.ParseURL(s.ts.RedisURL)
	s.Require().NoError(err)
	s.redisClient = redis.NewClient(opts)

	logger.ResetInstance()
	lp := sdklog.NewLoggerProvider()
	log, _ := logger.NewLogger("test", lp)
	obs, _ := observability.NewObservability("test", log)
	cacheMetrics, _ := observability.NewCacheMetrics("test")
	cacheStore := cache.NewCacheStore(s.redisClient, log, cacheMetrics)

	deps, err := startGRPCDeps(queries, cacheStore, log)
	s.Require().NoError(err)
	s.deps = deps

	// Repositories for seeding
	s.userRepo = deps.userRepo.UserCommand()
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

	// Seed Customer
	customer, err := s.userRepo.CreateUser(context.Background(), &requests.CreateUserRequest{
		FirstName: "Transaction",
		LastName:  "Customer",
		Email:     "customer@transaction.com",
		Password:  "password123",
	})
	s.Require().NoError(err)

	cCard, err := s.cardRepo.CardCommand.CreateCard(context.Background(), &requests.CreateCardRequest{
		UserID:       int(customer.UserID),
		CardType:     "debit",
		ExpireDate:   time.Now().AddDate(1, 0, 0),
		CVV:          "123",
		CardProvider: "visa",
	})
	s.Require().NoError(err)
	s.customerCardNumber = cCard.CardNumber

	_, err = s.saldoRepo.CreateSaldo(context.Background(), &requests.CreateSaldoRequest{
		CardNumber:   s.customerCardNumber,
		TotalBalance: 1000000,
	})
	s.Require().NoError(err)

	// Seed Merchant
	owner, err := s.userRepo.CreateUser(context.Background(), &requests.CreateUserRequest{
		FirstName: "Merchant",
		LastName:  "Owner",
		Email:     "merchant.owner@transaction.com",
		Password:  "password123",
	})
	s.Require().NoError(err)

	merchant, err := s.merchantRepo.CreateMerchant(context.Background(), &requests.CreateMerchantRequest{
		UserID: int(owner.UserID),
		Name:   "Transaction Merchant",
	})
	s.Require().NoError(err)
	s.merchantID = int(merchant.MerchantID)

	_, err = s.merchantRepo.UpdateMerchantStatus(context.Background(), &requests.UpdateMerchantStatusRequest{
		MerchantID: &s.merchantID,
		Status:     "active",
	})
	s.Require().NoError(err)

	mFull, _ := s.merchantRepo.FindByMerchantId(context.Background(), s.merchantID)
	s.merchantApiKey = mFull.ApiKey

	mCard, err := s.cardRepo.CardCommand.CreateCard(context.Background(), &requests.CreateCardRequest{
		UserID:       int(owner.UserID),
		CardType:     "debit",
		ExpireDate:   time.Now().AddDate(1, 0, 0),
		CVV:          "321",
		CardProvider: "mastercard",
	})
	s.Require().NoError(err)
	s.merchantCardNumber = mCard.CardNumber

	_, err = s.saldoRepo.CreateSaldo(context.Background(), &requests.CreateSaldoRequest{
		CardNumber:   s.merchantCardNumber,
		TotalBalance: 0,
	})
	s.Require().NoError(err)

	// Start gRPC Server
	transactionHandlerGapi := handler.NewHandler(transactionService)

	server := grpc.NewServer()
	pb.RegisterTransactionCommandServiceServer(server, transactionHandlerGapi)
	pb.RegisterTransactionQueryServiceServer(server, transactionHandlerGapi)
	s.grpcServer = server

	lis, err := net.Listen("tcp", ":0")
	s.Require().NoError(err)
	go func() { _ = server.Serve(lis) }()

	conn, err := grpc.NewClient(lis.Addr().String(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	s.Require().NoError(err)
	s.conn = conn
	s.commandClient = pb.NewTransactionCommandServiceClient(conn)
	s.queryClient = pb.NewTransactionQueryServiceClient(conn)
	s.merchantClient = pb_merchant.NewMerchantCommandServiceClient(conn)

	// Setup Echo
	s.router = echo.New()
	apiErrorHandler := errors.NewApiHandler(obs, log)

	// Seed Merchant Cache to bypass Kafka validation
	merchantCache := mencache.NewMerchantCache(cacheStore)
	merchantCache.SetMerchantCache(context.Background(), strconv.Itoa(s.merchantID), s.merchantApiKey)

	// Use Refactored Handlers
	api_transaction.RegisterTransactionHandler(&api_transaction.DepsTransaction{
		Client:          conn,
		E:               s.router,
		Kafka:           nil,
		Logger:          log,
		Cache:           cacheStore,
		ApiHandler:      apiErrorHandler,
		CacheApiGateway: mencache.NewCacheApiGateway(cacheStore),
	})
}

func (s *TransactionHandlerTestSuite) TearDownSuite() {
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

func (s *TransactionHandlerTestSuite) Test1_CreateTransaction() {
	createReq := map[string]interface{}{
		"card_number":      s.customerCardNumber,
		"amount":           50000,
		"payment_method":   "visa",
		"merchant_id":      s.merchantID,
		"transaction_time": time.Now().Format(time.RFC3339),
	}
	body, _ := json.Marshal(createReq)

	request := httptest.NewRequest(http.MethodPost, "/api/transaction-command/create", bytes.NewBuffer(body))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	request.Header.Set("X-API-Key", s.merchantApiKey)
	rec := httptest.NewRecorder()

	s.router.ServeHTTP(rec, request)
	s.Require().Equal(http.StatusOK, rec.Code, rec.Body.String())

	var res map[string]interface{}
	json.Unmarshal(rec.Body.Bytes(), &res)
	data := res["data"].(map[string]interface{})
	s.transactionID = int(data["id"].(float64))

	// Verify balances
	customerSaldo, _ := s.saldoRepo.FindByCardNumber(context.Background(), s.customerCardNumber)
	s.Equal(int32(950000), customerSaldo.TotalBalance)

	merchantSaldo, _ := s.saldoRepo.FindByCardNumber(context.Background(), s.merchantCardNumber)
	s.Equal(int32(50000), merchantSaldo.TotalBalance)
}

func (s *TransactionHandlerTestSuite) Test2_FindTransactionById() {
	s.Require().NotZero(s.transactionID)
	request := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/transaction-query/%d", s.transactionID), nil)
	rec := httptest.NewRecorder()

	s.router.ServeHTTP(rec, request)
	s.Require().Equal(http.StatusOK, rec.Code)
}

func (s *TransactionHandlerTestSuite) Test3_FindAllTransactions() {
	request := httptest.NewRequest(http.MethodGet, "/api/transaction-query?page=1&page_size=10", nil)
	rec := httptest.NewRecorder()

	s.router.ServeHTTP(rec, request)
	s.Require().Equal(http.StatusOK, rec.Code)
}

func (s *TransactionHandlerTestSuite) Test4_UpdateTransaction() {
	s.Require().NotZero(s.transactionID)
	updateReq := map[string]interface{}{
		"card_number":      s.customerCardNumber,
		"amount":           60000, // Increase by 10000
		"payment_method":   "visa",
		"merchant_id":      s.merchantID,
		"transaction_time": time.Now().Format(time.RFC3339),
	}
	body, _ := json.Marshal(updateReq)

	request := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/api/transaction-command/update/%d", s.transactionID), bytes.NewBuffer(body))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	request.Header.Set("X-API-Key", s.merchantApiKey)
	rec := httptest.NewRecorder()

	s.router.ServeHTTP(rec, request)
	s.Require().Equal(http.StatusOK, rec.Code, rec.Body.String())

	// Verify adjusted balance (950k - 10k = 940k)
	customerSaldo, _ := s.saldoRepo.FindByCardNumber(context.Background(), s.customerCardNumber)
	s.Equal(int32(940000), customerSaldo.TotalBalance)
}

func (s *TransactionHandlerTestSuite) Test5_TrashedTransaction() {
	s.Require().NotZero(s.transactionID)
	request := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/api/transaction-command/trashed/%d", s.transactionID), nil)
	rec := httptest.NewRecorder()

	s.router.ServeHTTP(rec, request)
	s.Require().Equal(http.StatusOK, rec.Code)
}

func (s *TransactionHandlerTestSuite) Test6_RestoreTransaction() {
	s.Require().NotZero(s.transactionID)
	request := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/api/transaction-command/restore/%d", s.transactionID), nil)
	rec := httptest.NewRecorder()

	s.router.ServeHTTP(rec, request)
	s.Require().Equal(http.StatusOK, rec.Code)
}

func (s *TransactionHandlerTestSuite) Test7_PermanentDeleteTransaction() {
	s.Require().NotZero(s.transactionID)
	request := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/api/transaction-command/permanent/%d", s.transactionID), nil)
	rec := httptest.NewRecorder()

	s.router.ServeHTTP(rec, request)
	s.Require().Equal(http.StatusOK, rec.Code)
}

func TestTransactionHandlerSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	suite.Run(t, new(TransactionHandlerTestSuite))
}
