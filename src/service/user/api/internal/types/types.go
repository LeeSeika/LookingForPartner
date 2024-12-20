// Code generated by goctl. DO NOT EDIT.
package types

type GetUserInfoRequest struct {
	WxUid string `path:"wxUid"`
}

type GetUserInfoResponse struct {
	Avatar       string `json:"avatar"`
	School       string `json:"school"`
	Grade        int64  `json:"grade"`
	Introduction string `json:"introduction"`
	PostCount    int64  `json:"post_count"`
}

type NewProject struct {
	Name          string `json:"name"`
	Introduction  string `json:"introduction"`
	MaintainerID  string `json:"maintainer_id"`
	Role          string `json:"role"`
	HeadCountInfo string `json:"head_count_info"`
	Progress      string `json:"progress"`
}

type PaginationParams struct {
	Page  int64  `form:"page, optional"`
	Size  int64  `form:"size, optional"`
	Order string `form:"order, optional"`
}

type Paginator struct {
	TotalRecord int64 `json:"total_record"`
	TotalPage   int   `json:"total_page"`
	Offset      int   `json:"offset"`
	Limit       int   `json:"limit"`
	CurrPage    int   `json:"curr_page"`
	PrevPage    int   `json:"prev_page"`
	NextPage    int   `json:"next_page"`
}

type Post struct {
	PostID    string   `json:"post_id"`
	CreatedAt string   `json:"created_at"`
	Title     string   `json:"title"`
	Project   Project  `json:"project"`
	Content   string   `json:"content"`
	Author    UserInfo `json:"author"`
}

type Project struct {
	ProjectID     string   `json:"project_id"`
	Name          string   `json:"name"`
	Introduction  string   `json:"introduction"`
	Maintainer    UserInfo `json:"maintainer"`
	Role          string   `json:"role"`
	HeadCountInfo string   `json:"head_count_info"`
	Progress      string   `json:"progress"`
}

type RefreshTokenRequest struct {
}

type RefreshTokenResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type SetUserInfoRequest struct {
	WxUid        string `path:"wxUid"`
	School       string `json:"school, optional"`
	Grade        int64  `json:"grade, optional"`
	Introduction string `json:"introduction, optional"`
}

type SetUserInfoResponse struct {
	Avatar       string `json:"avatar"`
	School       string `json:"school"`
	Grade        int64  `json:"grade"`
	Introduction string `json:"introduction"`
	PostCount    int64  `json:"post_count"`
}

type UpdatedProject struct {
	Name          string `json:"name, optional"`
	Introduction  string `json:"introduction, optional"`
	Role          string `json:"role, optional"`
	HeadCountInfo string `json:"head_count_info, optional"`
	Progress      string `json:"progress, optional"`
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
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
	Gender   int    `json:"gender"`
}

type WxLoginResponse struct {
	Token        string   `json:"token"`
	RefreshToken string   `json:"refresh_token"`
	UserInfo     UserInfo `json:"user_info"`
}
