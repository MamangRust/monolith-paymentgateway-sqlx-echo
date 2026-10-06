package cardapimapper

import (
	pb "github.com/MamangRust/monolith-payment-gateway-pb/card"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

type cardCommandResponseMapper struct {
}

func NewCardCommandResponseMapper() CardCommandResponseMapper {
	return &cardCommandResponseMapper{}
}

// ToApiResponseCard maps the ApiResponseCard from the domain to the ApiResponseCard of the api.
//
// Args:
//   - card: A pointer to a pb.ApiResponseCard representing the ApiResponseCard from the domain.
//
// Returns:
//   - A pointer to a response.ApiResponseCard containing the mapped data, including status, message,
//     and data.
func (s *cardCommandResponseMapper) ToApiResponseCard(card *pb.ApiResponseCard) *response.ApiResponseCard {
	return &response.ApiResponseCard{
		Status:  card.Status,
		Message: card.Message,
		Data:    s.mapCardResponse(card.Data),
	}
}

func (s *cardCommandResponseMapper) ToApiResponseCardDeleteAt(card *pb.ApiResponseCardDeleteAt) *response.ApiResponseCardDeleteAt {
	return &response.ApiResponseCardDeleteAt{
		Status:  card.Status,
		Message: card.Message,
		Data:    s.mapCardResponseDeleteAt(card.Data),
	}
}

// ToApiResponseCardDelete maps the ApiResponseCardDelete from the domain to the ApiResponseCardDelete of the api.
//
// Args:
//   - card: A pointer to a pb.ApiResponseCardDelete representing the ApiResponseCardDelete from the domain.
//
// Returns:
//   - A pointer to a response.ApiResponseCardDelete containing the mapped data, including status and message.
func (s *cardCommandResponseMapper) ToApiResponseCardDelete(card *pb.ApiResponseCardDelete) *response.ApiResponseCardDelete {
	return &response.ApiResponseCardDelete{
		Status:  card.Status,
		Message: card.Message,
	}
}

// ToApiResponseCardAll maps the ApiResponseCardAll from the domain to the ApiResponseCardAll of the api.
//
// Args:
//   - card: A pointer to a pb.ApiResponseCardAll representing the ApiResponseCardAll from the domain.
//
// Returns:
//   - A pointer to a response.ApiResponseCardAll containing the mapped data, including status and message.
func (s *cardCommandResponseMapper) ToApiResponseCardAll(card *pb.ApiResponseCardAll) *response.ApiResponseCardAll {
	return &response.ApiResponseCardAll{
		Status:  card.Status,
		Message: card.Message,
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
func (s *cardCommandResponseMapper) mapCardResponse(card *pb.CardResponse) *response.CardResponse {
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

// mapCardResponseDeleteAt maps a CardResponseDeleteAt from the domain representation to the API response representation.
//
// Args:
//   - card: A pointer to a pb.CardResponseDeleteAt representing the domain CardResponseDeleteAt object.
//
// Returns:
//   - A pointer to a response.CardResponseDeleteAt containing the mapped data, including ID, UserID,
//     CardNumber, CardType, ExpireDate, CVV, CardProvider, CreatedAt, UpdatedAt, and DeletedAt.
func (s *cardCommandResponseMapper) mapCardResponseDeleteAt(card *pb.CardResponseDeleteAt) *response.CardResponseDeleteAt {
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
