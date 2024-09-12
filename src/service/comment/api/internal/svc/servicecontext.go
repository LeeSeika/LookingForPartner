package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"lookingforpartner/service/comment/api/internal/config"
	"lookingforpartner/service/comment/rpc/commentclient"
)

type ServiceContext struct {
	Config     config.Config
	CommentRpc commentclient.Comment
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		CommentRpc: commentclient.NewComment(zrpc.MustNewClient(c.CommentRpc)),
	}
}
