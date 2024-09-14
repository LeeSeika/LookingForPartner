package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"lookingforpartner/service/post/api/internal/config"
	"lookingforpartner/service/post/rpc/postclient"
)

type ServiceContext struct {
	Config  config.Config
	PostRpc postclient.Post
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		PostRpc: postclient.NewPost(zrpc.MustNewClient(c.PostRpc)),
	}
}
