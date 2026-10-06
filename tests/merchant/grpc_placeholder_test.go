package merchant_test

import (
	"net"

	pbrole "github.com/MamangRust/monolith-payment-gateway-pb/role"
	pbuser "github.com/MamangRust/monolith-payment-gateway-pb/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// newPlaceholderConn returns a lazy gRPC client connection pointed at a local
// listener with no server behind it. DB-only merchant tests use it to satisfy
// constructors that now require raw gRPC clients without starting a real
// dependency service (the corresponding RPCs are never issued).
func newPlaceholderConn() *grpc.ClientConn {
	lis, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		panic(err)
	}

	conn, err := grpc.NewClient(lis.Addr().String(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	return conn
}

func placeholderUserQueryClient() pbuser.UserQueryServiceClient {
	return pbuser.NewUserQueryServiceClient(newPlaceholderConn())
}

func placeholderRoleQueryClient() pbrole.RoleServiceClient {
	return pbrole.NewRoleServiceClient(newPlaceholderConn())
}

func placeholderRoleCommandClient() pbrole.RoleCommandServiceClient {
	return pbrole.NewRoleCommandServiceClient(newPlaceholderConn())
}
