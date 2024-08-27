package logic

import (
	"context"
	"errors"
	"lookingforpartner/common/errs"
	"lookingforpartner/pb/user"
	"lookingforpartner/service/user/api/internal/svc"
	"lookingforpartner/service/user/api/internal/types"

	"github.com/rs/zerolog/log"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{

		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo(req *types.GetUserInfoRequest) (resp *types.GetUserInfoResponse, err error) {
	// rpc call
	getUserInfoReq := user.GetUserInfoRequest{WxUid: req.ID}
	getUserInfoResp, err := l.svcCtx.UserRpc.GetUserInfo(l.ctx, &getUserInfoReq)
	if err != nil {
		log.Error().Msgf("[User][Api] GetUserInfo error, err: %v", err)
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
