package card_test

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

	card_handler "github.com/MamangRust/monolith-payment-gateway-apigateway/handler/card"
	"github.com/MamangRust/monolith-payment-gateway-card/handler"
	"github.com/MamangRust/monolith-payment-gateway-card/repository"
	"github.com/MamangRust/monolith-payment-gateway-card/service"
	pb "github.com/MamangRust/monolith-payment-gateway-pb/card"
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

type CardApiTestSuite struct {
	suite.Suite
	ts         *tests.TestSuite
	dbPool     *sqlx.DB
	echo       *echo.Echo
	grpcServer *grpc.Server
	lis        *bufconn.Listener
	userID     int32
	cardID     int32
	cardNumber string
	identity   *tests.IdentityStack
}

func (s *CardApiTestSuite) SetupSuite() {
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

	logger.ResetInstance()
	lp := sdklog.NewLoggerProvider()
	log, _ := logger.NewLogger("test", lp)
	hasher := hash.NewHashingPassword()
	cacheMetrics, _ := observability.NewCacheMetrics("test")
	cacheStore := cache.NewCacheStore(redisClient, log, cacheMetrics)

	s.identity, err = tests.StartIdentityStack(pool, cacheStore, log)
	s.Require().NoError(err)

	repos := repository.NewRepositories(queries, pbuser.NewUserQueryServiceClient(s.identity.User.Conn))
	userRepos := user_repository.NewRepositories(
		queries,
		pbrole.NewRoleServiceClient(s.identity.Role.Conn),
		pbrole.NewRoleCommandServiceClient(s.identity.Role.Conn),
	)

	cardSvc := service.NewService(&service.Deps{
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

	cardH := handler.NewHandler(cardSvc)
	userH := user_handler.NewHandler(userSvc)

	// Setup gRPC server
	s.lis = bufconn.Listen(1024 * 1024)
	s.grpcServer = grpc.NewServer()
	pb.RegisterCardQueryServiceServer(s.grpcServer, cardH)
	pb.RegisterCardCommandServiceServer(s.grpcServer, cardH)

	go func() {
		if err := s.grpcServer.Serve(s.lis); err != nil {
		}
	}()

	s.echo = echo.New()

	// Middleware for injecting userID in tests. The card query handler reads
	// c.Get("userId") (lowercase d), matching the real JWT middleware.
	s.echo.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if uid := c.Request().Header.Get("X-Test-User-ID"); uid != "" {
				c.Set("userId", uid)
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

	card_handler.RegisterCardHandler(&card_handler.DepsCard{
		Client:     conn,
		E:          s.echo,
		Logger:     log,
		Cache:      cacheStore,
		ApiHandler: apiHandler,
	})

	// Create a user for testing
	ctx := context.Background()
	userRes, err := userH.Create(ctx, &pbuser.CreateUserRequest{
		Firstname:       "Card",
		Lastname:        "Api",
		Email:           fmt.Sprintf("card.api.%d.%d@example.com", time.Now().UnixNano(), time.Now().UnixNano()%1000),
		Password:        "Password123!",
		ConfirmPassword: "Password123!",
	})
	s.Require().NoError(err)
	s.userID = userRes.Data.Id
}

func (s *CardApiTestSuite) TearDownSuite() {
	if s.identity != nil {
		s.identity.Close()
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

func (s *CardApiTestSuite) Test1_CardLifecycle() {
	// Create
	reqJSON := fmt.Sprintf(`{"user_id": %d, "card_type": "debit", "expire_date": "%s", "cvv": "123", "card_provider": "visa"}`, s.userID, time.Now().AddDate(5, 0, 0).Format(time.RFC3339))
	req := httptest.NewRequest(http.MethodPost, "/api/card-command/create", strings.NewReader(reqJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	s.echo.ServeHTTP(rec, req)

	s.Equal(http.StatusOK, rec.Code)
	var res map[string]interface{}
	err := json.Unmarshal(rec.Body.Bytes(), &res)
	s.NoError(err)
	data := res["data"].(map[string]interface{})
	s.cardID = int32(data["id"].(float64))
	s.cardNumber = data["card_number"].(string)
	s.Equal("debit", data["card_type"])

	// FindById
	req = httptest.NewRequest(http.MethodGet, "/api/card-query/"+strconv.Itoa(int(s.cardID)), nil)
	rec = httptest.NewRecorder()
	s.echo.ServeHTTP(rec, req)
	s.Equal(http.StatusOK, rec.Code)

	// Update
	updateJSON := fmt.Sprintf(`{"card_id": %d, "user_id": %d, "card_type": "credit", "expire_date": "%s", "cvv": "456", "card_provider": "mastercard"}`, s.cardID, s.userID, time.Now().AddDate(5, 0, 0).Format(time.RFC3339))
	req = httptest.NewRequest(http.MethodPost, "/api/card-command/update/"+strconv.Itoa(int(s.cardID)), strings.NewReader(updateJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec = httptest.NewRecorder()
	s.echo.ServeHTTP(rec, req)
	s.Equal(http.StatusOK, rec.Code)
}

func (s *CardApiTestSuite) Test2_QueryOperations() {
	// FindAll
	req := httptest.NewRequest(http.MethodGet, "/api/card-query?page=1&page_size=10", nil)
	rec := httptest.NewRecorder()
	s.echo.ServeHTTP(rec, req)
	s.Equal(http.StatusOK, rec.Code)

	// FindActive
	req = httptest.NewRequest(http.MethodGet, "/api/card-query/active?page=1&page_size=10", nil)
	rec = httptest.NewRecorder()
	s.echo.ServeHTTP(rec, req)
	s.Equal(http.StatusOK, rec.Code)

	// FindByUserID
	req = httptest.NewRequest(http.MethodGet, "/api/card-query/user", nil)
	req.Header.Set("X-Test-User-ID", strconv.Itoa(int(s.userID)))
	rec = httptest.NewRecorder()
	s.echo.ServeHTTP(rec, req)
	s.Equal(http.StatusOK, rec.Code)

	// FindByCardNumber
	req = httptest.NewRequest(http.MethodGet, "/api/card-query/card_number/"+s.cardNumber, nil)
	rec = httptest.NewRecorder()
	s.echo.ServeHTTP(rec, req)
	s.Equal(http.StatusOK, rec.Code)
}

func (s *CardApiTestSuite) Test3_TrashAndRestore() {
	// Trash
	req := httptest.NewRequest(http.MethodPost, "/api/card-command/trashed/"+strconv.Itoa(int(s.cardID)), nil)
	rec := httptest.NewRecorder()
	s.echo.ServeHTTP(rec, req)
	s.Equal(http.StatusOK, rec.Code)

	// FindTrashed
	req = httptest.NewRequest(http.MethodGet, "/api/card-query/trashed?page=1&page_size=10", nil)
	rec = httptest.NewRecorder()
	s.echo.ServeHTTP(rec, req)
	s.Equal(http.StatusOK, rec.Code)

	// Restore
	req = httptest.NewRequest(http.MethodPost, "/api/card-command/restore/"+strconv.Itoa(int(s.cardID)), nil)
	rec = httptest.NewRecorder()
	s.echo.ServeHTTP(rec, req)
	s.Equal(http.StatusOK, rec.Code)
}

func TestCardApiSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	suite.Run(t, new(CardApiTestSuite))
}
