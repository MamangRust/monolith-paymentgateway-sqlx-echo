package cardapimapper

import (
	pb "github.com/MamangRust/monolith-payment-gateway-pb/card"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

type cardStatsAmountResponseMapper struct {
}

// NewCardStatsAmountResponseMapper creates a new instance of CardStatsAmountResponseMapper.
//
// It returns a pointer to a cardStatsAmountResponseMapper which implements the
// CardStatsAmountResponseMapper interface for mapping database rows to domain
// models.
func NewCardStatsAmountResponseMapper() CardStatsAmountResponseMapper {
	return &cardStatsAmountResponseMapper{}
}

// ToApiResponseMonthlyAmounts maps the ApiResponseMonthlyAmount from the domain to the ApiResponseMonthlyAmount of the api.
//
// Args:
//   - cards: A pointer to a pb.ApiResponseMonthlyAmount representing the ApiResponseMonthlyAmount from the domain.
//
// Returns:
//   - A pointer to a response.ApiResponseMonthlyAmount containing the mapped data, including status, message, and data.
func (s *cardStatsAmountResponseMapper) ToApiResponseMonthlyAmounts(cards *pb.ApiResponseMonthlyAmount) *response.ApiResponseMonthlyAmount {
	return &response.ApiResponseMonthlyAmount{
		Status:  cards.Status,
		Message: cards.Message,
		Data:    s.mapMonthlyAmounts(cards.Data),
	}
}

// ToApiResponseYearlyAmounts maps the ApiResponseYearlyAmount from the domain to the ApiResponseYearlyAmount of the api.
//
// Args:
//   - cards: A pointer to a pb.ApiResponseYearlyAmount representing the ApiResponseYearlyAmount from the domain.
//
// Returns:
//   - A pointer to a response.ApiResponseYearlyAmount containing the mapped data, including status, message, and data.
func (s *cardStatsAmountResponseMapper) ToApiResponseYearlyAmounts(cards *pb.ApiResponseYearlyAmount) *response.ApiResponseYearlyAmount {
	return &response.ApiResponseYearlyAmount{
		Status:  cards.Status,
		Message: cards.Message,
		Data:    s.mapYearlyAmounts(cards.Data),
	}
}

// mapMonthlyAmount maps a CardResponseMonthlyAmount from the domain representation to the API response representation.
//
// Args:
//   - cards: A pointer to a pb.CardResponseMonthlyAmount representing the domain CardResponseMonthlyAmount object.
//
// Returns:
//   - A pointer to a response.CardResponseMonthAmount containing the mapped data, including month and total amount.
func (s *cardStatsAmountResponseMapper) mapMonthlyAmount(cards *pb.CardResponseMonthlyAmount) *response.CardResponseMonthAmount {
	return &response.CardResponseMonthAmount{
		Month:       cards.Month,
		TotalAmount: cards.TotalAmount,
	}
}

// mapMonthlyAmounts maps a slice of CardResponseMonthlyAmount from the domain representation to the API response representation.
//
// Args:
//   - cards: A slice of pointers to pb.CardResponseMonthlyAmount representing the domain CardResponseMonthlyAmount objects.
//
// Returns:
//   - A slice of pointers to response.CardResponseMonthAmount containing the mapped data, including month and total amount.
func (s *cardStatsAmountResponseMapper) mapMonthlyAmounts(cards []*pb.CardResponseMonthlyAmount) []*response.CardResponseMonthAmount {
	var responseCards []*response.CardResponseMonthAmount

	for _, role := range cards {
		responseCards = append(responseCards, s.mapMonthlyAmount(role))
	}

	return responseCards
}

// mapYearlyAmount maps a CardResponseYearlyAmount from the domain representation to the API response representation.
//
// Args:
//   - cards: A pointer to a pb.CardResponseYearlyAmount representing the domain CardResponseYearlyAmount object.
//
// Returns:
//   - A pointer to a response.CardResponseYearAmount containing the mapped data, including year and total amount.
func (s *cardStatsAmountResponseMapper) mapYearlyAmount(cards *pb.CardResponseYearlyAmount) *response.CardResponseYearAmount {
	return &response.CardResponseYearAmount{
		Year:        cards.Year,
		TotalAmount: cards.TotalAmount,
	}
}

// mapYearlyAmounts maps a slice of CardResponseYearlyAmount from the domain representation to the API response representation.
//
// Args:
//   - cards: A slice of pointers to pb.CardResponseYearlyAmount representing the domain CardResponseYearlyAmount objects.
//
// Returns:
//   - A slice of pointers to response.CardResponseYearAmount containing the mapped data, including year and total amount.
func (s *cardStatsAmountResponseMapper) mapYearlyAmounts(cards []*pb.CardResponseYearlyAmount) []*response.CardResponseYearAmount {
	var responseCards []*response.CardResponseYearAmount

	for _, role := range cards {
		responseCards = append(responseCards, s.mapYearlyAmount(role))
	}

	return responseCards
}
