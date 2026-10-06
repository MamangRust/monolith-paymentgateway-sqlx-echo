package merchantapimapper

import (
	pb "github.com/MamangRust/monolith-payment-gateway-pb/merchant"
	pbstats "github.com/MamangRust/monolith-payment-gateway-pb/merchant/stats"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

type MerchantBaseResponseMapper interface {
	// Maps a single merchant gRPC response to an HTTP API response.
	ToApiResponseMerchant(merchants *pb.ApiResponseMerchant) *response.ApiResponseMerchant
}

type MerchantQueryResponseMapper interface {
	MerchantBaseResponseMapper

	// Maps multiple merchant gRPC responses to a slice-based API response.
	ToApiResponseMerchants(merchants *pb.ApiResponsesMerchant) *response.ApiResponsesMerchant

	// Maps a paginated list of merchants to a paginated HTTP API response.
	ToApiResponsesMerchant(merchants *pb.ApiResponsePaginationMerchant) *response.ApiResponsePaginationMerchant

	// Maps a paginated list of soft-deleted merchants to an HTTP API response.
	ToApiResponsesMerchantDeleteAt(merchants *pb.ApiResponsePaginationMerchantDeleteAt) *response.ApiResponsePaginationMerchantDeleteAt
}

type MerchantCommandResponseMapper interface {
	MerchantBaseResponseMapper

	// Maps a soft-deleted merchant response to an HTTP API format.
	ToApiResponseMerchantDeleteAt(card *pb.ApiResponseMerchantDeleteAt) *response.ApiResponseMerchantDeleteAt

	// Maps a response containing all merchants to an HTTP API format.
	ToApiResponseMerchantAll(card *pb.ApiResponseMerchantAll) *response.ApiResponseMerchantAll

	ToApiResponseMerchantDelete(card *pb.ApiResponseMerchantDelete) *response.ApiResponseMerchantDelete
}

type MerchantTransactionResponseMapper interface {
	// Maps a paginated response of merchant transactions to an HTTP API response.
	ToApiResponseMerchantsTransactionResponse(merchants *pb.ApiResponsePaginationMerchantTransaction) *response.ApiResponsePaginationMerchantTransaction
}

type MerchantStatsMethodResponseMapper interface {
	// Maps monthly payment method statistics of a merchant to an API response.
	ToApiResponseMonthlyPaymentMethods(ms *pbstats.ApiResponseMerchantMonthlyPaymentMethod) *response.ApiResponseMerchantMonthlyPaymentMethod

	// Maps yearly payment method statistics of a merchant to an API response.
	ToApiResponseYearlyPaymentMethods(ms *pbstats.ApiResponseMerchantYearlyPaymentMethod) *response.ApiResponseMerchantYearlyPaymentMethod
}

type MerchantStatsAmountResponseMapper interface {
	// Maps monthly financial amounts (e.g., transactions, top-ups) to an API response.
	ToApiResponseMonthlyAmounts(ms *pbstats.ApiResponseMerchantMonthlyAmount) *response.ApiResponseMerchantMonthlyAmount

	// Maps yearly financial amounts (e.g., transactions, top-ups) to an API response.
	ToApiResponseYearlyAmounts(ms *pbstats.ApiResponseMerchantYearlyAmount) *response.ApiResponseMerchantYearlyAmount
}

type MerchantStatsTotalAmountResponseMapper interface {
	// Maps monthly total financial statistics across merchants to an API response.
	ToApiResponseMonthlyTotalAmounts(ms *pbstats.ApiResponseMerchantMonthlyTotalAmount) *response.ApiResponseMerchantMonthlyTotalAmount

	// Maps yearly total financial statistics across merchants to an API response.
	ToApiResponseYearlyTotalAmounts(ms *pbstats.ApiResponseMerchantYearlyTotalAmount) *response.ApiResponseMerchantYearlyTotalAmount
}
