package logic

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"lookingforpartner/common/errs"

	"lookingforpartner/service/user/rpc/internal/svc"
	"lookingforpartner/service/user/rpc/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *user.GetUserInfoRequest) (*user.GetUserInfoResponse, error) {
	u, err := l.svcCtx.UserInterface.GetUser(in.WxUid)
	if err != nil {
		l.Logger.Errorf("[User][Rpc] GetUser error, err: %+v", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.RpcNotFound
		}
		return nil, errs.RpcUnknown
	}

	userInfo := user.UserInfo{
		PostCount:    u.PostCount,
		School:       u.School,
		Grade:        u.Grade,
		Avatar:       u.Avatar,
		Introduction: u.Introduction,
		Username:     u.Username,
	}

	return &user.GetUserInfoResponse{UserInfo: &userInfo}, nil
}
