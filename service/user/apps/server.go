package apps

import (
	"fmt"
	"time"

	pbrole "github.com/MamangRust/monolith-payment-gateway-pb/role"
	pb "github.com/MamangRust/monolith-payment-gateway-pb/user"
	"github.com/MamangRust/monolith-payment-gateway-pkg/adapter"
	"github.com/MamangRust/monolith-payment-gateway-pkg/hash"
	"github.com/MamangRust/monolith-payment-gateway-pkg/resilience"
	"github.com/MamangRust/monolith-payment-gateway-pkg/server"
	"github.com/MamangRust/monolith-payment-gateway-user/handler"
	"github.com/MamangRust/monolith-payment-gateway-user/repository"
	"github.com/MamangRust/monolith-payment-gateway-user/service"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewServer(cfg *server.Config) (*server.GRPCServer, error) {
	srv, err := server.New(cfg)
	if err != nil {
		return nil, err
	}

	// gRPC client to the role service so users can be assigned their default
	// ROLE_ADMIN without the user service owning the roles/user_roles tables.
	connRole, err := grpc.NewClient(
		viper.GetString("GRPC_ROLE_ADDR"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		resilience.GuardDialOption("role", srv.Logger),
	)
	if err != nil {
		return nil, fmt.Errorf("connect to role service: %w", err)
	}
	srv.AddCleanupHook(func() error {
		return connRole.Close()
	})

	roleQueryClient := pbrole.NewRoleServiceClient(connRole)
	roleCommandClient := pbrole.NewRoleCommandServiceClient(connRole)

	guardRole := resilience.NewDependencyGuard("role", 5, 30, 100, 3*time.Second, srv.Logger)
	guardUserRole := resilience.NewDependencyGuard("user_role", 5, 30, 100, 3*time.Second, srv.Logger)

	repos := repository.NewRepositories(
		srv.DB,
		roleQueryClient,
		roleCommandClient,
		repository.GuardOptions{
			Role: []adapter.GuardOption{
				adapter.WithDependencyGuard(guardRole),
			},
			UserRole: []adapter.GuardOption{
				adapter.WithDependencyGuard(guardUserRole),
			},
		},
	)
	svc := service.NewService(&service.Deps{
		Cache:           srv.CacheStore,
		Repositories:    repos,
		RoleAdapter:     repos.RoleAdapter(),
		UserRoleAdapter: repos.UserRoleAdapter(),
		Hash:            hash.NewHashingPassword(),
		Logger:          srv.Logger,
	})
	h := handler.NewHandler(svc)

	srv.RegisterServices = func(gs *grpc.Server) {
		pb.RegisterUserQueryServiceServer(gs, h)
		pb.RegisterUserCommandServiceServer(gs, h)
	}

	return srv, nil
}
