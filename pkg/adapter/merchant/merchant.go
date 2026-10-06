// Package merchant adapts the Merchant service gRPC API into the shared
// pkg/database sqlc row types used by the consumer repositories.
package merchant

import (
	"context"
	"time"

	pbmerchant "github.com/MamangRust/monolith-payment-gateway-pb/merchant"
	"github.com/MamangRust/monolith-payment-gateway-pkg/adapter"
	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-pkg/resilience"
	"github.com/jackc/pgx/v5/pgtype"
)

// MerchantAdapter reaches the merchant service over gRPC and maps its response
// back to the local SQLC row type, so the transaction service depends on the
// merchant service instead of its tables.
type MerchantAdapter struct {
	QueryClient pbmerchant.MerchantQueryServiceClient
	guard       *resilience.DependencyGuard
}

// NewAdapter wraps the generated merchant query client. Resilience is applied
// per call through the optional dependency guard; a nil guard is a passthrough.
func NewAdapter(queryClient pbmerchant.MerchantQueryServiceClient, opts ...adapter.GuardOption) *MerchantAdapter {
	a := &MerchantAdapter{QueryClient: queryClient}
	for _, opt := range opts {
		opt(a)
	}
	return a
}

// SetGuard implements adapter.GuardSetter.
func (a *MerchantAdapter) SetGuard(g *resilience.DependencyGuard) { a.guard = g }

func (a *MerchantAdapter) FindByApiKey(ctx context.Context, api_key string) (*db.GetMerchantByApiKeyRow, error) {
	resp, err := adapter.Call(a.guard, ctx, func(ctx context.Context) (*pbmerchant.ApiResponseMerchant, error) {
		return a.QueryClient.FindByApiKey(ctx, &pbmerchant.FindByApiKeyRequest{
			ApiKey: api_key,
		})
	})
	if err != nil {
		return nil, err
	}
	if resp == nil || resp.Data == nil {
		return nil, nil
	}

	return &db.GetMerchantByApiKeyRow{
		MerchantID: resp.Data.Id,
		Name:       resp.Data.Name,
		ApiKey:     resp.Data.ApiKey,
		UserID:     resp.Data.UserId,
		Status:     resp.Data.Status,
		CreatedAt:  parseTimestampValue(resp.Data.CreatedAt),
		UpdatedAt:  parseTimestampValue(resp.Data.UpdatedAt),
	}, nil
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
