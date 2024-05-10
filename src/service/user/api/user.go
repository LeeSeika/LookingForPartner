package main

import (
	"context"
	"flag"
	"fmt"
	"go.uber.org/zap"
	"log"
	"lookingforpartner/common/logger"
	"lookingforpartner/common/rpcclient"
	"lookingforpartner/service/user/api/internal/config"
	"lookingforpartner/service/user/api/internal/routes"
	"lookingforpartner/service/user/api/internal/svc"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

func main() {
	var configFilePath string
	workDir, _ := os.Getwd()
	defaultConfigPath := "/service/user/api/etc/user.yaml"
	if strings.HasSuffix(workDir, "/") {
		defaultConfigPath = "etc/user.yaml"
	}
	flag.StringVar(&configFilePath, "conf", workDir+defaultConfigPath, "程序读取的配置文件全路径")

	var c config.Config
	c = config.MustLoad(configFilePath)

	logger.MustInit(c.Log.FileName, c.Log.Level, c.Log.Mode, c.Log.MaxSize, c.Log.MaxBackups, c.Log.MaxAge)
	defer zap.L().Sync()

	rpcclient.InitResolverBuilder(c.Etcd.Address)

	svcCtx := svc.NewServiceContext(c)
	engine := routes.SetupRouter(&c, svcCtx)

	// 开启 HTTP 服务
	port := c.Server.Port
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: engine,
	}
	go func() {
		// 开启一个goroutine启动服务
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.L().Fatal("http listen failed", zap.Error(err))
		}
	}()

	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	log.Println("Shutdown Server ...")
	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := server.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server Shutdown: ", zap.Error(err))
	}
}
