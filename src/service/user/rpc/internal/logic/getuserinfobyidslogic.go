package logic

import (
	"context"
	"lookingforpartner/common/errs"
	"lookingforpartner/service/user/rpc/internal/converter"

	"lookingforpartner/pb/user"
	"lookingforpartner/service/user/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoByIDsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoByIDsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoByIDsLogic {
	return &GetUserInfoByIDsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoByIDsLogic) GetUserInfoByIDs(in *user.GetUserInfoByIDsRequest) (*user.GetUserInfoByIDsResponse, error) {

	users, err := l.svcCtx.UserInterface.GetUsersByIDs(l.ctx, in.WechatIDs)
	if err != nil {
		l.Logger.Errorf("cannot get users by ids, err:%+v", err)
		return nil, errs.FormatRpcUnknownError(err.Error())
	}

	userRpcs := make([]*user.UserInfo, 0, len(users))
	for _, userDB := range users {
		userRpc := converter.UserDBToRpc(userDB)
		userRpcs = append(userRpcs, userRpc)
	}

	return &user.GetUserInfoByIDsResponse{UserInfos: userRpcs}, nil
}
