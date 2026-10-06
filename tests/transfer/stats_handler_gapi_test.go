package transfer_test

import (
	"context"
	"net"
	"testing"
	"time"

	pbtransfer "github.com/MamangRust/monolith-payment-gateway-pb/transfer"
	pb "github.com/MamangRust/monolith-payment-gateway-pb/transfer/stats"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/cache"
	"github.com/MamangRust/monolith-payment-gateway-transfer/handler"
	"github.com/MamangRust/monolith-payment-gateway-transfer/repository"
	"github.com/MamangRust/monolith-payment-gateway-transfer/service"
	tests "github.com/MamangRust/monolith-payment-gateway-test"
	"github.com/jmoiron/sqlx"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

type TransferStatsHandlerGapiTestSuite struct {
	suite.Suite
	ts          *tests.TestSuite
	dbPool      *sqlx.DB
	lis         *bufconn.Listener
	conn        *grpc.ClientConn
	client      pb.TransferStatsStatusServiceClient
	deps        *grpcDeps
	userID      int32
	cardNumber1 string
	cardNumber2 string
	testYear    int
	testMonth   int
}

func (s *TransferStatsHandlerGapiTestSuite) SetupSuite() {
	ts, err := tests.SetupTestSuite()
	s.Require().NoError(err)
	s.ts = ts

	pool, err := sqlx.ConnectContext(s.ts.Ctx, "pgx", s.ts.DBURL)
	s.Require().NoError(err)
	s.dbPool = pool
	queries := pool

	zapLog := zap.NewNop()
	myLogger := &logger.Logger{Log: zapLog}

	redisOption, err := redis.ParseURL(s.ts.RedisURL)
	s.Require().NoError(err)
	redisClient := redis.NewClient(redisOption)
	cacheStore := cache.NewCacheStore(redisClient, myLogger, &dummyCacheMetrics{})

	deps, err := startGRPCDeps(queries, cacheStore, myLogger)
	s.Require().NoError(err)
	s.deps = deps

	repos := repository.NewRepositories(
		queries,
		deps.cardQueryClient,
		deps.cardCommandClient,
		deps.saldoQueryClient,
		deps.saldoCommandClient,
	)

	svc := service.NewService(&service.Deps{
		Kafka:        nil,
		Repositories: repos,
		Logger:       myLogger,
		Cache:        cacheStore,
	})

	h := handler.NewHandler(svc)

	s.lis = bufconn.Listen(1024 * 1024)
	server := grpc.NewServer()
	pb.RegisterTransferStatsStatusServiceServer(server, h)

	go func() {
		if err := server.Serve(s.lis); err != nil {
		}
	}()

	conn, err := grpc.NewClient("passthrough://bufnet",
		grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
			return s.lis.Dial()
		}),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	s.Require().NoError(err)
	s.conn = conn
	s.client = pb.NewTransferStatsStatusServiceClient(conn)

	s.testYear = time.Now().Year()
	s.testMonth = int(time.Now().Month())

	ctx := context.Background()
	err = s.dbPool.QueryRowContext(ctx, "INSERT INTO users (firstname, lastname, email, password, verification_code, is_verified) VALUES ('TransferGapi', 'Stats', 'transfer_gapi_stats@example.com', 'pass', '123', true) RETURNING user_id").Scan(&s.userID)
	s.Require().NoError(err)

	s.cardNumber1 = "1111222233334444"
	s.cardNumber2 = "5555666677778888"
	_, err = s.dbPool.ExecContext(ctx, "INSERT INTO cards (user_id, card_number, card_type, cvv, card_provider, expire_date) VALUES ($1, $2, 'debit', '123', 'visa', '2030-01-01')", s.userID, s.cardNumber1)
	s.Require().NoError(err)
	_, err = s.dbPool.ExecContext(ctx, "INSERT INTO cards (user_id, card_number, card_type, cvv, card_provider, expire_date) VALUES ($1, $2, 'debit', '123', 'mastercard', '2030-01-01')", s.userID, s.cardNumber2)
	s.Require().NoError(err)

	_, err = s.dbPool.ExecContext(ctx, "INSERT INTO transfers (transfer_from, transfer_to, transfer_amount, transfer_time, status) VALUES ($1, $2, $3, $4, 'success')",
		s.cardNumber1, s.cardNumber2, 150000, time.Date(s.testYear, time.Month(s.testMonth), 10, 10, 0, 0, 0, time.UTC))
	s.Require().NoError(err)
}

func (s *TransferStatsHandlerGapiTestSuite) TearDownSuite() {
	if s.deps != nil {
		s.deps.Close()
	}
	if s.conn != nil {
		s.conn.Close()
	}
	if s.lis != nil {
		s.lis.Close()
	}
	if s.dbPool != nil {
		s.dbPool.Close()
	}
	s.ts.Teardown()
}

func (s *TransferStatsHandlerGapiTestSuite) TestFindMonthlyTransferStatusSuccess() {
	ctx := context.Background()
	res, err := s.client.FindMonthlyTransferStatusSuccess(ctx, &pbtransfer.FindMonthlyTransferStatus{
		Year:  int32(s.testYear),
		Month: int32(s.testMonth),
	})
	s.NoError(err)
	s.Equal("success", res.Status)
	s.NotEmpty(res.Data)
}

func TestTransferStatsHandlerGapiSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	suite.Run(t, new(TransferStatsHandlerGapiTestSuite))
}
