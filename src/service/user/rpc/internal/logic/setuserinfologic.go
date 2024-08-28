package logic

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"lookingforpartner/common/logger"

	"lookingforpartner/common/errs"
	"lookingforpartner/model"
	"lookingforpartner/pb/user"
	"lookingforpartner/service/user/rpc/internal/converter"
	"lookingforpartner/service/user/rpc/internal/svc"

	"gorm.io/gorm"
)

type SetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetUserInfoLogic {
	return &SetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logger.NewLogger(ctx, "user"),
	}
}

func (l *SetUserInfoLogic) SetUserInfo(in *user.SetUserInfoRequest) (*user.SetUserInfoResponse, error) {
	u := model.User{
		WxUid:        in.WxUid,
		School:       in.School,
		Grade:        in.Grade,
		Introduction: in.Introduction,
	}
	err := l.svcCtx.UserInterface.SetUser(&u)
	if err != nil {
		l.Logger.Errorf("cannot update user, err: %+v", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.RpcNotFound
		}
		return nil, errs.RpcUnknown
	}

	return &user.SetUserInfoResponse{UserInfo: converter.UserDBToRpc(&u)}, nil
}
