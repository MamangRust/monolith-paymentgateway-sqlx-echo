package resilience

import (
	"context"
	"time"

	"github.com/MamangRust/monolith-payment-gateway-pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Default guard settings for downstream gRPC dependencies: 5 consecutive
// failures open the breaker for 30 seconds, at most 20 calls in flight, and
// every call gets a 10 second deadline.
const (
	guardFailureThreshold   = 5
	guardBreakerTimeoutSecs = 30
	guardMaxConcurrent      = 20
	guardCallTimeout        = 10 * time.Second
)

// GuardDialOption returns a grpc.DialOption that wraps every unary RPC with a
// DependencyGuard keyed by the given dependency name (circuit breaker +
// bulkhead + per-call timeout), so a failing dependency is refused locally
// instead of piling up requests.
func GuardDialOption(name string, log logger.LoggerInterface) grpc.DialOption {
	guard := NewDependencyGuard(
		name,
		guardFailureThreshold,
		guardBreakerTimeoutSecs,
		guardMaxConcurrent,
		guardCallTimeout,
		log,
	)

	return grpc.WithChainUnaryInterceptor(func(
		ctx context.Context,
		method string,
		req, reply any,
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		err := guard.Call(ctx, func(ctx context.Context) error {
			return invoker(ctx, method, req, reply, cc, opts...)
		})

		switch err {
		case nil:
			return nil
		case ErrDependencyGuardOpen:
			return status.Errorf(codes.Unavailable, "dependency %q is temporarily unavailable (circuit breaker open)", name)
		case ErrDependencyGuardBulkhead:
			return status.Errorf(codes.ResourceExhausted, "dependency %q has too many concurrent requests", name)
		default:
			return err
		}
	})
}
