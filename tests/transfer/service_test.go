package transfer_test

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
	"github.com/MamangRust/monolith-payment-gateway-transfer/repository"
	"github.com/MamangRust/monolith-payment-gateway-transfer/service"
	user_repo_impl "github.com/MamangRust/monolith-payment-gateway-user/repository"
	tests "github.com/MamangRust/monolith-payment-gateway-test"
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

	getCardByCardNumberSQL = `
SELECT
    card_id,
    user_id,
    card_number,
    card_type,
    expire_date,
    cvv,
    card_provider,
    status,
    credit_limit,
    outstanding_balance,
    reward_points,
    created_at,
    updated_at
FROM cards
WHERE
    card_number = $1
    AND deleted_at IS NULL
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

func (r *testCardRepo) FindCardByCardNumber(ctx context.Context, card_number string) (*db.GetCardByCardNumberRow, error) {
	var row db.GetCardByCardNumberRow
	if err := sqlx.GetContext(ctx, r.db, &row, getCardByCardNumberSQL, card_number); err != nil {
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

type TransferServiceTestSuite struct {
	suite.Suite
	ts              *tests.TestSuite
	transferService service.Service
	dbPool          *sqlx.DB
	deps            *grpcDeps
	transferID      int
	senderCard      string
	receiverCard    string
	userID          int
}

func (s *TransferServiceTestSuite) SetupSuite() {
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
	)

	s.transferService = service.NewService(&service.Deps{
		Kafka:        nil,
		Repositories: repos,
		Logger:       log,
		Cache:        cacheStore,
	})

	// Seed User
	userRepo := user_repo_impl.NewUserCommandRepository(queries)
	user, err := userRepo.CreateUser(context.Background(), &requests.CreateUserRequest{
		FirstName: "Transfer",
		LastName:  "Tester",
		Email:     fmt.Sprintf("transfer.tester.%d@example.com", time.Now().UnixNano()),
		Password:  "password123",
	})
	s.Require().NoError(err)
	s.userID = int(user.UserID)

	// Seed Sender Card & Saldo
	cardCmdRepo := card_repo_impl.NewCardCommandRepository(queries)
	sender, err := cardCmdRepo.CreateCard(context.Background(), &requests.CreateCardRequest{
		UserID:       s.userID,
		CardType:     "debit",
		ExpireDate:   time.Now().AddDate(5, 0, 0),
		CVV:          "123",
		CardProvider: "visa",
	})
	s.Require().NoError(err)
	s.senderCard = sender.CardNumber

	saldoCmdRepo := saldo_repo_impl.NewSaldoCommandRepository(queries)
	_, err = saldoCmdRepo.CreateSaldo(context.Background(), &requests.CreateSaldoRequest{
		CardNumber:   s.senderCard,
		TotalBalance: 1000000,
	})
	s.Require().NoError(err)

	// Seed Receiver Card & Saldo
	receiver, err := cardCmdRepo.CreateCard(context.Background(), &requests.CreateCardRequest{
		UserID:       s.userID,
		CardType:     "debit",
		ExpireDate:   time.Now().AddDate(5, 0, 0),
		CVV:          "456",
		CardProvider: "mastercard",
	})
	s.Require().NoError(err)
	s.receiverCard = receiver.CardNumber

	_, err = saldoCmdRepo.CreateSaldo(context.Background(), &requests.CreateSaldoRequest{
		CardNumber:   s.receiverCard,
		TotalBalance: 0,
	})
	s.Require().NoError(err)
}

func (s *TransferServiceTestSuite) TearDownSuite() {
	if s.deps != nil {
		s.deps.Close()
	}
	s.dbPool.Close()
	s.ts.Teardown()
}

func (s *TransferServiceTestSuite) Test1_TransferLifecycle() {
	ctx := context.Background()

	// Create Transfer
	createReq := &requests.CreateTransferRequest{
		TransferFrom:   s.senderCard,
		TransferTo:     s.receiverCard,
		TransferAmount: 100000,
	}
	res, err := s.transferService.CreateTransaction(ctx, createReq)
	s.NoError(err)
	s.NotNil(res)
	s.transferID = int(res.TransferID)
	s.Equal("success", res.Status)

	// FindById
	found, err := s.transferService.FindById(ctx, s.transferID)
	s.NoError(err)
	s.NotNil(found)
	s.Equal(int32(s.transferID), found.TransferID)

	// Update Transfer
	updateReq := &requests.UpdateTransferRequest{
		TransferID:     &s.transferID,
		TransferFrom:   s.senderCard,
		TransferTo:     s.receiverCard,
		TransferAmount: 200000,
	}
	updated, err := s.transferService.UpdateTransaction(ctx, updateReq)
	s.NoError(err)
	s.NotNil(updated)
	s.Equal(int32(200000), updated.TransferAmount)
}

func (s *TransferServiceTestSuite) Test2_QueryOperations() {
	ctx := context.Background()

	// FindAll
	all, total, err := s.transferService.FindAll(ctx, &requests.FindAllTransfers{
		Page:     1,
		PageSize: 10,
	})
	s.NoError(err)
	s.NotNil(all)
	s.GreaterOrEqual(*total, 1)

	// FindByActive
	active, totalActive, err := s.transferService.FindByActive(ctx, &requests.FindAllTransfers{
		Page:     1,
		PageSize: 10,
	})
	s.NoError(err)
	s.NotNil(active)
	s.GreaterOrEqual(*totalActive, 1)

	// FindByTransferFrom
	from, err := s.transferService.FindTransferByTransferFrom(ctx, s.senderCard)
	s.NoError(err)
	s.NotNil(from)
	s.GreaterOrEqual(len(from), 1)

	// FindByTransferTo
	to, err := s.transferService.FindTransferByTransferTo(ctx, s.receiverCard)
	s.NoError(err)
	s.NotNil(to)
	s.GreaterOrEqual(len(to), 1)
}

func (s *TransferServiceTestSuite) Test3_TrashAndRestore() {
	ctx := context.Background()
	s.Require().NotZero(s.transferID)

	// Trash
	trashed, err := s.transferService.TrashedTransfer(ctx, s.transferID)
	s.NoError(err)
	s.NotNil(trashed)

	// FindByTrashed
	trashedList, totalTrashed, err := s.transferService.FindByTrashed(ctx, &requests.FindAllTransfers{
		Page:     1,
		PageSize: 10,
	})
	s.NoError(err)
	s.NotNil(trashedList)
	s.GreaterOrEqual(*totalTrashed, 1)

	// Restore
	restored, err := s.transferService.RestoreTransfer(ctx, s.transferID)
	s.NoError(err)
	s.NotNil(restored)
}

func (s *TransferServiceTestSuite) Test4_BulkOperations() {
	ctx := context.Background()

	// Restore All
	ok, err := s.transferService.RestoreAllTransfer(ctx)
	s.NoError(err)
	s.True(ok)

	// Delete All Permanent
	ok, err = s.transferService.DeleteAllTransferPermanent(ctx)
	s.NoError(err)
	s.True(ok)
}

func (s *TransferServiceTestSuite) Test5_TransferIdempotentReplay() {
	ctx := context.Background()
	key := fmt.Sprintf("transfer:create:user:%d:op:%d", s.userID, time.Now().UnixNano())
	amount := 50000

	var senderBefore db.Saldo
	err := sqlx.GetContext(ctx, s.dbPool, &senderBefore, getSaldoByCardNumberSQL, s.senderCard)
	s.Require().NoError(err)
	var receiverBefore db.Saldo
	err = sqlx.GetContext(ctx, s.dbPool, &receiverBefore, getSaldoByCardNumberSQL, s.receiverCard)
	s.Require().NoError(err)

	req := &requests.CreateTransferRequest{
		TransferFrom:   s.senderCard,
		TransferTo:     s.receiverCard,
		TransferAmount: amount,
		IdempotencyKey: key,
	}

	// A replay with the same key must return the original settlement and must
	// NOT move money a second time.
	first, err := s.transferService.CreateTransaction(ctx, req)
	s.Require().NoError(err)
	s.Require().NotNil(first)

	second, err := s.transferService.CreateTransaction(ctx, req)
	s.Require().NoError(err)
	s.Require().NotNil(second)
	s.Equal(first.TransferID, second.TransferID, "replay must return the original transfer")

	var senderAfter db.Saldo
	err = sqlx.GetContext(ctx, s.dbPool, &senderAfter, getSaldoByCardNumberSQL, s.senderCard)
	s.Require().NoError(err)
	var receiverAfter db.Saldo
	err = sqlx.GetContext(ctx, s.dbPool, &receiverAfter, getSaldoByCardNumberSQL, s.receiverCard)
	s.Require().NoError(err)
	s.Equal(senderBefore.TotalBalance-int32(amount), senderAfter.TotalBalance, "sender must be debited exactly once")
	s.Equal(receiverBefore.TotalBalance+int32(amount), receiverAfter.TotalBalance, "receiver must be credited exactly once")
}

func (s *TransferServiceTestSuite) TestZZ_NotFound404() {
	ctx := context.Background()
	_, err := s.transferService.FindById(ctx, 999999999)
	s.Require().Error(err)
	s.Contains(err.Error(), "not found")
}

func (s *TransferServiceTestSuite) Test6_FindByIdNotFound() {
	ctx := context.Background()
	_, err := s.transferService.FindById(ctx, 999999999)
	s.Require().Error(err, "a non-existent transfer must not resolve")
	s.Contains(err.Error(), "not found", "missing transfer must surface as a 404-style not-found error")
}

func TestTransferServiceSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	suite.Run(t, new(TransferServiceTestSuite))
}
