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
			return nil, errs.FormattedApiNotFound()
		}
		return nil, errs.FormattedApiInternal()
	}

	resp = &types.GetUserInfoResponse{
		Avatar:       getUserInfoResp.UserInfo.Avatar,
		School:       getUserInfoResp.UserInfo.School,
		Grade:        getUserInfoResp.UserInfo.Grade,
		Introduction: getUserInfoResp.UserInfo.Introduction,
		PostCount:    getUserInfoResp.UserInfo.PostCount,
	}

	return resp, nil
}
