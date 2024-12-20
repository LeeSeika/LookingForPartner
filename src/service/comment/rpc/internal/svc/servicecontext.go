package svc

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/zrpc"
	"log"
	"lookingforpartner/common/constant"
	"lookingforpartner/common/localqueue"
	"lookingforpartner/service/comment/rpc/internal/config"
	"lookingforpartner/service/comment/rpc/internal/dao"
	"lookingforpartner/service/comment/rpc/internal/dao/mysql"
	"lookingforpartner/service/post/rpc/postclient"
)

type ServiceContext struct {
	Config              config.Config
	CommentInterface    dao.CommentInterface
	PostRpc             postclient.Post
	KqDeleteRootComment *kq.Pusher
	LocalQueue          *localqueue.Queue
}

func NewServiceContext(c config.Config) *ServiceContext {
	commentInterface, err := mysql.NewMysqlInterface(c.Mysql.DataBase, c.Mysql.Username, c.Mysql.Password, c.Mysql.Host, c.Mysql.Port, c.Mysql.MaxIdleConns, c.Mysql.MaxOpenConns, c.Mysql.ConnMaxLifeTime)
	if err != nil {
		log.Printf("failed to create post interface, err: %v\n", err)
		return nil
	}
	return &ServiceContext{
		Config:              c,
		CommentInterface:    commentInterface,
		PostRpc:             postclient.NewPost(zrpc.MustNewClient(c.PostRpc)),
		KqDeleteRootComment: kq.NewPusher(c.KqDeleteRootCommentPusherConf.Brokers, c.KqDeleteRootCommentPusherConf.Topic),
		LocalQueue:          localqueue.NewQueue(constant.DefaultLocalQueueChanCap, constant.DefaultLocalQueueDataCap),
	}
}
