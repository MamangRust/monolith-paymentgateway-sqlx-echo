package merchant_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/MamangRust/monolith-payment-gateway-merchant/handler"
	"github.com/MamangRust/monolith-payment-gateway-merchant/repository"
	"github.com/MamangRust/monolith-payment-gateway-merchant/service"
	pb "github.com/MamangRust/monolith-payment-gateway-pb/merchant"
	pbrole "github.com/MamangRust/monolith-payment-gateway-pb/role"
	pbuser "github.com/MamangRust/monolith-payment-gateway-pb/user"
	"github.com/MamangRust/monolith-payment-gateway-pkg/hash"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
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
	"google.golang.org/protobuf/types/known/emptypb"
)

type MerchantGapiTestSuite struct {
	suite.Suite
	ts         *tests.TestSuite
	dbPool     *sqlx.DB
	merchantH  handler.Handler
	userH      user_handler.Handler
	userID     int32
	merchantID int32
	identity   *tests.IdentityStack
}

func (s *MerchantGapiTestSuite) SetupSuite() {
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

	s.merchantH = handler.NewHandler(merchantSvc)
	s.userH = user_handler.NewHandler(userSvc)

	// Create a user for testing
	ctx := context.Background()
	userRes, err := s.userH.Create(ctx, &pbuser.CreateUserRequest{
		Firstname:       "Merchant",
		Lastname:        "User",
		Email:           fmt.Sprintf("merchant.user.%d@example.com", time.Now().UnixNano()),
		Password:        "Password123!",
		ConfirmPassword: "Password123!",
	})
	s.Require().NoError(err)
	s.userID = userRes.Data.Id
}

func (s *MerchantGapiTestSuite) TearDownSuite() {
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

func (s *MerchantGapiTestSuite) Test1_MerchantLifecycle() {
	ctx := context.Background()

	// Create
	createReq := &pb.CreateMerchantRequest{
		Name:   "Test Merchant",
		UserId: s.userID,
	}
	res, err := s.merchantH.CreateMerchant(ctx, createReq)
	s.NoError(err)
	s.NotNil(res)
	s.merchantID = res.Data.Id
	s.Equal("Test Merchant", res.Data.Name)

	// FindById
	findReq := &pb.FindByIdMerchantRequest{MerchantId: s.merchantID}
	resF, err := s.merchantH.FindByIdMerchant(ctx, findReq)
	s.NoError(err)
	s.Equal("Test Merchant", resF.Data.Name)

	// Update
	updateReq := &pb.UpdateMerchantRequest{
		MerchantId: s.merchantID,
		Name:       "Updated Merchant",
		UserId:     s.userID,
		Status:     "active",
	}
	resU, err := s.merchantH.UpdateMerchant(ctx, updateReq)
	s.NoError(err)
	s.Equal("Updated Merchant", resU.Data.Name)

	// UpdateMerchantStatus
	statusReq := &pb.UpdateMerchantStatusRequest{
		MerchantId: s.merchantID,
		Status:     "inactive",
	}
	resS, err := s.merchantH.UpdateMerchantStatus(ctx, statusReq)
	s.NoError(err)
	s.Equal("inactive", resS.Data.Status)

	// UpdateMerchantStatus with invalid merchant ID
	_, err = s.merchantH.UpdateMerchantStatus(ctx, &pb.UpdateMerchantStatusRequest{MerchantId: 0, Status: "active"})
	s.Error(err)
}

func (s *MerchantGapiTestSuite) Test2_QueryOperations() {
	ctx := context.Background()
	s.Require().NotZero(s.merchantID)

	// FindAll
	allReq := &pb.FindAllMerchantRequest{Page: 1, PageSize: 10}
	resA, err := s.merchantH.FindAllMerchant(ctx, allReq)
	s.NoError(err)
	s.GreaterOrEqual(resA.PaginationMeta.TotalRecords, int32(1))

	// FindByActive
	resAc, err := s.merchantH.FindByActive(ctx, allReq)
	s.NoError(err)
	s.GreaterOrEqual(resAc.PaginationMeta.TotalRecords, int32(1))
}

func (s *MerchantGapiTestSuite) Test3_TrashAndRestore() {
	ctx := context.Background()
	s.Require().NotZero(s.merchantID)

	// Trash
	resT, err := s.merchantH.TrashedMerchant(ctx, &pb.FindByIdMerchantRequest{MerchantId: s.merchantID})
	s.NoError(err)
	s.NotNil(resT)

	// FindByTrashed
	resTL, err := s.merchantH.FindByTrashed(ctx, &pb.FindAllMerchantRequest{Page: 1, PageSize: 10})
	s.NoError(err)
	s.GreaterOrEqual(resTL.PaginationMeta.TotalRecords, int32(1))

	// Restore
	resR, err := s.merchantH.RestoreMerchant(ctx, &pb.FindByIdMerchantRequest{MerchantId: s.merchantID})
	s.NoError(err)
	s.NotNil(resR)
}

func (s *MerchantGapiTestSuite) Test4_BulkOperations() {
	ctx := context.Background()

	// Restore All
	resR, err := s.merchantH.RestoreAllMerchant(ctx, &emptypb.Empty{})
	s.NoError(err)
	s.Equal("success", resR.Status)

	// Delete All Permanent
	resD, err := s.merchantH.DeleteAllMerchantPermanent(ctx, &emptypb.Empty{})
	s.NoError(err)
	s.Equal("success", resD.Status)
}

func TestMerchantGapiSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	suite.Run(t, new(MerchantGapiTestSuite))
}
