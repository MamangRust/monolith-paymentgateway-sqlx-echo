package cardapimapper

import (
	pb "github.com/MamangRust/monolith-payment-gateway-pb/card"
	pbstats "github.com/MamangRust/monolith-payment-gateway-pb/card/stats"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

type CardBaseResponseMapper interface {
	// ToApiResponseCard maps a single gRPC card response to an HTTP API response.
	ToApiResponseCard(card *pb.ApiResponseCard) *response.ApiResponseCard
}

type CardQueryResponseMapper interface {
	CardBaseResponseMapper

	// ToApiResponsesCard maps a paginated list of gRPC card responses
	// to a paginated HTTP API response.
	ToApiResponsesCard(cards *pb.ApiResponsePaginationCard) *response.ApiResponsePaginationCard

	// ToApiResponsesCardDeletedAt maps a paginated list of soft-deleted gRPC cards
	// to a paginated HTTP API response.
	ToApiResponsesCardDeletedAt(cards *pb.ApiResponsePaginationCardDeleteAt) *response.ApiResponsePaginationCardDeleteAt
}

type CardCommandResponseMapper interface {
	CardBaseResponseMapper

	ToApiResponseCardDeleteAt(card *pb.ApiResponseCardDeleteAt) *response.ApiResponseCardDeleteAt

	ToApiResponseCardDelete(card *pb.ApiResponseCardDelete) *response.ApiResponseCardDelete

	// ToApiResponseCardAll maps a gRPC response containing all cards
	// to an HTTP API response format.
	ToApiResponseCardAll(card *pb.ApiResponseCardAll) *response.ApiResponseCardAll
}

type CardDashboardResponseMapper interface {
	// ToApiResponseDashboardCard maps a gRPC response of general card dashboard statistics
	// to an HTTP API response.
	ToApiResponseDashboardCard(dash *pb.ApiResponseDashboardCard) *response.ApiResponseDashboardCard

	// ToApiResponseDashboardCardCardNumber maps a gRPC response of card-number-specific dashboard statistics
	// to an HTTP API response.
	ToApiResponseDashboardCardCardNumber(dash *pb.ApiResponseDashboardCardNumber) *response.ApiResponseDashboardCardNumber
}

type CardStatsBalanceResponseMapper interface {
	// ToApiResponseMonthlyBalances maps a gRPC response containing monthly balance statistics
	// to an HTTP API response format.
	ToApiResponseMonthlyBalances(cards *pbstats.ApiResponseMonthlyBalance) *response.ApiResponseMonthlyBalance

	// ToApiResponseYearlyBalances maps a gRPC response containing yearly balance statistics
	// to an HTTP API response format.
	ToApiResponseYearlyBalances(cards *pbstats.ApiResponseYearlyBalance) *response.ApiResponseYearlyBalance
}

type CardStatsAmountResponseMapper interface {
	// ToApiResponseMonthlyAmounts maps a gRPC response containing monthly financial amounts
	// (e.g., top-up, withdrawal, transaction) to an HTTP API response format.
	ToApiResponseMonthlyAmounts(cards *pb.ApiResponseMonthlyAmount) *response.ApiResponseMonthlyAmount

	// ToApiResponseYearlyAmounts maps a gRPC response containing yearly financial amounts
	// (e.g., top-up, withdrawal, transaction) to an HTTP API response format.
	ToApiResponseYearlyAmounts(cards *pb.ApiResponseYearlyAmount) *response.ApiResponseYearlyAmount
}
