package logic

import (
	"context"

	"lookingforpartner/service/user/rpc/internal/svc"
	"lookingforpartner/service/user/rpc/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserBaseInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserBaseInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserBaseInfoLogic {
	return &GetUserBaseInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserBaseInfoLogic) GetUserBaseInfo(in *user.GetUserBaseInfoRequest) (*user.GetUserBaseInfoResponse, error) {
	// todo: add your logic here and delete this line

	return &user.GetUserBaseInfoResponse{}, nil
}
