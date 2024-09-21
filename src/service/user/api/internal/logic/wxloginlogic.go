package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"io/ioutil"
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
	// wx api call
	appID := l.svcCtx.Config.AppID
	appSecret := l.svcCtx.Config.AppSecret

	authUrl := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", appID, appSecret, req.Code)
	authReq, err := http.NewRequest(http.MethodGet, authUrl, nil)
	if err != nil {
		l.Logger.Errorf("cannot new wechat login request, err: %+v", err)
		return nil, errs.FormattedApiInternal()
	}

	client := &http.Client{}
	authResp, err := client.Do(authReq)
	if err != nil {
		l.Logger.Errorf("cannot send wechat login request, err: %+v", err)
		return nil, errs.FormatApiError(authResp.StatusCode, errs.ApiProcessWxLoginFailed)
	}

	respBodyData, err := ioutil.ReadAll(authResp.Body)
	if err != nil {
		l.Logger.Errorf("cannot read reponse body, err: %+v", err)
		return nil, errs.FormattedApiInternal()
	}

	type respBody struct {
		SessionKey string `json:"session_key"`
		Openid     string `json:"openid"`
		ErrCode    int    `json:"errcode"`
		ErrMsg     string `json:"errMsg"`
	}

	var rb respBody
	err = json.Unmarshal(respBodyData, &rb)
	if err != nil {
		l.Logger.Errorf("cannot unmarshal json, err: %+v", err)
		return nil, errs.FormattedApiInternal()
	}

	// rpc call
	wxLoginReq := user.WxLoginRequest{
		Username: req.NickName,
		WxUid:    rb.Openid,
	}
	wxLoginResp, err := l.svcCtx.UserRpc.WxLogin(l.ctx, &wxLoginReq)
	if err != nil {
		l.Logger.Errorf("cannot call WxLogin rpc, err: %+v", err)
		return nil, errs.FormattedApiInternal()
	}

	// generate token
	accessExpire := l.svcCtx.Config.Auth.AccessExpire
	refreshExpire := l.svcCtx.Config.Auth.RefreshExpire
	accessSecret := l.svcCtx.Config.Auth.AccessSecret
	accessToken, refreshToken, err := common.CreateTokenAndRefreshToken(rb.Openid, accessExpire, refreshExpire, accessSecret)
	if err != nil {
		l.Logger.Errorf("cannot generate token, err: %+v", err)
		return nil, errs.FormattedApiGenTokenFailed()
	}

	userInfo := types.UserInfo{
		WxUid:        rb.Openid,
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
