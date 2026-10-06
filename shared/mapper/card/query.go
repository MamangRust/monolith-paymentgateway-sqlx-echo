package cardapimapper

import (
	pb "github.com/MamangRust/monolith-payment-gateway-pb/card"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
	apimapper "github.com/MamangRust/monolith-payment-gateway-shared/mapper"
)

type cardQueryResponseMapper struct {
}

func NewCardQueryResponseMapper() CardQueryResponseMapper {
	return &cardQueryResponseMapper{}
}

// ToApiResponseCard maps the ApiResponseCard from the domain to the ApiResponseCard of the api.
//
// Args:
//   - card: A pointer to a pb.ApiResponseCard representing the ApiResponseCard from the domain.
//
// Returns:
//   - A pointer to a response.ApiResponseCard containing the mapped data, including status, message,
//     and data.
func (s *cardQueryResponseMapper) ToApiResponseCard(card *pb.ApiResponseCard) *response.ApiResponseCard {
	return &response.ApiResponseCard{
		Status:  card.Status,
		Message: card.Message,
		Data:    s.mapCardResponse(card.Data),
	}
}

// ToApiResponsesCard maps the ApiResponsePaginationCard from the domain to the ApiResponsePaginationCard of the api.
//
// Args:
//   - cards: A pointer to a pb.ApiResponsePaginationCard representing the ApiResponsePaginationCard from the domain.
//
// Returns:
//   - A pointer to a response.ApiResponsePaginationCard containing the mapped data, including status, message, data, and pagination details.
func (s *cardQueryResponseMapper) ToApiResponsesCard(cards *pb.ApiResponsePaginationCard) *response.ApiResponsePaginationCard {
	return &response.ApiResponsePaginationCard{
		Status:     cards.Status,
		Message:    cards.Message,
		Data:       s.mapCardResponses(cards.Data),
		Pagination: apimapper.MapPaginationMeta(cards.PaginationMeta),
	}

}

// ToApiResponsesCardDeletedAt maps the ApiResponsePaginationCardDeleteAt from the domain to the
// ApiResponsePaginationCardDeleteAt of the api.
//
// Args:
//   - cards: A pointer to a pb.ApiResponsePaginationCardDeleteAt representing the ApiResponsePaginationCardDeleteAt from the domain.
//
// Returns:
//   - A pointer to a response.ApiResponsePaginationCardDeleteAt containing the mapped data, including status, message, data, and pagination details.
func (s *cardQueryResponseMapper) ToApiResponsesCardDeletedAt(cards *pb.ApiResponsePaginationCardDeleteAt) *response.ApiResponsePaginationCardDeleteAt {
	return &response.ApiResponsePaginationCardDeleteAt{
		Status:     cards.Status,
		Message:    cards.Message,
		Data:       s.mapCardResponsesDeleteAt(cards.Data),
		Pagination: apimapper.MapPaginationMeta(cards.PaginationMeta),
	}
}

// mapCardResponse maps a CardResponse from the domain representation to the API response representation.
//
// Args:
//   - card: A pointer to a pb.CardResponse representing the domain CardResponse object.
//
// Returns:
//   - A pointer to a response.CardResponse containing the mapped data, including ID, UserID,
//     CardNumber, CardType, ExpireDate, CVV, CardProvider, CreatedAt, and UpdatedAt.
func (s *cardQueryResponseMapper) mapCardResponse(card *pb.CardResponse) *response.CardResponse {
	return &response.CardResponse{
		ID:           int(card.Id),
		UserID:       int(card.UserId),
		CardNumber:   card.CardNumber,
		CardType:     card.CardType,
		ExpireDate:   card.ExpireDate,
		CVV:          card.Cvv,
		CardProvider: card.CardProvider,
		CreatedAt:    card.CreatedAt,
		UpdatedAt:    card.UpdatedAt,
	}
}

// mapCardResponses maps a slice of CardResponse from the domain representation to the API response representation.
//
// Args:
//   - cards: A pointer to a slice of pb.CardResponse representing the domain CardResponse objects.
//
// Returns:
//   - A pointer to a slice of response.CardResponse containing the mapped data, including ID, UserID,
//     CardNumber, CardType, ExpireDate, CVV, CardProvider, CreatedAt, and UpdatedAt.
func (s *cardQueryResponseMapper) mapCardResponses(cards []*pb.CardResponse) []*response.CardResponse {
	var responseCards []*response.CardResponse

	for _, role := range cards {
		responseCards = append(responseCards, s.mapCardResponse(role))
	}

	return responseCards
}

// mapCardResponseDeleteAt maps a CardResponseDeleteAt from the domain representation to the API response representation.
//
// Args:
//   - card: A pointer to a pb.CardResponseDeleteAt representing the domain CardResponseDeleteAt object.
//
// Returns:
//   - A pointer to a response.CardResponseDeleteAt containing the mapped data, including ID, UserID,
//     CardNumber, CardType, ExpireDate, CVV, CardProvider, CreatedAt, UpdatedAt, and DeletedAt.
func (s *cardQueryResponseMapper) mapCardResponseDeleteAt(card *pb.CardResponseDeleteAt) *response.CardResponseDeleteAt {
	var deletedAt string
	if card.DeletedAt != nil {
		deletedAt = card.DeletedAt.Value
	}

	return &response.CardResponseDeleteAt{
		ID:           int(card.Id),
		UserID:       int(card.UserId),
		CardNumber:   card.CardNumber,
		CardType:     card.CardType,
		ExpireDate:   card.ExpireDate,
		CVV:          card.Cvv,
		CardProvider: card.CardProvider,
		CreatedAt:    card.CreatedAt,
		UpdatedAt:    card.UpdatedAt,
		DeletedAt:    &deletedAt,
	}
}

// mapCardResponsesDeleteAt maps a slice of CardResponseDeleteAt from the domain representation to the API response representation.
//
// Args:
//   - cards: A pointer to a slice of pb.CardResponseDeleteAt representing the domain CardResponseDeleteAt objects.
//
// Returns:
//   - A pointer to a slice of response.CardResponseDeleteAt containing the mapped data, including ID, UserID,
//     CardNumber, CardType, ExpireDate, CVV, CardProvider, CreatedAt, UpdatedAt, and DeletedAt.
func (s *cardQueryResponseMapper) mapCardResponsesDeleteAt(cards []*pb.CardResponseDeleteAt) []*response.CardResponseDeleteAt {
	var responseCards []*response.CardResponseDeleteAt

	for _, role := range cards {
		responseCards = append(responseCards, s.mapCardResponseDeleteAt(role))
	}

	return responseCards
}
