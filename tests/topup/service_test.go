package topup_test

import (
	"context"
	"fmt"
	"net"
	"sync"
	"testing"
	"time"

	card_handler "github.com/MamangRust/monolith-payment-gateway-card/handler"
	card_repo_impl "github.com/MamangRust/monolith-payment-gateway-card/repository"
	card_service "github.com/MamangRust/monolith-payment-gateway-card/service"
	pbcard "github.com/MamangRust/monolith-payment-gateway-pb/card"
	pbuser "github.com/MamangRust/monolith-payment-gateway-pb/user"
	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	saldo_repo_impl "github.com/MamangRust/monolith-payment-gateway-saldo/repository"
	"github.com/MamangRust/monolith-payment-gateway-shared/cache"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/observability"
	"github.com/MamangRust/monolith-payment-gateway-topup/repository"
	"github.com/MamangRust/monolith-payment-gateway-topup/service"
	user_repo "github.com/MamangRust/monolith-payment-gateway-user/repository"
	tests "github.com/MamangRust/monolith-payment-gateway-test"
	"github.com/jmoiron/sqlx"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/suite"
	sdklog "go.opentelemetry.io/otel/sdk/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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

	getTopupByIDSQL = `
SELECT
    topup_id,
    topup_no,
    card_number,
    topup_amount,
    topup_method,
    topup_time,
    status,
    created_at,
    updated_at
FROM topups
WHERE
    topup_id = $1
    AND deleted_at IS NULL
`
)

type TopupServiceTestSuite struct {
	suite.Suite
	ts           *tests.TestSuite
	topupService service.Service
	dbPool       *sqlx.DB
	cardServer   *grpc.Server
	identity     *tests.IdentityStack
	topupID      int
	cardNumber   string
	userID       int
}

func (s *TopupServiceTestSuite) SetupSuite() {
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

	// The topup service reaches the card service through the card adapter
	// (FindUserCardByCardNumber/UpdateCard), so the card service is exposed over
	// an in-process gRPC server. The card command service resolves the card
	// owner through the user service, hence the identity stack.
	s.identity, err = tests.StartIdentityStack(queries, cacheStore, log)
	s.Require().NoError(err)

	lis, err := net.Listen("tcp", "localhost:0")
	s.Require().NoError(err)

	cardConn, err := grpc.NewClient(lis.Addr().String(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	s.Require().NoError(err)

	cardSvc := card_service.NewService(&card_service.Deps{
		Cache:        cacheStore,
		Repositories: card_repo_impl.NewRepositories(queries, pbuser.NewUserQueryServiceClient(s.identity.User.Conn)),
		Logger:       log,
		Kafka:        nil,
	})
	cardH := card_handler.NewHandler(cardSvc)

	s.cardServer = grpc.NewServer()
	pbcard.RegisterCardQueryServiceServer(s.cardServer, cardH)
	pbcard.RegisterCardCommandServiceServer(s.cardServer, cardH)

	go func() {
		_ = s.cardServer.Serve(lis)
	}()

	// The topup service never calls the saldo repository, so its clients are nil.
	repos := repository.NewRepositories(
		queries,
		pbcard.NewCardQueryServiceClient(cardConn),
		pbcard.NewCardCommandServiceClient(cardConn),
		nil,
		nil,
	)

	s.topupService = service.NewService(&service.Deps{
		Kafka:        nil,
		Repositories: repos,
		Logger:       log,
		Cache:        cacheStore,
	})

	// Seed User
	userRepo := user_repo.NewUserCommandRepository(queries)
	user, err := userRepo.CreateUser(context.Background(), &requests.CreateUserRequest{
		FirstName: "Topup",
		LastName:  "Tester",
		Email:     fmt.Sprintf("topup.tester.%d@example.com", time.Now().UnixNano()),
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

func (s *TopupServiceTestSuite) TearDownSuite() {
	if s.cardServer != nil {
		s.cardServer.Stop()
	}
	if s.identity != nil {
		s.identity.Close()
	}
	s.dbPool.Close()
	s.ts.Teardown()
}

func (s *TopupServiceTestSuite) Test1_TopupLifecycle() {
	ctx := context.Background()

	// Create Topup
	createReq := &requests.CreateTopupRequest{
		CardNumber:  s.cardNumber,
		TopupAmount: 100000,
		TopupMethod: "gopay",
	}
	res, err := s.topupService.CreateTopup(ctx, createReq)
	s.NoError(err)
	s.NotNil(res)
	s.topupID = int(res.TopupID)
	s.Equal("success", res.Status)

	// FindById
	found, err := s.topupService.FindById(ctx, s.topupID)
	s.NoError(err)
	s.NotNil(found)
	s.Equal(int32(s.topupID), found.TopupID)

	// Update Topup
	updateReq := &requests.UpdateTopupRequest{
		TopupID:     &s.topupID,
		CardNumber:  s.cardNumber,
		TopupAmount: 200000,
		TopupMethod: "dana",
	}
	updated, err := s.topupService.UpdateTopup(ctx, updateReq)
	s.NoError(err)
	s.NotNil(updated)
	s.Equal("success", updated.Status)
}

func (s *TopupServiceTestSuite) Test2_QueryOperations() {
	ctx := context.Background()

	// FindAll
	all, total, err := s.topupService.FindAll(ctx, &requests.FindAllTopups{
		Page:     1,
		PageSize: 10,
	})
	s.NoError(err)
	s.NotNil(all)
	s.GreaterOrEqual(*total, 1)

	// FindByActive
	active, totalActive, err := s.topupService.FindByActive(ctx, &requests.FindAllTopups{
		Page:     1,
		PageSize: 10,
	})
	s.NoError(err)
	s.NotNil(active)
	s.GreaterOrEqual(*totalActive, 1)

	// FindByCardNumber
	byCard, _, err := s.topupService.FindAllByCardNumber(ctx, &requests.FindAllTopupsByCardNumber{
		CardNumber: s.cardNumber,
		Page:       1,
		PageSize:   10,
	})
	s.NoError(err)
	s.NotNil(byCard)
	s.GreaterOrEqual(len(byCard), 1)
}

func (s *TopupServiceTestSuite) Test3_TrashAndRestore() {
	ctx := context.Background()
	s.Require().NotZero(s.topupID)

	// Trash
	trashed, err := s.topupService.TrashedTopup(ctx, s.topupID)
	s.NoError(err)
	s.NotNil(trashed)

	// FindByTrashed
	trashedList, totalTrashed, err := s.topupService.FindByTrashed(ctx, &requests.FindAllTopups{
		Page:     1,
		PageSize: 10,
	})
	s.NoError(err)
	s.NotNil(trashedList)
	s.GreaterOrEqual(*totalTrashed, 1)

	// Restore
	restored, err := s.topupService.RestoreTopup(ctx, s.topupID)
	s.NoError(err)
	s.NotNil(restored)
}

func (s *TopupServiceTestSuite) Test4_BulkOperations() {
	ctx := context.Background()

	// Restore All
	ok, err := s.topupService.RestoreAllTopup(ctx)
	s.NoError(err)
	s.True(ok)

	// Delete All Permanent
	ok, err = s.topupService.DeleteAllTopupPermanent(ctx)
	s.NoError(err)
	s.True(ok)
}

func (s *TopupServiceTestSuite) Test5_TopupIdempotentReplay() {
	ctx := context.Background()
	key := fmt.Sprintf("topup:create:user:%d:op:%d", s.userID, time.Now().UnixNano())

	req := &requests.CreateTopupRequest{
		CardNumber:     s.cardNumber,
		TopupAmount:    150000,
		TopupMethod:    "gopay",
		IdempotencyKey: key,
	}

	first, err := s.topupService.CreateTopup(ctx, req)
	s.Require().NoError(err)
	s.Require().NotNil(first)
	s.Equal("success", first.Status)

	var balanceAfterCreate int64
	err = s.dbPool.QueryRowContext(ctx, "SELECT total_balance FROM saldos WHERE card_number = $1", s.cardNumber).Scan(&balanceAfterCreate)
	s.Require().NoError(err)

	// Replay with the same key: the original record is returned and the
	// balance must not be credited a second time (fast-path replay over the
	// unique idempotency_key index).
	second, err := s.topupService.CreateTopup(ctx, req)
	s.Require().NoError(err)
	s.Require().NotNil(second)
	s.Equal(first.TopupID, second.TopupID, "replay must return the original record")

	var balanceAfterReplay int64
	err = s.dbPool.QueryRowContext(ctx, "SELECT total_balance FROM saldos WHERE card_number = $1", s.cardNumber).Scan(&balanceAfterReplay)
	s.Require().NoError(err)
	s.Equal(balanceAfterCreate, balanceAfterReplay, "replay must not credit the balance twice")
}

func (s *TopupServiceTestSuite) Test6_TopupConcurrentIdempotentReplay() {
	ctx := context.Background()
	key := fmt.Sprintf("topup:create:user:%d:concurrent:%d", s.userID, time.Now().UnixNano())
	amount := 120000

	var before db.Saldo
	err := sqlx.GetContext(ctx, s.dbPool, &before, getSaldoByCardNumberSQL, s.cardNumber)
	s.Require().NoError(err)

	// Two concurrent creates with the SAME idempotency key: both must succeed
	// (the loser hits the unique index and falls back to returning the winner's
	// record — replay, not 409) and only one record/credit may exist.
	requestsReady := make(chan struct{})
	results := make(chan error, 2)
	ids := make(chan int32, 2)
	var wg sync.WaitGroup
	wg.Add(2)
	for i := 0; i < 2; i++ {
		go func() {
			defer wg.Done()
			<-requestsReady
			res, err := s.topupService.CreateTopup(ctx, &requests.CreateTopupRequest{
				CardNumber:     s.cardNumber,
				TopupAmount:    amount,
				TopupMethod:    "gopay",
				IdempotencyKey: key,
			})
			results <- err
			if err == nil {
				ids <- res.TopupID
			}
		}()
	}
	close(requestsReady)
	wg.Wait()
	close(results)
	close(ids)

	successes := 0
	for err := range results {
		if err == nil {
			successes++
		}
	}
	s.Equal(2, successes, "concurrent duplicate with the same key must both succeed (replay, not 409)")

	var firstID int32
	first := true
	for id := range ids {
		if first {
			firstID = id
			first = false
		} else {
			s.Equal(firstID, id, "both concurrent requests must return the same record id")
		}
	}

	var rowCount int
	err = s.dbPool.QueryRowContext(ctx, "SELECT COUNT(*) FROM topups WHERE idempotency_key = $1", key).Scan(&rowCount)
	s.Require().NoError(err)
	s.Equal(1, rowCount, "exactly one record may exist for the key")

	var after db.Saldo
	err = sqlx.GetContext(ctx, s.dbPool, &after, getSaldoByCardNumberSQL, s.cardNumber)
	s.Require().NoError(err)
	s.Equal(before.TotalBalance+int32(amount), after.TotalBalance, "balance must be credited exactly once")
}

func (s *TopupServiceTestSuite) Test7_TopupConcurrentUpdatesApplyAllDeltas() {
	ctx := context.Background()

	// Seed a topup to update concurrently.
	key := fmt.Sprintf("topup:create:user:%d:update:%d", s.userID, time.Now().UnixNano())
	res, err := s.topupService.CreateTopup(ctx, &requests.CreateTopupRequest{
		CardNumber:     s.cardNumber,
		TopupAmount:    100000,
		TopupMethod:    "gopay",
		IdempotencyKey: key,
	})
	s.Require().NoError(err)
	topupID := int(res.TopupID)

	var before db.Saldo
	err = sqlx.GetContext(ctx, s.dbPool, &before, getSaldoByCardNumberSQL, s.cardNumber)
	s.Require().NoError(err)

	// Two concurrent updates with different target amounts. The atomic delta
	// statement serializes the updates on the topup row (FOR UPDATE), so the
	// final balance must ALWAYS equal the balance implied by the final topup
	// amount: before - original_amount + final_amount. Depending on which
	// goroutine wins the row lock the final amount is 150000 or 200000 (and
	// the implied delta 50000 or 100000) — both are consistent. A
	// read-modify-write implementation would lose a delta (balance below the
	// implied value) or double-apply it (balance above), and is rejected here.
	requestsReady := make(chan struct{})
	results := make(chan error, 2)
	var wg sync.WaitGroup
	wg.Add(2)
	for _, newAmount := range []int{150000, 200000} {
		newAmount := newAmount
		go func() {
			defer wg.Done()
			<-requestsReady
			_, err := s.topupService.UpdateTopup(ctx, &requests.UpdateTopupRequest{
				TopupID:     &topupID,
				CardNumber:  s.cardNumber,
				TopupAmount: newAmount,
				TopupMethod: "gopay",
			})
			results <- err
		}()
	}
	close(requestsReady)
	wg.Wait()
	close(results)

	for err := range results {
		s.NoError(err, "concurrent updates must not fail")
	}

	var final db.GetTopupByIDRow
	err = sqlx.GetContext(ctx, s.dbPool, &final, getTopupByIDSQL, int32(topupID))
	s.Require().NoError(err)
	var after db.Saldo
	err = sqlx.GetContext(ctx, s.dbPool, &after, getSaldoByCardNumberSQL, s.cardNumber)
	s.Require().NoError(err)
	expected := before.TotalBalance - 100000 + final.TopupAmount
	s.Equal(expected, after.TotalBalance, "balance must be consistent with the final topup amount (no lost update, no double apply)")
}

func (s *TopupServiceTestSuite) TestZZ_NotFound404() {
	ctx := context.Background()
	_, err := s.topupService.FindById(ctx, 999999999)
	s.Require().Error(err)
	s.Contains(err.Error(), "not found")
}

func (s *TopupServiceTestSuite) Test8_FindByIdNotFound() {
	ctx := context.Background()
	_, err := s.topupService.FindById(ctx, 999999999)
	s.Require().Error(err, "a non-existent topup must not resolve")
	s.Contains(err.Error(), "not found", "missing topup must surface as a 404-style not-found error")
}

func TestTopupServiceSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	suite.Run(t, new(TopupServiceTestSuite))
}
