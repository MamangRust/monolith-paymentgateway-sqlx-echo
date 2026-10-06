package service

import (
	"context"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
)

type TransferQueryService interface {
	FindAll(ctx context.Context, req *requests.FindAllTransfers) ([]*db.GetTransfersRow, *int, error)
	FindById(ctx context.Context, transferId int) (*db.GetTransferByIDRow, error)
	FindByActive(ctx context.Context, req *requests.FindAllTransfers) ([]*db.GetActiveTransfersRow, *int, error)
	FindByTrashed(ctx context.Context, req *requests.FindAllTransfers) ([]*db.GetTrashedTransfersRow, *int, error)
	FindTransferByTransferFrom(ctx context.Context, transfer_from string) ([]*db.GetTransfersBySourceCardRow, error)
	FindTransferByTransferTo(ctx context.Context, transfer_to string) ([]*db.GetTransfersByDestinationCardRow, error)
}

type TransferCommandService interface {
	CreateTransaction(ctx context.Context, request *requests.CreateTransferRequest) (*db.UpdateTransferStatusRow, error)
	UpdateTransaction(ctx context.Context, request *requests.UpdateTransferRequest) (*db.UpdateTransferRow, error)
	TrashedTransfer(ctx context.Context, transfer_id int) (*db.Transfer, error)
	RestoreTransfer(ctx context.Context, transfer_id int) (*db.Transfer, error)
	DeleteTransferPermanent(ctx context.Context, transfer_id int) (bool, error)

	RestoreAllTransfer(ctx context.Context) (bool, error)
	DeleteAllTransferPermanent(ctx context.Context) (bool, error)
}
