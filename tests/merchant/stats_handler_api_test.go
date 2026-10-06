package merchant_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	merchant_handler "github.com/MamangRust/monolith-payment-gateway-apigateway/handler/merchant"
	merchantstatshandler "github.com/MamangRust/monolith-payment-gateway-merchant/handler/stats"
	stats_cache "github.com/MamangRust/monolith-payment-gateway-merchant/redis/stats"
	apikey_cache "github.com/MamangRust/monolith-payment-gateway-merchant/redis/statsbyapikey"
	merchant_cache "github.com/MamangRust/monolith-payment-gateway-merchant/redis/statsbymerchant"
	stats_repo "github.com/MamangRust/monolith-payment-gateway-merchant/repository/stats"
	apikey_repo "github.com/MamangRust/monolith-payment-gateway-merchant/repository/statsbyapikey"
	merchant_repo "github.com/MamangRust/monolith-payment-gateway-merchant/repository/statsbymerchant"
	stats_service "github.com/MamangRust/monolith-payment-gateway-merchant/service/stats"
	apikey_service "github.com/MamangRust/monolith-payment-gateway-merchant/service/statsbyapikey"
	merchant_service "github.com/MamangRust/monolith-payment-gateway-merchant/service/statsbymerchant"
	pb "github.com/MamangRust/monolith-payment-gateway-pb/merchant/stats"
	logger_pkg "github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/cache"
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
	"github.com/MamangRust/monolith-payment-gateway-shared/observability"
	tests "github.com/MamangRust/monolith-payment-gateway-test"
	"github.com/jmoiron/sqlx"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/suite"
	sdklog "go.opentelemetry.io/otel/sdk/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

type MerchantStatsHandlerApiTestSuite struct {
	suite.Suite
	ts         *tests.TestSuite
	dbPool     *sqlx.DB
	echo       *echo.Echo
	lis        *bufconn.Listener
	grpcServer *grpc.Server
	merchantID int32
	apiKey     string
	testYear   int
}

func (s *MerchantStatsHandlerApiTestSuite) SetupSuite() {
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

	// Dependency Injection (Merchant Service side)
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

	amountHandler := merchantstatshandler.NewMerchantStatsAmountHandler(svc, m_svc, a_svc)
	methodHandler := merchantstatshandler.NewMerchantStatsMethodHandler(svc, m_svc, a_svc)
	totalAmountHandler := merchantstatshandler.NewMerchantStatsTotalAmountHandler(svc, m_svc, a_svc)

	// gRPC Setup
	s.lis = bufconn.Listen(1024 * 1024)
	s.grpcServer = grpc.NewServer()
	pb.RegisterMerchantStatsAmountServiceServer(s.grpcServer, amountHandler)
	pb.RegisterMerchantStatsMethodServiceServer(s.grpcServer, methodHandler)
	pb.RegisterMerchantStatsTotalAmountServiceServer(s.grpcServer, totalAmountHandler)

	go func() {
		if err := s.grpcServer.Serve(s.lis); err != nil {
			// use print instead of fatal in test suite goroutine
		}
	}()

	conn, err := grpc.DialContext(context.Background(), "bufnet",
		grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
			return s.lis.Dial()
		}),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	s.Require().NoError(err)

	// Echo Setup
	s.echo = echo.New()
	apiHandler := errors.NewApiHandler(obs, logObj)

	merchant_handler.RegisterMerchantHandler(&merchant_handler.DepsMerchant{
		Client:     conn,
		E:          s.echo,
		Logger:     logObj,
		Cache:      cacheStore,
		ApiHandler: apiHandler,
	})

	s.testYear = time.Now().Year()

	// Seed Data
	ctx := context.Background()
	var userID int32
	s.dbPool.QueryRowContext(ctx, "INSERT INTO users (firstname, lastname, email, password, verification_code, is_verified) VALUES ('Api', 'Stats', 'api_stats@example.com', 'pass', '123', true) RETURNING user_id").Scan(&userID)

	s.apiKey = "api-key-stats"
	s.dbPool.QueryRowContext(ctx, "INSERT INTO merchants (name, api_key, user_id, status) VALUES ('Api Stats Merchant', $1, $2, 'active') RETURNING merchant_id", s.apiKey, userID).Scan(&s.merchantID)

	s.dbPool.ExecContext(ctx, "INSERT INTO cards (user_id, card_number, card_type, cvv, card_provider, expire_date) VALUES ($1, '9999999999999999', 'debit', '123', 'visa', '2030-01-01')", userID)
	s.dbPool.ExecContext(ctx, "INSERT INTO transactions (card_number, amount, payment_method, merchant_id, transaction_time, status) VALUES ('9999999999999999', 3000, 'mastercard', $1, $2, 'success')", s.merchantID, time.Date(s.testYear, 2, 1, 10, 0, 0, 0, time.UTC))
}

func (s *MerchantStatsHandlerApiTestSuite) TearDownSuite() {
	if s.grpcServer != nil {
		s.grpcServer.Stop()
	}
	if s.lis != nil {
		s.lis.Close()
	}
	s.ts.Teardown()
}

func (s *MerchantStatsHandlerApiTestSuite) TestAmountApi() {
	// 1. Global Monthly
	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/merchant-stats-amount/monthly-amount?year=%d", s.testYear), nil)
	rec := httptest.NewRecorder()
	s.echo.ServeHTTP(rec, req)

	s.Equal(http.StatusOK, rec.Code)
	var res map[string]interface{}
	json.Unmarshal(rec.Body.Bytes(), &res)
	s.Equal("success", res["status"])

	data := res["data"].([]interface{})
	s.NotEmpty(data)
	// Feb is index 1
	febData := data[1].(map[string]interface{})
	s.EqualValues(3000, febData["total_amount"])
}

func (s *MerchantStatsHandlerApiTestSuite) TestMethodApi() {
	// 1. Global Monthly
	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/merchant-stats-method/monthly-payment-methods?year=%d", s.testYear), nil)
	rec := httptest.NewRecorder()
	s.echo.ServeHTTP(rec, req)

	s.Equal(http.StatusOK, rec.Code)
	var res map[string]interface{}
	json.Unmarshal(rec.Body.Bytes(), &res)
	s.Equal("success", res["status"])

	data := res["data"].([]interface{})
	s.NotEmpty(data)

	// Check for 'mastercard' in Feb
	found := false
	for _, d := range data {
		m := d.(map[string]interface{})
		if m["payment_method"] == "mastercard" && m["month"] == "Feb" {
			found = true
			break
		}
	}
	s.True(found)
}

func (s *MerchantStatsHandlerApiTestSuite) TestTotalAmountApi() {
	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/merchant-stats-totalamount/monthly-total-amount?year=%d", s.testYear), nil)
	rec := httptest.NewRecorder()
	s.echo.ServeHTTP(rec, req)

	s.Equal(http.StatusOK, rec.Code)
	var res map[string]interface{}
	json.Unmarshal(rec.Body.Bytes(), &res)
	s.Equal("success", res["status"])

	data := res["data"].([]interface{})
	s.NotEmpty(data)
	febData := data[1].(map[string]interface{})
	s.EqualValues(3000, febData["total_amount"])
}

func TestMerchantStatsHandlerApiSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	suite.Run(t, new(MerchantStatsHandlerApiTestSuite))
}
