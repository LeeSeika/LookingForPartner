package svc

import (
	"github.com/gin-gonic/gin"
	"lookingforpartner/common"
	"lookingforpartner/common/rpcclient"
	"lookingforpartner/idl/pb/user"
	"lookingforpartner/service/user/api/internal/config"
)

type ServiceContext struct {
	Config     config.Config
	UserClient user.UserServiceClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		UserClient: user.NewUserServiceClient(rpcclient.MustInitGrpcConn(c.UserRpc.Key)),
	}
}

func GetServiceContext(c *gin.Context) *ServiceContext {
	v, _ := c.Get(common.SvcCtx)
	svcCtx := v.(*ServiceContext)
	return svcCtx
}
