// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package userclient

import (
	"context"

	"lookingforpartner/pb/user"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	GetUserInfoByIDsRequest     = user.GetUserInfoByIDsRequest
	GetUserInfoByIDsResponse    = user.GetUserInfoByIDsResponse
	GetUserInfoRequest          = user.GetUserInfoRequest
	GetUserInfoResponse         = user.GetUserInfoResponse
	SetUserInfoRequest          = user.SetUserInfoRequest
	SetUserInfoResponse         = user.SetUserInfoResponse
	UpdateUserPostCountRequest  = user.UpdateUserPostCountRequest
	UpdateUserPostCountResponse = user.UpdateUserPostCountResponse
	UserInfo                    = user.UserInfo
	WxLoginRequest              = user.WxLoginRequest
	WxLoginResponse             = user.WxLoginResponse

	User interface {
		WxLogin(ctx context.Context, in *WxLoginRequest, opts ...grpc.CallOption) (*WxLoginResponse, error)
		SetUserInfo(ctx context.Context, in *SetUserInfoRequest, opts ...grpc.CallOption) (*SetUserInfoResponse, error)
		GetUserInfo(ctx context.Context, in *GetUserInfoRequest, opts ...grpc.CallOption) (*GetUserInfoResponse, error)
		GetUserInfoByIDs(ctx context.Context, in *GetUserInfoByIDsRequest, opts ...grpc.CallOption) (*GetUserInfoByIDsResponse, error)
		UpdateUserPostCount(ctx context.Context, in *UpdateUserPostCountRequest, opts ...grpc.CallOption) (*UpdateUserPostCountResponse, error)
	}

	defaultUser struct {
		cli zrpc.Client
	}
)

func NewUser(cli zrpc.Client) User {
	return &defaultUser{
		cli: cli,
	}
}

func (m *defaultUser) WxLogin(ctx context.Context, in *WxLoginRequest, opts ...grpc.CallOption) (*WxLoginResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.WxLogin(ctx, in, opts...)
}

func (m *defaultUser) SetUserInfo(ctx context.Context, in *SetUserInfoRequest, opts ...grpc.CallOption) (*SetUserInfoResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.SetUserInfo(ctx, in, opts...)
}

func (m *defaultUser) GetUserInfo(ctx context.Context, in *GetUserInfoRequest, opts ...grpc.CallOption) (*GetUserInfoResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.GetUserInfo(ctx, in, opts...)
}

func (m *defaultUser) GetUserInfoByIDs(ctx context.Context, in *GetUserInfoByIDsRequest, opts ...grpc.CallOption) (*GetUserInfoByIDsResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.GetUserInfoByIDs(ctx, in, opts...)
}

func (m *defaultUser) UpdateUserPostCount(ctx context.Context, in *UpdateUserPostCountRequest, opts ...grpc.CallOption) (*UpdateUserPostCountResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.UpdateUserPostCount(ctx, in, opts...)
}
