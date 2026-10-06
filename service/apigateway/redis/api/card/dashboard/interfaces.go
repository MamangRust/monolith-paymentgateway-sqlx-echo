package card_dashboard_cache

import (
	"context"

	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

type CardDashboardTotalCache interface {
	GetDashboardCardCache(ctx context.Context) (*response.ApiResponseDashboardCard, bool)
	SetDashboardCardCache(ctx context.Context, data *response.ApiResponseDashboardCard)
	DeleteDashboardCardCache(ctx context.Context)
}

type CardDashboardByCardNumberCache interface {
	SetDashboardCardCardNumberCache(ctx context.Context, cardNumber string, data *response.ApiResponseDashboardCardNumber)
	GetDashboardCardCardNumberCache(ctx context.Context, cardNumber string) (*response.ApiResponseDashboardCardNumber, bool)
	DeleteDashboardCardCardNumberCache(ctx context.Context, cardNumber string)
}
