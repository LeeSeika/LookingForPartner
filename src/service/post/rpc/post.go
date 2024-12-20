package main

import (
	"flag"
	"fmt"
	"lookingforpartner/common/logger"

	"lookingforpartner/pb/post"
	"lookingforpartner/service/post/rpc/internal/config"
	"lookingforpartner/service/post/rpc/internal/server"
	"lookingforpartner/service/post/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/post.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	logger.SetupLogger("post-rpc")

	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		post.RegisterPostServer(grpcServer, server.NewPostServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting post rpc server at %s...\n", c.ListenOn)
	s.Start()
}
