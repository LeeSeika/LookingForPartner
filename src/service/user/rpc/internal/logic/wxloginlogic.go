package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"lookingforpartner/common/logger"

	"lookingforpartner/common/constant"
	"lookingforpartner/common/errs"
	"lookingforpartner/model"
	"lookingforpartner/pb/user"
	"lookingforpartner/service/user/rpc/internal/converter"
	"lookingforpartner/service/user/rpc/internal/svc"
)

type WxLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewWxLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WxLoginLogic {
	return &WxLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logger.NewLogger(ctx, "user"),
	}
}

func (l *WxLoginLogic) WxLogin(in *user.WxLoginRequest) (*user.WxLoginResponse, error) {
	u := model.User{
		WxUid:    constant.NanoidPrefixUser + in.WxUid,
		Username: in.Username,
	}
	err := l.svcCtx.UserInterface.FirstOrCreateUser(&u)
	if err != nil {
		l.Logger.Errorf("cannot get or create user, err: %+v", err)
		return nil, errs.RpcUnknown
	}

	return &user.WxLoginResponse{UserInfo: converter.UserDBToRpc(&u)}, nil
}
