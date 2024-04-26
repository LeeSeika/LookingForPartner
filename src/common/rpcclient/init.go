package rpcclient

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
	"lookingforpartner/pkg/discovery"
	"time"
)

const dialTimeout = 10

var resolverBuilder *discovery.Resolver

func InitResolverBuilder(etcdAddrs []string) {
	resolverBuilder = discovery.NewResolver(etcdAddrs)
}

func MustInitGrpcConn(serviceKey string) *grpc.ClientConn {

	resolver.Register(resolverBuilder)
	// defer resolverBuilder.Close()

	zap.L().Info("Begin initializing Service RPC Connection", zap.String("service", serviceKey))

	conn, err := connectRpcServer(serviceKey)
	if err != nil {
		panic(fmt.Sprintf("Connect to %s rpc server failed, err:%v", serviceKey, err))
	}
	return conn
}

func connectRpcServer(serviceKey string) (*grpc.ClientConn, error) {
	dialCtx, cancel := context.WithTimeout(context.Background(), dialTimeout*time.Second)
	defer cancel()

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	addr := fmt.Sprintf("%s:///%s", resolverBuilder.Scheme(), serviceKey)

	// Load balance
	opts = append(opts, grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, "round_robin")))

	conn, err := grpc.DialContext(dialCtx, addr, opts...)
	return conn, err
}
