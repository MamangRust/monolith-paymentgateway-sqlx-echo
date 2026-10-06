package service

import (
	"context"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
)

func (s *merchantCommandService) invalidateMerchantCaches(ctx context.Context, merchantID, userID int32, apiKey string) {
	if s.cache == nil {
		return
	}

	s.cache.DeleteCachedMerchant(ctx, int(merchantID))
	s.cache.InvalidateMerchantListCaches(ctx)
	s.cache.DeleteCachedMerchantByUserID(ctx, int(userID))
	if apiKey != "" {
		s.cache.DeleteCachedMerchantByAPIKey(ctx, apiKey)
	}
}

func (s *merchantCommandService) invalidateMerchantCachesForRow(ctx context.Context, merchant *db.Merchant) {
	if merchant == nil {
		return
	}
	s.invalidateMerchantCaches(ctx, merchant.MerchantID, merchant.UserID, merchant.ApiKey)
}
