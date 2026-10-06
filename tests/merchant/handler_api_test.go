package merchant_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	merchant_handler "github.com/MamangRust/monolith-payment-gateway-apigateway/handler/merchant"
	"github.com/MamangRust/monolith-payment-gateway-merchant/handler"
	"github.com/MamangRust/monolith-payment-gateway-merchant/repository"
	"github.com/MamangRust/monolith-payment-gateway-merchant/service"
	pb "github.com/MamangRust/monolith-payment-gateway-pb/merchant"
	pbrole "github.com/MamangRust/monolith-payment-gateway-pb/role"
	pbuser "github.com/MamangRust/monolith-payment-gateway-pb/user"
	"github.com/MamangRust/monolith-payment-gateway-pkg/hash"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/cache"
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
	"github.com/MamangRust/monolith-payment-gateway-shared/observability"
	tests "github.com/MamangRust/monolith-payment-gateway-test"
	user_handler "github.com/MamangRust/monolith-payment-gateway-user/handler"
	user_repository "github.com/MamangRust/monolith-payment-gateway-user/repository"
	user_service "github.com/MamangRust/monolith-payment-gateway-user/service"
	"github.com/jmoiron/sqlx"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/labstack/echo/v4"
	redis_client "github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/suite"
	sdklog "go.opentelemetry.io/otel/sdk/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

type MerchantApiTestSuite struct {
	suite.Suite
	ts         *tests.TestSuite
	dbPool     *sqlx.DB
	echo       *echo.Echo
	grpcServer *grpc.Server
	lis        *bufconn.Listener
	userID     int32
	merchantID int32
	apiKey     string
	identity   *tests.IdentityStack
}

func (s *MerchantApiTestSuite) SetupSuite() {
	ts, err := tests.SetupTestSuite()
	s.Require().NoError(err)
	s.ts = ts

	pool, err := sqlx.ConnectContext(s.ts.Ctx, "pgx", s.ts.DBURL)
	s.Require().NoError(err)
	s.dbPool = pool

	opts, err := redis_client.ParseURL(s.ts.RedisURL)
	s.Require().NoError(err)
	redisClient := redis_client.NewClient(opts)

	logger.ResetInstance()
	lp := sdklog.NewLoggerProvider()
	log, _ := logger.NewLogger("test", lp)
	hasher := hash.NewHashingPassword()
	cacheMetrics, _ := observability.NewCacheMetrics("test")
	cacheStore := cache.NewCacheStore(redisClient, log, cacheMetrics)

	// Start the role + user gRPC dependencies. Merchant lookups resolve users
	// through the user query client, and user creation assigns ROLE_ADMIN
	// through the role adapters.
	s.identity, err = tests.StartIdentityStack(pool, cacheStore, log)
	s.Require().NoError(err)

	queries := pool
	repos := repository.NewRepositories(queries, pbuser.NewUserQueryServiceClient(s.identity.User.Conn))
	userRepos := user_repository.NewRepositories(
		queries,
		pbrole.NewRoleServiceClient(s.identity.Role.Conn),
		pbrole.NewRoleCommandServiceClient(s.identity.Role.Conn),
	)

	merchantSvc := service.NewService(&service.Deps{
		Cache:        cacheStore,
		Repositories: repos,
		Logger:       log,
		Kafka:        nil,
	})

	userSvc := user_service.NewService(&user_service.Deps{
		Cache:           cacheStore,
		Repositories:    userRepos,
		Hash:            hasher,
		Logger:          log,
		RoleAdapter:     s.identity.Role.RoleAdapter,
		UserRoleAdapter: s.identity.Role.UserRoleAdapter,
	})

	merchantH := handler.NewHandler(merchantSvc)
	userH := user_handler.NewHandler(userSvc)

	// Setup gRPC server
	s.lis = bufconn.Listen(1024 * 1024)
	s.grpcServer = grpc.NewServer()
	pb.RegisterMerchantQueryServiceServer(s.grpcServer, merchantH)
	pb.RegisterMerchantCommandServiceServer(s.grpcServer, merchantH)

	go func() {
		if err := s.grpcServer.Serve(s.lis); err != nil {
		}
	}()

	s.echo = echo.New()

	// Middleware for injecting userID in tests
	s.echo.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if uid := c.Request().Header.Get("X-Test-User-ID"); uid != "" {
				uidInt, _ := strconv.Atoi(uid)
				c.Set("user_id", int32(uidInt))
			}
			return next(c)
		}
	})

	conn, err := grpc.DialContext(context.Background(), "bufnet",
		grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
			return s.lis.Dial()
		}),
		grpc.WithInsecure())
	s.Require().NoError(err)

	obs, err := observability.NewObservability("test", log)
	s.Require().NoError(err)
	apiHandler := errors.NewApiHandler(obs, log)

	merchant_handler.RegisterMerchantHandler(&merchant_handler.DepsMerchant{
		Client:     conn,
		E:          s.echo,
		Logger:     log,
		Cache:      cacheStore,
		ApiHandler: apiHandler,
	})

	// Create a user for testing
	ctx := context.Background()
	userRes, err := userH.Create(ctx, &pbuser.CreateUserRequest{
		Firstname:       "Merchant",
		Lastname:        "Api",
		Email:           fmt.Sprintf("merchant.api.%d@example.com", time.Now().UnixNano()),
		Password:        "Password123!",
		ConfirmPassword: "Password123!",
	})
	s.Require().NoError(err)
	s.userID = userRes.Data.Id
}

func (s *MerchantApiTestSuite) TearDownSuite() {
	if s.grpcServer != nil {
		s.grpcServer.Stop()
	}
	if s.identity != nil {
		s.identity.Close()
	}
	if s.dbPool != nil {
		s.dbPool.Close()
	}
	if s.ts != nil {
		s.ts.Teardown()
	}
}

func (s *MerchantApiTestSuite) Test1_MerchantLifecycle() {
	// Create
	reqJSON := fmt.Sprintf(`{"name": "Api Merchant", "user_id": %d}`, s.userID)
	req := httptest.NewRequest(http.MethodPost, "/api/merchant-command/create", strings.NewReader(reqJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	s.echo.ServeHTTP(rec, req)

	s.Equal(http.StatusOK, rec.Code)
	var res map[string]interface{}
	err := json.Unmarshal(rec.Body.Bytes(), &res)
	s.NoError(err)
	data := res["data"].(map[string]interface{})
	s.merchantID = int32(data["id"].(float64))
	s.apiKey = data["api_key"].(string)
	s.Equal("Api Merchant", data["name"])

	// FindById
	req = httptest.NewRequest(http.MethodGet, "/api/merchant-query/"+strconv.Itoa(int(s.merchantID)), nil)
	rec = httptest.NewRecorder()
	s.echo.ServeHTTP(rec, req)
	s.Equal(http.StatusOK, rec.Code)

	// Update (route is /update/:id, not /updates/:id)
	updateJSON := fmt.Sprintf(`{"name": "Updated Api Merchant", "user_id": %d, "status": "active"}`, s.userID)
	req = httptest.NewRequest(http.MethodPost, "/api/merchant-command/update/"+strconv.Itoa(int(s.merchantID)), strings.NewReader(updateJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec = httptest.NewRecorder()
	s.echo.ServeHTTP(rec, req)
	s.Equal(http.StatusOK, rec.Code)
}

func (s *MerchantApiTestSuite) Test2_QueryOperations() {
	// FindAll
	req := httptest.NewRequest(http.MethodGet, "/api/merchant-query?page=1&page_size=10", nil)
	rec := httptest.NewRecorder()
	s.echo.ServeHTTP(rec, req)
	s.Equal(http.StatusOK, rec.Code)

	// FindActive
	req = httptest.NewRequest(http.MethodGet, "/api/merchant-query/active?page=1&page_size=10", nil)
	rec = httptest.NewRecorder()
	s.echo.ServeHTTP(rec, req)
	s.Equal(http.StatusOK, rec.Code)

	// FindByApiKey
	req = httptest.NewRequest(http.MethodGet, "/api/merchant-query/api-key?api_key="+s.apiKey, nil)
	rec = httptest.NewRecorder()
	s.echo.ServeHTTP(rec, req)
	s.Equal(http.StatusOK, rec.Code)

	// FindByMerchantUserId
	req = httptest.NewRequest(http.MethodGet, "/api/merchant-query/merchant-user", nil)
	req.Header.Set("X-Test-User-ID", strconv.Itoa(int(s.userID)))
	rec = httptest.NewRecorder()
	s.echo.ServeHTTP(rec, req)
	s.Equal(http.StatusOK, rec.Code)
}

func (s *MerchantApiTestSuite) Test3_TrashAndRestore() {
	// Trash
	req := httptest.NewRequest(http.MethodPost, "/api/merchant-command/trashed/"+strconv.Itoa(int(s.merchantID)), nil)
	rec := httptest.NewRecorder()
	s.echo.ServeHTTP(rec, req)
	s.Equal(http.StatusOK, rec.Code)

	// FindTrashed
	req = httptest.NewRequest(http.MethodGet, "/api/merchant-query/trashed?page=1&page_size=10", nil)
	rec = httptest.NewRecorder()
	s.echo.ServeHTTP(rec, req)
	s.Equal(http.StatusOK, rec.Code)

	// Restore
	req = httptest.NewRequest(http.MethodPost, "/api/merchant-command/restore/"+strconv.Itoa(int(s.merchantID)), nil)
	rec = httptest.NewRecorder()
	s.echo.ServeHTTP(rec, req)
	s.Equal(http.StatusOK, rec.Code)
}

func TestMerchantApiSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	suite.Run(t, new(MerchantApiTestSuite))
}
