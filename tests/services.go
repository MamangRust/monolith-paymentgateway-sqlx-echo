package tests

import (
	roleadapter "github.com/MamangRust/monolith-payment-gateway-pkg/adapter/role"
	useradapter "github.com/MamangRust/monolith-payment-gateway-pkg/adapter/user"
	"github.com/MamangRust/monolith-payment-gateway-pkg/hash"
	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"github.com/MamangRust/monolith-payment-gateway-shared/cache"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	card_handler "github.com/MamangRust/monolith-payment-gateway-card/handler"
	card_repo "github.com/MamangRust/monolith-payment-gateway-card/repository"
	card_service "github.com/MamangRust/monolith-payment-gateway-card/service"
	merchant_handler "github.com/MamangRust/monolith-payment-gateway-merchant/handler"
	merchant_repo "github.com/MamangRust/monolith-payment-gateway-merchant/repository"
	merchant_service "github.com/MamangRust/monolith-payment-gateway-merchant/service"
	pbcard "github.com/MamangRust/monolith-payment-gateway-pb/card"
	cardstats "github.com/MamangRust/monolith-payment-gateway-pb/card/stats"
	pbmerchant "github.com/MamangRust/monolith-payment-gateway-pb/merchant"
	merchantstats "github.com/MamangRust/monolith-payment-gateway-pb/merchant/stats"
	pbdocument "github.com/MamangRust/monolith-payment-gateway-pb/merchant_document"
	pbrole "github.com/MamangRust/monolith-payment-gateway-pb/role"
	pbsaldo "github.com/MamangRust/monolith-payment-gateway-pb/saldo"
	saldostats "github.com/MamangRust/monolith-payment-gateway-pb/saldo/stats"
	pbuser "github.com/MamangRust/monolith-payment-gateway-pb/user"
	role_handler "github.com/MamangRust/monolith-payment-gateway-role/handler"
	role_repo "github.com/MamangRust/monolith-payment-gateway-role/repository"
	role_service "github.com/MamangRust/monolith-payment-gateway-role/service"
	saldo_handler "github.com/MamangRust/monolith-payment-gateway-saldo/handler"
	saldo_repo "github.com/MamangRust/monolith-payment-gateway-saldo/repository"
	saldo_service "github.com/MamangRust/monolith-payment-gateway-saldo/service"
	user_handler "github.com/MamangRust/monolith-payment-gateway-user/handler"
	user_repo "github.com/MamangRust/monolith-payment-gateway-user/repository"
	user_service "github.com/MamangRust/monolith-payment-gateway-user/service"
)

// RoleGRPC bundles an in-process role gRPC service with the role adapters that
// other services depend on. Services such as user and auth resolve ROLE_ADMIN
// and assign it to new users through these adapters.
type RoleGRPC struct {
	Conn            *grpc.ClientConn
	Server          *grpc.Server
	RoleAdapter     roleadapter.RoleAdapter
	UserRoleAdapter roleadapter.UserRoleAdapter
}

// StartRoleGRPC starts an in-process role service backed by db. The caller owns
// the returned resources and must call Close when finished.
func StartRoleGRPC(db *sqlx.DB, cacheStore *cache.CacheStore, log logger.LoggerInterface) (*RoleGRPC, error) {
	svc := role_service.NewService(&role_service.Deps{
		Cache:        cacheStore,
		Repositories: role_repo.NewRepositories(db),
		Logger:       log,
	})
	h := role_handler.NewHandler(svc)

	server := grpc.NewServer()
	pbrole.RegisterRoleServiceServer(server, h.RoleQuery)
	pbrole.RegisterRoleCommandServiceServer(server, h.RoleCommand)

	addr, err := RunGRPCServer(server)
	if err != nil {
		return nil, err
	}

	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		server.Stop()
		return nil, err
	}

	return &RoleGRPC{
		Conn:            conn,
		Server:          server,
		RoleAdapter:     roleadapter.NewRoleAdapter(pbrole.NewRoleServiceClient(conn)),
		UserRoleAdapter: roleadapter.NewUserRoleAdapter(pbrole.NewRoleCommandServiceClient(conn)),
	}, nil
}

func (r *RoleGRPC) Close() {
	if r == nil {
		return
	}
	if r.Conn != nil {
		_ = r.Conn.Close()
	}
	if r.Server != nil {
		r.Server.Stop()
	}
}

// UserGRPC bundles an in-process user gRPC service with a user adapter. The
// service is wired to the role dependency so user creation can assign the
// default ROLE_ADMIN.
type UserGRPC struct {
	Conn    *grpc.ClientConn
	Server  *grpc.Server
	Service user_service.Service
	Adapter useradapter.QueryRepository
}

// StartUserGRPC starts an in-process user service backed by db and role.
func StartUserGRPC(db *sqlx.DB, cacheStore *cache.CacheStore, log logger.LoggerInterface, role *RoleGRPC) (*UserGRPC, error) {
	svc := user_service.NewService(&user_service.Deps{
		Cache: cacheStore,
		Repositories: user_repo.NewRepositories(
			db,
			pbrole.NewRoleServiceClient(role.Conn),
			pbrole.NewRoleCommandServiceClient(role.Conn),
		),
		Hash:            hash.NewHashingPassword(),
		Logger:          log,
		RoleAdapter:     role.RoleAdapter,
		UserRoleAdapter: role.UserRoleAdapter,
	})
	h := user_handler.NewHandler(svc)

	server := grpc.NewServer()
	pbuser.RegisterUserQueryServiceServer(server, h)
	pbuser.RegisterUserCommandServiceServer(server, h)

	addr, err := RunGRPCServer(server)
	if err != nil {
		return nil, err
	}

	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		server.Stop()
		return nil, err
	}

	return &UserGRPC{
		Conn:    conn,
		Server:  server,
		Service: svc,
		Adapter: useradapter.NewQueryAdapter(pbuser.NewUserQueryServiceClient(conn)),
	}, nil
}

func (u *UserGRPC) Close() {
	if u == nil {
		return
	}
	if u.Conn != nil {
		_ = u.Conn.Close()
	}
	if u.Server != nil {
		u.Server.Stop()
	}
}

// IdentityStack is the role + user dependency pair required by services that
// validate users or assign roles. It is the test equivalent of the gRPC
// dependencies wired in each service's apps/server.go.
type IdentityStack struct {
	Role *RoleGRPC
	User *UserGRPC
}

// StartIdentityStack starts a role service and a user service wired to it.
func StartIdentityStack(db *sqlx.DB, cacheStore *cache.CacheStore, log logger.LoggerInterface) (*IdentityStack, error) {
	role, err := StartRoleGRPC(db, cacheStore, log)
	if err != nil {
		return nil, err
	}

	user, err := StartUserGRPC(db, cacheStore, log, role)
	if err != nil {
		role.Close()
		return nil, err
	}

	return &IdentityStack{Role: role, User: user}, nil
}

func (s *IdentityStack) Close() {
	if s == nil {
		return
	}
	s.User.Close()
	s.Role.Close()
}

// CardGRPC bundles an in-process card gRPC service with its repositories. The
// card service validates ownership through the user dependency, so user must be
// started first.
type CardGRPC struct {
	Conn         *grpc.ClientConn
	Server       *grpc.Server
	Repositories *card_repo.Repositories
}

func StartCardGRPC(db *sqlx.DB, cacheStore *cache.CacheStore, log logger.LoggerInterface, user *UserGRPC) (*CardGRPC, error) {
	repos := card_repo.NewRepositories(db, pbuser.NewUserQueryServiceClient(user.Conn))
	svc := card_service.NewService(&card_service.Deps{
		Cache:        cacheStore,
		Repositories: repos,
		Logger:       log,
	})
	h := card_handler.NewHandler(svc)

	server := grpc.NewServer()
	pbcard.RegisterCardQueryServiceServer(server, h)
	pbcard.RegisterCardCommandServiceServer(server, h)
	pbcard.RegisterCardDashboardServiceServer(server, h)
	cardstats.RegisterCardStatsBalanceServiceServer(server, h)
	cardstats.RegisterCardStatsTransferServiceServer(server, h)
	cardstats.RegisterCardStatsWithdrawServiceServer(server, h)
	cardstats.RegisterCardStatsTopupServiceServer(server, h)
	cardstats.RegisterCardStatsTransactionServiceServer(server, h)

	addr, err := RunGRPCServer(server)
	if err != nil {
		return nil, err
	}
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		server.Stop()
		return nil, err
	}

	return &CardGRPC{Conn: conn, Server: server, Repositories: repos}, nil
}

func (c *CardGRPC) Close() {
	if c == nil {
		return
	}
	if c.Conn != nil {
		_ = c.Conn.Close()
	}
	if c.Server != nil {
		c.Server.Stop()
	}
}

// SaldoGRPC bundles an in-process saldo gRPC service. The saldo service reads
// card state through the card dependency, so card must be started first.
type SaldoGRPC struct {
	Conn         *grpc.ClientConn
	Server       *grpc.Server
	Repositories saldo_repo.Repositories
}

func StartSaldoGRPC(db *sqlx.DB, cacheStore *cache.CacheStore, log logger.LoggerInterface, card *CardGRPC) (*SaldoGRPC, error) {
	repos := saldo_repo.NewRepositories(
		db,
		pbcard.NewCardQueryServiceClient(card.Conn),
		pbcard.NewCardCommandServiceClient(card.Conn),
	)
	svc := saldo_service.NewService(&saldo_service.Deps{
		Cache:        cacheStore,
		Repositories: repos,
		Logger:       log,
	})
	h := saldo_handler.NewHandler(svc)

	server := grpc.NewServer()
	pbsaldo.RegisterSaldoQueryServiceServer(server, h)
	pbsaldo.RegisterSaldoCommandServiceServer(server, h)
	saldostats.RegisterSaldoStatsBalanceServiceServer(server, h)

	addr, err := RunGRPCServer(server)
	if err != nil {
		return nil, err
	}
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		server.Stop()
		return nil, err
	}

	return &SaldoGRPC{Conn: conn, Server: server, Repositories: repos}, nil
}

func (s *SaldoGRPC) Close() {
	if s == nil {
		return
	}
	if s.Conn != nil {
		_ = s.Conn.Close()
	}
	if s.Server != nil {
		s.Server.Stop()
	}
}

// MerchantGRPC bundles an in-process merchant gRPC service. Merchant lookups
// resolve users through the user dependency, so user must be started first.
type MerchantGRPC struct {
	Conn         *grpc.ClientConn
	Server       *grpc.Server
	Repositories merchant_repo.Repositories
}

func StartMerchantGRPC(db *sqlx.DB, cacheStore *cache.CacheStore, log logger.LoggerInterface, user *UserGRPC) (*MerchantGRPC, error) {
	repos := merchant_repo.NewRepositories(db, pbuser.NewUserQueryServiceClient(user.Conn))
	svc := merchant_service.NewService(&merchant_service.Deps{
		Cache:        cacheStore,
		Repositories: repos,
		Logger:       log,
	})
	h := merchant_handler.NewHandler(svc)

	server := grpc.NewServer()
	pbmerchant.RegisterMerchantQueryServiceServer(server, h)
	pbmerchant.RegisterMerchantCommandServiceServer(server, h)
	pbmerchant.RegisterMerchantTransactionServiceServer(server, h)
	pbdocument.RegisterMerchantDocumentQueryServiceServer(server, h)
	pbdocument.RegisterMerchantDocumentCommandServiceServer(server, h)
	merchantstats.RegisterMerchantStatsAmountServiceServer(server, h)
	merchantstats.RegisterMerchantStatsMethodServiceServer(server, h)
	merchantstats.RegisterMerchantStatsTotalAmountServiceServer(server, h)

	addr, err := RunGRPCServer(server)
	if err != nil {
		return nil, err
	}
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		server.Stop()
		return nil, err
	}

	return &MerchantGRPC{Conn: conn, Server: server, Repositories: repos}, nil
}

func (m *MerchantGRPC) Close() {
	if m == nil {
		return
	}
	if m.Conn != nil {
		_ = m.Conn.Close()
	}
	if m.Server != nil {
		m.Server.Stop()
	}
}

// Stack is the full in-process dependency graph (role -> user -> card -> saldo,
// and user -> merchant) required by the domain suites. Each service runs on its
// own ephemeral port; callers dial the exposed Conn fields to build clients.
type Stack struct {
	Role     *RoleGRPC
	User     *UserGRPC
	Card     *CardGRPC
	Saldo    *SaldoGRPC
	Merchant *MerchantGRPC
}

func StartStack(db *sqlx.DB, cacheStore *cache.CacheStore, log logger.LoggerInterface) (*Stack, error) {
	role, err := StartRoleGRPC(db, cacheStore, log)
	if err != nil {
		return nil, err
	}
	user, err := StartUserGRPC(db, cacheStore, log, role)
	if err != nil {
		role.Close()
		return nil, err
	}
	card, err := StartCardGRPC(db, cacheStore, log, user)
	if err != nil {
		user.Close()
		role.Close()
		return nil, err
	}
	saldo, err := StartSaldoGRPC(db, cacheStore, log, card)
	if err != nil {
		card.Close()
		user.Close()
		role.Close()
		return nil, err
	}
	merchant, err := StartMerchantGRPC(db, cacheStore, log, user)
	if err != nil {
		saldo.Close()
		card.Close()
		user.Close()
		role.Close()
		return nil, err
	}

	return &Stack{Role: role, User: user, Card: card, Saldo: saldo, Merchant: merchant}, nil
}

func (s *Stack) Close() {
	if s == nil {
		return
	}
	s.Merchant.Close()
	s.Saldo.Close()
	s.Card.Close()
	s.User.Close()
	s.Role.Close()
}

// Client accessors keep test call sites free of pb import noise.
func (s *Stack) CardQueryClient() pbcard.CardQueryServiceClient {
	return pbcard.NewCardQueryServiceClient(s.Card.Conn)
}

func (s *Stack) CardCommandClient() pbcard.CardCommandServiceClient {
	return pbcard.NewCardCommandServiceClient(s.Card.Conn)
}

func (s *Stack) SaldoQueryClient() pbsaldo.SaldoQueryServiceClient {
	return pbsaldo.NewSaldoQueryServiceClient(s.Saldo.Conn)
}

func (s *Stack) SaldoCommandClient() pbsaldo.SaldoCommandServiceClient {
	return pbsaldo.NewSaldoCommandServiceClient(s.Saldo.Conn)
}

func (s *Stack) MerchantQueryClient() pbmerchant.MerchantQueryServiceClient {
	return pbmerchant.NewMerchantQueryServiceClient(s.Merchant.Conn)
}

func (s *Stack) MerchantCommandClient() pbmerchant.MerchantCommandServiceClient {
	return pbmerchant.NewMerchantCommandServiceClient(s.Merchant.Conn)
}

func (s *Stack) UserQueryClient() pbuser.UserQueryServiceClient {
	return pbuser.NewUserQueryServiceClient(s.User.Conn)
}

func (s *Stack) RoleQueryClient() pbrole.RoleServiceClient {
	return pbrole.NewRoleServiceClient(s.Role.Conn)
}

func (s *Stack) RoleCommandClient() pbrole.RoleCommandServiceClient {
	return pbrole.NewRoleCommandServiceClient(s.Role.Conn)
}
