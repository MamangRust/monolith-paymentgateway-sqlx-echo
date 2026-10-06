package mencache

import (
	"context"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
)

type CardQueryCache interface {
	GetByIdCache(ctx context.Context, cardID int) (*db.GetCardByIDRow, bool)
	GetByUserIDCache(ctx context.Context, userID int) (*db.GetCardByUserIDRow, bool)
	GetByCardNumberCache(ctx context.Context, cardNumber string) (*db.GetCardByCardNumberRow, bool)
	GetUserCardByCardNumberCache(ctx context.Context, cardNumber string) (*db.GetUserEmailByCardNumberRow, bool)
	GetFindAllCache(ctx context.Context, req *requests.FindAllCards) ([]*db.GetCardsRow, *int, bool)
	GetByActiveCache(ctx context.Context, req *requests.FindAllCards) ([]*db.GetActiveCardsWithCountRow, *int, bool)
	GetByTrashedCache(ctx context.Context, req *requests.FindAllCards) ([]*db.GetTrashedCardsWithCountRow, *int, bool)
	SetByIdCache(ctx context.Context, cardID int, data *db.GetCardByIDRow)
	SetByUserIDCache(ctx context.Context, userID int, data *db.GetCardByUserIDRow)
	SetByCardNumberCache(ctx context.Context, cardNumber string, data *db.GetCardByCardNumberRow)
	SetFindAllCache(ctx context.Context, req *requests.FindAllCards, data []*db.GetCardsRow, totalRecords *int)
	SetByActiveCache(ctx context.Context, req *requests.FindAllCards, data []*db.GetActiveCardsWithCountRow, totalRecords *int)
	SetUserCardByCardNumberCache(ctx context.Context, cardNumber string, data *db.GetUserEmailByCardNumberRow)
	SetByTrashedCache(ctx context.Context, req *requests.FindAllCards, data []*db.GetTrashedCardsWithCountRow, totalRecords *int)
	DeleteByIdCache(ctx context.Context, cardID int)
	DeleteByUserIDCache(ctx context.Context, userID int)
	DeleteByCardNumberCache(ctx context.Context, cardNumber string)
}

type CardCommandCache interface {
	DeleteCardCommandCache(ctx context.Context, id int)
	DeleteCardCache(ctx context.Context, id, userID int, cardNumber string)
}
