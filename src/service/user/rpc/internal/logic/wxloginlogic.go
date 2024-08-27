package logic

import (
	"context"

	"lookingforpartner/common/constant"
	"lookingforpartner/common/errs"
	"lookingforpartner/model"
	"lookingforpartner/pb/user"
	"lookingforpartner/service/user/rpc/internal/converter"
	"lookingforpartner/service/user/rpc/internal/svc"

	"github.com/rs/zerolog/log"
)

type WxLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWxLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WxLoginLogic {
	return &WxLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WxLoginLogic) WxLogin(in *user.WxLoginRequest) (*user.WxLoginResponse, error) {
	u := model.User{
		WxUid:    constant.NanoidPrefixUser + in.WxUid,
		Username: in.Username,
	}
	err := l.svcCtx.UserInterface.FirstOrCreateUser(&u)
	if err != nil {
		log.Error().Msgf("cannot get or create user, err: %+v", err)
		return nil, errs.RpcUnknown
	}

	return &user.WxLoginResponse{UserInfo: converter.UserDBToRpc(&u)}, nil
}
