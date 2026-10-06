package service

import (
	"context"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
)

// TopupQueryService defines the read-only operations for querying topup data.
type TopupQueryService interface {
	FindAll(ctx context.Context, req *requests.FindAllTopups) ([]*db.GetTopupsRow, *int, error)
	FindAllByCardNumber(ctx context.Context, req *requests.FindAllTopupsByCardNumber) ([]*db.GetTopupsByCardNumberRow, *int, error)
	FindById(ctx context.Context, topupID int) (*db.GetTopupByIDRow, error)
	FindByActive(ctx context.Context, req *requests.FindAllTopups) ([]*db.GetActiveTopupsRow, *int, error)
	FindByTrashed(ctx context.Context, req *requests.FindAllTopups) ([]*db.GetTrashedTopupsRow, *int, error)
}

type TopupCommandService interface {
	CreateTopup(ctx context.Context, request *requests.CreateTopupRequest) (*db.UpdateTopupStatusRow, error)
	UpdateTopup(ctx context.Context, request *requests.UpdateTopupRequest) (*db.UpdateTopupStatusRow, error)
	TrashedTopup(ctx context.Context, topup_id int) (*db.Topup, error)
	RestoreTopup(ctx context.Context, topup_id int) (*db.Topup, error)
	DeleteTopupPermanent(ctx context.Context, topup_id int) (bool, error)

	RestoreAllTopup(ctx context.Context) (bool, error)
	DeleteAllTopupPermanent(ctx context.Context) (bool, error)
}
