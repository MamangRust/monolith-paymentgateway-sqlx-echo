package transfer_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	transferhandler "github.com/MamangRust/monolith-payment-gateway-apigateway/handler/transfer"
	card_repo "github.com/MamangRust/monolith-payment-gateway-card/repository"
	pb "github.com/MamangRust/monolith-payment-gateway-pb/transfer"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	saldo_repo "github.com/MamangRust/monolith-payment-gateway-saldo/repository"
	"github.com/MamangRust/monolith-payment-gateway-shared/cache"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	app_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors"
	"github.com/MamangRust/monolith-payment-gateway-shared/observability"
	tests "github.com/MamangRust/monolith-payment-gateway-test"
	"github.com/MamangRust/monolith-payment-gateway-transfer/handler"
	"github.com/MamangRust/monolith-payment-gateway-transfer/repository"
	"github.com/MamangRust/monolith-payment-gateway-transfer/service"
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

type TransferHandlerTestSuite struct {
	suite.Suite
	ts            *tests.TestSuite
	dbPool        *sqlx.DB
	redisClient   *redis.Client
	grpcServer    *grpc.Server
	commandClient pb.TransferCommandServiceClient
	queryClient   pb.TransferQueryServiceClient
	conn          *grpc.ClientConn
	router        *echo.Echo
	repos         repository.Repositories
	deps          *grpcDeps
	userRepo      user_repo.UserCommandRepository
	cardRepo      card_repo.Repositories
	saldoRepo     saldo_repo.Repositories

	senderCardNumber   string
	receiverCardNumber string
	transferID         int
}

func (s *TransferHandlerTestSuite) SetupSuite() {
	ts, err := tests.SetupTestSuite()
	s.Require().NoError(err)
	s.ts = ts

	pool, err := sqlx.ConnectContext(s.ts.Ctx, "pgx", s.ts.DBURL)
	s.Require().NoError(err)
	s.dbPool = pool

	queries := pool

	logger.ResetInstance()
	lp := sdklog.NewLoggerProvider()
	log, _ := logger.NewLogger("test", lp)
	obs, _ := observability.NewObservability("test", log)
	cacheMetrics, _ := observability.NewCacheMetrics("test")

	opts, err := redis.ParseURL(s.ts.RedisURL)
	s.Require().NoError(err)
	s.redisClient = redis.NewClient(opts)
	cacheStore := cache.NewCacheStore(s.redisClient, log, cacheMetrics)

	deps, err := startGRPCDeps(queries, cacheStore, log)
	s.Require().NoError(err)
	s.deps = deps

	// Repositories for seeding
	s.userRepo = deps.userRepo.UserCommand()
	s.cardRepo = *deps.cardRepo
	s.saldoRepo = deps.saldoRepo

	// Transfer repos
	s.repos = repository.NewRepositories(
		queries,
		deps.cardQueryClient,
		deps.cardCommandClient,
		deps.saldoQueryClient,
		deps.saldoCommandClient,
	)

	transferService := service.NewService(&service.Deps{
		Kafka:        nil,
		Repositories: s.repos,
		Logger:       log,
		Cache:        cacheStore,
	})

	// Seed Sender
	sender, err := s.userRepo.CreateUser(context.Background(), &requests.CreateUserRequest{
		FirstName: "Sender",
		LastName:  "User",
		Email:     "sender@transfer.com",
		Password:  "password123",
	})
	s.Require().NoError(err)

	sCard, err := s.cardRepo.CardCommand.CreateCard(context.Background(), &requests.CreateCardRequest{
		UserID:       int(sender.UserID),
		CardType:     "debit",
		ExpireDate:   time.Now().AddDate(1, 0, 0),
		CVV:          "111",
		CardProvider: "visa",
	})
	s.Require().NoError(err)
	s.senderCardNumber = sCard.CardNumber

	_, err = s.saldoRepo.CreateSaldo(context.Background(), &requests.CreateSaldoRequest{
		CardNumber:   s.senderCardNumber,
		TotalBalance: 1000000,
	})
	s.Require().NoError(err)

	// Seed Receiver
	receiver, err := s.userRepo.CreateUser(context.Background(), &requests.CreateUserRequest{
		FirstName: "Receiver",
		LastName:  "User",
		Email:     "receiver@transfer.com",
		Password:  "password123",
	})
	s.Require().NoError(err)

	rCard, err := s.cardRepo.CardCommand.CreateCard(context.Background(), &requests.CreateCardRequest{
		UserID:       int(receiver.UserID),
		CardType:     "debit",
		ExpireDate:   time.Now().AddDate(1, 0, 0),
		CVV:          "222",
		CardProvider: "mastercard",
	})
	s.Require().NoError(err)
	s.receiverCardNumber = rCard.CardNumber

	_, err = s.saldoRepo.CreateSaldo(context.Background(), &requests.CreateSaldoRequest{
		CardNumber:   s.receiverCardNumber,
		TotalBalance: 0,
	})
	s.Require().NoError(err)

	// Start gRPC Server
	transferHandlerGapi := handler.NewHandler(transferService)

	server := grpc.NewServer()
	pb.RegisterTransferCommandServiceServer(server, transferHandlerGapi)
	pb.RegisterTransferQueryServiceServer(server, transferHandlerGapi)
	s.grpcServer = server

	lis, err := net.Listen("tcp", ":0")
	s.Require().NoError(err)

	go func() {
		_ = server.Serve(lis)
	}()

	// Create gRPC Client
	conn, err := grpc.NewClient(lis.Addr().String(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	s.Require().NoError(err)
	s.conn = conn
	s.commandClient = pb.NewTransferCommandServiceClient(conn)
	s.queryClient = pb.NewTransferQueryServiceClient(conn)

	// Setup Echo
	s.router = echo.New()
	apiErrorHandler := app_errors.NewApiHandler(obs, log)

	transferhandler.RegisterTransferHandler(&transferhandler.DepsTransfer{
		Client:     conn,
		E:          s.router,
		Logger:     log,
		Cache:      cacheStore,
		ApiHandler: apiErrorHandler,
	})
}

func (s *TransferHandlerTestSuite) TearDownSuite() {
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

func (s *TransferHandlerTestSuite) Test1_CreateTransfer() {
	createReq := map[string]interface{}{
		"transfer_from":   s.senderCardNumber,
		"transfer_to":     s.receiverCardNumber,
		"transfer_amount": 100000,
		"transfer_time":   time.Now().Format(time.RFC3339),
	}
	body, _ := json.Marshal(createReq)

	request := httptest.NewRequest(http.MethodPost, "/api/transfer-command/create", bytes.NewBuffer(body))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	s.router.ServeHTTP(rec, request)
	s.Require().Equal(http.StatusOK, rec.Code, rec.Body.String())

	var createRes map[string]interface{}
	json.Unmarshal(rec.Body.Bytes(), &createRes)
	transferData := createRes["data"].(map[string]interface{})
	s.transferID = int(transferData["id"].(float64))

	// Verify balances
	senderSaldo, _ := s.saldoRepo.FindByCardNumber(context.Background(), s.senderCardNumber)
	s.Equal(int32(900000), senderSaldo.TotalBalance)

	receiverSaldo, _ := s.saldoRepo.FindByCardNumber(context.Background(), s.receiverCardNumber)
	s.Equal(int32(100000), receiverSaldo.TotalBalance)
}

func (s *TransferHandlerTestSuite) Test2_FindTransferById() {
	s.Require().NotZero(s.transferID)

	request := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/transfer-query/%d", s.transferID), nil)
	rec := httptest.NewRecorder()
	s.router.ServeHTTP(rec, request)
	s.Equal(http.StatusOK, rec.Code)
}

func (s *TransferHandlerTestSuite) Test3_FindAllTransfers() {
	request := httptest.NewRequest(http.MethodGet, "/api/transfer-query?page=1&page_size=10", nil)
	rec := httptest.NewRecorder()
	s.router.ServeHTTP(rec, request)
	s.Equal(http.StatusOK, rec.Code)
}

func (s *TransferHandlerTestSuite) Test4_UpdateTransfer() {
	s.Require().NotZero(s.transferID)

	updateReq := map[string]interface{}{
		"transfer_from":   s.senderCardNumber,
		"transfer_to":     s.receiverCardNumber,
		"transfer_amount": 150000, // Increase by 50000
		"transfer_time":   time.Now().Format(time.RFC3339),
	}
	updateBody, _ := json.Marshal(updateReq)
	request := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/api/transfer-command/update/%d", s.transferID), bytes.NewBuffer(updateBody))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	s.router.ServeHTTP(rec, request)
	s.Require().Equal(http.StatusOK, rec.Code, rec.Body.String())

	// Verify adjusted balances (Sender 900k - 50k = 850k, Receiver 100k + 50k = 150k)
	senderSaldo, _ := s.saldoRepo.FindByCardNumber(context.Background(), s.senderCardNumber)
	s.Equal(int32(850000), senderSaldo.TotalBalance)

	receiverSaldo, _ := s.saldoRepo.FindByCardNumber(context.Background(), s.receiverCardNumber)
	s.Equal(int32(150000), receiverSaldo.TotalBalance)
}

func (s *TransferHandlerTestSuite) Test5_TrashedTransfer() {
	s.Require().NotZero(s.transferID)

	request := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/api/transfer-command/trashed/%d", s.transferID), nil)
	rec := httptest.NewRecorder()
	s.router.ServeHTTP(rec, request)
	s.Equal(http.StatusOK, rec.Code)
}

func (s *TransferHandlerTestSuite) Test6_RestoreTransfer() {
	s.Require().NotZero(s.transferID)

	request := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/api/transfer-command/restore/%d", s.transferID), nil)
	rec := httptest.NewRecorder()
	s.router.ServeHTTP(rec, request)
	s.Equal(http.StatusOK, rec.Code)
}

func (s *TransferHandlerTestSuite) Test7_PermanentDeleteTransfer() {
	s.Require().NotZero(s.transferID)

	request := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/api/transfer-command/permanent/%d", s.transferID), nil)
	rec := httptest.NewRecorder()
	s.router.ServeHTTP(rec, request)
	s.Equal(http.StatusOK, rec.Code)
}

func TestTransferHandlerSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	suite.Run(t, new(TransferHandlerTestSuite))
}
