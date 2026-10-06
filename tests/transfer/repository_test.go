package transfer_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	card_repo "github.com/MamangRust/monolith-payment-gateway-card/repository"
	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/cache"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/observability"
	"github.com/MamangRust/monolith-payment-gateway-transfer/repository"
	user_repo "github.com/MamangRust/monolith-payment-gateway-user/repository"
	tests "github.com/MamangRust/monolith-payment-gateway-test"

	"github.com/jmoiron/sqlx"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/suite"
	sdklog "go.opentelemetry.io/otel/sdk/log"
)

type TransferRepositoryTestSuite struct {
	suite.Suite
	ts          *tests.TestSuite
	repo        repository.Repositories
	cardRepo    *card_repo.Repositories
	userRepo    user_repo.Repositories
	deps        *grpcDeps
	redisClient *redis.Client
	userID      int
}

func (s *TransferRepositoryTestSuite) SetupSuite() {
	ts, err := tests.SetupTestSuite()
	s.Require().NoError(err)
	s.ts = ts

	pool, err := sqlx.ConnectContext(s.ts.Ctx, "pgx", s.ts.DBURL)
	s.Require().NoError(err)

	queries := pool

	opts, err := redis.ParseURL(s.ts.RedisURL)
	s.Require().NoError(err)
	s.redisClient = redis.NewClient(opts)

	logger.ResetInstance()
	lp := sdklog.NewLoggerProvider()
	log, _ := logger.NewLogger("test", lp)
	cacheMetrics, _ := observability.NewCacheMetrics("test")
	cacheStore := cache.NewCacheStore(s.redisClient, log, cacheMetrics)

	deps, err := startGRPCDeps(queries, cacheStore, log)
	s.Require().NoError(err)
	s.deps = deps

	s.userRepo = deps.userRepo
	s.cardRepo = deps.cardRepo
	s.repo = repository.NewRepositories(
		queries,
		deps.cardQueryClient,
		deps.cardCommandClient,
		deps.saldoQueryClient,
		deps.saldoCommandClient,
	)

	// Create user
	user, err := s.userRepo.UserCommand().CreateUser(context.Background(), &requests.CreateUserRequest{
		FirstName: "Transfer",
		LastName:  "Tester",
		Email:     fmt.Sprintf("transfer.tester-%d@example.com", time.Now().UnixNano()),
		Password:  "password123",
	})
	s.Require().NoError(err)
	s.userID = int(user.UserID)
}

func (s *TransferRepositoryTestSuite) TearDownSuite() {
	if s.deps != nil {
		s.deps.Close()
	}
	if s.redisClient != nil {
		s.redisClient.Close()
	}
	s.ts.Teardown()
}

func (s *TransferRepositoryTestSuite) createSeedTransfer() (*db.CreateTransferRow, error) {
	fromCard, err := s.cardRepo.CardCommand.CreateCard(context.Background(), &requests.CreateCardRequest{
		UserID:       s.userID,
		CardType:     "debit",
		ExpireDate:   time.Now().AddDate(5, 0, 0),
		CVV:          "111",
		CardProvider: "Visa",
	})
	if err != nil {
		return nil, err
	}

	toCard, err := s.cardRepo.CardCommand.CreateCard(context.Background(), &requests.CreateCardRequest{
		UserID:       s.userID,
		CardType:     "debit",
		ExpireDate:   time.Now().AddDate(5, 0, 0),
		CVV:          "222",
		CardProvider: "MasterCard",
	})
	if err != nil {
		return nil, err
	}

	return s.repo.CreateTransfer(context.Background(), &requests.CreateTransferRequest{
		TransferFrom:   fromCard.CardNumber,
		TransferTo:     toCard.CardNumber,
		TransferAmount: 100000,
	})
}

func (s *TransferRepositoryTestSuite) TestCreateTransfer() {
	ctx := context.Background()

	fromCard, _ := s.cardRepo.CardCommand.CreateCard(ctx, &requests.CreateCardRequest{
		UserID:       s.userID,
		CardType:     "debit",
		ExpireDate:   time.Now().AddDate(5, 0, 0),
		CVV:          "111",
		CardProvider: "Visa",
	})

	toCard, _ := s.cardRepo.CardCommand.CreateCard(ctx, &requests.CreateCardRequest{
		UserID:       s.userID,
		CardType:     "debit",
		ExpireDate:   time.Now().AddDate(5, 0, 0),
		CVV:          "222",
		CardProvider: "MasterCard",
	})

	req := &requests.CreateTransferRequest{
		TransferFrom:   fromCard.CardNumber,
		TransferTo:     toCard.CardNumber,
		TransferAmount: 100000,
	}

	res, err := s.repo.CreateTransfer(ctx, req)
	s.NoError(err)
	s.NotNil(res)
}

func (s *TransferRepositoryTestSuite) TestFindAllTransfers() {
	_, err := s.createSeedTransfer()
	s.Require().NoError(err)
	ctx := context.Background()

	res, err := s.repo.FindAll(ctx, &requests.FindAllTransfers{
		Page:     1,
		PageSize: 10,
		Search:   "",
	})
	s.NoError(err)
	s.GreaterOrEqual(len(res), 1)
}

func (s *TransferRepositoryTestSuite) TestFindById() {
	transfer, err := s.createSeedTransfer()
	s.Require().NoError(err)
	ctx := context.Background()

	found, err := s.repo.FindById(ctx, int(transfer.TransferID))
	s.NoError(err)
	s.NotNil(found)
	s.Equal(transfer.TransferID, found.TransferID)
}

func (s *TransferRepositoryTestSuite) TestFindByActive() {
	_, err := s.createSeedTransfer()
	s.Require().NoError(err)
	ctx := context.Background()

	res, err := s.repo.FindByActive(ctx, &requests.FindAllTransfers{
		Page:     1,
		PageSize: 10,
		Search:   "",
	})
	s.NoError(err)
	s.GreaterOrEqual(len(res), 1)
}

func (s *TransferRepositoryTestSuite) TestFindByTrashed() {
	transfer, err := s.createSeedTransfer()
	s.Require().NoError(err)
	ctx := context.Background()

	_, err = s.repo.TrashedTransfer(ctx, int(transfer.TransferID))
	s.Require().NoError(err)

	res, err := s.repo.FindByTrashed(ctx, &requests.FindAllTransfers{
		Page:     1,
		PageSize: 10,
		Search:   "",
	})
	s.NoError(err)
	s.GreaterOrEqual(len(res), 1)
}

func (s *TransferRepositoryTestSuite) TestUpdateTransfer() {
	transfer, err := s.createSeedTransfer()
	s.Require().NoError(err)
	ctx := context.Background()

	id := int(transfer.TransferID)
	req := &requests.UpdateTransferRequest{
		TransferID:     &id,
		TransferFrom:   transfer.TransferFrom,
		TransferTo:     transfer.TransferTo,
		TransferAmount: 200000,
	}

	res, err := s.repo.UpdateTransfer(ctx, req)
	s.NoError(err)
	s.NotNil(res)
}

func (s *TransferRepositoryTestSuite) TestTrashTransfer() {
	transfer, err := s.createSeedTransfer()
	s.Require().NoError(err)
	ctx := context.Background()

	trashed, err := s.repo.TrashedTransfer(ctx, int(transfer.TransferID))
	s.NoError(err)
	s.NotNil(trashed)
}

func (s *TransferRepositoryTestSuite) TestRestoreTransfer() {
	transfer, err := s.createSeedTransfer()
	s.Require().NoError(err)
	ctx := context.Background()

	_, err = s.repo.TrashedTransfer(ctx, int(transfer.TransferID))
	s.Require().NoError(err)

	restored, err := s.repo.RestoreTransfer(ctx, int(transfer.TransferID))
	s.NoError(err)
	s.NotNil(restored)
}

func (s *TransferRepositoryTestSuite) TestDeleteTransferPermanent() {
	transfer, err := s.createSeedTransfer()
	s.Require().NoError(err)
	ctx := context.Background()

	_, err = s.repo.TrashedTransfer(ctx, int(transfer.TransferID))
	s.Require().NoError(err)

	success, err := s.repo.DeleteTransferPermanent(ctx, int(transfer.TransferID))
	s.NoError(err)
	s.True(success)
}

func (s *TransferRepositoryTestSuite) TestRestoreAllTransfer() {
	transfer, err := s.createSeedTransfer()
	s.Require().NoError(err)
	ctx := context.Background()

	_, err = s.repo.TrashedTransfer(ctx, int(transfer.TransferID))
	s.Require().NoError(err)

	success, err := s.repo.RestoreAllTransfer(ctx)
	s.NoError(err)
	s.True(success)
}

func (s *TransferRepositoryTestSuite) TestDeleteAllTransferPermanent() {
	_, err := s.createSeedTransfer()
	s.Require().NoError(err)
	ctx := context.Background()

	success, err := s.repo.DeleteAllTransferPermanent(ctx)
	s.NoError(err)
	s.True(success)
}

func TestTransferRepositorySuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	suite.Run(t, new(TransferRepositoryTestSuite))
}
