package logic

import (
	"context"
	"lookingforpartner/service/user/api/internal/svc"
	"lookingforpartner/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserBaseInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserBaseInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserBaseInfoLogic {
	return &GetUserBaseInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserBaseInfoLogic) GetUserBaseInfo(req *types.GetUserBaseInfoRequest) (resp *types.GetUserBaseInfoResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
