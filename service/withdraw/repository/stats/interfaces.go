package withdrawstatsrepository

import (
	"context"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
)

type WithdrawStatsStatusRepository interface {
	GetMonthWithdrawStatusSuccess(ctx context.Context, req *requests.MonthStatusWithdraw) ([]*db.GetMonthWithdrawStatusSuccessRow, error)
	GetYearlyWithdrawStatusSuccess(ctx context.Context, year int) ([]*db.GetYearlyWithdrawStatusSuccessRow, error)
	GetMonthWithdrawStatusFailed(ctx context.Context, req *requests.MonthStatusWithdraw) ([]*db.GetMonthWithdrawStatusFailedRow, error)
	GetYearlyWithdrawStatusFailed(ctx context.Context, year int) ([]*db.GetYearlyWithdrawStatusFailedRow, error)
}

type WithdrawStatsAmountRepository interface {
	GetMonthlyWithdraws(ctx context.Context, year int) ([]*db.GetMonthlyWithdrawsRow, error)
	GetYearlyWithdraws(ctx context.Context, year int) ([]*db.GetYearlyWithdrawsRow, error)
}
