// Code migrated from sqlc-generated sources; domain types kept for repository layer.
// source: card.sql

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
	"time"
)

type AddRewardPointsParams struct {
	CardID       int32 `db:"card_id" json:"card_id"`
	RewardPoints int32 `db:"reward_points" json:"reward_points"`
}

type AddRewardPointsRow struct {
	CardID             int32            `db:"card_id" json:"card_id"`
	UserID             int32            `db:"user_id" json:"user_id"`
	CardNumber         string           `db:"card_number" json:"card_number"`
	CardType           string           `db:"card_type" json:"card_type"`
	ExpireDate         pgtype.Date      `db:"expire_date" json:"expire_date"`
	Cvv                string           `db:"cvv" json:"cvv"`
	CardProvider       string           `db:"card_provider" json:"card_provider"`
	Status             string           `db:"status" json:"status"`
	CreditLimit        int32            `db:"credit_limit" json:"credit_limit"`
	OutstandingBalance int32            `db:"outstanding_balance" json:"outstanding_balance"`
	RewardPoints       int32            `db:"reward_points" json:"reward_points"`
	CreatedAt          pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt          pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type CountRecentAuthByCardNumberParams struct {
	CardNumber string           `db:"card_number" json:"card_number"`
	CreatedAt  pgtype.Timestamp `db:"created_at" json:"created_at"`
}

type CreateCardParams struct {
	UserID       int32       `db:"user_id" json:"user_id"`
	CardNumber   string      `db:"card_number" json:"card_number"`
	CardType     string      `db:"card_type" json:"card_type"`
	ExpireDate   pgtype.Date `db:"expire_date" json:"expire_date"`
	Cvv          string      `db:"cvv" json:"cvv"`
	CardProvider string      `db:"card_provider" json:"card_provider"`
}

type CreateCardRow struct {
	CardID       int32            `db:"card_id" json:"card_id"`
	UserID       int32            `db:"user_id" json:"user_id"`
	CardNumber   string           `db:"card_number" json:"card_number"`
	CardType     string           `db:"card_type" json:"card_type"`
	ExpireDate   pgtype.Date      `db:"expire_date" json:"expire_date"`
	Cvv          string           `db:"cvv" json:"cvv"`
	CardProvider string           `db:"card_provider" json:"card_provider"`
	CreatedAt    pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt    pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type GetActiveCardsWithCountParams struct {
	Column1 string `db:"column1" json:"column_1"`
	Limit   int32  `db:"limit" json:"limit"`
	Offset  int32  `db:"offset" json:"offset"`
}

type GetActiveCardsWithCountRow struct {
	CardID       int32            `db:"card_id" json:"card_id"`
	UserID       int32            `db:"user_id" json:"user_id"`
	CardNumber   string           `db:"card_number" json:"card_number"`
	CardType     string           `db:"card_type" json:"card_type"`
	ExpireDate   pgtype.Date      `db:"expire_date" json:"expire_date"`
	Cvv          string           `db:"cvv" json:"cvv"`
	CardProvider string           `db:"card_provider" json:"card_provider"`
	CreatedAt    pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt    pgtype.Timestamp `db:"updated_at" json:"updated_at"`
	DeletedAt    pgtype.Timestamp `db:"deleted_at" json:"deleted_at"`
	TotalCount   int64            `db:"total_count" json:"total_count"`
}

type GetAuthTransactionsByCardNumberParams struct {
	CardNumber string `db:"card_number" json:"card_number"`
	Limit      int32  `db:"limit" json:"limit"`
	Offset     int32  `db:"offset" json:"offset"`
}

type GetAuthTransactionsByCardNumberRow struct {
	AuthID         int32            `db:"auth_id" json:"auth_id"`
	TxnID          string           `db:"txn_id" json:"txn_id"`
	CardNumber     string           `db:"card_number" json:"card_number"`
	MerchantID     int32            `db:"merchant_id" json:"merchant_id"`
	Amount         int64            `db:"amount" json:"amount"`
	Currency       string           `db:"currency" json:"currency"`
	Mcc            string           `db:"mcc" json:"mcc"`
	PosEntryMode   string           `db:"pos_entry_mode" json:"pos_entry_mode"`
	Status         string           `db:"status" json:"status"`
	IdempotencyKey string           `db:"idempotency_key" json:"idempotency_key"`
	RiskScore      int32            `db:"risk_score" json:"risk_score"`
	CreatedAt      pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt      pgtype.Timestamp `db:"updated_at" json:"updated_at"`
	TotalCount     int64            `db:"total_count" json:"total_count"`
}

type GetCardByCardNumberRow struct {
	CardID             int32            `db:"card_id" json:"card_id"`
	UserID             int32            `db:"user_id" json:"user_id"`
	CardNumber         string           `db:"card_number" json:"card_number"`
	CardType           string           `db:"card_type" json:"card_type"`
	ExpireDate         pgtype.Date      `db:"expire_date" json:"expire_date"`
	Cvv                string           `db:"cvv" json:"cvv"`
	CardProvider       string           `db:"card_provider" json:"card_provider"`
	Status             string           `db:"status" json:"status"`
	CreditLimit        int32            `db:"credit_limit" json:"credit_limit"`
	OutstandingBalance int32            `db:"outstanding_balance" json:"outstanding_balance"`
	RewardPoints       int32            `db:"reward_points" json:"reward_points"`
	CreatedAt          pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt          pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type GetCardByCardNumberWithStatusRow struct {
	CardID             int32            `db:"card_id" json:"card_id"`
	UserID             int32            `db:"user_id" json:"user_id"`
	CardNumber         string           `db:"card_number" json:"card_number"`
	CardType           string           `db:"card_type" json:"card_type"`
	ExpireDate         pgtype.Date      `db:"expire_date" json:"expire_date"`
	Cvv                string           `db:"cvv" json:"cvv"`
	CardProvider       string           `db:"card_provider" json:"card_provider"`
	Status             string           `db:"status" json:"status"`
	CreditLimit        int32            `db:"credit_limit" json:"credit_limit"`
	OutstandingBalance int32            `db:"outstanding_balance" json:"outstanding_balance"`
	RewardPoints       int32            `db:"reward_points" json:"reward_points"`
	CreatedAt          pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt          pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type GetCardByIDRow struct {
	CardID             int32            `db:"card_id" json:"card_id"`
	UserID             int32            `db:"user_id" json:"user_id"`
	CardNumber         string           `db:"card_number" json:"card_number"`
	CardType           string           `db:"card_type" json:"card_type"`
	ExpireDate         pgtype.Date      `db:"expire_date" json:"expire_date"`
	Cvv                string           `db:"cvv" json:"cvv"`
	CardProvider       string           `db:"card_provider" json:"card_provider"`
	Status             string           `db:"status" json:"status"`
	CreditLimit        int32            `db:"credit_limit" json:"credit_limit"`
	OutstandingBalance int32            `db:"outstanding_balance" json:"outstanding_balance"`
	RewardPoints       int32            `db:"reward_points" json:"reward_points"`
	CreatedAt          pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt          pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type GetCardByUserIDRow struct {
	CardID             int32            `db:"card_id" json:"card_id"`
	UserID             int32            `db:"user_id" json:"user_id"`
	CardNumber         string           `db:"card_number" json:"card_number"`
	CardType           string           `db:"card_type" json:"card_type"`
	ExpireDate         pgtype.Date      `db:"expire_date" json:"expire_date"`
	Cvv                string           `db:"cvv" json:"cvv"`
	CardProvider       string           `db:"card_provider" json:"card_provider"`
	Status             string           `db:"status" json:"status"`
	CreditLimit        int32            `db:"credit_limit" json:"credit_limit"`
	OutstandingBalance int32            `db:"outstanding_balance" json:"outstanding_balance"`
	RewardPoints       int32            `db:"reward_points" json:"reward_points"`
	CreatedAt          pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt          pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type GetCardPaymentsByCardNumberParams struct {
	CardNumber string `db:"card_number" json:"card_number"`
	Limit      int32  `db:"limit" json:"limit"`
	Offset     int32  `db:"offset" json:"offset"`
}

type GetCardPaymentsByCardNumberRow struct {
	PaymentID      int32            `db:"payment_id" json:"payment_id"`
	PaymentUuid    string           `db:"payment_uuid" json:"payment_uuid"`
	CardNumber     string           `db:"card_number" json:"card_number"`
	BillingID      *int32           `db:"billing_id" json:"billing_id"`
	Amount         int64            `db:"amount" json:"amount"`
	PaymentChannel string           `db:"payment_channel" json:"payment_channel"`
	ReferenceID    string           `db:"reference_id" json:"reference_id"`
	Status         string           `db:"status" json:"status"`
	CreatedAt      pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt      pgtype.Timestamp `db:"updated_at" json:"updated_at"`
	TotalCount     int64            `db:"total_count" json:"total_count"`
}

type GetCardsParams struct {
	Column1 string `db:"column1" json:"column_1"`
	Limit   int32  `db:"limit" json:"limit"`
	Offset  int32  `db:"offset" json:"offset"`
}

type GetCardsRow struct {
	CardID       int32            `db:"card_id" json:"card_id"`
	UserID       int32            `db:"user_id" json:"user_id"`
	CardNumber   string           `db:"card_number" json:"card_number"`
	CardType     string           `db:"card_type" json:"card_type"`
	ExpireDate   pgtype.Date      `db:"expire_date" json:"expire_date"`
	Cvv          string           `db:"cvv" json:"cvv"`
	CardProvider string           `db:"card_provider" json:"card_provider"`
	CreatedAt    pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt    pgtype.Timestamp `db:"updated_at" json:"updated_at"`
	TotalCount   int64            `db:"total_count" json:"total_count"`
}

type GetMonthlyBalancesRow struct {
	Month        string `db:"month" json:"month"`
	TotalBalance int32  `db:"total_balance" json:"total_balance"`
}

type GetMonthlyBalancesByCardNumberParams struct {
	Column1    time.Time `db:"column1" json:"column_1"`
	CardNumber string    `db:"card_number" json:"card_number"`
}

type GetMonthlyBalancesByCardNumberRow struct {
	Month        string `db:"month" json:"month"`
	TotalBalance int32  `db:"total_balance" json:"total_balance"`
}

type GetMonthlyTopupAmountRow struct {
	Month            string `db:"month" json:"month"`
	TotalTopupAmount int32  `db:"total_topup_amount" json:"total_topup_amount"`
}

type GetMonthlyTopupAmountByCardNumberParams struct {
	CardNumber string    `db:"card_number" json:"card_number"`
	Column2    time.Time `db:"column2" json:"column_2"`
}

type GetMonthlyTopupAmountByCardNumberRow struct {
	Month            string `db:"month" json:"month"`
	TotalTopupAmount int32  `db:"total_topup_amount" json:"total_topup_amount"`
}

type GetMonthlyTransactionAmountRow struct {
	Month                  string `db:"month" json:"month"`
	TotalTransactionAmount int32  `db:"total_transaction_amount" json:"total_transaction_amount"`
}

type GetMonthlyTransactionAmountByCardNumberParams struct {
	CardNumber string    `db:"card_number" json:"card_number"`
	Column2    time.Time `db:"column2" json:"column_2"`
}

type GetMonthlyTransactionAmountByCardNumberRow struct {
	Month                  string `db:"month" json:"month"`
	TotalTransactionAmount int32  `db:"total_transaction_amount" json:"total_transaction_amount"`
}

type GetMonthlyTransferAmountByReceiverParams struct {
	TransferTo string    `db:"transfer_to" json:"transfer_to"`
	Column2    time.Time `db:"column2" json:"column_2"`
}

type GetMonthlyTransferAmountByReceiverRow struct {
	Month               string `db:"month" json:"month"`
	TotalReceivedAmount int32  `db:"total_received_amount" json:"total_received_amount"`
}

type GetMonthlyTransferAmountBySenderParams struct {
	TransferFrom string    `db:"transfer_from" json:"transfer_from"`
	Column2      time.Time `db:"column2" json:"column_2"`
}

type GetMonthlyTransferAmountBySenderRow struct {
	Month           string `db:"month" json:"month"`
	TotalSentAmount int32  `db:"total_sent_amount" json:"total_sent_amount"`
}

type GetMonthlyTransferAmountReceiverRow struct {
	Month               string `db:"month" json:"month"`
	TotalReceivedAmount int32  `db:"total_received_amount" json:"total_received_amount"`
}

type GetMonthlyTransferAmountSenderRow struct {
	Month           string `db:"month" json:"month"`
	TotalSentAmount int32  `db:"total_sent_amount" json:"total_sent_amount"`
}

type GetMonthlyWithdrawAmountRow struct {
	Month               string `db:"month" json:"month"`
	TotalWithdrawAmount int32  `db:"total_withdraw_amount" json:"total_withdraw_amount"`
}

type GetMonthlyWithdrawAmountByCardNumberParams struct {
	CardNumber string    `db:"card_number" json:"card_number"`
	Column2    time.Time `db:"column2" json:"column_2"`
}

type GetMonthlyWithdrawAmountByCardNumberRow struct {
	Month               string `db:"month" json:"month"`
	TotalWithdrawAmount int32  `db:"total_withdraw_amount" json:"total_withdraw_amount"`
}

type GetRedeemableRewardIdsRow struct {
	RewardID     int32 `db:"reward_id" json:"reward_id"`
	PointsEarned int32 `db:"points_earned" json:"points_earned"`
}

type GetTrashedCardsWithCountParams struct {
	Column1 string `db:"column1" json:"column_1"`
	Limit   int32  `db:"limit" json:"limit"`
	Offset  int32  `db:"offset" json:"offset"`
}

type GetTrashedCardsWithCountRow struct {
	CardID       int32            `db:"card_id" json:"card_id"`
	UserID       int32            `db:"user_id" json:"user_id"`
	CardNumber   string           `db:"card_number" json:"card_number"`
	CardType     string           `db:"card_type" json:"card_type"`
	ExpireDate   pgtype.Date      `db:"expire_date" json:"expire_date"`
	Cvv          string           `db:"cvv" json:"cvv"`
	CardProvider string           `db:"card_provider" json:"card_provider"`
	CreatedAt    pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt    pgtype.Timestamp `db:"updated_at" json:"updated_at"`
	DeletedAt    pgtype.Timestamp `db:"deleted_at" json:"deleted_at"`
	TotalCount   int64            `db:"total_count" json:"total_count"`
}

type GetUserEmailByCardNumberRow struct {
	CardID       int32            `db:"card_id" json:"card_id"`
	Email        string           `db:"email" json:"email"`
	UserID       int32            `db:"user_id" json:"user_id"`
	CardNumber   string           `db:"card_number" json:"card_number"`
	CardType     string           `db:"card_type" json:"card_type"`
	ExpireDate   pgtype.Date      `db:"expire_date" json:"expire_date"`
	Cvv          string           `db:"cvv" json:"cvv"`
	CardProvider string           `db:"card_provider" json:"card_provider"`
	CreatedAt    pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt    pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type GetYearlyBalancesRow struct {
	Year         pgtype.Numeric `db:"year" json:"year"`
	TotalBalance int64          `db:"total_balance" json:"total_balance"`
}

type GetYearlyBalancesByCardNumberParams struct {
	Column1    interface{} `db:"column1" json:"column_1"`
	CardNumber string      `db:"card_number" json:"card_number"`
}

type GetYearlyBalancesByCardNumberRow struct {
	Year         pgtype.Numeric `db:"year" json:"year"`
	TotalBalance int64          `db:"total_balance" json:"total_balance"`
}

type GetYearlyTopupAmountRow struct {
	Year             pgtype.Numeric `db:"year" json:"year"`
	TotalTopupAmount int64          `db:"total_topup_amount" json:"total_topup_amount"`
}

type GetYearlyTopupAmountByCardNumberParams struct {
	CardNumber string      `db:"card_number" json:"card_number"`
	Column2    interface{} `db:"column2" json:"column_2"`
}

type GetYearlyTopupAmountByCardNumberRow struct {
	Year             pgtype.Numeric `db:"year" json:"year"`
	TotalTopupAmount int64          `db:"total_topup_amount" json:"total_topup_amount"`
}

type GetYearlyTransactionAmountRow struct {
	Year                   pgtype.Numeric `db:"year" json:"year"`
	TotalTransactionAmount int64          `db:"total_transaction_amount" json:"total_transaction_amount"`
}

type GetYearlyTransactionAmountByCardNumberParams struct {
	CardNumber string      `db:"card_number" json:"card_number"`
	Column2    interface{} `db:"column2" json:"column_2"`
}

type GetYearlyTransactionAmountByCardNumberRow struct {
	Year                   pgtype.Numeric `db:"year" json:"year"`
	TotalTransactionAmount int64          `db:"total_transaction_amount" json:"total_transaction_amount"`
}

type GetYearlyTransferAmountByReceiverParams struct {
	TransferTo string      `db:"transfer_to" json:"transfer_to"`
	Column2    interface{} `db:"column2" json:"column_2"`
}

type GetYearlyTransferAmountByReceiverRow struct {
	Year                pgtype.Numeric `db:"year" json:"year"`
	TotalReceivedAmount int64          `db:"total_received_amount" json:"total_received_amount"`
}

type GetYearlyTransferAmountBySenderParams struct {
	TransferFrom string      `db:"transfer_from" json:"transfer_from"`
	Column2      interface{} `db:"column2" json:"column_2"`
}

type GetYearlyTransferAmountBySenderRow struct {
	Year            pgtype.Numeric `db:"year" json:"year"`
	TotalSentAmount int64          `db:"total_sent_amount" json:"total_sent_amount"`
}

type GetYearlyTransferAmountReceiverRow struct {
	Year                pgtype.Numeric `db:"year" json:"year"`
	TotalReceivedAmount int64          `db:"total_received_amount" json:"total_received_amount"`
}

type GetYearlyTransferAmountSenderRow struct {
	Year            pgtype.Numeric `db:"year" json:"year"`
	TotalSentAmount int64          `db:"total_sent_amount" json:"total_sent_amount"`
}

type GetYearlyWithdrawAmountRow struct {
	Year                pgtype.Numeric `db:"year" json:"year"`
	TotalWithdrawAmount int64          `db:"total_withdraw_amount" json:"total_withdraw_amount"`
}

type GetYearlyWithdrawAmountByCardNumberParams struct {
	CardNumber string      `db:"card_number" json:"card_number"`
	Column2    interface{} `db:"column2" json:"column_2"`
}

type GetYearlyWithdrawAmountByCardNumberRow struct {
	Year                pgtype.Numeric `db:"year" json:"year"`
	TotalWithdrawAmount int64          `db:"total_withdraw_amount" json:"total_withdraw_amount"`
}

type InsertAuthTransactionParams struct {
	TxnID          string `db:"txn_id" json:"txn_id"`
	CardNumber     string `db:"card_number" json:"card_number"`
	MerchantID     int32  `db:"merchant_id" json:"merchant_id"`
	Amount         int64  `db:"amount" json:"amount"`
	Currency       string `db:"currency" json:"currency"`
	Mcc            string `db:"mcc" json:"mcc"`
	PosEntryMode   string `db:"pos_entry_mode" json:"pos_entry_mode"`
	IdempotencyKey string `db:"idempotency_key" json:"idempotency_key"`
}

type InsertAuthTransactionRow struct {
	AuthID         int32            `db:"auth_id" json:"auth_id"`
	TxnID          string           `db:"txn_id" json:"txn_id"`
	CardNumber     string           `db:"card_number" json:"card_number"`
	MerchantID     int32            `db:"merchant_id" json:"merchant_id"`
	Amount         int64            `db:"amount" json:"amount"`
	Currency       string           `db:"currency" json:"currency"`
	Mcc            string           `db:"mcc" json:"mcc"`
	PosEntryMode   string           `db:"pos_entry_mode" json:"pos_entry_mode"`
	Status         string           `db:"status" json:"status"`
	IdempotencyKey string           `db:"idempotency_key" json:"idempotency_key"`
	RiskScore      int32            `db:"risk_score" json:"risk_score"`
	CreatedAt      pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt      pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type InsertBillingCycleParams struct {
	CardNumber string    `db:"card_number" json:"card_number"`
	CycleStart time.Time `db:"cycle_start" json:"cycle_start"`
	CycleEnd   time.Time `db:"cycle_end" json:"cycle_end"`
	AmountDue  int32     `db:"amount_due" json:"amount_due"`
	DueDate    time.Time `db:"due_date" json:"due_date"`
}

type InsertCardPaymentParams struct {
	PaymentUuid    string `db:"payment_uuid" json:"payment_uuid"`
	CardNumber     string `db:"card_number" json:"card_number"`
	BillingID      *int32 `db:"billing_id" json:"billing_id"`
	Amount         int64  `db:"amount" json:"amount"`
	PaymentChannel string `db:"payment_channel" json:"payment_channel"`
	ReferenceID    string `db:"reference_id" json:"reference_id"`
}

type InsertCardPaymentRow struct {
	PaymentID      int32            `db:"payment_id" json:"payment_id"`
	PaymentUuid    string           `db:"payment_uuid" json:"payment_uuid"`
	CardNumber     string           `db:"card_number" json:"card_number"`
	BillingID      *int32           `db:"billing_id" json:"billing_id"`
	Amount         int64            `db:"amount" json:"amount"`
	PaymentChannel string           `db:"payment_channel" json:"payment_channel"`
	ReferenceID    string           `db:"reference_id" json:"reference_id"`
	Status         string           `db:"status" json:"status"`
	CreatedAt      pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt      pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type InsertCardRewardParams struct {
	CardNumber   string    `db:"card_number" json:"card_number"`
	TxnID        string    `db:"txn_id" json:"txn_id"`
	Amount       int64     `db:"amount" json:"amount"`
	Mcc          string    `db:"mcc" json:"mcc"`
	PointsEarned int32     `db:"points_earned" json:"points_earned"`
	ExpiresAt    time.Time `db:"expires_at" json:"expires_at"`
}

type ProcessBillingCyclesRow struct {
	BillingID  int32            `db:"billing_id" json:"billing_id"`
	CardNumber string           `db:"card_number" json:"card_number"`
	CycleStart time.Time        `db:"cycle_start" json:"cycle_start"`
	CycleEnd   time.Time        `db:"cycle_end" json:"cycle_end"`
	AmountDue  int32            `db:"amount_due" json:"amount_due"`
	DueDate    time.Time        `db:"due_date" json:"due_date"`
	Status     string           `db:"status" json:"status"`
	CreatedAt  pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt  pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type RedeemRewardPointsParams struct {
	CardID       int32 `db:"card_id" json:"card_id"`
	RewardPoints int32 `db:"reward_points" json:"reward_points"`
}

type RedeemRewardPointsRow struct {
	CardID             int32            `db:"card_id" json:"card_id"`
	UserID             int32            `db:"user_id" json:"user_id"`
	CardNumber         string           `db:"card_number" json:"card_number"`
	CardType           string           `db:"card_type" json:"card_type"`
	ExpireDate         pgtype.Date      `db:"expire_date" json:"expire_date"`
	Cvv                string           `db:"cvv" json:"cvv"`
	CardProvider       string           `db:"card_provider" json:"card_provider"`
	Status             string           `db:"status" json:"status"`
	CreditLimit        int32            `db:"credit_limit" json:"credit_limit"`
	OutstandingBalance int32            `db:"outstanding_balance" json:"outstanding_balance"`
	RewardPoints       int32            `db:"reward_points" json:"reward_points"`
	CreatedAt          pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt          pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type RestoreCardRow struct {
	CardID       int32            `db:"card_id" json:"card_id"`
	UserID       int32            `db:"user_id" json:"user_id"`
	CardNumber   string           `db:"card_number" json:"card_number"`
	CardType     string           `db:"card_type" json:"card_type"`
	ExpireDate   pgtype.Date      `db:"expire_date" json:"expire_date"`
	Cvv          string           `db:"cvv" json:"cvv"`
	CardProvider string           `db:"card_provider" json:"card_provider"`
	CreatedAt    pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt    pgtype.Timestamp `db:"updated_at" json:"updated_at"`
	DeletedAt    pgtype.Timestamp `db:"deleted_at" json:"deleted_at"`
}

type TrashCardRow struct {
	CardID             int32            `db:"card_id" json:"card_id"`
	UserID             int32            `db:"user_id" json:"user_id"`
	CardNumber         string           `db:"card_number" json:"card_number"`
	CardType           string           `db:"card_type" json:"card_type"`
	ExpireDate         pgtype.Date      `db:"expire_date" json:"expire_date"`
	Cvv                string           `db:"cvv" json:"cvv"`
	CardProvider       string           `db:"card_provider" json:"card_provider"`
	Status             string           `db:"status" json:"status"`
	CreditLimit        int32            `db:"credit_limit" json:"credit_limit"`
	OutstandingBalance int32            `db:"outstanding_balance" json:"outstanding_balance"`
	RewardPoints       int32            `db:"reward_points" json:"reward_points"`
	CreatedAt          pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt          pgtype.Timestamp `db:"updated_at" json:"updated_at"`
	DeletedAt          pgtype.Timestamp `db:"deleted_at" json:"deleted_at"`
}

type UpdateAuthTransactionRiskScoreParams struct {
	TxnID     string `db:"txn_id" json:"txn_id"`
	RiskScore int32  `db:"risk_score" json:"risk_score"`
}

type UpdateBillingCycleStatusParams struct {
	BillingID int32  `db:"billing_id" json:"billing_id"`
	Status    string `db:"status" json:"status"`
	Status_2  string `db:"status_2" json:"status_2"`
}

type UpdateCardParams struct {
	CardID       int32       `db:"card_id" json:"card_id"`
	CardType     string      `db:"card_type" json:"card_type"`
	ExpireDate   pgtype.Date `db:"expire_date" json:"expire_date"`
	Cvv          string      `db:"cvv" json:"cvv"`
	CardProvider string      `db:"card_provider" json:"card_provider"`
}

type UpdateCardRow struct {
	CardID       int32            `db:"card_id" json:"card_id"`
	UserID       int32            `db:"user_id" json:"user_id"`
	CardNumber   string           `db:"card_number" json:"card_number"`
	CardType     string           `db:"card_type" json:"card_type"`
	ExpireDate   pgtype.Date      `db:"expire_date" json:"expire_date"`
	Cvv          string           `db:"cvv" json:"cvv"`
	CardProvider string           `db:"card_provider" json:"card_provider"`
	CreatedAt    pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt    pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type UpdateCardStatusParams struct {
	CardID int32  `db:"card_id" json:"card_id"`
	Status string `db:"status" json:"status"`
}

type UpdateCardStatusRow struct {
	CardID             int32            `db:"card_id" json:"card_id"`
	UserID             int32            `db:"user_id" json:"user_id"`
	CardNumber         string           `db:"card_number" json:"card_number"`
	CardType           string           `db:"card_type" json:"card_type"`
	ExpireDate         pgtype.Date      `db:"expire_date" json:"expire_date"`
	Cvv                string           `db:"cvv" json:"cvv"`
	CardProvider       string           `db:"card_provider" json:"card_provider"`
	Status             string           `db:"status" json:"status"`
	CreditLimit        int32            `db:"credit_limit" json:"credit_limit"`
	OutstandingBalance int32            `db:"outstanding_balance" json:"outstanding_balance"`
	RewardPoints       int32            `db:"reward_points" json:"reward_points"`
	CreatedAt          pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt          pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type UpdateCreditLimitParams struct {
	CardID      int32 `db:"card_id" json:"card_id"`
	CreditLimit int32 `db:"credit_limit" json:"credit_limit"`
}

type UpdateCreditLimitRow struct {
	CardID             int32            `db:"card_id" json:"card_id"`
	UserID             int32            `db:"user_id" json:"user_id"`
	CardNumber         string           `db:"card_number" json:"card_number"`
	CardType           string           `db:"card_type" json:"card_type"`
	ExpireDate         pgtype.Date      `db:"expire_date" json:"expire_date"`
	Cvv                string           `db:"cvv" json:"cvv"`
	CardProvider       string           `db:"card_provider" json:"card_provider"`
	Status             string           `db:"status" json:"status"`
	CreditLimit        int32            `db:"credit_limit" json:"credit_limit"`
	OutstandingBalance int32            `db:"outstanding_balance" json:"outstanding_balance"`
	RewardPoints       int32            `db:"reward_points" json:"reward_points"`
	CreatedAt          pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt          pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type UpdateOutstandingBalanceParams struct {
	CardID             int32 `db:"card_id" json:"card_id"`
	OutstandingBalance int32 `db:"outstanding_balance" json:"outstanding_balance"`
}

type UpdateOutstandingBalanceRow struct {
	CardID             int32            `db:"card_id" json:"card_id"`
	UserID             int32            `db:"user_id" json:"user_id"`
	CardNumber         string           `db:"card_number" json:"card_number"`
	CardType           string           `db:"card_type" json:"card_type"`
	ExpireDate         pgtype.Date      `db:"expire_date" json:"expire_date"`
	Cvv                string           `db:"cvv" json:"cvv"`
	CardProvider       string           `db:"card_provider" json:"card_provider"`
	Status             string           `db:"status" json:"status"`
	CreditLimit        int32            `db:"credit_limit" json:"credit_limit"`
	OutstandingBalance int32            `db:"outstanding_balance" json:"outstanding_balance"`
	RewardPoints       int32            `db:"reward_points" json:"reward_points"`
	CreatedAt          pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt          pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}
