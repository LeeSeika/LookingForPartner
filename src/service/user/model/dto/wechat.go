package dto

type WechatLoginResponseBody struct {
	SessionKey string `json:"session_key"`
	Openid     string `json:"openid"`
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errMsg"`
}
