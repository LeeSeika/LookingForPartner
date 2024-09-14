package main

import (
	"flag"
	"fmt"

	"lookingforpartner/service/post/api/internal/config"
	"lookingforpartner/service/post/api/internal/handler"
	"lookingforpartner/service/post/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/post.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting post api server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
