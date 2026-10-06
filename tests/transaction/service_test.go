package transaction_test

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	card_repo_impl "github.com/MamangRust/monolith-payment-gateway-card/repository"
	merchant_repo_impl "github.com/MamangRust/monolith-payment-gateway-merchant/repository"
	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	saldo_repo_impl "github.com/MamangRust/monolith-payment-gateway-saldo/repository"
	"github.com/MamangRust/monolith-payment-gateway-shared/cache"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/observability"
	"github.com/MamangRust/monolith-payment-gateway-transaction/repository"
	"github.com/MamangRust/monolith-payment-gateway-transaction/service"
	user_repo_impl "github.com/MamangRust/monolith-payment-gateway-user/repository"
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

	getCardByUserIDSQL = `
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
    user_id = $1
    AND deleted_at IS NULL
LIMIT 1
`

	getMerchantByApiKeySQL = `
SELECT
    merchant_id,
    merchant_no,
    name,
    api_key,
    user_id,
    status,
    created_at,
    updated_at
FROM merchants
WHERE
    api_key = $1
    AND deleted_at IS NULL
`

	updateCardSQL = `
UPDATE cards
SET
    card_type = $2,
    expire_date = $3,
    cvv = $4,
    card_provider = $5,
    updated_at = current_timestamp
WHERE
    card_id = $1
    AND deleted_at IS NULL
RETURNING
    card_id,
    user_id,
    card_number,
    card_type,
    expire_date,
    cvv,
    card_provider,
    created_at,
    updated_at
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

	addRewardPointsSQL = `
UPDATE cards
SET
    reward_points = reward_points + $2,
    updated_at = CURRENT_TIMESTAMP
WHERE
    card_id = $1
    AND deleted_at IS NULL
RETURNING
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
`

	updateOutstandingBalanceSQL = `
UPDATE cards
SET
    outstanding_balance = $2,
    updated_at = CURRENT_TIMESTAMP
WHERE
    card_id = $1
    AND deleted_at IS NULL
RETURNING
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
`

	createTransactionSQL = `
INSERT INTO
    transactions (
        card_number,
        amount,
        payment_method,
        merchant_id,
        transaction_time,
        created_at,
        updated_at
    )
VALUES (
        $1,
        $2,
        $3,
        $4,
        $5,
        current_timestamp,
        current_timestamp
    )
RETURNING
    transaction_id,
    transaction_no,
    card_number,
    amount,
    payment_method,
    merchant_id,
    transaction_time,
    status,
    created_at,
    updated_at
`

	updateTransactionStatusDirectSQL = `
UPDATE transactions
SET
    status = $2,
    updated_at = CURRENT_TIMESTAMP
WHERE
    transaction_id = $1
    AND deleted_at IS NULL
RETURNING
    transaction_id,
    transaction_no,
    card_number,
    amount,
    payment_method,
    merchant_id,
    transaction_time,
    status,
    fraud_score,
    created_at,
    updated_at
`

	getTransactionsByMerchantIDSQL = `
SELECT
    transaction_id,
    transaction_no,
    card_number,
    amount,
    payment_method,
    merchant_id,
    transaction_time,
    status,
    created_at,
    updated_at
FROM transactions
WHERE
    merchant_id = $1
    AND deleted_at IS NULL
ORDER BY transaction_time DESC
`
)

type testMerchantRepo struct {
	db *sqlx.DB
}

func (r *testMerchantRepo) FindByApiKey(ctx context.Context, api_key string) (*db.GetMerchantByApiKeyRow, error) {
	var row db.GetMerchantByApiKeyRow
	if err := sqlx.GetContext(ctx, r.db, &row, getMerchantByApiKeySQL, api_key); err != nil {
		return nil, err
	}
	return &row, nil
}

type testCardRepo struct {
	db *sqlx.DB
}

func (r *testCardRepo) FindCardByUserId(ctx context.Context, user_id int) (*db.GetCardByUserIDRow, error) {
	var row db.GetCardByUserIDRow
	if err := sqlx.GetContext(ctx, r.db, &row, getCardByUserIDSQL, int32(user_id)); err != nil {
		return nil, err
	}
	return &row, nil
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

func (r *testCardRepo) UpdateCard(ctx context.Context, request *requests.UpdateCardRequest) (*db.UpdateCardRow, error) {
	expireDate := pgtype.Date{
		Time:  request.ExpireDate,
		Valid: true,
	}
	var row db.UpdateCardRow
	if err := sqlx.GetContext(ctx, r.db, &row, updateCardSQL,
		int32(request.CardID), request.CardType, expireDate, request.CVV, request.CardProvider); err != nil {
		return nil, err
	}
	return &row, nil
}

func (r *testCardRepo) UpdateCardOutstandingBalance(ctx context.Context, cardID int, outstandingBalance int) (*db.UpdateOutstandingBalanceRow, error) {
	var row db.UpdateOutstandingBalanceRow
	if err := sqlx.GetContext(ctx, r.db, &row, updateOutstandingBalanceSQL,
		int32(cardID), int32(outstandingBalance)); err != nil {
		return nil, err
	}
	return &row, nil
}

func (r *testCardRepo) AddRewardPoints(ctx context.Context, cardID int, points int) (*db.AddRewardPointsRow, error) {
	var row db.AddRewardPointsRow
	if err := sqlx.GetContext(ctx, r.db, &row, addRewardPointsSQL,
		int32(cardID), int32(points)); err != nil {
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

type TransactionServiceTestSuite struct {
	suite.Suite
	ts                  *tests.TestSuite
	transactionService  service.Service
	dbPool              *sqlx.DB
	deps                *grpcDeps
	transactionID       int
	cardNumber          string
	merchantID          int
	merchantApiKey      string
	otherMerchantApiKey string
	userID              int
	merchantUserID      int
}

func (s *TransactionServiceTestSuite) SetupSuite() {
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
		deps.merchantQueryClient,
	)

	s.transactionService = service.NewService(&service.Deps{
		Kafka:        nil,
		Repositories: repos,
		Logger:       log,
		Cache:        cacheStore,
	})

	// Seed User
	userRepo := user_repo_impl.NewUserCommandRepository(queries)
	user, err := userRepo.CreateUser(context.Background(), &requests.CreateUserRequest{
		FirstName: "Transaction",
		LastName:  "Tester",
		Email:     fmt.Sprintf("transaction.tester.%d@example.com", time.Now().UnixNano()),
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

	// Seed a separate merchant owner. The customer and merchant must not share
	// the same saldo account because settlement debits one account and credits
	// the other atomically.
	merchantOwner, err := userRepo.CreateUser(context.Background(), &requests.CreateUserRequest{
		FirstName: "Merchant",
		LastName:  "Owner",
		Email:     fmt.Sprintf("merchant.owner.%d@example.com", time.Now().UnixNano()),
		Password:  "password123",
	})
	s.Require().NoError(err)
	s.merchantUserID = int(merchantOwner.UserID)

	// Seed Merchant
	merchantRepoImpl := merchant_repo_impl.NewMerchantCommandRepository(queries)
	merchant, err := merchantRepoImpl.CreateMerchant(context.Background(), &requests.CreateMerchantRequest{
		UserID: s.merchantUserID,
		Name:   "Test Merchant",
	})
	s.Require().NoError(err)
	s.merchantID = int(merchant.MerchantID)
	s.merchantApiKey = merchant.ApiKey

	otherMerchantOwner, err := userRepo.CreateUser(context.Background(), &requests.CreateUserRequest{
		FirstName: "Other Merchant",
		LastName:  "Owner",
		Email:     fmt.Sprintf("other.merchant.owner.%d@example.com", time.Now().UnixNano()),
		Password:  "password123",
	})
	s.Require().NoError(err)
	otherMerchant, err := merchantRepoImpl.CreateMerchant(context.Background(), &requests.CreateMerchantRequest{
		UserID: int(otherMerchantOwner.UserID),
		Name:   "Other Test Merchant",
	})
	s.Require().NoError(err)
	s.otherMerchantApiKey = otherMerchant.ApiKey

	// Seed Merchant Card & Saldo (needed for transaction credit)
	merchantCard, err := cardCmdRepo.CreateCard(context.Background(), &requests.CreateCardRequest{
		UserID:       s.merchantUserID,
		CardType:     "debit",
		ExpireDate:   time.Now().AddDate(5, 0, 0),
		CVV:          "456",
		CardProvider: "mastercard",
	})
	s.Require().NoError(err)
	_, err = saldoCmdRepo.CreateSaldo(context.Background(), &requests.CreateSaldoRequest{
		CardNumber:   merchantCard.CardNumber,
		TotalBalance: 1000000,
	})
	s.Require().NoError(err)
}

func (s *TransactionServiceTestSuite) TearDownSuite() {
	if s.deps != nil {
		s.deps.Close()
	}
	s.dbPool.Close()
	s.ts.Teardown()
}

func (s *TransactionServiceTestSuite) Test1_TransactionLifecycle() {
	ctx := context.Background()

	// Create Transaction
	merchantID := int(s.merchantID)
	createReq := &requests.CreateTransactionRequest{
		CardNumber:      s.cardNumber,
		Amount:          100000,
		PaymentMethod:   "gopay",
		MerchantID:      &merchantID,
		TransactionTime: time.Now(),
	}
	res, err := s.transactionService.Create(ctx, s.merchantApiKey, createReq)
	s.NoError(err)
	s.NotNil(res)
	s.transactionID = int(res.TransactionID)
	s.Equal("success", res.Status)

	userSaldo, err := (&testSaldoRepo{db: s.dbPool}).FindByCardNumber(ctx, s.cardNumber)
	s.Require().NoError(err)
	s.Equal(int32(900000), userSaldo.TotalBalance)

	// FindById
	found, err := s.transactionService.FindById(ctx, s.transactionID)
	s.NoError(err)
	s.NotNil(found)
	s.Equal(int32(s.transactionID), found.TransactionID)

	// Update Transaction
	updateReq := &requests.UpdateTransactionRequest{
		TransactionID:   &s.transactionID,
		CardNumber:      s.cardNumber,
		Amount:          200000,
		PaymentMethod:   "dana",
		MerchantID:      &merchantID,
		TransactionTime: time.Now(),
	}
	updated, err := s.transactionService.Update(ctx, s.merchantApiKey, updateReq)
	s.NoError(err)
	s.NotNil(updated)
	s.Equal(int32(200000), updated.Amount)
}

func (s *TransactionServiceTestSuite) TestOwnershipChecksRejectVoidAndRefund() {
	ctx := context.Background()
	merchantID := int32(s.merchantID)

	var beforeUser db.Saldo
	err := sqlx.GetContext(ctx, s.dbPool, &beforeUser, getSaldoByCardNumberSQL, s.cardNumber)
	s.Require().NoError(err)
	var merchantCard db.GetCardByUserIDRow
	err = sqlx.GetContext(ctx, s.dbPool, &merchantCard, getCardByUserIDSQL, int32(s.merchantUserID))
	s.Require().NoError(err)
	var beforeMerchant db.Saldo
	err = sqlx.GetContext(ctx, s.dbPool, &beforeMerchant, getSaldoByCardNumberSQL, merchantCard.CardNumber)
	s.Require().NoError(err)

	var voidTransaction db.CreateTransactionRow
	err = sqlx.GetContext(ctx, s.dbPool, &voidTransaction, createTransactionSQL,
		s.cardNumber, int32(100), "gopay", merchantID, time.Now())
	s.Require().NoError(err)
	_, err = s.dbPool.ExecContext(ctx, updateTransactionStatusDirectSQL,
		voidTransaction.TransactionID, "authorized")
	s.Require().NoError(err)

	_, err = s.transactionService.VoidTransaction(ctx, s.otherMerchantApiKey, int(voidTransaction.TransactionID))
	s.Require().Error(err)
	s.Contains(err.Error(), "merchant does not own transaction")

	var refundTransaction db.CreateTransactionRow
	err = sqlx.GetContext(ctx, s.dbPool, &refundTransaction, createTransactionSQL,
		s.cardNumber, int32(100), "gopay", merchantID, time.Now())
	s.Require().NoError(err)
	_, err = s.dbPool.ExecContext(ctx, updateTransactionStatusDirectSQL,
		refundTransaction.TransactionID, "captured")
	s.Require().NoError(err)

	_, err = s.transactionService.RefundTransaction(ctx, s.otherMerchantApiKey, int(refundTransaction.TransactionID))
	s.Require().Error(err)
	s.Contains(err.Error(), "merchant does not own transaction")

	var afterUser db.Saldo
	err = sqlx.GetContext(ctx, s.dbPool, &afterUser, getSaldoByCardNumberSQL, s.cardNumber)
	s.Require().NoError(err)
	s.Equal(beforeUser.TotalBalance, afterUser.TotalBalance)
	var afterMerchant db.Saldo
	err = sqlx.GetContext(ctx, s.dbPool, &afterMerchant, getSaldoByCardNumberSQL, merchantCard.CardNumber)
	s.Require().NoError(err)
	s.Equal(beforeMerchant.TotalBalance, afterMerchant.TotalBalance)
}

func (s *TransactionServiceTestSuite) Test2_AtomicCreateRejectsInsufficientBalance() {
	ctx := context.Background()
	queries := s.dbPool
	merchantID := s.merchantID
	beforeUser, err := (&testSaldoRepo{db: queries}).FindByCardNumber(ctx, s.cardNumber)
	s.Require().NoError(err)

	merchantCard, err := (&testCardRepo{db: queries}).FindCardByUserId(ctx, s.merchantUserID)
	s.Require().NoError(err)
	beforeMerchant, err := (&testSaldoRepo{db: queries}).FindByCardNumber(ctx, merchantCard.CardNumber)
	s.Require().NoError(err)
	var beforeTransactions []*db.GetTransactionsByMerchantIDRow
	err = sqlx.SelectContext(ctx, s.dbPool, &beforeTransactions, getTransactionsByMerchantIDSQL, int32(s.merchantID))
	s.Require().NoError(err)

	res, err := s.transactionService.Create(ctx, s.merchantApiKey, &requests.CreateTransactionRequest{
		CardNumber:      s.cardNumber,
		Amount:          int(beforeUser.TotalBalance) + 1,
		PaymentMethod:   "gopay",
		MerchantID:      &merchantID,
		TransactionTime: time.Now(),
	})
	s.Error(err)
	s.Nil(res)

	afterUser, err := (&testSaldoRepo{db: queries}).FindByCardNumber(ctx, s.cardNumber)
	s.Require().NoError(err)
	s.Equal(beforeUser.TotalBalance, afterUser.TotalBalance)
	afterMerchant, err := (&testSaldoRepo{db: queries}).FindByCardNumber(ctx, merchantCard.CardNumber)
	s.Require().NoError(err)
	s.Equal(beforeMerchant.TotalBalance, afterMerchant.TotalBalance)
	var afterTransactions []*db.GetTransactionsByMerchantIDRow
	err = sqlx.SelectContext(ctx, s.dbPool, &afterTransactions, getTransactionsByMerchantIDSQL, int32(s.merchantID))
	s.Require().NoError(err)
	s.Len(afterTransactions, len(beforeTransactions))
}

func (s *TransactionServiceTestSuite) Test3_ConcurrentAtomicCreatesSettleOnce() {
	ctx := context.Background()
	queries := s.dbPool
	merchantID := s.merchantID
	beforeUser, err := (&testSaldoRepo{db: queries}).FindByCardNumber(ctx, s.cardNumber)
	s.Require().NoError(err)
	merchantCard, err := (&testCardRepo{db: queries}).FindCardByUserId(ctx, s.merchantUserID)
	s.Require().NoError(err)
	beforeMerchant, err := (&testSaldoRepo{db: queries}).FindByCardNumber(ctx, merchantCard.CardNumber)
	s.Require().NoError(err)
	var beforeTransactions []*db.GetTransactionsByMerchantIDRow
	err = sqlx.SelectContext(ctx, s.dbPool, &beforeTransactions, getTransactionsByMerchantIDSQL, int32(s.merchantID))
	s.Require().NoError(err)

	// Each request must be individually affordable, while two concurrent
	// requests together must exceed the available balance.
	amount := int(beforeUser.TotalBalance)/2 + 1
	s.Require().GreaterOrEqual(int(beforeUser.TotalBalance), amount)
	s.Require().Less(int(beforeUser.TotalBalance), amount*2)

	requestsReady := make(chan struct{})
	results := make(chan error, 2)
	var wg sync.WaitGroup
	wg.Add(2)
	for i := 0; i < 2; i++ {
		go func() {
			defer wg.Done()
			<-requestsReady
			_, err := s.transactionService.Create(ctx, s.merchantApiKey, &requests.CreateTransactionRequest{
				CardNumber:      s.cardNumber,
				Amount:          amount,
				PaymentMethod:   "gopay",
				MerchantID:      &merchantID,
				TransactionTime: time.Now(),
			})
			results <- err
		}()
	}
	close(requestsReady)
	wg.Wait()
	close(results)

	successes := 0
	for err := range results {
		if err == nil {
			successes++
		}
	}
	s.Equal(1, successes)

	afterUser, err := (&testSaldoRepo{db: queries}).FindByCardNumber(ctx, s.cardNumber)
	s.Require().NoError(err)
	s.Equal(beforeUser.TotalBalance-int32(amount), afterUser.TotalBalance)
	afterMerchant, err := (&testSaldoRepo{db: queries}).FindByCardNumber(ctx, merchantCard.CardNumber)
	s.Require().NoError(err)
	s.Equal(beforeMerchant.TotalBalance+int32(amount), afterMerchant.TotalBalance)
	var afterTransactions []*db.GetTransactionsByMerchantIDRow
	err = sqlx.SelectContext(ctx, s.dbPool, &afterTransactions, getTransactionsByMerchantIDSQL, int32(s.merchantID))
	s.Require().NoError(err)
	s.Len(afterTransactions, len(beforeTransactions)+1)
}

func (s *TransactionServiceTestSuite) Test4_QueryOperations() {
	ctx := context.Background()

	// FindAll
	all, total, err := s.transactionService.FindAll(ctx, &requests.FindAllTransactions{
		Page:     1,
		PageSize: 10,
	})
	s.NoError(err)
	s.NotNil(all)
	s.GreaterOrEqual(*total, 1)

	// FindByActive
	active, totalActive, err := s.transactionService.FindByActive(ctx, &requests.FindAllTransactions{
		Page:     1,
		PageSize: 10,
	})
	s.NoError(err)
	s.NotNil(active)
	s.GreaterOrEqual(*totalActive, 1)

	// FindByCardNumber
	byCard, totalCard, err := s.transactionService.FindAllByCardNumber(ctx, &requests.FindAllTransactionCardNumber{
		CardNumber: s.cardNumber,
		Page:       1,
		PageSize:   10,
	})
	s.NoError(err)
	s.NotNil(byCard)
	s.GreaterOrEqual(*totalCard, 1)

	// FindByMerchantId
	byMerchant, err := s.transactionService.FindTransactionByMerchantId(ctx, s.merchantID)
	s.NoError(err)
	s.NotNil(byMerchant)
	s.GreaterOrEqual(len(byMerchant), 1)
}

func (s *TransactionServiceTestSuite) Test5_TrashAndRestore() {
	ctx := context.Background()
	s.Require().NotZero(s.transactionID)

	// Trash
	trashed, err := s.transactionService.TrashedTransaction(ctx, s.transactionID)
	s.NoError(err)
	s.NotNil(trashed)

	// FindByTrashed
	trashedList, totalTrashed, err := s.transactionService.FindByTrashed(ctx, &requests.FindAllTransactions{
		Page:     1,
		PageSize: 10,
	})
	s.NoError(err)
	s.NotNil(trashedList)
	s.GreaterOrEqual(*totalTrashed, 1)

	// Restore
	restored, err := s.transactionService.RestoreTransaction(ctx, s.transactionID)
	s.NoError(err)
	s.NotNil(restored)
}

func (s *TransactionServiceTestSuite) Test6_BulkOperations() {
	ctx := context.Background()

	// Restore All
	ok, err := s.transactionService.RestoreAllTransaction(ctx)
	s.NoError(err)
	s.True(ok)

	// Delete All Permanent
	ok, err = s.transactionService.DeleteAllTransactionPermanent(ctx)
	s.NoError(err)
	s.True(ok)
}

func (s *TransactionServiceTestSuite) Test7_TransactionIdempotentReplay() {
	ctx := context.Background()
	key := fmt.Sprintf("transaction:create:merchant:%d:op:%d", s.merchantID, time.Now().UnixNano())
	amount := 50000
	merchantID := s.merchantID

	var userBefore db.Saldo
	err := sqlx.GetContext(ctx, s.dbPool, &userBefore, getSaldoByCardNumberSQL, s.cardNumber)
	s.Require().NoError(err)

	req := &requests.CreateTransactionRequest{
		CardNumber:      s.cardNumber,
		Amount:          amount,
		PaymentMethod:   "gopay",
		MerchantID:      &merchantID,
		TransactionTime: time.Now(),
		IdempotencyKey:  key,
	}

	// A replay with the same key must return the original transaction and must
	// NOT settle a second time.
	first, err := s.transactionService.Create(ctx, s.merchantApiKey, req)
	s.Require().NoError(err)
	s.Require().NotNil(first)

	second, err := s.transactionService.Create(ctx, s.merchantApiKey, req)
	s.Require().NoError(err)
	s.Require().NotNil(second)
	s.Equal(first.TransactionID, second.TransactionID, "replay must return the original transaction")

	var userAfter db.Saldo
	err = sqlx.GetContext(ctx, s.dbPool, &userAfter, getSaldoByCardNumberSQL, s.cardNumber)
	s.Require().NoError(err)
	s.Equal(userBefore.TotalBalance-int32(amount), userAfter.TotalBalance, "user must be debited exactly once")
}

func (s *TransactionServiceTestSuite) Test8_InvalidStateTransitionRejected400() {
	ctx := context.Background()
	merchantID := s.merchantID

	var merchantCard db.GetCardByUserIDRow
	err := sqlx.GetContext(ctx, s.dbPool, &merchantCard, getCardByUserIDSQL, int32(s.merchantUserID))
	s.Require().NoError(err)
	var userBefore db.Saldo
	err = sqlx.GetContext(ctx, s.dbPool, &userBefore, getSaldoByCardNumberSQL, s.cardNumber)
	s.Require().NoError(err)
	var merchantBefore db.Saldo
	err = sqlx.GetContext(ctx, s.dbPool, &merchantBefore, getSaldoByCardNumberSQL, merchantCard.CardNumber)
	s.Require().NoError(err)

	// The atomic create settles immediately into "success"; lifecycle
	// operations whose source state precondition is not met must be rejected
	// 400 via the state machine WITHOUT moving any balance.
	key := fmt.Sprintf("transaction:create:merchant:%d:transition:%d", s.merchantID, time.Now().UnixNano())
	res, err := s.transactionService.Create(ctx, s.merchantApiKey, &requests.CreateTransactionRequest{
		CardNumber:      s.cardNumber,
		Amount:          50000,
		PaymentMethod:   "gopay",
		MerchantID:      &merchantID,
		TransactionTime: time.Now(),
		IdempotencyKey:  key,
	})
	s.Require().NoError(err)
	txID := int(res.TransactionID)

	_, err = s.transactionService.CaptureTransaction(ctx, s.merchantApiKey, txID)
	s.Require().Error(err)
	s.Contains(err.Error(), "invalid status transition", "capture from success must be 400")

	_, err = s.transactionService.VoidTransaction(ctx, s.merchantApiKey, txID)
	s.Require().Error(err)
	s.Contains(err.Error(), "invalid status transition", "void from success must be 400")

	_, err = s.transactionService.RefundTransaction(ctx, s.merchantApiKey, txID)
	s.Require().Error(err)
	s.Contains(err.Error(), "invalid status transition", "refund from success must be 400")

	var userAfter db.Saldo
	err = sqlx.GetContext(ctx, s.dbPool, &userAfter, getSaldoByCardNumberSQL, s.cardNumber)
	s.Require().NoError(err)
	var merchantAfter db.Saldo
	err = sqlx.GetContext(ctx, s.dbPool, &merchantAfter, getSaldoByCardNumberSQL, merchantCard.CardNumber)
	s.Require().NoError(err)
	s.Equal(userBefore.TotalBalance-int32(50000), userAfter.TotalBalance, "only the original settlement may move money")
	s.Equal(merchantBefore.TotalBalance+int32(50000), merchantAfter.TotalBalance)
}

func (s *TransactionServiceTestSuite) TestZZ_NotFound404() {
	ctx := context.Background()
	_, err := s.transactionService.FindById(ctx, 999999999)
	s.Require().Error(err)
	s.Contains(err.Error(), "not found")
}

func (s *TransactionServiceTestSuite) Test9_FindByIdNotFound() {
	ctx := context.Background()
	_, err := s.transactionService.FindById(ctx, 999999999)
	s.Require().Error(err, "a non-existent transaction must not resolve")
	s.Contains(err.Error(), "not found", "missing transaction must surface as a 404-style not-found error")
}

func TestTransactionServiceSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	suite.Run(t, new(TransactionServiceTestSuite))
}
