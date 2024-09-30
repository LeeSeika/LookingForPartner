package logic

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"lookingforpartner/service/user/model/entity"

	"lookingforpartner/common/errs"
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
		Logger: logx.WithContext(ctx),
	}
}

func (l *SetUserInfoLogic) SetUserInfo(in *user.SetUserInfoRequest) (*user.SetUserInfoResponse, error) {
	u := &entity.User{
		WxUid:        in.WxUid,
		School:       in.School,
		Grade:        in.Grade,
		Introduction: in.Introduction,
	}
	u, err := l.svcCtx.UserInterface.UpdateUser(l.ctx, u)
	if err != nil {
		l.Logger.Errorf("cannot update user, err: %+v", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.RpcNotFound
		}
		return nil, errs.FormatRpcUnknownError(err.Error())
	}

	return &user.SetUserInfoResponse{UserInfo: converter.UserDBToRpc(u)}, nil
}
