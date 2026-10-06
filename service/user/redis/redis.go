package mencache

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/cache"
)

type Mencache interface {
	UserQueryCache
	UserCommandCache
}

// Mencache is a struct that holds user query and user command caches.
type mencache struct {
	UserQueryCache
	UserCommandCache
}

// NewMencache creates a new instance of Mencache using the given dependencies.
// It creates a new cache store using the given context, redis client, and logger,
// and returns a Mencache struct with initialized caches for user query and user command.
func NewMencache(cacheStore *cache.CacheStore) Mencache {
	return &mencache{
		UserQueryCache:   NewUserQueryCache(cacheStore),
		UserCommandCache: NewUserCommandCache(cacheStore),
	}
}
