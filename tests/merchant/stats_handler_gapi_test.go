package merchant_test

import (
	"context"
	"log"
	"net"
	"testing"
	"time"

	"github.com/MamangRust/monolith-payment-gateway-merchant/handler/stats"
	stats_cache "github.com/MamangRust/monolith-payment-gateway-merchant/redis/stats"
	apikey_cache "github.com/MamangRust/monolith-payment-gateway-merchant/redis/statsbyapikey"
	merchant_cache "github.com/MamangRust/monolith-payment-gateway-merchant/redis/statsbymerchant"
	stats_repo "github.com/MamangRust/monolith-payment-gateway-merchant/repository/stats"
	apikey_repo "github.com/MamangRust/monolith-payment-gateway-merchant/repository/statsbyapikey"
	merchant_repo "github.com/MamangRust/monolith-payment-gateway-merchant/repository/statsbymerchant"
	stats_service "github.com/MamangRust/monolith-payment-gateway-merchant/service/stats"
	apikey_service "github.com/MamangRust/monolith-payment-gateway-merchant/service/statsbyapikey"
	merchant_service "github.com/MamangRust/monolith-payment-gateway-merchant/service/statsbymerchant"
	pb_merchant "github.com/MamangRust/monolith-payment-gateway-pb/merchant"
	pb "github.com/MamangRust/monolith-payment-gateway-pb/merchant/stats"
	logger_pkg "github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/cache"
	"github.com/MamangRust/monolith-payment-gateway-shared/observability"
	tests "github.com/MamangRust/monolith-payment-gateway-test"
	"github.com/jmoiron/sqlx"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/suite"
	sdklog "go.opentelemetry.io/otel/sdk/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

type MerchantStatsHandlerGapiTestSuite struct {
	suite.Suite
	ts         *tests.TestSuite
	dbPool     *sqlx.DB
	lis        *bufconn.Listener
	conn       *grpc.ClientConn
	merchantID int32
	apiKey     string
	testYear   int
}

func (s *MerchantStatsHandlerGapiTestSuite) SetupSuite() {
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

	// Logger & Observability
	logger_pkg.ResetInstance()
	lp := sdklog.NewLoggerProvider()
	logObj, _ := logger_pkg.NewLogger("test", lp)
	obs, _ := observability.NewObservability("test", logObj)
	cacheMetrics, _ := observability.NewCacheMetrics("test")
	cacheStore := cache.NewCacheStore(redisClient, logObj, cacheMetrics)

	// Dependency Injection
	s_cache := stats_cache.NewMerchantStatsCache(cacheStore)
	a_cache := apikey_cache.NewMerchantStatsByApiKeyCache(cacheStore)
	m_cache := merchant_cache.NewMerchantStatsByMerchantCache(cacheStore)

	s_repo := stats_repo.NewMerchantStatsRepository(queries)
	a_repo := apikey_repo.NewMerchantStatsByApiKeyRepository(queries)
	m_repo := merchant_repo.NewMerchantStatsByMerchantRepository(queries)

	svc := stats_service.NewMerchantStatsService(&stats_service.DepsStats{
		Mencache:      s_cache,
		Logger:        logObj,
		Repository:    s_repo,
		Observability: obs,
	})
	a_svc := apikey_service.NewMerchantStatsByApiKeyService(&apikey_service.DepsStatsByApiKey{
		Mencache:      a_cache,
		Logger:        logObj,
		Repository:    a_repo,
		Observability: obs,
	})
	m_svc := merchant_service.NewMerchantStatsByMerchantService(&merchant_service.DepsStatsByMerchant{
		Mencache:      m_cache,
		Logger:        logObj,
		Repository:    m_repo,
		Observability: obs,
	})

	// Wrap in service.Service if needed, but handler constructor takes interfaces
	// Actually, MerchantStatsHandler constructor takes service.Service which might be too complex to mock entirely here
	// I'll check what service.Service has.
	// Oh, I see NewMerchantStatsHandler(service service.Service)
	// I'll create a mock or a partial service.Service.

	// I'll just use the real handler constructor if possible or create individual handlers.
	amountHandler := merchantstatshandler.NewMerchantStatsAmountHandler(svc, m_svc, a_svc)
	methodHandler := merchantstatshandler.NewMerchantStatsMethodHandler(svc, m_svc, a_svc)
	totalAmountHandler := merchantstatshandler.NewMerchantStatsTotalAmountHandler(svc, m_svc, a_svc)

	// gRPC Setup
	s.lis = bufconn.Listen(1024 * 1024)
	grpcServer := grpc.NewServer()
	pb.RegisterMerchantStatsAmountServiceServer(grpcServer, amountHandler)
	pb.RegisterMerchantStatsMethodServiceServer(grpcServer, methodHandler)
	pb.RegisterMerchantStatsTotalAmountServiceServer(grpcServer, totalAmountHandler)

	go func() {
		if err := grpcServer.Serve(s.lis); err != nil {
			log.Printf("Server exited: %v", err)
		}
	}()

	conn, err := grpc.DialContext(context.Background(), "bufnet",
		grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
			return s.lis.Dial()
		}),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	s.Require().NoError(err)
	s.conn = conn

	s.testYear = time.Now().Year()

	// Seed Data
	ctx := context.Background()
	var userID int32
	s.dbPool.QueryRowContext(ctx, "INSERT INTO users (firstname, lastname, email, password, verification_code, is_verified) VALUES ('Gapi', 'Stats', 'gapi_stats@example.com', 'pass', '123', true) RETURNING user_id").Scan(&userID)

	s.apiKey = "gapi-key-1"
	s.dbPool.QueryRowContext(ctx, "INSERT INTO merchants (name, api_key, user_id, status) VALUES ('Gapi Merchant', $1, $2, 'active') RETURNING merchant_id", s.apiKey, userID).Scan(&s.merchantID)

	s.dbPool.ExecContext(ctx, "INSERT INTO cards (user_id, card_number, card_type, cvv, card_provider, expire_date) VALUES ($1, '1212121212121212', 'debit', '123', 'visa', '2030-01-01')", userID)
	s.dbPool.ExecContext(ctx, "INSERT INTO transactions (card_number, amount, payment_method, merchant_id, transaction_time, status) VALUES ('1212121212121212', 2000, 'visa', $1, $2, 'success')", s.merchantID, time.Date(s.testYear, 1, 1, 10, 0, 0, 0, time.UTC))
}

func (s *MerchantStatsHandlerGapiTestSuite) TearDownSuite() {
	s.conn.Close()
	s.lis.Close()
	s.ts.Teardown()
}

func (s *MerchantStatsHandlerGapiTestSuite) TestAmountHandlers() {
	client := pb.NewMerchantStatsAmountServiceClient(s.conn)
	ctx := context.Background()

	// 1. Global Monthly
	res1, err := client.FindMonthlyAmountMerchant(ctx, &pb_merchant.FindYearMerchant{Year: int32(s.testYear)})
	s.Require().NoError(err)
	s.Require().Equal("success", res1.Status, "Message: %s", res1.Message)
	s.Require().NotEmpty(res1.Data)
	s.Require().EqualValues(2000, res1.Data[0].TotalAmount)

	// 2. By Merchant
	res2, err := client.FindMonthlyAmountByMerchants(ctx, &pb_merchant.FindYearMerchantById{
		Year:       int32(s.testYear),
		MerchantId: s.merchantID,
	})
	s.Require().NoError(err)
	s.Require().Equal("success", res2.Status, "Message: %s", res2.Message)
	s.Require().NotEmpty(res2.Data)
	s.Require().EqualValues(2000, res2.Data[0].TotalAmount)

	// 3. By Apikey
	res3, err := client.FindMonthlyAmountByApikey(ctx, &pb_merchant.FindYearMerchantByApikey{
		Year:   int32(s.testYear),
		ApiKey: s.apiKey,
	})
	s.Require().NoError(err)
	s.Require().Equal("success", res3.Status, "Message: %s", res3.Message)
	s.Require().NotEmpty(res3.Data)
	s.Require().EqualValues(2000, res3.Data[0].TotalAmount)
}

func (s *MerchantStatsHandlerGapiTestSuite) TestMethodHandlers() {
	client := pb.NewMerchantStatsMethodServiceClient(s.conn)
	ctx := context.Background()

	// 1. Global Monthly
	res1, err := client.FindMonthlyPaymentMethodsMerchant(ctx, &pb_merchant.FindYearMerchant{Year: int32(s.testYear)})
	s.Require().NoError(err)
	s.Require().Equal("success", res1.Status)
	s.Require().NotEmpty(res1.Data)
	// 'visa' should be at least one of the methods
	found := false
	for _, d := range res1.Data {
		if d.PaymentMethod == "visa" {
			found = true
			break
		}
	}
	s.True(found)
}

func (s *MerchantStatsHandlerGapiTestSuite) TestTotalAmountHandlers() {
	client := pb.NewMerchantStatsTotalAmountServiceClient(s.conn)
	ctx := context.Background()

	// 1. Global Monthly
	res1, err := client.FindMonthlyTotalAmountMerchant(ctx, &pb_merchant.FindYearMerchant{Year: int32(s.testYear)})
	s.Require().NoError(err)
	s.Require().Equal("success", res1.Status)
	s.Require().NotEmpty(res1.Data)
	s.Require().EqualValues(2000, res1.Data[0].TotalAmount)
}

func TestMerchantStatsHandlerGapiSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	suite.Run(t, new(MerchantStatsHandlerGapiTestSuite))
}
