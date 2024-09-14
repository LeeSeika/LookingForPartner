package config

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Mysql struct {
		DataBase        string
		Username        string
		Password        string
		Host            string
		Port            string
		MaxIdleConns    int
		MaxOpenConns    int
		ConnMaxLifeTime int
	}
	PostRpc zrpc.RpcClientConf

	KqDeleteCommentsByIDPusherConf struct {
		Brokers []string
		Topic   string
	}
	KqDeleteCommentsByIDConsumerConf kq.KqConf
}
