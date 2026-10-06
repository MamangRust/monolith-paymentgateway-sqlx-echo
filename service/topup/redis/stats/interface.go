package topupstatscache

import (
	"context"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
)

type TopupStatsStatusCache interface {
	GetMonthTopupStatusSuccessCache(ctx context.Context, req *requests.MonthTopupStatus) ([]*db.GetMonthTopupStatusSuccessRow, bool)
	SetMonthTopupStatusSuccessCache(ctx context.Context, req *requests.MonthTopupStatus, data []*db.GetMonthTopupStatusSuccessRow)

	GetYearlyTopupStatusSuccessCache(ctx context.Context, year int) ([]*db.GetYearlyTopupStatusSuccessRow, bool)
	SetYearlyTopupStatusSuccessCache(ctx context.Context, year int, data []*db.GetYearlyTopupStatusSuccessRow)

	GetMonthTopupStatusFailedCache(ctx context.Context, req *requests.MonthTopupStatus) ([]*db.GetMonthTopupStatusFailedRow, bool)
	SetMonthTopupStatusFailedCache(ctx context.Context, req *requests.MonthTopupStatus, data []*db.GetMonthTopupStatusFailedRow)

	GetYearlyTopupStatusFailedCache(ctx context.Context, year int) ([]*db.GetYearlyTopupStatusFailedRow, bool)
	SetYearlyTopupStatusFailedCache(ctx context.Context, year int, data []*db.GetYearlyTopupStatusFailedRow)
}

type TopupStatsMethodCache interface {
	GetMonthlyTopupMethodsCache(ctx context.Context, year int) ([]*db.GetMonthlyTopupMethodsRow, bool)
	SetMonthlyTopupMethodsCache(ctx context.Context, year int, data []*db.GetMonthlyTopupMethodsRow)

	GetYearlyTopupMethodsCache(ctx context.Context, year int) ([]*db.GetYearlyTopupMethodsRow, bool)
	SetYearlyTopupMethodsCache(ctx context.Context, year int, data []*db.GetYearlyTopupMethodsRow)
}

type TopupStatsAmountCache interface {
	GetMonthlyTopupAmountsCache(ctx context.Context, year int) ([]*db.GetMonthlyTopupAmountsRow, bool)
	SetMonthlyTopupAmountsCache(ctx context.Context, year int, data []*db.GetMonthlyTopupAmountsRow)

	GetYearlyTopupAmountsCache(ctx context.Context, year int) ([]*db.GetYearlyTopupAmountsRow, bool)
	SetYearlyTopupAmountsCache(ctx context.Context, year int, data []*db.GetYearlyTopupAmountsRow)
}
