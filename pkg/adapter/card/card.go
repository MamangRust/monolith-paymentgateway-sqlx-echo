// Package card adapts the Card service gRPC API into the shared pkg/database
// sqlc row types used by the consumer repositories.
package card

import (
	"context"
	"time"

	pbcard "github.com/MamangRust/monolith-payment-gateway-pb/card"
	"github.com/MamangRust/monolith-payment-gateway-pkg/adapter"
	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-pkg/resilience"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/jackc/pgx/v5/pgtype"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// CardAdapter reaches the card service over gRPC and maps its responses back to
// the local SQLC row types, so services that need card data depend on the card
// service instead of its tables. It deliberately implements the union of the
// consumer-owned CardRepository interfaces (topup, transaction, transfer,
// withdraw); each consumer only calls the subset it declares.
type CardAdapter struct {
	QueryClient   pbcard.CardQueryServiceClient
	CommandClient pbcard.CardCommandServiceClient
	guard         *resilience.DependencyGuard
}

// NewAdapter wraps the generated card clients. Resilience is applied per call
// through the optional dependency guard; a nil guard is a passthrough.
func NewAdapter(queryClient pbcard.CardQueryServiceClient, commandClient pbcard.CardCommandServiceClient, opts ...adapter.GuardOption) *CardAdapter {
	a := &CardAdapter{
		QueryClient:   queryClient,
		CommandClient: commandClient,
	}
	for _, opt := range opts {
		opt(a)
	}
	return a
}

// SetGuard implements adapter.GuardSetter.
func (a *CardAdapter) SetGuard(g *resilience.DependencyGuard) { a.guard = g }

func (a *CardAdapter) FindCardByUserId(ctx context.Context, user_id int) (*db.GetCardByUserIDRow, error) {
	resp, err := adapter.Call(a.guard, ctx, func(ctx context.Context) (*pbcard.ApiResponseCard, error) {
		return a.QueryClient.FindByUserIdCard(ctx, &pbcard.FindByUserIdCardRequest{
			UserId: int32(user_id),
		})
	})
	if err != nil {
		return nil, err
	}
	if resp == nil || resp.Data == nil {
		return nil, nil
	}

	return &db.GetCardByUserIDRow{
		CardID:             resp.Data.Id,
		UserID:             resp.Data.UserId,
		CardNumber:         resp.Data.CardNumber,
		CardType:           resp.Data.CardType,
		ExpireDate:         parseDate(resp.Data.ExpireDate),
		Cvv:                resp.Data.Cvv,
		CardProvider:       resp.Data.CardProvider,
		Status:             resp.Data.Status,
		CreditLimit:        resp.Data.CreditLimit,
		OutstandingBalance: resp.Data.OutstandingBalance,
		RewardPoints:       resp.Data.RewardPoints,
	}, nil
}

func (a *CardAdapter) FindUserCardByCardNumber(ctx context.Context, card_number string) (*db.GetUserEmailByCardNumberRow, error) {
	resp, err := adapter.Call(a.guard, ctx, func(ctx context.Context) (*pbcard.CardWithEmailResponse, error) {
		return a.QueryClient.FindUserCardByCardNumber(ctx, &pbcard.FindByCardNumberRequest{
			CardNumber: card_number,
		})
	})
	if err != nil {
		return nil, err
	}
	if resp == nil {
		return nil, nil
	}

	return &db.GetUserEmailByCardNumberRow{
		CardID:       resp.Id,
		UserID:       resp.UserId,
		CardNumber:   resp.CardNumber,
		CardType:     resp.CardType,
		ExpireDate:   parseDate(resp.ExpireDate),
		Cvv:          resp.Cvv,
		CardProvider: resp.CardProvider,
		Email:        resp.Email,
	}, nil
}

func (a *CardAdapter) FindCardByCardNumber(ctx context.Context, card_number string) (*db.GetCardByCardNumberRow, error) {
	resp, err := adapter.Call(a.guard, ctx, func(ctx context.Context) (*pbcard.ApiResponseCard, error) {
		return a.QueryClient.FindByCardNumber(ctx, &pbcard.FindByCardNumberRequest{
			CardNumber: card_number,
		})
	})
	if err != nil {
		return nil, err
	}
	if resp == nil || resp.Data == nil {
		return nil, nil
	}

	return &db.GetCardByCardNumberRow{
		CardID:             resp.Data.Id,
		UserID:             resp.Data.UserId,
		CardNumber:         resp.Data.CardNumber,
		CardType:           resp.Data.CardType,
		ExpireDate:         parseDate(resp.Data.ExpireDate),
		Cvv:                resp.Data.Cvv,
		CardProvider:       resp.Data.CardProvider,
		Status:             resp.Data.Status,
		CreditLimit:        resp.Data.CreditLimit,
		OutstandingBalance: resp.Data.OutstandingBalance,
		RewardPoints:       resp.Data.RewardPoints,
	}, nil
}

func (a *CardAdapter) UpdateCard(ctx context.Context, request *requests.UpdateCardRequest) (*db.UpdateCardRow, error) {
	resp, err := adapter.Call(a.guard, ctx, func(ctx context.Context) (*pbcard.ApiResponseCard, error) {
		return a.CommandClient.UpdateCard(ctx, &pbcard.UpdateCardRequest{
			CardId:       int32(request.CardID),
			UserId:       int32(request.UserID),
			CardType:     request.CardType,
			ExpireDate:   timestamppb.New(request.ExpireDate),
			Cvv:          request.CVV,
			CardProvider: request.CardProvider,
		})
	})
	if err != nil {
		return nil, err
	}
	if resp == nil || resp.Data == nil {
		return nil, nil
	}

	return &db.UpdateCardRow{
		CardID:       resp.Data.Id,
		UserID:       resp.Data.UserId,
		CardNumber:   resp.Data.CardNumber,
		CardType:     resp.Data.CardType,
		ExpireDate:   parseDate(resp.Data.ExpireDate),
		Cvv:          resp.Data.Cvv,
		CardProvider: resp.Data.CardProvider,
	}, nil
}

// UpdateCardOutstandingBalance is declared by the transaction service's
// CardRepository but has no caller and no dedicated RPC on the card service.
// Outstanding balance is managed inside the card service during
// authorization/capture, so this stays a compile-satisfying no-op rather than
// reaching into the cards table. Add a card RPC first if a real caller appears.
func (a *CardAdapter) UpdateCardOutstandingBalance(ctx context.Context, cardID int, outstandingBalance int) (*db.UpdateOutstandingBalanceRow, error) {
	return &db.UpdateOutstandingBalanceRow{}, nil
}

// AddRewardPoints is likewise declared but unused. The card service exposes
// points changes through its own commands (e.g. RedeemPoints), so there is no
// safe RPC to map "add points" onto; keep it a no-op instead of misusing
// UpdateCreditLimit.
func (a *CardAdapter) AddRewardPoints(ctx context.Context, cardID int, points int) (*db.AddRewardPointsRow, error) {
	return &db.AddRewardPointsRow{}, nil
}

// parseDate converts the card handlers' RFC3339 wire format into the
// pgtype.Date used by the local SQLC rows.
func parseDate(ts string) pgtype.Date {
	if ts == "" {
		return pgtype.Date{Valid: false}
	}

	t, err := time.Parse(time.RFC3339, ts)
	if err != nil {
		return pgtype.Date{Valid: false}
	}

	return pgtype.Date{Time: t, Valid: true}
}
