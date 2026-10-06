package card_test

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"testing"
	"time"

	"github.com/MamangRust/monolith-payment-gateway-card/repository"
	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	sharederrors "github.com/MamangRust/monolith-payment-gateway-shared/errors"
	user_repo "github.com/MamangRust/monolith-payment-gateway-user/repository"
	tests "github.com/MamangRust/monolith-payment-gateway-test"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jmoiron/sqlx"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/stretchr/testify/suite"
)

type CardRepositoryTestSuite struct {
	suite.Suite
	ts       *tests.TestSuite
	pool     *sqlx.DB
	repo     *repository.Repositories
	userRepo user_repo.Repositories
	userID   int
}

func (s *CardRepositoryTestSuite) SetupSuite() {
	ts, err := tests.SetupTestSuite()
	s.Require().NoError(err)
	s.ts = ts

	pool, err := sqlx.ConnectContext(s.ts.Ctx, "pgx", s.ts.DBURL)
	s.Require().NoError(err)
	s.pool = pool

	queries := pool
	s.repo = repository.NewRepositories(queries, placeholderUserQueryClient())
	s.userRepo = user_repo.NewRepositories(
		queries,
		placeholderRoleQueryClient(),
		placeholderRoleCommandClient(),
	)

	// Create a user for card ownership
	user, err := s.userRepo.UserCommand().CreateUser(context.Background(), &requests.CreateUserRequest{
		FirstName: "Card",
		LastName:  "Owner",
		Email:     fmt.Sprintf("card.owner-%d-%d@example.com", time.Now().UnixNano(), time.Now().UnixNano()%10000),
		Password:  "password123",
	})
	s.Require().NoError(err)
	s.userID = int(user.UserID)
}

func (s *CardRepositoryTestSuite) TearDownSuite() {
	if s.pool != nil {
		s.pool.Close()
	}
	s.ts.Teardown()
}

func (s *CardRepositoryTestSuite) createSeedCard() (*db.CreateCardRow, error) {
	return s.repo.CardCommand.CreateCard(context.Background(), &requests.CreateCardRequest{
		UserID:       s.userID,
		CardType:     "debit",
		ExpireDate:   time.Now().AddDate(5, 0, 0),
		CVV:          "123",
		CardProvider: "Visa",
	})
}

func (s *CardRepositoryTestSuite) TestCreateCard() {
	ctx := context.Background()
	req := &requests.CreateCardRequest{
		UserID:       s.userID,
		CardType:     "debit",
		ExpireDate:   time.Now().AddDate(5, 0, 0),
		CVV:          "123",
		CardProvider: "Visa",
	}

	res, err := s.repo.CardCommand.CreateCard(ctx, req)
	s.NoError(err)
	s.NotNil(res)
}

func (s *CardRepositoryTestSuite) TestFindAllCards() {
	_, err := s.createSeedCard()
	s.Require().NoError(err)
	ctx := context.Background()

	res, err := s.repo.CardQuery.FindAllCards(ctx, &requests.FindAllCards{
		Page:     1,
		PageSize: 10,
		Search:   "",
	})
	s.NoError(err)
	s.GreaterOrEqual(len(res), 1)
}

func (s *CardRepositoryTestSuite) TestFindById() {
	card, err := s.createSeedCard()
	s.Require().NoError(err)
	ctx := context.Background()

	found, err := s.repo.CardQuery.FindById(ctx, int(card.CardID))
	s.NoError(err)
	s.NotNil(found)
	s.Equal(card.CardID, found.CardID)
}

func (s *CardRepositoryTestSuite) TestFindByActive() {
	_, err := s.createSeedCard()
	s.Require().NoError(err)
	ctx := context.Background()

	res, err := s.repo.CardQuery.FindByActive(ctx, &requests.FindAllCards{
		Page:     1,
		PageSize: 10,
		Search:   "",
	})
	s.NoError(err)
	s.GreaterOrEqual(len(res), 1)
}

func (s *CardRepositoryTestSuite) TestFindByTrashed() {
	card, err := s.createSeedCard()
	s.Require().NoError(err)
	ctx := context.Background()

	_, err = s.repo.CardCommand.TrashedCard(ctx, int(card.CardID))
	s.Require().NoError(err)

	res, err := s.repo.CardQuery.FindByTrashed(ctx, &requests.FindAllCards{
		Page:     1,
		PageSize: 10,
		Search:   "",
	})
	s.NoError(err)
	s.GreaterOrEqual(len(res), 1)
}

func (s *CardRepositoryTestSuite) TestUpdateCard() {
	card, err := s.createSeedCard()
	s.Require().NoError(err)
	ctx := context.Background()

	req := &requests.UpdateCardRequest{
		CardID:       int(card.CardID),
		UserID:       s.userID,
		CardType:     "credit",
		ExpireDate:   time.Now().AddDate(6, 0, 0),
		CVV:          "456",
		CardProvider: "MasterCard",
	}

	res, err := s.repo.CardCommand.UpdateCard(ctx, req)
	s.NoError(err)
	s.NotNil(res)
}

func (s *CardRepositoryTestSuite) TestTrashCard() {
	card, err := s.createSeedCard()
	s.Require().NoError(err)
	ctx := context.Background()

	trashed, err := s.repo.CardCommand.TrashedCard(ctx, int(card.CardID))
	s.NoError(err)
	s.NotNil(trashed)
}

func (s *CardRepositoryTestSuite) TestRestoreCard() {
	card, err := s.createSeedCard()
	s.Require().NoError(err)
	ctx := context.Background()

	_, err = s.repo.CardCommand.TrashedCard(ctx, int(card.CardID))
	s.Require().NoError(err)

	restored, err := s.repo.CardCommand.RestoreCard(ctx, int(card.CardID))
	s.NoError(err)
	s.NotNil(restored)
}

func (s *CardRepositoryTestSuite) TestDeleteCardPermanent() {
	card, err := s.createSeedCard()
	s.Require().NoError(err)
	ctx := context.Background()

	_, err = s.repo.CardCommand.TrashedCard(ctx, int(card.CardID))
	s.Require().NoError(err)

	success, err := s.repo.CardCommand.DeleteCardPermanent(ctx, int(card.CardID))
	s.NoError(err)
	s.True(success)
}

func (s *CardRepositoryTestSuite) TestRestoreAllCard() {
	card, err := s.createSeedCard()
	s.Require().NoError(err)
	ctx := context.Background()

	_, err = s.repo.CardCommand.TrashedCard(ctx, int(card.CardID))
	s.Require().NoError(err)

	success, err := s.repo.CardCommand.RestoreAllCard(ctx)
	s.NoError(err)
	s.True(success)
}

func (s *CardRepositoryTestSuite) TestDeleteAllCardPermanent() {
	_, err := s.createSeedCard()
	s.Require().NoError(err)
	ctx := context.Background()

	success, err := s.repo.CardCommand.DeleteAllCardPermanent(ctx)
	s.NoError(err)
	s.True(success)
}

func (s *CardRepositoryTestSuite) TestDuplicateCardNumberConflict() {
	ctx := context.Background()

	// Same statement the service repository issues (cards INSERT plus the saldo
	// outbox event) so the DB-level behavior matches production.
	const createCardSQL = `
WITH inserted_card AS (
    INSERT INTO cards (
        user_id, card_number, card_type, expire_date, cvv, card_provider, created_at, updated_at
    )
    VALUES ($1, $2, $3, $4, $5, $6, current_timestamp, current_timestamp)
    RETURNING card_id, user_id, card_number, card_type, expire_date, cvv, card_provider, created_at, updated_at
), queued_event AS (
    INSERT INTO outbox_events (event_key, topic, message_key, payload)
    SELECT
        'card-create-saldo:' || card_id::text,
        'saldo-service-topic-create-saldo',
        card_id::text,
        jsonb_build_object('card_number', card_number, 'total_balance', 0)
    FROM inserted_card
    ON CONFLICT (event_key) DO NOTHING
    RETURNING event_key
)
SELECT card_id, user_id, card_number, card_type, expire_date, cvv, card_provider, created_at, updated_at
FROM inserted_card
`

	// card_number is auto-generated by the service, so the unique constraint is
	// exercised directly at the DB layer: the second insert with the same
	// number must violate UNIQUE (SQLSTATE 23505) and map to a 409 conflict.
	seed := db.CreateCardParams{
		UserID:       int32(s.userID),
		CardNumber:   "9999999999999999",
		CardType:     "debit",
		ExpireDate:   pgtype.Date{Time: time.Now().AddDate(5, 0, 0), Valid: true},
		Cvv:          "123",
		CardProvider: "Visa",
	}
	cardArgs := []any{
		seed.UserID, seed.CardNumber, seed.CardType, seed.ExpireDate, seed.Cvv, seed.CardProvider,
	}

	var first db.CreateCardRow
	err := sqlx.GetContext(ctx, s.pool, &first, createCardSQL, cardArgs...)
	s.Require().NoError(err)
	s.Require().NotNil(&first)

	_, err = s.pool.ExecContext(ctx, createCardSQL, cardArgs...)
	s.Require().Error(err, "duplicate card_number must violate the unique constraint")

	appErr := sharederrors.ErrConstraintOrFailed(err, "Card", "create card")
	s.Require().NotNil(appErr)
	s.Equal(http.StatusConflict, appErr.Code, "23505 must map to 409")
	s.Contains(appErr.Message, "already exists")
}

func (s *CardRepositoryTestSuite) TestConcurrentCardUpdatesNoCorruption() {
	ctx := context.Background()
	card, err := s.createSeedCard()
	s.Require().NoError(err)

	base := requests.UpdateCardRequest{
		CardID:       int(card.CardID),
		UserID:       s.userID,
		CardType:     "debit",
		ExpireDate:   time.Now().AddDate(5, 0, 0),
		CVV:          "123",
		CardProvider: "Visa",
	}
	g1 := base
	g1.CardType = "credit"
	g2 := base
	g2.CVV = "999"

	// Concurrent full updates serialize at the row level: both must succeed and
	// the final row must be one of the two complete request states (last-write-
	// wins), never a torn/mixed or corrupted combination.
	var wg sync.WaitGroup
	errs := make(chan error, 2)
	wg.Add(2)
	go func() {
		defer wg.Done()
		_, e := s.repo.CardCommand.UpdateCard(ctx, &g1)
		errs <- e
	}()
	go func() {
		defer wg.Done()
		_, e := s.repo.CardCommand.UpdateCard(ctx, &g2)
		errs <- e
	}()
	wg.Wait()
	close(errs)
	for err := range errs {
		s.NoError(err)
	}

	found, err := s.repo.CardQuery.FindById(ctx, int(card.CardID))
	s.Require().NoError(err)
	valid := (found.CardType == "credit" && found.Cvv == "123") ||
		(found.CardType == "debit" && found.Cvv == "999")
	s.True(valid, "torn card state after concurrent updates: %s/%s", found.CardType, found.Cvv)
}

func TestCardRepositorySuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	suite.Run(t, new(CardRepositoryTestSuite))
}
