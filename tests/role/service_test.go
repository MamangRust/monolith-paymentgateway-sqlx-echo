package role_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-role/repository"
	"github.com/MamangRust/monolith-payment-gateway-role/service"
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

type RoleServiceTestSuite struct {
	suite.Suite
	ts          *tests.TestSuite
	dbPool      *sqlx.DB
	redisClient *redis.Client
	roleService *service.Service
	roleID      int
}

func (s *RoleServiceTestSuite) SetupSuite() {
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
	repos := repository.NewRepositories(queries)

	logger.ResetInstance()
	lp := sdklog.NewLoggerProvider()
	log, _ := logger.NewLogger("test", lp)
	cacheMetrics, _ := observability.NewCacheMetrics("test")
	cacheStore := cache.NewCacheStore(s.redisClient, log, cacheMetrics)

	s.roleService = service.NewService(&service.Deps{
		Repositories: repos,
		Logger:       log,
		Cache:        cacheStore,
	})

	// Flush Redis to ensure isolation
	s.redisClient.FlushAll(s.ts.Ctx)
}

func (s *RoleServiceTestSuite) TearDownSuite() {
	if s.redisClient != nil {
		s.redisClient.Close()
	}
	if s.dbPool != nil {
		s.dbPool.Close()
	}
	if s.ts != nil {
		s.ts.Teardown()
	}
}

func (s *RoleServiceTestSuite) Test1_CreateRole() {
	ctx := context.Background()
	req := &requests.CreateRoleRequest{
		Name: fmt.Sprintf("Service Role-%d", time.Now().UnixNano()),
	}

	res, err := s.roleService.RoleCommand.CreateRole(ctx, req)
	s.NoError(err)
	s.NotNil(res)
	s.Equal(req.Name, res.RoleName)
	s.roleID = int(res.RoleID)
}

func (s *RoleServiceTestSuite) Test2_FindById() {
	s.Require().NotZero(s.roleID)
	ctx := context.Background()

	found, err := s.roleService.RoleQuery.FindById(ctx, s.roleID)
	s.NoError(err)
	s.NotNil(found)
	s.Equal(s.roleID, int(found.RoleID))
}

func (s *RoleServiceTestSuite) Test3_FindAll() {
	ctx := context.Background()

	req := &requests.FindAllRoles{
		Search:   "Role",
		Page:     1,
		PageSize: 10,
	}
	roles, total, err := s.roleService.RoleQuery.FindAll(ctx, req)
	s.NoError(err)
	s.NotNil(roles)
	s.NotZero(*total)
}

func (s *RoleServiceTestSuite) Test4_FindByActive() {
	ctx := context.Background()

	req := &requests.FindAllRoles{
		Search:   "Role",
		Page:     1,
		PageSize: 10,
	}
	roles, total, err := s.roleService.RoleQuery.FindByActiveRole(ctx, req)
	s.NoError(err)
	s.NotNil(roles)
	s.NotZero(*total)
}

func (s *RoleServiceTestSuite) Test5_UpdateRole() {
	s.Require().NotZero(s.roleID)
	ctx := context.Background()

	updateName := fmt.Sprintf("Updated Role-%d", time.Now().UnixNano())
	req := &requests.UpdateRoleRequest{
		ID:   &s.roleID,
		Name: updateName,
	}

	res, err := s.roleService.RoleCommand.UpdateRole(ctx, req)
	s.NoError(err)
	s.NotNil(res)
	s.Equal(updateName, res.RoleName)
}

func (s *RoleServiceTestSuite) Test6_TrashAndRestore() {
	s.Require().NotZero(s.roleID)
	ctx := context.Background()

	// Trash
	_, err := s.roleService.RoleCommand.TrashedRole(ctx, s.roleID)
	s.NoError(err)

	// Find By Trashed
	req := &requests.FindAllRoles{
		Page:     1,
		PageSize: 10,
	}
	trashed, total, err := s.roleService.RoleQuery.FindByTrashedRole(ctx, req)
	s.NoError(err)
	s.NotNil(trashed)
	s.NotZero(*total)

	// Restore
	_, err = s.roleService.RoleCommand.RestoreRole(ctx, s.roleID)
	s.NoError(err)
}

func (s *RoleServiceTestSuite) Test7_BulkOperations() {
	ctx := context.Background()

	// Trash for bulk test
	_, err := s.roleService.RoleCommand.TrashedRole(ctx, s.roleID)
	s.NoError(err)

	// Restore All
	success, err := s.roleService.RoleCommand.RestoreAllRole(ctx)
	s.NoError(err)
	s.True(success)

	// Trash again for delete all
	_, err = s.roleService.RoleCommand.TrashedRole(ctx, s.roleID)
	s.NoError(err)

	// Delete All Permanent
	success, err = s.roleService.RoleCommand.DeleteAllRolePermanent(ctx)
	s.NoError(err)
	s.True(success)

	s.roleID = 0
}

func (s *RoleServiceTestSuite) Test8_DeletePermanent() {
	ctx := context.Background()

	// Create another role
	req := &requests.CreateRoleRequest{
		Name: fmt.Sprintf("DeleteMe-%d", time.Now().UnixNano()),
	}
	res, err := s.roleService.RoleCommand.CreateRole(ctx, req)
	s.NoError(err)

	rid := int(res.RoleID)

	// Trash first
	_, err = s.roleService.RoleCommand.TrashedRole(ctx, rid)
	s.NoError(err)

	// Delete Permanent
	success, err := s.roleService.RoleCommand.DeleteRolePermanent(ctx, rid)
	s.NoError(err)
	s.True(success)
}

func (s *RoleServiceTestSuite) TestZZ_NotFound404() {
	ctx := context.Background()
	_, err := s.roleService.RoleQuery.FindById(ctx, 999999999)
	s.Require().Error(err)
	s.Contains(err.Error(), "not found")
}

func TestRoleServiceSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	suite.Run(t, new(RoleServiceTestSuite))
}
