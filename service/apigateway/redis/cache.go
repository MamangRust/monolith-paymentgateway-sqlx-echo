package mencache

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/cache"
)

type CacheApiGateway interface {
	MerchantCache
	RoleCache
}

type mencacheApiGateay struct {
	MerchantCache
	RoleCache
}

func NewCacheApiGateway(cacheStore *cache.CacheStore) CacheApiGateway {

	return &mencacheApiGateay{
		MerchantCache: NewMerchantCache(cacheStore),
		RoleCache:     NewRoleCache(cacheStore),
	}
}
