package handler

import (
	"context"
	"math"

	pbutils "github.com/MamangRust/monolith-payment-gateway-pb/common"
	pb "github.com/MamangRust/monolith-payment-gateway-pb/role"
	"github.com/MamangRust/monolith-payment-gateway-role/service"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/errors"
	role_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/role_errors/grpc"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type roleQueryHandleGrpc struct {
	pb.UnimplementedRoleServiceServer
	roleQuery service.RoleQueryService
}

func NewRoleQueryHandleGrpc(roleQuery service.RoleQueryService) RoleQueryHandlerGrpc {
	return &roleQueryHandleGrpc{
		roleQuery: roleQuery,
	}
}

func (s *roleQueryHandleGrpc) FindAllRole(ctx context.Context, req *pb.FindAllRoleRequest) (*pb.ApiResponsePaginationRole, error) {
	page := int(req.GetPage())
	pageSize := int(req.GetPageSize())
	search := req.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	reqService := requests.FindAllRoles{
		Page:     page,
		PageSize: pageSize,
		Search:   search,
	}

	roles, totalRecords, err := s.roleQuery.FindAll(ctx, &reqService)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoRoles := make([]*pb.RoleResponse, len(roles))
	for i, role := range roles {
		protoRoles[i] = &pb.RoleResponse{
			Id:        int32(role.RoleID),
			Name:      role.RoleName,
			CreatedAt: role.CreatedAt.Time.Format("2006-01-02"),
			UpdatedAt: role.UpdatedAt.Time.Format("2006-01-02"),
		}
	}

	totalPages := int(math.Ceil(float64(*totalRecords) / float64(pageSize)))

	paginationMeta := &pbutils.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(*totalRecords),
	}

	return &pb.ApiResponsePaginationRole{
		Status:         "success",
		Message:        "Successfully fetched role records",
		Data:           protoRoles,
		PaginationMeta: paginationMeta,
	}, nil
}

func (s *roleQueryHandleGrpc) FindByActive(ctx context.Context, req *pb.FindAllRoleRequest) (*pb.ApiResponsePaginationRoleDeleteAt, error) {
	page := int(req.GetPage())
	pageSize := int(req.GetPageSize())
	search := req.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	reqService := requests.FindAllRoles{
		Page:     page,
		PageSize: pageSize,
		Search:   search,
	}

	roles, totalRecords, err := s.roleQuery.FindByActiveRole(ctx, &reqService)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoRoles := make([]*pb.RoleResponseDeleteAt, len(roles))
	for i, role := range roles {
		protoRoles[i] = &pb.RoleResponseDeleteAt{
			Id:        int32(role.RoleID),
			Name:      role.RoleName,
			CreatedAt: role.CreatedAt.Time.Format("2006-01-02"),
			UpdatedAt: role.UpdatedAt.Time.Format("2006-01-02"),
			DeletedAt: wrapperspb.String(role.DeletedAt.Time.Format("2006-01-02")),
		}
	}

	totalPages := int(math.Ceil(float64(*totalRecords) / float64(pageSize)))

	paginationMeta := &pbutils.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(*totalRecords),
	}

	return &pb.ApiResponsePaginationRoleDeleteAt{
		Status:         "success",
		Message:        "Successfully fetched active roles",
		Data:           protoRoles,
		PaginationMeta: paginationMeta,
	}, nil
}

func (s *roleQueryHandleGrpc) FindByTrashed(ctx context.Context, req *pb.FindAllRoleRequest) (*pb.ApiResponsePaginationRoleDeleteAt, error) {
	page := int(req.GetPage())
	pageSize := int(req.GetPageSize())
	search := req.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	reqService := requests.FindAllRoles{
		Page:     page,
		PageSize: pageSize,
		Search:   search,
	}

	roles, totalRecords, err := s.roleQuery.FindByTrashedRole(ctx, &reqService)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoRoles := make([]*pb.RoleResponseDeleteAt, len(roles))
	for i, role := range roles {
		protoRoles[i] = &pb.RoleResponseDeleteAt{
			Id:        int32(role.RoleID),
			Name:      role.RoleName,
			CreatedAt: role.CreatedAt.Time.Format("2006-01-02"),
			UpdatedAt: role.UpdatedAt.Time.Format("2006-01-02"),
			DeletedAt: wrapperspb.String(role.DeletedAt.Time.Format("2006-01-02")),
		}
	}

	totalPages := int(math.Ceil(float64(*totalRecords) / float64(pageSize)))

	paginationMeta := &pbutils.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(*totalRecords),
	}

	return &pb.ApiResponsePaginationRoleDeleteAt{
		Status:         "success",
		Message:        "Successfully fetched trashed roles",
		Data:           protoRoles,
		PaginationMeta: paginationMeta,
	}, nil
}

func (s *roleQueryHandleGrpc) FindByIdRole(ctx context.Context, req *pb.FindByIdRoleRequest) (*pb.ApiResponseRole, error) {
	roleID := int(req.GetRoleId())

	if roleID <= 0 {
		return nil, role_errors.ErrGrpcRoleInvalidId
	}

	role, err := s.roleQuery.FindById(ctx, roleID)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoRole := &pb.RoleResponse{
		Id:        int32(role.RoleID),
		Name:      role.RoleName,
		CreatedAt: role.CreatedAt.Time.Format("2006-01-02"),
		UpdatedAt: role.UpdatedAt.Time.Format("2006-01-02"),
	}

	return &pb.ApiResponseRole{
		Status:  "success",
		Message: "Successfully fetched role",
		Data:    protoRole,
	}, nil
}

func (s *roleQueryHandleGrpc) FindByNameRole(ctx context.Context, req *pb.FindByNameRoleRequest) (*pb.ApiResponseRole, error) {
	name := req.GetName()

	if name == "" {
		return nil, role_errors.ErrGrpcRoleInvalidName
	}

	role, err := s.roleQuery.FindByName(ctx, name)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoRole := &pb.RoleResponse{
		Id:        int32(role.RoleID),
		Name:      role.RoleName,
		CreatedAt: role.CreatedAt.Time.Format("2006-01-02"),
		UpdatedAt: role.UpdatedAt.Time.Format("2006-01-02"),
	}

	return &pb.ApiResponseRole{
		Status:  "success",
		Message: "Successfully fetched role by name",
		Data:    protoRole,
	}, nil
}

func (s *roleQueryHandleGrpc) FindByUserId(ctx context.Context, req *pb.FindByIdUserRoleRequest) (*pb.ApiResponsesRole, error) {
	userID := int(req.GetUserId())

	if userID <= 0 {
		return nil, role_errors.ErrGrpcRoleInvalidId
	}

	roles, err := s.roleQuery.FindByUserId(ctx, userID)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoRoles := make([]*pb.RoleResponse, len(roles))
	for i, role := range roles {
		protoRoles[i] = &pb.RoleResponse{
			Id:        int32(role.RoleID),
			Name:      role.RoleName,
			CreatedAt: role.CreatedAt.Time.Format("2006-01-02"),
			UpdatedAt: role.UpdatedAt.Time.Format("2006-01-02"),
		}
	}

	return &pb.ApiResponsesRole{
		Status:  "success",
		Message: "Successfully fetched role by user id",
		Data:    protoRoles,
	}, nil
}
