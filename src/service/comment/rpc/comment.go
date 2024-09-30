package main

import (
	"flag"
	"fmt"
	"log"
	"lookingforpartner/common/logger"
	"lookingforpartner/service/comment/rpc/internal/mqs"

	"lookingforpartner/pb/comment"
	"lookingforpartner/service/comment/rpc/internal/config"
	"lookingforpartner/service/comment/rpc/internal/server"
	"lookingforpartner/service/comment/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/comment.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	logger.SetupLogger("comment-rpc")

	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		comment.RegisterCommentServer(grpcServer, server.NewCommentServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	// mq
	go func() {
		defer func() {
			if p := recover(); p != nil {
				log.Printf("recovered from mq panic:%+v\n", p)
			}
		}()

		serviceGroup := service.NewServiceGroup()
		defer serviceGroup.Stop()

		for _, mq := range mqs.Consumers(c, ctx) {
			serviceGroup.Add(mq)
		}
		serviceGroup.Start()

	}()

	fmt.Printf("Starting comment rpc server at %s...\n", c.ListenOn)
	s.Start()
}
