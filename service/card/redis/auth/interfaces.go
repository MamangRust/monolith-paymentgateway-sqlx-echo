package cardauthmencache

import (
	"context"
	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
)

type CardAuthCache interface {
	GetByTxnID(ctx context.Context, txnID string) (*db.CardAuthTransaction, bool)
	SetByTxnID(ctx context.Context, txnID string, data *db.CardAuthTransaction)
	DeleteByTxnID(ctx context.Context, txnID string)
}
