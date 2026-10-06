package withdrawstatscache

import (
	"context"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
)

type WithdrawStatsStatusCache interface {
	GetCachedMonthWithdrawStatusSuccessCache(ctx context.Context, req *requests.MonthStatusWithdraw) ([]*db.GetMonthWithdrawStatusSuccessRow, bool)
	SetCachedMonthWithdrawStatusSuccessCache(ctx context.Context, req *requests.MonthStatusWithdraw, data []*db.GetMonthWithdrawStatusSuccessRow)

	GetCachedYearlyWithdrawStatusSuccessCache(ctx context.Context, year int) ([]*db.GetYearlyWithdrawStatusSuccessRow, bool)
	SetCachedYearlyWithdrawStatusSuccessCache(ctx context.Context, year int, data []*db.GetYearlyWithdrawStatusSuccessRow)

	GetCachedMonthWithdrawStatusFailedCache(ctx context.Context, req *requests.MonthStatusWithdraw) ([]*db.GetMonthWithdrawStatusFailedRow, bool)
	SetCachedMonthWithdrawStatusFailedCache(ctx context.Context, req *requests.MonthStatusWithdraw, data []*db.GetMonthWithdrawStatusFailedRow)

	GetCachedYearlyWithdrawStatusFailedCache(ctx context.Context, year int) ([]*db.GetYearlyWithdrawStatusFailedRow, bool)
	SetCachedYearlyWithdrawStatusFailedCache(ctx context.Context, year int, data []*db.GetYearlyWithdrawStatusFailedRow)
}

type WithdrawStatsAmountCache interface {
	GetCachedMonthlyWithdraws(ctx context.Context, year int) ([]*db.GetMonthlyWithdrawsRow, bool)
	SetCachedMonthlyWithdraws(ctx context.Context, year int, data []*db.GetMonthlyWithdrawsRow)

	GetCachedYearlyWithdraws(ctx context.Context, year int) ([]*db.GetYearlyWithdrawsRow, bool)
	SetCachedYearlyWithdraws(ctx context.Context, year int, data []*db.GetYearlyWithdrawsRow)
}
