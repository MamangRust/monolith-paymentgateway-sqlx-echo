package topup_test

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
	pb "github.com/MamangRust/monolith-payment-gateway-pb/topup"
	pbuser "github.com/MamangRust/monolith-payment-gateway-pb/user"
	"github.com/MamangRust/monolith-payment-gateway-pkg/hash"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	saldo_handler "github.com/MamangRust/monolith-payment-gateway-saldo/handler"
	saldo_repository "github.com/MamangRust/monolith-payment-gateway-saldo/repository"
	saldo_service "github.com/MamangRust/monolith-payment-gateway-saldo/service"
	"github.com/MamangRust/monolith-payment-gateway-shared/cache"
	"github.com/MamangRust/monolith-payment-gateway-shared/observability"
	"github.com/MamangRust/monolith-payment-gateway-topup/handler"
	"github.com/MamangRust/monolith-payment-gateway-topup/repository"
	"github.com/MamangRust/monolith-payment-gateway-topup/service"
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
	"google.golang.org/grpc/test/bufconn"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type TopupGapiTestSuite struct {
	suite.Suite
	ts         *tests.TestSuite
	dbPool     *sqlx.DB
	topupH     handler.Handler
	userID     int32
	cardID     int32
	cardNumber string
	topupID    int32
	grpcServer *grpc.Server
	lis        *bufconn.Listener
	conn       *grpc.ClientConn
	identity   *tests.IdentityStack
}

func (s *TopupGapiTestSuite) SetupSuite() {
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

	logger.ResetInstance()
	lp := sdklog.NewLoggerProvider()
	log, _ := logger.NewLogger("test", lp)
	hasher := hash.NewHashingPassword()
	cacheMetrics, _ := observability.NewCacheMetrics("test")
	cacheStore := cache.NewCacheStore(redisClient, log, cacheMetrics)

	s.identity, err = tests.StartIdentityStack(pool, cacheStore, log)
	s.Require().NoError(err)

	// bufconn gRPC server for the card/saldo adapters. grpc dials lazily, so
	// create the listener/conn first, build the services, register, then serve.
	s.lis = bufconn.Listen(1024 * 1024)
	s.grpcServer = grpc.NewServer()
	s.conn, err = grpc.DialContext(context.Background(), "bufnet",
		grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
			return s.lis.Dial()
		}),
		grpc.WithInsecure())
	s.Require().NoError(err)

	cardQueryClient := pbcard.NewCardQueryServiceClient(s.conn)
	cardCommandClient := pbcard.NewCardCommandServiceClient(s.conn)
	saldoQueryClient := pbsaldo.NewSaldoQueryServiceClient(s.conn)
	saldoCommandClient := pbsaldo.NewSaldoCommandServiceClient(s.conn)

	// Dependency services handlers (local implementation)
	cardSvc := card_service.NewService(&card_service.Deps{
		Cache:        cacheStore,
		Repositories: card_repository.NewRepositories(queries, pbuser.NewUserQueryServiceClient(s.identity.User.Conn)),
		Logger:       log,
		Kafka:        nil,
	})
	cardH := card_handler.NewHandler(cardSvc)

	saldoSvc := saldo_service.NewService(&saldo_service.Deps{
		Cache:        cacheStore,
		Repositories: saldo_repository.NewRepositories(queries, cardQueryClient, cardCommandClient),
		Logger:       log,
	})
	saldoH := saldo_handler.NewHandler(saldoSvc)

	userSvc := user_service.NewService(&user_service.Deps{
		Cache: cacheStore,
		Repositories: user_repository.NewRepositories(
			queries,
			pbrole.NewRoleServiceClient(s.identity.Role.Conn),
			pbrole.NewRoleCommandServiceClient(s.identity.Role.Conn),
		),
		Hash:            hasher,
		Logger:          log,
		RoleAdapter:     s.identity.Role.RoleAdapter,
		UserRoleAdapter: s.identity.Role.UserRoleAdapter,
	})
	userH := user_handler.NewHandler(userSvc)

	pbcard.RegisterCardQueryServiceServer(s.grpcServer, cardH)
	pbcard.RegisterCardCommandServiceServer(s.grpcServer, cardH)
	pbsaldo.RegisterSaldoQueryServiceServer(s.grpcServer, saldoH)
	pbsaldo.RegisterSaldoCommandServiceServer(s.grpcServer, saldoH)

	go func() {
		if err := s.grpcServer.Serve(s.lis); err != nil {
		}
	}()

	// Topup Repositories with raw clients
	repos := repository.NewRepositories(queries, cardQueryClient, cardCommandClient, saldoQueryClient, saldoCommandClient)

	topupSvc := service.NewService(&service.Deps{
		Kafka:        nil,
		Cache:        cacheStore,
		Repositories: repos,
		Logger:       log,
	})
	s.topupH = handler.NewHandler(topupSvc)

	// Create test environment (User, Card, Saldo)
	ctx := context.Background()
	userRes, err := userH.Create(ctx, &pbuser.CreateUserRequest{
		Firstname:       "Topup",
		Lastname:        "Gapi",
		Email:           fmt.Sprintf("topup.gapi.%d@example.com", time.Now().UnixNano()),
		Password:        "Password123!",
		ConfirmPassword: "Password123!",
	})
	s.Require().NoError(err)
	s.userID = userRes.Data.Id

	cardRes, err := cardH.CreateCard(ctx, &pbcard.CreateCardRequest{
		UserId:       s.userID,
		CardType:     "debit",
		ExpireDate:   timestamppb.New(time.Now().AddDate(5, 0, 0)),
		Cvv:          "123",
		CardProvider: "visa",
	})
	s.Require().NoError(err)
	s.cardID = cardRes.Data.Id
	s.cardNumber = cardRes.Data.CardNumber

	_, err = saldoH.CreateSaldo(ctx, &pbsaldo.CreateSaldoRequest{
		CardNumber:   s.cardNumber,
		TotalBalance: 100000,
	})
	s.Require().NoError(err)
}

func (s *TopupGapiTestSuite) TearDownSuite() {
	if s.identity != nil {
		s.identity.Close()
	}
	if s.conn != nil {
		s.conn.Close()
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

func (s *TopupGapiTestSuite) Test1_TopupLifecycle() {
	ctx := context.Background()

	// Create Topup
	createReq := &pb.CreateTopupRequest{
		CardNumber:  s.cardNumber,
		TopupAmount: 50000,
		TopupMethod: "bca",
	}
	res, err := s.topupH.CreateTopup(ctx, createReq)
	s.Require().NoError(err)
	s.Require().NotNil(res)
	s.topupID = res.Data.Id
	s.Equal(int32(50000), res.Data.TopupAmount)

	// FindById
	resF, err := s.topupH.FindByIdTopup(ctx, &pb.FindByIdTopupRequest{TopupId: s.topupID})
	s.Require().NoError(err)
	s.Equal(int32(50000), resF.Data.TopupAmount)

	// Update Topup
	updateReq := &pb.UpdateTopupRequest{
		TopupId:     s.topupID,
		CardNumber:  s.cardNumber,
		TopupAmount: 75000,
		TopupMethod: "mandiri",
	}
	resU, err := s.topupH.UpdateTopup(ctx, updateReq)
	s.Require().NoError(err)
	s.Equal(int32(75000), resU.Data.TopupAmount)
}

func (s *TopupGapiTestSuite) Test2_QueryOperations() {
	ctx := context.Background()
	s.Require().NotZero(s.topupID)

	allReq := &pb.FindAllTopupRequest{Page: 1, PageSize: 10}

	// FindAll
	resA, err := s.topupH.FindAllTopup(ctx, allReq)
	s.NoError(err)
	s.GreaterOrEqual(resA.PaginationMeta.TotalRecords, int32(1))

	// FindByActive
	resAc, err := s.topupH.FindByActive(ctx, allReq)
	s.NoError(err)
	s.GreaterOrEqual(resAc.PaginationMeta.TotalRecords, int32(1))
}

func (s *TopupGapiTestSuite) Test3_TrashAndRestore() {
	ctx := context.Background()
	s.Require().NotZero(s.topupID)

	// Trashed
	resT, err := s.topupH.TrashedTopup(ctx, &pb.FindByIdTopupRequest{TopupId: s.topupID})
	s.NoError(err)
	s.NotNil(resT)

	// FindByTrashed
	resTL, err := s.topupH.FindByTrashed(ctx, &pb.FindAllTopupRequest{Page: 1, PageSize: 10})
	s.NoError(err)
	s.GreaterOrEqual(resTL.PaginationMeta.TotalRecords, int32(1))

	// Restore
	resR, err := s.topupH.RestoreTopup(ctx, &pb.FindByIdTopupRequest{TopupId: s.topupID})
	s.NoError(err)
	s.NotNil(resR)
}

func (s *TopupGapiTestSuite) Test4_BulkOperations() {
	ctx := context.Background()

	// Restore All
	resR, err := s.topupH.RestoreAllTopup(ctx, &emptypb.Empty{})
	s.NoError(err)
	s.Equal("success", resR.Status)

	// Delete All Permanent
	resD, err := s.topupH.DeleteAllTopupPermanent(ctx, &emptypb.Empty{})
	s.NoError(err)
	s.Equal("success", resD.Status)
}

func TestTopupGapiSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	suite.Run(t, new(TopupGapiTestSuite))
}
