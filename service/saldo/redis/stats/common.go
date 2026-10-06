package saldostatscache

import "time"

const (
	saldoMonthTotalBalanceCacheKey = "saldo:month_total_balance:month:%d:year:%d"
	saldoYearTotalBalanceCacheKey  = "saldo:year_total_balance:year:%d"
	saldoYearlyBalanceCacheKey     = "saldo:yearly_balance:year:%d"
	saldoMonthBalanceCacheKey      = "saldo:month_balance:year:%d"

	ttlDefault = 5 * time.Minute
)
