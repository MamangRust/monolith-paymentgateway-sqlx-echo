package cardstatsbycardmencache

import "time"

// cardStatisticByCardCache is a struct that represents the cache store
const (
	expirationCardStatistic = 10 * time.Minute

	cacheKeyMonthlyBalanceByCard  = "stat:monthly_balance_by_card:card_number:%s:year:%d"
	cacheKeyYearlyBalanceByCard   = "stat:yearly_balance_by_card:card_number:%s:year:%d"
	cacheKeyMonthlyTopupByCard    = "stat:monthly_topup_by_card:card_number:%s:year:%d"
	cacheKeyYearlyTopupByCard     = "stat:yearly_topup_by_card:card_number:%s:year:%d"
	cacheKeyMonthlyWithdrawByCard = "stat:monthly_withdraw_by_card:card_number:%s:year:%d"
	cacheKeyYearlyWithdrawByCard  = "stat:yearly_withdraw_by_card:card_number:%s:year:%d"
	cacheKeyMonthlyTxnByCard      = "stat:monthly_txn_by_card:card_number:%s:year:%d"
	cacheKeyYearlyTxnByCard       = "stat:yearly_txn_by_card:card_number:%s:year:%d"
	cacheKeyMonthlySenderByCard   = "stat:monthly_sender_by_card:card_number:%s:year:%d"
	cacheKeyYearlySenderByCard    = "stat:yearly_sender_by_card:card_number:%s:year:%d"
	cacheKeyMonthlyReceiverByCard = "stat:monthly_receiver_by_card:card_number:%s:year:%d"
	cacheKeyYearlyReceiverByCard  = "stat:yearly_receiver_by_card:card_number:%s:year:%d"
)
