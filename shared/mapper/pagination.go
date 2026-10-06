package apimapper

import (
	pb "github.com/MamangRust/monolith-payment-gateway-pb/common"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

// mapPaginationMeta maps a gRPC PaginationMeta to an HTTP-compatible API response PaginationMeta.
//
// Args:
//   - s: A pointer to a pb.PaginationMeta containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.PaginationMeta containing the mapped data, including current page, page size,
//     total records, and total pages.
func MapPaginationMeta(s *pb.PaginationMeta) *response.PaginationMeta {
	return &response.PaginationMeta{
		CurrentPage:  int(s.CurrentPage),
		PageSize:     int(s.PageSize),
		TotalRecords: int(s.TotalRecords),
		TotalPages:   int(s.TotalPages),
	}
}
