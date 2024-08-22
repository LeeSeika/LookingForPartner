package logic

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"lookingforpartner/common/errs"
	"lookingforpartner/service/user/model"
	"lookingforpartner/service/user/rpc/internal/converter"

	"lookingforpartner/service/user/rpc/internal/svc"
	"lookingforpartner/service/user/rpc/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
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
		Logger: logx.WithContext(ctx),
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
		l.Logger.Errorf("[User][Rpc] SetUser error, err: %+v", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.RpcNotFound
		}
		return nil, errs.RpcUnknown
	}

	return &user.SetUserInfoResponse{User: converter.UserDB2Rpc(&u)}, nil
}
