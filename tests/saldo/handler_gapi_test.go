package saldo_test

import (
	"context"
	"fmt"
	"net"
	"testing"
	"time"

	card_handler "github.com/MamangRust/monolith-payment-gateway-card/handler"
	card_repository "github.com/MamangRust/monolith-payment-gateway-card/repository"
	card_service "github.com/MamangRust/monolith-payment-gateway-card/service"
	pbcard "github.com/MamangRust/monolith-payment-gateway-pb/card"
	pbrole "github.com/MamangRust/monolith-payment-gateway-pb/role"
	pbsaldo "github.com/MamangRust/monolith-payment-gateway-pb/saldo"
	pbuser "github.com/MamangRust/monolith-payment-gateway-pb/user"
	"github.com/MamangRust/monolith-payment-gateway-pkg/hash"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-saldo/handler"
	"github.com/MamangRust/monolith-payment-gateway-saldo/repository"
	"github.com/MamangRust/monolith-payment-gateway-saldo/service"
	"github.com/MamangRust/monolith-payment-gateway-shared/cache"
	"github.com/MamangRust/monolith-payment-gateway-shared/observability"
	user_handler "github.com/MamangRust/monolith-payment-gateway-user/handler"
	user_repository "github.com/MamangRust/monolith-payment-gateway-user/repository"
	user_service "github.com/MamangRust/monolith-payment-gateway-user/service"
	tests "github.com/MamangRust/monolith-payment-gateway-test"
	"github.com/jmoiron/sqlx"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/suite"
	sdklog "go.opentelemetry.io/otel/sdk/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type SaldoGapiTestSuite struct {
	suite.Suite
	ts         *tests.TestSuite
	dbPool     *sqlx.DB
	saldoH     handler.Handler
	cardH      card_handler.Handler
	userH      user_handler.Handler
	cardServer *grpc.Server
	userID     int32
	cardID     int32
	cardNumber string
	saldoID    int32
	identity   *tests.IdentityStack
}

func (s *SaldoGapiTestSuite) SetupSuite() {
	ts, err := tests.SetupTestSuite()
	s.Require().NoError(err)
	s.ts = ts

	pool, err := sqlx.ConnectContext(s.ts.Ctx, "pgx", s.ts.DBURL)
	s.Require().NoError(err)
	s.dbPool = pool

	opts, err := redis.ParseURL(s.ts.RedisURL)
	s.Require().NoError(err)
	redisClient := redis.NewClient(opts)

	logger.ResetInstance()
	lp := sdklog.NewLoggerProvider()
	log, _ := logger.NewLogger("test", lp)
	hasher := hash.NewHashingPassword()
	cacheMetrics, _ := observability.NewCacheMetrics("test")
	cacheStore := cache.NewCacheStore(redisClient, log, cacheMetrics)

	s.identity, err = tests.StartIdentityStack(pool, cacheStore, log)
	s.Require().NoError(err)

	queries := pool

	// The saldo repository validates cards through the card adapter, so the
	// card service is exposed over an in-process gRPC server. grpc.NewClient
	// dials lazily, so the listener/conn are created before the repositories
	// that use the adapter, and the service is registered before Serve.
	lis, err := net.Listen("tcp", "localhost:0")
	s.Require().NoError(err)

	cardConn, err := grpc.NewClient(lis.Addr().String(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	s.Require().NoError(err)

	repos := repository.NewRepositories(
		queries,
		pbcard.NewCardQueryServiceClient(cardConn),
		pbcard.NewCardCommandServiceClient(cardConn),
	)
	cardRepos := card_repository.NewRepositories(queries, pbuser.NewUserQueryServiceClient(s.identity.User.Conn))
	userRepos := user_repository.NewRepositories(
		queries,
		pbrole.NewRoleServiceClient(s.identity.Role.Conn),
		pbrole.NewRoleCommandServiceClient(s.identity.Role.Conn),
	)

	saldoSvc := service.NewService(&service.Deps{
		Cache:        cacheStore,
		Repositories: repos,
		Logger:       log,
	})

	cardSvc := card_service.NewService(&card_service.Deps{
		Cache:        cacheStore,
		Repositories: cardRepos,
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

	s.saldoH = handler.NewHandler(saldoSvc)
	s.cardH = card_handler.NewHandler(cardSvc)
	s.userH = user_handler.NewHandler(userSvc)

	s.cardServer = grpc.NewServer()
	pbcard.RegisterCardQueryServiceServer(s.cardServer, s.cardH)
	pbcard.RegisterCardCommandServiceServer(s.cardServer, s.cardH)

	go func() {
		_ = s.cardServer.Serve(lis)
	}()

	// Create user
	ctx := context.Background()
	userRes, err := s.userH.Create(ctx, &pbuser.CreateUserRequest{
		Firstname:       "Saldo",
		Lastname:        "User",
		Email:           fmt.Sprintf("saldo.user.%d@example.com", time.Now().UnixNano()),
		Password:        "Password123!",
		ConfirmPassword: "Password123!",
	})
	s.Require().NoError(err)
	s.userID = userRes.Data.Id

	// Create card
	cardRes, err := s.cardH.CreateCard(ctx, &pbcard.CreateCardRequest{
		UserId:       s.userID,
		CardType:     "debit",
		ExpireDate:   timestamppb.New(time.Now().AddDate(5, 0, 0)),
		Cvv:          "123",
		CardProvider: "visa",
	})
	s.Require().NoError(err)
	s.cardID = cardRes.Data.Id
	s.cardNumber = cardRes.Data.CardNumber
}

func (s *SaldoGapiTestSuite) TearDownSuite() {
	if s.cardServer != nil {
		s.cardServer.Stop()
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

func (s *SaldoGapiTestSuite) Test1_SaldoLifecycle() {
	ctx := context.Background()

	// Create
	createReq := &pbsaldo.CreateSaldoRequest{
		CardNumber:   s.cardNumber,
		TotalBalance: 100000,
	}
	res, err := s.saldoH.CreateSaldo(ctx, createReq)
	s.NoError(err)
	s.NotNil(res)
	s.saldoID = res.Data.SaldoId
	s.Equal(int32(100000), res.Data.TotalBalance)

	// FindById
	resF, err := s.saldoH.FindByIdSaldo(ctx, &pbsaldo.FindByIdSaldoRequest{SaldoId: s.saldoID})
	s.NoError(err)
	s.Equal(int32(100000), resF.Data.TotalBalance)

	// Update
	updateReq := &pbsaldo.UpdateSaldoRequest{
		SaldoId:      s.saldoID,
		CardNumber:   s.cardNumber,
		TotalBalance: 200000,
	}
	resU, err := s.saldoH.UpdateSaldo(ctx, updateReq)
	s.NoError(err)
	s.Equal(int32(200000), resU.Data.TotalBalance)
}

func (s *SaldoGapiTestSuite) Test2_QueryOperations() {
	ctx := context.Background()
	s.Require().NotZero(s.saldoID)

	// FindAll
	allReq := &pbsaldo.FindAllSaldoRequest{Page: 1, PageSize: 10}
	resA, err := s.saldoH.FindAllSaldo(ctx, allReq)
	s.NoError(err)
	s.GreaterOrEqual(resA.PaginationMeta.TotalRecords, int32(1))

	// FindByActive
	resAc, err := s.saldoH.FindByActive(ctx, allReq)
	s.NoError(err)
	s.GreaterOrEqual(resAc.PaginationMeta.TotalRecords, int32(1))

	// FindByCardNumber
	resC, err := s.saldoH.FindByCardNumber(ctx, &pbcard.FindByCardNumberRequest{CardNumber: s.cardNumber})
	s.NoError(err)
	s.Equal(s.cardNumber, resC.Data.CardNumber)
}

func (s *SaldoGapiTestSuite) Test3_TrashAndRestore() {
	ctx := context.Background()
	s.Require().NotZero(s.saldoID)

	// Trash
	resT, err := s.saldoH.TrashedSaldo(ctx, &pbsaldo.FindByIdSaldoRequest{SaldoId: s.saldoID})
	s.NoError(err)
	s.NotNil(resT)

	// FindByTrashed
	resTL, err := s.saldoH.FindByTrashed(ctx, &pbsaldo.FindAllSaldoRequest{Page: 1, PageSize: 10})
	s.NoError(err)
	s.GreaterOrEqual(resTL.PaginationMeta.TotalRecords, int32(1))

	// Restore
	resR, err := s.saldoH.RestoreSaldo(ctx, &pbsaldo.FindByIdSaldoRequest{SaldoId: s.saldoID})
	s.NoError(err)
	s.NotNil(resR)
}

func (s *SaldoGapiTestSuite) Test4_BulkOperations() {
	ctx := context.Background()

	// Restore All
	resR, err := s.saldoH.RestoreAllSaldo(ctx, &emptypb.Empty{})
	s.NoError(err)
	s.Equal("success", resR.Status)

	// Delete All Permanent
	resD, err := s.saldoH.DeleteAllSaldoPermanent(ctx, &emptypb.Empty{})
	s.NoError(err)
	s.Equal("success", resD.Status)
}

func TestSaldoGapiSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	suite.Run(t, new(SaldoGapiTestSuite))
}
