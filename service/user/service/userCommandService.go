package service

import (
	"context"
	"errors"

	roleadapter "github.com/MamangRust/monolith-payment-gateway-pkg/adapter/role"
	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-pkg/hash"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-payment-gateway-shared/errorhandler"
	user_repo_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/user_errors/repository"
	user_errors "github.com/MamangRust/monolith-payment-gateway-shared/errors/user_errors/service"
	"github.com/MamangRust/monolith-payment-gateway-shared/observability"

	mencache "github.com/MamangRust/monolith-payment-gateway-user/redis"
	"github.com/MamangRust/monolith-payment-gateway-user/repository"
	"go.opentelemetry.io/otel/attribute"
	"go.uber.org/zap"
)

// userCommandDeps defines dependencies for userCommandService.
type UserCommandDeps struct {
	Cache                 mencache.UserCommandCache
	UserQueryRepository   repository.UserQueryRepository
	UserCommandRepository repository.UserCommandRepository
	RoleAdapter           roleadapter.RoleAdapter
	UserRoleAdapter       roleadapter.UserRoleAdapter
	Logger                logger.LoggerInterface
	Hashing               hash.HashPassword
	Observability         observability.TraceLoggerObservability
}

// userCommandService implements user command operations.
type userCommandService struct {
	cache                 mencache.UserCommandCache
	userQueryRepository   repository.UserQueryRepository
	userCommandRepository repository.UserCommandRepository
	roleAdapter           roleadapter.RoleAdapter
	userRoleAdapter       roleadapter.UserRoleAdapter
	logger                logger.LoggerInterface
	hashing               hash.HashPassword
	observability         observability.TraceLoggerObservability
}

// NewUserCommandService creates a new UserCommandService.
func NewUserCommandService(
	params *UserCommandDeps,
) UserCommandService {
	return &userCommandService{
		cache:                 params.Cache,
		userQueryRepository:   params.UserQueryRepository,
		userCommandRepository: params.UserCommandRepository,
		roleAdapter:           params.RoleAdapter,
		userRoleAdapter:       params.UserRoleAdapter,
		logger:                params.Logger,
		hashing:               params.Hashing,
		observability:         params.Observability,
	}
}

func (s *userCommandService) CreateUser(ctx context.Context, request *requests.CreateUserRequest) (*db.CreateUserRow, error) {
	const method = "CreateUser"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.String("email", request.Email))

	defer func() {
		end(status)
	}()

	s.logger.Debug("Creating new user", zap.String("email", request.Email))

	existingUser, err := s.userQueryRepository.FindByEmail(ctx, request.Email)
	if err != nil {
		if errors.Is(err, user_repo_errors.ErrUserNotFound) {
			s.logger.Debug("Email is available, proceeding to create user", zap.String("email", request.Email))
		} else {
			status = "error"
			return errorhandler.HandleError[*db.CreateUserRow](
				s.logger,
				err,
				method,
				span,
				zap.String("email", request.Email),
			)
		}
	} else if existingUser != nil {
		status = "error"
		return errorhandler.HandleError[*db.CreateUserRow](
			s.logger,
			user_errors.ErrUserEmailAlready,
			method,
			span,
			zap.String("email", request.Email),
		)
	}

	hash, err := s.hashing.HashPassword(request.Password)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.CreateUserRow](
			s.logger,
			user_errors.ErrUserPassword,
			method,
			span,
		)
	}

	request.Password = hash

	res, err := s.userCommandRepository.CreateUser(ctx, request)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.CreateUserRow](
			s.logger,
			err,
			method,
			span,
		)
	}

	// Assign the default ROLE_ADMIN to every newly created user via the role
	// service gRPC adapter. The role contract lives only in the role service.
	role, err := s.roleAdapter.FindByName(ctx, "ROLE_ADMIN")
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.CreateUserRow](
			s.logger,
			err,
			method,
			span,
			zap.String("role", "ROLE_ADMIN"),
		)
	}

	if _, err := s.userRoleAdapter.AssignRoleToUser(ctx, &requests.CreateUserRoleRequest{
		UserId: int(res.UserID),
		RoleId: int(role.RoleID),
	}); err != nil {
		status = "error"
		return errorhandler.HandleError[*db.CreateUserRow](
			s.logger,
			err,
			method,
			span,
			zap.Int("user_id", int(res.UserID)),
			zap.Int("role_id", int(role.RoleID)),
		)
	}

	s.cache.DeleteUserListCache(ctx)
	logSuccess("Successfully created new user", zap.String("email", res.Email), zap.Int("user_id", int(res.UserID)))

	return res, nil
}

func (s *userCommandService) UpdateUser(ctx context.Context, request *requests.UpdateUserRequest) (*db.UpdateUserRow, error) {
	const method = "UpdateUser"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("user_id", *request.UserID))

	defer func() {
		end(status)
	}()

	s.logger.Debug("Updating user", zap.Int("user_id", *request.UserID))

	existingUser, err := s.userQueryRepository.FindById(ctx, *request.UserID)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.UpdateUserRow](
			s.logger,
			err,
			method,
			span,

			zap.Int("user_id", *request.UserID),
		)
	}

	if request.Email != "" && request.Email != existingUser.Email {
		duplicateUser, _ := s.userQueryRepository.FindByEmail(ctx, request.Email)
		if duplicateUser != nil {
			status = "error"
			return errorhandler.HandleError[*db.UpdateUserRow](
				s.logger,
				user_errors.ErrUserEmailAlready,
				method,
				span,
				zap.String("email", request.Email),
			)
		}
		existingUser.Email = request.Email
	}

	if request.Password != "" {
		hash, err := s.hashing.HashPassword(request.Password)
		if err != nil {
			status = "error"
			return errorhandler.HandleError[*db.UpdateUserRow](
				s.logger,
				user_errors.ErrUserPassword,
				method,
				span,
			)
		}
		existingUser.Password = hash
	}

	res, err := s.userCommandRepository.UpdateUser(ctx, request)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.UpdateUserRow](
			s.logger,
			err,
			method,
			span,

			zap.Int("user_id", *request.UserID),
		)
	}

	s.cache.DeleteUserCache(ctx, int(res.UserID))
	s.cache.DeleteUserListCache(ctx)
	logSuccess("Successfully updated user", zap.Int("user_id", int(res.UserID)))

	return res, nil
}

func (s *userCommandService) TrashedUser(ctx context.Context, user_id int) (*db.TrashUserRow, error) {
	const method = "TrashedUser"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("user_id", user_id))

	defer func() {
		end(status)
	}()

	s.logger.Debug("Trashing user", zap.Int("user_id", user_id))

	res, err := s.userCommandRepository.TrashedUser(ctx, user_id)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.TrashUserRow](
			s.logger,
			err,
			method,
			span,

			zap.Int("user_id", user_id),
		)
	}

	s.cache.DeleteUserCache(ctx, user_id)
	s.cache.DeleteUserListCache(ctx)
	logSuccess("Successfully trashed user", zap.Int("user_id", user_id))

	return res, nil
}

func (s *userCommandService) RestoreUser(ctx context.Context, user_id int) (*db.RestoreUserRow, error) {
	const method = "RestoreUser"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("user_id", user_id))

	defer func() {
		end(status)
	}()

	s.logger.Debug("Restoring user", zap.Int("user_id", user_id))

	res, err := s.userCommandRepository.RestoreUser(ctx, user_id)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.RestoreUserRow](
			s.logger,
			err,
			method,
			span,

			zap.Int("user_id", user_id),
		)
	}

	s.cache.DeleteUserCache(ctx, user_id)
	s.cache.DeleteUserListCache(ctx)
	logSuccess("Successfully restored user", zap.Int("user_id", user_id))

	return res, nil
}

func (s *userCommandService) DeleteUserPermanent(ctx context.Context, user_id int) (bool, error) {
	const method = "DeleteUserPermanent"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("user_id", user_id))

	defer func() {
		end(status)
	}()

	s.logger.Debug("Deleting user permanently", zap.Int("user_id", user_id))

	_, err := s.userCommandRepository.DeleteUserPermanent(ctx, user_id)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[bool](
			s.logger,
			err,
			method,
			span,

			zap.Int("user_id", user_id),
		)
	}

	s.cache.DeleteUserCache(ctx, user_id)
	s.cache.DeleteUserListCache(ctx)
	logSuccess("Successfully deleted user permanently", zap.Int("user_id", user_id))

	return true, nil
}

func (s *userCommandService) RestoreAllUser(ctx context.Context) (bool, error) {
	const method = "RestoreAllUser"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method)

	defer func() {
		end(status)
	}()

	s.logger.Debug("Restoring all users")

	_, err := s.userCommandRepository.RestoreAllUser(ctx)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[bool](
			s.logger,
			err,
			method,
			span,
		)
	}

	s.cache.DeleteUserListCache(ctx)
	logSuccess("Successfully restored all users")

	return true, nil
}

func (s *userCommandService) DeleteAllUserPermanent(ctx context.Context) (bool, error) {
	const method = "DeleteAllUserPermanent"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method)

	defer func() {
		end(status)
	}()

	s.logger.Debug("Permanently deleting all users")

	_, err := s.userCommandRepository.DeleteAllUserPermanent(ctx)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[bool](
			s.logger,
			err,
			method,
			span,
		)
	}

	s.cache.DeleteUserListCache(ctx)
	logSuccess("Successfully deleted all users permanently")

	return true, nil
}
