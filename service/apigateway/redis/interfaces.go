package mencache

import "context"

type MerchantCache interface {
	GetMerchantCache(ctx context.Context, apiKey string) (string, bool)
	SetMerchantCache(ctx context.Context, merchantID string, apiKey string)
}

type RoleCache interface {
	GetRoleCache(ctx context.Context, userID string) ([]string, bool)
	SetRoleCache(ctx context.Context, userID string, roles []string)
}
