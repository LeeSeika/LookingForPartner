package mqs

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/service"
	"lookingforpartner/service/comment/rpc/internal/config"
	"lookingforpartner/service/comment/rpc/internal/svc"
)

func Consumers(c config.Config, svcContext *svc.ServiceContext) []service.Service {

	return []service.Service{
		//Listening for changes in consumption flow status
		kq.MustNewQueue(c.KqDeleteCommentsByIDConsumerConf, NewDeleteCommentsByID(svcContext)),
	}

}
