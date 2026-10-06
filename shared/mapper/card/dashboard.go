package cardapimapper

import (
	pb "github.com/MamangRust/monolith-payment-gateway-pb/card"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

// CardDashboardResponseMapper maps a protobuf message of ApiResponseDashboardCard to an API response of ApiResponseDashboardCard and ApiResponseDashboardCardNumber.
type cardDashboardResponseMapper struct {
}

// NewCardDashboardResponseMapper returns a new instance of a CardDashboardResponseMapper.
//
// This function can be used to create a new instance of a CardDashboardResponseMapper
// for mapping a protobuf message of ApiResponseDashboardCard to an API response of
// ApiResponseDashboardCard and ApiResponseDashboardCardNumber.
func NewCardDashboardResponseMapper() CardDashboardResponseMapper {
	return &cardDashboardResponseMapper{}
}

// ToApiResponseDashboardCard maps the ApiResponseDashboardCard from the domain to the ApiResponseDashboardCard of the api.
//
// Args:
//   - dash: A pointer to a pb.ApiResponseDashboardCard representing the ApiResponseDashboardCard from the domain.
//
// Returns:
//   - A pointer to a response.ApiResponseDashboardCard containing the mapped data, including status, message, and data.
func (s *cardDashboardResponseMapper) ToApiResponseDashboardCard(dash *pb.ApiResponseDashboardCard) *response.ApiResponseDashboardCard {
	return &response.ApiResponseDashboardCard{
		Status:  dash.Status,
		Message: dash.Message,
		Data:    s.mapDashboardCard(dash.Data),
	}
}

// ToApiResponseDashboardCardCardNumber maps the ApiResponseDashboardCardNumber from the domain to the ApiResponseDashboardCardNumber of the api.
//
// Args:
//   - dash: A pointer to a pb.ApiResponseDashboardCardNumber representing the ApiResponseDashboardCardNumber from the domain.
//
// Returns:
//   - A pointer to a response.ApiResponseDashboardCardNumber containing the mapped data, including status, message, and data.
func (s *cardDashboardResponseMapper) ToApiResponseDashboardCardCardNumber(dash *pb.ApiResponseDashboardCardNumber) *response.ApiResponseDashboardCardNumber {
	return &response.ApiResponseDashboardCardNumber{
		Status:  dash.Status,
		Message: dash.Message,
		Data:    s.mapDashboardCardCardNumber(dash.Data),
	}
}

// mapDashboardCard maps a CardResponseDashboard from the domain representation to the API response representation.
//
// Args:
//   - dash: A pointer to a pb.CardResponseDashboard representing the domain CardResponseDashboard object.
//
// Returns:
//   - A pointer to a response.DashboardCard containing the mapped data, including total balance,
//     total withdraw, total topup, total transfer, and total transaction.
func (s *cardDashboardResponseMapper) mapDashboardCard(dash *pb.CardResponseDashboard) *response.DashboardCard {
	return &response.DashboardCard{
		TotalBalance:     &dash.TotalBalance,
		TotalWithdraw:    &dash.TotalWithdraw,
		TotalTopup:       &dash.TotalTopup,
		TotalTransfer:    &dash.TotalTransfer,
		TotalTransaction: &dash.TotalTransaction,
	}
}

// mapDashboardCardCardNumber maps a CardResponseDashboardCardNumber from the domain representation to the API response representation.
//
// Args:
//   - dash: A pointer to a pb.CardResponseDashboardCardNumber representing the domain CardResponseDashboardCardNumber object.
//
// Returns:
//   - A pointer to a response.DashboardCardCardNumber containing the mapped data, including total balance,
//     total withdraw, total topup, total transfer send, total transfer receiver, and total transaction.
func (s *cardDashboardResponseMapper) mapDashboardCardCardNumber(dash *pb.CardResponseDashboardCardNumber) *response.DashboardCardCardNumber {
	return &response.DashboardCardCardNumber{
		TotalBalance:          &dash.TotalBalance,
		TotalWithdraw:         &dash.TotalWithdraw,
		TotalTopup:            &dash.TotalTopup,
		TotalTransferSend:     &dash.TotalTransferSend,
		TotalTransferReceiver: &dash.TotalTransferReceiver,
		TotalTransaction:      &dash.TotalTransaction,
	}
}
