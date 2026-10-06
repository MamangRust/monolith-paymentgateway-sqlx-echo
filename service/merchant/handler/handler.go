package handler

import (
	merchantstatshandler "github.com/MamangRust/monolith-payment-gateway-merchant/handler/stats"
	"github.com/MamangRust/monolith-payment-gateway-merchant/service"
)

type Handler interface {
	MerchantQueryHandleGrpc
	MerchantCommandHandleGrpc
	MerchantDocumentQueryHandleGrpc
	MerchantDocumentCommandHandleGrpc
	MerchantTransactionHandleGrpc
	merchantstatshandler.HandleStats
}

// Handler contains the gRPC handlers for merchant and merchant document operations.
type handler struct {
	MerchantQueryHandleGrpc
	MerchantCommandHandleGrpc
	MerchantDocumentQueryHandleGrpc
	MerchantDocumentCommandHandleGrpc
	MerchantTransactionHandleGrpc
	merchantstatshandler.HandleStats
}

func NewHandler(service service.Service) Handler {
	return &handler{
		MerchantQueryHandleGrpc:           NewMerchantQueryHandleGrpc(service.MerchantQueryService()),
		MerchantCommandHandleGrpc:         NewMerchantCommandHandleGrpc(service.MerchantCommandService()),
		MerchantDocumentQueryHandleGrpc:   NewMerchantDocumentQueryHandleGrpc(service.MerchantDocumentQueryService()),
		MerchantDocumentCommandHandleGrpc: NewMerchantDocumentCommandHandleGrpc(service.MerchantDocumentCommandService()),
		MerchantTransactionHandleGrpc: NewMerchantTransactionHandleGrpc(
			service.MerchantTransactionService()),
		HandleStats: merchantstatshandler.NewMerchantStatsHandler(service),
	}
}
