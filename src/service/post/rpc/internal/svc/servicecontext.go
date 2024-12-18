package svc

import (
	"github.com/zeromicro/go-queue/kq"
	"log"
	"lookingforpartner/common/constant"
	"lookingforpartner/common/localqueue"
	"lookingforpartner/service/post/rpc/internal/config"
	"lookingforpartner/service/post/rpc/internal/dao"
	"lookingforpartner/service/post/rpc/internal/dao/mysql"
)

type ServiceContext struct {
	Config             config.Config
	PostInterface      dao.PostInterface
	KqCreatePostPusher *kq.Pusher
	KqDeletePostPusher *kq.Pusher
	LocalQueue         *localqueue.Queue
}

func NewServiceContext(c config.Config) *ServiceContext {
	postInterface, err := mysql.NewMysqlInterface(c.Mysql.DataBase, c.Mysql.Username, c.Mysql.Password, c.Mysql.Host, c.Mysql.Port, c.Mysql.MaxIdleConns, c.Mysql.MaxOpenConns, c.Mysql.ConnMaxLifeTime)
	if err != nil {
		log.Printf("failed to create post interface, err: %v\n", err)
		return nil
	}

	return &ServiceContext{
		Config:             c,
		PostInterface:      postInterface,
		KqCreatePostPusher: kq.NewPusher(c.KqCreatePostPusherConf.Brokers, c.KqCreatePostPusherConf.Topic),
		KqDeletePostPusher: kq.NewPusher(c.KqDeletePostPusherConf.Brokers, c.KqDeletePostPusherConf.Topic),
		LocalQueue:         localqueue.NewQueue(constant.DefaultLocalQueueChanCap, constant.DefaultLocalQueueDataCap),
	}
}
