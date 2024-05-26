package logic

import (
	"context"

	"lookingforpartner/service/user/rpc/internal/svc"
	"lookingforpartner/service/user/rpc/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserExtraInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserExtraInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserExtraInfoLogic {
	return &GetUserExtraInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserExtraInfoLogic) GetUserExtraInfo(in *user.GetUserExtraInfoRequest) (*user.GetUserExtraInfoRequest, error) {
	// todo: add your logic here and delete this line

	return &user.GetUserExtraInfoRequest{}, nil
}
