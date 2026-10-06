// Package saldo adapts the Saldo service gRPC API into the shared pkg/database
// sqlc row types used by the consumer repositories.
package saldo

import (
	"context"
	"time"

	pbcard "github.com/MamangRust/monolith-payment-gateway-pb/card"
	pbsaldo "github.com/MamangRust/monolith-payment-gateway-pb/saldo"
	"github.com/MamangRust/monolith-payment-gateway-pkg/adapter"
	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-pkg/resilience"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/jackc/pgx/v5/pgtype"
)

// SaldoAdapter reaches the saldo service over gRPC and maps its responses back
// to local SQLC types, so topup/transaction/transfer/withdraw depend on the
// saldo service instead of its tables.
type SaldoAdapter struct {
	QueryClient   pbsaldo.SaldoQueryServiceClient
	CommandClient pbsaldo.SaldoCommandServiceClient
	guard         *resilience.DependencyGuard
}

// NewAdapter wraps the generated saldo clients. Resilience is applied per call
// through the optional dependency guard; a nil guard is a passthrough.
func NewAdapter(queryClient pbsaldo.SaldoQueryServiceClient, commandClient pbsaldo.SaldoCommandServiceClient, opts ...adapter.GuardOption) *SaldoAdapter {
	a := &SaldoAdapter{
		QueryClient:   queryClient,
		CommandClient: commandClient,
	}
	for _, opt := range opts {
		opt(a)
	}
	return a
}

// SetGuard implements adapter.GuardSetter.
func (a *SaldoAdapter) SetGuard(g *resilience.DependencyGuard) { a.guard = g }

func (a *SaldoAdapter) FindByCardNumber(ctx context.Context, card_number string) (*db.Saldo, error) {
	resp, err := adapter.Call(a.guard, ctx, func(ctx context.Context) (*pbsaldo.ApiResponseSaldo, error) {
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

	return MapSaldoResponseToDB(resp.Data), nil
}

func (a *SaldoAdapter) UpdateSaldoBalance(ctx context.Context, request *requests.UpdateSaldoBalance) (*db.UpdateSaldoBalanceRow, error) {
	resp, err := adapter.Call(a.guard, ctx, func(ctx context.Context) (*pbsaldo.ApiResponseSaldo, error) {
		return a.CommandClient.UpdateSaldoBalance(ctx, &pbsaldo.UpdateSaldoBalanceRequest{
			CardNumber:   request.CardNumber,
			TotalBalance: int32(request.TotalBalance),
		})
	})
	if err != nil {
		return nil, err
	}
	if resp == nil || resp.Data == nil {
		return nil, nil
	}

	return &db.UpdateSaldoBalanceRow{
		SaldoID:        resp.Data.SaldoId,
		CardNumber:     resp.Data.CardNumber,
		TotalBalance:   resp.Data.TotalBalance,
		WithdrawAmount: withdrawAmountPtr(resp.Data.WithdrawAmount),
		WithdrawTime:   parseTimestampValue(resp.Data.WithdrawTime),
		CreatedAt:      parseTimestampValue(resp.Data.CreatedAt),
		UpdatedAt:      parseTimestampValue(resp.Data.UpdatedAt),
	}, nil
}

func (a *SaldoAdapter) UpdateSaldoWithdraw(ctx context.Context, request *requests.UpdateSaldoWithdraw) (*db.UpdateSaldoWithdrawRow, error) {
	withdrawAmount := 0
	if request.WithdrawAmount != nil {
		withdrawAmount = *request.WithdrawAmount
	}

	withdrawTime := time.Time{}
	if request.WithdrawTime != nil {
		withdrawTime = *request.WithdrawTime
	}

	resp, err := adapter.Call(a.guard, ctx, func(ctx context.Context) (*pbsaldo.ApiResponseSaldo, error) {
		return a.CommandClient.UpdateSaldoWithdraw(ctx, &pbsaldo.UpdateSaldoWithdrawRequest{
			CardNumber:     request.CardNumber,
			TotalBalance:   int32(request.TotalBalance),
			WithdrawAmount: int32(withdrawAmount),
			WithdrawTime:   withdrawTime.Format(time.RFC3339),
		})
	})
	if err != nil {
		return nil, err
	}
	if resp == nil || resp.Data == nil {
		return nil, nil
	}

	return &db.UpdateSaldoWithdrawRow{
		SaldoID:        resp.Data.SaldoId,
		CardNumber:     resp.Data.CardNumber,
		TotalBalance:   resp.Data.TotalBalance,
		WithdrawAmount: withdrawAmountPtr(resp.Data.WithdrawAmount),
		WithdrawTime:   parseTimestampValue(resp.Data.WithdrawTime),
		CreatedAt:      parseTimestampValue(resp.Data.CreatedAt),
		UpdatedAt:      parseTimestampValue(resp.Data.UpdatedAt),
	}, nil
}

// MapSaldoResponseToDB converts a saldo gRPC response into the local SQLC Saldo
// model so consumers can treat remote reads like local rows.
func MapSaldoResponseToDB(s *pbsaldo.SaldoResponse) *db.Saldo {
	if s == nil {
		return nil
	}

	return &db.Saldo{
		SaldoID:        s.SaldoId,
		CardNumber:     s.CardNumber,
		TotalBalance:   s.TotalBalance,
		WithdrawAmount: withdrawAmountPtr(s.WithdrawAmount),
		WithdrawTime:   parseTimestampValue(s.WithdrawTime),
		CreatedAt:      parseTimestampValue(s.CreatedAt),
		UpdatedAt:      parseTimestampValue(s.UpdatedAt),
	}
}

func withdrawAmountPtr(amount int32) *int32 {
	if amount == 0 {
		return nil
	}
	v := amount
	return &v
}

func parseTimestampValue(ts string) pgtype.Timestamp {
	if ts == "" {
		return pgtype.Timestamp{Valid: false}
	}

	t, err := time.Parse(time.RFC3339, ts)
	if err != nil {
		return pgtype.Timestamp{Valid: false}
	}

	return pgtype.Timestamp{Time: t, Valid: true}
}
