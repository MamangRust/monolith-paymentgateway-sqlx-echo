package cardapimapper

import (
	pb "github.com/MamangRust/monolith-payment-gateway-pb/card/stats"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

// CardStatsBalanceResponseMapper is an interface for mapping CardStatisticBalanceResponse from the domain to the API response representation.
type cardStatsBalanceResponseMapper struct {
}

// NewCardStatsBalanceResponseMapper creates a new instance of CardStatisticBalanceResponseMapper.
//
// It returns a pointer to a cardStatsBalanceResponseMapper which implements the
// CardStatisticBalanceResponseMapper interface for mapping database rows to domain
// models.
func NewCardStatsBalanceResponseMapper() CardStatsBalanceResponseMapper {
	return &cardStatsBalanceResponseMapper{}
}

// ToApiResponseMonthlyBalances maps the ApiResponseMonthlyBalance from the domain to the ApiResponseMonthlyBalance of the api.
//
// Args:
//   - cards: A pointer to a pb.ApiResponseMonthlyBalance representing the ApiResponseMonthlyBalance from the domain.
//
// Returns:
//   - A pointer to a response.ApiResponseMonthlyBalance containing the mapped data, including status, message, and data.
func (s *cardStatsBalanceResponseMapper) ToApiResponseMonthlyBalances(cards *pb.ApiResponseMonthlyBalance) *response.ApiResponseMonthlyBalance {
	return &response.ApiResponseMonthlyBalance{
		Status:  cards.Status,
		Message: cards.Message,
		Data:    s.mapMonthlyBalances(cards.Data),
	}
}

// ToApiResponseYearlyBalances maps the ApiResponseYearlyBalance from the domain to the ApiResponseYearlyBalance of the api.
//
// Args:
//   - cards: A pointer to a pb.ApiResponseYearlyBalance representing the ApiResponseYearlyBalance from the domain.
//
// Returns:
//   - A pointer to a response.ApiResponseYearlyBalance containing the mapped data, including status, message, and data.
func (s *cardStatsBalanceResponseMapper) ToApiResponseYearlyBalances(cards *pb.ApiResponseYearlyBalance) *response.ApiResponseYearlyBalance {
	return &response.ApiResponseYearlyBalance{
		Status:  cards.Status,
		Message: cards.Message,
		Data:    s.mapYearlyBalances(cards.Data),
	}
}

// mapMonthlyBalance maps a CardResponseMonthlyBalance from the domain representation to the API response representation.
//
// Args:
//   - cards: A pointer to a pb.CardResponseMonthlyBalance representing the domain CardResponseMonthlyBalance object.
//
// Returns:
//   - A pointer to a response.CardResponseMonthBalance containing the mapped data, including month and total balance.
func (s *cardStatsBalanceResponseMapper) mapMonthlyBalance(cards *pb.CardResponseMonthlyBalance) *response.CardResponseMonthBalance {
	return &response.CardResponseMonthBalance{
		Month:        cards.Month,
		TotalBalance: cards.TotalBalance,
	}
}

// mapMonthlyBalances maps a slice of CardResponseMonthlyBalance from the domain representation to the API response representation.
//
// Args:
//   - cards: A slice of pointers to pb.CardResponseMonthlyBalance representing the domain CardResponseMonthlyBalance objects.
//
// Returns:
//   - A slice of pointers to response.CardResponseMonthBalance containing the mapped data, including month and total balance.
func (s *cardStatsBalanceResponseMapper) mapMonthlyBalances(cards []*pb.CardResponseMonthlyBalance) []*response.CardResponseMonthBalance {
	var responseCards []*response.CardResponseMonthBalance

	for _, role := range cards {
		responseCards = append(responseCards, s.mapMonthlyBalance(role))
	}

	return responseCards
}

// mapYearlyBalance maps a CardResponseYearlyBalance from the domain representation to the API response representation.
//
// Args:
//   - cards: A pointer to a pb.CardResponseYearlyBalance representing the domain CardResponseYearlyBalance object.
//
// Returns:
//   - A pointer to a response.CardResponseYearlyBalance containing the mapped data, including year and total balance.
func (s *cardStatsBalanceResponseMapper) mapYearlyBalance(cards *pb.CardResponseYearlyBalance) *response.CardResponseYearlyBalance {
	return &response.CardResponseYearlyBalance{
		Year:         cards.Year,
		TotalBalance: cards.TotalBalance,
	}
}

// mapYearlyBalances maps a slice of CardResponseYearlyBalance from the domain representation to the API response representation.
//
// Args:
//   - cards: A slice of pointers to pb.CardResponseYearlyBalance representing the domain CardResponseYearlyBalance objects.
//
// Returns:
//   - A slice of pointers to response.CardResponseYearlyBalance containing the mapped data, including year and total balance.
func (s *cardStatsBalanceResponseMapper) mapYearlyBalances(cards []*pb.CardResponseYearlyBalance) []*response.CardResponseYearlyBalance {
	var responseCards []*response.CardResponseYearlyBalance

	for _, role := range cards {
		responseCards = append(responseCards, s.mapYearlyBalance(role))
	}

	return responseCards
}
