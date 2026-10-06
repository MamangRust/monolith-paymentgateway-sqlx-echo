package role_test

import (
	"context"
	"testing"

	pb "github.com/MamangRust/monolith-payment-gateway-pb/role"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-role/handler"
	"github.com/MamangRust/monolith-payment-gateway-role/repository"
	"github.com/MamangRust/monolith-payment-gateway-role/service"
	"github.com/MamangRust/monolith-payment-gateway-shared/cache"
	"github.com/MamangRust/monolith-payment-gateway-shared/observability"
	tests "github.com/MamangRust/monolith-payment-gateway-test"
	"github.com/jmoiron/sqlx"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/suite"
	sdklog "go.opentelemetry.io/otel/sdk/log"
	"google.golang.org/protobuf/types/known/emptypb"
)

type RoleGapiTestSuite struct {
	suite.Suite
	ts      *tests.TestSuite
	handler *handler.Handler
	dbPool  *sqlx.DB
	roleID  int32
}

func (s *RoleGapiTestSuite) SetupSuite() {
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

	s.handler = handler.NewHandler(svc)
}

func (s *RoleGapiTestSuite) TearDownSuite() {
	if s.dbPool != nil {
		s.dbPool.Close()
	}
	if s.ts != nil {
		s.ts.Teardown()
	}
}

func (s *RoleGapiTestSuite) Test1_RoleLifecycle() {
	ctx := context.Background()

	// Create
	req := &pb.CreateRoleRequest{
		Name: "Test Gapi Role",
	}
	res, err := s.handler.RoleCommand.CreateRole(ctx, req)
	s.NoError(err)
	s.NotNil(res)
	s.roleID = res.Data.Id
	s.Equal(req.Name, res.Data.Name)

	// FindById
	found, err := s.handler.RoleQuery.FindByIdRole(ctx, &pb.FindByIdRoleRequest{RoleId: s.roleID})
	s.NoError(err)
	s.NotNil(found)
	s.Equal(s.roleID, found.Data.Id)

	// Update
	updateReq := &pb.UpdateRoleRequest{
		Id:   s.roleID,
		Name: "Updated Gapi Role",
	}
	updated, err := s.handler.RoleCommand.UpdateRole(ctx, updateReq)
	s.NoError(err)
	s.NotNil(updated)
	s.Equal(updateReq.Name, updated.Data.Name)
}

func (s *RoleGapiTestSuite) Test2_QueryOperations() {
	ctx := context.Background()

	// FindAll
	all, err := s.handler.RoleQuery.FindAllRole(ctx, &pb.FindAllRoleRequest{
		Page:     1,
		PageSize: 10,
	})
	s.NoError(err)
	s.NotNil(all)
	s.GreaterOrEqual(all.PaginationMeta.TotalRecords, int32(1))

	// FindByActive
	active, err := s.handler.RoleQuery.FindByActive(ctx, &pb.FindAllRoleRequest{
		Page:     1,
		PageSize: 10,
	})
	s.NoError(err)
	s.NotNil(active)
	s.GreaterOrEqual(active.PaginationMeta.TotalRecords, int32(1))
}

func (s *RoleGapiTestSuite) Test3_TrashAndRestore() {
	ctx := context.Background()
	s.Require().NotZero(s.roleID)

	// Trash
	trashed, err := s.handler.RoleCommand.TrashedRole(ctx, &pb.FindByIdRoleRequest{RoleId: s.roleID})
	s.NoError(err)
	s.NotNil(trashed)

	// FindByTrashed
	trashedList, err := s.handler.RoleQuery.FindByTrashed(ctx, &pb.FindAllRoleRequest{
		Page:     1,
		PageSize: 10,
	})
	s.NoError(err)
	s.NotNil(trashedList)
	s.GreaterOrEqual(trashedList.PaginationMeta.TotalRecords, int32(1))

	// Restore
	restored, err := s.handler.RoleCommand.RestoreRole(ctx, &pb.FindByIdRoleRequest{RoleId: s.roleID})
	s.NoError(err)
	s.NotNil(restored)
}

func (s *RoleGapiTestSuite) Test4_BulkOperations() {
	ctx := context.Background()

	// Restore All
	ok, err := s.handler.RoleCommand.RestoreAllRole(ctx, &emptypb.Empty{})
	s.NoError(err)
	s.NotNil(ok)

	// Delete All Permanent
	ok, err = s.handler.RoleCommand.DeleteAllRolePermanent(ctx, &emptypb.Empty{})
	s.NoError(err)
	s.NotNil(ok)
}

func TestRoleGapiSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	suite.Run(t, new(RoleGapiTestSuite))
}
