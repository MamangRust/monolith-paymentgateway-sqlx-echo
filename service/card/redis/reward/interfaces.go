package cardrewardmencache

import (
    "context"
    db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
)

type CardRewardCache interface {
    GetBalance(ctx context.Context, cardNumber string) (int64, bool)
    SetBalance(ctx context.Context, cardNumber string, balance int64)
    DeleteBalance(ctx context.Context, cardNumber string)
    
    GetHistory(ctx context.Context, cardNumber string) ([]*db.CardReward, bool)
    SetHistory(ctx context.Context, cardNumber string, data []*db.CardReward)
    DeleteHistory(ctx context.Context, cardNumber string)
}
