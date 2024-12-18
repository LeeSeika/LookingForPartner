package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"lookingforpartner/common/logger"
	"lookingforpartner/service/user/api/internal/config"
	"lookingforpartner/service/user/api/internal/handler"
	"lookingforpartner/service/user/api/internal/svc"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	logger.SetupLogger("user-api")

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting user api server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
