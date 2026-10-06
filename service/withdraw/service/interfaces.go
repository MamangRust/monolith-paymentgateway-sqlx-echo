package service

import (
	"context"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
)

// WithdrawQueryService defines query operations for fetching withdraw data.
type WithdrawQueryService interface {
	FindAll(ctx context.Context, req *requests.FindAllWithdraws) ([]*db.GetWithdrawsRow, *int, error)
	FindAllByCardNumber(ctx context.Context, req *requests.FindAllWithdrawCardNumber) ([]*db.GetWithdrawsByCardNumberRow, *int, error)
	FindById(ctx context.Context, withdrawID int) (*db.GetWithdrawByIDRow, error)
	FindByActive(ctx context.Context, req *requests.FindAllWithdraws) ([]*db.GetActiveWithdrawsRow, *int, error)
	FindByTrashed(ctx context.Context, req *requests.FindAllWithdraws) ([]*db.GetTrashedWithdrawsRow, *int, error)
}

type WithdrawCommandService interface {
	Create(ctx context.Context, request *requests.CreateWithdrawRequest) (*db.UpdateWithdrawStatusRow, error)
	Update(ctx context.Context, request *requests.UpdateWithdrawRequest) (*db.UpdateWithdrawRow, error)
	TrashedWithdraw(ctx context.Context, withdraw_id int) (*db.Withdraw, error)
	RestoreWithdraw(ctx context.Context, withdraw_id int) (*db.Withdraw, error)
	DeleteWithdrawPermanent(ctx context.Context, withdraw_id int) (bool, error)

	RestoreAllWithdraw(ctx context.Context) (bool, error)
	DeleteAllWithdrawPermanent(ctx context.Context) (bool, error)
}
