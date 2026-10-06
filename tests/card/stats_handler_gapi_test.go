package card_test

import (
	"context"
	"net"
	"testing"
	"time"

	"github.com/MamangRust/monolith-payment-gateway-card/handler"
	"github.com/MamangRust/monolith-payment-gateway-card/repository"
	"github.com/MamangRust/monolith-payment-gateway-card/service"
	pb "github.com/MamangRust/monolith-payment-gateway-pb/card"
	pbstats "github.com/MamangRust/monolith-payment-gateway-pb/card/stats"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/cache"
	"github.com/MamangRust/monolith-payment-gateway-shared/observability"
	tests "github.com/MamangRust/monolith-payment-gateway-test"
	"github.com/jmoiron/sqlx"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/suite"
	sdklog "go.opentelemetry.io/otel/sdk/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

type CardStatsGapiTestSuite struct {
	suite.Suite
	ts          *tests.TestSuite
	dbPool      *sqlx.DB
	cardH       handler.Handler
	cardNumber1 string
	cardNumber2 string
	testYear    int
	grpcServer  *grpc.Server
	lis         *bufconn.Listener
	conn        *grpc.ClientConn

	// Clients
	balanceClient  pbstats.CardStatsBalanceServiceClient
	topupClient    pbstats.CardStatsTopupServiceClient
	transferClient pbstats.CardStatsTransferServiceClient
	withdrawClient pbstats.CardStatsWithdrawServiceClient
}

func (s *CardStatsGapiTestSuite) SetupSuite() {
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
	repos := repository.NewRepositories(queries, placeholderUserQueryClient())

	logger.ResetInstance()
	lp := sdklog.NewLoggerProvider()
	log, _ := logger.NewLogger("test", lp)
	cacheMetrics, _ := observability.NewCacheMetrics("test")
	cacheStore := cache.NewCacheStore(redisClient, log, cacheMetrics)

	cardSvc := service.NewService(&service.Deps{
		Repositories: repos,
		Logger:       log,
		Cache:        cacheStore,
		Kafka:        nil,
	})
	s.cardH = handler.NewHandler(cardSvc)

	// Setup gRPC server
	s.lis = bufconn.Listen(1024 * 1024)
	s.grpcServer = grpc.NewServer()

	pb.RegisterCardQueryServiceServer(s.grpcServer, s.cardH)
	pbstats.RegisterCardStatsBalanceServiceServer(s.grpcServer, s.cardH)
	pbstats.RegisterCardStatsTopupServiceServer(s.grpcServer, s.cardH)
	pbstats.RegisterCardStatsTransferServiceServer(s.grpcServer, s.cardH)
	pbstats.RegisterCardStatsWithdrawServiceServer(s.grpcServer, s.cardH)
	pbstats.RegisterCardStatsTransactionServiceServer(s.grpcServer, s.cardH)

	go func() {
		if err := s.grpcServer.Serve(s.lis); err != nil {
		}
	}()

	s.conn, err = grpc.DialContext(context.Background(), "bufnet",
		grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
			return s.lis.Dial()
		}),
		grpc.WithInsecure())
	s.Require().NoError(err)

	s.balanceClient = pbstats.NewCardStatsBalanceServiceClient(s.conn)
	s.topupClient = pbstats.NewCardStatsTopupServiceClient(s.conn)
	s.transferClient = pbstats.NewCardStatsTransferServiceClient(s.conn)
	s.withdrawClient = pbstats.NewCardStatsWithdrawServiceClient(s.conn)

	s.testYear = time.Now().Year()

	// Seed data (same as service test)
	ctx := context.Background()
	s.cardNumber1 = "3333444455556666"
	s.cardNumber2 = "7777888899990000"

	var userID int32
	err = s.dbPool.QueryRowContext(ctx, "INSERT INTO users (firstname, lastname, email, password, verification_code, is_verified) VALUES ('Gapi', 'Stats', 'gapi_stats@example.com', 'pass', '123', true) RETURNING user_id").Scan(&userID)
	s.Require().NoError(err)

	s.dbPool.ExecContext(ctx, "INSERT INTO cards (user_id, card_number, card_type, cvv, card_provider, expire_date) VALUES ($1, $2, 'debit', '123', 'visa', '2030-01-01')", userID, s.cardNumber1)
	s.dbPool.ExecContext(ctx, "INSERT INTO cards (user_id, card_number, card_type, cvv, card_provider, expire_date) VALUES ($1, $2, 'credit', '456', 'mastercard', '2030-01-01')", userID, s.cardNumber2)

	s.seedHistoricalData()
}

func (s *CardStatsGapiTestSuite) seedHistoricalData() {
	s.insertSaldo(s.cardNumber1, 1000, time.Date(s.testYear, 1, 15, 10, 0, 0, 0, time.UTC))
	s.insertSaldo(s.cardNumber1, 2000, time.Date(s.testYear, 2, 15, 10, 0, 0, 0, time.UTC))
	s.insertSaldo(s.cardNumber2, 3000, time.Date(s.testYear, 2, 20, 10, 0, 0, 0, time.UTC))

	s.insertTopup(s.cardNumber1, 500, time.Date(s.testYear, 1, 10, 10, 0, 0, 0, time.UTC))
	s.insertTopup(s.cardNumber1, 1000, time.Date(s.testYear, 2, 10, 10, 0, 0, 0, time.UTC))
	s.insertTopup(s.cardNumber2, 1500, time.Date(s.testYear, 2, 12, 10, 0, 0, 0, time.UTC))

	s.insertWithdraw(s.cardNumber1, 100, time.Date(s.testYear, 1, 20, 10, 0, 0, 0, time.UTC))
	s.insertWithdraw(s.cardNumber1, 200, time.Date(s.testYear, 2, 20, 10, 0, 0, 0, time.UTC))
	s.insertWithdraw(s.cardNumber2, 300, time.Date(s.testYear, 2, 25, 10, 0, 0, 0, time.UTC))

	s.insertTransfer(s.cardNumber1, s.cardNumber2, 50, time.Date(s.testYear, 1, 25, 10, 0, 0, 0, time.UTC))
	s.insertTransfer(s.cardNumber1, s.cardNumber2, 100, time.Date(s.testYear, 2, 28, 10, 0, 0, 0, time.UTC))
}

func (s *CardStatsGapiTestSuite) insertSaldo(cardNumber string, amount int, t time.Time) {
	// Saldo keeps a single active snapshot per card (idx_saldos_card_number_active);
	// superseded snapshots are soft-deleted so the stats history still counts them.
	s.dbPool.ExecContext(context.Background(), "UPDATE saldos SET deleted_at = current_timestamp WHERE card_number = $1 AND deleted_at IS NULL", cardNumber)
	s.dbPool.ExecContext(context.Background(), "INSERT INTO saldos (card_number, total_balance, created_at) VALUES ($1, $2, $3)", cardNumber, amount, t)
}

func (s *CardStatsGapiTestSuite) insertTopup(cardNumber string, amount int, t time.Time) {
	s.dbPool.ExecContext(context.Background(), "INSERT INTO topups (card_number, topup_amount, topup_time, topup_method, status) VALUES ($1, $2, $3, 'pb', 'success')", cardNumber, amount, t)
}

func (s *CardStatsGapiTestSuite) insertWithdraw(cardNumber string, amount int, t time.Time) {
	s.dbPool.ExecContext(context.Background(), "INSERT INTO withdraws (card_number, withdraw_amount, withdraw_time, status) VALUES ($1, $2, $3, 'success')", cardNumber, amount, t)
}

func (s *CardStatsGapiTestSuite) insertTransfer(from, to string, amount int, t time.Time) {
	s.dbPool.ExecContext(context.Background(), "INSERT INTO transfers (transfer_from, transfer_to, transfer_amount, transfer_time, status) VALUES ($1, $2, $3, $4, 'success')", from, to, amount, t)
}

func (s *CardStatsGapiTestSuite) TearDownSuite() {
	if s.conn != nil {
		s.conn.Close()
	}
	if s.grpcServer != nil {
		s.grpcServer.Stop()
	}
	if s.dbPool != nil {
		s.dbPool.Close()
	}
	if s.ts != nil {
		s.ts.Teardown()
	}
}

func (s *CardStatsGapiTestSuite) TestBalanceGapi() {
	ctx := context.Background()

	// Global Monthly
	res, err := s.balanceClient.FindMonthlyBalance(ctx, &pbstats.FindYearBalance{Year: int32(s.testYear)})
	s.NoError(err)
	s.Equal("success", res.Status)
	s.Equal(int64(1000), res.Data[0].TotalBalance) // Jan
	s.Equal(int64(5000), res.Data[1].TotalBalance) // Feb

	// By Card Monthly
	resByCard, err := s.balanceClient.FindMonthlyBalanceByCardNumber(ctx, &pbstats.FindYearBalanceCardNumber{CardNumber: s.cardNumber1, Year: int32(s.testYear)})
	s.NoError(err)
	s.Equal(int64(1000), resByCard.Data[0].TotalBalance) // Jan Card 1
	s.Equal(int64(2000), resByCard.Data[1].TotalBalance) // Feb Card 1
}

func (s *CardStatsGapiTestSuite) TestTopupGapi() {
	ctx := context.Background()

	res, err := s.topupClient.FindMonthlyTopupAmount(ctx, &pb.FindYearAmount{Year: int32(s.testYear)})
	s.NoError(err)
	s.Equal(int64(500), res.Data[0].TotalAmount)
	s.Equal(int64(2500), res.Data[1].TotalAmount)

	resByCard, err := s.topupClient.FindMonthlyTopupAmountByCardNumber(ctx, &pb.FindYearAmountCardNumber{CardNumber: s.cardNumber1, Year: int32(s.testYear)})
	s.NoError(err)
	s.Equal(int64(500), resByCard.Data[0].TotalAmount)
}

func TestCardStatsGapiSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	suite.Run(t, new(CardStatsGapiTestSuite))
}
