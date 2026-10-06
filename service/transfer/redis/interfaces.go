package mencache

import (
	"context"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
)

type TransferQueryCache interface {
	GetCachedTransfersCache(ctx context.Context, req *requests.FindAllTransfers) ([]*db.GetTransfersRow, *int, bool)
	SetCachedTransfersCache(ctx context.Context, req *requests.FindAllTransfers, data []*db.GetTransfersRow, total *int)

	GetCachedTransferActiveCache(ctx context.Context, req *requests.FindAllTransfers) ([]*db.GetActiveTransfersRow, *int, bool)
	SetCachedTransferActiveCache(ctx context.Context, req *requests.FindAllTransfers, data []*db.GetActiveTransfersRow, total *int)

	GetCachedTransferTrashedCache(ctx context.Context, req *requests.FindAllTransfers) ([]*db.GetTrashedTransfersRow, *int, bool)
	SetCachedTransferTrashedCache(ctx context.Context, req *requests.FindAllTransfers, data []*db.GetTrashedTransfersRow, total *int)

	GetCachedTransferCache(ctx context.Context, id int) (*db.GetTransferByIDRow, bool)
	SetCachedTransferCache(ctx context.Context, data *db.GetTransferByIDRow)

	GetCachedTransferByFrom(ctx context.Context, from string) ([]*db.GetTransfersBySourceCardRow, bool)
	SetCachedTransferByFrom(ctx context.Context, from string, data []*db.GetTransfersBySourceCardRow)

	GetCachedTransferByTo(ctx context.Context, to string) ([]*db.GetTransfersByDestinationCardRow, bool)
	SetCachedTransferByTo(ctx context.Context, to string, data []*db.GetTransfersByDestinationCardRow)
}

type TransferCommandCache interface {
	DeleteTransferCache(ctx context.Context, id int)
}
