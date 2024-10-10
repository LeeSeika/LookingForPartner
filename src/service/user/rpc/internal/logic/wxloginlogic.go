package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"io/ioutil"
	"lookingforpartner/service/user/model/dto"
	"lookingforpartner/service/user/model/entity"
	"net/http"

	"lookingforpartner/common/constant"
	"lookingforpartner/common/errs"
	"lookingforpartner/pb/user"
	"lookingforpartner/service/user/rpc/internal/converter"
	"lookingforpartner/service/user/rpc/internal/svc"
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
		Logger: logx.WithContext(ctx),
	}
}

func (l *WxLoginLogic) WxLogin(in *user.WxLoginRequest) (*user.WxLoginResponse, error) {
	// wechat api call
	appID := l.svcCtx.Config.AppID
	appSecret := l.svcCtx.Config.AppSecret

	authUrl := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", appID, appSecret, in.Code)
	authReq, err := http.NewRequest(http.MethodGet, authUrl, nil)
	if err != nil {
		l.Logger.Errorf("cannot new wechat login request, err: %+v", err)
		return nil, errs.FormatRpcUnknownError(err.Error())
	}

	client := &http.Client{}
	authResp, err := client.Do(authReq)
	if err != nil {
		l.Logger.Errorf("cannot send wechat login request, err: %+v", err)
		return nil, errs.FormatRpcUnknownError(err.Error())
	}

	respBodyData, err := ioutil.ReadAll(authResp.Body)
	if err != nil {
		l.Logger.Errorf("cannot read response body from wechat, err: %+v", err)
		return nil, errs.FormatRpcUnknownError(err.Error())
	}

	var rb dto.WechatLoginResponseBody
	err = json.Unmarshal(respBodyData, &rb)
	if err != nil {
		l.Logger.Errorf("cannot unmarshal json when login wechat, err: %+v", err)
		return nil, errs.FormatRpcUnknownError(err.Error())
	}

	if rb.ErrCode != 0 {
		l.Logger.
			WithFields(logx.Field("wechat response error msg", rb.ErrMsg)).
			WithFields(logx.Field("wechat response error code", rb.ErrCode)).
			Errorf("cannot get wechat session")

		return &user.WxLoginResponse{WechatResponseCode: int32(rb.ErrCode)}, nil
	}

	// create user
	u := &entity.User{
		WxUid:    constant.NanoidPrefixUser + rb.Openid,
		Username: in.Username,
	}
	u, err = l.svcCtx.UserInterface.FirstOrCreateUser(l.ctx, u)
	if err != nil {
		l.Logger.Errorf("cannot get or create user, err: %+v", err)
		return nil, errs.FormatRpcUnknownError(err.Error())
	}

	return &user.WxLoginResponse{UserInfo: converter.UserDBToRpc(u)}, nil
}
