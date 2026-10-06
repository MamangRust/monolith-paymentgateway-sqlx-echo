package repository

import (
	"context"
	"time"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	sharedErrors "github.com/MamangRust/monolith-payment-gateway-shared/errors"
	card_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/card_errors/repository"
	"github.com/jmoiron/sqlx"
)

const insertCardReward = `INSERT INTO card_rewards (
    card_number,
    txn_id,
    amount,
    mcc,
    points_earned,
    expires_at,
    redeemed
) VALUES ($1, $2, $3, $4, $5, $6, FALSE)
RETURNING reward_id, card_number, txn_id, amount, mcc, points_earned, expires_at, redeemed, created_at`

const getCardRewardBalance = `SELECT COALESCE(SUM(points_earned), 0)::BIGINT AS total_points
FROM card_rewards
WHERE card_number = $1
  AND redeemed = FALSE
  AND expires_at > CURRENT_TIMESTAMP`

const getCardRewardHistory = `SELECT reward_id, card_number, txn_id, amount, mcc, points_earned, expires_at, redeemed, created_at
FROM card_rewards
WHERE card_number = $1
ORDER BY created_at DESC`

const getRedeemableRewardIds = `SELECT reward_id, points_earned
FROM card_rewards
WHERE card_number = $1
  AND redeemed = FALSE
  AND expires_at > CURRENT_TIMESTAMP
ORDER BY created_at`

const markRewardsRedeemed = `WITH updated AS (
    UPDATE card_rewards
    SET redeemed = TRUE
    WHERE reward_id = ANY($1::INT[])
    RETURNING points_earned
)
SELECT COALESCE(SUM(points_earned), 0)::BIGINT AS redeemed_points FROM updated`

type cardRewardRepository struct {
	db *sqlx.DB
}

func NewCardRewardRepository(q *sqlx.DB) CardRewardRepository {
	return &cardRewardRepository{db: q}
}

// earnPoints calculates points based on amount and MCC category.
func earnPoints(amount int64, mcc string) int32 {
	base := int32(amount / 10000)
	switch mcc {
	case "5812", "4511", "5541":
		base *= 2
	}
	if base < 1 {
		base = 1
	}
	return base
}

func (r *cardRewardRepository) EarnRewards(ctx context.Context, req *requests.EarnRewardsRequest) (*db.CardReward, error) {
	points := earnPoints(req.Amount, req.Mcc)

	var res db.CardReward
	err := r.db.GetContext(ctx, &res, insertCardReward,
		req.CardNumber,
		req.TxnID,
		req.Amount,
		req.Mcc,
		points,
		time.Now().AddDate(1, 0, 0),
	)
	if err != nil {
		return nil, card_errors.ErrEarnRewardsFailed.WithInternal(err)
	}

	return &res, nil
}

func (r *cardRewardRepository) GetBalance(ctx context.Context, cardNumber string) (int64, error) {
	var balance int64
	err := r.db.GetContext(ctx, &balance, getCardRewardBalance, cardNumber)
	if err != nil {
		return 0, card_errors.ErrGetRewardBalanceFailed.WithInternal(err)
	}

	return balance, nil
}

func (r *cardRewardRepository) GetHistory(ctx context.Context, cardNumber string) ([]*db.CardReward, error) {
	var res []*db.CardReward
	err := r.db.SelectContext(ctx, &res, getCardRewardHistory, cardNumber)
	if err != nil {
		return nil, card_errors.ErrGetRewardHistoryFailed.WithInternal(err)
	}

	return res, nil
}

func (r *cardRewardRepository) RedeemRewards(ctx context.Context, cardNumber string, points int64) (int64, error) {
	var rows []*db.GetRedeemableRewardIdsRow
	err := r.db.SelectContext(ctx, &rows, getRedeemableRewardIds, cardNumber)
	if err != nil {
		return 0, card_errors.ErrRedeemRewardsFailed.WithInternal(err)
	}

	var ids []int32
	var total int64
	for _, row := range rows {
		if total >= points {
			break
		}
		ids = append(ids, row.RewardID)
		total += int64(row.PointsEarned)
	}

	if total < points {
		return 0, sharedErrors.ErrBadRequest.WithMessage("Insufficient reward points")
	}

	var redeemed int64
	err = r.db.GetContext(ctx, &redeemed, markRewardsRedeemed, ids)
	if err != nil {
		return 0, card_errors.ErrRedeemRewardsFailed.WithInternal(err)
	}

	return redeemed, nil
}
