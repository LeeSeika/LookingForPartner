// Code generated by goctl. DO NOT EDIT.
package types

type GetUserInfoRequest struct {
	ID string `path:"wxUid"`
}

type GetUserInfoResponse struct {
	Avatar       string `json:"avatar"`
	School       string `json:"school"`
	Grade        int64  `json:"grade"`
	Introduction string `json:"introduction"`
	PostCount    int64  `json:"post_count"`
}

type RefreshTokenReqeust struct {
}

type RefreshTokenResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type SetUserInfoRequest struct {
	ID           string `path:"wxUid"`
	School       string `json:"school"`
	Grade        int64  `json:"grade"`
	Introduction string `json:"introduction"`
}

type SetUserInfoResponse struct {
	Avatar       string `json:"avatar"`
	School       string `json:"school"`
	Grade        int64  `json:"grade"`
	Introduction string `json:"introduction"`
	PostCount    int64  `json:"post_count"`
}

type UserInfo struct {
	WxUid        string `json:"wx_uid"`
	Avatar       string `json:"avatar"`
	School       string `json:"school"`
	Grade        int64  `json:"grade"`
	Introduction string `json:"introduction"`
	PostCount    int64  `json:"post_count"`
	Username     string `json:"username"`
}

type WxLoginRequest struct {
	Code     string `json:"code"`
	NickName string `json:"nickname"`
}

type WxLoginResponse struct {
	Token        string   `json:"token"`
	RefreshToken string   `json:"refresh_token"`
	UserInfo     UserInfo `json:"user_info"`
}
