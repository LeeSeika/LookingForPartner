package config

import (
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
	UserRpc                         zrpc.RpcClientConf
	KqUpdateUserPostCountPusherConf struct {
		Brokers []string
		Topic   string
	}
	KqDeleteSubjectPusherConf struct {
		Brokers []string
		Topic   string
	}
}
