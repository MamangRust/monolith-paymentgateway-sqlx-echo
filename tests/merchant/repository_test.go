package merchant_test

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"testing"
	"time"

	"github.com/MamangRust/monolith-payment-gateway-merchant/repository"
	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	sharederrors "github.com/MamangRust/monolith-payment-gateway-shared/errors"
	user_repo "github.com/MamangRust/monolith-payment-gateway-user/repository"
	tests "github.com/MamangRust/monolith-payment-gateway-test"

	"github.com/jmoiron/sqlx"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/stretchr/testify/suite"
)

type MerchantRepositoryTestSuite struct {
	suite.Suite
	ts       *tests.TestSuite
	pool     *sqlx.DB
	repo     repository.Repositories
	userRepo user_repo.Repositories
	userID   int
	apiKey   string
}

func (s *MerchantRepositoryTestSuite) SetupSuite() {
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

	// Seed User
	user, err := s.userRepo.UserCommand().CreateUser(context.Background(), &requests.CreateUserRequest{
		FirstName: "Merchant",
		LastName:  "Owner",
		Email:     fmt.Sprintf("merchant.owner-%d@example.com", time.Now().UnixNano()),
		Password:  "password123",
	})
	s.Require().NoError(err)
	s.userID = int(user.UserID)
}

func (s *MerchantRepositoryTestSuite) TearDownSuite() {
	if s.pool != nil {
		s.pool.Close()
	}
	s.ts.Teardown()
}

func (s *MerchantRepositoryTestSuite) createSeedMerchant() (*db.CreateMerchantRow, error) {
	m, err := s.repo.CreateMerchant(context.Background(), &requests.CreateMerchantRequest{
		Name:   fmt.Sprintf("Test Merchant-%d", time.Now().UnixNano()),
		UserID: s.userID,
	})
	if err == nil && s.apiKey == "" {
		// grab the api_key for later transaction tests
		found, ferr := s.repo.FindByMerchantId(context.Background(), int(m.MerchantID))
		if ferr == nil {
			s.apiKey = found.ApiKey
		}
	}
	return m, err
}

func (s *MerchantRepositoryTestSuite) createSeedDocument(merchantID int) (*db.CreateMerchantDocumentRow, error) {
	return s.repo.CreateMerchantDocument(context.Background(), &requests.CreateMerchantDocumentRequest{
		MerchantID:   merchantID,
		DocumentType: "Identity Proof",
		DocumentUrl:  fmt.Sprintf("http://example.com/doc-%d.pdf", time.Now().UnixNano()),
	})
}

// ────────────────────────────────────────────────────────────────────────────────
// MerchantQuery tests
// ────────────────────────────────────────────────────────────────────────────────

func (s *MerchantRepositoryTestSuite) TestCreateMerchant() {
	ctx := context.Background()
	res, err := s.repo.CreateMerchant(ctx, &requests.CreateMerchantRequest{
		Name:   "Test Merchant",
		UserID: s.userID,
	})
	s.NoError(err)
	s.NotNil(res)
	s.Equal(int32(s.userID), res.UserID)
}

func (s *MerchantRepositoryTestSuite) TestFindAllMerchants() {
	_, err := s.createSeedMerchant()
	s.Require().NoError(err)
	ctx := context.Background()

	res, err := s.repo.FindAllMerchants(ctx, &requests.FindAllMerchants{
		Page:     1,
		PageSize: 10,
		Search:   "",
	})
	s.NoError(err)
	s.GreaterOrEqual(len(res), 1)
}

func (s *MerchantRepositoryTestSuite) TestFindByMerchantId() {
	merchant, err := s.createSeedMerchant()
	s.Require().NoError(err)
	ctx := context.Background()

	found, err := s.repo.FindByMerchantId(ctx, int(merchant.MerchantID))
	s.NoError(err)
	s.NotNil(found)
	s.Equal(merchant.MerchantID, found.MerchantID)
}

func (s *MerchantRepositoryTestSuite) TestFindByApiKey() {
	merchant, err := s.createSeedMerchant()
	s.Require().NoError(err)
	ctx := context.Background()

	full, err := s.repo.FindByMerchantId(ctx, int(merchant.MerchantID))
	s.Require().NoError(err)

	found, err := s.repo.FindByApiKey(ctx, full.ApiKey)
	s.NoError(err)
	s.NotNil(found)
	s.Equal(full.ApiKey, found.ApiKey)
}

func (s *MerchantRepositoryTestSuite) TestFindByName() {
	merchant, err := s.createSeedMerchant()
	s.Require().NoError(err)
	ctx := context.Background()

	found, err := s.repo.FindByName(ctx, merchant.Name)
	s.NoError(err)
	s.NotNil(found)
	s.Equal(merchant.Name, found.Name)
}

func (s *MerchantRepositoryTestSuite) TestFindByMerchantUserId() {
	_, err := s.createSeedMerchant()
	s.Require().NoError(err)
	ctx := context.Background()

	res, err := s.repo.FindByMerchantUserId(ctx, s.userID)
	s.NoError(err)
	s.GreaterOrEqual(len(res), 1)
}

func (s *MerchantRepositoryTestSuite) TestFindByActive() {
	_, err := s.createSeedMerchant()
	s.Require().NoError(err)
	ctx := context.Background()

	res, err := s.repo.FindByActive(ctx, &requests.FindAllMerchants{
		Page:     1,
		PageSize: 10,
		Search:   "",
	})
	s.NoError(err)
	s.GreaterOrEqual(len(res), 1)
}

func (s *MerchantRepositoryTestSuite) TestFindByTrashed() {
	merchant, err := s.createSeedMerchant()
	s.Require().NoError(err)
	ctx := context.Background()

	_, err = s.repo.TrashedMerchant(ctx, int(merchant.MerchantID))
	s.Require().NoError(err)

	res, err := s.repo.FindByTrashed(ctx, &requests.FindAllMerchants{
		Page:     1,
		PageSize: 10,
		Search:   "",
	})
	s.NoError(err)
	s.GreaterOrEqual(len(res), 1)
}

// ────────────────────────────────────────────────────────────────────────────────
// MerchantCommand tests
// ────────────────────────────────────────────────────────────────────────────────

func (s *MerchantRepositoryTestSuite) TestUpdateMerchant() {
	merchant, err := s.createSeedMerchant()
	s.Require().NoError(err)
	ctx := context.Background()

	id := int(merchant.MerchantID)
	res, err := s.repo.UpdateMerchant(ctx, &requests.UpdateMerchantRequest{
		MerchantID: &id,
		Name:       "Updated Merchant",
		UserID:     s.userID,
		Status:     "active",
	})
	s.NoError(err)
	s.NotNil(res)
	s.Equal("Updated Merchant", res.Name)
}

func (s *MerchantRepositoryTestSuite) TestUpdateMerchantStatus() {
	merchant, err := s.createSeedMerchant()
	s.Require().NoError(err)
	ctx := context.Background()

	id := int(merchant.MerchantID)
	res, err := s.repo.UpdateMerchantStatus(ctx, &requests.UpdateMerchantStatusRequest{
		MerchantID: &id,
		Status:     "inactive",
	})
	s.NoError(err)
	s.NotNil(res)
	s.Equal("inactive", res.Status)
}

func (s *MerchantRepositoryTestSuite) TestTrashMerchant() {
	merchant, err := s.createSeedMerchant()
	s.Require().NoError(err)
	ctx := context.Background()

	trashed, err := s.repo.TrashedMerchant(ctx, int(merchant.MerchantID))
	s.NoError(err)
	s.NotNil(trashed)
}

func (s *MerchantRepositoryTestSuite) TestRestoreMerchant() {
	merchant, err := s.createSeedMerchant()
	s.Require().NoError(err)
	ctx := context.Background()

	_, err = s.repo.TrashedMerchant(ctx, int(merchant.MerchantID))
	s.Require().NoError(err)

	restored, err := s.repo.RestoreMerchant(ctx, int(merchant.MerchantID))
	s.NoError(err)
	s.NotNil(restored)
}

func (s *MerchantRepositoryTestSuite) TestDeleteMerchantPermanent() {
	merchant, err := s.createSeedMerchant()
	s.Require().NoError(err)
	ctx := context.Background()

	_, err = s.repo.TrashedMerchant(ctx, int(merchant.MerchantID))
	s.Require().NoError(err)

	ok, err := s.repo.DeleteMerchantPermanent(ctx, int(merchant.MerchantID))
	s.NoError(err)
	s.True(ok)
}

func (s *MerchantRepositoryTestSuite) TestRestoreAllMerchant() {
	merchant, err := s.createSeedMerchant()
	s.Require().NoError(err)
	ctx := context.Background()

	_, err = s.repo.TrashedMerchant(ctx, int(merchant.MerchantID))
	s.Require().NoError(err)

	ok, err := s.repo.RestoreAllMerchant(ctx)
	s.NoError(err)
	s.True(ok)
}

func (s *MerchantRepositoryTestSuite) TestDeleteAllMerchantPermanent() {
	_, err := s.createSeedMerchant()
	s.Require().NoError(err)
	ctx := context.Background()

	ok, err := s.repo.DeleteAllMerchantPermanent(ctx)
	s.NoError(err)
	s.True(ok)
}

// ────────────────────────────────────────────────────────────────────────────────
// MerchantDocumentQuery tests
// ────────────────────────────────────────────────────────────────────────────────

func (s *MerchantRepositoryTestSuite) TestCreateMerchantDocument() {
	merchant, err := s.createSeedMerchant()
	s.Require().NoError(err)
	ctx := context.Background()

	doc, err := s.repo.CreateMerchantDocument(ctx, &requests.CreateMerchantDocumentRequest{
		MerchantID:   int(merchant.MerchantID),
		DocumentType: "Identity Proof",
		DocumentUrl:  "http://example.com/doc.pdf",
	})
	s.NoError(err)
	s.NotNil(doc)
	s.Equal(merchant.MerchantID, doc.MerchantID)
}

func (s *MerchantRepositoryTestSuite) TestFindAllDocuments() {
	merchant, err := s.createSeedMerchant()
	s.Require().NoError(err)
	_, err = s.createSeedDocument(int(merchant.MerchantID))
	s.Require().NoError(err)
	ctx := context.Background()

	res, err := s.repo.FindAllDocuments(ctx, &requests.FindAllMerchantDocuments{
		Page:     1,
		PageSize: 10,
		Search:   "",
	})
	s.NoError(err)
	s.GreaterOrEqual(len(res), 1)
}

func (s *MerchantRepositoryTestSuite) TestFindByIdDocument() {
	merchant, err := s.createSeedMerchant()
	s.Require().NoError(err)
	doc, err := s.createSeedDocument(int(merchant.MerchantID))
	s.Require().NoError(err)
	ctx := context.Background()

	found, err := s.repo.FindByIdDocument(ctx, int(doc.DocumentID))
	s.NoError(err)
	s.NotNil(found)
	s.Equal(doc.DocumentID, found.DocumentID)
}

func (s *MerchantRepositoryTestSuite) TestFindByActiveDocuments() {
	merchant, err := s.createSeedMerchant()
	s.Require().NoError(err)
	_, err = s.createSeedDocument(int(merchant.MerchantID))
	s.Require().NoError(err)
	ctx := context.Background()

	res, err := s.repo.FindByActiveDocuments(ctx, &requests.FindAllMerchantDocuments{
		Page:     1,
		PageSize: 10,
		Search:   "",
	})
	s.NoError(err)
	s.GreaterOrEqual(len(res), 1)
}

func (s *MerchantRepositoryTestSuite) TestFindByTrashedDocuments() {
	merchant, err := s.createSeedMerchant()
	s.Require().NoError(err)
	doc, err := s.createSeedDocument(int(merchant.MerchantID))
	s.Require().NoError(err)
	ctx := context.Background()

	_, err = s.repo.TrashedMerchantDocument(ctx, int(doc.DocumentID))
	s.Require().NoError(err)

	res, err := s.repo.FindByTrashedDocuments(ctx, &requests.FindAllMerchantDocuments{
		Page:     1,
		PageSize: 10,
		Search:   "",
	})
	s.NoError(err)
	s.GreaterOrEqual(len(res), 1)
}

// ────────────────────────────────────────────────────────────────────────────────
// MerchantDocumentCommand tests
// ────────────────────────────────────────────────────────────────────────────────

func (s *MerchantRepositoryTestSuite) TestUpdateMerchantDocument() {
	merchant, err := s.createSeedMerchant()
	s.Require().NoError(err)
	doc, err := s.createSeedDocument(int(merchant.MerchantID))
	s.Require().NoError(err)
	ctx := context.Background()

	id := int(doc.DocumentID)
	res, err := s.repo.UpdateMerchantDocument(ctx, &requests.UpdateMerchantDocumentRequest{
		DocumentID:   &id,
		MerchantID:   int(merchant.MerchantID),
		DocumentType: "Updated Type",
		DocumentUrl:  "http://example.com/updated.pdf",
		Status:       "pending",
		Note:         "Please re-upload",
	})
	s.NoError(err)
	s.NotNil(res)
	s.Equal("Updated Type", res.DocumentType)
}

func (s *MerchantRepositoryTestSuite) TestUpdateMerchantDocumentStatus() {
	merchant, err := s.createSeedMerchant()
	s.Require().NoError(err)
	doc, err := s.createSeedDocument(int(merchant.MerchantID))
	s.Require().NoError(err)
	ctx := context.Background()

	id := int(doc.DocumentID)
	res, err := s.repo.UpdateMerchantDocumentStatus(ctx, &requests.UpdateMerchantDocumentStatusRequest{
		DocumentID: &id,
		MerchantID: int(merchant.MerchantID),
		Status:     "approved",
		Note:       "All good",
	})
	s.NoError(err)
	s.NotNil(res)
	s.Equal("approved", res.Status)
}

func (s *MerchantRepositoryTestSuite) TestTrashedMerchantDocument() {
	merchant, err := s.createSeedMerchant()
	s.Require().NoError(err)
	doc, err := s.createSeedDocument(int(merchant.MerchantID))
	s.Require().NoError(err)
	ctx := context.Background()

	trashed, err := s.repo.TrashedMerchantDocument(ctx, int(doc.DocumentID))
	s.NoError(err)
	s.NotNil(trashed)
}

func (s *MerchantRepositoryTestSuite) TestRestoreMerchantDocument() {
	merchant, err := s.createSeedMerchant()
	s.Require().NoError(err)
	doc, err := s.createSeedDocument(int(merchant.MerchantID))
	s.Require().NoError(err)
	ctx := context.Background()

	_, err = s.repo.TrashedMerchantDocument(ctx, int(doc.DocumentID))
	s.Require().NoError(err)

	restored, err := s.repo.RestoreMerchantDocument(ctx, int(doc.DocumentID))
	s.NoError(err)
	s.NotNil(restored)
}

func (s *MerchantRepositoryTestSuite) TestDeleteMerchantDocumentPermanent() {
	merchant, err := s.createSeedMerchant()
	s.Require().NoError(err)
	doc, err := s.createSeedDocument(int(merchant.MerchantID))
	s.Require().NoError(err)
	ctx := context.Background()

	_, err = s.repo.TrashedMerchantDocument(ctx, int(doc.DocumentID))
	s.Require().NoError(err)

	ok, err := s.repo.DeleteMerchantDocumentPermanent(ctx, int(doc.DocumentID))
	s.NoError(err)
	s.True(ok)
}

func (s *MerchantRepositoryTestSuite) TestRestoreAllMerchantDocument() {
	merchant, err := s.createSeedMerchant()
	s.Require().NoError(err)
	doc, err := s.createSeedDocument(int(merchant.MerchantID))
	s.Require().NoError(err)
	ctx := context.Background()

	_, err = s.repo.TrashedMerchantDocument(ctx, int(doc.DocumentID))
	s.Require().NoError(err)

	ok, err := s.repo.RestoreAllMerchantDocument(ctx)
	s.NoError(err)
	s.True(ok)
}

func (s *MerchantRepositoryTestSuite) TestDeleteAllMerchantDocumentPermanent() {
	merchant, err := s.createSeedMerchant()
	s.Require().NoError(err)
	_, err = s.createSeedDocument(int(merchant.MerchantID))
	s.Require().NoError(err)
	ctx := context.Background()

	ok, err := s.repo.DeleteAllMerchantDocumentPermanent(ctx)
	s.NoError(err)
	s.True(ok)
}

// ────────────────────────────────────────────────────────────────────────────────
// MerchantTransaction tests
// ────────────────────────────────────────────────────────────────────────────────

func (s *MerchantRepositoryTestSuite) seedTransaction(cardNumber string, merchantID int32, apiKey string) {
	_, err := s.pool.ExecContext(context.Background(),
		"INSERT INTO transactions (card_number, merchant_id, amount, payment_method, transaction_time, status) VALUES ($1, $2, $3, 'bank_transfer', $4, 'success')",
		cardNumber, merchantID, 100000, time.Now(),
	)
	s.Require().NoError(err)
}

func (s *MerchantRepositoryTestSuite) TestFindAllTransactions() {
	ctx := context.Background()

	// Seed card + merchant + transaction
	var userID int32
	err := s.pool.QueryRowContext(ctx, "INSERT INTO users (firstname, lastname, email, password, verification_code, is_verified) VALUES ('Trx', 'Test', $1, 'pass', '123', true) RETURNING user_id",
		fmt.Sprintf("trx.test-%d@example.com", time.Now().UnixNano())).Scan(&userID)
	s.Require().NoError(err)

	cardNumber := fmt.Sprintf("%016d", time.Now().UnixNano()%1e16)
	_, err = s.pool.ExecContext(ctx, "INSERT INTO cards (user_id, card_number, card_type, cvv, card_provider, expire_date) VALUES ($1, $2, 'debit', '123', 'visa', '2030-01-01')", userID, cardNumber)
	s.Require().NoError(err)

	merchant, err := s.createSeedMerchant()
	s.Require().NoError(err)
	full, err := s.repo.FindByMerchantId(ctx, int(merchant.MerchantID))
	s.Require().NoError(err)

	s.seedTransaction(cardNumber, merchant.MerchantID, full.ApiKey)

	res, err := s.repo.FindAllTransactions(ctx, &requests.FindAllMerchantTransactions{
		Page:     1,
		PageSize: 10,
		Search:   "",
	})
	s.NoError(err)
	s.GreaterOrEqual(len(res), 1)
}

func (s *MerchantRepositoryTestSuite) TestFindAllTransactionsByMerchant() {
	ctx := context.Background()

	var userID int32
	err := s.pool.QueryRowContext(ctx, "INSERT INTO users (firstname, lastname, email, password, verification_code, is_verified) VALUES ('TrxMerchant', 'Test', $1, 'pass', '123', true) RETURNING user_id",
		fmt.Sprintf("trxmerchant.test-%d@example.com", time.Now().UnixNano())).Scan(&userID)
	s.Require().NoError(err)

	cardNumber := fmt.Sprintf("%016d", (time.Now().UnixNano()+1)%1e16)
	_, err = s.pool.ExecContext(ctx, "INSERT INTO cards (user_id, card_number, card_type, cvv, card_provider, expire_date) VALUES ($1, $2, 'debit', '321', 'mastercard', '2030-01-01')", userID, cardNumber)
	s.Require().NoError(err)

	merchant, err := s.createSeedMerchant()
	s.Require().NoError(err)
	full, err := s.repo.FindByMerchantId(ctx, int(merchant.MerchantID))
	s.Require().NoError(err)

	s.seedTransaction(cardNumber, merchant.MerchantID, full.ApiKey)

	res, err := s.repo.FindAllTransactionsByMerchant(ctx, &requests.FindAllMerchantTransactionsById{
		MerchantID: int(merchant.MerchantID),
		Page:       1,
		PageSize:   10,
		Search:     "",
	})
	s.NoError(err)
	s.GreaterOrEqual(len(res), 1)
}

func (s *MerchantRepositoryTestSuite) TestFindAllTransactionsByApikey() {
	ctx := context.Background()

	var userID int32
	err := s.pool.QueryRowContext(ctx, "INSERT INTO users (firstname, lastname, email, password, verification_code, is_verified) VALUES ('TrxApiKey', 'Test', $1, 'pass', '123', true) RETURNING user_id",
		fmt.Sprintf("trxapikey.test-%d@example.com", time.Now().UnixNano())).Scan(&userID)
	s.Require().NoError(err)

	cardNumber := fmt.Sprintf("%016d", (time.Now().UnixNano()+2)%1e16)
	_, err = s.pool.ExecContext(ctx, "INSERT INTO cards (user_id, card_number, card_type, cvv, card_provider, expire_date) VALUES ($1, $2, 'debit', '456', 'visa', '2030-01-01')", userID, cardNumber)
	s.Require().NoError(err)

	merchant, err := s.createSeedMerchant()
	s.Require().NoError(err)
	full, err := s.repo.FindByMerchantId(ctx, int(merchant.MerchantID))
	s.Require().NoError(err)

	s.seedTransaction(cardNumber, merchant.MerchantID, full.ApiKey)

	res, err := s.repo.FindAllTransactionsByApikey(ctx, &requests.FindAllMerchantTransactionsByApiKey{
		ApiKey:   full.ApiKey,
		Page:     1,
		PageSize: 10,
		Search:   "",
	})
	s.NoError(err)
	s.GreaterOrEqual(len(res), 1)
}

func (s *MerchantRepositoryTestSuite) TestDuplicateApiKeyConflict() {
	ctx := context.Background()

	// Same statement the service repository issues (merchants INSERT plus the
	// email outbox event) so the DB-level behavior matches production.
	const createMerchantSQL = `
WITH valid_user AS (
    SELECT user_id, email
    FROM users
    WHERE users.user_id = $3 AND users.deleted_at IS NULL
), inserted_merchant AS (
    INSERT INTO merchants (name, api_key, user_id, status, created_at, updated_at)
    SELECT $1, $2, valid_user.user_id, $4, current_timestamp, current_timestamp
    FROM valid_user
    RETURNING merchant_id, merchant_no, name, api_key, user_id, status, created_at, updated_at
), queued_event AS (
    INSERT INTO outbox_events (event_key, topic, message_key, payload)
    SELECT
        'merchant-create:' || merchant_id::text,
        'email-service-topic-merchant-create',
        merchant_id::text,
        jsonb_build_object(
            'event_id', gen_random_uuid()::text,
            'schema_version', 1,
            'event_type', 'merchant.created',
            'occurred_at', now(),
            'email', u.email,
            'subject', 'Initial Verification - SanEdge',
            'body', 'Your merchant account has been created successfully. Please upload the required documents for verification.'
        )
    FROM inserted_merchant m
    JOIN valid_user u ON u.user_id = m.user_id
    ON CONFLICT (event_key) DO NOTHING
    RETURNING event_key
)
SELECT merchant_id, merchant_no, name, api_key, user_id, status, created_at, updated_at
FROM inserted_merchant
`

	// api_key is auto-generated by the service, so the unique constraint is
	// exercised directly at the DB layer: the second insert with the same key
	// must violate UNIQUE (SQLSTATE 23505) and map to a 409 conflict.
	seed := db.CreateMerchantParams{
		Name:   "Duplicate Key Merchant",
		ApiKey: "dup-api-key-409",
		UserID: int32(s.userID),
		Status: "inactive",
	}
	merchantArgs := []any{
		seed.Name, seed.ApiKey, seed.UserID, seed.Status,
	}

	var first db.CreateMerchantRow
	err := sqlx.GetContext(ctx, s.pool, &first, createMerchantSQL, merchantArgs...)
	s.Require().NoError(err)
	s.Require().NotNil(&first)

	_, err = s.pool.ExecContext(ctx, createMerchantSQL, merchantArgs...)
	s.Require().Error(err, "duplicate api_key must violate the unique constraint")

	appErr := sharederrors.ErrConstraintOrFailed(err, "Merchant", "create merchant")
	s.Require().NotNil(appErr)
	s.Equal(http.StatusConflict, appErr.Code, "23505 must map to 409")
	s.Contains(appErr.Message, "already exists")
}

func (s *MerchantRepositoryTestSuite) TestConcurrentMerchantUpdatesNoCorruption() {
	ctx := context.Background()
	merchant, err := s.createSeedMerchant()
	s.Require().NoError(err)
	id := int(merchant.MerchantID)

	g1 := &requests.UpdateMerchantRequest{MerchantID: &id, Name: "Concurrent Merchant A", UserID: s.userID, Status: "active"}
	g2 := &requests.UpdateMerchantRequest{MerchantID: &id, Name: "Concurrent Merchant B", UserID: s.userID, Status: "inactive"}

	// Concurrent full updates serialize at the row level: both must succeed and
	// the final row must be one of the two complete request states (last-write-
	// wins), never a torn/mixed or corrupted combination.
	var wg sync.WaitGroup
	errs := make(chan error, 2)
	wg.Add(2)
	go func() {
		defer wg.Done()
		_, e := s.repo.UpdateMerchant(ctx, g1)
		errs <- e
	}()
	go func() {
		defer wg.Done()
		_, e := s.repo.UpdateMerchant(ctx, g2)
		errs <- e
	}()
	wg.Wait()
	close(errs)
	for err := range errs {
		s.NoError(err)
	}

	found, err := s.repo.FindByMerchantId(ctx, id)
	s.Require().NoError(err)
	valid := (found.Name == "Concurrent Merchant A" && found.Status == "active") ||
		(found.Name == "Concurrent Merchant B" && found.Status == "inactive")
	s.True(valid, "torn merchant state after concurrent updates: %s/%s", found.Name, found.Status)
}

func TestMerchantRepositorySuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	suite.Run(t, new(MerchantRepositoryTestSuite))
}
