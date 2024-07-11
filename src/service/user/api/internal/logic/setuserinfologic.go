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

type SetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetUserInfoLogic {
	return &SetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetUserInfoLogic) SetUserInfo(req *types.SetUserInfoRequest) (resp *types.SetUserInfoResponse, err error) {
	// validate
	uid, ok := l.ctx.Value("uid").(string)
	if !ok {
		return nil, errs.FormattedApiUnAuthorized()
	}
	if uid != req.ID {
		return nil, errs.FormattedApiUnAuthorized()
	}

	// rpc call
	setUserInfoReq := user.SetUserInfoRequest{
		WxUid:        req.ID,
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

	getUserInfoReq := user.GetUserInfoRequest{WxUid: req.ID}
	getUserInfoResp, err := l.svcCtx.UserRpc.GetUserInfo(l.ctx, &getUserInfoReq)
	if err != nil {
		l.Logger.Errorf("[User][Api] GetUserInfo error, err: %v", err)
		if errors.Is(err, errs.RpcNotFound) {
			return nil, errs.FormattedApiNotFound()
		}
		return nil, errs.FormattedApiInternal()
	}

	resp = &types.SetUserInfoResponse{
		Avatar:       getUserInfoResp.User.Avatar,
		School:       getUserInfoResp.User.School,
		Grade:        getUserInfoResp.User.Grade,
		Introduction: getUserInfoResp.User.Introduction,
		PostCount:    getUserInfoResp.User.PostCount,
	}

	return resp, nil
}
