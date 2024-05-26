package logic

import (
	"context"
	"lookingforpartner/service/user/api/internal/svc"
	"lookingforpartner/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetUserBaseInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetUserBaseInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetUserBaseInfoLogic {
	return &SetUserBaseInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetUserBaseInfoLogic) SetUserBaseInfo(req *types.SetUserBaseInfoRequest) (resp *types.SetUserBaseInfoResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
