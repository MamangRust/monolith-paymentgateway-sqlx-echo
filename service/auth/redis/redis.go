package mencache

import (
	sharedcachehelpers "github.com/MamangRust/monolith-payment-gateway-shared/cache"
)

// Mencache is a struct that holds various cache interfaces.
type Mencache struct {
	IdentityCache      IdentityCache
	LoginCache         LoginCache
	PasswordResetCache PasswordResetCache
	RegisterCache      RegisterCache
}

func NewMencache(cacheStore *sharedcachehelpers.CacheStore) *Mencache {

	return &Mencache{
		IdentityCache:      NewidentityCache(cacheStore),
		LoginCache:         NewLoginCache(cacheStore),
		PasswordResetCache: NewPasswordResetCache(cacheStore),
		RegisterCache:      NewRegisterCache(cacheStore),
	}
}
