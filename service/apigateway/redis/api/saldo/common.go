package saldo_cache

import "time"

const (
	saldoAllCacheKey     = "saldo:all:page:%d:pageSize:%d:search:%s"
	saldoActiveCacheKey  = "saldo:active:page:%d:pageSize:%d:search:%s"
	saldoTrashedCacheKey = "saldo:trashed:page:%d:pageSize:%d:search:%s"
	saldoByIdCacheKey    = "saldo:id:%d"
	saldoByCardNumberKey = "saldo:card_number:%s"

	ttlDefault = 5 * time.Minute
)
