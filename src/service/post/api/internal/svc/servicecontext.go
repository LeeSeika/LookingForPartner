package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"lookingforpartner/service/comment/rpc/commentclient"
	"lookingforpartner/service/post/api/internal/config"
	"lookingforpartner/service/post/rpc/postclient"
	"lookingforpartner/service/user/rpc/userclient"
)

type ServiceContext struct {
	Config     config.Config
	PostRpc    postclient.Post
	CommentRpc commentclient.Comment
	UserRpc    userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		PostRpc:    postclient.NewPost(zrpc.MustNewClient(c.PostRpc)),
		CommentRpc: commentclient.NewComment(zrpc.MustNewClient(c.CommentRpc)),
		UserRpc:    userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
