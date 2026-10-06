package cardbillingmencache

import (
    "context"
    db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
)

type CardBillingCache interface {
    GetByCardNumber(ctx context.Context, cardNumber string) ([]*db.BillingCycle, bool)
    SetByCardNumber(ctx context.Context, cardNumber string, data []*db.BillingCycle)
    DeleteByCardNumber(ctx context.Context, cardNumber string)
}
