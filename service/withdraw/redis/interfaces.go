package mencache

import (
	"context"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
)

type WithdrawQueryCache interface {
	GetCachedWithdrawsCache(ctx context.Context, req *requests.FindAllWithdraws) ([]*db.GetWithdrawsRow, *int, bool)
	SetCachedWithdrawsCache(ctx context.Context, req *requests.FindAllWithdraws, data []*db.GetWithdrawsRow, total *int)

	GetCachedWithdrawByCardCache(ctx context.Context, req *requests.FindAllWithdrawCardNumber) ([]*db.GetWithdrawsByCardNumberRow, *int, bool)
	SetCachedWithdrawByCardCache(ctx context.Context, req *requests.FindAllWithdrawCardNumber, data []*db.GetWithdrawsByCardNumberRow, total *int)

	GetCachedWithdrawActiveCache(ctx context.Context, req *requests.FindAllWithdraws) ([]*db.GetActiveWithdrawsRow, *int, bool)
	SetCachedWithdrawActiveCache(ctx context.Context, req *requests.FindAllWithdraws, data []*db.GetActiveWithdrawsRow, total *int)

	GetCachedWithdrawTrashedCache(ctx context.Context, req *requests.FindAllWithdraws) ([]*db.GetTrashedWithdrawsRow, *int, bool)
	SetCachedWithdrawTrashedCache(ctx context.Context, req *requests.FindAllWithdraws, data []*db.GetTrashedWithdrawsRow, total *int)

	GetCachedWithdrawCache(ctx context.Context, id int) (*db.GetWithdrawByIDRow, bool)
	SetCachedWithdrawCache(ctx context.Context, data *db.GetWithdrawByIDRow)
}

type WithdrawCommandCache interface {
	DeleteCachedWithdrawCache(ctx context.Context, id int)
}
