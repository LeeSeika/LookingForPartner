package logic

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"lookingforpartner/common/errs"
	"lookingforpartner/common/logger"
	"lookingforpartner/pb/user"
	"lookingforpartner/service/user/api/internal/svc"
	"lookingforpartner/service/user/api/internal/types"
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

func (l *SetUserInfoLogic) SetUserInfo(req *types.SetUserInfoRequest) (resp *types.SetUserInfoResponse, err error) {
	// validate
	uid, ok := l.ctx.Value("uid").(string)
	if !ok {
		return nil, errs.FormattedApiUnAuthorized()
	}
	if uid != req.WxUid {
		return nil, errs.FormattedApiUnAuthorized()
	}

	// rpc call
	setUserInfoReq := user.SetUserInfoRequest{
		WxUid:        req.WxUid,
		School:       req.School,
		Grade:        req.Grade,
		Introduction: req.Introduction,
	}
	_, err = l.svcCtx.UserRpc.SetUserInfo(l.ctx, &setUserInfoReq)
	if err != nil {
		l.Logger.Errorf("[User][Api] SetUserInfo error, err: %v", err)
		if errors.Is(err, errs.RpcNotFound) {
			return nil, errs.FormattedApiNotFound()
		}
		return nil, errs.FormattedApiInternal()
	}

	getUserInfoReq := user.GetUserInfoRequest{WxUid: req.WxUid}
	getUserInfoResp, err := l.svcCtx.UserRpc.GetUserInfo(l.ctx, &getUserInfoReq)
	if err != nil {
		l.Logger.Errorf("[User][Api] GetUserInfo error, err: %v", err)
		if errors.Is(err, errs.RpcNotFound) {
			return nil, errs.FormattedApiNotFound()
		}
		return nil, errs.FormattedApiInternal()
	}

	resp = &types.SetUserInfoResponse{
		Avatar:       getUserInfoResp.UserInfo.Avatar,
		School:       getUserInfoResp.UserInfo.School,
		Grade:        getUserInfoResp.UserInfo.Grade,
		Introduction: getUserInfoResp.UserInfo.Introduction,
		PostCount:    getUserInfoResp.UserInfo.PostCount,
	}

	return resp, nil
}
