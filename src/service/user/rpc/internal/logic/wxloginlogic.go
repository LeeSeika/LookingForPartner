package logic

import (
	"context"
	"lookingforpartner/common/constant"
	"lookingforpartner/common/errs"
	"lookingforpartner/model"
	"lookingforpartner/service/user/rpc/internal/converter"

	"lookingforpartner/pb/user"
	"lookingforpartner/service/user/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
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
		Logger: logx.WithContext(ctx),
	}
}

func (l *WxLoginLogic) WxLogin(in *user.WxLoginRequest) (*user.WxLoginResponse, error) {
	u := model.User{
		WxUid:    constant.NanoidPrefixUser + in.WxUid,
		Username: in.Username,
	}
	err := l.svcCtx.UserInterface.FirstOrCreateUser(&u)
	if err != nil {
		l.Logger.Errorf("[User][Rpc] FirstOrCreateUser error, err: %+v", err)
		return nil, errs.RpcUnknown
	}

	return &user.WxLoginResponse{UserInfo: converter.UserDBToRpc(&u)}, nil
}
