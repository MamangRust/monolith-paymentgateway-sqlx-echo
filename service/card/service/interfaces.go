package service

import (
	"context"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

type CardAuthorizationService interface {
	Authorize(ctx context.Context, request *requests.AuthorizeCardRequest) (*db.CardAuthTransaction, error)
	Reverse(ctx context.Context, request *requests.ReverseTransactionRequest) (*db.CardAuthTransaction, error)
	GetAuthTransaction(ctx context.Context, txnID string) (*db.CardAuthTransaction, error)
	GetAuthTransactionsByCardNumber(ctx context.Context, cardNumber string, page, pageSize int) ([]*db.GetAuthTransactionsByCardNumberRow, error)
}

type CardPaymentService interface {
	PostPayment(ctx context.Context, request *requests.PostPaymentRequest) (*db.CardPayment, error)
	GetPaymentHistory(ctx context.Context, cardNumber string, page, pageSize int) ([]*db.GetCardPaymentsByCardNumberRow, error)
	CountPayments(ctx context.Context, cardNumber string) (int, error)
}

type CardRewardService interface {
	EarnRewards(ctx context.Context, request *requests.EarnRewardsRequest) (*db.CardReward, error)
	GetBalance(ctx context.Context, cardNumber string) (int64, error)
	GetHistory(ctx context.Context, cardNumber string) ([]*db.CardReward, error)
	RedeemRewards(ctx context.Context, cardNumber string, points int64) (int64, error)
}

type BillingEngineService interface {
	TriggerBillingCycle(ctx context.Context, billingCycleDay int) (int, error)
	GetStatement(ctx context.Context, cardNumber string) (*db.BillingCycle, error)
	GetStatementsByCard(ctx context.Context, cardNumber string, page, pageSize int) ([]*db.BillingCycle, error)
	GetBillingCyclesByCardNumber(ctx context.Context, cardNumber string) ([]*db.BillingCycle, error)
}

type CardQueryService interface {
	FindAll(ctx context.Context, req *requests.FindAllCards) ([]*db.GetCardsRow, *int, error)
	FindByActive(ctx context.Context, req *requests.FindAllCards) ([]*db.GetActiveCardsWithCountRow, *int, error)
	FindByTrashed(ctx context.Context, req *requests.FindAllCards) ([]*db.GetTrashedCardsWithCountRow, *int, error)
	FindById(ctx context.Context, card_id int) (*db.GetCardByIDRow, error)
	FindByUserID(ctx context.Context, userID int) (*db.GetCardByUserIDRow, error)
	FindByCardNumber(ctx context.Context, card_number string) (*db.GetCardByCardNumberRow, error)
	FindUserCardByCardNumber(ctx context.Context, card_number string) (*db.GetUserEmailByCardNumberRow, error)
}

type CardDashboardService interface {
	DashboardCard(ctx context.Context) (*response.DashboardCard, error)
	DashboardCardCardNumber(ctx context.Context, cardNumber string) (*response.DashboardCardCardNumber, error)
}

type CardCommandService interface {
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
	ProcessBillingCycles(ctx context.Context) error
}
