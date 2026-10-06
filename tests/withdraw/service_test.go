package withdraw_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	card_repo_impl "github.com/MamangRust/monolith-payment-gateway-card/repository"
	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	saldo_repo_impl "github.com/MamangRust/monolith-payment-gateway-saldo/repository"
	"github.com/MamangRust/monolith-payment-gateway-shared/cache"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/observability"
	user_repo_impl "github.com/MamangRust/monolith-payment-gateway-user/repository"
	"github.com/MamangRust/monolith-payment-gateway-withdraw/repository"
	"github.com/MamangRust/monolith-payment-gateway-withdraw/service"
	tests "github.com/MamangRust/monolith-payment-gateway-test"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jmoiron/sqlx"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/suite"
	sdklog "go.opentelemetry.io/otel/sdk/log"
)

const (
	getSaldoByCardNumberSQL = `
SELECT *
FROM saldos
WHERE card_number = $1
  AND deleted_at IS NULL
ORDER BY saldo_id DESC
LIMIT 1
`

	getUserEmailByCardNumberSQL = `
SELECT c.card_id, u.email, c.user_id, c.card_number, c.card_type, c.expire_date, c.cvv, c.card_provider, c.created_at, c.updated_at
FROM cards c
    JOIN users u ON u.user_id = c.user_id
WHERE
    c.card_number = $1
    AND c.deleted_at IS NULL
    AND u.deleted_at IS NULL
`

	updateSaldoBalanceSQL = `
UPDATE saldos
SET
    total_balance = $2,
    updated_at = current_timestamp
WHERE
    card_number = $1
    AND deleted_at IS NULL
RETURNING
    saldo_id,
    card_number,
    total_balance,
    withdraw_amount,
    withdraw_time,
    created_at,
    updated_at
`

	updateSaldoWithdrawSQL = `
UPDATE saldos
SET
    withdraw_amount = $2,
    total_balance = total_balance - $2,
    withdraw_time = $3,
    updated_at = current_timestamp
WHERE
    card_number = $1
    AND deleted_at IS NULL
    AND total_balance >= $2
RETURNING
    saldo_id,
    card_number,
    total_balance,
    withdraw_amount,
    withdraw_time,
    created_at,
    updated_at
`
)

type testCardRepo struct {
	db *sqlx.DB
}

func (r *testCardRepo) FindUserCardByCardNumber(ctx context.Context, card_number string) (*db.GetUserEmailByCardNumberRow, error) {
	var row db.GetUserEmailByCardNumberRow
	if err := sqlx.GetContext(ctx, r.db, &row, getUserEmailByCardNumberSQL, card_number); err != nil {
		return nil, err
	}
	return &row, nil
}

type testSaldoRepo struct {
	db *sqlx.DB
}

func (r *testSaldoRepo) FindByCardNumber(ctx context.Context, card_number string) (*db.Saldo, error) {
	var saldo db.Saldo
	if err := sqlx.GetContext(ctx, r.db, &saldo, getSaldoByCardNumberSQL, card_number); err != nil {
		return nil, err
	}
	return &saldo, nil
}

func (r *testSaldoRepo) UpdateSaldoBalance(ctx context.Context, request *requests.UpdateSaldoBalance) (*db.UpdateSaldoBalanceRow, error) {
	var row db.UpdateSaldoBalanceRow
	if err := sqlx.GetContext(ctx, r.db, &row, updateSaldoBalanceSQL, request.CardNumber, int32(request.TotalBalance)); err != nil {
		return nil, err
	}
	return &row, nil
}

func (r *testSaldoRepo) UpdateSaldoWithdraw(ctx context.Context, request *requests.UpdateSaldoWithdraw) (*db.UpdateSaldoWithdrawRow, error) {
	withdrawTime := pgtype.Timestamp{}
	if request.WithdrawTime != nil {
		withdrawTime.Time = *request.WithdrawTime
		withdrawTime.Valid = true
	}

	var withdrawAmount *int32
	if request.WithdrawAmount != nil {
		amount := int32(*request.WithdrawAmount)
		withdrawAmount = &amount
	}

	var row db.UpdateSaldoWithdrawRow
	if err := sqlx.GetContext(ctx, r.db, &row, updateSaldoWithdrawSQL, request.CardNumber, withdrawAmount, withdrawTime); err != nil {
		return nil, err
	}
	return &row, nil
}

type WithdrawServiceTestSuite struct {
	suite.Suite
	ts              *tests.TestSuite
	withdrawService service.Service
	dbPool          *sqlx.DB
	deps            *grpcDeps
	withdrawID      int
	cardNumber      string
	userID          int
}

func (s *WithdrawServiceTestSuite) SetupSuite() {
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
	cacheMetrics, _ := observability.NewCacheMetrics("test")
	cacheStore := cache.NewCacheStore(redisClient, log, cacheMetrics)

	deps, err := startGRPCDeps(queries, cacheStore, log)
	s.Require().NoError(err)
	s.deps = deps

	repos := repository.NewRepositories(
		queries,
		deps.cardQueryClient,
		deps.cardCommandClient,
		deps.saldoQueryClient,
		deps.saldoCommandClient,
		testDailyWithdrawLimit,
	)

	s.withdrawService = service.NewService(&service.Deps{
		Kafka:        nil,
		Repositories: repos,
		Logger:       log,
		Cache:        cacheStore,
	})

	// Seed User
	userRepo := user_repo_impl.NewUserCommandRepository(queries)
	user, err := userRepo.CreateUser(context.Background(), &requests.CreateUserRequest{
		FirstName: "Withdraw",
		LastName:  "Tester",
		Email:     fmt.Sprintf("withdraw.tester.%d@example.com", time.Now().UnixNano()),
		Password:  "password123",
	})
	s.Require().NoError(err)
	s.userID = int(user.UserID)

	// Seed Card
	cardCmdRepo := card_repo_impl.NewCardCommandRepository(queries)
	card, err := cardCmdRepo.CreateCard(context.Background(), &requests.CreateCardRequest{
		UserID:       s.userID,
		CardType:     "debit",
		ExpireDate:   time.Now().AddDate(5, 0, 0),
		CVV:          "123",
		CardProvider: "visa",
	})
	s.Require().NoError(err)
	s.cardNumber = card.CardNumber

	// Seed Saldo
	saldoCmdRepo := saldo_repo_impl.NewSaldoCommandRepository(queries)
	_, err = saldoCmdRepo.CreateSaldo(context.Background(), &requests.CreateSaldoRequest{
		CardNumber:   s.cardNumber,
		TotalBalance: 1000000,
	})
	s.Require().NoError(err)
}

func (s *WithdrawServiceTestSuite) TearDownSuite() {
	if s.deps != nil {
		s.deps.Close()
	}
	s.dbPool.Close()
	s.ts.Teardown()
}

func (s *WithdrawServiceTestSuite) Test1_WithdrawLifecycle() {
	ctx := context.Background()

	// Create Withdraw
	createReq := &requests.CreateWithdrawRequest{
		CardNumber:     s.cardNumber,
		WithdrawAmount: 100000,
		WithdrawTime:   time.Now(),
	}
	res, err := s.withdrawService.Create(ctx, createReq)
	s.NoError(err)
	s.NotNil(res)
	s.withdrawID = int(res.WithdrawID)
	s.Equal("success", res.Status)

	// FindById
	found, err := s.withdrawService.FindById(ctx, s.withdrawID)
	s.NoError(err)
	s.NotNil(found)
	s.Equal(int32(s.withdrawID), found.WithdrawID)

	// Update Withdraw
	updateReq := &requests.UpdateWithdrawRequest{
		WithdrawID:     &s.withdrawID,
		CardNumber:     s.cardNumber,
		WithdrawAmount: 200000,
		WithdrawTime:   time.Now(),
	}
	updated, err := s.withdrawService.Update(ctx, updateReq)
	s.NoError(err)
	s.NotNil(updated)
	s.Equal(int32(200000), updated.WithdrawAmount)
}

func (s *WithdrawServiceTestSuite) Test2_QueryOperations() {
	ctx := context.Background()

	// FindAll
	all, total, err := s.withdrawService.FindAll(ctx, &requests.FindAllWithdraws{
		Page:     1,
		PageSize: 10,
	})
	s.NoError(err)
	s.NotNil(all)
	s.GreaterOrEqual(*total, 1)

	// FindByActive
	active, totalActive, err := s.withdrawService.FindByActive(ctx, &requests.FindAllWithdraws{
		Page:     1,
		PageSize: 10,
	})
	s.NoError(err)
	s.NotNil(active)
	s.GreaterOrEqual(*totalActive, 1)

	// FindByCardNumber
	byCard, totalCard, err := s.withdrawService.FindAllByCardNumber(ctx, &requests.FindAllWithdrawCardNumber{
		CardNumber: s.cardNumber,
		Page:       1,
		PageSize:   10,
	})
	s.NoError(err)
	s.NotNil(byCard)
	s.GreaterOrEqual(*totalCard, 1)
}

func (s *WithdrawServiceTestSuite) Test3_TrashAndRestore() {
	ctx := context.Background()
	s.Require().NotZero(s.withdrawID)

	// Trash
	trashed, err := s.withdrawService.TrashedWithdraw(ctx, s.withdrawID)
	s.NoError(err)
	s.NotNil(trashed)

	// FindByTrashed
	trashedList, totalTrashed, err := s.withdrawService.FindByTrashed(ctx, &requests.FindAllWithdraws{
		Page:     1,
		PageSize: 10,
	})
	s.NoError(err)
	s.NotNil(trashedList)
	s.GreaterOrEqual(*totalTrashed, 1)

	// Restore
	restored, err := s.withdrawService.RestoreWithdraw(ctx, s.withdrawID)
	s.NoError(err)
	s.NotNil(restored)
}

func (s *WithdrawServiceTestSuite) Test4_BulkOperations() {
	ctx := context.Background()

	// Restore All
	ok, err := s.withdrawService.RestoreAllWithdraw(ctx)
	s.NoError(err)
	s.True(ok)

	// Delete All Permanent
	ok, err = s.withdrawService.DeleteAllWithdrawPermanent(ctx)
	s.NoError(err)
	s.True(ok)
}

func (s *WithdrawServiceTestSuite) Test5_WithdrawIdempotentReplay() {
	ctx := context.Background()
	key := fmt.Sprintf("withdraw:create:user:%d:op:%d", s.userID, time.Now().UnixNano())
	amount := 50000

	var before db.Saldo
	err := sqlx.GetContext(ctx, s.dbPool, &before, getSaldoByCardNumberSQL, s.cardNumber)
	s.Require().NoError(err)

	req := &requests.CreateWithdrawRequest{
		CardNumber:     s.cardNumber,
		WithdrawAmount: amount,
		WithdrawTime:   time.Now(),
		IdempotencyKey: key,
	}

	// A replay with the same key must return the original withdrawal and must
	// NOT debit the balance a second time.
	first, err := s.withdrawService.Create(ctx, req)
	s.Require().NoError(err)
	s.Require().NotNil(first)

	second, err := s.withdrawService.Create(ctx, req)
	s.Require().NoError(err)
	s.Require().NotNil(second)
	s.Equal(first.WithdrawID, second.WithdrawID, "replay must return the original withdrawal")

	var after db.Saldo
	err = sqlx.GetContext(ctx, s.dbPool, &after, getSaldoByCardNumberSQL, s.cardNumber)
	s.Require().NoError(err)
	s.Equal(before.TotalBalance-int32(amount), after.TotalBalance, "balance must be debited exactly once")
}

func (s *WithdrawServiceTestSuite) TestZZ_NotFound404() {
	ctx := context.Background()
	_, err := s.withdrawService.FindById(ctx, 999999999)
	s.Require().Error(err)
	s.Contains(err.Error(), "not found")
}

func (s *WithdrawServiceTestSuite) Test6_FindByIdNotFound() {
	ctx := context.Background()
	_, err := s.withdrawService.FindById(ctx, 999999999)
	s.Require().Error(err, "a non-existent withdraw must not resolve")
	s.Contains(err.Error(), "not found", "missing withdraw must surface as a 404-style not-found error")
}

func TestWithdrawServiceSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	suite.Run(t, new(WithdrawServiceTestSuite))
}
