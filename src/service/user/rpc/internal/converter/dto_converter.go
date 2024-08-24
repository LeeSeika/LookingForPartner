package converter

import (
	"lookingforpartner/model"
	"lookingforpartner/pb/user"
)

func UserDB2Rpc(u *model.User) *user.UserInfo {

	userInfo := user.UserInfo{
		PostCount:    u.PostCount,
		School:       u.School,
		Grade:        u.Grade,
		Avatar:       u.Avatar,
		Introduction: u.Introduction,
		Username:     u.Username,
	}
	return &userInfo
}
