package logic

import (
	"context"
	"lookingforpartner/service/user/api/internal/svc"
	"lookingforpartner/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetUserExtraInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetUserExtraInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetUserExtraInfoLogic {
	return &SetUserExtraInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetUserExtraInfoLogic) SetUserExtraInfo(req *types.SetUserExtraInfoRequest) (resp *types.SetUserExtraInfoResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
