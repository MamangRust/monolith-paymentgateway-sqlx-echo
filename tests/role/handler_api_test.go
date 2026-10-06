package role_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	role_handler "github.com/MamangRust/monolith-payment-gateway-apigateway/handler/role"
	"github.com/MamangRust/monolith-payment-gateway-apigateway/redis"
	pb "github.com/MamangRust/monolith-payment-gateway-pb/role"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-role/handler"
	"github.com/MamangRust/monolith-payment-gateway-role/repository"
	"github.com/MamangRust/monolith-payment-gateway-role/service"
	"github.com/MamangRust/monolith-payment-gateway-shared/cache"
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
	"github.com/MamangRust/monolith-payment-gateway-shared/observability"
	tests "github.com/MamangRust/monolith-payment-gateway-test"
	"github.com/jmoiron/sqlx"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/labstack/echo/v4"
	redis_client "github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/suite"
	sdklog "go.opentelemetry.io/otel/sdk/log"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"net"
)

type RoleApiTestSuite struct {
	suite.Suite
	ts         *tests.TestSuite
	dbPool     *sqlx.DB
	echo       *echo.Echo
	grpcServer *grpc.Server
	lis        *bufconn.Listener
	roleID     int32
}

func (s *RoleApiTestSuite) SetupSuite() {
	ts, err := tests.SetupTestSuite()
	s.Require().NoError(err)
	s.ts = ts

	pool, err := sqlx.ConnectContext(s.ts.Ctx, "pgx", s.ts.DBURL)
	s.Require().NoError(err)
	s.dbPool = pool

	opts, err := redis_client.ParseURL(s.ts.RedisURL)
	s.Require().NoError(err)
	redisClient := redis_client.NewClient(opts)

	queries := pool
	repos := repository.NewRepositories(queries)

	logger.ResetInstance()
	lp := sdklog.NewLoggerProvider()
	log, _ := logger.NewLogger("test", lp)
	cacheMetrics, _ := observability.NewCacheMetrics("test")
	cacheStore := cache.NewCacheStore(redisClient, log, cacheMetrics)

	svc := service.NewService(&service.Deps{
		Repositories: repos,
		Logger:       log,
		Cache:        cacheStore,
	})

	roleHandler := handler.NewHandler(svc)

	// Setup gRPC server
	s.lis = bufconn.Listen(1024 * 1024)
	s.grpcServer = grpc.NewServer()
	pb.RegisterRoleServiceServer(s.grpcServer, roleHandler.RoleQuery)
	pb.RegisterRoleCommandServiceServer(s.grpcServer, roleHandler.RoleCommand)

	go func() {
		if err := s.grpcServer.Serve(s.lis); err != nil {
			// fmt.Printf("grpc server error: %v\n", err)
		}
	}()

	// Setup Echo and register apigateway handler
	s.echo = echo.New()

	// Inject user_id middleware
	s.echo.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("user_id", 1)
			return next(c)
		}
	})

	// Pre-populate Role Cache to bypass Kafka
	roleMencache := mencache.NewRoleCache(cacheStore)
	roleMencache.SetRoleCache(context.Background(), "1", []string{"Admin_Admin_14", "Admin_Role_10"})

	// Verify cache
	if roles, found := roleMencache.GetRoleCache(context.Background(), "1"); found {
		log.Info("Cache verified in SetupSuite", zap.Strings("roles", roles))
	} else {
		log.Warn("Cache NOT found in SetupSuite!")
	}

	conn, err := grpc.DialContext(context.Background(), "bufnet",
		grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
			return s.lis.Dial()
		}),
		grpc.WithInsecure())
	s.Require().NoError(err)

	obs, err := observability.NewObservability("test", log)
	s.Require().NoError(err)
	apiHandler := errors.NewApiHandler(obs, log)

	// Register via exported RegisterRoleHandler
	role_handler.RegisterRoleHandler(&role_handler.DepsRole{
		Client:     conn,
		E:          s.echo,
		Logger:     log,
		CacheStore: cacheStore,
		Cache:      roleMencache,
		Kafka:      nil, // Passing nil Kafka
		ApiHandler: apiHandler,
	})
}

func (s *RoleApiTestSuite) TearDownSuite() {
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

func (s *RoleApiTestSuite) Test1_RoleLifecycle() {
	// Create
	reqJSON := `{"name": "Test API Role"}`
	req := httptest.NewRequest(http.MethodPost, "/api/role", strings.NewReader(reqJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	s.echo.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		s.T().Errorf("Response Body: %s", rec.Body.String())
		return
	}

	s.Equal(http.StatusOK, rec.Code)
	var res map[string]interface{}
	err := json.Unmarshal(rec.Body.Bytes(), &res)
	s.NoError(err)

	data, ok := res["data"].(map[string]interface{})
	if !ok {
		s.T().Errorf("Data not found in response: %v", res)
		return
	}
	s.roleID = int32(data["id"].(float64))
	s.Equal("Test API Role", data["name"])

	// FindById
	req = httptest.NewRequest(http.MethodGet, "/api/role-query/"+strconv.Itoa(int(s.roleID)), nil)
	rec = httptest.NewRecorder()
	s.echo.ServeHTTP(rec, req)
	s.Equal(http.StatusOK, rec.Code)

	// Update
	updateJSON := `{"name": "Updated API Role"}`
	req = httptest.NewRequest(http.MethodPost, "/api/role/"+strconv.Itoa(int(s.roleID)), strings.NewReader(updateJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec = httptest.NewRecorder()
	s.echo.ServeHTTP(rec, req)
	s.Equal(http.StatusOK, rec.Code)
}

func (s *RoleApiTestSuite) Test2_QueryOperations() {
	// FindAll
	req := httptest.NewRequest(http.MethodGet, "/api/role-query?page=1&page_size=10", nil)
	rec := httptest.NewRecorder()
	s.echo.ServeHTTP(rec, req)

	s.Equal(http.StatusOK, rec.Code)
	var res map[string]interface{}
	err := json.Unmarshal(rec.Body.Bytes(), &res)
	s.NoError(err)
	s.Equal("success", res["status"])
}

func TestRoleApiSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	suite.Run(t, new(RoleApiTestSuite))
}
