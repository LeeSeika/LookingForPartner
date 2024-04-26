package main

import (
	"flag"
	"lookingforpartner/service/user/rpc/internal/config"
	"lookingforpartner/service/user/rpc/internal/server"
	"lookingforpartner/service/user/rpc/internal/svc"
	"os"
	"strings"
)

func main() {
	var configFilePath string
	workDir, _ := os.Getwd()
	defaultConfigPath := "/etc/user.yaml"
	if strings.HasSuffix(workDir, "/") {
		defaultConfigPath = "etc/user.yaml"
	}
	flag.StringVar(&configFilePath, "conf", workDir+defaultConfigPath, "程序读取的配置文件全路径")

	var c config.Config
	c = config.MustLoad(configFilePath)

	svcCtx := svc.NewServiceContext(c)
	userServer := server.NewUserServer(svcCtx)

	userServer.MustStart()
	defer userServer.Stop()
}
