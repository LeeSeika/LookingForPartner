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

	KqDeleteRootCommentPusherConf struct {
		Brokers []string
		Topic   string
	}
	KqDeleteRootCommentConsumerConf kq.KqConf
	KqDeletePostConsumerConf        kq.KqConf
	KqCreatePostConsumerConf        kq.KqConf
}
