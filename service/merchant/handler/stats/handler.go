package merchantstatshandler

import "github.com/MamangRust/monolith-payment-gateway-merchant/service"

type HandleStats interface {
	MerchantStatsAmountHandleGrpc
	MerchantStatsMethodHandleGrpc
	MerchantStatsTotalAmountHandleGrpc
}

type handlerStats struct {
	MerchantStatsAmountHandleGrpc
	MerchantStatsMethodHandleGrpc
	MerchantStatsTotalAmountHandleGrpc
}

func NewMerchantStatsHandler(service service.Service) HandleStats {
	return &handlerStats{
		MerchantStatsAmountHandleGrpc:      NewMerchantStatsAmountHandler(service.MerchantStatsService(), service.MerchantStatsByMerchantService(), service.MerchantStatsByApiKeyService()),
		MerchantStatsMethodHandleGrpc:      NewMerchantStatsMethodHandler(service.MerchantStatsService(), service.MerchantStatsByMerchantService(), service.MerchantStatsByApiKeyService()),
		MerchantStatsTotalAmountHandleGrpc: NewMerchantStatsTotalAmountHandler(service.MerchantStatsService(), service.MerchantStatsByMerchantService(), service.MerchantStatsByApiKeyService()),
	}
}
