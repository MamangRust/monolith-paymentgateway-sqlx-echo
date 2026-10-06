package carddashboardmencache

import (
	"context"

	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

type CardDashboardTotalCache interface {
	GetDashboardCardCache(ctx context.Context) (*response.DashboardCard, bool)
	SetDashboardCardCache(ctx context.Context, data *response.DashboardCard)
	DeleteDashboardCardCache(ctx context.Context)
}

type CardDashboardByCardNumberCache interface {
	SetDashboardCardCardNumberCache(ctx context.Context, cardNumber string, data *response.DashboardCardCardNumber)
	GetDashboardCardCardNumberCache(ctx context.Context, cardNumber string) (*response.DashboardCardCardNumber, bool)
	DeleteDashboardCardCardNumberCache(ctx context.Context, cardNumber string)
}
