package repository

import (
	"context"
	"time"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
)

//go:generate mockgen -source=interfaces.go -destination=mocks/card.go
type CardCommandRepository interface {
	CreateCard(ctx context.Context, request *requests.CreateCardRequest) (*db.CreateCardRow, error)
	UpdateCard(ctx context.Context, request *requests.UpdateCardRequest) (*db.UpdateCardRow, error)
	TrashedCard(ctx context.Context, cardId int) (*db.TrashCardRow, error)
	RestoreCard(ctx context.Context, cardId int) (*db.RestoreCardRow, error)
	DeleteCardPermanent(ctx context.Context, cardId int) (bool, error)
	RestoreAllCard(ctx context.Context) (bool, error)
	DeleteAllCardPermanent(ctx context.Context) (bool, error)
	ToggleCardStatus(ctx context.Context, request *requests.ToggleCardStatusRequest) (*db.UpdateCardStatusRow, error)
	UpdateCreditLimit(ctx context.Context, request *requests.UpdateCreditLimitRequest) (*db.UpdateCreditLimitRow, error)
	RedeemPoints(ctx context.Context, request *requests.RedeemPointsRequest) (*db.RedeemRewardPointsRow, error)
	ProcessBillingCycles(ctx context.Context, billingCycleDay int) ([]*db.BillingCycle, error)
}

type CardQueryRepository interface {
	FindAllCards(ctx context.Context, req *requests.FindAllCards) ([]*db.GetCardsRow, error)
	FindByActive(ctx context.Context, req *requests.FindAllCards) ([]*db.GetActiveCardsWithCountRow, error)
	FindByTrashed(ctx context.Context, req *requests.FindAllCards) ([]*db.GetTrashedCardsWithCountRow, error)
	FindById(ctx context.Context, card_id int) (*db.GetCardByIDRow, error)
	FindCardByUserId(ctx context.Context, user_id int) (*db.GetCardByUserIDRow, error)
	FindCardByCardNumber(ctx context.Context, card_number string) (*db.GetCardByCardNumberRow, error)
	FindUserCardByCardNumber(ctx context.Context, card_number string) (*db.GetUserEmailByCardNumberRow, error)
}

type UserRepository interface {
	FindById(ctx context.Context, user_id int) (*db.GetUserByIDRow, error)
}

// CardAuthTransactionRepository handles auth transaction persistence.
type CardAuthTransactionRepository interface {
	InsertPending(ctx context.Context, req *requests.AuthorizeCardRequest) (*db.CardAuthTransaction, error)
	Approve(ctx context.Context, txnID string) (*db.CardAuthTransaction, error)
	Decline(ctx context.Context, txnID string) (*db.CardAuthTransaction, error)
	Reverse(ctx context.Context, txnID string) (*db.CardAuthTransaction, error)
	FindByIdempotencyKey(ctx context.Context, key string) (*db.CardAuthTransaction, error)
	FindByTxnID(ctx context.Context, txnID string) (*db.CardAuthTransaction, error)
	FindByCardNumber(ctx context.Context, cardNumber string, page, pageSize int) ([]*db.GetAuthTransactionsByCardNumberRow, error)
	CountRecentByCardNumber(ctx context.Context, cardNumber string, since time.Time) (int, error)
	UpdateRiskScore(ctx context.Context, txnID string, score int) error
}

// CardPaymentResult contains a payment and whether it was returned from an
// idempotent replay. Replays must not publish duplicate side-effects.
type CardPaymentResult struct {
	Payment  *db.CardPayment
	Replayed bool
}

// CardPaymentRepository handles card payment persistence.
type CardPaymentRepository interface {
	PostPayment(ctx context.Context, req *requests.PostPaymentRequest) (*db.CardPayment, error)
	PostPaymentIdempotent(ctx context.Context, req *requests.PostPaymentRequest, expectedBillingStatus, targetBillingStatus string) (*CardPaymentResult, error)
	GetPaymentByReferenceID(ctx context.Context, referenceID string) (*db.CardPayment, error)
	GetPaymentHistory(ctx context.Context, cardNumber string, page, pageSize int) ([]*db.GetCardPaymentsByCardNumberRow, error)
	CountPayments(ctx context.Context, cardNumber string) (int, error)
}

// CardRewardRepository handles reward points persistence.
type CardRewardRepository interface {
	EarnRewards(ctx context.Context, req *requests.EarnRewardsRequest) (*db.CardReward, error)
	GetBalance(ctx context.Context, cardNumber string) (int64, error)
	GetHistory(ctx context.Context, cardNumber string) ([]*db.CardReward, error)
	RedeemRewards(ctx context.Context, cardNumber string, points int64) (int64, error)
}

// BillingCycleRepository handles billing cycle persistence.
type BillingCycleRepository interface {
	GetBillingCyclesByCardNumber(ctx context.Context, cardNumber string) ([]*db.BillingCycle, error)
	GetBillingCycleByID(ctx context.Context, billingID int) (*db.BillingCycle, error)
}
