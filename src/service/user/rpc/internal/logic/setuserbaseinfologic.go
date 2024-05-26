package logic

import (
	"context"

	"lookingforpartner/service/user/rpc/internal/svc"
	"lookingforpartner/service/user/rpc/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetUserBaseInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetUserBaseInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetUserBaseInfoLogic {
	return &SetUserBaseInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SetUserBaseInfoLogic) SetUserBaseInfo(in *user.SetUserBaseInfoRequest) (*user.SetUserBaseInfoResponse, error) {
	// todo: add your logic here and delete this line

	return &user.SetUserBaseInfoResponse{}, nil
}
