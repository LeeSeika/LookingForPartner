package logic

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"lookingforpartner/common/errs"
	"lookingforpartner/service/user/api/internal/svc"
	"lookingforpartner/service/user/api/internal/types"
	"lookingforpartner/service/user/rpc/pb/user"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo(req *types.GetUserInfoRequest) (resp *types.GetUserInfoResponse, err error) {
	// rpc call
	getUserInfoReq := user.GetUserInfoRequest{WxUid: req.ID}
	getUserInfoResp, err := l.svcCtx.UserRpc.GetUserInfo(l.ctx, &getUserInfoReq)
	if err != nil {
		l.Logger.Errorf("[User][Api] GetUserInfo error, err: %v", err)
		if errors.Is(err, errs.RpcNotFound) {
			return nil, errs.FormattedNotFound()
		}
		return nil, errs.FormattedUnknown()
	}

	resp = &types.GetUserInfoResponse{
		Avatar:       getUserInfoResp.User.Avatar,
		School:       getUserInfoResp.User.School,
		Grade:        getUserInfoResp.User.Grade,
		Introduction: getUserInfoResp.User.Introduction,
		PostCount:    getUserInfoResp.User.PostCount,
	}

	return resp, nil
}
