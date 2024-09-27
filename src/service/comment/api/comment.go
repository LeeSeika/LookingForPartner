package main

import (
	"flag"
	"fmt"

	"lookingforpartner/service/comment/api/internal/config"
	"lookingforpartner/service/comment/api/internal/handler"
	"lookingforpartner/service/comment/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/comment.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting comment api server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
