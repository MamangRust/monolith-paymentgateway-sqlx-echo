package topup_test

import (
	"context"
	"log"
	"net"
	"testing"
	"time"

	pbtopup "github.com/MamangRust/monolith-payment-gateway-pb/topup"
	pb "github.com/MamangRust/monolith-payment-gateway-pb/topup/stats"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/cache"
	"github.com/MamangRust/monolith-payment-gateway-topup/handler"
	"github.com/MamangRust/monolith-payment-gateway-topup/repository"
	"github.com/MamangRust/monolith-payment-gateway-topup/service"
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

type TopupStatsHandlerGapiTestSuite struct {
	suite.Suite
	ts           *tests.TestSuite
	dbPool       *sqlx.DB
	lis          *bufconn.Listener
	conn         *grpc.ClientConn
	amountClient pb.TopupStatsAmountServiceClient
	statusClient pb.TopupStatsStatusServiceClient
	userID       int32
	cardNumber1  string
	testYear     int
}

func (s *TopupStatsHandlerGapiTestSuite) SetupSuite() {
	ts, err := tests.SetupTestSuite()
	s.Require().NoError(err)
	s.ts = ts

	pool, err := sqlx.ConnectContext(s.ts.Ctx, "pgx", s.ts.DBURL)
	s.Require().NoError(err)
	s.dbPool = pool
	queries := pool
	// Stats handlers never reach the card/saldo adapters; their clients are unused.
	repos := repository.NewRepositories(queries, nil, nil, nil, nil)

	zapLog := zap.NewNop()
	myLogger := &logger.Logger{Log: zapLog}

	redisOption, err := redis.ParseURL(s.ts.RedisURL)
	s.Require().NoError(err)
	redisClient := redis.NewClient(redisOption)
	cacheStore := cache.NewCacheStore(redisClient, myLogger, &dummyCacheMetrics{})

	svc := service.NewService(&service.Deps{
		Kafka:        nil,
		Repositories: repos,
		Logger:       myLogger,
		Cache:        cacheStore,
	})

	h := handler.NewHandler(svc)

	s.lis = bufconn.Listen(1024 * 1024)
	server := grpc.NewServer()
	pb.RegisterTopupStatsAmountServiceServer(server, h)
	pb.RegisterTopupStatsStatusServiceServer(server, h)

	go func() {
		if err := server.Serve(s.lis); err != nil {
			log.Printf("Server exited with error: %v", err)
		}
	}()

	conn, err := grpc.NewClient("passthrough://bufnet",
		grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
			return s.lis.Dial()
		}),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	s.Require().NoError(err)
	s.conn = conn

	s.amountClient = pb.NewTopupStatsAmountServiceClient(conn)
	s.statusClient = pb.NewTopupStatsStatusServiceClient(conn)

	s.testYear = time.Now().Year()

	// Seed Data
	ctx := context.Background()
	err = s.dbPool.QueryRowContext(ctx, "INSERT INTO users (firstname, lastname, email, password, verification_code, is_verified) VALUES ('TopupGapi', 'Stats', 'topup_gapi_stats@example.com', 'pass', '123', true) RETURNING user_id").Scan(&s.userID)
	s.Require().NoError(err)

	s.cardNumber1 = "3333444455556666"
	_, err = s.dbPool.ExecContext(ctx, "INSERT INTO cards (user_id, card_number, card_type, cvv, card_provider, expire_date) VALUES ($1, $2, 'debit', '123', 'visa', '2030-01-01')", s.userID, s.cardNumber1)
	s.Require().NoError(err)

	s.dbPool.ExecContext(ctx, "INSERT INTO topups (card_number, topup_amount, topup_method, topup_time, status) VALUES ($1, $2, $3, $4, 'success')", s.cardNumber1, 5000, "bank_transfer", time.Date(s.testYear, 1, 10, 10, 0, 0, 0, time.UTC))
}

func (s *TopupStatsHandlerGapiTestSuite) TearDownSuite() {
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

func (s *TopupStatsHandlerGapiTestSuite) TestFindMonthlyTopupAmounts() {
	ctx := context.Background()
	resp, err := s.amountClient.FindMonthlyTopupAmounts(ctx, &pbtopup.FindYearTopupStatus{
		Year: int32(s.testYear),
	})
	s.NoError(err)
	s.Equal("success", resp.Status)
	s.NotEmpty(resp.Data)
	s.Equal(int32(5000), resp.Data[0].TotalAmount)
}

func (s *TopupStatsHandlerGapiTestSuite) TestFindMonthlyTopupAmountsByCardNumber() {
	ctx := context.Background()
	resp, err := s.amountClient.FindMonthlyTopupAmountsByCardNumber(ctx, &pbtopup.FindYearTopupCardNumber{
		Year:       int32(s.testYear),
		CardNumber: s.cardNumber1,
	})
	s.NoError(err)
	s.Equal("success", resp.Status)
	s.NotEmpty(resp.Data)
	s.Equal(int32(5000), resp.Data[0].TotalAmount)
}

func TestTopupStatsHandlerGapiSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	suite.Run(t, new(TopupStatsHandlerGapiTestSuite))
}
