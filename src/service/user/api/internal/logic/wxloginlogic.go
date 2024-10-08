package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"lookingforpartner/common/errs"
	"lookingforpartner/common/logger"
	"lookingforpartner/pb/user"
	"lookingforpartner/service/user/api/internal/common"
	"lookingforpartner/service/user/api/internal/svc"
	"lookingforpartner/service/user/api/internal/types"
	"net/http"
)

type WxLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewWxLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WxLoginLogic {
	return &WxLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logger.NewLogger(ctx, "user-api"),
	}
}

func (l *WxLoginLogic) WxLogin(req *types.WxLoginRequest) (resp *types.WxLoginResponse, err error) {

	wxLoginReq := user.WxLoginRequest{
		Username: req.NickName,
		Code:     req.Code,
	}
	wxLoginResp, err := l.svcCtx.UserRpc.WxLogin(l.ctx, &wxLoginReq)
	if err != nil {
		l.Logger.Errorf("cannot login, err: %+v", err)
		if wxLoginResp.WechatResponseCode != 0 {

			if wxLoginResp.WechatResponseCode == errs.WechatLoginInvalidCode {
				return nil, errs.FormatApiError(http.StatusBadRequest, "invalid js_code")
			} else if wxLoginResp.WechatResponseCode == errs.WechatLoginReachedRateLimit {
				return nil, errs.FormatApiError(http.StatusTooManyRequests, "too many login requests")
			} else if wxLoginResp.WechatResponseCode == errs.WechatLoginBlockedUser {
				return nil, errs.FormatApiError(http.StatusForbidden, "this account has been blocked by wechat")
			} else if wxLoginResp.WechatResponseCode == errs.WechatLoginSystemError {
				return nil, errs.FormatApiError(http.StatusServiceUnavailable, "wechat server unavailable")
			}
		}

		return nil, errs.FormattedApiInternal()
	}

	// generate token
	accessExpire := l.svcCtx.Config.Auth.AccessExpire
	refreshExpire := l.svcCtx.Config.Auth.RefreshExpire
	accessSecret := l.svcCtx.Config.Auth.AccessSecret
	accessToken, refreshToken, err := common.CreateTokenAndRefreshToken(wxLoginResp.UserInfo.WxUid, accessExpire, refreshExpire, accessSecret)
	if err != nil {
		l.Logger.Errorf("cannot generate token, err: %+v", err)
		return nil, errs.FormattedApiGenTokenFailed()
	}

	userInfo := types.UserInfo{
		WxUid:        wxLoginResp.UserInfo.WxUid,
		Avatar:       wxLoginResp.UserInfo.Avatar,
		School:       wxLoginResp.UserInfo.School,
		Grade:        wxLoginResp.UserInfo.Grade,
		Introduction: wxLoginResp.UserInfo.Introduction,
		PostCount:    wxLoginResp.UserInfo.PostCount,
		Username:     wxLoginResp.UserInfo.Username,
	}
	resp = &types.WxLoginResponse{
		Token:        accessToken,
		RefreshToken: refreshToken,
		UserInfo:     userInfo,
	}

	return resp, nil
}
