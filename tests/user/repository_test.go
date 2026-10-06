package user_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	pbrole "github.com/MamangRust/monolith-payment-gateway-pb/role"
	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/cache"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/observability"
	"github.com/MamangRust/monolith-payment-gateway-user/repository"
	tests "github.com/MamangRust/monolith-payment-gateway-test"

	"github.com/jmoiron/sqlx"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/suite"
	sdklog "go.opentelemetry.io/otel/sdk/log"
)

type UserRepositoryTestSuite struct {
	suite.Suite
	ts          *tests.TestSuite
	dbPool      *sqlx.DB
	redisClient *redis.Client
	repo        repository.Repositories
	identity    *tests.IdentityStack
}

func (s *UserRepositoryTestSuite) SetupSuite() {
	ts, err := tests.SetupTestSuite()
	s.Require().NoError(err)
	s.ts = ts

	pool, err := sqlx.ConnectContext(s.ts.Ctx, "pgx", s.ts.DBURL)
	s.Require().NoError(err)
	s.dbPool = pool

	opts, err := redis.ParseURL(s.ts.RedisURL)
	s.Require().NoError(err)
	s.redisClient = redis.NewClient(opts)

	logger.ResetInstance()
	lp := sdklog.NewLoggerProvider()
	log, _ := logger.NewLogger("test", lp)
	cacheMetrics, _ := observability.NewCacheMetrics("test")
	cacheStore := cache.NewCacheStore(s.redisClient, log, cacheMetrics)

	s.identity, err = tests.StartIdentityStack(pool, cacheStore, log)
	s.Require().NoError(err)

	queries := pool
	s.repo = repository.NewRepositories(
		queries,
		pbrole.NewRoleServiceClient(s.identity.Role.Conn),
		pbrole.NewRoleCommandServiceClient(s.identity.Role.Conn),
	)
}

func (s *UserRepositoryTestSuite) TearDownSuite() {
	if s.identity != nil {
		s.identity.Close()
	}
	if s.redisClient != nil {
		s.redisClient.Close()
	}
	if s.dbPool != nil {
		s.dbPool.Close()
	}
	s.ts.Teardown()
}

func (s *UserRepositoryTestSuite) createSeedUser() (*db.CreateUserRow, error) {
	return s.repo.UserCommand().CreateUser(context.Background(), &requests.CreateUserRequest{
		FirstName:       "User",
		LastName:        "Tester",
		Email:           fmt.Sprintf("user.tester-%d@example.com", time.Now().UnixNano()),
		Password:        "password123",
		ConfirmPassword: "password123",
	})
}

func (s *UserRepositoryTestSuite) TestCreateUser() {
	ctx := context.Background()
	req := &requests.CreateUserRequest{
		FirstName:       "User",
		LastName:        "Tester",
		Email:           fmt.Sprintf("user.tester.%d@example.com", time.Now().UnixNano()),
		Password:        "password123",
		ConfirmPassword: "password123",
	}

	res, err := s.repo.UserCommand().CreateUser(ctx, req)
	s.NoError(err)
	s.NotNil(res)
}

func (s *UserRepositoryTestSuite) TestFindAllUsers() {
	_, err := s.createSeedUser()
	s.Require().NoError(err)
	ctx := context.Background()

	res, err := s.repo.UserQuery().FindAllUsers(ctx, &requests.FindAllUsers{
		Page:     1,
		PageSize: 10,
		Search:   "",
	})
	s.NoError(err)
	s.GreaterOrEqual(len(res), 1)
}

func (s *UserRepositoryTestSuite) TestFindById() {
	user, err := s.createSeedUser()
	s.Require().NoError(err)
	ctx := context.Background()

	found, err := s.repo.UserQuery().FindById(ctx, int(user.UserID))
	s.NoError(err)
	s.NotNil(found)
	s.Equal(user.UserID, found.UserID)
}

func (s *UserRepositoryTestSuite) TestFindByActive() {
	_, err := s.createSeedUser()
	s.Require().NoError(err)
	ctx := context.Background()

	res, err := s.repo.UserQuery().FindByActive(ctx, &requests.FindAllUsers{
		Page:     1,
		PageSize: 10,
		Search:   "",
	})
	s.NoError(err)
	s.GreaterOrEqual(len(res), 1)
}

func (s *UserRepositoryTestSuite) TestFindByTrashed() {
	user, err := s.createSeedUser()
	s.Require().NoError(err)
	ctx := context.Background()

	_, err = s.repo.UserCommand().TrashedUser(ctx, int(user.UserID))
	s.Require().NoError(err)

	res, err := s.repo.UserQuery().FindByTrashed(ctx, &requests.FindAllUsers{
		Page:     1,
		PageSize: 10,
		Search:   "",
	})
	s.NoError(err)
	s.GreaterOrEqual(len(res), 1)
}

func (s *UserRepositoryTestSuite) TestUpdateUser() {
	user, err := s.createSeedUser()
	s.Require().NoError(err)
	ctx := context.Background()

	id := int(user.UserID)
	req := &requests.UpdateUserRequest{
		UserID:          &id,
		FirstName:       "Updated",
		LastName:        "User",
		Email:           user.Email,
		Password:        "newpassword123",
		ConfirmPassword: "newpassword123",
	}

	res, err := s.repo.UserCommand().UpdateUser(ctx, req)
	s.NoError(err)
	s.NotNil(res)
}

func (s *UserRepositoryTestSuite) TestTrashUser() {
	user, err := s.createSeedUser()
	s.Require().NoError(err)
	ctx := context.Background()

	trashed, err := s.repo.UserCommand().TrashedUser(ctx, int(user.UserID))
	s.NoError(err)
	s.NotNil(trashed)
}

func (s *UserRepositoryTestSuite) TestRestoreUser() {
	user, err := s.createSeedUser()
	s.Require().NoError(err)
	ctx := context.Background()

	_, err = s.repo.UserCommand().TrashedUser(ctx, int(user.UserID))
	s.Require().NoError(err)

	restored, err := s.repo.UserCommand().RestoreUser(ctx, int(user.UserID))
	s.NoError(err)
	s.NotNil(restored)
}

func (s *UserRepositoryTestSuite) TestDeleteUserPermanent() {
	user, err := s.createSeedUser()
	s.Require().NoError(err)
	ctx := context.Background()

	_, err = s.repo.UserCommand().TrashedUser(ctx, int(user.UserID))
	s.Require().NoError(err)

	success, err := s.repo.UserCommand().DeleteUserPermanent(ctx, int(user.UserID))
	s.NoError(err)
	s.True(success)
}

func (s *UserRepositoryTestSuite) TestRestoreAllUser() {
	user, err := s.createSeedUser()
	s.Require().NoError(err)
	ctx := context.Background()

	_, err = s.repo.UserCommand().TrashedUser(ctx, int(user.UserID))
	s.Require().NoError(err)

	success, err := s.repo.UserCommand().RestoreAllUser(ctx)
	s.NoError(err)
	s.True(success)
}

func (s *UserRepositoryTestSuite) TestDeleteAllUserPermanent() {
	_, err := s.createSeedUser()
	s.Require().NoError(err)
	ctx := context.Background()

	success, err := s.repo.UserCommand().DeleteAllUserPermanent(ctx)
	s.NoError(err)
	s.True(success)
}

func TestUserRepositorySuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	suite.Run(t, new(UserRepositoryTestSuite))
}
