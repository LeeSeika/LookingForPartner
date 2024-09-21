package svc

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/zrpc"
	"log"
	"lookingforpartner/common/constant"
	"lookingforpartner/common/localqueue"
	"lookingforpartner/service/comment/rpc/commentclient"
	"lookingforpartner/service/post/rpc/internal/config"
	"lookingforpartner/service/post/rpc/internal/dao"
	"lookingforpartner/service/post/rpc/internal/dao/mysql"
	"lookingforpartner/service/user/rpc/userclient"
)

type ServiceContext struct {
	Config                      config.Config
	PostInterface               dao.PostInterface
	UserRpc                     userclient.User
	CommentRpc                  commentclient.Comment
	KqUpdateUserPostCountPusher *kq.Pusher
	KqDeleteSubjectPusher       *kq.Pusher
	LocalQueue                  *localqueue.Queue
}

func NewServiceContext(c config.Config) *ServiceContext {
	postInterface, err := mysql.NewMysqlInterface(c.Mysql.DataBase, c.Mysql.Username, c.Mysql.Password, c.Mysql.Host, c.Mysql.Port, c.Mysql.MaxIdleConns, c.Mysql.MaxOpenConns, c.Mysql.ConnMaxLifeTime)
	if err != nil {
		log.Printf("failed to create post interface, err: %v\n", err)
		return nil
	}

	return &ServiceContext{
		Config:                      c,
		PostInterface:               postInterface,
		UserRpc:                     userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		CommentRpc:                  commentclient.NewComment(zrpc.MustNewClient(c.CommentRpc)),
		KqUpdateUserPostCountPusher: kq.NewPusher(c.KqUpdateUserPostCountPusherConf.Brokers, c.KqUpdateUserPostCountPusherConf.Topic),
		KqDeleteSubjectPusher:       kq.NewPusher(c.KqDeleteSubjectPusherConf.Brokers, c.KqDeleteSubjectPusherConf.Topic),
		LocalQueue:                  localqueue.NewQueue(constant.DefaultLocalQueueChanCap, constant.DefaultLocalQueueDataCap),
	}
}
