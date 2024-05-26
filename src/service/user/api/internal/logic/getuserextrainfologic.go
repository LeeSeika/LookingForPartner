package logic

import (
	"context"
	"lookingforpartner/service/user/api/internal/svc"
	"lookingforpartner/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserExtraInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserExtraInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserExtraInfoLogic {
	return &GetUserExtraInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserExtraInfoLogic) GetUserExtraInfo(req *types.GetUserExtraInfoRequest) (resp *types.GetUserExtraInfoResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
