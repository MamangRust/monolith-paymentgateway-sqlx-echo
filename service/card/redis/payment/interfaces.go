package cardpaymentmencache

import (
    "context"
    db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
)

type CardPaymentCache interface {
    GetPaymentHistory(ctx context.Context, cardNumber string, page, pageSize int) ([]*db.GetCardPaymentsByCardNumberRow, bool)
    SetPaymentHistory(ctx context.Context, cardNumber string, page, pageSize int, data []*db.GetCardPaymentsByCardNumberRow)
    DeletePaymentHistory(ctx context.Context, cardNumber string)
    
    GetPaymentCount(ctx context.Context, cardNumber string) (int, bool)
    SetPaymentCount(ctx context.Context, cardNumber string, count int)
    DeletePaymentCount(ctx context.Context, cardNumber string)
}
