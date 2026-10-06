package service

import (
	mencache "github.com/MamangRust/monolith-payment-gateway-auth/redis"
	"github.com/MamangRust/monolith-payment-gateway-auth/repository"
	roleadapter "github.com/MamangRust/monolith-payment-gateway-pkg/adapter/role"
	"github.com/MamangRust/monolith-payment-gateway-pkg/auth"
	"github.com/MamangRust/monolith-payment-gateway-pkg/hash"
	"github.com/MamangRust/monolith-payment-gateway-pkg/kafka"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/cache"
	"github.com/MamangRust/monolith-payment-gateway-shared/observability"
)

// Service aggregates authentication and identity-related services.
type Service struct {
	Login         LoginService
	Register      RegistrationService
	PasswordReset PasswordResetService
	Identify      IdentifyService
}

// Deps defines dependencies required to initialize Service.
type Deps struct {
	Cache           *cache.CacheStore
	Repositories    *repository.Repositories
	RoleAdapter     roleadapter.RoleAdapter
	UserRoleAdapter roleadapter.UserRoleAdapter
	Token           auth.TokenManager
	Hash            hash.HashPassword
	Logger          logger.LoggerInterface
	// Kafka is retained for backwards-compatible test/application wiring.
	// Domain email delivery is persisted through the SQLC outbox.
	Kafka *kafka.Kafka
}

// NewService initializes and returns the core authentication service bundle.
func NewService(deps *Deps) *Service {
	observability, _ := observability.NewObservability("auth-server", deps.Logger)

	cache := mencache.NewMencache(deps.Cache)

	tokenService := NewTokenService(&tokenServiceDeps{
		Token:         deps.Token,
		RefreshToken:  deps.Repositories.RefreshToken,
		Logger:        deps.Logger,
		Observability: observability,
	})

	return &Service{
		Login:         newLogin(deps, tokenService, observability, cache.LoginCache),
		Register:      newRegister(deps, observability, cache.RegisterCache),
		PasswordReset: newPasswordReset(deps, observability, cache.PasswordResetCache),
		Identify:      newIdentity(deps, tokenService, observability, cache.IdentityCache),
	}
}

// newLogin initializes and returns the LoginService.
func newLogin(deps *Deps, tokenService *tokenService, observability observability.TraceLoggerObservability, cache mencache.LoginCache) LoginService {
	return NewLoginService(&LoginServiceDeps{
		Cache:          cache,
		Logger:         deps.Logger,
		Hash:           deps.Hash,
		UserRepository: deps.Repositories.User,
		RefreshToken:   deps.Repositories.RefreshToken,
		Token:          deps.Token,
		TokenService:   tokenService,
		Observability:  observability,
	})
}

// newRegister initializes and returns the RegistrationService.
func newRegister(deps *Deps, observability observability.TraceLoggerObservability, cache mencache.RegisterCache) RegistrationService {
	return NewRegisterService(&RegisterServiceDeps{
		Cache:           cache,
		User:            deps.Repositories.User,
		RoleAdapter:     deps.RoleAdapter,
		UserRoleAdapter: deps.UserRoleAdapter,
		Hash:            deps.Hash,
		Logger:          deps.Logger,
		Observability:   observability,
	})
}

// newPasswordReset initializes the reset forgot password, reset password and verify code services.
func newPasswordReset(deps *Deps, observability observability.TraceLoggerObservability, cache mencache.PasswordResetCache) PasswordResetService {
	return NewPasswordResetService(&PasswordResetServiceDeps{
		Cache:         cache,
		Logger:        deps.Logger,
		User:          deps.Repositories.User,
		ResetToken:    deps.Repositories.ResetToken,
		Observability: observability,
	})
}

// newIdentity initializes the IdentifyService for identity verification and token refresh.
func newIdentity(deps *Deps, tokenService *tokenService, observability observability.TraceLoggerObservability, cache mencache.IdentityCache) IdentifyService {
	return NewIdentityService(&IdentityServiceDeps{
		Cache:         cache,
		Token:         deps.Token,
		RefreshToken:  deps.Repositories.RefreshToken,
		User:          deps.Repositories.User,
		Logger:        deps.Logger,
		TokenService:  tokenService,
		Observability: observability,
	})
}
