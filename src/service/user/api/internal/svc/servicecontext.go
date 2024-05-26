package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"lookingforpartner/service/user/api/internal/config"
	"lookingforpartner/service/user/rpc/userclient"
)

type ServiceContext struct {
	UserRpc userclient.User
	Config  config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		Config:  c,
	}
}
