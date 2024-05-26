package logic

import (
	"context"

	"lookingforpartner/service/user/rpc/internal/svc"
	"lookingforpartner/service/user/rpc/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetUserExtraInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetUserExtraInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetUserExtraInfoLogic {
	return &SetUserExtraInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SetUserExtraInfoLogic) SetUserExtraInfo(in *user.SetUserExtraInfoRequest) (*user.SetUserExtraInfoRequest, error) {
	// todo: add your logic here and delete this line

	return &user.SetUserExtraInfoRequest{}, nil
}
