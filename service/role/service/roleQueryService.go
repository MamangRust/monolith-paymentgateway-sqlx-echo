package service

import (
	"context"

	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"

	mencache "github.com/MamangRust/monolith-payment-gateway-role/redis"
	"github.com/MamangRust/monolith-payment-gateway-role/repository"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/errorhandler"
	role_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/role_errors/service"
	"github.com/MamangRust/monolith-payment-gateway-shared/observability"
	"go.opentelemetry.io/otel/attribute"
	"go.uber.org/zap"
)

// roleQueryDeps defines dependencies for roleQueryService.
type RoleQueryDeps struct {
	Ctx           context.Context
	Cache         mencache.RoleQueryCache
	Repository    repository.RoleQueryRepository
	Logger        logger.LoggerInterface
	Observability observability.TraceLoggerObservability
}

// roleQueryService implements role query operations.
type roleQueryService struct {
	cache         mencache.RoleQueryCache
	roleQuery     repository.RoleQueryRepository
	logger        logger.LoggerInterface
	observability observability.TraceLoggerObservability
}

// NewRoleQueryService creates a new roleQueryService.
func NewRoleQueryService(params *RoleQueryDeps) *roleQueryService {
	return &roleQueryService{
		cache:         params.Cache,
		roleQuery:     params.Repository,
		logger:        params.Logger,
		observability: params.Observability,
	}
}

func (s *roleQueryService) FindAll(ctx context.Context, req *requests.FindAllRoles) ([]*db.GetRolesRow, *int, error) {
	const method = "FindAll"

	page, pageSize := s.normalizePagination(req.Page, req.PageSize)
	search := req.Search

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("page", page),
		attribute.Int("pageSize", pageSize),
		attribute.String("search", search))

	defer func() {
		end(status)
	}()

	if data, total, found := s.cache.GetCachedRoles(ctx, req); found {
		logSuccess("Successfully retrieved all role records from cache", zap.Int("totalRecords", *total), zap.Int("page", page), zap.Int("pageSize", pageSize))
		return data, total, nil
	}

	res, err := s.roleQuery.FindAllRoles(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandlerErrorPagination[[]*db.GetRolesRow](
			s.logger,
			role_errors.ErrFailedFindAll,
			method,
			span,

			zap.Int("page", req.Page),
			zap.Int("page_size", req.PageSize),
			zap.String("search", req.Search),
		)
	}

	var totalCount int

	if len(res) > 0 {
		totalCount = int(res[0].TotalCount)
	} else {
		totalCount = 0
	}

	s.cache.SetCachedRoles(ctx, req, res, &totalCount)

	logSuccess("Successfully fetched role",
		zap.Int("totalRecords", totalCount),
		zap.Int("page", page),
		zap.Int("pageSize", pageSize))

	return res, &totalCount, nil
}

func (s *roleQueryService) FindById(ctx context.Context, id int) (*db.Role, error) {
	const method = "FindById"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("id", id))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetCachedRoleById(ctx, id); found {
		logSuccess("Successfully retrieved role from cache", zap.Int("id", id))

		return data, nil
	}

	res, err := s.roleQuery.FindById(ctx, id)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.Role](
			s.logger,
			err,
			method,
			span,

			zap.Int("role_id", id),
		)
	}

	s.cache.SetCachedRoleById(ctx, id, res)

	logSuccess("Successfully fetched role", zap.Int("id", id))

	return res, nil
}

func (s *roleQueryService) FindByName(ctx context.Context, name string) (*db.Role, error) {
	const method = "FindByName"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.String("name", name))

	defer func() {
		end(status)
	}()

	res, err := s.roleQuery.FindByName(ctx, name)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.Role](
			s.logger,
			err,
			method,
			span,
			zap.String("name", name),
		)
	}

	logSuccess("Successfully fetched role by name", zap.String("name", name))

	return res, nil
}

func (s *roleQueryService) FindByUserId(ctx context.Context, id int) ([]*db.Role, error) {
	const method = "FindByUserId"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("user_id", id))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetCachedRoleByUserId(ctx, id); found {
		logSuccess("Successfully fetched role by user ID from cache", zap.Int("id", id))
		return data, nil
	}

	res, err := s.roleQuery.FindByUserId(ctx, id)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.Role](
			s.logger,
			err,
			method,
			span,

			zap.Int("user_id", id),
		)
	}

	s.cache.SetCachedRoleByUserId(ctx, id, res)

	logSuccess("Successfully fetched role by user ID", zap.Int("id", id))

	return res, nil
}

func (s *roleQueryService) FindByActiveRole(ctx context.Context, req *requests.FindAllRoles) ([]*db.GetActiveRolesRow, *int, error) {
	const method = "FindByActiveRole"

	page, pageSize := s.normalizePagination(req.Page, req.PageSize)
	search := req.Search

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("page", page),
		attribute.Int("pageSize", pageSize),
		attribute.String("search", search))

	defer func() {
		end(status)
	}()

	res, err := s.roleQuery.FindByActiveRole(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandlerErrorPagination[[]*db.GetActiveRolesRow](
			s.logger,
			role_errors.ErrFailedFindActive,
			method,
			span,

			zap.Int("page", page),
			zap.Int("page_size", pageSize),
			zap.String("search", search),
		)
	}
	var totalCount int

	if len(res) > 0 {
		totalCount = int(res[0].TotalCount)
	} else {
		totalCount = 0
	}

	s.cache.SetCachedRoleActive(ctx, req, res, &totalCount)

	logSuccess("Successfully fetched active role",
		zap.Int("totalRecords", totalCount),
		zap.Int("page", page),
		zap.Int("pageSize", pageSize))

	return res, &totalCount, nil
}

func (s *roleQueryService) FindByTrashedRole(ctx context.Context, req *requests.FindAllRoles) ([]*db.GetTrashedRolesRow, *int, error) {
	const method = "FindByTrashedRole"

	page, pageSize := s.normalizePagination(req.Page, req.PageSize)
	search := req.Search

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("page", page),
		attribute.Int("pageSize", pageSize),
		attribute.String("search", search))

	defer func() {
		end(status)
	}()

	if data, total, found := s.cache.GetCachedRoleTrashed(ctx, req); found {
		logSuccess("Successfully fetched trashed role from cache", zap.Int("totalRecords", *total), zap.Int("page", page), zap.Int("pageSize", pageSize))
		return data, total, nil
	}

	res, err := s.roleQuery.FindByTrashedRole(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandlerErrorPagination[[]*db.GetTrashedRolesRow](
			s.logger,
			role_errors.ErrFailedFindTrashed,
			method,
			span,

			zap.Int("page", page),
			zap.Int("pageSize", pageSize),
			zap.String("search", search),
		)
	}

	var totalCount int

	if len(res) > 0 {
		totalCount = int(res[0].TotalCount)
	} else {
		totalCount = 0
	}

	s.cache.SetCachedRoleTrashed(ctx, req, res, &totalCount)

	logSuccess("Successfully fetched trashed role",
		zap.Int("totalRecords", totalCount),
		zap.Int("page", page),
		zap.Int("pageSize", pageSize))

	return res, &totalCount, nil
}

// normalizePagination validates and normalizes pagination parameters.
// Ensures the page is set to at least 1 and pageSize to a default of 10 if
// they are not positive. Returns the normalized page and pageSize values.
//
// Parameters:
//   - page: The requested page number.
//   - pageSize: The number of items per page.
//
// Returns:
//   - The normalized page and pageSize values.
func (s *roleQueryService) normalizePagination(page, pageSize int) (int, int) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	return page, pageSize
}
