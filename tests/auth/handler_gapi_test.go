package auth_test

import (
	"context"
	"encoding/json"
	"net"
	"strconv"
	"testing"

	"github.com/MamangRust/monolith-payment-gateway-auth/handler"
	"github.com/MamangRust/monolith-payment-gateway-auth/repository"
	"github.com/MamangRust/monolith-payment-gateway-auth/service"
	pb "github.com/MamangRust/monolith-payment-gateway-pb"
	"github.com/MamangRust/monolith-payment-gateway-pkg/auth"
	"github.com/MamangRust/monolith-payment-gateway-pkg/hash"
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
	"google.golang.org/grpc/credentials/insecure"
)

type AuthHandlerGapiTestSuite struct {
	suite.Suite
	ts          *tests.TestSuite
	dbPool      *sqlx.DB
	redisClient *redis.Client
	client      pb.AuthServiceClient
	conn        *grpc.ClientConn
	grpcServer  *grpc.Server
	identity    *tests.IdentityStack
	email       string
	password    string
	accessToken string
}

func (s *AuthHandlerGapiTestSuite) SetupSuite() {
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

	svc := service.NewService(&service.Deps{
		Repositories:    repos,
		Logger:          log,
		Cache:           cacheStore,
		RoleAdapter:     s.identity.Role.RoleAdapter,
		UserRoleAdapter: s.identity.Role.UserRoleAdapter,
		Token:           tokenManager,
		Hash:            hasher,
		Kafka:           nil,
	})

	h := handler.NewAuthHandleGrpc(svc, log)

	s.grpcServer = grpc.NewServer()
	pb.RegisterAuthServiceServer(s.grpcServer, h)

	lis, err := net.Listen("tcp", "localhost:0")
	s.Require().NoError(err)

	go func() {
		_ = s.grpcServer.Serve(lis)
	}()

	conn, err := grpc.NewClient(lis.Addr().String(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	s.Require().NoError(err)
	s.conn = conn
	s.client = pb.NewAuthServiceClient(conn)

	s.email = "auth.handler.gapi.test@example.com"
	s.password = "password123"

	// Seed ROLE_ADMIN
	_, _ = pool.ExecContext(context.Background(), "INSERT INTO roles (role_name) VALUES ('ROLE_ADMIN')")
}

func (s *AuthHandlerGapiTestSuite) TearDownSuite() {
	if s.conn != nil {
		s.conn.Close()
	}
	if s.grpcServer != nil {
		s.grpcServer.Stop()
	}
	s.identity.Close()
	if s.redisClient != nil {
		s.redisClient.Close()
	}
	if s.dbPool != nil {
		s.dbPool.Close()
	}
	s.ts.Teardown()
}

func (s *AuthHandlerGapiTestSuite) Test1_Register() {
	ctx := context.Background()
	req := &pb.RegisterRequest{
		Firstname:       "Auth",
		Lastname:        "Handler",
		Email:           s.email,
		Password:        s.password,
		ConfirmPassword: s.password,
	}

	res, err := s.client.RegisterUser(ctx, req)
	s.NoError(err)
	s.NotNil(res)
	s.Equal("success", res.Status)
	s.Equal(s.email, res.Data.Email)

	// Login requires a verified account — complete verification via gRPC
	// VerifyCode using the code Register cached in Redis.
	s.verifyUser(s.email)
}

func (s *AuthHandlerGapiTestSuite) Test2_Login() {
	ctx := context.Background()
	req := &pb.LoginRequest{
		Email:    s.email,
		Password: s.password,
	}

	res, err := s.client.LoginUser(ctx, req)
	s.NoError(err)
	s.NotNil(res)
	s.Equal("success", res.Status)
	s.NotEmpty(res.Data.AccessToken)
	s.accessToken = res.Data.AccessToken
}

func (s *AuthHandlerGapiTestSuite) Test4_LoginLockout() {
	ctx := context.Background()
	email := "locked.gapi@example.com"
	password := "wrongpassword"

	// Register user first
	regReq := &pb.RegisterRequest{
		Firstname:       "Locked",
		Lastname:        "Gapi",
		Email:           email,
		Password:        "correctpassword",
		ConfirmPassword: "correctpassword",
	}
	_, err := s.client.RegisterUser(ctx, regReq)
	s.NoError(err)

	s.verifyUser(email)

	loginReq := &pb.LoginRequest{
		Email:    email,
		Password: password,
	}

	// Fail login 5 times
	for i := 0; i < 5; i++ {
		_, err := s.client.LoginUser(ctx, loginReq)
		s.Error(err)
	}

	// 6th attempt should return error
	_, err = s.client.LoginUser(ctx, loginReq)
	s.Error(err)
	s.Contains(err.Error(), "Account temporarily locked")
}

func (s *AuthHandlerGapiTestSuite) Test3_GetMe() {
	s.Require().NotEmpty(s.accessToken)
	ctx := context.Background()

	tokenManager, _ := auth.NewManager("mysecret")
	userIdStr, err := tokenManager.ValidateToken(s.accessToken)
	s.NoError(err)

	userId, err := strconv.Atoi(userIdStr)
	s.NoError(err)

	res, err := s.client.GetMe(ctx, &pb.GetMeRequest{UserId: int32(userId)})
	s.NoError(err)
	s.NotNil(res)
	s.Equal("success", res.Status)
	s.Equal(s.email, res.Data.Email)
}

// verifyUser marks a just-registered user as verified via the gRPC VerifyCode
// call, mirroring the real email-verification flow.
func (s *AuthHandlerGapiTestSuite) verifyUser(email string) {
	ctx := context.Background()
	raw, err := s.redisClient.Get(ctx, "register:verify_code:"+email).Result()
	s.Require().NoError(err, "verification code should be cached after register")
	var code string
	s.Require().NoError(json.Unmarshal([]byte(raw), &code), "cached code is JSON-encoded")
	res, err := s.client.VerifyCode(ctx, &pb.VerifyCodeRequest{Code: code})
	s.Require().NoError(err)
	s.Require().Equal("success", res.Status)
}

func TestAuthHandlerGapiSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	suite.Run(t, new(AuthHandlerGapiTestSuite))
}
