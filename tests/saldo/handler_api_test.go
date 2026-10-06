package saldo_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	saldo_handler "github.com/MamangRust/monolith-payment-gateway-apigateway/handler/saldo"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/cache"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
	"github.com/MamangRust/monolith-payment-gateway-shared/observability"
	tests "github.com/MamangRust/monolith-payment-gateway-test"
	"github.com/jmoiron/sqlx"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/labstack/echo/v4"
	redis_client "github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/suite"
	sdklog "go.opentelemetry.io/otel/sdk/log"
)

type SaldoApiTestSuite struct {
	suite.Suite
	ts         *tests.TestSuite
	dbPool     *sqlx.DB
	echo       *echo.Echo
	stack      *tests.Stack
	userID     int32
	cardNumber string
	saldoID    int32
}

func (s *SaldoApiTestSuite) SetupSuite() {
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
	cacheMetrics, _ := observability.NewCacheMetrics("test")
	cacheStore := cache.NewCacheStore(redisClient, log, cacheMetrics)

	// The saldo API validates cards through the card gRPC service, so bring up
	// the full in-process role -> user -> card -> saldo graph backed by this db.
	stack, err := tests.StartStack(pool, cacheStore, log)
	s.Require().NoError(err)
	s.stack = stack

	s.echo = echo.New()

	obs, err := observability.NewObservability("test", log)
	s.Require().NoError(err)
	apiHandler := errors.NewApiHandler(obs, log)

	saldo_handler.RegisterSaldoHandler(&saldo_handler.DepsSaldo{
		Client:     stack.Saldo.Conn,
		E:          s.echo,
		Logger:     log,
		Cache:      cacheStore,
		ApiHandler: apiHandler,
	})

	// Create user
	ctx := context.Background()
	user, err := stack.User.Service.CreateUser(ctx, &requests.CreateUserRequest{
		FirstName:       "Saldo",
		LastName:        "Api",
		Email:           fmt.Sprintf("saldo.api.%d@example.com", time.Now().UnixNano()),
		Password:        "Password123!",
		ConfirmPassword: "Password123!",
	})
	s.Require().NoError(err)
	s.userID = user.UserID

	// Create card
	card, err := stack.Card.Repositories.CardCommand.CreateCard(ctx, &requests.CreateCardRequest{
		UserID:       int(s.userID),
		CardType:     "debit",
		ExpireDate:   time.Now().AddDate(5, 0, 0),
		CVV:          "123",
		CardProvider: "visa",
	})
	s.Require().NoError(err)
	s.cardNumber = card.CardNumber
}

func (s *SaldoApiTestSuite) TearDownSuite() {
	if s.stack != nil {
		s.stack.Close()
	}
	if s.dbPool != nil {
		s.dbPool.Close()
	}
	if s.ts != nil {
		s.ts.Teardown()
	}
}

func (s *SaldoApiTestSuite) Test1_SaldoLifecycle() {
	// Create
	reqJSON := fmt.Sprintf(`{"card_number": "%s", "total_balance": 100000}`, s.cardNumber)
	req := httptest.NewRequest(http.MethodPost, "/api/saldo-command/create", strings.NewReader(reqJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	s.echo.ServeHTTP(rec, req)

	s.Equal(http.StatusOK, rec.Code)
	var res map[string]interface{}
	err := json.Unmarshal(rec.Body.Bytes(), &res)
	s.NoError(err)
	data := res["data"].(map[string]interface{})
	s.saldoID = int32(data["id"].(float64))
	s.Equal(float64(100000), data["total_balance"])

	// FindById
	req = httptest.NewRequest(http.MethodGet, "/api/saldo-query/"+strconv.Itoa(int(s.saldoID)), nil)
	rec = httptest.NewRecorder()
	s.echo.ServeHTTP(rec, req)
	s.Equal(http.StatusOK, rec.Code)

	// Update
	updateJSON := fmt.Sprintf(`{"card_number": "%s", "total_balance": 200000}`, s.cardNumber)
	req = httptest.NewRequest(http.MethodPost, "/api/saldo-command/update/"+strconv.Itoa(int(s.saldoID)), strings.NewReader(updateJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec = httptest.NewRecorder()
	s.echo.ServeHTTP(rec, req)
	s.Equal(http.StatusOK, rec.Code)
}

func (s *SaldoApiTestSuite) Test2_QueryOperations() {
	// FindAll
	req := httptest.NewRequest(http.MethodGet, "/api/saldo-query?page=1&page_size=10&search="+s.cardNumber, nil)
	rec := httptest.NewRecorder()
	s.echo.ServeHTTP(rec, req)
	s.Equal(http.StatusOK, rec.Code)

	// FindActive
	req = httptest.NewRequest(http.MethodGet, "/api/saldo-query/active?page=1&page_size=10&search="+s.cardNumber, nil)
	rec = httptest.NewRecorder()
	s.echo.ServeHTTP(rec, req)
	s.Equal(http.StatusOK, rec.Code)

	// FindByCardNumber
	req = httptest.NewRequest(http.MethodGet, "/api/saldo-query/card_number/"+s.cardNumber, nil)
	rec = httptest.NewRecorder()
	s.echo.ServeHTTP(rec, req)
	s.Equal(http.StatusOK, rec.Code)
}

func (s *SaldoApiTestSuite) Test3_TrashAndRestore() {
	// Trash
	req := httptest.NewRequest(http.MethodPost, "/api/saldo-command/trashed/"+strconv.Itoa(int(s.saldoID)), nil)
	rec := httptest.NewRecorder()
	s.echo.ServeHTTP(rec, req)
	s.Equal(http.StatusOK, rec.Code)

	// FindTrashed
	req = httptest.NewRequest(http.MethodGet, "/api/saldo-query/trashed?page=1&page_size=10&search="+s.cardNumber, nil)
	rec = httptest.NewRecorder()
	s.echo.ServeHTTP(rec, req)
	s.Equal(http.StatusOK, rec.Code)

	// Restore
	req = httptest.NewRequest(http.MethodPost, "/api/saldo-command/restore/"+strconv.Itoa(int(s.saldoID)), nil)
	rec = httptest.NewRecorder()
	s.echo.ServeHTTP(rec, req)
	s.Equal(http.StatusOK, rec.Code)
}

func TestSaldoApiSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	suite.Run(t, new(SaldoApiTestSuite))
}
