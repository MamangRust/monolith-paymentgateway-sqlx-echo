package auth_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/MamangRust/monolith-payment-gateway-auth/repository"
	"github.com/MamangRust/monolith-payment-gateway-auth/service"
	"github.com/MamangRust/monolith-payment-gateway-pkg/auth"
	"github.com/MamangRust/monolith-payment-gateway-pkg/hash"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/cache"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/observability"
	tests "github.com/MamangRust/monolith-payment-gateway-test"

	"github.com/jmoiron/sqlx"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/suite"
	sdklog "go.opentelemetry.io/otel/sdk/log"
)

type AuthServiceTestSuite struct {
	suite.Suite
	ts          *tests.TestSuite
	dbPool      *sqlx.DB
	redisClient *redis.Client
	service     *service.Service
	identity    *tests.IdentityStack
	email       string
	password    string
}

func (s *AuthServiceTestSuite) SetupSuite() {
	ts, err := tests.SetupTestSuite()
	s.Require().NoError(err)
	s.ts = ts

	pool, err := sqlx.ConnectContext(s.ts.Ctx, "pgx", s.ts.DBURL)
	s.Require().NoError(err)
	s.dbPool = pool

	opts, err := redis.ParseURL(s.ts.RedisURL)
	s.Require().NoError(err)
	s.redisClient = redis.NewClient(opts)

	queries := pool
	repos := repository.NewRepositories(&repository.Deps{Db: queries})

	logger.ResetInstance()
	lp := sdklog.NewLoggerProvider()
	log, _ := logger.NewLogger("test", lp)
	cacheMetrics, _ := observability.NewCacheMetrics("test")
	cacheStore := cache.NewCacheStore(s.redisClient, log, cacheMetrics)

	tokenManager, _ := auth.NewManager("mysecret")
	hasher := hash.NewHashingPassword()

	s.identity, err = tests.StartIdentityStack(pool, cacheStore, log)
	s.Require().NoError(err)

	s.service = service.NewService(&service.Deps{
		Repositories:    repos,
		Logger:          log,
		Cache:           cacheStore,
		RoleAdapter:     s.identity.Role.RoleAdapter,
		UserRoleAdapter: s.identity.Role.UserRoleAdapter,
		Token:           tokenManager,
		Hash:            hasher,
		Kafka:           nil,
	})

	s.email = "auth.service.test@example.com"
	s.password = "password123"
}

func (s *AuthServiceTestSuite) TearDownSuite() {
	s.identity.Close()
	s.redisClient.Close()
	s.dbPool.Close()
	s.ts.Teardown()
}

func (s *AuthServiceTestSuite) Test1_Register() {
	ctx := context.Background()
	req := &requests.RegisterRequest{
		FirstName:       "Auth",
		LastName:        "Service",
		Email:           s.email,
		Password:        s.password,
		ConfirmPassword: s.password,
	}

	res, err := s.service.Register.Register(ctx, req)
	s.NoError(err)
	s.NotNil(res)
	s.Equal(s.email, res.Email)

	// Login requires a verified account — complete the verification flow
	// using the code Register cached in Redis.
	s.verifyUser(s.email)
}

func (s *AuthServiceTestSuite) Test2_Login() {
	ctx := context.Background()
	req := &requests.AuthRequest{
		Email:    s.email,
		Password: s.password,
	}

	res, err := s.service.Login.Login(ctx, req)
	s.NoError(err)
	s.NotNil(res)
	s.NotEmpty(res.AccessToken)
	s.NotEmpty(res.RefreshToken)
}

func (s *AuthServiceTestSuite) Test4_LoginLockout() {
	ctx := context.Background()
	email := "locked.user@example.com"
	password := "wrongpassword"

	// Register user first
	regReq := &requests.RegisterRequest{
		FirstName:       "Locked",
		LastName:        "User",
		Email:           email,
		Password:        "correctpassword",
		ConfirmPassword: "correctpassword",
	}
	_, err := s.service.Register.Register(ctx, regReq)
	s.NoError(err)

	s.verifyUser(email)

	loginReq := &requests.AuthRequest{
		Email:    email,
		Password: password,
	}

	// Fail login 5 times
	for i := 0; i < 5; i++ {
		_, err := s.service.Login.Login(ctx, loginReq)
		s.Error(err)
	}

	// 6th attempt should return ErrAccountLocked
	_, err = s.service.Login.Login(ctx, loginReq)
	s.Error(err)
	s.Contains(err.Error(), "Account temporarily locked")
}

func (s *AuthServiceTestSuite) Test3_ForgotPassword() {
	ctx := context.Background()

	success, err := s.service.PasswordReset.ForgotPassword(ctx, s.email)
	s.NoError(err)
	s.True(success)
}

// verifyUser marks a just-registered user as verified, mirroring the real
// email-verification flow (Register caches the code in Redis under
// "register:verify_code:<email>").
func (s *AuthServiceTestSuite) verifyUser(email string) {
	ctx := context.Background()
	raw, err := s.redisClient.Get(ctx, "register:verify_code:"+email).Result()
	s.Require().NoError(err, "verification code should be cached after register")
	var code string
	s.Require().NoError(json.Unmarshal([]byte(raw), &code), "cached code is JSON-encoded")
	ok, err := s.service.PasswordReset.VerifyCode(ctx, code)
	s.Require().NoError(err)
	s.Require().True(ok)
}

func TestAuthServiceSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	suite.Run(t, new(AuthServiceTestSuite))
}
