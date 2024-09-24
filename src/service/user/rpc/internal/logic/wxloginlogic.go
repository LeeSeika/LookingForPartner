package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"lookingforpartner/common/logger"
	"lookingforpartner/service/user/model/entity"

	"lookingforpartner/common/constant"
	"lookingforpartner/common/errs"
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
		Logger: logger.NewLogger(ctx, "user-rpc"),
	}
}

func (l *WxLoginLogic) WxLogin(in *user.WxLoginRequest) (*user.WxLoginResponse, error) {
	u := &entity.User{
		WxUid:    constant.NanoidPrefixUser + in.WxUid,
		Username: in.Username,
	}
	u, err := l.svcCtx.UserInterface.FirstOrCreateUser(l.ctx, u)
	if err != nil {
		l.Logger.Errorf("cannot get or create user, err: %+v", err)
		return nil, errs.FormatRpcUnknownError(err.Error())
	}

	return &user.WxLoginResponse{UserInfo: converter.UserDBToRpc(u)}, nil
}
